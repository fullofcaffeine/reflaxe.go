#!/usr/bin/env python3

from __future__ import annotations

import re
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SEMANTIC_ROOTS = [ROOT / "test/semantic_diff", ROOT / "test/semantic_diff_lanes"]


class SemanticDiffHarnessContractTest(unittest.TestCase):
    def test_fixtures_keep_one_portable_program_as_the_reference(self) -> None:
        forbidden_patterns = {
            "a go.* import": re.compile(r"(?m)^\s*import\s+go\."),
            "a typed Go extern": re.compile(r"@:go\.import\b"),
            "a direct go.* API reference": re.compile(
                r"(?<![A-Za-z0-9_])go\.(?:Chan|Go|Map|NativeSlice|Result|Select|Slice)\b"
            ),
            "a go_output-only fallback": re.compile(r"\bgo_output\b"),
        }

        violations: list[str] = []
        for semantic_root in SEMANTIC_ROOTS:
            for source in sorted(semantic_root.rglob("*.hx")):
                text = source.read_text(encoding="utf-8")
                for label, pattern in forbidden_patterns.items():
                    if pattern.search(text):
                        violations.append(f"{source.relative_to(ROOT)} contains {label}")

        self.assertEqual(
            [],
            violations,
            "Semantic diff must execute one genuinely portable program under both "
            "Haxe --interp and generated Go. Move target-native behavior to a "
            "runtime snapshot; see docs/semantic-diff-guide.md.\n"
            + "\n".join(violations),
        )


if __name__ == "__main__":
    unittest.main()
