#!/usr/bin/env python3
"""Build a provenance-bound evidence bundle for the Haxe.Go GPT-5.6 Pro review."""

from __future__ import annotations

import argparse
import hashlib
import json
import os
import re
import shutil
import stat
import subprocess
import sys
import tarfile
import tempfile
import zipfile
from pathlib import Path, PurePosixPath
from typing import Any, Iterable, Sequence


ROOT = Path(__file__).resolve().parents[2]
ZIP_TIMESTAMP = (1980, 1, 1, 0, 0, 0)
PINNED_REPOMIX_PACKAGE = "repomix@1.14.0"
PRIMARY_EXCLUSIONS = (
    ".beads/issues.jsonl",
    ".beads/interactions.jsonl",
)
SECRET_PATH_SUFFIXES = (".pem", ".key")

# These are deliberately narrow evidence slices, not sibling source dependencies.
# Every path is read through `git archive <commit>`, so dirty sibling worktrees cannot
# leak into the review bundle.
REFERENCE_PATHS: dict[str, tuple[str, ...]] = {
    "haxe.rust": (
        ".github/workflows/ci.yml",
        ".github/workflows/release-repair.yml",
        ".github/workflows/weekly-ci-evidence.yml",
        "package.json",
        "package-lock.json",
        "haxelib.json",
        "release-manifest.json",
        "release.config.js",
        "docs/oracle-gpt-5.6-production-readiness-review.md",
        "docs/production-readiness-audit-2026-07-13.md",
        "docs/public-compatibility-manifest.json",
        "docs/release-reference-architecture.md",
        "docs/semver-release-posture.md",
        "scripts/ci/upstream-stdlib-boundary-check.js",
        "scripts/release",
        "src/reflaxe/rust/CompilerBootstrap.hx",
        "src/reflaxe/rust/analyze/NoHxrtEligibilityAnalyzer.hx",
        "src/reflaxe/rust/analyze/RuntimeRequirementAnalyzer.hx",
        "src/reflaxe/rust/analyze/SurfaceContractRegistry.hx",
        "src/reflaxe/rust/passes/NoHxrtPass.hx",
        "std/rust/_std",
        "test/fixtures/release-lifecycle-plugin.cjs",
        "test/scripts/release-artifact.test.js",
        "test/scripts/release-lifecycle.test.js",
        "test/scripts/release-notes.test.js",
        "test/scripts/release-policy.test.js",
        "test/scripts/release-provenance.test.js",
        "test/scripts/release-workflow.test.js",
        "test/snapshot/portable_facade_contract_report",
        "test/snapshot/portable_facade_native_option_result",
    ),
    "haxe.ruby": (
        ".github/workflows/ci.yml",
        ".github/workflows/release-repair.yml",
        "package.json",
        "package-lock.json",
        "haxelib.json",
        "hxruby.gemspec",
        "docs/release-hosting-and-repair.md",
        "docs/release-version-policy.md",
        "scripts/ci/release-artifact-reproducibility-check.js",
        "scripts/ci/release-contracts-check.js",
        "scripts/ci/release-hosting-check.mjs",
        "scripts/ci/release-version-policy-check.mjs",
        "scripts/ci/release-workflow-check.js",
        "scripts/release",
    ),
    "haxe.elixir.codex": (
        ".github/workflows/ci.yml",
        ".github/workflows/release-repair.yml",
        "package.json",
        "package-lock.json",
        "haxelib.json",
        "release.config.js",
        "release/manifest.json",
        "scripts/ci/check-stdlib-source-layout.sh",
        "scripts/ci/haxelib-package-smoke.sh",
        "scripts/release",
        "src/reflaxe/elixir/CompilerBootstrap.hx",
        "src/reflaxe/elixir/CompilerInit.hx",
        "std/elixir/_std",
    ),
}

PRIMARY_EVIDENCE_INDEX: dict[str, tuple[str, ...]] = {
    "product_and_architecture": (
        "AGENTS.md",
        "README.md",
        "docs/start-here.md",
        "docs/profiles.md",
        "docs/profile-semantics-guide.md",
        "docs/known-gaps.md",
        "src/reflaxe/go/**",
        "runtime/hxrt/**",
        "std/**",
        "vendor/reflaxe/**",
    ),
    "generated_portable_and_metal": (
        "examples/**/generated/portable/**",
        "examples/**/generated/metal/**",
        "test/snapshot/**/intended/**",
        "benchmarks/pure_go/**",
    ),
    "compatibility_and_ownership": (
        "docs/feature-support-matrix.md",
        "docs/ownership-rubric.md",
        "docs/portable-module-mapping-contract.md",
        "docs/portable-stdlib-parity-program.md",
        "docs/stdlib-provenance-ledger.json",
        "docs/stdlib-shim-migration-log.md",
        "docs/stdlib-shim-rationale.md",
        "test/portable_stdlib_inventory.json",
    ),
    "release_and_security": (
        ".releaserc.json",
        ".github/workflows/**",
        "package.json",
        "haxelib.json",
        "scripts/release/**",
        "scripts/security/**",
        "docs/release-readiness-checklist.md",
        "docs/release-visibility.md",
        "docs/security-dependency-audit.md",
    ),
    "executable_evidence": (
        "test/run-ci.py",
        "test/run-snapshots.py",
        "test/run-semantic-diff.py",
        "test/run-upstream-stdlib-sweep.py",
        "test/run-portable-stdlib-inventory.py",
        "test/run-portable-parity-closure.py",
        "test/semantic_diff/**",
        "test/semantic_diff_lanes/**",
        "test/snapshot/**",
    ),
}

