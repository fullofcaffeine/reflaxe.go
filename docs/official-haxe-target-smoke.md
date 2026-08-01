# Official Haxe target smoke

## What this proves

This lane is an **official-suite smoke** for the portable compiler scorecard.
It executes three official Haxe behavior owners whose selected bodies and
expected values remain unmodified. They pass through a repository-authored
narrow target adapter, the installed Haxe.Go package, generated Go, `gofmt`,
`go test`, `go build`, and `go run`:

- top-level `unit.TestNumericSuffixes.testFloatSuffixes`;
- generated `unitstd/IntIterator.unit.hx`;
- general regression `unit.issues.Issue6369.test`.

The checkout is pinned to Haxe commit
`e0b355c6be312c1b17382603f018cf52522ec651` (Haxe 4.3.7), and the runner is
pinned to utest commit `a94f8812e8786f2b5fec52ce9f26927591d26327`.
The machine authority is
[`test/official_haxe_target_smoke/manifest.json`](../test/official_haxe_target_smoke/manifest.json).

This is not the complete applicable official `tests/unit` contract. It does
not broaden release compatibility, qualify the metal preset or explicit
Go-native APIs, or turn the current repo-local stdlib sweep into official-suite
evidence.

The adapter invokes the exact selected methods and uses pinned utest in
fail-fast mode. The unitstd body is expanded by the pinned official
`unit.UnitBuilder.read` macro. It deliberately does not select utest's
asynchronous runner, reports, timers, or stack-trace behavior; a green smoke
therefore cannot advance claims for those independent surfaces.

The selected official methods call `eq`, `t`, and `f`. A local assertion
carrier delegates exactly those calls to pinned utest but does not implement
utest's asynchronous discovery interface. For class-based owners, the runner
creates a temporary copy after hash verification, changes only package, class,
and base scaffolding, and adds that assertion carrier. The complete official
class body—including every expected value—remains unchanged; the temporary
adapted source is neither committed nor published.
Machine-readable lines leave the
program through a repository-owned `@:goNative` typed `hxrt` display observer; that output
transport proves only that the target executable ran and is not counted as
portable or metal API evidence.

## Why active runtime records matter

Official Haxe sources contain target guards, dummy assertions, and dynamically
generated cases. File counts therefore overstate coverage. The runner compares
the active runtime test identities and positive assertion counts with the
manifest. A source hash change, missing or renamed source, missing/extra active
test, dummy-only result, failed assertion, Go build error, runtime crash, or
timeout fails the lane.

## Provenance and artifact flow

The runner fetches or reuses ignored exact-commit checkouts. It verifies commit,
selected-source SHA-256, and MIT license evidence before compiling. No official
test body is committed here. In particular, the security-sensitive upstream
TLS fixture and its private-key material are not selected, copied, or included
in artifacts.

The compiler is consumed from a newly built ZIP installed into an isolated
Haxelib repository. Every Haxe application directory is created below that
repository, and the runner repeats `haxelib path reflaxe.go` from the exact
compile working directory. It fails if resolution names the source checkout or
does not name the isolated repository. Reports record this confinement result,
the source Git identity, dirty-state digest, toolchain versions, upstream pins,
active runtime results, generated-file hashes, and each command outcome under
`.cache/official-haxe-target-smoke`.

The result also records the byte length and SHA-256 of the deterministic
Haxelib ZIP that was installed. This binds a dirty local proof—and a clean
exact-SHA CI proof—to the precise package bytes consumed by Haxe, rather than
only to a list of changed paths.

Before upload, the runner scans every generated file, result record, and
stdout/stderr log for repository, temporary-workspace, isolated-sandbox, and
pinned-checkout absolute paths. A successful tool warning therefore cannot
leak a machine-local path into a green evidence artifact.
If confinement or any earlier lane stage fails, the runner removes the artifact
directory before returning nonzero. The workflows may still attempt their
`if: always()` upload, but there is no rejected bundle left to publish.

Run the required lane with:

```bash
npm run test:official-haxe-smoke
```

The required command also exercises deliberate assertion, Go-build, runtime,
timeout, and missing-selected-source failures. The last control preserves the
valid pinned license and other selected sources, removes one actual selected
test, and requires the selected-source check itself to reject it. These controls
must each be observed as nonzero before the overall lane can pass.

## Deferred expansion

The full applicable active inventory remains intentionally deferred after this
tracer: all shared top-level classes, portable `unitstd`, general issues,
generic historical `hxcpp_issues`, and capability shards must be classified and
executed before wording can advance beyond “official-suite smoke.”
