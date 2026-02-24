#!/usr/bin/env python3

from __future__ import annotations

import argparse
import shutil
import subprocess
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
FIXTURE_ROOT = ROOT / "test" / "fixtures" / "goextern"
CACHE_ROOT = ROOT / "test" / ".test-cache" / "goextern-fixtures"
DEFAULT_PACKAGES = ("context", "errors", "fmt", "sync", "time")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Validate deterministic goextern fixtures")
    parser.add_argument("--update", action="store_true", help="Refresh committed fixtures from current generator output")
    parser.add_argument(
        "--package",
        action="append",
        default=[],
        help="Go package import path to include (repeatable). Default: context, errors, fmt, sync, time",
    )
    parser.add_argument("--timeout", type=int, default=120, help="Timeout per generator invocation in seconds")
    return parser.parse_args()


def run_command(cmd: list[str], cwd: Path, timeout_s: int) -> subprocess.CompletedProcess[str]:
    return subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout_s)


def all_files(root: Path) -> list[Path]:
    if not root.exists():
        return []
    return [path for path in sorted(root.rglob("*")) if path.is_file()]


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


def main() -> int:
    args = parse_args()

    packages = tuple(sorted({pkg.strip() for pkg in (args.package or DEFAULT_PACKAGES) if pkg.strip()}))
    if not packages:
        print("No packages selected")
        return 0

    if CACHE_ROOT.exists():
        shutil.rmtree(CACHE_ROOT)
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)

    for go_pkg in packages:
        cmd = [
            "bash",
            "scripts/dev/goextern.sh",
            "--package",
            go_pkg,
            "--out",
            str(CACHE_ROOT),
            "--haxe-package",
            "goextern",
        ]
        proc = run_command(cmd, cwd=ROOT, timeout_s=args.timeout)
        if proc.returncode != 0:
            print(f"[FAIL] generate {go_pkg}")
            if proc.stdout.strip():
                print(proc.stdout.strip())
            if proc.stderr.strip():
                print(proc.stderr.strip())
            return 1
        print(f"[PASS] generate {go_pkg}")

    if args.update:
        if FIXTURE_ROOT.exists():
            shutil.rmtree(FIXTURE_ROOT)
        shutil.copytree(CACHE_ROOT, FIXTURE_ROOT)
        print(f"Updated fixtures: {FIXTURE_ROOT}")
        return 0

    deltas = collect_tree_deltas(FIXTURE_ROOT, CACHE_ROOT)
    if deltas:
        print("[FAIL] goextern fixture drift detected")
        for line in deltas[:50]:
            print(line)
        if len(deltas) > 50:
            print(f"... and {len(deltas) - 50} more")
        print("Run: python3 test/run-goextern-fixtures.py --update")
        return 1

    print("[PASS] goextern fixtures are up to date")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
