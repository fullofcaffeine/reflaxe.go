# Serializer Typed-Accessor Second-Pass Review

Date: 2026-07-27  
Bead: `haxe_go-vfp.10.5.1`  
Reviewer: independent `gpt-5.6-sol` agent at `xhigh`

## Decision

Accept the shared Reflect/Type metadata design. Serialization no longer owns an
unsafe field lift, a private invocation bridge, or duplicate metadata. Generated
member discovery and field access are typed; erased invocation still uses the
existing safe `Reflect.callMethod` runtime boundary.

This is a bounded extension of the AST-first generated-metadata pipeline. It
does not justify a universal semantic IR.

## Findings and dispositions

1. Integral wire tokens could fail to populate declared `Float` fields because
   generated assignment originally accepted only exact Go types.
   - Fixed with one generated `int` to `float64` branch.
   - `serializer_typed_accessor_contract` proves `i3` becomes `3`, not zero.
2. The initial footprint wording omitted the shared `runtime/hxrt/reflect.go`
   dependency and the selective-runtime performance harness did not exercise
   serialization.
   - Documentation now names the safe shared dependency.
   - The harness and baseline now include a private-field serialization round
     trip with full-versus-selective source and linked-binary measurements.
3. The decision trail overstated historical red evidence and fully typed
   invocation.
   - Append-only correction rows distinguish current replayable evidence from
     observed history and distinguish typed access from shared erased calls.

The reviewer re-ran a fresh read-only pass after these changes and reported no
remaining findings.

## Evidence

- 12 serializer semantic-diff contracts pass.
- The serialization snapshot and runtime smoke pass.
- The selective-runtime hard budget passes with the serialization lane.
- Compiler debt is ratcheted at `go_unsafe=2` and `compiler_shim=8`.
- Full validation and exact-commit release evidence are recorded in the Bead
  before closure.

The review exchange was provided through the agent mailbox; the host did not
expose a separate transcript file to commit.
