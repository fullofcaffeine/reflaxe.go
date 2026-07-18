#!/usr/bin/env python3

from __future__ import annotations

import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent


@unittest.skipUnless(shutil.which("go"), "requires Go")
class SocketRuntimeCrossBuildTest(unittest.TestCase):
    def test_socket_runtime_builds_for_posix_and_windows(self) -> None:
        targets = (
            ("linux", "amd64", "hxrt-linux.test"),
            ("windows", "amd64", "hxrt-windows.test.exe"),
        )
        with tempfile.TemporaryDirectory(prefix="haxe-go-socket-cross-build-") as raw:
            output_dir = Path(raw)
            for goos, goarch, name in targets:
                with self.subTest(goos=goos, goarch=goarch):
                    environment = os.environ.copy()
                    environment.update(
                        {
                            "CGO_ENABLED": "0",
                            "GO111MODULE": "off",
                            "GOOS": goos,
                            "GOARCH": goarch,
                        }
                    )
                    process = subprocess.run(
                        [
                            "go",
                            "test",
                            "-c",
                            "-o",
                            str(output_dir / name),
                            "./runtime/hxrt",
                        ],
                        cwd=ROOT,
                        env=environment,
                        capture_output=True,
                        text=True,
                        timeout=120,
                    )
                    self.assertEqual(
                        0,
                        process.returncode,
                        process.stdout + process.stderr,
                    )


if __name__ == "__main__":
    unittest.main()
