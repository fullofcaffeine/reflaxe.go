#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SPEC = REPO_ROOT / "docs" / "existing-go-module-mode.md"


class ExistingModuleModeSpecContractTest(unittest.TestCase):
    def setUp(self) -> None:
        self.text = SPEC.read_text(encoding="utf-8")
        self.plain_text = " ".join(self.text.split())

    def test_spec_defines_the_complete_mode_boundary(self) -> None:
        index = (REPO_ROOT / "docs" / "index.md").read_text(encoding="utf-8")
        self.assertIn("existing-go-module-mode.md", index)

        for heading in [
            "## Before and after",
            "## Typed project manifest",
            "## Validation before writes",
            "## Generated-file ownership",
            "## Package and entry point",
            "## Build behavior",
            "## Diagnostics and reports",
            "## Migration",
            "## Implementation order",
        ]:
            self.assertIn(heading, self.text)

        for task in ["M03-02", "M03-03", "M03-04", "M03-07", "M03-08"]:
            self.assertIn(task, self.text)

    def test_spec_preserves_caller_files_and_standalone_behavior(self) -> None:
        for contract in [
            "never creates, changes, or removes `go.mod` or `go.sum`",
            "Standalone mode remains the default",
            "reads the module path from `go.mod`",
            "A changed generated file is a conflict",
            "machine-local absolute paths",
        ]:
            self.assertIn(contract, self.plain_text)

    def test_spec_names_typed_inputs_and_entry_point_policies(self) -> None:
        for contract in [
            "`moduleRoot`",
            "`packageDir`",
            "`packageName`",
            "`runtimeDir`",
            "`compiler-main`",
            "`caller-bridge`",
            "`build`",
            "`go_output`",
            "`go_module`",
        ]:
            self.assertIn(contract, self.plain_text)


if __name__ == "__main__":
    unittest.main()
