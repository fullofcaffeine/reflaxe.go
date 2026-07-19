#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "metal-preset-retention-decision.md"
USAGE_EVIDENCE = (
    ROOT
    / "docs"
    / "reviews"
    / "gpt-5.6-pro"
    / "metal-preset-usage-evidence-vfp-6.6.json"
)
REVIEW = (
    ROOT
    / "docs"
    / "reviews"
    / "gpt-5.6-pro"
    / "review-vfp-6.6-metal-preset.md"
)
REVIEW_PROMPT = (
    ROOT
    / "docs"
    / "reviews"
    / "gpt-5.6-pro"
    / "review-prompt-vfp-6.6-metal-preset.md"
)
REVIEW_PROVENANCE = (
    ROOT
    / "docs"
    / "reviews"
    / "gpt-5.6-pro"
    / "review-vfp-6.6-metal-preset.provenance.json"
)
METAL_FIXTURE = ROOT / "test" / "snapshot" / "go_native" / "metal_preset_equivalence"
EXPLICIT_FIXTURE = (
    ROOT / "test" / "snapshot" / "go_native" / "explicit_policy_equivalence"
)


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def tracked_files(pattern: str) -> list[Path]:
    result = subprocess.run(
        ["git", "ls-files", pattern],
        cwd=ROOT,
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        check=True,
    )
    return [ROOT / line for line in result.stdout.splitlines() if line]


def git_lines(*args: str) -> list[str]:
    result = subprocess.run(
        ["git", *args],
        cwd=ROOT,
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        check=True,
    )
    return [line for line in result.stdout.splitlines() if line]


