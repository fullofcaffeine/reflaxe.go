# Portable beta candidate evidence: `c22af0ea`

This record says something deliberately narrow: one Haxe.Go source commit passed the complete portable-beta candidate matrix, and the proof was frozen into a reproducible ZIP. It does not say that every Haxe or Go feature is supported, and it does not publish a release.

The machine-readable companion is [portable-beta-candidate-c22af0ea.json](portable-beta-candidate-c22af0ea.json).

## Why the exact commit matters

Think of the evidence ZIP as a sealed box. The source, test logs, policies, issue state, and sibling references inside it all have hashes. Every selected hosted workflow ran against the same exact source commit, `c22af0ea82e5e481e23277e513ed5b7c6b5c770b`. We do not combine a compiler test from one commit with a security result from another and call the mixture a release candidate.

The authoritative source is the embedded Git archive. The Repomix XML is a line-numbered reading aid, not a replacement source tree.

## What this proves

- The admitted portable compiler surface passed snapshots, semantic comparisons, strict standard-library inventory, diagnostics and generated-output contracts.
- The representative target smoke uses Haxe through the custom Haxe.Go backend, then strictly builds/tests the generated Go and executes the resulting program. It is not only a generated-text inspection.
- The two supported Go versions, 1.25.12 and 1.26.5, passed their governed tooling and dependency-audit lanes.
- The package contracts produced deterministic Haxelib artifacts and exercised isolated installs. Portable and metal are separate scorecards: a green metal/native result does not advance the portable claim.
- Static analysis, vulnerability checks, secret scanning, raw-injection hygiene, race-sensitive tests, license policy, public API policy, and release-readiness contracts passed.
- Two builds of the review ZIP, made with the same inputs and output basename, were byte-for-byte identical.

The public claim is a pre-1.0 portable beta on Linux/amd64 for exactly 38 admitted operations covering 173 named symbols. Its compatibility-surface digest is `99625a5bcb401561a8393ddcef5675ba552c1a84ab288ea4ed0e1cc950bac0d0`.

## What this does not prove

- This is not a stable 1.x claim.
- This is not a publication record. The candidate has not been released by this task.
- HTTP remains excluded behind `haxe_go-vfp.10.8`.
- Go-native APIs remain a separate, non-admitted product surface behind `haxe_go-vfp.9.1`.
- Untrusted or adversarial source compilation remains excluded. The admitted beta assumes reviewed Haxe source, metadata, dependencies, plugins, macros, and selected child tools; Haxe.Go is not a security sandbox.
- Other operating systems, architectures, unlisted members, and unlisted error or concurrency paths do not become supported because Linux portable tests are green.
- The `metal` compatibility preset is not a second semantic product and its evidence cannot widen the portable claim.
- Existing admitted file, process, thread, TCP, UDP, and TLS operations are only the exact member subsets named in the compatibility manifest; this record does not turn them into blanket whole-module promises.

The rule is simple: if an operation is not explicitly marked `release_admitted`, it is excluded.

## Evidence matrix

| Evidence owner | Result | What it protects |
| --- | --- | --- |
| Compiler snapshots and focused contracts | Pass | Generated Go shape, diagnostics, known regressions, and post-generation fatal-build handling |
| Portable semantic diff | Pass | Haxe-visible behavior for the portable cases that have a meaningful interpreter/reference oracle |
| Official Haxe target inventory | Pass on Go 1.25.12 and 1.26.5 | Complete applicable upstream target modules, compiled through the real custom backend |
| Examples and target smoke | Pass | Haxe compile → generated Go build/test → target execution and expected output |
| Runtime, race, and checkptr owners | Pass | Generated runtime/stdlib behavior and resource/concurrency contracts |
| Package artifact and isolated install | Pass | Deterministic packaging and downstream use from a clean Haxelib repository |
| Security, license, API, compatibility, and release policy | Pass | Supply-chain posture and truthful scope; these checks do not create compiler semantics by themselves |

Warm/compiler-server paths are accelerators only. Cold, clean hosted lanes remain part of the proof.

## Hosted exact-SHA runs

- [Examples Artifacts 30743182325](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30743182325)
- [Security - Gitleaks 30743182328](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30743182328)
- [Security - Static Analysis 30743182367](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30743182367)
- [CI Harness 30743182369](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30743182369)
- [CI - Quality 30743182373](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30743182373)

All five report success for `c22af0ea82e5e481e23277e513ed5b7c6b5c770b`. Their 20 hosted artifacts record SHA-256 digests and remain available through October 31, 2026. The evidence ZIP preserves the artifact index, job graphs, and sanitized logs; it does not silently pretend that hosted artifacts live forever.

## Artifact identity

- File: `haxe-go-portable-beta-c22af0ea.zip`
- Size: 27,529,201 bytes
- SHA-256: `7d76936576bb9e7e654f73ac010a474305e23b2c08ec2fe2f406441af633eb66`
- Internal payloads: 383
- Internal `SHA256SUMS` entries verified: 384
- Embedded source commit returned by `git get-tar-commit-id`: `c22af0ea82e5e481e23277e513ed5b7c6b5c770b`

The complete ZIP passed `unzip -t`, Gitleaks, and the machine-local path check. Repomix omitted nine security-sensitive test fixtures from its XML view; the authoritative Git archive still contains them, and the bundle copies them as explicit raw companions after scanning.

## How to verify the ZIP

Use the handed-off ZIP whose filename matches the identity above:

```bash
shasum -a 256 haxe-go-portable-beta-c22af0ea.zip
unzip -t haxe-go-portable-beta-c22af0ea.zip

verify_dir="$(mktemp -d)"
unzip -q haxe-go-portable-beta-c22af0ea.zip -d "$verify_dir"
(
  cd "$verify_dir"
  shasum -a 256 -c SHA256SUMS
  git get-tar-commit-id < primary/haxe.go-source-c22af0ea.tar
)
```

The outer SHA must equal the value above, every internal checksum must say `OK`, and the Git command must print the exact candidate commit.

## Rebuild inputs

The sibling references are pinned for learning and architectural comparison, not used as substitute test oracles:

- Haxe.Rust `5b8c9416f963e541229e633a2bb655a93e3e9c16`
- Haxe.Ruby `dbb70af0e48e252e413645b7bf16197a4776f0f8`
- Haxe.Elixir `68625fa91ffff48c5ffb269bff01c6f3e716128c`

Rebuilding also captures live GitHub and Beads metadata. Those live facts may change later, so a later rebuild is new dated evidence rather than a reason to rewrite this frozen record.

## Next gate

`haxe_go-vfp.12.5` is the independent final release-admission review. That review is worth doing because it checks a provenance-sensitive release claim, not because every completed task needs another agent. If it accepts the exact evidence and narrow claim, `haxe_go-vfp.12.6` owns same-SHA publication. Stable 1.x remains separate under `haxe_go-vfp.12.7`.
