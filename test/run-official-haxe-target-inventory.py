#!/usr/bin/env python3

"""Run the complete applicable pinned Haxe tests/unit inventory in Go shards."""

from __future__ import annotations

import argparse
from collections import Counter, defaultdict
import hashlib
import importlib.util
import json
import os
from pathlib import Path, PurePosixPath
import re
import shutil
import subprocess
import sys
import tempfile
import time


ROOT = Path(__file__).resolve().parent.parent
INVENTORY_ROOT = ROOT / "test" / "official_haxe_target_inventory"
INVENTORY_SOURCE = INVENTORY_ROOT / "src"
MANIFEST_PATH = INVENTORY_ROOT / "manifest.json"
SMOKE_RUNNER_PATH = ROOT / "test" / "run-official-haxe-target-smoke.py"
DEFAULT_CACHE = ROOT / ".cache" / "official-haxe-target-inventory"
RECORD_PREFIX = "OFFICIAL_HAXE_INVENTORY_RECORD\t"
SUMMARY_PREFIX = "OFFICIAL_HAXE_INVENTORY_SUMMARY\t"
TEST_BASE_MARKER = "@:keepSub\nclass Test"
TEST_BASE_HELPERS = (
    "eq", "feq", "aeq", "t", "f", "assert", "exc", "unspec", "allow",
    "noAssert", "hf", "nhf", "hsf", "nhsf",
)


class InventoryError(RuntimeError):
    def __init__(self, stage: str, message: str) -> None:
        super().__init__(message)
        self.stage = stage


def load_smoke_runner():
    spec = importlib.util.spec_from_file_location("official_haxe_smoke_shared", SMOKE_RUNNER_PATH)
    if spec is None or spec.loader is None:
        raise InventoryError("environment", "cannot load the official smoke authority")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


SMOKE = load_smoke_runner()


def sha256_bytes(value: bytes) -> str:
    return hashlib.sha256(value).hexdigest()


def sha256_file(path: Path) -> str:
    return sha256_bytes(path.read_bytes())


def canonical_json_sha256(value: object) -> str:
    return sha256_bytes(
        (json.dumps(value, sort_keys=True, separators=(",", ":")) + "\n").encode("utf-8")
    )


def adapt_test_base_text(source: str, selected_helpers: set[str]) -> str:
    """Build a typed harness base from exact selected upstream helper bodies."""

    if source.count(TEST_BASE_MARKER) != 1:
        raise InventoryError(
            "adapter", "pinned unit.Test scaffolding no longer has one @:keepSub marker"
        )
    if source.count("\tpublic function new()") != 1:
        raise InventoryError("adapter", "pinned unit.Test constructor scaffolding drifted")
    chunks: dict[str, str] = {}
    for helper in TEST_BASE_HELPERS:
        match = re.search(
            rf"\n\tfunction {re.escape(helper)}(?:<[^>\n]+>)?\s*\(", source
        )
        if match is None:
            raise InventoryError("adapter", f"pinned unit.Test helper scaffolding drifted: {helper}")
        next_method = re.search(r"\n\tfunction [A-Za-z_]", source[match.end():])
        end = (
            match.end() + next_method.start()
            if next_method is not None
            else source.rfind("\n}")
        )
        if helper in selected_helpers:
            chunk = source[match.start() + 1 : end].rstrip()
            chunks[helper] = chunk.replace("\tfunction ", "\tpublic function ", 1)
    methods = "\n\n".join(chunks[name] for name in TEST_BASE_HELPERS if name in chunks)
    if methods:
        methods = "\n\n" + methods
    return (
        "import utest.Assert;\n\n"
        "@:keepSub\n"
        "class OfficialInventoryTestBase implements OfficialInventoryCase {\n"
        "\tpublic function new() {}\n\n"
        "\tpublic function __initializeUtest__():utest.TestData.InitializeUtest {\n"
        "\t\treturn {tests: [], dependencies: [], accessories: {}};\n"
        "\t}"
        + methods
        + "\n}\n"
    )


