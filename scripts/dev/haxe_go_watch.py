#!/usr/bin/env python3
"""Managed Haxe -> Go development loop.

The existing ``go-hx.sh`` command remains the authority for selecting and
running a build.  This module adds only the long-lived concerns: discovering
inputs, coalescing edits, retaining the last successful program, and cleaning
up processes that it starts.
"""

from __future__ import annotations

import argparse
from dataclasses import dataclass
import json
import os
from pathlib import Path
import shlex
import signal
import socket
import subprocess
import sys
import tempfile
import time
from typing import Callable, Mapping, Sequence


TOOL_DIR = Path(__file__).resolve().parent
DEFAULT_WRAPPER = TOOL_DIR / "go-hx.sh"
SOURCE_SUFFIXES = {".hx"}
CONFIG_NAMES = {"haxelib.json", ".haxerc"}


@dataclass(frozen=True)
class InputGraph:
    """The files and directories whose observable bytes can affect a build."""

    files: set[Path]
    roots: set[Path]
    extra_roots: set[Path]


@dataclass(frozen=True)
class BuildPlan:
    project: Path
    hxml: Path
    go_output: Path
    action: str


def _clean_hxml_line(line: str) -> list[str]:
    """Parse one HXML line without interpreting shell substitutions."""

    stripped = line.strip()
    if not stripped or stripped.startswith("#"):
        return []
    try:
        return shlex.split(stripped, comments=True, posix=True)
    except ValueError:
        return [stripped]


def discover_input_graph(
    hxml: Path,
    extra_roots: Sequence[Path],
    *,
    initial_cwd: Path | None = None,
) -> InputGraph:
    """Follow project-owned HXML includes, classpaths, and resource inputs.

    Haxelib and macro internals can depend on arbitrary external state, so this
    intentionally claims only the visible project graph plus paths explicitly
    supplied with ``--watch-dir``.
    """

    files: set[Path] = set()
    roots: set[Path] = set()
    extras: set[Path] = set()

    def add_path(raw: str, cwd: Path, target: set[Path]) -> Path:
        path = Path(raw).expanduser()
        if not path.is_absolute():
            path = cwd / path
        resolved = path.resolve()
        target.add(resolved)
        return resolved

    def scan(path: Path, cwd: Path, active: frozenset[Path]) -> Path:
        path = path.resolve()
        cwd = cwd.resolve()
        if path in active:
            return cwd
        files.add(path)
        try:
            lines = path.read_text(encoding="utf-8").splitlines()
        except OSError:
            return cwd

        nested_active = active | {path}
        for line in lines:
            tokens = _clean_hxml_line(line)
            if not tokens:
                continue
            option = tokens[0]
            value = tokens[1] if len(tokens) > 1 else ""
            if option in {"-C", "--cwd"} and value:
                next_cwd = Path(value).expanduser()
                cwd = (next_cwd if next_cwd.is_absolute() else cwd / next_cwd).resolve()
            elif option in {"-cp", "-p", "--class-path"} and value:
                add_path(value, cwd, roots)
            elif option in {"-resource", "--resource"} and value:
                add_path(value.split("@", 1)[0], cwd, files)
            elif len(tokens) == 1 and option.lower().endswith(".hxml"):
                included = add_path(option, cwd, files)
                cwd = scan(included, cwd, nested_active)
        return cwd

    hxml = hxml.resolve()
    project_cwd = (initial_cwd or hxml.parent).resolve()
    scan(hxml, project_cwd, frozenset())
    for name in CONFIG_NAMES:
        candidate = project_cwd / name
        if candidate.exists():
            files.add(candidate.resolve())
    for raw in extra_roots:
        path = raw.expanduser()
        if not path.is_absolute():
            path = project_cwd / path
        path = path.resolve()
        if path.is_dir():
            extras.add(path)
        else:
            files.add(path)
    return InputGraph(files=files, roots=roots, extra_roots=extras)


