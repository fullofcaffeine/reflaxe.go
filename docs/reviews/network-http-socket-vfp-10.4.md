# HTTP and socket lifecycle audit (`haxe_go-vfp.10.4`)

## Decision

The existing staged-source/runtime split remains correct. Haxe-visible
request, callback, response, upload, stream, and error semantics stay in
ordinary typed Haxe source. Go's standard networking packages stay behind
small typed `hxrt` capabilities that own native connections, deadlines,
request bodies, TLS sessions, and cleanup.

This audit closes concrete implementation gaps. It does not by itself widen
the release manifest's portable-networking promise.

## Evidence matrix

| Risk | Contract | Expected result |
| --- | --- | --- |
| Multipart buffers or ignores the supplied `Input` | `http_multipart_streaming_contract`; `TestHttpRequestMultipartStreamsDeclaredInputWithPartialChunks` | The exact declared bytes arrive after multiple partial reads; the runtime requests bounded chunks and does not retain a generated Haxe object. |
| Upload ends early or its source fails | `http_multipart_streaming_contract`; `TestHttpRequestMultipartEarlyEOFAbortsTheExchangeAndReleasesTheServer` | The exchange aborts, the server observes termination, and a source exception keeps its message. |
| HTTP timeout leaks a live request | `TestHttpRequestTimeout` | The client returns within the configured bound and the local server observes request-context cancellation. |
| Direct `customRequest` duplicates callbacks or leaks its output | `http_custom_request_lifecycle_contract` | Success writes and closes the Output with status-only callbacks; HTTP error reports status then error and leaves the Output open; response fields remain unset. |
| Socket write assumes full progress | `TestSocketWriteReportsPartialProgressForSourceOwnedWriteFullBytes` | Each native write reports its real progress and the source-owned write loop can send the remainder. |
| Peer closes after a partial read | `TestSocketPeerClosePreservesPartialReadThenReportsEOF` | The available bytes are returned before the following read reports EOF. |
| Accept or UDP receive waits forever | `TestSocketAcceptAndUDPReadHonorConfiguredTimeouts` | Both operations honor the configured socket timeout. |
| TLS peer accepts TCP but stalls the handshake | `TestSslSocketConnectHonorsTheConfiguredHandshakeTimeout` | `Socket.setTimeout` bounds both dialing and TLS handshake establishment. |
| Concurrent lifecycle races | `go test -race ./runtime/hxrt` | The complete runtime package passes under the Go race detector. |
| Platform-specific socket code stops compiling | `test_socket_runtime_cross_build.py` | Supported cross-build targets compile their build-tagged adapters. |
| TLS/SNI behavior regresses | `stdlib/sys_ssl_socket_direct`, `stdlib/sys_ssl_socket_sni_direct` | Local TLS I/O, peer certificates, accepted SSL identity, default certificate, and SNI selection remain runnable. |

## Ownership check

- `std/go/_std/sys/Http.hx` owns public request selection, multipart metadata,
  Input error translation, response fields, callback order, and Output
  lifecycle.
- `std/hxrt/http/NativeHttp.hx` exposes a typed chunk callback and opaque
  request/response handles.
- `runtime/hxrt/http.go` owns multipart framing, `net/http` execution,
  cancellation, response-body closure, and idle-transport cleanup.
- staged `sys.net` and `sys.ssl` types own Haxe streams, errors, identity, and
  configuration.
- `runtime/hxrt/socket.go` and `socket_ssl.go` own concrete native progress,
  deadlines, connections, and TLS resources.

No new compiler-owned stdlib shim, raw `__go__` island, `Dynamic` native
handle, or profile-name semantic branch was introduced.

## Scope boundary

The local matrix is deliberately deterministic and hostile enough to cover
partial progress, cancellation, early closure, and a stalled handshake. It
does not exercise every external HTTP proxy, public CA store, DNS
configuration, operating system, or real Internet failure mode. A release
policy may use this evidence when deciding whether to admit the broader
portable-networking surface, but that admission is a separate
provenance-sensitive decision.

## Validation result

The implementation matrix passed on 2026-07-28:

- `npm test`: 304/304 snapshots plus compiler, governance, cross-build, and
  terminal contracts;
- `npm run test:semantic-diff`: 149/149 portable parity cases;
- `npm run test:stdlib-sweep:go-test`: 55/55 upstream modules;
- `npm run test:examples`: 12/12 runnable example/profile lanes;
- `go test -race ./runtime/hxrt` and the repository Go tooling matrix
  (race, checkptr, vet, and staticcheck);
- `npm run security:deps` with supported Go 1.25.12: no reachable Go
  vulnerabilities. The machine's older Go 1.25.6 correctly failed closed on
  standard-library advisories that are fixed in the supported patch release.

The npm audit reports one moderate `ajv` advisory; the governed dependency
audit fails at high severity and above, so this is recorded but is not a
release-blocking finding under the current policy.
