# Approach C Baseline Freeze (Historical M0 Snapshot)

Date captured: 2026-02-25
Status: Archived historical snapshot (superseded by current implementation)

This file preserves the pre-consolidation baseline used during Approach C planning.
It does not describe the current architecture.

## Current Architecture Status (as of 2026-03-05)

The baseline gaps recorded in this spike are now partially or fully superseded:

1. Unified build context exists:
   - `src/reflaxe/go/compiler/GoBuildContext.hx`
   - `src/reflaxe/go/compiler/GoBuildContextResolver.hx`
2. Compile pipeline consumes the resolved context and emits deterministic reports:
   - `src/reflaxe/go/GoReflaxeCompiler.hx`
   - optional artifacts: `profile_contract.*`, `hxrt_plan.*`, `optimizer_plan.*`
3. Module lanes are analyzed and propagated in compile state:
   - `src/reflaxe/go/analyze/MetalLaneAnalyzer.hx`
4. Lane-specific semantic-diff coverage exists:
   - `python3 test/run-semantic-diff.py --suite lanes`

At the time of this snapshot, remaining work was concentrated in parity closure and unsupported-expression cleanup, not in replacing the profile/context model.
That status has since been superseded by current parity evidence:

- `docs/known-gaps.md` summarizes the live readiness posture.
- `test/.test-cache/portable_parity_closure_summary.json` is the generated machine-readable parity authority.
- `test/.test-cache/portable_parity_closure_summary.md` is the generated human-readable parity summary.

## Historical Baseline Notes (kept for provenance)

At capture time (2026-02-25), this spike recorded:

- profile-shaped contract selector (`portable|metal`) as the visible interface,
- strict defaults and fallback policy spread across multiple entrypoints,
- runtime selective-copy behavior being additive and orthogonal to profile,
- strict-enforcer allowlist concerns (including stale `std/go/metal` assumptions),
- missing centralized typed context/report spine at that historical point.

Those observations are retained for change-history traceability only.

## Use Instead

- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/portable-stdlib-parity-program.md`
- `docs/known-gaps.md`
