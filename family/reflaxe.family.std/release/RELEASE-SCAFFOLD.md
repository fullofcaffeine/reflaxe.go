# Release Scaffold

Status: blocked until the immutable family-core boundary is reconciled across
at least two consumers.

This file describes a planned release, not a release that currently exists.
The Go and Rust bootstraps both use `0.1.0-bootstrap.1` while containing
different payloads, so neither copy may be promoted as the shared package.

Planned standalone release model for the `reflaxe.family.std` core:

1. SemVer tags (`vMAJOR.MINOR.PATCH`).
2. Contract version tuple tracked in release notes:
   - `portable-semantics` version
   - allowlist schema version
   - conformance mapping schema version
3. Every release includes:
   - immutable core content digest
   - schema validation report
   - manifest integrity report
   - consumer migration notes (if contract behavior changes)
4. Every consumer verifies the same digest for the pinned core version.
5. Target-specific allowlists, ownership mappings, fixture bindings, deviations,
   and implementation evidence use separately identified target overlays.
6. A version is never republished with different core contents.
