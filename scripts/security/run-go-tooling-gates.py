#!/usr/bin/env python3

from __future__ import annotations

import json
import os
import shlex
import shutil
import subprocess
import sys
import tempfile
import time
from dataclasses import asdict, dataclass
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
STATICCHECK_VERSION = "v0.7.0"
DEFAULT_COMMAND_TIMEOUT_SECONDS = 6 * 60


@dataclass(frozen=True)
class Target:
    name: str
    source: str
    fixture: str
    module: str | None = None


@dataclass(frozen=True)
class Gate:
    name: str
    command: tuple[str, ...]


@dataclass
class GateResult:
    target: str
    source: str
    gate: str
    command: str
    result: str
    exit_code: int
    duration_seconds: float
    report: str


TARGETS = (
    Target(
        name="hxrt",
        source="runtime/hxrt",
        fixture="test/go_tooling/hxrt_gate_test.go",
        module="reflaxe_go_hxrt_tooling_gate",
    ),
    Target(
        name="pulseforge-portable",
        source="examples/pulseforge/generated/portable",
        fixture="test/go_tooling/flagship_gate_test.go",
    ),
    Target(
        name="pulseforge-metal",
        source="examples/pulseforge/generated/metal",
        fixture="test/go_tooling/flagship_gate_test.go",
    ),
    Target(
        name="fluxproxy-portable",
        source="examples/fluxproxy/generated/portable",
        fixture="test/go_tooling/flagship_gate_test.go",
    ),
    Target(
        name="fluxproxy-metal",
        source="examples/fluxproxy/generated/metal",
        fixture="test/go_tooling/flagship_gate_test.go",
    ),
)


def fail(message: str) -> int:
    print(f"[go-tooling] error: {message}", file=sys.stderr)
    return 1


def positive_int_from_env(name: str, default: int) -> int:
    raw = os.environ.get(name, str(default))
    try:
        value = int(raw)
    except ValueError as error:
        raise ValueError(f"{name} must be a positive integer") from error
    if value <= 0:
        raise ValueError(f"{name} must be a positive integer")
    return value


def executable(value: str) -> str | None:
    if os.sep in value:
        path = Path(value).resolve()
        return str(path) if path.is_file() and os.access(path, os.X_OK) else None
    return shutil.which(value)


def output_text(value: str | bytes | None) -> str:
    if value is None:
        return ""
    if isinstance(value, bytes):
        return value.decode("utf-8", errors="replace")
    return value


def run_command(
    command: list[str],
    *,
    cwd: Path,
    timeout_seconds: int,
) -> tuple[str, int, str, float]:
    started = time.monotonic()
    try:
        proc = subprocess.run(
            command,
            cwd=cwd,
            env=os.environ.copy(),
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            text=True,
            timeout=timeout_seconds,
            check=False,
        )
        duration = time.monotonic() - started
        result = "pass" if proc.returncode == 0 else "fail"
        return result, proc.returncode, proc.stdout, duration
    except subprocess.TimeoutExpired as error:
        duration = time.monotonic() - started
        return "timeout", 124, output_text(error.stdout), duration
    except OSError as error:
        duration = time.monotonic() - started
        return "fail", 127, f"unable to execute command: {error}\n", duration