class MetalPresetRetentionContractTest(unittest.TestCase):
    def test_decision_retains_the_preset_without_authorizing_deprecation(self) -> None:
        text = DECISION.read_text(encoding="utf-8")
        normalized = " ".join(text.replace("`", "").split()).lower()

        self.assertIn("decision: retain", normalized)
        self.assertIn("supported without a deprecation warning", normalized)
        self.assertIn("does not authorize deprecation", normalized)
        self.assertIn("haxe_go-vfp.6.6", text)
        for heading in (
            "## Remaining preset-only behavior",
            "## Observed usage and migration impact",
            "## Explicit-axis replacement",
            "## Sibling-compiler precedent",
            "## SemVer and migration timeline",
            "## Reopen and rollback criteria",
            "## Independent review",
        ):
            self.assertIn(heading, text)

    def test_profile_branches_stay_confined_to_mapping_and_reports(self) -> None:
        source_root = ROOT / "src" / "reflaxe" / "go"
        haxe_sources = list(source_root.rglob("*.hx"))
        profile_branch_files = {
            path.relative_to(ROOT).as_posix()
            for path in haxe_sources
            if "GoProfile.Metal" in path.read_text(encoding="utf-8")
        }
        self.assertEqual(
            {
                "src/reflaxe/go/ProfileResolver.hx",
                "src/reflaxe/go/GoReflaxeCompiler.hx",
                "src/reflaxe/go/compiler/GoPolicyPreset.hx",
            },
            profile_branch_files,
        )

        preset_literal_counts = {
            path.relative_to(ROOT).as_posix(): text.count(
                "GoPolicyPreset.MetalCompatibility"
            )
            for path in haxe_sources
            if (
                text := path.read_text(encoding="utf-8")
            ).count("GoPolicyPreset.MetalCompatibility")
        }
        self.assertEqual(
            {
                "src/reflaxe/go/compiler/GoBuildContext.hx": 1,
                "src/reflaxe/go/compiler/GoBuildContextResolver.hx": 2,
                "src/reflaxe/go/compiler/GoNativeAuthorityPolicy.hx": 1,
                "src/reflaxe/go/compiler/GoNativeFallbackPolicy.hx": 1,
                "src/reflaxe/go/compiler/GoNativeSpecializationPolicy.hx": 1,
            },
            preset_literal_counts,
        )

        compatibility_predicate_counts = {
            path.relative_to(ROOT).as_posix(): text.count(
                "usesMetalCompatibilityPreset"
            )
            for path in haxe_sources
            if (
                text := path.read_text(encoding="utf-8")
            ).count("usesMetalCompatibilityPreset")
        }
        self.assertEqual(
            {
                "src/reflaxe/go/GoReflaxeCompiler.hx": 4,
                "src/reflaxe/go/compiler/GoBuildContext.hx": 2,
            },
            compatibility_predicate_counts,
        )

        policy_preset_files = {
            path.relative_to(ROOT).as_posix()
            for path in haxe_sources
            if "policyPreset" in path.read_text(encoding="utf-8")
        }
        self.assertEqual(
            {
                "src/reflaxe/go/GoReflaxeCompiler.hx",
                "src/reflaxe/go/ast/transformers/registry/GoASTPassRegistry.hx",
                "src/reflaxe/go/compiler/GoBuildContext.hx",
            },
            policy_preset_files,
        )

        report_compiler = (
            ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")
        self.assertEqual(3, report_compiler.count("buildContext.policyPreset.label()"))
        planner = (
            ROOT
            / "src"
            / "reflaxe"
            / "go"
            / "ast"
            / "transformers"
            / "registry"
            / "GoASTPassRegistry.hx"
        ).read_text(encoding="utf-8")
        self.assertEqual(1, planner.count("context.buildContext.policyPreset.label()"))
        self.assertEqual(1, planner.count("policyPreset"))

        metal_contract_predicate_files = {
            path.relative_to(ROOT).as_posix()
            for path in haxe_sources
            if "isMetalContract(" in path.read_text(encoding="utf-8")
        }
        self.assertEqual(
            {"src/reflaxe/go/compiler/GoBuildContext.hx"},
            metal_contract_predicate_files,
        )

        compiler = (ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx").read_text(
            encoding="utf-8"
        )
        self.assertNotIn("GoProfile.Metal", compiler)
        self.assertNotIn("GoPolicyPreset.MetalCompatibility", compiler)
        self.assertNotIn("usesMetalCompatibilityPreset", compiler)
        self.assertNotIn("isMetalContract", compiler)
        self.assertNotIn("policyPreset", compiler)

    def test_repository_usage_inventory_is_current(self) -> None:
        evidence = json.loads(USAGE_EVIDENCE.read_text(encoding="utf-8"))
        inventory = evidence["repositoryInventory"]

        inventory_commit = evidence["repositoryInventoryCommit"]
        self.assertRegex(inventory_commit, r"^[0-9a-f]{40}$")
        self.assertEqual(
            evidence["repositoryInventoryTree"],
            git_lines("rev-parse", f"{inventory_commit}^{{tree}}")[0],
        )
        pinned_profile_files = [
            line.split(":", 1)[1]
            for line in git_lines(
                "grep",
                "-l",
                "reflaxe_go_profile=metal",
                inventory_commit,
                "--",
                "*.hxml",
            )
        ]
        self.assertEqual(
            inventory["trackedMetalProfileHxmlFiles"], len(pinned_profile_files)
        )
        self.assertEqual(
            inventory["exampleMetalProfileHxmlFiles"],
            sum(path.startswith("examples/") for path in pinned_profile_files),
        )
        self.assertEqual(
            inventory["snapshotMetalProfileHxmlFiles"],
            sum(path.startswith("test/snapshot/") for path in pinned_profile_files),
        )
        self.assertEqual(
            inventory["templateMetalProfileHxmlFiles"],
            sum(path.startswith("templates/") for path in pinned_profile_files),
        )

        hxml_files = tracked_files("*.hxml")
        profile_files = [
            path
            for path in hxml_files
            if "reflaxe_go_profile=metal" in path.read_text(encoding="utf-8")
        ]
        self.assertEqual(inventory["trackedMetalProfileHxmlFiles"], len(profile_files))
        self.assertEqual(
            inventory["exampleMetalProfileHxmlFiles"],
            sum("/examples/" in f"/{path.relative_to(ROOT).as_posix()}" for path in profile_files),
        )
        self.assertEqual(
            inventory["snapshotMetalProfileHxmlFiles"],
            sum("/test/snapshot/" in f"/{path.relative_to(ROOT).as_posix()}" for path in profile_files),
        )
        self.assertEqual(
            inventory["templateMetalProfileHxmlFiles"],
            sum("/templates/" in f"/{path.relative_to(ROOT).as_posix()}" for path in profile_files),
        )

        metadata_files = [
            path
            for path in tracked_files("*.hx")
            if "@:goMetal" in path.read_text(encoding="utf-8")
        ]
        self.assertEqual(inventory["trackedLegacyGoMetalMetadataFiles"], len(metadata_files))
        self.assertEqual(
            inventory["committedGeneratedMetalExampleTrees"],
            len(list((ROOT / "examples").glob("*/generated/metal"))),
        )
        self.assertEqual(
            inventory["metalExpectedStdoutFiles"],
            len(list((ROOT / "examples").glob("*/expected/metal*.stdout"))),
        )

        public_search = evidence["publicCodeSearch"]
        self.assertIn("lower bound", public_search["limitation"].lower())
        self.assertFalse(public_search["authorizesRemoval"])
        self.assertTrue(all("query" in row and "resultCount" in row for row in public_search["queries"]))

    def test_explicit_policy_bundle_matches_metal_generated_go(self) -> None:
        self.assertEqual(
            (METAL_FIXTURE / "Main.hx").read_bytes(),
            (EXPLICIT_FIXTURE / "Main.hx").read_bytes(),
        )
        self.assertEqual(
            (METAL_FIXTURE / "expected.stdout").read_bytes(),
            (EXPLICIT_FIXTURE / "expected.stdout").read_bytes(),
        )

        metal_hxml = (METAL_FIXTURE / "compile.hxml").read_text(encoding="utf-8")
        explicit_hxml = (EXPLICIT_FIXTURE / "compile.hxml").read_text(encoding="utf-8")
        self.assertIn("reflaxe_go_profile=metal", metal_hxml)
        self.assertNotIn("reflaxe_go_native_authority", metal_hxml)
        for define in (
            "reflaxe_go_profile=portable",
            "reflaxe_go_native_authority=explicit",
            "reflaxe_go_native_specialization=eager",
            "reflaxe_go_native_fallback=error",
            "reflaxe_go_strict_policy=on",
        ):
            self.assertIn(define, explicit_hxml)

        metal_intended = METAL_FIXTURE / "intended"
        explicit_intended = EXPLICIT_FIXTURE / "intended"
        metal_outputs = sorted(
            path.relative_to(metal_intended)
            for path in metal_intended.rglob("*")
            if path.is_file() and (path.suffix == ".go" or path.name == "go.mod")
        )
        explicit_outputs = sorted(
            path.relative_to(explicit_intended)
            for path in explicit_intended.rglob("*")
            if path.is_file() and (path.suffix == ".go" or path.name == "go.mod")
        )
        self.assertEqual(metal_outputs, explicit_outputs)
        for relative in metal_outputs:
            self.assertEqual(
                (metal_intended / relative).read_bytes(),
                (explicit_intended / relative).read_bytes(),
                relative.as_posix(),
            )

        metal_report = json.loads(
            (metal_intended / "profile_contract.json").read_text(encoding="utf-8")
        )
        explicit_report = json.loads(
            (explicit_intended / "profile_contract.json").read_text(encoding="utf-8")
        )
        for field, expected in (
            ("nativeAuthorityPolicy", "explicit"),
            ("nativeSpecializationPolicy", "eager"),
            ("nativeFallbackPolicy", "error"),
            ("strictUserBoundaries", True),
        ):
            self.assertEqual(expected, metal_report[field], field)
            self.assertEqual(expected, explicit_report[field], field)
        self.assertEqual(
            metal_report["loweringDecisions"], explicit_report["loweringDecisions"]
        )

    def test_independent_review_has_exact_model_and_artifact_provenance(self) -> None:
        provenance = json.loads(REVIEW_PROVENANCE.read_text(encoding="utf-8"))
        self.assertEqual("gpt-5.6-sol", provenance["actualModelRoute"])
        self.assertEqual("xhigh", provenance["reasoningEffort"])
        self.assertEqual("complete", provenance["status"])
        self.assertRegex(provenance["reviewedCommit"], r"^[0-9a-f]{40}$")
        self.assertEqual(sha256(REVIEW_PROMPT), provenance["promptSha256"])
        self.assertEqual(sha256(REVIEW), provenance["reviewSha256"])
        self.assertIn("selector-deprecation verdict", REVIEW.read_text(encoding="utf-8").lower())


if __name__ == "__main__":
    unittest.main()
