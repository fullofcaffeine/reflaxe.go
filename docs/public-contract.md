# Public Contract and Semantic Versioning Boundary

This document defines what a haxe.go release number protects. It is deliberately
an index over the files and tests that already own product truth, not a second
database of every Haxe declaration.

The short version is:

1. Users and reviewers decide whether a change affects a documented public
   contract.
2. The commit records that impact as a Conventional Commit.
3. Semantic Release maps the commit to the next version under the reviewed
   [`0.x` or stable-major policy](semver-lifecycle-policy.md).
4. CI checks the underlying compatibility, packaging, behavior, and release
   evidence. It does not guess intent from a source-code diff.

That means an agent may propose the classification, but it does not have
unbounded authority over versioning. The public boundary below, executable
contracts, review, and the explicit stable-major approval gate constrain the
decision.

## What SemVer protects

Only the documented or explicitly admitted part of a surface is protected.
The authoritative inventory and evidence column identifies where its exact
boundary lives.

| Surface | Protected treatment | Authoritative inventory and evidence |
| --- | --- | --- |
| Package and installation | Published package identities, public Haxe import roots, documented entrypoints, required Haxelib layout, and release sidecar identities are release contracts. | `haxelib.json`, `package.json`, `scripts/release/build-haxelib-artifact.py`, and `test/test_haxelib_release_install.py` |
| Toolchains and platforms | Only the exact Haxe, Go, Node, operating-system, and architecture combinations admitted by policy are compatibility promises. A passing unlisted environment is not silently admitted. | `docs/toolchain-policy.json`, `docs/compatibility-support-manifest.json`, `test/test_toolchain_policy_contract.py`, and `npm run compatibility:verify` |
| Profiles and source controls | Documented defines, metadata, macro entrypoints, `portable|metal`, compatibility aliases, and `@:goNative` behavior are public inputs. `metal` remains a convenience policy preset, not a second semantic product. | `docs/native-policy-presets.md`, `docs/profile-semantics-guide.md`, `docs/defines-reference.md`, and `test/test_metal_preset_retention_contract.py` |
| Portable Haxe semantics | A portable behavior is protected only at the operation/member, platform, and trust boundary explicitly admitted by the compatibility manifest. A stdlib inventory row alone is evidence, not blanket support. | `docs/compatibility-support-manifest.json`, `test/portable_stdlib_inventory.json`, `test/run-semantic-diff.py`, and `npm run compatibility:verify` |
| Go-native source APIs | Documented `go.*` facades, extern metadata, macros, and native-boundary behavior are consumer-visible source APIs at their declared evidence state. Experimental entries are not stable admission, but incompatible changes still require explicit breaking-change classification and release notes. | `docs/goextern.md`, `std/go`, `docs/compatibility-support-manifest.json`, and `test/run-goextern-fixtures.py` |
| Generated Go and runtime ABI | Documented shapes consumed by handwritten Go are protected where a contract explicitly says so. Compiler-generated code and its matching `hxrt` package must remain mutually compatible within one artifact. The typed Go IR and unadvertised helper names remain internal. | `docs/typed-go-ir.md`, `docs/class-dispatch-abi.md`, `runtime/hxrt`, and `test/snapshot/core/ast_typed_type_operator_printer` |
| Reports and versioned data | Documented report fields, schemas, release identity files, checksums, and fail-closed reader rules are public data contracts. Ordering and fields explicitly marked internal are not. | `docs/profile-semantics-guide.md`, `test/snapshot/core/report_artifacts_basic`, release artifact manifests, and report-schema snapshots |
| Commands and diagnostics | Documented user commands, flags, failure meaning, source locations, and remediation are public behavior. Exact prose and repository-only task composition are not machine-parsing APIs. | `docs/start-here.md`, `docs/defines-reference.md`, `test/test_post_generation_build_runner.py`, and `test/snapshot/negative` |

Support status still matters. `supported` or release-admitted behavior receives
the stated compatibility promise. `experimental`, `compile-only`,
`compatibility-only`, and `excluded` retain the narrower meanings in
`docs/compatibility-support-manifest.json`. Calling a surface public for change
classification does not upgrade its support status.

## What is internal

The following can change without a public API bump when the protected behavior
above remains intact:

- compiler-private types, passes, temporary metadata, and vendored layout;
- the typed Go IR described as internal in `docs/typed-go-ir.md`;
- exact whitespace, private locals, file splitting, and incidental generated
  formatting not named by an interop or debugging contract;
- unadvertised `hxrt` helper names when compiler, runtime, and packaged output
  change together without affecting consumers;
- snapshot case names, CI job names, repository-only script composition, and
  tracker structure; and
- exact diagnostic wording when meaning, source location, and actionable
  remediation remain equivalent.

Readable generated Go is an important debugging and quality property. A
readable emitted byte or private helper is not a public API merely because it
can be inspected.

