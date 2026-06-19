#!/usr/bin/env python3

from __future__ import annotations

import re
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RAW_INJECTION_RE = re.compile(r"\b__go__\s*\(")
GO_INJECTION_RE = re.compile(r"\bGoInjection\s*\.")
GO_ALLOW_RAW_RE = re.compile(r"@:goAllowRaw\b")

ALLOWED_TEST_FIXTURE_PREFIXES = (
    "test/snapshot/negative/",
    "test/snapshot/core/strict_mode_injection_allow_raw/",
)
IGNORED_TEST_PREFIXES = (
    "test/.test-cache/",
)


class RawInjectionHygieneContractTest(unittest.TestCase):
    def relative(self, path: Path) -> str:
        return path.relative_to(REPO_ROOT).as_posix()

    def matching_lines(self, path: Path, patterns: tuple[re.Pattern[str], ...]) -> list[str]:
        text = path.read_text(encoding="utf-8")
        hits: list[str] = []
        for lineno, line in enumerate(text.splitlines(), start=1):
            if any(pattern.search(line) for pattern in patterns):
                hits.append(f"{self.relative(path)}:{lineno}: {line.strip()}")
        return hits

    def test_examples_do_not_use_raw_go_injection(self) -> None:
        """Examples are product docs, so they must teach wrappers or typed externs."""
        violations: list[str] = []
        for path in sorted((REPO_ROOT / "examples").rglob("*.hx")):
            violations.extend(
                self.matching_lines(
                    path,
                    (RAW_INJECTION_RE, GO_INJECTION_RE, GO_ALLOW_RAW_RE),
                )
            )

        self.assertFalse(
            violations,
            "examples must not use raw __go__, GoInjection, or @:goAllowRaw; "
            "use staged std, hxrt helpers, go.* facades, or typed externs instead:\n"
            + "\n".join(violations),
        )

    def test_test_fixture_raw_injection_is_explicitly_scoped(self) -> None:
        """Intentional raw-injection fixtures are allowed only in boundary-test folders."""
        violations: list[str] = []
        for path in sorted((REPO_ROOT / "test").rglob("*.hx")):
            rel = self.relative(path)
            if rel.startswith(IGNORED_TEST_PREFIXES):
                continue
            if rel.startswith(ALLOWED_TEST_FIXTURE_PREFIXES):
                continue
            violations.extend(
                self.matching_lines(
                    path,
                    (RAW_INJECTION_RE, GO_INJECTION_RE, GO_ALLOW_RAW_RE),
                )
            )

        self.assertFalse(
            violations,
            "ordinary test fixtures must not normalize raw Go injection; "
            "put intentional boundary probes under test/snapshot/negative/ or a named allow-raw contract fixture:\n"
            + "\n".join(violations),
        )

    def test_release_contracts_run_this_hygiene_check(self) -> None:
        runner = (REPO_ROOT / "test" / "run-release-contracts.py").read_text(encoding="utf-8")
        self.assertIn("test/test_raw_injection_hygiene_contract.py", runner)

    def test_agents_tell_future_agents_to_run_the_contract(self) -> None:
        agents = (REPO_ROOT / "AGENTS.md").read_text(encoding="utf-8")
        self.assertIn("test/test_raw_injection_hygiene_contract.py", agents)
        self.assertIn("Intentional boundary fixtures", agents)


if __name__ == "__main__":
    unittest.main()
