#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
import unittest


ROOT = Path(__file__).resolve().parent.parent
BEHAVIOR_COMMAND = "python3 test/test_sys_get_char_terminal.py"


class SysGetCharTerminalContractTest(unittest.TestCase):
    def test_behavior_gate_is_wired_into_local_changed_and_ci_surfaces(self) -> None:
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package["scripts"]
        self.assertEqual(BEHAVIOR_COMMAND, scripts.get("test:sys-get-char-terminal"))
        self.assertIn("npm run test:sys-get-char-terminal", scripts["test"])
        self.assertIn("npm run test:sys-get-char-terminal", scripts["test:changed"])

        ci_runner = (ROOT / "test" / "run-ci.py").read_text(encoding="utf-8")
        self.assertIn("build_sys_get_char_terminal_command", ci_runner)
        self.assertIn("Sys.getChar terminal contract stage", ci_runner)

        release_runner = (ROOT / "test" / "run-release-contracts.py").read_text(
            encoding="utf-8"
        )
        self.assertIn("test/test_sys_get_char_terminal_contract.py", release_runner)

    def test_staged_source_owns_eof_and_echo_while_hxrt_owns_terminal_state(self) -> None:
        staged_sys = (ROOT / "std" / "go" / "_std" / "Sys.hx").read_text(
            encoding="utf-8"
        )
        native_terminal = (
            ROOT / "std" / "hxrt" / "sys" / "NativeTerminal.hx"
        ).read_text(
            encoding="utf-8"
        )
        file_runtime = (ROOT / "runtime" / "hxrt" / "file.go").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (
            ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")

        self.assertIn("var value = NativeTerminal.readChar();", staged_sys)
        self.assertIn("throw new Eof();", staged_sys)
        self.assertIn("stdout().writeByte(value);", staged_sys)
        self.assertIn('@:go.name("SysReadCharValue")', native_terminal)
        self.assertNotIn("SysGetChar", file_runtime)

        terminal_files = {
            "terminal.go",
            "terminal_darwin.go",
            "terminal_linux.go",
            "terminal_posix.go",
            "terminal_unsupported.go",
            "terminal_windows.go",
        }
        for file_name in terminal_files:
            self.assertTrue((ROOT / "runtime" / "hxrt" / file_name).is_file(), file_name)
            self.assertIn(f'"{file_name}"', feature_analyzer)
        self.assertIn('var HxrtTerminal = "terminal";', feature_analyzer)
        self.assertIn('path == "hxrt.sys.NativeTerminal"', feature_analyzer)

        runtime_sources = {
            path.name: path.read_text(encoding="utf-8")
            for path in (ROOT / "runtime" / "hxrt").glob("*.go")
        }
        self.assertFalse(
            any("golang.org/x/term" in source for source in runtime_sources.values())
        )
        unsafe_pointer_files = {
            name for name, source in runtime_sources.items() if "unsafe.Pointer" in source
        }
        self.assertEqual({"terminal_posix.go"}, unsafe_pointer_files)


if __name__ == "__main__":
    unittest.main()
