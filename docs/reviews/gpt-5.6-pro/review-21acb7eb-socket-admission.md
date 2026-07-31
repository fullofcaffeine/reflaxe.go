# Independent advanced socket/TLS admission review

This is a normalized durable record of the review supplied by the user. It is
not presented as a verbatim transcript. The declared reviewer was GPT-5.6 Pro
at the deepest reasoning mode available in that session. The serving route was
not exposed and is not inferred here.

## Reviewed source

- repository: `fullofcaffeine/reflaxe.go`;
- candidate commit: `21acb7eb8d0c1ce23b4b2d2bde095d1e538d2962`;
- candidate tree: `ca4ef127c4c1d4bc69c981a42ab693bfb077460c`;
- candidate parent: `4851ebb913c372a626e3866fc01b947f6c0f88f5`;
- implementation commit: `dcbbf08323e68ec4fee4dbac1d7bc138eb6976e1`;
- previous network review commit: `5073a210e197c64b7fa09002214a2ea5085f6a29`;
- tracker item: `haxe_go-vfp.10.9.8`;
- parent blocker: `haxe_go-vfp.10.9`;
- review bundle: `haxe-go-socket-admission-21acb7eb.zip`;
- bundle SHA-256: `34e6103897208d33f99398cdebf83d10b5b7adf5705e446b356be33876eb5727`.

The reviewer reported that every bundled checksum passed and independently
confirmed that the embedded Git archive identifies the candidate commit. The
archive was treated as source authority; Repomix was only a navigation aid.

## Evidence boundary

The reviewer separated source inspection, personally rerun tests,
implementing-agent reports, and supplied exact-SHA CI evidence. The local
review environment had Go 1.23.2 on Linux/amd64, Python 3.13.5, Node 22.16.0,
and no Haxe executable. Its Go results were therefore supplementary, not
release-toolchain evidence. The supplied CI established Linux/amd64 execution
on Go 1.25.12 and 1.26.5. Windows remained compile-only. The macOS exact-SHA
quality job did not run `runtime/hxrt`, so Darwin resource convergence remained
local/reported evidence rather than release evidence.

## Verdict

**SPLIT.**

The broad advanced-socket surface could not be admitted at the reviewed
commit. The reviewer recommended preserving the existing blocking IPv4 TCP
client claim with narrower concurrency wording, splitting every other claim by
operation, admitting only already-safe children, and fixing concrete native
ownership defects before server, readiness, UDP, or TLS promotion.

## Findings at the reviewed commit

### Blockers

1. `close()` was not a lifecycle barrier against in-progress TCP/TLS connect,
   TCP bind, or UDP bind. An operation begun before close could later install a
   resource. The required design was a reusable lifecycle generation plus
   typed dial cancellation and stale-result closure—not a permanent closed
   flag.
2. If readiness `RawConn.Control` failed before invoking its callback,
   descriptor zero could be closed even though no duplicate descriptor was
   owned. Duplicate ownership needed an explicit Boolean because descriptor
   zero can also be a legitimate successful duplicate.
3. Failed public, implicit peer-certificate, and accepted-server TLS
   handshakes retained the failed connection. The exact failed connection had
   to be detached and closed without allowing a stale failure to detach a
   replacement.

### High severity

1. If a listener was created but retained deadline application failed, the
   handle kept stale closed-listener state. The failed transition needed to
   leave a documented empty/unbound handle.
2. A finite `Socket.select` could wait indefinitely for `readMu` behind a
   concurrent blocking read, so the public timeout did not include snapshot
   preparation.
3. UDP connection creation, bind, deadline installation, and broadcast-option
   installation were not one transaction. Failure could retain a descriptor,
   and the first `setBroadcast` could apply the option twice.

### Medium severity

1. The convergence suite was useful aggregate evidence, but its documentation
   overstated individual resource accounting and said accept loops stopped
   before measurement even though they remained in the baseline. Garbage
   collection was secondary convergence evidence, not proof of explicit
   ownership cleanup.
2. Compatibility records combined unrelated members, used whole UDP/TLS class
   names, and linked some rows to imprecise evidence. The reviewer required
   exact operation children, symbols, evidence identifiers, exclusions, and
   generated-artifact equality.

### Low severity

The POSIX `select` adapters used fixed-capacity `FdSet` values. A process with
high descriptor pressure could receive a deterministic capacity error. The
reviewer accepted this as a documented beta exclusion; a future scalable
`poll`/`ppoll`/`kqueue` adapter can remove it.

## Recommended release shape

At the reviewed commit, the reviewer recommended:

- keeping the serial blocking IPv4 TCP client core admitted;
- separately admitting established plain-TCP timeout/nonblocking controls,
  plain-TCP shutdown/fast-send, and `Host.localhost` after a policy-only split;
- keeping server/listener, readiness, UDP, and direct TLS false until the
  corresponding ownership defects were fixed;
- keeping named and reverse DNS deliberately excluded unless a separately
  typed bounded-resolver API is designed;
- retaining Linux/amd64 as the only release runtime platform;
- never treating a whole `sys.net.UdpSocket` or `sys.ssl.Socket` class as an
  admitted symbol.

The review explicitly approved the architecture: staged Haxe owns public
semantics and object identity, while typed `hxrt` owns native resources,
descriptors, cancellation, TLS, deadlines, and transactional cleanup. It found
no reason for a universal compiler IR, compiler-owned socket shims, raw Go
injection, `Dynamic` native handles, or portable-versus-metal network
semantics.

The repository's finding-by-finding response is recorded in
[`../socket-admission-oracle-disposition-vfp-10.9.8.md`](../socket-admission-oracle-disposition-vfp-10.9.8.md).