ANSI_ESCAPE = re.compile(r"\x1b\[[0-?]*[ -/]*[@-~]")
USERS_ROOT = "/" + "Users" + "/"
GITHUB_RUNNER_WORK_ROOT = "/" + "home/runner/work" + "/"
MACOS_PRIVATE_TEMP_ROOT = "/" + "private/var/folders" + "/"
MACHINE_PATH_PATTERNS = (
    re.compile(re.escape(USERS_ROOT) + r"[A-Za-z0-9._-]+/(?:[^/\s]+/)*?(?:haxe\.go|reflaxe\.go)/"),
    re.compile(re.escape(GITHUB_RUNNER_WORK_ROOT) + r"[A-Za-z0-9._-]+/[A-Za-z0-9._-]+/"),
    re.compile(r"[A-Za-z]:\\a\\[^\\\s]+\\[^\\\s]+\\"),
    re.compile(re.escape(MACOS_PRIVATE_TEMP_ROOT) + r"[A-Za-z0-9._-]+/[^\s]+"),
)


class EvidenceError(RuntimeError):
    pass


def run(
    command: Sequence[str],
    *,
    cwd: Path = ROOT,
    check: bool = True,
    text: bool = True,
) -> subprocess.CompletedProcess[str] | subprocess.CompletedProcess[bytes]:
    completed = subprocess.run(
        list(command),
        cwd=cwd,
        capture_output=True,
        text=text,
    )
    if check and completed.returncode != 0:
        stdout = completed.stdout if isinstance(completed.stdout, str) else completed.stdout.decode("utf-8", "replace")
        stderr = completed.stderr if isinstance(completed.stderr, str) else completed.stderr.decode("utf-8", "replace")
        raise EvidenceError(
            f"command failed ({completed.returncode}): {' '.join(command)}\n"
            f"stdout:\n{redact_machine_paths(stdout)}\n"
            f"stderr:\n{redact_machine_paths(stderr)}"
        )
    return completed


def write_json(path: Path, value: Any) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2, sort_keys=True) + "\n", encoding="utf-8")


