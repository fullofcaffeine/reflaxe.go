#!/usr/bin/env python3

from __future__ import annotations

import argparse
import os
import subprocess
from dataclasses import dataclass
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent


@dataclass(frozen=True)
class GoexternFixtureStage:
    kind: str
    reason: str
    command: list[str]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Run stable CI command surface for reflaxe.go")
    parser.add_argument("--chunk", default="", help="Deterministic shard in i/n form (e.g. 0/4)")
    parser.add_argument("--failed", action="store_true", help="Re-run only previously failing snapshot cases")
    parser.add_argument("--changed", action="store_true", help="Run only snapshot cases touched by git diff")
    parser.add_argument("--pattern", default="", help="Regex filter over snapshot case ids")
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per command in seconds")
    parser.add_argument(
        "--snapshot-lock-timeout",
        type=int,
        default=30,
        help="Seconds to wait for snapshot harness lock (passed to run-snapshots.py --lock-timeout)",
    )
    parser.add_argument("--skip-stdlib-sweep", action="store_true", help="Skip upstream stdlib sweep stage")
    parser.add_argument("--force-stdlib-sweep", action="store_true", help="Run stdlib sweep even for chunked/filtered runs")
    parser.add_argument("--stdlib-compile-only", action="store_true", help="Run stdlib sweep without go test stage")
    parser.add_argument("--skip-stdlib-full-sweep", action="store_true", help="Skip full portable-eligible stdlib sweep stage")
    parser.add_argument("--force-stdlib-full-sweep", action="store_true", help="Run full portable-eligible stdlib sweep even for chunked/filtered runs")
    parser.add_argument("--stdlib-full-go-test", action="store_true", help="Run full portable-eligible stdlib sweep with go test stage")
    parser.add_argument("--skip-stdlib-inventory", action="store_true", help="Skip portable stdlib inventory validation stage")
    parser.add_argument("--skip-portable-allowlist", action="store_true", help="Skip portable allowlist validation stage")
    parser.add_argument("--skip-portable-conformance", action="store_true", help="Skip portable Tier1 conformance stage")
    parser.add_argument(
        "--force-portable-conformance",
        action="store_true",
        help="Run portable Tier1 conformance stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-portable-parity-closure", action="store_true", help="Skip portable parity-closure summary stage")
    parser.add_argument("--skip-family-stdlib-bootstrap", action="store_true", help="Skip family std sync/verify stage")
    parser.add_argument("--skip-stdlib-governance", action="store_true", help="Skip stdlib provenance/boundary governance stage")
    parser.add_argument("--skip-optimizer-matrix", action="store_true", help="Skip optimizer-plan matrix snapshot stage")
    parser.add_argument(
        "--force-optimizer-matrix",
        action="store_true",
        help="Run optimizer-plan matrix snapshot stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-auto-planner-schema", action="store_true", help="Skip auto planner report-schema gate stage")
    parser.add_argument(
        "--force-auto-planner-schema",
        action="store_true",
        help="Run auto planner report-schema gate stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-release-contracts", action="store_true", help="Skip release contracts stage")
    parser.add_argument(
        "--force-release-contracts",
        action="store_true",
        help="Run release contracts stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-semantic-diff", action="store_true", help="Skip semantic differential stage")
    parser.add_argument("--force-semantic-diff", action="store_true", help="Run semantic differential stage even for chunked/filtered runs")
    parser.add_argument("--skip-semantic-diff-optimizer-matrix", action="store_true", help="Skip semantic-diff optimizer matrix stage")
    parser.add_argument(
        "--force-semantic-diff-optimizer-matrix",
        action="store_true",
        help="Run semantic-diff optimizer matrix stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-semantic-diff-lanes", action="store_true", help="Skip lane semantic differential stage")
    parser.add_argument(
        "--force-semantic-diff-lanes",
        action="store_true",
        help="Run lane semantic differential stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-metal-fallback-diagnostics", action="store_true", help="Skip metal fallback diagnostics stage")
    parser.add_argument(
        "--force-metal-fallback-diagnostics",
        action="store_true",
        help="Run metal fallback diagnostics stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-examples", action="store_true", help="Skip examples stage")
    parser.add_argument("--force-examples", action="store_true", help="Run examples even for chunked/filtered runs")
    parser.add_argument("--examples-compile-only", action="store_true", help="Run examples compile/go-test checks without go run stdout checks")
    parser.add_argument("--skip-metal-example-boundary", action="store_true", help="Skip metal example boundary stage")
    parser.add_argument("--skip-goextern-fixtures", action="store_true", help="Skip goextern fixture drift stage")
    return parser.parse_args()


