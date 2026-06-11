#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path
from typing import Any

REPO_ROOT = Path(__file__).resolve().parent.parent
INVENTORY_FILE = REPO_ROOT / "test" / "portable_stdlib_inventory.json"
ALLOWLIST_FILE = REPO_ROOT / "test" / "portable_allowlist.json"
BEADS_FILE = REPO_ROOT / ".beads" / "issues.jsonl"
MAPPING_DOC = REPO_ROOT / "docs" / "portable-module-mapping-contract.md"
FEATURE_MATRIX_DOC = REPO_ROOT / "docs" / "feature-support-matrix.md"
KNOWN_GAPS_DOC = REPO_ROOT / "docs" / "known-gaps.md"
PHASE2_ROADMAP_DOC = REPO_ROOT / "docs" / "phase2-roadmap.md"
PARITY_PROGRAM_DOC = REPO_ROOT / "docs" / "portable-stdlib-parity-program.md"
RELEASE_CHECKLIST_DOC = REPO_ROOT / "docs" / "release-readiness-checklist.md"
APPROACH_C_BASELINE_DOC = REPO_ROOT / "docs" / "spikes" / "approach-c-baseline-freeze.md"
STDLIB_MIGRATION_LOG_DOC = REPO_ROOT / "docs" / "stdlib-shim-migration-log.md"

OWNER_COMPATIBILITY = {
    "staged_std": {"haxe_source", "mixed"},
    "runtime_hxrt": {"runtime_binding", "mixed"},
    "compiler_shim": {"compiler_intrinsic", "mixed"},
    "mixed": {"haxe_source", "runtime_binding", "compiler_intrinsic", "mixed"},
}


def load_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def load_beads_statuses() -> dict[str, str]:
    statuses: dict[str, str] = {}
    for line in BEADS_FILE.read_text(encoding="utf-8").splitlines():
        if not line.strip():
            continue
        row = json.loads(line)
        issue_id = row.get("id")
        status = row.get("status")
        if isinstance(issue_id, str) and isinstance(status, str):
            statuses[issue_id] = status
    return statuses


def parse_mapping_table() -> dict[str, str]:
    text = MAPPING_DOC.read_text(encoding="utf-8")
    rows: dict[str, str] = {}
    in_table = False
    for line in text.splitlines():
        if line.startswith("| Module | Ownership class |"):
            in_table = True
            continue
        if in_table:
            if not line.startswith("|"):
                break
            if line.startswith("| ---"):
                continue
            parts = [part.strip() for part in line.strip().strip("|").split("|")]
            if len(parts) < 2:
                continue
            module = parts[0].strip().strip("`")
            ownership = normalize_mapping_owner(parts[1])
            if module:
                rows[module] = ownership
    return rows


def normalize_mapping_owner(raw: str) -> str:
    value = raw.strip().strip("`")
    if value.startswith("mixed"):
        return "mixed"
    return value.split()[0]


