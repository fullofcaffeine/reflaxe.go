#!/usr/bin/env python3

from __future__ import annotations

import argparse
import dataclasses
import os
from pathlib import Path
import shlex
import shutil
import subprocess
import time
import json

from git_changes import GitChangeDiscoveryError, collect_changed_paths

ROOT = Path(__file__).resolve().parent.parent
EXAMPLES_ROOT = ROOT / "examples"
PROFILES = ("portable", "metal")
EXCLUDE_NAMES = {"go.sum", "_GeneratedFiles.json", ".DS_Store"}
EXCLUDE_DIRS = {".cache"}
TELEMETRY_DIR = ROOT / ".cache" / "generated-output-telemetry"
QA_MANIFEST = EXAMPLES_ROOT / "qa-manifest.json"
COMPATIBILITY_SOURCE = ROOT / "docs" / "compatibility-support-source.json"


@dataclasses.dataclass(frozen=True)
class ExampleLaneMetadata:
    product_surfaces: tuple[str, ...]
    evidence_modes: tuple[str, ...]
    release_claim_bearing: bool
    compatibility_operations: tuple[str, ...]


@dataclasses.dataclass(frozen=True)
class ExampleMetadata:
    example_id: str
    tier: str
    claim_bearing: bool
    profiles: tuple[str, ...]
    test_command: str
    lanes: dict[str, ExampleLaneMetadata]


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
    metadata: ExampleMetadata

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
    telemetry: list[dict] = dataclasses.field(default_factory=list)


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


