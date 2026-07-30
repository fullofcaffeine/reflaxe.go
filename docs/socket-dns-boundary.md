# Socket DNS and timeout boundary

## Decision: preserve eager `Host` resolution

Haxe.Go keeps the Haxe 4.3.7 `sys.net.Host` contract: constructing
`new Host("hostname")` synchronously resolves the name and fills the public
`ip` field before the constructor returns. It may throw if the name cannot be
resolved.

This resolution happens before a `Socket` participates. A socket timeout
therefore cannot cancel, limit, or retroactively change it. Hostname lookup,
reverse lookup, and local-host discovery remain outside the beta networking
claim. The admitted client operation continues to require an
application-controlled numeric IPv4 endpoint.

## Why this is the compatible choice

The public API fixes the order:

1. `new Host(name)` creates a host and immediately makes its numeric `ip`
   available.
2. `new Socket()` creates a separate unconnected socket.
3. `socket.setTimeout(...)` configures that socket.
4. `socket.connect(host, port)` receives the already-created `Host`.

There is no socket or timeout value at step 1. Delaying resolution until
`connect` would change when exceptions occur, when `ip` becomes meaningful,
and what `toString()` and `reverse()` can do. Starting a hidden background
resolver would still provide no public cancellation handle and could outlive
the constructor that started it.

A future context-bounded resolver would need to be a separate, explicit API
with a typed cancellation owner. It must not silently redefine
`sys.net.Host`.

## What each timeout value currently controls

These are current target behaviors, not broader socket-control release
admission:

- A positive `Socket.setTimeout` value configures the socket's later TCP dial
  and TLS handshake and installs an absolute deadline on its native connection
  or listener.
- `setTimeout(0)` installs an immediate deadline for later socket operations
  and connection setup.
- A negative timeout clears the configured socket deadline. A separately
  nonblocking socket still uses its immediate nonblocking deadline policy.

All three forms act only on that `SocketHandle`. A socket timeout does not
retroactively bound `new Host("hostname")`, `Host.reverse()`, or
`Host.localhost()`.

The exact progress, nonblocking, readiness, and zero/negative timeout contract
remains release-excluded under `haxe_go-vfp.10.9`; this decision only prevents
those controls from being mistaken for DNS cancellation.

## Operation-by-operation boundary

| Operation | What happens | Beta release status |
| --- | --- | --- |
| `new Host("127.0.0.1")` | Parses the numeric IPv4 literal locally and fills `ip`. | Admitted only as part of the named blocking IPv4 TCP client core. |
| `new Host("hostname")` | Calls the Go resolver synchronously during construction. | Experimental; lookup timeout, cancellation, resolver configuration, search domains, and split-horizon behavior are excluded. |
| `Host.reverse()` | Calls reverse DNS synchronously for the stored IPv4 address. | Experimental; timeout, cancellation, and result policy are excluded. |
| `Host.localhost()` | Reads the operating-system hostname synchronously; it is not controlled by a socket timeout. | Experimental and outside the admitted client operation. |
| `Socket.connect(host, port)` | Dials `host.toString()`, which is already a numeric address. | Only the documented blocking numeric-IPv4 client members are admitted. |

## Practical guidance

For the admitted beta path, construct `Host` with an application-controlled
numeric IPv4 literal and then connect. If an application requires bounded DNS,
resolver selection, or cancellation, perform that work through an explicit
application-owned native boundary and pass the resulting numeric IPv4 address
into `Host`. That resolver behavior is not part of the portable Haxe.Go release
claim.

## Executable policy

`test/test_socket_dns_boundary_contract.py` protects this decision. It checks
the eager staged-source sequence, the absence of DNS work from
`Socket.connect`, the operation-level compatibility exclusion, and the
`portable-socket-advanced` readiness blocker. The generated compatibility
manifest remains the release authority.
