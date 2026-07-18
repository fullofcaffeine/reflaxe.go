#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
LEDGER_PATH = ROOT / "docs" / "stdlib-provenance-ledger.json"
CHECKER_PATH = ROOT / "scripts" / "ci" / "stdlib-provenance-ledger-check.js"
RELEASE_RUNNER_PATH = ROOT / "test" / "run-release-contracts.py"
PACKAGE_PATH = ROOT / "package.json"

OWNERSHIP_CLASSES = {
    "upstream_std_override",
    "staged_support",
    "hxrt_binding",
    "public_go_facade",
    "obsolete",
    "intentional_boundary_fixture",
}

TARGET_SUPPORT_OWNERS = {
    "staged_support",
    "hxrt_binding",
    "public_go_facade",
}


def canonical_override_destination(source_path: str) -> str:
    if source_path.startswith("std/go/_std/"):
        return source_path
    relative = source_path.removeprefix("std/")
    relative = relative.removeprefix("_std/")
    if relative.endswith(".cross.hx"):
        relative = relative.removesuffix(".cross.hx") + ".hx"
    return "std/go/_std/" + relative


SOURCE_SPECIAL_DESTINATIONS = {
    "std/go/Context.hx": ("public_go_facade", "std/go/Context.hx"),
    "std/go/ContextPkg.hx": ("public_go_facade", "std/go/ContextPkg.hx"),
    "std/go/Fmt.hx": ("public_go_facade", "std/go/Fmt.hx"),
    "std/go/Http.hx": ("public_go_facade", "std/go/Http.hx"),
    "std/go/Time.hx": ("public_go_facade", "std/go/Time.hx"),
    "std/haxe/GoSerializationBridge.hx": (
        "staged_support",
        "std/haxe/GoSerializationBridge.hx",
    ),
    "std/sys/thread/ElasticThreadPoolWorker.cross.hx": (
        "staged_support",
        "std/sys/thread/ElasticThreadPoolWorker.hx",
    ),
    "std/sys/thread/FixedThreadPoolShutdownException.cross.hx": (
        "staged_support",
        "std/sys/thread/FixedThreadPoolShutdownException.hx",
    ),
    "std/sys/thread/FixedThreadPoolWorker.cross.hx": (
        "staged_support",
        "std/sys/thread/FixedThreadPoolWorker.hx",
    ),
    "std/sys/net/_SocketIO.hx": (
        "staged_support",
        "std/sys/net/_SocketIO.hx",
    ),
    "std/_std/haxe/iterators/GoStringRuntime.cross.hx": (
        "hxrt_binding",
        "std/hxrt/string/GoStringRuntime.hx",
    ),
    "std/hxrt/atomic/AtomicIntHandle.hx": (
        "hxrt_binding",
        "std/hxrt/atomic/AtomicIntHandle.hx",
    ),
    "std/hxrt/atomic/AtomicObjectHandle.hx": (
        "hxrt_binding",
        "std/hxrt/atomic/AtomicObjectHandle.hx",
    ),
    "std/hxrt/atomic/NativeAtomicInt.hx": (
        "hxrt_binding",
        "std/hxrt/atomic/NativeAtomicInt.hx",
    ),
    "std/hxrt/atomic/NativeAtomicObject.hx": (
        "hxrt_binding",
        "std/hxrt/atomic/NativeAtomicObject.hx",
    ),
    "std/hxrt/collections/IntMapHandle.hx": (
        "hxrt_binding",
        "std/hxrt/collections/IntMapHandle.hx",
    ),
    "std/hxrt/collections/NativeEnumValue.hx": (
        "hxrt_binding",
        "std/hxrt/collections/NativeEnumValue.hx",
    ),
    "std/hxrt/collections/NativeIntMap.hx": (
        "hxrt_binding",
        "std/hxrt/collections/NativeIntMap.hx",
    ),
    "std/hxrt/collections/NativeObjectMap.hx": (
        "hxrt_binding",
        "std/hxrt/collections/NativeObjectMap.hx",
    ),
    "std/hxrt/collections/NativeStringMap.hx": (
        "hxrt_binding",
        "std/hxrt/collections/NativeStringMap.hx",
    ),
    "std/hxrt/collections/ObjectMapHandle.hx": (
        "hxrt_binding",
        "std/hxrt/collections/ObjectMapHandle.hx",
    ),
    "std/hxrt/collections/StringMapHandle.hx": (
        "hxrt_binding",
        "std/hxrt/collections/StringMapHandle.hx",
    ),
    "std/hxrt/http/HttpRequestHandle.hx": (
        "hxrt_binding",
        "std/hxrt/http/HttpRequestHandle.hx",
    ),
    "std/hxrt/http/HttpResponseHandle.hx": (
        "hxrt_binding",
        "std/hxrt/http/HttpResponseHandle.hx",
    ),
    "std/hxrt/http/NativeHttp.hx": (
        "hxrt_binding",
        "std/hxrt/http/NativeHttp.hx",
    ),
    "std/hxrt/crypto/NativeCrypto.hx": (
        "hxrt_binding",
        "std/hxrt/crypto/NativeCrypto.hx",
    ),
    "std/hxrt/io/ByteView.hx": (
        "hxrt_binding",
        "std/hxrt/io/ByteView.hx",
    ),
    "std/hxrt/io/NativeBytes.hx": (
        "hxrt_binding",
        "std/hxrt/io/NativeBytes.hx",
    ),
    "std/hxrt/io/NativeFloatBits.hx": (
        "hxrt_binding",
        "std/hxrt/io/NativeFloatBits.hx",
    ),
    "std/hxrt/date/DateParts.hx": (
        "hxrt_binding",
        "std/hxrt/date/DateParts.hx",
    ),
    "std/hxrt/date/NativeDate.hx": (
        "hxrt_binding",
        "std/hxrt/date/NativeDate.hx",
    ),
    "std/hxrt/math/NativeMath.hx": (
        "hxrt_binding",
        "std/hxrt/math/NativeMath.hx",
    ),
    "std/hxrt/math/NativeMathInt.hx": (
        "hxrt_binding",
        "std/hxrt/math/NativeMathInt.hx",
    ),
    "std/hxrt/math/NativeRandom.hx": (
        "hxrt_binding",
        "std/hxrt/math/NativeRandom.hx",
    ),
    "std/hxrt/zip/NativeZip.hx": (
        "hxrt_binding",
        "std/hxrt/zip/NativeZip.hx",
    ),
    "std/hxrt/fs/FileSystemStat.hx": (
        "hxrt_binding",
        "std/hxrt/fs/FileSystemStat.hx",
    ),
    "std/hxrt/fs/NativeFileSystem.hx": (
        "hxrt_binding",
        "std/hxrt/fs/NativeFileSystem.hx",
    ),
    "std/hxrt/fs/FileInputHandle.hx": (
        "hxrt_binding",
        "std/hxrt/fs/FileInputHandle.hx",
    ),
    "std/hxrt/fs/FileOutputHandle.hx": (
        "hxrt_binding",
        "std/hxrt/fs/FileOutputHandle.hx",
    ),
    "std/hxrt/fs/NativeFile.hx": (
        "hxrt_binding",
        "std/hxrt/fs/NativeFile.hx",
    ),
    "std/hxrt/sys/NativeConsole.hx": (
        "hxrt_binding",
        "std/hxrt/sys/NativeConsole.hx",
    ),
    "std/hxrt/sys/NativeSys.hx": (
        "hxrt_binding",
        "std/hxrt/sys/NativeSys.hx",
    ),
    "std/hxrt/sys/NativeTerminal.hx": (
        "hxrt_binding",
        "std/hxrt/sys/NativeTerminal.hx",
    ),
    "std/hxrt/sys/SysEnvironmentEntry.hx": (
        "hxrt_binding",
        "std/hxrt/sys/SysEnvironmentEntry.hx",
    ),
    "std/hxrt/process/NativeProcess.hx": (
        "hxrt_binding",
        "std/hxrt/process/NativeProcess.hx",
    ),
    "std/hxrt/process/ProcessExitStatus.hx": (
        "hxrt_binding",
        "std/hxrt/process/ProcessExitStatus.hx",
    ),
    "std/hxrt/process/ProcessHandle.hx": (
        "hxrt_binding",
        "std/hxrt/process/ProcessHandle.hx",
    ),
    "std/hxrt/process/ProcessInputHandle.hx": (
        "hxrt_binding",
        "std/hxrt/process/ProcessInputHandle.hx",
    ),
    "std/hxrt/process/ProcessOutputHandle.hx": (
        "hxrt_binding",
        "std/hxrt/process/ProcessOutputHandle.hx",
    ),
    "std/hxrt/regex/NativeRegex.hx": (
        "hxrt_binding",
        "std/hxrt/regex/NativeRegex.hx",
    ),
    "std/hxrt/regex/RegexHandle.hx": (
        "hxrt_binding",
        "std/hxrt/regex/RegexHandle.hx",
    ),
    "std/hxrt/regex/RegexMatch.hx": (
        "hxrt_binding",
        "std/hxrt/regex/RegexMatch.hx",
    ),
    "std/hxrt/serialization/NativeSerialization.hx": (
        "hxrt_binding",
        "std/hxrt/serialization/NativeSerialization.hx",
    ),
    "std/hxrt/serialization/SerializationField.hx": (
        "hxrt_binding",
        "std/hxrt/serialization/SerializationField.hx",
    ),
    "std/hxrt/net/NativeSocket.hx": (
        "hxrt_binding",
        "std/hxrt/net/NativeSocket.hx",
    ),
    "std/hxrt/net/SocketAcceptResult.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketAcceptResult.hx",
    ),
    "std/hxrt/net/SocketAddress.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketAddress.hx",
    ),
    "std/hxrt/net/SocketDatagramResult.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketDatagramResult.hx",
    ),
    "std/hxrt/net/SocketHandle.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketHandle.hx",
    ),
    "std/hxrt/net/SocketIOResult.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketIOResult.hx",
    ),
    "std/hxrt/net/SocketSelectResult.hx": (
        "hxrt_binding",
        "std/hxrt/net/SocketSelectResult.hx",
    ),
    "std/hxrt/ssl/CertificateHandle.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/CertificateHandle.hx",
    ),
    "std/hxrt/ssl/KeyHandle.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/KeyHandle.hx",
    ),
    "std/hxrt/ssl/NativeDigest.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/NativeDigest.hx",
    ),
    "std/hxrt/ssl/NativeKey.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/NativeKey.hx",
    ),
    "std/hxrt/ssl/NativeCertificate.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/NativeCertificate.hx",
    ),
    "std/hxrt/ssl/NativeSocket.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/NativeSocket.hx",
    ),
    "std/hxrt/ssl/SNIConfigHandle.hx": (
        "hxrt_binding",
        "std/hxrt/ssl/SNIConfigHandle.hx",
    ),
    "std/_std/hxrt/stack/NativeStack.hx": ("hxrt_binding", "std/hxrt/stack/NativeStack.hx"),
    "std/_std/hxrt/stack/NativeStackFrame.hx": (
        "hxrt_binding",
        "std/hxrt/stack/NativeStackFrame.hx",
    ),
    "std/hxrt/template/NativeTemplate.hx": (
        "hxrt_binding",
        "std/hxrt/template/NativeTemplate.hx",
    ),
    "std/_std/hxrt/thread/ConditionHandle.hx": (
        "hxrt_binding",
        "std/hxrt/thread/ConditionHandle.hx",
    ),
    "std/_std/hxrt/thread/EventLoopHandle.hx": (
        "hxrt_binding",
        "std/hxrt/thread/EventLoopHandle.hx",
    ),
    "std/_std/hxrt/thread/EventLoopProgress.hx": (
        "hxrt_binding",
        "std/hxrt/thread/EventLoopProgress.hx",
    ),
    "std/_std/hxrt/thread/LockHandle.hx": ("hxrt_binding", "std/hxrt/thread/LockHandle.hx"),
    "std/_std/hxrt/thread/MutexHandle.hx": ("hxrt_binding", "std/hxrt/thread/MutexHandle.hx"),
    "std/_std/hxrt/thread/NativeThread.hx": ("hxrt_binding", "std/hxrt/thread/NativeThread.hx"),
    "std/_std/hxrt/thread/SemaphoreHandle.hx": (
        "hxrt_binding",
        "std/hxrt/thread/SemaphoreHandle.hx",
    ),
    "std/_std/hxrt/thread/ThreadLocalHandle.hx": (
        "hxrt_binding",
        "std/hxrt/thread/ThreadLocalHandle.hx",
    ),
}