def load_example_metadata() -> dict[str, ExampleMetadata]:
    if not QA_MANIFEST.exists():
        raise RuntimeError(f"missing examples QA manifest: {QA_MANIFEST}")
    payload = json.loads(QA_MANIFEST.read_text(encoding="utf-8"))
    if payload.get("schemaVersion") != 2 or not isinstance(payload.get("examples"), list):
        raise RuntimeError(f"invalid examples QA manifest schema: {QA_MANIFEST}")

    if not COMPATIBILITY_SOURCE.exists():
        raise RuntimeError(f"missing compatibility authority: {COMPATIBILITY_SOURCE}")
    compatibility = json.loads(COMPATIBILITY_SOURCE.read_text(encoding="utf-8"))
    known_operations: set[str] = set()
    admitted_operations: set[str] = set()
    for surface in compatibility.get("surfaces", []):
        surface_id = str(surface.get("id", "")).strip()
        for operation in surface.get("operations", []):
            operation_id = str(operation.get("id", "")).strip()
            if not surface_id or not operation_id:
                raise RuntimeError(
                    f"invalid compatibility operation identity in {COMPATIBILITY_SOURCE}"
                )
            reference = f"{surface_id}/{operation_id}"
            if reference in known_operations:
                raise RuntimeError(f"duplicate compatibility operation: {reference}")
            known_operations.add(reference)
            if operation.get("release_admitted") is True:
                admitted_operations.add(reference)

    records: dict[str, ExampleMetadata] = {}
    for raw in payload["examples"]:
        example_id = str(raw.get("id", "")).strip()
        if not example_id or example_id in records:
            raise RuntimeError(f"invalid or duplicate example id in {QA_MANIFEST}: {example_id!r}")
        claim_bearing = raw.get("claimBearing")
        if not isinstance(claim_bearing, bool):
            raise RuntimeError(f"example {example_id} must declare claimBearing")
        execution = raw.get("execution", [])
        if claim_bearing and execution != [
            "haxe-custom-backend",
            "gofmt",
            "go-test",
            "go-run",
            "expected-output",
        ]:
            raise RuntimeError(f"claim-bearing example {example_id} does not declare the full executable chain")
        raw_lanes = raw.get("lanes")
        if not isinstance(raw_lanes, dict) or set(raw_lanes) != {"default", "ci"}:
            raise RuntimeError(
                f"example {example_id} must declare exactly the default and ci execution lanes"
            )
        lanes: dict[str, ExampleLaneMetadata] = {}
        for lane_id, raw_lane in raw_lanes.items():
            if not isinstance(raw_lane, dict):
                raise RuntimeError(f"example {example_id} lane {lane_id} must be an object")
            product_surfaces = tuple(str(item) for item in raw_lane.get("productSurfaces", []))
            evidence_modes = tuple(str(item) for item in raw_lane.get("evidenceModes", []))
            if not product_surfaces or not evidence_modes:
                raise RuntimeError(
                    f"example {example_id} lane {lane_id} must declare surfaces and evidence modes"
                )
            release_claim_bearing = raw_lane.get("releaseClaimBearing")
            if not isinstance(release_claim_bearing, bool):
                raise RuntimeError(
                    f"example {example_id} lane {lane_id} must declare releaseClaimBearing"
                )
            raw_operations = raw_lane.get("compatibilityOperations")
            if not isinstance(raw_operations, list) or not all(
                isinstance(item, str) and item.strip() for item in raw_operations
            ):
                raise RuntimeError(
                    f"example {example_id} lane {lane_id} must declare compatibilityOperations as a string list"
                )
            compatibility_operations = tuple(item.strip() for item in raw_operations)
            if len(set(compatibility_operations)) != len(compatibility_operations):
                raise RuntimeError(
                    f"example {example_id} lane {lane_id} repeats a compatibility operation"
                )
            if release_claim_bearing:
                if not claim_bearing:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} must belong to a claim-bearing example"
                    )
                if "portable-compiler" not in product_surfaces:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} must own portable-compiler evidence"
                    )
                if "go-native-metal" in product_surfaces or "native-metal" in evidence_modes:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} cannot borrow native/metal evidence"
                    )
                if not compatibility_operations:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} has no compatibility operations"
                    )
                unknown = sorted(set(compatibility_operations) - known_operations)
                if unknown:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} names unknown compatibility operations: {unknown}"
                    )
                excluded = sorted(set(compatibility_operations) - admitted_operations)
                if excluded:
                    raise RuntimeError(
                        f"release-bearing example {example_id} lane {lane_id} names operations that are not release-admitted: {excluded}"
                    )
            elif compatibility_operations:
                raise RuntimeError(
                    f"non-release-bearing example {example_id} lane {lane_id} cannot publish compatibility operations"
                )
            lanes[lane_id] = ExampleLaneMetadata(
                product_surfaces=product_surfaces,
                evidence_modes=evidence_modes,
                release_claim_bearing=release_claim_bearing,
                compatibility_operations=compatibility_operations,
            )
        profiles = tuple(str(item) for item in raw.get("profiles", []))
        if any(lane.release_claim_bearing for lane in lanes.values()) and profiles != ("portable",):
            raise RuntimeError(
                f"release-bearing example {example_id} must be portable-only"
            )
        records[example_id] = ExampleMetadata(
            example_id=example_id,
            tier=str(raw.get("tier", "")),
            claim_bearing=claim_bearing,
            profiles=profiles,
            test_command=str(raw.get("testCommand", "")),
            lanes=lanes,
        )
    return records


def discover_cases() -> list[ExampleProfileCase]:
    cases: list[ExampleProfileCase] = []
    if not EXAMPLES_ROOT.exists():
        return cases

    metadata = load_example_metadata()
    maintained = {
        path.name
        for path in EXAMPLES_ROOT.iterdir()
        if path.is_dir() and not path.name.startswith(".") and (path / "README.md").exists()
    }
    if maintained != set(metadata):
        missing = sorted(maintained - set(metadata))
        stale = sorted(set(metadata) - maintained)
        raise RuntimeError(f"examples QA manifest drift: missing={missing}, stale={stale}")

    for example_dir in sorted(EXAMPLES_ROOT.iterdir()):
        if not example_dir.is_dir():
            continue
        example_metadata = metadata.get(example_dir.name)
        if example_metadata is None:
            continue
        discovered_profiles: set[str] = set()
        for profile in PROFILES:
            compile_hxml = example_dir / f"compile.{profile}.hxml"
            compile_ci_hxml = example_dir / f"compile.{profile}.ci.hxml"
            if not compile_hxml.exists() or not compile_ci_hxml.exists():
                continue
            discovered_profiles.add(profile)
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
                    metadata=example_metadata,
                )
            )
        if discovered_profiles != set(example_metadata.profiles):
            raise RuntimeError(
                f"example profile declaration drift for {example_dir.name}: "
                f"manifest={sorted(example_metadata.profiles)}, discovered={sorted(discovered_profiles)}"
            )
    return cases


