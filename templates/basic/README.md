# reflaxe.go basic template

This is a minimal starter layout for a **Haxe -> Go** project using `reflaxe.go`.

From the `reflaxe.go` repository, scaffold this template with:

```bash
npm run dev:new-project -- ./my_haxe_go_app
```

## Setup

From this folder:

```bash
npm install
npm run setup
```

What `setup` does:

- creates a local lix scope if missing
- installs `reflaxe.go` (default source: `github:fullofcaffeine/reflaxe.go`)
- downloads and pins Haxe 4.3.7 in scope

Compile behavior:

- plain `haxe compile*.hxml` builds a Go binary by default after codegen
- pass `-D go_no_build` for codegen-only workflows

Override install source (for local compiler development):

```bash
REFLAXE_GO_SOURCE="path:/absolute/path/to/reflaxe.go" npm run setup:reflaxe-go
```

## Build and run

Start the normal edit-and-restart loop:

```bash
npm run dev
```

The last working program stays alive when a new edit does not compile. Press
`Ctrl+C` to stop the watcher and its program. Use `npm run hx:run` below for a
single direct build and run.

Portable profile:

```bash
npm run hx:run
```

Metal profile:

```bash
npm run hx:run:metal
```

Build binaries:

```bash
npm run hx:build
npm run hx:build:metal
```

Output locations:

- `out`, `out_metal` for generated Go modules
- `bin/hx_app*` for compiled binaries