def run(cmd: list[str]) -> int:
    print("$", " ".join(cmd))
    proc = subprocess.run(cmd, cwd=ROOT)
    return proc.returncode


def build_canonical_std_layout_command() -> list[str]:
    return ["python3", "test/test_canonical_std_layout_contract.py"]


def build_affected_plan_command() -> list[str]:
    return ["python3", "test/run-test-plan.py"]


def build_testing_strategy_command() -> list[str]:
    return ["npm", "run", "test:strategy"]


def build_official_haxe_target_smoke_command() -> list[str]:
    return ["npm", "run", "test:official-haxe-smoke"]


def build_output_confinement_command() -> list[str]:
    return ["python3", "test/test_generated_output_confinement.py"]


def build_snapshot_command(args: argparse.Namespace) -> list[str]:
    cmd = [
        "python3",
        "test/run-snapshots.py",
        "--timeout",
        str(args.timeout),
        "--lock-timeout",
        str(args.snapshot_lock_timeout),
    ]
    if args.chunk:
        cmd.extend(["--chunk", args.chunk])
    if args.failed:
        cmd.append("--failed")
    if args.changed:
        cmd.append("--changed")
    if args.pattern:
        cmd.extend(["--pattern", args.pattern])
    return cmd


def build_sys_get_char_terminal_command() -> list[str]:
    return ["npm", "run", "test:sys-get-char-terminal"]


def should_run_stdlib_sweep(args: argparse.Namespace) -> bool:
    if args.skip_stdlib_sweep:
        return False
    if args.force_stdlib_sweep:
        return True

    # For CI shards and focused reruns, snapshots are the intended signal.
    # Keep stdlib sweep on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_stdlib_command(args: argparse.Namespace) -> list[str]:
    cmd = ["python3", "test/run-upstream-stdlib-sweep.py", "--strict"]
    if not args.stdlib_compile_only:
        cmd.append("--go-test")
    return cmd


def should_run_stdlib_full_sweep(args: argparse.Namespace) -> bool:
    if args.skip_stdlib_full_sweep:
        return False
    if args.force_stdlib_full_sweep:
        return True

    # Full portable-eligible sweep is intended for complete CI runs.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_stdlib_full_command(args: argparse.Namespace) -> list[str]:
    cmd = [
        "python3",
        "test/run-upstream-stdlib-sweep.py",
        "--modules-file",
        "test/upstream_std_modules_full.txt",
        "--strict",
    ]
    if args.stdlib_full_go_test:
        cmd.append("--go-test")
    return cmd


def should_run_stdlib_inventory(args: argparse.Namespace) -> bool:
    return not args.skip_stdlib_inventory


def build_stdlib_inventory_command() -> list[str]:
    return ["python3", "test/run-portable-stdlib-inventory.py"]


def should_run_portable_allowlist(args: argparse.Namespace) -> bool:
    return not args.skip_portable_allowlist


def build_portable_allowlist_command() -> list[str]:
    return ["python3", "test/run-portable-allowlist.py"]


def should_run_portable_conformance(args: argparse.Namespace) -> bool:
    if args.skip_portable_conformance:
        return False
    if args.force_portable_conformance:
        return True

    # Keep portable conformance on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_portable_conformance_command(args: argparse.Namespace) -> list[str]:
    return ["python3", "test/run-portable-conformance.py", "--timeout", str(args.timeout)]


def should_run_portable_parity_closure(args: argparse.Namespace) -> bool:
    return not args.skip_portable_parity_closure


def build_portable_parity_closure_command() -> list[str]:
    return ["python3", "test/run-portable-parity-closure.py"]


def should_run_family_std_bootstrap(args: argparse.Namespace) -> bool:
    return not args.skip_family_stdlib_bootstrap


def build_family_std_bootstrap_command() -> list[str]:
    return ["python3", "tools/family_std_sync.py", "--mode", "verify"]


def should_run_stdlib_governance(args: argparse.Namespace) -> bool:
    return not args.skip_stdlib_governance


