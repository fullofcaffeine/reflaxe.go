# M03-08 Oracle disposition

Request: `orq_20260820T214220Z_3faf90d0`

Processor: `gpt-5.6-sol`, reasoning `xhigh`.

Oracle advice is independent evidence. The local agent remains responsible for
the plan, implementation, tests, and merge decision.

## Local baseline

The response matches the durable request:

- selected revision: `b98ed0d54c214955330eda964df46eff5b213dbd`;
- prompt SHA-256: `c51f94dcf7d57682ec2894f067d863c3c518e9affcc770c4f4fc1811255a2452`;
- bundle SHA-256: `54d9d4ba17e01fa3283d5585f6c872d7a8aee757ff8ef02da63b7f29add0a61b`;
- conversation handle: `convo_d3e00a4229fff7de`;
- completion binding: `834954da846cd9510b5822bfca753aaa3d0f3d7dc258d5d568e03cd19a7482ed`;
- valid final `Pro` model proof with no attention flag.

Current repository evidence:

- Existing-module output is written directly into a caller-owned Go module.
- Reflaxe's generic `_GeneratedFiles.json` stores paths but no content digests.
  Current replacement, deletion, and package ownership decisions trust those
  paths.
- Three new focused tests are red on the selected revision. They prove that a
  rerun overwrites a caller-modified generated file, forged legacy inventory
  deletes caller source, and forged inventory grants compiler-main ownership.
- Seven existing focused confinement tests are green. They prove traversal,
  output-root and runtime symlink rejection, poisoned traversal metadata
  rejection, reserved-file alias rejection, and protected-byte preservation.
- Package ownership is checked during project-manifest resolution. Therefore,
  the new verified ownership snapshot must be available before generation.
- `docs/existing-go-module-mode.md` already defines
  `packageDir/.reflaxe-go-owned.json` as the public ownership location and says
  legacy `_GeneratedFiles.json` is not mixed-tree authority.
- Haxe eval exposes typed SHA-256, exclusive file-open/lock flags, `fsync`, and
  rename primitives. No shell or raw target injection is necessary.
- `GoPostBuildRunner` provides the local test pattern: a typed optional effect
  function exercised through a direct `--interp` fixture.

Tests actually run before reconciliation:

- three new individual package-output tests: expected failures reproduced;
- `test/test_generated_output_confinement.py`: 5 passed;
- existing-module symlink preservation: 1 passed;
- reserved module-file alias rejection: 1 passed;
- `git diff --check`: passed before the final test addition.

Oracle did not run tests because its selective bundle was not a runnable Haxe
checkout.

## Oracle claim matrix

