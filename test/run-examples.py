#!/usr/bin/env python3

from __future__ import annotations

import argparse
import dataclasses
import os
from pathlib import Path
import shutil
import subprocess
import time

ROOT = Path(__file__).resolve().parent.parent
EXAMPLES_ROOT = ROOT / "examples"
PROFILES = ("portable", "gopher", "metal")
EXCLUDE_NAMES = {"go.sum", "_GeneratedFiles.json"}
EXCLUDE_DIRS = {".cache"}


@dataclasses.dataclass(frozen=True)
class ExampleProfileCase:
    example: str
    profile: str
    example_dir: Path
    compile_hxml: Path
    compile_ci_hxml: Path
    out_dir: Path
    out_ci_dir: Path
    expected_stdout: Path
    expected_ci_stdout: Path
    generated_dir: Path

    @property
    def case_id(self) -> str:
        return f"{self.example}/{self.profile}"


@dataclasses.dataclass
class CaseResult:
    case_id: str
    ok: bool
    stage: str
    message: str
    duration_s: float


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run reflaxe.go examples profile matrix")
    parser.add_argument("--list", action="store_true", help="List discovered example/profile cases")
    parser.add_argument("--example", action="append", default=[], help="Example name filter (repeatable)")
    parser.add_argument("--profile", action="append", default=[], help="Profile filter (repeatable)")
    parser.add_argument("--changed", action="store_true", help="Run only changed examples")
    parser.add_argument("--compile-only", action="store_true", help="Skip go run stdout checks")
    parser.add_argument("--bless-generated", action="store_true", help="Refresh generated/<profile> from out_<profile>")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    return parser.parse_args()


def discover_cases() -> list[ExampleProfileCase]:
    cases: list[ExampleProfileCase] = []
    if not EXAMPLES_ROOT.exists():
        return cases

    for example_dir in sorted(EXAMPLES_ROOT.iterdir()):
        if not example_dir.is_dir():
            continue
        for profile in PROFILES:
            compile_hxml = example_dir / f"compile.{profile}.hxml"
            compile_ci_hxml = example_dir / f"compile.{profile}.ci.hxml"
            if not compile_hxml.exists() or not compile_ci_hxml.exists():
                continue
            cases.append(
                ExampleProfileCase(
                    example=example_dir.name,
                    profile=profile,
                    example_dir=example_dir,
                    compile_hxml=compile_hxml,
                    compile_ci_hxml=compile_ci_hxml,
                    out_dir=example_dir / f"out_{profile}",
                    out_ci_dir=example_dir / f"out_{profile}_ci",
                    expected_stdout=example_dir / "expected" / f"{profile}.stdout",
                    expected_ci_stdout=example_dir / "expected" / f"{profile}.ci.stdout",
                    generated_dir=example_dir / "generated" / profile,
                )
            )
    return cases


def changed_examples() -> set[str]:
    cmd = ["git", "diff", "--name-only", "--", "examples"]
    try:
        proc = subprocess.run(cmd, cwd=ROOT, capture_output=True, text=True, check=True)
    except (FileNotFoundError, subprocess.CalledProcessError):
        return set()

    out: set[str] = set()
    for line in proc.stdout.splitlines():
        path = Path(line.strip())
        parts = path.parts
        if len(parts) >= 2 and parts[0] == "examples":
            out.add(parts[1])
    return out


def apply_filters(cases: list[ExampleProfileCase], args: argparse.Namespace) -> list[ExampleProfileCase]:
    selected = list(cases)

    if args.example:
        wanted = {item.strip() for item in args.example if item.strip()}
        selected = [case for case in selected if case.example in wanted]

    if args.profile:
        wanted = {item.strip() for item in args.profile if item.strip()}
        selected = [case for case in selected if case.profile in wanted]

    if args.changed:
        changed = changed_examples()
        selected = [case for case in selected if case.example in changed]

    return selected


def run_command(cmd: list[str], cwd: Path, timeout_s: int, env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
    merged = os.environ.copy()
    if env:
        merged.update(env)
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s, env=merged)


def command_output(proc: subprocess.CompletedProcess[str]) -> str:
    chunks: list[str] = []
    if proc.stdout:
        chunks.append(proc.stdout.strip())
    if proc.stderr:
        chunks.append(proc.stderr.strip())
    return "\n".join(chunk for chunk in chunks if chunk)


def all_files(root: Path) -> list[Path]:
    if not root.exists():
        return []
    files: list[Path] = []
    for path in sorted(root.rglob("*")):
        if path.is_dir():
            continue
        if path.name in EXCLUDE_NAMES:
            continue
        if any(part in EXCLUDE_DIRS for part in path.parts):
            continue
        files.append(path)
    return files


def collect_tree_deltas(left: Path, right: Path) -> list[str]:
    left_files = {path.relative_to(left): path for path in all_files(left)} if left.exists() else {}
    right_files = {path.relative_to(right): path for path in all_files(right)} if right.exists() else {}

    rels = sorted(set(left_files) | set(right_files))
    deltas: list[str] = []
    for rel in rels:
        l = left_files.get(rel)
        r = right_files.get(rel)
        if l is None:
            deltas.append(f"Only in {right}: {rel.as_posix()}")
            continue
        if r is None:
            deltas.append(f"Only in {left}: {rel.as_posix()}")
            continue
        ltxt = l.read_text(encoding="utf-8", errors="replace")
        rtxt = r.read_text(encoding="utf-8", errors="replace")
        if ltxt != rtxt:
            deltas.append(f"Diff: {rel.as_posix()}")
    return deltas


def copy_tree(source: Path, target: Path) -> None:
    if target.exists():
        shutil.rmtree(target)
    target.mkdir(parents=True, exist_ok=True)
    for path in all_files(source):
        rel = path.relative_to(source)
        dest = target / rel
        dest.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(path, dest)