def build_stdlib_governance_command() -> list[str]:
    return ["npm", "run", "test:stdlib:governance"]


def should_run_optimizer_matrix(args: argparse.Namespace) -> bool:
    if args.skip_optimizer_matrix:
        return False
    if args.force_optimizer_matrix:
        return True

    # Keep optimizer matrix on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_optimizer_matrix_command() -> list[str]:
    return ["npm", "run", "test:optimizer:matrix"]


def should_run_auto_planner_schema(args: argparse.Namespace) -> bool:
    if args.skip_auto_planner_schema:
        return False
    if args.force_auto_planner_schema:
        return True

    # Keep schema gate on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_auto_planner_schema_command() -> list[str]:
    return ["npm", "run", "test:auto-planner:schema"]


def should_run_release_contracts(args: argparse.Namespace) -> bool:
    if args.skip_release_contracts:
        return False
    if args.force_release_contracts:
        return True

    # Keep release contracts on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_release_contracts_command() -> list[str]:
    return ["python3", "test/run-release-contracts.py"]


def should_run_semantic_diff(args: argparse.Namespace) -> bool:
    if args.skip_semantic_diff:
        return False
    if args.force_semantic_diff:
        return True

    # Keep semantic diff on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_semantic_diff_command(args: argparse.Namespace) -> list[str]:
    cmd = [
        "python3",
        "test/run-semantic-diff.py",
        "--timeout",
        str(args.timeout),
    ]
    if args.changed:
        cmd.append("--changed")
    return cmd


def should_run_semantic_diff_optimizer_matrix(args: argparse.Namespace) -> bool:
    if args.skip_semantic_diff_optimizer_matrix:
        return False
    if args.force_semantic_diff_optimizer_matrix:
        return True

    # Keep semantic-diff optimizer matrix on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_semantic_diff_optimizer_matrix_command() -> list[str]:
    return ["npm", "run", "test:semantic-diff:optimizer-matrix"]


def should_run_semantic_diff_lanes(args: argparse.Namespace) -> bool:
    if args.skip_semantic_diff_lanes:
        return False
    if args.force_semantic_diff_lanes:
        return True

    # Keep lane semantic diff on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_semantic_diff_lanes_command(args: argparse.Namespace) -> list[str]:
    cmd = [
        "python3",
        "test/run-semantic-diff.py",
        "--suite",
        "lanes",
        "--timeout",
        str(args.timeout),
    ]
    if args.changed:
        cmd.append("--changed")
    return cmd


def should_run_metal_fallback_diagnostics(args: argparse.Namespace) -> bool:
    if args.skip_metal_fallback_diagnostics:
        return False
    if args.force_metal_fallback_diagnostics:
        return True

    # Keep fallback diagnostics on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_metal_fallback_diagnostics_command(args: argparse.Namespace) -> list[str]:
    return [
        "python3",
        "test/run-snapshots.py",
        "--timeout",
        str(args.timeout),
        "--lock-timeout",
        str(args.snapshot_lock_timeout),
        "--case",
        "core/report_artifacts_lane_fallback",
    ]


def should_run_examples(args: argparse.Namespace) -> bool:
    if args.skip_examples:
        return False
    if args.force_examples:
        return True

    # Keep examples on full runs by default.
    return not (args.chunk or args.failed or args.changed or args.pattern)


def build_examples_command(args: argparse.Namespace) -> list[str]:
    cmd = [
        "python3",
        "test/run-examples.py",
        "--timeout",
        str(args.timeout),
    ]
    if args.examples_compile_only:
        cmd.append("--compile-only")
    if args.changed:
        cmd.append("--changed")
    return cmd


def should_run_metal_example_boundary(args: argparse.Namespace) -> bool:
    return not args.skip_metal_example_boundary


def build_metal_example_boundary_command() -> list[str]:
    return ["python3", "test/run-metal-example-boundary.py"]


def build_goextern_fixtures_command() -> list[str]:
    return ["python3", "test/run-goextern-fixtures.py"]


def build_goextern_fixture_smoke_command() -> list[str]:
    return ["python3", "test/run-goextern-fixtures.py", "--smoke"]


def build_goextern_command() -> list[str]:
    return ["npm", "run", "test:goextern"]


