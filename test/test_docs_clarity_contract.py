#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class DocsClarityContractTest(unittest.TestCase):
    def test_docs_index_and_glossary_exist(self) -> None:
        self.assertTrue((REPO_ROOT / "docs" / "index.md").exists(), "docs/index.md must exist")
        self.assertTrue((REPO_ROOT / "docs" / "glossary.md").exists(), "docs/glossary.md must exist")

    def test_readme_links_docs_index_and_glossary(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        self.assertIn("docs/index.md", readme)
        self.assertIn("docs/glossary.md", readme)

    def test_readme_no_longer_claims_portable_only_examples_are_dual_profile(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        self.assertNotIn("examples/tui_todo --profile metal", readme)
        self.assertNotIn("complex single-codebase app compiled across both profiles", readme)
        self.assertNotIn("compact profile adapter/storyboard reference", readme)

    def test_core_docs_have_terms_and_related_sections(self) -> None:
        targets = [
            "docs/start-here.md",
            "docs/profiles.md",
            "docs/profile-semantics-guide.md",
            "docs/semantic-diff-guide.md",
            "docs/examples-matrix.md",
            "docs/hxrt-runtime.md",
        ]
        for rel in targets:
            text = (REPO_ROOT / rel).read_text(encoding="utf-8")
            self.assertIn("## Terms", text, f"{rel} should define terms for newcomers")
            self.assertIn("## Related docs", text, f"{rel} should include related-doc links")

    def test_profile_and_start_here_reference_glossary(self) -> None:
        profiles = (REPO_ROOT / "docs" / "profiles.md").read_text(encoding="utf-8")
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        self.assertIn("docs/glossary.md", profiles)
        self.assertIn("docs/glossary.md", start_here)


if __name__ == "__main__":
    unittest.main()
