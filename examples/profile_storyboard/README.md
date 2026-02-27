# profile_storyboard

Compact profile walkthrough example compiled from one Haxe codebase to `portable` and `metal`.

## Why this example exists

- Gives a small, readable app where profile behavior is visible quickly.
- Demonstrates profile adapter selection via `profile/RuntimeFactory.hx`.
- Shows why two profiles are useful without splitting app/domain logic.

## Profile behavior in this example

- `portable`: portability-first semantic lane.
- `metal`: explicit Go-first lane with strict boundary defaults.
- Both lanes preserve the same storyboard contract; profile differences are mainly code shape and adapter wiring.

## When to choose each profile here

- Choose `portable` when this pattern will be shared across other Haxe targets.
- Choose `metal` when this module is part of a Go-native/perf-focused pipeline and you want stricter metal checks.

## Tradeoffs shown by this example

- You get one shared Haxe codebase and deterministic output in both profiles.
- Generated Go may look similar when portable surfaces dominate; that is expected.
- Profile-specific runtime adapters still make intent explicit in source control and CI.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_metal && go run .)
```

## Generated Go diff inspection

Quick whole-tree diff:

```bash
diff -ru generated/portable generated/metal
```

High-signal files:

- `generated/portable/module_profile_runtimefactory.go`
- `generated/metal/module_profile_runtimefactory.go`
- `generated/portable/main.go`
- `generated/metal/main.go`

## What it renders

- Health block: readiness progress bar, card mix, open load, velocity, ETA.
- Board block: `TODO`, `DOING`, `DONE` lanes with owners/tags.
- Risk block: high-risk open work and release-tagged open count.
- Profile signal line: profile-specific telemetry from runtime adapters.
- Decision line: release recommendation based on computed risk.

## Related docs

- `docs/profiles.md`
- `docs/profile-semantics-guide.md`
- `docs/examples-matrix.md`
