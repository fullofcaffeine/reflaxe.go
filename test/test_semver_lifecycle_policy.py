#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
POLICY = ROOT / "release" / "policy.json"
POLICY_DOC = ROOT / "docs" / "semver-lifecycle-policy.md"

STABLE_REQUIREMENTS = {
    "public-contract": ("complete", "haxe_go-vfp.6.3"),
    "stable-scope-and-support": ("pending", "haxe_go-vfp.12.7"),
    "release-readiness": ("pending", "haxe_go-vfp.6.5"),
    "published-beta-baseline": ("pending", "haxe_go-vfp.4.8"),
    "consumer-upgrade-rehearsal": ("pending", "haxe_go-vfp.12.4"),
    "security-and-licensing": ("pending", "haxe_go-vfp.6.5"),
    "independent-stable-review": ("pending", "haxe_go-vfp.12.7"),
}


class SemverLifecyclePolicyTest(unittest.TestCase):
    def load_policy(self) -> dict[str, object]:
        return json.loads(POLICY.read_text(encoding="utf-8"))

    def test_machine_policy_keeps_beta_channels_and_breaking_rules_explicit(self) -> None:
        policy = self.load_policy()

        self.assertEqual(policy["schemaVersion"], 1)
        self.assertEqual(
            policy["normalChannel"],
            {
                "branch": "master",
                "tagFormat": "v${version}",
                "prerelease": False,
                "productMaturityAuthority": "docs/compatibility-support-manifest.json",
            },
        )
        self.assertEqual(
            policy["releaseLines"]["0"],
            {"stage": "initial-development", "breakingBump": "minor"},
        )

        deprecation = policy["deprecationPolicy"]
        self.assertEqual(deprecation["majorZero"]["noticeRelease"], "minor")
        self.assertEqual(deprecation["majorZero"]["minimumFunctionalMinorReleasesAfterNotice"], 1)
        self.assertEqual(deprecation["majorZero"]["earliestRemovalMinorOffset"], 2)
        self.assertEqual(deprecation["stable"]["removalRelease"], "next-major")

        experimental = policy["experimentalPolicy"]
        self.assertEqual(experimental["compatibilityPromise"], "excluded-unless-admitted")
        self.assertEqual(experimental["minimumChangeRelease"], "minor")
        self.assertEqual(
            experimental["stableMinorException"],
            "requires-explicit-surface-proof",
        )

    def test_stable_one_is_blocked_by_objective_requirements_and_human_approval(self) -> None:
        line = self.load_policy()["releaseLines"]["1"]

        self.assertEqual(line["stage"], "stable")
        self.assertIsNone(line["approval"])
        requirements = {
            entry["id"]: (entry["status"], entry["record"])
            for entry in line["requirements"]
        }
        self.assertEqual(requirements, STABLE_REQUIREMENTS)
        self.assertTrue(any(status == "pending" for status, _record in requirements.values()))

    def test_human_policy_covers_every_requested_decision(self) -> None:
        document = POLICY_DOC.read_text(encoding="utf-8")

        for heading in (
            "## What beta means",
            "## Change classes during 0.x",
            "## Experimental surfaces",
            "## Deprecation windows",
            "## Profiles, metadata, and generated output",
            "## Release channels",
            "## Stable 1.x admission",
            "## Who decides",
            "## Sibling lessons",
        ):
            self.assertIn(heading, document)

        for phrase in (
            "notice in minor N",
            "functional throughout minor N+1",
            "removal no earlier than minor N+2",
            "explicit human approval",
            "does not automatically authorize `1.0.0`",
            "not automatically a SemVer prerelease",
            "requires explicit surface proof",
            "`fix(experimental):`",
            "`fix(deprecation):`",
            "only the target approval field",
        ):
            self.assertIn(phrase, document)

    def test_release_analyzer_and_docs_consume_the_policy(self) -> None:
        config = json.loads((ROOT / ".releaserc.json").read_text(encoding="utf-8"))
        analyzer_options = config["plugins"][0][1]
        self.assertEqual(analyzer_options, {"policyPath": "release/policy.json"})

        analyzer = (ROOT / "scripts" / "release" / "analyze-commits.mjs").read_text(
            encoding="utf-8"
        )
        self.assertIn("loadReleasePolicy", analyzer)
        self.assertIn("validateReleasePolicy", analyzer)
        self.assertNotIn("pluginConfig?.approvedStableMajors", analyzer)

        for relative in (
            "docs/public-contract.md",
            "docs/release-version-policy.md",
            "docs/release-readiness-checklist.md",
            "docs/index.md",
        ):
            self.assertIn("semver-lifecycle-policy.md", (ROOT / relative).read_text(encoding="utf-8"))

        release_runner = (ROOT / "test" / "run-release-contracts.py").read_text(encoding="utf-8")
        evidence_builder = (ROOT / "scripts" / "review" / "build_gpt56_evidence.py").read_text(
            encoding="utf-8"
        )
        self.assertIn("test/test_semver_lifecycle_policy.py", release_runner)
        self.assertIn('"release/policy.json"', evidence_builder)
        self.assertIn('"docs/semver-lifecycle-policy.md"', evidence_builder)


if __name__ == "__main__":
    unittest.main()
