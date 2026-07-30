# Socket readiness and nonblocking controls

This document defines what `sys.net.Socket.select`, `waitForRead`, and
`setBlocking` mean on Haxe.Go. It also records what remains outside the release
promise.

## What the API reports

`Socket.select(read, write, others, timeout)` returns the exact `Socket` objects
from the caller's three arrays:

- **read readiness** means a read can make progress now. That includes bytes
  already buffered by Haxe.Go, bytes available in the operating system, and an
  EOF or reset that the next read must surface.
- **write readiness** means the operating system currently has room for at
  least some output. It does not promise that an arbitrarily large write will
  finish in one call.
- **exceptional readiness** is the operating system's exceptional-condition
  set. On TCP this is primarily urgent/out-of-band data. A new, disconnected,
  or closed Haxe socket is not fabricated as exceptional merely because it has
  no active connection.

Repeated entries stay repeated. If the same socket appears twice in an input
array and is ready, both source indexes are returned, so object identity and
the public `custom` value are preserved.

`waitForRead()` uses the same read-readiness path with no timeout.

## Why native readiness is necessary

The old implementation treated every connected socket as writable and every
disconnected socket as exceptional. That was a check of object state, not
readiness. In particular, it could claim a socket was writable after its
kernel send buffer had filled.

Haxe.Go now snapshots typed socket resources and calls an operating-system
`select` adapter. Go guarantees a raw connection descriptor only while its
`Control` callback is running, so the snapshot duplicates the descriptor
inside that callback, marks the duplicate close-on-exec atomically, and closes
it after each bounded poll slice. This prevents a concurrent source close,
descriptor-number reuse, or child-process launch from turning the temporary
poll capability into false readiness or an inherited resource. The native
adapter returns descriptor readiness only to the typed runtime; generated Haxe
receives indexes, never file descriptors. The runtime also checks its
`bufio.Reader` before polling so buffered bytes remain visible and are not lost
by a destructive probe.

The implementation waits in short bounded slices. This lets it resnapshot a
handle after another thread closes it rather than waiting forever on a stale
descriptor.

## Timeout behavior

The readiness timeout has three forms:

1. An omitted timeout waits until some requested socket is ready.
2. A zero timeout is a single immediate poll.
3. A positive timeout waits up to that many seconds.

For compatibility with Haxe's native targets, a negative `select` timeout is
treated like an omitted timeout. Separately, a negative
`Socket.setTimeout(value)` clears the socket's operation deadline;
`setTimeout(0)` requests an immediate operation deadline; and a positive value
sets the per-operation deadline in seconds. None of these values can bound the
earlier synchronous DNS work performed by `new Host("hostname")`; see the
[socket DNS boundary](socket-dns-boundary.md).

## `setBlocking(false)`

For an already connected socket or listening socket, `setBlocking(false)`
makes reads, writes, and accepts return promptly. If no progress is available,
staged Haxe raises `haxe.io.Error.Blocked`. A write may report a partial write
before a later attempt reports `Blocked`; callers must preserve that progress
and retry after write readiness.

The Go runtime uses a one-millisecond operation probe rather than permanently
changing the descriptor's `O_NONBLOCK` flag. This is an implementation detail:
the public contract is prompt progress-or-`Blocked` behavior, and focused
runtime tests prove it for connected read, write, and accept.

Nonblocking connect remains excluded. Go's `net.Dialer` closes a timed-out
connection attempt, so it cannot preserve an in-progress socket for a later
`select` call to complete. Claiming parity here would require a separate
native-connect state machine.

## Platform and release boundary

Linux and Darwin have build-tagged read, write, and exceptional readiness
adapters. Linux is the only release platform; Darwin tests are local runtime
evidence, not a release admission. Windows is compile-only evidence and the
runtime adapter reports that native readiness is unavailable rather than
returning invented results. Other operating systems are likewise unadmitted.

TLS readiness remains excluded. A `tls.Conn` can hold decrypted application
bytes inside TLS-specific buffers that are not represented by the underlying
descriptor set, so plain TCP evidence is not enough to claim TLS readiness.

`shutdown` and `setFastSend` have their own evidence-backed candidate contract;
see [socket shutdown and fast-send controls](socket-tls-controls.md). That
evidence does not expand readiness to TLS. DNS, UDP, hostile-peer behavior,
long-running resource convergence, and runtime support outside the admitted
platform retain their own exclusions.

This work does not admit these controls for release. The compatibility
manifest stays fail-closed under `haxe_go-vfp.10.9` until the complete advanced
socket review decides the exact member-level release surface.

## How the contract is checked

- `test/semantic_diff/socket_readiness_contract` checks portable `select`
  behavior, duplicate indexes, identity, buffered bytes, and ordinary write
  readiness against Haxe.
- `test/snapshot/sys/socket_readiness_contract` compiles and runs generated Go,
  including connected nonblocking read/write/accept behavior.
- `runtime/hxrt/socket_test.go` fills a TCP send buffer to native `EAGAIN`,
  verifies that write readiness disappears, drains the peer, and verifies that
  readiness returns. It also covers reset, EOF, duplicates, timeout forms, and
  buffered reads.
- `runtime/hxrt/socket_readiness_posix_test.go` checks native saturation and
  exceptional readiness with out-of-band data. It also closes the source
  socket while a snapshot exists and proves that the duplicated descriptor
  is close-on-exec and remains valid only until explicit snapshot release.
- `test/test_socket_runtime_cross_build.py` is compile-only evidence for
  non-runtime platforms; it is not described as runtime support.
