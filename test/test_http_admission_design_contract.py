#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DESIGN = ROOT / "docs" / "http-client-admission-design.md"
FEATURE_MATRIX = ROOT / "docs" / "feature-support-matrix.md"


class HttpAdmissionDesignContractTest(unittest.TestCase):
    def text(self) -> str:
        self.assertTrue(DESIGN.is_file(), f"missing HTTP admission design: {DESIGN}")
        return DESIGN.read_text(encoding="utf-8")

    def test_decision_leads_with_the_user_visible_outcome_and_authority(self) -> None:
        text = self.text()
        for heading in (
            "# Portable HTTP client admission design",
            "## Outcome",
            "## Why the old seam was insufficient",
            "## Source and native ownership",
            "## Request contract",
            "## Response contract",
            "## Upload contract",
            "## Error precedence",
            "## Timeout contract",
            "## Second-pass challenge review",
            "## Admission boundary",
        ):
            self.assertIn(heading, text)

        self.assertIn("Haxe 4.3.7", text)
        self.assertIn("staged `sys.Http`", text)
        self.assertIn("generated Haxe is never called from a native goroutine", text)
        self.assertIn("does not require a universal compiler IR", text)

    def test_typed_boundary_and_lifetimes_are_explicit(self) -> None:
        text = self.text()
        for capability in (
            "`HttpRequestHandle`",
            "`HttpExchangeHandle`",
            "`HttpUploadSinkHandle`",
            "`HttpReadResultHandle`",
            "`startExchange`",
            "`awaitResponse`",
            "`writeUploadChunk`",
            "`finishUpload`",
            "`readResponseChunk`",
            "`closeExchange`",
            "`cancelExchange`",
        ):
            self.assertIn(capability, text)

        self.assertIn("one owner", text)
        self.assertIn("bounded response chunk", text)
        self.assertIn("first terminal event wins", text)

    def test_sequences_and_blocking_input_limit_are_not_implicit(self) -> None:
        text = self.text()
        for event in (
            "headers -> `onStatus` -> `prepare` -> body writes",
            "complete body -> status classification -> Output close",
            "partial body -> transfer error -> `onError`",
            "`request()` retains every byte already written",
            "`cnxTimeout < 0`",
            "`cnxTimeout == 0`",
            "`cnxTimeout > 0`",
        ):
            self.assertIn(event, text)

        self.assertIn("arbitrary custom `Input.readBytes`", text)
        self.assertIn("cannot be forcibly interrupted", text)
        self.assertIn("explicitly excluded", text)
        self.assertIn("no reads after public return", text)
        self.assertIn("Abort never waits for the sink's writer mutex", text)

    def test_every_oracle_finding_has_one_implementation_owner(self) -> None:
        text = self.text()
        expected = {
            "B1": "haxe_go-vfp.10.8.3",
            "B2": "haxe_go-vfp.10.8.4",
            "H1": "haxe_go-vfp.10.8.2",
            "H2": "haxe_go-vfp.10.8.2",
            "H3": "haxe_go-vfp.10.8.5",
            "L2": "haxe_go-vfp.10.8.2",
        }
        for finding, owner in expected.items():
            with self.subTest(finding=finding):
                self.assertIn(f"| {finding} |", text)
                self.assertIn(f"`{owner}`", text)

        self.assertIn("haxe_go-vfp.10.8.6", text)
        self.assertIn("HTTP remains release-excluded", text)

    def test_implementation_inventory_links_the_canonical_design(self) -> None:
        matrix = FEATURE_MATRIX.read_text(encoding="utf-8")
        self.assertIn(
            "[portable HTTP admission design](http-client-admission-design.md)",
            matrix,
        )

    def test_final_review_candidate_is_operation_split_and_fail_closed(self) -> None:
        text = self.text()
        for value in (
            "## Resource convergence evidence",
            "## Candidate operation-level disposition",
            "`http-ipv4-blocking-client-core`",
            "`http-ipv4-multipart-upload`",
            "`http-data-url-client`",
            "`http-proxy-and-custom-transport`",
            "`https-client`",
            "zero active server connections",
            "Linux/amd64 runtime evidence",
            "Windows compile-only evidence",
            "Release admission remains fail-closed",
            "commit-pinned independent review",
        ):
            self.assertIn(value, text)


if __name__ == "__main__":
    unittest.main()
