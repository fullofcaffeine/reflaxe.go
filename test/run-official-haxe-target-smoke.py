#!/usr/bin/env python3

"""Run a pinned, active three-family official Haxe target smoke on Go."""

from __future__ import annotations

import argparse
import hashlib
import json
import os
from pathlib import Path, PurePosixPath
import shutil
import subprocess
import sys
import tempfile
import time


ROOT = Path(__file__).resolve().parent.parent
SMOKE_ROOT = ROOT / "test" / "official_haxe_target_smoke"
MANIFEST_PATH = SMOKE_ROOT / "manifest.json"
DEFAULT_CACHE = ROOT / ".cache" / "official-haxe-target-smoke"
RECORD_PREFIX = "OFFICIAL_HAXE_SMOKE_RECORD\t"
SUMMARY_PREFIX = "OFFICIAL_HAXE_SMOKE_SUMMARY\t"


class SmokeError(RuntimeError):
    def __init__(self, stage: str, message: str) -> None:
        super().__init__(message)
        self.stage = stage


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--artifact-dir", type=Path, default=DEFAULT_CACHE / "artifacts")
    parser.add_argument("--checkout-cache", type=Path, default=DEFAULT_CACHE / "upstream")
    parser.add_argument("--haxe-checkout", type=Path)
    parser.add_argument("--utest-checkout", type=Path)
    parser.add_argument("--timeout", type=int, default=180)
    parser.add_argument(
        "--verify-failure-propagation",
        action="store_true",
        help="Prove assertion, Go-build, runtime, timeout, and missing-source failures",
    )
    return parser.parse_args()


