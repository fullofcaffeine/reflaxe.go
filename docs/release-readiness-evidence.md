# Release Readiness Evidence

## What it is

Release readiness is the machine decision that one exact haxe.go commit may be
published with one exact public claim and artifact set. It does not replace the
focused compatibility, API, security, licensing, or artifact checks. It joins
their results and rejects a green result assembled from different commits.

The policy authority is `release/readiness-policy.json`. The evidence document
uses `kind: haxe.go-release-readiness-evidence` and `schemaVersion: 1`.

## Why it exists

Before this gate, `npm run release:status` proved that release tooling was
wired correctly, but it was deliberately permissive about the latest hosted
release. Healthy wiring is not proof that a particular candidate is ready.
For example, a candidate could still have:

- a tag, manifest, or artifact built from a different commit;
- a public support claim broader than the compatibility evidence;
- a missing checksum or provenance statement;
- a vulnerable reachable dependency or failed API policy check;
- an unresolved licensing question;
- an applicable high-priority blocker; or
- hosted assets that differ from the locally verified bundle.

The readiness gate makes those contradictions one fail-closed verdict while
leaving each detailed authority in its existing focused file.

## How it works

1. The CI Harness runs quality, secret scanning, dependency audit, Go tooling,
   and performance jobs for `${{ github.sha }}`.
2. The release job can start only after every required job passes.
   `collect-upstream-release-evidence.py` records the actual
   `needs.*.result` values, resolved Haxe and Node versions, exact Go matrix,
   the exact hosted runner image reported by the Linux quality job, public-API
   result, security lanes, and tested SHA in a structured document.
   A missing or non-success job result cannot become passing evidence.
3. The workflow records the same commit as `RELEASE_TESTED_SHA` and
   `RELEASE_UPSTREAM_GATES_SHA`; the same-SHA wrapper also requires the
   structured upstream document through `RELEASE_UPSTREAM_EVIDENCE`.
4. The artifact builder creates the Haxelib ZIP, checksum, content manifest,
   provenance, and `release-assets.json`. The independent asset verifier checks
   their bytes and identities.
5. The workflow creates one blocker-evidence file from an isolated Beads client
   pointed at the configured remote. Collection fails if `refs/dolt/data`
   advances while those records are being read. The same immutable file is
   then reused for both release phases, so a tracker update halfway through a
   release cannot make the two decisions describe different tracker states.
   The final portable-admission Bead also carries a structured review record:
   the Oracle request, frozen packet digest, Oracle-reviewed candidate SHA and
   verdict stay separate from the local correction disposition for the exact
   SHA that may ship.
6. The collector creates `candidate` evidence from governed repository files,
   the structured upstream results, the verified asset manifest, and that
   remote blocker evidence.
7. The readiness verifier checks the candidate after Semantic Release selects
   the exact tag but before hosted assets are reconciled.
8. Release reconciliation creates or completes only the allowed GitHub state.
9. The collector creates `published` evidence. In `live` mode the verifier
   ignores caller-supplied hosted facts and queries the GitHub API. The tag
   commit, immutable state, asset names, and API digests must exactly match the
   tested SHA and local asset manifest.

The transaction remains retry-safe. A failed candidate check prevents hosted
release reconciliation (the exact tag may already exist for a safe retry). A
failed published check reports incomplete or contradictory hosted state; the
existing reconciliation contract permits only safe missing-state completion
and never overwrites a conflicting immutable asset.

## Evidence fields

| Field | Meaning |
| --- | --- |
| `release` | Version, tag, tested SHA, source SHA, and artifact-manifest tag must agree. |
| `compatibility` | Exact lifecycle statement, admitted scopes, owned exclusions, and a canonical SHA-256 over every admitted operation/member, symbol, evidence row, platform, preset, and required trust assumption. |
| `publicApi` | The public SemVer boundary passed for the tested SHA. |
| `platform` | The exact admitted OS/architecture and hosted runner image OS/version reported by the quality job. |
| `toolchains` | Actual resolved Haxe and Node versions plus the exact successful Go matrix, checked against governed support lines. |
| `security` | Structured GitHub job results for the tested SHA and zero reachable vulnerabilities. |
| `licensing` | Approved scope digest and no unresolved questions, matching `license-policy.json`. |
| `blockers` | Status, priority, and scopes used to block only unresolved P0/P1 work intersecting admitted scope. The final admission owner additionally binds Oracle provenance and the local correction disposition to the exact release SHA. |
| `artifacts` | Required roles, names, SHA-256 digests, and provenance subjects for the tested SHA. |
| `github` | `null` for a candidate; live GitHub API truth for a published release. |

