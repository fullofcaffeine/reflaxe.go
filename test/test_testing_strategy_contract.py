#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
import subprocess
import sys
import unittest


ROOT = Path(__file__).resolve().parent.parent
STRATEGY = ROOT / "test" / "testing-strategy.json"
EXAMPLES = ROOT / "examples" / "qa-manifest.json"
DOC = ROOT / "docs" / "testing-strategy.md"
PACKAGE = ROOT / "package.json"


class TestingStrategyContractTest(unittest.TestCase):
    def test_portfolio_counts_active_harness_owners_instead_of_raw_files(self) -> None:
        strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))
        inventory = strategy["portfolio"]["currentInventory"]

        def listed(*args: str) -> int:
            process = subprocess.run(
                [sys.executable, *args],
                cwd=ROOT,
                capture_output=True,
                text=True,
                check=True,
            )
            return len([line for line in process.stdout.splitlines() if line.strip()])

        stdlib = json.loads(
            (ROOT / "test" / "portable_stdlib_inventory.json").read_text(encoding="utf-8")
        )
        self.assertEqual(listed("test/run-snapshots.py", "--list"), inventory["snapshotOwners"])
        self.assertEqual(
            listed("test/run-semantic-diff.py", "--list"),
            inventory["portableSemanticDiffOwners"],
        )
        self.assertEqual(
            listed("test/run-semantic-diff.py", "--suite", "lanes", "--list"),
            inventory["laneSemanticDiffOwners"],
        )
        self.assertEqual(listed("test/run-examples.py", "--list"), inventory["exampleProfileOwners"])
        self.assertEqual(
            len([module for module in stdlib["modules"] if module["in_strict_sweep"]]),
            inventory["strictStdlibModules"],
        )
        self.assertEqual(
            len([module for module in stdlib["modules"] if module["portable_eligible"]]),
            inventory["fullPortableEligibleStdlibModules"],
        )

    def test_product_surfaces_are_independent_and_complete(self) -> None:
        strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))
        example_manifest = json.loads(EXAMPLES.read_text(encoding="utf-8"))
        surfaces = {item["id"]: item for item in strategy["productSurfaces"]}
        self.assertEqual(
            {
                "portable-compiler",
                "go-native-metal",
                "runtime-stdlib",
                "diagnostics-tooling",
                "package-downstream-examples",
            },
            set(surfaces),
        )
        for surface in surfaces.values():
            self.assertTrue(surface["claims"])
            self.assertTrue(surface["focusedOwners"])
            self.assertTrue(surface["verticalOwners"])
            self.assertTrue(surface["fullBackstopCommand"])
            self.assertTrue(surface["releaseCommand"])
            self.assertIn("status", surface)
            self.assertIn("residualRisks", surface)

        self.assertTrue(surfaces["portable-compiler"]["officialHaxeTargetSuiteApplies"])
        metal_qualification = surfaces["portable-compiler"]["profileQualifications"]["metal"]
        self.assertIn("preset-invariance", metal_qualification)
        self.assertIn("not Go-native", metal_qualification)
        for surface_id, surface in surfaces.items():
            if surface_id != "portable-compiler":
                self.assertFalse(surface["officialHaxeTargetSuiteApplies"])

        expected_examples = {
            surface_id: {
                example["id"]
                for example in example_manifest["examples"]
                if any(
                    surface_id in lane["productSurfaces"]
                    for lane in example["lanes"].values()
                )
            }
            for surface_id in surfaces
        }
        for surface_id, surface in surfaces.items():
            self.assertEqual(expected_examples[surface_id], set(surface["examples"]))

        for example in example_manifest["examples"]:
            if not example["claimBearing"]:
                continue
            for profile in example["profiles"]:
                for lane_id, lane in example["lanes"].items():
                    for surface_id in lane["productSurfaces"]:
                        self.assertIn(
                            profile,
                            surfaces[surface_id]["supportedProfiles"],
                            f"{example['id']}/{profile}/{lane_id} claims unsupported {surface_id}",
                        )
                        self.assertIn(
                            profile,
                            surfaces[surface_id]["testedProfiles"],
                            f"{example['id']}/{profile}/{lane_id} claims untested {surface_id}",
                        )

    def test_feedback_rings_and_review_policy_are_explicit(self) -> None:
        strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))
        self.assertEqual(["R0", "R1", "R2", "R3", "R4", "R5"], [ring["id"] for ring in strategy["feedbackRings"]])
        self.assertEqual("observation", strategy["affectedSelection"]["mode"])
        self.assertTrue(strategy["affectedSelection"]["alwaysRunOwners"])
        self.assertTrue(strategy["affectedSelection"]["unknownPathsExpandToFull"])
        self.assertTrue(strategy["highRiskReview"]["distinctFromImplementation"])
        self.assertTrue(strategy["representativeWorkflow"]["redState"]["command"])
        self.assertTrue(strategy["representativeWorkflow"]["oracle"]["independent"])
        self.assertTrue(strategy["representativeWorkflow"]["tracerBullet"])

    def test_weekly_ring_names_the_harness_subset_instead_of_unscheduled_quality(self) -> None:
        strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))
        r4 = next(ring for ring in strategy["feedbackRings"] if ring["id"] == "R4")
        self.assertIn("weekly CI Harness", r4["command"])
        self.assertIn("Quality runs on pull requests and master pushes", r4["coldPolicy"])
        doc = (ROOT / "docs" / "testing-strategy.md").read_text(encoding="utf-8")
        self.assertIn("weekly CI Harness subset", doc)

    def test_required_strategy_command_runs_example_claim_state_regressions(self) -> None:
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        self.assertIn(
            "python3 test/test_generated_output_telemetry.py",
            package["scripts"]["test:strategy"],
        )

    def test_examples_declare_tier_surface_profiles_and_real_execution(self) -> None:
        manifest = json.loads(EXAMPLES.read_text(encoding="utf-8"))
        records = {item["id"]: item for item in manifest["examples"]}
        maintained = {
            path.name
            for path in (ROOT / "examples").iterdir()
            if path.is_dir() and (path / "README.md").exists()
        }
        self.assertEqual(maintained, set(records))
        for record in records.values():
            self.assertIn(record["tier"], {"flagship-application", "capability-showcase", "compile-only-snippet"})
            self.assertTrue(record["profiles"])
            self.assertTrue(record["testCommand"])
            self.assertTrue(record["oracle"]["source"])
            self.assertEqual({"default", "ci"}, set(record["lanes"]))
            for lane in record["lanes"].values():
                self.assertTrue(lane["productSurfaces"])
                self.assertTrue(lane["evidenceModes"])
            if record["claimBearing"]:
                self.assertEqual(
                    ["haxe-custom-backend", "gofmt", "go-test", "go-run", "expected-output"],
                    record["execution"],
                )

    def test_strategy_document_explains_no_evidence_laundering(self) -> None:
        text = DOC.read_text(encoding="utf-8")
        self.assertIn("One green surface cannot advance another surface", text)
        self.assertIn("red for the intended reason", text)
        self.assertIn("independent oracle", text)
        self.assertIn("tracer bullet", text)
        self.assertIn("R0", text)
        self.assertIn("R5", text)


if __name__ == "__main__":
    unittest.main()
