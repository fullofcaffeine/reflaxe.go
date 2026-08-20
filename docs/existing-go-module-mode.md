# Existing Go module mode

haxe.go currently creates a complete Go module in `go_output`. Many projects
already have a `go.mod`, native Go files, build tags, and release commands.
Existing Go module mode lets generated Haxe code join that project safely.

This mode has one primary rule. The caller owns the Go module, and the compiler
owns only the generated files that its ownership record names.

## Terms

A caller-owned module is a directory that already contains `go.mod`. The
project build or a developer creates this file before Haxe compilation.

A project manifest is a JSON file with the inputs for one generated package.
The compiler parses this file into typed Haxe values before it writes files.

An ownership record lists each compiler-owned file and its content digest.
The digest detects a local change before the compiler replaces or removes it.

## Before and after

Standalone mode currently produces this structure:

```text
go_output/
  go.mod                 compiler-owned
  main.go                compiler-owned
  hxrt/                  compiler-owned
```

Existing-module mode joins a structure like this one:

```text
project/                 moduleRoot, caller-owned
  go.mod                 caller-owned
  go.sum                 caller-owned
  cmd/tool/              packageDir
    main.go              caller-owned bridge
    haxego_generated_main.go
                          compiler-owned
  internal/haxe_hxrt/    runtimeDir, compiler-owned
```

Existing-module mode never creates, changes, or removes `go.mod` or `go.sum`.
It reads the module path from `go.mod` and derives runtime imports from it.

## Typed project manifest

The define `reflaxe_go_project=<path>` activates existing-module mode. Its
value names one JSON project manifest. Paths inside the manifest are relative
to the manifest file unless this document says otherwise.

```json
{
  "schemaVersion": 1,
  "mode": "existing-module",
  "moduleRoot": ".",
  "packageDir": "cmd/tool",
  "packageName": "main",
  "runtimeDir": "internal/haxe_hxrt",
  "entrypoint": {
    "kind": "caller-bridge",
    "symbol": "RunHaxeMain"
  },
  "build": {
    "kind": "none"
  }
}
```

The parser rejects unknown fields and unsupported schema versions. It converts
the JSON values into these typed concepts:

```haxe
enum GoProjectMode {
    Standalone;
    ExistingModule(project:ExistingGoModuleProject);
}

enum GoEntrypointPolicy {
    CompilerMain;
    CallerBridge(symbol:String);
}

enum GoBuildPolicy {
    NoBuild;
    GoBuild(request:GoBuildRequest);
}
```

`ExistingGoModuleProject` contains `moduleRoot`, `packageDir`, `packageName`,
`runtimeDir`, the entry-point policy, and `build`. Domain logic does not read
raw JSON values.

The existing `go_output` define remains necessary during the migration. In
this mode, its canonical path must equal `moduleRoot/packageDir`. The compiler
uses it as an assertion and not as ownership permission.

The optional `go_module` define becomes a module-path assertion. If present,
it must equal the module path that the compiler reads from `go.mod`.

## Validation before writes

The compiler completes all project validation before its first output write.
It applies these rules:

1. `moduleRoot` must contain a regular `go.mod` file.
2. `packageDir` and `runtimeDir` must stay below the canonical `moduleRoot`.
3. Relative paths must not contain traversal, device, or mixed-separator forms.
4. Symlinks must not move an owned path outside `moduleRoot`.
5. Existing Go files in `packageDir` must declare `packageName`.
6. The selected entry point must not conflict with an existing Go symbol.
7. Existing ownership data must have a supported version and valid digests.
8. Legacy defines must not conflict with the typed project manifest.

A validation error leaves `go.mod`, `go.sum`, native Go files, and prior
generated files unchanged. M03-02 adds byte-exact tests for this rule.

## Generated-file ownership

Existing-module mode writes `packageDir/.reflaxe-go-owned.json`. This versioned
record stores paths relative to `moduleRoot`, content digests, package
identity, runtime location, and the manifest schema version.

The compiler can replace or remove a file only when the record owns that file.
The recorded digest must also match the current bytes. A changed generated file
is a conflict, because the compiler cannot know whether the caller owns the edit.

On the first generation, a destination file must not exist. The `runtimeDir`
must be absent or empty. These rules prevent the compiler from adopting files
that another tool created.

The compiler never uses the standalone `_GeneratedFiles.json` list to remove
files in a mixed-owner module. M03-08 owns traversal, symlink, stale-record,
changed-file, and interrupted-write tests.

