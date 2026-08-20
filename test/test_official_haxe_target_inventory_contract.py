#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import json
from pathlib import Path
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
INVENTORY_ROOT = ROOT / "test" / "official_haxe_target_inventory"
MANIFEST = INVENTORY_ROOT / "manifest.json"
RUNNER = ROOT / "test" / "run-official-haxe-target-inventory.py"
DOC = ROOT / "docs" / "official-haxe-target-inventory.md"
PACKAGE = ROOT / "package.json"
STRATEGY = ROOT / "test" / "testing-strategy.json"
CI_HARNESS = ROOT / ".github" / "workflows" / "ci-harness.yml"


def load_runner():
    spec = importlib.util.spec_from_file_location("official_haxe_target_inventory", RUNNER)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load official Haxe target inventory runner")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


class OfficialHaxeTargetInventoryContractTest(unittest.TestCase):
    def test_complete_candidate_inventory_is_pinned_and_surface_scoped(self) -> None:
        manifest = json.loads(MANIFEST.read_text(encoding="utf-8"))
        self.assertEqual(1, manifest["schemaVersion"])
        self.assertEqual("portable-compiler", manifest["productSurface"])
        self.assertEqual(
            {"top-level": 57, "unitstd": 67, "issue": 1079, "hxcpp-issue": 8},
            manifest["candidateInventory"]["counts"],
        )
        self.assertRegex(manifest["candidateInventory"]["sha256"], r"^[0-9a-f]{64}$")
        runtime_reference = manifest["activeRuntimeBaselineFile"]
        self.assertEqual("active-runtime-baseline.json", runtime_reference["path"])
        self.assertRegex(runtime_reference["sha256"], r"^[0-9a-f]{64}$")
        runtime_baseline = json.loads(
            (INVENTORY_ROOT / runtime_reference["path"]).read_text(encoding="utf-8")
        )
        self.assertEqual(650, len(runtime_baseline["records"]))
        self.assertEqual(
            "33963f192031a60c24aef5efda0d5a39a23cf82acb3a53961f5402233d606451",
            runtime_baseline["sha256"],
        )
        self.assertEqual(
            {"active": 607, "blocked": 480, "inapplicable": 124},
            manifest["classificationBaseline"]["counts"],
        )
        self.assertRegex(manifest["classificationBaseline"]["sha256"], r"^[0-9a-f]{64}$")
        self.assertIn("blockedOwners", manifest)
        self.assertIn("inapplicableRules", manifest)
        reference = manifest["blockedOwnerFiles"][0]
        self.assertEqual("blocked-issues.json", reference["path"])
        self.assertRegex(reference["sha256"], r"^[0-9a-f]{64}$")
        external = json.loads((INVENTORY_ROOT / reference["path"]).read_text(encoding="utf-8"))
        self.assertEqual(402, len(external["records"]))
        self.assertEqual(
            "06c590f6cf74abff8510bcb8fff9b3ae604067a7175d1a9d93c4aa9fa1908eb8",
            external["sourceProposalSha256"],
        )
        zero_reference = manifest["reviewedZeroRuntimeOwnerFiles"][0]
        zero_runtime = json.loads(
            (INVENTORY_ROOT / zero_reference["path"]).read_text(encoding="utf-8")
        )
        self.assertEqual(117, len(zero_runtime["records"]))
        self.assertEqual(
            {"blocked": 7, "inapplicable": 110},
            {
                status: sum(item["status"] == status for item in zero_runtime["records"])
                for status in ("blocked", "inapplicable")
            },
        )
        compile_only = {
            item["path"]
            for item in zero_runtime["records"]
            if item["status"] == "blocked"
        }
        self.assertEqual(
            {
                "unit/issues/Issue4359.hx",
                "unit/issues/Issue4535.hx",
                "unit/issues/Issue6059.hx",
                "unit/issues/Issue6772.hx",
                "unit/issues/Issue7332.hx",
                "unit/issues/Issue7389.hx",
                "unit/issues/Issue9309.hx",
            },
            compile_only,
        )

    def test_upstream_license_provenance_does_not_guess_one_license_for_tests(self) -> None:
        manifest = json.loads(MANIFEST.read_text(encoding="utf-8"))
        haxe = manifest["upstream"]["haxe"]
        self.assertEqual("mixed", haxe["license"])
        self.assertIn("tests/unit", haxe["licenseScopeNote"])
        self.assertIn("does not assign", haxe["licenseScopeNote"])
        evidence = {item["path"]: item["sha256"] for item in haxe["licenseEvidence"]}
        self.assertEqual(
            {
                "README.md": "708b5a3386662e375c8bf90493687a764e9b7d251c500db854c8f76c570685ae",
                "extra/LICENSE.txt": "f84691d619932ebfd4fa3568f8311f87ed4bf12e747e9aaa619a92cb1d2d359d",
            },
            evidence,
        )

    def test_synthetic_addition_and_runtime_drift_fail_closed(self) -> None:
        module = load_runner()
        left, right = module.split_shard(
            {"id": "sample", "owners": [{"path": str(i)} for i in range(5)]}
        )
        self.assertEqual(["0", "1"], [item["path"] for item in left["owners"]])
        self.assertEqual(["2", "3", "4"], [item["path"] for item in right["owners"]])
        helpers = "\n".join(
            f"\tfunction {name}() {{}}" for name in module.TEST_BASE_HELPERS
        )
        synthetic = (
            "package unit;\n@:keepSub\nclass Test {\n\tpublic function new() {}\n"
            + helpers
            + "\n}\n"
        )
        adapted = module.adapt_test_base_text(synthetic, {"eq"})
        self.assertIn("@:keepSub", adapted)
        self.assertIn("public function new", adapted)
        self.assertIn("function eq(", adapted)
        self.assertNotIn("function aeq(", adapted)
        owner = module.adapt_owner_text(
            "package unit.issues;\nclass Issue extends unit.Test { function test() {} }\n"
        )
        self.assertIn("import OfficialInventoryTestBase;", owner)
        self.assertIn("extends OfficialInventoryTestBase", owner)
        self.assertIn("implements utest.ITest", owner)
        self.assertNotIn("extends unit.Test", owner)
        self.assertIn(
            "import OfficialInventoryTestBase;",
            module.adapt_owner_text("\ufeffpackage unit;\nclass TestA extends Test {}\n"),
        )
        with self.assertRaises(module.InventoryError):
            module.adapt_test_base_text("package unit;\nclass Test {}\n", {"eq"})
        with tempfile.TemporaryDirectory(prefix="official-inventory-contract-") as raw:
            root = Path(raw)
            (root / "unit").mkdir()
            source = root / "unit" / "TestAlpha.hx"
            source.write_text("class TestAlpha {}\n", encoding="utf-8")
            records = module.discover_candidate_inventory(root)
            expected = module.canonical_inventory_sha256(records)
            self.assertEqual(1, len(records))
            (root / "unit" / "TestBeta.hx").write_text(
                "class TestBeta {}\n", encoding="utf-8"
            )
            changed = module.discover_candidate_inventory(root)
            self.assertNotEqual(expected, module.canonical_inventory_sha256(changed))

        baseline = {
            "records": [{"id": "unit.TestAlpha.test", "assertions": 2}],
            "sha256": module.canonical_runtime_sha256(
                [{"id": "unit.TestAlpha.test", "assertions": 2}]
            ),
        }
        module.validate_runtime_baseline(
            baseline, [{"id": "unit.TestAlpha.test", "status": "pass", "assertions": 2}]
        )
        for drift in (
            [],
            [{"id": "unit.TestAlpha.renamed", "status": "pass", "assertions": 2}],
            [{"id": "unit.TestAlpha.test", "status": "pass", "assertions": 1}],
        ):
            with self.assertRaises(module.InventoryError):
                module.validate_runtime_baseline(baseline, drift)
        duplicate = [
            {"id": "unit.TestAlpha.test", "status": "pass", "assertions": 2},
            {"id": "unit.TestAlpha.test", "status": "pass", "assertions": 2},
        ]
        duplicate_baseline = {
            "records": module.canonical_runtime_records(duplicate),
            "sha256": module.canonical_runtime_sha256(duplicate),
        }
        with self.assertRaises(module.InventoryError):
            module.validate_runtime_baseline(duplicate_baseline, duplicate)

        classification = [
            {"family": "top-level", "path": "unit/TestAlpha.hx", "status": "active"},
            {"family": "top-level", "path": "unit/TestBeta.hx", "status": "blocked"},
        ]
        classification_baseline = {
            "counts": {"active": 1, "blocked": 1},
            "sha256": module.canonical_classification_sha256(classification),
        }
        module.validate_classification_baseline(classification_baseline, classification)
        with self.assertRaises(module.InventoryError):
            module.validate_classification_baseline(
                classification_baseline,
                [{**classification[0], "status": "inapplicable"}, classification[1]],
            )
        workspace = Path("/tmp/haxe-go-official-inventory-random-suffix")
        sandbox = workspace / "sandbox"
        redacted = module.redact(
            f"{sandbox}/generated/main.go {ROOT}/package.json",
            [ROOT, workspace, sandbox],
        )
        self.assertEqual(
            "<REDACTED:installed-package>/generated/main.go "
            "<REDACTED:repository>/package.json",
            redacted,
        )

    def test_extended_lane_is_separate_from_required_smoke(self) -> None:
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        self.assertEqual(
            "python3 test/run-official-haxe-target-inventory.py",
            package["scripts"]["test:official-haxe-inventory"],
        )
        self.assertNotEqual(
            package["scripts"]["test:official-haxe-smoke"],
            package["scripts"]["test:official-haxe-inventory"],
        )
        self.assertIn(
            "python3 test/test_official_haxe_target_inventory_contract.py",
            package["scripts"]["test:strategy"],
        )
        text = DOC.read_text(encoding="utf-8")
        self.assertIn("extended", text.lower())
        self.assertIn("portable compiler scorecard", text)
        self.assertIn("does not admit", text)

        strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))
        owners = {item["id"]: item for item in strategy["testOwners"]}
        inventory = owners["official-haxe-target-inventory"]
        self.assertEqual(["portable-compiler"], inventory["surfaces"])
        self.assertEqual(["R4", "R5"], inventory["rings"])
        self.assertFalse(inventory["alwaysRun"])
        portable = next(
            item for item in strategy["productSurfaces"] if item["id"] == "portable-compiler"
        )
        self.assertIn("official-haxe-target-inventory", portable["verticalOwners"])
        self.assertIn("official-haxe-target-inventory", portable["realRuntimeOrSystemOwners"])

    def test_supported_cold_ci_lanes_publish_exact_inventory_evidence(self) -> None:
        workflow = CI_HARNESS.read_text(encoding="utf-8")
        self.assertIn("official-haxe-target-inventory:", workflow)
        self.assertIn('go: ["1.25.13", "1.26.6"]', workflow)
        self.assertIn("npm run test:official-haxe-inventory", workflow)
        self.assertIn("npm run test:official-haxe-inventory -- --require-clean-source", workflow)
        self.assertIn("Verify clean exact source", workflow)
        self.assertIn(
            'rm -rf "haxe-${HAXE_VERSION}-linux64" "neko-2.4.0-linux64"',
            workflow,
        )
        self.assertIn("official-haxe-target-inventory-${{ matrix.go }}", workflow)
        self.assertIn("- official-haxe-target-inventory", workflow)

    def test_runner_uses_the_pinned_utest_failure_contract(self) -> None:
        module = load_runner()
        runner = RUNNER.read_text(encoding="utf-8")
        self.assertIn('"UTEST_FAILURE_THROW"', runner)
        main = (INVENTORY_ROOT / "src" / "OfficialInventoryMain.hx").read_text(
            encoding="utf-8"
        )
        self.assertNotIn("new Runner()", main)
        self.assertIn("__initializeUtest__", main)
        self.assertNotIn("cast testCase", main)
        self.assertNotIn("UnitBuilder.generateSpec", main)
        self.assertTrue((INVENTORY_ROOT / "src" / "OfficialInventoryUnitStd.hx").is_file())
        self.assertIn("--resource", runner)
        self.assertIn("claimEvidence", runner)
        self.assertIn("diagnosticSelection", runner)
        self.assertIn("--require-clean-source", runner)
        self.assertTrue(module.is_claim_evidence(False, False, False, False))
        self.assertFalse(module.is_claim_evidence(True, False, False, False))
        self.assertFalse(module.is_claim_evidence(False, True, False, False))
        self.assertFalse(module.is_claim_evidence(False, False, True, False))
        self.assertFalse(module.is_claim_evidence(False, False, False, True))

    def test_zero_runtime_runnable_owner_fails_instead_of_becoming_inapplicable(self) -> None:
        module = load_runner()
        owner = {
            "family": "issue",
            "path": "unit/issues/IssueExample.hx",
            "runtimeClass": "unit.issues.IssueExample",
        }
        with self.assertRaisesRegex(
            module.InventoryError, "zero runtime records.*explicitly classified"
        ):
            module.classify_executed_owner(owner, [])


if __name__ == "__main__":
    unittest.main()
