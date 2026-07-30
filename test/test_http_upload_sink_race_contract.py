#!/usr/bin/env python3

from __future__ import annotations

import json
import os
from pathlib import Path
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parent.parent
CASE = ROOT / "test" / "snapshot" / "sys" / "http_upload_sink_lifecycle_contract"
PACKAGE_JSON = ROOT / "package.json"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"
SCRIPT_NAME = "test:http-upload-sink-race"


def run_checked(
    command: list[str],
    *,
    cwd: Path,
    env: dict[str, str] | None = None,
    timeout: int = 120,
) -> subprocess.CompletedProcess[str]:
    merged_env = os.environ.copy()
    if env:
        merged_env.update(env)
    process = subprocess.run(
        command,
        cwd=cwd,
        env=merged_env,
        capture_output=True,
        text=True,
        timeout=timeout,
    )
    if process.returncode == 0:
        return process
    output = "\n".join(part for part in (process.stdout, process.stderr) if part)
    raise AssertionError(f"{' '.join(command)} failed with {process.returncode}\n{output}")


def verify_gate_wiring() -> None:
    scripts = json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))["scripts"]
    expected = "python3 test/test_http_upload_sink_race_contract.py"
    if scripts.get(SCRIPT_NAME) != expected:
        raise AssertionError(f"{SCRIPT_NAME} must run the generated-Haxe upload race contract")
    for aggregate in ("test", "test:changed"):
        if f"npm run {SCRIPT_NAME}" not in scripts.get(aggregate, ""):
            raise AssertionError(f"{aggregate} must include npm run {SCRIPT_NAME}")
    if "test/test_http_upload_sink_race_contract.py" not in RELEASE_RUNNER.read_text(
        encoding="utf-8"
    ):
        raise AssertionError("release contracts must include the generated-Haxe upload race contract")


def main() -> int:
    verify_gate_wiring()
    expected = (CASE / "expected.stdout").read_text(encoding="utf-8").replace("\r\n", "\n")

    with tempfile.TemporaryDirectory(prefix="haxe-go-http-upload-race-") as raw_temp:
        output = Path(raw_temp) / "out"
        run_checked(
            [
                "haxe",
                "compile.hxml",
                "-D",
                "go_no_build",
                "-D",
                f"go_output={output}",
            ],
            cwd=CASE,
            env={"HAXE_NO_SERVER": "1"},
        )
        go_files = sorted(str(path) for path in output.rglob("*.go"))
        if not go_files:
            raise AssertionError("the upload lifecycle fixture generated no Go files")
        run_checked(["gofmt", "-w", *go_files], cwd=output)
        run_checked(["go", "test", "./..."], cwd=output)
        raced = run_checked(
            ["go", "run", "-race", "."],
            cwd=output,
            env={"GORACE": "halt_on_error=1"},
            timeout=180,
        )

        actual = raced.stdout.replace("\r\n", "\n")
        if actual != expected:
            raise AssertionError(
                "generated-Haxe upload race stdout mismatch\n"
                f"expected:\n{expected}\nactual:\n{actual}"
            )
        if "DATA RACE" in raced.stderr:
            raise AssertionError(f"Go race detector reported a data race\n{raced.stderr}")

    print("HTTP upload sink generated-Haxe race contract passed")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
