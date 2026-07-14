#!/usr/bin/env python3

"""Fail closed when release-version or tested-SHA wiring drifts."""

from __future__ import annotations

import json
import sys
from pathlib import Path
from typing import NoReturn


ROOT = Path(__file__).resolve().parents[2]
DEVELOPMENT_VERSION = "0.0.0"
TAG_FORMAT = "v$" + "{version}"
ANALYZER = "./scripts/release/analyze-commits.mjs"
GITHUB_PLUGIN = "@semantic-release/github"


def fail(message: str) -> NoReturn:
    raise RuntimeError(message)


def load_json(path: Path) -> dict[str, object]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as exc:
        fail(f"cannot read {path.relative_to(ROOT)}: {exc}")
    if not isinstance(value, dict):
        fail(f"{path.relative_to(ROOT)} must contain a JSON object")
    return value


def require_file(relative_path: str) -> None:
    if not (ROOT / relative_path).is_file():
        fail(f"required release-policy file is missing: {relative_path}")


def verify_source_manifests() -> None:
    package = load_json(ROOT / "package.json")
    package_lock = load_json(ROOT / "package-lock.json")
    haxelib = load_json(ROOT / "haxelib.json")
    lock_packages = package_lock.get("packages")
    if not isinstance(lock_packages, dict) or not isinstance(lock_packages.get(""), dict):
        fail("package-lock.json is missing its root package record")

    versions = {
        "package.json": package.get("version"),
        "haxelib.json": haxelib.get("version"),
        "package-lock.json": package_lock.get("version"),
        "package-lock.json root package": lock_packages[""].get("version"),
    }
    for label, version in versions.items():
        if version != DEVELOPMENT_VERSION:
            fail(
                f"{label} must use the {DEVELOPMENT_VERSION} development sentinel, "
                f"got {version!r}"
            )
    if package.get("private") is not True:
        fail("package.json must remain private because npm is tooling, not version authority")
    if "Development checkout" not in str(haxelib.get("releasenote", "")):
        fail("haxelib.json releasenote must identify a Development checkout")


def verify_release_config() -> None:
    config = load_json(ROOT / ".releaserc.json")
    if config.get("branches") != ["master"]:
        fail("semantic-release branches must be exactly ['master']")
    if config.get("tagFormat") != TAG_FORMAT:
        fail(f"semantic-release tagFormat must be {TAG_FORMAT!r}")

    plugins = config.get("plugins")
    expected_names = [ANALYZER, GITHUB_PLUGIN]
    if not isinstance(plugins, list):
        fail("semantic-release plugins must be a list")
    names = [
        entry[0] if isinstance(entry, list) and entry else entry
        for entry in plugins
    ]
    if names != expected_names:
        fail(
            "semantic-release plugins must be the mutation-free analyzer and "
            f"GitHub publisher only; got {names!r}"
        )
    if (
        not isinstance(plugins[0], list)
        or len(plugins[0]) != 2
        or plugins[0][1] != {"approvedStableMajors": []}
    ):
        fail("release analyzer must declare the reviewed approvedStableMajors list")

    expected_github_options = {
        "successCommentCondition": False,
        "failCommentCondition": False,
        "releasedLabels": False,
        "addReleases": False,
    }
    if (
        not isinstance(plugins[1], list)
        or len(plugins[1]) != 2
        or plugins[1][1] != expected_github_options
    ):
        fail("GitHub release plugin options must disable issue and pull-request mutation")

    package = load_json(ROOT / "package.json")
    scripts = package.get("scripts")
    if not isinstance(scripts, dict):
        fail("package.json scripts must be an object")
    expected_scripts = {
        "release": "bash scripts/release/run-same-sha-release.sh",
        "release:dry-run": "bash scripts/release/run-same-sha-release.sh --dry-run",
        "release:stage-metadata": "python3 scripts/release/stage-release-metadata.py",
        "release:license-policy": "python3 scripts/release/verify-license-policy.py --mode release",
        "release:policy": "python3 scripts/release/verify-release-policy.py",
        "test:release-version-policy": "node test/test_release_version_policy.mjs",
    }
    for name, expected in expected_scripts.items():
        if scripts.get(name) != expected:
            fail(f"package script {name!r} must be {expected!r}")

    dependencies = package.get("devDependencies")
    if not isinstance(dependencies, dict):
        fail("package.json devDependencies must be an object")
    for forbidden in ("@semantic-release/changelog", "@semantic-release/git"):
        if forbidden in dependencies:
            fail(f"tracked-checkout release mutator remains installed: {forbidden}")
    if "semver" not in dependencies:
        fail("semver must be a direct locked dependency of the release policy")