Snapshot = dict[Path, tuple[int, int] | None]


def _record_file(snapshot: Snapshot, path: Path) -> None:
    try:
        stat = path.stat()
        snapshot[path.resolve()] = (stat.st_mtime_ns, stat.st_size)
    except OSError:
        snapshot[path.resolve()] = None


def capture_snapshot(graph: InputGraph) -> Snapshot:
    """Capture cheap file metadata for polling without reading source bytes."""

    snapshot: Snapshot = {}
    for path in graph.files:
        _record_file(snapshot, path)
    for root in graph.roots:
        if not root.exists():
            _record_file(snapshot, root)
            continue
        for path in root.rglob("*"):
            if path.is_file() and path.suffix.lower() in SOURCE_SUFFIXES:
                _record_file(snapshot, path)
    for root in graph.extra_roots:
        if not root.exists():
            _record_file(snapshot, root)
            continue
        for path in root.rglob("*"):
            if path.is_file():
                _record_file(snapshot, path)
    return snapshot


class ChangeDebouncer:
    """Turn a burst of changing snapshots into one stable rebuild request."""

    def __init__(self, quiet_seconds: float) -> None:
        self.quiet_seconds = quiet_seconds
        self._candidate: object | None = None
        self._changed_at: float | None = None
        self._ready: object | None = None

    def observe(self, now: float, snapshot: object) -> bool:
        if snapshot != self._candidate:
            self._candidate = snapshot
            self._changed_at = now
            return False
        if self._changed_at is None or now - self._changed_at < self.quiet_seconds:
            return False
        self._ready = self._candidate
        self._candidate = None
        self._changed_at = None
        return True

    def consume(self):
        ready = self._ready
        self._ready = None
        return ready


def start_owned_process(command: Sequence[str], cwd: Path) -> subprocess.Popen:
    """Start a child in a separate process group owned by this watcher."""

    kwargs: dict[str, object] = {"cwd": str(cwd)}
    if os.name == "posix":
        kwargs["start_new_session"] = True
    elif os.name == "nt":
        kwargs["creationflags"] = subprocess.CREATE_NEW_PROCESS_GROUP
    return subprocess.Popen(list(command), **kwargs)


def stop_owned_process(process: subprocess.Popen, timeout: float = 3.0) -> None:
    """Stop a child and its descendants, escalating only after a grace period."""

    if os.name != "posix":
        if process.poll() is not None:
            return
        process.terminate()
        try:
            process.wait(timeout=timeout)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait(timeout=timeout)
        return

    group = process.pid
    deadline = time.monotonic() + timeout

    def signal_group(signum: int) -> bool:
        try:
            os.killpg(group, signum)
            return True
        except (PermissionError, ProcessLookupError):
            return False

    def group_alive() -> bool:
        try:
            os.killpg(group, 0)
            return True
        except (PermissionError, ProcessLookupError):
            return False

    signal_group(signal.SIGTERM)
    try:
        if process.poll() is None:
            process.wait(timeout=max(0.0, deadline - time.monotonic()))
    except subprocess.TimeoutExpired:
        signal_group(signal.SIGKILL)
        process.wait(timeout=max(0.1, timeout))

    # The leader can exit on SIGTERM while a descendant ignores it. Wait only
    # within the same bounded grace period, then kill the remaining owned group.
    while group_alive() and time.monotonic() < deadline:
        time.sleep(0.01)
    if group_alive():
        signal_group(signal.SIGKILL)


class RunLifecycle:
    """Publish a new runnable process only after its build succeeds."""

    def __init__(
        self,
        *,
        build: Callable[[], bool],
        start: Callable[[], subprocess.Popen],
        stop: Callable[[subprocess.Popen], None],
    ) -> None:
        self.build = build
        self.start = start
        self.stop = stop
        self.process: subprocess.Popen | None = None

    def rebuild(self) -> bool:
        if not self.build():
            return False
        previous = self.process
        if previous is not None:
            self.stop(previous)
        self.process = self.start()
        return True

    def close(self) -> None:
        if self.process is not None:
            self.stop(self.process)
            self.process = None


