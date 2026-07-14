#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
POLICY = ROOT / "license-policy.json"
POLICY_DOC = ROOT / "LICENSING.md"
VERIFIER = ROOT / "scripts" / "release" / "verify-license-policy.py"
HAXE_STDLIB_LICENSE = ROOT / "licenses" / "HAXE-STDLIB-MIT.txt"
REFLAXE_LICENSE = ROOT / "vendor" / "reflaxe" / "LICENSE"


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def canonical_scope_digest(policy: dict[str, object]) -> str:
    scope = {
        key: policy[key]
        for key in (
            "shippedSourcePatterns",
            "components",
            "generatedOutputClasses",
            "releasePackage",
        )
    }
    encoded = json.dumps(
        scope,
        ensure_ascii=False,
        separators=(",", ":"),
        sort_keys=True,
    ).encode("utf-8")
    return hashlib.sha256(encoded).hexdigest()


def run_verifier(
    root: Path,
    *,
    mode: str,
    package_root: Path | None = None,
) -> subprocess.CompletedProcess[str]:
    command = [
        "python3",
        str(VERIFIER),
        "--root",
        str(root),
        "--mode",
        mode,
    ]
    if package_root is not None:
        command.extend(["--package-root", str(package_root)])
    return subprocess.run(
        command,
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=30,
    )


def write_json(path: Path, value: object) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2) + "\n", encoding="utf-8")


def write_fixture(root: Path, *, status: str) -> dict[str, object]:
    license_path = root / "LICENSE"
    license_path.parent.mkdir(parents=True, exist_ok=True)
    license_path.write_text("fixture license\n", encoding="utf-8")
    source_path = root / "src" / "fixture.txt"
    source_path.parent.mkdir(parents=True, exist_ok=True)
    source_path.write_text("fixture source\n", encoding="utf-8")

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
        "shippedSourcePatterns": ["src/**/*.txt"],
        "components": [
            {
                "id": "fixture-component",
                "provenance": "repository-authored",
                "sourcePatterns": ["src/**/*.txt"],
                "declaredLicenses": ["MIT"],
                "generatedOutputTreatment": "not-in-generated-output",
            }
        ],
        "generatedOutputClasses": [
            {
                "id": "fixture-output",
                "origin": "fixture",
                "licenseTreatment": "project-does-not-assert-ownership",
                "requiredArtifacts": [],
            }
        ],
        "releasePackage": {
            "requiredFiles": [
                {
                    "sourcePath": "LICENSE",
                    "packagePath": "LICENSE",
                    "sha256": sha256(license_path),
                }
            ]
        },
        "unresolvedQuestions": [] if status == "approved" else ["fixture decision"],
    }
    if status == "approved":
        approval = policy["approval"]
        assert isinstance(approval, dict)
        approval.update(
            {
                "decidedBy": "Fixture Owner",
                "authority": "project-copyright-owner",
                "decisionDate": "2026-07-14",
                "decisionRecord": "fixture-decision-1",
                "scopeSha256": canonical_scope_digest(policy),
            }
        )
    write_json(root / "license-policy.json", policy)
    return policy