def initialize_report_dir(value: str) -> tuple[Path, Path]:
    report_dir = Path(value).expanduser().resolve()
    protected = {
        Path("/").resolve(),
        Path.home().resolve(),
        ROOT.resolve(),
        ROOT.parent.resolve(),
    }
    if report_dir in protected or len(report_dir.parts) < 3:
        raise ValueError("GO_TOOLING_REPORT_DIR must be a dedicated output directory")

    report_dir.mkdir(parents=True, exist_ok=True)
    reports_dir = report_dir / "reports"
    if reports_dir.is_symlink():
        raise ValueError("GO_TOOLING_REPORT_DIR reports path must not be a symlink")
    reports_dir.mkdir(exist_ok=True)
    for stale_report in reports_dir.glob("*.txt"):
        if stale_report.is_symlink():
            raise ValueError(
                "GO_TOOLING_REPORT_DIR contains an unsafe report symlink: "
                f"{stale_report.name}"
            )
        if stale_report.is_file():
            stale_report.unlink()
        elif stale_report.exists():
            raise ValueError(
                "GO_TOOLING_REPORT_DIR contains a non-file report path: "
                f"{stale_report.name}"
            )

    for name in (
        "go-version.txt",
        "staticcheck-install.txt",
        "staticcheck-version.txt",
        "manifest.json",
        "summary.md",
    ):
        stale_path = report_dir / name
        if stale_path.is_symlink():
            raise ValueError(f"GO_TOOLING_REPORT_DIR contains an unsafe symlink: {name}")
        if stale_path.is_file():
            stale_path.unlink()
        elif stale_path.exists():
            raise ValueError(f"GO_TOOLING_REPORT_DIR contains a non-file path: {name}")
    return report_dir, reports_dir


def write_report(
    path: Path,
    *,
    target: str,
    source: str,
    gate: str,
    command: str,
    result: str,
    exit_code: int,
    duration_seconds: float,
    output: str,
) -> None:
    path.write_text(
        "\n".join(
            [
                f"target={target}",
                f"source={source}",
                f"gate={gate}",
                f"command={command}",
                f"result={result}",
                f"exit_code={exit_code}",
                f"duration_seconds={duration_seconds:.3f}",
                "",
                "--- output ---",
                output.rstrip(),
                "",
            ]
        ),
        encoding="utf-8",
    )


def prepare_target(target: Target, work_root: Path) -> Path:
    source = ROOT / target.source
    fixture = ROOT / target.fixture
    if not source.is_dir():
        raise FileNotFoundError(f"release scope is missing: {target.source}")
    if not fixture.is_file():
        raise FileNotFoundError(f"release fixture is missing: {target.fixture}")

    destination = work_root / target.name
    shutil.copytree(source, destination)
    shutil.copy2(fixture, destination / "go_tooling_gate_test.go")
    if target.module is not None:
        (destination / "go.mod").write_text(
            f"module {target.module}\n\ngo 1.22\n",
            encoding="utf-8",
        )
    return destination


def install_or_resolve_staticcheck(
    *,
    go_bin: str,
    work_root: Path,
    report_dir: Path,
    timeout_seconds: int,
) -> tuple[str | None, dict[str, object]]:
    override = os.environ.get("STATICCHECK_BIN", "")
    if override:
        staticcheck_bin = executable(override)
        if staticcheck_bin is None:
            return None, {
                "result": "fail",
                "reason": "STATICCHECK_BIN is not executable",
                "version": STATICCHECK_VERSION,
            }
    else:
        bin_dir = work_root / "bin"
        bin_dir.mkdir()
        staticcheck_bin = str(bin_dir / "staticcheck")
        command = [
            go_bin,
            "install",
            f"honnef.co/go/tools/cmd/staticcheck@{STATICCHECK_VERSION}",
        ]
        env_before = os.environ.get("GOBIN")
        os.environ["GOBIN"] = str(bin_dir)
        try:
            result, exit_code, output, duration = run_command(
                command,
                cwd=ROOT,
                timeout_seconds=timeout_seconds,
            )
        finally:
            if env_before is None:
                os.environ.pop("GOBIN", None)
            else:
                os.environ["GOBIN"] = env_before
        install_report = report_dir / "staticcheck-install.txt"
        write_report(
            install_report,
            target="tool",
            source="honnef.co/go/tools/cmd/staticcheck",
            gate="install",
            command=f"go install honnef.co/go/tools/cmd/staticcheck@{STATICCHECK_VERSION}",
            result=result,
            exit_code=exit_code,
            duration_seconds=duration,
            output=output,
        )
        if result != "pass" or not Path(staticcheck_bin).is_file():
            return None, {
                "result": result,
                "reason": "pinned staticcheck installation failed",
                "version": STATICCHECK_VERSION,
                "report": install_report.relative_to(report_dir).as_posix(),
            }

    result, exit_code, output, duration = run_command(
        [staticcheck_bin, "-version"],
        cwd=ROOT,
        timeout_seconds=timeout_seconds,
    )
    version_report = report_dir / "staticcheck-version.txt"
    write_report(
        version_report,
        target="tool",
        source="honnef.co/go/tools/cmd/staticcheck",
        gate="version",
        command="staticcheck -version",
        result=result,
        exit_code=exit_code,
        duration_seconds=duration,
        output=output,
    )
    if result != "pass" or f"({STATICCHECK_VERSION})" not in output:
        return None, {
            "result": "fail",
            "reason": f"staticcheck must report exact version {STATICCHECK_VERSION}",
            "version": STATICCHECK_VERSION,
            "reported": output.strip(),
            "report": version_report.relative_to(report_dir).as_posix(),
        }
    return staticcheck_bin, {
        "result": "pass",
        "version": STATICCHECK_VERSION,
        "reported": output.strip(),
        "report": version_report.relative_to(report_dir).as_posix(),
    }


