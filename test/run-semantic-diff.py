#!/usr/bin/env python3

from __future__ import annotations

import argparse
import dataclasses
import difflib
import json
import os
from pathlib import Path
import shutil
import subprocess
import time
from typing import Iterable

ROOT = Path(__file__).resolve().parent.parent
SEMANTIC_ROOT = ROOT / "test" / "semantic_diff"
CACHE_ROOT = ROOT / "test" / ".test-cache"
LAST_FAILED = CACHE_ROOT / "semantic_diff_last_failed.txt"
LAST_RUN = CACHE_ROOT / "semantic_diff_last_run.json"


@dataclasses.dataclass(frozen=True)
class SemanticCase:
    case_id: str
    case_path: Path
    main_hx: Path


@dataclasses.dataclass
class CaseResult:
    case_id: str
    ok: bool
    stage: str
    message: str
    duration_s: float


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run semantic differential tests (Haxe --interp vs reflaxe.go portable output)"
    )
    parser.add_argument("--list", action="store_true", help="List discovered semantic diff cases")
    parser.add_argument("--case", action="append", default=[], help="Run specific case id(s)")
    parser.add_argument("--pattern", default="", help="Regex filter over case ids")
    parser.add_argument("--changed", action="store_true", help="Run only semantic diff cases touched by git diff")
    parser.add_argument("--failed", action="store_true", help="Re-run only previously failing cases")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    return parser.parse_args()


def discover_cases() -> list[SemanticCase]:
    cases: list[SemanticCase] = []
    if not SEMANTIC_ROOT.exists():
        return cases

    for case_dir in sorted(SEMANTIC_ROOT.iterdir()):
        if not case_dir.is_dir():
            continue
        main_hx = case_dir / "Main.hx"
        if not main_hx.exists():
            continue
        cases.append(
            SemanticCase(
                case_id=case_dir.name,
                case_path=case_dir,
                main_hx=main_hx,
            )
        )

    return cases


def changed_case_ids() -> set[str]:
    cmd = ["git", "diff", "--name-only", "--", "test/semantic_diff"]
    try:
        proc = subprocess.run(cmd, cwd=ROOT, capture_output=True, text=True, check=True)
    except (FileNotFoundError, subprocess.CalledProcessError):
        return set()

    ids: set[str] = set()
    for line in proc.stdout.splitlines():
        path = Path(line.strip())
        parts = path.parts
        if len(parts) >= 3 and parts[0] == "test" and parts[1] == "semantic_diff":
            ids.add(parts[2])
    return ids


def read_last_failed() -> list[str]:
    if not LAST_FAILED.exists():
        return []
    return [line.strip() for line in LAST_FAILED.read_text(encoding="utf-8").splitlines() if line.strip()]


def apply_filters(cases: Iterable[SemanticCase], args: argparse.Namespace) -> list[SemanticCase]:
    selected = list(cases)

    if args.failed:
        failed = set(read_last_failed())
        selected = [case for case in selected if case.case_id in failed]

    if args.changed:
        changed = changed_case_ids()
        selected = [case for case in selected if case.case_id in changed]

    if args.case:
        wanted = {item.strip() for item in args.case if item.strip()}
        selected = [case for case in selected if case.case_id in wanted]

    if args.pattern:
        import re

        regex = re.compile(args.pattern)
        selected = [case for case in selected if regex.search(case.case_id)]

    return selected


def normalize_stdout(text: str) -> str:
    return text.replace("\r\n", "\n")


def run_command(cmd: list[str], cwd: Path, timeout_s: int, env: dict[str, str] | None = None) -> subprocess.CompletedProcess[str]:
    merged_env = os.environ.copy()
    merged_env["HAXE_NO_SERVER"] = "1"
    if env:
        merged_env.update(env)
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s, env=merged_env)


def command_output(proc: subprocess.CompletedProcess[str]) -> str:
    chunks: list[str] = []
    if proc.stdout:
        chunks.append(proc.stdout.strip())
    if proc.stderr:
        chunks.append(proc.stderr.strip())
    return "\n".join(chunk for chunk in chunks if chunk)


def build_interp_cmd(case: SemanticCase) -> list[str]:
    return [
        "haxe",
        "-cp",
        str(case.case_path),
        "-D",
        "no-traces",
        "-D",
        "no_traces",
        "-main",
        "Main",
        "--interp",
    ]


