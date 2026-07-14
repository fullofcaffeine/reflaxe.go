# GPT-5.6 Pro evidence record: `cd79624f`

This is the human-readable index for the canonical machine-readable [evidence record](evidence-cd79624f.json). The review source is Haxe.Go commit `cd79624f855521dbf320ac2b7524d889ca388c0e` on `origin/master`.

## Artifact identity

- File: `dist/review/haxe-go-gpt56-evidence-cd79624f.zip`
- Size: 13,502,320 bytes
- SHA-256: `ab4b0a1097229ad2202ca7da6c092b2b85cba537522951fb625f6b1312c4b511`
- Builder: commit `531e5bc1368c2b72a41cd5b46fc2449e7ed90393`; builder-file SHA-256 `447be705cb2a4ad73116f83256dcac125d2b4e745d195584b5819a8dd4272a78`

Two clean builds produced byte-identical ZIPs. `unzip -t` passed, all 379 entries in the internal `SHA256SUMS` verified, the full payload passed Gitleaks, and 375 UTF-8 payload files passed the machine-local workspace-path check.

## Evidence boundaries

The Git archive is authoritative for all 5,368 included tracked files. The line-numbered Repomix view is only a navigation aid. Repomix's scanner omitted eight tracked TLS/proxy test fixtures; those files remain in the authoritative archive and are also included as explicit raw companion payloads after the complete bundle passed Gitleaks.

The exact source and sibling Git payloads are commit-deterministic. GitHub Actions logs, release metadata, and host controls are live capture-at-build evidence. They describe the state recorded on 2026-07-14 and must not be assumed to describe later repository state.

## Material facts for the reviewer

- All five selected CI/security runs succeeded at the exact source commit.
- The current release is `v0.53.1`, whose tag resolves to `4adb9ae8209c24850495110e79ba6c5a8e1fa2bd`.
- That GitHub release has zero assets and is not immutable.
- The repository reports no rulesets, and `master` is not protected.
- The source inventory contains 68 checked-in `.cross.hx` files, 142 generated metal files, 195 generated portable files, 712 Haxe source files, and 3,624 intended snapshot files.
- Sibling precedent is pinned to Haxe.Rust `5b8c9416…`, Haxe.Ruby `08faba04…`, and Haxe.Elixir `68625fa9…`.
- The roadmap snapshot contains 663 issues rooted at `haxe_go-vfp`, at Dolt commit `d4n7bp04tc9k1srifgn0ipbb8thkgld9`, with zero dependency cycles.

CI success is evidence of the current contracts, not proof that the product is production-ready. In particular, release artifact absence, mutable releases, and missing host enforcement are explicit inputs for the deep review rather than facts to smooth over.

## Storage and reproduction

The ZIP is intentionally ignored under `dist/review/`; upload or transfer only the file matching the size and SHA-256 above. The exact argument vector and all intentional exclusions are preserved in the JSON record. General construction and verification rules are documented in [README.md](README.md).
