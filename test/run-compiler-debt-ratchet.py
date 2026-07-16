#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
import re
import sys
from collections import defaultdict
from functools import lru_cache
from pathlib import Path
from typing import Any, Iterable


ROOT = Path(__file__).resolve().parent.parent
DEFAULT_POLICY = ROOT / "test" / "compiler_debt_policy.json"
DEFAULT_REPORT_DIR = ROOT / ".cache" / "compiler-debt"
DEFAULT_STDLIB_INTRINSIC_REGISTRY = ROOT / "docs" / "compiler-stdlib-intrinsics.json"

GUARDED_METRICS = (
    "go_raw",
    "haxe_dynamic",
    "haxe_any",
    "go_unsafe",
    "go_reflection",
    "compiler_shim",
)
ROW_KEY_FIELDS = (
    "metric",
    "file",
    "context",
    "owner",
    "capability",
    "profile",
    "surface",
    "classification",
    "exception_id",
)

HAXE_TOKEN_METRICS = {
    "Dynamic": "haxe_dynamic",
    "Any": "haxe_any",
}
GO_SELECTOR_METRICS = {
    "unsafe": "go_unsafe",
    "reflect": "go_reflection",
}
FUNCTION_RE = re.compile(r"\bfunction\s+([A-Za-z_][A-Za-z0-9_]*)\s*\(")
QUALIFIED_GO_RAW_RE = re.compile(r"\bGo(?:Stmt|Expr)\.GoRaw\s*\(")
BARE_GO_RAW_RE = re.compile(r"(?<![A-Za-z0-9_.])GoRaw\s*\(")
COMPILER_SHIM_RE = re.compile(
    r"\bfunction\s+((?:lower[A-Za-z0-9_]*ShimDecls)|reflectFieldsShimDecl)\s*\("
)

SHIM_CAPABILITIES = {
    "lowerStdlibShimDecls": "stdlib_dispatch",
    "lowerIoStdlibShimDecls": "io",
    "lowerGoConcurrencyShimDecls": "go_concurrency",
    "lowerTypedGoConcurrencyShimDecls": "go_concurrency",
    "lowerTypedGoCollectionShimDecls": "go_collections",
    "lowerGoResultShimDecls": "go_result",
    "lowerTypedGoResultShimDecls": "go_result",
    "lowerHttpStdlibShimDecls": "http",
    "lowerFileSystemShimDecls": "filesystem",
    "lowerStdlibSymbolShimDecls": "stdlib_symbols",
    "reflectFieldsShimDecl": "reflection",
    "lowerTypeReflectionShimDecls": "reflection",
    "lowerRegexSerializerShimDecls": "regex_serializer",
    "lowerNetSocketShimDecls": "net_socket",
}


def relative_path(root: Path, path: Path) -> str:
    return path.resolve().relative_to(root.resolve()).as_posix()


def mask_comments_and_strings(text: str, *, preserve_strings: bool = False) -> str:
    """Replace comments and optionally strings while preserving line layout."""
    out: list[str] = []
    index = 0
    state = "code"
    quote = ""
    while index < len(text):
        char = text[index]
        following = text[index + 1] if index + 1 < len(text) else ""

        if state == "code":
            if char == "/" and following == "/":
                out.extend((" ", " "))
                index += 2
                state = "line_comment"
                continue
            if char == "/" and following == "*":
                out.extend((" ", " "))
                index += 2
                state = "block_comment"
                continue
            if char in ('"', "'", "`"):
                quote = char
                out.append(char if preserve_strings else " ")
                index += 1
                state = "string"
                continue
            out.append(char)
            index += 1
            continue

        if state == "line_comment":
            if char == "\n":
                out.append("\n")
                state = "code"
            else:
                out.append(" ")
            index += 1
            continue

        if state == "block_comment":
            if char == "*" and following == "/":
                out.extend((" ", " "))
                index += 2
                state = "code"
                continue
            out.append("\n" if char == "\n" else " ")
            index += 1
            continue

        if state == "string":
            if char == "\n" and quote != "`":
                out.append("\n")
                index += 1
                state = "code"
                continue
            if char == "\\" and quote != "`" and following:
                if preserve_strings:
                    out.extend((char, following))
                else:
                    out.extend((" ", "\n" if following == "\n" else " "))
                index += 2
                continue
            if char == quote:
                out.append(char if preserve_strings else " ")
                index += 1
                state = "code"
                continue
            out.append(char if preserve_strings else ("\n" if char == "\n" else " "))
            index += 1
            continue

    return "".join(out)


