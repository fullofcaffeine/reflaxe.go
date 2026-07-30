<!-- generated; edit compatibility-support-source.json and run npm run compatibility:generate -->
# Compatibility and Support Matrix

Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads on the admitted toolchain, platform, and operation/member surface.

This is the human rendering of
[`compatibility-support-manifest.json`](compatibility-support-manifest.json).
The JSON manifest is authoritative. A module-level green result is evidence inventory,
not blanket admission of every member or error path.

Default disposition: Any unlisted module operation, member, platform, architecture, native API, error path, or trust model is excluded from the admitted release scope.

## Evidence states

| State | Release eligible | Meaning |
| --- | --- | --- |
| `semantic-diff-supported` | yes | Named deterministic behavior is compared between Haxe --interp and generated Go, which also compiles, passes go test, and runs. |
| `compile-go-test-run-supported` | yes | The named target-sensitive behavior compiles, passes go test, runs on Go, and matches an explicit expected contract without claiming interpreter parity. |
| `compile-only` | no | The named surface compiles or cross-builds, but runtime and semantic support are not admitted. |
| `experimental` | no | The named surface is implemented and tested in some lanes but remains outside the admitted release contract. |
| `compatibility-only` | no | The input or behavior is retained for compatibility and does not define a separate semantic product or release admission. |
| `excluded` | no | The named surface is deliberately unsupported by the admitted release scope. |

## Toolchains

| Role | Admitted value |
| --- | --- |
| Haxe compiler | `4.3.7` |
| Generated Go language floor | `1.22` |
| Go build lines | `1.25, 1.26`; latest patch required |
| Node tooling line | `24` |

## Platforms and architectures

| ID | OS | Architecture | State | Release admitted | Qualification |
| --- | --- | --- | --- | --- | --- |
| `linux-amd64` | `linux` | `amd64` | `compile-go-test-run-supported` | yes | Canonical beta evidence platform; release evidence must record the exact hosted runner image and resolved toolchain patches. |
| `darwin-runner-default` | `darwin` | `runner-default-unfrozen` | `experimental` | no | The quality workflow runs macos-latest, but the architecture and image are not frozen as a release compatibility promise. |
| `linux-arm64` | `linux` | `arm64` | `compile-only` | no | Representative examples cross-build; no release runtime, race, or security matrix is admitted. |
| `darwin-arm64` | `darwin` | `arm64` | `compile-only` | no | Representative examples cross-build; the moving macos-latest lane is tracked separately. |
| `windows-amd64` | `windows` | `amd64` | `compile-only` | no | Representative examples cross-build; Windows runtime behavior is not in the admitted beta matrix. |
| `other-os-architecture` | `other` | `other` | `excluded` | no | No implicit operating-system or architecture support is inferred. |

## Compatibility presets

| Preset | Selector | State | Release admitted | Qualification |
| --- | --- | --- | --- | --- |
| `portable` | `-D reflaxe_go_profile=portable` | `semantic-diff-supported` | yes | Default product path. Only operations and members explicitly admitted below are part of the beta claim. |
| `metal` | `-D reflaxe_go_profile=metal` | `compatibility-only` | no | Supported compatibility policy preset, not a second semantic product and not an independent native-authority claim. |

## Operation/member admission and native-surface inventory

### `portable-language-core`

