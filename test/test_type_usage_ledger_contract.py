#!/usr/bin/env python3

from __future__ import annotations

import json
import shutil
import subprocess
import unittest
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parent.parent
FIXTURE = ROOT / "test" / "snapshot" / "core" / "type_usage_ledger"
OUT = FIXTURE / "out"
COMPILER_INIT = ROOT / "src" / "reflaxe" / "go" / "CompilerInit.hx"
COMPILATION_CONTEXT = ROOT / "src" / "reflaxe" / "go" / "CompilationContext.hx"
REFLAXE_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx"
LEDGER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoTypeUsageLedger.hx"


def flatten_strings(value: Any) -> list[str]:
    if isinstance(value, str):
        return [value]
    if isinstance(value, list):
        return [item for entry in value for item in flatten_strings(entry)]
    if isinstance(value, dict):
        return [item for entry in value.values() for item in flatten_strings(entry)]
    return []


class TypeUsageLedgerContractTest(unittest.TestCase):
    def compile_fixture(self) -> dict[str, Any]:
        shutil.rmtree(OUT, ignore_errors=True)
        proc = subprocess.run(
            ["haxe", "compile.hxml"],
            cwd=FIXTURE,
            capture_output=True,
            text=True,
            check=False,
        )
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)

        report = OUT / "type_usage.json"
        self.assertTrue(report.is_file(), "typed usage report was not emitted")
        return json.loads(report.read_text(encoding="utf-8"))

    def test_tracker_is_enabled_and_report_is_deterministic(self) -> None:
        compiler_init = COMPILER_INIT.read_text(encoding="utf-8")
        self.assertIn("trackUsedTypes: true", compiler_init)

        first = self.compile_fixture()
        first_bytes = (OUT / "type_usage.json").read_bytes()
        second = self.compile_fixture()
        second_bytes = (OUT / "type_usage.json").read_bytes()

        self.assertEqual(first, second)
        self.assertEqual(first_bytes, second_bytes)
        self.assertTrue((OUT / "type_usage.md").is_file())

    def test_context_publishes_one_deeply_read_only_snapshot(self) -> None:
        context_source = COMPILATION_CONTEXT.read_text(encoding="utf-8")
        compiler_source = REFLAXE_COMPILER.read_text(encoding="utf-8")
        ledger_source = LEDGER.read_text(encoding="utf-8")

        self.assertIn(
            "public final typedUsageLedger:GoTypeUsageLedgerSnapshot",
            context_source,
        )
        self.assertNotIn("context.typedUsageLedger =", compiler_source)
        self.assertEqual(
            compiler_source.count("CompilationContext.fromBuildContext("),
            1,
        )
        self.assertIn("GoImmutableList<GoTypeUsageModuleEvidence>", ledger_source)
        self.assertIn("GoImmutableList<GoTypeShape>", ledger_source)
        self.assertIn("enum GoTypeShape", ledger_source)
        self.assertIn(
            "final usageKind:GoNativeImportUsageKind",
            ledger_source,
        )
        self.assertNotIn("final modules:Array<GoTypeUsageModuleEvidence>", ledger_source)

    def test_report_preserves_shapes_dce_owners_and_native_imports(self) -> None:
        report = self.compile_fixture()

        self.assertEqual(report["schemaVersion"], 1)
        self.assertEqual(report["source"], "reflaxe_type_usage_tracker")
        self.assertEqual(
            report["scannerFallback"],
            "transitional_contract_diagnostics_only",
        )

        modules = report["modules"]
        module_names = [entry["module"] for entry in modules]
        self.assertEqual(module_names, sorted(module_names))
        self.assertIn("Main", module_names)
        self.assertIn("UsedBox", module_names)
        self.assertIn("UsedBox.UsedSibling", module_names)
        self.assertNotIn("DeadBox", module_names)
        self.assertFalse(any("DeadBox" in value for value in flatten_strings(report)))
        self.assertFalse((OUT / "module_deadbox.go").exists())

        main = next(entry for entry in modules if entry["module"] == "Main")
        self.assertTrue(
            any(
                entry["level"] == "constructed"
                and entry["shape"]["kind"] == "class"
                and entry["shape"]["path"] == "UsedBox"
                and [
                    parameter["path"]
                    for parameter in entry["shape"]["parameters"]
                ]
                == ["StdTypes.Int"]
                for entry in main["typeUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["shape"]["kind"] == "class"
                and entry["shape"]["path"] == "UsedBox"
                and len(entry["shape"]["parameters"]) == 1
                and entry["shape"]["parameters"][0]["path"] == "UsedBox"
                and [
                    parameter["path"]
                    for parameter in entry["shape"]["parameters"][0]["parameters"]
                ]
                == ["StdTypes.Int"]
                for entry in main["typeUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["shape"]["kind"] == "function"
                and [
                    argument["shape"]["path"]
                    for argument in entry["shape"]["arguments"]
                ]
                == ["StdTypes.Int"]
                and entry["shape"]["returnType"]["path"] == "String"
                for entry in main["typeUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["shape"]["kind"] == "anonymous"
                and [
                    field["name"]
                    for field in entry["shape"]["fields"]
                ]
                == ["count", "label"]
                for entry in main["typeUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["kind"] == "call"
                and entry["target"] == "go.Fmt"
                and entry["member"] == "println"
                for entry in main["memberUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["kind"] == "field_access"
                and entry["target"] == "UsedBox"
                and entry["member"] == "value"
                for entry in main["memberUsages"]
            )
        )
        self.assertTrue(
            any(
                entry["target"] == "go.Fmt"
                and entry["metadataImportPath"] == "fmt"
                and entry["resolvedImportPath"] == "fmt"
                for entry in main["nativeImports"]
            )
        )
        self.assertTrue(
            any(
                entry["target"] == "hxrt.string.GoStringRuntime"
                and entry["metadataImportPath"] == "hxrt"
                and entry["resolvedImportPath"] == "snapshot/hxrt"
                and entry["usageKind"] == "call"
                for entry in main["nativeImports"]
            )
        )
        self.assertTrue(
            any(entry["feature"] == "core" for entry in report["capabilities"])
        )

        self.assertEqual(
            report["typeUsageCount"],
            sum(len(entry["typeUsages"]) for entry in modules),
        )
        self.assertEqual(
            report["memberUsageCount"],
            sum(len(entry["memberUsages"]) for entry in modules),
        )
        self.assertEqual(
            report["nativeImportCount"],
            sum(len(entry["nativeImports"]) for entry in modules),
        )
        self.assertEqual(report["capabilityCount"], len(report["capabilities"]))

        for value in flatten_strings(report):
            self.assertNotIn(str(ROOT), value)
            self.assertFalse(value.startswith("/"), value)


if __name__ == "__main__":
    unittest.main()
