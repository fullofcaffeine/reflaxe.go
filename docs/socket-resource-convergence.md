# Socket and TLS resource convergence

This document defines the cleanup evidence used by the advanced socket
admission review. It measures whether repeated success and failure leave
resources behind; it does not decide release admission by itself.

## What is repeated

The native POSIX suite performs a warm-up and then 20 repetitions of one
combined lifecycle:

1. blocking IPv4 TCP connect, write, read, close, and repeated close;
2. a read timeout against a peer that accepts data but never responds;
3. a TCP reset that can race with connection-policy installation;
4. a stalled TLS handshake bounded by `Socket.setTimeout`;
5. empty and non-empty UDP loopback datagrams;
6. bind, listen, accept, native readiness, a blocked read canceled by
   concurrent close, and listener cleanup;
7. TLS I/O, write-side `close_notify`, peer EOF, preserved reverse response,
   and complete close.

The reset case found one real rollback defect: a peer could reset after native
connect but before deadline/fast-send policy was applied. The failed
connection was installed and the public call threw without detaching or
closing it. Connection installation is now transactional, and a focused
regression requires the failed resource to be both closed and absent from the
typed handle.

## What is measured

After the repetitions, the suite waits for all three tracked server modes to
return to zero active accepted connections. It then requires:

- a bounded goroutine count relative to the post-warm-up baseline;
- a bounded Linux file-descriptor count read from `/proc/self/fd`;
- no blocked read, accept loop, TLS handshake, or per-connection handler left
  running.

Darwin runs the lifecycle and goroutine checks locally. Linux additionally
requires the descriptor source to exist and is the release runtime lane.
Windows remains compile-only evidence and is not described as a runtime
resource result.

The limits are convergence bounds, not promises that Go uses an exact fixed
number of internal descriptors or goroutines. The test permits a small
runtime delta while requiring active application connections to reach exactly
zero.

## What this evidence supports

The suite adds cleanup evidence to:

- the already admitted blocking IPv4 TCP client core;
- the candidate TCP server lifecycle;
- connected timeout, nonblocking, and readiness controls;
- shutdown and fast-send controls;
- UDP loopback behavior;
- direct TLS lifecycle and inherited controls.

It does not cover DNS cancellation, nonblocking connect, TLS readiness,
IPv6, public certificate stores, public or hostile networks, arbitrary load,
or runtime behavior outside the reviewed platform. Twenty bounded
repetitions are deterministic regression evidence, not a production soak or a
claim about indefinite hostile peers.

## Admission boundary

Every advanced operation remains `release_admitted=false` behind
`haxe_go-vfp.10.9` in the pre-review compatibility source. Resource cleanup is
necessary but not sufficient: semantics, platform evidence, exclusions, and
the usefulness of each proposed boundary still require an independent
commit-pinned review.

This evidence therefore does not admit an operation, remove the
`portable-socket-advanced` release exclusion, or close the parent blocker on
its own.

