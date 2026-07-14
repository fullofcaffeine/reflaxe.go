#!/usr/bin/env python3

"""Verify the exact origin and reconstruction of the vendored Reflaxe tree."""

from __future__ import annotations

import argparse
import hashlib
import json
import re
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
MANIFEST_PATH = ROOT / "provenance" / "reflaxe" / "vendor-manifest.json"
PATCH_PATH = ROOT / "provenance" / "reflaxe" / "upstream-to-supplier.patch"
DEFAULT_VENDOR_DIR = ROOT / "vendor" / "reflaxe"


class VerificationError(RuntimeError):
    """A provenance invariant did not match the committed evidence."""


def sha256_file(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def inventory(root: Path) -> dict[str, str]:
    if not root.is_dir():
        raise VerificationError(f"directory not found: {root}")

    result: dict[str, str] = {}
    for path in sorted(root.rglob("*")):
        relative = path.relative_to(root).as_posix()
        if path.is_symlink():
            raise VerificationError(f"{relative}: symlinks are not permitted")
        if path.is_file():
            result[relative] = sha256_file(path)
        elif not path.is_dir():
            raise VerificationError(f"{relative}: unsupported filesystem entry")
    return result


def tree_sha256(files: dict[str, str]) -> str:
    records = "".join(
        f"{digest}  {relative}\n" for relative, digest in sorted(files.items())
    )
    return hashlib.sha256(records.encode("utf-8")).hexdigest()


def git_object_sha1(kind: bytes, payload: bytes) -> bytes:
    header = kind + b" " + str(len(payload)).encode("ascii") + b"\0"
    return hashlib.sha1(header + payload).digest()


def git_tree_sha1(root: Path) -> str:
    entries: list[bytes] = []
    children = sorted(
        root.iterdir(),
        key=lambda path: path.name.encode("utf-8")
        + (b"/" if path.is_dir() else b""),
    )
    for path in children:
        name = path.name.encode("utf-8")
        if path.is_symlink():
            raise VerificationError(f"{path.name}: symlinks are not permitted")
        if path.is_dir():
            mode = b"40000"
            digest = bytes.fromhex(git_tree_sha1(path))
        elif path.is_file():
            mode = b"100755" if path.stat().st_mode & 0o111 else b"100644"
            digest = git_object_sha1(b"blob", path.read_bytes())
        else:
            raise VerificationError(f"{path.name}: unsupported filesystem entry")
        entries.append(mode + b" " + name + b"\0" + digest)
    return git_object_sha1(b"tree", b"".join(entries)).hex()


def expected_inventory(manifest: dict[str, object]) -> dict[str, str]:
    shipped = manifest.get("shipped_tree")
    if not isinstance(shipped, dict):
        raise VerificationError("manifest shipped_tree must be an object")
    entries = shipped.get("files")
    if not isinstance(entries, list):
        raise VerificationError("manifest shipped_tree.files must be a list")

    expected: dict[str, str] = {}
    for entry in entries:
        if not isinstance(entry, dict):
            raise VerificationError("manifest file entry must be an object")
        relative = entry.get("path")
        digest = entry.get("sha256")
        if not isinstance(relative, str) or not isinstance(digest, str):
            raise VerificationError("manifest file entry requires path and sha256")
        if relative in expected:
            raise VerificationError(f"duplicate manifest path: {relative}")
        expected[relative] = digest
    return expected


def compare_inventory(
    expected: dict[str, str],
    actual: dict[str, str],
    *,
    label: str,
) -> None:
    missing = sorted(set(expected) - set(actual))
    if missing:
        raise VerificationError(f"{label}: missing file: {missing[0]}")
    unexpected = sorted(set(actual) - set(expected))
    if unexpected:
        raise VerificationError(f"{label}: unexpected file: {unexpected[0]}")
    for relative in sorted(expected):
        if expected[relative] != actual[relative]:
            raise VerificationError(
                f"{relative}: digest mismatch "
                f"(expected {expected[relative]}, got {actual[relative]})"
            )


def run_git(arguments: list[str], *, cwd: Path) -> str:
    proc = subprocess.run(
        ["git", *arguments],
        cwd=cwd,
        capture_output=True,
        text=True,
    )
    if proc.returncode != 0:
        detail = proc.stderr.strip() or proc.stdout.strip() or "unknown git error"
        raise VerificationError(f"git {' '.join(arguments)} failed: {detail}")
    return proc.stdout.strip()


def verify_patch(manifest: dict[str, object], vendor_dir: Path) -> None:
    patch_set = manifest["patch_set"]
    official = manifest["official_upstream"]
    supplier = manifest["supplier_snapshot"]
    assert isinstance(patch_set, dict)
    assert isinstance(official, dict)
    assert isinstance(supplier, dict)

    expected_patch_sha = patch_set["sha256"]
    actual_patch_sha = sha256_file(PATCH_PATH)
    if actual_patch_sha != expected_patch_sha:
        raise VerificationError(
            "upstream-to-supplier.patch: digest mismatch "
            f"(expected {expected_patch_sha}, got {actual_patch_sha})"
        )

    patch_text = PATCH_PATH.read_text(encoding="utf-8")
    modified = re.findall(
        r"^diff --git a/src/reflaxe/(.+) b/src/reflaxe/\1$",
        patch_text,
        flags=re.MULTILINE,
    )
    expected_modified = patch_set["modified_source_files"]
    if modified != expected_modified:
        raise VerificationError("patch modified-file inventory does not match manifest")

    with tempfile.TemporaryDirectory(prefix="reflaxe-provenance-") as temp_name:
        work = Path(temp_name)
        source = work / "src" / "reflaxe"
        source.parent.mkdir(parents=True)
        shutil.copytree(vendor_dir / "src" / "reflaxe", source)
        shutil.copy2(PATCH_PATH, work / "change.patch")

        run_git(["apply", "--check", "-R", "change.patch"], cwd=work)
        run_git(["apply", "-R", "change.patch"], cwd=work)
        reverted = inventory(source)
        if len(reverted) != official["source_file_count"]:
            raise VerificationError("reconstructed official source file count mismatch")
        if tree_sha256(reverted) != official["source_tree_sha256"]:
            raise VerificationError("reconstructed official source digest mismatch")

        run_git(["apply", "--check", "change.patch"], cwd=work)
        run_git(["apply", "change.patch"], cwd=work)
        restored = inventory(source)
        if len(restored) != supplier["source_file_count"]:
            raise VerificationError("restored supplier source file count mismatch")
        if tree_sha256(restored) != supplier["source_tree_sha256"]:
            raise VerificationError("restored supplier source digest mismatch")


def checkout_exact(repository: str, commit: str, destination: Path) -> None:
    destination.mkdir()
    run_git(["init", "-q"], cwd=destination)
    run_git(["remote", "add", "origin", repository], cwd=destination)
    run_git(["fetch", "--depth=1", "origin", commit], cwd=destination)
    run_git(["checkout", "--detach", "-q", "FETCH_HEAD"], cwd=destination)
    actual = run_git(["rev-parse", "HEAD"], cwd=destination)
    if actual != commit:
        raise VerificationError(
            f"checkout resolved {actual}, expected immutable commit {commit}"
        )


def verify_network_reconstruction(
    manifest: dict[str, object],
    expected_vendor: dict[str, str],
) -> None:
    official = manifest["official_upstream"]
    supplier = manifest["supplier_snapshot"]
    assert isinstance(official, dict)
    assert isinstance(supplier, dict)

    with tempfile.TemporaryDirectory(prefix="reflaxe-reconstruct-") as temp_name:
        temp = Path(temp_name)
        supplier_checkout = temp / "supplier"
        official_checkout = temp / "official"
        checkout_exact(
            str(supplier["repository"]),
            str(supplier["commit"]),
            supplier_checkout,
        )
        checkout_exact(
            str(official["repository"]),
            str(official["commit"]),
            official_checkout,
        )

        supplier_tree = run_git(
            [
                "rev-parse",
                f"HEAD:{supplier['path']}",
            ],
            cwd=supplier_checkout,
        )
        if supplier_tree != supplier["git_tree_sha1"]:
            raise VerificationError(
                f"supplier Git tree mismatch: expected {supplier['git_tree_sha1']}, "
                f"got {supplier_tree}"
            )

        fetched_vendor = supplier_checkout / str(supplier["path"])
        compare_inventory(
            expected_vendor,
            inventory(fetched_vendor),
            label="supplier snapshot",
        )

        official_source = official_checkout / str(official["source_path"])
        official_files = inventory(official_source)
        if len(official_files) != official["source_file_count"]:
            raise VerificationError("official checkout source file count mismatch")
        if tree_sha256(official_files) != official["source_tree_sha256"]:
            raise VerificationError("official checkout source digest mismatch")
        if sha256_file(official_checkout / str(official["license_path"])) != official[
            "license_sha256"
        ]:
            raise VerificationError("official checkout license digest mismatch")

        shutil.copy2(PATCH_PATH, official_checkout / "change.patch")
        run_git(["apply", "--check", "change.patch"], cwd=official_checkout)
        run_git(["apply", "change.patch"], cwd=official_checkout)
        reconstructed = inventory(official_source)
        supplier_source = inventory(
            supplier_checkout / str(supplier["source_path"])
        )
        compare_inventory(
            supplier_source,
            reconstructed,
            label="patched official source",
        )


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Verify the shipped Reflaxe tree and its pinned origin chain."
    )
    parser.add_argument(
        "--vendor-dir",
        type=Path,
        default=DEFAULT_VENDOR_DIR,
        help="vendored Reflaxe directory to verify (default: repository vendor tree)",
    )
    parser.add_argument(
        "--reconstruct",
        action="store_true",
        help="fetch both pinned repositories and reconstruct the shipped source",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        manifest = json.loads(MANIFEST_PATH.read_text(encoding="utf-8"))
        if manifest.get("schema_version") != 1:
            raise VerificationError("unsupported manifest schema")

        vendor_dir = args.vendor_dir.resolve()
        expected = expected_inventory(manifest)
        actual = inventory(vendor_dir)
        compare_inventory(expected, actual, label="shipped vendor tree")

        shipped = manifest["shipped_tree"]
        assert isinstance(shipped, dict)
        if len(actual) != shipped["file_count"]:
            raise VerificationError("shipped vendor file count mismatch")
        actual_tree = tree_sha256(actual)
        if actual_tree != shipped["sha256"]:
            raise VerificationError(
                f"shipped vendor tree digest mismatch: expected {shipped['sha256']}, "
                f"got {actual_tree}"
            )
        supplier = manifest["supplier_snapshot"]
        assert isinstance(supplier, dict)
        actual_git_tree = git_tree_sha1(vendor_dir)
        if actual_git_tree != supplier["git_tree_sha1"]:
            raise VerificationError(
                f"shipped vendor Git tree mismatch: expected "
                f"{supplier['git_tree_sha1']}, got {actual_git_tree}"
            )
        print(
            "[vendor-provenance] shipped tree: OK "
            f"({len(actual)} files, sha256 {actual_tree}, git {actual_git_tree})"
        )

        verify_patch(manifest, vendor_dir)
        modified_count = len(manifest["patch_set"]["modified_source_files"])
        print(
            "[vendor-provenance] patch round-trip: OK "
            f"({modified_count} modified source files)"
        )

        if args.reconstruct:
            verify_network_reconstruction(manifest, expected)
            print("[vendor-provenance] network reconstruction: OK")
    except (OSError, KeyError, TypeError, ValueError, VerificationError) as error:
        print(f"[vendor-provenance] error: {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
