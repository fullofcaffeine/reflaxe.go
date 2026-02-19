#!/usr/bin/env python3

from __future__ import annotations

import argparse
import concurrent.futures
import contextlib
import dataclasses
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
import time
from typing import Iterable

try:
    import fcntl  # type: ignore[attr-defined]
except ImportError:
    fcntl = None

ROOT = Path(__file__).resolve().parent.parent
SNAPSHOT_ROOT = ROOT / "test" / "snapshot"
CACHE_ROOT = ROOT / "test" / ".test-cache"
LAST_FAILED = CACHE_ROOT / "last_failed.txt"
LAST_RUN = CACHE_ROOT / "last_run.json"
RUN_LOCK = CACHE_ROOT / "run-snapshots.lock"
EXCLUDE_NAMES = {"go.sum", "_GeneratedFiles.json", ".DS_Store", ".gitkeep"}
EXCLUDE_DIRS = {".cache"}


@dataclasses.dataclass(frozen=True)
class SnapshotCase:
    case_id: str
    case_path: Path
    compile_hxml: Path


@dataclasses.dataclass
class CaseResult:
    case_id: str
    ok: bool
    duration_s: float
    stage: str
    message: str


@dataclasses.dataclass(frozen=True)
class TreeDelta:
    rel_path: Path
    kind: str  # added | removed | modified


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run reflaxe.go snapshot tests")
    parser.add_argument("--list", action="store_true", help="List discovered cases")
    parser.add_argument("--category", default="", help="Comma-separated category filter")
    parser.add_argument("--case", action="append", default=[], help="Run specific case (category/name)")
    parser.add_argument("--pattern", default="", help="Regex filter over case ids")
    parser.add_argument("--jobs", type=int, default=1, help="Parallel jobs")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    update_group = parser.add_mutually_exclusive_group()
    update_group.add_argument("--update", action="store_true", help="Replace intended/ outputs from out/")
    update_group.add_argument("--bless", action="store_true", help="Update only changed files in intended/")
    parser.add_argument("--runtime", action="store_true", help="Run runtime smoke when expected.stdout exists")
    parser.add_argument("--failed", action="store_true", help="Re-run only previously failing cases")
    parser.add_argument("--changed", action="store_true", help="Run cases touched by git diff")
    parser.add_argument("--chunk", default="", help="Deterministic shard in i/n form (e.g. 0/4)")
    parser.add_argument("--lock-timeout", type=int, default=30, help="Seconds to wait for harness lock (0 = fail fast)")
    return parser.parse_args()


@contextlib.contextmanager
def acquire_run_lock(timeout_s: int):
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    if fcntl is not None:
        lock_file = RUN_LOCK.open("a+", encoding="utf-8")
        start = time.monotonic()
        while True:
            try:
                fcntl.flock(lock_file.fileno(), fcntl.LOCK_EX | fcntl.LOCK_NB)
                break
            except BlockingIOError:
                if timeout_s <= 0 or (time.monotonic() - start) >= timeout_s:
                    lock_file.close()
                    raise SystemExit(
                        f"Another snapshot run is active (lock: {RUN_LOCK}). "
                        "Wait and retry, or set --lock-timeout to a larger value."
                    )
                time.sleep(0.2)

        try:
            lock_file.seek(0)
            lock_file.truncate(0)
            lock_file.write(f"pid={os.getpid()}\n")
            lock_file.flush()
            yield
        finally:
            try:
                lock_file.seek(0)
                lock_file.truncate(0)
                lock_file.flush()
            finally:
                fcntl.flock(lock_file.fileno(), fcntl.LOCK_UN)
                lock_file.close()
        return

    # Fallback path (platforms without fcntl).
    start = time.monotonic()
    while True:
        try:
            fd = os.open(str(RUN_LOCK), os.O_CREAT | os.O_EXCL | os.O_WRONLY)
            break
        except FileExistsError:
            if timeout_s <= 0 or (time.monotonic() - start) >= timeout_s:
                raise SystemExit(
                    f"Another snapshot run is active (lock: {RUN_LOCK}). "
                    "Wait and retry, or set --lock-timeout to a larger value."
                )
            time.sleep(0.2)

    try:
        os.write(fd, f"pid={os.getpid()}\n".encode("utf-8"))
        yield
    finally:
        os.close(fd)
        try:
            RUN_LOCK.unlink()
        except FileNotFoundError:
            pass


