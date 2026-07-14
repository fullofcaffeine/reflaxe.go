#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import subprocess
import tempfile
import unittest
import zipfile
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
BUILDER = REPO_ROOT / "scripts" / "review" / "build_gpt56_evidence.py"
REVIEW_README = REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "README.md"
RELEASE_CONTRACT_RUNNER = REPO_ROOT / "test" / "run-release-contracts.py"


def load_builder():
    spec = importlib.util.spec_from_file_location("build_gpt56_evidence", BUILDER)
    if spec is None or spec.loader is None:
        raise AssertionError(f"cannot import evidence builder: {BUILDER}")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


class ReviewEvidenceBundleContractTest(unittest.TestCase):
    def test_builder_help_requires_explicit_provenance_inputs(self) -> None:
        self.assertTrue(BUILDER.exists(), "review evidence builder must exist")
        completed = subprocess.run(
            ["python3", str(BUILDER), "--help"],
            cwd=REPO_ROOT,
            capture_output=True,
            text=True,
        )
        self.assertEqual(completed.returncode, 0, completed.stderr)
        for option in [
            "--source-ref",
            "--rust-ref",
            "--ruby-ref",
            "--elixir-ref",
            "--ci-run",
            "--release-tag",
            "--output",
        ]:
            self.assertIn(option, completed.stdout)

    def test_machine_paths_are_redacted_without_hiding_repository_paths(self) -> None:
        builder = load_builder()
        users_root = "/" + "Users" + "/"
        github_work_root = "/" + "home/runner/work" + "/"
        raw = "\n".join(
            [
                users_root + "alice/work/haxe.go/src/Main.hx:12",
                github_work_root + "reflaxe.go/reflaxe.go/test/run-ci.py:9",
                r"D:\a\reflaxe.go\reflaxe.go\src\Main.hx:3",
                f"{Path.home().as_posix()}/.config/repomix/repomix.config.json",
                "src/reflaxe/go/GoCompiler.hx:42",
            ]
        )
        redacted = builder.redact_machine_paths(raw)
        self.assertNotIn(users_root + "alice", redacted)
        self.assertNotIn(github_work_root, redacted)
        self.assertNotIn(r"D:\a", redacted)
        self.assertIn("<local-home>/.config/repomix/repomix.config.json", redacted)
        self.assertIn("<local-workspace>/src/Main.hx:12", redacted)
        self.assertIn("<github-workspace>/test/run-ci.py:9", redacted)
        self.assertIn("src/reflaxe/go/GoCompiler.hx:42", redacted)

    def test_deterministic_zip_is_sorted_and_timestamp_independent(self) -> None:
        builder = load_builder()
        with tempfile.TemporaryDirectory() as raw_tmp:
            tmp = Path(raw_tmp)
            payload = tmp / "payload"
            payload.mkdir()
            (payload / "z.txt").write_text("zeta\n", encoding="utf-8")
            (payload / "a.txt").write_text("alpha\n", encoding="utf-8")
            first = tmp / "first.zip"
            second = tmp / "second.zip"

            builder.write_deterministic_zip(payload, first)
            (payload / "a.txt").touch()
            builder.write_deterministic_zip(payload, second)

            self.assertEqual(first.read_bytes(), second.read_bytes())
            with zipfile.ZipFile(first) as archive:
                self.assertEqual(archive.namelist(), ["a.txt", "z.txt"])
                self.assertTrue(all(info.date_time == (1980, 1, 1, 0, 0, 0) for info in archive.infolist()))

    def test_tool_logs_drop_progress_time_and_temporary_paths(self) -> None:
        builder = load_builder()
        private_temp_root = "/" + "private/var/folders" + "/"
        repomix_log = """⠸ Collect file... (98/100) src/Main.hx
✔ Packing completed successfully!
8 suspicious file(s) detected and excluded from the output:
1. test/fixture.pem
   - 1 security issue detected
  Total Files: 99 files
       Output: PRIVATE_TEMP_ROOTxx/random/output.xml
     Security: 1 suspicious file(s) detected and excluded
🎉 All Done!
""".replace("PRIVATE_TEMP_ROOT", private_temp_root)
        normalized_repomix = builder.normalize_repomix_log(repomix_log)
        self.assertNotIn("Collect file", normalized_repomix)
        self.assertNotIn(private_temp_root, normalized_repomix)
        self.assertIn("test/fixture.pem", normalized_repomix)
        self.assertIn("Output: <bundle-output>", normalized_repomix)
        self.assertEqual(builder.parse_repomix_security_exclusions(normalized_repomix), ["test/fixture.pem"])

        gitleaks_log = "6:12PM INF scanned ~123 bytes in 0.42s\n6:12PM INF no leaks found\n"
        normalized_gitleaks = builder.normalize_gitleaks_log(gitleaks_log)
        self.assertEqual(normalized_gitleaks, "PASS: no leaks found (max-archive-depth=1).\n")

    def test_reference_sets_cover_the_three_sibling_decisions(self) -> None:
        builder = load_builder()
        self.assertEqual(builder.PINNED_REPOMIX_PACKAGE, "repomix@1.14.0")
        references = builder.REFERENCE_PATHS
        self.assertIn("haxe.rust", references)
        self.assertIn("haxe.ruby", references)
        self.assertIn("haxe.elixir.codex", references)

        rust_paths = "\n".join(references["haxe.rust"])
        self.assertIn("SurfaceContractRegistry.hx", rust_paths)
        self.assertIn("RuntimeRequirementAnalyzer.hx", rust_paths)
        self.assertIn("NoHxrtEligibilityAnalyzer.hx", rust_paths)
        self.assertIn("std/rust/_std", rust_paths)
        self.assertIn("release-reference-architecture.md", rust_paths)

        ruby_paths = "\n".join(references["haxe.ruby"])
        self.assertIn("release-version-policy.md", ruby_paths)
        self.assertIn("release-hosting-and-repair.md", ruby_paths)

        elixir_paths = "\n".join(references["haxe.elixir.codex"])
        self.assertIn("std/elixir/_std", elixir_paths)
        self.assertIn("check-stdlib-source-layout.sh", elixir_paths)

    def test_review_docs_define_contents_exclusions_and_reproduction(self) -> None:
        text = REVIEW_README.read_text(encoding="utf-8")
        for phrase in [
            "exact Git commit",
            "generated portable and metal",
            "GitHub Actions logs",
            "release and tag state",
            "compatibility and ownership inventories",
            "Rust-family portable specialization",
            "canonical `_std`",
            "machine-local paths",
            ".beads/issues.jsonl",
            "scripts/review/build_gpt56_evidence.py",
            "SHA-256",
        ]:
            self.assertIn(phrase, text)

    def test_contract_is_registered_in_release_checks(self) -> None:
        runner = RELEASE_CONTRACT_RUNNER.read_text(encoding="utf-8")
        self.assertIn("test/test_review_evidence_bundle_contract.py", runner)


if __name__ == "__main__":
    unittest.main()