def sha256_file(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as handle:
        for chunk in iter(lambda: handle.read(1024 * 1024), b""):
            digest.update(chunk)
    return digest.hexdigest()


def write_deterministic_zip(source_dir: Path, output_path: Path) -> None:
    """Write sorted files with stable timestamps, modes, and compression."""
    output_path.parent.mkdir(parents=True, exist_ok=True)
    files = sorted(path for path in source_dir.rglob("*") if path.is_file())
    with zipfile.ZipFile(output_path, "w", compression=zipfile.ZIP_DEFLATED, compresslevel=9) as archive:
        for path in files:
            relative = path.relative_to(source_dir).as_posix()
            info = zipfile.ZipInfo(relative, date_time=ZIP_TIMESTAMP)
            info.compress_type = zipfile.ZIP_DEFLATED
            info.create_system = 3
            mode = stat.S_IMODE(path.stat().st_mode)
            info.external_attr = (stat.S_IFREG | mode) << 16
            archive.writestr(info, path.read_bytes(), compress_type=zipfile.ZIP_DEFLATED, compresslevel=9)


def redact_machine_paths(text: str) -> str:
    text = ANSI_ESCAPE.sub("", text)
    home = Path.home().as_posix().rstrip("/")
    if home and home != "/":
        text = text.replace(home + "/", "<local-home>/")
    text = MACHINE_PATH_PATTERNS[0].sub("<local-workspace>/", text)
    text = MACHINE_PATH_PATTERNS[1].sub("<github-workspace>/", text)
    text = re.sub(re.escape(GITHUB_RUNNER_WORK_ROOT) + r"[^\s\"'<>]+", "<github-runner-work>", text)
    text = MACHINE_PATH_PATTERNS[2].sub("<github-workspace>/", text)
    text = MACHINE_PATH_PATTERNS[3].sub("<temporary-path>", text)
    text = text.replace("\\", "/") if "<github-workspace>/" in text else text
    return text


def find_machine_path(text: str) -> str | None:
    home = Path.home().as_posix().rstrip("/")
    if home and home != "/" and home + "/" in text:
        return home + "/"
    probes = (
        MACHINE_PATH_PATTERNS[0],
        MACHINE_PATH_PATTERNS[1],
        re.compile(re.escape(GITHUB_RUNNER_WORK_ROOT) + r"_[A-Za-z0-9._-]+/"),
        re.compile(r"[A-Za-z]:\\a\\"),
        re.compile(re.escape(MACOS_PRIVATE_TEMP_ROOT) + r"[A-Za-z0-9._-]+/"),
    )
    for probe in probes:
        match = probe.search(text)
        if match:
            return match.group(0)
    return None


def normalize_repomix_log(log: str) -> str:
    """Keep stable completion/security evidence and discard spinner progress."""
    marker = "✔ Packing completed successfully!"
    marker_index = log.find(marker)
    if marker_index < 0:
        raise EvidenceError("Repomix log does not contain a successful completion marker")
    stable_lines: list[str] = []
    for line in log[marker_index:].splitlines():
        if re.match(r"^\s*Output:", line):
            stable_lines.append("       Output: <bundle-output>")
        else:
            stable_lines.append(line.rstrip())
    return "\n".join(stable_lines).strip() + "\n"


def parse_repomix_security_exclusions(log: str) -> list[str]:
    exclusions: list[str] = []
    in_exclusion_section = False
    for line in log.splitlines():
        stripped = line.strip()
        if re.fullmatch(r"\d+ suspicious file\(s\) detected and excluded from the output:", stripped):
            in_exclusion_section = True
            continue
        if not in_exclusion_section:
            continue
        if stripped.startswith("These files have been excluded") or stripped.startswith("📊 Pack Summary"):
            break
        match = re.fullmatch(r"\d+\.\s+(.+)", stripped)
        if match:
            path = match.group(1)
            posix_path = PurePosixPath(path)
            if posix_path.is_absolute() or ".." in posix_path.parts:
                raise EvidenceError(f"unsafe Repomix security-exclusion path: {path}")
            exclusions.append(path)
    return exclusions


def normalize_gitleaks_log(log: str) -> str:
    if "no leaks found" not in log.lower():
        raise EvidenceError("Gitleaks completed without its expected no-leaks confirmation")
    return "PASS: no leaks found (max-archive-depth=1).\n"


def resolve_commit(repo: Path, ref: str) -> str:
    completed = run(["git", "rev-parse", f"{ref}^{{commit}}"], cwd=repo)
    assert isinstance(completed.stdout, str)
    commit = completed.stdout.strip()
    if not re.fullmatch(r"[0-9a-f]{40}", commit):
        raise EvidenceError(f"could not resolve full commit for {repo}: {ref!r}")
    return commit


def remote_branches_containing(repo: Path, commit: str) -> list[str]:
    completed = run(
        ["git", "branch", "-r", "--contains", commit, "--format=%(refname:short)"],
        cwd=repo,
    )
    assert isinstance(completed.stdout, str)
    return sorted(line.strip() for line in completed.stdout.splitlines() if line.strip())


def require_origin_containment(repo: Path, commit: str) -> list[str]:
    branches = remote_branches_containing(repo, commit)
    if not any(branch.startswith("origin/") for branch in branches):
        raise EvidenceError(f"commit {commit} in {repo} is not contained by an origin branch")
    return branches


def list_git_tree(repo: Path, commit: str) -> list[dict[str, str]]:
    completed = run(
        ["git", "ls-tree", "-rz", "--full-tree", commit],
        cwd=repo,
        text=False,
    )
    assert isinstance(completed.stdout, bytes)
    entries: list[dict[str, str]] = []
    for raw_entry in completed.stdout.split(b"\0"):
        if not raw_entry:
            continue
        metadata, raw_path = raw_entry.split(b"\t", 1)
        mode, object_type, object_id = metadata.decode("ascii").split(" ")
        entries.append(
            {
                "mode": mode,
                "type": object_type,
                "object_id": object_id,
                "path": raw_path.decode("utf-8", "surrogateescape"),
            }
        )
    return entries


def is_primary_excluded(path: str) -> bool:
    if path in PRIMARY_EXCLUSIONS:
        return True
    if path.startswith("infra/secrets/"):
        return True
    return path.endswith(SECRET_PATH_SUFFIXES)


def require_no_secret_paths(entries: Iterable[dict[str, str]]) -> None:
    secret_paths = [entry["path"] for entry in entries if is_primary_excluded(entry["path"]) and entry["path"] not in PRIMARY_EXCLUSIONS]
    if secret_paths:
        raise EvidenceError("tracked secret-bearing paths cannot enter evidence: " + ", ".join(secret_paths))


def write_git_archive(
    repo: Path,
    commit: str,
    output_path: Path,
    *,
    prefix: str,
    paths: Sequence[str] | None = None,
    exclusions: Sequence[str] = (),
) -> None:
    output_path.parent.mkdir(parents=True, exist_ok=True)
    command = ["git", "archive", "--format=tar", f"--prefix={prefix.rstrip('/')}/", commit]
    if paths:
        command.extend(["--", *paths])
    elif exclusions:
        command.extend(["--", ".", *[f":(exclude){path}" for path in exclusions]])
    with output_path.open("wb") as output:
        completed = subprocess.run(command, cwd=repo, stdout=output, stderr=subprocess.PIPE)
    if completed.returncode != 0:
        stderr = completed.stderr.decode("utf-8", "replace")
        raise EvidenceError(f"git archive failed for {repo}: {redact_machine_paths(stderr)}")


def safe_extract_tar(archive_path: Path, destination: Path) -> None:
    destination.mkdir(parents=True, exist_ok=True)
    destination_resolved = destination.resolve()
    with tarfile.open(archive_path, "r:") as archive:
        for member in archive.getmembers():
            member_path = PurePosixPath(member.name)
            if member_path.is_absolute() or ".." in member_path.parts:
                raise EvidenceError(f"unsafe archive member: {member.name}")
            resolved = (destination / member.name).resolve()
            if destination_resolved not in (resolved, *resolved.parents):
                raise EvidenceError(f"archive member escapes destination: {member.name}")
            if member.issym() or member.islnk():
                link_path = PurePosixPath(member.linkname)
                if link_path.is_absolute() or ".." in link_path.parts:
                    raise EvidenceError(f"unsafe archive link: {member.name} -> {member.linkname}")
        archive.extractall(destination, filter="data")


def run_repomix(repomix_command: Sequence[str], input_dir: Path, output_path: Path, header: str) -> dict[str, Any]:
    version_result = run([*repomix_command, "--version"])
    assert isinstance(version_result.stdout, str)
    command = [
        *repomix_command,
        ".",
        "--output",
        str(output_path),
        "--style",
        "xml",
        "--parsable-style",
        "--output-show-line-numbers",
        "--no-default-patterns",
        "--ignore",
        ".git/**,node_modules/**,dist/**,repomix-output*,.beads/issues.jsonl,.beads/interactions.jsonl,**/*.pem,**/*.key,infra/secrets/**",
        "--header-text",
        header,
    ]
    completed = run(command, cwd=input_dir)
    assert isinstance(completed.stdout, str)
    assert isinstance(completed.stderr, str)
    log = normalize_repomix_log(redact_machine_paths(completed.stdout + completed.stderr))
    output_text = output_path.read_text(encoding="utf-8")
    leaked_path = find_machine_path(output_text)
    if leaked_path:
        raise EvidenceError(f"Repomix output contains machine-local path: {leaked_path}")
    return {
        "version": version_result.stdout.strip(),
        "command": [
            *repomix_command,
            ".",
            "--style",
            "xml",
            "--parsable-style",
            "--output-show-line-numbers",
            "--no-default-patterns",
        ],
        "security_excluded_paths": parse_repomix_security_exclusions(log),
        "log": log,
    }


def copy_repomix_security_exclusions(source_root: Path, primary_dir: Path, paths: Sequence[str]) -> None:
    destination_root = primary_dir / "repomix-security-exclusions"
    for relative in paths:
        source = source_root / relative
        if not source.is_file():
            raise EvidenceError(f"Repomix reported an exclusion that is not a source file: {relative}")
        destination = destination_root / relative
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, destination)
    write_json(
        primary_dir / "repomix-security-exclusions.json",
        {
            "reason": (
                "Repomix omitted these tracked test fixtures during its security scan. "
                "They remain in the canonical git archive and are copied here for reviewer completeness; "
                "the complete bundle passes Gitleaks."
            ),
            "paths": list(paths),
        },
    )


