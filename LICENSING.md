# Licensing and Generated-Output Policy

## Status

This policy is **unresolved and release-blocking**. It records verified source
facts and the decisions that still require an authorized project owner or
qualified legal reviewer. It does not make a legal conclusion by implication,
and it is not legal advice.

The machine-readable authority is `license-policy.json`. A Haxelib package may
be assembled for testing while this record is unresolved, but the publication
entrypoints must reject it.

## Verified source facts

- The repository `haxelib.json` and `package.json` declare
  `GPL-3.0-only`, and the repository ships the corresponding root `LICENSE`.
  That declaration does not, by itself, answer how generated user code or
  copied runtime code should be treated.
- `runtime/hxrt/*.go` is repository-authored source that the compiler copies
  verbatim into generated Go modules as `hxrt/*.go`.
- The vendored Reflaxe framework is reconstructed from official upstream
  commit `430b4187a6bf4813cf618fc3a73ccf494a2ab9f5` plus the supplier and local
  patch chain recorded in `provenance/reflaxe/vendor-manifest.json`. That
  upstream commit declares MIT and its exact license text is preserved at
  `vendor/reflaxe/LICENSE`.
- Haxe 4.3.7 states that its standard library is MIT-licensed. The applicable
  notice is preserved at `licenses/HAXE-STDLIB-MIT.txt`. Four staged RTTI
  overrides retain the Haxe Foundation notice in-file; the wider staged
  override tree contains both repository-authored target code and code adapted
  from Haxe standard-library contracts. The exact outbound treatment of that
  mixed tree and its lowered Go remains unresolved.

Authoritative upstream references:

- Haxe 4.3.7 license statement and text:
  <https://github.com/HaxeFoundation/haxe/blob/4.3.7/extra/LICENSE.txt>
- Reflaxe license at the recorded upstream commit:
  <https://github.com/SomeRanDev/reflaxe/blob/430b4187a6bf4813cf618fc3a73ccf494a2ab9f5/LICENSE>

## Generated-output classes

The decision must address each output class separately:

1. `lowered-user-program`: Go produced from user-owned Haxe input.
2. `compiler-emitted-framework-support`: declarations or helpers emitted from
   compiler-owned templates and lowering logic.
3. `copied-hxrt-source`: verbatim runtime files copied from `runtime/hxrt`.
4. `lowered-haxe-standard-library`: Go emitted from staged or upstream Haxe
   standard-library source selected into the program.

No class inherits a conclusion from the repository-level license declaration.
Each needs an explicit `licenseTreatment` and a deterministic list of required
notices, license files, or source-offer material before release approval.

## Required authorization

Approval must record all of the following in `license-policy.json`:

- `decidedBy`: the reviewer's real name or accountable project identity;
- `authority`: either `project-copyright-owner` or
  `qualified-legal-review`;
- `decisionDate`: an ISO `YYYY-MM-DD` date;
- `decisionRecord`: an immutable Beads comment, reviewed document, or other
  durable decision identifier;
- `scopeSha256`: the verifier-produced digest of the shipped-source coverage,
  components, generated-output classes, and required package material that was
  actually reviewed.

Every generated-output class must then replace `unresolved` with the approved
treatment and replace `requiredArtifacts: null` with an explicit array. Every
component's generated-output treatment must likewise be resolved, and
`unresolvedQuestions` must be empty.

Changing the reviewed scope invalidates the digest and blocks publication until
the authorized reviewer approves the new scope. Missing or changed license
material also blocks publication.

## Enforcement

`python3 scripts/release/verify-license-policy.py --mode audit` validates the
inventory, source coverage, notice hashes, and optional staged-package bytes
without pretending that an unresolved policy is approved.

`python3 scripts/release/verify-license-policy.py --mode release` adds the
authorization checks and fails while any decision is unresolved. Both the
same-SHA release wrapper and the GitHub publication job invoke release mode
before semantic-release can create or publish a tag.

The Haxelib package includes this document, the machine policy, the project
license, the Reflaxe license, and the Haxe standard-library notice. Candidate
artifact construction remains possible so those bytes can be tested; public
release remains impossible until this policy is explicitly approved.
