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

## Corrected readiness diagnosis and repair

The migration-time observation was real: `bd ready` changed from 11 ready
issues under 1.0.4 to one under 1.1.0 while the dependency tuple hash stayed
the same. The initial explanation was wrong, however. A minimal isolated
fixture proved that an ordinary `parent-child` link leaves both parent and
child ready. Adding an external active blocker to the parent suppresses the
child until that blocker closes.

Importing the same later logical export into isolated 1.0.4 and 1.1.0
databases reproduced the difference as nine ready issues versus two. The cause
was a hierarchy feedback loop in the haxe.go graph: root epic `haxe_go-vfp`
was the structural ancestor of its child epics and also had explicit `blocks`
dependencies on those same children. The children blocked the root, and 1.1.0
then propagated the root's blocked state through its descendants. The older
client's less complete propagation leaked some grandchildren into the ready
list. This was not migration corruption and was not evidence that normal
parent-child organization directly blocks work.

Upstream documents parent-child links as organizational while supporting
blocked-parent propagation to descendants. It also rejects new blocking edges
between parents and descendants, but its current cycle/doctor coverage can
miss the mirror direction used by this historical graph. See upstream issues
[#104](https://github.com/steveyegge/beads/issues/104),
[#19](https://github.com/steveyegge/beads/issues/19), and
[#4814](https://github.com/steveyegge/beads/issues/4814), plus pull requests
[#4034](https://github.com/steveyegge/beads/pull/4034) and
[#4035](https://github.com/steveyegge/beads/pull/4035).

`haxe_go-7hiq` repaired only the eight root-to-child `blocks` edges whose two
endpoints were still active. All structural `parent-child` edges and five
closed historical hierarchy-shadow edges were retained. The repair was first
rehearsed against an isolated database and then checked against preserved
before/after exports:

- before: 727 issues, 13 memories, and 1,007 dependency tuples;
- after: 727 issues, 13 memories, and 999 dependency tuples;
- issue ID-set SHA-256:
  `9e8294ece8dbbcf8cb89d45c151e829be188fe64d38bfe7ad1110b4f1589c8b9`;
- memory key-set SHA-256:
  `8ed8a7c8bbe8c3b659c8bdb1a9df051154224b62574c274ebdbac083c426dc0d`;
- after dependency-tuple SHA-256:
  `b5365b7dbbf7609b1cbe3d80776f76702c14c2d5c246d82852ad5d2cfbbbadab`;
- exactly eight intended `blocks` tuples removed, zero tuples added, every
  parent-child tuple unchanged, every memory unchanged, and every other issue
  field unchanged apart from derived dependency counts;
- `bd ready` increased from two to 16 on the live pre/post repair state.

The repository now runs
`scripts/beads/check-hierarchy-deadlocks.py` from its tracker health gate. The
guard rejects active blocking edges between any ancestor and descendant while
allowing sibling ordering and inactive historical provenance.
