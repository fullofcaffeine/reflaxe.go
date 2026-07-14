#!/usr/bin/env python3

"""Validate canonical Reflaxe source and packaged standard-library layouts.

Source checkouts own target overrides as ordinary ``.hx`` files under
``std/go/_std``. Reflaxe package staging flattens that root into ``src`` and
renames those files to ``.cross.hx``. Keeping both shapes explicit prevents a
checkout-only bootstrap from hiding a broken installed package.
"""

from __future__ import annotations

import argparse
from dataclasses import dataclass
import hashlib
import json
from pathlib import Path, PurePosixPath
import re
from typing import Iterable


EXPECTED_STD_PATHS = ["std", "std/go/_std"]
EXPECTED_SOURCE_CLASS_PATHS = [
    "${SCOPE_DIR}/src",
    "${SCOPE_DIR}/std",
    "${SCOPE_DIR}/std/go/_std",
]
LEGACY_SUPPORT_CLASS_PATH = "${SCOPE_DIR}/std/_std"
EXPECTED_REFLAXE_CLASS_PATH = "${SCOPE_DIR}/vendor/reflaxe/src"
PACKAGE_MAP_MANIFEST = "reflaxe-package-manifest.json"
EXPECTED_ARCHIVE_POLICY = {
    "compression": "stored",
    "fileMode": "0644",
    "ordering": "utf8-bytewise",
    "timestamp": "2000-01-01T00:00:00Z",
}
EXPECTED_PACKAGE_ENTRY_KINDS = {
    "class-path",
    "metadata",
    "package-runner",
    "runtime",
    "stdlib",
    "stdlib-override",
    "vendored-reflaxe",
}
SHA256 = re.compile(r"^[0-9a-f]{64}$")
TEXT_SUFFIXES = {
    ".go",
    ".hxml",
    ".hx",
    ".json",
    ".md",
    ".mjs",
    ".py",
    ".sh",
    ".txt",
}
POSIX_HOME_PATH = re.compile(r"(?<![A-Za-z0-9])/(?:Users|home)/[^\s\"'`]+")
WINDOWS_HOME_PATH = re.compile(r"(?i)(?<![A-Za-z0-9])[A-Z]:[\\/]+Users[\\/]+[^\s\"'`]+")


@dataclass(frozen=True, order=True)
class Violation:
    code: str
    path: str
    message: str

    def render(self) -> str:
        location = f" [{self.path}]" if self.path else ""
        return f"{self.code}{location}: {self.message}"


def relative_display(path: Path, root: Path) -> str:
    try:
        return path.relative_to(root).as_posix()
    except ValueError:
        return path.name


def normalized_hxml_class_paths(path: Path) -> list[str]:
    if not path.is_file():
        return []

    class_paths: list[str] = []
    for raw_line in path.read_text(encoding="utf-8").splitlines():
        line = raw_line.split("#", 1)[0].strip()
        if not line:
            continue
        if line.startswith("-cp "):
            value = line[4:].strip()
        elif line.startswith("-p "):
            value = line[3:].strip()
        elif line.startswith("--class-path "):
            value = line[len("--class-path ") :].strip()
        else:
            continue
        class_paths.append(value.rstrip("/\\"))
    return class_paths


def hxml_libraries(path: Path) -> list[str]:
    if not path.is_file():
        return []

    libraries: list[str] = []
    for raw_line in path.read_text(encoding="utf-8").splitlines():
        line = raw_line.split("#", 1)[0].strip()
        if line.startswith("-lib "):
            libraries.append(line[5:].strip())
        elif line.startswith("--library "):
            libraries.append(line[len("--library ") :].strip())
    return libraries


def has_ordered_subsequence(values: list[str], expected: list[str]) -> bool:
    cursor = 0
    for value in values:
        if cursor < len(expected) and value == expected[cursor]:
            cursor += 1
    return cursor == len(expected)


def read_json_object(path: Path) -> dict[str, object] | None:
    if not path.is_file():
        return None
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError):
        return None
    return value if isinstance(value, dict) else None