Named Haxe language operations

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Haxe if control flow` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Deterministic typed control flow within the snapshot and semantic-diff corpus. |
| `Haxe switch control flow` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Deterministic typed control flow within the snapshot and semantic-diff corpus. |
| `Haxe numeric for control flow` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Deterministic typed control flow within the snapshot and semantic-diff corpus. |
| `Haxe while control flow` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Deterministic typed control flow within the snapshot and semantic-diff corpus. |
| `Haxe try/catch control flow` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Deterministic typed control flow within the snapshot and semantic-diff corpus. |
| `Haxe class construction and instance dispatch` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Covered typed construction, dispatch, pattern matching, capture, and function-value behavior only. |
| `Haxe interface dispatch` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Covered typed construction, dispatch, pattern matching, capture, and function-value behavior only. |
| `Haxe enum construction and pattern matching` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Covered typed construction, dispatch, pattern matching, capture, and function-value behavior only. |
| `Haxe closure capture and invocation` | `operation` | `semantic-diff-supported` | yes | `npm:test-semantic-diff`, `npm:test-ci` | Covered typed construction, dispatch, pattern matching, capture, and function-value behavior only. |
| `throw Dynamic -> catch haxe.Exception.message` | `operation` | `semantic-diff-supported` | yes | `semantic:exceptions` | Haxe-thrown exception carrier behavior; foreign Go panic behavior remains a separate native contract. |
### `portable-collections-and-text`

Named core collection and text members

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Array.length` | `member` | `semantic-diff-supported` | yes | `semantic:array-string` | Core indexed and iterator behavior exercised by the named contract. |
| `Array[index] read` | `member` | `semantic-diff-supported` | yes | `semantic:array-string` | Core indexed and iterator behavior exercised by the named contract. |
| `Array.push` | `member` | `semantic-diff-supported` | yes | `semantic:array-string` | Core indexed and iterator behavior exercised by the named contract. |
| `Array.pop` | `member` | `semantic-diff-supported` | yes | `semantic:array-string` | Core indexed and iterator behavior exercised by the named contract. |
| `String.length` | `member` | `semantic-diff-supported` | yes | `semantic:array-string`, `semantic:string-edge` | Named string members including out-of-range charCodeAt behavior. |
| `String.charAt` | `member` | `semantic-diff-supported` | yes | `semantic:array-string`, `semantic:string-edge` | Named string members including out-of-range charCodeAt behavior. |
| `String.charCodeAt` | `member` | `semantic-diff-supported` | yes | `semantic:array-string`, `semantic:string-edge` | Named string members including out-of-range charCodeAt behavior. |
| `String.substring` | `member` | `semantic-diff-supported` | yes | `semantic:array-string`, `semantic:string-edge` | Named string members including out-of-range charCodeAt behavior. |
| `String.fromCharCode` | `member` | `semantic-diff-supported` | yes | `semantic:array-string`, `semantic:string-edge` | Named string members including out-of-range charCodeAt behavior. |
| `StringTools.trim` | `member` | `semantic-diff-supported` | yes | `semantic:stringtools` | Only members exercised by the cross-stdlib contract. |
| `StringTools.startsWith` | `member` | `semantic-diff-supported` | yes | `semantic:stringtools` | Only members exercised by the cross-stdlib contract. |
| `StringTools.replace` | `member` | `semantic-diff-supported` | yes | `semantic:stringtools` | Only members exercised by the cross-stdlib contract. |
| `StringTools.contains` | `member` | `semantic-diff-supported` | yes | `semantic:stringtools` | Only members exercised by the cross-stdlib contract. |
| `StringTools.endsWith` | `member` | `semantic-diff-supported` | yes | `semantic:stringtools` | Only members exercised by the cross-stdlib contract. |
### `portable-data`