def current_go_release() -> str | None:
    proc = subprocess.run(
        ["go", "env", "GOVERSION"],
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=20,
    )
    if proc.returncode != 0:
        return None

    raw = proc.stdout.strip().lower()
    if not raw.startswith("go"):
        return None

    parts = raw[2:].split(".")
    if len(parts) < 2:
        return None
    if not parts[0].isdigit() or not parts[1].isdigit():
        return None

    return f"{int(parts[0])}.{int(parts[1])}"


def goextern_fixture_target_release() -> str:
    target_release = os.environ.get("GOEXTERN_FIXTURE_GO_VERSION", "1.23").strip()
    if not target_release:
        target_release = "1.23"
    return target_release


def resolve_goextern_fixture_stage(skip: bool, current_release: str | None, target_release: str) -> GoexternFixtureStage:
    if skip:
        return GoexternFixtureStage("skipped", "explicitly skipped via --skip-goextern-fixtures", [])

    if current_release is None:
        return GoexternFixtureStage("smoke", "unable to detect current Go release; running current toolchain smoke", build_goextern_fixture_smoke_command())

    if current_release != target_release:
        return GoexternFixtureStage(
            "smoke",
            f"fixtures are pinned to Go {target_release}; current toolchain is Go {current_release}; running current toolchain smoke",
            build_goextern_fixture_smoke_command(),
        )

    return GoexternFixtureStage(
        "fixtures",
        f"Go {current_release} matches pinned fixture toolchain",
        build_goextern_fixtures_command(),
    )