def changed_examples() -> set[str]:
    all_examples = set(load_example_metadata())
    base = os.environ.get("TEST_PLAN_BASE_REF", "").strip()
    if not base and os.environ.get("GITHUB_BASE_REF", "").strip():
        base = f"origin/{os.environ['GITHUB_BASE_REF'].strip()}"
    try:
        changed_paths = collect_changed_paths(ROOT, ["examples"], base=base)
    except GitChangeDiscoveryError:
        # A focused command may cost more after a Git failure; it must never
        # silently cost us the example whose change could not be discovered.
        return all_examples

    out: set[str] = set()
    for raw_path in changed_paths:
        path = Path(raw_path)
        parts = path.parts
        if len(parts) >= 2 and parts[0] == "examples":
            candidate = parts[1]
            if candidate not in all_examples:
                # Root-level manifests and future shared example authorities
                # can affect every lane. Treat them conservatively instead of
                # mistaking their filename for an example identifier.
                return all_examples
            out.add(candidate)
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


def read_run_args(path: Path) -> list[str]:
    if not path.exists():
        return []
    raw = path.read_text(encoding="utf-8", errors="replace").strip()
    if raw == "":
        return []
    try:
        return shlex.split(raw)
    except ValueError as exc:
        raise RuntimeError(f"invalid run args file {path}: {exc}") from exc


def command_output(proc: subprocess.CompletedProcess[str]) -> str:
    chunks: list[str] = []
    if proc.stdout:
        chunks.append(proc.stdout.strip())
    if proc.stderr:
        chunks.append(proc.stderr.strip())
    return "\n".join(chunk for chunk in chunks if chunk)


def go_module_binary_names(root: Path) -> set[str]:
    go_mod = root / "go.mod"
    if not go_mod.exists():
        return set()

    try:
        lines = go_mod.read_text(encoding="utf-8", errors="replace").splitlines()
    except OSError:
        return set()

    for line in lines:
        trimmed = line.strip()
        if not trimmed.startswith("module "):
            continue
        parts = trimmed.split(maxsplit=1)
        if len(parts) != 2:
            continue
        module_path = parts[1].strip()
        if module_path == "":
            continue
        base = module_path.rstrip("/").split("/")[-1]
        if base == "":
            continue
        return {base, base + ".exe"}

    return set()


def all_files(root: Path, extra_exclude_names: set[str] | None = None) -> list[Path]:
    if not root.exists():
        return []
    files: list[Path] = []
    exclude_names = set(EXCLUDE_NAMES)
    if extra_exclude_names:
        exclude_names.update(extra_exclude_names)
    for path in sorted(root.rglob("*")):
        if path.is_dir():
            continue
        if path.name in exclude_names:
            continue
        if any(part in EXCLUDE_DIRS for part in path.parts):
            continue
        files.append(path)
    return files


def collect_tree_deltas(left: Path, right: Path) -> list[str]:
    dynamic_excludes = go_module_binary_names(left) | go_module_binary_names(right)
    left_files = {path.relative_to(left): path for path in all_files(left, dynamic_excludes)} if left.exists() else {}
    right_files = {path.relative_to(right): path for path in all_files(right, dynamic_excludes)} if right.exists() else {}

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
    for path in all_files(source, go_module_binary_names(source)):
        rel = path.relative_to(source)
        dest = target / rel
        dest.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(path, dest)


def clean_out_dirs(case: ExampleProfileCase) -> None:
    if case.out_dir.exists():
        shutil.rmtree(case.out_dir)
    if case.out_ci_dir.exists():
        shutil.rmtree(case.out_ci_dir)


def collect_output_telemetry(out_dir: Path) -> dict:
    go_files = sorted(out_dir.rglob("*.go")) if out_dir.exists() else []
    total_bytes = 0
    largest_path: Path | None = None
    largest_bytes = 0
    for path in go_files:
        try:
            size = path.stat().st_size
        except OSError:
            size = 0
        total_bytes += size
        if size > largest_bytes or largest_path is None:
            largest_path = path
            largest_bytes = size

    return {
        "outputDir": relpath(out_dir),
        "goFileCount": len(go_files),
        "totalGoBytes": total_bytes,
        "largestGoFile": relpath(largest_path) if largest_path is not None else None,
        "largestGoFileBytes": largest_bytes if largest_path is not None else 0,
        "goTestElapsedMs": None,
        "goTestOk": None,
        "goTestCommand": "go test ./...",
    }


