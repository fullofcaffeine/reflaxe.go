#!/usr/bin/env python3

from __future__ import annotations

import re
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
DYNAMIC_CATCH_RE = re.compile(r"catch\s*\([^)]*:\s*Dynamic\s*\)")

STD_DYNAMIC_CATCH_ALLOWLIST = {
    Path("std/go/_std/haxe/Template.hx"),
    Path("std/sys/thread/ElasticThreadPoolWorker.cross.hx"),
}


class TypedExceptionHygieneContract(unittest.TestCase):
    def test_examples_do_not_use_dynamic_catches(self) -> None:
        hits: list[str] = []
        for path in sorted((REPO_ROOT / "examples").rglob("*.hx")):
            text = path.read_text(encoding="utf-8")
            for line_no, line in enumerate(text.splitlines(), start=1):
                if DYNAMIC_CATCH_RE.search(line):
                    hits.append(f"{path.relative_to(REPO_ROOT)}:{line_no}: {line.strip()}")
        self.assertEqual([], hits, "Examples should use typed catches or haxe.Exception, not catch Dynamic.")

    def test_staged_std_dynamic_catches_are_documented_and_allowlisted(self) -> None:
        hits: list[str] = []
        for path in sorted((REPO_ROOT / "std").rglob("*.hx")):
            rel = path.relative_to(REPO_ROOT)
            text = path.read_text(encoding="utf-8")
            lines = text.splitlines()
            for line_no, line in enumerate(lines, start=1):
                if not DYNAMIC_CATCH_RE.search(line):
                    continue
                if rel not in STD_DYNAMIC_CATCH_ALLOWLIST:
                    hits.append(f"{rel}:{line_no}: not allowlisted: {line.strip()}")
                    continue
                context = "\n".join(lines[max(0, line_no - 4):line_no + 1]).lower()
                if "dynamic catch is intentional" not in context:
                    hits.append(f"{rel}:{line_no}: missing 'Dynamic catch is intentional' comment")
        self.assertEqual([], hits, "Staged std Dynamic catches must be rare, allowlisted, and documented.")


if __name__ == "__main__":
    unittest.main()
