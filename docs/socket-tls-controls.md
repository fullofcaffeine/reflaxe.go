# Socket shutdown and fast-send controls

This document defines how `sys.net.Socket.shutdown` and
`sys.net.Socket.setFastSend` behave on plain TCP and on inherited
`sys.ssl.Socket` calls. It also records the evidence and the release boundary.

## Shutdown behavior

`shutdown(read, write)` closes only the requested direction when the transport
can express that operation:

| Transport | `read=true, write=false` | `read=false, write=true` | both `true` |
| --- | --- | --- | --- |
| plain TCP | stop the local read half with `CloseRead` | stop the local write half with `CloseWrite`; peer reads EOF | stop both TCP halves |
| TLS over TCP | raise `TLS read-only shutdown is unsupported` | send TLS `close_notify`; peer reads EOF while the caller's read side remains available | close the complete TLS connection and release the handle |

Passing `false, false` does nothing. Repeating a completed TLS write shutdown
or closing an already closed handle is safe.

The TLS asymmetry is deliberate. Go's TLS connection has a protocol-level
`CloseWrite`, but no safe `CloseRead` equivalent. Pretending that read-only
shutdown succeeded would leave the connection unchanged; closing the whole
connection would violate the caller's request. Haxe.Go therefore reports the
unsupported direction exactly instead of choosing either misleading behavior.

## Fast-send behavior

`setFastSend(true)` enables TCP `TCP_NODELAY`, which favors prompt small writes
over waiting to combine them. `setFastSend(false)` disables that option.

TLS is a framing layer around TCP. The typed runtime follows the TLS
connection's `NetConn` capability until it reaches the underlying TCP
connection, then updates that socket. Generated Haxe never receives a
`net.Conn`, `tls.Conn`, or file descriptor. A connected transport that does not
expose TCP control reports `socket fast-send requires a TCP connection`
instead of silently accepting the request.

The preference is also retained before connection or acceptance. Once a TCP
connection is installed, the runtime applies the stored value to that direct
or TLS-wrapped transport.

## Why the proof is wire-visible

A call that merely does not throw does not prove either control worked. The
tests therefore check effects outside the caller:

- a real TLS peer receives application bytes followed by `close_notify` and
  EOF, then sends a response that the caller can still read;
- read-only TLS shutdown returns the exact unsupported error and the
  connection remains usable;
- operating-system `TCP_NODELAY` state changes on the TCP socket beneath TLS;
- plain TCP half-close remains bidirectional;
- repeated write shutdown, repeated full close, and concurrent close/shutdown
  complete without a panic or duplicate Haxe-visible error.

The generated-Haxe TLS snapshot performs the public shutdown sequence and
observes both the peer EOF and the preserved response path. The native POSIX
test measures `TCP_NODELAY` directly; the snapshot's `setFastSend` call is not
option-state proof. It is not a no-throw test.

## Architecture and platform boundary

Staged `sys.net.Socket` and `sys.ssl.Socket` continue to own public methods,
Haxe exceptions, and stream behavior. The typed `hxrt` boundary owns the
native connection, TLS `close_notify`, and TCP socket option. This is a small
runtime seam and does not require compiler lowering, a universal IR, raw Go
injection, `Dynamic`, or a profile-specific implementation.

The focused native controls test is buildable and runnable on Linux and
Darwin. Linux is the only release platform; Darwin is local runtime evidence.
Windows remains compile-only and no Windows runtime result is claimed.
TLS read-only shutdown remains explicitly unsupported. Hostile-peer behavior,
production-soak behavior, and runtime support outside the reviewed platform
remain excluded. A bounded 20-repetition matrix now covers successful TLS I/O,
write-side close-notify, complete close, and stalled-handshake cleanup; see
[socket and TLS resource convergence](socket-resource-convergence.md).

Native TLS ownership follows four failure invariants: close cancels a TLS dial
and prevents stale installation; a failed TLS handshake detaches and closes
the exact failed connection; a failed accepted handshake cannot leave an
unreachable connection attached; and a stale failed handshake cannot detach a
replacement connection. Public handshake, implicit `peerCertificate`
handshake, and accepted-server handshake share that transaction.

The compatibility manifest now admits three separate Linux/amd64 member
groups: direct application-controlled TLS clients, application-controlled
server/SNI behavior, and write/full-shutdown plus fast-send controls. It does
not admit the whole `sys.ssl.Socket` class. Public CA-store portability, client
certificates, TLS readiness, TLS read-only shutdown, hostile/load behavior,
and runtime support outside Linux/amd64 remain excluded.