def adapt_owner_text(source: str) -> str:
    """Replace only direct upstream test-harness inheritance in one owner."""

    adapted, count = re.subn(
        r"(\bclass\s+[A-Za-z_][A-Za-z0-9_]*[^\n{]*\s)extends\s+(?:unit\.)?Test\b",
        r"\1extends OfficialInventoryTestBase implements utest.ITest",
        source,
    )
    if count == 0:
        raise InventoryError("adapter", "selected owner has no direct unit.Test inheritance")
    package = re.search(r"(?m)^(?:\ufeff)?package(?:\s+[^;]+)?;\n", adapted)
    if package is None:
        raise InventoryError("adapter", "selected owner has no package declaration")
    return adapted[: package.end()] + "import OfficialInventoryTestBase;\n" + adapted[package.end() :]


def prepare_owner_adapters(app: Path, haxe_root: Path, owners: list[dict]) -> dict:
    source_root = haxe_root / "tests" / "unit" / "src"
    records: list[dict] = []
    for owner in owners:
        if owner["family"] == "unitstd":
            continue
        source_path = source_root / owner["path"]
        source = source_path.read_text(encoding="utf-8")
        adapted = adapt_owner_text(source)
        destination = app / owner["path"]
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(adapted, encoding="utf-8")
        records.append(
            {
                "path": owner["path"],
                "sourceSha256": sha256_bytes(source.encode("utf-8")),
                "adaptedSha256": sha256_bytes(adapted.encode("utf-8")),
            }
        )
    return {
        "count": len(records),
        "sha256": canonical_json_sha256(records),
        "transformation": "replace direct unit.Test inheritance with typed harness base",
    }


def prepare_unitstd_adapters(app: Path, owners: list[dict]) -> dict:
    records: list[dict] = []
    for owner in owners:
        if owner["family"] != "unitstd":
            continue
        runtime_parts = owner["runtimeClass"].split(".")
        package = ".".join(runtime_parts[:-1])
        class_name = runtime_parts[-1]
        relative = Path(owner["path"]).relative_to("unitstd").as_posix()
        source = (
            f"package {package};\n\n"
            "import OfficialInventoryTestBase;\n\n"
            f"class {class_name} extends OfficialInventoryTestBase implements utest.ITest {{\n"
            "\tpublic function new() { super(); }\n\n"
            "\tpublic function test():Void {\n"
            f'\t\tOfficialInventoryUnitStd.body("official-unitstd/{relative}");\n'
            "\t}\n"
            "}\n"
        )
        destination = app.joinpath(*runtime_parts[:-1], f"{class_name}.hx")
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(source, encoding="utf-8")
        records.append(
            {
                "path": owner["path"],
                "runtimeClass": owner["runtimeClass"],
                "adapterSha256": sha256_bytes(source.encode("utf-8")),
            }
        )
    return {"count": len(records), "sha256": canonical_json_sha256(records)}


def prepare_test_base_adapter(app: Path, haxe_root: Path, owners: list[dict]) -> dict:
    source_path = haxe_root / "tests" / "unit" / "src" / "unit" / "Test.hx"
    source = source_path.read_text(encoding="utf-8")
    selected_helpers: set[str] = set()
    source_root = haxe_root / "tests" / "unit" / "src"
    for owner in owners:
        owner_source = (source_root / owner["path"]).read_text(encoding="utf-8")
        for helper in TEST_BASE_HELPERS:
            if re.search(rf"\b{re.escape(helper)}\s*\(", owner_source):
                selected_helpers.add(helper)
    if any(owner["family"] == "unitstd" for owner in owners):
        selected_helpers.update({"eq", "feq", "t", "f", "exc"})
    if "unspec" in selected_helpers:
        selected_helpers.add("noAssert")
    if selected_helpers.intersection({"unspec", "noAssert", "hf", "hsf"}):
        selected_helpers.add("t")
    if selected_helpers.intersection({"nhf", "nhsf"}):
        selected_helpers.add("f")
    adapted = adapt_test_base_text(source, selected_helpers)
    base_destination = app / "OfficialInventoryTestBase.hx"
    base_destination.write_text(adapted, encoding="utf-8")
    case_interface = (
        "interface OfficialInventoryCase {\n"
        "\tfunction __initializeUtest__():utest.TestData.InitializeUtest;\n"
        "}\n"
    )
    (app / "OfficialInventoryCase.hx").write_text(case_interface, encoding="utf-8")
    return {
        "source": "tests/unit/src/unit/Test.hx",
        "sourceSha256": sha256_bytes(source.encode("utf-8")),
        "adaptedSha256": sha256_bytes(adapted.encode("utf-8")),
        "caseInterfaceSha256": sha256_bytes(case_interface.encode("utf-8")),
        "selectedHelpers": sorted(selected_helpers),
        "transformation": "move selected helper bodies to a typed harness base",
    }