def relpath(path: Path | None) -> str | None:
    if path is None:
        return None
    try:
        return path.resolve().relative_to(ROOT).as_posix()
    except ValueError:
        return path.as_posix()


def run_go_checks(out_dir: Path, timeout_s: int) -> tuple[bool, str, str, dict]:
    telemetry = collect_output_telemetry(out_dir)
    go_files = sorted(out_dir.rglob("*.go")) if out_dir.exists() else []
    if not go_files:
        return False, "go", "no generated .go files", telemetry

    gofmt_cmd = ["gofmt", "-w"] + [str(path) for path in go_files]
    gofmt_proc = run_command(gofmt_cmd, cwd=out_dir, timeout_s=timeout_s)
    if gofmt_proc.returncode != 0:
        return False, "gofmt", command_output(gofmt_proc), telemetry

    started = time.monotonic()
    gotest_proc = run_command(["go", "test", "./..."], cwd=out_dir, timeout_s=timeout_s)
    telemetry["goTestElapsedMs"] = round((time.monotonic() - started) * 1000.0, 3)
    telemetry["goTestOk"] = gotest_proc.returncode == 0
    if gotest_proc.returncode != 0:
        return False, "go test", command_output(gotest_proc), telemetry

    return True, "go", "", telemetry


def annotate_telemetry(
    case: ExampleProfileCase,
    lane: str,
    telemetry: dict,
    *,
    compile_only: bool,
) -> dict:
    lane_metadata = case.metadata.lanes[lane]
    out = dict(telemetry)
    out.update(
        {
            "caseId": case.case_id,
            "example": case.example,
            "profile": case.profile,
            "lane": lane,
            "tier": case.metadata.tier,
            "declaredClaimBearing": case.metadata.claim_bearing,
            "declaredProductSurfaces": list(lane_metadata.product_surfaces),
            "declaredEvidenceModes": list(lane_metadata.evidence_modes),
            "declaredReleaseClaimBearing": lane_metadata.release_claim_bearing,
            "declaredCompatibilityOperations": list(lane_metadata.compatibility_operations),
            "claimBearing": False,
            "productSurfaces": [],
            "evidenceModes": [],
            "releaseClaimBearing": False,
            "compatibilityOperations": [],
            "runtimeStatus": "skipped" if compile_only else "pending",
            "stdoutStatus": "skipped" if compile_only else "pending",
            "claimStatus": "diagnostic-only" if compile_only else "pending",
            "releaseClaimStatus": (
                "diagnostic-only"
                if compile_only and lane_metadata.release_claim_bearing
                else "pending" if lane_metadata.release_claim_bearing else "not-declared"
            ),
        }
    )
    if telemetry.get("goTestOk") is False:
        out["claimStatus"] = "go-test-failed"
    return out