def shim_capability(context: str) -> str:
    return SHIM_CAPABILITIES.get(context, "stdlib_shim")


@lru_cache(maxsize=None)
def compiler_shim_policy_by_context(root_value: str) -> dict[str, tuple[str, str]]:
    registry_path = Path(root_value) / "docs" / "compiler-stdlib-intrinsics.json"
    registry = json.loads(registry_path.read_text(encoding="utf-8"))
    policies: dict[str, tuple[str, str]] = {}

    def register(context: str, status: str, label: str) -> None:
        if context in policies:
            raise ValueError(f"duplicate compiler shim registry context {context}: {label}")
        if status == "migration_required":
            policies[context] = ("avoidable", "compiler_stdlib_migration_debt")
        elif status == "approved_intrinsic":
            policies[context] = ("required", "compiler_stdlib_intrinsic_boundary")
        elif status == "native_api":
            policies[context] = ("required", "native_compiler_shim_boundary")
        else:
            raise ValueError(f"unknown compiler shim registry status {status}: {label}")

    dispatcher = registry["dispatcher"]
    register(dispatcher["context"], dispatcher["status"], "dispatcher")
    for group in registry["groups"]:
        register(group["entryPoint"], group["status"], group["group"])
        for support in group.get("supportEntryPoints", []):
            register(support["context"], support["status"], group["group"])
    for group in registry["nativeCompilerGroups"]:
        for context in group["debtContexts"]:
            register(context, "native_api", group["group"])
    return policies


def compiler_shim_dimensions(root: Path, context: str) -> dict[str, str]:
    policies = compiler_shim_policy_by_context(str(root.resolve()))
    if context not in policies:
        raise ValueError(
            f"unregistered compiler shim entry point {context}; update "
            f"{DEFAULT_STDLIB_INTRINSIC_REGISTRY.relative_to(ROOT)} first"
        )
    classification, exception_id = policies[context]
    return {
        "owner": "compiler_shim",
        "capability": shim_capability(context),
        "profile": "shared",
        "surface": "compiler",
        "classification": classification,
        "exception_id": exception_id,
    }


def source_capability(file: str) -> str:
    if file.startswith("runtime/hxrt/terminal"):
        return "sys"
    if file.endswith("/Sys.hx") or "/hxrt/sys/" in file:
        return "sys"
    if file.endswith("atomic_object.go"):
        return "atomic"
    if "/atomic/" in file:
        return "atomic"
    if file.endswith(("/enum_value.go", "/map_int.go", "/map_object.go", "/map_string.go")):
        return "collections"
    if "/collections/" in file:
        return "collections"
    if "/thread/" in file or file.endswith("Thread.hx"):
        return "threading"
    if "/ssl/" in file:
        return "ssl"
    if "/rtti/" in file or file.endswith("NativeStackTrace.hx"):
        return "reflection"
    if "Json" in file:
        return "json"
    if file.endswith("Template.hx"):
        return "template"
    if "/iterators/" in file or "/ds/" in file:
        return "collections"
    return "dynamic_values"


def go_raw_dimensions(file: str, context: str) -> dict[str, str]:
    if file == "src/reflaxe/go/GoCompiler.hx":
        if context in SHIM_CAPABILITIES:
            return {
                "owner": "compiler_shim",
                "capability": shim_capability(context),
            }
        return {"owner": "compiler_core", "capability": "typed_lowering"}
    if file.endswith("GoRegexSerializerEmitter.hx"):
        return {"owner": "compiler_shim", "capability": "regex_serializer"}
    if file.endswith("GoNetSocketEmitter.hx"):
        return {"owner": "compiler_shim", "capability": "net_socket"}
    if file.endswith("GoTypeReflectionEmitter.hx") or file.endswith("GoRttiMetadataEmitter.hx"):
        return {"owner": "compiler_shim", "capability": "reflection"}
    if file.endswith("GoLambdaIterableLowering.hx"):
        return {"owner": "compiler_core", "capability": "collection_lowering"}
    if file.endswith("GoTestAstFixtureEmitter.hx"):
        return {"owner": "test_infrastructure", "capability": "ast_fixture"}
    if "/ast/transformers/" in file:
        return {"owner": "compiler_infrastructure", "capability": "ast_transform"}
    return {"owner": "compiler_core", "capability": "typed_lowering"}