## How a release number is chosen

The person or agent authoring a change evaluates its consumer impact against
this document and writes the matching Conventional Commit. Reviewers correct a
misclassification just as they would correct code or documentation. The
release analyzer then applies the mechanical policy in
[Release Version and Source-Identity Policy](release-version-policy.md):

| Consumer impact | Commit signal | Version effect |
| --- | --- | --- |
| Compatible bug, performance, or security fix | `fix:` or `perf:` | patch |
| Additive public capability | `feat:` | minor |
| Incompatible public change | `!` in the header or `BREAKING CHANGE:` footer | breaking change on `0.x` becomes a minor; a stable line requires the next approved major |
| Internal refactor with no user-visible change | `refactor:`, `test:`, `build:`, or `chore:` as appropriate | no release by the current analyzer |
| Documentation only | `docs:` | no release |

CI can prove that manifests, tests, and exact-SHA release mechanics agree. It
cannot infer whether a renamed public concept is acceptable to users. That is
why the commit classification remains an accountable authored decision rather
than an automatic declaration diff.

Promotion to `1.0.0`, or to any later stable major, requires explicit human approval
in release policy. Semantic Release cannot turn an ordinary breaking
change on `0.x` into `1.0.0` by itself.

## Deprecation and compatibility

The detailed windows, experimental treatment, and stable admission checklist
live in [SemVer and Compatibility Lifecycle Policy](semver-lifecycle-policy.md).

During `0.x`, an incompatible public change advances the minor version and
must carry migration guidance. A documented compatibility alias or selector is
still public even when it does not define separate semantics. In particular,
`metal` and `@:goMetal` remain compatibility inputs under their recorded
retention policies.

For a stable line, a public removal or incompatible behavior/schema change is
major-only unless an alias, adapter, or backward-compatible reader preserves
existing consumers. A normal rename or removal should first ship a documented,
actionable deprecation in an earlier minor release. Security or correctness
emergencies may use a reviewed exception with prominent migration notes.

## Why this is federated

haxe.go already has two detailed authorities:

- `docs/compatibility-support-manifest.json` owns admitted platform,
  preset, operation/member, evidence-state, and trust-boundary truth.
- `test/portable_stdlib_inventory.json` owns the portable module evidence
  inventory, while explicitly not widening operation-level release admission.

Packaging, reports, externs, generated ABI, and diagnostics likewise have
focused executable contracts. Copying all of them into a universal declaration
manifest would introduce another review surface that could disagree with the
owners. This document instead makes those owners compose into one public
boundary, and `test/test_public_contract.py` fails when a required owner or
release-document link disappears.

This choice follows commit-pinned sibling evidence:

| Sibling | Observed design | Lesson for haxe.go |
| --- | --- | --- |
| `haxe.ruby` at `ded7f02d666612350440d2d31e52dfe48449f9b9` | `docs/public-contract.md` explicitly keeps its public inventory federated across machine-checked owners. | Closest precedent: one readable policy plus focused inventories and upgrade/package evidence. |
| `haxe.elixir.codex` at `2030abea264dac770915dbeff427acc349ff082e` | `release/manifest.json` governs release lines and stable admission; its analyzer delegates Conventional Commit parsing to the official analyzer. | Keep version policy small and let executable feature contracts own feature truth. |
| `haxe.rust` at `85067736d0b929dfc67d6684d59b7e2bd3bae6ea` | Its 37,735-line public compatibility manifest enumerates 411 Haxe types plus metadata, defines, contracts, and evidence. | Useful when declaration-level coverage is the product problem, but too costly to duplicate here without evidence of recurring accidental API breaks. |
| `haxe.c` at `3a650c1481c2072e87a8ba89f9cdaaba29a244de` | Its bootstrap capability inventory gates an incomplete compiler milestone. | A capability inventory answers bootstrap readiness, not the general public-SemVer question. |

## When to revisit declaration-level diffing

Do not add a universal API manifest merely because one can be generated.
Reopen the decision only when at least one of these is true:

- stable `1.x` admission is close enough that compatibility with an immutable
  release baseline must be rehearsed;
- two or more real accidental public declaration breaks escape focused tests,
  showing a recurring gap rather than a hypothetical risk; or
- external consumers need a supported machine-readable declaration index that
  the existing owners cannot provide.

If reopened, compare typed Haxe compiler output against the last published
release. Do not use regular expressions over source, and do not let the new
artifact replace behavior, package, or runtime evidence. Its maintenance and
failure mode must be justified before it becomes release-blocking.

## Validation

Run the focused contract and its release suite with:

```bash
python3 test/test_public_contract.py
npm run test:release-contracts
```

Release candidates must also follow
[Release Readiness Checklist](release-readiness-checklist.md), which composes
this boundary with compatibility, packaging, security, and exact-SHA release
evidence.
