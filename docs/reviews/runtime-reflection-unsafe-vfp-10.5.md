# Runtime Reflection, Serializer, and Unsafe Boundary Review

Date: 2026-07-27
Bead: `haxe_go-vfp.10.5`
Reasoning level: `xhigh`

## Decision

Close the parent boundary after the current validation and independent review
pass. Haxe.Go has one production `unsafe.Pointer` operation: a typed
`syscall.Termios` pointer passed to the POSIX terminal ioctl. The serializer no
longer adds unsafe access or its own reflection mechanism. Remaining Go
reflection is confined to named runtime helpers and generated adapters for
values whose concrete Go type is intentionally erased.

This is not a claim that reflection is free or that every possible serializer
payload is supported. It is a claim that the admitted boundaries are explicit,
selective, dynamically tested, and prevented from growing silently.

## What, why, and how

- **What:** inventory every production unsafe operation and each remaining
  reflection ownership island; map the serializer, reflection, toolchain, and
  checkptr evidence to the parent acceptance criteria.
- **Why:** snapshots can compile while an unsafe conversion is invalid on a new
  Go release, and erased values can panic only when a particular nil,
  non-comparable, interface, cycle, or error path executes.
- **How:** remove avoidable unsafe access, ratchet the remaining imports and
  selectors by file, run supported Go lines with race/checkptr/vet/Staticcheck,
  execute the real unsafe terminal path in a checkptr-instrumented PTY program,
  and compare portable serializer/reflection behavior with Haxe 4.3.7.

## Unsafe inventory

| Location | Owner | Operation and invariant | Why retained | Dynamic evidence |
| --- | --- | --- | --- | --- |
| `runtime/hxrt/terminal_posix.go::terminalIoctlTermios` | build-tagged `hxrt` terminal capability selected by `Sys.getChar` | Convert exactly one live `*syscall.Termios` to `unsafe.Pointer`, pass it directly to `syscall.SYS_IOCTL`, and call `runtime.KeepAlive` after the syscall. No arithmetic, retained pointer, alias reconstruction, or untyped payload is permitted. | The frozen standard `syscall` API exposes the required structure and constants but no typed ioctl wrapper. A current `x/term` raises the generated Go floor; the compatible older dependency has a known advisory. | `test/test_sys_get_char_terminal.py` builds with `-gcflags=all=-d=checkptr=2` and drives no-newline input, echo off/on, and restoration through a real PTY. Redirected input/EOF and Linux, Darwin, Windows, and unsupported-host builds cover the surrounding typed capability. |

The compiler-debt policy permits exactly one `unsafe` import and one selector in
that file (`go_unsafe=2`). Generated program/runtime examples have an
independent zero-unsafe policy. `test/test_sys_get_char_terminal_contract.py`
also fails if another runtime file acquires `unsafe.Pointer`.

The serializer's former `reflect.NewAt`/`unsafe.Pointer` private-field lift was
avoidable and is gone. `haxe_go-vfp.10.5.1` replaced it with shared typed
same-package field/method metadata and constructor-free Type helpers. The
second-pass disposition is
[`serializer-accessors-vfp-10.5.1.md`](serializer-accessors-vfp-10.5.1.md).

## Reflection ownership inventory

Reflection is used only when Go's static operations cannot represent an erased
Haxe contract safely. It never grants access to package-private fields.

