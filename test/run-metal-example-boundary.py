#!/usr/bin/env python3

from __future__ import annotations

import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent

TARGET_GLOBS = [
    "examples/*/profile/MetalRuntime.hx",
    "examples/*/app/runtime/GoNativeRuntime.hx",
]

FORBIDDEN_IMPORTS = [
    re.compile(r"^\s*import\s+haxe\.ds\.List\s*;"),
    re.compile(r"^\s*import\s+haxe\.ds\.IntMap\s*;"),
]


def iter_target_files() -> list[Path]:
    files: list[Path] = []
    for pattern in TARGET_GLOBS:
        files.extend(sorted(ROOT.glob(pattern)))
    return files


def main() -> int:
    violations: list[str] = []
    targets = iter_target_files()

    for path in targets:
        text = path.read_text(encoding="utf-8")
        for lineno, line in enumerate(text.splitlines(), start=1):
            for pattern in FORBIDDEN_IMPORTS:
                if pattern.search(line):
                    rel = path.relative_to(ROOT).as_posix()
                    violations.append(f"{rel}:{lineno}: forbidden portable collection import in metal-only example module: {line.strip()}")

    if violations:
        print("Metal example boundary check failed:")
        for entry in violations:
            print(f"- {entry}")
        return 1

    print(f"Metal example boundary check passed ({len(targets)} files scanned).")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
