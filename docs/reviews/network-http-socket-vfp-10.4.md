# HTTP and socket lifecycle audit (`haxe_go-vfp.10.4`)

## Decision

The existing staged-source/runtime split remains correct. Haxe-visible
request, callback, response, upload, stream, and error semantics stay in
ordinary typed Haxe source. Go's standard networking packages stay behind
small typed `hxrt` capabilities that own native connections, deadlines,
request bodies, TLS sessions, and cleanup.

This audit closes concrete implementation gaps. An independent second-pass
review found that some evidence rows originally overstated semantic parity.
The [review disposition](network-admission-oracle-disposition-vfp-10.4.md)
admits only a narrow blocking IPv4 TCP client core and transfers HTTP plus
advanced socket/TLS work to `haxe_go-vfp.10.8` and `haxe_go-vfp.10.9`.

## Evidence matrix

| Risk | Contract | Expected result |
| --- | --- | --- |
| Multipart buffers or ignores the supplied `Input` | `http_multipart_streaming_contract`; `TestHttpRequestMultipartStreamsDeclaredInputWithPartialChunks` | The exact declared bytes arrive after multiple partial reads and native chunk requests are bounded. This does not prove cancellation of a blocked Haxe `Input` or safe callback execution context. |
| Upload ends early or its source fails | `http_multipart_streaming_contract`; `TestHttpRequestMultipartEarlyEOFAbortsTheExchangeAndReleasesTheServer` | The tested early-EOF/source-error cases abort. This does not prove that an early response or timeout unblocks a blocked source read or prevents callbacks after return. |
| HTTP timeout leaks a live request | `TestHttpRequestTimeout` | The client returns within the configured bound and the local server observes request-context cancellation. |
| Direct `customRequest` final-state behavior | `http_custom_request_lifecycle_contract` | For small complete responses, success writes and closes the Output with status-only callbacks; HTTP error reports status then error and leaves the Output open. Because native execution first buffers the complete body, this is not streamed lifecycle parity and does not preserve partial bytes after a transfer failure. |
| Socket write assumes full progress | `TestSocketWriteReportsPartialProgressForSourceOwnedWriteFullBytes` | Each native write reports its real progress and the source-owned write loop can send the remainder. |
| Peer closes after a partial read | `TestSocketPeerClosePreservesPartialReadThenReportsEOF` | The available bytes are returned before the following read reports EOF. |
| Accept or UDP receive waits forever | `TestSocketAcceptAndUDPReadHonorConfiguredTimeouts` | Both operations honor the configured socket timeout. |
| TLS peer accepts TCP but stalls the handshake | `TestSslSocketConnectHonorsTheConfiguredHandshakeTimeout` | `Socket.setTimeout` bounds both dialing and TLS handshake establishment. |
| Concurrent lifecycle races | `go test -race ./runtime/hxrt` | The complete runtime package passes under the Go race detector. |
| Platform-specific socket code stops compiling | `test_socket_runtime_cross_build.py` | Linux/amd64 and Windows/amd64 build-tagged adapters compile. This is compile-only evidence, not Windows runtime evidence. |
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

The local matrix is deterministic and covers valuable bounded cases, including
partial TCP progress, request timeout, early upload EOF, concurrent close, and
a stalled TLS handshake. It does not prove streamed HTTP response choreography,
cancellation of a blocked upload source, real socket write/exception readiness,
server backlog semantics, default TLS hostname identity, bounded DNS, zero-byte
UDP sender identity, resource convergence under repetition, or runtime behavior
outside Linux/amd64.

Consequently, only the named blocking IPv4 TCP client members are admitted.
HTTP and advanced socket, DNS, UDP, server, readiness, and TLS operations remain
release-excluded under their follow-up blockers.

Follow-up note (2026-07-30): `haxe_go-vfp.10.9.2` now preserves sender identity
for zero-byte UDP datagrams, and `haxe_go-vfp.10.9.3` now preserves the original
logical host for default direct-TLS verification and SNI. Those focused repairs
do not change this review's narrow TCP-only release disposition; the remaining
UDP and TLS contracts stay excluded under `haxe_go-vfp.10.9`.

## Validation result

The implementation matrix was reported as passing on 2026-07-28. These results
remain regression evidence for the bounded hardening work; they do not override
the operation-level release exclusions:

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