| Boundary | Exact responsibility | Why a typed ordinary operation is insufficient | Containment |
| --- | --- | --- | --- |
| `runtime/hxrt/string.go` | typed-nil recognition for erased values | A non-nil Go interface can contain a nil pointer, slice, map, function, or channel. | `isAnyNil`; selected with shared string/equality support |
| `runtime/hxrt/equality.go` | reference identity for non-comparable erased values | Go interface `==` panics for maps/slices/functions, while Haxe requires identity semantics. | `referenceEqual`; exact per-file debt ceiling and checkptr fixture |
| `runtime/hxrt/map_object.go` | stable identity keys for `ObjectMap` | Erased reference-shaped keys cannot be put directly into an ordinary comparable Go key. | `objectMapIdentityOf`; rejects unsupported non-comparable values |
| `runtime/hxrt/enum_value.go` | recognize the generated enum carrier | Erasure hides the nominal Haxe enum type from generic collection algorithms. | structural `tag`/`params` predicate only; algorithms stay staged |
| `runtime/hxrt/json.go` | active-path identity for maps, slices, and Array carriers | Cycle detection needs reference identity before recursive JSON normalization. | local `jsonVisit` set; repeated aliases outside the active path remain valid |
| `runtime/hxrt/reflect.go` | exported Go field/method fallback, safe assignment/calls, comparison, copying, and classification | Public Haxe `Reflect` accepts `Dynamic`; concrete host shapes are unknown at compile time. | staged `Reflect` owns policy; generated members use typed same-package adapters; no `unsafe` or unexported-field lift |
| `runtime/hxrt/template.go` | erased slice/object/function representation for staged `haxe.Template` | Template inputs and callbacks are dynamic by contract. | three narrow helpers; parsing, lookup order, macros, iteration, errors, and rendering remain staged |
| generated Type/Reflect adapters | classify or invoke an already-selected erased value and perform deep equality where the portable contract requires it | Some reachable `Dynamic`, function, and interface carriers have no one static Go type. | feature-selected closed-world emitters registered in `docs/compiler-stdlib-intrinsics.json`; typed per-class field/method facts remain separate |
| explicit `go.*` channel fallback helpers | reflect on a native generic channel and use `reflect.Select` for non-blocking operations | An erased native channel has no statically available element type at that boundary. | explicit Go-native API only; typed fast paths remain preferred and the debt is independently measured |

`test/compiler_debt_policy.json` gives every checked-in `hxrt` and generated
example reflection file an exact import and selector ceiling. A new file or one
extra selector fails `npm run test:compiler-debt`; the policy cannot silently
turn one reviewed island into a general runtime mechanism.

## Serializer and reflection behavior evidence

The serializer is ordinary staged Haxe. It owns tokens, recursion, cache order,
resolvers, construction sequencing, custom hooks, and errors. Native
`runtime/hxrt/serialization.go` owns only bounded `strconv.ParseFloat`.

| Acceptance shape | Evidence |
| --- | --- |
| `null` and typed nil behavior | `serializer_wire_contract`; `reflect_extended_contract`; direct `hxrt` nil checks |
| cycles and repeated references | `serializer_cache_reference_contract`; `serializer_reference_stress_contract`; `serializer_boundary_matrix_contract`; direct cyclic/repeated JSON tests |
| interface-typed values | `serializer_boundary_matrix_contract` round-trips an interface-typed field and calls the restored implementation |
| enums and class/enum references | `serializer_class_enum_contract`; `serializer_extended_tokens_contract`; resolver contracts |
| erased generics | `serializer_boundary_matrix_contract` checks 64 deterministic `GenericBox<Int>` round trips and a cached generic self-cycle |
| errors | malformed token, truncated string, and unknown class in `serializer_boundary_matrix_contract`; resolver-null paths; `serializeException` in `serializer_extended_tokens_contract` |
| private/inherited fields and custom hooks | `serializer_typed_accessor_contract` across three class levels, including `@:transient`, virtual dispatch, hook calls, and integral-token-to-`Float` assignment |
| public `Reflect` and `Type` behavior | `reflect_compare`, `reflect_field_ops`, `reflect_extended_contract`, `type_reflection_contract`, `type_reflection_extended_contract`, and direct `runtime/hxrt/reflect_test.go` |

The 64-value loop is a bounded deterministic property test rather than a random
fuzzer. It is reproducible in CI and complements the token/reference stress
fixtures without pretending to cover hostile or unbounded input.

## Toolchain and checkptr evidence