def sha256(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as handle:
        for chunk in iter(lambda: handle.read(1024 * 1024), b""):
            digest.update(chunk)
    return digest.hexdigest()


def safe_manifest_path(value: object) -> str | None:
    if not isinstance(value, str) or not value or "\\" in value or "\0" in value:
        return None
    path = PurePosixPath(value)
    if (
        path.is_absolute()
        or path.as_posix() != value
        or any(segment in {"", ".", ".."} for segment in path.parts)
        or (len(value) >= 2 and value[0].isalpha() and value[1] == ":")
    ):
        return None
    return value


def text_files(root: Path) -> Iterable[Path]:
    if not root.exists():
        return []
    return (
        path
        for path in sorted(root.rglob("*"))
        if path.is_file() and path.suffix.lower() in TEXT_SUFFIXES
    )


def find_absolute_path_leaks(root: Path, forbidden_roots: Iterable[Path]) -> list[str]:
    exact_needles = {
        str(path.resolve())
        for path in forbidden_roots
        if str(path) not in {"", "."} and path.exists()
    }
    leaks: list[str] = []
    for path in text_files(root):
        relative = relative_display(path, root)
        content = path.read_text(encoding="utf-8", errors="replace")
        if any(needle and needle in content for needle in exact_needles):
            leaks.append(relative)
            continue
        # Vendored Reflaxe documents generic absolute-path examples as part of
        # its path-normalization implementation. Exact checkout/build roots are
        # still rejected there; only generic examples receive this exemption.
        if not relative.startswith("vendor/") and (
            POSIX_HOME_PATH.search(content) or WINDOWS_HOME_PATH.search(content)
        ):
            leaks.append(relative)
    return leaks


def audit_source_layout(root: Path) -> list[Violation]:
    root = root.resolve()
    violations: list[Violation] = []
    manifest_path = root / "haxelib.json"
    manifest = read_json_object(manifest_path)
    reflaxe = manifest.get("reflaxe") if manifest else None
    std_paths = reflaxe.get("stdPaths") if isinstance(reflaxe, dict) else None
    if std_paths != EXPECTED_STD_PATHS:
        violations.append(
            Violation(
                "source-std-paths",
                "haxelib.json",
                f"reflaxe.stdPaths must be {EXPECTED_STD_PATHS!r}; found {std_paths!r}",
            )
        )

    canonical_root = root / "std" / "go" / "_std"
    canonical_haxe = sorted(canonical_root.rglob("*.hx")) if canonical_root.is_dir() else []
    if not canonical_haxe:
        violations.append(
            Violation(
                "source-canonical-root-missing",
                "std/go/_std",
                "canonical source override root must contain ordinary .hx files",
            )
        )

    cross_files = sorted(
        path
        for base in (root / "src", root / "std")
        if base.exists()
        for path in base.rglob("*.cross.hx")
        if path.is_file()
    )
    if cross_files:
        sample = ", ".join(relative_display(path, root) for path in cross_files[:5])
        suffix = f", ... ({len(cross_files) - 5} more)" if len(cross_files) > 5 else ""
        violations.append(
            Violation(
                "source-cross-files",
                "std",
                f"source authority must not contain .cross.hx files; found {len(cross_files)}: {sample}{suffix}",
            )
        )

    legacy_support_root = root / "std" / "_std"
    legacy_support_sources = (
        sorted(path for path in legacy_support_root.rglob("*.hx") if path.is_file())
        if legacy_support_root.is_dir()
        else []
    )
    if legacy_support_sources:
        sample = ", ".join(
            relative_display(path, root) for path in legacy_support_sources[:5]
        )
        suffix = (
            f", ... ({len(legacy_support_sources) - 5} more)"
            if len(legacy_support_sources) > 5
            else ""
        )
        violations.append(
            Violation(
                "source-legacy-support-root",
                "std/_std",
                "target support must live in ordinary std or std/hxrt modules; "
                f"found {len(legacy_support_sources)} legacy source file(s): {sample}{suffix}",
            )
        )

    hxml_path = root / "haxe_libraries" / "reflaxe.go.hxml"
    class_paths = normalized_hxml_class_paths(hxml_path)
    if not has_ordered_subsequence(class_paths, EXPECTED_SOURCE_CLASS_PATHS):
        violations.append(
            Violation(
                "source-classpath-precedence",
                "haxe_libraries/reflaxe.go.hxml",
                "initial classpaths must declare ${SCOPE_DIR}/src, then std, then std/go/_std "
                f"so the target override has effective precedence; found {class_paths!r}",
            )
        )
    if LEGACY_SUPPORT_CLASS_PATH in class_paths:
        violations.append(
            Violation(
                "source-legacy-support-classpath",
                "haxe_libraries/reflaxe.go.hxml",
                "initial classpaths must not declare the retired ${SCOPE_DIR}/std/_std support root",
            )
        )

    reflaxe_hxml_path = root / "haxe_libraries" / "reflaxe.hxml"
    reflaxe_class_paths = normalized_hxml_class_paths(reflaxe_hxml_path)
    if (
        "reflaxe" not in hxml_libraries(hxml_path)
        or EXPECTED_REFLAXE_CLASS_PATH not in reflaxe_class_paths
    ):
        violations.append(
            Violation(
                "source-vendored-reflaxe-pretyping",
                "haxe_libraries/reflaxe.go.hxml, haxe_libraries/reflaxe.hxml",
                "source builds must load vendored Reflaxe through an explicit initial library "
                f"classpath ({EXPECTED_REFLAXE_CLASS_PATH})",
            )
        )

    config_paths = [manifest_path, root / "extraParams.hxml", hxml_path, reflaxe_hxml_path]
    leaked_configs: list[str] = []
    for path in config_paths:
        if not path.is_file():
            continue
        content = path.read_text(encoding="utf-8", errors="replace")
        if str(root) in content or POSIX_HOME_PATH.search(content) or WINDOWS_HOME_PATH.search(content):
            leaked_configs.append(relative_display(path, root))
    leaked_configs.extend(
        f"std/go/_std/{path}"
        for path in find_absolute_path_leaks(canonical_root, [root])
    )
    if leaked_configs:
        violations.append(
            Violation(
                "absolute-path-leak",
                ", ".join(leaked_configs),
                "source configuration contains an absolute machine-local path",
            )
        )

    return sorted(violations)


def expected_packaged_cross_files(source_root: Path) -> set[str]:
    canonical_root = source_root / "std" / "go" / "_std"
    if not canonical_root.is_dir():
        return set()
    expected: set[str] = set()
    for source in canonical_root.rglob("*.hx"):
        relative = source.relative_to(canonical_root)
        expected.add((Path("src") / relative.with_suffix(".cross.hx")).as_posix())
    return expected


def expected_packaged_ordinary_files(source_root: Path) -> set[str]:
    std_root = source_root / "std"
    canonical_root = std_root / "go" / "_std"
    if not std_root.is_dir():
        return set()

    expected: set[str] = set()
    for source in std_root.rglob("*.hx"):
        if source.is_relative_to(canonical_root):
            continue
        relative = source.relative_to(std_root)
        expected.add((Path("src") / relative).as_posix())
    return expected


def audit_package_map_manifest(package_root: Path, source_root: Path) -> list[Violation]:
    manifest_path = package_root / PACKAGE_MAP_MANIFEST
    manifest = read_json_object(manifest_path)
    if manifest is None:
        return [
            Violation(
                "package-map-manifest",
                PACKAGE_MAP_MANIFEST,
                "package must contain a valid deterministic source-to-package manifest",
            )
        ]

    problems: list[str] = []
    if (
        manifest.get("schemaVersion") != 1
        or manifest.get("format") != "reflaxe.go-haxelib-package"
        or manifest.get("classPath") != "src"
        or manifest.get("archive") != EXPECTED_ARCHIVE_POLICY
    ):
        problems.append("header or archive policy is not canonical")

    raw_entries = manifest.get("entries")
    if not isinstance(raw_entries, list):
        problems.append("entries must be an array")
        raw_entries = []

    package_paths: list[str] = []
    entries_by_source: dict[str, dict[str, object]] = {}
    for index, raw_entry in enumerate(raw_entries):
        if not isinstance(raw_entry, dict):
            problems.append(f"entry {index} is not an object")
            continue
        source_path = safe_manifest_path(raw_entry.get("sourcePath"))
        package_path = safe_manifest_path(raw_entry.get("packagePath"))
        kind = raw_entry.get("kind")
        source_digest = raw_entry.get("sourceSha256")
        package_digest = raw_entry.get("packageSha256")
        size = raw_entry.get("size")
        if source_path is None or package_path is None:
            problems.append(f"entry {index} contains an unsafe path")
            continue
        package_paths.append(package_path)
        if source_path in entries_by_source:
            problems.append(f"duplicate source mapping: {source_path}")
        entries_by_source[source_path] = raw_entry
        if kind not in EXPECTED_PACKAGE_ENTRY_KINDS:
            problems.append(f"entry {package_path} has an unknown kind")
        if not isinstance(source_digest, str) or not SHA256.fullmatch(source_digest):
            problems.append(f"entry {package_path} has an invalid source hash")
        if not isinstance(package_digest, str) or not SHA256.fullmatch(package_digest):
            problems.append(f"entry {package_path} has an invalid package hash")
        if not isinstance(size, int) or isinstance(size, bool) or size < 0:
            problems.append(f"entry {package_path} has an invalid size")

        source_file = source_root.joinpath(*PurePosixPath(source_path).parts)
        package_file = package_root.joinpath(*PurePosixPath(package_path).parts)
        if not source_file.is_file():
            problems.append(f"entry source is missing: {source_path}")
        elif isinstance(source_digest, str) and sha256(source_file) != source_digest:
            problems.append(f"entry source hash differs: {source_path}")
        if not package_file.is_file():
            problems.append(f"entry package file is missing: {package_path}")
        else:
            if isinstance(package_digest, str) and sha256(package_file) != package_digest:
                problems.append(f"entry package hash differs: {package_path}")
            if isinstance(size, int) and not isinstance(size, bool) and package_file.stat().st_size != size:
                problems.append(f"entry package size differs: {package_path}")

    expected_order = sorted(package_paths, key=lambda value: value.encode("utf-8"))
    if package_paths != expected_order:
        problems.append("entries are not sorted by UTF-8 package path")
    if len(package_paths) != len(set(package_paths)):
        problems.append("package paths are not unique")

    actual_package_files = {
        relative_display(path, package_root)
        for path in package_root.rglob("*")
        if path.is_file() and path != manifest_path
    }
    if set(package_paths) != actual_package_files:
        missing = sorted(actual_package_files - set(package_paths))
        stale = sorted(set(package_paths) - actual_package_files)
        problems.append(f"file coverage differs (unmapped={missing!r}, stale={stale!r})")

    expected_source_mappings: dict[str, tuple[str, str]] = {}
    class_path_root = source_root / "src"
    if class_path_root.is_dir():
        for source in sorted(class_path_root.rglob("*.hx")):
            source_path = relative_display(source, source_root)
            expected_source_mappings[source_path] = (source_path, "class-path")
    canonical_root = source_root / "std" / "go" / "_std"
    if canonical_root.is_dir():
        for source in sorted(canonical_root.rglob("*.hx")):
            source_path = relative_display(source, source_root)
            package_path = (
                Path("src") / source.relative_to(canonical_root).with_suffix(".cross.hx")
            ).as_posix()
            expected_source_mappings[source_path] = (package_path, "stdlib-override")
    std_root = source_root / "std"
    if std_root.is_dir():
        for source in sorted(std_root.rglob("*.hx")):
            if source.is_relative_to(canonical_root):
                continue
            source_path = relative_display(source, source_root)
            package_path = (Path("src") / source.relative_to(std_root)).as_posix()
            expected_source_mappings[source_path] = (package_path, "stdlib")
    runtime_root = source_root / "runtime"
    if runtime_root.is_dir():
        for source in sorted(runtime_root.rglob("*.go")):
            if source.name.endswith("_test.go"):
                continue
            source_path = relative_display(source, source_root)
            expected_source_mappings[source_path] = (source_path, "runtime")
    vendored_source_root = source_root / "vendor" / "reflaxe" / "src"
    if vendored_source_root.is_dir():
        for source in sorted(vendored_source_root.rglob("*.hx")):
            source_path = relative_display(source, source_root)
            expected_source_mappings[source_path] = (source_path, "vendored-reflaxe")
    for source_path in (
        "vendor/reflaxe/FUTURE_MODIFICATIONS.md",
        "vendor/reflaxe/LICENSE",
        "vendor/reflaxe/PATCHES.md",
        "vendor/reflaxe/haxelib.json",
    ):
        if (source_root / source_path).is_file():
            expected_source_mappings[source_path] = (source_path, "vendored-reflaxe")
    for source_path, kind in (
        ("LICENSE", "metadata"),
        ("LICENSING.md", "metadata"),
        ("README.md", "metadata"),
        ("Run.hx", "package-runner"),
        ("extraParams.hxml", "metadata"),
        ("haxelib.json", "metadata"),
        ("license-policy.json", "metadata"),
        ("licenses/HAXE-STDLIB-MIT.txt", "metadata"),
    ):
        if (source_root / source_path).is_file():
            expected_source_mappings[source_path] = (source_path, kind)
    for source_path, (package_path, kind) in expected_source_mappings.items():
        entry = entries_by_source.get(source_path)
        if entry is None:
            problems.append(f"manifest omits declared source mapping: {source_path}")
        elif entry.get("packagePath") != package_path or entry.get("kind") != kind:
            problems.append(f"manifest misclassifies declared source mapping: {source_path}")
    unexpected_sources = sorted(set(entries_by_source) - set(expected_source_mappings))
    if unexpected_sources:
        problems.append(f"manifest includes undeclared source mappings: {unexpected_sources!r}")

    if not problems:
        return []
    displayed = "; ".join(problems[:6])
    if len(problems) > 6:
        displayed += f"; ... ({len(problems) - 6} more)"
    return [Violation("package-map-manifest", PACKAGE_MAP_MANIFEST, displayed)]


def audit_package_layout(package_root: Path, source_root: Path) -> list[Violation]:
    package_root = package_root.resolve()
    source_root = source_root.resolve()
    violations: list[Violation] = []
    manifest = read_json_object(package_root / "haxelib.json")
    if manifest is None or manifest.get("classPath") != "src" or "reflaxe" in manifest:
        violations.append(
            Violation(
                "package-manifest",
                "haxelib.json",
                "packaged manifest must use classPath=src and omit source-only reflaxe metadata",
            )
        )

    if (package_root / "std").exists() or (package_root / "src" / "go" / "_std").exists():
        violations.append(
            Violation(
                "package-unflattened-std",
                "std",
                "package must flatten std paths into src and must not retain std/go/_std",
            )
        )

    expected = expected_packaged_cross_files(source_root)
    actual = {
        relative_display(path, package_root)
        for path in (package_root / "src").rglob("*.cross.hx")
        if path.is_file()
    } if (package_root / "src").is_dir() else set()
    plain_counterparts = {
        path.removesuffix(".cross.hx") + ".hx"
        for path in expected
        if (package_root / (path.removesuffix(".cross.hx") + ".hx")).exists()
    }
    if actual != expected or plain_counterparts:
        missing = sorted(expected - actual)
        unexpected = sorted(actual - expected)
        violations.append(
            Violation(
                "package-cross-mapping",
                "src",
                "packaged .cross.hx set must exactly map ordinary std/go/_std sources; "
                f"missing={missing!r}, unexpected={unexpected!r}, plain={sorted(plain_counterparts)!r}",
            )
        )

    expected_ordinary = expected_packaged_ordinary_files(source_root)
    actual_ordinary = {
        relative_display(path, package_root)
        for path in (package_root / "src").rglob("*.hx")
        if path.is_file() and not path.name.endswith(".cross.hx")
    } if (package_root / "src").is_dir() else set()
    missing_ordinary = sorted(expected_ordinary - actual_ordinary)
    cross_ordinary = sorted(
        path.removesuffix(".hx") + ".cross.hx"
        for path in expected_ordinary
        if (package_root / (path.removesuffix(".hx") + ".cross.hx")).is_file()
    )
    if missing_ordinary or cross_ordinary:
        violations.append(
            Violation(
                "package-ordinary-support-mapping",
                "src",
                "public facades and target support must remain ordinary .hx package modules; "
                f"missing={missing_ordinary!r}, converted={cross_ordinary!r}",
            )
        )

    leaks = find_absolute_path_leaks(package_root, [source_root, package_root])
    if leaks:
        violations.append(
            Violation(
                "absolute-path-leak",
                ", ".join(leaks[:5]),
                f"package contains absolute machine-local paths in {len(leaks)} file(s)",
            )
        )

    violations.extend(audit_package_map_manifest(package_root, source_root))

    return sorted(violations)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source-root", type=Path, default=Path.cwd())
    parser.add_argument("--package-root", type=Path)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    violations = audit_source_layout(args.source_root)
    if args.package_root is not None:
        violations.extend(audit_package_layout(args.package_root, args.source_root))

    if violations:
        for violation in sorted(violations):
            print(f"[canonical-std] ERROR: {violation.render()}")
        return 1

    scope = "source/package" if args.package_root is not None else "source"
    print(f"[canonical-std] OK: canonical {scope} layout")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
