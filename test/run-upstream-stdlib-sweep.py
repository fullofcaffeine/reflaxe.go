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
DEFAULT_EXPECTED_MISSING_FILE = ROOT / "test" / "upstream_std_expected_missing.json"
LAST_RUN = CACHE_ROOT / "upstream_std_sweep_last.json"

PROBE_TYPE_OVERRIDES = {
    "haxe.ds.IntMap": "haxe.ds.IntMap<Dynamic>",
    "haxe.ds.Map": "haxe.ds.Map<Dynamic, Dynamic>",
    "haxe.ds.StringMap": "haxe.ds.StringMap<Dynamic>",
    "haxe.ds.ObjectMap": "haxe.ds.ObjectMap<Dynamic, Dynamic>",
    "haxe.ds.EnumValueMap": "haxe.ds.EnumValueMap<EnumValue, Dynamic>",
    "haxe.ds.List": "haxe.ds.List<Dynamic>",
    "haxe.ds.BalancedTree": "haxe.ds.BalancedTree<Dynamic, Dynamic>",
    "haxe.ds.Option": "haxe.ds.Option<Dynamic>",
    "haxe.ds.ReadOnlyArray": "haxe.ds.ReadOnlyArray<Dynamic>",
    "haxe.ds.Vector": "haxe.ds.Vector<Dynamic>",
}

PROBE_VALUE_OVERRIDES = {
    "haxe.Int32": "0",
}


@dataclasses.dataclass
class ExpectedMissingRule:
    module: str
    reason: str
    failure_regex: str
    min_inclusive: tuple[int, int, int] | None = None
    max_inclusive: tuple[int, int, int] | None = None
    max_exclusive: tuple[int, int, int] | None = None


@dataclasses.dataclass
class SweepResult:
    module: str
    ok: bool
    outcome: str
    stage: str
    duration_s: float
    message: str
    policy_note: str = ""


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run upstream stdlib module sweep against reflaxe.go")
    parser.add_argument("--modules-file", default=str(DEFAULT_MODULES_FILE), help="Path to module list file")
    parser.add_argument("--module", action="append", default=[], help="Run only specific module(s)")
    parser.add_argument("--pattern", default="", help="Regex filter for module names")
    parser.add_argument("--list", action="store_true", help="List selected modules and exit")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    parser.add_argument("--go-test", action="store_true", help="Also run `go test ./...` in generated out/")
    parser.add_argument(
        "--expected-missing-file",
        default=str(DEFAULT_EXPECTED_MISSING_FILE),
        help="Path to expected-missing policy file (JSON)",
    )
    parser.add_argument("--no-expected-missing-policy", action="store_true", help="Disable expected-missing policy")
    parser.add_argument("--haxe-version", default="", help="Override detected Haxe version for policy matching")
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


def parse_version(raw: str) -> tuple[int, int, int] | None:
    match = re.search(r"(\d+)(?:\.(\d+))?(?:\.(\d+))?", raw.strip())
    if not match:
        return None
    return (
        int(match.group(1)),
        int(match.group(2) or 0),
        int(match.group(3) or 0),
    )


def detect_haxe_version(args: argparse.Namespace) -> tuple[str, tuple[int, int, int]]:
    if args.haxe_version.strip():
        raw = args.haxe_version.strip()
    else:
        proc = subprocess.run(
            ["haxe", "--version"],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=args.timeout,
            env={**os.environ, "HAXE_NO_SERVER": "1"},
        )
        if proc.returncode != 0:
            raise SystemExit(f"Failed to detect Haxe version:\n{command_output(proc)}")
        raw_output = (proc.stdout or proc.stderr or "").strip()
        raw = raw_output.splitlines()[-1] if raw_output else ""

    parsed = parse_version(raw)
    if parsed is None:
        raise SystemExit(f"Unable to parse Haxe version: {raw!r}")
    return raw, parsed


def parse_rule_version(value: object, *, field_name: str, module: str) -> tuple[int, int, int] | None:
    if value is None:
        return None
    parsed = parse_version(str(value))
    if parsed is None:
        raise SystemExit(
            f"Invalid expected-missing policy version {value!r} for {module} ({field_name})"
        )
    return parsed