def verify_same_sha_workflow() -> None:
    workflow_path = ROOT / ".github" / "workflows" / "ci-harness.yml"
    try:
        workflow = workflow_path.read_text(encoding="utf-8")
    except OSError as exc:
        fail(f"cannot read .github/workflows/ci-harness.yml: {exc}")
    marker = "\n  semantic-release:"
    if marker not in workflow:
        fail("ci-harness workflow has no semantic-release job")
    release_job = workflow.split(marker, 1)[1]
    github_sha = "$" + "{{ github.sha }}"
    manual_condition = (
        "if: github.event_name == 'workflow_dispatch' && "
        "inputs.publish_release && github.ref == 'refs/heads/master'"
    )
    if "workflow_dispatch:\n    inputs:\n      publish_release:" not in workflow:
        fail("ci-harness must declare the explicit publish_release input")
    required = (
        manual_condition,
        "permissions:\n      contents: write",
        "fetch-depth: 0",
        "ref: " + github_sha,
        "run: npm run release:policy",
        "run: npm run release:license-policy",
        "RELEASE_TESTED_SHA: " + github_sha,
        "run: npm run release",
    )
    for phrase in required:
        if phrase not in release_job:
            fail(f"semantic-release job is missing exact tested-SHA wiring: {phrase}")
    for forbidden in ("issues: write", "pull-requests: write"):
        if forbidden in release_job:
            fail(f"semantic-release job retains unnecessary permission: {forbidden}")
    if "github.event_name == 'push'" in release_job:
        fail("ordinary master pushes must not publish development releases")
    if "continue-on-error:" in release_job:
        fail("semantic-release job must not weaken a release-blocking step")
    for dependency in (
        "quality",
        "gitleaks",
        "dependency-audit",
        "go-tooling",
        "perf-go",
        "perf-apps",
    ):
        if "- " + dependency not in release_job:
            fail(f"semantic-release job no longer waits for required gate: {dependency}")


def main() -> int:
    try:
        for path in (
            "scripts/release/analyze-commits.mjs",
            "scripts/release/run-same-sha-release.sh",
            "scripts/release/stage-release-metadata.py",
            "scripts/release/verify-license-policy.py",
            "test/test_release_identity_contract.py",
            "test/test_release_version_policy.mjs",
            "test/test_same_sha_release_wrapper.py",
            "docs/release-version-policy.md",
            "LICENSING.md",
            "license-policy.json",
        ):
            require_file(path)
        wrapper = (ROOT / "scripts" / "release" / "run-same-sha-release.sh").read_text(
            encoding="utf-8"
        )
        if "verify-license-policy.py --mode release" not in wrapper:
            fail("same-SHA release wrapper does not enforce the licensing gate")
        verify_source_manifests()
        verify_release_config()
        verify_same_sha_workflow()
    except RuntimeError as exc:
        print(f"[release-policy] error: {exc}", file=sys.stderr)
        return 1

    print(
        "[release-policy] tag-owned version lineage: OK "
        f"(source sentinel {DEVELOPMENT_VERSION})"
    )
    print("[release-policy] mutation-free semantic-release config: OK")
    print("[release-policy] same-tested-SHA workflow: OK")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