def discover_cases() -> list[SnapshotCase]:
    cases: list[SnapshotCase] = []
    for hxml in sorted(SNAPSHOT_ROOT.rglob("compile.hxml")):
        if "_archive" in hxml.parts:
            continue
        case_dir = hxml.parent
        rel = case_dir.relative_to(SNAPSHOT_ROOT)
        if len(rel.parts) < 2:
            continue
        case_id = f"{rel.parts[0]}/{rel.parts[1]}"
        cases.append(SnapshotCase(case_id=case_id, case_path=case_dir, compile_hxml=hxml))
    cases.sort(key=lambda c: c.case_id)
    return cases


def apply_filters(cases: Iterable[SnapshotCase], args: argparse.Namespace) -> list[SnapshotCase]:
    selected = list(cases)

    if args.failed:
        selected_ids = set(read_last_failed())
        selected = [c for c in selected if c.case_id in selected_ids]

    if args.changed:
        selected_ids = changed_case_ids()
        selected = [c for c in selected if c.case_id in selected_ids]

    if args.category:
        allowed = {token.strip() for token in args.category.split(",") if token.strip()}
        selected = [c for c in selected if c.case_id.split("/", 1)[0] in allowed]

    if args.case:
        wanted = set(args.case)
        selected = [c for c in selected if c.case_id in wanted]

    if args.pattern:
        regex = re.compile(args.pattern)
        selected = [c for c in selected if regex.search(c.case_id)]

    if args.chunk:
        selected = select_chunk(selected, args.chunk)

    return selected


def select_chunk(cases: list[SnapshotCase], chunk_spec: str) -> list[SnapshotCase]:
    match = re.fullmatch(r"(\d+)/(\d+)", chunk_spec.strip())
    if not match:
        raise SystemExit(f"Invalid --chunk value '{chunk_spec}', expected i/n")

    idx = int(match.group(1))
    total = int(match.group(2))
    if total <= 0 or idx < 0 or idx >= total:
        raise SystemExit(f"Invalid --chunk value '{chunk_spec}', expected 0 <= i < n")

    out: list[SnapshotCase] = []
    for case in cases:
        digest = hashlib.sha256(case.case_id.encode("utf-8")).hexdigest()
        bucket = int(digest[:8], 16) % total
        if bucket == idx:
            out.append(case)
    return out


def changed_case_ids() -> set[str]:
    cmd = ["git", "diff", "--name-only", "--", "test/snapshot"]
    try:
        proc = subprocess.run(cmd, cwd=ROOT, capture_output=True, text=True, check=True)
    except (FileNotFoundError, subprocess.CalledProcessError):
        return set()

    case_ids: set[str] = set()
    for line in proc.stdout.splitlines():
        path = Path(line.strip())
        parts = path.parts
        if len(parts) >= 4 and parts[0] == "test" and parts[1] == "snapshot":
            case_ids.add(f"{parts[2]}/{parts[3]}")
    return case_ids


def read_last_failed() -> list[str]:
    if not LAST_FAILED.exists():
        return []
    return [line.strip() for line in LAST_FAILED.read_text(encoding="utf-8").splitlines() if line.strip()]


def run_command(cmd: list[str], cwd: Path, timeout_s: int, env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
    merged_env = os.environ.copy()
    if env:
        merged_env.update(env)
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s, env=merged_env)


def command_output(proc: subprocess.CompletedProcess[str]) -> str:
    chunks = []
    if proc.stdout:
        chunks.append(proc.stdout.strip())
    if proc.stderr:
        chunks.append(proc.stderr.strip())
    return "\n".join(chunk for chunk in chunks if chunk)


def should_expect_compile_failure(case: SnapshotCase) -> bool:
    return (case.case_path / "expected.error.txt").exists() or case.case_id.startswith("negative/")


def validate_expected_error(case: SnapshotCase, output: str) -> tuple[bool, str]:
    expected_file = case.case_path / "expected.error.txt"
    if not expected_file.exists():
        return True, ""

    needle = expected_file.read_text(encoding="utf-8").strip()
    if not needle:
        return True, ""

    if needle in output:
        return True, ""

    return False, f"compile error mismatch, expected substring not found: {needle!r}"