class CompilerServer:
    """Own one opt-in Haxe ``--wait`` process for this watch session."""

    def __init__(self, haxe_bin: str, project: Path) -> None:
        self.haxe_bin = haxe_bin
        self.project = project
        self.process: subprocess.Popen | None = None
        self.port: int | None = None

    def start(self, cancelled: Callable[[], bool] = lambda: False) -> int:
        probe = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        probe.bind(("127.0.0.1", 0))
        self.port = int(probe.getsockname()[1])
        probe.close()
        self.process = start_owned_process(
            [self.haxe_bin, "--wait", str(self.port)], cwd=self.project
        )
        deadline = time.monotonic() + 5.0
        while time.monotonic() < deadline:
            if cancelled():
                self.close()
                raise InterruptedError("compiler server startup was cancelled")
            if self.process.poll() is not None:
                raise RuntimeError("Haxe compiler server exited before becoming ready")
            try:
                with socket.create_connection(("127.0.0.1", self.port), timeout=0.1):
                    return self.port
            except OSError:
                time.sleep(0.03)
        self.close()
        raise RuntimeError("Haxe compiler server did not become ready within 5 seconds")

    def close(self) -> None:
        if self.process is not None:
            stop_owned_process(self.process)
            self.process = None


def _run(command: Sequence[str], cwd: Path, *, capture: bool = False) -> subprocess.CompletedProcess:
    return subprocess.run(
        list(command),
        cwd=cwd,
        text=True,
        capture_output=capture,
        check=False,
    )


def _run_interruptible(
    command: Sequence[str], cwd: Path, cancelled: Callable[[], bool]
) -> subprocess.CompletedProcess:
    """Run one build as an owned group so watcher shutdown can interrupt it."""

    process = start_owned_process(command, cwd)
    while process.poll() is None:
        if cancelled():
            stop_owned_process(process)
            return subprocess.CompletedProcess(list(command), 130)
        time.sleep(0.03)
    return subprocess.CompletedProcess(list(command), process.returncode)


def inspect_plan(wrapper: Path, wrapper_args: Sequence[str]) -> BuildPlan:
    result = _run(
        ["bash", str(wrapper), *wrapper_args, "--print-plan-json"],
        Path.cwd(),
        capture=True,
    )
    if result.returncode != 0:
        if result.stderr:
            sys.stderr.write(result.stderr)
        raise RuntimeError("could not resolve the Haxe -> Go development build")
    payload = json.loads(result.stdout)
    return BuildPlan(
        project=Path(payload["project"]),
        hxml=Path(payload["hxml"]),
        go_output=Path(payload["goOutput"]),
        action=str(payload["action"]),
    )


def _wrapper_args(args: argparse.Namespace, *, action: str | None = None) -> list[str]:
    command = ["--project", args.project, "--action", action or args.action]
    if args.profile:
        command.extend(["--profile", args.profile])
    if args.hxml:
        command.extend(["--hxml", args.hxml])
    if args.ci:
        command.append("--ci")
    if args.out:
        command.extend(["--out", args.out])
    if args.binary and (action or args.action) == "build":
        command.extend(["--binary", args.binary])
    if args.haxe_bin:
        command.extend(["--haxe-bin", args.haxe_bin])
    if args.go_bin:
        command.extend(["--go-bin", args.go_bin])
    for define in args.define:
        command.extend(["--define", define])
    return command