def load_expected_missing_rules(path: Path) -> list[ExpectedMissingRule]:
    if not path.exists():
        return []

    try:
        payload = json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        raise SystemExit(f"Invalid JSON in expected-missing policy file {path}: {exc}") from exc

    raw_rules = payload.get("rules")
    if not isinstance(raw_rules, list):
        raise SystemExit(f"Expected 'rules' list in expected-missing policy file: {path}")

    rules: list[ExpectedMissingRule] = []
    seen_modules: set[str] = set()
    for index, raw_rule in enumerate(raw_rules):
        if not isinstance(raw_rule, dict):
            raise SystemExit(f"Expected object at expected-missing policy rules[{index}]")

        module = raw_rule.get("module", "")
        if not isinstance(module, str) or not module.strip():
            raise SystemExit(f"Missing/invalid module name at expected-missing policy rules[{index}]")
        module = module.strip()
        if module in seen_modules:
            raise SystemExit(f"Duplicate expected-missing policy module entry: {module}")
        seen_modules.add(module)

        reason = raw_rule.get("reason", "")
        if not isinstance(reason, str):
            raise SystemExit(f"Invalid reason for expected-missing policy module {module}")

        failure_regex = raw_rule.get("failure_regex", "")
        if not isinstance(failure_regex, str) or not failure_regex.strip():
            failure_regex = rf"Type not found\s*:\s*{re.escape(module)}"
        try:
            re.compile(failure_regex)
        except re.error as exc:
            raise SystemExit(f"Invalid failure_regex for expected-missing policy module {module}: {exc}") from exc

        raw_version = raw_rule.get("haxe_version", {})
        if raw_version is None:
            raw_version = {}
        if not isinstance(raw_version, dict):
            raise SystemExit(f"Invalid haxe_version object for expected-missing policy module {module}")

        min_inclusive = parse_rule_version(raw_version.get("min_inclusive"), field_name="min_inclusive", module=module)
        max_inclusive = parse_rule_version(raw_version.get("max_inclusive"), field_name="max_inclusive", module=module)
        max_exclusive = parse_rule_version(raw_version.get("max_exclusive"), field_name="max_exclusive", module=module)

        if max_inclusive is not None and max_exclusive is not None:
            raise SystemExit(
                f"Expected-missing policy module {module} cannot set both max_inclusive and max_exclusive"
            )

        rules.append(
            ExpectedMissingRule(
                module=module,
                reason=reason.strip(),
                failure_regex=failure_regex,
                min_inclusive=min_inclusive,
                max_inclusive=max_inclusive,
                max_exclusive=max_exclusive,
            )
        )

    return rules


def version_in_range(version: tuple[int, int, int], rule: ExpectedMissingRule) -> bool:
    if rule.min_inclusive is not None and version < rule.min_inclusive:
        return False
    if rule.max_inclusive is not None and version > rule.max_inclusive:
        return False
    if rule.max_exclusive is not None and version >= rule.max_exclusive:
        return False
    return True


def active_expected_missing_rules(
    rules: list[ExpectedMissingRule],
    haxe_version: tuple[int, int, int],
) -> dict[str, ExpectedMissingRule]:
    active: dict[str, ExpectedMissingRule] = {}
    for rule in rules:
        if version_in_range(haxe_version, rule):
            active[rule.module] = rule
    return active


def is_expected_missing_failure(rule: ExpectedMissingRule, message: str) -> bool:
    return re.search(rule.failure_regex, message, flags=re.MULTILINE) is not None


def run_command(cmd: list[str], cwd: Path, timeout_s: int) -> subprocess.CompletedProcess[str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s, env=env)


def safe_name(module: str) -> str:
    return re.sub(r"[^A-Za-z0-9_.-]", "_", module)


def probe_type_for(module: str) -> str:
    return PROBE_TYPE_OVERRIDES.get(module, module)


def probe_value_for(module: str) -> str:
    return PROBE_VALUE_OVERRIDES.get(module, "null")


def write_case(case_dir: Path, module: str) -> None:
    if case_dir.exists():
        shutil.rmtree(case_dir)
    case_dir.mkdir(parents=True, exist_ok=True)

    probe_type = probe_type_for(module)
    probe_value = probe_value_for(module)
    main_hx = (
        f"import {module};\n\n"
        "class Main {\n"
        f"  static var __probe:{probe_type} = {probe_value};\n\n"
        "  static function main() {\n"
        "    Sys.println(__probe);\n"
        "  }\n"
        "}\n"
    )
    compile_hxml = (
        "-cp .\n"
        "-lib reflaxe.go\n"
        "-D go_output=out\n"
        "-D go_no_build\n"
        "-D reflaxe_go_strict_examples\n"
        "-D reflaxe.dont_output_metadata_id\n"
        "-D no-traces\n"
        "-D no_traces\n"
        "-main Main\n"
    )

    (case_dir / "Main.hx").write_text(main_hx, encoding="utf-8")
    (case_dir / "compile.hxml").write_text(compile_hxml, encoding="utf-8")


