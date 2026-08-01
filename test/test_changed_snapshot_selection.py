#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest
from unittest import mock


ROOT = Path(__file__).resolve().parent.parent
RUNNER = ROOT / "test" / "run-snapshots.py"


def load_runner():
    test_dir = str(RUNNER.parent)
    if test_dir not in sys.path:
        sys.path.insert(0, test_dir)
    spec = importlib.util.spec_from_file_location("run_snapshots_changed_contract", RUNNER)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load run-snapshots.py")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


def git(root: Path, *args: str) -> None:
    subprocess.run(["git", *args], cwd=root, check=True, capture_output=True, text=True)


def write(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text, encoding="utf-8")


class ChangedSnapshotSelectionTest(unittest.TestCase):
    def test_worktree_index_rename_delete_and_untracked_cases_are_unioned(self) -> None:
        module = load_runner()
        with tempfile.TemporaryDirectory(prefix="haxe-go-changed-snapshots-") as raw:
            repo = Path(raw)
            git(repo, "init", "-q")
            git(repo, "config", "user.name", "Snapshot Contract")
            git(repo, "config", "user.email", "snapshot@example.invalid")

            for case in ("unstaged", "renamed", "deleted"):
                write(repo / "test" / "snapshot" / "core" / case / "compile.hxml", "-main Main\n")
            write(repo / "outside.txt", "ignored\n")
            git(repo, "add", ".")
            git(repo, "commit", "-qm", "baseline")

            write(repo / "test" / "snapshot" / "core" / "unstaged" / "Main.hx", "class Main {}\n")
            git(
                repo,
                "mv",
                "test/snapshot/core/renamed",
                "test/snapshot/core/renamed_now",
            )
            (repo / "test" / "snapshot" / "core" / "deleted" / "compile.hxml").unlink()
            write(repo / "test" / "snapshot" / "stdlib" / "staged" / "compile.hxml", "-main Main\n")
            git(repo, "add", "test/snapshot/stdlib/staged")
            write(repo / "test" / "snapshot" / "sys" / "untracked" / "compile.hxml", "-main Main\n")
            write(repo / "untracked-outside.txt", "ignored\n")

            original_root = module.ROOT
            try:
                module.ROOT = repo
                selected = module.changed_case_ids()
            finally:
                module.ROOT = original_root

        self.assertEqual(
            {
                "core/unstaged",
                "core/renamed",
                "core/renamed_now",
                "core/deleted",
                "stdlib/staged",
                "sys/untracked",
            },
            selected,
        )

    def test_git_discovery_failure_expands_to_every_snapshot(self) -> None:
        module = load_runner()
        expected = {case.case_id for case in module.discover_cases()}
        with mock.patch.object(
            module,
            "collect_changed_paths",
            side_effect=module.GitChangeDiscoveryError("deliberate discovery failure"),
        ):
            selected = module.changed_case_ids()
        self.assertEqual(expected, selected)


if __name__ == "__main__":
    unittest.main()