def haxe_token_dimensions(file: str) -> dict[str, str]:
    if file.startswith("src/go/"):
        return {
            "owner": "go_native_api",
            "capability": "native_interop",
            "profile": "native_opt_in",
            "surface": "source_api",
            "classification": "required",
            "exception_id": "native_dynamic_contract",
        }
    if file.startswith("std/go/_std/"):
        return {
            "owner": "staged_std",
            "capability": source_capability(file),
            "profile": "shared",
            "surface": "staged_std",
            "classification": "required",
            "exception_id": "staged_std_dynamic_contract",
        }
    if file.startswith("std/"):
        return {
            "owner": "runtime_facade",
            "capability": source_capability(file),
            "profile": "shared",
            "surface": "staged_std",
            "classification": "required",
            "exception_id": "runtime_facade_dynamic_contract",
        }
    return {
        "owner": "compiler_adapter",
        "capability": "macro_and_tooling_boundary",
        "profile": "shared",
        "surface": "compiler",
        "classification": "required",
        "exception_id": "compiler_dynamic_boundary",
    }


def go_selector_dimensions(file: str, metric: str) -> dict[str, str]:
    is_unsafe = metric == "go_unsafe"
    if file.startswith("runtime/hxrt/"):
        is_admitted_terminal_boundary = (
            is_unsafe and file == "runtime/hxrt/terminal_posix.go"
        )
        return {
            "owner": "runtime_hxrt",
            "capability": source_capability(file),
            "profile": "shared",
            "surface": "runtime",
            "classification": (
                "required"
                if is_admitted_terminal_boundary or not is_unsafe
                else "avoidable"
            ),
            "exception_id": "runtime_unsafe_boundary" if is_unsafe else "runtime_reflection_boundary",
        }

    parts = file.split("/")
    profile = "shared"
    if "generated" in parts:
        generated_index = parts.index("generated")
        if generated_index + 1 < len(parts):
            profile = parts[generated_index + 1]
    generated_runtime = "/hxrt/" in file
    return {
        "owner": "runtime_hxrt" if generated_runtime else "generated_output",
        "capability": "runtime_support" if generated_runtime else "dynamic_semantics",
        "profile": profile,
        "surface": "generated_runtime" if generated_runtime else "generated_program",
        "classification": "avoidable" if is_unsafe else "required",
        "exception_id": "generated_unsafe_boundary" if is_unsafe else "generated_reflection_boundary",
    }


def finding(
    metric: str,
    file: str,
    context: str,
    dimensions: dict[str, str],
) -> dict[str, str]:
    return {
        "metric": metric,
        "file": file,
        "context": context,
        **dimensions,
    }


