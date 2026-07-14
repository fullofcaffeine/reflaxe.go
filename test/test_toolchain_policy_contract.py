#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
POLICY_JSON = REPO_ROOT / "docs" / "toolchain-policy.json"
POLICY_DOC = REPO_ROOT / "docs" / "toolchain-policy.md"
RELEASE_STATUS = REPO_ROOT / "scripts" / "release" / "check-release-state.sh"
WORKFLOW_DIR = REPO_ROOT / ".github" / "workflows"
HAXE_ACTION_REF = (
    "uses: krdlab/setup-haxe@d93667502be3b4f31a94a3308a74388f2e178a8d"
    " # v2.1.0"
)


class ToolchainPolicyContractTest(unittest.TestCase):
    def load_policy(self) -> dict[str, object]:
        return json.loads(POLICY_JSON.read_text(encoding="utf-8"))

    def test_machine_policy_names_language_floor_and_supported_build_lines(self) -> None:
        policy = self.load_policy()

        self.assertEqual(policy["schema_version"], 1)
        self.assertEqual(policy["haxe"]["supported_versions"], ["4.3.7"])
        self.assertEqual(policy["haxe"]["recommended_version"], "4.3.7")
        self.assertEqual(policy["go"]["generated_language_floor"], "1.22")
        self.assertEqual(policy["go"]["supported_build_lines"], ["1.25", "1.26"])
        self.assertEqual(policy["go"]["recommended_build_line"], "1.26")
        self.assertEqual(policy["go"]["ci_selectors"], ["1.25.x", "1.26.x"])
        self.assertEqual(policy["node"]["supported_tooling_lines"], ["24"])
        self.assertEqual(policy["node"]["recommended_tooling_line"], "24")
        self.assertTrue(policy["release"]["require_exact_patch_evidence"])

    def test_quality_matrix_covers_both_supported_go_lines(self) -> None:
        workflow = (WORKFLOW_DIR / "ci-quality.yml").read_text(encoding="utf-8")
        matrix_lines = re.findall(r'^\s+go: "([0-9]+\.[0-9]+\.x)"$', workflow, flags=re.MULTILINE)

        self.assertEqual(matrix_lines, ["1.25.x", "1.26.x", "1.26.x"])
        self.assertIn('NODE_VERSION: "24"', workflow)
        self.assertIn('HAXE_VERSION: "4.3.7"', workflow)
        self.assertIn(HAXE_ACTION_REF, workflow)
        self.assertIn("haxe-version: ${{ env.HAXE_VERSION }}", workflow)
        self.assertNotIn("brew install haxe", workflow)

    def test_other_product_workflows_use_recommended_toolchains(self) -> None:
        harness = (WORKFLOW_DIR / "ci-harness.yml").read_text(encoding="utf-8")
        security = (WORKFLOW_DIR / "security-static-analysis.yml").read_text(encoding="utf-8")
        examples = (WORKFLOW_DIR / "examples-artifacts.yml").read_text(encoding="utf-8")

        for workflow in (harness, security):
            self.assertIn('NODE_VERSION: "24"', workflow)
            self.assertIn('GO_VERSION: "1.26.x"', workflow)
        self.assertIn('GO_VERSION: "1.26.x"', examples)
        self.assertIn("go-version: ${{ env.GO_VERSION }}", examples)

        for path in sorted(WORKFLOW_DIR.glob("*.yml")):
            text = path.read_text(encoding="utf-8")
            with self.subTest(workflow=path.name):
                self.assertNotIn('NODE_VERSION: "20"', text)
                self.assertNotIn('go-version: "1.22.x"', text)
                self.assertNotIn('go-version: "1.23.x"', text)
                self.assertNotIn('GO_VERSION: "1.23.x"', text)

    def test_generated_go_language_floor_is_deliberate_and_unchanged(self) -> None:
        policy = self.load_policy()
        compiler = (REPO_ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx").read_text(encoding="utf-8")
        iterator = (REPO_ROOT / "src" / "reflaxe" / "go" / "GoOutputIterator.hx").read_text(encoding="utf-8")
        representative_go_mod = (REPO_ROOT / "test" / "snapshot" / "core" / "hello_trace" / "intended" / "go.mod").read_text(
            encoding="utf-8"
        )

        self.assertEqual(policy["go"]["generated_language_floor"], "1.22")
        self.assertNotEqual(policy["go"]["generated_language_floor"], policy["go"]["recommended_build_line"])
        self.assertIn('"go 1.22"', compiler)
        self.assertIn('"go 1.22"', iterator)
        self.assertIn("\ngo 1.22\n", representative_go_mod)

    def test_goextern_go123_fixture_is_non_production_compatibility_evidence(self) -> None:
        policy = self.load_policy()
        probes = {probe["id"]: probe for probe in policy["compatibility_probes"]}
        probe = probes["goextern-fixtures-go1.23"]
        goextern_doc = (REPO_ROOT / "docs" / "goextern.md").read_text(encoding="utf-8")

        self.assertEqual(probe["toolchain_line"], "1.23")
        self.assertFalse(probe["production_supported"])
        self.assertFalse(probe["security_supported"])
        self.assertIn("non-production compatibility fixture", goextern_doc)
        self.assertIn("does not establish security support", goextern_doc)

    def test_release_readiness_consumes_the_toolchain_policy(self) -> None:
        policy_doc = POLICY_DOC.read_text(encoding="utf-8")
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        release_status = RELEASE_STATUS.read_text(encoding="utf-8")

        self.assertIn("# Supported Toolchain Policy", policy_doc)
        self.assertIn("https://go.dev/doc/devel/release", policy_doc)
        self.assertIn("https://github.com/nodejs/Release", policy_doc)
        self.assertIn("https://haxe.org/download/list/", policy_doc)
        self.assertIn("[supported toolchain policy](toolchain-policy.md)", checklist)
        self.assertIn("exact resolved patch versions", checklist)
        self.assertIn("does not make an old compatibility fixture production-supported", checklist)
        self.assertIn('require_file "docs/toolchain-policy.json"', release_status)
        self.assertIn("toolchain policy wiring", release_status)


if __name__ == "__main__":
    unittest.main()