## Package and entry point

All generated Go files use `packageName`. The compiler rejects a package
directory that contains another package declaration.

The `compiler-main` policy emits `func main()` and requires `packageName` to be
`main`. Until the digest-backed ownership record lands, the compiler admits
this policy only in an empty package directory or one that contains only files
in the current generated-file inventory.

The `caller-bridge` policy emits one function with the configured symbol. It does
not emit `func main()`. An exported symbol lets another package call the bridge;
an unexported symbol can be used by caller Go files in the same package. A native
Go `main` function can therefore keep ownership of process startup.

M03-03 owns package declarations, both entry-point policies, symbol validation,
and the generated runtime import path.

## Build behavior

The `build` field selects `none` or a typed Go build request. A request contains
a package target, output path, tags, linker flags, `trimpath`, race mode, and
an approved argument list.

```json
{
  "kind": "go-build",
  "packageTarget": "./cmd/tool",
  "output": "dist/tool",
  "tags": ["gms_pure_go"],
  "ldflags": ["-s", "-w", "-X", "main.Build=release candidate"],
  "trimpath": true,
  "race": false,
  "arguments": ["-buildvcs=false"]
}
```

Every field is required. `packageTarget` is `.` or one exact `./` package below
the module root. `output` is a module-relative file path. Tags are sorted and
deduplicated. Linker flags remain ordered tokens; the compiler encodes them
with Go's quoted-argument rules before passing one `-ldflags` process argument.
The output file's parent directory must already exist, as required by `go build`.

Additional arguments are closed to these forms: `-a`, `-v`, `-x`,
`-buildvcs=auto|false|true`, the Go `-buildmode` values, `-mod=readonly|vendor`,
and a positive `-p` value. Fields already modeled by the request cannot be
repeated through `arguments`. The compiler adds `-mod=readonly` when the list
does not select a module mode, which protects caller-owned module files.

The compiler invokes Go from `moduleRoot`. It passes each argument as a process
argument and never constructs an unrestricted shell string. M03-04 owns this
request. It records the effective command and argument array in the generated
`reflaxe_go_build.json`; the report uses `.` for the working directory and does
not contain the machine-local module path.

M03-05 owns the explicit environment allowlist for CGO and cross-compilation.
The project manifest does not inherit arbitrary environment variables.

## Diagnostics and reports

Existing-module diagnostics use stable codes for these error groups:

- invalid project manifest
- path escape or symlink escape
- missing or invalid module file
- package or entry-point conflict
- ownership conflict
- build-policy conflict

Diagnostics identify paths relative to `moduleRoot`. Reports must not contain
machine-local absolute paths. A report names the project mode, package,
runtime import, ownership version, and effective build target.

## Migration

Standalone mode remains the default. It keeps the current `go.mod`, `hxrt`,
entry-point, cleanup, and build behavior.

Existing-module mode starts only when `reflaxe_go_project` names a valid
manifest. Existing users do not need to change their HXML files.

In existing-module mode, legacy defines have these meanings:

- `go_output` asserts the package directory.
- `go_module` asserts the module path.
- `go_no_build` and `go_codegen_only` can select `build.kind=none`.
- `go_cmd` and `go_build_output` are invalid.

The compiler rejects conflicting old and new settings before it writes files.
M03-07 proves that standalone behavior remains compatible.

## Implementation order

1. M03-02 protects `go.mod` and `go.sum` before other existing-module writes.
2. M03-03 adds package selection and the two entry-point policies.
3. M03-04 adds the typed build request and module-root invocation.
4. M03-08 adds ownership, confinement, cleanup, and interruption tests.
5. M03-07 runs the complete standalone compatibility suite.

The first vertical test uses a temporary Go module and a native `main.go`.
Haxe generation emits a caller bridge into that package. The test builds and
runs the module, then compares all caller-owned files with their original bytes.

## Non-goals

This mode does not run `go mod tidy`. It does not edit module dependencies.
It does not add Beads-specific paths, package names, or build rules to haxe.go.
It does not permit one manifest to generate several Go packages.

A later measured requirement can add multi-package generation. The first mode
keeps one generated package and one explicit runtime directory.

## Related docs

- [Generated-output confinement](generated-output-confinement.md)
- [Defines reference](defines-reference.md)
- [Runtime package](hxrt-runtime.md)
- [Compiler target template](compiler-target-template.md)