def collect_haxe_findings(root: Path, path: Path) -> list[dict[str, str]]:
    file = relative_path(root, path)
    masked = mask_comments_and_strings(path.read_text(encoding="utf-8"))
    lines = masked.splitlines()
    function_indents = [
        len(line.expandtabs(4)) - len(line.expandtabs(4).lstrip())
        for line in lines
        if FUNCTION_RE.search(line)
    ]
    owning_function_indent = min(function_indents, default=0)
    findings: list[dict[str, str]] = []
    current_function = "<module>"
    for line in lines:
        function_match = FUNCTION_RE.search(line)
        line_indent = len(line.expandtabs(4)) - len(line.expandtabs(4).lstrip())
        if function_match and line_indent == owning_function_indent:
            current_function = function_match.group(1)

        bare_go_raw_calls = [
            match
            for match in BARE_GO_RAW_RE.finditer(line)
            if not re.search(r"\bcase\s*$", line[: match.start()])
            and not (
                file == "src/reflaxe/go/ast/GoAST.hx"
                and re.match(r"^\s*GoRaw\s*\([^)]*\)\s*;\s*$", line)
            )
        ]
        for _ in [*QUALIFIED_GO_RAW_RE.finditer(line), *bare_go_raw_calls]:
            dimensions = {
                **go_raw_dimensions(file, current_function),
                "profile": "shared",
                "surface": "compiler",
                "classification": "avoidable",
                "exception_id": "typed_ast_gap",
            }
            findings.append(finding("go_raw", file, current_function, dimensions))

        for token, metric in HAXE_TOKEN_METRICS.items():
            for _ in re.finditer(rf"\b{token}\b", line):
                findings.append(finding(metric, file, current_function, haxe_token_dimensions(file)))

        if file == "src/reflaxe/go/GoCompiler.hx":
            for match in COMPILER_SHIM_RE.finditer(line):
                context = match.group(1)
                findings.append(
                    finding(
                        "compiler_shim",
                        file,
                        context,
                        compiler_shim_dimensions(root, context),
                    )
                )
    return findings


def go_package_imports(text: str) -> list[tuple[str, str]]:
    comments_masked = mask_comments_and_strings(text, preserve_strings=True)
    imports: list[tuple[str, str]] = []
    in_block = False
    single_re = re.compile(
        r'^\s*import\s+(?:(?P<alias>[A-Za-z_][A-Za-z0-9_]*|\.)\s+)?"(?P<package>reflect|unsafe)"\s*$'
    )
    block_re = re.compile(
        r'^\s*(?:(?P<alias>[A-Za-z_][A-Za-z0-9_]*|\.)\s+)?"(?P<package>reflect|unsafe)"\s*$'
    )
    for line in comments_masked.splitlines():
        stripped = line.strip()
        if not in_block:
            if re.match(r"^import\s*\(\s*$", stripped):
                in_block = True
                continue
            match = single_re.match(line)
        else:
            if stripped == ")":
                in_block = False
                continue
            match = block_re.match(line)
        if match is None:
            continue
        package = match.group("package")
        alias = match.group("alias") or package
        imports.append((package, alias))
    return imports


def collect_go_findings(root: Path, path: Path) -> list[dict[str, str]]:
    file = relative_path(root, path)
    text = path.read_text(encoding="utf-8")
    masked = mask_comments_and_strings(text)
    findings: list[dict[str, str]] = []
    imports = go_package_imports(text)
    for package, _alias in imports:
        metric = GO_SELECTOR_METRICS[package]
        findings.append(finding(metric, file, "<import>", go_selector_dimensions(file, metric)))
    for package, alias in imports:
        if alias in {"_", "."}:
            continue
        metric = GO_SELECTOR_METRICS[package]
        for line in masked.splitlines():
            for _ in re.finditer(rf"\b{re.escape(alias)}\s*\.", line):
                findings.append(finding(metric, file, "<module>", go_selector_dimensions(file, metric)))
    return findings


def unique_files(paths: Iterable[Path]) -> list[Path]:
    return sorted(set(path.resolve() for path in paths if path.is_file()))


def collect_findings(root: Path) -> list[dict[str, str]]:
    root = root.resolve()
    haxe_files = unique_files(
        list((root / "src" / "reflaxe" / "go").rglob("*.hx"))
        + list((root / "src" / "go").rglob("*.hx"))
        + list((root / "std").rglob("*.hx"))
    )
    go_files = unique_files(
        list((root / "runtime" / "hxrt").rglob("*.go"))
        + list((root / "examples").glob("*/generated/portable/**/*.go"))
        + list((root / "examples").glob("*/generated/metal/**/*.go"))
    )

    findings: list[dict[str, str]] = []
    for path in haxe_files:
        findings.extend(collect_haxe_findings(root, path))
    for path in go_files:
        findings.extend(collect_go_findings(root, path))
    return findings


