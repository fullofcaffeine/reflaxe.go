#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import unittest
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parent.parent
REGISTRY_PATH = ROOT / "docs" / "compiler-stdlib-intrinsics.json"
COMPILER_PATH = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
CLASSIFIER_PATH = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoStdlibShimClassifier.hx"
OWNERSHIP_PATH = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoStdlibOwnership.hx"
PLANNER_PATH = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSourceOwnedStdlibPlanner.hx"
RTTI_EMITTER_PATH = ROOT / "src" / "reflaxe" / "go" / "compiler" / "emit" / "GoRttiMetadataEmitter.hx"
DEBT_POLICY_PATH = ROOT / "test" / "compiler_debt_policy.json"
PACKAGE_PATH = ROOT / "package.json"
RELEASE_RUNNER_PATH = ROOT / "test" / "run-release-contracts.py"

ALLOWED_STATUS = {"approved_intrinsic", "migration_required"}
ALLOWED_SYMBOL_SCOPE = {"complete_type", "partial_type", "member"}
FOLLOW_UP_RE = re.compile(r"^haxe_go-[a-z0-9.-]+$")


def load_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def function_source(source: str, name: str) -> str:
    match = re.search(rf"\bfunction\s+{re.escape(name)}\s*\([^)]*\)[^{{]*{{", source)
    if match is None:
        raise AssertionError(f"missing function {name}")

    start = match.start()
    index = match.end() - 1
    depth = 0
    state = "code"
    quote = ""
    while index < len(source):
        char = source[index]
        following = source[index + 1] if index + 1 < len(source) else ""
        if state == "code":
            if char == "/" and following == "/":
                state = "line_comment"
                index += 2
                continue
            if char == "/" and following == "*":
                state = "block_comment"
                index += 2
                continue
            if char in {'"', "'"}:
                state = "string"
                quote = char
                index += 1
                continue
            if char == "{":
                depth += 1
            elif char == "}":
                depth -= 1
                if depth == 0:
                    return source[start : index + 1]
            index += 1
            continue
        if state == "line_comment":
            if char == "\n":
                state = "code"
            index += 1
            continue
        if state == "block_comment":
            if char == "*" and following == "/":
                state = "code"
                index += 2
                continue
            index += 1
            continue
        if char == "\\":
            index += 2
            continue
        if char == quote:
            state = "code"
        index += 1
    raise AssertionError(f"unterminated function {name}")


def strip_comments(source: str) -> str:
    source = re.sub(r"/\*.*?\*/", "", source, flags=re.DOTALL)
    return re.sub(r"//[^\n]*", "", source)


def dispatch_groups(source: str) -> dict[str, str]:
    body = function_source(source, "lowerStdlibShimDecls")
    pairs = re.findall(
        r'requiredStdlibShimGroups\.exists\("([^"]+)"\).*?concat\((lower[A-Za-z0-9_]+ShimDecls)\(\)\)',
        body,
        flags=re.DOTALL,
    )
    return dict(pairs)


def classifier_surfaces(source: str) -> dict[tuple[str, str], tuple[str, ...]]:
    entries: dict[tuple[str, str], tuple[str, ...]] = {}
    for kind, path, raw_groups in re.findall(
        r'\{\s*kind:\s*"(class|enum)",\s*path:\s*"([^"]+)",\s*groups:\s*\[([^]]*)\]\s*\}',
        source,
    ):
        groups = tuple(re.findall(r'"([^"]+)"', raw_groups))
        entries[(kind, path)] = groups
    return entries


def planner_surfaces(source: str) -> dict[str, tuple[str, ...]]:
    body = function_source(source, "noteSourceOwnedStdlibUsage")
    entries: dict[str, tuple[str, ...]] = {}
    cases = list(re.finditer(r'\bcase\s+((?:"[^"]+"\s*,?\s*)+):', body))
    for index, case in enumerate(cases):
        end = cases[index + 1].start() if index + 1 < len(cases) else body.find("case _:", case.end())
        if end < 0:
            end = len(body)
        case_body = body[case.end() : end]
        groups = tuple(sorted(set(re.findall(r'requireStdlibShimGroup\("([^"]+)"\)', case_body))))
        if not groups:
            continue
        for symbol in re.findall(r'"([^"]+)"', case.group(1)):
            entries[symbol] = groups
    return entries


