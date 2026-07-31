#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SOURCE = ROOT / "docs" / "compatibility-support-source.json"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketAdmissionSplitContractTest(unittest.TestCase):
    def operations(self) -> dict[str, dict[str, object]]:
        source = json.loads(SOURCE.read_text(encoding="utf-8"))
        networking = next(
            surface
            for surface in source["surfaces"]
            if surface["id"] == "portable-networking"
        )
        return {operation["id"]: operation for operation in networking["operations"]}

    def test_advanced_networking_is_split_by_exact_member_contract(self) -> None:
        operations = self.operations()
        removed_broad_rows = {
            "tcp-server-and-listener-controls",
            "socket-timeout-nonblocking-readiness-controls",
            "socket-shutdown-fast-send-controls",
            "host-dns-and-reverse-lookup",
            "udp-ipv4",
            "tls-socket",
        }
        self.assertTrue(removed_broad_rows.isdisjoint(operations))

        admitted = {
            "tcp-ipv4-blocking-server-core",
            "tcp-ipv4-connected-timeout-nonblocking-controls",
            "tcp-ipv4-connected-nonblocking-accept",
            "tcp-ipv4-readiness-controls",
            "plain-tcp-shutdown-fast-send-controls",
            "host-localhost",
            "udp-ipv4-datagram-core",
            "udp-ipv4-broadcast-control",
            "tls-ipv4-direct-client",
            "tls-ipv4-server-sni",
            "tls-shutdown-fast-send-controls",
        }
        for operation_id in admitted:
            operation = operations[operation_id]
            self.assertTrue(operation["release_admitted"], operation_id)
            self.assertFalse(operation["blockers"], operation_id)
            self.assertIn("Linux/amd64", operation["qualification"], operation_id)

        for operation_id in ("host-named-resolution", "host-reverse-lookup"):
            operation = operations[operation_id]
            self.assertFalse(operation["release_admitted"], operation_id)
            self.assertEqual("experimental", operation["state"])
            self.assertFalse(operation["blockers"], operation_id)

    def test_whole_udp_and_tls_classes_are_never_admitted_as_symbols(self) -> None:
        operations = self.operations()
        symbols = {
            symbol
            for operation in operations.values()
            for symbol in operation["symbols"]
        }
        self.assertNotIn("sys.net.UdpSocket", symbols)
        self.assertNotIn("sys.ssl.Socket", symbols)
        self.assertEqual(
            [
                "new sys.net.UdpSocket",
                "sys.net.UdpSocket.bind (numeric IPv4)",
                "sys.net.UdpSocket.host",
                "sys.net.UdpSocket.sendTo (blocking IPv4)",
                "sys.net.UdpSocket.readFrom (blocking IPv4)",
                "sys.net.UdpSocket.close",
            ],
            operations["udp-ipv4-datagram-core"]["symbols"],
        )
        self.assertEqual(
            [
                "new sys.ssl.Socket",
                "sys.ssl.Socket.verifyCert/setCA/setHostname",
                "sys.ssl.Socket.connect (numeric IPv4 route with logical host)",
                "sys.ssl.Socket.handshake",
                "sys.ssl.Socket.input.readByte/readBytes (blocking)",
                "sys.ssl.Socket.output.writeByte/writeBytes (blocking)",
                "sys.ssl.Socket.output.flush",
                "sys.ssl.Socket.peerCertificate",
                "sys.ssl.Socket.close",
            ],
            operations["tls-ipv4-direct-client"]["symbols"],
        )

    def test_exact_evidence_and_exclusions_follow_the_review(self) -> None:
        operations = self.operations()
        self.assertIn(
            "snapshot:socket-udp",
            operations["udp-ipv4-datagram-core"]["evidence_ids"],
        )
        self.assertIn(
            "semantic:host-basic",
            operations["host-localhost"]["evidence_ids"],
        )
        readiness_exclusions = " ".join(
            operations["tcp-ipv4-readiness-controls"]["exclusions"]
        )
        self.assertIn("FdSet", readiness_exclusions)
        self.assertIn("TLS readiness", readiness_exclusions)
        self.assertIn("nonblocking connect", readiness_exclusions.lower())
        baseline = operations["tcp-ipv4-blocking-client-core"]
        self.assertIn("already-installed connection", baseline["qualification"])
        self.assertTrue(
            any("in-progress connect" in item for item in baseline["exclusions"])
        )

    def test_release_runner_executes_this_contract(self) -> None:
        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_admission_split_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
