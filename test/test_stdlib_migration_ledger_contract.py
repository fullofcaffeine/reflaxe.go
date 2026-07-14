#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
import subprocess
import unittest


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
    def test_every_tracked_std_source_has_one_resolved_owner_and_destination(self) -> None:
        ledger = load_ledger()
        self.assertEqual(2, ledger.get("schemaVersion"))
        entries = ledger.get("entries")
        self.assertIsInstance(entries, list)

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

            expected_bead = (
                "haxe_go-vfp.5.3" if ownership == "upstream_std_override" else "haxe_go-vfp.5.4"
            )
            self.assertEqual(expected_bead, entry.get("migrationBead"), source_path)

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
            {"ds", "http", "io", "net_socket", "stdlib_symbols", "template_support"},
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
            self.assertIsInstance(audit.get("decision"), str, group)
            self.assertTrue(audit["decision"], group)
            references = audit.get("references")
            self.assertIsInstance(references, list, group)
            self.assertTrue(references, group)
            for reference in references:
                self.assertTrue((ROOT / reference).is_file(), f"{group}: missing audit reference {reference}")

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
        self.assertIn("test/test_stdlib_migration_ledger_contract.py", RELEASE_RUNNER_PATH.read_text(encoding="utf-8"))
        self.assertIn("stdlib-provenance-ledger.json", CHECKER_PATH.read_text(encoding="utf-8"))


if __name__ == "__main__":
    unittest.main()
