# Advanced socket/TLS Oracle disposition (`haxe_go-vfp.10.9.8`)

## Outcome

The GPT-5.6 Pro verdict is accepted: **split the surface and correct the native
ownership defects before admission**.

In plain language, networking is not one on/off feature. A blocking TCP client,
a TCP server, UDP datagrams, descriptor readiness, and TLS each have different
failure modes. The compatibility manifest therefore names only the exact
member groups supported by deterministic Linux/amd64 evidence. Whole classes
and other platforms remain outside the promise.

The review described commit `21acb7eb`. The fixes below are follow-up work and
must pass the repository's complete gates before their release flags are final.

## Finding-by-finding response

| Review finding | Decision and correction | Bead |
| --- | --- | --- |
| Close can return before an older TCP/TLS connect or TCP/UDP bind installs its resource. | Accepted. `SocketHandle` now advances a reusable lifecycle generation on close, owns cancellation functions for in-flight dials, rejects stale installation, and closes stale acquired resources. A later deliberately started operation uses the new generation and may succeed. | `haxe_go-vfp.10.9.8.1` |
| Readiness can close unrelated descriptor zero when `RawConn.Control` fails before its callback. | Accepted. Descriptor duplication now records ownership explicitly and closes only a successfully owned duplicate, including the valid case where that duplicate is descriptor zero. | `haxe_go-vfp.10.9.8.2` |
| Failed TLS handshakes retain unusable native connections. | Accepted. Public handshake, implicit `peerCertificate`, and accepted-server handshake share a transaction that detaches and closes the exact failed connection. A stale failure cannot detach a replacement. | `haxe_go-vfp.10.9.8.3` |
| Failed retained listener-policy installation leaves stale listener state. | Accepted. A failed post-listen deadline transition clears the listener, deadline capability, and wrapper, closes the new listener, and leaves the handle empty so the caller must bind again. | `haxe_go-vfp.10.9.8.4` |
| UDP creation/bind and deadline/broadcast policy are not transactional. | Accepted. Deadline and broadcast policy are applied before publication; failure closes the new descriptor; explicit bind uses the lifecycle generation; first `setBroadcast` applies exactly once; failed mutation preserves prior state. | `haxe_go-vfp.10.9.8.5` |
| Finite select can block behind a concurrent read's lock. | Accepted. Readiness snapshot preparation uses nonblocking lock probes and bounded retry slices. A signaling regression proves a 20 ms select returns within its allowance beside a read that has entered the native call. | `haxe_go-vfp.10.9.8.6` |
| Convergence documentation overstates individual accounting and accept-loop shutdown. | Accepted. The documentation now says the accept loops remain in the post-warm-up baseline, the counts are aggregate tolerances, GC is secondary evidence, and focused pre-GC ownership tests—not the aggregate count—prove each corrected transaction. | `haxe_go-vfp.10.9.8.7` |
| Compatibility operations are overbroad or linked to imprecise evidence. | Accepted. Broad rows were replaced with exact client, server, connected-control, readiness, shutdown, host, UDP, and TLS children. Whole-class UDP/TLS symbols are forbidden by an executable contract. Dedicated host and UDP evidence IDs were added and all generated outputs are regenerated from the source. | `haxe_go-vfp.10.9.8.7` |
| POSIX readiness has a fixed descriptor-capacity limit. | Accepted as an explicit narrow exclusion. High duplicated descriptor numbers fail deterministically; scalable polling is optional future hardening and is not implied by the beta claim. | no release blocker while excluded |

## Test-first evidence

The repaired paths have deterministic native regressions for:

- close cancellation of TCP and TLS dial work;
- stale TCP/TLS connection, TCP bind, and UDP bind rejection;
- successful intentional reuse after close;
- readiness `Control` failure before callback, duplicate failure, and an owned
  descriptor-zero duplicate;
- public, peer-certificate, and accepted TLS handshake failure;
- a failed handshake racing replacement;
- failed listener policy installation and subsequent clean reuse;
- UDP lazy creation, bind, and existing-option mutation failure;
- exactly-once broadcast application;
- finite select beside an entered blocking read.

The aggregate POSIX lifecycle matrix also uses an operation-entry signal for
its blocked-read cancellation case instead of relying on a scheduler sleep.

## Exact operation policy

Subject to the generated manifest's toolchain, trust, and platform boundary,
the split admits only these Linux/amd64 groups:

- blocking IPv4 TCP client core;
- blocking IPv4 TCP server core;
- established plain-TCP timeout and blocking-mode controls;
- connected nonblocking accept;
- plain-TCP readiness and `waitForRead`;
- established plain-TCP shutdown and fast-send;
- `Host.localhost`;
- blocking IPv4 UDP datagram core and broadcast control;
- application-controlled direct TLS client and server/SNI groups;
- TLS write/full-shutdown and fast-send controls.

The following remain excluded:

- named and reverse DNS timeout/cancellation guarantees;
- nonblocking connect;
- TLS and UDP readiness;
- high-descriptor readiness beyond fixed `FdSet` capacity;
- TLS read-only shutdown;
- IPv6, multicast, public CA-store portability, client certificates, public or
  hostile networks, production soak, and runtime support outside Linux/amd64;
- any member, error path, or concurrency pattern not explicitly named by the
  generated manifest.

Windows cross-builds remain compile-only evidence. Darwin local execution does
not widen the Linux/amd64 release platform.

## Architecture decision

The review strengthens the current boundary instead of changing it:

1. staged Haxe owns public objects, streams, exceptions, source identity, and
   Haxe-visible sequencing;
2. typed `hxrt` owns native connection acquisition, cancellation, descriptors,
   deadlines, socket options, readiness, TLS handshakes, and cleanup;
3. lifecycle generations and detach-if-current helpers are native ownership
   facts, not compiler lowering rules.

No universal IR, compiler networking shim, `Dynamic` native resource, raw Go
injection, or profile-specific network implementation is introduced.

## Closure rule

Neither `haxe_go-vfp.10.9.8` nor its parent closes merely because the code was
edited. Closure requires the focused tests, supported Go ordinary/race/checkptr
lanes, Haxe snapshots and semantic fixtures, compatibility/release contracts,
family mirror verification, and a post-fix second-pass review at the exact
landed commit. Until then, this document records the accepted design and
pending verification rather than claiming completion.