def complete_lane_telemetry(entry: dict, *, runtime_ok: bool, stdout_ok: bool | None) -> None:
    """Publish claim metadata only after the real runtime and oracle both pass."""
    if not runtime_ok:
        entry["runtimeStatus"] = "failed"
        entry["stdoutStatus"] = "not-run"
        entry["claimStatus"] = "runtime-failed"
        if entry["declaredReleaseClaimBearing"]:
            entry["releaseClaimStatus"] = "runtime-failed"
    elif stdout_ok is not True:
        entry["runtimeStatus"] = "passed"
        entry["stdoutStatus"] = "failed" if stdout_ok is False else "not-run"
        entry["claimStatus"] = "stdout-failed" if stdout_ok is False else "incomplete"
        if entry["declaredReleaseClaimBearing"]:
            entry["releaseClaimStatus"] = entry["claimStatus"]
    else:
        entry["runtimeStatus"] = "passed"
        entry["stdoutStatus"] = "passed"
        entry["claimBearing"] = bool(entry["declaredClaimBearing"])
        entry["productSurfaces"] = list(entry["declaredProductSurfaces"])
        entry["evidenceModes"] = list(entry["declaredEvidenceModes"])
        entry["claimStatus"] = "supported" if entry["claimBearing"] else "non-claim-bearing-pass"
        if entry["declaredReleaseClaimBearing"]:
            entry["releaseClaimBearing"] = True
            entry["compatibilityOperations"] = list(entry["declaredCompatibilityOperations"])
            entry["releaseClaimStatus"] = "supported"


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
    telemetry_entries: list[dict] = []
    try:
        run_ci_args = read_run_args(case.example_dir / "run.ci.args")
        run_args = read_run_args(case.example_dir / "run.args")
        clean_out_dirs(case)

        compile_ci_proc = run_command(
            ["haxe", case.compile_ci_hxml.name, "-D", "go_no_build"],
            cwd=case.example_dir,
            timeout_s=args.timeout,
            env={"HAXE_NO_SERVER": "1"},
        )
        if compile_ci_proc.returncode != 0:
            return CaseResult(case.case_id, False, "compile_ci", command_output(compile_ci_proc), time.monotonic() - started, telemetry_entries)

        ok, stage, msg, telemetry = run_go_checks(case.out_ci_dir, args.timeout)
        ci_telemetry = annotate_telemetry(case, "ci", telemetry, compile_only=args.compile_only)
        telemetry_entries.append(ci_telemetry)
        if not ok:
            return CaseResult(case.case_id, False, stage + "_ci", msg, time.monotonic() - started, telemetry_entries)

        if not args.compile_only:
            run_ci_proc = run_command(["go", "run", ".", *run_ci_args], cwd=case.out_ci_dir, timeout_s=args.timeout)
            if run_ci_proc.returncode != 0:
                complete_lane_telemetry(ci_telemetry, runtime_ok=False, stdout_ok=None)
                return CaseResult(case.case_id, False, "runtime_ci", command_output(run_ci_proc), time.monotonic() - started, telemetry_entries)
            ok_stdout, msg_stdout = compare_stdout(case.expected_ci_stdout, run_ci_proc.stdout)
            if not ok_stdout:
                complete_lane_telemetry(ci_telemetry, runtime_ok=True, stdout_ok=False)
                return CaseResult(case.case_id, False, "stdout_ci", msg_stdout, time.monotonic() - started, telemetry_entries)
            complete_lane_telemetry(ci_telemetry, runtime_ok=True, stdout_ok=True)

        compile_proc = run_command(
            ["haxe", case.compile_hxml.name, "-D", "go_no_build"],
            cwd=case.example_dir,
            timeout_s=args.timeout,
            env={"HAXE_NO_SERVER": "1"},
        )
        if compile_proc.returncode != 0:
            return CaseResult(case.case_id, False, "compile", command_output(compile_proc), time.monotonic() - started, telemetry_entries)

        ok, stage, msg, telemetry = run_go_checks(case.out_dir, args.timeout)
        default_telemetry = annotate_telemetry(case, "default", telemetry, compile_only=args.compile_only)
        telemetry_entries.append(default_telemetry)
        if not ok:
            return CaseResult(case.case_id, False, stage, msg, time.monotonic() - started, telemetry_entries)

        if not args.compile_only:
            run_proc = run_command(["go", "run", ".", *run_args], cwd=case.out_dir, timeout_s=args.timeout)
            if run_proc.returncode != 0:
                complete_lane_telemetry(default_telemetry, runtime_ok=False, stdout_ok=None)
                return CaseResult(case.case_id, False, "runtime", command_output(run_proc), time.monotonic() - started, telemetry_entries)
            ok_stdout, msg_stdout = compare_stdout(case.expected_stdout, run_proc.stdout)
            if not ok_stdout:
                complete_lane_telemetry(default_telemetry, runtime_ok=True, stdout_ok=False)
                return CaseResult(case.case_id, False, "stdout", msg_stdout, time.monotonic() - started, telemetry_entries)
            complete_lane_telemetry(default_telemetry, runtime_ok=True, stdout_ok=True)

        if args.bless_generated:
            copy_tree(case.out_dir, case.generated_dir)

        if not case.generated_dir.exists():
            return CaseResult(case.case_id, False, "generated", f"missing generated directory: {case.generated_dir} (use --bless-generated)", time.monotonic() - started, telemetry_entries)

        deltas = collect_tree_deltas(case.generated_dir, case.out_dir)
        if deltas:
            preview = "\n".join(deltas[:20])
            if len(deltas) > 20:
                preview += f"\n... and {len(deltas) - 20} more"
            return CaseResult(case.case_id, False, "generated", preview, time.monotonic() - started, telemetry_entries)

        return CaseResult(case.case_id, True, "done", "ok", time.monotonic() - started, telemetry_entries)

    except subprocess.TimeoutExpired as exc:
        for entry in telemetry_entries:
            if entry.get("claimStatus") == "pending":
                complete_lane_telemetry(entry, runtime_ok=False, stdout_ok=None)
                entry["runtimeStatus"] = "timeout"
                entry["claimStatus"] = "runtime-timeout"
        return CaseResult(case.case_id, False, "timeout", f"command timed out after {args.timeout}s: {exc.cmd}", time.monotonic() - started, telemetry_entries)
    except FileNotFoundError as exc:
        return CaseResult(case.case_id, False, "tool", f"missing tool: {exc}", time.monotonic() - started, telemetry_entries)
    except RuntimeError as exc:
        return CaseResult(case.case_id, False, "config", str(exc), time.monotonic() - started, telemetry_entries)


