# Release Scaffold

Planned standalone release model for `reflaxe.family.std`:

1. SemVer tags (`vMAJOR.MINOR.PATCH`).
2. Contract version tuple tracked in release notes:
   - `portable-semantics` version
   - allowlist schema version
   - conformance mapping schema version
3. Every release includes:
   - schema validation report
   - manifest integrity report
   - consumer migration notes (if contract behavior changes)