`docs/toolchain-policy.json` admits Go `1.25.12` and `1.26.5`. CI Harness run
[`30256931224`](https://github.com/fullofcaffeine/reflaxe.go/actions/runs/30256931224)
passed the release-blocking Go tooling jobs on both versions for commit
`4c4a2c1a65e2cc92101b5197ec8af026a435469a`. Each job ran seven scopes through:

1. `go test -race`;
2. `go test -gcflags=all=-d=checkptr=2`;
3. `go vet -stdmethods=false`; and
4. Staticcheck `SA*`.

That commit contains the final production unsafe/reflection/serializer code
from child `haxe_go-vfp.10.5.1`. This parent additionally makes the real PTY
behavior contract checkptr-instrumented and adds the interface/generic
serializer matrix. Exact-commit CI for this parent is recorded in the Bead
closeout after push.

Current local validation:

- `npm test`: 301/301 snapshots and all governance stages passed;
- `npm run test:semantic-diff`: 142/142 cases passed;
- `npm run test:stdlib-sweep:go-test`: 55/55 modules passed;
- `npm run security:go-tooling`: all 28 local
  race/checkptr/vet/Staticcheck gates passed;
- `npm run test:release-contracts`: passed; and
- `npm run test:compiler-debt`: passed with `go_unsafe=2`,
  `go_reflection=716`, and `compiler_shim=8`.

## Scope decisions and follow-ups

- No universal semantic IR is justified by this audit. Each retained boundary
  has one representation-specific responsibility and already fits the
  AST-first builder/lowering/pass/printer architecture.
- Arbitrary external package-private field access remains unsupported. The
  compiler generates typed facts only for its own reachable generated types.
- Two unrelated compiler-lowering defects found while drafting the new
  serializer fixture are tracked separately:
  - `haxe_go-vfp.8.12`: erased generic field in direct string concatenation;
  - `haxe_go-vfp.8.13`: exhaustive `try`/`catch` returns emitted without a Go
    return proof.
- `IntMapSnapshot`, `StringMapSnapshot`, and `ObjectMapSnapshot` have no current
  staged Serializer call site; Serializer uses the typed Keys/Get paths. Their
  source comments still contain migration-era “until Serializer moves”
  wording. Correcting those copied comments alone changes nearly 700 committed
  generated files, while removing or retaining the unused exported helpers
  needs a compatibility decision. That wording and disposition are explicitly
  deferred to the existing runtime-slicing task `haxe_go-vfp.10.6`; it does not
  describe an active unsafe, reflection, or serializer-policy boundary.
- Random/adversarial serializer fuzzing can be added if the input trust model
  expands. It is not needed to claim the current deterministic portable
  contract.

## Independent second pass

A fresh-context `gpt-5.6-sol` reviewer at `xhigh` challenged the source,
policy, tests, and audit rather than trusting this document. The first pass
found one material stale statement in `docs/compiler-debt-ratchet.md` that
still described the retired serializer unsafe lift, plus the low-risk
`*MapSnapshot` migration-era comments. The debt document is corrected; the map
comment/removal decision is explicitly and accurately deferred to
`haxe_go-vfp.10.6`.

After those dispositions and an append-only tracker correction, the reviewer
returned **PASS with no closure-blocking findings**. The review exchange was
provided through the agent mailbox; the host did not expose a standalone
transcript file.

## Closure checklist

- [x] Avoidable serializer unsafe access removed.
- [x] Sole retained unsafe operation has an owner, invariant, justification,
  exact ratchet, and real checkptr execution.
- [x] Remaining reflection islands are named, selective, safe, and documented.
- [x] Nil, cycles, interfaces, enums, generics, and applicable error paths have
  direct round-trip or failure evidence.
- [x] Supported Go lines passed race/checkptr/vet/Staticcheck for the final
  production boundary implementation.
- [x] Current parent change passes full local validation.
- [x] Independent xhigh review passes.
- [ ] Exact-commit remote CI passes before the Bead is closed.
