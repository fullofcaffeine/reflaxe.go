# Beads v53 Migration Integrity Record

The Beads database schema was migrated from version 32 (`bd` 1.0.4) to version
53 (`bd` 1.1.0) on 2026-07-19 so `haxe_go-vfp.6.3` could be claimed and updated.
The migration was performed only after creating an operator-owned backup
outside the repository.

## Before migration

- remote `refs/dolt/data`: `cf377a9628711e2cd60d03636a25d5b2feb4cc2b`
- issues: 723
- memories: 12
- dependency edges: 1,004
- issue ID-set SHA-256:
  `44eee704a6397f1e20dd651a52cd2a6448d615d2cfbc3e564be9f31a76a14b58`
- memory ID-set SHA-256:
  `ecfa0b90e55309b527514d086378498176b1fbc0880badb2d086a201b8a40fc5`
- dependency-tuple SHA-256:
  `c1335edf010e611ed5ee31cee313dbc08073be35314e479524b0602aea2eac8c`

The external recovery bundle contains both the raw version-32 database archive
and a logical issues-and-memories export. It is intentionally not committed and
no machine-local backup path is recorded here.

## After migration

- remote `refs/dolt/data`: `2c7b7b6d5c94f5546f169a8957ec119e93ad52ca`
- issues: 723
- memories: 12
- dependency edges: 1,004
- issue, memory, and dependency ID/tuple hashes: unchanged
- local version-53 logical export SHA-256:
  `693db4450053a5b76ac794724b939c402e526fcf384a27648fc633236b4d460b`
- fresh-clone remote logical export SHA-256:
  `693db4450053a5b76ac794724b939c402e526fcf384a27648fc633236b4d460b`

The local and fresh-clone version-53 exports are byte-for-byte identical. A
field-level comparison with an independently built `bd` 1.0.4 export found no
semantic record changes after excluding derived serialization fields,
`updated_at`, and regenerated legacy comment IDs. Comment author, text, order,
and timestamp values were preserved.

## Separate client behavior change

`bd ready` changed from 11 ready issues under 1.0.4 to one under 1.1.0 even
though the dependency graph is identical. The new client treats structural
parent-child links as affecting readiness. This is tracked as `haxe_go-7hiq`;
valid dependency data was not deleted or rewritten to hide the symptom.
