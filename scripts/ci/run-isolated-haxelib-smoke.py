#!/usr/bin/env python3

"""Install a real package ZIP in a new Haxelib repo and run one portable app."""

from __future__ import annotations

import argparse
import json
import os
from pathlib import Path, PurePosixPath
import shutil
import subprocess
import sys
import tempfile
import zipfile


ROOT = Path(__file__).resolve().parents[2]
DEFAULT_FIXTURE = ROOT / "test" / "snapshot" / "stdlib" / "stringtools_cross_std_basic"
PACKAGE_MANIFEST = "reflaxe-package-manifest.json"


class SmokeError(RuntimeError):
    """A classified isolated-package stage failed."""

    def __init__(self, stage: str, message: str) -> None:
        super().__init__(message)
        self.stage = stage


def run_checked(
    command: list[str],
    *,
    cwd: Path,
    stage: str,
    environment: dict[str, str],
    timeout: int = 180,
) -> subprocess.CompletedProcess[str]:
    try:
        process = subprocess.run(
            command,
            cwd=cwd,
            env=environment,
            capture_output=True,
            text=True,
            timeout=timeout,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        raise SmokeError(stage, f"cannot run {' '.join(command)}: {error}") from error
    if process.returncode != 0:
        detail = (process.stdout + process.stderr).strip()
        raise SmokeError(
            stage,
            f"command failed ({' '.join(command)}):\n{detail}",
        )
    return process


def deterministic_environment(sandbox: Path) -> dict[str, str]:
    go_cache = sandbox / "go-cache"
    go_path = sandbox / "go-path"
    for directory in (go_cache, go_path):
        directory.mkdir(parents=True)
    environment = dict(os.environ)
    environment.update(
        {
            "GOCACHE": str(go_cache),
            "GOPATH": str(go_path),
            "LC_ALL": "C",
            "NO_COLOR": "1",
            "PYTHONHASHSEED": "0",
            "TZ": "UTC",
        }
    )
    return environment


def require_tool(name: str) -> str:
    executable = shutil.which(name)
    if executable is None:
        raise SmokeError("environment", f"required executable is unavailable: {name}")
    return executable


def safe_relative_path(value: object, label: str) -> str:
    if not isinstance(value, str) or not value or "\\" in value or "\0" in value:
        raise SmokeError("package", f"{label} is not a safe relative path")
    path = PurePosixPath(value)
    if (
        path.is_absolute()
        or path.as_posix() != value
        or any(part in {"", ".", ".."} for part in path.parts)
        or (len(value) >= 2 and value[0].isalpha() and value[1] == ":")
    ):
        raise SmokeError("package", f"{label} is not a safe relative path: {value!r}")
    return value


def parse_json_bytes(value: bytes, label: str) -> dict[str, object]:
    try:
        parsed = json.loads(value)
    except (json.JSONDecodeError, UnicodeDecodeError) as error:
        raise SmokeError("package", f"{label} is not valid JSON") from error
    if not isinstance(parsed, dict):
        raise SmokeError("package", f"{label} must contain a JSON object")
    return parsed


def inspect_archive(archive_path: Path) -> dict[str, object]:
    try:
        with zipfile.ZipFile(archive_path) as archive:
            infos = archive.infolist()
            names = [info.filename for info in infos]
            if any(info.is_dir() for info in infos):
                raise SmokeError("package", "package ZIP must not contain directory entries")
            for name in names:
                safe_relative_path(name, "package member")
            if len(names) != len(set(names)):
                raise SmokeError("package", "package ZIP contains duplicate members")
            haxelib = parse_json_bytes(archive.read("haxelib.json"), "haxelib.json")
            manifest = parse_json_bytes(
                archive.read(PACKAGE_MANIFEST),
                PACKAGE_MANIFEST,
            )
    except (OSError, KeyError, zipfile.BadZipFile) as error:
        raise SmokeError("package", f"cannot inspect package ZIP: {error}") from error

    if haxelib.get("name") != "reflaxe.go":
        raise SmokeError("package", "package haxelib.json must name reflaxe.go")
    version = haxelib.get("version")
    if not isinstance(version, str) or not version:
        raise SmokeError("package", "package haxelib.json must declare a version")
    if haxelib.get("classPath") != "src" or "reflaxe" in haxelib:
        raise SmokeError(
            "package",
            "installed package metadata must use classPath=src without source-only reflaxe fields",
        )
    if (
        manifest.get("schemaVersion") != 1
        or manifest.get("format") != "reflaxe.go-haxelib-package"
        or manifest.get("classPath") != "src"
    ):
        raise SmokeError("package", "embedded package manifest header is not canonical")
    entries = manifest.get("entries")
    if not isinstance(entries, list):
        raise SmokeError("package", "embedded package manifest entries must be an array")

    package_paths: list[str] = []
    override_count = 0
    for index, entry in enumerate(entries):
        if not isinstance(entry, dict):
            raise SmokeError("package", f"embedded package manifest entry {index} is invalid")
        source_path = safe_relative_path(entry.get("sourcePath"), "manifest source path")
        package_path = safe_relative_path(entry.get("packagePath"), "manifest package path")
        package_paths.append(package_path)
        if entry.get("kind") == "stdlib-override":
            override_count += 1
            if (
                not source_path.startswith("std/go/_std/")
                or not package_path.startswith("src/")
                or not package_path.endswith(".cross.hx")
            ):
                raise SmokeError(
                    "package",
                    f"stdlib override mapping is not canonical: {source_path} -> {package_path}",
                )
        elif package_path.endswith(".cross.hx"):
            raise SmokeError(
                "package",
                f"package .cross.hx lacks stdlib-override ownership: {package_path}",
            )

    if package_paths != sorted(package_paths, key=lambda item: item.encode("utf-8")):
        raise SmokeError("package", "embedded package manifest entries are not sorted")
    if len(package_paths) != len(set(package_paths)):
        raise SmokeError("package", "embedded package manifest paths are not unique")
    if set(package_paths) != set(names) - {PACKAGE_MANIFEST}:
        raise SmokeError("package", "embedded package manifest does not cover the ZIP exactly")

    return {
        "archiveMembers": len(names),
        "manifestEntries": len(entries),
        "stdlibOverrides": override_count,
        "version": version,
    }


def build_archive(
    source_root: Path,
    workspace: Path,
    environment: dict[str, str],
) -> Path:
    package_script = source_root / "scripts" / "release" / "package-haxelib.sh"
    if not package_script.is_file():
        raise SmokeError(
            "package",
            "source root lacks scripts/release/package-haxelib.sh",
        )
    archive = workspace / "reflaxe.go.zip"
    run_checked(
        ["bash", str(package_script), str(archive), str(source_root)],
        cwd=source_root,
        stage="package",
        environment=environment,
    )
    return archive


def fixture_name(fixture_root: Path) -> str:
    snapshot_root = ROOT / "test" / "snapshot"
    try:
        return fixture_root.resolve().relative_to(snapshot_root.resolve()).as_posix()
    except ValueError:
        return fixture_root.name


def find_path_leaks(root: Path, needles: list[Path]) -> list[str]:
    encoded_needles = [str(path.resolve()).encode("utf-8") for path in needles]
    leaks: list[str] = []
    for path in sorted(item for item in root.rglob("*") if item.is_file()):
        content = path.read_bytes()
        if any(needle in content for needle in encoded_needles):
            leaks.append(path.relative_to(root).as_posix())
    return leaks


def run_installed_smoke(
    *,
    archive: Path,
    source_root: Path,
    fixture_root: Path,
    sandbox: Path,
    environment: dict[str, str],
) -> dict[str, object]:
    haxelib = require_tool("haxelib")
    haxe = require_tool("haxe")
    go = require_tool("go")

    run_checked(
        [haxelib, "newrepo"],
        cwd=sandbox,
        stage="install",
        environment=environment,
    )
    run_checked(
        [haxelib, "install", str(archive), "--always", "--skip-dependencies"],
        cwd=sandbox,
        stage="install",
        environment=environment,
    )
    installed_path = run_checked(
        [haxelib, "path", "reflaxe.go"],
        cwd=sandbox,
        stage="install",
        environment=environment,
    ).stdout
    if str(source_root.resolve()) in installed_path:
        raise SmokeError(
            "install",
            "haxelib path resolved the source checkout instead of the isolated repository",
        )
    if str((sandbox / ".haxelib").resolve()) not in installed_path:
        raise SmokeError("install", "haxelib path did not resolve the isolated repository")

    main_source = fixture_root / "Main.hx"
    expected_stdout_path = fixture_root / "expected.stdout"
    if not main_source.is_file() or not expected_stdout_path.is_file():
        raise SmokeError(
            "fixture",
            "isolated smoke fixture must contain Main.hx and expected.stdout",
        )
    app = sandbox / "app"
    app.mkdir()
    shutil.copyfile(main_source, app / "Main.hx")
    output = app / "out"
    run_checked(
        [
            haxe,
            "-cp",
            ".",
            "-lib",
            "reflaxe.go",
            "-D",
            f"go_output={output}",
            "-D",
            "reflaxe_go_profile=portable",
            "-D",
            "reflaxe_go_strict_examples",
            "-D",
            "reflaxe.dont_output_metadata_id",
            "-D",
            "go_no_build",
            "-D",
            "no-traces",
            "-D",
            "no_traces",
            "-main",
            "Main",
        ],
        cwd=app,
        stage="compile",
        environment=environment,
    )
    run_checked(
        [go, "test", "./..."],
        cwd=output,
        stage="go-test",
        environment=environment,
    )
    actual_stdout = run_checked(
        [go, "run", "."],
        cwd=output,
        stage="run",
        environment=environment,
    ).stdout
    expected_stdout = expected_stdout_path.read_text(encoding="utf-8")
    if actual_stdout != expected_stdout:
        raise SmokeError(
            "stdout",
            f"generated app output differs: expected {expected_stdout!r}, got {actual_stdout!r}",
        )

    leaks = find_path_leaks(output, [source_root, sandbox])
    if leaks:
        raise SmokeError(
            "path-scan",
            "generated output contains checkout or sandbox paths: " + ", ".join(leaks),
        )

    return {
        "checks": {
            "checkoutClasspathsAbsent": True,
            "generatedPathLeaksAbsent": True,
            "goRun": "pass",
            "goTest": "pass",
            "haxeCompile": "pass",
            "haxelibInstall": "pass",
            "stdout": "pass",
        },
        "fixture": {
            "name": fixture_name(fixture_root),
            "profile": "portable",
        },
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--archive", type=Path)
    parser.add_argument("--source-root", type=Path, default=ROOT)
    parser.add_argument("--fixture-root", type=Path, default=DEFAULT_FIXTURE)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    source_root = args.source_root.resolve()
    fixture_root = args.fixture_root.resolve()
    try:
        with tempfile.TemporaryDirectory(prefix="haxe-go-isolated-haxelib-") as raw:
            workspace = Path(raw)
            sandbox = workspace / "sandbox"
            sandbox.mkdir()
            environment = deterministic_environment(sandbox)
            if args.archive is None:
                require_tool("haxe")
                archive = build_archive(source_root, workspace, environment)
            else:
                archive = args.archive.resolve()
            package_evidence = inspect_archive(archive)
            smoke_evidence = run_installed_smoke(
                archive=archive,
                source_root=source_root,
                fixture_root=fixture_root,
                sandbox=sandbox,
                environment=environment,
            )
    except SmokeError as error:
        print(
            f"[isolated-haxelib-smoke] ERROR [{error.stage}]: {error}",
            file=sys.stderr,
        )
        return 2
    except OSError as error:
        print(
            f"[isolated-haxelib-smoke] ERROR [environment]: {error}",
            file=sys.stderr,
        )
        return 2

    evidence = {
        "schemaVersion": 1,
        "kind": "haxe.go-canonical-std-isolated-package-smoke",
        "package": package_evidence,
        **smoke_evidence,
    }
    print(json.dumps(evidence, indent=2, sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