def _parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Compile Haxe to Go, watch the real project inputs, and safely restart after success."
    )
    parser.add_argument("--project", default=".")
    parser.add_argument("--profile", choices=("portable", "metal"))
    parser.add_argument("--hxml")
    parser.add_argument("--ci", action="store_true")
    parser.add_argument(
        "--action", choices=("compile", "run", "build", "test", "vet", "fmt"), default="run"
    )
    parser.add_argument("--out")
    parser.add_argument("--binary")
    parser.add_argument("--haxe-bin", default=os.environ.get("HAXE_BIN", "haxe"))
    parser.add_argument("--go-bin", default=os.environ.get("GO_BIN", "go"))
    parser.add_argument("--define", action="append", default=[])
    parser.add_argument("--watch-dir", action="append", default=[])
    parser.add_argument("--debounce-ms", type=int, default=120)
    parser.add_argument("--poll-ms", type=int, default=80)
    parser.add_argument("--server", choices=("off", "on"), default="on")
    parser.add_argument("--once", action="store_true")
    parser.add_argument("--max-cycles", type=int, help=argparse.SUPPRESS)
    parser.add_argument("--wrapper", type=Path, default=DEFAULT_WRAPPER, help=argparse.SUPPRESS)
    return parser


def _validate_durations(args: argparse.Namespace, parser: argparse.ArgumentParser) -> None:
    if args.debounce_ms < 0:
        parser.error("--debounce-ms must be non-negative")
    if args.poll_ms <= 0:
        parser.error("--poll-ms must be positive")