def build_go_cmd(case: SemanticCase) -> list[str]:
    out_dir = case.case_path / "out"
    return [
        "haxe",
        "-cp",
        str(case.case_path),
        "-cp",
        str(ROOT / "src"),
        "--macro",
        "reflaxe.go.CompilerBootstrap.Start()",
        "--macro",
        "reflaxe.go.CompilerInit.Start()",
        "-D",
        "go_output=" + str(out_dir),
        "-D",
        "reflaxe_go_profile=portable",
        "-D",
        "reflaxe_go_strict_examples",
        "-D",
        "reflaxe.dont_output_metadata_id",
        "-D",
        "no-traces",
        "-D",
        "no_traces",
        "-main",
        "Main",
    ]


def ensure_no_out(case: SemanticCase) -> None:
    out_dir = case.case_path / "out"
    if out_dir.exists():
        shutil.rmtree(out_dir)


def build_stdout_diff(reference: str, actual: str) -> str:
    ref_lines = reference.splitlines(keepends=True)
    out_lines = actual.splitlines(keepends=True)
    diff_lines = list(
        difflib.unified_diff(
            ref_lines,
            out_lines,
            fromfile="reference(--interp)",
            tofile="reflaxe.go(portable)",
        )
    )
    if not diff_lines:
        return "stdout mismatch (no textual diff available)"
    preview = "".join(diff_lines[:200]).rstrip("\n")
    return preview


def run_case(case: SemanticCase, args: argparse.Namespace) -> CaseResult:
    started = time.monotonic()
    try:
        ensure_no_out(case)

        ref_proc = run_command(build_interp_cmd(case), cwd=ROOT, timeout_s=args.timeout)
        if ref_proc.returncode != 0:
            return CaseResult(case.case_id, False, "reference", command_output(ref_proc), time.monotonic() - started)
        reference_stdout = normalize_stdout(ref_proc.stdout)

        go_compile_proc = run_command(build_go_cmd(case), cwd=ROOT, timeout_s=args.timeout)
        if go_compile_proc.returncode != 0:
            return CaseResult(case.case_id, False, "compile", command_output(go_compile_proc), time.monotonic() - started)

        out_dir = case.case_path / "out"
        go_test_proc = run_command(["go", "test", "./..."], cwd=out_dir, timeout_s=args.timeout)
        if go_test_proc.returncode != 0:
            return CaseResult(case.case_id, False, "go test", command_output(go_test_proc), time.monotonic() - started)

        go_run_proc = run_command(["go", "run", "."], cwd=out_dir, timeout_s=args.timeout)
        if go_run_proc.returncode != 0:
            return CaseResult(case.case_id, False, "runtime", command_output(go_run_proc), time.monotonic() - started)
        go_stdout = normalize_stdout(go_run_proc.stdout)

        if reference_stdout != go_stdout:
            return CaseResult(
                case.case_id,
                False,
                "diff",
                build_stdout_diff(reference_stdout, go_stdout),
                time.monotonic() - started,
            )

        return CaseResult(case.case_id, True, "done", "ok", time.monotonic() - started)
    except subprocess.TimeoutExpired as exc:
        return CaseResult(
            case.case_id,
            False,
            "timeout",
            f"command timed out after {args.timeout}s: {exc.cmd}",
            time.monotonic() - started,
        )


def write_last_failed(results: list[CaseResult]) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    failed = sorted(result.case_id for result in results if not result.ok)
    LAST_FAILED.write_text("\n".join(failed) + ("\n" if failed else ""), encoding="utf-8")


def write_last_run(results: list[CaseResult]) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    payload = {
        "generated_at_epoch": int(time.time()),
        "total": len(results),
        "passed": sum(1 for result in results if result.ok),
        "failed": sum(1 for result in results if not result.ok),
        "results": [dataclasses.asdict(result) for result in results],
    }
    LAST_RUN.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()
    cases = discover_cases()

    if args.list:
        for case in cases:
            print(case.case_id)
        return 0

    selected = apply_filters(cases, args)
    if not selected:
        print("No semantic diff cases selected")
        return 0

    results: list[CaseResult] = []
    for case in selected:
        print(f"==> {case.case_id}")
        result = run_case(case, args)
        results.append(result)
        status = "PASS" if result.ok else "FAIL"
        print(f"[{status}] {case.case_id} ({result.stage}, {result.duration_s:.2f}s)")
        if not result.ok and result.message:
            print(result.message)

    write_last_failed(results)
    write_last_run(results)

    passed = sum(1 for result in results if result.ok)
    failed = len(results) - passed
    print(f"\nSummary: {passed} passed, {failed} failed, {len(results)} total")
    print(f"Last run report: {LAST_RUN}")

    return 1 if failed else 0


if __name__ == "__main__":
    raise SystemExit(main())
