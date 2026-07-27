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
        ["python3", "test/test_beads_workflow_contract.py"],
        ["python3", "test/test_array_remove_insert_lowering_contract.py"],
        ["python3", "test/test_canonical_std_closeout_contract.py"],
        ["python3", "test/test_ci_haxe_setup_action.py"],
        ["python3", "test/test_compatibility_support_manifest.py"],
        ["python3", "test/test_compiler_debt_ratchet.py"],
        ["python3", "test/test_compiler_stdlib_intrinsic_registry.py"],
        ["python3", "test/test_dependency_audit_gate_semantics.py"],
        ["python3", "test/test_examples_qa_contract.py"],
        ["python3", "test/test_generated_output_telemetry.py"],
        ["python3", "test/test_generated_output_confinement.py"],
        ["python3", "test/test_go_tooling_gate_contract.py"],
        ["python3", "test/test_go_tooling_gate_semantics.py"],
        ["python3", "test/test_goextern_ci_gate_contract.py"],
        ["python3", "test/test_github_governance_policy.py"],
        ["python3", "test/test_haxelib_package_runner.py"],
        ["python3", "test/test_haxelib_release_artifact.py"],
        ["python3", "test/test_haxelib_release_install.py"],
        ["python3", "test/test_haxe_dynamic_boundaries.py"],
        ["python3", "test/test_haxe_macro_lifecycle.py"],
        ["python3", "test/test_haxe_typed_identifiers.py"],
        ["python3", "test/test_haxe_warning_ratchet.py"],
        ["python3", "test/test_inline_throw_accessor_result_type_contract.py"],
        ["python3", "test/test_language_hard_fail_inventory_contract.py"],
        ["python3", "test/test_license_policy_contract.py"],
        ["python3", "test/test_lambda_iterable_lowering_ownership_contract.py"],
        ["python3", "test/test_markdown_internal_links_contract.py"],
        ["python3", "test/test_metal_preset_retention_contract.py"],
        ["python3", "test/test_metal_graduation_contract.py"],
        ["python3", "test/test_multi_package_output_decision_contract.py"],
        ["python3", "test/test_perf_budget_policy_contract.py"],
        ["python3", "test/test_perf_delta_dry_run.py"],
        ["python3", "test/test_perf_warning_summary.py"],
        ["python3", "test/test_pinned_npm_bootstrap.py"],
        ["python3", "test/test_portable_parity_closure_contract.py"],
        ["python3", "test/test_portable_governance_contract.py"],
        ["python3", "test/test_post_generation_build_runner.py"],
        ["python3", "test/test_public_contract.py"],
        ["python3", "test/test_raw_injection_hygiene_contract.py"],
        ["python3", "test/test_release_readiness_checklist_contract.py"],
        ["python3", "test/test_release_readiness_gate.py"],
        ["python3", "test/test_release_identity_contract.py"],
        ["python3", "test/test_release_package_metadata_contract.py"],
        ["node", "test/test_release_reconciliation.mjs"],
        ["node", "test/test_release_version_policy.mjs"],
        ["python3", "test/test_review_evidence_bundle_contract.py"],
        ["python3", "test/test_supply_chain_contract.py"],
        ["python3", "test/test_same_sha_release_wrapper.py"],
        ["python3", "test/test_semantic_release_dependency_boundary.py"],
        ["python3", "test/test_semver_lifecycle_policy.py"],
        ["python3", "test/test_sibling_target_classpath_guard.py"],
        ["python3", "test/test_stdlib_migration_ledger_contract.py"],
        ["python3", "test/test_sys_get_char_terminal_contract.py"],
        ["python3", "test/test_toolchain_policy_contract.py"],
        ["python3", "test/test_typed_exception_hygiene_contract.py"],
        ["python3", "test/test_vendor_reflaxe_provenance.py"],
    ]
    for command in commands:
        exit_code = run(command)
        if exit_code != 0:
            return exit_code
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