Named data, serialization, and standard collection members

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `haxe.Json.parse` | `member` | `semantic-diff-supported` | yes | `semantic:json` | Deterministic parse/stringify shapes in the named JSON contract. |
| `haxe.Json.stringify` | `member` | `semantic-diff-supported` | yes | `semantic:json` | Deterministic parse/stringify shapes in the named JSON contract. |
| `haxe.io.Bytes.ofString` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.io.Bytes.blit` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.io.Bytes.fill` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.io.Bytes.sub` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.io.Bytes.compare` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.io.Bytes.toString` | `member` | `semantic-diff-supported` | yes | `semantic:bytes` | Named byte-buffer operations and bounds behavior. |
| `haxe.ds.StringMap.set` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.StringMap.get` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.StringMap.exists` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.StringMap.remove` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.IntMap.set` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.IntMap.get` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.IntMap.exists` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.IntMap.remove` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.ObjectMap.set` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.ObjectMap.get` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.ObjectMap.exists` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.ObjectMap.remove` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.EnumValueMap.set` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.EnumValueMap.get` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.EnumValueMap.exists` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.EnumValueMap.remove` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.add` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.push` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.length` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.first` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.last` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `haxe.ds.List.pop` | `member` | `semantic-diff-supported` | yes | `semantic:collections` | Core operations for the map/list shapes exercised by the contract. |
| `Date.fromString` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Only operations exercised by option_date_path and linked direct contracts. |
| `Date.getFullYear` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Only operations exercised by option_date_path and linked direct contracts. |
| `Date.getMonth` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Only operations exercised by option_date_path and linked direct contracts. |
| `Date.getDate` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Only operations exercised by option_date_path and linked direct contracts. |
| `Date.getHours` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Only operations exercised by option_date_path and linked direct contracts. |
| `haxe.ds.Option Some/None pattern matching` | `operation` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Some and None construction plus exhaustive switch matching exercised by option_date_path. |
| `haxe.io.Path.join` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Path joining and parsed dir/file/ext fields exercised by option_date_path. |
| `haxe.io.Path.dir` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Path joining and parsed dir/file/ext fields exercised by option_date_path. |
| `haxe.io.Path.file` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Path joining and parsed dir/file/ext fields exercised by option_date_path. |
| `haxe.io.Path.ext` | `member` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Path joining and parsed dir/file/ext fields exercised by option_date_path. |
| `new haxe.io.Path` | `operation` | `semantic-diff-supported` | yes | `semantic:date-option-path` | Construction from the literal POSIX path shape exercised by option_date_path. |
### `portable-reflection`

Complete Haxe 4.3.7 Reflect API and closed-world Type metadata over supported generated and native carriers

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Type.getClass` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getSuperClass` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getClassName` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getClassFields` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getInstanceFields` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getEnum` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getEnumConstructs` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.createEmptyInstance` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.allEnums` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.enumConstructor` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.typeof` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.enumParameters` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Type.getEnumName` | `member` | `semantic-diff-supported` | yes | `semantic:type-reflection` | Closed-world public declarations and the exact Type members exercised by the extended contract. |
| `Reflect.callMethod` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.compare` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.compareMethods` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.copy` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.deleteField` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.field` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.fields` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.getProperty` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.hasField` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.isEnumValue` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.isFunction` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.isObject` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.makeVarArgs` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.setField` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
| `Reflect.setProperty` | `member` | `semantic-diff-supported` | yes | `semantic:reflect-compare`, `semantic:reflect-extended`, `semantic:reflect-fields` | All 15 public Reflect entrypoints are present. Admission is evidence-scoped to generated field/property and enum cases, Haxe anonymous objects, supported maps/exported Go fields and methods, covered numeric/string comparisons, and the exercised generated function carriers; lookup precedence is RTTI, native field, generated field, generated method, then native method. |
### `portable-root-sys`

Interactive standard-input character contract

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Sys.getChar` | `member` | `compile-go-test-run-supported` | yes | `contract:sys-get-char-terminal` | On admitted linux-amd64 terminals, reads one byte without newline, suppresses host echo, restores terminal state, preserves redirected EOF, and performs requested staged echo exactly once. |
### `portable-file-io`

Narrow file-content contract

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `sys.io.File.getContent` | `member` | `semantic-diff-supported` | yes | `semantic:file-content`, `semantic:file-errors` | Text content reads plus missing, directory, and permission failure categories on the admitted platform. |
| `sys.io.File.saveContent` | `member` | `semantic-diff-supported` | yes | `semantic:file-content`, `semantic:file-errors` | Text content writes plus directory and permission failure categories on the admitted platform. |
### `portable-process`

Narrow child-process lifecycle contract

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `sys.io.Process.stdout.readLine` | `member` | `semantic-diff-supported` | yes | `semantic:process`, `semantic:process-errors` | Line, empty-line, long-line, and EOF separation for admitted local commands. |
| `sys.io.Process.stdin.writeString` | `member` | `semantic-diff-supported` | yes | `semantic:process-errors` | Named pipe write/read, positive pid, and close behavior from the error-semantics contract. |
| `sys.io.Process.stdin.close` | `member` | `semantic-diff-supported` | yes | `semantic:process-errors` | Named pipe write/read, positive pid, and close behavior from the error-semantics contract. |
| `sys.io.Process.stderr.readLine` | `member` | `semantic-diff-supported` | yes | `semantic:process-errors` | Named pipe write/read, positive pid, and close behavior from the error-semantics contract. |
| `sys.io.Process.getPid` | `member` | `semantic-diff-supported` | yes | `semantic:process-errors` | Named pipe write/read, positive pid, and close behavior from the error-semantics contract. |
| `sys.io.Process.close` | `member` | `semantic-diff-supported` | yes | `semantic:process-errors` | Named pipe write/read, positive pid, and close behavior from the error-semantics contract. |
| `sys.io.Process.exitCode` | `member` | `semantic-diff-supported` | yes | `semantic:process`, `semantic:process-errors` | Zero and nonzero exit state remains distinct from EOF and launch failure. |
| `new sys.io.Process(..., detached=true)` | `operation` | `excluded` | no | `semantic:process-errors` | Detached process construction intentionally throws on this target. |
### `portable-filesystem`