def parse_json_output(command: Sequence[str], *, cwd: Path = ROOT) -> Any:
    completed = run(command, cwd=cwd)
    assert isinstance(completed.stdout, str)
    try:
        return json.loads(completed.stdout)
    except json.JSONDecodeError as error:
        raise EvidenceError(f"command did not return JSON: {' '.join(command)}: {error}") from error


def capture_endpoint(command: Sequence[str]) -> dict[str, Any]:
    completed = run(command, check=False)
    assert isinstance(completed.stdout, str)
    assert isinstance(completed.stderr, str)
    if completed.returncode == 0:
        try:
            value = json.loads(completed.stdout)
        except json.JSONDecodeError:
            value = {"raw": redact_machine_paths(completed.stdout)}
        return {"ok": True, "value": value}
    return {
        "ok": False,
        "exit_code": completed.returncode,
        "stderr": redact_machine_paths(completed.stderr).strip(),
    }


def capture_github_evidence(
    bundle_root: Path,
    *,
    github_repo: str,
    source_commit: str,
    ci_runs: Sequence[int],
    release_tag: str,
) -> dict[str, Any]:
    live_dir = bundle_root / "live" / "github"
    runs_dir = live_dir / "runs"
    runs_dir.mkdir(parents=True, exist_ok=True)

    repository = parse_json_output(["gh", "api", f"repos/{github_repo}"])
    write_json(live_dir / "repository.json", repository)
    default_branch = repository.get("default_branch", "master")

    host_controls = {
        "rulesets": capture_endpoint(["gh", "api", f"repos/{github_repo}/rulesets?includes_parents=true"]),
        "default_branch_protection": capture_endpoint(
            ["gh", "api", f"repos/{github_repo}/branches/{default_branch}/protection"]
        ),
    }
    write_json(live_dir / "host-controls.json", host_controls)

    release = parse_json_output(["gh", "api", f"repos/{github_repo}/releases/tags/{release_tag}"])
    tag_ref = parse_json_output(["gh", "api", f"repos/{github_repo}/git/ref/tags/{release_tag}"])
    write_json(live_dir / f"release-{release_tag}.json", release)
    write_json(live_dir / f"tag-{release_tag}.json", tag_ref)

    run_summaries: list[dict[str, Any]] = []
    for run_id in sorted(set(ci_runs)):
        summary = parse_json_output(
            [
                "gh",
                "run",
                "view",
                str(run_id),
                "--repo",
                github_repo,
                "--json",
                "databaseId,name,displayTitle,workflowName,status,conclusion,headBranch,headSha,event,createdAt,updatedAt,url,jobs",
            ]
        )
        if summary.get("headSha") != source_commit:
            raise EvidenceError(
                f"GitHub run {run_id} is for {summary.get('headSha')}, expected {source_commit}"
            )
        if summary.get("status") != "completed":
            raise EvidenceError(f"GitHub run {run_id} is not complete: {summary.get('status')}")
        if summary.get("conclusion") != "success":
            raise EvidenceError(f"GitHub run {run_id} did not succeed: {summary.get('conclusion')}")

        jobs = parse_json_output(
            ["gh", "api", f"repos/{github_repo}/actions/runs/{run_id}/jobs?per_page=100"]
        )
        artifacts = parse_json_output(
            ["gh", "api", f"repos/{github_repo}/actions/runs/{run_id}/artifacts?per_page=100"]
        )
        write_json(runs_dir / f"{run_id}.json", {"summary": summary, "jobs": jobs, "artifacts": artifacts})

        log_result = run(["gh", "run", "view", str(run_id), "--repo", github_repo, "--log"])
        assert isinstance(log_result.stdout, str)
        assert isinstance(log_result.stderr, str)
        sanitized_log = redact_machine_paths(log_result.stdout + log_result.stderr)
        leaked_path = find_machine_path(sanitized_log)
        if leaked_path:
            raise EvidenceError(f"sanitized GitHub log {run_id} still contains {leaked_path}")
        (runs_dir / f"{run_id}.log").write_text(sanitized_log, encoding="utf-8")
        run_summaries.append(summary)

    return {
        "repository": github_repo,
        "default_branch": default_branch,
        "release_tag": release_tag,
        "release_id": release.get("id"),
        "release_immutable": release.get("immutable"),
        "release_asset_count": len(release.get("assets", [])),
        "release_tag_object": tag_ref.get("object", {}),
        "ci_runs": [
            {
                "database_id": summary.get("databaseId"),
                "workflow": summary.get("workflowName"),
                "conclusion": summary.get("conclusion"),
                "head_sha": summary.get("headSha"),
                "url": summary.get("url"),
            }
            for summary in run_summaries
        ],
    }