def clean_out_dirs(case: ExampleProfileCase) -> None:
    if case.out_dir.exists():
        shutil.rmtree(case.out_dir)
    if case.out_ci_dir.exists():
        shutil.rmtree(case.out_ci_dir)


def run_go_checks(out_dir: Path, timeout_s: int) -> tuple[bool, str, str]:
    go_files = sorted(out_dir.rglob("*.go")) if out_dir.exists() else []
    if not go_files:
        return False, "go", "no generated .go files"

    gofmt_cmd = ["gofmt", "-w"] + [str(path) for path in go_files]
    gofmt_proc = run_command(gofmt_cmd, cwd=out_dir, timeout_s=timeout_s)
    if gofmt_proc.returncode != 0:
        return False, "gofmt", command_output(gofmt_proc)

    gotest_proc = run_command(["go", "test", "./..."], cwd=out_dir, timeout_s=timeout_s)
    if gotest_proc.returncode != 0:
        return False, "go test", command_output(gotest_proc)

    return True, "go", ""


def compare_stdout(expected_file: Path, actual: str) -> tuple[bool, str]:
    if not expected_file.exists():
        return False, f"missing expected file: {expected_file}"

    expected = expected_file.read_text(encoding="utf-8", errors="replace").replace("\r\n", "\n")
    actual_norm = actual.replace("\r\n", "\n")
    if expected == actual_norm:
        return True, ""

    return False, "stdout mismatch"


def run_case(case: ExampleProfileCase, args: argparse.Namespace) -> CaseResult:
    started = time.monotonic()
    try:
        clean_out_dirs(case)

        compile_ci_proc = run_command(["haxe", case.compile_ci_hxml.name], cwd=case.example_dir, timeout_s=args.timeout, env={"HAXE_NO_SERVER": "1"})
        if compile_ci_proc.returncode != 0:
            return CaseResult(case.case_id, False, "compile_ci", command_output(compile_ci_proc), time.monotonic() - started)

        ok, stage, msg = run_go_checks(case.out_ci_dir, args.timeout)
        if not ok:
            return CaseResult(case.case_id, False, stage + "_ci", msg, time.monotonic() - started)

        if not args.compile_only:
            run_ci_proc = run_command(["go", "run", "."], cwd=case.out_ci_dir, timeout_s=args.timeout)
            if run_ci_proc.returncode != 0:
                return CaseResult(case.case_id, False, "runtime_ci", command_output(run_ci_proc), time.monotonic() - started)
            ok_stdout, msg_stdout = compare_stdout(case.expected_ci_stdout, run_ci_proc.stdout)
            if not ok_stdout:
                return CaseResult(case.case_id, False, "stdout_ci", msg_stdout, time.monotonic() - started)

        compile_proc = run_command(["haxe", case.compile_hxml.name], cwd=case.example_dir, timeout_s=args.timeout, env={"HAXE_NO_SERVER": "1"})
        if compile_proc.returncode != 0:
            return CaseResult(case.case_id, False, "compile", command_output(compile_proc), time.monotonic() - started)

        ok, stage, msg = run_go_checks(case.out_dir, args.timeout)
        if not ok:
            return CaseResult(case.case_id, False, stage, msg, time.monotonic() - started)

        if not args.compile_only:
            run_proc = run_command(["go", "run", "."], cwd=case.out_dir, timeout_s=args.timeout)
            if run_proc.returncode != 0:
                return CaseResult(case.case_id, False, "runtime", command_output(run_proc), time.monotonic() - started)
            ok_stdout, msg_stdout = compare_stdout(case.expected_stdout, run_proc.stdout)
            if not ok_stdout:
                return CaseResult(case.case_id, False, "stdout", msg_stdout, time.monotonic() - started)

        if args.bless_generated:
            copy_tree(case.out_dir, case.generated_dir)

        if not case.generated_dir.exists():
            return CaseResult(case.case_id, False, "generated", f"missing generated directory: {case.generated_dir} (use --bless-generated)", time.monotonic() - started)

        deltas = collect_tree_deltas(case.generated_dir, case.out_dir)
        if deltas:
            preview = "\n".join(deltas[:20])
            if len(deltas) > 20:
                preview += f"\n... and {len(deltas) - 20} more"
            return CaseResult(case.case_id, False, "generated", preview, time.monotonic() - started)

        return CaseResult(case.case_id, True, "done", "ok", time.monotonic() - started)

    except subprocess.TimeoutExpired as exc:
        return CaseResult(case.case_id, False, "timeout", f"command timed out after {args.timeout}s: {exc.cmd}", time.monotonic() - started)
    except FileNotFoundError as exc:
        return CaseResult(case.case_id, False, "tool", f"missing tool: {exc}", time.monotonic() - started)


def main() -> int:
    args = parse_args()

    cases = discover_cases()
    if args.list:
        for case in cases:
            print(case.case_id)
        return 0

    selected = apply_filters(cases, args)
    if not selected:
        print("No example cases selected")
        return 0

    results: list[CaseResult] = []
    for case in selected:
        print(f"==> {case.case_id}")
        result = run_case(case, args)
        results.append(result)
        status = "PASS" if result.ok else "FAIL"
        print(f"[{status}] {case.case_id} ({result.stage}, {result.duration_s:.2f}s)")
        if result.message and (not result.ok):
            print(result.message)

    passed = sum(1 for result in results if result.ok)
    failed = len(results) - passed
    print(f"\nSummary: {passed} passed, {failed} failed, {len(results)} total")

    return 1 if failed else 0


if __name__ == "__main__":
    raise SystemExit(main())
