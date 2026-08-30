#!/usr/bin/env python3

from __future__ import annotations

from pathlib import Path
import unittest


ROOT = Path(__file__).resolve().parent.parent


class GeneratedLintContractTests(unittest.TestCase):
    def test_runtime_documents_intentional_security_findings_at_the_call(self) -> None:
        file_runtime = (ROOT / "runtime/hxrt/file.go").read_text(encoding="utf-8")
        rationale = (
            "//nolint:gosec // Haxe file APIs create ordinary user-visible files "
            "and honor the process umask."
        )
        self.assertEqual(4, file_runtime.count(rationale))

    def test_int32_lowering_avoids_redundant_native_conversions(self) -> None:
        bytes_output = (
            ROOT
            / "test/snapshot/stdlib/io_type_smoke/intended/module_haxe_io_bytesbuffer.go"
        ).read_text(encoding="utf-8")
        self.assertNotIn("int(int32((hxrt.Int32Wrap", bytes_output)
        self.assertNotIn("int(int32(int32(", bytes_output)

        unary_output = (
            ROOT
            / "test/snapshot/core/arithmetic/intended/main.go"
        ).read_text(encoding="utf-8")
        self.assertNotIn("int(int32(-int32(", unary_output)
        self.assertNotIn("int(int32(^int32(", unary_output)
        self.assertIn("int(-int32(", unary_output)
        self.assertIn("int(^int32(", unary_output)


if __name__ == "__main__":
    unittest.main()
