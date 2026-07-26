#!/usr/bin/env python3

from __future__ import annotations

import json
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
PACKAGE_PATH = ROOT / "package.json"
LOCK_PATH = ROOT / "package-lock.json"
RELEASE_CONFIG_PATH = ROOT / ".releaserc.json"
RELEASE_CONTRACT_RUNNER_PATH = ROOT / "test" / "run-release-contracts.py"
SENTINEL_DIR = ROOT / "scripts" / "release" / "disabled-semantic-release-npm"
SENTINEL_SPEC = "file:scripts/release/disabled-semantic-release-npm"


def load_json(path: Path) -> dict:
    return json.loads(path.read_text(encoding="utf-8"))


class SemanticReleaseDependencyBoundaryContract(unittest.TestCase):
    def test_unused_npm_publish_plugin_is_a_local_fail_closed_sentinel(self) -> None:
        package = load_json(PACKAGE_PATH)
        self.assertEqual(
            SENTINEL_SPEC,
            package["devDependencies"].get("@semantic-release/npm"),
        )

        sentinel_package = load_json(SENTINEL_DIR / "package.json")
        self.assertEqual("@semantic-release/npm", sentinel_package.get("name"))
        self.assertEqual("13.1.5+disabled", sentinel_package.get("version"))
        self.assertTrue(sentinel_package.get("private"))
        self.assertEqual("GPL-3.0-only", sentinel_package.get("license"))
        self.assertEqual("./index.js", sentinel_package.get("exports"))

        config = load_json(RELEASE_CONFIG_PATH)
        plugin_names = [
            entry[0] if isinstance(entry, list) else entry
            for entry in config["plugins"]
        ]
        self.assertNotIn("@semantic-release/npm", plugin_names)

    def test_sentinel_throws_if_a_publish_hook_is_activated(self) -> None:
        script = """
        import * as plugin from './scripts/release/disabled-semantic-release-npm/index.js';
        const expected = ['addChannel', 'prepare', 'publish', 'verifyConditions'];
        if (JSON.stringify(Object.keys(plugin).sort()) !== JSON.stringify(expected)) {
          throw new Error(`unexpected exports: ${Object.keys(plugin).sort().join(',')}`);
        }
        for (const hook of expected) {
          try {
            await plugin[hook]();
            throw new Error(`${hook} unexpectedly succeeded`);
          } catch (error) {
            if (!String(error.message).includes('disabled by haxe.go release policy')) {
              throw error;
            }
          }
        }
        """
        result = subprocess.run(
            ["node", "--input-type=module", "--eval", script],
            cwd=ROOT,
            text=True,
            capture_output=True,
        )
        self.assertEqual(0, result.returncode, result.stderr)

    def test_lockfile_resolves_the_sentinel_without_bundled_npm(self) -> None:
        lock = load_json(LOCK_PATH)
        self.assertEqual(
            SENTINEL_SPEC,
            lock["packages"][""]["devDependencies"].get("@semantic-release/npm"),
        )
        sentinel = lock["packages"].get("node_modules/@semantic-release/npm")
        self.assertEqual(
            {
                "resolved": "scripts/release/disabled-semantic-release-npm",
                "link": True,
            },
            sentinel,
        )
        self.assertEqual(
            "13.1.5+disabled",
            lock["packages"][
                "scripts/release/disabled-semantic-release-npm"
            ].get("version"),
        )
        installed_npm_paths = [
            path
            for path in lock["packages"]
            if path == "node_modules/npm" or path.endswith("/node_modules/npm")
        ]
        self.assertEqual([], installed_npm_paths)

    def test_boundary_contract_is_wired_into_the_release_gate(self) -> None:
        release_contract_runner = RELEASE_CONTRACT_RUNNER_PATH.read_text(
            encoding="utf-8"
        )
        self.assertIn(
            '["python3", "test/test_semantic_release_dependency_boundary.py"]',
            release_contract_runner,
        )


if __name__ == "__main__":
    unittest.main()
