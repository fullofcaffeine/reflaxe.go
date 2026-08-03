#!/usr/bin/env python3
"""Build exact, scope-bounded GitHub release notes from compatibility policy."""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parents[2]
DEFAULT_MANIFEST = ROOT / "docs" / "compatibility-support-manifest.json"
STABLE_VERSION = re.compile(r"^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$")
FULL_SHA = re.compile(r"^[0-9a-f]{40}$")
REPOSITORY = re.compile(r"^[A-Za-z0-9_.-]+/[A-Za-z0-9_.-]+$")


class NotesError(ValueError):
    """A release-notes input or compatibility-policy error."""


def load_manifest(path: Path) -> dict[str, Any]:
    try:
        document = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as error:
        raise NotesError(f"cannot read compatibility manifest {path}: {error}") from error
    if document.get("kind") != "haxe.go-compatibility-support-manifest":
        raise NotesError("compatibility manifest has the wrong kind")
    return document


def require_mapping(document: dict[str, Any], key: str) -> dict[str, Any]:
    value = document.get(key)
    if not isinstance(value, dict):
        raise NotesError(f"compatibility manifest {key} must be an object")
    return value


def require_list(document: dict[str, Any], key: str) -> list[Any]:
    value = document.get(key)
    if not isinstance(value, list):
        raise NotesError(f"compatibility manifest {key} must be an array")
    return value


def require_text(document: dict[str, Any], key: str, owner: str) -> str:
    value = document.get(key)
    if not isinstance(value, str) or not value.strip():
        raise NotesError(f"{owner}.{key} must be non-empty text")
    return value.strip()


def natural_join(values: list[str]) -> str:
    if not values:
        return ""
    if len(values) == 1:
        return values[0]
    if len(values) == 2:
        return f"{values[0]} and {values[1]}"
    return f"{', '.join(values[:-1])}, and {values[-1]}"


