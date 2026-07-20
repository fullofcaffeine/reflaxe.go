#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
PACKAGE = ROOT / "package.json"
PACKAGE_LOCK = ROOT / "package-lock.json"
HAXELIB = ROOT / "haxelib.json"
RELEASE_CONFIG = ROOT / ".releaserc.json"
WORKFLOW = ROOT / ".github" / "workflows" / "ci-harness.yml"
WRAPPER = ROOT / "scripts" / "release" / "run-same-sha-release.sh"
STAGER = ROOT / "scripts" / "release" / "stage-release-metadata.py"
POLICY_CHECK = ROOT / "scripts" / "release" / "verify-release-policy.py"
RELEASE_STATUS = ROOT / "scripts" / "release" / "check-release-state.sh"
POLICY_DOC = ROOT / "docs" / "release-version-policy.md"
RELEASE_CONTRACTS = ROOT / "test" / "run-release-contracts.py"
DEVELOPMENT_VERSION = "0.0.0"


def digest(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


class ReleaseIdentityContractTest(unittest.TestCase):
    def test_source_manifests_are_development_sentinels_and_lock_matches(self) -> None:
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        lock = json.loads(PACKAGE_LOCK.read_text(encoding="utf-8"))
        haxelib = json.loads(HAXELIB.read_text(encoding="utf-8"))

        self.assertTrue(package["private"])
        self.assertEqual(package["version"], DEVELOPMENT_VERSION)
        self.assertEqual(haxelib["version"], DEVELOPMENT_VERSION)
        self.assertEqual(lock["version"], DEVELOPMENT_VERSION)
        self.assertEqual(lock["packages"][""]["version"], DEVELOPMENT_VERSION)
        self.assertIn("Development checkout", haxelib["releasenote"])

    def test_semantic_release_uses_tag_policy_without_checkout_mutators(self) -> None:
        config = json.loads(RELEASE_CONFIG.read_text(encoding="utf-8"))
        plugins = config["plugins"]
        names = [entry[0] if isinstance(entry, list) else entry for entry in plugins]

        self.assertEqual(
            names,
            ["./scripts/release/analyze-commits.mjs", "@semantic-release/github"],
        )
        self.assertEqual(plugins[0][1], {"policyPath": "release/policy.json"})
        self.assertNotIn("@semantic-release/git", names)
        self.assertNotIn("@semantic-release/changelog", names)

        github_options = plugins[1][1]
        self.assertFalse(github_options["successCommentCondition"])
        self.assertFalse(github_options["failCommentCondition"])
        self.assertFalse(github_options["releasedLabels"])
        self.assertFalse(github_options["addReleases"])

    def test_release_scripts_and_workflow_bind_the_exact_tested_sha(self) -> None:
        package = json.loads(PACKAGE.read_text(encoding="utf-8"))
        scripts = package["scripts"]
        self.assertEqual(
            scripts["release"],
            "bash scripts/release/run-same-sha-release.sh",
        )
        self.assertEqual(
            scripts["release:dry-run"],
            "bash scripts/release/run-same-sha-release.sh --dry-run",
        )
        self.assertEqual(
            scripts["release:stage-metadata"],
            "python3 scripts/release/stage-release-metadata.py",
        )
        self.assertEqual(
            scripts["release:policy"],
            "python3 scripts/release/verify-release-policy.py",
        )
        self.assertEqual(
            scripts["release:license-policy"],
            "python3 scripts/release/verify-license-policy.py --mode release",
        )

        workflow = WORKFLOW.read_text(encoding="utf-8")
        release_job = workflow.split("\n  semantic-release:", 1)[1]
        github_sha = "$" + "{{ github.sha }}"
        self.assertIn("workflow_dispatch:\n    inputs:\n      publish_release:", workflow)
        self.assertIn(
            "if: github.event_name == 'workflow_dispatch' && "
            "inputs.publish_release && github.ref == 'refs/heads/master'",
            release_job,
        )
        self.assertNotIn("github.event_name == 'push'", release_job)
        self.assertIn("ref: " + github_sha, release_job)
        self.assertIn("RELEASE_TESTED_SHA: " + github_sha, release_job)
        self.assertIn("run: npm run release:license-policy", release_job)
        self.assertIn("run: npm run release", release_job)
        self.assertNotIn("issues: write", release_job)
        self.assertNotIn("pull-requests: write", release_job)
        self.assertNotIn("continue-on-error:", release_job)
        for dependency in (
            "quality",
            "gitleaks",
            "dependency-audit",
            "go-tooling",
            "perf-go",
            "perf-apps",
        ):
            self.assertIn("- " + dependency, release_job)

        wrapper = WRAPPER.read_text(encoding="utf-8")
        for phrase in (
            "RELEASE_TESTED_SHA",
            "git rev-parse HEAD",
            "--untracked-files=no",
            "git tag --list",
            "git ls-remote",
            "semantic-release",
        ):
            self.assertIn(phrase, wrapper)

        release_status = RELEASE_STATUS.read_text(encoding="utf-8")
        self.assertIn("scripts/release/verify-release-policy.py", release_status)
        self.assertIn("release identity policy: OK", release_status)

    def test_staged_metadata_matches_version_tag_and_sha_without_source_mutation(self) -> None:
        before = {path: digest(path) for path in (PACKAGE, HAXELIB)}
        with tempfile.TemporaryDirectory() as temp_name:
            output = Path(temp_name) / "metadata"
            source_sha = "a" * 40
            proc = subprocess.run(
                [
                    "python3",
                    str(STAGER),
                    "--version",
                    "0.54.0",
                    "--source-sha",
                    source_sha,
                    "--release-note",
                    "Same tested source release",
                    "--output-dir",
                    str(output),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
            )
            self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)

            staged_package = json.loads(
                (output / "package.json").read_text(encoding="utf-8")
            )
            staged_haxelib = json.loads(
                (output / "haxelib.json").read_text(encoding="utf-8")
            )
            identity_path = output / "release-identity.json"
            identity = json.loads(identity_path.read_text(encoding="utf-8"))
            self.assertEqual(staged_package["version"], "0.54.0")
            self.assertEqual(staged_haxelib["version"], "0.54.0")
            self.assertEqual(staged_haxelib["releasenote"], "Same tested source release")
            self.assertEqual(identity["schema_version"], 1)
            self.assertEqual(identity["version"], "0.54.0")
            self.assertEqual(identity["tag"], "v0.54.0")
            self.assertEqual(identity["source_commit"], source_sha)
            self.assertEqual(identity["version_authority"], "git-tag")
            self.assertEqual(
                identity["files"]["package.json"]["sha256"],
                digest(output / "package.json"),
            )
            self.assertEqual(
                identity["files"]["haxelib.json"]["sha256"],
                digest(output / "haxelib.json"),
            )
            self.assertNotIn(str(ROOT), identity_path.read_text(encoding="utf-8"))

            second = subprocess.run(
                [
                    "python3",
                    str(STAGER),
                    "--version",
                    "0.54.0",
                    "--source-sha",
                    source_sha,
                    "--output-dir",
                    str(output),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
            )
            self.assertEqual(second.returncode, 1)
            self.assertIn("already exists", second.stderr)

        self.assertEqual(before, {path: digest(path) for path in (PACKAGE, HAXELIB)})

    def test_stager_rejects_noncanonical_version_and_source_sha(self) -> None:
        cases = [
            (["--version", "01.2.3", "--source-sha", "a" * 40], "version"),
            (
                ["--version", "0.0.0", "--source-sha", "a" * 40],
                "development sentinel",
            ),
            (["--version", "0.54.0", "--source-sha", "abc"], "source SHA"),
        ]
        for index, (arguments, diagnostic) in enumerate(cases):
            with self.subTest(arguments=arguments):
                with tempfile.TemporaryDirectory() as temp_name:
                    output = Path(temp_name) / f"metadata-{index}"
                    proc = subprocess.run(
                        [
                            "python3",
                            str(STAGER),
                            *arguments,
                            "--output-dir",
                            str(output),
                        ],
                        cwd=ROOT,
                        capture_output=True,
                        text=True,
                    )
                    self.assertEqual(proc.returncode, 1)
                    self.assertIn(diagnostic, proc.stderr)
                    self.assertFalse(output.exists())

    def test_policy_verifier_and_document_are_release_contracts(self) -> None:
        proc = subprocess.run(
            ["python3", str(POLICY_CHECK)],
            cwd=ROOT,
            capture_output=True,
            text=True,
        )
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("tag-owned version lineage: OK", proc.stdout)
        self.assertIn("same-tested-SHA workflow: OK", proc.stdout)

        doc = POLICY_DOC.read_text(encoding="utf-8")
        for phrase in (
            "# Release Version and Source-Identity Policy",
            "Git tags are the only version authority",
            "0.0.0 development sentinel",
            "breaking 0.x change advances the minor version",
            "never rewrites or commits CHANGELOG.md",
            "exact CI-tested SHA",
            "explicit manual release request",
            "npm run test:release-version-policy",
        ):
            self.assertIn(phrase, doc)

        runner = RELEASE_CONTRACTS.read_text(encoding="utf-8")
        self.assertIn("test/test_release_identity_contract.py", runner)
        self.assertIn("test/test_same_sha_release_wrapper.py", runner)
        self.assertIn("test/test_release_version_policy.mjs", runner)


if __name__ == "__main__":
    unittest.main()
