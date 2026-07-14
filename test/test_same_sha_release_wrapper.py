#!/usr/bin/env python3

from __future__ import annotations

import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
WRAPPER = ROOT / "scripts" / "release" / "run-same-sha-release.sh"


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
        (self.repo / "tracked.txt").write_text("baseline\n", encoding="utf-8")
        (self.repo / "fake-release.sh").write_text(
            "#!/usr/bin/env bash\nexit 0\n",
            encoding="utf-8",
        )
        (self.repo / "fake-release.sh").chmod(0o755)
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

    def test_new_tag_must_point_at_tested_sha(self) -> None:
        self.set_fake_release("git tag v0.54.0\n")
        proc = self.run_wrapper()
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("v0.54.0", proc.stdout)
        self.assertEqual(self.git("rev-parse", "v0.54.0^{commit}"), self.tested_sha)

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

    def test_untested_input_sha_fails_before_release(self) -> None:
        proc = self.run_wrapper(tested_sha="b" * 40)
        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("does not match RELEASE_TESTED_SHA", proc.stderr)
        self.assertNotIn("v0.54.0", self.git("tag", "--list"))


if __name__ == "__main__":
    unittest.main()