Unknown top-level fields fail. Missing lists do not mean “none”; they fail
schema validation. Exclusions must be visible and owned, and an excluded scope
cannot be advertised as supported.

## Roadmap blockers and exclusions

The admitted beta surface is `preset:portable` on
`platform:linux-amd64`, bounded further by the operation/member manifest.
Networking is split at that operation boundary: only the exact TCP, UDP, and
TLS member groups marked `release_admitted` participate. Named/reverse DNS,
HTTP, unlisted members, and the manifest's other explicit exclusions remain
outside the claim. Go-native, portable HTTP, and stable-1.x retain named Beads
owners, so open work in those excluded scopes does not silently block the
bounded beta.

The blocker evidence records the real priority and state of every compatibility
owner, plus any temporary review blocker that the readiness policy attaches to
an already-admitted preset or platform. A temporary review blocker cannot
invent a new scope or masquerade as a permanent compatibility exclusion. This
lets the tracker close an admission review and unblock release without a
follow-up source commit, while an open P0/P1 review still fails closed.

The final admission owner is stricter than an ordinary closed blocker. Its
`releaseAdmission` metadata has two independent parts:

- `oracleReview` identifies what the external reviewer actually saw: the
  reviewed source SHA, request ID, frozen packet SHA-256, and returned verdict.
- `localDisposition` identifies the correction-focused pass over the source
  that can actually ship: its exact SHA, accepted verdict, disposition-document
  SHA-256, processor model, and reasoning level.

The readiness verifier requires the local reviewed SHA to equal the release
tested SHA. The surrounding blocker evidence is itself pinned to one remote
`refs/dolt/data` commit, so the admission metadata and its tracker revision are
one immutable input to both candidate and published decisions. This avoids two
misleading shortcuts: claiming Oracle reviewed the successor when it reviewed
only the predecessor, or treating a closed issue with no exact-SHA disposition
as release approval.

The release workflow generates this evidence with
`scripts/release/refresh-readiness-blockers.py` from an isolated client of the
configured remote and records the exact `refs/dolt/data` commit. The completed
hosted-artifact owner (`haxe_go-vfp.4.8`) therefore remains visible as closed
rather than being fabricated or silently dropped. Omitting a compatibility
owner, using an unadmitted scope for an additional review owner, or failing to
account for every policy blocker makes verification fail. Changing a blocker's
state is picked up by the next workflow run; the current run deliberately keeps
using its one captured tracker state through candidate and published
verification.

If policy admits one of those scopes, the exclusion must be removed and its
applicable P0/P1 blocker must be closed. Changing a claim without changing its
evidence and ownership fails the gate.

## Commands

Run the deterministic contract without a real release:

```bash
python3 test/test_release_readiness_gate.py
python3 scripts/release/refresh-readiness-blockers.py \
  --observed-at "$(date -u +%F)" \
  --output /tmp/haxe-go-release-blockers.json
npm run release:readiness -- \
  --evidence test/fixtures/release_readiness/published-pass.json \
  --mode fixture
```

Production operators do not hand-build evidence or use fixture mode. A green
push to `master` starts the release job. The same-SHA wrapper collects and
verifies both release phases automatically.

The live GitHub-settings audit is separate because the workflow token cannot
read repository administration settings. Run
`npm run security:github-governance:live` before release-infrastructure changes
and during periodic repository security maintenance.

Related contracts:

- [Release readiness checklist](release-readiness-checklist.md)
- [Compatibility release status](compatibility-release-status.md)
- [Public contract and SemVer boundary](public-contract.md)
- [Release version and source identity](release-version-policy.md)
- [Release retry and reconciliation](release-reconciliation.md)
- [GitHub governance policy](github-governance-policy.md)
- [Independent xhigh second-pass review](reviews/release-readiness-vfp-6.5.md)