def load_manifest() -> dict:
    try:
        value = json.loads(MANIFEST_PATH.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as error:
        raise SmokeError("manifest", f"cannot load {MANIFEST_PATH}: {error}") from error
    if value.get("schemaVersion") != 1:
        raise SmokeError("manifest", "unsupported official Haxe smoke manifest schema")
    return value


def sha256_bytes(value: bytes) -> str:
    return hashlib.sha256(value).hexdigest()


def sha256_file(path: Path) -> str:
    return sha256_bytes(path.read_bytes())


def require_tool(name: str) -> str:
    executable = shutil.which(name)
    if executable is None:
        raise SmokeError("environment", f"required executable is unavailable: {name}")
    return executable


def run_process(
    command: list[str],
    *,
    cwd: Path,
    environment: dict[str, str],
    timeout: float,
) -> subprocess.CompletedProcess[str]:
    try:
        return subprocess.run(
            command,
            cwd=cwd,
            env=environment,
            capture_output=True,
            text=True,
            timeout=timeout,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        raise SmokeError("command", f"cannot run {' '.join(command)}: {error}") from error


def run_checked(
    command: list[str],
    *,
    cwd: Path,
    environment: dict[str, str],
    timeout: float,
    stage: str,
) -> subprocess.CompletedProcess[str]:
    process = run_process(command, cwd=cwd, environment=environment, timeout=timeout)
    if process.returncode != 0:
        detail = (process.stdout + process.stderr).strip()
        raise SmokeError(stage, f"command failed ({' '.join(command)}):\n{detail}")
    return process


def git_text(checkout: Path, *args: str) -> str:
    process = run_checked(
        [require_tool("git"), *args],
        cwd=checkout,
        environment=dict(os.environ),
        timeout=120,
        stage="provenance",
    )
    return process.stdout.strip()


def ensure_checkout(spec: dict, explicit: Path | None, cache_root: Path, name: str) -> Path:
    commit = str(spec["commit"])
    if explicit is not None:
        checkout = explicit.resolve()
    else:
        checkout = (cache_root / f"{name}-{commit[:12]}").resolve()
        if not checkout.exists():
            checkout.parent.mkdir(parents=True, exist_ok=True)
            process = run_process(
                [
                    require_tool("git"),
                    "clone",
                    "--filter=blob:none",
                    "--no-checkout",
                    str(spec["repository"]),
                    str(checkout),
                ],
                cwd=ROOT,
                environment=dict(os.environ),
                timeout=300,
            )
            if process.returncode != 0:
                raise SmokeError("provenance", process.stdout + process.stderr)
        if not (checkout / ".git").is_dir():
            raise SmokeError("provenance", f"checkout cache is not a Git repository: {checkout}")
        fetch = run_process(
            [require_tool("git"), "fetch", "--depth", "1", "origin", commit],
            cwd=checkout,
            environment=dict(os.environ),
            timeout=300,
        )
        if fetch.returncode != 0:
            raise SmokeError("provenance", fetch.stdout + fetch.stderr)
        run_checked(
            [require_tool("git"), "checkout", "--detach", "--force", commit],
            cwd=checkout,
            environment=dict(os.environ),
            timeout=180,
            stage="provenance",
        )

    if not checkout.is_dir() or not (checkout / ".git").exists():
        raise SmokeError("provenance", f"missing exact checkout: {checkout}")
    if git_text(checkout, "rev-parse", "HEAD") != commit:
        raise SmokeError("provenance", f"{name} checkout is not at {commit}")
    if git_text(checkout, "status", "--porcelain"):
        raise SmokeError("provenance", f"{name} checkout is dirty: {checkout}")
    return checkout


def verify_upstream(manifest: dict, haxe_root: Path, utest_root: Path) -> list[dict]:
    haxe = manifest["upstream"]["haxe"]
    utest = manifest["upstream"]["utest"]
    haxe_evidence = []
    for record in haxe.get(
        "licenseEvidence",
        [{"path": haxe["licensePath"], "sha256": haxe["licenseSha256"]}],
    ):
        relative = PurePosixPath(str(record["path"]))
        if relative.is_absolute() or ".." in relative.parts:
            raise SmokeError("provenance", f"license evidence escapes checkout: {relative}")
        haxe_evidence.append(
            {
                "kind": "haxe-license",
                "path": str(relative),
                "sha256": sha256_file(haxe_root.joinpath(*relative.parts)),
                "expectedSha256": record["sha256"],
            }
        )
    evidence = [
        *haxe_evidence,
        {
            "kind": "utest-license-declaration",
            "path": utest["licenseEvidencePath"],
            "sha256": sha256_file(utest_root / utest["licenseEvidencePath"]),
            "expectedSha256": utest["licenseEvidenceSha256"],
        },
    ]
    for item in evidence:
        if item["sha256"] != item["expectedSha256"]:
            raise SmokeError("provenance", f"upstream license evidence drift: {item['path']}")

    for record in manifest["activeSmokeRecords"]:
        path = haxe_root / record["upstreamPath"]
        if not path.is_file():
            raise SmokeError("provenance", f"missing selected upstream source: {record['upstreamPath']}")
        if sha256_file(path) != record["sha256"]:
            raise SmokeError("provenance", f"selected upstream source drift: {record['upstreamPath']}")
    return evidence


def deterministic_environment(sandbox: Path) -> dict[str, str]:
    environment = dict(os.environ)
    for name in ("go-cache", "go-path"):
        (sandbox / name).mkdir(parents=True, exist_ok=True)
    environment.update(
        {
            "GOCACHE": str(sandbox / "go-cache"),
            "GOPATH": str(sandbox / "go-path"),
            "HAXE_NO_SERVER": "1",
            "LC_ALL": "C",
            "NO_COLOR": "1",
            "PYTHONHASHSEED": "0",
            "TZ": "UTC",
        }
    )
    return environment


def prepare_installed_package(
    workspace: Path, environment: dict[str, str], timeout: float
) -> tuple[Path, dict]:
    archive = workspace / "reflaxe.go.zip"
    run_checked(
        ["bash", "scripts/release/package-haxelib.sh", str(archive)],
        cwd=ROOT,
        environment=environment,
        timeout=timeout,
        stage="package",
    )
    sandbox = workspace / "sandbox"
    sandbox.mkdir()
    haxelib = require_tool("haxelib")
    run_checked(
        [haxelib, "newrepo"],
        cwd=sandbox,
        environment=environment,
        timeout=timeout,
        stage="install",
    )
    run_checked(
        [haxelib, "install", str(archive), "--always", "--skip-dependencies"],
        cwd=sandbox,
        environment=environment,
        timeout=timeout,
        stage="install",
    )
    installed = run_checked(
        [haxelib, "path", "reflaxe.go"],
        cwd=sandbox,
        environment=environment,
        timeout=timeout,
        stage="install",
    ).stdout
    if str(ROOT.resolve()) in installed or str((sandbox / ".haxelib").resolve()) not in installed:
        raise SmokeError("install", "reflaxe.go did not resolve from the isolated Haxelib repository")
    return sandbox, {
        "kind": "deterministic-haxelib-zip",
        "sha256": sha256_file(archive),
        "bytes": archive.stat().st_size,
    }


def verify_compile_package_resolution(
    app: Path,
    sandbox: Path,
    environment: dict[str, str],
    timeout: float,
) -> None:
    """Prove that the compile working directory resolves the packaged target."""
    try:
        app.resolve().relative_to(sandbox.resolve())
    except ValueError as error:
        raise SmokeError(
            "install",
            "the Haxe compile directory is outside the isolated Haxelib repository tree",
        ) from error

    installed = run_checked(
        [require_tool("haxelib"), "path", "reflaxe.go"],
        cwd=app,
        environment=environment,
        timeout=timeout,
        stage="install",
    ).stdout
    isolated_repository = str((sandbox / ".haxelib").resolve())
    if str(ROOT.resolve()) in installed or isolated_repository not in installed:
        raise SmokeError(
            "install",
            "the Haxe compile directory did not resolve reflaxe.go from its isolated repository",
        )


ASSERTION_ADAPTER_METHODS = """
\n\tpublic function new() {}

\tpublic function eq<T>(expected:T, actual:T, ?pos:haxe.PosInfos):Void {
\t\tutest.Assert.equals(expected, actual, pos);
\t}

\tpublic function t(value:Bool, ?pos:haxe.PosInfos):Void {
\t\tutest.Assert.isTrue(value, pos);
\t}

\tpublic function f(value:Bool, ?pos:haxe.PosInfos):Void {
\t\tutest.Assert.isFalse(value, pos);
\t}
"""


def adapt_official_class_source(
    source: Path,
    destination: Path,
    *,
    package_declaration: str,
    class_declaration: str,
    adapted_declaration: str,
) -> None:
    """Change only module/class scaffolding; retain the pinned official method bodies."""
    text = source.read_text(encoding="utf-8")
    if text.count(package_declaration) != 1 or text.count(class_declaration) != 1:
        raise SmokeError("adapter", f"official class scaffolding drift: {source}")
    text = text.replace(package_declaration, "package;", 1)
    text = text.replace(class_declaration, adapted_declaration, 1)
    closing = text.rfind("}")
    if closing < 0:
        raise SmokeError("adapter", f"official class has no closing brace: {source}")
    destination.write_text(
        text[:closing] + ASSERTION_ADAPTER_METHODS + text[closing:], encoding="utf-8"
    )


def prepare_official_class_adapters(app: Path, haxe_root: Path) -> Path:
    adapted = app / "adapted-official"
    adapted.mkdir()
    adapt_official_class_source(
        haxe_root / "tests/unit/src/unit/TestNumericSuffixes.hx",
        adapted / "OfficialNumericSuffixCase.hx",
        package_declaration="package unit;",
        class_declaration="class TestNumericSuffixes extends Test",
        adapted_declaration="class OfficialNumericSuffixCase",
    )
    adapt_official_class_source(
        haxe_root / "tests/unit/src/unit/issues/Issue6369.hx",
        adapted / "OfficialIssue6369Case.hx",
        package_declaration="package unit.issues;",
        class_declaration="class Issue6369 extends unit.Test",
        adapted_declaration="class OfficialIssue6369Case",
    )
    return adapted


def compile_program(
    *,
    sandbox: Path,
    haxe_root: Path,
    utest_root: Path,
    environment: dict[str, str],
    timeout: float,
    name: str,
    defines: list[str] | None = None,
) -> tuple[Path, subprocess.CompletedProcess[str]]:
    app = sandbox / name
    app.mkdir()
    verify_compile_package_resolution(app, sandbox, environment, timeout)
    adapted = prepare_official_class_adapters(app, haxe_root)
    unitstd = app / "official-unitstd"
    unitstd.mkdir()
    selected_spec = haxe_root / "tests" / "unit" / "src" / "unitstd" / "IntIterator.unit.hx"
    (unitstd / selected_spec.name).symlink_to(selected_spec)
    output = app / "generated"
    command = [
        require_tool("haxe"),
        "-cp",
        str(haxe_root / "tests" / "unit" / "src"),
        "-cp",
        str(utest_root / "src"),
        "-cp",
        str(SMOKE_ROOT / "src"),
        "-cp",
        str(adapted),
        "-lib",
        "reflaxe.go",
        "-main",
        "OfficialTargetSmokeMain",
        "-D",
        "reflaxe_go_profile=portable",
        "-D",
        f"go_output={output}",
        "-D",
        f"go_module=official_haxe_target_smoke_{name.replace('-', '_')}",
        "-D",
        "go_no_build",
        "-D",
        "UTEST_FAILURE_THROW",
        "-D",
        "official_haxe_smoke_unitstd_path=official-unitstd/IntIterator.unit.hx",
        "--dce",
        "full",
        "-D",
        "analyzer-optimize",
        "-D",
        "analyzer-user-var-fusion",
    ]
    for define in defines or []:
        command.extend(["-D", define])
    process = run_checked(
        command,
        cwd=app,
        environment=environment,
        timeout=timeout,
        stage=f"haxe-{name}",
    )
    if str(ROOT.resolve()).encode() in b"".join(path.read_bytes() for path in output.rglob("*.go")):
        raise SmokeError("confinement", "generated Go contains an absolute repository path")
    return output, process


def go_format_build_test(
    output: Path,
    *,
    environment: dict[str, str],
    timeout: float,
) -> dict[str, subprocess.CompletedProcess[str]]:
    go_files = sorted(str(path) for path in output.rglob("*.go"))
    if not go_files:
        raise SmokeError("go", "the official target smoke generated no Go source")
    results = {
        "gofmt": run_checked(
            [require_tool("gofmt"), "-w", *go_files],
            cwd=output,
            environment=environment,
            timeout=timeout,
            stage="gofmt",
        ),
        "go-test": run_checked(
            [require_tool("go"), "test", "./..."],
            cwd=output,
            environment=environment,
            timeout=timeout,
            stage="go-test",
        ),
        "go-build": run_checked(
            [require_tool("go"), "build", "./..."],
            cwd=output,
            environment=environment,
            timeout=timeout,
            stage="go-build",
        ),
    }
    return results


def execute_program(
    output: Path,
    *,
    environment: dict[str, str],
    timeout: float,
) -> subprocess.CompletedProcess[str]:
    return run_process(
        [require_tool("go"), "run", "."],
        cwd=output,
        environment=environment,
        timeout=timeout,
    )


def parse_runtime_records(stdout: str) -> tuple[list[dict], dict]:
    records: list[dict] = []
    summaries: list[dict] = []
    for line in stdout.splitlines():
        record_at = line.find(RECORD_PREFIX)
        summary_at = line.find(SUMMARY_PREFIX)
        if record_at >= 0:
            fields = line[record_at + len(RECORD_PREFIX) :].split("\t")
            if len(fields) != 3:
                raise SmokeError("runtime", f"invalid runtime record: {line}")
            records.append(
                {"id": fields[0], "status": fields[1], "assertions": int(fields[2])}
            )
        elif summary_at >= 0:
            fields = line[summary_at + len(SUMMARY_PREFIX) :].split("\t")
            if len(fields) != 3:
                raise SmokeError("runtime", f"invalid runtime summary: {line}")
            summaries.append(
                {"completed": int(fields[0]), "expected": int(fields[1]), "status": fields[2]}
            )
    if len(summaries) != 1:
        raise SmokeError("runtime", f"expected one runtime summary, found {len(summaries)}")
    return records, summaries[0]


def validate_active_inventory(manifest: dict, records: list[dict], summary: dict) -> None:
    expected: dict[str, int] = {}
    for source in manifest["activeSmokeRecords"]:
        for test_id in source["expectedActiveTests"]:
            expected[test_id] = int(source["minimumAssertions"])
    actual = {str(record.get("id")): record for record in records}
    if set(actual) != set(expected):
        raise SmokeError(
            "active-inventory",
            f"active test drift: missing={sorted(set(expected) - set(actual))}, "
            f"added={sorted(set(actual) - set(expected))}",
        )
    for test_id, minimum in expected.items():
        record = actual[test_id]
        if record.get("status") != "pass" or int(record.get("assertions", 0)) < minimum:
            raise SmokeError("active-inventory", f"inactive, dummy, or failed runtime record: {record}")
    if summary.get("status") != "pass" or summary.get("completed") != len(expected):
        raise SmokeError("active-inventory", f"invalid runtime summary: {summary}")


def write_process_logs(artifact: Path, name: str, process: subprocess.CompletedProcess[str]) -> None:
    logs = artifact / "logs"
    logs.mkdir(exist_ok=True)
    (logs / f"{name}.stdout").write_text(process.stdout, encoding="utf-8")
    (logs / f"{name}.stderr").write_text(process.stderr, encoding="utf-8")


def verify_artifact_path_confinement(artifact: Path, forbidden_roots: list[Path]) -> None:
    """Reject machine-local paths anywhere in the evidence uploaded by CI.

    Generated Go is only one artifact member. Successful tools may print warnings
    to stdout/stderr, and result or inventory files can also retain an ephemeral
    checkout path. Scan bytes so non-UTF-8 output cannot bypass the boundary.
    """
    forbidden: set[bytes] = set()
    for path in forbidden_roots:
        if not str(path).strip():
            continue
        # macOS commonly exposes both /var/... and /private/var/... spellings.
        forbidden.add(str(path.absolute()).encode())
        forbidden.add(str(path.resolve()).encode())
    for path in sorted(artifact.rglob("*")):
        if not path.is_file():
            continue
        content = path.read_bytes()
        for value in forbidden:
            if value and value in content:
                relative = path.relative_to(artifact).as_posix()
                # Both required workflows upload with `if: always()`. Remove the
                # whole untrusted bundle before raising so the upload step cannot
                # publish the very path the confinement gate rejected.
                shutil.rmtree(artifact, ignore_errors=True)
                raise SmokeError(
                    "confinement",
                    f"uploaded artifact {relative} contains an ephemeral absolute path",
                )


def generated_inventory(output: Path) -> list[dict]:
    return [
        {
            "path": path.relative_to(output).as_posix(),
            "sha256": sha256_file(path),
            "bytes": path.stat().st_size,
        }
        for path in sorted(output.rglob("*"))
        if path.is_file()
    ]


def verify_missing_selected_source_control(
    manifest: dict,
    haxe_root: Path,
    utest_root: Path,
    workspace: Path,
) -> dict:
    """Remove one selected test while preserving every earlier provenance check."""
    missing_root = workspace / "missing-selected-source"
    haxe = manifest["upstream"]["haxe"]
    for record in haxe.get(
        "licenseEvidence",
        [{"path": haxe["licensePath"], "sha256": haxe["licenseSha256"]}],
    ):
        evidence_path = Path(record["path"])
        (missing_root / evidence_path).parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(haxe_root / evidence_path, missing_root / evidence_path)

    records = manifest["activeSmokeRecords"]
    if not records:
        raise SmokeError("failure-propagation", "missing-source control has no selected source")
    for record in records:
        relative = Path(record["upstreamPath"])
        destination = missing_root / relative
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(haxe_root / relative, destination)

    missing_selected_source = missing_root / records[0]["upstreamPath"]
    missing_selected_source.unlink()
    try:
        verify_upstream(manifest, missing_root, utest_root)
        detected = False
    except SmokeError as error:
        detected = error.stage == "provenance" and "missing selected upstream source" in str(error)
    return {
        "kind": "missing",
        "stage": "selected-source",
        "observedNonzero": 2,
        "detected": detected,
    }


def verify_failure_propagation(
    *,
    workspace: Path,
    sandbox: Path,
    haxe_root: Path,
    utest_root: Path,
    environment: dict[str, str],
    timeout: float,
    success_output: Path,
) -> list[dict]:
    evidence: list[dict] = []

    assertion_output, _ = compile_program(
        sandbox=sandbox,
        haxe_root=haxe_root,
        utest_root=utest_root,
        environment=environment,
        timeout=timeout,
        name="failure-assertion",
        defines=["official_haxe_smoke_assertion_failure"],
    )
    go_format_build_test(assertion_output, environment=environment, timeout=timeout)
    assertion = execute_program(assertion_output, environment=environment, timeout=timeout)
    assertion_detected = (
        assertion.returncode != 0
        and "OFFICIAL_HAXE_SMOKE_CONTROL\tassertion" in assertion.stdout
    )
    evidence.append({"kind": "assertion", "observedNonzero": assertion.returncode, "detected": assertion_detected})

    broken_build = workspace / "failure-go-build"
    shutil.copytree(success_output, broken_build)
    marker = broken_build / "official_smoke_deliberate_build_failure.go"
    marker.write_text("package main\nfunc official smoke deliberate build failure\n", encoding="utf-8")
    build = run_process(
        [require_tool("go"), "test", "./..."],
        cwd=broken_build,
        environment=environment,
        timeout=timeout,
    )
    build_detected = build.returncode != 0 and marker.name in (build.stdout + build.stderr)
    evidence.append({"kind": "go-build", "observedNonzero": build.returncode, "detected": build_detected})

    runtime_output, _ = compile_program(
        sandbox=sandbox,
        haxe_root=haxe_root,
        utest_root=utest_root,
        environment=environment,
        timeout=timeout,
        name="failure-runtime",
        defines=["official_haxe_smoke_runtime_failure"],
    )
    go_format_build_test(runtime_output, environment=environment, timeout=timeout)
    runtime = execute_program(runtime_output, environment=environment, timeout=timeout)
    runtime_detected = (
        runtime.returncode != 0
        and "OFFICIAL_HAXE_SMOKE_CONTROL\truntime" in runtime.stdout
    )
    evidence.append({"kind": "runtime", "observedNonzero": runtime.returncode, "detected": runtime_detected})

    timeout_output, _ = compile_program(
        sandbox=sandbox,
        haxe_root=haxe_root,
        utest_root=utest_root,
        environment=environment,
        timeout=timeout,
        name="failure-timeout",
        defines=["official_haxe_smoke_timeout_failure"],
    )
    go_format_build_test(timeout_output, environment=environment, timeout=timeout)
    timeout_binary = timeout_output / "official-smoke-timeout-control"
    run_checked(
        [require_tool("go"), "build", "-o", str(timeout_binary), "."],
        cwd=timeout_output,
        environment=environment,
        timeout=timeout,
        stage="failure-timeout-build",
    )
    try:
        subprocess.run(
            [str(timeout_binary)],
            cwd=timeout_output,
            env=environment,
            capture_output=True,
            text=True,
            timeout=0.25,
        )
        timeout_detected = False
    except subprocess.TimeoutExpired as error:
        captured = error.stdout or ""
        if isinstance(captured, bytes):
            captured = captured.decode("utf-8", errors="replace")
        timeout_detected = "OFFICIAL_HAXE_SMOKE_CONTROL\ttimeout" in captured
    evidence.append({"kind": "timeout", "observedNonzero": 124, "detected": timeout_detected})

    evidence.append(
        verify_missing_selected_source_control(
            load_manifest(), haxe_root, utest_root, workspace
        )
    )

    failures = [item["kind"] for item in evidence if not item["detected"] or item["observedNonzero"] == 0]
    if failures:
        raise SmokeError("failure-propagation", f"negative controls were not detected: {failures}")
    return evidence


def source_identity() -> dict:
    git = require_tool("git")
    head = run_checked(
        [git, "rev-parse", "HEAD"],
        cwd=ROOT,
        environment=dict(os.environ),
        timeout=30,
        stage="source-identity",
    ).stdout.strip()
    status = run_checked(
        [git, "status", "--porcelain=v1"],
        cwd=ROOT,
        environment=dict(os.environ),
        timeout=30,
        stage="source-identity",
    ).stdout
    return {
        "commit": head,
        "dirty": bool(status),
        "dirtyStatusSha256": sha256_bytes(status.encode("utf-8")) if status else None,
    }


def toolchain_identity(environment: dict[str, str]) -> dict:
    haxe = run_checked(
        [require_tool("haxe"), "--version"],
        cwd=ROOT,
        environment=environment,
        timeout=30,
        stage="toolchain",
    ).stdout.strip()
    go = run_checked(
        [require_tool("go"), "version"],
        cwd=ROOT,
        environment=environment,
        timeout=30,
        stage="toolchain",
    ).stdout.strip()
    if haxe != "4.3.7":
        raise SmokeError("toolchain", f"official Haxe 4.3.7 smoke requires Haxe 4.3.7, found {haxe}")
    return {"haxe": haxe, "go": go}


def main() -> int:
    args = parse_args()
    started = time.monotonic()
    artifact = args.artifact_dir.resolve()
    try:
        manifest = load_manifest()
        cache_root = args.checkout_cache.resolve()
        haxe_root = ensure_checkout(
            manifest["upstream"]["haxe"], args.haxe_checkout, cache_root, "haxe"
        )
        utest_root = ensure_checkout(
            manifest["upstream"]["utest"], args.utest_checkout, cache_root, "utest"
        )
        provenance = verify_upstream(manifest, haxe_root, utest_root)

        if artifact.exists():
            shutil.rmtree(artifact)
        artifact.mkdir(parents=True)
        with tempfile.TemporaryDirectory(prefix="haxe-go-official-smoke-") as raw:
            workspace = Path(raw)
            environment = deterministic_environment(workspace)
            toolchains = toolchain_identity(environment)
            sandbox, package_artifact = prepare_installed_package(
                workspace, environment, args.timeout
            )
            output, haxe_process = compile_program(
                sandbox=sandbox,
                haxe_root=haxe_root,
                utest_root=utest_root,
                environment=environment,
                timeout=args.timeout,
                name="success",
            )
            go_results = go_format_build_test(output, environment=environment, timeout=args.timeout)
            runtime = execute_program(output, environment=environment, timeout=args.timeout)
            write_process_logs(artifact, "haxe", haxe_process)
            for name, process in go_results.items():
                write_process_logs(artifact, name, process)
            write_process_logs(artifact, "runtime", runtime)
            if runtime.returncode != 0:
                raise SmokeError("runtime", runtime.stdout + runtime.stderr)
            records, summary = parse_runtime_records(runtime.stdout)
            validate_active_inventory(manifest, records, summary)

            failure_evidence = []
            if args.verify_failure_propagation:
                failure_evidence = verify_failure_propagation(
                    workspace=workspace,
                    sandbox=sandbox,
                    haxe_root=haxe_root,
                    utest_root=utest_root,
                    environment=environment,
                    timeout=args.timeout,
                    success_output=output,
                )

            shutil.copytree(output, artifact / "generated")
            result = {
                "schemaVersion": 1,
                "kind": "haxe.go-official-haxe-target-smoke",
                "claim": manifest["claim"],
                "source": source_identity(),
                "upstream": {
                    "haxe": manifest["upstream"]["haxe"],
                    "utest": manifest["upstream"]["utest"],
                    "licenseEvidence": provenance,
                },
                "toolchains": toolchains,
                "packageArtifact": package_artifact,
                "packageResolution": {
                    "mode": "isolated-haxelib-repository",
                    "compileWorkingDirectory": "sandbox/success",
                    "sourceCheckoutExcluded": True,
                },
                "activeRuntimeRecords": records,
                "runtimeSummary": summary,
                "failurePropagation": failure_evidence,
                "generatedFiles": generated_inventory(output),
                "elapsedSeconds": round(time.monotonic() - started, 3),
            }
            (artifact / "result.json").write_text(
                json.dumps(result, indent=2, sort_keys=True) + "\n", encoding="utf-8"
            )
            verify_artifact_path_confinement(
                artifact,
                [ROOT, workspace, sandbox, haxe_root, utest_root],
            )
        print(
            f"Official Haxe target smoke passed: {len(result['activeRuntimeRecords'])} "
            f"active tests, artifacts={artifact}"
        )
        return 0
    except SmokeError as error:
        shutil.rmtree(artifact, ignore_errors=True)
        print(f"[official-haxe-target-smoke] ERROR [{error.stage}]: {error}", file=sys.stderr)
        return 2
    except (OSError, json.JSONDecodeError) as error:
        shutil.rmtree(artifact, ignore_errors=True)
        print(f"[official-haxe-target-smoke] ERROR [environment]: {error}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
