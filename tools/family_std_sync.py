#!/usr/bin/env python3

from __future__ import annotations

import argparse
import hashlib
import json
import subprocess
from dataclasses import dataclass
from pathlib import Path
from typing import Any

REPO_ROOT = Path(__file__).resolve().parent.parent
FAMILY_ROOT = REPO_ROOT / "family" / "reflaxe.family.std"
VERIFY_SCRIPT = FAMILY_ROOT / "tools" / "verify_family_std.py"
PIN_FILE = REPO_ROOT / "family" / "family_std_pin.json"
CACHE_ROOT = REPO_ROOT / "test" / ".test-cache"
DUAL_RUN_REPORT_JSON = CACHE_ROOT / "family_std_dual_run_report.json"
DUAL_RUN_REPORT_MD = CACHE_ROOT / "family_std_dual_run_report.md"


@dataclass(frozen=True)
class Mapping:
    canonical: Path
    family: Path
    kind: str  # text | json


MAPPINGS = [
    Mapping(
        canonical=REPO_ROOT / "docs" / "portable-semantics-v1.md",
        family=FAMILY_ROOT / "contracts" / "portable-semantics" / "v1.md",
        kind="text",
    ),
    Mapping(
        canonical=REPO_ROOT / "test" / "portable_allowlist.json",
        family=FAMILY_ROOT / "allowlists" / "portable_allowlist.v1.json",
        kind="json",
    ),
    Mapping(
        canonical=REPO_ROOT / "test" / "portable_conformance_tier1.json",
        family=FAMILY_ROOT / "conformance" / "tier1" / "portable_conformance_tier1.v1.json",
        kind="json",
    ),
    Mapping(
        canonical=REPO_ROOT / "docs" / "portable-module-mapping-contract.md",
        family=FAMILY_ROOT / "docs" / "module-mapping-contract.v1.md",
        kind="text",
    ),
]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Sync canonical portable artifacts with family std bootstrap snapshot."
    )
    parser.add_argument(
        "--mode",
        choices=["export", "import", "verify"],
        default="verify",
        help="export: canonical -> family, import: family -> canonical, verify: drift check + family validation",
    )
    parser.add_argument("--dry-run", action="store_true", help="Show actions without writing files")
    return parser.parse_args()


def fail(message: str) -> None:
    raise SystemExit(message)


def canonicalize_json_text(raw: str, path: Path) -> str:
    try:
        obj = json.loads(raw)
    except json.JSONDecodeError as exc:
        raise SystemExit(f"invalid JSON: {path} ({exc})") from exc
    return json.dumps(obj, indent=2, sort_keys=True) + "\n"


def normalized_content(path: Path, kind: str) -> str:
    raw = path.read_text(encoding="utf-8")
    if kind == "json":
        return canonicalize_json_text(raw, path)
    return raw