def capture_roadmap(bundle_root: Path, roadmap_id: str) -> dict[str, Any]:
    roadmap = parse_json_output(["bd", "show", roadmap_id, "--json"])
    dolt_status = parse_json_output(["bd", "vc", "status", "--json"])
    stats = parse_json_output(["bd", "stats", "--json"])
    cycles = parse_json_output(["bd", "dep", "cycles", "--json"])
    value = {
        "roadmap": roadmap,
        "dolt_status": dolt_status,
        "stats": stats,
        "dependency_cycles": cycles,
    }
    text = json.dumps(value, sort_keys=True)
    leaked_path = find_machine_path(text)
    if leaked_path:
        raise EvidenceError(f"roadmap evidence contains machine-local path: {leaked_path}")
    write_json(bundle_root / "roadmap" / "haxe-go-next.json", value)
    return {
        "root_issue": roadmap_id,
        "dolt_commit": dolt_status.get("commit"),
        "total_issues": stats.get("summary", {}).get("total_issues"),
        "dependency_cycle_count": len(cycles) if isinstance(cycles, list) else None,
    }


def expand_selected_paths(repo: Path, commit: str, selectors: Sequence[str]) -> list[str]:
    all_paths = [entry["path"] for entry in list_git_tree(repo, commit)]
    selected: set[str] = set()
    for selector in selectors:
        prefix = selector.rstrip("/") + "/"
        matches = [path for path in all_paths if path == selector or path.startswith(prefix)]
        if not matches:
            raise EvidenceError(f"reference selector does not exist at {commit}: {selector}")
        selected.update(matches)
    return sorted(selected)


def capture_reference(
    references_root: Path,
    *,
    name: str,
    repo: Path,
    ref: str,
    selectors: Sequence[str],
    temp_root: Path,
) -> dict[str, Any]:
    commit = resolve_commit(repo, ref)
    remote_branches = require_origin_containment(repo, commit)
    remote_url_result = run(["git", "remote", "get-url", "origin"], cwd=repo)
    assert isinstance(remote_url_result.stdout, str)
    selected_paths = expand_selected_paths(repo, commit, selectors)
    archive_path = temp_root / f"{name}.tar"
    write_git_archive(repo, commit, archive_path, prefix=name, paths=selectors)
    safe_extract_tar(archive_path, references_root)
    return {
        "name": name,
        "remote": remote_url_result.stdout.strip(),
        "commit": commit,
        "requested_ref": ref,
        "remote_branches_containing_commit": remote_branches,
        "selectors": list(selectors),
        "selected_file_count": len(selected_paths),
        "selected_files": selected_paths,
    }