def group_dependencies(source: str) -> dict[str, tuple[str, ...]]:
    # A fully retired dependency layer has no hook to parse. The exact registry
    # comparison below still fails closed if any group declares dependencies.
    if re.search(r"\bfunction\s+applyStdlibShimGroupDependencies\s*\(", source) is None:
        return {}
    body = function_source(source, "applyStdlibShimGroupDependencies")
    matches = list(re.finditer(r'requiredStdlibShimGroups\.exists\("([^"]+)"\)', body))
    dependencies: dict[str, tuple[str, ...]] = {}
    for index, match in enumerate(matches):
        end = matches[index + 1].start() if index + 1 < len(matches) else len(body)
        block = body[match.end() : end]
        dependencies[match.group(1)] = tuple(
            sorted(set(re.findall(r'requireStdlibShimGroup\("([^"]+)"\)', block)))
        )
    return dependencies


def compiler_owned_authorities(source: str) -> set[str]:
    body = function_source(source, "isCompilerOwnedAuthority")
    required_case = re.search(r"\bcase\s+(.*?):\s*\n?\s*true;", body, flags=re.DOTALL)
    if required_case is None:
        # An empty authority set is a valid end state; any registered authority
        # still makes the exact set comparison fail.
        return set()
    return set(re.findall(r'"([^"]+)"', required_case.group(1)))


def static_call_symbol(class_name: str, raw_pack: str, field_name: str) -> str:
    package = ".".join(re.findall(r'"([^"]+)"', raw_pack))
    owner = f"{package}.{class_name}" if package else class_name
    return f"{owner}.{field_name}"


def direct_call_lowerings(source: str) -> set[str]:
    functions = (
        "lowerCall",
        "lowerDsSortHelperCall",
        "lowerLambdaSourceCallAdapter",
        "lowerRestAbstractCall",
    )
    symbols: set[str] = set()
    for function in functions:
        body = strip_comments(function_source(source, function))
        for class_name, raw_pack, field_name in re.findall(
            r'isStaticCall\(callee,\s*"([^"]+)",\s*\[([^]]*)\],\s*"([^"]+)"\)',
            body,
        ):
            symbols.add(static_call_symbol(class_name, raw_pack, field_name))
    return symbols


def special_stdlib_intrinsics(compiler: str, rtti_emitter: str) -> set[str]:
    symbols: set[str] = set()
    if 'fullClassName(classType) == "haxe.Resource" && field.name == "content"' in compiler:
        symbols.add("haxe.Resource.content")
    if 'classType.name == "Sys" && fieldName == "cpuTime"' in compiler:
        symbols.add("Sys.cpuTime")
    if "asHaxeExceptionMessageGetterTarget(callee)" in compiler:
        symbols.add("haxe.Exception.message")
    if "asHaxeExceptionToStringTarget(callee)" in compiler:
        symbols.add("haxe.Exception.toString")
    if "asHaxeValueExceptionUnwrapTarget(callee)" in compiler:
        symbols.add("haxe.ValueException.value")
    if 'case \\"__meta__\\"' in rtti_emitter:
        symbols.add("generated_type.__meta__")
    if 'case \\"__rtti\\"' in rtti_emitter:
        symbols.add("generated_type.__rtti")
    return symbols


def decision_fields(
    entry: dict[str, Any],
    label: str,
    decisions: dict[str, dict[str, Any]] | None = None,
) -> None:
    if decisions is not None:
        decision_id = entry.get("decisionId")
        if not isinstance(decision_id, str) or decision_id not in decisions:
            raise AssertionError(f"{label}: missing registered decisionId")
        entry = decisions[decision_id]
    for field in ("what", "why", "how", "reviewCondition"):
        value = entry.get(field)
        if not isinstance(value, str) or not value.strip():
            raise AssertionError(f"{label}: missing {field}")
    evidence = entry.get("evidence")
    if not isinstance(evidence, list) or not evidence or not all(isinstance(item, str) and item for item in evidence):
        raise AssertionError(f"{label}: evidence must be a non-empty string list")