SPECIAL_DESTINATIONS = dict(SOURCE_SPECIAL_DESTINATIONS)
for special_owner, special_destination in SOURCE_SPECIAL_DESTINATIONS.values():
    SPECIAL_DESTINATIONS[special_destination] = (special_owner, special_destination)

SOURCE_EXPECTED_SHIM_GROUPS = {
    "std/haxe/Template.cross.hx": ["stdlib_symbols"],
    "std/haxe/ds/BalancedTree.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Certificate.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Digest.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Key.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Socket.cross.hx": ["stdlib_symbols"],
}

EXPECTED_SHIM_GROUPS: dict[str, list[str]] = {}
for shim_source, shim_groups in SOURCE_EXPECTED_SHIM_GROUPS.items():
    shim_destination = SPECIAL_DESTINATIONS.get(
        shim_source,
        ("upstream_std_override", canonical_override_destination(shim_source)),
    )[1]
    EXPECTED_SHIM_GROUPS[shim_source] = shim_groups
    EXPECTED_SHIM_GROUPS[shim_destination] = shim_groups

EXPECTED_SHIM_AUDIT_DECISIONS = {
    "stdlib_symbols": "split_migration_debt_from_exact_intrinsics_haxe_go_vfp_8_7_15",
}

SHIM_AUDIT_AUTHORITY_REFERENCES = {
    "docs/compiler-stdlib-intrinsics.json",
    "test/test_compiler_stdlib_intrinsic_registry.py",
}


def load_ledger() -> dict[str, object]:
    return json.loads(LEDGER_PATH.read_text(encoding="utf-8"))


def tracked_std_sources() -> set[str]:
    process = subprocess.run(
        ["git", "ls-files", "-z", "--", "std"],
        cwd=ROOT,
        check=True,
        capture_output=True,
    )
    return {
        raw.decode("utf-8")
        for raw in process.stdout.split(b"\0")
        if raw and raw.decode("utf-8").endswith((".hx", ".cross.hx"))
    }

