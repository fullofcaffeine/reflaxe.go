#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
PUBLIC_CONTRACT = ROOT / "docs" / "public-contract.md"

SURFACE_CONTRACTS = {
    "Package and installation": {
        "authorities": (
            "haxelib.json",
            "package.json",
            "scripts/release/build-haxelib-artifact.py",
        ),
        "evidence": ("test/test_haxelib_release_install.py",),
    },
    "Toolchains and platforms": {
        "authorities": (
            "docs/toolchain-policy.json",
            "docs/compatibility-support-manifest.json",
        ),
        "evidence": ("test/test_toolchain_policy_contract.py",),
    },
    "Profiles and source controls": {
        "authorities": (
            "docs/native-policy-presets.md",
            "docs/profile-semantics-guide.md",
        ),
        "evidence": ("test/test_metal_preset_retention_contract.py",),
    },
    "Portable Haxe semantics": {
        "authorities": (
            "docs/compatibility-support-manifest.json",
            "test/portable_stdlib_inventory.json",
        ),
        "evidence": ("test/run-semantic-diff.py",),
    },
    "Go-native source APIs": {
        "authorities": (
            "docs/goextern.md",
            "std/go",
        ),
        "evidence": ("test/run-goextern-fixtures.py",),
    },
    "Generated Go and runtime ABI": {
        "authorities": (
            "docs/typed-go-ir.md",
            "runtime/hxrt",
        ),
        "evidence": ("test/snapshot/core/ast_typed_type_operator_printer",),
    },
    "Reports and versioned data": {
        "authorities": ("docs/profile-semantics-guide.md",),
        "evidence": ("test/snapshot/core/report_artifacts_basic",),
    },
    "Commands and diagnostics": {
        "authorities": ("docs/start-here.md",),
        "evidence": (
            "test/test_post_generation_build_runner.py",
            "test/snapshot/negative",
        ),
    },
}


class PublicContractTest(unittest.TestCase):
    def test_every_public_category_names_live_authority_and_evidence(self) -> None:
        contract = PUBLIC_CONTRACT.read_text(encoding="utf-8")

        for category, contract_parts in SURFACE_CONTRACTS.items():
            with self.subTest(category=category):
                row = next(
                    (line for line in contract.splitlines() if line.startswith(f"| {category} |")),
                    None,
                )
                self.assertIsNotNone(row, category)
                for part in ("authorities", "evidence"):
                    self.assertTrue(contract_parts[part], f"{category} has no {part}")
                    for path in contract_parts[part]:
                        self.assertIn(f"`{path}`", row)
                        self.assertTrue((ROOT / path).exists(), path)

    def test_contract_draws_semver_and_internal_boundaries(self) -> None:
        contract = PUBLIC_CONTRACT.read_text(encoding="utf-8")

        for heading in (
            "## What SemVer protects",
            "## What is internal",
            "## How a release number is chosen",
            "## Why this is federated",
            "## When to revisit declaration-level diffing",
        ):
            self.assertIn(heading, contract)

        self.assertIn("breaking change on `0.x`", contract)
        self.assertIn("explicit human approval", contract)
        self.assertIn("Conventional Commit", contract)
        self.assertIn("not a second semantic product", contract)
        self.assertIn("not a public API", contract)

    def test_sibling_lessons_are_pinned_and_actionable(self) -> None:
        contract = PUBLIC_CONTRACT.read_text(encoding="utf-8")

        expected_revisions = {
            "haxe.ruby": "ded7f02d666612350440d2d31e52dfe48449f9b9",
            "haxe.elixir.codex": "2030abea264dac770915dbeff427acc349ff082e",
            "haxe.rust": "85067736d0b929dfc67d6684d59b7e2bd3bae6ea",
            "haxe.c": "3a650c1481c2072e87a8ba89f9cdaaba29a244de",
        }
        for sibling, revision in expected_revisions.items():
            self.assertIn(sibling, contract)
            self.assertIn(revision, contract)

        self.assertIn("37,735", contract)
        self.assertIn("federated", contract)
        self.assertIn("capability inventory", contract)

    def test_release_docs_and_ci_consume_the_contract(self) -> None:
        release_policy = (ROOT / "docs" / "release-version-policy.md").read_text(encoding="utf-8")
        checklist = (ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        docs_index = (ROOT / "docs" / "index.md").read_text(encoding="utf-8")
        runner = (ROOT / "test" / "run-release-contracts.py").read_text(encoding="utf-8")
        evidence_builder = (ROOT / "scripts" / "review" / "build_gpt56_evidence.py").read_text(
            encoding="utf-8"
        )

        for consumer in (release_policy, checklist, docs_index):
            self.assertIn("public-contract.md", consumer)
        self.assertIn("test/test_public_contract.py", runner)
        self.assertIn('"docs/public-contract.md"', evidence_builder)
        self.assertIn('"test/test_public_contract.py"', evidence_builder)


if __name__ == "__main__":
    unittest.main()
