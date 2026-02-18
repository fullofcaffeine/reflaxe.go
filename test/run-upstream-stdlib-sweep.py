#!/usr/bin/env python3

from __future__ import annotations

import argparse
import dataclasses
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import time

ROOT = Path(__file__).resolve().parent.parent
CACHE_ROOT = ROOT / "test" / ".test-cache"
WORK_ROOT = CACHE_ROOT / "upstream-stdlib-sweep"
DEFAULT_MODULES_FILE = ROOT / "test" / "upstream_std_modules.txt"
LAST_RUN = CACHE_ROOT / "upstream_std_sweep_last.json"

PROBE_TYPE_OVERRIDES = {
    "haxe.ds.IntMap": "haxe.ds.IntMap<Dynamic>",
    "haxe.ds.StringMap": "haxe.ds.StringMap<Dynamic>",
    "haxe.ds.ObjectMap": "haxe.ds.ObjectMap<Dynamic, Dynamic>",
    "haxe.ds.EnumValueMap": "haxe.ds.EnumValueMap<EnumValue, Dynamic>",
    "haxe.ds.List": "haxe.ds.List<Dynamic>",
}


@dataclasses.dataclass
class SweepResult:
    module: str
    ok: bool
    stage: str
    duration_s: float
    message: str


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run upstream stdlib module sweep against reflaxe.go")
    parser.add_argument("--modules-file", default=str(DEFAULT_MODULES_FILE), help="Path to module list file")
    parser.add_argument("--module", action="append", default=[], help="Run only specific module(s)")
    parser.add_argument("--pattern", default="", help="Regex filter for module names")
    parser.add_argument("--list", action="store_true", help="List selected modules and exit")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    parser.add_argument("--go-test", action="store_true", help="Also run `go test ./...` in generated out/")
    parser.add_argument("--strict", action="store_true", help="Exit non-zero if any module fails")
    return parser.parse_args()


def load_modules(path: Path) -> list[str]:
    if not path.exists():
        raise SystemExit(f"Modules file not found: {path}")

    modules: list[str] = []
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        modules.append(line)

    # Keep first occurrence order while deduplicating.
    seen: set[str] = set()
    unique: list[str] = []
    for module in modules:
        if module in seen:
            continue
        seen.add(module)
        unique.append(module)
    return unique


def select_modules(all_modules: list[str], args: argparse.Namespace) -> list[str]:
    selected = list(all_modules)

    if args.module:
        wanted = set(args.module)
        selected = [module for module in selected if module in wanted]

    if args.pattern:
        regex = re.compile(args.pattern)
        selected = [module for module in selected if regex.search(module)]

    return selected


def command_output(proc: subprocess.CompletedProcess[str]) -> str:
    chunks = []
    if proc.stdout:
        chunks.append(proc.stdout.strip())
    if proc.stderr:
        chunks.append(proc.stderr.strip())
    return "\n".join(chunk for chunk in chunks if chunk)


def run_command(cmd: list[str], cwd: Path, timeout_s: int) -> subprocess.CompletedProcess[str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s, env=env)


def safe_name(module: str) -> str:
    return re.sub(r"[^A-Za-z0-9_.-]", "_", module)


def probe_type_for(module: str) -> str:
    return PROBE_TYPE_OVERRIDES.get(module, module)


def write_case(case_dir: Path, module: str) -> None:
    if case_dir.exists():
        shutil.rmtree(case_dir)
    case_dir.mkdir(parents=True, exist_ok=True)

    probe_type = probe_type_for(module)
    main_hx = (
        f"import {module};\n\n"
        "class Main {\n"
        f"  static var __probe:{probe_type} = null;\n\n"
        "  static function main() {\n"
        "    Sys.println(__probe);\n"
        "  }\n"
        "}\n"
    )
    compile_hxml = (
        "-cp .\n"
        "-lib reflaxe.go\n"
        "-D go_output=out\n"
        "-D reflaxe_go_strict_examples\n"
        "-D reflaxe.dont_output_metadata_id\n"
        "-D no-traces\n"
        "-D no_traces\n"
        "-main Main\n"
    )

    (case_dir / "Main.hx").write_text(main_hx, encoding="utf-8")
    (case_dir / "compile.hxml").write_text(compile_hxml, encoding="utf-8")


def run_module(module: str, args: argparse.Namespace) -> SweepResult:
    started = time.monotonic()
    case_dir = WORK_ROOT / safe_name(module)

    try:
        write_case(case_dir, module)

        compile_proc = run_command(["haxe", "compile.hxml"], case_dir, args.timeout)
        if compile_proc.returncode != 0:
            return SweepResult(
                module=module,
                ok=False,
                stage="compile",
                duration_s=time.monotonic() - started,
                message=command_output(compile_proc),
            )

        if args.go_test:
            out_dir = case_dir / "out"
            go_proc = subprocess.run(
                ["go", "test", "./..."],
                cwd=out_dir,
                capture_output=True,
                text=True,
                timeout=args.timeout,
            )
            if go_proc.returncode != 0:
                return SweepResult(
                    module=module,
                    ok=False,
                    stage="go test",
                    duration_s=time.monotonic() - started,
                    message=command_output(go_proc),
                )

        return SweepResult(
            module=module,
            ok=True,
            stage="done",
            duration_s=time.monotonic() - started,
            message="ok",
        )
    except subprocess.TimeoutExpired as exc:
        return SweepResult(
            module=module,
            ok=False,
            stage="timeout",
            duration_s=time.monotonic() - started,
            message=f"command timed out after {args.timeout}s: {exc.cmd}",
        )


def write_last_run(results: list[SweepResult]) -> None:
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

    modules_file = Path(args.modules_file)
    all_modules = load_modules(modules_file)
    selected = select_modules(all_modules, args)

    if args.list:
        for module in selected:
            print(module)
        return 0

    if not selected:
        print("No modules selected")
        return 0

    results: list[SweepResult] = []
    for module in selected:
        print(f"==> {module}")
        result = run_module(module, args)
        results.append(result)
        status = "PASS" if result.ok else "FAIL"
        print(f"[{status}] {module} ({result.stage}, {result.duration_s:.2f}s)")
        if not result.ok and result.message:
            print(result.message)

    write_last_run(results)

    passed = sum(1 for result in results if result.ok)
    failed = len(results) - passed
    print(f"\nSummary: {passed} passed, {failed} failed, {len(results)} total")
    print(f"Last run report: {LAST_RUN}")

    if args.strict and failed > 0:
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
