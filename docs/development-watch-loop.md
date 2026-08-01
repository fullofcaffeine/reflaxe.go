# Development watch loop

This page explains the one-command edit, rebuild, and restart loop for a Haxe
project that targets Go.

## Start here

From this repository, run:

```bash
npm run dev -- --project examples/tui_todo --profile portable
```

From a project created with `npm run dev:new-project`, run:

```bash
npm run dev
```

The command performs one real Haxe-to-Go build, strictly builds the generated
Go program, starts it, and then watches the inputs named by the selected HXML
file. Save a `.hx` file and it builds again. Press `Ctrl+C` once to stop the
watcher and every program or compiler process it owns.

An HXML file is the small text build configuration passed to Haxe. It names the
main class, source directories, target library, defines, and generated-output
directory, so the watcher reuses it instead of asking you to repeat that setup.

## What happens when a build fails?

The last working program keeps running. The watcher prints the compiler error
and continues watching. After the next successful build, it stops the old
program and starts the newly built one. This avoids replacing a useful local
app with half-generated or uncompilable output.

In other words, a file save is only a request to try a build. A successful
Haxe compile and Go build are what authorize a restart.

## What is watched?

The watcher starts from the same HXML selected by `dev:hx`. It follows nested
HXML files, direct Haxe classpaths, project `haxelib.json`/`.haxerc` files, and
declared HXML resources. Add macro inputs or other project configuration with a
repeatable override:

```bash
npm run dev -- --project ./my_app --watch-dir config --watch-dir schemas
```

Several quick file events are combined into one build. The defaults are a
120 ms quiet period and an 80 ms portable polling interval. They can be tuned:

```bash
npm run dev -- --project ./my_app --debounce-ms 200 --poll-ms 100
```

## Direct builds and the optional compiler server

To make the edit loop fast, the default watcher owns one Haxe compilation
server for its lifetime:

```bash
npm run dev -- --project ./my_app
```

The server keeps Haxe's parser and typer cache in memory between edits. It is an
accelerator, not a separate compiler mode. If it cannot start, the watcher says
so and falls back to direct builds. Stopping the watcher also stops its server.

Turn the accelerator off when you want every edit to start a fresh compiler:

```bash
npm run dev -- --project ./my_app --server off
```

That direct build remains the correctness baseline. Ordinary `dev:hx`, CI, and
releases use it, and the watch-loop tests retain a cold direct lane.

Use a direct one-shot build when diagnosing a cache-sensitive problem:

```bash
npm run dev -- --project ./my_app --once
```

## Other actions

`run` is the default and owns a restartable program. The remaining actions
re-run after changes and exit after each successful check:

```bash
npm run dev -- --project ./my_app --action test
npm run dev -- --project ./my_app --action vet
npm run dev -- --project ./my_app --action build --binary bin/my_app
```

You can use the same `--hxml`, `--profile`, `--out`, `--define`, `--haxe-bin`,
and `--go-bin` options supported by `dev:hx`. `portable` remains the normal
default product path; selecting `metal` only selects the existing compatibility
policy preset.

## Supported hosts and limits

The managed process-group cleanup contract is tested on POSIX hosts and is
intended for Linux and macOS development. Windows continues to have compile-only
project evidence; this task does not widen that release or runtime claim.

Polling intentionally avoids another package dependency and works on both
Linux and macOS. It observes project-owned files and explicit extra roots. An
undeclared file read by a custom macro cannot be discovered automatically, so
add its directory with `--watch-dir`.

The measured baseline and server comparison for this implementation are
recorded with tracker item `haxe_go-vfp.12.8`.

### Measured development cycle

The tracer used the basic template on an Apple M2 Pro with Haxe 4.3.7, Go
1.25.6, and Darwin arm64. Go 1.25.6 is a local development toolchain, not one
of the exact release-admitted Go patch versions, so these numbers are DevEx
measurements rather than release evidence.

| Build path | First build | Rebuild after an input event |
| --- | ---: | ---: |
| Direct compiler | 3.63 s | 3.06 s |
| Managed Haxe server | 3.60 s | 1.01 s |

The server-backed edit was about 67% faster in this small sample. The practical
regression budget is intentionally soft: over five local repetitions, the
server-backed median edit must remain faster than the direct median. A result
more than 15% slower than the last recorded server median should be investigated
and remeasured on an idle host, not automatically turned into a release failure.
Performance numbers describe this sample project and machine, not a universal
speed promise.