Named filesystem operations

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `sys.FileSystem.exists` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.rename` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.stat` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.fullPath` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.absolutePath` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.createDirectory` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.isDirectory` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.readDirectory` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.deleteFile` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
| `sys.FileSystem.deleteDirectory` | `member` | `semantic-diff-supported` | yes | `semantic:filesystem` | Complete Haxe 4.3.7 surface over a local temporary tree, including metadata size, canonical existing paths, and absolute paths that need not exist. |
### `portable-concurrency`

Admitted portable synchronization, foreground-thread, event-loop, and pool operations

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `sys.thread.Lock` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named primitive behavior has interpreter parity; supported Go race evidence covers ownership, visibility, per-waiter condition generations, and missed/duplicate wakeup regressions. |
| `sys.thread.Mutex` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named primitive behavior has interpreter parity; supported Go race evidence covers ownership, visibility, per-waiter condition generations, and missed/duplicate wakeup regressions. |
| `sys.thread.Condition` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named primitive behavior has interpreter parity; supported Go race evidence covers ownership, visibility, per-waiter condition generations, and missed/duplicate wakeup regressions. |
| `sys.thread.Semaphore` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named primitive behavior has interpreter parity; supported Go race evidence covers ownership, visibility, per-waiter condition generations, and missed/duplicate wakeup regressions. |
| `sys.thread.Deque` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named primitive behavior has interpreter parity; supported Go race evidence covers ownership, visibility, per-waiter condition generations, and missed/duplicate wakeup regressions. |
| `sys.thread.Tls` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-primitives`, `runtime:thread-regressions`, `snapshot:detached-thread-lifecycle`, `npm:security-go-tooling`, `contract:concurrency` | Get, set, clear, and main/worker isolation have interpreter parity. Runtime churn proves completed portable workers and compiler-owned detached go.Go.spawn callbacks release logical identity and TLS values on return and native-panic unwind; supported Go race, checkptr, vet, and staticcheck lanes exercise the boundary. |
| `sys.thread.Thread` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-runtime`, `runtime:thread-regressions`, `fixture:thread-pool-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named foreground-thread and event-loop behavior has interpreter parity; deterministic supported-Go stress proves pool admission/shutdown exactly once, failed-worker replacement, cancellation cleanup, wakeup deadline recomputation, and bounded portable identity state. |
| `sys.thread.EventLoop` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-runtime`, `runtime:thread-regressions`, `fixture:thread-pool-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named foreground-thread and event-loop behavior has interpreter parity; deterministic supported-Go stress proves pool admission/shutdown exactly once, failed-worker replacement, cancellation cleanup, wakeup deadline recomputation, and bounded portable identity state. |
| `sys.thread.FixedThreadPool` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-runtime`, `runtime:thread-regressions`, `fixture:thread-pool-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named foreground-thread and event-loop behavior has interpreter parity; deterministic supported-Go stress proves pool admission/shutdown exactly once, failed-worker replacement, cancellation cleanup, wakeup deadline recomputation, and bounded portable identity state. |
| `sys.thread.ElasticThreadPool` | `surface` | `semantic-diff-supported` | yes | `semantic:thread-runtime`, `runtime:thread-regressions`, `fixture:thread-pool-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:concurrency-audit` | Named foreground-thread and event-loop behavior has interpreter parity; deterministic supported-Go stress proves pool admission/shutdown exactly once, failed-worker replacement, cancellation cleanup, wakeup deadline recomputation, and bounded portable identity state. |
### `portable-networking`

Operation-split network support; only explicitly admitted blocking IPv4 TCP client members participate in the beta release

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `new sys.net.Host("IPv4 literal")` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Host.toString (IPv4 literal)` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.connect (blocking IPv4 literal)` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.input.readByte/readBytes (blocking)` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.output.writeByte/writeBytes (blocking)` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.output.flush` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.host` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.peer` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `sys.net.Socket.close` | `operation` | `semantic-diff-supported` | yes | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | On the portable preset, Linux/amd64, Haxe 4.3.7, and Go 1.25.12 or 1.26.5, the admitted core is blocking IPv4 TCP client I/O to application-controlled, pre-resolved numeric endpoints. Deterministic loopback evidence covers connection, source-owned stream reads and writes with partial-progress preservation, flush, local and peer addresses, and close. |
| `haxe.Http` | `surface` | `experimental` | no | `semantic:http`, `runtime:http-regressions`, `snapshot:http-selective-runtime`, `bead:network-audit` | Canonical staged haxe.Http/sys.Http owns public callback and Output policy over typed native resources. Current loopback tests prove bounded multipart reads and final-state behavior for small complete responses, but they do not prove streamed response lifecycle, partial-body preservation, cancellable uploads, or exact request fidelity. |
| `sys.Http` | `surface` | `experimental` | no | `semantic:http`, `runtime:http-regressions`, `snapshot:http-selective-runtime`, `bead:network-audit` | Canonical staged haxe.Http/sys.Http owns public callback and Output policy over typed native resources. Current loopback tests prove bounded multipart reads and final-state behavior for small complete responses, but they do not prove streamed response lifecycle, partial-body preservation, cancellable uploads, or exact request fidelity. |
| `sys.net.Socket.bind` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Loopback fixtures exercise the current server path, but bind currently starts listening and listen does not apply its public backlog argument, so server lifecycle parity is not admitted. |
| `sys.net.Socket.listen` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Loopback fixtures exercise the current server path, but bind currently starts listening and listen does not apply its public backlog argument, so server lifecycle parity is not admitted. |
| `sys.net.Socket.accept` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Loopback fixtures exercise the current server path, but bind currently starts listening and listen does not apply its public backlog argument, so server lifecycle parity is not admitted. |
| `sys.net.Socket.setTimeout` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `sys.net.Socket.setBlocking` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `sys.net.Socket.select` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `sys.net.Socket.waitForRead` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `sys.net.Socket.shutdown` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `sys.net.Socket.setFastSend` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Deterministic fixtures exercise current timeout and select behavior, but connection presence is not real write readiness, exceptional readiness is unresolved, and the remaining control semantics are not yet admitted. |
| `new sys.net.Host("hostname")` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Host currently resolves names eagerly through the Go resolver before a Socket timeout or cancellation policy can apply. Compile and loopback evidence do not admit bounded DNS or reverse lookup behavior. |
| `sys.net.Host.reverse` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Host currently resolves names eagerly through the Go resolver before a Socket timeout or cancellation policy can apply. Compile and loopback evidence do not admit bounded DNS or reverse lookup behavior. |
| `sys.net.Host.localhost` | `operation` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `bead:network-audit` | Host currently resolves names eagerly through the Go resolver before a Socket timeout or cancellation policy can apply. Compile and loopback evidence do not admit bounded DNS or reverse lookup behavior. |
| `sys.net.UdpSocket` | `surface` | `experimental` | no | `semantic:socket`, `runtime:socket-regressions`, `runtime:socket-cross-build`, `bead:network-audit` | Loopback and compile fixtures exercise current IPv4 UDP behavior, but zero-byte datagram sender identity and complete staged-Haxe parity remain unresolved. |
| `sys.ssl.Socket` | `surface` | `experimental` | no | `snapshot:socket-tls`, `snapshot:socket-sni`, `runtime:socket-regressions`, `runtime:socket-cross-build`, `bead:network-audit` | Linux loopback tests and Linux/Windows compile checks exercise current TLS code. They do not prove default logical-host verification and SNI, TLS control semantics, public certificate stores, hostile peers, or runtime behavior on Windows or Darwin. |
### `go-native`

