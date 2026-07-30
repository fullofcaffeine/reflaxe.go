# Socket server lifecycle and backlog

## What changed

Haxe.Go now preserves the Haxe 4.3.7 TCP server sequence:

1. `socket.bind(host, port)` reserves the local IPv4 endpoint.
2. `socket.listen(connections)` starts listening and supplies the pending
   connection backlog.
3. `socket.accept()` returns one connected child socket.

Binding does not accept connections. This distinction matters when an
application must finish configuration before making a server reachable, and it
keeps the public `connections` argument from becoming decorative.

The operation remains release-excluded until the advanced socket parent review
closes. Correcting the implementation is evidence for an admission candidate;
it is not, by itself, a production-server promise.

## Why Go needs a narrow native adapter

Go's ordinary `net.Listen` API performs socket creation, bind, and listen as one
operation. It also chooses its own backlog. Calling it from Haxe's `bind` method
therefore made the endpoint reachable too early, while the later `listen`
method had no work left to do.

Haxe.Go instead retains a typed, opaque pre-listen socket inside
`SocketHandle`. Build-tagged runtime files perform the platform descriptor
operations:

- `socket_listener_posix.go` owns the POSIX descriptor;
- `socket_listener_windows.go` owns the Winsock handle;
- `socket_listener_unsupported.go` returns an explicit unsupported error.

Generated Haxe never receives a descriptor or native listener. After the OS
listen transition, the adapter converts the socket into Go's pollable
`net.TCPListener`, so normal deadlines, concurrent close, TLS wrapping, and
accepted stream I/O continue to use Go's network runtime.

An unbounded Go-side queue would not be equivalent to an operating-system
backlog, so Haxe.Go does not emulate one.

## Exact lifecycle

| State | Allowed transition | Result |
| --- | --- | --- |
| New | `bind(host, port)` | The numeric IPv4 endpoint is reserved but not listening. `host()` reports the reserved address, including an OS-selected port when `port` is zero. |
| Bound | `listen(connections)` | A nonnegative backlog is passed to the operating system and the handle becomes a listener. |
| Listening | `listen(connections)` | The new nonnegative backlog is passed to the existing native listener. This makes repeated calls deterministic without silently ignoring the argument. |
| Bound or listening | `close()` | The descriptor/listener is released. A blocked `accept()` is interrupted after listening has started. |
| New or closed | `listen(connections)` | Fails because no bound endpoint exists. |
| Any | `listen(-1)` | Fails before an OS call because a negative pending-connection count is not a valid Haxe.Go request. |

The backlog is passed exactly as an integer, but the operating system may cap
or normalize its queue. Haxe.Go promises not to replace the requested value
with Go's default; it does not promise that every kernel admits precisely the
same number of simultaneous handshakes.

## Accepted-socket policy

An accepted child retains the listener's configured timeout, blocking, and
fast-send policy, and installs those settings on its own connection. The
current blocking accept and timeout evidence is deterministic.

`haxe_go-vfp.10.9.6` subsequently added real OS readiness and proved the
connected nonblocking accept progress-or-`Blocked` behavior. Nonblocking
connect and the broader server surface remain release-excluded under the
parent `haxe_go-vfp.10.9`; see the
[socket readiness contract](socket-readiness-nonblocking.md).

## TLS servers

`sys.ssl.Socket.bind` now builds and retains its typed TLS server policy while
reserving the TCP endpoint. The inherited `listen(connections)` performs the
real TCP listen transition and then wraps that listener with TLS. Certificate
selection and SNI remain native TLS policy, while public sequencing stays in
staged Haxe.

This means TLS bind no longer starts a server early and no longer loses the
public backlog value. Broader TLS admission, shutdown, fast-send, public trust,
hostile-peer, and cross-platform runtime guarantees remain separate.

## Platform and release boundary

- Linux/amd64 is the only platform eligible for the current beta networking
  claim. The Linux runtime test exercises a deliberately small backlog without
  accepting and requires the pending queue to become bounded.
- POSIX and Windows adapters are compile-checked. A Windows cross-build is
  compile evidence, not Windows runtime evidence.
- No Darwin, Windows, or other operating system is admitted by this change.
- Listener load, hostile peers, nonblocking connect, and long-duration server
  behavior remain excluded pending the parent review. Real readiness and
  connected nonblocking accept now have candidate evidence, but are not
  release-admitted by this document.

## Evidence

- `test/semantic_diff/socket_server_lifecycle_contract` compares the public
  bind/listen sequence and a loopback round trip with Haxe 4.3.7.
- `test/snapshot/sys/socket_server_lifecycle_contract` runs the generated Go
  form and guards the runtime file selection.
- `runtime/hxrt/socket_test.go` covers pre-listen refusal, lifecycle errors,
  duplicate listen, close-before-listen, inherited policy, TLS composition,
  and concurrent listen/close.
- `runtime/hxrt/socket_listener_linux_test.go` exercises a bounded Linux
  pending queue.
- `test/test_socket_runtime_cross_build.py` keeps Linux and Windows builds
  compiling while preserving the runtime-versus-compile-only distinction.
