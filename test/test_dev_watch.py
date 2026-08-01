#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import json
import os
from pathlib import Path
import signal
import shlex
import shutil
import subprocess
import sys
import tempfile
import time
import unittest


ROOT = Path(__file__).resolve().parent.parent
WATCHER = ROOT / "scripts" / "dev" / "haxe_go_watch.py"
WRAPPER = ROOT / "scripts" / "dev" / "go-hx.sh"
PACKAGE = ROOT / "package.json"
TEMPLATE_PACKAGE = ROOT / "templates" / "basic" / "package.json"
TEMPLATE_WATCHER = ROOT / "templates" / "basic" / "scripts" / "dev" / "watch.py"
DOC = ROOT / "docs" / "development-watch-loop.md"


def load_watcher():
    spec = importlib.util.spec_from_file_location("haxe_go_watch", WATCHER)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load development watcher")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


class DevWatchContractTest(unittest.TestCase):
    def wait_for_text(self, path: Path, expected: str, timeout: float = 8.0) -> str:
        deadline = time.monotonic() + timeout
        text = ""
        while time.monotonic() < deadline:
            if path.exists():
                text = path.read_text(encoding="utf-8")
                if expected in text:
                    return text
            time.sleep(0.02)
        self.fail(f"timed out waiting for {expected!r} in {path}; observed:\n{text}")

    def assert_pid_gone(self, pid: int, timeout: float = 2.0) -> None:
        deadline = time.monotonic() + timeout
        while time.monotonic() < deadline:
            try:
                os.kill(pid, 0)
            except ProcessLookupError:
                return
            time.sleep(0.01)
        self.fail(f"process {pid} remained visible after cleanup")

    def stop_watcher(self, watcher: subprocess.Popen) -> None:
        """Let the watcher clean its owned groups before using a last-resort kill."""

        if watcher.poll() is not None:
            return
        watcher.send_signal(signal.SIGTERM)
        try:
            watcher.wait(timeout=5)
        except subprocess.TimeoutExpired:
            watcher.kill()
            watcher.wait(timeout=2)

    def test_input_graph_follows_nested_hxml_classpaths_and_resources(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-inputs-") as raw:
            project = Path(raw)
            (project / "src").mkdir()
            (project / "src" / "Main.hx").write_text("class Main {}\n", encoding="utf-8")
            (project / "shared").mkdir()
            (project / "shared" / "Shared.hx").write_text("class Shared {}\n", encoding="utf-8")
            (project / "config").mkdir()
            resource = project / "config" / "app.json"
            resource.write_text("{}\n", encoding="utf-8")
            (project / "nested").mkdir()
            nested = project / "nested" / "options.hxml"
            nested.write_text("-cp shared\n", encoding="utf-8")
            build = project / "compile.hxml"
            build.write_text(
                "-cp src\nnested/options.hxml\n-resource config/app.json@app\n",
                encoding="utf-8",
            )

            graph = module.discover_input_graph(build, [])

            self.assertEqual({build.resolve(), nested.resolve(), resource.resolve()}, graph.files)
            self.assertEqual({(project / "src").resolve(), (project / "shared").resolve()}, graph.roots)
            snapshot = module.capture_snapshot(graph)
            self.assertIn((project / "src" / "Main.hx").resolve(), snapshot)
            self.assertIn((project / "shared" / "Shared.hx").resolve(), snapshot)

    def test_input_graph_uses_project_cwd_and_propagates_nested_cwd_changes(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-hxml-cwd-") as raw:
            project = Path(raw)
            (project / "config").mkdir()
            (project / "src").mkdir()
            (project / ".haxerc").write_text("{}\n", encoding="utf-8")
            schemas = project / "schemas"
            schemas.mkdir()
            schema = schemas / "api.proto"
            schema.write_text('syntax = "proto3";\n', encoding="utf-8")
            shifted = project / "shifted"
            (shifted / "shared").mkdir(parents=True)
            (shifted / "data").mkdir()
            resource = shifted / "data" / "app.json"
            resource.write_text("{}\n", encoding="utf-8")
            nested = project / "config" / "nested.hxml"
            nested.write_text("-C shifted\n-cp shared\n", encoding="utf-8")
            build = project / "config" / "compile.hxml"
            build.write_text(
                "-cp src\nconfig/nested.hxml\n-resource data/app.json@app\n",
                encoding="utf-8",
            )

            graph = module.discover_input_graph(
                build, [Path("schemas")], initial_cwd=project
            )

            self.assertIn((project / "src").resolve(), graph.roots)
            self.assertIn((shifted / "shared").resolve(), graph.roots)
            self.assertIn(resource.resolve(), graph.files)
            self.assertIn((project / ".haxerc").resolve(), graph.files)
            self.assertIn(schemas.resolve(), graph.extra_roots)
            self.assertIn(schema.resolve(), module.capture_snapshot(graph))

    def test_explicit_watch_directory_includes_arbitrary_macro_inputs(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-extra-") as raw:
            project = Path(raw)
            build = project / "compile.hxml"
            build.write_text("-cp src\n", encoding="utf-8")
            (project / "src").mkdir()
            schemas = project / "schemas"
            schemas.mkdir()
            proto = schemas / "api.proto"
            proto.write_text("syntax = \"proto3\";\n", encoding="utf-8")

            graph = module.discover_input_graph(build, [schemas])
            snapshot = module.capture_snapshot(graph)

            self.assertIn(proto.resolve(), snapshot)

    def test_debouncer_coalesces_a_burst_and_reports_the_latest_snapshot(self) -> None:
        module = load_watcher()
        debouncer = module.ChangeDebouncer(0.1)
        self.assertFalse(debouncer.observe(0.00, {"a": 1}))
        self.assertFalse(debouncer.observe(0.05, {"a": 2}))
        self.assertFalse(debouncer.observe(0.14, {"a": 2}))
        self.assertTrue(debouncer.observe(0.16, {"a": 2}))
        self.assertEqual({"a": 2}, debouncer.consume())
        self.assertIsNone(debouncer.consume())

    def test_long_lived_watch_defaults_to_managed_server_but_once_stays_direct(self) -> None:
        module = load_watcher()
        parsed = module._parser().parse_args([])
        self.assertEqual("on", parsed.server)
        self.assertFalse(parsed.once)

    def test_failed_rebuild_keeps_last_good_process_and_success_restarts_it(self) -> None:
        module = load_watcher()

        class FakeProcess:
            def __init__(self, name: str) -> None:
                self.name = name
                self.stopped = False

            def poll(self):
                return 0 if self.stopped else None

        started: list[FakeProcess] = []
        stopped: list[FakeProcess] = []
        outcomes = iter([True, False, True])

        def build() -> bool:
            return next(outcomes)

        def start() -> FakeProcess:
            process = FakeProcess(f"run-{len(started) + 1}")
            started.append(process)
            return process

        def stop(process: FakeProcess) -> None:
            process.stopped = True
            stopped.append(process)

        lifecycle = module.RunLifecycle(build=build, start=start, stop=stop)
        self.assertTrue(lifecycle.rebuild())
        first = lifecycle.process
        self.assertIs(first, started[0])
        self.assertFalse(lifecycle.rebuild())
        self.assertIs(first, lifecycle.process)
        self.assertEqual([], stopped)
        self.assertTrue(lifecycle.rebuild())
        self.assertIs(started[1], lifecycle.process)
        self.assertEqual([first], stopped)
        lifecycle.close()
        self.assertEqual([first, started[1]], stopped)

    @unittest.skipUnless(os.name == "posix", "process-group lifecycle is POSIX-only")
    def test_process_owner_stops_the_complete_child_process_group(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-process-") as raw:
            pid_file = Path(raw) / "child.pid"
            parent = module.start_owned_process(
                [
                    sys.executable,
                    "-c",
                    (
                        "import pathlib,subprocess,sys,time; "
                        "p=subprocess.Popen([sys.executable,'-c','import time; time.sleep(60)']); "
                        "pathlib.Path(sys.argv[1]).write_text(str(p.pid)); time.sleep(60)"
                    ),
                    str(pid_file),
                ],
                cwd=Path(raw),
            )
            try:
                deadline = time.monotonic() + 5
                while not pid_file.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(pid_file.exists())
                child_pid = int(pid_file.read_text(encoding="utf-8"))

                module.stop_owned_process(parent, timeout=2)

                self.assertIsNotNone(parent.poll())
                self.assert_pid_gone(child_pid)
            finally:
                try:
                    os.killpg(parent.pid, signal.SIGKILL)
                except (PermissionError, ProcessLookupError):
                    pass

    @unittest.skipUnless(os.name == "posix", "process-group lifecycle is POSIX-only")
    def test_process_owner_accepts_an_already_exited_group(self) -> None:
        module = load_watcher()
        process = module.start_owned_process(
            [sys.executable, "-c", "pass"], cwd=ROOT
        )
        process.wait(timeout=2)
        module.stop_owned_process(process, timeout=1)
        self.assertEqual(0, process.returncode)

    @unittest.skipUnless(os.name == "posix", "process-group lifecycle is POSIX-only")
    def test_process_owner_kills_descendant_that_ignores_term_after_leader_exits(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-stubborn-") as raw:
            pid_file = Path(raw) / "child.pid"
            leader = module.start_owned_process(
                [
                    sys.executable,
                    "-c",
                    (
                        "import pathlib,signal,subprocess,sys,time; "
                        "p=subprocess.Popen([sys.executable,'-c',"
                        "'import signal,time; signal.signal(signal.SIGTERM,signal.SIG_IGN); time.sleep(60)']); "
                        "pathlib.Path(sys.argv[1]).write_text(str(p.pid)); time.sleep(60)"
                    ),
                    str(pid_file),
                ],
                cwd=Path(raw),
            )
            child_pid: int | None = None
            try:
                deadline = time.monotonic() + 5
                while not pid_file.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(pid_file.exists())
                child_pid = int(pid_file.read_text(encoding="utf-8"))
                time.sleep(0.05)
                module.stop_owned_process(leader, timeout=0.2)
                self.assert_pid_gone(child_pid)
            finally:
                try:
                    os.killpg(leader.pid, signal.SIGKILL)
                except (PermissionError, ProcessLookupError):
                    pass

    @unittest.skipUnless(os.name == "posix", "managed compiler server is a POSIX dev contract")
    def test_compiler_server_is_owned_and_closed(self) -> None:
        module = load_watcher()
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-server-") as raw:
            root = Path(raw)
            fake_haxe = root / "fake-haxe"
            fake_haxe.write_text(
                "#!/usr/bin/env python3\n"
                "import socket,sys\n"
                "port=int(sys.argv[2])\n"
                "server=socket.socket(); server.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)\n"
                "server.bind(('127.0.0.1',port)); server.listen()\n"
                "while True:\n"
                "    connection,_=server.accept(); connection.close()\n",
                encoding="utf-8",
            )
            fake_haxe.chmod(0o755)
            server = module.CompilerServer(str(fake_haxe), root)
            process = None
            try:
                port = server.start()
                process = server.process
                self.assertGreater(port, 0)
                self.assertIsNotNone(process)
                self.assertIsNone(process.poll())
            finally:
                server.close()
            self.assertIsNotNone(process)
            self.assertIsNotNone(process.poll())

    @unittest.skipUnless(os.name == "posix", "startup signal lifecycle is POSIX-only")
    def test_signal_during_server_startup_cleans_the_owned_server(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-startup-signal-") as raw:
            root = Path(raw)
            hxml = root / "compile.hxml"
            hxml.write_text("-cp .\n-D go_output=out\n", encoding="utf-8")
            plan = json.dumps(
                {"project": str(root), "hxml": str(hxml), "goOutput": str(root / "out"), "action": "compile"}
            )
            wrapper = root / "wrapper.sh"
            wrapper.write_text(
                "#!/usr/bin/env bash\n"
                f"printf '%s\\n' {shlex.quote(plan)}\n",
                encoding="utf-8",
            )
            wrapper.chmod(0o755)
            server_pid = root / "server.pid"
            fake_haxe = root / "fake-haxe"
            fake_haxe.write_text(
                "#!/usr/bin/env python3\n"
                "import pathlib,signal,sys,time\n"
                f"pathlib.Path({str(server_pid)!r}).write_text(str(__import__('os').getpid()))\n"
                "signal.signal(signal.SIGTERM, signal.SIG_IGN)\n"
                "time.sleep(60)\n",
                encoding="utf-8",
            )
            fake_haxe.chmod(0o755)
            watcher = subprocess.Popen(
                [
                    sys.executable,
                    str(WATCHER),
                    "--project",
                    str(root),
                    "--wrapper",
                    str(wrapper),
                    "--haxe-bin",
                    str(fake_haxe),
                    "--action",
                    "compile",
                ],
                cwd=ROOT,
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            owned_pid: int | None = None
            try:
                deadline = time.monotonic() + 5
                while not server_pid.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(server_pid.exists())
                owned_pid = int(server_pid.read_text(encoding="utf-8"))
                watcher.send_signal(signal.SIGTERM)
                watcher.wait(timeout=5)
                self.assertEqual(0, watcher.returncode)
                self.assert_pid_gone(owned_pid)
            finally:
                self.stop_watcher(watcher)
                if owned_pid is not None:
                    try:
                        os.kill(owned_pid, signal.SIGKILL)
                    except ProcessLookupError:
                        pass

    @unittest.skipUnless(os.name == "posix", "active-build signal lifecycle is POSIX-only")
    def test_signal_during_active_build_stops_the_owned_build_process(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-active-signal-") as raw:
            root = Path(raw)
            hxml = root / "compile.hxml"
            hxml.write_text("-cp .\n-D go_output=out\n", encoding="utf-8")
            plan = json.dumps(
                {"project": str(root), "hxml": str(hxml), "goOutput": str(root / "out"), "action": "compile"}
            )
            build_pid = root / "build.pid"
            wrapper = root / "wrapper.sh"
            wrapper.write_text(
                "#!/usr/bin/env bash\n"
                "set -euo pipefail\n"
                "for arg in \"$@\"; do\n"
                f"  if [[ \"$arg\" == --print-plan-json ]]; then printf '%s\\n' {shlex.quote(plan)}; exit 0; fi\n"
                "done\n"
                f"printf '%s' \"$$\" > {shlex.quote(str(build_pid))}\n"
                "exec sleep 60\n",
                encoding="utf-8",
            )
            wrapper.chmod(0o755)
            watcher = subprocess.Popen(
                [
                    sys.executable,
                    str(WATCHER),
                    "--project",
                    str(root),
                    "--wrapper",
                    str(wrapper),
                    "--server",
                    "off",
                    "--action",
                    "compile",
                ],
                cwd=ROOT,
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            owned_pid: int | None = None
            try:
                deadline = time.monotonic() + 5
                while not build_pid.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(build_pid.exists())
                owned_pid = int(build_pid.read_text(encoding="utf-8"))
                watcher.send_signal(signal.SIGTERM)
                watcher.wait(timeout=3)
                self.assertEqual(0, watcher.returncode)
                self.assert_pid_gone(owned_pid)
            finally:
                self.stop_watcher(watcher)
                if owned_pid is not None:
                    try:
                        os.killpg(owned_pid, signal.SIGKILL)
                    except (PermissionError, ProcessLookupError):
                        pass

    @unittest.skipUnless(os.name == "posix", "server fallback lifecycle is POSIX-only")
    def test_dead_server_falls_back_once_and_later_builds_stay_direct(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-server-fallback-") as raw:
            root = Path(raw)
            source = root / "Main.hx"
            source.write_text("ONE\n", encoding="utf-8")
            hxml = root / "compile.hxml"
            hxml.write_text("-cp .\n-D go_output=out\n", encoding="utf-8")
            calls = root / "calls.log"
            plan = json.dumps(
                {"project": str(root), "hxml": str(hxml), "goOutput": str(root / "out"), "action": "compile"}
            )
            wrapper = root / "wrapper.sh"
            wrapper.write_text(
                "#!/usr/bin/env bash\n"
                "set -euo pipefail\n"
                "for arg in \"$@\"; do\n"
                f"  if [[ \"$arg\" == --print-plan-json ]]; then printf '%s\\n' {shlex.quote(plan)}; exit 0; fi\n"
                "done\n"
                f"if [[ \" $* \" == *' --connect '* ]]; then echo CONNECT >> {shlex.quote(str(calls))}; exit 1; fi\n"
                f"echo DIRECT >> {shlex.quote(str(calls))}\n",
                encoding="utf-8",
            )
            wrapper.chmod(0o755)
            fake_haxe = root / "fake-haxe"
            fake_haxe.write_text(
                "#!/usr/bin/env python3\n"
                "import socket,sys\n"
                "server=socket.socket(); server.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)\n"
                "server.bind(('127.0.0.1',int(sys.argv[2]))); server.listen(); connection,_=server.accept(); connection.close()\n",
                encoding="utf-8",
            )
            fake_haxe.chmod(0o755)
            log = root / "watch.log"
            environment = dict(os.environ, PYTHONUNBUFFERED="1")
            with log.open("w", encoding="utf-8") as output:
                watcher = subprocess.Popen(
                    [
                        sys.executable,
                        str(WATCHER),
                        "--project",
                        str(root),
                        "--wrapper",
                        str(wrapper),
                        "--haxe-bin",
                        str(fake_haxe),
                        "--action",
                        "compile",
                        "--poll-ms",
                        "20",
                        "--debounce-ms",
                        "40",
                        "--max-cycles",
                        "1",
                    ],
                    cwd=ROOT,
                    env=environment,
                    stdout=output,
                    stderr=subprocess.STDOUT,
                )
            try:
                self.wait_for_text(log, "watching")
                source.write_text("TWO\n", encoding="utf-8")
                watcher.wait(timeout=8)
                self.assertEqual(0, watcher.returncode, log.read_text(encoding="utf-8"))
                self.assertEqual(["CONNECT", "DIRECT", "DIRECT"], calls.read_text().splitlines())
            finally:
                self.stop_watcher(watcher)

    @unittest.skipUnless(os.name == "posix", "restart lifecycle is a POSIX dev contract")
    def test_real_loop_recovers_from_failure_without_stopping_last_good_program(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-loop-") as raw:
            root = Path(raw)
            source = root / "Main.hx"
            source.write_text("GOOD-1\n", encoding="utf-8")
            hxml = root / "compile.hxml"
            hxml.write_text("-cp .\n-D go_output=out\n", encoding="utf-8")
            wrapper = root / "fake-wrapper.sh"
            plan = json.dumps(
                {
                    "project": str(root),
                    "hxml": str(hxml),
                    "goOutput": str(root / "out"),
                    "action": "run",
                }
            )
            wrapper.write_text(
                "#!/usr/bin/env bash\n"
                "set -euo pipefail\n"
                "for arg in \"$@\"; do\n"
                f"  if [[ \"$arg\" == --print-plan-json ]]; then printf '%s\\n' {shlex.quote(plan)}; exit 0; fi\n"
                "done\n"
                "binary=''\n"
                "while [[ $# -gt 0 ]]; do\n"
                "  if [[ \"$1\" == --binary ]]; then binary=\"$2\"; shift 2; else shift; fi\n"
                "done\n"
                f"if grep -q FAIL {shlex.quote(str(source))}; then exit 1; fi\n"
                "mkdir -p \"$(dirname \"$binary\")\"\n"
                f"value=$(head -n 1 {shlex.quote(str(source))})\n"
                "printf '#!/usr/bin/env bash\\nprintf \"RUN:%%s:%%s\\n\" \"$$\" \"%s\"\\nwhile :; do sleep 1; done\\n' \"$value\" > \"$binary\"\n"
                "chmod +x \"$binary\"\n",
                encoding="utf-8",
            )
            wrapper.chmod(0o755)
            log = root / "watch.log"
            environment = dict(os.environ)
            environment["PYTHONUNBUFFERED"] = "1"
            with log.open("w", encoding="utf-8") as output:
                watcher = subprocess.Popen(
                    [
                        sys.executable,
                        str(WATCHER),
                        "--project",
                        str(root),
                        "--wrapper",
                        str(wrapper),
                        "--server",
                        "off",
                        "--poll-ms",
                        "20",
                        "--debounce-ms",
                        "40",
                    ],
                    cwd=ROOT,
                    env=environment,
                    stdout=output,
                    stderr=subprocess.STDOUT,
                )
            owned_pids: list[int] = []
            try:
                first_log = self.wait_for_text(log, "RUN:")
                first_pid = int(first_log.split("RUN:", 1)[1].split(":", 1)[0])
                owned_pids.append(first_pid)
                source.write_text("FAIL\n", encoding="utf-8")
                self.wait_for_text(log, "last working program keeps running")
                os.kill(first_pid, 0)
                source.write_text("GOOD-2\n", encoding="utf-8")
                second_log = self.wait_for_text(log, "GOOD-2")
                second_pid = int(second_log.rsplit("RUN:", 1)[1].split(":", 1)[0])
                owned_pids.append(second_pid)
                watcher.send_signal(signal.SIGTERM)
                watcher.wait(timeout=10)
                self.assertEqual(0, watcher.returncode, log.read_text(encoding="utf-8"))
                for pid in owned_pids:
                    self.assert_pid_gone(pid)
            finally:
                self.stop_watcher(watcher)
                for pid in owned_pids:
                    try:
                        os.killpg(pid, signal.SIGKILL)
                    except (PermissionError, ProcessLookupError):
                        pass

    @unittest.skipUnless(os.name == "posix", "initial-build edit lifecycle is POSIX-only")
    def test_edit_during_initial_build_is_rebuilt_before_becoming_quiet(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-initial-edit-") as raw:
            root = Path(raw)
            source = root / "Main.hx"
            source.write_text("OLD\n", encoding="utf-8")
            hxml = root / "compile.hxml"
            hxml.write_text("-cp .\n-D go_output=out\n", encoding="utf-8")
            calls = root / "calls.log"
            entered = root / "entered"
            release = root / "release"
            plan = json.dumps(
                {"project": str(root), "hxml": str(hxml), "goOutput": str(root / "out"), "action": "compile"}
            )
            wrapper = root / "wrapper.sh"
            wrapper.write_text(
                "#!/usr/bin/env bash\n"
                "set -euo pipefail\n"
                "for arg in \"$@\"; do\n"
                f"  if [[ \"$arg\" == --print-plan-json ]]; then printf '%s\\n' {shlex.quote(plan)}; exit 0; fi\n"
                "done\n"
                f"head -n 1 {shlex.quote(str(source))} >> {shlex.quote(str(calls))}\n"
                f"if [[ $(wc -l < {shlex.quote(str(calls))}) -eq 1 ]]; then touch {shlex.quote(str(entered))}; "
                f"while [[ ! -e {shlex.quote(str(release))} ]]; do sleep 0.02; done; fi\n",
                encoding="utf-8",
            )
            wrapper.chmod(0o755)
            watcher = subprocess.Popen(
                [
                    sys.executable,
                    str(WATCHER),
                    "--project",
                    str(root),
                    "--wrapper",
                    str(wrapper),
                    "--server",
                    "off",
                    "--action",
                    "compile",
                    "--poll-ms",
                    "20",
                    "--debounce-ms",
                    "40",
                    "--max-cycles",
                    "1",
                ],
                cwd=ROOT,
                stdout=subprocess.DEVNULL,
                stderr=subprocess.DEVNULL,
            )
            try:
                deadline = time.monotonic() + 5
                while not entered.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(entered.exists())
                source.write_text("NEW\n", encoding="utf-8")
                release.touch()
                watcher.wait(timeout=8)
                self.assertEqual(0, watcher.returncode)
                self.assertEqual(["OLD", "NEW"], calls.read_text().splitlines())
            finally:
                self.stop_watcher(watcher)

    def test_one_shot_wrapper_exposes_the_selected_plan_without_compiling(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-watch-plan-") as raw:
            project = Path(raw)
            (project / "src").mkdir()
            (project / "compile.hxml").write_text(
                "-cp src\n-D go_output=generated\n-main Main\n", encoding="utf-8"
            )
            result = subprocess.run(
                [
                    "bash",
                    str(WRAPPER),
                    "--project",
                    str(project),
                    "--action",
                    "run",
                    "--print-plan-json",
                ],
                cwd=ROOT,
                text=True,
                capture_output=True,
                check=True,
            )
            plan = json.loads(result.stdout)
            self.assertEqual(project.resolve(), Path(plan["project"]))
            self.assertEqual((project / "compile.hxml").resolve(), Path(plan["hxml"]))
            self.assertEqual((project / "generated").resolve(), Path(plan["goOutput"]))
            self.assertEqual("run", plan["action"])

    def test_commands_template_and_friendly_guide_are_wired(self) -> None:
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        template = json.loads(TEMPLATE_PACKAGE.read_text(encoding="utf-8"))
        doc = DOC.read_text(encoding="utf-8")
        self.assertEqual("python3 scripts/dev/haxe_go_watch.py", package["scripts"]["dev"])
        self.assertEqual(package["scripts"]["dev"], package["scripts"]["dev:watch"])
        self.assertEqual("python3 test/test_dev_watch.py", package["scripts"]["test:dev-watch"])
        self.assertIn("npm run test:dev-watch", package["scripts"]["test:strategy"])
        self.assertEqual("python3 scripts/dev/watch.py", template["scripts"]["dev"])
        self.assertTrue(TEMPLATE_WATCHER.exists())
        self.assertIn("## Start here", doc)
        self.assertIn("npm run dev -- --project", doc)
        self.assertIn("last working program keeps running", doc)
        self.assertIn("direct build remains the correctness baseline", doc)
        self.assertIn("haxe_go-vfp.12.8", doc)

    @unittest.skipUnless(
        shutil.which("haxe") and shutil.which("haxelib") and shutil.which("go"),
        "requires the real Haxe and Go toolchains",
    )
    def test_scaffold_command_compiles_through_the_real_target_and_runs_go(self) -> None:
        result = subprocess.run(
            [
                sys.executable,
                str(TEMPLATE_WATCHER),
                "--once",
                "--server",
                "off",
                "--action",
                "run",
            ],
            cwd=ROOT,
            text=True,
            capture_output=True,
            timeout=60,
            check=False,
        )
        self.assertEqual(0, result.returncode, result.stdout + result.stderr)
        self.assertIn("[hx-go] hxml=compile.hxml", result.stdout)
        self.assertIn("[hx-go] go output=./out", result.stdout)
        self.assertIn("hello from reflaxe.go", result.stdout)


if __name__ == "__main__":
    unittest.main()