class CompilerStdlibIntrinsicRegistryContract(unittest.TestCase):
    def setUp(self) -> None:
        self.registry = load_json(REGISTRY_PATH)
        self.compiler = COMPILER_PATH.read_text(encoding="utf-8")

    def test_registry_schema_and_decisions_are_explicit(self) -> None:
        registry = self.registry
        self.assertEqual(1, registry.get("schemaVersion"))
        self.assertEqual(
            "exact_haxe_symbols_and_compiler_entry_points",
            registry.get("inventoryUnit"),
        )
        self.assertEqual(
            "portable_haxe_stdlib",
            registry.get("scope"),
        )

        decisions = registry.get("decisions")
        self.assertIsInstance(decisions, dict)
        self.assertEqual(sorted(decisions), list(decisions))
        for decision_id, decision in decisions.items():
            decision_fields(decision, f"decision {decision_id}")

        dispatcher = registry.get("dispatcher")
        self.assertIsInstance(dispatcher, dict)
        self.assertEqual("lowerStdlibShimDecls", dispatcher.get("context"))
        self.assertEqual("migration_required", dispatcher.get("status"))
        decision_fields(dispatcher, "dispatcher", decisions)

        groups = registry.get("groups")
        self.assertIsInstance(groups, list)
        group_names = [entry.get("group") for entry in groups]
        self.assertEqual(sorted(set(group_names)), group_names)
        self.assertTrue(group_names)

        owned_symbols: set[str] = set()
        authority_names: set[str] = set()
        for group in groups:
            label = f"group {group.get('group')}"
            self.assertIn(group.get("status"), ALLOWED_STATUS, label)
            decision_fields(group, label, decisions)
            if group["status"] == "migration_required":
                self.assertRegex(group.get("followUpBead", ""), FOLLOW_UP_RE, label)
            else:
                self.assertNotIn("followUpBead", group, label)

            symbols = group.get("ownedSymbols")
            self.assertIsInstance(symbols, list, label)
            self.assertTrue(symbols, label)
            for symbol in symbols:
                self.assertIsInstance(symbol, dict, label)
                name = symbol.get("symbol")
                self.assertIsInstance(name, str, label)
                self.assertTrue(name, label)
                self.assertNotIn(name, owned_symbols, name)
                owned_symbols.add(name)
                self.assertIn(symbol.get("scope"), ALLOWED_SYMBOL_SCOPE, name)
                if symbol["scope"] == "partial_type":
                    members = symbol.get("members")
                    self.assertIsInstance(members, list, name)
                    self.assertEqual(sorted(set(members)), members, name)
                    self.assertTrue(members, name)

            authorities = group.get("compilerOwnedAuthorities")
            self.assertIsInstance(authorities, list, label)
            self.assertEqual(sorted(set(authorities)), authorities, label)
            for authority in authorities:
                self.assertNotIn(authority, authority_names, authority)
                authority_names.add(authority)

            selectors = group.get("classifierSelections")
            self.assertIsInstance(selectors, list, label)
            for selector in selectors:
                self.assertIn(selector.get("kind"), {"class", "enum"}, label)
                self.assertIsInstance(selector.get("path"), str, label)

            planner = group.get("plannerSelections")
            self.assertIsInstance(planner, list, label)
            self.assertEqual(sorted(set(planner)), planner, label)
            dependencies = group.get("dependencies")
            self.assertIsInstance(dependencies, list, label)
            self.assertEqual(sorted(set(dependencies)), dependencies, label)
            self.assertIsInstance(group.get("declaresDependencyBlock"), bool, label)

        direct = registry.get("directLowerings")
        self.assertIsInstance(direct, list)
        direct_names = [entry.get("symbol") for entry in direct]
        self.assertEqual(sorted(set(direct_names)), direct_names)
        for entry in direct:
            label = f"direct lowering {entry.get('symbol')}"
            self.assertIn(entry.get("status"), ALLOWED_STATUS, label)
            decision_fields(entry, label, decisions)
            if entry["status"] == "migration_required":
                self.assertRegex(entry.get("followUpBead", ""), FOLLOW_UP_RE, label)
            else:
                self.assertNotIn("followUpBead", entry, label)

        special = registry.get("specialIntrinsics")
        self.assertIsInstance(special, list)
        special_names = [entry.get("symbol") for entry in special]
        self.assertEqual(sorted(set(special_names)), special_names)
        for entry in special:
            label = f"special intrinsic {entry.get('symbol')}"
            self.assertIn(entry.get("status"), ALLOWED_STATUS, label)
            decision_fields(entry, label, decisions)
            if entry["status"] == "migration_required":
                self.assertRegex(entry.get("followUpBead", ""), FOLLOW_UP_RE, label)
            else:
                self.assertNotIn("followUpBead", entry, label)

    def test_dispatch_and_entry_points_are_bidirectional(self) -> None:
        actual_dispatch = dispatch_groups(self.compiler)
        registered_groups = {
            entry["group"]: entry["entryPoint"] for entry in self.registry["groups"]
        }
        native_groups = {
            entry["group"]: entry["entryPoint"] for entry in self.registry["nativeCompilerGroups"]
        }
        self.assertEqual(actual_dispatch, {**registered_groups, **native_groups})

        portable_contexts = {self.registry["dispatcher"]["context"]}
        context_status = {
            self.registry["dispatcher"]["context"]: self.registry["dispatcher"]["status"]
        }
        for group in self.registry["groups"]:
            portable_contexts.add(group["entryPoint"])
            context_status[group["entryPoint"]] = group["status"]
            for support in group.get("supportEntryPoints", []):
                portable_contexts.add(support["context"])
                context_status[support["context"]] = support["status"]
                decision_fields(
                    support,
                    f"support entry point {support['context']}",
                    self.registry["decisions"],
                )
                if support["status"] == "migration_required":
                    self.assertRegex(
                        support.get("followUpBead", ""),
                        FOLLOW_UP_RE,
                        support["context"],
                    )
                else:
                    self.assertNotIn("followUpBead", support, support["context"])

        native_contexts = {
            context
            for group in self.registry["nativeCompilerGroups"]
            for context in group["debtContexts"]
        }
        policy = load_json(DEBT_POLICY_PATH)
        shim_limits = {
            entry["context"]: entry
            for entry in policy["limits"]
            if entry.get("metric") == "compiler_shim"
        }
        self.assertEqual(portable_contexts | native_contexts, set(shim_limits))
        for context, status in context_status.items():
            limit = shim_limits[context]
            if status == "migration_required":
                self.assertEqual("avoidable", limit["classification"], context)
                self.assertEqual("compiler_stdlib_migration_debt", limit["exception_id"], context)
            else:
                self.assertEqual("required", limit["classification"], context)
                self.assertEqual("compiler_stdlib_intrinsic_boundary", limit["exception_id"], context)
        for context in native_contexts:
            self.assertEqual("required", shim_limits[context]["classification"], context)
            self.assertEqual("native_compiler_shim_boundary", shim_limits[context]["exception_id"], context)

        self.assertNotIn("compiler_shim_boundary", policy["exceptions"])
        for exception_id in (
            "compiler_stdlib_migration_debt",
            "compiler_stdlib_intrinsic_boundary",
            "native_compiler_shim_boundary",
        ):
            decision_fields(policy["exceptions"][exception_id], exception_id)

    def test_classifier_planner_dependencies_and_authorities_are_exact(self) -> None:
        registered_classifier: dict[tuple[str, str], tuple[str, ...]] = {}
        registered_planner: dict[str, set[str]] = {}
        registered_dependencies: dict[str, tuple[str, ...]] = {}
        registered_authorities: set[str] = set()
        for group in self.registry["groups"]:
            name = group["group"]
            for selector in group["classifierSelections"]:
                key = (selector["kind"], selector["path"])
                registered_classifier.setdefault(key, tuple())
                registered_classifier[key] = tuple(sorted(set(registered_classifier[key]) | {name}))
            for symbol in group["plannerSelections"]:
                registered_planner.setdefault(symbol, set()).add(name)
            if group["declaresDependencyBlock"]:
                registered_dependencies[name] = tuple(group["dependencies"])
            registered_authorities.update(group["compilerOwnedAuthorities"])

        actual_classifier = {
            key: tuple(sorted(groups))
            for key, groups in classifier_surfaces(CLASSIFIER_PATH.read_text(encoding="utf-8")).items()
        }
        self.assertEqual(registered_classifier, actual_classifier)
        self.assertEqual(
            {symbol: tuple(sorted(groups)) for symbol, groups in registered_planner.items()},
            planner_surfaces(PLANNER_PATH.read_text(encoding="utf-8")),
        )
        self.assertEqual(registered_dependencies, group_dependencies(self.compiler))
        self.assertEqual(
            registered_authorities,
            compiler_owned_authorities(OWNERSHIP_PATH.read_text(encoding="utf-8")),
        )

    def test_direct_stdlib_call_rewrites_are_exact(self) -> None:
        registered = {entry["symbol"] for entry in self.registry["directLowerings"]}
        native = set(self.registry["nativeDirectLowerings"])
        self.assertEqual(registered | native, direct_call_lowerings(self.compiler))

    def test_special_data_diagnostic_and_carrier_intrinsics_are_exact(self) -> None:
        registered = {entry["symbol"] for entry in self.registry["specialIntrinsics"]}
        self.assertEqual(
            registered,
            special_stdlib_intrinsics(
                self.compiler,
                RTTI_EMITTER_PATH.read_text(encoding="utf-8"),
            ),
        )

    def test_registry_gate_runs_in_normal_changed_governance_and_release_paths(self) -> None:
        scripts = load_json(PACKAGE_PATH)["scripts"]
        command = "python3 test/test_compiler_stdlib_intrinsic_registry.py"
        self.assertEqual(command, scripts.get("test:stdlib:intrinsics"))
        self.assertIn("npm run test:stdlib:intrinsics", scripts["test"])
        self.assertIn("npm run test:stdlib:intrinsics", scripts["test:changed"])
        self.assertIn("npm run test:stdlib:intrinsics", scripts["test:stdlib:governance"])
        self.assertIn(
            "test/test_compiler_stdlib_intrinsic_registry.py",
            RELEASE_RUNNER_PATH.read_text(encoding="utf-8"),
        )


if __name__ == "__main__":
    unittest.main()
