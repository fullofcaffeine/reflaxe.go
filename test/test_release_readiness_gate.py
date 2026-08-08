#!/usr/bin/env python3

from __future__ import annotations

import copy
import json
import os
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
VERIFIER = ROOT / "scripts" / "release" / "verify-release-readiness.py"
COLLECTOR = ROOT / "scripts" / "release" / "collect-release-readiness.py"
UPSTREAM_COLLECTOR = (
    ROOT / "scripts" / "release" / "collect-upstream-release-evidence.py"
)
POLICY = ROOT / "release" / "readiness-policy.json"
PASS_FIXTURE = (
    ROOT / "test" / "fixtures" / "release_readiness" / "published-pass.json"
)


class ReleaseReadinessGateTest(unittest.TestCase):
    maxDiff = None

    def evidence(self) -> dict[str, object]:
        return json.loads(PASS_FIXTURE.read_text(encoding="utf-8"))

    def run_evidence(
        self,
        evidence: dict[str, object],
        *,
        mode: str = "fixture",
        env: dict[str, str] | None = None,
        policy: dict[str, object] | None = None,
    ) -> subprocess.CompletedProcess[str]:
        with tempfile.TemporaryDirectory(prefix="haxe-go-readiness-") as raw:
            path = Path(raw) / "evidence.json"
            path.write_text(json.dumps(evidence), encoding="utf-8")
            policy_path = POLICY
            if policy is not None:
                policy_path = Path(raw) / "policy.json"
                policy_path.write_text(json.dumps(policy), encoding="utf-8")
            return subprocess.run(
                [
                    "python3",
                    str(VERIFIER),
                    "--policy",
                    str(policy_path),
                    "--evidence",
                    str(path),
                    "--mode",
                    mode,
                ],
                cwd=ROOT,
                text=True,
                capture_output=True,
                timeout=30,
                env=env,
            )

    def assert_rejected(
        self, evidence: dict[str, object], expected_message: str
    ) -> None:
        result = self.run_evidence(evidence)
        self.assertNotEqual(0, result.returncode, result.stdout + result.stderr)
        self.assertIn(expected_message, result.stderr)

    def test_published_fixture_passes(self) -> None:
        result = self.run_evidence(self.evidence())
        self.assertEqual(0, result.returncode, result.stdout + result.stderr)
        self.assertIn("published release evidence: READY", result.stdout)

    def test_network_policy_splits_admitted_operations_from_open_blockers(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        compatibility = policy["compatibility"]

        self.assertEqual(
            ["platform:linux-amd64", "preset:portable"],
            compatibility["admittedScopes"],
        )
        self.assertNotIn(
            "portable-networking",
            compatibility["requiredExclusions"],
        )
        self.assertEqual(
            "haxe_go-vfp.10.8",
            compatibility["requiredExclusions"]["portable-http"],
        )
        self.assertNotIn(
            "portable-socket-advanced",
            compatibility["requiredExclusions"],
        )
        self.assertNotIn("haxe_go-vfp.10.4", compatibility["blockerScopes"])
        self.assertEqual(
            "portable-http",
            compatibility["blockerScopes"]["haxe_go-vfp.10.8"],
        )
        self.assertEqual(
            "preset:portable",
            compatibility["blockerScopes"]["haxe_go-vfp.10.9"],
        )

    def test_open_socket_admission_review_blocks_the_portable_release(self) -> None:
        evidence = self.evidence()
        parent = next(
            blocker
            for blocker in evidence["blockers"]["records"]
            if blocker["id"] == "haxe_go-vfp.10.9"
        )
        parent["status"] = "open"
        self.assert_rejected(
            evidence,
            "applicable unresolved P0/P1 blocker: haxe_go-vfp.10.9",
        )

    def test_final_admission_record_binds_oracle_and_local_review_to_release_sha(
        self,
    ) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        admission_policy = policy["finalAdmission"]
        self.assertEqual("haxe_go-vfp.12.5", admission_policy["owner"])
        self.assertEqual(
            "orq_20260803T080027Z_615d041e",
            admission_policy["oracleRequestId"],
        )
        self.assertEqual(
            "ce8cc8cca65d48dceae11155e9fc651dbf8bef7f611270e7b0a43c194e33d5f9",
            admission_policy["frozenPacketSha256"],
        )
        self.assertEqual(
            "docs/reviews/gpt-5.6-pro/disposition-c22af0ea-portable-beta.md",
            admission_policy["dispositionPath"],
        )
        self.assertEqual(
            "preset:portable",
            policy["compatibility"]["blockerScopes"]["haxe_go-vfp.12.5"],
        )

        evidence = self.evidence()
        final_record = next(
            blocker
            for blocker in evidence["blockers"]["records"]
            if blocker["id"] == "haxe_go-vfp.12.5"
        )
        admission = final_record["admission"]
        self.assertEqual(
            "c22af0ea82e5e481e23277e513ed5b7c6b5c770b",
            admission["oracleReview"]["reviewedSourceSha"],
        )
        self.assertEqual(
            evidence["release"]["testedSha"],
            admission["localDisposition"]["reviewedSourceSha"],
        )

        mutations = (
            (
                "Oracle verdict",
                lambda value: value["oracleReview"].update(
                    verdict="READY"
                ),
            ),
            (
                "Oracle reviewed SHA",
                lambda value: value["oracleReview"].update(
                    reviewedSourceSha="2" * 40
                ),
            ),
            (
                "Oracle request",
                lambda value: value["oracleReview"].update(
                    requestId="orq_wrong"
                ),
            ),
            (
                "frozen packet",
                lambda value: value["oracleReview"].update(
                    frozenPacketSha256="0" * 64
                ),
            ),
            (
                "local admission reviewed SHA",
                lambda value: value["localDisposition"].update(
                    reviewedSourceSha="2" * 40
                ),
            ),
            (
                "local admission disposition",
                lambda value: value["localDisposition"].update(
                    dispositionSha256="0" * 64
                ),
            ),
        )
        for message, mutate in mutations:
            with self.subTest(message=message):
                mutated = self.evidence()
                record = next(
                    blocker
                    for blocker in mutated["blockers"]["records"]
                    if blocker["id"] == "haxe_go-vfp.12.5"
                )
                mutate(record["admission"])
                self.assert_rejected(mutated, message)

    def test_additional_review_blocker_cannot_invent_a_scope(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        policy["compatibility"]["blockerScopes"]["haxe_go-vfp.10.9"] = (
            "portable-socket-advanced"
        )
        result = self.run_evidence(self.evidence(), policy=policy)
        self.assertNotEqual(0, result.returncode, result.stdout + result.stderr)
        self.assertIn(
            "additional readiness blocker does not govern an admitted scope",
            result.stderr,
        )

    def test_candidate_fixture_passes_without_pretending_assets_are_hosted(self) -> None:
        evidence = self.evidence()
        evidence["phase"] = "candidate"
        evidence["github"] = None
        result = self.run_evidence(evidence)
        self.assertEqual(0, result.returncode, result.stdout + result.stderr)
        self.assertIn("candidate release evidence: READY", result.stdout)

    def test_release_identity_mismatches_fail(self) -> None:
        mutations = (
            ("manifest tag", lambda value: value["release"].update(manifestTag="v0.53.0")),
            ("release tag", lambda value: value["release"].update(tag="v0.53.0")),
            (
                "tested SHA",
                lambda value: value["release"].update(
                    sourceSha="2222222222222222222222222222222222222222"
                ),
            ),
        )
        for message, mutate in mutations:
            with self.subTest(message=message):
                evidence = self.evidence()
                mutate(evidence)
                self.assert_rejected(evidence, message)

    def test_artifact_and_provenance_failures_are_rejected(self) -> None:
        evidence = self.evidence()
        evidence["artifacts"]["assets"] = evidence["artifacts"]["assets"][:-1]
        self.assert_rejected(evidence, "required artifact roles")

        evidence = self.evidence()
        evidence["artifacts"]["provenanceSubjects"].remove(
            "reflaxe.go-0.54.0.zip"
        )
        self.assert_rejected(evidence, "provenance subjects")

    def test_toolchain_security_license_and_api_failures_are_rejected(self) -> None:
        mutations = (
            (
                "supported toolchain",
                lambda value: value["toolchains"]["go"]["resolved"].append(
                    "1.24.0"
                ),
            ),
            (
                "reachable vulnerabilities",
                lambda value: value["security"]["reachableVulnerabilities"].append(
                    "GO-2099-0001"
                ),
            ),
            (
                "licensing approval",
                lambda value: value["licensing"].update(status="pending"),
            ),
            (
                "unresolved licensing questions",
                lambda value: value["licensing"]["unresolvedQuestions"].append(
                    "ownership"
                ),
            ),
            (
                "public API policy",
                lambda value: value["publicApi"].update(result="fail"),
            ),
        )
        for message, mutate in mutations:
            with self.subTest(message=message):
                evidence = self.evidence()
                mutate(evidence)
                self.assert_rejected(evidence, message)

    def test_exact_hosted_runner_image_is_required(self) -> None:
        evidence = self.evidence()
        evidence["platform"]["runnerImage"]["version"] = ""
        self.assert_rejected(evidence, "hosted runner image")

        evidence = self.evidence()
        evidence["platform"]["architecture"] = "arm64"
        self.assert_rejected(evidence, "release platform")

    def test_only_applicable_p0_or_p1_blockers_block(self) -> None:
        evidence = self.evidence()
        evidence["blockers"]["records"].append(
            {
                "id": "haxe_go-test",
                "priority": 1,
                "status": "open",
                "scopes": ["preset:portable"],
            }
        )
        self.assert_rejected(evidence, "applicable unresolved P0/P1 blocker")

        result = self.run_evidence(self.evidence())
        self.assertEqual(0, result.returncode, result.stdout + result.stderr)

    def test_exclusions_must_be_owned_and_cannot_be_advertised(self) -> None:
        evidence = self.evidence()
        evidence["compatibility"]["exclusions"][0]["owner"] = ""
        self.assert_rejected(evidence, "unowned exclusion")

        evidence = self.evidence()
        evidence["compatibility"]["exclusions"][0][
            "advertisedAsSupported"
        ] = True
        self.assert_rejected(evidence, "excluded scope advertised as supported")

    def test_claim_cannot_exceed_exact_evidence(self) -> None:
        evidence = self.evidence()
        evidence["compatibility"]["admittedScopes"].append("go-native")
        self.assert_rejected(evidence, "admitted scopes")

        evidence = self.evidence()
        evidence["compatibility"]["evidencedScopes"].remove(
            "platform:linux-amd64"
        )
        self.assert_rejected(evidence, "claim exceeds evidence")

        evidence = self.evidence()
        evidence["compatibility"]["surfaceSha256"] = "0" * 64
        self.assert_rejected(evidence, "operation/member surface")

    def test_every_gate_must_pass_for_the_exact_tested_sha(self) -> None:
        evidence = self.evidence()
        evidence["security"]["gates"][0]["result"] = "fail"
        self.assert_rejected(evidence, "required security gate")

        evidence = self.evidence()
        evidence["security"]["gates"][0][
            "testedSha"
        ] = "2222222222222222222222222222222222222222"
        self.assert_rejected(evidence, "gate tested SHA")

    def test_published_phase_uses_github_api_truth(self) -> None:
        evidence = self.evidence()
        evidence["github"]["apiAuthoritative"] = False
        self.assert_rejected(evidence, "GitHub API state must be authoritative")

        evidence = self.evidence()
        evidence["github"]["assets"][0][
            "digest"
        ] = "sha256:eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
        self.assert_rejected(evidence, "hosted assets")

        evidence = self.evidence()
        evidence["github"]["targetCommit"] = (
            "2222222222222222222222222222222222222222"
        )
        self.assert_rejected(evidence, "hosted release target")

    def test_live_mode_replaces_supplied_host_state_with_github_api_truth(
        self,
    ) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-readiness-gh-") as raw:
            bin_dir = Path(raw)
            fake_gh = bin_dir / "gh"
            fake_gh.write_text(
                """#!/usr/bin/env python3
import json
import sys

endpoint = sys.argv[-1]
if endpoint.endswith("/releases/tags/v0.54.0"):
    print(json.dumps({
        "tag_name": "v0.54.0",
        "draft": False,
        "immutable": True,
        "assets": [
            {"name": "reflaxe.go-0.54.0.zip", "digest": "sha256:" + "a" * 64},
            {"name": "reflaxe.go-0.54.0.zip.sha256", "digest": "sha256:" + "b" * 64},
            {"name": "reflaxe.go-0.54.0.manifest.json", "digest": "sha256:" + "c" * 64},
            {"name": "reflaxe.go-0.54.0.provenance.json", "digest": "sha256:" + "d" * 64}
        ]
    }))
elif endpoint.endswith("/git/ref/tags/v0.54.0"):
    print(json.dumps({
        "object": {"type": "commit", "sha": "1" * 40}
    }))
else:
    print("unexpected endpoint: " + endpoint, file=sys.stderr)
    raise SystemExit(2)
""",
                encoding="utf-8",
            )
            fake_gh.chmod(0o755)
            evidence = self.evidence()
            evidence["github"]["apiAuthoritative"] = False
            environment = os.environ.copy()
            environment["PATH"] = f"{bin_dir}{os.pathsep}{environment['PATH']}"
            result = self.run_evidence(evidence, mode="live", env=environment)
            self.assertEqual(0, result.returncode, result.stdout + result.stderr)

    def test_schema_rejects_unknown_top_level_truth(self) -> None:
        evidence = copy.deepcopy(self.evidence())
        evidence["assumeReady"] = True
        self.assert_rejected(evidence, "evidence fields must be exactly")

    def test_collector_builds_candidate_evidence_from_governed_sources(
        self,
    ) -> None:
        evidence = self.evidence()
        with tempfile.TemporaryDirectory(prefix="haxe-go-readiness-collect-") as raw:
            root = Path(raw)
            assets = root / "release-assets.json"
            assets.write_text(
                json.dumps(
                    {
                        "schemaVersion": 1,
                        "tag": evidence["release"]["tag"],
                        "sourceSha": evidence["release"]["testedSha"],
                        "assets": [
                            {
                                "name": asset["name"],
                                "path": asset["name"],
                                "size": 1,
                                "digest": asset["digest"],
                            }
                            for asset in evidence["artifacts"]["assets"]
                        ],
                    }
                ),
                encoding="utf-8",
            )
            output = root / "candidate.json"
            upstream = root / "upstream.json"
            upstream.write_text(
                json.dumps(
                    {
                        "schemaVersion": 1,
                        "kind": "haxe.go-upstream-gate-evidence",
                        "testedSha": evidence["release"]["testedSha"],
                        "publicApi": {"result": "pass"},
                        "platform": evidence["platform"],
                        "toolchains": evidence["toolchains"],
                        "security": evidence["security"],
                    }
                ),
                encoding="utf-8",
            )
            blockers = root / "blockers.json"
            blockers.write_text(
                json.dumps(evidence["blockers"]), encoding="utf-8"
            )
            collected = subprocess.run(
                [
                    "python3",
                    str(COLLECTOR),
                    "--phase",
                    "candidate",
                    "--version",
                    evidence["release"]["version"],
                    "--tag",
                    evidence["release"]["tag"],
                    "--tested-sha",
                    evidence["release"]["testedSha"],
                    "--upstream-evidence",
                    str(upstream),
                    "--blocker-evidence",
                    str(blockers),
                    "--assets",
                    str(assets),
                    "--output",
                    str(output),
                ],
                cwd=ROOT,
                text=True,
                capture_output=True,
                timeout=30,
            )
            self.assertEqual(
                0, collected.returncode, collected.stdout + collected.stderr
            )
            candidate = json.loads(output.read_text(encoding="utf-8"))
            self.assertEqual("candidate", candidate["phase"])
            self.assertIsNone(candidate["github"])
            self.assertRegex(
                candidate["compatibility"]["surfaceSha256"], r"^[0-9a-f]{64}$"
            )
            self.assertGreater(
                candidate["compatibility"]["admittedOperationCount"], 0
            )
            verified = self.run_evidence(candidate)
            self.assertEqual(
                0, verified.returncode, verified.stdout + verified.stderr
            )

    def test_collector_rejects_gate_sha_or_asset_identity_mismatch(self) -> None:
        evidence = self.evidence()
        with tempfile.TemporaryDirectory(prefix="haxe-go-readiness-collect-") as raw:
            root = Path(raw)
            assets = root / "release-assets.json"
            assets.write_text(
                json.dumps(
                    {
                        "schemaVersion": 1,
                        "tag": evidence["release"]["tag"],
                        "sourceSha": evidence["release"]["testedSha"],
                        "assets": [],
                    }
                ),
                encoding="utf-8",
            )
            upstream = root / "upstream.json"
            upstream.write_text(
                json.dumps(
                    {
                        "schemaVersion": 1,
                        "kind": "haxe.go-upstream-gate-evidence",
                        "testedSha": "2" * 40,
                        "publicApi": {"result": "pass"},
                        "platform": evidence["platform"],
                        "toolchains": evidence["toolchains"],
                        "security": evidence["security"],
                    }
                ),
                encoding="utf-8",
            )
            blockers = root / "blockers.json"
            blockers.write_text(
                json.dumps(evidence["blockers"]), encoding="utf-8"
            )
            result = subprocess.run(
                [
                    "python3",
                    str(COLLECTOR),
                    "--phase",
                    "candidate",
                    "--version",
                    evidence["release"]["version"],
                    "--tag",
                    evidence["release"]["tag"],
                    "--tested-sha",
                    evidence["release"]["testedSha"],
                    "--upstream-evidence",
                    str(upstream),
                    "--blocker-evidence",
                    str(blockers),
                    "--assets",
                    str(assets),
                    "--output",
                    str(root / "candidate.json"),
                ],
                cwd=ROOT,
                text=True,
                capture_output=True,
                timeout=30,
            )
            self.assertNotEqual(0, result.returncode)
            self.assertIn("upstream evidence SHA", result.stderr)

    def test_upstream_collector_records_workflow_results_and_resolved_tools(
        self,
    ) -> None:
        tested_sha = "1" * 40
        with tempfile.TemporaryDirectory(prefix="haxe-go-upstream-evidence-") as raw:
            output = Path(raw) / "upstream.json"
            command = [
                "python3",
                str(UPSTREAM_COLLECTOR),
                "--tested-sha",
                tested_sha,
                "--quality-result",
                "success",
                "--gitleaks-result",
                "success",
                "--dependency-audit-result",
                "success",
                "--go-tooling-result",
                "success",
                "--haxe-version",
                "4.3.7",
                "--node-version",
                "v24.11.1",
                "--runner-image-os",
                "ubuntu24",
                "--runner-image-version",
                "20260720.1.0",
                "--output",
                str(output),
            ]
            result = subprocess.run(
                command,
                cwd=ROOT,
                text=True,
                capture_output=True,
                timeout=30,
            )
            self.assertEqual(0, result.returncode, result.stdout + result.stderr)
            evidence = json.loads(output.read_text(encoding="utf-8"))
            self.assertEqual(tested_sha, evidence["testedSha"])
            self.assertEqual({"result": "pass"}, evidence["publicApi"])
            self.assertEqual(
                {
                    "id": "linux-amd64",
                    "os": "linux",
                    "architecture": "amd64",
                    "runnerImage": {
                        "os": "ubuntu24",
                        "version": "20260720.1.0",
                    },
                },
                evidence["platform"],
            )
            self.assertEqual(
                {"resolved": ["24.11.1"]}, evidence["toolchains"]["node"]
            )
            self.assertTrue(
                all(
                    gate["result"] == "pass"
                    and gate["testedSha"] == tested_sha
                    for gate in evidence["security"]["gates"]
                )
            )
            self.assertNotIn(
                "github-governance-live",
                {gate["id"] for gate in evidence["security"]["gates"]},
            )

            command[command.index("success", command.index("--gitleaks-result"))] = (
                "failure"
            )
            command[-1] = str(Path(raw) / "failed.json")
            failed = subprocess.run(
                command,
                cwd=ROOT,
                text=True,
                capture_output=True,
                timeout=30,
            )
            self.assertNotEqual(0, failed.returncode)
            self.assertIn("upstream workflow gate did not succeed", failed.stderr)


if __name__ == "__main__":
    unittest.main()
