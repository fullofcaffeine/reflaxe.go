#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import re
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
RECORD = ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "portable-beta-candidate-c22af0ea.json"
GUIDE = ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "portable-beta-candidate-c22af0ea.md"
CANDIDATE = "c22af0ea82e5e481e23277e513ed5b7c6b5c770b"
SHA256 = re.compile(r"^[0-9a-f]{64}$")


def candidate_file_sha256(relative: str) -> str:
    completed = subprocess.run(
        ["git", "show", f"{CANDIDATE}:{relative}"],
        cwd=ROOT,
        capture_output=True,
        check=True,
    )
    return hashlib.sha256(completed.stdout).hexdigest()


class PortableBetaCandidateEvidenceTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        cls.record = json.loads(RECORD.read_text(encoding="utf-8"))

    def test_record_pins_one_exact_unpublished_candidate(self) -> None:
        record = self.record
        self.assertEqual(record["schema_version"], 1)
        self.assertEqual(record["kind"], "haxe.go-portable-beta-candidate-evidence")
        self.assertEqual(record["tracker"], "haxe_go-vfp.12.4")
        self.assertEqual(record["source"]["commit"], CANDIDATE)
        self.assertEqual(record["source"]["tree"], "fbdb80c6ce39ba8d89c029334452b151a6cce4a3")
        self.assertEqual(record["source"]["parent"], "cc40b388779f6bcc265e74e235ed9c929c6bf77c")
        self.assertEqual(record["source"]["remote_branch"], "origin/master")
        self.assertFalse(record["publication"]["published"])
        self.assertEqual(record["publication"]["next_gate"], "haxe_go-vfp.12.5")

    def test_archive_is_hashed_reproducible_and_internally_verified(self) -> None:
        artifact = self.record["artifact"]
        self.assertEqual(artifact["filename"], "haxe-go-portable-beta-c22af0ea.zip")
        self.assertEqual(artifact["bytes"], 27_529_201)
        self.assertEqual(
            artifact["sha256"],
            "7d76936576bb9e7e654f73ac010a474305e23b2c08ec2fe2f406441af633eb66",
        )
        self.assertEqual(artifact["payload_count"], 383)
        validation = artifact["validation"]
        self.assertEqual(validation["deterministic_build_count"], 2)
        self.assertTrue(validation["outer_zip_byte_identical"])
        self.assertTrue(validation["zip_integrity_passed"])
        self.assertTrue(validation["source_archive_commit_verified"])
        self.assertTrue(validation["gitleaks_passed"])
        self.assertTrue(validation["machine_local_path_check_passed"])
        self.assertEqual(validation["internal_sha256_entries_verified"], 384)

    def test_every_hosted_workflow_is_green_on_the_same_commit(self) -> None:
        expected = {
            30743182325: "Examples Artifacts",
            30743182328: "Security - Gitleaks",
            30743182367: "Security - Static Analysis",
            30743182369: "CI Harness",
            30743182373: "CI - Quality",
        }
        runs = self.record["hosted_evidence"]["runs"]
        self.assertEqual({run["database_id"]: run["workflow"] for run in runs}, expected)
        self.assertTrue(all(run["head_sha"] == CANDIDATE for run in runs))
        self.assertTrue(all(run["conclusion"] == "success" for run in runs))
        self.assertEqual(self.record["hosted_evidence"]["artifact_count"], 20)
        self.assertEqual(self.record["hosted_evidence"]["artifacts_expire_at"], "2026-10-31T10:09:21Z")

    def test_release_claim_is_narrow_and_surface_separated(self) -> None:
        claim = self.record["claim"]
        self.assertEqual(claim["lifecycle"], "pre-1.0 beta")
        self.assertEqual(claim["preset"], "portable")
        self.assertEqual(claim["platform"], "linux-amd64")
        self.assertEqual(claim["operation_count"], 38)
        self.assertEqual(claim["symbol_count"], 173)
        self.assertEqual(
            claim["surface_sha256"],
            "99625a5bcb401561a8393ddcef5675ba552c1a84ab288ea4ed0e1cc950bac0d0",
        )
        self.assertEqual(
            claim["required_exclusions"],
            {
                "go-native": "haxe_go-vfp.9.1",
                "portable-http": "haxe_go-vfp.10.8",
                "stable-1.x": "haxe_go-vfp.12.7",
            },
        )
        self.assertEqual(claim["default_disposition"], "unlisted-is-excluded")
        self.assertEqual(claim["trust_model"], "trusted-source-only")
        self.assertFalse(claim["metal_results_advance_portable_claim"])

    def test_authority_hashes_and_license_approval_are_frozen(self) -> None:
        authorities = self.record["authorities"]
        expected = {
            "docs/compatibility-support-manifest.json": "3773df9cf7ade3a729ad6601e9e5cf96501cea0fc086b9ec00a05d64b5f3e2e8",
            "release/readiness-policy.json": "972c0612ab64c2d5230b842aab458bc67043714d597396c34623e6e4259ec719",
            "license-policy.json": "35476c8df6a9f461f539ad14a6cdea085e2d541f2b8f0cb72529f605c7397698",
            "scripts/review/build_gpt56_evidence.py": "450af38f55e71c81c9d18c9b716b889b167b1362d647c92a352b9e852cc9be65",
        }
        self.assertEqual({item["path"]: item["sha256"] for item in authorities}, expected)
        for path, digest in expected.items():
            self.assertRegex(digest, SHA256)
            self.assertEqual(candidate_file_sha256(path), digest)
        self.assertEqual(
            self.record["license_approval"]["scope_sha256"],
            "56e4dc13b3676562cd2ed6ac23f1706141a2098fdd67f408ceba026d84848c8a",
        )
        self.assertEqual(self.record["license_approval"]["approved_by"], "Marcelo Serpa")

    def test_rebuild_inputs_and_frozen_roadmap_are_explicit(self) -> None:
        references = {item["name"]: item["commit"] for item in self.record["references"]}
        self.assertEqual(
            references,
            {
                "haxe.rust": "5b8c9416f963e541229e633a2bb655a93e3e9c16",
                "haxe.ruby": "dbb70af0e48e252e413645b7bf16197a4776f0f8",
                "haxe.elixir.codex": "68625fa91ffff48c5ffb269bff01c6f3e716128c",
            },
        )
        roadmap = self.record["roadmap_snapshot"]
        self.assertEqual(roadmap["root_issue"], "haxe_go-vfp.12")
        self.assertEqual(roadmap["dependency_cycle_count"], 0)
        command = self.record["reproduction"]["command"]
        self.assertIn(CANDIDATE, command)
        self.assertEqual(command.count("--ci-run"), 5)
        self.assertEqual(command[-1], "haxe-go-portable-beta-c22af0ea.zip")

    def test_guide_explains_what_the_evidence_does_and_does_not_mean(self) -> None:
        text = GUIDE.read_text(encoding="utf-8")
        normalized = text.lower()
        for phrase in [
            "What this proves",
            "What this does not prove",
            "same exact source commit",
            "portable and metal are separate scorecards",
            "not a stable 1.x claim",
            "not a publication record",
            "HTTP remains excluded",
            "untrusted or adversarial source compilation remains excluded",
            "How to verify the ZIP",
            "haxe_go-vfp.12.5",
        ]:
            self.assertIn(phrase.lower(), normalized)


if __name__ == "__main__":
    unittest.main()