def discover_candidate_inventory(unit_source_root: Path) -> list[dict]:
    """Discover test owners from the four pinned upstream authority roots."""

    unit = unit_source_root / "unit"
    groups = (
        ("top-level", sorted(unit.glob("Test*.hx"))),
        ("unitstd", sorted((unit_source_root / "unitstd").rglob("*.unit.hx"))),
        ("issue", sorted((unit / "issues").glob("*.hx"))),
        ("hxcpp-issue", sorted((unit / "hxcpp_issues").glob("*.hx"))),
    )
    records: list[dict] = []
    for family, paths in groups:
        for path in paths:
            relative = path.relative_to(unit_source_root).as_posix()
            if family == "unitstd":
                spec_relative = path.relative_to(unit_source_root / "unitstd")
                package = ["unit", "spec", *spec_relative.parts[:-1]]
                # UnitBuilder uses the text before the first dot as the class
                # stem, including files such as ComplexTypeTools.macro.unit.hx.
                stem = spec_relative.name.split(".", 1)[0]
                runtime_class = ".".join([*package, f"Test{stem}"])
            elif family == "top-level":
                runtime_class = f"unit.{path.stem}"
            elif family == "issue":
                runtime_class = f"unit.issues.{path.stem}"
            else:
                runtime_class = f"unit.hxcpp_issues.{path.stem}"
            records.append(
                {
                    "family": family,
                    "path": relative,
                    "runtimeClass": runtime_class,
                    "sha256": sha256_file(path),
                }
            )
    return records


def canonical_inventory_sha256(records: list[dict]) -> str:
    return canonical_json_sha256(sorted(records, key=lambda item: (item["family"], item["path"])))


def canonical_runtime_records(records: list[dict]) -> list[dict]:
    return sorted(
        [
            {"id": str(item["id"]), "assertions": int(item["assertions"])}
            for item in records
        ],
        key=lambda item: item["id"],
    )


def canonical_runtime_sha256(records: list[dict]) -> str:
    return canonical_json_sha256(canonical_runtime_records(records))


def canonical_classification_records(records: list[dict]) -> list[dict]:
    return sorted(
        [
            {
                "family": str(item["family"]),
                "path": str(item["path"]),
                "status": str(item["status"]),
            }
            for item in records
        ],
        key=lambda item: (item["family"], item["path"]),
    )


def canonical_classification_sha256(records: list[dict]) -> str:
    return canonical_json_sha256(canonical_classification_records(records))


def validate_runtime_baseline(baseline: dict, actual: list[dict]) -> None:
    ids = [str(record.get("id")) for record in actual]
    if len(ids) != len(set(ids)):
        raise InventoryError("active-inventory", "runtime test identities are not unique")
    for record in actual:
        if record.get("status") != "pass" or int(record.get("assertions", 0)) <= 0:
            raise InventoryError("active-inventory", f"failed or assertion-free record: {record}")
    normalized = canonical_runtime_records(actual)
    if normalized != baseline.get("records"):
        expected_ids = {item["id"] for item in baseline.get("records", [])}
        actual_ids = {item["id"] for item in normalized}
        raise InventoryError(
            "active-inventory",
            f"runtime drift: missing={sorted(expected_ids - actual_ids)}, "
            f"added={sorted(actual_ids - expected_ids)}, assertion totals may also differ",
        )
    digest = canonical_runtime_sha256(normalized)
    if digest != baseline.get("sha256"):
        raise InventoryError("active-inventory", "runtime baseline digest differs")


def validate_classification_baseline(baseline: dict, actual: list[dict]) -> None:
    normalized = canonical_classification_records(actual)
    allowed = {"active", "blocked", "inapplicable"}
    unexpected = sorted({item["status"] for item in normalized} - allowed)
    if unexpected:
        raise InventoryError("classification", f"unexpected classification states: {unexpected}")
    counts = dict(sorted(Counter(item["status"] for item in normalized).items()))
    if counts != baseline.get("counts"):
        raise InventoryError(
            "classification",
            f"classification count drift: expected={baseline.get('counts')}, actual={counts}",
        )
    if canonical_classification_sha256(normalized) != baseline.get("sha256"):
        raise InventoryError("classification", "owner classification digest differs")


