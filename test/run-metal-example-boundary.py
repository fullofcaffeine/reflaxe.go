#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent

BOUNDARY_TARGET_GLOBS = [
    "examples/*/app/runtime/GoNativeRuntime.hx",
    "examples/*/profile/MetalRuntime.hx",
]

EXAMPLES_MANIFEST = Path("examples/qa-manifest.json")

FORBIDDEN_IMPORTS = [
    re.compile(r"^\s*import\s+haxe\.ds\.List\s*;"),
    re.compile(r"^\s*import\s+haxe\.ds\.IntMap\s*;"),
    re.compile(r"^\s*import\s+haxe\.ds\.StringMap\s*;"),
]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Enforce and audit metal example collection-boundary policy.")
    parser.add_argument("--root", type=Path, default=ROOT, help="Repository root to scan.")
    parser.add_argument(
        "--scope",
        choices=["boundary", "full"],
        default="boundary",
        help="boundary=metal adapter modules only, full=all modules in examples that declare a metal profile.",
    )
    parser.add_argument(
        "--mode",
        choices=["enforce", "audit"],
        default="enforce",
        help="enforce fails on violations, audit always exits 0 and reports findings.",
    )
    parser.add_argument(
        "--max-violations",
        type=int,
        help="Optional violation threshold; when provided in audit mode, exit 1 if violations exceed this count.",
    )
    parser.add_argument("--report", type=Path, help="Optional JSON report output path.")
    return parser.parse_args()


def metal_example_ids(root: Path) -> list[str]:
    manifest = root / EXAMPLES_MANIFEST
    if not manifest.exists():
        raise RuntimeError(f"missing examples QA manifest: {manifest}")
    payload = json.loads(manifest.read_text(encoding="utf-8"))
    if payload.get("schemaVersion") != 2 or not isinstance(payload.get("examples"), list):
        raise RuntimeError(f"invalid examples QA manifest schema: {manifest}")

    result: list[str] = []
    seen: set[str] = set()
    for raw in payload["examples"]:
        if not isinstance(raw, dict):
            raise RuntimeError(f"invalid example record in {manifest}")
        example_id = str(raw.get("id", "")).strip()
        profiles = raw.get("profiles")
        if not re.fullmatch(r"[A-Za-z0-9_-]+", example_id) or example_id in seen:
            raise RuntimeError(f"invalid or duplicate example id in {manifest}: {example_id!r}")
        if not isinstance(profiles, list) or not all(isinstance(item, str) for item in profiles):
            raise RuntimeError(f"example {example_id} has invalid profiles in {manifest}")
        seen.add(example_id)
        if "metal" in profiles:
            result.append(example_id)
    return sorted(result)


def iter_target_files(root: Path, scope: str) -> list[Path]:
    files: list[Path] = []
    seen: set[Path] = set()
    patterns = BOUNDARY_TARGET_GLOBS
    if scope == "full":
        patterns = [f"examples/{example_id}/**/*.hx" for example_id in metal_example_ids(root)]
    for pattern in patterns:
        for path in sorted(root.glob(pattern)):
            resolved = path.resolve()
            if resolved in seen:
                continue
            seen.add(resolved)
            files.append(path)
    return files


def relative_posix(path: Path, root: Path) -> str:
    try:
        return path.relative_to(root).as_posix()
    except ValueError:
        return path.as_posix()


def collect_violations(root: Path, targets: list[Path]) -> list[dict[str, object]]:
    violations: list[dict[str, object]] = []
    for path in targets:
        text = path.read_text(encoding="utf-8")
        for lineno, line in enumerate(text.splitlines(), start=1):
            for pattern in FORBIDDEN_IMPORTS:
                if pattern.search(line):
                    violations.append(
                        {
                            "path": relative_posix(path, root),
                            "line": lineno,
                            "import": line.strip(),
                            "reason": "forbidden portable collection import in metal example module",
                        }
                    )
    violations.sort(key=lambda entry: (str(entry["path"]), int(entry["line"]), str(entry["import"])))
    return violations


def write_report(report_path: Path, payload: dict[str, object]) -> None:
    report_path.parent.mkdir(parents=True, exist_ok=True)
    report_path.write_text(json.dumps(payload, indent=2, sort_keys=True) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()
    if args.max_violations is not None and args.max_violations < 0:
        print("--max-violations must be >= 0")
        return 2

    root = args.root.resolve()
    try:
        targets = iter_target_files(root, args.scope)
    except (json.JSONDecodeError, RuntimeError) as error:
        print(f"Metal example boundary target selection failed: {error}")
        return 2
    violations = collect_violations(root, targets)
    violation_count = len(violations)
    threshold_exceeded = args.max_violations is not None and violation_count > args.max_violations

    payload: dict[str, object] = {
        "schemaVersion": 1,
        "scope": args.scope,
        "mode": args.mode,
        "scannedFiles": [relative_posix(path, root) for path in targets],
        "violations": violations,
        "summary": {
            "scannedFileCount": len(targets),
            "violationCount": violation_count,
            "maxViolations": args.max_violations,
            "thresholdExceeded": threshold_exceeded,
        },
    }

    if args.report is not None:
        write_report(args.report, payload)

    if args.mode == "audit":
        print(f"Metal example boundary audit report: {violation_count} violation(s), {len(targets)} file(s) scanned.")
        if threshold_exceeded:
            print(f"Metal example boundary violation threshold exceeded: {violation_count} > {args.max_violations}")
        if args.report is not None:
            print(f"Report written to: {args.report}")
        return 1 if threshold_exceeded else 0

    if violations:
        print("Metal example boundary check failed:")
        for entry in violations:
            print(f"- {entry['path']}:{entry['line']}: {entry['reason']}: {entry['import']}")
        if args.report is not None:
            print(f"Report written to: {args.report}")
        return 1

    print(f"Metal example boundary check passed ({len(targets)} files scanned).")
    if args.report is not None:
        print(f"Report written to: {args.report}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
