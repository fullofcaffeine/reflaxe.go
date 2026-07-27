"""Shared canonicalization for release-readiness evidence."""

from __future__ import annotations

import hashlib
import json
from typing import Any


def canonical_sha256(value: Any) -> str:
    encoded = json.dumps(
        value,
        ensure_ascii=False,
        separators=(",", ":"),
        sort_keys=True,
    ).encode("utf-8")
    return hashlib.sha256(encoded).hexdigest()


def compatibility_surface(manifest: dict[str, Any]) -> tuple[str, int, int]:
    """Hash the exact admitted operation/member surface and its evidence rows."""

    admitted_platforms = sorted(
        (
            {
                key: platform.get(key)
                for key in (
                    "id",
                    "os",
                    "architecture",
                    "state",
                    "qualification",
                    "evidence_ids",
                    "exclusions",
                )
            }
            for platform in manifest.get("platforms", [])
            if isinstance(platform, dict)
            and platform.get("release_admitted") is True
        ),
        key=lambda item: str(item["id"]),
    )
    admitted_presets = sorted(
        (
            {
                key: preset.get(key)
                for key in (
                    "id",
                    "selector",
                    "state",
                    "qualification",
                    "evidence_ids",
                    "exclusions",
                )
            }
            for preset in manifest.get("presets", [])
            if isinstance(preset, dict) and preset.get("release_admitted") is True
        ),
        key=lambda item: str(item["id"]),
    )
    admitted_operations: list[dict[str, Any]] = []
    referenced_evidence: set[str] = set()
    for item in admitted_platforms + admitted_presets:
        referenced_evidence.update(item.get("evidence_ids") or [])
    for surface in manifest.get("surfaces", []):
        if not isinstance(surface, dict):
            continue
        for operation in surface.get("operations", []):
            if not isinstance(operation, dict) or operation.get("release_admitted") is not True:
                continue
            record = {
                "surface_id": surface.get("id"),
                "operation_id": operation.get("id"),
                "symbols": operation.get("symbols"),
                "granularity": operation.get("granularity"),
                "state": operation.get("state"),
                "qualification": operation.get("qualification"),
                "evidence_ids": operation.get("evidence_ids"),
                "exclusions": operation.get("exclusions"),
                "blockers": operation.get("blockers"),
            }
            admitted_operations.append(record)
            referenced_evidence.update(operation.get("evidence_ids") or [])
    admitted_operations.sort(
        key=lambda item: (str(item["surface_id"]), str(item["operation_id"]))
    )
    evidence_by_id = {
        item.get("id"): item
        for item in manifest.get("evidence", [])
        if isinstance(item, dict) and isinstance(item.get("id"), str)
    }
    evidence = [evidence_by_id[item] for item in sorted(referenced_evidence)]
    required_trust = sorted(
        (
            item
            for item in manifest.get("trust_assumptions", [])
            if isinstance(item, dict) and item.get("release_required") is True
        ),
        key=lambda item: str(item.get("id")),
    )
    payload = {
        "release_claim": manifest.get("release_claim"),
        "platforms": admitted_platforms,
        "presets": admitted_presets,
        "operations": admitted_operations,
        "evidence": evidence,
        "trust_assumptions": required_trust,
    }
    symbol_count = sum(len(item.get("symbols") or []) for item in admitted_operations)
    return canonical_sha256(payload), len(admitted_operations), symbol_count