class LicensePolicyContractTest(unittest.TestCase):
    maxDiff = None

    def test_repository_inventory_is_auditable_but_release_blocking(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        self.assertEqual(1, policy.get("schemaVersion"))
        self.assertEqual("haxe.go-license-policy", policy.get("kind"))
        self.assertEqual("unresolved", policy.get("status"))
        self.assertEqual(
            {
                "haxe-standard-library-overrides",
                "reflaxe-framework",
                "reflaxe-go-compiler-and-library",
                "reflaxe-go-hxrt",
            },
            {component["id"] for component in policy["components"]},
        )
        self.assertEqual(
            {
                "compiler-emitted-framework-support",
                "copied-hxrt-source",
                "lowered-haxe-standard-library",
                "lowered-user-program",
            },
            {output["id"] for output in policy["generatedOutputClasses"]},
        )
        self.assertTrue(policy["unresolvedQuestions"])
        self.assertTrue(
            all(value is None for value in policy["approval"].values()),
            "an unresolved policy must not carry an apparent approval",
        )

        audit = run_verifier(ROOT, mode="audit")
        self.assertEqual(0, audit.returncode, audit.stdout + audit.stderr)
        self.assertIn("inventory and license material: OK", audit.stdout)
        self.assertIn(canonical_scope_digest(policy), audit.stdout)

        release = run_verifier(ROOT, mode="release")
        self.assertEqual(1, release.returncode, release.stdout + release.stderr)
        self.assertIn("license policy is unresolved", release.stderr)

    def test_authoritative_license_material_is_preserved_and_packaged(self) -> None:
        self.assertEqual(
            "68f17b4096da082538e5a6e2c050cd510f47fe33574522f3dc397750da600e79",
            sha256(REFLAXE_LICENSE),
        )
        self.assertEqual(
            "61c9e5c8ca48e1f6e27f66cc6fb2eb11865a08672e1c793a13cfdaa89ad1bb74",
            sha256(HAXE_STDLIB_LICENSE),
        )
        self.assertIn("Copyright (c) 2022 Robert Borghese", REFLAXE_LICENSE.read_text())
        self.assertIn(
            "Copyright (C)2005-2016 Haxe Foundation",
            HAXE_STDLIB_LICENSE.read_text(),
        )

        run_hx = (ROOT / "Run.hx").read_text(encoding="utf-8")
        package_audit = (
            ROOT / "scripts" / "ci" / "canonical_stdlib_layout_check.py"
        ).read_text(encoding="utf-8")
        artifact_verifier = (
            ROOT / "scripts" / "release" / "verify-haxelib-artifact.py"
        ).read_text(encoding="utf-8")
        for path in (
            "LICENSE",
            "LICENSING.md",
            "license-policy.json",
            "licenses/HAXE-STDLIB-MIT.txt",
        ):
            self.assertIn(path, run_hx)
            self.assertIn(path, package_audit)
            self.assertIn(path, artifact_verifier)
        self.assertIn('copyVendoredReflaxeFile("LICENSE")', run_hx)
        self.assertIn("vendor/reflaxe/LICENSE", package_audit)
        self.assertIn("vendor/reflaxe/LICENSE", artifact_verifier)

    def test_release_entrypoints_cannot_bypass_the_policy_gate(self) -> None:
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        self.assertEqual(
            "python3 scripts/release/verify-license-policy.py --mode release",
            package["scripts"].get("release:license-policy"),
        )
        wrapper = (
            ROOT / "scripts" / "release" / "run-same-sha-release.sh"
        ).read_text(encoding="utf-8")
        self.assertIn("verify-license-policy.py --mode release", wrapper)
        self.assertNotIn("LICENSE_POLICY_VERIFIER", wrapper)

        workflow = (
            ROOT / ".github" / "workflows" / "ci-harness.yml"
        ).read_text(encoding="utf-8")
        release_job = workflow.split("\n  semantic-release:", 1)[1]
        self.assertIn("run: npm run release:license-policy", release_job)
        self.assertLess(
            release_job.index("run: npm run release:license-policy"),
            release_job.index("run: npm run release\n"),
        )

        release_contracts = (
            ROOT / "test" / "run-release-contracts.py"
        ).read_text(encoding="utf-8")
        self.assertIn('"test/test_license_policy_contract.py"', release_contracts)

    def test_approved_policy_is_bound_to_scope_and_notice_bytes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-license-policy-") as raw:
            fixture = Path(raw)
            policy = write_fixture(fixture, status="approved")

            approved = run_verifier(fixture, mode="release")
            self.assertEqual(0, approved.returncode, approved.stdout + approved.stderr)
            self.assertIn("approved release policy: OK", approved.stdout)

            policy["generatedOutputClasses"][0]["requiredArtifacts"] = ["NOTICE"]
            write_json(fixture / "license-policy.json", policy)
            stale_scope = run_verifier(fixture, mode="release")
            self.assertEqual(1, stale_scope.returncode, stale_scope.stdout + stale_scope.stderr)
            self.assertIn("approval scope digest is stale", stale_scope.stderr)

            write_fixture(fixture, status="approved")
            (fixture / "LICENSE").write_text("tampered license\n", encoding="utf-8")
            stale_notice = run_verifier(fixture, mode="release")
            self.assertEqual(1, stale_notice.returncode, stale_notice.stdout + stale_notice.stderr)
            self.assertIn("license material hash differs", stale_notice.stderr)

    def test_package_root_notice_audit_rejects_missing_material(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-license-package-") as raw:
            fixture = Path(raw) / "source"
            package_root = Path(raw) / "package"
            policy = write_fixture(fixture, status="approved")
            package_root.mkdir()
            for record in policy["releasePackage"]["requiredFiles"]:
                source = fixture / record["sourcePath"]
                destination = package_root / record["packagePath"]
                destination.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(source, destination)

            complete = run_verifier(
                fixture,
                mode="audit",
                package_root=package_root,
            )
            self.assertEqual(0, complete.returncode, complete.stdout + complete.stderr)

            (package_root / "LICENSE").unlink()
            missing = run_verifier(
                fixture,
                mode="audit",
                package_root=package_root,
            )
            self.assertEqual(1, missing.returncode, missing.stdout + missing.stderr)
            self.assertIn("required package license material is missing", missing.stderr)

    def test_unclassified_shipped_source_is_rejected(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-license-coverage-") as raw:
            fixture = Path(raw)
            policy = write_fixture(fixture, status="approved")
            extra = fixture / "extra" / "unclassified.txt"
            extra.parent.mkdir()
            extra.write_text("unclassified\n", encoding="utf-8")
            policy["shippedSourcePatterns"].append("extra/**/*.txt")
            policy["approval"]["scopeSha256"] = canonical_scope_digest(policy)
            write_json(fixture / "license-policy.json", policy)

            rejected = run_verifier(fixture, mode="audit")
            self.assertEqual(1, rejected.returncode, rejected.stdout + rejected.stderr)
            self.assertIn("shipped source is not assigned to a component", rejected.stderr)


if __name__ == "__main__":
    unittest.main()