def sha256_text(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()


def write_normalized(path: Path, kind: str, content: str, dry_run: bool) -> bool:
    existing = path.read_text(encoding="utf-8") if path.exists() else ""
    if existing == content:
        return False
    if dry_run:
        return True
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")
    return True


def update_manifest(dry_run: bool) -> bool:
    manifest = FAMILY_ROOT / "MANIFEST.v1.txt"
    files: list[str] = []
    for path in FAMILY_ROOT.rglob("*"):
        if path.is_dir():
            continue
        rel = path.relative_to(FAMILY_ROOT).as_posix()
        if "/__pycache__/" in f"/{rel}/":
            continue
        files.append(rel)
    files = sorted(files)
    rendered = "\n".join(files) + "\n"
    existing = manifest.read_text(encoding="utf-8") if manifest.exists() else ""
    if existing == rendered:
        return False
    if not dry_run:
        manifest.write_text(rendered, encoding="utf-8")
    return True


def run_export(dry_run: bool) -> int:
    changed: list[Path] = []
    for mapping in MAPPINGS:
        content = normalized_content(mapping.canonical, mapping.kind)
        if write_normalized(mapping.family, mapping.kind, content, dry_run):
            changed.append(mapping.family)

    if update_manifest(dry_run):
        changed.append(FAMILY_ROOT / "MANIFEST.v1.txt")

    if changed:
        print("[PASS] export complete")
        for path in changed:
            print(f"- updated: {path.relative_to(REPO_ROOT)}")
    else:
        print("[PASS] export complete (no changes)")
    return 0


def run_import(dry_run: bool) -> int:
    changed: list[Path] = []
    for mapping in MAPPINGS:
        content = normalized_content(mapping.family, mapping.kind)
        if write_normalized(mapping.canonical, mapping.kind, content, dry_run):
            changed.append(mapping.canonical)

    if changed:
        print("[PASS] import complete")
        for path in changed:
            print(f"- updated: {path.relative_to(REPO_ROOT)}")
    else:
        print("[PASS] import complete (no changes)")
    return 0


def load_pin() -> dict[str, Any]:
    try:
        pin = json.loads(PIN_FILE.read_text(encoding="utf-8"))
    except FileNotFoundError as exc:
        fail(f"missing pin file: {PIN_FILE}")
        raise exc
    except json.JSONDecodeError as exc:
        fail(f"invalid pin JSON: {PIN_FILE} ({exc})")
        raise exc
    if not isinstance(pin, dict):
        fail("family_std_pin.json root must be object")
    if pin.get("schema_version") != 1:
        fail("family_std_pin.json schema_version must be 1")
    version = pin.get("version")
    if not isinstance(version, str) or not version.strip():
        fail("family_std_pin.json version must be non-empty string")
    return pin


def verify_pin_version(pin: dict[str, Any]) -> None:
    family_version = (FAMILY_ROOT / "VERSION").read_text(encoding="utf-8").strip()
    pinned = str(pin["version"]).strip()
    if family_version != pinned:
        fail(
            "family pin/version mismatch: "
            f"pin={pinned!r}, family VERSION={family_version!r}. "
            "fix: python3 tools/family_std_sync.py --mode export"
        )


def write_dual_run_report(pin: dict[str, Any], rows: list[dict[str, Any]], mismatch_count: int) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    payload = {
        "schema_version": 1,
        "pin": pin,
        "entry_count": len(rows),
        "mismatch_count": mismatch_count,
        "entries": rows,
    }
    DUAL_RUN_REPORT_JSON.write_text(json.dumps(payload, indent=2, sort_keys=True) + "\n", encoding="utf-8")

    lines = [
        "# Family Std Dual-Run Report",
        "",
        f"- pinned_name: `{pin.get('name', 'reflaxe.family.std')}`",
        f"- pinned_version: `{pin.get('version')}`",
        f"- source: `{pin.get('source', '')}`",
        f"- migration_mode: `{pin.get('migration_window', {}).get('mode', '')}`",
        f"- entry_count: `{len(rows)}`",
        f"- mismatch_count: `{mismatch_count}`",
        "",
        "## Entries",
    ]
    for row in rows:
        status = "match" if row["match"] else "mismatch"
        lines.append(
            f"- `{row['canonical']}` <-> `{row['family']}`: `{status}` "
            f"(sha256={row['canonical_sha256']})"
        )
    lines.append("")
    lines.append("Artifacts:")
    lines.append(f"- `{DUAL_RUN_REPORT_JSON.relative_to(REPO_ROOT)}`")
    lines.append(f"- `{DUAL_RUN_REPORT_MD.relative_to(REPO_ROOT)}`")
    DUAL_RUN_REPORT_MD.write_text("\n".join(lines) + "\n", encoding="utf-8")


def run_verify() -> int:
    pin = load_pin()
    verify_pin_version(pin)

    mismatches: list[Mapping] = []
    rows: list[dict[str, Any]] = []
    for mapping in MAPPINGS:
        canonical = normalized_content(mapping.canonical, mapping.kind)
        family = normalized_content(mapping.family, mapping.kind)
        canonical_sha = sha256_text(canonical)
        family_sha = sha256_text(family)
        match = canonical == family
        rows.append(
            {
                "canonical": mapping.canonical.relative_to(REPO_ROOT).as_posix(),
                "family": mapping.family.relative_to(REPO_ROOT).as_posix(),
                "kind": mapping.kind,
                "canonical_sha256": canonical_sha,
                "family_sha256": family_sha,
                "match": match,
            }
        )
        if canonical != family:
            mismatches.append(mapping)

    write_dual_run_report(pin, rows, len(mismatches))

    if mismatches:
        print("[FAIL] family std sync drift detected")
        for mapping in mismatches:
            print(f"- canonical: {mapping.canonical.relative_to(REPO_ROOT)}")
            print(f"  family:    {mapping.family.relative_to(REPO_ROOT)}")
        print("fix: python3 tools/family_std_sync.py --mode export")
        print(f"report: {DUAL_RUN_REPORT_JSON.relative_to(REPO_ROOT)}")
        return 1

    proc = subprocess.run(["python3", str(VERIFY_SCRIPT)], cwd=REPO_ROOT)
    if proc.returncode != 0:
        return proc.returncode

    print("[PASS] family std sync verify")
    print(f"[PASS] report: {DUAL_RUN_REPORT_JSON.relative_to(REPO_ROOT)}")
    return 0


def main() -> int:
    args = parse_args()
    if args.mode == "export":
        return run_export(args.dry_run)
    if args.mode == "import":
        return run_import(args.dry_run)
    if args.dry_run:
        raise SystemExit("--dry-run is only supported for export/import")
    return run_verify()


if __name__ == "__main__":
    raise SystemExit(main())