def clean_out_dir(case: SnapshotCase) -> None:
    out_dir = case.case_path / "out"
    if out_dir.exists():
        shutil.rmtree(out_dir, ignore_errors=True)


def all_files(root: Path) -> list[Path]:
    if not root.exists():
        return []

    files: list[Path] = []
    for path in sorted(root.rglob("*")):
        if path.is_dir():
            if path.name in EXCLUDE_DIRS:
                continue
            continue
        if path.name in EXCLUDE_NAMES:
            continue
        if any(part in EXCLUDE_DIRS for part in path.parts):
            continue
        files.append(path)
    return files


def collect_tree_deltas(left: Path, right: Path) -> list[TreeDelta]:
    left_files = {p.relative_to(left): p for p in all_files(left)}
    right_files = {p.relative_to(right): p for p in all_files(right)}
    rels = sorted(set(left_files) | set(right_files))

    deltas: list[TreeDelta] = []
    for rel in rels:
        l = left_files.get(rel)
        r = right_files.get(rel)
        if l is None:
            deltas.append(TreeDelta(rel, "added"))
            continue
        if r is None:
            deltas.append(TreeDelta(rel, "removed"))
            continue

        ltxt = l.read_text(encoding="utf-8", errors="replace")
        rtxt = r.read_text(encoding="utf-8", errors="replace")
        if ltxt != rtxt:
            deltas.append(TreeDelta(rel, "modified"))
    return deltas


def diff_trees(left: Path, right: Path) -> tuple[bool, str]:
    deltas = collect_tree_deltas(left, right)
    if not deltas:
        return True, ""

    lines: list[str] = []
    for delta in deltas:
        if delta.kind == "added":
            lines.append(f"Only in {right}: {delta.rel_path.as_posix()}")
        elif delta.kind == "removed":
            lines.append(f"Only in {left}: {delta.rel_path.as_posix()}")
        else:
            lines.append(f"Diff: {delta.rel_path.as_posix()}")
    return False, "\n".join(lines)


def update_intended(case: SnapshotCase) -> None:
    out_dir = case.case_path / "out"
    intended = case.case_path / "intended"
    if intended.exists():
        shutil.rmtree(intended)
    intended.mkdir(parents=True, exist_ok=True)

    if not out_dir.exists():
        return

    for source in all_files(out_dir):
        rel = source.relative_to(out_dir)
        target = intended / rel
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)


def read_text_or_none(path: Path) -> str | None:
    if not path.exists():
        return None
    return path.read_text(encoding="utf-8", errors="replace")


def collapse_whitespace(text: str) -> str:
    return "".join(text.split())


def infer_bless_causes(rel_path: Path, old_text: str | None, new_text: str | None) -> tuple[bool, bool, bool]:
    runtime_change = bool(rel_path.parts) and rel_path.parts[0] == "hxrt"
    combined = f"{old_text or ''}\n{new_text or ''}"
    if not runtime_change and ("/hxrt" in combined or '"hxrt"' in combined or "hxrt." in combined):
        runtime_change = True

    gofmt_change = False
    if old_text is not None and new_text is not None and old_text != new_text:
        gofmt_change = collapse_whitespace(old_text) == collapse_whitespace(new_text)

    naming_change = False
    if old_text is not None and new_text is not None and old_text != new_text:
        old_tokens = set(re.findall(r"\b[A-Za-z_][A-Za-z0-9_]*\b", old_text))
        new_tokens = set(re.findall(r"\b[A-Za-z_][A-Za-z0-9_]*\b", new_text))
        changed_tokens = (old_tokens - new_tokens) | (new_tokens - old_tokens)
        naming_change = any(
            token.startswith(("Haxe_", "__hx_", "Go_", "Std_", "Main_"))
            for token in changed_tokens
        )

    return gofmt_change, naming_change, runtime_change


def format_bless_checklist(rel_path: Path, old_text: str | None, new_text: str | None) -> str:
    gofmt_change, naming_change, runtime_change = infer_bless_causes(rel_path, old_text, new_text)
    return (
        f"[gofmt={'x' if gofmt_change else ' '}] "
        f"[naming={'x' if naming_change else ' '}] "
        f"[runtime={'x' if runtime_change else ' '}]"
    )