def aggregate_rows(rows: list[dict[str, Any]], field: str) -> list[dict[str, Any]]:
    totals: dict[str, int] = defaultdict(int)
    for row in rows:
        totals[str(row[field])] += int(row["count"])
    return [{field: key, "count": totals[key]} for key in sorted(totals)]


def build_report(findings: list[dict[str, str]]) -> dict[str, Any]:
    grouped: dict[tuple[str, ...], int] = defaultdict(int)
    for item in findings:
        grouped[tuple(item[field] for field in ROW_KEY_FIELDS)] += 1

    rows: list[dict[str, Any]] = []
    for key in sorted(grouped):
        row = dict(zip(ROW_KEY_FIELDS, key))
        row["count"] = grouped[key]
        rows.append(row)

    totals = {metric: 0 for metric in GUARDED_METRICS}
    for row in rows:
        totals[row["metric"]] += row["count"]

    return {
        "schema_version": 1,
        "guarded_metrics": list(GUARDED_METRICS),
        "totals": totals,
        "by_metric": aggregate_rows(rows, "metric"),
        "by_owner": aggregate_rows(rows, "owner"),
        "by_capability": aggregate_rows(rows, "capability"),
        "by_profile": aggregate_rows(rows, "profile"),
        "by_surface": aggregate_rows(rows, "surface"),
        "by_classification": aggregate_rows(rows, "classification"),
        "by_file": rows,
    }


def row_key(row: dict[str, Any]) -> tuple[str, ...]:
    return tuple(str(row[field]) for field in ROW_KEY_FIELDS)


def validate_policy(policy: dict[str, Any]) -> list[str]:
    errors: list[str] = []
    if policy.get("schema_version") != 1:
        errors.append("policy schema_version must be 1")
    if tuple(policy.get("guarded_metrics", [])) != GUARDED_METRICS:
        errors.append("policy guarded_metrics must exactly match the runner's ordered metric set")

    exceptions = policy.get("exceptions")
    if not isinstance(exceptions, dict) or not exceptions:
        errors.append("policy exceptions must be a non-empty object")
        exceptions = {}
    else:
        for exception_id, exception in sorted(exceptions.items()):
            if not isinstance(exception, dict):
                errors.append(f"exception {exception_id} must be an object")
                continue
            for field in ("owner", "classification", "what", "why", "how"):
                value = exception.get(field)
                if not isinstance(value, str) or not value.strip():
                    errors.append(f"exception {exception_id} is missing {field}")
            if exception.get("classification") not in {"required", "avoidable"}:
                errors.append(f"exception {exception_id} has invalid classification")

    limits = policy.get("limits")
    if not isinstance(limits, list):
        errors.append("policy limits must be an array")
        return errors

    seen: set[tuple[str, ...]] = set()
    for index, limit in enumerate(limits):
        if not isinstance(limit, dict):
            errors.append(f"limit {index} must be an object")
            continue
        missing = [field for field in ROW_KEY_FIELDS if not isinstance(limit.get(field), str)]
        if missing:
            errors.append(f"limit {index} is missing key fields: {', '.join(missing)}")
            continue
        key = row_key(limit)
        if key in seen:
            errors.append(f"duplicate limit for {' | '.join(key)}")
        seen.add(key)
        if limit["metric"] not in GUARDED_METRICS:
            errors.append(f"limit {index} has unguarded metric {limit['metric']}")
        if Path(limit["file"]).is_absolute():
            errors.append(f"limit {index} contains an absolute file path")
        if not isinstance(limit.get("max_count"), int) or limit["max_count"] < 0:
            errors.append(f"limit {index} max_count must be a non-negative integer")
        exception = exceptions.get(limit["exception_id"])
        if exception is None:
            errors.append(f"limit {index} references unknown exception {limit['exception_id']}")
        elif exception.get("classification") != limit["classification"]:
            errors.append(f"limit {index} classification disagrees with its exception")
    return errors


