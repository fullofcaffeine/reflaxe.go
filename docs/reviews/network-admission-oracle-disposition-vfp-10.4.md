# Independent network admission review disposition (`haxe_go-vfp.10.4`)

## What was reviewed

The user supplied an independent review that identified itself as GPT-5.6 Pro.
The serving route was not exposed, so this record does not invent one. The
review used the commit-pinned bundle for:

- repository commit `5073a210e197c64b7fa09002214a2ea5085f6a29`;
- commit subject `fix(network): harden HTTP and socket lifecycles`;
- tracker item `haxe_go-vfp.10.4`;
- bundle source `haxe-go-network-admission-5073a210.repomix.xml`.

The reviewer reported that the bundle checksums and its 90 advertised files
matched. The reviewer inspected the supplied source and evidence but did not
independently rerun the complete Haxe, Go, release, race, static-analysis, or
vulnerability suites.

## Decision

The verdict is **split**.

The implementation is useful hardening, but the original evidence described a
broader semantic result than the tests prove. The beta release now admits only
the blocking IPv4 TCP client core described in the generated compatibility
manifest. HTTP and every advanced socket, DNS, UDP, server, readiness, and TLS
operation remain explicitly excluded.

In plain language: a portable Haxe program may use the named blocking TCP
client operations against an application-controlled numeric IPv4 address on
the supported Linux release platform. The release does not yet promise that
HTTP, hostnames, servers, UDP, TLS, nonblocking sockets, or readiness behave
correctly in all the ways their public APIs imply.

## Finding disposition

| Review finding | Local verification | Durable owner |
| --- | --- | --- |
| Responses are fully buffered before staged `sys.Http` sees status/body progress, and partial bytes are lost on a later read failure. | Accepted. `HttpRequestExecute` uses `io.ReadAll`; the current fixture checks only the final state of small completed responses. | `haxe_go-vfp.10.8` |
| Multipart upload lets `net/http` call generated Haxe from its body-reading goroutine and does not provide a close operation that can unblock a waiting source read. | Accepted. The current boundary is an `io.Reader` callback, not a synchronized cancellable upload sink. | `haxe_go-vfp.10.8` |
| Parameter/header order and repeated values are collapsed; form/query behavior diverges. | Accepted. Native maps and `Set` calls erase ordered multiplicity, and the POST form begins with URL query values. | `haxe_go-vfp.10.8` |
| Multipart framing uses one public fixed boundary and can disagree with a caller-supplied content type. | Accepted. Boundary generation and content type are not one atomic native result. | `haxe_go-vfp.10.8` |
| Method/body selection, redirect, compression, status, timeout, `requestUrl`, and data-URL behavior inherit unreviewed `net/http` policy or differ from Haxe 4.3.7. | Accepted. These are deterministic source-contract gaps, not merely Internet hardening. | `haxe_go-vfp.10.8` |
| Multipart metadata permits control characters in part headers. | Accepted as HTTP hardening required before admission. | `haxe_go-vfp.10.8` |
| `Socket.select` reports connection presence as write readiness and absence as an exceptional condition. | Accepted. The existing easy-case fixture does not prove operating-system readiness. | `haxe_go-vfp.10.9` |
| Direct TLS loses the original logical hostname unless callers separately set one. | Accepted. The numeric resolved address is passed to TLS, so default certificate identity and SNI are not proved. | `haxe_go-vfp.10.9` |
| DNS is eager and outside socket timeout/cancellation policy. | Accepted as a required explicit decision: implement bounded resolution or keep it excluded. | `haxe_go-vfp.10.9` |
| `bind` starts listening and `listen(backlog)` is a no-op. | Accepted. Current loopback evidence does not prove the public lifecycle or backlog. | `haxe_go-vfp.10.9` |
| TLS inherits shutdown and fast-send APIs whose native implementation only handles plain TCP connections. | Accepted. Silent no-op behavior is not admissible API evidence. | `haxe_go-vfp.10.9` |
| Windows evidence is compile-only, resource convergence is not proved, and a zero-byte UDP datagram loses its sender. | Accepted. Release runtime admission remains Linux/amd64 only; extra operating systems require real runtime lanes. | `haxe_go-vfp.10.9` |