def prune_empty_dirs(start: Path, stop_at: Path) -> None:
    cur = start
    while cur != stop_at and cur.exists():
        try:
            cur.rmdir()
        except OSError:
            return
        cur = cur.parent


def bless_intended(case: SnapshotCase, deltas: list[TreeDelta]) -> str:
    out_dir = case.case_path / "out"
    intended = case.case_path / "intended"
    intended.mkdir(parents=True, exist_ok=True)

    lines: list[str] = []
    max_lines = 20

    for idx, delta in enumerate(deltas):
        rel = delta.rel_path
        source = out_dir / rel
        target = intended / rel
        before = read_text_or_none(target)
        after = read_text_or_none(source) if delta.kind != "removed" else None

        if delta.kind in {"added", "modified"}:
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(source, target)
        else:
            if target.exists():
                target.unlink()
                prune_empty_dirs(target.parent, intended)

        if idx < max_lines:
            checklist = format_bless_checklist(rel, before, after)
            lines.append(f"{delta.kind:8s} {rel.as_posix()} {checklist}")

    extra = len(deltas) - max_lines
    if extra > 0:
        lines.append(f"... and {extra} more changed file(s)")

    headline = f"blessed {len(deltas)} changed file(s)"
    if not lines:
        return headline
    return f"{headline}\n" + "\n".join(lines)


def maybe_cleanup_artifacts(case: SnapshotCase, success: bool) -> None:
    keep_artifacts = os.environ.get("KEEP_ARTIFACTS", "") == "1"
    if keep_artifacts:
        return
    if success:
        out_dir = case.case_path / "out"
        if out_dir.exists():
            shutil.rmtree(out_dir, ignore_errors=True)


def run_case(case: SnapshotCase, args: argparse.Namespace) -> CaseResult:
    started = time.monotonic()

    try:
        clean_out_dir(case)

        compile_proc = run_command(
            ["haxe", "compile.hxml", "-D", "go_no_build"],
            cwd=case.case_path,
            timeout_s=args.timeout,
            env={"HAXE_NO_SERVER": "1"},
        )
        compile_output = command_output(compile_proc)
        expects_failure = should_expect_compile_failure(case)

        if compile_proc.returncode != 0:
            if expects_failure:
                ok, msg = validate_expected_error(case, compile_output)
                duration = time.monotonic() - started
                if ok:
                    return CaseResult(case.case_id, True, duration, "compile", "expected compile failure")
                return CaseResult(case.case_id, False, duration, "compile", msg + "\n" + compile_output)
            duration = time.monotonic() - started
            return CaseResult(case.case_id, False, duration, "compile", compile_output)

        if expects_failure:
            duration = time.monotonic() - started
            return CaseResult(case.case_id, False, duration, "compile", "expected compile failure, but compile succeeded")

        out_dir = case.case_path / "out"
        go_files = list(out_dir.rglob("*.go")) if out_dir.exists() else []

        if go_files:
            gofmt_proc = run_command(["gofmt", "-w", *[str(p) for p in sorted(go_files)]], cwd=case.case_path, timeout_s=args.timeout)
            if gofmt_proc.returncode != 0:
                duration = time.monotonic() - started
                return CaseResult(case.case_id, False, duration, "gofmt", command_output(gofmt_proc))

            gotest_proc = run_command(["go", "test", "./..."], cwd=out_dir, timeout_s=args.timeout)
            if gotest_proc.returncode != 0:
                duration = time.monotonic() - started
                return CaseResult(case.case_id, False, duration, "go test", command_output(gotest_proc))

        intended = case.case_path / "intended"
        if args.update:
            update_intended(case)
        elif args.bless:
            deltas = collect_tree_deltas(intended, out_dir)
            if deltas:
                bless_message = bless_intended(case, deltas)
                ok, diff = diff_trees(intended, out_dir)
                if not ok:
                    duration = time.monotonic() - started
                    return CaseResult(case.case_id, False, duration, "bless", f"bless verification failed\n{diff}")
                duration = time.monotonic() - started
                return CaseResult(case.case_id, True, duration, "bless", bless_message)
        elif intended.exists():
            ok, diff = diff_trees(intended, out_dir)
            if not ok:
                duration = time.monotonic() - started
                return CaseResult(case.case_id, False, duration, "diff", diff)
        else:
            duration = time.monotonic() - started
            return CaseResult(case.case_id, False, duration, "diff", "missing intended/ (use --update)")

        if args.runtime:
            expected_stdout = case.case_path / "expected.stdout"
            if expected_stdout.exists():
                run_proc = run_command(["go", "run", "."], cwd=out_dir, timeout_s=args.timeout)
                if run_proc.returncode != 0:
                    duration = time.monotonic() - started
                    return CaseResult(case.case_id, False, duration, "runtime", command_output(run_proc))

                expected = expected_stdout.read_text(encoding="utf-8").replace("\r\n", "\n")
                actual = run_proc.stdout.replace("\r\n", "\n")
                if actual != expected:
                    duration = time.monotonic() - started
                    return CaseResult(case.case_id, False, duration, "runtime", "stdout mismatch")

        duration = time.monotonic() - started
        return CaseResult(case.case_id, True, duration, "done", "ok")

    except subprocess.TimeoutExpired as exc:
        duration = time.monotonic() - started
        return CaseResult(case.case_id, False, duration, "timeout", f"command timed out after {args.timeout}s: {exc.cmd}")
    except FileNotFoundError as exc:
        duration = time.monotonic() - started
        return CaseResult(case.case_id, False, duration, "tool", f"missing tool: {exc}")
    finally:
        # Clean only on success unless KEEP_ARTIFACTS=1.
        pass


