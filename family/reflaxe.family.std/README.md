# reflaxe.family.std (Bootstrap Snapshot)

This directory is a bootstrap snapshot for the future standalone `reflaxe.family.std` repository.

It packages family-shared portable contract artifacts extracted from `haxe.go`:

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
