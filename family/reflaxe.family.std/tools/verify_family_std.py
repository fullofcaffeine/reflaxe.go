#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path

FAMILY_ROOT = Path(__file__).resolve().parents[1]
REPO_ROOT = Path(__file__).resolve().parents[3]
MANIFEST_FILE = FAMILY_ROOT / "MANIFEST.v1.txt"


def fail(message: str) -> None:
    raise SystemExit(message)


def load_json(path: Path) -> dict:
    try:
        data = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as exc:
        fail(f"missing file: {path}")
        raise exc
    except json.JSONDecodeError as exc:
        fail(f"invalid JSON in {path}: {exc}")
        raise exc
    if not isinstance(data, dict):
        fail(f"JSON root must be object: {path}")
    return data


def read_manifest(path: Path) -> list[str]:
    if not path.exists():
        fail(f"missing manifest: {path}")
    entries: list[str] = []
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        entries.append(line)
    if entries != sorted(entries):
        fail("MANIFEST.v1.txt entries must be sorted")
    return entries


def actual_file_list(root: Path) -> list[str]:
    out: list[str] = []
    for path in root.rglob("*"):
        if path.is_dir():
            continue
        rel = path.relative_to(root).as_posix()
        if "/__pycache__/" in f"/{rel}/":
            continue
        out.append(rel)
    return sorted(out)


def validate_allowlist(path: Path) -> None:
    data = load_json(path)
    if data.get("schema_version") != 1:
        fail("allowlist schema_version must be 1")
    tiers = data.get("tiers")
    if not isinstance(tiers, dict):
        fail("allowlist tiers must be object")
    for tier_name, modules in tiers.items():
        if not isinstance(modules, list):
            fail(f"allowlist tier `{tier_name}` must be array")
        if modules != sorted(modules):
            fail(f"allowlist tier `{tier_name}` modules must be sorted")


def validate_conformance(path: Path) -> None:
    data = load_json(path)
    if data.get("schema_version") != 1:
        fail("conformance schema_version must be 1")
    if data.get("tier") != "tier1":
        fail("conformance tier must be tier1")
    modules = data.get("modules")
    if not isinstance(modules, dict):
        fail("conformance modules must be object")
    keys = list(modules.keys())
    if keys != sorted(keys):
        fail("conformance module keys must be sorted")
    for module, cases in modules.items():
        if not isinstance(cases, list) or not cases:
            fail(f"conformance module `{module}` must map to non-empty array")
        if cases != sorted(cases):
            fail(f"conformance cases for `{module}` must be sorted")
        if len(cases) != len(set(cases)):
            fail(f"conformance cases for `{module}` must be unique")


def validate_boundary_policy(path: Path) -> None:
    data = load_json(path)
    if data.get("schema_version") != 1:
        fail("boundary policy schema_version must be 1")
    approved = data.get("approved_override_roots")
    if not isinstance(approved, list) or approved != sorted(approved):
        fail("boundary policy approved_override_roots must be a sorted array")


def validate_sync() -> None:
    mirrors = [
        (
            REPO_ROOT / "docs/portable-semantics-v1.md",
            FAMILY_ROOT / "contracts/portable-semantics/v1.md",
        ),
        (
            REPO_ROOT / "test/portable_allowlist.json",
            FAMILY_ROOT / "allowlists/portable_allowlist.v1.json",
        ),
        (
            REPO_ROOT / "test/portable_conformance_tier1.json",
            FAMILY_ROOT / "conformance/tier1/portable_conformance_tier1.v1.json",
        ),
        (
            REPO_ROOT / "docs/portable-module-mapping-contract.md",
            FAMILY_ROOT / "docs/module-mapping-contract.v1.md",
        ),
    ]
    for source, mirror in mirrors:
        source_text = source.read_text(encoding="utf-8")
        mirror_text = mirror.read_text(encoding="utf-8")
        if source_text != mirror_text:
            fail(f"mirror drift detected: {mirror.relative_to(REPO_ROOT)} != {source.relative_to(REPO_ROOT)}")


def main() -> int:
    validate_allowlist(FAMILY_ROOT / "allowlists/portable_allowlist.v1.json")
    validate_conformance(FAMILY_ROOT / "conformance/tier1/portable_conformance_tier1.v1.json")
    validate_boundary_policy(FAMILY_ROOT / "provenance/upstream-boundary-policy.v1.json")

    manifest_entries = read_manifest(MANIFEST_FILE)
    files = actual_file_list(FAMILY_ROOT)
    if manifest_entries != files:
        missing = sorted(set(files) - set(manifest_entries))
        extra = sorted(set(manifest_entries) - set(files))
        parts: list[str] = []
        if missing:
            parts.append("missing entries: " + ", ".join(missing))
        if extra:
            parts.append("stale entries: " + ", ".join(extra))
        fail("manifest mismatch (" + "; ".join(parts) + ")")

    validate_sync()

    print("[PASS] family std bootstrap snapshot validated")
    print(f"[PASS] manifest entries: {len(files)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