def build_primary_inventory(repo: Path, commit: str) -> dict[str, Any]:
    entries = list_git_tree(repo, commit)
    require_no_secret_paths(entries)
    included = [entry for entry in entries if not is_primary_excluded(entry["path"])]
    paths = [entry["path"] for entry in included]
    return {
        "commit": commit,
        "file_count": len(included),
        "excluded_tracked_files": list(PRIMARY_EXCLUSIONS),
        "counts": {
            "haxe_source_files": sum(path.endswith(".hx") for path in paths),
            "checked_in_cross_hx_files": sum(path.endswith(".cross.hx") for path in paths),
            "generated_portable_files": sum("/generated/portable/" in path for path in paths),
            "generated_metal_files": sum("/generated/metal/" in path for path in paths),
            "snapshot_intended_files": sum("/intended/" in path for path in paths),
            "workflow_files": sum(path.startswith(".github/workflows/") for path in paths),
        },
        "evidence_index": PRIMARY_EVIDENCE_INDEX,
        "files": included,
    }


def assert_no_machine_paths(root: Path) -> int:
    checked = 0
    binary_suffixes = {".tar", ".zip", ".png", ".jpg", ".jpeg", ".gif"}
    for path in sorted(item for item in root.rglob("*") if item.is_file()):
        if path.suffix.lower() in binary_suffixes:
            continue
        try:
            text = path.read_text(encoding="utf-8")
        except UnicodeDecodeError:
            continue
        checked += 1
        leaked_path = find_machine_path(text)
        if leaked_path:
            raise EvidenceError(
                f"machine-local path remains in {path.relative_to(root).as_posix()}: {leaked_path}"
            )
    return checked


def run_gitleaks(bundle_root: Path) -> str:
    command = [
        "gitleaks",
        "dir",
        "--no-banner",
        "--no-color",
        "--redact",
        "--max-archive-depth",
        "1",
        "--config",
        str(ROOT / ".gitleaks.toml"),
        str(bundle_root),
    ]
    completed = run(command)
    assert isinstance(completed.stdout, str)
    assert isinstance(completed.stderr, str)
    return normalize_gitleaks_log(redact_machine_paths(completed.stdout + completed.stderr))


def payload_records(bundle_root: Path) -> list[dict[str, Any]]:
    excluded = {"MANIFEST.json", "SHA256SUMS"}
    records: list[dict[str, Any]] = []
    for path in sorted(item for item in bundle_root.rglob("*") if item.is_file()):
        relative = path.relative_to(bundle_root).as_posix()
        if relative in excluded:
            continue
        records.append(
            {
                "path": relative,
                "bytes": path.stat().st_size,
                "sha256": sha256_file(path),
            }
        )
    return records


def write_sha256sums(bundle_root: Path) -> None:
    lines: list[str] = []
    for path in sorted(item for item in bundle_root.rglob("*") if item.is_file()):
        relative = path.relative_to(bundle_root).as_posix()
        if relative == "SHA256SUMS":
            continue
        lines.append(f"{sha256_file(path)}  {relative}")
    (bundle_root / "SHA256SUMS").write_text("\n".join(lines) + "\n", encoding="utf-8")


def bundle_readme(
    *,
    source_commit: str,
    source_branches: Sequence[str],
    github_repo: str,
    release_tag: str,
    ci_runs: Sequence[int],
    references: Sequence[dict[str, Any]],
    output_name: str,
) -> str:
    reference_lines = "\n".join(
        f"- `{reference['name']}` at `{reference['commit']}` ({reference['selected_file_count']} selected files)"
        for reference in references
    )
    run_args = " ".join(f"--ci-run {run_id}" for run_id in sorted(set(ci_runs)))
    return f"""# Haxe.Go GPT-5.6 Pro review evidence

This bundle is an evidence snapshot, not a source checkout and not a release artifact.

## Primary authority

- Repository: `{github_repo}`
- Exact source commit: `{source_commit}`
- Remote branches containing the commit: `{', '.join(source_branches)}`
- Release state captured for: `{release_tag}`

`primary/haxe.go-source-{source_commit[:8]}.tar` is produced directly by `git archive`.
`primary/haxe.go-source-{source_commit[:8]}.xml` is the line-numbered Repomix review view of that archive.

## Contents

- exact source plus compiler/runtime/std/tests/docs/workflows;
- committed generated portable and metal example output and intended snapshot output;
- compatibility, stdlib, ownership, profile, and release inventories;
- sanitized logs and job metadata for GitHub Actions runs `{', '.join(str(run_id) for run_id in sorted(set(ci_runs)))}`;
- live release, tag, ruleset, and default-branch-protection API state;
- a read-only Haxe.Go Next roadmap and Dolt status snapshot;
- filtered sibling reference designs:
{reference_lines}

## Exclusions and redactions

- `.git`, `node_modules`, caches, build output, `dist`, and untracked files never enter `git archive`;
- `.beads/issues.jsonl` is excluded because it is the immutable pre-Dolt history archive, not compiler evidence;
- `.beads/interactions.jsonl` is excluded because the operational roadmap snapshot is supplied separately;
- secret-bearing path classes (`*.pem`, `*.key`, `infra/secrets/**`) fail the build;
- local home and hosted-runner workspace paths in captured logs are replaced with `<local-home>`, `<local-workspace>`, or `<github-workspace>`;
- Repomix security scanning and a final Gitleaks scan run before packaging;
- tracked fixtures omitted by Repomix's heuristic scanner are listed and copied under `primary/repomix-security-exclusions/` so an upload omission cannot be misreported as a source defect.

## Integrity

`MANIFEST.json` records authorities, exclusions, tool versions, and every payload SHA-256.
`SHA256SUMS` covers every bundle file except itself. The outer ZIP digest is recorded in the
repository review record rather than recursively inside the ZIP.

## Reproduction

Run from a checkout containing `scripts/review/build_gpt56_evidence.py` with the sibling repositories
available at the documented relative paths:

```bash
python3 scripts/review/build_gpt56_evidence.py \\
  --source-ref {source_commit} \\
  --rust-ref {next(reference['commit'] for reference in references if reference['name'] == 'haxe.rust')} \\
  --ruby-ref {next(reference['commit'] for reference in references if reference['name'] == 'haxe.ruby')} \\
  --elixir-ref {next(reference['commit'] for reference in references if reference['name'] == 'haxe.elixir.codex')} \\
  {run_args} \\
  --release-tag {release_tag} \\
  --output dist/review/{output_name}
```
"""


