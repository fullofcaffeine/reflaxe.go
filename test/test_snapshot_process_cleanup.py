#!/usr/bin/env python3
"""Exercise snapshot command ownership with real launcher and child processes."""

from __future__ import annotations

import importlib.util
import json
import os
from pathlib import Path
import signal
import socket
import subprocess
import sys
import tempfile
import time
import unittest

ROOT = Path(__file__).resolve().parent.parent
sys.path.insert(0, str(ROOT / "test"))
SPEC = importlib.util.spec_from_file_location("snapshot_process_contract", ROOT / "test" / "run-snapshots.py")
assert SPEC is not None and SPEC.loader is not None
RUNNER = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = RUNNER
SPEC.loader.exec_module(RUNNER)

# Each process owns a listener. Connection refusal observes that its live resource
# has closed, without mistaking a terminated but unreaped process for a live one.
LAUNCHER = r'''
import json, os, pathlib, socket, subprocess, sys, threading, time
root = pathlib.Path(sys.argv[1])
listener = socket.socket()
listener.bind(("127.0.0.1", 0))
listener.listen()
def serve():
    while True:
        connection, _ = listener.accept()
        with connection:
            connection.sendall(b"alive")
threading.Thread(target=serve, daemon=True).start()
kind = sys.argv[2]
record = root / (kind + ".json")
record.with_suffix(".pending").write_text(json.dumps({"pid": os.getpid(), "port": listener.getsockname()[1]}))
record.with_suffix(".pending").replace(record)
if kind == "child":
    time.sleep(60)
else:
    child = subprocess.Popen([sys.executable, __file__, str(root), "child"])
    deadline = time.monotonic() + 10
    while not (root / "child.json").exists():
        if child.poll() is not None or time.monotonic() > deadline:
            raise RuntimeError("child did not become ready")
        time.sleep(0.01)
    print("launcher ready", flush=True)
    print("launcher diagnostic", file=sys.stderr, flush=True)
    if kind == "exited":
        sys.exit(0)
    child.wait()
'''


def listener_open(port: int) -> bool:
    with socket.socket() as probe:
        probe.settimeout(0.1)
        try:
            probe.connect(("127.0.0.1", port))
            return probe.recv(5) == b"alive"
        except (ConnectionError, TimeoutError):
            return False


