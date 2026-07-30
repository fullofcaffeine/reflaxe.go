#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "socket-tls-controls.md"
COMPATIBILITY = ROOT / "docs" / "compatibility-support-source.json"
RUNTIME = ROOT / "runtime" / "hxrt" / "socket.go"
STAGED_TLS = ROOT / "std" / "go" / "_std" / "sys" / "ssl" / "Socket.hx"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketTLSControlsContractTest(unittest.TestCase):
    def test_runtime_uses_typed_tls_and_tcp_control_boundaries(self) -> None:
        runtime = RUNTIME.read_text(encoding="utf-8")
        shutdown = runtime.split("func SocketShutdown(", 1)[1].split(
            "func SocketPeer(", 1
        )[0]
        fast_send = runtime.split("func SocketSetFastSend(", 1)[1].split(
            "func SocketReadValues(", 1
        )[0]

        self.assertIn("conn.(socketNestedConnection)", shutdown)
        self.assertIn("conn.(socketCloseWriter)", shutdown)
        self.assertIn('"TLS read-only shutdown is unsupported"', shutdown)
        self.assertIn("SocketClose(handle)", shutdown)
        self.assertIn("socketTCPConnection(handle.conn)", runtime)
        self.assertIn("tcpConn.SetNoDelay(handle.fastSend)", runtime)
        self.assertIn("applyFastSendLocked(true)", fast_send)
        self.assertNotIn("any", shutdown)

    def test_staged_tls_documents_inherited_public_behavior(self) -> None:
        staged = STAGED_TLS.read_text(encoding="utf-8")
        for phrase in (
            "override public function shutdown",
            "TLS write shutdown must send a protocol `close_notify`",
            "read-only raises a deterministic unsupported error",
            "override public function setFastSend",
            "updates the underlying TCP socket",
        ):
            self.assertIn(phrase, staged)

    def test_compatibility_claim_is_precise_and_remains_fail_closed(self) -> None:
        source = json.loads(COMPATIBILITY.read_text(encoding="utf-8"))
        networking = next(
            item for item in source["surfaces"] if item["id"] == "portable-networking"
        )
        controls = next(
            item
            for item in networking["operations"]
            if item["id"] == "socket-shutdown-fast-send-controls"
        )
        tls = next(
            item for item in networking["operations"] if item["id"] == "tls-socket"
        )

        self.assertEqual("compile-go-test-run-supported", controls["state"])
        self.assertFalse(controls["release_admitted"])
        self.assertEqual(["haxe_go-vfp.10.9"], controls["blockers"])
        for evidence in (
            "semantic:socket",
            "snapshot:socket-tls",
            "runtime:socket-tls-controls-posix",
            "runtime:socket-resource-convergence-posix",
            "policy:socket-tls-controls",
            "policy:socket-resource-convergence",
        ):
            self.assertIn(evidence, controls["evidence_ids"])

        qualification = controls["qualification"]
        for phrase in (
            "TLS close_notify",
            "TLS read-only shutdown is explicitly unsupported",
            "TCP_NODELAY",
            "Release admission stays fail-closed",
        ):
            self.assertIn(phrase, qualification)
        exclusions = " ".join(controls["exclusions"])
        for phrase in (
            "TLS read-only shutdown",
            "production-soak",
            "Windows runtime",
        ):
            self.assertIn(phrase, exclusions)

        self.assertIn("runtime:socket-tls-controls-posix", tls["evidence_ids"])
        self.assertNotIn("inherited control semantics", tls["qualification"])

    def test_documentation_explains_behavior_proof_and_boundary(self) -> None:
        decision = " ".join(DECISION.read_text(encoding="utf-8").split())
        for phrase in (
            "TLS `close_notify`",
            "read side remains available",
            "`TLS read-only shutdown is unsupported`",
            "`TCP_NODELAY`",
            "not a no-throw test",
            "concurrent close",
            "Windows remains compile-only",
            "does not admit these controls for release",
        ):
            self.assertIn(phrase, decision)

        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_tls_controls_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
