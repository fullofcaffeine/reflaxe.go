#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
WRAPPER = ROOT / "scripts" / "release" / "run-same-sha-release.sh"
LICENSE_POLICY_VERIFIER = ROOT / "scripts" / "release" / "verify-license-policy.py"


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def scope_digest(policy: dict[str, object]) -> str:
    scope = {
        key: policy[key]
        for key in (
            "shippedSourcePatterns",
            "components",
            "generatedOutputClasses",
            "releasePackage",
        )
    }
    return hashlib.sha256(
        json.dumps(scope, separators=(",", ":"), sort_keys=True).encode("utf-8")
    ).hexdigest()


class SameShaReleaseWrapperTest(unittest.TestCase):
    def setUp(self) -> None:
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.repo = Path(self.temp.name)
        subprocess.run(
            ["git", "init", "-q", "-b", "master"],
            cwd=self.repo,
            check=True,
        )
        subprocess.run(
            ["git", "config", "user.name", "Release Contract"],
            cwd=self.repo,
            check=True,
        )
        subprocess.run(
            ["git", "config", "user.email", "release@example.test"],
            cwd=self.repo,
            check=True,
        )
        (self.repo / "scripts" / "release").mkdir(parents=True)
        shutil.copy2(WRAPPER, self.repo / "scripts" / "release" / WRAPPER.name)
        shutil.copy2(
            LICENSE_POLICY_VERIFIER,
            self.repo / "scripts" / "release" / LICENSE_POLICY_VERIFIER.name,
        )
        (self.repo / "tracked.txt").write_text("baseline\n", encoding="utf-8")
        (self.repo / "LICENSE").write_text("fixture license\n", encoding="utf-8")
        (self.repo / "fake-release.sh").write_text(
            "#!/usr/bin/env bash\nexit 0\n",
            encoding="utf-8",
        )
        (self.repo / "fake-release.sh").chmod(0o755)
        (self.repo / "fake-builder.py").write_text(
            """#!/usr/bin/env python3
import json, os, pathlib, sys
args = sys.argv[1:]
out = pathlib.Path(args[args.index('--output-dir') + 1])
tag = args[args.index('--tag') + 1]
sha = args[args.index('--source-sha') + 1]
out.mkdir(parents=True)
(out / 'release-assets.json').write_text(json.dumps({
    'schemaVersion': 1, 'tag': tag, 'sourceSha': sha, 'assets': []
}) + '\\n')
with open(os.environ['RELEASE_PIPELINE_LOG'], 'a') as stream:
    stream.write('build ' + ' '.join(args) + '\\n')
""",
            encoding="utf-8",
        )
        (self.repo / "fake-verifier.py").write_text(
            """#!/usr/bin/env python3
import os, sys
with open(os.environ['RELEASE_PIPELINE_LOG'], 'a') as stream:
    stream.write('verify ' + ' '.join(sys.argv[1:]) + '\\n')
""",
            encoding="utf-8",
        )
        (self.repo / "fake-reconciler.mjs").write_text(
            """import fs from 'node:fs';
fs.appendFileSync(process.env.RELEASE_PIPELINE_LOG, 'reconcile ' + process.argv.slice(2).join(' ') + '\\n');
""",
            encoding="utf-8",
        )
        (self.repo / "fake-readiness-collector.py").write_text(
            """#!/usr/bin/env python3
import os, pathlib, sys
args = sys.argv[1:]
output = pathlib.Path(args[args.index('--output') + 1])
output.write_text('{}\\n')
with open(os.environ['RELEASE_PIPELINE_LOG'], 'a') as stream:
    stream.write('collect ' + ' '.join(args) + '\\n')
""",
            encoding="utf-8",
        )
        (self.repo / "fake-readiness-verifier.py").write_text(
            """#!/usr/bin/env python3
import os, sys
with open(os.environ['RELEASE_PIPELINE_LOG'], 'a') as stream:
    stream.write('readiness ' + ' '.join(sys.argv[1:]) + '\\n')
""",
            encoding="utf-8",
        )
        self.pipeline_log = self.repo / "pipeline.log"
        self.upstream_evidence = self.repo / "upstream-evidence.json"
        self.upstream_evidence.write_text("{}\n", encoding="utf-8")
        self.blocker_evidence = self.repo / "blocker-evidence.json"
        self.blocker_evidence.write_text("{}\n", encoding="utf-8")
        self.write_license_policy("approved")
        subprocess.run(["git", "add", "."], cwd=self.repo, check=True)
        subprocess.run(
            ["git", "commit", "-qm", "chore: baseline"],
            cwd=self.repo,
            check=True,
        )
        subprocess.run(["git", "tag", "v0.53.1"], cwd=self.repo, check=True)
        (self.repo / "tracked.txt").write_text(
            "baseline\ncandidate\n",
            encoding="utf-8",
        )
        subprocess.run(["git", "add", "tracked.txt"], cwd=self.repo, check=True)
        subprocess.run(
            ["git", "commit", "-qm", "fix: candidate"],
            cwd=self.repo,
            check=True,
        )
        self.tested_sha = self.git("rev-parse", "HEAD")

    def git(self, *arguments: str) -> str:
        return subprocess.run(
            ["git", *arguments],
            cwd=self.repo,
            check=True,
            capture_output=True,
            text=True,
        ).stdout.strip()

    def set_fake_release(self, body: str) -> None:
        path = self.repo / "fake-release.sh"
        path.write_text(
            "#!/usr/bin/env bash\nset -euo pipefail\n" + body,
            encoding="utf-8",
        )
        path.chmod(0o755)
        subprocess.run(["git", "add", "fake-release.sh"], cwd=self.repo, check=True)
        subprocess.run(
            ["git", "commit", "-qm", "test: configure fake release"],
            cwd=self.repo,
            check=True,
        )
        self.tested_sha = self.git("rev-parse", "HEAD")

    def write_license_policy(self, status: str) -> None:
        approved = status == "approved"
        policy: dict[str, object] = {
            "schemaVersion": 1,
            "kind": "haxe.go-license-policy",
            "status": status,
            "approval": {
                "decidedBy": None,
                "authority": None,
                "decisionDate": None,
                "decisionRecord": None,
                "scopeSha256": None,
            },
            "shippedSourcePatterns": ["tracked.txt"],
            "components": [
                {
                    "id": "fixture",
                    "provenance": "repository-authored",
                    "sourcePatterns": ["tracked.txt"],
                    "declaredLicenses": ["MIT"],
                    "generatedOutputTreatment": (
                        "not-in-generated-output" if approved else "unresolved"
                    ),
                }
            ],
            "generatedOutputClasses": [
                {
                    "id": "fixture-output",
                    "origin": "fixture",
                    "licenseTreatment": (
                        "project-does-not-assert-ownership"
                        if approved
                        else "unresolved"
                    ),
                    "requiredArtifacts": [] if approved else None,
                }
            ],
            "releasePackage": {
                "requiredFiles": [
                    {
                        "sourcePath": "LICENSE",
                        "packagePath": "LICENSE",
                        "sha256": sha256(self.repo / "LICENSE"),
                    }
                ]
            },
            "unresolvedQuestions": [] if approved else ["fixture decision"],
        }
        if approved:
            policy["approval"] = {
                "decidedBy": "Fixture Owner",
                "authority": "project-copyright-owner",
                "decisionDate": "2026-07-14",
                "decisionRecord": "fixture-decision",
                "scopeSha256": scope_digest(policy),
            }
        (self.repo / "license-policy.json").write_text(
            json.dumps(policy, indent=2) + "\n",
            encoding="utf-8",
        )

    def run_wrapper(
        self,
        *,
        tested_sha: str | None = None,
    ) -> subprocess.CompletedProcess[str]:
        env = {
            **os.environ,
            "CI": "false",
            "RELEASE_TESTED_SHA": tested_sha or self.tested_sha,
            "SEMANTIC_RELEASE_BIN": str(self.repo / "fake-release.sh"),
            "GITHUB_REPOSITORY": "owner/repo",
            "RELEASE_ARTIFACT_BUILDER": str(self.repo / "fake-builder.py"),
            "RELEASE_ASSET_VERIFIER": str(self.repo / "fake-verifier.py"),
            "RELEASE_RECONCILER": str(self.repo / "fake-reconciler.mjs"),
            "RELEASE_READINESS_COLLECTOR": str(
                self.repo / "fake-readiness-collector.py"
            ),
            "RELEASE_READINESS_VERIFIER": str(
                self.repo / "fake-readiness-verifier.py"
            ),
            "RELEASE_READINESS_POLICY": str(self.repo / "license-policy.json"),
            "RELEASE_UPSTREAM_GATES_SHA": tested_sha or self.tested_sha,
            "RELEASE_UPSTREAM_EVIDENCE": str(self.upstream_evidence),
            "RELEASE_BLOCKER_EVIDENCE": str(self.blocker_evidence),
            "RELEASE_PIPELINE_LOG": str(self.pipeline_log),
        }
        return subprocess.run(
            ["bash", "scripts/release/run-same-sha-release.sh"],
            cwd=self.repo,
            env=env,
            capture_output=True,
            text=True,
        )

    def test_no_release_is_a_clean_same_sha_noop(self) -> None:
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn(
            "same-tested-SHA release contract: OK (no new release tag)",
            proc.stdout,
        )
        self.assertEqual(self.git("rev-parse", "HEAD"), self.tested_sha)
        self.assertFalse(self.pipeline_log.exists())

    def test_unapproved_license_policy_fails_before_release(self) -> None:
        self.write_license_policy("unresolved")
        subprocess.run(
            ["git", "add", "license-policy.json"], cwd=self.repo, check=True
        )
        subprocess.run(
            ["git", "commit", "-qm", "test: revoke fixture license approval"],
            cwd=self.repo,
            check=True,
        )
        self.tested_sha = self.git("rev-parse", "HEAD")
        proc = self.run_wrapper()
        self.assertNotEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("license policy is unresolved", proc.stderr)
        self.assertNotIn("same-tested-SHA release contract", proc.stdout)

    def test_new_tag_must_point_at_tested_sha(self) -> None:
        self.set_fake_release("git tag v0.54.0\n")
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("v0.54.0", proc.stdout)
        self.assertEqual(self.git("rev-parse", "v0.54.0^{commit}"), self.tested_sha)
        pipeline = self.pipeline_log.read_text(encoding="utf-8").splitlines()
        self.assertEqual(7, len(pipeline), pipeline)
        self.assertIn("build --version 0.54.0 --tag v0.54.0", pipeline[0])
        self.assertIn("verify --assets", pipeline[1])
        self.assertIn("collect --phase candidate", pipeline[2])
        self.assertIn("readiness --policy", pipeline[3])
        self.assertIn("reconcile --repository owner/repo --assets", pipeline[4])
        self.assertIn("collect --phase published", pipeline[5])
        self.assertIn("readiness --policy", pipeline[6])

    def test_existing_exact_tag_completes_interrupted_release(self) -> None:
        subprocess.run(["git", "tag", "v0.54.0"], cwd=self.repo, check=True)
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("completing or verifying existing exact tag v0.54.0", proc.stdout)
        pipeline = self.pipeline_log.read_text(encoding="utf-8").splitlines()
        self.assertEqual(7, len(pipeline), pipeline)

    def test_multiple_existing_exact_tags_fail_before_asset_build(self) -> None:
        subprocess.run(["git", "tag", "v0.54.0"], cwd=self.repo, check=True)
        subprocess.run(["git", "tag", "v0.54.1"], cwd=self.repo, check=True)
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("multiple canonical release tags", proc.stderr)
        self.assertFalse(self.pipeline_log.exists())

    def test_tracked_checkout_mutation_fails(self) -> None:
        self.set_fake_release("printf 'mutated\\n' >> tracked.txt\n")
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("tracked checkout changed", proc.stderr)

    def test_tag_on_another_commit_fails(self) -> None:
        self.set_fake_release("git tag v0.54.0 HEAD^\n")
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("does not match tested SHA", proc.stderr)

    def test_development_sentinel_cannot_become_a_release_tag(self) -> None:
        self.set_fake_release("git tag v0.0.0\n")
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("development sentinel", proc.stderr)

    def test_ci_requires_explicit_tested_sha_even_when_github_sha_exists(self) -> None:
        env = {
            **os.environ,
            "CI": "true",
            "GITHUB_SHA": self.tested_sha,
            "SEMANTIC_RELEASE_BIN": str(self.repo / "fake-release.sh"),
        }
        env.pop("RELEASE_TESTED_SHA", None)
        proc = subprocess.run(
            ["bash", "scripts/release/run-same-sha-release.sh"],
            cwd=self.repo,
            env=env,
            capture_output=True,
            text=True,
        )
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("RELEASE_TESTED_SHA is required in CI", proc.stderr)

    def test_ci_requires_upstream_gate_sha_before_semantic_release(self) -> None:
        release_marker = self.repo / "semantic-release-ran"
        self.set_fake_release(f"touch {release_marker}\n")
        env = {
            **os.environ,
            "CI": "true",
            "RELEASE_TESTED_SHA": self.tested_sha,
            "SEMANTIC_RELEASE_BIN": str(self.repo / "fake-release.sh"),
        }
        env.pop("RELEASE_UPSTREAM_GATES_SHA", None)
        proc = subprocess.run(
            ["bash", "scripts/release/run-same-sha-release.sh"],
            cwd=self.repo,
            env=env,
            capture_output=True,
            text=True,
        )
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn(
            "RELEASE_UPSTREAM_GATES_SHA must match tested SHA", proc.stderr
        )
        self.assertFalse(release_marker.exists())

    def test_untested_input_sha_fails_before_release(self) -> None:
        proc = self.run_wrapper(tested_sha="b" * 40)
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("does not match RELEASE_TESTED_SHA", proc.stderr)
        self.assertNotIn("v0.54.0", self.git("tag", "--list"))


if __name__ == "__main__":
    unittest.main()
