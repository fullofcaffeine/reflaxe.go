<!-- generated; edit compatibility-support-source.json and run npm run compatibility:generate -->
# Compatibility Release Status

Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads on the admitted toolchain, platform, and operation/member surface.

This status is generated from the same governed source as the machine manifest and
must be used as the compatibility paragraph in release notes.

## Admitted scope

- preset: `portable` (`semantic-diff-supported`);
- platform: `linux/amd64`;
- named operations/members: 124;
- toolchains: exact Haxe version and latest patched supported Go/Node lines from `toolchain-policy.json`;
- trust: reviewed application source, locked tooling, and application-controlled local file/process boundaries.

No module-level inventory row expands this scope. Unlisted operations and error paths
take the default excluded disposition.

## Not admitted by this release scope

- `portable-process` / `new sys.io.Process(..., detached=true)`: `excluded`.
- `portable-networking` / `haxe.Http request callbacks`, `sys.Http request callbacks`: `experimental`; blockers: haxe_go-vfp.10.4.
- `portable-networking` / `sys.net.Socket`, `sys.net.UdpSocket`, `sys.net.Host`, `sys.ssl.Socket`: `experimental`; blockers: haxe_go-vfp.10.4.
- `go-native` / `go.Slice<T>`, `go.Map<K,V>`, `go.Result<T>`: `experimental`; blockers: haxe_go-vfp.9.1.
- `go-native` / `go.Go`, `go.Chan<T>`, `go.Select`: `experimental`; blockers: haxe_go-vfp.9.1.
- `go-native` / `@:go.import`, `@:go.name`, `@:go.receiver`, `@:go.valueError`: `experimental`; blockers: haxe_go-vfp.9.1.
- `compiler-input-trust` / `Compiler use as a sandbox for untrusted Haxe source`: `excluded`.
- `distribution` / `Deterministic local Haxelib package and isolated install`: `compile-go-test-run-supported`; blockers: haxe_go-vfp.4.8.
- `distribution` / `Published checksummed same-SHA Haxelib release assets`: `excluded`; blockers: haxe_go-vfp.4.8.
- non-canonical operating-system/architecture combinations and moving runner identities;
- a stable 1.x compatibility claim or published validated beta-baseline artifact.

## Known blockers

- `haxe_go-vfp.10.4`: Network, HTTP, socket, timeout, cancellation, and cleanup closure (blocks portable-networking).
- `haxe_go-vfp.9.1`: Governed Go-native capability matrix (blocks go-native).
- `haxe_go-vfp.4.8`: Checksummed same-SHA hosted release assets and provenance (blocks published beta-baseline artifact).
- `haxe_go-vfp.6.3`: Generated public API manifest (blocks public API admission).
- `haxe_go-vfp.6.4`: Pre-1.0 SemVer, deprecation, and stable-1.x admission policy (blocks compatibility policy closure).
- `haxe_go-vfp.12.7`: Separate stable-1.x admission decision (blocks any stable 1.x claim).

## Machine authority

- `docs/compatibility-support-manifest.json`
- verify with `npm run compatibility:verify`
