#!/usr/bin/env python3

"""Deterministic Git change discovery shared by affected-test planners."""

from __future__ import annotations

from pathlib import Path
import subprocess


class GitChangeDiscoveryError(RuntimeError):
    """Raised when Git cannot establish a trustworthy changed-path inventory."""


def _run_git(root: Path, args: list[str]) -> bytes:
    proc = subprocess.run(
        ["git", *args],
        cwd=root,
        capture_output=True,
        check=True,
    )
    return proc.stdout


def _parse_name_status_z(raw: bytes) -> set[str]:
    tokens = [token.decode("utf-8", errors="surrogateescape") for token in raw.split(b"\0") if token]
    paths: set[str] = set()
    index = 0
    while index < len(tokens):
        status = tokens[index]
        index += 1
        if index >= len(tokens):
            break
        paths.add(tokens[index])
        index += 1
        if status.startswith(("R", "C")) and index < len(tokens):
            paths.add(tokens[index])
            index += 1
    return paths


def collect_changed_paths(root: Path, pathspecs: list[str] | None = None, base: str = "") -> set[str]:
    """Return committed, worktree, index, rename, delete, and untracked paths.

    The result is deliberately a union. A selector must never silently lose a
    staged fixture merely because an unstaged diff also exists.
    """

    suffix = ["--", *(pathspecs or [])]
    paths: set[str] = set()
    try:
        if base:
            paths.update(
                _parse_name_status_z(
                    _run_git(root, ["diff", "--name-status", "-z", "--find-renames", f"{base}...HEAD", *suffix])
                )
            )
        paths.update(
            _parse_name_status_z(
                _run_git(root, ["diff", "--name-status", "-z", "--find-renames", *suffix])
            )
        )
        paths.update(
            _parse_name_status_z(
                _run_git(root, ["diff", "--cached", "--name-status", "-z", "--find-renames", *suffix])
            )
        )
        untracked = _run_git(root, ["ls-files", "--others", "--exclude-standard", "-z", *suffix])
        paths.update(
            token.decode("utf-8", errors="surrogateescape")
            for token in untracked.split(b"\0")
            if token
        )
    except (FileNotFoundError, subprocess.CalledProcessError) as error:
        raise GitChangeDiscoveryError(f"Git change discovery failed: {error}") from error
    return paths