Typed target-native authoring surfaces

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `go.Slice<T>` | `surface` | `experimental` | no | `snapshot:native-collections`, `bead:native-matrix` | Typed snapshots and runtime smokes exist, but the governed native capability matrix is not complete. |
| `go.Map<K,V>` | `surface` | `experimental` | no | `snapshot:native-collections`, `bead:native-matrix` | Typed snapshots and runtime smokes exist, but the governed native capability matrix is not complete. |
| `go.Result<T>` | `surface` | `experimental` | no | `snapshot:native-collections`, `bead:native-matrix` | Typed snapshots and runtime smokes exist, but the governed native capability matrix is not complete. |
| `go.Go` | `surface` | `experimental` | no | `snapshot:native-concurrency`, `fixture:channel-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:native-matrix` | Go-specific goroutine and channel lifecycle behavior, including closed-versus-empty receive and native close/send panics, is tested; true native select and the complete public native capability contract remain open. |
| `go.Chan<T>` | `surface` | `experimental` | no | `snapshot:native-concurrency`, `fixture:channel-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:native-matrix` | Go-specific goroutine and channel lifecycle behavior, including closed-versus-empty receive and native close/send panics, is tested; true native select and the complete public native capability contract remain open. |
| `go.Select` | `surface` | `experimental` | no | `snapshot:native-concurrency`, `fixture:channel-gate`, `npm:security-go-tooling`, `contract:concurrency`, `bead:native-matrix` | Go-specific goroutine and channel lifecycle behavior, including closed-versus-empty receive and native close/send panics, is tested; true native select and the complete public native capability contract remain open. |
| `@:go.import` | `surface` | `experimental` | no | `snapshot:native-extern`, `bead:native-matrix` | The named typed extern forms work for covered fixtures; broad callbacks, generics, unsafe shapes, and public API stability remain open. |
| `@:go.name` | `surface` | `experimental` | no | `snapshot:native-extern`, `bead:native-matrix` | The named typed extern forms work for covered fixtures; broad callbacks, generics, unsafe shapes, and public API stability remain open. |
| `@:go.receiver` | `surface` | `experimental` | no | `snapshot:native-extern`, `bead:native-matrix` | The named typed extern forms work for covered fixtures; broad callbacks, generics, unsafe shapes, and public API stability remain open. |
| `@:go.valueError` | `surface` | `experimental` | no | `snapshot:native-extern`, `bead:native-matrix` | The named typed extern forms work for covered fixtures; broad callbacks, generics, unsafe shapes, and public API stability remain open. |
### `compiler-input-trust`

