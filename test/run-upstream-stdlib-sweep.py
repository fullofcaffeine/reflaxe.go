#!/usr/bin/env python3

from __future__ import annotations

import argparse
import contextlib
import dataclasses
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import time

try:
    import fcntl  # type: ignore[attr-defined]
except ImportError:
    fcntl = None

ROOT = Path(__file__).resolve().parent.parent
CACHE_ROOT = ROOT / "test" / ".test-cache"
DEFAULT_WORK_ROOT = CACHE_ROOT / "upstream-stdlib-sweep"
DEFAULT_MODULES_FILE = ROOT / "test" / "upstream_std_modules.txt"
DEFAULT_EXPECTED_MISSING_FILE = ROOT / "test" / "upstream_std_expected_missing.json"
DEFAULT_EXPECTED_UNAVAILABLE_FILE = ROOT / "test" / "upstream_std_expected_unavailable.json"
LAST_RUN = CACHE_ROOT / "upstream_std_sweep_last.json"
RUN_LOCK = CACHE_ROOT / "upstream-stdlib-sweep.lock"

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
class PolicyRule:
    module: str
    reason: str
    failure_regex: str
    source: str
    stage: str = "compile"
    min_inclusive: tuple[int, int, int] | None = None
    max_inclusive: tuple[int, int, int] | None = None
    max_exclusive: tuple[int, int, int] | None = None