def render_notes(
    *,
    manifest: dict[str, Any],
    version: str,
    tag: str,
    source_sha: str,
    repository: str,
) -> str:
    claim = require_mapping(manifest, "release_claim")
    statement = require_text(claim, "statement", "release_claim")
    default_disposition = require_text(
        claim, "default_disposition", "release_claim"
    )
    admitted_preset = require_text(claim, "admitted_preset", "release_claim")
    admitted_platform = require_text(claim, "admitted_platform", "release_claim")

    presets = require_list(manifest, "presets")
    platforms = require_list(manifest, "platforms")
    surfaces = require_list(manifest, "surfaces")
    toolchains = require_mapping(manifest, "toolchains")
    trust_assumptions = require_list(manifest, "trust_assumptions")

    platform = next(
        (
            entry
            for entry in platforms
            if isinstance(entry, dict) and entry.get("id") == admitted_platform
        ),
        None,
    )
    if not isinstance(platform, dict) or platform.get("release_admitted") is not True:
        raise NotesError("admitted platform is missing or not release-admitted")
    platform_label = (
        f"{require_text(platform, 'os', 'admitted platform').title()}/"
        f"{require_text(platform, 'architecture', 'admitted platform')}"
    )

    admitted_symbols = 0
    excluded_surfaces: list[tuple[str, list[str]]] = []
    for surface in surfaces:
        if not isinstance(surface, dict):
            raise NotesError("compatibility surface entries must be objects")
        surface_id = require_text(surface, "id", "surface")
        operations = require_list(surface, "operations")
        excluded_symbols: list[str] = []
        for operation in operations:
            if not isinstance(operation, dict):
                raise NotesError(f"surface {surface_id} operation must be an object")
            symbols = operation.get("symbols")
            if not isinstance(symbols, list) or not all(
                isinstance(symbol, str) and symbol for symbol in symbols
            ):
                raise NotesError(f"surface {surface_id} operation symbols are invalid")
            if operation.get("release_admitted") is True:
                admitted_symbols += len(symbols)
            elif surface_id != "distribution":
                excluded_symbols.extend(symbols)
        if excluded_symbols:
            excluded_surfaces.append((surface_id, excluded_symbols))

    haxe = require_mapping(toolchains, "haxe")
    go = require_mapping(toolchains, "go")
    node = require_mapping(toolchains, "node")
    haxe_versions = haxe.get("supported_versions")
    go_versions = go.get("ci_versions")
    node_lines = node.get("supported_tooling_lines")
    if not isinstance(haxe_versions, list) or not all(
        isinstance(value, str) for value in haxe_versions
    ):
        raise NotesError("toolchains.haxe.supported_versions is invalid")
    if not isinstance(go_versions, list) or not all(
        isinstance(value, str) for value in go_versions
    ):
        raise NotesError("toolchains.go.ci_versions is invalid")
    if not isinstance(node_lines, list) or not all(
        isinstance(value, str) for value in node_lines
    ):
        raise NotesError("toolchains.node.supported_tooling_lines is invalid")

    lines = [
        f"# Haxe.Go {version}",
        "",
        statement,
        "",
        f"This release is tag `{tag}` at exact source commit `{source_sha}`.",
        "",
        "## What this beta admits",
        "",
        f"- Policy preset: `{admitted_preset}`.",
        f"- Runtime platform: {platform_label}.",
        f"- Haxe {natural_join(haxe_versions)}.",
        f"- Go {natural_join(go_versions)}.",
        f"- Node {natural_join(node_lines)} for build and release tooling.",
        f"- Exactly {admitted_symbols} named operations or members from the governed support manifest.",
        "",
        "The admission assumes:",
        "",
    ]
    for assumption in trust_assumptions:
        if not isinstance(assumption, dict) or assumption.get("release_required") is not True:
            continue
        lines.append(f"- {require_text(assumption, 'statement', 'trust assumption')}")

    lines.extend(
        [
            "",
            "## Important exclusions",
            "",
            f"- {default_disposition}",
        ]
    )
    for preset in presets:
        if not isinstance(preset, dict) or preset.get("release_admitted") is True:
            continue
        preset_id = require_text(preset, "id", "preset")
        qualification = require_text(preset, "qualification", f"preset {preset_id}")
        lines.append(f"- `{preset_id}` preset: {qualification}")

    excluded_platforms = [
        require_text(entry, "id", "platform")
        for entry in platforms
        if isinstance(entry, dict) and entry.get("release_admitted") is not True
    ]
    if excluded_platforms:
        lines.append(
            "- Other platforms and architectures are not runtime-admitted: "
            f"{', '.join(f'`{value}`' for value in excluded_platforms)}."
        )
    for surface_id, symbols in excluded_surfaces:
        lines.append(
            f"- `{surface_id}` remains outside this beta for: "
            f"{'; '.join(symbols)}."
        )
    lines.extend(
        [
            "- This is not a stable 1.x release.",
            "",
            "## Verified release assets",
            "",
            "The GitHub release contains a deterministic Haxelib ZIP, its SHA-256 sidecar, "
            "a complete content manifest, and an in-toto/SLSA provenance statement. "
            "The release workflow verifies their local and hosted digests before publication.",
            "",
            "## Scope authority",
            "",
            f"- [Compatibility support manifest](https://github.com/{repository}/blob/{tag}/docs/compatibility-support-manifest.json)",
            f"- [Human-readable support matrix](https://github.com/{repository}/blob/{tag}/docs/compatibility-support-matrix.md)",
            f"- [Known gaps and production caveats](https://github.com/{repository}/blob/{tag}/docs/known-gaps.md)",
            f"- [Exact source](https://github.com/{repository}/commit/{source_sha})",
            "",
        ]
    )
    return "\n".join(lines)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--version", required=True)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--source-sha", required=True)
    parser.add_argument("--repository", required=True)
    parser.add_argument("--manifest", type=Path, default=DEFAULT_MANIFEST)
    parser.add_argument("--output", type=Path, required=True)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    if not STABLE_VERSION.fullmatch(args.version):
        raise NotesError("version must be canonical stable SemVer")
    if args.tag != f"v{args.version}":
        raise NotesError("tag must equal v<version>")
    if not FULL_SHA.fullmatch(args.source_sha):
        raise NotesError("source SHA must be 40 lowercase hexadecimal characters")
    if not REPOSITORY.fullmatch(args.repository):
        raise NotesError("repository must use OWNER/NAME form")
    if args.output.exists():
        raise NotesError(f"output already exists: {args.output}")
    manifest = load_manifest(args.manifest)
    notes = render_notes(
        manifest=manifest,
        version=args.version,
        tag=args.tag,
        source_sha=args.source_sha,
        repository=args.repository,
    )
    args.output.parent.mkdir(parents=True, exist_ok=True)
    args.output.write_text(notes, encoding="utf-8")
    print(f"[release-notes] wrote bounded notes: {args.output}")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except NotesError as error:
        raise SystemExit(f"[release-notes] error: {error}") from error
