# Portable beta Oracle disposition

## Local baseline

Before applying the response, the repository already had a narrow generated
compatibility authority: 38 admitted operations / 173 symbols for the portable
Linux/amd64 pre-1.0 beta, with Go-native/metal and unlisted operations kept
separate. The frozen review source was an ancestor of the current branch. Local
inspection reproduced four evidence/governance defects but no defect in the
admitted compiler/runtime implementation and no reason for a universal IR or a
new profile-specific semantic backend.

## Outcome

The Oracle response is accepted as advisory evidence with four retained
release-evidence findings. It did not identify a defect in the admitted
compiler/runtime behavior or a reason to replace the typed staged-Haxe / `hxrt`
architecture.

The local disposition is `ACCEPTED_AFTER_CORRECTIONS` for the narrow pre-1.0
portable beta only. It is not approval for stable 1.x, Go-native/metal,
unlisted compatibility operations, another platform, or a broader trust model.

## Provenance kept intact

- Oracle request: `orq_20260803T080027Z_615d041e`
- Oracle-reviewed source: `c22af0ea82e5e481e23277e513ed5b7c6b5c770b`
- Frozen outer packet SHA-256:
  `ce8cc8cca65d48dceae11155e9fc651dbf8bef7f611270e7b0a43c194e33d5f9`
- Oracle verdict: `READY_AFTER_BOUNDED_FIXES` for the portable beta,
  `READY_WITH_LIMITS` for bounded production use, and `NOT_READY` for stable
  1.x.

The exact corrected source SHA is deliberately not written into this file:
doing so would change the commit it was trying to name. Instead,
`haxe_go-vfp.12.5` stores a structured `releaseAdmission` record after the
correction-focused gates pass. The release workflow reads that record from one
immutable `refs/dolt/data` revision and requires its local reviewed SHA to equal
the release tested SHA. Oracle's original reviewed SHA remains a separate field,
so the project never claims Oracle inspected the successor.

## Finding dispositions

| Finding | Local decision | Correction |
| --- | --- | --- |
| Final review was not a machine release blocker | Retained | `haxe_go-vfp.12.5` now governs `preset:portable`. Release evidence requires the Oracle provenance plus an exact-SHA local disposition and the SHA-256 of this document. |
| Public/package wording hid admitted concurrency and networking subsets | Retained | README and known-gaps language now distinguishes named admitted subsets from the excluded remainder. The Haxelib package ships the manifest, matrix, release status, known gaps, toolchain policy, and readiness checklist that explain the claim. |
| Classic branch-protection 404 was summarized as no protection | Retained | Evidence generation now reports the classic endpoint separately from active repository rulesets. A checksum-bound erratum preserves the frozen packet and records both active rulesets. |
| “12 examples” mixed program and profile units | Retained | Evidence generation derives and labels 8 maintained program directories, 13 profile cases, and 1 release-claim-bearing case from the exact-source QA manifest. The frozen packet has a checksum-bound erratum. |

## Recommendation not adopted

The Oracle suggested another independent review of this bounded correction
delta. No new Oracle request is dispatched. The requester explicitly prohibited
another dispatch, and the repository policy permits a documented local xhigh
correction-focused second pass for this situation. That pass must challenge the
new tests' mutation sensitivity, the independence of the package/evidence
oracles, exact-SHA joins, unsupported claim expansion, and preservation of the
original Oracle provenance before the tracker admission record is written.

## Verification and remaining limits

Processor: `gpt-5.6-sol` at `xhigh`, matching the ledger policy.

The retained changes were verified locally with:

- focused red-to-green mutation contracts for final admission, host-control
  summaries, example units, the checksum-bound erratum, and packaged support
  authorities;
- `npm run test:changed`;
- `npm test`: 313 snapshot cases passed;
- `npm run test:examples`: 13 real Haxe-to-Go build/run profile cases passed;
- `npm run test:official-haxe-smoke`: 3 installed-target smoke cases passed;
- `npm run compatibility:verify` and `npm run release:status`;
- `python3 test/test_haxelib_release_artifact.py`: 5 deterministic exact-commit
  artifact contracts passed;
- `npm run test:haxelib-release-install`: 2 isolated Haxelib install/run
  contracts passed; and
- `npm run test:release-contracts`: the complete release contract catalog
  passed on the corrected commit.

The first exact-commit artifact run exposed a test-only variable-scope error;
that test was corrected and the exact artifact, isolated install, and release
catalog were rerun successfully. This is recorded rather than hidden because it
demonstrates that the release proof was sensitive to the new package members.

Remaining limits are unchanged: stable 1.x, Go-native/metal admission, HTTP,
named/reverse DNS, IPv6, non-Linux runtime admission, untrusted-source
compilation, and public/hostile-network guarantees still require their own
evidence or remain explicitly excluded. Oracle's proposed additional delta
review was not run and is not represented as evidence.

## Closure rule

This disposition justifies processing and archiving the captured Oracle request
only after the focused contracts, package/install proof, release contracts, and
the repository's required full gates pass for the corrected source. Archive
means that this response has been reconciled; it does not mean every optional
roadmap item or stable-release requirement is complete.
