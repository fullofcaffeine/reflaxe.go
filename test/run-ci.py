#!/usr/bin/env python3

from __future__ import annotations

import argparse
import subprocess
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent


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
    parser.add_argument("--skip-semantic-diff", action="store_true", help="Skip semantic differential stage")
    parser.add_argument("--force-semantic-diff", action="store_true", help="Run semantic differential stage even for chunked/filtered runs")
    parser.add_argument("--skip-semantic-diff-lanes", action="store_true", help="Skip lane semantic differential stage")
    parser.add_argument(
        "--force-semantic-diff-lanes",
        action="store_true",
        help="Run lane semantic differential stage even for chunked/filtered runs",
    )
    parser.add_argument("--skip-examples", action="store_true", help="Skip examples stage")
    parser.add_argument("--force-examples", action="store_true", help="Run examples even for chunked/filtered runs")
    parser.add_argument("--examples-compile-only", action="store_true", help="Run examples compile/go-test checks without go run stdout checks")
    parser.add_argument("--skip-goextern-fixtures", action="store_true", help="Skip goextern fixture drift stage")
    return parser.parse_args()


def run(cmd: list[str]) -> int:
    print("$", " ".join(cmd))
    proc = subprocess.run(cmd, cwd=ROOT)
    return proc.returncode


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


def build_goextern_fixtures_command() -> list[str]:
    return ["python3", "test/run-goextern-fixtures.py"]


def main() -> int:
    args = parse_args()

    print("==> Snapshot stage")
    snapshot_code = run(build_snapshot_command(args))
    if snapshot_code != 0:
        return snapshot_code

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

    if should_run_semantic_diff(args):
        print("==> Semantic diff stage")
        semantic_diff_code = run(build_semantic_diff_command(args))
        if semantic_diff_code != 0:
            return semantic_diff_code
    else:
        print("==> Skipping semantic diff stage")

    if should_run_semantic_diff_lanes(args):
        print("==> Semantic diff lanes stage")
        semantic_diff_lanes_code = run(build_semantic_diff_lanes_command(args))
        if semantic_diff_lanes_code != 0:
            return semantic_diff_lanes_code
    else:
        print("==> Skipping semantic diff lanes stage")

    if args.skip_goextern_fixtures:
        print("==> Skipping goextern fixtures stage")
    else:
        print("==> goextern fixtures stage")
        goextern_code = run(build_goextern_fixtures_command())
        if goextern_code != 0:
            return goextern_code

    if not should_run_examples(args):
        print("==> Skipping examples stage")
        return 0

    print("==> Examples stage")
    return run(build_examples_command(args))


if __name__ == "__main__":
    raise SystemExit(main())