class SnapshotProcessCleanupTest(unittest.TestCase):
    def test_success_and_nonzero_results_preserve_command_cwd_environment_and_text(self) -> None:
        with tempfile.TemporaryDirectory(prefix="snapshot-command-") as raw:
            root = Path(raw)
            script = "import os,sys; print(os.getcwd()); print(os.environ['SNAPSHOT_CONTRACT']); print('diagnostic', file=sys.stderr); sys.exit(7)"
            cmd = [sys.executable, "-c", script]
            result = RUNNER.run_command(cmd, root, 10, {"SNAPSHOT_CONTRACT": "owned"})
            self.assertEqual(cmd, result.args)
            self.assertEqual(7, result.returncode)
            directory, value = result.stdout.splitlines()
            # Windows may preserve an equivalent 8.3 spelling in os.getcwd().
            self.assertTrue(root.samefile(directory))
            self.assertEqual("owned", value)
            self.assertEqual("diagnostic\n", result.stderr)
            success = RUNNER.run_command([sys.executable, "-c", "print('ok')"], root, 10)
            self.assertEqual(0, success.returncode)
            self.assertEqual("ok\n", success.stdout)

    @unittest.skipUnless(os.name == "nt", "Windows DWORD exit status contract")
    def test_windows_preserves_unsigned_native_exit_status(self) -> None:
        with tempfile.TemporaryDirectory(prefix="snapshot-native-status-") as raw:
            cmd = [sys.executable, "-c", "import ctypes; ctypes.windll.kernel32.ExitProcess(0xC0000005)"]
            result = RUNNER.run_command(cmd, Path(raw), 10)
            self.assertEqual(0xC0000005, result.returncode)

    def test_spawn_errors_preserve_native_identity(self) -> None:
        with tempfile.TemporaryDirectory(prefix="snapshot-spawn-error-") as raw:
            root = Path(raw)
            for cmd, cwd in (
                ([str(root / "missing-executable")], root),
                ([sys.executable, "-c", "pass"], root / "missing-directory"),
            ):
                with self.subTest(command=cmd, cwd=cwd):
                    with self.assertRaises(OSError) as baseline:
                        subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=10)
                    with self.assertRaises(type(baseline.exception)) as actual:
                        RUNNER.run_command(cmd, cwd, 10)
                    for attribute in ("errno", "strerror", "filename", "filename2", "winerror"):
                        self.assertEqual(
                            getattr(baseline.exception, attribute, None),
                            getattr(actual.exception, attribute, None), attribute,
                        )

    @unittest.skipUnless(os.name in ("posix", "nt"), "native process ownership contract")
    def test_timeout_stops_launcher_and_child(self) -> None:
        self.check_timeout_tree("launcher")

    @unittest.skipUnless(os.name in ("posix", "nt"), "native process ownership contract")
    def test_timeout_stops_child_after_launcher_exits_with_inherited_pipes(self) -> None:
        self.check_timeout_tree("exited")

    @unittest.skipUnless(os.name in ("posix", "nt"), "native process ownership contract")
    def test_timeout_does_not_stop_an_unrelated_process(self) -> None:
        with tempfile.TemporaryDirectory(prefix="snapshot-outside-tree-") as raw:
            root = Path(raw)
            script = root / "launcher.py"
            script.write_text(LAUNCHER, encoding="utf-8")
            outside = subprocess.Popen(
                [sys.executable, str(script), str(root), "child"],
                stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL,
            )
            try:
                deadline = time.monotonic() + 10
                record_file = root / "child.json"
                while not record_file.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(record_file.exists(), "outside process never became ready")
                record = json.loads(record_file.read_text(encoding="utf-8"))
                self.check_timeout_tree("launcher")
                self.assertIsNone(outside.poll())
                self.assertTrue(listener_open(record["port"]), "unrelated process was stopped")
            finally:
                outside.kill()
                outside.wait(timeout=10)

    def check_timeout_tree(self, kind: str) -> None:
        with tempfile.TemporaryDirectory(prefix="snapshot-timeout-") as raw:
            root = Path(raw)
            script = root / "launcher.py"
            script.write_text(LAUNCHER, encoding="utf-8")
            cmd = [sys.executable, str(script), str(root), kind]
            records = []
            try:
                # Windows starts an additional Python job supervisor. This is
                # only the fixture's budget; production command budgets do not change.
                timeout_s = 10 if os.name == "nt" else 2
                if os.name == "nt" and kind == "exited":
                    # A completed Windows supervisor closes its job, reclaiming
                    # leftover children while preserving the launcher's status.
                    result = RUNNER.run_command(cmd, root, timeout_s)
                    self.assertEqual(0, result.returncode)
                    self.assertIn("launcher ready", result.stdout)
                    self.assertIn("launcher diagnostic", result.stderr)
                    error = None
                else:
                    with self.assertRaises(subprocess.TimeoutExpired) as caught:
                        RUNNER.run_command(cmd, root, timeout_s)
                    error = caught.exception
                for name in (kind, "child"):
                    record_file = root / (name + ".json")
                    self.assertTrue(record_file.exists(), f"{name} never became ready")
                    records.append(json.loads(record_file.read_text(encoding="utf-8")))
                if error is not None:
                    self.assertEqual(cmd, error.cmd)
                    self.assertEqual(timeout_s, error.timeout)
                    output_type = str if os.name == "nt" else bytes
                    self.assertIsInstance(error.output, output_type)
                    self.assertIsInstance(error.stderr, output_type)
                    expected_stdout = "launcher ready" if os.name == "nt" else b"launcher ready"
                    expected_stderr = "launcher diagnostic" if os.name == "nt" else b"launcher diagnostic"
                    self.assertIn(expected_stdout, error.output)
                    self.assertIn(expected_stderr, error.stderr)
                deadline = time.monotonic() + 2
                while any(listener_open(record["port"]) for record in records) and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertFalse(listener_open(records[0]["port"]), "launcher survived its timeout")
                self.assertFalse(listener_open(records[1]["port"]), "native-child analogue survived launcher timeout")
            finally:
                # The RED run leaves a child alive. Clean only the exact PIDs this
                # fixture just created; never use process-name matching.
                for record_file in (root / (kind + ".json"), root / "child.json"):
                    if record_file.exists():
                        record = json.loads(record_file.read_text(encoding="utf-8"))
                        if listener_open(record["port"]):
                            try:
                                os.kill(record["pid"], signal.SIGKILL if os.name == "posix" else signal.SIGTERM)
                            except ProcessLookupError:
                                pass


if __name__ == "__main__":
    unittest.main()
