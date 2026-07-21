# reflaxe.family.std (Bootstrap Snapshot)

This directory is a bootstrap snapshot for the future standalone `reflaxe.family.std` repository.

It is currently a **target-local mirror**, not an installed release or a
cross-repository source of truth. The verifier proves that this directory agrees
with `haxe.go`'s canonical local artifacts. It does not prove that another
compiler has the same files merely because it uses the same bootstrap version.

The 2026-07-20 sibling comparison found that `haxe.rust` also uses
`0.1.0-bootstrap.1`, but several of its mirrored payload files differ. Genes and
`haxe.elixir.codex` use strong repo-local compatibility governance without a
family directory. See
`docs/spikes/reflaxe-family-stdlib-extraction-spike.md` for the evidence and the
revised extraction decision.

It packages candidate family-core and target-adapter artifacts extracted from
`haxe.go`:

- portable semantics contract (`contracts/portable-semantics/v1.md`)
- portable allowlist (`allowlists/portable_allowlist.v1.json`)
- tier1 conformance mapping (`conformance/tier1/portable_conformance_tier1.v1.json`)
- portable module ownership mapping (`docs/module-mapping-contract.v1.md`)
- provenance schema and boundary policy (`provenance/*`)

Validation:

```bash
python3 family/reflaxe.family.std/tools/verify_family_std.py
```

This snapshot is CI-gated in `haxe.go` until extraction to an external repo is completed.

A future external release must separate one immutable family core from
target-qualified adapter overlays. Every consumer of a given core version must
verify the same core content digest; target-specific allowlists, mappings,
fixture bindings, deviations, and implementation evidence must not silently
change that core payload.
