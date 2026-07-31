#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "socket-readiness-nonblocking.md"
COMPATIBILITY = ROOT / "docs" / "compatibility-support-source.json"
RUNTIME = ROOT / "runtime" / "hxrt" / "socket.go"
FEATURE_ANALYZER = (
    ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
)
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketReadinessContractTest(unittest.TestCase):
    def test_runtime_delegates_to_typed_native_readiness(self) -> None:
        runtime = RUNTIME.read_text(encoding="utf-8")
        select = runtime.split("func SocketSelect(", 1)[1].split(
            "func SocketWaitForRead(", 1
        )[0]

        self.assertIn("socketSelectOnce(read, write, others, wait)", select)
        self.assertIn("socketSelectNative(socketNativeSelectRequest{", runtime)
        self.assertIn("socketSelectIndexes(", runtime)
        self.assertIn("socketDuplicateDescriptor(value)", runtime)
        self.assertIn("defer socketReleaseReadinessSnapshots(cache)", runtime)
        self.assertIn("reader.Buffered() > 0", runtime)
        self.assertNotIn("socketConnected", runtime)
        self.assertNotIn("pollRead", runtime)

    def test_platform_adapters_are_packaged_with_socket_runtime(self) -> None:
        analyzer = FEATURE_ANALYZER.read_text(encoding="utf-8")
        for name in (
            "socket_readiness_darwin.go",
            "socket_readiness_linux_32.go",
            "socket_readiness_linux_64.go",
            "socket_readiness_unsupported.go",
        ):
            self.assertIn(f'"{name}"', analyzer)
            self.assertTrue((ROOT / "runtime" / "hxrt" / name).is_file())
        linux = (
            ROOT / "runtime" / "hxrt" / "socket_readiness_linux_64.go"
        ).read_text(encoding="utf-8")
        self.assertIn("syscall.F_DUPFD_CLOEXEC", linux)

    def test_compatibility_claim_splits_connected_controls_from_readiness(self) -> None:
        source = json.loads(COMPATIBILITY.read_text(encoding="utf-8"))
        networking = next(
            item for item in source["surfaces"] if item["id"] == "portable-networking"
        )
        operations = {item["id"]: item for item in networking["operations"]}
        operation = operations["tcp-ipv4-readiness-controls"]
        connected = operations["tcp-ipv4-connected-timeout-nonblocking-controls"]
        shutdown = operations["plain-tcp-shutdown-fast-send-controls"]

        self.assertEqual("semantic-diff-supported", operation["state"])
        self.assertTrue(operation["release_admitted"])
        self.assertEqual([], operation["blockers"])
        self.assertNotIn("sys.net.Socket.shutdown", operation["symbols"])
        self.assertNotIn("sys.net.Socket.setFastSend", operation["symbols"])
        for evidence in (
            "semantic:socket-readiness",
            "snapshot:socket-readiness",
            "runtime:socket-readiness-posix",
            "runtime:socket-cross-build",
            "policy:socket-readiness",
        ):
            self.assertIn(evidence, operation["evidence_ids"])

        qualification = operation["qualification"]
        self.assertIn("read/write/exception readiness", qualification)
        self.assertIn("snapshot preparation", qualification)
        exclusions = " ".join(operation["exclusions"])
        for phrase in (
            "Nonblocking connect",
            "TLS readiness",
            "FdSet",
        ):
            self.assertIn(phrase, exclusions)

        self.assertTrue(connected["release_admitted"])
        self.assertNotIn("sys.net.Socket.select", connected["symbols"])

        self.assertEqual(
            [
                "sys.net.Socket.shutdown (established plain TCP)",
                "sys.net.Socket.setFastSend (established plain TCP)",
            ],
            shutdown["symbols"],
        )
        self.assertEqual("compile-go-test-run-supported", shutdown["state"])
        self.assertTrue(shutdown["release_admitted"])
        self.assertEqual([], shutdown["blockers"])
        self.assertIn("runtime:socket-tls-controls-posix", shutdown["evidence_ids"])
        self.assertIn("policy:socket-tls-controls", shutdown["evidence_ids"])

    def test_documentation_explains_behavior_and_release_boundary(self) -> None:
        decision = " ".join(DECISION.read_text(encoding="utf-8").split())
        for phrase in (
            "read readiness",
            "write readiness",
            "exceptional readiness",
            "buffered bytes",
            "partial write",
            "`haxe.io.Error.Blocked`",
            "Nonblocking connect remains excluded",
            "TLS readiness remains excluded",
            "Windows is compile-only evidence",
            "Linux/amd64",
        ):
            self.assertIn(phrase, decision)

        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_readiness_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
