#!/usr/bin/env python3

"""Fail closed when dependency, action, or vendor provenance drifts."""

from __future__ import annotations

import json
import re
import subprocess
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
PACKAGE_JSON = ROOT / "package.json"
PACKAGE_LOCK = ROOT / "package-lock.json"
ACTIONS_LOCK = ROOT / ".github" / "actions-lock.json"
DEPENDABOT = ROOT / ".github" / "dependabot.yml"
WORKFLOW_ROOT = ROOT / ".github"
VENDOR_VERIFIER = ROOT / "scripts" / "vendor" / "verify-reflaxe-provenance.py"
USES_RE = re.compile(r"^\s*uses:\s*([^\s#]+)(?:\s+#\s*(\S+))?\s*$")
COMMIT_RE = re.compile(r"^[0-9a-f]{40}$")
VERSION_RE = re.compile(r"^v\d+\.\d+\.\d+(?:[-+][0-9A-Za-z.-]+)?$")
PACKAGE_MANAGER_RE = re.compile(r"^npm@\d+\.\d+\.\d+$")


class VerificationError(RuntimeError):
    """A committed supply-chain contract was missing or inconsistent."""


def load_json(path: Path) -> dict[str, object]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as error:
        raise VerificationError(f"cannot load {path.relative_to(ROOT)}: {error}") from error
    if not isinstance(value, dict):
        raise VerificationError(f"{path.relative_to(ROOT)} must contain an object")
    return value


def verify_package_lock() -> None:
    package = load_json(PACKAGE_JSON)
    lock = load_json(PACKAGE_LOCK)
    packages = lock.get("packages")
    if not isinstance(packages, dict) or not isinstance(packages.get(""), dict):
        raise VerificationError("package-lock.json has no root package metadata")
    root = packages[""]

    if lock.get("lockfileVersion") != 3:
        raise VerificationError("package-lock.json must use lockfileVersion 3")
    for key in ("name", "version"):
        if lock.get(key) != package.get(key) or root.get(key) != package.get(key):
            raise VerificationError(f"package-lock.json root {key} does not match package.json")
    for key in ("dependencies", "devDependencies"):
        if root.get(key, {}) != package.get(key, {}):
            raise VerificationError(f"package-lock.json root {key} is stale")
    package_manager = package.get("packageManager")
    if not isinstance(package_manager, str) or not PACKAGE_MANAGER_RE.fullmatch(
        package_manager
    ):
        raise VerificationError("package.json packageManager must pin an exact npm version")
    print(
        "[supply-chain] package-lock.json: OK "
        f"(lockfileVersion 3; {package_manager})"
    )


def action_manifest() -> dict[str, tuple[str, str]]:
    manifest = load_json(ACTIONS_LOCK)
    if manifest.get("schema_version") != 1:
        raise VerificationError(".github/actions-lock.json has unsupported schema")
    entries = manifest.get("actions")
    if not isinstance(entries, list):
        raise VerificationError(".github/actions-lock.json actions must be a list")

    locked: dict[str, tuple[str, str]] = {}
    order: list[str] = []
    for entry in entries:
        if not isinstance(entry, dict):
            raise VerificationError("action lock entry must be an object")
        repository = entry.get("repository")
        version = entry.get("version")
        commit = entry.get("commit")
        if not all(isinstance(value, str) for value in (repository, version, commit)):
            raise VerificationError("action lock entry requires repository, version, and commit")
        assert isinstance(repository, str)
        assert isinstance(version, str)
        assert isinstance(commit, str)
        if repository in locked:
            raise VerificationError(f"duplicate locked action: {repository}")
        if not VERSION_RE.fullmatch(version):
            raise VerificationError(f"{repository}: version must be an exact release tag")
        if not COMMIT_RE.fullmatch(commit):
            raise VerificationError(f"{repository}: commit must be a 40-character SHA")
        locked[repository] = (version, commit)
        order.append(repository)
    if order != sorted(order):
        raise VerificationError(".github/actions-lock.json actions must be sorted")
    return locked


def workflow_files() -> list[Path]:
    return sorted(
        {
            *WORKFLOW_ROOT.rglob("*.yml"),
            *WORKFLOW_ROOT.rglob("*.yaml"),
        }
    )


def verify_actions_and_installs() -> None:
    locked = action_manifest()
    seen: set[str] = set()
    npm_ci_count = 0
    npm_install_re = re.compile(r"\bnpm\s+(?:install|i)\b")

    for path in workflow_files():
        for line_number, line in enumerate(
            path.read_text(encoding="utf-8").splitlines(), 1
        ):
            if npm_install_re.search(line):
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: use npm ci, not npm install"
                )
            if re.search(r"\bnpm\s+ci\b", line):
                npm_ci_count += 1

            match = USES_RE.match(line)
            if match is None:
                continue
            action, version_comment = match.groups()
            if action.startswith("./"):
                continue
            if "@" not in action:
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: action has no immutable ref"
                )
            action_path, commit = action.rsplit("@", 1)
            components = action_path.split("/")
            if len(components) < 2:
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: unsupported action reference"
                )
            repository = "/".join(components[:2])
            if repository not in locked:
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: {repository} is not manifested"
                )
            expected_version, expected_commit = locked[repository]
            if commit != expected_commit:
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: {repository} commit drift"
                )
            if version_comment != expected_version:
                raise VerificationError(
                    f"{path.relative_to(ROOT)}:{line_number}: "
                    f"{repository} version comment drift"
                )
            seen.add(repository)

    unused = sorted(set(locked) - seen)
    if unused:
        raise VerificationError(f"manifested action is unused: {unused[0]}")
    if npm_ci_count < 1:
        raise VerificationError("no workflow performs a clean npm ci install")
    print(
        "[supply-chain] GitHub Actions: OK "
        f"({len(locked)} repositories pinned; {npm_ci_count} npm ci lanes)"
    )


def verify_update_automation() -> None:
    try:
        config = DEPENDABOT.read_text(encoding="utf-8")
    except OSError as error:
        raise VerificationError(f"cannot load .github/dependabot.yml: {error}") from error
    for ecosystem in ("github-actions", "npm"):
        if f'package-ecosystem: "{ecosystem}"' not in config:
            raise VerificationError(f"Dependabot does not track {ecosystem}")
    if config.count('interval: "weekly"') < 2:
        raise VerificationError("Dependabot update cadence must be weekly")
    print("[supply-chain] Dependabot coverage: OK (GitHub Actions and npm)")


def verify_vendor() -> None:
    proc = subprocess.run([sys.executable, str(VENDOR_VERIFIER)], cwd=ROOT)
    if proc.returncode != 0:
        raise VerificationError("vendored Reflaxe provenance verification failed")


def main() -> int:
    try:
        verify_package_lock()
        verify_actions_and_installs()
        verify_update_automation()
        verify_vendor()
    except (OSError, VerificationError) as error:
        print(f"[supply-chain] error: {error}", file=sys.stderr)
        return 1
    print("[supply-chain] supply-chain verification: OK")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