| Claim | Disposition | Local evidence and consequence |
| --- | --- | --- |
| Generic Reflaxe inventory conflates output listing with mixed-tree ownership. | Retained | The three red tests reproduce overwrite, deletion, and false package ownership from path-only metadata. |
| ExistingModule needs one dedicated digest-backed transaction; Standalone stays on the generic manager. | Retained | This isolates the stronger policy at the lowest owner and preserves Standalone compatibility. |
| Collect every artifact before mutation through a typed sink. | Retained | Current preflight omits licenses, runtime files, optional reports, and `extraFiles`. A complete plan is required to make the conflict check atomic. |
| Rerooting the generic manager into a private stage is sufficient. | Rejected from the provisional local plan | It avoids direct caller mutation but keeps generic metadata and emission rules inside the authority path. A typed artifact sink makes completeness and duplicate detection directly testable with fewer implicit states. |
| Use the documented package-local ownership manifest with exact paths and SHA-256 digests. | Retained | The location is already public. A closed deterministic schema is sufficient; timestamps, PIDs, phases, and absolute paths add no ownership evidence. |
| Require both an immutable journal and digest verification. | Retained | A digest-only partial commit is indistinguishable from a caller edit; a journal without digests can restore or delete the wrong bytes. |
| The installed ownership manifest is the commit marker; journal recovery classifies exact old, absent-with-backup, exact new, or conflict. | Retained | This provides one durable authority without a mutable phase machine. Unknown state preserves evidence and stops. |
| Automatic recovery requires an ephemeral exclusive writer lock. | Rejected for M03-08 | The existing output boundary explicitly supports one non-hostile compiler writer. A cross-host process-lock adapter would create a new concurrency contract that Beadshx does not need for its first vertical tracer. Simultaneous writers remain documented as unsupported. |
| Package inspection must use a verified digest snapshot, never prefix plus legacy inventory. | Retained | Manifest resolution performs this check before output generation, and the new forged-ownership test demonstrates the current hole. |
| Strict closed JSON decoding is required for manifest and journal. | Retained | Unknown versions and fields, duplicate or case-aliased file paths, malformed digests, reserved paths, and wrong project identity fail before mutation. The host JSON escape remains inside one typed codec. |
| The project manifest must be reserved and stable during commit. | Narrowed | Generated output cannot target the project manifest. Concurrent hostile mutation remains outside the documented one-writer boundary, so M03-08 does not add a second file guard beyond `go.mod` and `go.sum`. |
| Structured build output is outside the generated-artifact transaction. | Rejected from this prerequisite | The structured build report is generated output and joins the transaction. A caller-selected `go build -o` binary remains a governed build effect under the existing interface; changing its ownership semantics would broaden this final prerequisite. |
| Existing legacy output may be adopted from `_GeneratedFiles.json` or matching newly generated bytes. | Rejected | Neither source proves current ownership. Migration is manual removal and regeneration; any future adoption command must require explicit user authority. |
| Package relocation should be automatic now. | Deferred | The public manifest is package-local. Relocation remains an explicit migration until a concrete compatibility requirement justifies cross-package discovery. |
| Full power-loss durability is part of this task. | Deferred | M03-08 requires deterministic process interruption. File and directory synchronization under sudden power loss is a distinct contract and needs separate platform evidence. |
| Defend against an unrelated hostile process swapping filesystem components between every check and effect. | Deferred | Current scope serializes haxe.go writers and rejects observed symlinks. Descriptor-relative hostile-filesystem defense is a stronger security boundary than the documented local non-hostile caller model. |
| Remove empty output directories during rollback or stale cleanup. | Rejected | Directory removal is unnecessary for correctness and increases ownership surface. Empty directories may remain. |

## Integrated conclusion

The following was the broad Oracle recommendation recorded before
implementation. The final local implementation deliberately narrows it in the
post-implementation section below.

Implement one ExistingModule-only `GoExistingModuleOutputTransaction` and one
typed artifact-emission seam. Do not change Reflaxe or Standalone behavior.

1. Add a shared `GoGeneratedArtifactSink`. The Standalone implementation uses
   the existing boundary and `OutputManager`. The ExistingModule builder
   converts every generated source, runtime file, license, report, structured
   build report, and `extraFiles` value into exact bytes and one validated plan
   entry. Preserve current `extraFiles` priority and concatenation behavior.
2. Add typed digest, ownership-manifest, snapshot, journal, operation, and
   strict codec values. Reserve `go.mod`, `go.sum`, the project manifest,
   `_GeneratedFiles.json`, ownership metadata, transaction control paths, and
   file/descendant or case aliases.
3. Change `GoPackageDirectoryInspector` to consume a verified ownership
   snapshot. Remove legacy metadata and filename-prefix authority from
   ExistingModule.
4. Add a narrow eval filesystem adapter. Prove process-lifetime exclusive
   locking, create-new, no-replace same-filesystem move, close-before-return,
   no-follow state inspection, confinement recheck, and digest-checked removal.
   Stop implementation if a supported host cannot prove the needed primitive;
   never silently weaken it.