## Architecture result

The review confirms the existing ownership direction:

1. staged Haxe owns Haxe-visible callbacks, exceptions, streams, public fields,
   request policy, and socket identity;
2. typed `hxrt` capabilities own Go and operating-system resources;
3. the HTTP seam should become a typed streaming exchange plus a typed upload
   sink;
4. the socket seam needs typed logical-host and operating-system readiness
   capabilities.

These are small boundary changes. They do not justify a compiler-wide
intermediate representation, a compiler-owned networking shim, profile-based
network semantics, raw injection, or a second network backend.

## Release-policy amendment

One recommendation was adapted to the repository's existing release
architecture. `release/readiness-policy.json` keeps only the coarse admitted
scope names `preset:portable` and `platform:linux-amd64`. Exact operation and
member admission is already canonicalized and hashed from
`docs/compatibility-support-manifest.json`; adding an operation name to
`admittedScopes` would create a second authority.

The governed changes are therefore:

- admit `tcp-ipv4-blocking-client-core` in the operation manifest;
- replace the old `portable-networking` exclusion with `portable-http`, owned
  by `haxe_go-vfp.10.8`;
- add `portable-socket-advanced`, owned by `haxe_go-vfp.10.9`;
- preserve the default rule that everything not explicitly admitted is
  excluded.

This document is the required written second-pass disposition for the
`thinking:xhigh` closure of `haxe_go-vfp.10.4`.

## Follow-up resolution status

Two accepted defects have since received focused repairs without widening the
release claim:

- `haxe_go-vfp.10.9.2` preserves the sender address for a valid zero-byte UDP
  datagram in native and generated-Haxe loopback evidence.
- `haxe_go-vfp.10.9.3` introduces a typed endpoint that dials the resolved
  address while retaining the original logical hostname for default direct-TLS
  verification and SNI. Deterministic local-CA evidence covers success,
  mismatch rejection, explicit `setHostname` override, and SNI selection.
- `haxe_go-vfp.10.9.4` preserves eager Haxe 4.3.7 `Host` construction and
  records the exact boundary: socket timeout values cannot cancel or bound DNS
  that completes before a Socket exists. The operation stays release-excluded
  behind an executable compatibility-policy contract.
- `haxe_go-vfp.10.9.5` restores the Haxe TCP server lifecycle through a typed
  build-tagged OS boundary: bind reserves without listening, listen applies
  the requested nonnegative backlog, repeated listen reapplies it, and close
  releases either state. Generated-Haxe, TLS-wrapper, concurrency, and Linux
  bounded-queue regressions are present; server admission still awaits the
  advanced-socket parent review, and Windows evidence remains compile-only.
- `haxe_go-vfp.10.9.6` replaces connection-presence guesses with typed,
  build-tagged native readiness. Saturation/drain, buffered bytes, EOF/reset,
  duplicate identity, POSIX urgent data, timeout forms, and connected
  nonblocking read/write/accept have native and generated-Haxe evidence.
  Nonblocking connect, TLS readiness, and Windows runtime readiness remain
  explicitly excluded; shutdown and fast-send stay in their own operation so
  this evidence cannot imply their support.
- `haxe_go-vfp.10.9.7` makes inherited TLS controls truthful. Write shutdown
  sends TLS `close_notify` while retaining reads, read-only TLS shutdown
  reports an exact unsupported error, full close is idempotent, and
  `setFastSend` updates the TCP transport beneath TLS through a typed
  `NetConn` boundary. Wire-visible generated-Haxe and native option-state
  tests cover repetition and concurrent close. The operation remains
  release-excluded pending the advanced-socket parent review.

UDP and TLS remain release-excluded until their complete operation-level
contracts and the parent `haxe_go-vfp.10.9` review are complete.
