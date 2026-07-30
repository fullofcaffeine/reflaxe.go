#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "socket-dns-boundary.md"
COMPATIBILITY = ROOT / "docs" / "compatibility-support-source.json"
READINESS = ROOT / "release" / "readiness-policy.json"
HOST_SOURCE = ROOT / "std" / "go" / "_std" / "sys" / "net" / "Host.hx"
SOCKET_SOURCE = ROOT / "std" / "go" / "_std" / "sys" / "net" / "Socket.hx"
RUNTIME = ROOT / "runtime" / "hxrt" / "socket.go"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketDnsBoundaryContractTest(unittest.TestCase):
    def network_operations(self) -> dict[str, dict[str, object]]:
        source = json.loads(COMPATIBILITY.read_text(encoding="utf-8"))
        networking = next(
            surface
            for surface in source["surfaces"]
            if surface["id"] == "portable-networking"
        )
        return {operation["id"]: operation for operation in networking["operations"]}

    def test_upstream_eager_host_construction_remains_the_source_contract(self) -> None:
        host = HOST_SOURCE.read_text(encoding="utf-8")
        socket = SOCKET_SOURCE.read_text(encoding="utf-8")
        runtime = RUNTIME.read_text(encoding="utf-8")

        constructor = host.split("public function new(name:String):Void", 1)[1].split(
            "public function toString", 1
        )[0]
        self.assertIn("host = name;", constructor)
        self.assertIn("ip = NativeSocket.hostResolve(name);", constructor)
        self.assertNotIn("SocketHandle", constructor)
        self.assertNotIn("setTimeout", constructor)

        connect = socket.split("public function connect(host:Host, port:Int):Void", 1)[
            1
        ].split("public function listen", 1)[0]
        self.assertIn("host.toString()", connect)
        self.assertNotIn("hostResolve", connect)

        self.assertIn("func HostResolve(name *string) int", runtime)
        self.assertIn("net.LookupIP(name)", runtime)
        self.assertIn(
            "// SocketSetTimeout updates the deadline policy; a negative value clears it.",
            runtime,
        )

    def test_decision_names_every_phase_and_timeout_value(self) -> None:
        text = DECISION.read_text(encoding="utf-8")
        normalized = " ".join(text.split())
        for phrase in (
            "Decision: preserve eager `Host` resolution",
            "`new Host(\"hostname\")`",
            "`Host.reverse()`",
            "`Host.localhost()`",
            "positive `Socket.setTimeout`",
            "`setTimeout(0)`",
            "negative timeout",
            "does not retroactively bound",
            "application-controlled numeric IPv4",
            "haxe_go-vfp.10.9",
        ):
            self.assertIn(phrase, normalized)

    def test_compatibility_and_readiness_cannot_admit_dns_accidentally(self) -> None:
        operations = self.network_operations()
        tcp = operations["tcp-ipv4-blocking-client-core"]
        controls = operations["socket-timeout-nonblocking-readiness-controls"]
        dns = operations["host-dns-and-reverse-lookup"]

        self.assertTrue(tcp["release_admitted"])
        self.assertIn("pre-resolved numeric endpoint", tcp["qualification"])
        self.assertTrue(
            any("Hostname and DNS behavior" in item for item in tcp["exclusions"])
        )

        self.assertFalse(controls["release_admitted"])
        self.assertIn("does not bound Host construction", controls["qualification"])

        self.assertEqual("experimental", dns["state"])
        self.assertFalse(dns["release_admitted"])
        self.assertEqual(["haxe_go-vfp.10.9"], dns["blockers"])
        self.assertIn("synchronously during construction", dns["qualification"])
        self.assertIn("before a Socket exists", dns["qualification"])
        self.assertIn("policy:socket-dns-boundary", dns["evidence_ids"])
        exclusions = " ".join(dns["exclusions"])
        self.assertIn("timeout", exclusions)
        self.assertIn("cancellation", exclusions)
        self.assertIn("reverse lookup", exclusions)
        self.assertIn("resolver configuration", exclusions)

        readiness = json.loads(READINESS.read_text(encoding="utf-8"))[
            "compatibility"
        ]
        self.assertEqual(
            "haxe_go-vfp.10.9",
            readiness["requiredExclusions"]["portable-socket-advanced"],
        )
        self.assertEqual(
            "portable-socket-advanced",
            readiness["blockerScopes"]["haxe_go-vfp.10.9"],
        )

    def test_release_contract_runner_keeps_the_decision_executable(self) -> None:
        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_dns_boundary_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
