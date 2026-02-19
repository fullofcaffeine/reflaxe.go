# profile_storyboard

Showcase "release command center" example compiled from one Haxe codebase to `portable`, `gopher`, and `metal`.

## Why this exists

- Produces a demo-friendly output that still represents useful project telemetry.
- Shows a shared abstraction with profile overlays via `profile/RuntimeFactory.hx`.
- Makes profile differences obvious without forking core app logic.
- Serves as a high-signal docs/demo artifact for reflaxe.go.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.gopher.hxml
haxe compile.metal.hxml
```

## Run

```bash
(cd out_portable && go run .)
(cd out_gopher && go run .)
(cd out_metal && go run .)
```

## What it shows

- Health block: readiness progress bar, card mix, open load, velocity, ETA.
- Board block: `TODO`, `DOING`, `DONE` lanes with owners/tags.
- Risk block: high-risk open work and release-tagged open count.
- Profile signal line: profile-specific telemetry from runtime adapters.
- Decision line: simple release recommendation based on computed risk.