def write_state(results: list[CaseResult]) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)

    failed = [r.case_id for r in results if not r.ok]
    LAST_FAILED.write_text("\n".join(failed) + ("\n" if failed else ""), encoding="utf-8")

    payload = {
        "generated_at_epoch": int(time.time()),
        "total": len(results),
        "passed": sum(1 for r in results if r.ok),
        "failed": len(failed),
        "results": [dataclasses.asdict(r) for r in results],
    }
    LAST_RUN.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()
    if args.jobs < 1:
        raise SystemExit("--jobs must be >= 1")

    all_cases = discover_cases()
    if args.list:
        for case in all_cases:
            print(case.case_id)
        return 0

    selected = apply_filters(all_cases, args)
    if not selected:
        print("No snapshot cases selected")
        return 0

    results: list[CaseResult] = []
    with acquire_run_lock(args.lock_timeout):
        if args.jobs == 1:
            for case in selected:
                print(f"==> {case.case_id}")
                result = run_case(case, args)
                results.append(result)
                status = "PASS" if result.ok else "FAIL"
                print(f"[{status}] {case.case_id} ({result.stage}, {result.duration_s:.2f}s)")
                if result.message and (not result.ok or result.stage == "bless"):
                    print(result.message)
        else:
            print(f"Running {len(selected)} case(s) with {args.jobs} workers")
            with concurrent.futures.ThreadPoolExecutor(max_workers=args.jobs) as executor:
                future_by_id: dict[str, concurrent.futures.Future[CaseResult]] = {}
                for case in selected:
                    future_by_id[case.case_id] = executor.submit(run_case, case, args)

                for case in selected:
                    print(f"==> {case.case_id}")
                    try:
                        result = future_by_id[case.case_id].result()
                    except Exception as exc:
                        result = CaseResult(case.case_id, False, 0.0, "internal", f"worker crashed: {exc}")
                    results.append(result)
                    status = "PASS" if result.ok else "FAIL"
                    print(f"[{status}] {case.case_id} ({result.stage}, {result.duration_s:.2f}s)")
                    if result.message and (not result.ok or result.stage == "bless"):
                        print(result.message)

    for result in results:
        maybe_cleanup_artifacts(next(c for c in selected if c.case_id == result.case_id), result.ok)

    write_state(results)

    passed = sum(1 for r in results if r.ok)
    failed = len(results) - passed
    print(f"\nSummary: {passed} passed, {failed} failed, {len(results)} total")

    return 1 if failed else 0


if __name__ == "__main__":
    sys.exit(main())
