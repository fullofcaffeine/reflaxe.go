# SSL And UDP Semantic-diff Spike

This spike records why `sys.net.UdpSocket` and the `sys.ssl` leaf surfaces stay
under snapshot/runtime evidence for now instead of semantic-diff evidence.

## Terms

- `semantic-diff`: the harness that runs the same Haxe program through Haxe
  `--interp` and generated Go, then compares stdout. See
  `/docs/semantic-diff-guide.md`.
- snapshot/runtime evidence: a test that proves generated Go compiles and runs a
  stable Go-side contract, without claiming byte-for-byte interpreter parity.
- leaf SSL surfaces: `sys.ssl.Certificate`, `sys.ssl.Digest`, and `sys.ssl.Key`.
  These are the parsing, hashing, signing, and key-loading APIs below
  `sys.ssl.Socket`.
- target-sensitive: behavior that depends on the runtime, operating system,
  scheduler, wall clock, stack frames, network, or TLS implementation.

## What was tested

### SSL digest probe

A narrow `sys.ssl.Digest` probe attempted to run this through Haxe `--interp`:

```haxe
import haxe.io.Bytes;
import sys.ssl.Digest;
import sys.ssl.DigestAlgorithm;

class Main {
	static function main() {
		var payload = Bytes.ofString("digest-payload");
		Sys.println(Digest.make(payload, DigestAlgorithm.SHA256).toHex());
	}
}
```

The interpreter failed with:

```text
Field index for make not found on prototype sys.ssl.Digest
```

That means `sys.ssl.Digest.make` is not available as an interpreter reference in
this harness shape. The generated Go implementation is real and covered by
`test/snapshot/stdlib/sys_ssl_leaf_direct`, but Haxe `--interp` cannot be used as
the reference oracle for that leaf behavior today.

### UDP loopback probe

The existing UDP loopback snapshot fixture was also tried through Haxe
`--interp`:

- `test/snapshot/stdlib/sys_net_udp_socket_direct/Main.hx`

The interpreter failed while constructing `sys.net.UdpSocket`:

```text
Uncaught exception Not available on this platform
```

That means `sys.net.UdpSocket` is not available on Haxe `--interp`, even for a
local loopback-only comparison. The generated Go fixture still has value: it
proves `bind`, `host`, `sendTo`, `readFrom`, peer address round-tripping, and
socket-option setup on Go. It is not interpreter-vs-Go semantic parity.

## Decision

Do not add semantic-diff fixtures for these surfaces yet:

- `sys.net.UdpSocket`
- `sys.ssl.Certificate`
- `sys.ssl.Digest`
- `sys.ssl.Key`
- `sys.ssl.Socket`

The current honest evidence remains:

- `test/snapshot/stdlib/sys_net_udp_socket_direct`
- `test/snapshot/stdlib/sys_ssl_leaf_direct`
- `test/snapshot/stdlib/sys_ssl_socket_direct`
- `test/snapshot/stdlib/sys_ssl_socket_sni_direct`
- `test/semantic_diff/sys_net_address_ssl_digest_algorithm_contract` only for
  the deterministic `sys.net.Address` and `sys.ssl.DigestAlgorithm` subset

## Why this is the right boundary

A semantic-diff fixture is useful only when the interpreter can serve as a
meaningful reference. For these APIs, the useful behavior is owned by the Go
runtime, Go crypto/x509, Go TLS, and OS networking. The interpreter either does
not expose the API in the tested shape or reports the surface as unavailable.

Adding semantic-diff tests anyway would create one of two bad outcomes:

- It would fail because the interpreter is not a valid oracle for the surface.
- It would test only constants or carriers that are already covered elsewhere,
  while giving the false impression that cryptographic or socket behavior has
  interpreter-vs-Go parity.

## Reopen trigger

Reopen this decision only when one of these becomes true:

1. the Haxe reference runtime exposes a deterministic implementation for the SSL
   leaf or UDP behavior being compared,
2. the harness can compare against a different explicit reference target instead
   of Haxe `--interp`,
3. the comparison is narrowed to a pure carrier/constant subset and the docs say
   clearly that crypto, TLS, and OS socket behavior are not covered,
4. a snapshot-only fixture starts hiding user-visible drift that a normalized
   runtime harness could catch.

Until then, snapshot/runtime evidence is the safer production contract.

## Related docs

- `/docs/known-gaps.md#target-sensitive-parity-policy`
- `/docs/portable-module-mapping-contract.md`
- `/docs/semantic-diff-guide.md`