@dataclasses.dataclass
class ProbeSpec:
    mode: str
    macro_expr: str | None
    probe_type: str | None
    probe_value: str | None


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
    parser.add_argument(
        "--expected-unavailable-file",
        default=str(DEFAULT_EXPECTED_UNAVAILABLE_FILE),
        help="Path to expected-unavailable policy file (JSON)",
    )
    parser.add_argument("--no-expected-missing-policy", action="store_true", help="Disable expected-missing policy")
    parser.add_argument(
        "--no-expected-unavailable-policy",
        action="store_true",
        help="Disable expected-unavailable policy",
    )
    parser.add_argument("--haxe-version", default="", help="Override detected Haxe version for policy matching")
    parser.add_argument("--strict", action="store_true", help="Exit non-zero if any module fails")
    parser.add_argument(
        "--work-root",
        default=str(DEFAULT_WORK_ROOT),
        help="Per-run work root parent (a unique run-id folder is created beneath it)",
    )
    parser.add_argument("--run-id", default="", help="Optional deterministic run-id subdirectory")
    parser.add_argument("--keep-work", action="store_true", help="Keep per-run work directory for debugging")
    parser.add_argument("--lock-timeout", type=int, default=30, help="Seconds to wait for sweep lock (0=fail fast)")
    parser.add_argument(
        "--probe-mode",
        default="include",
        choices=["include", "typed"],
        help="Probe strategy: include module macro (default) or typed var probe",
    )
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
                        f"Another stdlib sweep is active (lock: {RUN_LOCK}). "
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

    start = time.monotonic()
    while True:
        try:
            fd = os.open(str(RUN_LOCK), os.O_CREAT | os.O_EXCL | os.O_WRONLY)
            break
        except FileExistsError:
            if timeout_s <= 0 or (time.monotonic() - start) >= timeout_s:
                raise SystemExit(
                    f"Another stdlib sweep is active (lock: {RUN_LOCK}). "
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


def load_modules(path: Path) -> list[str]:
    if not path.exists():
        raise SystemExit(f"Modules file not found: {path}")

    modules: list[str] = []
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        modules.append(line)

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


def parse_rule_version(value: object, *, field_name: str, module: str, source: str) -> tuple[int, int, int] | None:
    if value is None:
        return None
    parsed = parse_version(str(value))
    if parsed is None:
        raise SystemExit(
            f"Invalid policy version {value!r} for {module} ({field_name}) in {source}"
        )
    return parsed


def load_policy_rules(path: Path, *, source_label: str) -> list[PolicyRule]:
    if not path.exists():
        return []

    try:
        payload = json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        raise SystemExit(f"Invalid JSON in {source_label} policy file {path}: {exc}") from exc

    raw_rules = payload.get("rules")
    if not isinstance(raw_rules, list):
        raise SystemExit(f"Expected 'rules' list in {source_label} policy file: {path}")

    rules: list[PolicyRule] = []
    seen_modules: set[str] = set()
    for index, raw_rule in enumerate(raw_rules):
        if not isinstance(raw_rule, dict):
            raise SystemExit(f"Expected object at {source_label} policy rules[{index}]")

        module = raw_rule.get("module", "")
        if not isinstance(module, str) or not module.strip():
            raise SystemExit(f"Missing/invalid module name at {source_label} policy rules[{index}]")
        module = module.strip()
        if module in seen_modules:
            raise SystemExit(f"Duplicate {source_label} policy module entry: {module}")
        seen_modules.add(module)

        reason = raw_rule.get("reason", "")
        if not isinstance(reason, str):
            raise SystemExit(f"Invalid reason for {source_label} policy module {module}")

        stage = raw_rule.get("stage", "compile")
        if not isinstance(stage, str):
            raise SystemExit(f"Invalid stage for {source_label} policy module {module}")
        stage = stage.strip()
        if stage not in {"compile", "go_test", "any"}:
            raise SystemExit(
                f"Invalid stage {stage!r} for {source_label} policy module {module}; "
                "expected one of: compile, go_test, any"
            )

        failure_regex = raw_rule.get("failure_regex", "")
        if not isinstance(failure_regex, str) or not failure_regex.strip():
            failure_regex = rf"Type not found\s*:\s*{re.escape(module)}"
        try:
            re.compile(failure_regex)
        except re.error as exc:
            raise SystemExit(
                f"Invalid failure_regex for {source_label} policy module {module}: {exc}"
            ) from exc

        raw_version = raw_rule.get("haxe_version", {})
        if raw_version is None:
            raw_version = {}
        if not isinstance(raw_version, dict):
            raise SystemExit(f"Invalid haxe_version object for {source_label} policy module {module}")

        min_inclusive = parse_rule_version(
            raw_version.get("min_inclusive"), field_name="min_inclusive", module=module, source=source_label
        )
        max_inclusive = parse_rule_version(
            raw_version.get("max_inclusive"), field_name="max_inclusive", module=module, source=source_label
        )
        max_exclusive = parse_rule_version(
            raw_version.get("max_exclusive"), field_name="max_exclusive", module=module, source=source_label
        )

        if max_inclusive is not None and max_exclusive is not None:
            raise SystemExit(
                f"Policy module {module} cannot set both max_inclusive and max_exclusive in {source_label}"
            )

        rules.append(
            PolicyRule(
                module=module,
                reason=reason.strip(),
                failure_regex=failure_regex,
                source=source_label,
                stage=stage,
                min_inclusive=min_inclusive,
                max_inclusive=max_inclusive,
                max_exclusive=max_exclusive,
            )
        )

    return rules


def merge_policy_rules(*rule_sets: list[PolicyRule]) -> list[PolicyRule]:
    merged: list[PolicyRule] = []
    seen: dict[str, str] = {}
    for rules in rule_sets:
        for rule in rules:
            if rule.module in seen:
                raise SystemExit(
                    f"Duplicate module policy for {rule.module}: seen in {seen[rule.module]} and {rule.source}"
                )
            seen[rule.module] = rule.source
            merged.append(rule)
    return merged


def version_in_range(version: tuple[int, int, int], rule: PolicyRule) -> bool:
    if rule.min_inclusive is not None and version < rule.min_inclusive:
        return False
    if rule.max_inclusive is not None and version > rule.max_inclusive:
        return False
    if rule.max_exclusive is not None and version >= rule.max_exclusive:
        return False
    return True


def active_policy_rules(
    rules: list[PolicyRule],
    haxe_version: tuple[int, int, int],
) -> dict[str, PolicyRule]:
    active: dict[str, PolicyRule] = {}
    for rule in rules:
        if version_in_range(haxe_version, rule):
            active[rule.module] = rule
    return active


def is_expected_policy_failure(rule: PolicyRule, message: str) -> bool:
    return re.search(rule.failure_regex, message, flags=re.MULTILINE) is not None


def rule_applies_to_stage(rule: PolicyRule, stage: str) -> bool:
    if rule.stage == "any":
        return True
    return rule.stage == stage


def rule_observable_in_run(rule: PolicyRule, args: argparse.Namespace) -> bool:
    if rule.stage == "any":
        return True
    if rule.stage == "compile":
        return True
    if rule.stage == "go_test":
        return args.go_test
    return False


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


def probe_spec_for(module: str, args: argparse.Namespace) -> ProbeSpec:
    if args.probe_mode == "typed":
        return ProbeSpec(
            mode="typed",
            macro_expr=None,
            probe_type=probe_type_for(module),
            probe_value=probe_value_for(module),
        )

    if module == "haxe.Json":
        # On some case-insensitive filesystems include('haxe.Json') can collide with haxe/json paths.
        return ProbeSpec(
            mode="macro_get_type",
            macro_expr="haxe.macro.Context.getType('haxe.Json')",
            probe_type=None,
            probe_value=None,
        )

    return ProbeSpec(
        mode="macro_include",
        macro_expr=f"include('{module}')",
        probe_type=None,
        probe_value=None,
    )


def compile_case(case_dir: Path, args: argparse.Namespace, probe_spec: ProbeSpec) -> subprocess.CompletedProcess[str]:
    compile_cmd = ["haxe", "compile.hxml"]
    if probe_spec.macro_expr:
        compile_cmd.extend(["--macro", probe_spec.macro_expr])
    return run_command(compile_cmd, case_dir, args.timeout)


def should_retry_with_typed_probe(probe_spec: ProbeSpec, compile_output: str) -> bool:
    if probe_spec.mode == "typed":
        return False
    return "Invalid commandline class" in compile_output


def write_case(case_dir: Path, module: str, probe_spec: ProbeSpec) -> None:
    if case_dir.exists():
        shutil.rmtree(case_dir)
    case_dir.mkdir(parents=True, exist_ok=True)

    if probe_spec.mode == "typed":
        assert probe_spec.probe_type is not None
        assert probe_spec.probe_value is not None
        main_hx = (
            f"import {module};\n\n"
            "class Main {\n"
            f"  static var __probe:{probe_spec.probe_type} = {probe_spec.probe_value};\n\n"
            "  static function main() {\n"
            "    Sys.println(__probe);\n"
            "  }\n"
            "}\n"
        )
    else:
        main_hx = (
            "class Main {\n"
            "  static function main() {\n"
            "    Sys.println(\"ok\");\n"
            "  }\n"
            "}\n"
        )

    compiler_src = (ROOT / "src").as_posix()
    reflaxe_src = (ROOT / "vendor" / "reflaxe" / "src").as_posix()
    compile_hxml = (
        "-cp .\n"
        f"-cp {compiler_src}\n"
        f"-cp {reflaxe_src}\n"
        "--macro reflaxe.go.CompilerBootstrap.Start()\n"
        "--macro reflaxe.go.CompilerInit.Start()\n"
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


def resolve_run_id(raw: str) -> str:
    candidate = raw.strip()
    if candidate == "":
        candidate = f"run-{int(time.time())}-{os.getpid()}"
    safe = re.sub(r"[^A-Za-z0-9_.-]", "_", candidate)
    if safe == "":
        safe = f"run-{int(time.time())}-{os.getpid()}"
    return safe


def prepare_work_root(args: argparse.Namespace) -> Path:
    parent = Path(args.work_root)
    run_id = resolve_run_id(args.run_id)
    work_root = parent / run_id
    if work_root.exists():
        shutil.rmtree(work_root)
    work_root.mkdir(parents=True, exist_ok=True)
    return work_root


def run_module(
    module: str,
    args: argparse.Namespace,
    expected_rule: PolicyRule | None,
    work_root: Path,
) -> SweepResult:
    started = time.monotonic()
    case_dir = work_root / safe_name(module)
    probe_spec = probe_spec_for(module, args)

    try:
        write_case(case_dir, module, probe_spec)

        compile_proc = compile_case(case_dir, args, probe_spec)
        compile_output = command_output(compile_proc)
        if compile_proc.returncode != 0 and should_retry_with_typed_probe(probe_spec, compile_output):
            fallback_probe = ProbeSpec(
                mode="typed",
                macro_expr=None,
                probe_type=probe_type_for(module),
                probe_value=probe_value_for(module),
            )
            write_case(case_dir, module, fallback_probe)
            fallback_proc = compile_case(case_dir, args, fallback_probe)
            fallback_output = command_output(fallback_proc)
            if fallback_proc.returncode == 0:
                compile_proc = fallback_proc
                compile_output = fallback_output
            else:
                compile_proc = fallback_proc
                compile_output = (
                    f"{compile_output}\n\n[typed probe fallback]\n{fallback_output}".strip()
                )
        if compile_proc.returncode != 0:
            if (
                expected_rule is not None
                and rule_applies_to_stage(expected_rule, "compile")
                and is_expected_policy_failure(expected_rule, compile_output)
            ):
                return SweepResult(
                    module=module,
                    ok=True,
                    outcome="expected_missing",
                    stage="compile",
                    duration_s=time.monotonic() - started,
                    message=compile_output,
                    policy_note=f"{expected_rule.source}: {expected_rule.reason}",
                )
            return SweepResult(
                module=module,
                ok=False,
                outcome="fail",
                stage="compile",
                duration_s=time.monotonic() - started,
                message=compile_output,
            )

        if args.go_test:
            out_dir = case_dir / "out"
            if not out_dir.exists():
                return SweepResult(
                    module=module,
                    ok=False,
                    outcome="fail",
                    stage="go test",
                    duration_s=time.monotonic() - started,
                    message=f"Missing generated output directory: {out_dir}",
                )
            go_proc = subprocess.run(
                ["go", "test", "./..."],
                cwd=out_dir,
                capture_output=True,
                text=True,
                timeout=args.timeout,
            )
            if go_proc.returncode != 0:
                go_output = command_output(go_proc)
                if (
                    expected_rule is not None
                    and rule_applies_to_stage(expected_rule, "go_test")
                    and is_expected_policy_failure(expected_rule, go_output)
                ):
                    return SweepResult(
                        module=module,
                        ok=True,
                        outcome="expected_missing",
                        stage="go test",
                        duration_s=time.monotonic() - started,
                        message=go_output,
                        policy_note=f"{expected_rule.source}: {expected_rule.reason}",
                    )
                return SweepResult(
                    module=module,
                    ok=False,
                    outcome="fail",
                    stage="go test",
                    duration_s=time.monotonic() - started,
                    message=go_output,
                )

        if expected_rule is not None and rule_observable_in_run(expected_rule, args):
            policy_note = expected_rule.reason
            if policy_note:
                policy_note = f"{expected_rule.source}: {policy_note} "
            policy_note += f"(policy expected failure for {module}; stage={expected_rule.stage})"
            return SweepResult(
                module=module,
                ok=False,
                outcome="unexpected_present",
                stage="done",
                duration_s=time.monotonic() - started,
                message=f"Module passed but policy marks it expected-failure: {module}",
                policy_note=policy_note,
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
    expected_unavailable_policy_file: Path | None,
    work_root: Path,
) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    counts = result_counts(results)
    payload = {
        "generated_at_epoch": int(time.time()),
        "haxe_version": haxe_version_raw,
        "expected_missing_policy_file": str(expected_missing_policy_file) if expected_missing_policy_file else None,
        "expected_unavailable_policy_file": (
            str(expected_unavailable_policy_file) if expected_unavailable_policy_file else None
        ),
        "work_root": str(work_root),
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

    policy_rules: dict[str, PolicyRule] = {}
    haxe_version_raw = ""
    expected_missing_policy_file: Path | None = None
    expected_unavailable_policy_file: Path | None = None

    missing_rules: list[PolicyRule] = []
    unavailable_rules: list[PolicyRule] = []

    if not args.no_expected_missing_policy:
        expected_missing_policy_file = Path(args.expected_missing_file)
        missing_rules = load_policy_rules(expected_missing_policy_file, source_label="expected-missing")

    if not args.no_expected_unavailable_policy:
        expected_unavailable_policy_file = Path(args.expected_unavailable_file)
        unavailable_rules = load_policy_rules(expected_unavailable_policy_file, source_label="expected-unavailable")

    all_rules = merge_policy_rules(missing_rules, unavailable_rules)
    if all_rules:
        haxe_version_raw, haxe_version = detect_haxe_version(args)
        policy_rules = active_policy_rules(all_rules, haxe_version)
        print(
            "Expected-failure policy:",
            f"{len(policy_rules)} active module rule(s)",
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

    with acquire_run_lock(args.lock_timeout):
        work_root = prepare_work_root(args)

        results: list[SweepResult] = []
        for module in selected:
            print(f"==> {module}")
            result = run_module(module, args, policy_rules.get(module), work_root)
            results.append(result)
            status_by_outcome = {
                "pass": "PASS",
                "expected_missing": "EXPECTED_POLICY",
                "fail": "FAIL",
                "unexpected_present": "UNEXPECTED_PRESENT",
            }
            status = status_by_outcome.get(result.outcome, "FAIL")
            print(f"[{status}] {module} ({result.stage}, {result.duration_s:.2f}s)")
            if result.policy_note:
                print(f"policy: {result.policy_note}")
            if result.message and result.message != "ok":
                print(result.message)

        write_last_run(
            results,
            haxe_version_raw,
            expected_missing_policy_file,
            expected_unavailable_policy_file,
            work_root,
        )

        counts = result_counts(results)
        blocking_failures = counts["failed"] + counts["unexpected_present"]
        print(
            "\nSummary:",
            f"{counts['passed']} passed,",
            f"{counts['expected_missing']} expected policy,",
            f"{counts['failed']} failed,",
            f"{counts['unexpected_present']} unexpected present,",
            f"{len(results)} total",
        )
        print(f"Last run report: {LAST_RUN}")

        if not args.keep_work:
            shutil.rmtree(work_root, ignore_errors=True)

        if args.strict and blocking_failures > 0:
            return 1
        return 0


if __name__ == "__main__":
    raise SystemExit(main())