def load_manifest_payload(reference: dict, expected_kind: str) -> dict:
    relative = PurePosixPath(str(reference["path"]))
    if relative.is_absolute() or ".." in relative.parts:
        raise InventoryError("manifest", "referenced inventory file escapes inventory root")
    path = INVENTORY_ROOT.joinpath(*relative.parts)
    raw = path.read_bytes()
    if sha256_bytes(raw) != reference["sha256"]:
        raise InventoryError("manifest", f"referenced inventory file digest differs: {relative}")
    payload = json.loads(raw)
    if payload.get("schemaVersion") != 1 or payload.get("kind") != expected_kind:
        raise InventoryError("manifest", f"invalid referenced inventory file: {relative}")
    return payload


def load_manifest() -> dict:
    try:
        manifest = json.loads(MANIFEST_PATH.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as error:
        raise InventoryError("manifest", f"cannot load inventory manifest: {error}") from error
    if manifest.get("schemaVersion") != 1:
        raise InventoryError("manifest", "unsupported inventory manifest schema")
    external: list[dict] = []
    for reference in manifest.get("blockedOwnerFiles", []):
        payload = load_manifest_payload(
            reference, "haxe.go-official-inventory-blocked-issue-owners"
        )
        if not isinstance(payload.get("records"), list):
            raise InventoryError("manifest", "blocked-owner records must be an array")
        external.extend(payload["records"])
    manifest["blockedOwners"] = [*manifest.get("blockedOwners", []), *external]
    runtime = load_manifest_payload(
        manifest["activeRuntimeBaselineFile"],
        "haxe.go-official-inventory-active-runtime-baseline",
    )
    if not isinstance(runtime.get("records"), list):
        raise InventoryError("manifest", "active runtime baseline records must be an array")
    if canonical_runtime_records(runtime["records"]) != runtime["records"]:
        raise InventoryError("manifest", "active runtime baseline is not canonical")
    if canonical_runtime_sha256(runtime["records"]) != runtime.get("sha256"):
        raise InventoryError("manifest", "active runtime baseline digest differs internally")
    manifest["activeRuntimeBaseline"] = runtime
    return manifest


def classify_candidates(manifest: dict, candidates: list[dict]) -> tuple[list[dict], list[dict]]:
    by_path = {item["path"]: dict(item) for item in candidates}
    if len(by_path) != len(candidates):
        raise InventoryError("candidate-inventory", "candidate paths are not unique")

    explicit: dict[str, tuple[str, str]] = {}
    for status, field in (("blocked", "blockedOwners"), ("inapplicable", "inapplicableOwners")):
        for item in manifest[field]:
            path = str(item["path"])
            if path in explicit:
                raise InventoryError("classification", f"owner classified twice: {path}")
            explicit[path] = (status, str(item["reason"]))

    runnable: list[dict] = []
    classified: list[dict] = []
    for path, record in sorted(by_path.items()):
        if path in explicit:
            status, reason = explicit[path]
            record.update({"status": status, "reason": reason})
            classified.append(record)
        else:
            record.update({"status": "candidate", "reason": "selected for target execution"})
            runnable.append(record)
    unknown = sorted(set(explicit) - set(by_path))
    if unknown:
        raise InventoryError("classification", f"classified owners are absent upstream: {unknown}")
    return runnable, classified


def shard_candidates(manifest: dict, candidates: list[dict]) -> list[dict]:
    shards: list[dict] = []
    top = [item for item in candidates if item["family"] == "top-level"]
    top_size = int(manifest["shards"]["topOwnersPerShard"])
    for offset in range(0, len(top), top_size):
        shards.append(
            {
                "id": f"top-{offset // top_size + 1:02d}",
                "owners": top[offset : offset + top_size],
            }
        )

    unitstd_groups: dict[str, list[dict]] = defaultdict(list)
    for item in candidates:
        if item["family"] != "unitstd":
            continue
        path = item["path"]
        if path.startswith("unitstd/sys/net/") or path == "unitstd/Ssl.unit.hx":
            group = "network"
        elif path.startswith("unitstd/sys/"):
            group = "system"
        elif ".macro.unit.hx" in path:
            group = "macro"
        else:
            group = "core"
        unitstd_groups[group].append(item)
    for group in sorted(unitstd_groups):
        shards.append({"id": f"unitstd-{group}", "owners": unitstd_groups[group]})

    issue_size = int(manifest["shards"]["issueOwnersPerShard"])
    issues = [item for item in candidates if item["family"] == "issue"]
    for offset in range(0, len(issues), issue_size):
        shards.append(
            {
                "id": f"issues-{offset // issue_size + 1:02d}",
                "owners": issues[offset : offset + issue_size],
            }
        )
    hxcpp = [item for item in candidates if item["family"] == "hxcpp-issue"]
    if hxcpp:
        shards.append({"id": "hxcpp-issues", "owners": hxcpp})
    return shards


def split_shard(shard: dict) -> tuple[dict, dict]:
    owners = shard["owners"]
    if len(owners) < 2:
        raise InventoryError("shard", "cannot split a single-owner shard")
    middle = len(owners) // 2
    return (
        {"id": f"{shard['id']}-a", "owners": owners[:middle]},
        {"id": f"{shard['id']}-b", "owners": owners[middle:]},
    )


def write_cases_source(path: Path, owners: list[dict]) -> None:
    constructors = [
        f"\t\t\tnew {item['runtimeClass']}(),"
        for item in owners
    ]
    path.write_text(
        "class OfficialInventoryCases {\n"
        "\tpublic static function build():Array<OfficialInventoryCase> {\n"
        "\t\treturn [\n"
        + "\n".join(constructors)
        + "\n\t\t];\n\t}\n}\n",
        encoding="utf-8",
    )


def prepare_unitstd_tree(app: Path, owners: list[dict], haxe_root: Path) -> None:
    target = app / "official-unitstd"
    target.mkdir()
    source_root = haxe_root / "tests" / "unit" / "src"
    for owner in owners:
        if owner["family"] != "unitstd":
            continue
        relative = Path(owner["path"]).relative_to("unitstd")
        destination = target / relative
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.symlink_to(source_root / owner["path"])


def official_resource_arguments(owners: list[dict], haxe_root: Path) -> list[str]:
    paths = {owner["path"] for owner in owners}
    unit_root = haxe_root / "tests" / "unit"
    resources: list[tuple[str, str]] = []
    if "unit/TestResource.hx" in paths:
        resources.extend(
            [
                ("res1.txt", 're/s?!%[]))("\'1.txt'),
                ("res2.bin", 're/s?!%[]))("\'1.bin'),
            ]
        )
    if "unit/TestSerializerCrossTarget.hx" in paths:
        resources.append(("serializedValues.txt", "serializedValues.txt"))
    arguments: list[str] = []
    for source, name in resources:
        arguments.extend(["--resource", f"{unit_root / source}@{name}"])
    return arguments


def parse_runtime(stdout: str) -> tuple[list[dict], dict]:
    records: list[dict] = []
    summaries: list[dict] = []
    for line in stdout.splitlines():
        if RECORD_PREFIX in line:
            fields = line.split(RECORD_PREFIX, 1)[1].split("\t")
            if len(fields) != 3:
                raise InventoryError("runtime", f"invalid inventory record: {line}")
            records.append(
                {"id": fields[0], "status": fields[1], "assertions": int(fields[2])}
            )
        elif SUMMARY_PREFIX in line:
            fields = line.split(SUMMARY_PREFIX, 1)[1].split("\t")
            if len(fields) != 3:
                raise InventoryError("runtime", f"invalid inventory summary: {line}")
            summaries.append(
                {"completed": int(fields[0]), "expected": int(fields[1]), "failures": int(fields[2])}
            )
    if len(summaries) != 1:
        raise InventoryError("runtime", f"expected one shard summary, found {len(summaries)}")
    summary = summaries[0]
    if summary["completed"] != summary["expected"] or summary["failures"] != 0:
        raise InventoryError("runtime", f"shard did not complete successfully: {summary}")
    return records, summary


def redact(value: str, roots: list[Path]) -> str:
    result = value
    replacements: list[tuple[str, str]] = []
    for index, root in enumerate(roots):
        if root.resolve() == ROOT.resolve():
            label = "repository"
        elif root.name == "sandbox":
            label = "installed-package"
        elif root.name.startswith("haxe-go-official-inventory-"):
            label = "workspace"
        elif root.name.startswith("haxe-"):
            label = "haxe-upstream"
        elif root.name.startswith("utest-"):
            label = "utest-upstream"
        else:
            label = f"isolated-root-{index}"
        for spelling in {str(root.absolute()), str(root.resolve())}:
            replacements.append((spelling, label))
    # Replace nested roots before their parents so a package sandbox does not
    # inherit a random TemporaryDirectory suffix from the workspace label.
    for spelling, label in sorted(replacements, key=lambda item: len(item[0]), reverse=True):
        result = result.replace(spelling, f"<REDACTED:{label}>")
    return result


def run_shard(
    shard: dict,
    *,
    sandbox: Path,
    haxe_root: Path,
    utest_root: Path,
    environment: dict[str, str],
    artifact: Path,
    timeout: float,
) -> tuple[list[dict], dict]:
    shard_id = str(shard["id"])
    app = sandbox / f"inventory-{shard_id}"
    app.mkdir()
    SMOKE.verify_compile_package_resolution(app, sandbox, environment, timeout)
    adapter = prepare_test_base_adapter(app, haxe_root, shard["owners"])
    owner_adapters = prepare_owner_adapters(app, haxe_root, shard["owners"])
    unitstd_adapters = prepare_unitstd_adapters(app, shard["owners"])
    write_cases_source(app / "OfficialInventoryCases.hx", shard["owners"])
    prepare_unitstd_tree(app, shard["owners"], haxe_root)
    output = app / "generated"
    command = [
        SMOKE.require_tool("haxe"),
        "-cp", str(haxe_root / "tests" / "unit" / "src"),
        "-cp", str(utest_root / "src"),
        "-cp", str(ROOT / "test" / "official_haxe_target_smoke" / "src"),
        "-cp", str(INVENTORY_SOURCE),
        # Haxe gives the later class path precedence. Keep the temporary
        # provenance-recorded unit.Test adapter ahead of the upstream copy.
        "-cp", str(app),
        "-lib", "reflaxe.go",
        "-main", "OfficialInventoryMain",
        "-D", "reflaxe_go_profile=portable",
        "-D", f"go_output={output}",
        "-D", f"go_module=official_haxe_inventory_{shard_id.replace('-', '_')}",
        "-D", "go_no_build",
        "-D", "UTEST_FAILURE_THROW",
        "--dce", "full",
    ]
    command.extend(official_resource_arguments(shard["owners"], haxe_root))
    started = time.monotonic()
    haxe = SMOKE.run_checked(
        command, cwd=app, environment=environment, timeout=timeout, stage=f"haxe-{shard_id}"
    )
    go_results = SMOKE.go_format_build_test(output, environment=environment, timeout=timeout)
    runtime = SMOKE.execute_program(output, environment=environment, timeout=timeout)
    logs = artifact / "logs" / shard_id
    logs.mkdir(parents=True)
    redaction_roots = [ROOT, sandbox, haxe_root, utest_root]
    for name, process in [("haxe", haxe), *go_results.items(), ("runtime", runtime)]:
        (logs / f"{name}.stdout").write_text(redact(process.stdout, redaction_roots), encoding="utf-8")
        (logs / f"{name}.stderr").write_text(redact(process.stderr, redaction_roots), encoding="utf-8")
    if runtime.returncode != 0:
        detail = redact((runtime.stdout + runtime.stderr).strip(), redaction_roots)
        raise InventoryError("runtime", f"shard {shard_id} failed:\n{detail}")
    records, summary = parse_runtime(runtime.stdout)
    return records, {
        "id": shard_id,
        "owners": len(shard["owners"]),
        "activeTests": len(records),
        "assertions": sum(item["assertions"] for item in records),
        "elapsedSeconds": round(time.monotonic() - started, 3),
        "summary": summary,
        "testBaseAdapter": adapter,
        "ownerAdapters": owner_adapters,
        "unitstdAdapters": unitstd_adapters,
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--artifact-dir", type=Path, default=DEFAULT_CACHE / "artifacts")
    parser.add_argument("--checkout-cache", type=Path, default=DEFAULT_CACHE / "upstream")
    parser.add_argument("--haxe-checkout", type=Path)
    parser.add_argument("--utest-checkout", type=Path)
    parser.add_argument("--timeout", type=int, default=300)
    parser.add_argument("--propose-baseline", action="store_true")
    parser.add_argument(
        "--discover-blockers",
        action="store_true",
        help="Bisect rejected shards to exact owners and emit only a blocker proposal",
    )
    parser.add_argument("--only-shard")
    parser.add_argument(
        "--owner-regex",
        help="Diagnostic-only filter over candidate source paths",
    )
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    artifact = args.artifact_dir.resolve()
    try:
        manifest = load_manifest()
        haxe_root = SMOKE.ensure_checkout(
            manifest["upstream"]["haxe"], args.haxe_checkout, args.checkout_cache, "haxe"
        )
        utest_root = SMOKE.ensure_checkout(
            manifest["upstream"]["utest"], args.utest_checkout, args.checkout_cache, "utest"
        )
        SMOKE.verify_upstream(
            {"upstream": manifest["upstream"], "activeSmokeRecords": []}, haxe_root, utest_root
        )
        candidates = discover_candidate_inventory(haxe_root / "tests" / "unit" / "src")
        counts = dict(sorted(Counter(item["family"] for item in candidates).items()))
        expected = manifest["candidateInventory"]
        if counts != expected["counts"] or canonical_inventory_sha256(candidates) != expected["sha256"]:
            raise InventoryError("candidate-inventory", "pinned upstream candidate inventory drifted")
        runnable, classified = classify_candidates(manifest, candidates)
        diagnostic_selection = bool(args.owner_regex or args.only_shard)
        if args.owner_regex:
            try:
                owner_pattern = re.compile(args.owner_regex)
            except re.error as error:
                raise InventoryError("selection", f"invalid owner regex: {error}") from error
            runnable = [item for item in runnable if owner_pattern.search(item["path"])]
            if not runnable:
                raise InventoryError("selection", "owner regex selected no runnable candidates")
        shards = shard_candidates(manifest, runnable)
        if args.only_shard:
            shards = [item for item in shards if item["id"] == args.only_shard]
            if not shards:
                raise InventoryError("shard", f"unknown shard: {args.only_shard}")

        if artifact.exists():
            shutil.rmtree(artifact)
        artifact.mkdir(parents=True)
        started = time.monotonic()
        with tempfile.TemporaryDirectory(prefix="haxe-go-official-inventory-") as raw:
            workspace = Path(raw)
            environment = SMOKE.deterministic_environment(workspace)
            toolchains = SMOKE.toolchain_identity(environment)
            sandbox, package_artifact = SMOKE.prepare_installed_package(
                workspace, environment, args.timeout
            )
            active: list[dict] = []
            shard_evidence: list[dict] = []
            failures: list[dict] = []
            active_by_class: dict[str, list[dict]] = defaultdict(list)
            pending = list(shards)
            while pending:
                shard = pending.pop(0)
                try:
                    records, evidence = run_shard(
                        shard,
                        sandbox=sandbox,
                        haxe_root=haxe_root,
                        utest_root=utest_root,
                        environment=environment,
                        artifact=artifact,
                        timeout=args.timeout,
                    )
                except (InventoryError, SMOKE.SmokeError) as error:
                    if not args.discover_blockers:
                        raise
                    if len(shard["owners"]) > 1:
                        left, right = split_shard(shard)
                        pending[0:0] = [left, right]
                        continue
                    redacted = redact(str(error), [ROOT, workspace, sandbox, haxe_root, utest_root])
                    failures.append(
                        {
                            "owner": shard["owners"][0]["path"],
                            "stage": getattr(error, "stage", "inventory"),
                            "detail": "\n".join(redacted.splitlines()[-60:]),
                        }
                    )
                    continue
                active.extend(records)
                shard_evidence.append(evidence)
                for record in records:
                    active_by_class[record["id"].rsplit(".", 1)[0]].append(record)

            if failures:
                rejected = {
                    "schemaVersion": 1,
                    "kind": "haxe.go-official-inventory-blocker-proposal",
                    "claimEvidence": False,
                    "source": SMOKE.source_identity(),
                    "upstream": manifest["upstream"],
                    "toolchains": toolchains,
                    "candidateInventory": {"counts": counts, "sha256": expected["sha256"]},
                    "failures": failures,
                    "passingSplitShards": shard_evidence,
                    "partialRuntimeRecords": canonical_runtime_records(active),
                }
                (artifact / "blocker-proposal.json").write_text(
                    json.dumps(rejected, indent=2, sort_keys=True) + "\n", encoding="utf-8"
                )
                SMOKE.verify_artifact_path_confinement(
                    artifact, [ROOT, workspace, sandbox, haxe_root, utest_root]
                )
                print(
                    f"Official inventory blocker discovery rejected {len(failures)} exact owners; "
                    f"proposal: {artifact / 'blocker-proposal.json'}",
                    file=sys.stderr,
                )
                for failure in failures:
                    print(
                        f"- {failure['owner']} [{failure['stage']}]: "
                        f"{failure['detail'].splitlines()[-1] if failure['detail'] else 'failed'}",
                        file=sys.stderr,
                    )
                return 2

            for owner in runnable:
                records = active_by_class.get(owner["runtimeClass"], [])
                if records:
                    owner.update(
                        {
                            "status": "active",
                            "reason": "positive target-runtime assertions",
                            "activeTests": [record["id"] for record in records],
                            "assertions": sum(record["assertions"] for record in records),
                        }
                    )
                else:
                    owner.update(
                        {
                            "status": "inapplicable",
                            "reason": "no active utest fixture under Go target defines",
                        }
                    )
                classified.append(owner)

            normalized = canonical_runtime_records(active)
            proposal = {
                "records": normalized,
                "sha256": canonical_runtime_sha256(normalized),
            }
            classification_proposal = {
                "counts": dict(
                    sorted(Counter(item["status"] for item in classified).items())
                ),
                "sha256": canonical_classification_sha256(classified),
            }
            if not args.propose_baseline and not diagnostic_selection:
                validate_runtime_baseline(manifest["activeRuntimeBaseline"], active)
                validate_classification_baseline(
                    manifest["classificationBaseline"], classified
                )

            result = {
                "schemaVersion": 1,
                "kind": "haxe.go-official-haxe-target-inventory",
                "claimEvidence": not args.propose_baseline and not diagnostic_selection,
                "diagnosticSelection": {
                    "active": diagnostic_selection,
                    "onlyShard": args.only_shard,
                    "ownerRegex": args.owner_regex,
                },
                "claim": manifest["claim"],
                "productSurface": manifest["productSurface"],
                "source": SMOKE.source_identity(),
                "upstream": manifest["upstream"],
                "toolchains": toolchains,
                "packageArtifact": package_artifact,
                "testBaseAdapters": [
                    {"shard": item["id"], **item["testBaseAdapter"]}
                    for item in shard_evidence
                ],
                "candidateInventory": {"counts": counts, "sha256": expected["sha256"]},
                "classification": sorted(classified, key=lambda item: (item["family"], item["path"])),
                "classificationCounts": dict(sorted(Counter(item["status"] for item in classified).items())),
                "classificationBaselineProposal": classification_proposal,
                "activeRuntimeBaselineProposal": proposal,
                "shards": shard_evidence,
                "elapsedSeconds": round(time.monotonic() - started, 3),
            }
            (artifact / "result.json").write_text(
                json.dumps(result, indent=2, sort_keys=True) + "\n", encoding="utf-8"
            )
            (artifact / "baseline-proposal.json").write_text(
                json.dumps(proposal, indent=2, sort_keys=True) + "\n", encoding="utf-8"
            )
            SMOKE.verify_artifact_path_confinement(
                artifact, [ROOT, workspace, sandbox, haxe_root, utest_root]
            )
        print(
            f"Official Haxe target inventory passed: {len(active)} active tests, "
            f"{len(classified)} classified owners, {len(shards)} shards"
        )
        return 0
    except (InventoryError, SMOKE.SmokeError) as error:
        shutil.rmtree(artifact, ignore_errors=True)
        stage = getattr(error, "stage", "inventory")
        print(f"[official-haxe-target-inventory] ERROR [{stage}]: {error}", file=sys.stderr)
        return 2
    except (OSError, json.JSONDecodeError, ValueError) as error:
        shutil.rmtree(artifact, ignore_errors=True)
        print(f"[official-haxe-target-inventory] ERROR [environment]: {error}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
