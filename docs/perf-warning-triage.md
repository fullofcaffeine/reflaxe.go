# Perf Warning Triage

This page records human decisions for performance warnings that are visible in
green CI.

Performance warnings are not automatically bugs. They mean a benchmark crossed a
soft budget. A soft budget is a visibility threshold: CI stays green, but a
human should check whether the warning is repeated and stable enough to act on.

Canonical policy: `docs/performance-budget-policy.md`.

## June 13, 2026 Triage

Bead: `haxe.go-nhh2.3`

Decision:

- Do not update baselines.
- Do not promote app perf warnings to hard gates.
- Keep Go profile metal microbench hard gating as-is.
- Implement follow-up `haxe.go-nhh2.4` for multi-run startup variance handling.

Why:

- The latest Go profile microbench run was clean: `warnings=0`,
  `metal_hard_failures=0`, and `delta_hard_failures=0`.
- The latest flagship app run had two warning-only FluxProxy startup signals:
  `fluxproxy::core::portable.startup_ratio_vs_pure` and
  `delta.fluxproxy::core.startup_delta`.
- The flagship app delta hard-gate dry run had `Candidate count: 0`.
- Recent successful CI runs did not show one stable repeated app regression.
  Warnings moved between FluxProxy core, FluxProxy Go-native, PulseForge core,
  portable metrics, metal metrics, and delta metrics.

Evidence sampled:

| CI Harness run | Result | Go profile signal | Flagship app signal |
| --- | --- | --- | --- |
| `27478348757` | success | `warnings=0` | FluxProxy core portable startup and startup delta warnings |
| `27477940525` | success | `warnings=0` | PulseForge core latency/throughput delta warnings |
| `27476266298` | success | `warnings=0` | no app warnings |
| `27474886300` | success | one portable array startup warning | FluxProxy core metal throughput/latency warnings |
| `27462939863` | success | `warnings=0` | no app warnings |
| `27442845666` | success | one portable generic startup warning | FluxProxy Go-native metal throughput/latency warnings |

Interpretation:

- `Go profile microbench` does not currently need a baseline update or new
  optimization bead from this triage. The latest two successful runs were clean.
- `FluxProxy` remains worth watching because startup-related warnings appear in
  recent history, but the exact metric and variant are not stable enough for a
  hard gate.
- The next useful improvement is better variance handling: collect more than one
  sample, report median and p95, and promote only warnings that repeat under
  that stronger evidence model.

Follow-up result:

- `haxe.go-nhh2.4` makes the flagship app harness collect multiple startup
  samples per binary.
- Startup ratios now use the median sample.
- Raw artifacts still report startup average and p95 so reviewers can see
  whether a warning is a stable median drift or just noisy startup variance.

## How To Triage The Next Warning

1. Open the latest CI Harness run.
2. Download or inspect these artifacts:
   - `go-profile-perf-results`
   - `go-app-perf-results`
3. Read:
   - `warning_history.md`
   - `warnings.txt`
   - `delta_hard_gate_dry_run.md`
4. Compare the latest warning with several successful CI runs.
5. Choose one outcome:
   - baseline update, only when the new number is intentional and documented;
   - optimization bead, when a stable regression points to a real code path;
   - hard-gate promotion, only when the criteria in
     `docs/performance-budget-policy.md` are met;
   - variance-handling work, when the warning moves around or depends on startup
     noise.

Do not treat a single soft warning as release-blocking unless the policy has
already promoted that harness and metric to a hard gate.
