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
    parser.add_argument("--skip-semantic-diff", action="store_true", help="Skip semantic differential stage")
    parser.add_argument("--force-semantic-diff", action="store_true", help="Run semantic differential stage even for chunked/filtered runs")
    parser.add_argument("--skip-examples", action="store_true", help="Skip examples stage")
    parser.add_argument("--force-examples", action="store_true", help="Run examples even for chunked/filtered runs")
    parser.add_argument("--examples-compile-only", action="store_true", help="Run examples compile/go-test checks without go run stdout checks")
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

    if should_run_semantic_diff(args):
        print("==> Semantic diff stage")
        semantic_diff_code = run(build_semantic_diff_command(args))
        if semantic_diff_code != 0:
            return semantic_diff_code
    else:
        print("==> Skipping semantic diff stage")

    if not should_run_examples(args):
        print("==> Skipping examples stage")
        return 0

    print("==> Examples stage")
    return run(build_examples_command(args))


if __name__ == "__main__":
    raise SystemExit(main())
