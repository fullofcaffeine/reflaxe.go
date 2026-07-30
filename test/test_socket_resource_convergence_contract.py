#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "socket-resource-convergence.md"
COMPATIBILITY = ROOT / "docs" / "compatibility-support-source.json"
NATIVE_TEST = ROOT / "runtime" / "hxrt" / "socket_resource_convergence_posix_test.go"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketResourceConvergenceContractTest(unittest.TestCase):
    def test_native_suite_measures_repeated_failure_and_cleanup(self) -> None:
        native = NATIVE_TEST.read_text(encoding="utf-8")
        for phrase in (
            "const attempts = 20",
            "runtime.NumGoroutine()",
            'os.ReadDir("/proc/self/fd")',
            "activeConnections",
            "SocketSetTimeout",
            "SslSocketConnect",
            "SocketUdpSendTo",
            "SocketSelect",
            "SocketClose",
            "resources did not converge",
        ):
            self.assertIn(phrase, native)

    def test_pre_review_compatibility_state_is_evidenced_and_fail_closed(self) -> None:
        source = json.loads(COMPATIBILITY.read_text(encoding="utf-8"))
        networking = next(
            item for item in source["surfaces"] if item["id"] == "portable-networking"
        )
        operations = {item["id"]: item for item in networking["operations"]}

        for operation_id in (
            "tcp-server-and-listener-controls",
            "socket-timeout-nonblocking-readiness-controls",
            "socket-shutdown-fast-send-controls",
            "udp-ipv4",
            "tls-socket",
        ):
            operation = operations[operation_id]
            self.assertIn(
                "runtime:socket-resource-convergence-posix",
                operation["evidence_ids"],
            )
            self.assertFalse(operation["release_admitted"])
            self.assertEqual(["haxe_go-vfp.10.9"], operation["blockers"])
            self.assertIn(
                "resource convergence",
                operation["qualification"].lower(),
            )

        evidence = {
            item["id"]: item
            for item in source["evidence"]
        }
        self.assertEqual(
            "runtime/hxrt/socket_resource_convergence_posix_test.go",
            evidence["runtime:socket-resource-convergence-posix"]["path"],
        )
        self.assertEqual(
            "docs/socket-resource-convergence.md",
            evidence["policy:socket-resource-convergence"]["path"],
        )

    def test_documentation_separates_measurement_from_admission(self) -> None:
        decision = " ".join(DECISION.read_text(encoding="utf-8").split())
        for phrase in (
            "20 repetitions",
            "active accepted connections",
            "goroutine count",
            "Linux file-descriptor count",
            "stalled TLS handshake",
            "TCP reset",
            "UDP",
            "readiness",
            "Windows remains compile-only",
            "does not admit",
            "independent commit-pinned review",
        ):
            self.assertIn(phrase, decision)

        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_resource_convergence_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
