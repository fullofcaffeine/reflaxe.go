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

Override install source (for local compiler development):

```bash
REFLAXE_GO_SOURCE="path:/absolute/path/to/reflaxe.go" npm run setup:reflaxe-go
```

## Build and run

Portable profile:

```bash
npm run hx:run
```

Go-first profile:

```bash
npm run hx:run:gopher
```

Metal profile:

```bash
npm run hx:run:metal
```

Build binaries:

```bash
npm run hx:build
npm run hx:build:gopher
npm run hx:build:metal
```

Output locations:

- `out`, `out_gopher`, `out_metal` for generated Go modules
- `bin/hx_app*` for compiled binaries