def main(argv: Sequence[str] | None = None) -> int:
    raw = list(sys.argv[1:] if argv is None else argv)
    go_args: list[str] = []
    if "--" in raw:
        divider = raw.index("--")
        go_args = raw[divider + 1 :]
        raw = raw[:divider]
    parser = _parser()
    args = parser.parse_args(raw)
    _validate_durations(args, parser)
    wrapper = args.wrapper.resolve()
    selected_args = _wrapper_args(args)
    plan = inspect_plan(wrapper, selected_args)

    if args.once:
        command = ["bash", str(wrapper), *selected_args]
        if go_args:
            command.extend(["--", *go_args])
        return _run(command, Path.cwd()).returncode

    stop_requested = False

    def request_stop(_signum, _frame) -> None:
        nonlocal stop_requested
        stop_requested = True

    previous_handlers: dict[int, object] = {}
    for signum in (signal.SIGINT, signal.SIGTERM):
        previous_handlers[signum] = signal.signal(signum, request_stop)

    server: CompilerServer | None = None
    connect_args: list[str] = []
    temp: tempfile.TemporaryDirectory[str] | None = None
    executable: Path | None = None
    lifecycle: RunLifecycle | None = None

    def build() -> bool:
        nonlocal connect_args, server
        if args.action == "run":
            if executable is None:
                raise RuntimeError("watch executable storage was not initialized")
            build_args = _wrapper_args(args, action="build")
            if args.binary:
                binary_index = build_args.index("--binary")
                del build_args[binary_index : binary_index + 2]
            build_args.extend(["--binary", str(executable), *connect_args])
            command = ["bash", str(wrapper), *build_args]
        else:
            command = ["bash", str(wrapper), *selected_args, *connect_args]
            if go_args:
                command.extend(["--", *go_args])
        started = time.monotonic()
        result = _run_interruptible(command, Path.cwd(), lambda: stop_requested)
        if (
            result.returncode != 0
            and connect_args
            and server is not None
            and server.process is not None
            and server.process.poll() is not None
        ):
            print("[hx-go-watch] compiler server stopped; retrying this build directly", file=sys.stderr)
            connect_index = command.index("--connect")
            del command[connect_index : connect_index + 2]
            server.close()
            server = None
            connect_args = []
            result = _run_interruptible(command, Path.cwd(), lambda: stop_requested)
        elapsed = time.monotonic() - started
        if result.returncode == 0:
            print(f"[hx-go-watch] build succeeded in {elapsed:.2f}s")
            return True
        print(
            f"[hx-go-watch] build failed in {elapsed:.2f}s; watching continues",
            file=sys.stderr,
        )
        return False

    def start_program() -> subprocess.Popen:
        if executable is None:
            raise RuntimeError("watch executable storage was not initialized")
        print("[hx-go-watch] starting the new successful program")
        return start_owned_process([str(executable), *go_args], cwd=plan.project)

    try:
        if args.server == "on":
            server = CompilerServer(args.haxe_bin, plan.project)
            try:
                port = server.start(lambda: stop_requested)
                connect_args = ["--connect", str(port)]
                print(f"[hx-go-watch] compiler server ready on localhost:{port}")
            except InterruptedError:
                return 0
            except (OSError, RuntimeError) as error:
                print(
                    f"[hx-go-watch] compiler server unavailable; using direct builds: {error}",
                    file=sys.stderr,
                )
                server.close()
                server = None

        if stop_requested:
            return 0

        temp = tempfile.TemporaryDirectory(prefix="haxe-go-watch-")
        executable = Path(temp.name) / ("app.exe" if os.name == "nt" else "app")
        if args.action == "run":
            lifecycle = RunLifecycle(build=build, start=start_program, stop=stop_owned_process)

        print(f"[hx-go-watch] project={plan.project}")
        print(f"[hx-go-watch] hxml={plan.hxml}")
        print(
            f"[hx-go-watch] mode={'server-backed' if connect_args else 'direct'} "
            f"action={args.action}"
        )
        graph = discover_input_graph(
            plan.hxml,
            [Path(path) for path in args.watch_dir],
            initial_cwd=plan.project,
        )
        before_initial = capture_snapshot(graph)
        succeeded = lifecycle.rebuild() if lifecycle is not None else build()
        if stop_requested:
            return 0
        if not succeeded:
            print("[hx-go-watch] fix the error and save; the watcher is still active")

        graph = discover_input_graph(
            plan.hxml,
            [Path(path) for path in args.watch_dir],
            initial_cwd=plan.project,
        )
        after_initial = capture_snapshot(graph)
        baseline: Mapping[Path, tuple[int, int] | None] = (
            before_initial if after_initial != before_initial else after_initial
        )
        print(
            f"[hx-go-watch] watching {len(graph.roots) + len(graph.extra_roots)} source root(s) "
            f"and {len(graph.files)} direct input(s); press Ctrl-C to stop"
        )
        debouncer = ChangeDebouncer(args.debounce_ms / 1000.0)
        completed_cycles = 0
        while not stop_requested:
            time.sleep(args.poll_ms / 1000.0)
            current = capture_snapshot(graph)
            if current == baseline:
                continue
            if not debouncer.observe(time.monotonic(), current):
                continue
            requested = debouncer.consume()
            if requested is None:
                continue
            changed_paths = sum(
                1 for path in set(baseline) | set(current) if baseline.get(path) != current.get(path)
            )
            baseline = requested
            print(f"[hx-go-watch] change detected; rebuilding ({changed_paths} changed input(s))")
            succeeded = lifecycle.rebuild() if lifecycle is not None else build()
            if not succeeded and lifecycle is not None and lifecycle.process is not None:
                print("[hx-go-watch] last working program keeps running")
            graph = discover_input_graph(
                plan.hxml,
                [Path(path) for path in args.watch_dir],
                initial_cwd=plan.project,
            )
            after_build = capture_snapshot(graph)
            # Keep the pre-build snapshot as the baseline when an edit arrived
            # during compilation. The next polling pass will schedule another
            # build instead of silently treating that edit as already compiled.
            baseline = requested if after_build != requested else after_build
            completed_cycles += 1
            if args.max_cycles is not None and completed_cycles >= args.max_cycles:
                break
        return 0
    finally:
        if lifecycle is not None:
            lifecycle.close()
        if server is not None:
            server.close()
        if temp is not None:
            temp.cleanup()
        for signum, handler in previous_handlers.items():
            signal.signal(signum, handler)
        print("[hx-go-watch] stopped; owned program and compiler server are closed")


if __name__ == "__main__":
    raise SystemExit(main())