def compare_report_to_policy(report: dict[str, Any], policy: dict[str, Any]) -> list[str]:
    limits = {row_key(limit): limit for limit in policy.get("limits", [])}
    errors: list[str] = []
    for row in report["by_file"]:
        key = row_key(row)
        limit = limits.get(key)
        label = f"{row['metric']} at {row['file']}#{row['context']}"
        if limit is None:
            errors.append(f"unexplained debt location: {label} has count {row['count']}")
            continue
        if row["count"] > limit["max_count"]:
            errors.append(
                f"{label} exceeds baseline: current {row['count']} > allowed {limit['max_count']}"
            )
    return errors


def render_table(rows: list[dict[str, Any]], field: str) -> list[str]:
    title = field.replace("_", " ").title()
    lines = [f"| {title} | Count |", "| --- | ---: |"]
    lines.extend(f"| `{row[field]}` | {row['count']} |" for row in rows)
    return lines


def render_markdown(report: dict[str, Any]) -> str:
    lines = [
        "# Compiler Debt Ratchet Report",
        "",
        "Counts are directional evidence, not correctness scores. See `docs/compiler-debt-ratchet.md`.",
        "",
    ]
    for heading, field in (
        ("By metric", "metric"),
        ("By owner", "owner"),
        ("By capability", "capability"),
        ("By profile", "profile"),
        ("By surface", "surface"),
        ("By classification", "classification"),
    ):
        lines.extend((f"## {heading}", ""))
        lines.extend(render_table(report[f"by_{field}"], field))
        lines.append("")

    lines.extend(
        (
            "## By file",
            "",
            "| Metric | File | Context | Owner | Capability | Profile | Surface | Classification | Count |",
            "| --- | --- | --- | --- | --- | --- | --- | --- | ---: |",
        )
    )
    for row in report["by_file"]:
        lines.append(
            "| "
            + " | ".join(
                f"`{row[field]}`"
                for field in (
                    "metric",
                    "file",
                    "context",
                    "owner",
                    "capability",
                    "profile",
                    "surface",
                    "classification",
                )
            )
            + f" | {row['count']} |"
        )
    lines.append("")
    return "\n".join(lines)


def update_policy_limits(policy: dict[str, Any], report: dict[str, Any]) -> dict[str, Any]:
    updated = dict(policy)
    updated["limits"] = [
        {**{field: row[field] for field in ROW_KEY_FIELDS}, "max_count": row["count"]}
        for row in report["by_file"]
    ]
    return updated


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Report and ratchet compiler/runtime dynamic and raw-emission debt"
    )
    parser.add_argument("--root", type=Path, default=ROOT)
    parser.add_argument("--policy", type=Path, default=DEFAULT_POLICY)
    parser.add_argument("--report-dir", type=Path, default=DEFAULT_REPORT_DIR)
    parser.add_argument(
        "--update-baseline",
        action="store_true",
        help="Replace policy ceilings with current counts; review the resulting diff before commit",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    policy = json.loads(args.policy.read_text(encoding="utf-8"))
    policy_errors = validate_policy(policy)
    if policy_errors:
        print("Compiler debt policy is invalid:", file=sys.stderr)
        for error in policy_errors:
            print(f"- {error}", file=sys.stderr)
        return 2

    report = build_report(collect_findings(args.root))
    args.report_dir.mkdir(parents=True, exist_ok=True)
    (args.report_dir / "report.json").write_text(
        json.dumps(report, indent=2, sort_keys=True) + "\n",
        encoding="utf-8",
    )
    (args.report_dir / "report.md").write_text(render_markdown(report), encoding="utf-8")

    if args.update_baseline:
        updated = update_policy_limits(policy, report)
        args.policy.write_text(json.dumps(updated, indent=2, sort_keys=True) + "\n", encoding="utf-8")
        print(f"Updated {args.policy}")
        return 0

    errors = compare_report_to_policy(report, policy)
    if errors:
        print("Compiler debt ratchet failed:", file=sys.stderr)
        for error in errors:
            print(f"- {error}", file=sys.stderr)
        print("Reduce the debt or update the reviewed policy rationale and baseline.", file=sys.stderr)
        return 1

    summary = ", ".join(f"{metric}={report['totals'][metric]}" for metric in GUARDED_METRICS)
    print(f"Compiler debt ratchet passed: {summary}")
    print("Reports: .cache/compiler-debt/report.json and .cache/compiler-debt/report.md")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
