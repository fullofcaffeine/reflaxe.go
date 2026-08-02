# Beads workflow for haxe.go

## What it is

Beads is the repository's issue tracker. The local Dolt database is the operational source of truth for `bd list`, `bd show`, `bd ready`, and all issue writes. Its off-machine history is stored separately from source branches under `refs/dolt/data` on the Git origin.

The Git checkout and the Beads database therefore have separate pull and push operations. A successful `git push` does not imply that a pending Dolt issue update was pushed, and a successful `bd dolt push` does not publish source commits.

### Tracker transport and authentication

Beads uses SSH for its Dolt remote. In plain terms, tracker updates travel over
the same authenticated GitHub connection commonly used for private Git pushes.
The earlier HTTPS transport could disconnect with HTTP 400 while uploading a
valid tracker commit; retrying it did not make that path reliable. SSH accepted
the same commit and preserved the exact remote history.

You need a GitHub SSH key that can access this repository. This setting applies
only to Beads history under `refs/dolt/data`; it does not change the source-code
Git remote. A fresh clone reads the SSH address from `.beads/config.yaml` when
`bd bootstrap --yes` initializes its local tracker.

Existing clones that still show an HTTPS `origin` need this one-time local
update after confirming their SSH access:

```bash
git ls-remote git@github.com:fullofcaffeine/reflaxe.go.git refs/dolt/data
bd dolt remote remove origin
bd dolt remote add origin git+ssh://git@github.com/fullofcaffeine/reflaxe.go.git
bd dolt pull
```

Removing this local remote name does not delete tracker records or remote
history. The health check rejects a local Dolt `origin` that differs from the
checked `sync.remote`, making stale clone configuration visible before a push.

## Why it exists

The tracker preserves the dependency graph, review decisions, acceptance evidence, and persistent memories needed for the long-running Haxe.Go Next program. Keeping issue history in Dolt allows cell-level history and clean bootstrap without turning issue updates into source commits.

This repository also predates the Dolt migration. Its tracked `.beads/issues.jsonl` is the canonical legacy provenance archive, not the active database and not a routine export target. The archive is immutable at:

- SHA-256: `0e34e32cb1ac25fdc8592aea85aa5630ca31ab59076b3e33faa6611a4e51911c`
- 579 unique records: 578 live legacy issues and tombstone `haxe.go-dsn`

The imported Dolt issues preserve live issue IDs, semantic fields, labels, dependency relationships, and comment content. The importer rounded some timestamps, lost the original timezone offset on legacy `closed_at` and dependency timestamps, and regenerated legacy comment IDs. Consult the canonical legacy provenance archive when exact historical metadata matters.

To prevent accidental replacement of that archive, `.beads/config.yaml` intentionally contains:

```yaml
export.auto: false
export.git-add: false
sync.require_confirmation_on_mass_delete: true
```

Do not re-enable auto-export until a reviewed migration gives current operational exports a different path and preserves the archive unchanged.

## How it works

At the start of a session:

```bash
bd prime
bd dolt pull
bd ready
bd show <id>
bd update <id> --claim
```

Use Beads for all work tracking. Add follow-up issues as they are discovered, attach evidence in comments or issue fields, and use `bd remember` for cross-session architectural knowledge. Do not create markdown task lists or `MEMORY.md` substitutes.

Before ending a session:

```bash
# Update or close the active bead first.
bd dolt pull
git pull --rebase
bd dolt push
git push
git status --short --branch
```

Both histories must be current. `git status` should report the source branch up to date with origin; `bd dolt push` must succeed for issue changes. Preserve unrelated user files and never force-push either history as a routine conflict workaround.

Useful health checks:

```bash
scripts/beads/check-health.sh

# After bd dolt push and git push, verify both remote histories exactly:
scripts/beads/check-health.sh --session-close

# Compare local and remote Dolt exports without enforcing Git cleanliness:
scripts/beads/check-health.sh --verify-remote
```

The script uses supported read-only commands: `bd config validate`, `bd dep cycles`, `bd lint`, `bd orphans`, `bd vc status --json`, and `bd stats`. Exact remote verification creates a disposable Dolt clone, exports both databases with memories, and compares them byte-for-byte.

It also runs `scripts/beads/check-hierarchy-deadlocks.py`, a repository guard for
one dependency shape that upstream cycle checks do not currently catch. The
guard reads a logical export and never writes to the tracker.

`bd doctor` is not supported in embedded mode in bd 1.0.4. `bd config validate` plus the graph/lint/orphan checks are the explicit substitute. `bd preflight --check` currently describes the upstream Beads repository's generic Go/Nix checklist rather than haxe.go's gates, so it is not used as tracker-health evidence here.

## Hierarchy and readiness

A `parent-child` dependency says where an issue belongs; by itself, it does not
mean that the child must wait for the parent or that the parent must wait for
the child. Use a real blocking dependency to order work between independent
issues, including sibling issues under the same parent.

Blocked state does propagate through the hierarchy: when an active blocker
holds up a parent, its descendants are not independently ready. This becomes a
feedback loop if an ancestor also has a `blocks`, `conditional-blocks`, or
`waits-for` edge to one of those descendants. The descendant blocks the
ancestor, then the blocked ancestor suppresses the descendant. Do not add
blocking edges in either direction between an ancestor and descendant.

The upstream `bd dep cycles` check does not detect every hierarchy feedback
loop. Run the repository guard directly when reviewing dependency changes:

```bash
python3 scripts/beads/check-hierarchy-deadlocks.py
```

The guard rejects active ancestor/descendant blocking edges, permits valid
sibling ordering, and reports closed or pinned historical edges without
failing. It is also part of `scripts/beads/check-health.sh`.

## Recovery

For a fresh clone or a missing/stale local database, bootstrap from the Git-hosted Dolt ref:

```bash
bd bootstrap --yes
bd stats
```

If bootstrap warns that tracker files are too broadly readable, run:

```bash
chmod 700 .beads
```

For an operator-owned point-in-time backup, use a path outside the repository or a separately administered DoltHub destination. Never commit credentials or a machine-local absolute backup path:

```bash
bd backup init <external-path-or-DoltHub-url>
bd backup sync

# In an initialized disposable recovery checkout:
bd backup restore <backup-path> --force
bd stats
```

The validated recovery hierarchy is:

1. `bd dolt pull` for normal collaboration;
2. `bd bootstrap --yes` for a fresh or missing database;
3. `bd backup restore` for an explicit point-in-time backup;
4. `.beads/issues.jsonl` only for exact pre-Dolt legacy provenance.

See the upstream [Beads sync concepts](https://github.com/gastownhall/beads/blob/main/docs/SYNC_CONCEPTS.md) for the distinction between Dolt synchronization and JSONL interchange.