5. Implement full preflight, immutable journal creation, stage verification,
   old-byte backup, new-byte installation, complete generation verification,
   and ownership-manifest-last commit. Discover every conflict before the
   journal and recheck immediately before the first destructive move.
6. Recover under the exclusive lock before a new transaction. Roll back in
   reverse order when the old manifest state is proved; finalize cleanup when
   the exact new manifest is installed; preserve all evidence and stop on a
   third digest, symlink, missing required backup, malformed state, or unknown
   version.
7. Keep the transaction handle through `onOutputComplete()`. Build structured
   binaries into the transaction staging area. A successful build installs its
   bytes through the same manifest authority; a failed build keeps committed
   generated sources and the prior verified binary, then cleans the completed
   source transaction before reporting the build error.
8. Sever ExistingModule from `OutputManager.generateFiles()`, `saveFile()`,
   stale deletion, and metadata recording. Keep those calls structurally
   unchanged for Standalone.
9. Update the public contract for schema, process-interruption recovery,
   lock/busy behavior, manual legacy and relocation migration, build-output
   ownership, and explicit non-goals.

No unresolved product-owner decision blocks the narrower implementation. The
final local choices are:

- process interruption is required now; power-loss durability is deferred;
- automatic recovery supports one non-hostile compiler writer and does not
  claim simultaneous-process serialization;
- generated structured-build reports join this ownership boundary, while
  caller-selected binaries keep their existing governed-build behavior;
- package relocation and legacy adoption remain explicit manual migrations.

## Post-implementation local refinement

The implementation keeps the smallest complete boundary needed by Beadshx:

1. Existing-module generation collects exact source, runtime, license, report,
   and macro-extra bytes in `GoExistingModuleOutputPlan`. Standalone generation
   remains on Reflaxe's existing output manager without a new shared sink layer.
2. `.reflaxe-go-owned.json` stores sorted module-relative paths and SHA-256
   digests. Package inspection, replacement, and stale cleanup require the exact
   path spelling and current digest; legacy metadata grants no authority.
3. `GoExistingModuleOutputTransaction` stages new bytes, backs up verified old
   bytes, publishes an immutable journal, installs artifacts, and writes the
   ownership manifest last. The next invocation rolls back a pre-commit state,
   cleans an exact committed state, or preserves unknown evidence and stops.
4. The transaction reuses `GoGeneratedOutputBoundary` for traversal,
   containment, symlink, reserved-path, and path-redacted diagnostics. No raw
   target injection, shell filesystem operation, new process lock, power-loss
   promise, or hostile-filesystem promise was added.
5. Package relocation, legacy adoption, caller-selected build binaries, and
   simultaneous writers remain explicit non-goals. They do not block the first
   authored-Haxe Beadshx tracer.

The mandatory xhigh second pass found one case-alias bug: snapshot lookup used
a lowercased key and could treat a differently cased live path as owned on a
case-sensitive filesystem. Lookup now requires the exact stored spelling, a
case-only output rename fails before mutation, and a Linux-capable regression
preserves both paths. No other blocking correctness finding remained after the
transaction-order, rollback, stale-cleanup, manifest-last, confinement, and
caller-byte review.

## Verification and unresolved gaps

The implementation is complete. Evidence run before the haxe.go PR:

- existing-module specification: 3 passed;
- existing-module preservation: 5 passed;
- existing-module package output: 22 passed locally, with the case-sensitive
  alias regression skipped on this case-insensitive host and available to Linux
  CI;
- existing-module structured build: 8 passed;
- generated-output confinement: 5 passed;
- `npm run test:changed`: passed end to end;
- `npm test`: 313 snapshots passed with all surrounding contract gates green;
- `npm run test:examples`: 13 passed;
- compiler-debt, warning, typed-identifier, dynamic-boundary, macro-lifecycle,
  canonical-stdlib, Haxelib-package, runtime race, socket, and terminal gates:
  passed;
- Haxe formatter check and `git diff --check`: passed.
