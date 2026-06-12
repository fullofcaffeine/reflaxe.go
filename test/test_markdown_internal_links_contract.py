#!/usr/bin/env python3

from __future__ import annotations

import re
import unittest
from pathlib import Path
from urllib.parse import unquote, urlparse


REPO_ROOT = Path(__file__).resolve().parent.parent
MARKDOWN_LINK_RE = re.compile(r"(?<!!)\[[^\]]+\]\(([^)\s]+)(?:\s+\"[^\"]*\")?\)")
EXTERNAL_SCHEMES = {"http", "https", "mailto", "app"}


class MarkdownInternalLinksContractTest(unittest.TestCase):
    def markdown_files(self) -> list[Path]:
        files = [REPO_ROOT / "README.md", REPO_ROOT / "AGENTS.md"]
        files.extend(sorted((REPO_ROOT / "docs").rglob("*.md")))
        files.extend(sorted((REPO_ROOT / "examples").rglob("*.md")))
        return [path for path in files if path.exists()]

    def test_markdown_links_do_not_use_root_docs_paths(self) -> None:
        failures: list[str] = []
        for path in self.markdown_files():
            rel = path.relative_to(REPO_ROOT)
            for line_no, line in enumerate(path.read_text(encoding="utf-8").splitlines(), start=1):
                for match in MARKDOWN_LINK_RE.finditer(line):
                    target = match.group(1)
                    if target.startswith("/docs/"):
                        failures.append(
                            f"{rel}:{line_no}: use a relative docs link instead of root-absolute {target!r}"
                        )
        self.assertEqual([], failures)

    def test_local_markdown_links_resolve(self) -> None:
        failures: list[str] = []
        for path in self.markdown_files():
            rel = path.relative_to(REPO_ROOT)
            text = path.read_text(encoding="utf-8")
            for line_no, line in enumerate(text.splitlines(), start=1):
                for match in MARKDOWN_LINK_RE.finditer(line):
                    raw_target = unquote(match.group(1))
                    parsed = urlparse(raw_target)
                    if parsed.scheme in EXTERNAL_SCHEMES or raw_target.startswith("#"):
                        continue
                    if raw_target.startswith("/"):
                        failures.append(f"{rel}:{line_no}: root-absolute local link is not portable: {raw_target}")
                        continue
                    link_path = parsed.path
                    if not link_path or not link_path.endswith(".md"):
                        continue
                    target_path = (path.parent / link_path).resolve()
                    try:
                        target_path.relative_to(REPO_ROOT)
                    except ValueError:
                        failures.append(f"{rel}:{line_no}: link escapes repository: {raw_target}")
                        continue
                    if not target_path.exists():
                        failures.append(f"{rel}:{line_no}: missing Markdown target: {raw_target}")
        self.assertEqual([], failures)


if __name__ == "__main__":
    unittest.main()
