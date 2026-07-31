<!-- generated; edit compatibility-support-source.json and run npm run compatibility:generate -->
# Compatibility Release Status

Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads on the admitted toolchain, platform, and operation/member surface.

This status is generated from the same governed source as the machine manifest and
must be used as the compatibility paragraph in release notes.

## Admitted scope

- preset: `portable` (`semantic-diff-supported`);
- platform: `linux/amd64`;
- named operations/members: 171;
- toolchains: exact Haxe version and latest patched supported Go/Node lines from `toolchain-policy.json`;
- trust: reviewed application source, locked tooling, application-controlled local file/process boundaries, and application-controlled numeric IPv4 endpoints, peers, and TLS identity/trust.

No module-level inventory row expands this scope. Unlisted operations and error paths
take the default excluded disposition.

## Not admitted by this release scope

- `portable-process` / `new sys.io.Process(..., detached=true)`: `excluded`.
- `portable-networking` / `haxe.Http/sys.Http request (plain HTTP numeric IPv4)`, `haxe.Http/sys.Http customRequest (plain HTTP numeric IPv4)`, `haxe.Http/sys.Http requestUrl (plain HTTP numeric IPv4)`, `haxe.Http/sys.Http header, parameter, postData, and postBytes APIs (plain HTTP numeric IPv4)`: `semantic-diff-supported`; blockers: haxe_go-vfp.10.8.
- `portable-networking` / `haxe.Http/sys.Http fileTransfer/fileTransfert with request/customRequest (plain HTTP numeric IPv4)`: `semantic-diff-supported`; blockers: haxe_go-vfp.10.8.
- `portable-networking` / `haxe.Http/sys.Http request/customRequest (data URL)`: `compile-go-test-run-supported`; blockers: haxe_go-vfp.10.8.
- `portable-networking` / `haxe.Http.PROXY and customRequest(..., socket)`: `experimental`.
- `portable-networking` / `haxe.Http/sys.Http request/customRequest/requestUrl (HTTPS)`: `experimental`.
- `portable-networking` / `new sys.net.Host("hostname")`: `experimental`.
- `portable-networking` / `sys.net.Host.reverse`: `experimental`.
- `go-native` / `go.Slice<T>`, `go.Map<K,V>`, `go.Result<T>`: `experimental`; blockers: haxe_go-vfp.9.1.
- `go-native` / `go.Go`, `go.Chan<T>`, `go.Select`: `experimental`; blockers: haxe_go-vfp.9.1.
- `go-native` / `@:go.import`, `@:go.name`, `@:go.receiver`, `@:go.valueError`: `experimental`; blockers: haxe_go-vfp.9.1.
- `compiler-input-trust` / `Compiler use as a sandbox for untrusted Haxe source`: `excluded`.
- `distribution` / `Deterministic local Haxelib package and isolated install`: `compile-go-test-run-supported`; blockers: haxe_go-vfp.4.8.
- `distribution` / `Published checksummed same-SHA Haxelib release assets`: `excluded`; blockers: haxe_go-vfp.4.8.
- non-canonical operating-system/architecture combinations and moving runner identities;
- a stable 1.x compatibility claim or published validated beta-baseline artifact.

## Known blockers

- `haxe_go-vfp.10.8`: Portable HTTP request fidelity, streamed responses, and cancellable uploads (blocks portable-http).
- `haxe_go-vfp.9.1`: Governed Go-native capability matrix (blocks go-native).
- `haxe_go-vfp.4.8`: Checksummed same-SHA hosted release assets and provenance (blocks published beta-baseline artifact).
- `haxe_go-vfp.12.7`: Separate stable-1.x admission decision (blocks any stable 1.x claim).

## Machine authority

- `docs/compatibility-support-manifest.json`
- public SemVer boundary: `docs/public-contract.md`
- lifecycle and stable admission: `docs/semver-lifecycle-policy.md`
- verify with `npm run compatibility:verify`
