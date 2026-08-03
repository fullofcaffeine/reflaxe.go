# Portable beta Oracle disposition

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

## Closure rule

This disposition justifies processing and archiving the captured Oracle request
only after the focused contracts, package/install proof, release contracts, and
the repository's required full gates pass for the corrected source. Archive
means that this response has been reconciled; it does not mean every optional
roadmap item or stable-release requirement is complete.