class PortableGovernanceContractTest(unittest.TestCase):
    def test_tier1_allowlist_modules_have_mapping_rows(self) -> None:
        allowlist = load_json(ALLOWLIST_FILE)
        inventory = load_json(INVENTORY_FILE)
        mapping_rows = parse_mapping_table()
        inventory_by_module = {row["module"]: row for row in inventory["modules"]}

        missing: list[str] = []
        incompatible: list[str] = []
        for module in allowlist["tiers"]["tier1"]:
            if module not in mapping_rows:
                missing.append(module)
                continue
            inventory_owner = inventory_by_module[module]["owner"]
            mapped_owner = mapping_rows[module]
            allowed = OWNER_COMPATIBILITY.get(inventory_owner, {mapped_owner})
            if mapped_owner not in allowed:
                incompatible.append(f"{module}: inventory={inventory_owner}, mapping={mapped_owner}")

        self.assertFalse(missing, f"tier1 mapping rows missing: {missing}")
        self.assertFalse(incompatible, f"tier1 ownership mismatch: {incompatible}")

    def test_compile_only_modules_reference_open_blocker_beads(self) -> None:
        inventory = load_json(INVENTORY_FILE)
        bead_statuses = load_beads_statuses()
        failures: list[str] = []
        for row in inventory["modules"]:
            if row["status"] != "compile-only":
                continue
            issue = row.get("blocker_issue")
            module = row["module"]
            if not isinstance(issue, str) or not issue:
                failures.append(f"{module}: missing blocker_issue")
                continue
            status = bead_statuses.get(issue)
            if status is None:
                failures.append(f"{module}: blocker issue {issue} not found in .beads/issues.jsonl")
            elif status == "closed":
                failures.append(f"{module}: blocker issue {issue} is closed")
        self.assertFalse(failures, "compile-only inventory must point at live blocker beads:\n" + "\n".join(failures))

    def test_docs_point_to_generated_blocker_authority(self) -> None:
        feature_matrix = FEATURE_MATRIX_DOC.read_text(encoding="utf-8")
        parity_program = PARITY_PROGRAM_DOC.read_text(encoding="utf-8")
        release_checklist = RELEASE_CHECKLIST_DOC.read_text(encoding="utf-8")

        self.assertIn("test/portable_stdlib_inventory.json", feature_matrix)
        self.assertIn("test/.test-cache/portable_parity_closure_summary.md", feature_matrix)
        self.assertNotIn("haxe.go-14as.25` to `haxe.go-14as.27", feature_matrix)
        self.assertNotIn("Active follow-up tracking", feature_matrix)
        self.assertNotIn("## Tracking", feature_matrix)
        self.assertIn("## Completed Tracking History", feature_matrix)
        self.assertNotIn("Continue replacing reachable typed-node gaps", feature_matrix)
        self.assertIn("Invariant guard strategy", feature_matrix)
        self.assertIn("Completed follow-up evidence", feature_matrix)

        phase2_roadmap = PHASE2_ROADMAP_DOC.read_text(encoding="utf-8")
        self.assertNotIn("This file is the active Phase-2 roadmap", phase2_roadmap)
        self.assertNotIn("Execution tracker:", phase2_roadmap)
        self.assertNotIn("Approach-C closure follow-up tracker:", phase2_roadmap)
        self.assertNotIn("Latest follow-up:", phase2_roadmap)
        self.assertNotIn("Continue unsupported-expression inventory reduction", phase2_roadmap)
        self.assertIn("historical Phase-2 roadmap summary", phase2_roadmap)
        self.assertIn("Historical landed follow-up:", phase2_roadmap)

        known_gaps = KNOWN_GAPS_DOC.read_text(encoding="utf-8")
        self.assertNotIn("Remaining stdlib follow-up", known_gaps)
        self.assertIn("0 actionable portable stdlib blockers", known_gaps)
        self.assertIn("policy-locked", known_gaps)

        approach_c_baseline = APPROACH_C_BASELINE_DOC.read_text(encoding="utf-8")
        self.assertNotIn("compile-only stdlib promotion", approach_c_baseline)
        self.assertIn("current parity evidence", approach_c_baseline)

        mapping_contract = MAPPING_DOC.read_text(encoding="utf-8")
        self.assertNotIn("  - Tracking:", mapping_contract)
        self.assertIn("  - Closed evidence:", mapping_contract)

        stdlib_migration_log = STDLIB_MIGRATION_LOG_DOC.read_text(encoding="utf-8")
        self.assertNotIn("Staged stdlib migration follow-ups continue under `haxe.go-cgk.*`", stdlib_migration_log)
        self.assertNotIn("event-loop surfaces are explicitly unsupported on Go today", stdlib_migration_log)
        self.assertIn("`haxe.go-cgk.*` planning work is historical context", stdlib_migration_log)
        self.assertIn("At that point, direct Haxe event-loop surfaces were explicitly unsupported", stdlib_migration_log)

        self.assertIn("test/.test-cache/portable_parity_closure_summary.json", parity_program)
        self.assertIn("authoritative live list", parity_program)
        self.assertIn("closure_policy", parity_program)
        self.assertIn("actionable", parity_program)

        self.assertIn("npm run test:stdlib:governance", release_checklist)
        self.assertIn("npm run test:release-contracts", release_checklist)

    def test_known_gaps_has_release_checkable_production_scoreboard(self) -> None:
        known_gaps = KNOWN_GAPS_DOC.read_text(encoding="utf-8")
        release_checklist = RELEASE_CHECKLIST_DOC.read_text(encoding="utf-8")

        self.assertIn("## Production hardening scoreboard", known_gaps)
        self.assertIn("| Caveat class | Owner | Current decision | Evidence | Reopen trigger |", known_gaps)
        self.assertIn("Multi-package output", known_gaps)
        self.assertIn("Advanced Go extern interop", known_gaps)
        self.assertIn("Target-sensitive parity surfaces", known_gaps)
        self.assertIn("Performance budget drift", known_gaps)
        self.assertIn("Strict production boundary policy", known_gaps)
        self.assertIn("file a new scoped Bead", known_gaps)
        self.assertIn("Current GA hardening beads", known_gaps)
        self.assertIn("haxe.go-1bo0.2", known_gaps)
        self.assertIn("haxe.go-1bo0.7", known_gaps)
        self.assertNotIn("haxe.go-hm3p", known_gaps)

        self.assertIn("Production caveat scoreboard", release_checklist)
        self.assertIn("docs/known-gaps.md#production-hardening-scoreboard", release_checklist)


if __name__ == "__main__":
    unittest.main()
