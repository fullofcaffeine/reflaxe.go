#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class DocsClarityContractTest(unittest.TestCase):
    def section_text(self, text: str, heading: str) -> str:
        start = text.index(heading)
        next_heading = text.find("\n## ", start + len(heading))
        if next_heading == -1:
            return text[start:]
        return text[start:next_heading]

    def test_docs_index_and_glossary_exist(self) -> None:
        self.assertTrue((REPO_ROOT / "docs" / "index.md").exists(), "docs/index.md must exist")
        self.assertTrue((REPO_ROOT / "docs" / "glossary.md").exists(), "docs/glossary.md must exist")

    def test_readme_links_docs_index_and_glossary(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        self.assertIn("docs/index.md", readme)
        self.assertIn("docs/glossary.md", readme)

    def test_readme_quick_start_separates_first_run_from_validation(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        quick_start = self.section_text(readme, "## Quick start")
        validation = self.section_text(readme, "## Validation commands")

        self.assertIn("npm install", quick_start)
        self.assertIn("npm run hooks:install", quick_start)
        self.assertIn("npm run dev:hx -- --project examples/tui_todo --profile portable --action run", quick_start)
        self.assertNotIn("npm test", quick_start)
        self.assertNotIn("python3 test/run-ci.py", quick_start)

        self.assertIn("npm test", validation)
        self.assertIn("python3 test/run-ci.py", validation)
        self.assertIn("python3 test/run-examples.py", validation)

    def test_readme_no_longer_claims_portable_only_examples_are_dual_profile(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        self.assertNotIn("examples/tui_todo --profile metal", readme)
        self.assertNotIn("complex single-codebase app compiled across both profiles", readme)
        self.assertNotIn("compact profile adapter/storyboard reference", readme)

    def test_readme_output_model_names_first_run_and_override_paths(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        output_model = self.section_text(readme, "## Output model")

        self.assertIn("examples/tui_todo/out_portable", output_model)
        self.assertIn("out_<profile>", output_model)
        self.assertIn("go_output", output_model)
        self.assertIn("--out", output_model)
        self.assertIn("safe to delete", output_model)

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
        self.assertIn("glossary.md", profiles)
        self.assertIn("glossary.md", start_here)

    def test_profiles_document_portable_as_main_product_not_slow_mode(self) -> None:
        profiles = (REPO_ROOT / "docs" / "profiles.md").read_text(encoding="utf-8")
        guide = (REPO_ROOT / "docs" / "profile-semantics-guide.md").read_text(encoding="utf-8")
        agents = (REPO_ROOT / "AGENTS.md").read_text(encoding="utf-8")

        for text in [profiles, guide, agents]:
            self.assertIn("portable is the default product path", text)
            self.assertIn("metal is an explicit Go-native authoring contract", text)
            self.assertIn("Portable by default, Go-native by opt-in", text)

    def test_remaining_profile_docs_define_go_native_without_metal_as_fast_only_mode(self) -> None:
        targets = [
            "README.md",
            "docs/start-here.md",
            "docs/glossary.md",
            "docs/performance-budget-policy.md",
            "docs/go-concurrency-interop-guide.md",
            "docs/examples-matrix.md",
            "examples/pulseforge/README.md",
            "examples/fluxproxy/README.md",
            "examples/worker_pool_select/README.md",
        ]

        for rel in targets:
            text = (REPO_ROOT / rel).read_text(encoding="utf-8")
            plain_text = text.replace("`", "")
            self.assertIn("metal is not required for good Go output", plain_text, rel)
            self.assertIn("Go-native", text, rel)

        glossary = (REPO_ROOT / "docs" / "glossary.md").read_text(encoding="utf-8")
        self.assertIn("## Go-native", glossary)
        self.assertIn("## Go-first", glossary)
        self.assertIn("## go_native", glossary)
        self.assertIn("## Hot path", glossary)

    def test_start_here_explains_why_std_folder_is_small(self) -> None:
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        self.assertIn("## Where is the stdlib?", start_here)
        self.assertIn("upstream Haxe stdlib", start_here)
        self.assertIn("runtime/hxrt", start_here)
        self.assertIn("src/reflaxe/go/GoCompiler.hx", start_here)
        self.assertIn("docs/stdlib-shim-rationale.md", start_here)

    def test_start_here_separates_first_run_from_validation_gates(self) -> None:
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        first_run = self.section_text(start_here, "## First successful run")
        validation = self.section_text(start_here, "## Validation after first run")

        self.assertIn("npm install", first_run)
        self.assertIn("npm run hooks:install", first_run)
        self.assertIn("npm run dev:hx -- --project examples/tui_todo --profile portable --action run", first_run)
        self.assertNotIn("python3 test/run-ci.py", first_run)
        self.assertNotIn("python3 test/run-snapshots.py", first_run)

        self.assertIn("python3 test/run-snapshots.py", validation)
        self.assertIn("python3 test/run-examples.py", validation)
        self.assertIn("python3 test/run-ci.py", validation)

    def test_start_here_explains_first_run_output_artifacts(self) -> None:
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        artifacts = self.section_text(start_here, "## Generated output and local artifacts")

        self.assertIn("examples/tui_todo/out_portable", artifacts)
        self.assertIn("-D go_output=out_portable", artifacts)
        self.assertIn("--out", artifacts)
        self.assertIn("safe to delete", artifacts)
        self.assertIn("regenerated", artifacts)

    def test_docs_internal_links_do_not_use_root_docs_prefix(self) -> None:
        targets = [
            "docs/index.md",
            "docs/start-here.md",
            "docs/profiles.md",
            "docs/profile-semantics-guide.md",
            "docs/semantic-diff-guide.md",
            "docs/examples-matrix.md",
            "docs/hxrt-runtime.md",
        ]
        for rel in targets:
            text = (REPO_ROOT / rel).read_text(encoding="utf-8")
            self.assertNotIn("](/docs/", text, f"{rel} should use relative docs links")

    def test_goextern_explains_advanced_signature_boundaries(self) -> None:
        goextern = (REPO_ROOT / "docs" / "goextern.md").read_text(encoding="utf-8")
        self.assertIn("## Advanced Signature Boundary Policy", goextern)
        self.assertIn("ordinary typed extern metadata", goextern)
        self.assertIn("single `(T,error)`", goextern)
        self.assertIn("typed facade wrapper", goextern)
        self.assertIn("tuple carrier", goextern)
        self.assertIn("@:go.tupleReturn", goextern)
        self.assertIn("must not use raw `__go__`", goextern)
        self.assertIn("tools/goextern/main_test.go", goextern)


if __name__ == "__main__":
    unittest.main()