def build_telemetry_report(results: list[CaseResult]) -> dict:
    entries: list[dict] = []
    for result in results:
        for raw_entry in result.telemetry:
            entry = dict(raw_entry)
            entry["caseOk"] = result.ok
            entry["caseStage"] = result.stage
            if not result.ok:
                entry["claimBearing"] = False
                entry["productSurfaces"] = []
                entry["evidenceModes"] = []
                entry["releaseClaimBearing"] = False
                entry["compatibilityOperations"] = []
                entry["claimStatus"] = "case-failed"
                if entry["declaredReleaseClaimBearing"]:
                    entry["releaseClaimStatus"] = "case-failed"
            entries.append(entry)
    entries.sort(key=lambda item: (item["caseId"], item["lane"]))
    return {
        "schemaVersion": 1,
        "scope": "examples",
        "entryCount": len(entries),
        "entries": entries,
    }


def render_telemetry_markdown(report: dict) -> str:
    lines = [
        "# Generated Output Telemetry",
        "",
        f"- Scope: `{report['scope']}`",
        f"- Entry count: `{report['entryCount']}`",
        "",
        "| Case | Lane | Claim status | Runtime | Stdout | Go files | Total Go bytes | Largest Go file | Largest bytes | go test ms |",
        "| --- | --- | --- | --- | --- | ---: | ---: | --- | ---: | ---: |",
    ]
    for entry in report["entries"]:
        elapsed = entry["goTestElapsedMs"]
        lines.append(
            "| "
            + " | ".join(
                [
                    entry["caseId"],
                    entry["lane"],
                    entry.get("claimStatus", "legacy"),
                    entry.get("runtimeStatus", "unknown"),
                    entry.get("stdoutStatus", "unknown"),
                    str(entry["goFileCount"]),
                    str(entry["totalGoBytes"]),
                    entry["largestGoFile"] or "-",
                    str(entry["largestGoFileBytes"]),
                    "-" if elapsed is None else f"{elapsed:.3f}",
                ]
            )
            + " |"
        )
    if not report["entries"]:
        lines.append("| - | - | - | - | - | 0 | 0 | - | 0 | - |")
    lines.append("")
    return "\n".join(lines) + "\n"


def write_telemetry_report(results: list[CaseResult]) -> None:
    report = build_telemetry_report(results)
    TELEMETRY_DIR.mkdir(parents=True, exist_ok=True)
    (TELEMETRY_DIR / "examples.json").write_text(json.dumps(report, indent=2, sort_keys=True) + "\n", encoding="utf-8")
    (TELEMETRY_DIR / "examples.md").write_text(render_telemetry_markdown(report), encoding="utf-8")


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
    write_telemetry_report(results)
    print(f"\nSummary: {passed} passed, {failed} failed, {len(results)} total")

    return 1 if failed else 0


if __name__ == "__main__":
    raise SystemExit(main())