def write_manifest(
    report_dir: Path,
    *,
    result: str,
    go_version: str,
    staticcheck: dict[str, object],
    timeout_seconds: int,
    runs: list[GateResult],
) -> None:
    manifest = {
        "schema_version": 1,
        "result": result,
        "go_version": go_version,
        "staticcheck": staticcheck,
        "command_timeout_seconds": timeout_seconds,
        "retry_policy": "none",
        "targets": [
            {"name": target.name, "source": target.source, "fixture": target.fixture}
            for target in TARGETS
        ],
        "runs": [asdict(run) for run in runs],
    }
    (report_dir / "manifest.json").write_text(
        json.dumps(manifest, indent=2, sort_keys=True) + "\n",
        encoding="utf-8",
    )

    failures = [run for run in runs if run.result != "pass"]
    lines = [
        "# Go Tooling Gate Summary",
        "",
        f"Result: **{result.upper()}**",
        "",
        f"Go: `{go_version}`",
        f"Staticcheck: `{staticcheck.get('reported', STATICCHECK_VERSION)}`",
        f"Command timeout: {timeout_seconds} seconds",
        "Retry policy: none",
        "",
        "| Target | Gate | Result | Seconds | Report |",
        "| --- | --- | --- | ---: | --- |",
    ]
    for run in runs:
        lines.append(
            f"| `{run.target}` | `{run.gate}` | {run.result} | "
            f"{run.duration_seconds:.3f} | `{run.report}` |"
        )
    if failures:
        lines.extend(
            [
                "",
                f"Blocking failures: {len(failures)}. No finding or timeout was retried or downgraded.",
            ]
        )
    (report_dir / "summary.md").write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> int:
    try:
        timeout_seconds = positive_int_from_env(
            "GO_TOOLING_COMMAND_TIMEOUT_SECONDS",
            DEFAULT_COMMAND_TIMEOUT_SECONDS,
        )
    except ValueError as error:
        return fail(str(error))

    go_value = os.environ.get("GO_BIN", "go")
    go_bin = executable(go_value)
    if go_bin is None:
        return fail(f"Go toolchain is required; not executable: {go_value}")

    try:
        report_dir, reports_dir = initialize_report_dir(
            os.environ.get(
                "GO_TOOLING_REPORT_DIR",
                str(ROOT / ".cache" / "security" / "go-tooling"),
            )
        )
    except (OSError, ValueError) as error:
        return fail(str(error))

    go_result, go_exit, go_output, go_duration = run_command(
        [go_bin, "version"],
        cwd=ROOT,
        timeout_seconds=timeout_seconds,
    )
    go_version = go_output.strip()
    write_report(
        report_dir / "go-version.txt",
        target="tool",
        source="go",
        gate="version",
        command="go version",
        result=go_result,
        exit_code=go_exit,
        duration_seconds=go_duration,
        output=go_output,
    )
    if go_result != "pass":
        staticcheck = {"result": "not_run", "version": STATICCHECK_VERSION}
        write_manifest(
            report_dir,
            result="fail",
            go_version=go_version or "unavailable",
            staticcheck=staticcheck,
            timeout_seconds=timeout_seconds,
            runs=[],
        )
        return fail("go version failed; see the persisted report")

    with tempfile.TemporaryDirectory(prefix="haxe-go-tooling-") as temp_name:
        work_root = Path(temp_name)
        staticcheck_bin, staticcheck = install_or_resolve_staticcheck(
            go_bin=go_bin,
            work_root=work_root,
            report_dir=report_dir,
            timeout_seconds=timeout_seconds,
        )
        if staticcheck_bin is None:
            write_manifest(
                report_dir,
                result="fail",
                go_version=go_version,
                staticcheck=staticcheck,
                timeout_seconds=timeout_seconds,
                runs=[],
            )
            return fail(str(staticcheck["reason"]))

        gates = (
            Gate("race", (go_bin, "test", "-race", "-count=1", "-timeout=5m", "./...")),
            Gate(
                "checkptr",
                (
                    go_bin,
                    "test",
                    "-gcflags=all=-d=checkptr=2",
                    "-count=1",
                    "-timeout=5m",
                    "./...",
                ),
            ),
            Gate("vet", (go_bin, "vet", "-stdmethods=false", "./...")),
            Gate("staticcheck", (staticcheck_bin, "-checks=SA*", "./...")),
        )
        runs: list[GateResult] = []
        try:
            prepared = [(target, prepare_target(target, work_root)) for target in TARGETS]
        except (FileNotFoundError, OSError) as error:
            write_manifest(
                report_dir,
                result="fail",
                go_version=go_version,
                staticcheck=staticcheck,
                timeout_seconds=timeout_seconds,
                runs=runs,
            )
            return fail(str(error))

        print(f"[go-tooling] Go: {go_version}")
        print(f"[go-tooling] Staticcheck: {staticcheck['reported']}")
        print(
            f"[go-tooling] scopes={len(prepared)} gates={len(gates)} "
            f"timeout={timeout_seconds}s retries=none"
        )
        for target, target_dir in prepared:
            for gate in gates:
                display_command = list(gate.command)
                display_command[0] = "staticcheck" if gate.name == "staticcheck" else "go"
                command_text = shlex.join(display_command)
                print(f"[go-tooling] {target.name}: {command_text}")
                result, exit_code, output, duration = run_command(
                    list(gate.command),
                    cwd=target_dir,
                    timeout_seconds=timeout_seconds,
                )
                report_name = f"{gate.name}-{target.name}.txt"
                report_path = reports_dir / report_name
                write_report(
                    report_path,
                    target=target.name,
                    source=target.source,
                    gate=gate.name,
                    command=command_text,
                    result=result,
                    exit_code=exit_code,
                    duration_seconds=duration,
                    output=output,
                )
                runs.append(
                    GateResult(
                        target=target.name,
                        source=target.source,
                        gate=gate.name,
                        command=command_text,
                        result=result,
                        exit_code=exit_code,
                        duration_seconds=duration,
                        report=f"reports/{report_name}",
                    )
                )
                print(
                    f"[go-tooling] {target.name}/{gate.name}: {result} "
                    f"({duration:.2f}s)"
                )

    overall = "pass" if all(run.result == "pass" for run in runs) else "fail"
    write_manifest(
        report_dir,
        result=overall,
        go_version=go_version,
        staticcheck=staticcheck,
        timeout_seconds=timeout_seconds,
        runs=runs,
    )
    if overall != "pass":
        return fail(f"one or more release gates failed; see {report_dir / 'summary.md'}")
    print(f"[go-tooling] all gates passed; reports: {report_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