Compiler and generator trust boundary

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Compiler use with trusted Haxe source and locked dependencies` | `operation` | `compile-go-test-run-supported` | yes | `npm:test-ci`, `workflow:release`, `contract:output-confinement` | The admitted beta assumes reviewed source and repository-controlled or otherwise trusted dependencies. |
| `Compiler use as a sandbox for untrusted Haxe source` | `operation` | `excluded` | no | `contract:output-confinement`, `policy:output-confinement`, `policy:known-gaps` | Compiler-managed source, runtime, report, metadata, extra-file, and package-staging destinations are confined, but Haxe macros, dependencies, plugins, and explicitly selected child tools execute with host authority; the compiler is not a sandbox for adversarial source. |
### `distribution`

Package and hosted release evidence

| Symbol | Granularity | State | Release admitted | Evidence | Qualification |
| --- | --- | --- | --- | --- | --- |
| `Deterministic local Haxelib package and isolated install` | `operation` | `compile-go-test-run-supported` | no | `npm:test-haxelib-release-install`, `bead:release-assets` | Local deterministic construction and isolated install are gated, but hosted same-SHA assets are a separate release prerequisite. |
| `Published checksummed same-SHA Haxelib release assets` | `operation` | `excluded` | no | `bead:release-assets`, `workflow:release` | The current beta-baseline publication gate is not complete. |

## Portable stdlib module evidence

Module evidence never admits every member. Release admission comes only from the operation/member inventory in surfaces.

| Derived state | Module count |
| --- | ---: |
| `compile-go-test-run-supported` | 11 |
| `excluded` | 2 |
| `semantic-diff-supported` | 162 |

## Trust assumptions

- `trusted-source`: Application Haxe source, build metadata, extern metadata, and dependency source are reviewed or otherwise trusted.
- `locked-tooling`: Repository tooling is installed from the committed npm lock and vendored Reflaxe provenance passes its offline verification.
- `local-system-boundaries`: Admitted file and process operations run on the canonical platform against application-controlled paths and child commands.
- `application-controlled-tcp-boundary`: Only the named blocking IPv4 TCP client core is admitted, and only for application-controlled, pre-resolved numeric endpoints; HTTP, DNS, server sockets, socket controls, UDP, TLS, public networks, and hostile-peer behavior remain excluded.

## Known blockers

| Bead | Scope | Blocks |
| --- | --- | --- |
| `haxe_go-vfp.10.8` | Portable HTTP request fidelity, streamed responses, and cancellable uploads | portable-http |
| `haxe_go-vfp.10.9` | Advanced portable socket, DNS, UDP, readiness, server, and TLS semantics | portable-socket-advanced |
| `haxe_go-vfp.9.1` | Governed Go-native capability matrix | go-native |
| `haxe_go-vfp.4.8` | Checksummed same-SHA hosted release assets and provenance | published beta-baseline artifact |
| `haxe_go-vfp.12.7` | Separate stable-1.x admission decision | any stable 1.x claim |