def run_module(
    module: str,
    args: argparse.Namespace,
    expected_missing_rule: ExpectedMissingRule | None,
) -> SweepResult:
    started = time.monotonic()
    case_dir = WORK_ROOT / safe_name(module)

    try:
        write_case(case_dir, module)

        compile_proc = run_command(["haxe", "compile.hxml"], case_dir, args.timeout)
        if compile_proc.returncode != 0:
            output = command_output(compile_proc)
            if expected_missing_rule is not None and is_expected_missing_failure(expected_missing_rule, output):
                return SweepResult(
                    module=module,
                    ok=True,
                    outcome="expected_missing",
                    stage="compile",
                    duration_s=time.monotonic() - started,
                    message=output,
                    policy_note=expected_missing_rule.reason,
                )
            return SweepResult(
                module=module,
                ok=False,
                outcome="fail",
                stage="compile",
                duration_s=time.monotonic() - started,
                message=output,
            )

        if expected_missing_rule is not None:
            policy_note = expected_missing_rule.reason
            if policy_note:
                policy_note = f"{policy_note} "
            policy_note += f"(policy expected missing for {module})"
            return SweepResult(
                module=module,
                ok=False,
                outcome="unexpected_present",
                stage="compile",
                duration_s=time.monotonic() - started,
                message=f"Module compiled successfully but policy marks it expected-missing: {module}",
                policy_note=policy_note,
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
                    outcome="fail",
                    stage="go test",
                    duration_s=time.monotonic() - started,
                    message=command_output(go_proc),
                )

        return SweepResult(
            module=module,
            ok=True,
            outcome="pass",
            stage="done",
            duration_s=time.monotonic() - started,
            message="ok",
        )
    except subprocess.TimeoutExpired as exc:
        return SweepResult(
            module=module,
            ok=False,
            outcome="fail",
            stage="timeout",
            duration_s=time.monotonic() - started,
            message=f"command timed out after {args.timeout}s: {exc.cmd}",
        )


def result_counts(results: list[SweepResult]) -> dict[str, int]:
    passed = sum(1 for result in results if result.outcome == "pass")
    expected_missing = sum(1 for result in results if result.outcome == "expected_missing")
    failed = sum(1 for result in results if result.outcome == "fail")
    unexpected_present = sum(1 for result in results if result.outcome == "unexpected_present")
    return {
        "passed": passed,
        "expected_missing": expected_missing,
        "failed": failed,
        "unexpected_present": unexpected_present,
    }


def write_last_run(
    results: list[SweepResult],
    haxe_version_raw: str,
    expected_missing_policy_file: Path | None,
) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    counts = result_counts(results)
    payload = {
        "generated_at_epoch": int(time.time()),
        "haxe_version": haxe_version_raw,
        "expected_missing_policy_file": str(expected_missing_policy_file) if expected_missing_policy_file else None,
        "total": len(results),
        "passed": counts["passed"],
        "expected_missing": counts["expected_missing"],
        "failed": counts["failed"],
        "unexpected_present": counts["unexpected_present"],
        "results": [dataclasses.asdict(result) for result in results],
    }
    LAST_RUN.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()

    expected_missing_rules: dict[str, ExpectedMissingRule] = {}
    haxe_version_raw = ""
    expected_missing_policy_file: Path | None = None
    if not args.no_expected_missing_policy:
        expected_missing_policy_file = Path(args.expected_missing_file)
        all_rules = load_expected_missing_rules(expected_missing_policy_file)
        if all_rules:
            haxe_version_raw, haxe_version = detect_haxe_version(args)
            expected_missing_rules = active_expected_missing_rules(all_rules, haxe_version)
            print(
                "Expected-missing policy:",
                f"{len(expected_missing_rules)} active module rule(s)",
                f"for Haxe {haxe_version_raw}",
            )

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
        result = run_module(module, args, expected_missing_rules.get(module))
        results.append(result)
        status_by_outcome = {
            "pass": "PASS",
            "expected_missing": "EXPECTED_MISSING",
            "fail": "FAIL",
            "unexpected_present": "UNEXPECTED_PRESENT",
        }
        status = status_by_outcome.get(result.outcome, "FAIL")
        print(f"[{status}] {module} ({result.stage}, {result.duration_s:.2f}s)")
        if result.policy_note:
            print(f"policy: {result.policy_note}")
        if result.message and result.message != "ok":
            print(result.message)

    write_last_run(results, haxe_version_raw, expected_missing_policy_file)

    counts = result_counts(results)
    blocking_failures = counts["failed"] + counts["unexpected_present"]
    print(
        "\nSummary:",
        f"{counts['passed']} passed,",
        f"{counts['expected_missing']} expected missing,",
        f"{counts['failed']} failed,",
        f"{counts['unexpected_present']} unexpected present,",
        f"{len(results)} total",
    )
    print(f"Last run report: {LAST_RUN}")

    if args.strict and blocking_failures > 0:
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
