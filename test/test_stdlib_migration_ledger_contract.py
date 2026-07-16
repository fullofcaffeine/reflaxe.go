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
    "std/haxe/io/GoIoHelpers.cross.hx": ("staged_support", "std/haxe/io/GoIoHelpers.hx"),
    "std/sys/GoHttpHelpers.cross.hx": ("staged_support", "std/sys/GoHttpHelpers.hx"),
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
    "std/_std/hxrt/stack/NativeStack.hx": ("hxrt_binding", "std/hxrt/stack/NativeStack.hx"),
    "std/_std/hxrt/stack/NativeStackFrame.hx": (
        "hxrt_binding",
        "std/hxrt/stack/NativeStackFrame.hx",
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
    "std/haxe/Constraints.cross.hx": ["ds"],
    "std/haxe/Template.cross.hx": ["stdlib_symbols", "template_support"],
    "std/haxe/ds/BalancedTree.cross.hx": ["stdlib_symbols"],
    "std/haxe/io/FPHelper.cross.hx": ["stdlib_symbols"],
    "std/haxe/io/GoIoHelpers.cross.hx": ["io"],
    "std/sys/GoHttpHelpers.cross.hx": ["http"],
    "std/sys/ssl/Certificate.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Digest.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Key.cross.hx": ["stdlib_symbols"],
    "std/sys/ssl/Socket.cross.hx": ["net_socket", "stdlib_symbols"],
    "std/_std/haxe/ds/IntMap.cross.hx": ["ds"],
    "std/_std/haxe/ds/ObjectMap.cross.hx": ["ds"],
    "std/_std/haxe/ds/StringMap.cross.hx": ["ds"],
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
    "ds": "migration_required_haxe_go_vfp_8_7_10",
    "http": "migration_required_haxe_go_vfp_8_7_12",
    "io": "migration_required_haxe_go_vfp_8_7_11",
    "net_socket": "migration_required_haxe_go_vfp_8_7_14",
    "stdlib_symbols": "split_migration_debt_from_exact_intrinsics_haxe_go_vfp_8_7_15",
    "template_support": "migration_required_haxe_go_vfp_8_7_16",
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
        self.assertTrue((atomic_object_slice / "atomic_object.go").is_file())
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
