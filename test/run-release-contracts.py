#!/usr/bin/env python3

from __future__ import annotations

import subprocess
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent


def run(cmd: list[str]) -> int:
    print("$", " ".join(cmd))
    proc = subprocess.run(cmd, cwd=ROOT)
    return proc.returncode


def main() -> int:
    commands = [
        ["python3", "test/test_metal_graduation_contract.py"],
        ["python3", "test/test_multi_package_output_decision_contract.py"],
        ["python3", "test/test_portable_governance_contract.py"],
        ["python3", "test/test_release_readiness_checklist_contract.py"],
    ]
    for command in commands:
        exit_code = run(command)
        if exit_code != 0:
            return exit_code
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
