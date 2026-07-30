#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DECISION = ROOT / "docs" / "socket-server-lifecycle.md"
COMPATIBILITY = ROOT / "docs" / "compatibility-support-source.json"
RUNTIME = ROOT / "runtime" / "hxrt" / "socket.go"
TLS_RUNTIME = ROOT / "runtime" / "hxrt" / "socket_ssl.go"
TLS_SOURCE = ROOT / "std" / "go" / "_std" / "sys" / "ssl" / "Socket.hx"
FEATURE_ANALYZER = (
    ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
)
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class SocketServerLifecycleContractTest(unittest.TestCase):
    def test_runtime_keeps_bind_and_listen_as_separate_typed_transitions(self) -> None:
        runtime = RUNTIME.read_text(encoding="utf-8")
        bind = runtime.split("func SocketBindTCP(", 1)[1].split(
            "func SocketListen(", 1
        )[0]
        listen = runtime.split("func SocketListen(", 1)[1].split(
            "func SocketAccept(", 1
        )[0]

        self.assertIn("socketBindTCP(handle, host, port, nil)", bind)
        self.assertNotIn("net.Listen(", bind)
        self.assertIn("bound.Listen(backlog)", listen)
        self.assertIn("socketRelistenTCP(deadlineListener, backlog)", listen)
        self.assertIn("backlog < 0", listen)
        self.assertIn("socket listen requires a bound socket", listen)

    def test_tls_bind_defers_listener_wrapping_until_inherited_listen(self) -> None:
        native = TLS_RUNTIME.read_text(encoding="utf-8")
        staged = TLS_SOURCE.read_text(encoding="utf-8")
        bind = native.split("func SslSocketBind(", 1)[1].split(
            "func SslSocketHandshake(", 1
        )[0]

        self.assertIn("socketBindTCP(handle, host, port", bind)
        self.assertIn("tls.NewListener(listener, config)", bind)
        self.assertNotIn("net.ListenTCP(", bind)
        self.assertIn("NativeSslSocket.bind(handle", staged)
        self.assertNotIn("NativeSslSocket.listen(handle", staged)

    def test_platform_adapters_are_selected_with_the_socket_runtime(self) -> None:
        analyzer = FEATURE_ANALYZER.read_text(encoding="utf-8")
        for name in (
            "socket_listener_posix.go",
            "socket_listener_unsupported.go",
            "socket_listener_windows.go",
        ):
            self.assertIn(f'"{name}"', analyzer)
            self.assertTrue((ROOT / "runtime" / "hxrt" / name).is_file())

    def test_compatibility_claim_stays_operation_scoped_and_fail_closed(self) -> None:
        source = json.loads(COMPATIBILITY.read_text(encoding="utf-8"))
        networking = next(
            item for item in source["surfaces"] if item["id"] == "portable-networking"
        )
        operation = next(
            item
            for item in networking["operations"]
            if item["id"] == "tcp-server-and-listener-controls"
        )

        self.assertEqual("semantic-diff-supported", operation["state"])
        self.assertFalse(operation["release_admitted"])
        self.assertEqual(["haxe_go-vfp.10.9"], operation["blockers"])
        self.assertIn("bind-then-listen", operation["qualification"])
        self.assertIn("Windows remains compile-only", operation["qualification"])
        for evidence in (
            "semantic:socket-server-lifecycle",
            "snapshot:socket-server-lifecycle",
            "runtime:socket-listener-linux",
            "policy:socket-server-lifecycle",
        ):
            self.assertIn(evidence, operation["evidence_ids"])

        exclusions = " ".join(operation["exclusions"])
        self.assertIn("Nonblocking accept", exclusions)
        self.assertIn("runtime support outside Linux/amd64", exclusions)
        self.assertIn("kernel", exclusions)

    def test_documentation_and_release_runner_keep_the_boundary_executable(self) -> None:
        decision = " ".join(DECISION.read_text(encoding="utf-8").split())
        for phrase in (
            "`socket.bind(host, port)` reserves",
            "`socket.listen(connections)` starts listening",
            "does not accept connections",
            "operating system may cap or normalize",
            "Windows cross-build is compile evidence",
            "haxe_go-vfp.10.9.6",
        ):
            self.assertIn(phrase, decision)

        runner = RELEASE_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_socket_server_lifecycle_contract.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
