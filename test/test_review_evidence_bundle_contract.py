#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import json
import re
import subprocess
import tempfile
import unittest
import zipfile
from pathlib import Path
from unittest import mock


REPO_ROOT = Path(__file__).resolve().parent.parent
BUILDER = REPO_ROOT / "scripts" / "review" / "build_gpt56_evidence.py"
REVIEW_README = REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "README.md"
EVIDENCE_RECORD = REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "evidence-cd79624f.json"
REVIEW_PROMPT = REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "review-prompt-cd79624f.md"
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

    def test_source_path_fixture_requires_an_explicit_narrow_allowance(self) -> None:
        builder = load_builder()
        fixture = r"D:\a\reflaxe.go\reflaxe.go\src\Main.hx:3"

        self.assertEqual("D:\\a\\", builder.find_machine_path(fixture))
        self.assertIsNone(
            builder.find_machine_path(
                fixture,
                allowed_literals=builder.PRIMARY_MACHINE_PATH_FIXTURES,
            )
        )
        self.assertEqual(
            "D:\\a\\",
            builder.find_machine_path(
                r"D:\a\other\other\src\Main.hx:3",
                allowed_literals=builder.PRIMARY_MACHINE_PATH_FIXTURES,
            ),
        )

    def test_github_workspace_redaction_precedes_nested_runner_home(self) -> None:
        builder = load_builder()
        runner_home = "/" + "home" + "/runner"
        raw = "\n".join(
            [
                runner_home + "/work/reflaxe.go/reflaxe.go/test/run-ci.py:9",
                runner_home + "/.config/review/token.json",
            ]
        )

        with mock.patch.object(builder.Path, "home", return_value=Path(runner_home)):
            redacted = builder.redact_machine_paths(raw)

        self.assertEqual(
            redacted,
            "<github-workspace>/test/run-ci.py:9\n<local-home>/.config/review/token.json",
        )

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

    def test_committed_evidence_record_preserves_the_audited_bundle_identity(self) -> None:
        record = json.loads(EVIDENCE_RECORD.read_text(encoding="utf-8"))
        sha256 = re.compile(r"^[0-9a-f]{64}$")

        self.assertEqual(record["schema_version"], 1)
        self.assertEqual(record["kind"], "haxe.go-gpt-5.6-pro-evidence-record")
        self.assertEqual(
            record["source"]["commit"],
            "cd79624f855521dbf320ac2b7524d889ca388c0e",
        )
        self.assertEqual(record["source"]["tracked_file_count"], 5368)
        self.assertEqual(record["artifact"]["bytes"], 13502320)
        self.assertRegex(record["artifact"]["sha256"], sha256)
        self.assertRegex(record["builder"]["file_sha256"], sha256)
        self.assertEqual(record["validation"]["deterministic_build_count"], 2)
        self.assertTrue(record["validation"]["outer_zip_byte_identical"])
        self.assertEqual(record["validation"]["internal_sha256_entries_verified"], 379)
        self.assertTrue(record["validation"]["gitleaks_passed"])
        self.assertTrue(record["validation"]["machine_local_path_check_passed"])
        self.assertEqual(len(record["ci_runs"]), 5)
        self.assertTrue(
            all(
                run["head_sha"] == record["source"]["commit"]
                and run["conclusion"] == "success"
                for run in record["ci_runs"]
            )
        )
        self.assertEqual(len(record["repomix"]["security_excluded_paths"]), 8)
        self.assertTrue(record["repomix"]["raw_companions_included"])
        self.assertEqual(record["roadmap"]["total_issues"], 663)
        self.assertEqual(record["roadmap"]["dependency_cycle_count"], 0)

    def test_review_prompt_is_evidence_bound_go_specific_and_adjudicable(self) -> None:
        text = REVIEW_PROMPT.read_text(encoding="utf-8")

        for exact_identity in [
            "cd79624f855521dbf320ac2b7524d889ca388c0e",
            "ab4b0a1097229ad2202ca7da6c092b2b85cba537522951fb625f6b1312c4b511",
            "haxe-go-gpt56-evidence-cd79624f.zip",
            "5b8c9416f963e541229e633a2bb655a93e3e9c16",
            "08faba040457165b883ae5327315581979ea07db",
            "68625fa91ffff48c5ffb269bff01c6f3e716128c",
        ]:
            self.assertIn(exact_identity, text)

        for verdict in [
            "bounded production use now",
            "stable 1.x compatibility promise",
            "compiler architecture",
            "portable-specialization direction",
            "Go-native authoring experience",
            "generated-Go output quality",
            "release and distribution integrity",
        ]:
            self.assertIn(verdict, text)

        for product_rule in [
            "Portable is the default product path",
            "Go-native by opt-in",
            "metal-like generated Go whenever",
            "typed externs/facades",
            "macros only for genuine Haxe syntax gaps",
            "AST-first",
            "staged target `_std`",
        ]:
            self.assertIn(product_rule, text)

        for audit_term in [
            "## Known facts",
            "## Open hypotheses",
            "canonical `_std`",
            "GoRaw",
            "Haxe 4.3.7",
            "Go memory model",
            "panic/recover",
            "go.Select",
            "go.Result",
            "gofmt",
            "go vet",
            "staticcheck",
            "race detector",
            "govulncheck",
            "SemVer",
            "false positive",
            "accepted / rejected / duplicate / deferred / evidence-gap",
        ]:
            self.assertIn(audit_term, text)

        for finding_field in [
            "violated invariant",
            "repository-relative file and exact line range",
            "concrete failure scenario",
            "minimal root-cause fix or scope disposition",
            "exact regression or empirical evidence required",
            "existing Bead or milestone placement",
        ]:
            self.assertIn(finding_field, text)

        for epic in range(4, 13):
            self.assertIn(f"`haxe_go-vfp.{epic}`", text)

        self.assertIn("The Git archive is the source authority", text)
        self.assertIn("Do not implement changes", text)
        self.assertIn("Do not write “add more tests”", text)
        self.assertIn("precedent, not proof", text)

    def test_contract_is_registered_in_release_checks(self) -> None:
        runner = RELEASE_CONTRACT_RUNNER.read_text(encoding="utf-8")
        self.assertIn("test/test_review_evidence_bundle_contract.py", runner)


if __name__ == "__main__":
    unittest.main()