def parse_args(argv: Sequence[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Build a commit-pinned, secret-scanned GPT-5.6 Pro review evidence bundle."
    )
    parser.add_argument("--source-ref", required=True, help="Exact haxe.go commit or ref to archive")
    parser.add_argument("--rust-repo", type=Path, default=Path("../haxe.rust"))
    parser.add_argument("--rust-ref", required=True, help="Exact haxe.rust reference commit")
    parser.add_argument("--ruby-repo", type=Path, default=Path("../haxe.ruby"))
    parser.add_argument("--ruby-ref", required=True, help="Exact haxe.ruby reference commit")
    parser.add_argument("--elixir-repo", type=Path, default=Path("../haxe.elixir.codex"))
    parser.add_argument("--elixir-ref", required=True, help="Exact haxe.elixir reference commit")
    parser.add_argument("--github-repo", default="fullofcaffeine/reflaxe.go")
    parser.add_argument("--ci-run", type=int, action="append", required=True, help="Successful exact-SHA Actions run ID")
    parser.add_argument("--release-tag", required=True, help="Release/tag whose live state should be captured")
    parser.add_argument("--roadmap-id", default="haxe_go-vfp")
    parser.add_argument(
        "--repomix-bin",
        help=f"Override the default pinned `npx --yes {PINNED_REPOMIX_PACKAGE}` command with one executable",
    )
    parser.add_argument("--output", type=Path, required=True, help="Destination ZIP (normally under ignored dist/review)")
    return parser.parse_args(argv)


def build(args: argparse.Namespace) -> dict[str, Any]:
    if args.output.suffix.lower() != ".zip":
        raise EvidenceError("--output must end in .zip")
    source_commit = resolve_commit(ROOT, args.source_ref)
    source_branches = require_origin_containment(ROOT, source_commit)
    source_remote_result = run(["git", "remote", "get-url", "origin"])
    assert isinstance(source_remote_result.stdout, str)

    reference_inputs = (
        ("haxe.rust", args.rust_repo.resolve(), args.rust_ref),
        ("haxe.ruby", args.ruby_repo.resolve(), args.ruby_ref),
        ("haxe.elixir.codex", args.elixir_repo.resolve(), args.elixir_ref),
    )
    for name, repo, _ in reference_inputs:
        if not (repo / ".git").exists():
            raise EvidenceError(f"{name} repository is unavailable: {repo}")

    output = args.output.resolve()
    repomix_command = [args.repomix_bin] if args.repomix_bin else ["npx", "--yes", PINNED_REPOMIX_PACKAGE]
    output.parent.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory(prefix="haxe-go-gpt56-evidence.") as raw_temp:
        temp_root = Path(raw_temp)
        bundle_name = f"haxe-go-gpt56-evidence-{source_commit[:8]}"
        bundle_root = temp_root / bundle_name
        primary_dir = bundle_root / "primary"
        primary_dir.mkdir(parents=True)

        primary_inventory = build_primary_inventory(ROOT, source_commit)
        write_json(primary_dir / "source-inventory.json", primary_inventory)
        source_archive = primary_dir / f"haxe.go-source-{source_commit[:8]}.tar"
        write_git_archive(
            ROOT,
            source_commit,
            source_archive,
            prefix="haxe.go",
            exclusions=PRIMARY_EXCLUSIONS,
        )

        extracted_source = temp_root / "source"
        safe_extract_tar(source_archive, extracted_source)
        source_repomix = primary_dir / f"haxe.go-source-{source_commit[:8]}.xml"
        source_repomix_meta = run_repomix(
            repomix_command,
            extracted_source / "haxe.go",
            source_repomix,
            (
                "Haxe.Go GPT-5.6 Pro primary source evidence. "
                f"Exact Git commit: {source_commit}. "
                "See bundle MANIFEST.json for exclusions and live evidence authorities."
            ),
        )
        source_repomix_exclusions = source_repomix_meta["security_excluded_paths"]
        copy_repomix_security_exclusions(
            extracted_source / "haxe.go",
            primary_dir,
            source_repomix_exclusions,
        )
        (primary_dir / "repomix.log").write_text(source_repomix_meta.pop("log"), encoding="utf-8")

        references_root = bundle_root / "references"
        references_root.mkdir(parents=True)
        references: list[dict[str, Any]] = []
        for name, repo, ref in reference_inputs:
            references.append(
                capture_reference(
                    references_root,
                    name=name,
                    repo=repo,
                    ref=ref,
                    selectors=REFERENCE_PATHS[name],
                    temp_root=temp_root,
                )
            )
        write_json(references_root / "reference-manifest.json", references)
        reference_repomix_temp = temp_root / "reflaxe-family-references.xml"
        reference_repomix_meta = run_repomix(
            repomix_command,
            references_root,
            reference_repomix_temp,
            (
                "Filtered, commit-pinned sibling evidence for Haxe.Go review: "
                "Rust portable/metal admission and release architecture, Ruby release architecture, "
                "and Elixir canonical Reflaxe _std layout."
            ),
        )
        reference_repomix = references_root / "reflaxe-family-references.xml"
        shutil.move(reference_repomix_temp, reference_repomix)
        (references_root / "repomix.log").write_text(reference_repomix_meta.pop("log"), encoding="utf-8")

        github = capture_github_evidence(
            bundle_root,
            github_repo=args.github_repo,
            source_commit=source_commit,
            ci_runs=args.ci_run,
            release_tag=args.release_tag,
        )
        roadmap = capture_roadmap(bundle_root, args.roadmap_id)

        readme = bundle_readme(
            source_commit=source_commit,
            source_branches=source_branches,
            github_repo=args.github_repo,
            release_tag=args.release_tag,
            ci_runs=args.ci_run,
            references=references,
            output_name=output.name,
        )
        (bundle_root / "README.md").write_text(readme, encoding="utf-8")
        tooling_dir = bundle_root / "tooling"
        tooling_dir.mkdir()
        shutil.copy2(Path(__file__), tooling_dir / Path(__file__).name)

        validation_dir = bundle_root / "validation"
        validation_dir.mkdir()
        checked_text_files = assert_no_machine_paths(bundle_root)
        (validation_dir / "local-paths.log").write_text(
            f"PASS: {checked_text_files} UTF-8 text files contain no unredacted machine-local workspace paths.\n",
            encoding="utf-8",
        )
        gitleaks_log = run_gitleaks(bundle_root)
        (validation_dir / "gitleaks.log").write_text(gitleaks_log, encoding="utf-8")

        manifest = {
            "schema_version": 1,
            "kind": "haxe.go-gpt-5.6-pro-review-evidence",
            "primary": {
                "remote": source_remote_result.stdout.strip(),
                "commit": source_commit,
                "requested_ref": args.source_ref,
                "remote_branches_containing_commit": source_branches,
                "source_archive": source_archive.relative_to(bundle_root).as_posix(),
                "source_repomix": source_repomix.relative_to(bundle_root).as_posix(),
                "source_repomix_security_exclusions": source_repomix_exclusions,
                "tracked_file_count": primary_inventory["file_count"],
                "excluded_tracked_files": list(PRIMARY_EXCLUSIONS),
            },
            "github": github,
            "roadmap": roadmap,
            "references": references,
            "tooling": {
                "builder": "scripts/review/build_gpt56_evidence.py",
                "builder_sha256": sha256_file(Path(__file__)),
                "python": sys.version.split()[0],
                "repomix": source_repomix_meta,
                "reference_repomix": reference_repomix_meta,
                "gitleaks": (run(["gitleaks", "version"]).stdout or "").strip(),
            },
            "payloads": payload_records(bundle_root),
        }
        write_json(bundle_root / "MANIFEST.json", manifest)
        write_sha256sums(bundle_root)
        assert_no_machine_paths(bundle_root)
        temporary_output = output.with_name(output.name + ".tmp")
        try:
            write_deterministic_zip(bundle_root, temporary_output)
            os.replace(temporary_output, output)
        finally:
            temporary_output.unlink(missing_ok=True)

    return {
        "output": str(args.output),
        "bytes": output.stat().st_size,
        "sha256": sha256_file(output),
        "source_commit": source_commit,
        "ci_runs": sorted(set(args.ci_run)),
        "release_tag": args.release_tag,
        "reference_commits": {reference["name"]: reference["commit"] for reference in references},
    }


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv if argv is not None else sys.argv[1:])
    try:
        result = build(args)
    except EvidenceError as error:
        print(f"[review-evidence] ERROR: {error}", file=sys.stderr)
        return 1
    print(json.dumps(result, indent=2, sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