def main() -> int:
    args = parse_args()

    print("==> Affected ownership plan stage (observation only)")
    affected_plan_code = run(build_affected_plan_command())
    if affected_plan_code != 0:
        return affected_plan_code

    print("==> Testing strategy contract stage")
    strategy_code = run(build_testing_strategy_command())
    if strategy_code != 0:
        return strategy_code

    print("==> Official Haxe target smoke stage")
    official_haxe_smoke_code = run(build_official_haxe_target_smoke_command())
    if official_haxe_smoke_code != 0:
        return official_haxe_smoke_code

    print("==> Canonical std layout contract stage")
    canonical_std_layout_code = run(build_canonical_std_layout_command())
    if canonical_std_layout_code != 0:
        return canonical_std_layout_code

    print("==> Generated-output confinement stage")
    output_confinement_code = run(build_output_confinement_command())
    if output_confinement_code != 0:
        return output_confinement_code

    print("==> Snapshot stage")
    snapshot_code = run(build_snapshot_command(args))
    if snapshot_code != 0:
        return snapshot_code

    print("==> Sys.getChar terminal contract stage")
    sys_get_char_terminal_code = run(build_sys_get_char_terminal_command())
    if sys_get_char_terminal_code != 0:
        return sys_get_char_terminal_code

    if should_run_stdlib_sweep(args):
        print("==> Upstream stdlib sweep stage")
        stdlib_code = run(build_stdlib_command(args))
        if stdlib_code != 0:
            return stdlib_code
    else:
        print("==> Skipping stdlib sweep stage")

    if should_run_stdlib_full_sweep(args):
        print("==> Full portable-eligible stdlib sweep stage")
        stdlib_full_code = run(build_stdlib_full_command(args))
        if stdlib_full_code != 0:
            return stdlib_full_code
    else:
        print("==> Skipping full portable-eligible stdlib sweep stage")

    if should_run_stdlib_inventory(args):
        print("==> Portable stdlib inventory stage")
        stdlib_inventory_code = run(build_stdlib_inventory_command())
        if stdlib_inventory_code != 0:
            return stdlib_inventory_code
    else:
        print("==> Skipping portable stdlib inventory stage")

    if should_run_portable_allowlist(args):
        print("==> Portable allowlist stage")
        portable_allowlist_code = run(build_portable_allowlist_command())
        if portable_allowlist_code != 0:
            return portable_allowlist_code
    else:
        print("==> Skipping portable allowlist stage")

    if should_run_portable_conformance(args):
        print("==> Portable conformance stage")
        portable_conformance_code = run(build_portable_conformance_command(args))
        if portable_conformance_code != 0:
            return portable_conformance_code
    else:
        print("==> Skipping portable conformance stage")

    if should_run_portable_parity_closure(args):
        print("==> Portable parity closure summary stage")
        portable_parity_closure_code = run(build_portable_parity_closure_command())
        if portable_parity_closure_code != 0:
            return portable_parity_closure_code
    else:
        print("==> Skipping portable parity closure summary stage")

    if should_run_family_std_bootstrap(args):
        print("==> Family std sync/verify stage")
        family_std_bootstrap_code = run(build_family_std_bootstrap_command())
        if family_std_bootstrap_code != 0:
            return family_std_bootstrap_code
    else:
        print("==> Skipping family std sync/verify stage")

    if should_run_stdlib_governance(args):
        print("==> Stdlib governance stage")
        stdlib_governance_code = run(build_stdlib_governance_command())
        if stdlib_governance_code != 0:
            return stdlib_governance_code
    else:
        print("==> Skipping stdlib governance stage")

    if should_run_optimizer_matrix(args):
        print("==> Optimizer matrix stage")
        optimizer_matrix_code = run(build_optimizer_matrix_command())
        if optimizer_matrix_code != 0:
            return optimizer_matrix_code
    else:
        print("==> Skipping optimizer matrix stage")

    if should_run_auto_planner_schema(args):
        print("==> Auto planner report schema stage")
        auto_planner_schema_code = run(build_auto_planner_schema_command())
        if auto_planner_schema_code != 0:
            return auto_planner_schema_code
    else:
        print("==> Skipping auto planner report schema stage")

    if should_run_release_contracts(args):
        print("==> Release contracts stage")
        release_contracts_code = run(build_release_contracts_command())
        if release_contracts_code != 0:
            return release_contracts_code
    else:
        print("==> Skipping release contracts stage")

    if should_run_semantic_diff(args):
        print("==> Semantic diff stage")
        semantic_diff_code = run(build_semantic_diff_command(args))
        if semantic_diff_code != 0:
            return semantic_diff_code
    else:
        print("==> Skipping semantic diff stage")

    if should_run_semantic_diff_optimizer_matrix(args):
        print("==> Semantic diff optimizer matrix stage")
        semantic_diff_optimizer_matrix_code = run(build_semantic_diff_optimizer_matrix_command())
        if semantic_diff_optimizer_matrix_code != 0:
            return semantic_diff_optimizer_matrix_code
    else:
        print("==> Skipping semantic diff optimizer matrix stage")

    if should_run_semantic_diff_lanes(args):
        print("==> Semantic diff lanes stage")
        semantic_diff_lanes_code = run(build_semantic_diff_lanes_command(args))
        if semantic_diff_lanes_code != 0:
            return semantic_diff_lanes_code
    else:
        print("==> Skipping semantic diff lanes stage")

    if should_run_metal_fallback_diagnostics(args):
        print("==> Metal fallback diagnostics stage")
        metal_fallback_diagnostics_code = run(build_metal_fallback_diagnostics_command(args))
        if metal_fallback_diagnostics_code != 0:
            return metal_fallback_diagnostics_code
    else:
        print("==> Skipping metal fallback diagnostics stage")

    if should_run_metal_example_boundary(args):
        print("==> Metal example boundary stage")
        metal_example_boundary_code = run(build_metal_example_boundary_command())
        if metal_example_boundary_code != 0:
            return metal_example_boundary_code
    else:
        print("==> Skipping metal example boundary stage")

    print("==> goextern confidence stage")
    goextern_code = run(build_goextern_command())
    if goextern_code != 0:
        return goextern_code

    goextern_stage = resolve_goextern_fixture_stage(
        skip=args.skip_goextern_fixtures,
        current_release=current_go_release(),
        target_release=goextern_fixture_target_release(),
    )
    if goextern_stage.kind == "skipped":
        print(f"==> Skipping goextern fixtures stage ({goextern_stage.reason})")
    elif goextern_stage.kind == "smoke":
        print(f"==> goextern current-toolchain smoke stage ({goextern_stage.reason})")
        goextern_smoke_code = run(goextern_stage.command)
        if goextern_smoke_code != 0:
            return goextern_smoke_code
    else:
        print(f"==> goextern fixtures stage ({goextern_stage.reason})")
        goextern_code = run(goextern_stage.command)
        if goextern_code != 0:
            return goextern_code

    if not should_run_examples(args):
        print("==> Skipping examples stage")
        return 0

    print("==> Examples stage")
    return run(build_examples_command(args))


if __name__ == "__main__":
    raise SystemExit(main())