class StdlibMigrationLedgerContractTest(unittest.TestCase):
    def test_upstream_overrides_are_canonical_documented_source(self) -> None:
        failures: list[str] = []
        for entry in load_ledger()["entries"]:
            if entry["ownershipClass"] != "upstream_std_override":
                continue
            source_path = entry["path"]
            destination = entry["destination"]
            if source_path != destination:
                failures.append(f"{source_path}: not migrated to {destination}")
                continue

            content = (ROOT / source_path).read_text(encoding="utf-8")
            haxedocs = re.findall(r"/\*\*(.*?)\*/", content, flags=re.DOTALL)
            contract_doc = next(
                (
                    doc
                    for doc in haxedocs
                    if re.search(r"\bWhat\s*:?(?:\s|$)", doc)
                    and re.search(r"\bWhy\s*:?(?:\s|$)", doc)
                    and re.search(r"\bHow\s*:?(?:\s|$)", doc)
                    and "mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`"
                    in " ".join(doc.split())
                ),
                None,
            )
            if contract_doc is None:
                failures.append(f"{source_path}: missing What/Why/How override HaxeDoc")

        self.assertFalse(failures, "canonical override migration is incomplete:\n" + "\n".join(failures))

    def test_target_support_and_public_facades_are_canonical_documented_source(self) -> None:
        failures: list[str] = []
        for entry in load_ledger()["entries"]:
            ownership = entry["ownershipClass"]
            if ownership not in TARGET_SUPPORT_OWNERS:
                continue

            source_path = entry["path"]
            destination = entry["destination"]
            if source_path != destination:
                failures.append(f"{source_path}: not migrated to {destination}")

            destination_path = ROOT / destination
            if not destination_path.is_file():
                failures.append(f"{destination}: canonical source is missing")
                continue
            if destination.endswith(".cross.hx"):
                failures.append(f"{destination}: target support must remain ordinary .hx source")

            content = destination_path.read_text(encoding="utf-8")
            haxedocs = re.findall(r"/\*\*(.*?)\*/", content, flags=re.DOTALL)
            if not any(
                re.search(r"\bWhat\s*:?(?:\s|$)", doc)
                and re.search(r"\bWhy\s*:?(?:\s|$)", doc)
                and re.search(r"\bHow\s*:?(?:\s|$)", doc)
                for doc in haxedocs
            ):
                failures.append(f"{destination}: missing What/Why/How ownership HaxeDoc")

            if ownership == "public_go_facade":
                primary_types = re.findall(
                    r"(?m)^(?:extern\s+)?(?:class|interface|enum|abstract|typedef)\s+([A-Za-z_]\w*)",
                    content,
                )
                if primary_types != [destination_path.stem]:
                    failures.append(
                        f"{destination}: expected one matching primary type; found {primary_types!r}"
                    )

        self.assertFalse(
            failures,
            "target-support migration is incomplete:\n" + "\n".join(failures),
        )

    def test_every_tracked_std_source_has_one_resolved_owner_and_destination(self) -> None:
        ledger = load_ledger()
        self.assertEqual(3, ledger.get("schemaVersion"))
        entries = ledger.get("entries")
        self.assertIsInstance(entries, list)
        migration_beads = ledger.get("migrationContract", {}).get("migrationBeads")
        self.assertIsInstance(migration_beads, dict)
        for ownership in OWNERSHIP_CLASSES:
            allowed_beads = migration_beads.get(ownership)
            self.assertIsInstance(allowed_beads, list, ownership)
            self.assertTrue(allowed_beads, ownership)
            self.assertEqual(sorted(set(allowed_beads)), allowed_beads, ownership)
            for bead in allowed_beads:
                self.assertRegex(bead, r"^haxe_go-[a-z0-9.-]+$", ownership)

        by_path = {entry["path"]: entry for entry in entries}
        self.assertEqual(len(entries), len(by_path), "migration ledger paths must be unique")
        self.assertEqual(tracked_std_sources(), set(by_path), "no tracked std source may be silently dropped")

        destinations: list[str] = []
        for source_path, entry in sorted(by_path.items()):
            ownership = entry.get("ownershipClass")
            destination = entry.get("destination")
            self.assertIn(ownership, OWNERSHIP_CLASSES, source_path)
            self.assertIsInstance(destination, str, source_path)
            self.assertTrue(destination, source_path)
            self.assertFalse(destination.endswith(".cross.hx"), source_path)
            destinations.append(destination)

            expected_owner, expected_destination = SPECIAL_DESTINATIONS.get(
                source_path,
                ("upstream_std_override", canonical_override_destination(source_path)),
            )
            self.assertEqual(expected_owner, ownership, source_path)
            self.assertEqual(expected_destination, destination, source_path)

            self.assertIn(entry.get("migrationBead"), migration_beads[ownership], source_path)

        self.assertEqual(len(destinations), len(set(destinations)), "migration destinations must be unique")

    def test_destination_roots_follow_ownership_boundaries(self) -> None:
        for entry in load_ledger()["entries"]:
            source_path = entry["path"]
            ownership = entry["ownershipClass"]
            destination = entry["destination"]
            if ownership == "upstream_std_override":
                self.assertTrue(destination.startswith("std/go/_std/"), source_path)
            elif ownership == "hxrt_binding":
                self.assertTrue(destination.startswith("std/hxrt/"), source_path)
            elif ownership == "public_go_facade":
                self.assertTrue(destination.startswith("std/go/"), source_path)
                self.assertNotIn("/_std/", destination, source_path)
            elif ownership == "staged_support":
                self.assertTrue(destination.startswith("std/"), source_path)
                self.assertNotIn("/_std/", destination, source_path)
                self.assertFalse(destination.startswith("std/hxrt/"), source_path)

    def test_compiler_shim_audit_is_exact_and_bidirectional(self) -> None:
        ledger = load_ledger()
        contract = ledger.get("migrationContract")
        self.assertIsInstance(contract, dict)
        audits = contract.get("compilerShimAudit")
        self.assertIsInstance(audits, list)
        audit_by_group = {audit["group"]: audit for audit in audits}
        self.assertEqual(
            set(EXPECTED_SHIM_AUDIT_DECISIONS),
            set(audit_by_group),
        )

        actual_paths_by_group: dict[str, set[str]] = {group: set() for group in audit_by_group}
        for entry in ledger["entries"]:
            groups = entry.get("compilerShimGroups")
            self.assertIsInstance(groups, list, entry["path"])
            self.assertEqual(sorted(set(groups)), groups, entry["path"])
            self.assertEqual(EXPECTED_SHIM_GROUPS.get(entry["path"], []), groups, entry["path"])
            for group in groups:
                self.assertIn(group, audit_by_group, entry["path"])
                actual_paths_by_group[group].add(entry["path"])

        for group, audit in sorted(audit_by_group.items()):
            self.assertEqual(actual_paths_by_group[group], set(audit.get("sourcePaths", [])), group)
            self.assertEqual(EXPECTED_SHIM_AUDIT_DECISIONS[group], audit.get("decision"), group)
            references = audit.get("references")
            self.assertIsInstance(references, list, group)
            self.assertTrue(SHIM_AUDIT_AUTHORITY_REFERENCES.issubset(references), group)
            for reference in references:
                self.assertTrue((ROOT / reference).is_file(), f"{group}: missing audit reference {reference}")

    def test_filesystem_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_owners = {
            "std/go/_std/sys/FileSystem.hx": "upstream_std_override",
            "std/hxrt/fs/FileSystemStat.hx": "hxrt_binding",
            "std/hxrt/fs/NativeFileSystem.hx": "hxrt_binding",
        }
        for source_path, expected_owner in expected_owners.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual("haxe_go-vfp.8.7.4", entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )
        filesystem_runtime = ROOT / "runtime/hxrt/filesystem.go"
        self.assertTrue(filesystem_runtime.is_file())
        self.assertNotIn(
            "func FileSystem",
            (ROOT / "runtime/hxrt/sys.go").read_text(encoding="utf-8"),
        )
        self.assertIn('case FEATURE_FILESYSTEM:', feature_analyzer)
        self.assertIn('["filesystem.go"]', feature_analyzer)
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )

        self.assertNotIn("lowerFileSystemShimDecls", compiler)
        self.assertNotIn('requiredStdlibShimGroups.exists("filesystem")', compiler)
        self.assertNotIn('GoStructDecl("sys__FileSystem"', compiler)
        self.assertNotIn('return ["filesystem"]', classifier)
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "filesystem"
                for limit in debt_policy["limits"]
            ),
            "source-owned sys.FileSystem must not retain a compiler-shim debt allowance",
        )

    def test_file_io_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_owners = {
            "std/go/_std/sys/io/File.hx": "upstream_std_override",
            "std/go/_std/sys/io/FileInput.hx": "upstream_std_override",
            "std/go/_std/sys/io/FileOutput.hx": "upstream_std_override",
            "std/go/_std/sys/io/FileSeek.hx": "upstream_std_override",
            "std/hxrt/fs/FileInputHandle.hx": "hxrt_binding",
            "std/hxrt/fs/FileOutputHandle.hx": "hxrt_binding",
            "std/hxrt/fs/NativeFile.hx": "hxrt_binding",
        }
        for source_path, expected_owner in expected_owners.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual("haxe_go-vfp.8.7.5", entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        ownership = (ROOT / "src/reflaxe/go/compiler/GoStdlibOwnership.hx").read_text(
            encoding="utf-8"
        )
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )

        self.assertTrue((ROOT / "runtime/hxrt/file.go").is_file())
        sys_runtime = (ROOT / "runtime/hxrt/sys.go").read_text(encoding="utf-8")
        self.assertNotIn("type FileInput struct", sys_runtime)
        self.assertNotIn("type FileOutput struct", sys_runtime)
        self.assertNotIn("func File", sys_runtime)
        self.assertNotIn("func OpenFile", sys_runtime)
        self.assertIn("FEATURE_FILE_IO", feature_analyzer)
        self.assertIn('["file.go"]', feature_analyzer)

        forbidden_compiler_fragments = [
            'GoStructDecl("sys__io__File"',
            'GoStructDecl("sys__io__FileInput"',
            'GoStructDecl("sys__io__FileOutput"',
            'GoStructDecl("sys__io__FileSeek"',
            'GoGlobalVarDecl("sys__io__fileInputHandles"',
            'GoGlobalVarDecl("sys__io__fileOutputHandles"',
            'GoFuncDecl("sys__io__fileSeekWhence"',
            'GoFuncDecl("sys__io__File_',
            'lowerIoInputSyntheticHelper("*sys__io__FileInput"',
            'lowerIoOutputSyntheticHelper("*sys__io__FileOutput"',
            'isSysFileInput',
            'isSysFileOutput',
        ]
        for fragment in forbidden_compiler_fragments:
            self.assertNotIn(fragment, compiler, fragment)

        for source_owned_name in (
            '"sys.io.File"',
            '"sys.io.FileInput"',
            '"sys.io.FileOutput"',
            '"sys.io.FileSeek"',
        ):
            self.assertNotIn(source_owned_name, ownership, source_owned_name)
        self.assertNotIn('classType.name == "File"', classifier)

    def test_root_sys_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_entries = {
            "std/go/_std/Sys.hx": ("upstream_std_override", "haxe_go-vfp.8.7.6"),
            "std/hxrt/sys/NativeConsole.hx": ("hxrt_binding", "haxe_go-vfp.8.7.6"),
            "std/hxrt/sys/NativeSys.hx": ("hxrt_binding", "haxe_go-vfp.8.7.6"),
            "std/hxrt/sys/NativeTerminal.hx": ("hxrt_binding", "haxe_go-vfp.8.7.3"),
            "std/hxrt/sys/SysEnvironmentEntry.hx": ("hxrt_binding", "haxe_go-vfp.8.7.6"),
        }
        for source_path, (expected_owner, expected_bead) in expected_entries.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual(expected_bead, entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )

        for fragment in (
            'GoStructDecl("Sys"',
            'GoFuncDecl("Sys_',
            'isStaticCall(callee, "Sys", [], "print")',
            'isStaticCall(callee, "Sys", [], "println")',
            "requiresSysCommandSurface",
            "lowerSysStdlibShimDecls",
        ):
            self.assertNotIn(fragment, compiler, fragment)

        self.assertNotIn('(pack == "" && classType.name == "Sys")', classifier)
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "sys"
                for limit in debt_policy["limits"]
            ),
            "source-owned root Sys must not retain a compiler-shim debt allowance",
        )

    def test_process_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_owners = {
            "std/go/_std/sys/io/Process.hx": "upstream_std_override",
            "std/hxrt/process/NativeProcess.hx": "hxrt_binding",
            "std/hxrt/process/ProcessExitStatus.hx": "hxrt_binding",
            "std/hxrt/process/ProcessHandle.hx": "hxrt_binding",
            "std/hxrt/process/ProcessInputHandle.hx": "hxrt_binding",
            "std/hxrt/process/ProcessOutputHandle.hx": "hxrt_binding",
        }
        for source_path, expected_owner in expected_owners.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual("haxe_go-vfp.8.7.7", entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )

        for fragment in (
            "lowerProcessStdlibShimDecls",
            'GoStructDecl("sys__io__Process',
            'GoFuncDecl("New_sys__io__Process"',
            'lowerIoInputSyntheticHelper("*sys__io__ProcessOutput"',
            'lowerIoOutputSyntheticHelper("*sys__io__ProcessInput"',
        ):
            self.assertNotIn(fragment, compiler, fragment)

        self.assertNotIn('classType.name == "Process"', classifier)
        self.assertNotIn('return ["process"]', classifier)
        self.assertNotIn('case "process":', feature_analyzer)
        self.assertIn('path == "hxrt.process.NativeProcess"', feature_analyzer)
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "process"
                for limit in debt_policy["limits"]
            ),
            "source-owned sys.io.Process must not retain a compiler-shim debt allowance",
        )

    def test_template_runtime_bridge_is_typed_hxrt_instead_of_a_compiler_shim(self) -> None:
        ledger = load_ledger()
        ledger_entries = {entry["path"]: entry for entry in ledger["entries"]}
        template_entry = ledger_entries.get("std/go/_std/haxe/Template.hx")
        native_entry = ledger_entries.get("std/hxrt/template/NativeTemplate.hx")

        self.assertIsNotNone(template_entry)
        self.assertEqual(["stdlib_symbols"], template_entry.get("compilerShimGroups"))
        self.assertEqual("hxrt_binding", native_entry.get("ownershipClass"))
        self.assertEqual("haxe_go-vfp.8.7.16", native_entry.get("migrationBead"))
        self.assertEqual([], native_entry.get("compilerShimGroups"))
        self.assertIn(
            "haxe_go-vfp.8.7.16",
            ledger["migrationContract"]["migrationBeads"]["hxrt_binding"],
        )
        self.assertFalse(
            any(
                entry.get("group") == "template_support"
                for entry in ledger["migrationContract"]["compilerShimAudit"]
            )
        )

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )
        debt_ratchet = (ROOT / "test/run-compiler-debt-ratchet.py").read_text(
            encoding="utf-8"
        )
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )
        staged_template = (ROOT / "std/go/_std/haxe/Template.hx").read_text(
            encoding="utf-8"
        )
        native_template = (
            ROOT / "std/hxrt/template/NativeTemplate.hx"
        ).read_text(encoding="utf-8")
        runtime_template = (ROOT / "runtime/hxrt/template.go").read_text(
            encoding="utf-8"
        )

        for fragment in (
            "lowerTemplateSupportShimDecls",
            'requiredStdlibShimGroups.exists("template_support")',
            "haxe__Template_anyArrayToSlice_runtime",
            'GoFuncDecl("Reflect_getProperty"',
            'GoFuncDecl("Reflect_isObject"',
            'GoFuncDecl("Reflect_callMethod"',
        ):
            self.assertNotIn(fragment, compiler, fragment)
        self.assertNotIn('requireStdlibShimGroup("template_support")', planner)

        for fragment in (
            "import hxrt.template.NativeTemplate",
            "NativeTemplate.arrayValues",
            "NativeTemplate.isObject",
            "NativeTemplate.call",
        ):
            self.assertIn(fragment, staged_template, fragment)
        for fragment in ("__go__", "GoInjection", "@:goAllowRaw"):
            self.assertNotIn(fragment, staged_template, fragment)

        for heading in ("What", "Why", "How"):
            self.assertIn(heading, native_template)
        for fragment in (
            '@:go.name("TemplateArrayValues")',
            '@:go.name("TemplateIsObject")',
            '@:go.name("TemplateCall")',
        ):
            self.assertIn(fragment, native_template, fragment)
        for signature in (
            "func TemplateArrayValues(value any) []any",
            "func TemplateIsObject(value any) bool",
            "func TemplateCall(funcValue any, args []any) any",
        ):
            self.assertIn(signature, runtime_template, signature)

        self.assertIn('FEATURE_TEMPLATE = "template"', feature_analyzer)
        self.assertIn('[FEATURE_CORE, FEATURE_ARRAY]', feature_analyzer)
        self.assertIn('path == "hxrt.template.NativeTemplate"', feature_analyzer)
        self.assertIn('["template.go"]', feature_analyzer)
        self.assertIn('case "template.go":', reflaxe_compiler)

        generated_root = ROOT / "test/snapshot/stdlib/haxe_template_basic/intended"
        generated_main = (generated_root / "main.go").read_text(encoding="utf-8")
        generated_template = (generated_root / "module_haxe_template.go").read_text(
            encoding="utf-8"
        )
        for fragment in (
            "haxe__Template_anyArrayToSlice_runtime",
            "func Reflect_getProperty",
            "func Reflect_isObject",
            "func Reflect_callMethod",
        ):
            self.assertNotIn(fragment, generated_main, fragment)
        for fragment in (
            "hxrt.TemplateArrayValues",
            "hxrt.TemplateIsObject",
            "hxrt.TemplateCall",
        ):
            self.assertIn(fragment, generated_template, fragment)
        self.assertTrue((generated_root / "hxrt/template.go").is_file())
        self.assertFalse(
            (
                ROOT
                / "test/snapshot/core/const_kinds_contract/intended/hxrt/template.go"
            ).exists(),
            "Template-only reflection support must not enter unrelated generated programs",
        )

        self.assertFalse(
            any(group.get("group") == "template_support" for group in registry["groups"])
        )
        self.assertNotIn("migration_template", registry["decisions"])
        self.assertFalse(
            any(
                limit.get("capability") == "template"
                and limit.get("metric") in {"compiler_shim", "go_raw"}
                for limit in debt_policy["limits"]
            )
        )
        self.assertNotIn('"lowerTemplateSupportShimDecls": "template"', debt_ratchet)

    def test_crypto_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")

        crypto_paths = (
            "haxe.crypto.Base64",
            "haxe.crypto.Md5",
            "haxe.crypto.Sha1",
            "haxe.crypto.Sha224",
            "haxe.crypto.Sha256",
        )
        crypto_symbols = (
            "haxe__crypto__Base64",
            "haxe__crypto__Md5",
            "haxe__crypto__Sha1",
            "haxe__crypto__Sha224",
            "haxe__crypto__Sha256",
        )

        stdlib_start = compiler.index("function lowerStdlibSymbolShimDecls")
        stdlib_end = compiler.index("function reflectFieldsShimDecl", stdlib_start)
        stdlib_emitter = compiler[stdlib_start:stdlib_end]
        for path, symbol in zip(crypto_paths, crypto_symbols):
            self.assertNotIn(f'path: "{path}"', classifier)
            self.assertIn(f'"{path}"', planner)
            self.assertNotIn(symbol, stdlib_emitter)

        stdlib_group = next(
            group for group in registry["groups"] if group["group"] == "stdlib_symbols"
        )
        owned_symbols = {entry["symbol"] for entry in stdlib_group["ownedSymbols"]}
        selected_paths = {
            entry["path"] for entry in stdlib_group["classifierSelections"]
        }
        for path in crypto_paths:
            self.assertNotIn(path, owned_symbols)
            self.assertNotIn(path, selected_paths)

        for source_path in (
            "std/go/_std/haxe/crypto/Base64.hx",
            "std/go/_std/haxe/crypto/Md5.hx",
            "std/go/_std/haxe/crypto/Sha1.hx",
            "std/go/_std/haxe/crypto/Sha224.hx",
            "std/go/_std/haxe/crypto/Sha256.hx",
        ):
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.15.1", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        native_entry = ledger_entries.get("std/hxrt/crypto/NativeCrypto.hx")
        self.assertIsNotNone(native_entry)
        self.assertEqual("hxrt_binding", native_entry.get("ownershipClass"))
        self.assertEqual("haxe_go-vfp.8.7.15.1", native_entry.get("migrationBead"))
        self.assertEqual([], native_entry.get("compilerShimGroups"))

        self.assertIn('FEATURE_CRYPTO = "crypto"', feature_analyzer)
        self.assertIn('["crypto.go"]', feature_analyzer)
        self.assertIn('case "crypto.go":', reflaxe_compiler)
        generated_root = ROOT / "test/snapshot/stdlib/crypto_xml_zip_basic/intended"
        self.assertTrue((generated_root / "hxrt/crypto.go").is_file())
        self.assertFalse(
            (ROOT / "test/snapshot/core/const_kinds_contract/intended/hxrt/crypto.go").exists()
        )

    def test_xml_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )

        xml_paths = ("Xml", "haxe.xml.Parser", "haxe.xml.Printer")
        for path in xml_paths:
            self.assertNotIn(f'path: "{path}"', classifier)
            self.assertIn(f'"{path}"', planner)

        stdlib_start = compiler.index("function lowerStdlibSymbolShimDecls")
        stdlib_end = compiler.index("function reflectFieldsShimDecl", stdlib_start)
        stdlib_emitter = compiler[stdlib_start:stdlib_end]
        for symbol in (
            'GoDecl.GoStructDecl("Xml"',
            'GoDecl.GoStructDecl("haxe__xml__Parser"',
            'GoDecl.GoStructDecl("haxe__xml__Printer"',
            'GoDecl.GoFuncDecl("haxe__xml__Parser_parse"',
            'GoDecl.GoFuncDecl("haxe__xml__Printer_print"',
        ):
            self.assertNotIn(symbol, stdlib_emitter)
        self.assertNotIn('imports.push("encoding/xml")', compiler)

        stdlib_group = next(
            group for group in registry["groups"] if group["group"] == "stdlib_symbols"
        )
        owned_symbols = {entry["symbol"] for entry in stdlib_group["ownedSymbols"]}
        selected_paths = {
            entry["path"] for entry in stdlib_group["classifierSelections"]
        }
        for path in xml_paths:
            self.assertNotIn(path, owned_symbols)
            self.assertNotIn(path, selected_paths)

        for source_path in (
            "std/go/_std/Xml.hx",
            "std/go/_std/haxe/xml/Printer.hx",
        ):
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.15.2", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        generated_root = ROOT / "test/snapshot/stdlib/xml_root_dom_basic/intended"
        self.assertTrue((generated_root / "module_xml.go").is_file())
        self.assertTrue((generated_root / "module_haxe_xml_parser.go").is_file())
        self.assertTrue((generated_root / "module_haxe_xml_printer.go").is_file())
        self.assertNotIn('"encoding/xml"', (generated_root / "main.go").read_text())

    def test_zip_is_source_owned_over_a_typed_runtime_capability(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )

        zip_paths = ("haxe.zip.Compress", "haxe.zip.Uncompress")
        for path in zip_paths:
            self.assertNotIn(f'path: "{path}"', classifier)
            self.assertIn(f'"{path}"', planner)

        stdlib_start = compiler.index("function lowerStdlibSymbolShimDecls")
        stdlib_end = compiler.index("function reflectFieldsShimDecl", stdlib_start)
        stdlib_emitter = compiler[stdlib_start:stdlib_end]
        for symbol in (
            'GoDecl.GoStructDecl("haxe__zip__Compress"',
            'GoDecl.GoStructDecl("haxe__zip__Uncompress"',
            'GoDecl.GoFuncDecl("haxe__zip__Compress_run"',
            'GoDecl.GoFuncDecl("haxe__zip__Uncompress_run"',
        ):
            self.assertNotIn(symbol, stdlib_emitter)

        build_imports_start = compiler.index("function buildSupportImports")
        import_start = compiler.index(
            'if (requiredStdlibShimGroups.exists("stdlib_symbols"))',
            build_imports_start,
        )
        import_end = compiler.index(
            'if (requiredStdlibShimGroups.exists("go_result"))', import_start
        )
        stdlib_imports = compiler[import_start:import_end]
        for native_import in ("bytes", "compress/zlib", "io"):
            self.assertNotIn(f'imports.push("{native_import}")', stdlib_imports)

        stdlib_group = next(
            group for group in registry["groups"] if group["group"] == "stdlib_symbols"
        )
        owned_symbols = {entry["symbol"] for entry in stdlib_group["ownedSymbols"]}
        selected_paths = {
            entry["path"] for entry in stdlib_group["classifierSelections"]
        }
        for path in zip_paths:
            self.assertNotIn(path, owned_symbols)
            self.assertNotIn(path, selected_paths)

        for source_path in (
            "std/go/_std/haxe/zip/Compress.hx",
            "std/go/_std/haxe/zip/Uncompress.hx",
        ):
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.15.3", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        native_entry = ledger_entries.get("std/hxrt/zip/NativeZip.hx")
        self.assertIsNotNone(native_entry)
        self.assertEqual("hxrt_binding", native_entry.get("ownershipClass"))
        self.assertEqual("haxe_go-vfp.8.7.15.3", native_entry.get("migrationBead"))
        self.assertEqual([], native_entry.get("compilerShimGroups"))

        runtime_zip = (ROOT / "runtime/hxrt/zip.go").read_text(encoding="utf-8")
        self.assertIn("func ZipCompress(values []int, level int) []int", runtime_zip)
        self.assertIn(
            "func ZipUncompress(values []int, raw bool, bufferSize int) []int",
            runtime_zip,
        )
        self.assertNotIn("haxe__io__Bytes", runtime_zip)
        self.assertNotIn("reflect.", runtime_zip)
        self.assertNotIn("unsafe.", runtime_zip)

        self.assertIn('FEATURE_ZIP = "zip"', feature_analyzer)
        self.assertIn('path == "hxrt.zip.NativeZip"', feature_analyzer)
        self.assertIn('path == "haxe.zip.Compress"', feature_analyzer)
        self.assertIn('path == "haxe.zip.Uncompress"', feature_analyzer)
        self.assertIn("case FEATURE_ZIP:", feature_analyzer)
        self.assertIn('["zip.go"]', feature_analyzer)
        self.assertIn('case "zip.go":', reflaxe_compiler)

        generated_root = ROOT / "test/snapshot/stdlib/crypto_xml_zip_basic/intended"
        self.assertTrue((generated_root / "module_haxe_zip_compress.go").is_file())
        self.assertTrue((generated_root / "module_haxe_zip_uncompress.go").is_file())
        self.assertTrue((generated_root / "hxrt/zip.go").is_file())
        self.assertFalse(
            (ROOT / "test/snapshot/core/const_kinds_contract/intended/hxrt/zip.go").exists()
        )

    def test_date_and_math_are_source_owned_over_typed_native_capabilities(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")
        serializer_source = (
            ROOT / "std/go/_std/haxe/Serializer.hx"
        ).read_text(encoding="utf-8")
        serialization_runtime = (
            ROOT / "runtime/hxrt/serialization.go"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )

        for path in ("Date", "Math"):
            self.assertNotIn(f'path: "{path}"', classifier)
            self.assertIn(f'case "{path}"', planner)

        stdlib_start = compiler.index("function lowerStdlibSymbolShimDecls")
        stdlib_end = compiler.index("function reflectFieldsShimDecl", stdlib_start)
        stdlib_emitter = compiler[stdlib_start:stdlib_end]
        for symbol in (
            'GoDecl.GoStructDecl("Date"',
            'GoDecl.GoFuncDecl("Date_',
            'GoDecl.GoFuncDecl("getFullYear"',
            'GoDecl.GoStructDecl("Math"',
            'GoDecl.GoFuncDecl("Math_',
        ):
            self.assertNotIn(symbol, stdlib_emitter)

        build_imports_start = compiler.index("function buildSupportImports")
        import_start = compiler.index(
            'if (requiredStdlibShimGroups.exists("stdlib_symbols"))',
            build_imports_start,
        )
        import_end = compiler.index(
            'if (requiredStdlibShimGroups.exists("go_result"))', import_start
        )
        stdlib_imports = compiler[import_start:import_end]
        self.assertNotIn('imports.push("math")', stdlib_imports)
        self.assertNotIn('imports.push("time")', stdlib_imports)

        stdlib_group = next(
            group for group in registry["groups"] if group["group"] == "stdlib_symbols"
        )
        owned_symbols = {entry["symbol"] for entry in stdlib_group["ownedSymbols"]}
        selected_paths = {
            entry["path"] for entry in stdlib_group["classifierSelections"]
        }
        for path in ("Date", "Math"):
            self.assertNotIn(path, owned_symbols)
            self.assertNotIn(path, selected_paths)

        for source_path in ("std/go/_std/Date.hx", "std/go/_std/Math.hx"):
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.15.4", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        for source_path in (
            "std/hxrt/date/DateParts.hx",
            "std/hxrt/date/NativeDate.hx",
            "std/hxrt/math/NativeMath.hx",
            "std/hxrt/math/NativeMathInt.hx",
            "std/hxrt/math/NativeRandom.hx",
        ):
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual("hxrt_binding", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.15.4", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        runtime_date = (ROOT / "runtime/hxrt/date.go").read_text(encoding="utf-8")
        self.assertIn("func DateLocalTime(", runtime_date)
        self.assertIn("func DateLocalParts(", runtime_date)
        self.assertIn("func DateUTCParts(", runtime_date)
        self.assertIn("func DateTimezoneOffset(", runtime_date)
        self.assertNotIn("type Date struct", runtime_date)
        self.assertNotIn("reflect.", runtime_date)
        self.assertNotIn("unsafe.", runtime_date)

        native_math = (ROOT / "std/hxrt/math/NativeMath.hx").read_text(
            encoding="utf-8"
        )
        native_random = (ROOT / "std/hxrt/math/NativeRandom.hx").read_text(
            encoding="utf-8"
        )
        self.assertIn('@:go.import("math")', native_math)
        self.assertIn('@:go.import("math/rand")', native_random)
        runtime_math = (ROOT / "runtime/hxrt/math.go").read_text(encoding="utf-8")
        self.assertIn("func MathFloorInt(value float64) int", runtime_math)
        self.assertIn("func MathCeilInt(value float64) int", runtime_math)
        self.assertIn("func MathRoundInt(value float64) int", runtime_math)
        self.assertNotIn("func MathSin", runtime_math)
        self.assertNotIn("func MathSqrt", runtime_math)

        self.assertIn('FEATURE_DATE = "date"', feature_analyzer)
        self.assertIn('FEATURE_MATH = "math"', feature_analyzer)
        self.assertIn('path == "Date"', feature_analyzer)
        self.assertIn('path == "Math"', feature_analyzer)
        self.assertIn('path == "hxrt.date.NativeDate"', feature_analyzer)
        self.assertIn('path == "hxrt.math.NativeMathInt"', feature_analyzer)
        self.assertIn("case FEATURE_DATE:", feature_analyzer)
        self.assertIn("case FEATURE_MATH:", feature_analyzer)
        self.assertIn('["date.go"]', feature_analyzer)
        self.assertIn('["math.go"]', feature_analyzer)
        self.assertIn('case "date.go":', reflaxe_compiler)
        self.assertIn('case "math.go":', reflaxe_compiler)

        self.assertIn('case "Date"', serializer_source)
        self.assertIn("date.getTime()", serializer_source)
        self.assertNotIn('fieldType.PkgPath() != "time"', serialization_runtime)

        generated_root = ROOT / "test/snapshot/stdlib/date_math_source_owned/intended"
        self.assertTrue((generated_root / "module_date.go").is_file())
        self.assertTrue((generated_root / "module_math.go").is_file())
        self.assertTrue((generated_root / "hxrt/date.go").is_file())
        self.assertTrue((generated_root / "hxrt/math.go").is_file())
        self.assertFalse(
            (ROOT / "test/snapshot/core/const_kinds_contract/intended/hxrt/date.go").exists()
        )
        self.assertFalse(
            (ROOT / "test/snapshot/core/const_kinds_contract/intended/hxrt/math.go").exists()
        )
        self.assertFalse(
            (
                ROOT
                / "test/snapshot/stdlib/math_float_native_no_hxrt/intended/hxrt/math.go"
            ).exists()
        )

    def test_unicode_string_algorithms_are_source_owned(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )

        for path in ("UnicodeString", "_UnicodeString.UnicodeString_Impl_"):
            self.assertNotIn(f'path: "{path}"', classifier)
        self.assertIn(
            'case "UnicodeString", "_UnicodeString.UnicodeString_Impl_"', planner
        )
        self.assertIn(
            'requireSourceOwnedStdlibClass("haxe.iterators.StringIteratorUnicode")',
            planner,
        )
        self.assertIn(
            'requireSourceOwnedStdlibClass("haxe.iterators.StringKeyValueIteratorUnicode")',
            planner,
        )

        stdlib_start = compiler.index("function lowerStdlibSymbolShimDecls")
        stdlib_end = compiler.index("function reflectFieldsShimDecl", stdlib_start)
        stdlib_emitter = compiler[stdlib_start:stdlib_end]
        self.assertNotIn("_UnicodeString__UnicodeString_Impl__", stdlib_emitter)

        self.assertNotIn("function shouldSkipStaticDefaultArgPadding", compiler)

        stdlib_group = next(
            group for group in registry["groups"] if group["group"] == "stdlib_symbols"
        )
        owned_symbols = {entry["symbol"] for entry in stdlib_group["ownedSymbols"]}
        selected_paths = {
            entry["path"] for entry in stdlib_group["classifierSelections"]
        }
        self.assertNotIn("UnicodeString", owned_symbols)
        self.assertNotIn("UnicodeString", selected_paths)
        self.assertNotIn("_UnicodeString.UnicodeString_Impl_", selected_paths)

        source_path = "std/go/_std/UnicodeString.hx"
        entry = ledger_entries.get(source_path)
        self.assertIsNotNone(entry, source_path)
        self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
        self.assertEqual("haxe_go-vfp.8.7.15.5", entry.get("migrationBead"))
        self.assertEqual([], entry.get("compilerShimGroups"))

        staged_source = (ROOT / source_path).read_text(encoding="utf-8")
        self.assertIn(
            "The mainstream Haxe stdlib implementation cannot be used unchanged",
            staged_source,
        )
        for function_name in (
            "charAt",
            "charCodeAt",
            "indexOf",
            "lastIndexOf",
            "substr",
            "substring",
            "validate",
        ):
            self.assertIn(f"function {function_name}(", staged_source)
        self.assertNotIn("@:op(A += B)", staged_source)
        self.assertNotIn("GoInjection", staged_source)
        self.assertNotIn("__go__", staged_source)

        runtime_binding = (ROOT / "std/hxrt/string/GoStringRuntime.hx").read_text(
            encoding="utf-8"
        )
        runtime_string = (ROOT / "runtime/hxrt/string.go").read_text(
            encoding="utf-8"
        )
        self.assertIn("function length(value:String):Int", runtime_binding)
        self.assertIn("function sliceCodePoints(", runtime_binding)
        self.assertIn("func StringSliceCodePointsStringPtr(", runtime_string)

        generated_root = ROOT / "test/snapshot/stdlib/unicode_string_basic/intended"
        self.assertTrue((generated_root / "module_unicodestring.go").is_file())
        self.assertNotIn(
            "_UnicodeString__UnicodeString_Impl__get_length",
            (generated_root / "main.go").read_text(encoding="utf-8"),
        )
        self.assertFalse(
            (ROOT / "test/snapshot/core/const_kinds_contract/intended/module_unicodestring.go").exists()
        )

    def test_atomic_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_owners = {
            "std/go/_std/haxe/atomic/AtomicInt.hx": "upstream_std_override",
            "std/go/_std/haxe/atomic/AtomicObject.hx": "upstream_std_override",
            "std/hxrt/atomic/AtomicIntHandle.hx": "hxrt_binding",
            "std/hxrt/atomic/AtomicObjectHandle.hx": "hxrt_binding",
            "std/hxrt/atomic/NativeAtomicInt.hx": "hxrt_binding",
            "std/hxrt/atomic/NativeAtomicObject.hx": "hxrt_binding",
        }
        for source_path, expected_owner in expected_owners.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual("haxe_go-vfp.8.7.9", entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        self.assertFalse(
            (ROOT / "std/go/_std/haxe/atomic/AtomicBool.hx").exists(),
            "the mainstream source-owned AtomicBool already works unchanged over AtomicInt",
        )

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )
        debt_ratchet = (ROOT / "test/run-compiler-debt-ratchet.py").read_text(
            encoding="utf-8"
        )
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )
        atomic_int_runtime = (ROOT / "runtime/hxrt/atomic_int.go").read_text(
            encoding="utf-8"
        )
        atomic_object_runtime = (ROOT / "runtime/hxrt/atomic_object.go").read_text(
            encoding="utf-8"
        )
        atomic_int_snapshot = (
            ROOT / "test/snapshot/stdlib/atomic_int_bool_basic/intended/main.go"
        ).read_text(encoding="utf-8")
        atomic_object_snapshot = (
            ROOT / "test/snapshot/stdlib/atomic_object_basic/intended/main.go"
        ).read_text(encoding="utf-8")

        for fragment in (
            "lowerAtomicStdlibShimDecls",
            'requiredStdlibShimGroups.exists("atomic")',
            'requireStdlibShimGroup("atomic")',
            "haxe__atomic",
        ):
            self.assertNotIn(fragment, compiler, fragment)

        self.assertNotIn('groups: ["atomic"]', classifier)
        self.assertNotIn('StringTools.startsWith(path, "haxe.atomic.")', feature_analyzer)
        self.assertNotIn('case "atomic":', feature_analyzer)
        self.assertIn('path == "hxrt.atomic.NativeAtomicInt"', feature_analyzer)
        self.assertIn('path == "hxrt.atomic.NativeAtomicObject"', feature_analyzer)

        self.assertIn("func AtomicIntNew(value int) *AtomicIntCell", atomic_int_runtime)
        self.assertIn("func AtomicIntLoad(cell *AtomicIntCell) int", atomic_int_runtime)
        self.assertIn("func AtomicObjectNew(value any) *AtomicObjectCell", atomic_object_runtime)
        self.assertIn("func AtomicObjectLoad(cell *AtomicObjectCell) any", atomic_object_runtime)

        for snapshot in (atomic_int_snapshot, atomic_object_snapshot):
            self.assertNotIn("haxe__atomic", snapshot)
        self.assertIn("*hxrt.AtomicIntCell", atomic_int_snapshot)
        self.assertIn("hxrt.AtomicIntAdd", atomic_int_snapshot)
        self.assertIn("*hxrt.AtomicObjectCell", atomic_object_snapshot)
        self.assertIn("hxrt.AtomicObjectCompareExchange", atomic_object_snapshot)

        atomic_int_slice = ROOT / "test/snapshot/core/runtime_hxrt_infer_atomic_int/intended/hxrt"
        atomic_object_slice = ROOT / "test/snapshot/core/runtime_hxrt_infer_atomic_object/intended/hxrt"
        self.assertTrue((atomic_int_slice / "atomic_int.go").is_file())
        self.assertFalse((atomic_int_slice / "atomic_object.go").exists())
        self.assertFalse((atomic_int_slice / "equality.go").exists())
        self.assertTrue((atomic_object_slice / "atomic_object.go").is_file())
        self.assertTrue((atomic_object_slice / "equality.go").is_file())
        self.assertFalse((atomic_object_slice / "atomic_int.go").exists())

        self.assertFalse(
            any(group.get("group") == "atomic" for group in registry["groups"]),
            "source-owned haxe.atomic must not remain registered as a compiler shim group",
        )
        self.assertFalse(
            any(entry.get("decisionId") == "migration_atomic" for entry in registry["directLowerings"]),
            "source-owned haxe.atomic must not retain direct compiler lowering entries",
        )
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "atomic"
                for limit in debt_policy["limits"]
            ),
            "source-owned haxe.atomic must not retain a compiler-shim debt allowance",
        )
        self.assertNotIn('"lowerAtomicStdlibShimDecls": "atomic"', debt_ratchet)

    def test_ds_collections_are_source_owned_instead_of_a_compiler_shim(self) -> None:
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        expected_owners = {
            "std/go/_std/haxe/ds/EnumValueMap.hx": "upstream_std_override",
            "std/go/_std/haxe/ds/IntMap.hx": "upstream_std_override",
            "std/go/_std/haxe/ds/List.hx": "upstream_std_override",
            "std/go/_std/haxe/ds/ObjectMap.hx": "upstream_std_override",
            "std/go/_std/haxe/ds/StringMap.hx": "upstream_std_override",
            "std/go/_std/haxe/iterators/MapKeyValueIterator.hx": "upstream_std_override",
            "std/hxrt/collections/IntMapHandle.hx": "hxrt_binding",
            "std/hxrt/collections/NativeEnumValue.hx": "hxrt_binding",
            "std/hxrt/collections/NativeIntMap.hx": "hxrt_binding",
            "std/hxrt/collections/NativeObjectMap.hx": "hxrt_binding",
            "std/hxrt/collections/NativeStringMap.hx": "hxrt_binding",
            "std/hxrt/collections/ObjectMapHandle.hx": "hxrt_binding",
            "std/hxrt/collections/StringMapHandle.hx": "hxrt_binding",
        }
        for source_path, expected_owner in expected_owners.items():
            self.assertTrue((ROOT / source_path).is_file(), source_path)
            entry = ledger_entries.get(source_path)
            self.assertIsNotNone(entry, source_path)
            self.assertEqual(expected_owner, entry.get("ownershipClass"), source_path)
            self.assertEqual("haxe_go-vfp.8.7.10", entry.get("migrationBead"), source_path)
            self.assertEqual([], entry.get("compilerShimGroups"), source_path)

        for source_path in (
            "std/go/_std/haxe/ds/EnumValueMap.hx",
            "std/go/_std/haxe/ds/IntMap.hx",
            "std/go/_std/haxe/ds/List.hx",
            "std/go/_std/haxe/ds/ObjectMap.hx",
            "std/go/_std/haxe/ds/StringMap.hx",
            "std/go/_std/haxe/iterators/MapKeyValueIterator.hx",
        ):
            content = (ROOT / source_path).read_text(encoding="utf-8")
            self.assertNotIn("extern class", content, source_path)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx").read_text(
            encoding="utf-8"
        )
        feature_analyzer = (ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx").read_text(
            encoding="utf-8"
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )
        debt_ratchet = (ROOT / "test/run-compiler-debt-ratchet.py").read_text(
            encoding="utf-8"
        )
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(encoding="utf-8")
        )

        for fragment in (
            "lowerDsStdlibShimDecls",
            'requiredStdlibShimGroups.exists("ds")',
            'requireStdlibShimGroup("ds")',
            'group == "ds"',
        ):
            self.assertNotIn(fragment, compiler, fragment)

        self.assertNotIn('groups: ["ds"]', classifier)
        self.assertIn('path == "hxrt.collections.NativeIntMap"', feature_analyzer)
        self.assertIn('path == "hxrt.collections.NativeStringMap"', feature_analyzer)
        self.assertIn('path == "hxrt.collections.NativeObjectMap"', feature_analyzer)
        self.assertIn('path == "hxrt.collections.NativeEnumValue"', feature_analyzer)

        selective_cases = {
            "runtime_hxrt_infer_enum_value": "enum_value.go",
            "runtime_hxrt_infer_map_int": "map_int.go",
            "runtime_hxrt_infer_map_object": "map_object.go",
            "runtime_hxrt_infer_map_string": "map_string.go",
        }
        collection_runtime_files = set(selective_cases.values())
        for case_name, selected_file in selective_cases.items():
            runtime_slice = ROOT / "test/snapshot/core" / case_name / "intended/hxrt"
            self.assertTrue((runtime_slice / selected_file).is_file(), case_name)
            for excluded_file in collection_runtime_files - {selected_file}:
                self.assertFalse((runtime_slice / excluded_file).exists(), case_name)

        self.assertFalse(
            any(group.get("group") == "ds" for group in registry["groups"]),
            "source-owned haxe.ds collections must not remain a compiler shim group",
        )
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "ds"
                for limit in debt_policy["limits"]
            ),
            "source-owned haxe.ds collections must not retain a compiler-shim debt allowance",
        )
        self.assertNotIn('"lowerDsStdlibShimDecls": "ds"', debt_ratchet)

    def test_ambiguities_cannot_exist_without_a_follow_up_bead(self) -> None:
        ledger = load_ledger()
        ambiguities = ledger.get("migrationContract", {}).get("ambiguities")
        self.assertIsInstance(ambiguities, list)
        paths = {entry["path"] for entry in ledger["entries"]}
        seen: set[str] = set()
        for ambiguity in ambiguities:
            path = ambiguity.get("path")
            follow_up = ambiguity.get("followUpBead")
            self.assertIn(path, paths)
            self.assertNotIn(path, seen)
            self.assertRegex(follow_up or "", r"^haxe_go-[a-z0-9.-]+$")
            seen.add(path)

    def test_regex_and_serialization_algorithms_are_source_owned(self) -> None:
        """Regex and token-stream policy must not return to a mixed emitter."""
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(
            encoding="utf-8"
        )
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        ownership = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibOwnership.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(
                encoding="utf-8"
            )
        )
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}

        staged_paths = (
            "std/go/_std/EReg.hx",
            "std/go/_std/haxe/Serializer.hx",
            "std/go/_std/haxe/Unserializer.hx",
        )
        for staged_path in staged_paths:
            source = (ROOT / staged_path).read_text(encoding="utf-8")
            self.assertNotIn("extern class", source, staged_path)
            self.assertNotIn("__go__", source, staged_path)
            for heading in ("What:", "Why:", "How:"):
                self.assertIn(heading, source, staged_path)

            entry = ledger_entries.get(staged_path)
            self.assertIsNotNone(entry, staged_path)
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.13", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

        support_path = "std/haxe/GoSerializationBridge.hx"
        support_source = (ROOT / support_path).read_text(encoding="utf-8")
        support_entry = ledger_entries.get(support_path)
        self.assertIsNotNone(support_entry, support_path)
        self.assertEqual("staged_support", support_entry.get("ownershipClass"))
        self.assertEqual("haxe_go-vfp.8.7.13", support_entry.get("migrationBead"))
        self.assertEqual([], support_entry.get("compilerShimGroups"))
        for heading in ("What:", "Why:", "How:"):
            self.assertIn(heading, support_source, support_path)

        runtime_bindings = (
            "std/hxrt/regex/NativeRegex.hx",
            "std/hxrt/regex/RegexHandle.hx",
            "std/hxrt/regex/RegexMatch.hx",
            "std/hxrt/serialization/NativeSerialization.hx",
            "std/hxrt/serialization/SerializationField.hx",
        )
        for binding_path in runtime_bindings:
            binding_source = (ROOT / binding_path).read_text(encoding="utf-8")
            binding_entry = ledger_entries.get(binding_path)
            self.assertIsNotNone(binding_entry, binding_path)
            self.assertEqual("hxrt_binding", binding_entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.13", binding_entry.get("migrationBead"))
            self.assertEqual([], binding_entry.get("compilerShimGroups"))
            for heading in ("What:", "Why:", "How:"):
                self.assertIn(heading, binding_source, binding_path)

        self.assertFalse(
            (ROOT / "src/reflaxe/go/compiler/emit/GoRegexSerializerEmitter.hx").exists()
        )
        for fragment in (
            "GoRegexSerializerEmitter",
            "lowerRegexSerializerShimDecls",
            'requiredStdlibShimGroups.exists("regex_serializer")',
            'requireStdlibShimGroup("regex_serializer")',
        ):
            self.assertNotIn(fragment, compiler + planner, fragment)
        self.assertNotIn('groups: ["regex_serializer"]', classifier)
        self.assertNotIn('"EReg"', ownership)
        self.assertIn("lowerSerializationSourceBridgeShimDecls", compiler)

        for source_path in ("EReg", "haxe.Serializer", "haxe.Unserializer"):
            self.assertIn(f'case "{source_path}"', planner)

        self.assertFalse(
            any(group.get("group") == "regex_serializer" for group in registry["groups"]),
            "source-owned regex and serialization must not retain a compiler shim group",
        )
        self.assertNotIn("migration_regex_serializer", registry["decisions"])
        self.assertIn("approved_serialization_source_bridge", registry["decisions"])
        support_entries = [
            support
            for group in registry["groups"]
            for support in group.get("supportEntryPoints", [])
        ]
        self.assertIn(
            {
                "context": "lowerSerializationSourceBridgeShimDecls",
                "status": "approved_intrinsic",
                "decisionId": "approved_serialization_source_bridge",
            },
            support_entries,
        )

        self.assertIn('FEATURE_REGEX = "regex"', feature_analyzer)
        self.assertIn('FEATURE_SERIALIZATION = "serialization"', feature_analyzer)
        self.assertIn('case FEATURE_REGEX:\n\t\t\t\t["regex.go"]', feature_analyzer)
        self.assertIn(
            'case FEATURE_SERIALIZATION:\n\t\t\t\t["serialization.go"]',
            feature_analyzer,
        )
        self.assertIn('case "regex.go":', reflaxe_compiler)
        self.assertIn('case "serialization.go":', reflaxe_compiler)

        regex_runtime = (
            ROOT / "test/snapshot/core/runtime_hxrt_infer_regex/intended/hxrt"
        )
        serialization_runtime = (
            ROOT
            / "test/snapshot/core/runtime_hxrt_infer_serialization/intended/hxrt"
        )
        self.assertTrue((regex_runtime / "regex.go").is_file())
        self.assertFalse((regex_runtime / "serialization.go").exists())
        self.assertTrue((serialization_runtime / "serialization.go").is_file())
        self.assertFalse((serialization_runtime / "regex.go").exists())

    def test_base_haxe_io_is_source_owned_instead_of_a_compiler_shim(self) -> None:
        """Public byte and stream policy must compile from canonical Haxe source."""
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(
            encoding="utf-8"
        )
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        ownership = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibOwnership.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(
                encoding="utf-8"
            )
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )
        debt_ratchet = (ROOT / "test/run-compiler-debt-ratchet.py").read_text(
            encoding="utf-8"
        )
        bytes_runtime = (ROOT / "runtime/hxrt/bytes.go").read_text(
            encoding="utf-8"
        )
        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}

        staged_types = (
            "BufferInput",
            "Bytes",
            "BytesBuffer",
            "BytesInput",
            "BytesOutput",
            "Encoding",
            "Eof",
            "Error",
            "Input",
            "Output",
            "StringInput",
        )
        for type_name in staged_types:
            staged_path = f"std/go/_std/haxe/io/{type_name}.hx"
            source_path = ROOT / staged_path
            self.assertTrue(source_path.is_file(), staged_path)
            if not source_path.is_file():
                continue
            source = source_path.read_text(encoding="utf-8")
            self.assertNotIn("extern class", source, staged_path)
            self.assertNotIn("__go__", source, staged_path)
            for heading in ("What:", "Why:", "How:"):
                self.assertIn(heading, source, staged_path)

            entry = ledger_entries.get(staged_path)
            self.assertIsNotNone(entry, staged_path)
            if entry is None:
                continue
            self.assertEqual("upstream_std_override", entry.get("ownershipClass"))
            self.assertEqual("haxe_go-vfp.8.7.11", entry.get("migrationBead"))
            self.assertEqual([], entry.get("compilerShimGroups"))

            self.assertIn(f'case "haxe.io.{type_name}"', planner, type_name)

        self.assertFalse((ROOT / "std/haxe/io/GoIoHelpers.hx").exists())
        for fragment in (
            "lowerIoStdlibShimDecls",
            'requiredStdlibShimGroups.exists("io")',
            'requireStdlibShimGroup("io")',
        ):
            self.assertNotIn(fragment, compiler, fragment)
        self.assertNotIn('groups: ["io"]', classifier)
        for type_name in staged_types:
            self.assertNotIn(f'"haxe.io.{type_name}"', ownership, type_name)

        self.assertFalse(
            any(group.get("group") == "io" for group in registry["groups"]),
            "source-owned haxe.io must not retain a compiler shim group",
        )
        self.assertNotIn(
            "migration_required_haxe_go_vfp_8_7_11", registry["decisions"]
        )
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "io"
                for limit in debt_policy["limits"]
            ),
            "source-owned haxe.io must not retain a compiler-shim debt allowance",
        )
        self.assertNotIn('"lowerIoStdlibShimDecls": "io"', debt_ratchet)
        for legacy_helper in (
            "BytesFromString",
            "BytesToString",
            "BytesOfHex",
            "BytesToHex",
            "BytesBufferLength",
        ):
            self.assertNotIn(
                f"func {legacy_helper}(",
                bytes_runtime,
                f"staged haxe.io owns the former {legacy_helper} policy",
            )

    def test_sys_http_is_staged_policy_over_a_typed_runtime_boundary(self) -> None:
        """Request choreography belongs in Haxe; hxrt owns only native transport."""
        staged_path = ROOT / "std/go/_std/sys/Http.hx"
        self.assertTrue(staged_path.is_file(), "sys.Http must be canonical staged source")
        if not staged_path.is_file():
            return

        staged = staged_path.read_text(encoding="utf-8")
        native_http_path = ROOT / "std/hxrt/http/NativeHttp.hx"
        request_handle_path = ROOT / "std/hxrt/http/HttpRequestHandle.hx"
        response_handle_path = ROOT / "std/hxrt/http/HttpResponseHandle.hx"
        runtime_path = ROOT / "runtime/hxrt/http.go"
        runtime_test_path = ROOT / "runtime/hxrt/http_test.go"
        for path in (
            native_http_path,
            request_handle_path,
            response_handle_path,
            runtime_path,
            runtime_test_path,
        ):
            self.assertTrue(path.is_file(), str(path.relative_to(ROOT)))

        for forbidden in ("__go__", "GoInjection", "@:goAllowRaw"):
            self.assertNotIn(forbidden, staged, forbidden)
        for heading in ("What", "Why", "How"):
            self.assertIn(heading, staged, heading)
        self.assertIn("class Http extends haxe.http.HttpBase", staged)
        self.assertIn("NativeHttp.execute", staged)
        self.assertIn("onStatus", staged)
        self.assertIn("onError", staged)
        self.assertIn("getResponseHeaderValues", staged)
        self.assertIn("NativeHttp.setSocket(request, sock.handle)", staged)

        socket_source = (
            ROOT / "std/go/_std/sys/net/Socket.hx"
        ).read_text(encoding="utf-8")
        self.assertIn("@:allow(sys.Http)", socket_source)
        self.assertNotIn("__hx_httpHandle", socket_source + staged)

        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(
            encoding="utf-8"
        )
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        ownership = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibOwnership.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        registry = json.loads(
            (ROOT / "docs/compiler-stdlib-intrinsics.json").read_text(
                encoding="utf-8"
            )
        )
        debt_policy = json.loads(
            (ROOT / "test/compiler_debt_policy.json").read_text(encoding="utf-8")
        )
        debt_ratchet = (ROOT / "test/run-compiler-debt-ratchet.py").read_text(
            encoding="utf-8"
        )

        self.assertFalse((ROOT / "std/sys/GoHttpHelpers.hx").exists())
        for fragment in (
            "lowerHttpStdlibShimDecls",
            'requiredStdlibShimGroups.exists("http")',
            'requireStdlibShimGroup("http")',
            'requireSourceOwnedStdlibClass("sys.GoHttpHelpers")',
        ):
            self.assertNotIn(fragment, compiler, fragment)
        self.assertNotIn('{kind: "class", path: "sys.Http", groups: ["http"]}', classifier)
        self.assertNotIn('case "sys.Http":', ownership)
        self.assertIn('case "sys.Http":', planner)
        self.assertIn('requireSourceOwnedStdlibClass("sys.Http")', planner)
        self.assertFalse(
            any(group.get("group") == "http" for group in registry["groups"]),
            "source-owned sys.Http must not retain a compiler shim group",
        )
        self.assertNotIn("migration_http", registry["decisions"])
        self.assertFalse(
            any(
                limit.get("metric") == "compiler_shim"
                and limit.get("capability") == "http"
                for limit in debt_policy["limits"]
            )
        )
        self.assertNotIn('"lowerHttpStdlibShimDecls": "http"', debt_ratchet)
        self.assertIn('FEATURE_HTTP = "http"', feature_analyzer)
        self.assertIn('case FEATURE_HTTP:\n\t\t\t\t["http.go"]', feature_analyzer)

        if native_http_path.is_file():
            native_http = native_http_path.read_text(encoding="utf-8")
            for forbidden in ("Dynamic", "Any", "__go__"):
                self.assertNotIn(forbidden, native_http, forbidden)
            self.assertIn("execute(request:HttpRequestHandle):HttpResponseHandle", native_http)

        if runtime_path.is_file():
            runtime = runtime_path.read_text(encoding="utf-8")
            for generated_layout in ("sys__Http", "haxe__io__Bytes"):
                self.assertNotIn(generated_layout, runtime, generated_layout)

        inferred_runtime = (
            ROOT / "test/snapshot/core/runtime_hxrt_infer_http/intended/hxrt"
        )
        for selected_file in ("http.go", "bytes.go", "socket.go", "string.go"):
            self.assertTrue((inferred_runtime / selected_file).is_file(), selected_file)
        for unrelated_file in ("filesystem.go", "process.go", "ssl.go"):
            self.assertFalse((inferred_runtime / unrelated_file).exists(), unrelated_file)
        self.assertFalse(
            (
                ROOT
                / "test/snapshot/core/runtime_hxrt_infer_json/intended/hxrt/http.go"
            ).exists(),
            "unrelated typed runtime use must not copy the HTTP capability",
        )

        callback_shape = (
            ROOT
            / "test/snapshot/sys/http_request_callbacks_smoke/intended/module_haxe_http_httpbase.go"
        ).read_text(encoding="utf-8")
        self.assertRegex(callback_shape, r"(?m)^\s+onData\s+func\(\*string\)$")
        self.assertNotRegex(callback_shape, r"(?m)^\s+onData\(data \*string\)$")
        self.assertFalse(
            (
                ROOT
                / "test/snapshot/sys/http_request_callbacks_smoke/intended/module_sys_gohttphelpers.go"
            ).exists()
        )

        ledger_entries = {entry["path"]: entry for entry in load_ledger()["entries"]}
        staged_entry = ledger_entries.get("std/go/_std/sys/Http.hx")
        self.assertIsNotNone(staged_entry)
        if staged_entry is not None:
            self.assertEqual("upstream_std_override", staged_entry["ownershipClass"])
            self.assertEqual("haxe_go-vfp.8.7.12", staged_entry["migrationBead"])
            self.assertEqual([], staged_entry["compilerShimGroups"])

        for binding in (
            "std/hxrt/http/NativeHttp.hx",
            "std/hxrt/http/HttpRequestHandle.hx",
            "std/hxrt/http/HttpResponseHandle.hx",
        ):
            entry = ledger_entries.get(binding)
            self.assertIsNotNone(entry, binding)
            if entry is not None:
                self.assertEqual("hxrt_binding", entry["ownershipClass"])
                self.assertEqual("haxe_go-vfp.8.7.12", entry["migrationBead"])
                self.assertEqual([], entry["compilerShimGroups"])

    def test_bytes_native_view_is_shared_with_crypto_without_layout_leak(self) -> None:
        """Byte consumers reuse the opaque cache rather than copying generated fields."""
        native_crypto = (ROOT / "std/hxrt/crypto/NativeCrypto.hx").read_text(
            encoding="utf-8"
        )
        crypto_runtime = (ROOT / "runtime/hxrt/crypto.go").read_text(
            encoding="utf-8"
        )
        self.assertIn("import hxrt.io.ByteView;", native_crypto)
        self.assertNotIn("go.NativeSlice", native_crypto)
        for method in (
            "md5Values",
            "sha1Values",
            "sha224Values",
            "sha256Values",
        ):
            self.assertIn(f"{method}(values:ByteView)", native_crypto, method)
        self.assertIn(
            "base64Encode(values:ByteView, urlSafe:Bool):String", native_crypto
        )
        self.assertIn("base64Decode(value:String, urlSafe:Bool):ByteView", native_crypto)

        for function in (
            "CryptoBase64Encode(values *ByteView",
            "CryptoBase64Decode(value *string, urlSafe bool) *ByteView",
            "CryptoMd5Values(values *ByteView) *ByteView",
            "CryptoSha1Values(values *ByteView) *ByteView",
            "CryptoSha224Values(values *ByteView) *ByteView",
            "CryptoSha256Values(values *ByteView) *ByteView",
        ):
            self.assertIn(function, crypto_runtime, function)
        self.assertNotIn("cryptoValuesToBytes", crypto_runtime)
        self.assertNotIn("cryptoBytesToValues", crypto_runtime)

        for module_name in ("Base64", "Md5", "Sha1", "Sha224", "Sha256"):
            source = (
                ROOT / f"std/go/_std/haxe/crypto/{module_name}.hx"
            ).read_text(encoding="utf-8")
            self.assertIn("__hx_nativeView()", source, module_name)
            self.assertIn("__hx_fromNativeView", source, module_name)
            self.assertNotIn("NativeSlice", source, module_name)
            self.assertNotIn("toValues", source, module_name)
            self.assertNotIn("fromValues", source, module_name)

    def test_socket_public_api_is_staged_over_typed_runtime_handles(self) -> None:
        """Socket lifecycle must not return to compiler-owned declaration shims."""
        compiler = (ROOT / "src/reflaxe/go/GoCompiler.hx").read_text(encoding="utf-8")
        classifier = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibShimClassifier.hx"
        ).read_text(encoding="utf-8")
        ownership = (
            ROOT / "src/reflaxe/go/compiler/GoStdlibOwnership.hx"
        ).read_text(encoding="utf-8")
        planner = (
            ROOT / "src/reflaxe/go/compiler/GoSourceOwnedStdlibPlanner.hx"
        ).read_text(encoding="utf-8")
        feature_analyzer = (
            ROOT / "src/reflaxe/go/compiler/GoHxrtFeatureAnalyzer.hx"
        ).read_text(encoding="utf-8")
        reflaxe_compiler = (
            ROOT / "src/reflaxe/go/GoReflaxeCompiler.hx"
        ).read_text(encoding="utf-8")

        for staged_path in (
            "std/go/_std/sys/net/Host.hx",
            "std/go/_std/sys/net/Socket.hx",
            "std/go/_std/sys/net/UdpSocket.hx",
            "std/sys/net/_SocketIO.hx",
        ):
            source = (ROOT / staged_path).read_text(encoding="utf-8")
            self.assertNotIn("extern class", source, staged_path)
            self.assertNotIn("__go__", source, staged_path)

        for binding_path in (
            "std/hxrt/net/SocketHandle.hx",
            "std/hxrt/net/SocketAddress.hx",
            "std/hxrt/net/SocketAcceptResult.hx",
            "std/hxrt/net/SocketIOResult.hx",
            "std/hxrt/net/SocketDatagramResult.hx",
            "std/hxrt/net/SocketSelectResult.hx",
            "std/hxrt/net/NativeSocket.hx",
        ):
            binding = (ROOT / binding_path).read_text(encoding="utf-8")
            self.assertIn('@:go.import("hxrt")', binding, binding_path)

        self.assertTrue((ROOT / "runtime/hxrt/socket.go").is_file())
        self.assertTrue((ROOT / "runtime/hxrt/socket_test.go").is_file())
        self.assertTrue((ROOT / "runtime/hxrt/socket_ssl.go").is_file())
        self.assertFalse(
            (ROOT / "src/reflaxe/go/compiler/emit/GoNetSocketEmitter.hx").exists()
        )

        for fragment in (
            "GoNetSocketEmitter",
            "lowerNetSocketShimDecls",
            'requiredStdlibShimGroups.exists("net_socket")',
            'requireStdlibShimGroup("net_socket")',
        ):
            self.assertNotIn(fragment, compiler + planner, fragment)
        self.assertNotIn('groups: ["net_socket"]', classifier)
        for authority in ("sys.net.Host", "sys.net.Socket", "sys.net.UdpSocket"):
            self.assertNotIn(f'"{authority}"', ownership)

        self.assertIn('FEATURE_SOCKET = "socket"', feature_analyzer)
        self.assertIn('FEATURE_SOCKET_SSL = "socket_ssl"', feature_analyzer)
        for runtime_file in (
            "socket.go",
            "socket_broadcast_posix.go",
            "socket_broadcast_unsupported.go",
            "socket_broadcast_windows.go",
        ):
            self.assertIn(f'"{runtime_file}"', feature_analyzer)
        self.assertIn('["socket_ssl.go"]', feature_analyzer)
        self.assertIn(
            'case "socket.go", "socket_broadcast_posix.go", "socket_broadcast_unsupported.go",',
            reflaxe_compiler,
        )
        self.assertIn('"socket_broadcast_windows.go":', reflaxe_compiler)
        self.assertIn('case "socket_ssl.go":', reflaxe_compiler)

        socket_runtime = ROOT / "test/snapshot/core/runtime_hxrt_infer_socket/intended/hxrt"
        ssl_runtime = ROOT / "test/snapshot/core/runtime_hxrt_infer_ssl/intended/hxrt"
        socket_ssl_runtime = ROOT / "test/snapshot/core/runtime_hxrt_infer_socket_ssl/intended/hxrt"
        self.assertTrue((socket_runtime / "socket.go").is_file())
        for runtime_file in (
            "socket_broadcast_posix.go",
            "socket_broadcast_unsupported.go",
            "socket_broadcast_windows.go",
        ):
            self.assertTrue((socket_runtime / runtime_file).is_file(), runtime_file)
        self.assertFalse((socket_runtime / "ssl.go").exists())
        self.assertFalse((socket_runtime / "socket_ssl.go").exists())
        self.assertTrue((ssl_runtime / "ssl.go").is_file())
        self.assertFalse((ssl_runtime / "socket.go").exists())
        self.assertFalse((ssl_runtime / "socket_ssl.go").exists())
        for runtime_file in ("ssl.go", "socket.go", "socket_ssl.go"):
            self.assertTrue((socket_ssl_runtime / runtime_file).is_file(), runtime_file)
        for runtime_file in (
            "socket.go",
            "socket_broadcast_posix.go",
            "socket_broadcast_unsupported.go",
            "socket_broadcast_windows.go",
            "socket_ssl.go",
        ):
            self.assertFalse(
                (
                    ROOT
                    / "test/snapshot/core/const_kinds_contract/intended/hxrt"
                    / runtime_file
                ).exists(),
                runtime_file,
            )

    def test_migration_and_governance_entrypoints_consume_the_ledger(self) -> None:
        package = json.loads(PACKAGE_PATH.read_text(encoding="utf-8"))
        scripts = package["scripts"]
        contract_command = "python3 test/test_stdlib_migration_ledger_contract.py"
        self.assertEqual(contract_command, scripts.get("test:stdlib:migration-ledger"))
        self.assertIn("npm run test:stdlib:provenance", scripts.get("test:stdlib:governance", ""))
        self.assertIn("npm run test:stdlib:migration-ledger", scripts.get("test:stdlib:governance", ""))
        self.assertIn("npm run test:stdlib:intrinsics", scripts.get("test:stdlib:governance", ""))
        self.assertIn("test/test_stdlib_migration_ledger_contract.py", RELEASE_RUNNER_PATH.read_text(encoding="utf-8"))
        self.assertIn("stdlib-provenance-ledger.json", CHECKER_PATH.read_text(encoding="utf-8"))


if __name__ == "__main__":
    unittest.main()
