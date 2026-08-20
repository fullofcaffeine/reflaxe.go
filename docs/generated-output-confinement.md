# Generated-output confinement

## What this contract protects

`reflaxe.go` treats every compiler-selected output name as an untrusted path
input, even when the current producer normally emits a fixed or sanitized name.
Before a compiler-owned file is written, copied, recorded, or selected for stale
file deletion, the path must pass the typed boundary in
`src/reflaxe/go/compiler/GoGeneratedOutputBoundary.hx`.

The configured `go_output` directory is the ownership root in standalone mode.
Compiler-managed Go source, `go.mod`, reports, copied `hxrt` files, Reflaxe
extra files, and `_GeneratedFiles.json` must remain below its canonical
filesystem location. Existing-module mode instead uses the caller module root
as its confinement root and a digest-backed package ownership record.

This contract exists because lexical joining is not confinement. A relative
name can contain traversal, a host can interpret separators differently, and a
path that looks nested can traverse a symbolic link to another tree.

## How the write flow works

For standalone output, the compiler establishes the boundary before Reflaxe
starts generation:

1. Resolve `go_output` to an absolute, canonical directory. Symlinked ancestor
   directories are resolved, but the configured output directory itself may not
   be a symbolic link.
2. Validate `_GeneratedFiles.json` and every old `filesGenerated` entry before
   Reflaxe can delete a stale file.
3. Validate every Reflaxe extra-file key before any ordinary generated file is
   written.
4. Route generated modules, `go.mod`, reports, and runtime copies through the
   same boundary immediately before `OutputManager.saveFile`.
5. Re-check the destination after creating any missing parent directory.

The legacy `GoOutputIterator` helpers use the same boundary for direct writes.
They are retained for compatibility, but they are not a second output policy.

Existing-module output does not invoke Reflaxe's generic writer or stale-file
deletion. It collects all artifacts in a typed plan, validates prior ownership
digests, then installs them through `GoExistingModuleOutputTransaction`. The
package-local `.reflaxe-go-owned.json` record is the commit marker. The legacy
`_GeneratedFiles.json` file has no replacement, cleanup, or package-ownership
authority in this mode.

Haxelib staging is a separate process with a separate configured root, so
`Run.hx` keeps its package-specific typed boundary. It validates POSIX archive
member names, resolves every destination below the canonical package root, and
performs a complete symlink/confinement preflight before recursive `--clean`
deletion. Source-tree symlinks are rejected rather than copied into a release
package.

| Writer or mutation | Path authority | Enforcement |
| --- | --- | --- |
| Generated `.go` modules | `GoCompiler` relative file names | `GoGeneratedOutputBoundary.saveFile` |
| `go.mod` and optional JSON/Markdown reports | Fixed compiler names | Same boundary; fixed names are not exempt |
| Selective or full `hxrt` copy | Runtime feature plan and directory entries | `copyManagedFile` through the same boundary |
| Reflaxe extra files | Compiler/plugin-provided keys | Whole set preflighted before generation |
| Stale output deletion | `_GeneratedFiles.json` entries | Metadata path and every entry preflighted before generation |
| Existing-module replacement and stale cleanup | `.reflaxe-go-owned.json` paths and SHA-256 digests | Complete artifact plan plus manifest-last output transaction |
| Legacy iterator output | `GoGeneratedFile.relativePath` and runtime entries | Direct boundary writer/copy methods |
| Haxelib package writes and clean deletion | Declared package mappings | `PackagePathTools.confinedDestination` and delete-tree preflight |

The vendored Reflaxe tree remains pinned to its independently verified supplier
snapshot. Its generic `OutputManager` is intentionally not presented as a safe
cross-target boundary: the Go target wraps and preflights every route it uses.
Changing the vendor bytes would require a new provenance layer and would not
improve other targets unless they adopted the same contract.

## Rejected path forms

Generated paths use one canonical POSIX-relative spelling. The boundary rejects:

- empty paths, absolute paths, and leading separators;
- `.` or `..` segments, repeated separators, and normalization aliases;
- Windows drive-qualified, drive-relative, UNC, and device namespace forms;
- backslashes, colons/alternate data streams, control characters, trailing dots
  or spaces, and Windows device basenames such as `NUL` or `COM1`;
- an existing directory where a generated file is expected; and
- every symbolic-link component below the configured root, including a broken
  final symlink and a link whose target would otherwise remain inside the root.

Rejecting all descendant symlinks gives the output tree one unambiguous owner.
It also avoids treating a currently contained link as permanently safe when its
target can later change.

Compiler confinement errors have stable `GO-OUTPUT-PATH-*` codes. They do not
echo the rejected value, canonical root, source checkout, or resolved external
destination, so diagnostics can be retained without publishing machine-local
paths.

## Explicit outputs that are not managed files

`go_build_output=<path>` is an explicit caller-authorized `go build -o` sink.
It is passed to the Go tool from inside `go_output`, may intentionally name a
location outside the generated source tree, and is not recorded or cleaned as a
compiler-managed file. Failure diagnostics redact absolute values. Likewise,
`go_cmd` selects an explicitly authorized child executable whose behavior is
outside a file-confinement guarantee.

Use `go_no_build` or `go_codegen_only` when another trusted build system owns
binary placement. Do not infer permission for compiler-managed source or copied
runtime files from either build control; those files remain confined.

## Trust boundary and non-goals

Output confinement is defense in depth for trusted builds. It does **not** make
the Haxe compiler a sandbox for adversarial source:

- Haxe macros and build macros execute with compiler-process authority;
- dependencies and compiler plugins can execute code;
- explicitly selected child tools can perform arbitrary host operations; and
- a concurrent hostile local process that can replace output components between
  validation and the filesystem write is outside this process-local guard.

Therefore the compatibility manifest continues to exclude “compiler use as a
sandbox for untrusted Haxe source.” The exclusion is now permanent and
explanatory rather than blocked on output confinement. Builds should use
reviewed source and dependencies and an output directory owned by the build
process. A real adversarial-source service needs an operating-system sandbox,
separate credentials, resource limits, and isolated storage in addition to this
compiler invariant.

## Verification

Run the focused contract with:

```bash
npm run test:output-confinement
```

The contract covers traversal, absolute, drive, UNC, mixed-separator, device,
existing-link, contained-link, and broken-link paths. It also runs the real Go
compiler against an `hxrt` output symlink, poisons Reflaxe managed-file metadata
with a traversal deletion, checks that external sentinels survive, and exercises
package-clean symlink rejection through the Haxelib package tests.

The focused contract is part of `npm test`, `npm run test:changed`, the full CI
driver, and the release-contract driver.
