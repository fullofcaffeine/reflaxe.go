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
        ["python3", "test/test_ci_haxe_setup_action.py"],
        ["python3", "test/test_examples_qa_contract.py"],
        ["python3", "test/test_generated_output_telemetry.py"],
        ["python3", "test/test_goextern_ci_gate_contract.py"],
        ["python3", "test/test_language_hard_fail_inventory_contract.py"],
        ["python3", "test/test_lambda_iterable_lowering_ownership_contract.py"],
        ["python3", "test/test_markdown_internal_links_contract.py"],
        ["python3", "test/test_metal_graduation_contract.py"],
        ["python3", "test/test_multi_package_output_decision_contract.py"],
        ["python3", "test/test_perf_budget_policy_contract.py"],
        ["python3", "test/test_perf_delta_dry_run.py"],
        ["python3", "test/test_perf_warning_summary.py"],
        ["python3", "test/test_portable_parity_closure_contract.py"],
        ["python3", "test/test_portable_governance_contract.py"],
        ["python3", "test/test_release_readiness_checklist_contract.py"],
        ["python3", "test/test_release_package_metadata_contract.py"],
        ["python3", "test/test_typed_exception_hygiene_contract.py"],
    ]
    for command in commands:
        exit_code = run(command)
        if exit_code != 0:
            return exit_code
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
