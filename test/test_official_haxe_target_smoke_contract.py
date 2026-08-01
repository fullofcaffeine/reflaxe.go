#!/usr/bin/env python3

from __future__ import annotations

import json
import importlib.util
from pathlib import Path
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
SMOKE_ROOT = ROOT / "test" / "official_haxe_target_smoke"
MANIFEST = SMOKE_ROOT / "manifest.json"
RUNNER = ROOT / "test" / "run-official-haxe-target-smoke.py"
PACKAGE = ROOT / "package.json"
RUN_CI = ROOT / "test" / "run-ci.py"
DOC = ROOT / "docs" / "official-haxe-target-smoke.md"
CI_HARNESS = ROOT / ".github" / "workflows" / "ci-harness.yml"
CI_QUALITY = ROOT / ".github" / "workflows" / "ci-quality.yml"

HAXE_COMMIT = "e0b355c6be312c1b17382603f018cf52522ec651"
UTEST_COMMIT = "a94f8812e8786f2b5fec52ce9f26927591d26327"


def load_runner():
    spec = importlib.util.spec_from_file_location("official_haxe_target_smoke", RUNNER)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load official Haxe target smoke runner")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


class OfficialHaxeTargetSmokeContractTest(unittest.TestCase):
    def test_manifest_pins_provenance_and_three_active_official_families(self) -> None:
        manifest = json.loads(MANIFEST.read_text(encoding="utf-8"))
        self.assertEqual(1, manifest["schemaVersion"])
        self.assertEqual(HAXE_COMMIT, manifest["upstream"]["haxe"]["commit"])
        self.assertEqual(UTEST_COMMIT, manifest["upstream"]["utest"]["commit"])
        self.assertEqual("MIT", manifest["upstream"]["haxe"]["license"])
        self.assertEqual("MIT", manifest["upstream"]["utest"]["license"])

        records = manifest["activeSmokeRecords"]
        self.assertEqual({"top-level", "unitstd", "issue"}, {record["family"] for record in records})
        for record in records:
            self.assertEqual("portable-compiler", record["productSurface"])
            self.assertEqual("official-body-unmodified", record["provenance"])
            self.assertEqual("target-runtime", record["observer"])
            self.assertTrue(record["upstreamPath"])
            self.assertRegex(record["sha256"], r"^[0-9a-f]{64}$")
            self.assertTrue(record["expectedActiveTests"])
            self.assertGreater(record["minimumAssertions"], 0)

        harness = manifest["harness"]
        self.assertIn("pinned utest", harness["assertionAuthority"])
        self.assertIn("unit.UnitBuilder.read", harness["unitstdExpansionAuthority"])
        self.assertIn("temporary", harness["officialClassExpansionAuthority"])
        self.assertIn("not portable evidence", harness["observerBoundary"])

    def test_runner_and_required_ci_entrypoint_are_wired(self) -> None:
        runner = RUNNER.read_text(encoding="utf-8")
        main = (SMOKE_ROOT / "src" / "OfficialTargetSmokeMain.hx").read_text(
            encoding="utf-8"
        )
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        run_ci = RUN_CI.read_text(encoding="utf-8")
        self.assertEqual(
            "python3 test/run-official-haxe-target-smoke.py --verify-failure-propagation",
            package["scripts"]["test:official-haxe-smoke"],
        )
        self.assertIn("build_official_haxe_target_smoke_command", run_ci)
        self.assertIn("test:official-haxe-smoke", run_ci)
        self.assertIn("--verify-failure-propagation", runner)
        self.assertIn("adapt_official_class_source", runner)
        self.assertIn("official_haxe_smoke_unitstd_path", runner)
        self.assertIn("UTEST_FAILURE_THROW", runner)
        self.assertIn("official_haxe_smoke_timeout_failure", runner)
        self.assertIn("official_haxe_smoke_timeout_failure", main)
        for failure in ("assertion", "go-build", "runtime", "timeout", "missing"):
            self.assertIn(failure, runner)
        self.assertIn("def verify_missing_selected_source_control(", runner)
        self.assertIn("missing_selected_source.unlink()", runner)
        self.assertIn('"stage": "selected-source"', runner)

    def test_haxe_compile_is_confined_to_the_isolated_haxelib_repository(self) -> None:
        runner = RUNNER.read_text(encoding="utf-8")
        self.assertIn("def verify_compile_package_resolution(", runner)
        self.assertIn("app = sandbox / name", runner)
        self.assertIn("cwd=app", runner)
        self.assertIn(
            "verify_compile_package_resolution(app, sandbox, environment, timeout)",
            runner,
        )
        self.assertIn('"packageResolution": {', runner)
        self.assertIn('"sourceCheckoutExcluded": True', runner)

    def test_every_uploaded_artifact_rejects_ephemeral_absolute_paths(self) -> None:
        module = load_runner()
        with tempfile.TemporaryDirectory(prefix="haxe-go-artifact-contract-") as raw:
            root = Path(raw)
            artifact = root / "artifact"
            artifact.mkdir()
            forbidden = root / "ephemeral-workspace"
            (artifact / "warning.stderr").write_text(
                f"warning from {forbidden}/generated/Main.go\n",
                encoding="utf-8",
            )
            with self.assertRaises(module.SmokeError) as caught:
                module.verify_artifact_path_confinement(artifact, [forbidden])
            self.assertFalse(artifact.exists(), "a rejected artifact must not remain uploadable")
        self.assertEqual("confinement", caught.exception.stage)

    def test_result_binds_the_installed_package_archive_by_content(self) -> None:
        runner = RUNNER.read_text(encoding="utf-8")
        self.assertIn('"packageArtifact": package_artifact', runner)
        self.assertIn('"sha256": sha256_file(archive)', runner)
        self.assertIn('"bytes": archive.stat().st_size', runner)

    def test_upstream_bodies_and_security_sensitive_tls_fixture_are_not_vendored(self) -> None:
        upstream_haxe = list(SMOKE_ROOT.rglob("*.unit.hx"))
        upstream_issue = list(SMOKE_ROOT.rglob("Issue*.hx"))
        self.assertEqual([], upstream_haxe)
        self.assertEqual([], upstream_issue)
        self.assertFalse((SMOKE_ROOT / "Ssl.unit.hx").exists())

    def test_exact_sha_artifacts_are_uploaded_from_each_required_runtime_lane(self) -> None:
        for workflow in (CI_HARNESS, CI_QUALITY):
            text = workflow.read_text(encoding="utf-8")
            self.assertIn(".cache/official-haxe-target-smoke/artifacts", text)
            self.assertIn("official-haxe-target-smoke", text)
            self.assertIn("if: always()", text)

    def test_documentation_keeps_the_claim_narrow(self) -> None:
        text = DOC.read_text(encoding="utf-8")
        self.assertIn("official-suite smoke", text)
        self.assertIn(HAXE_COMMIT, text)
        self.assertIn(UTEST_COMMIT, text)
        self.assertIn("not the complete applicable", text)
        self.assertIn("portable compiler scorecard", text)
        self.assertIn("active runtime", text)

    def test_documentation_explains_the_contributor_path_before_internals(self) -> None:
        text = DOC.read_text(encoding="utf-8")
        opening = "\n".join(text.splitlines()[:70])
        self.assertIn("## Run it", opening)
        self.assertIn("npm run test:official-haxe-smoke", opening)
        self.assertIn("In plain terms", opening)
        self.assertIn("Active inventory means", text)
        self.assertIn("haxe_go-vfp.12.10", text)


if __name__ == "__main__":
    unittest.main()
