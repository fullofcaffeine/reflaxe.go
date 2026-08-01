# Official Haxe target smoke

In plain terms, this test checks the compiler the way a user would receive it:
it builds an installable Haxelib package, installs that package in a clean
location, compiles selected official Haxe tests to Go, asks Go to format and
build the result, and runs the program. It catches packaging and target-wiring
bugs that a test run directly from the source checkout can miss.

## Run it

```bash
npm run test:official-haxe-smoke
```

A successful run means all of these steps happened:

1. exact pinned Haxe and utest sources were verified;
2. a new Haxe.Go ZIP was installed into an isolated Haxelib repository;
3. Haxe resolved the compiler from that installation, not this checkout;
4. three selected official tests compiled through Haxe.Go;
5. generated Go passed `gofmt`, `go test`, and `go build`;
6. the program ran three expected test methods and completed 16 assertions;
7. deliberate assertion, build, runtime, timeout, and missing-source failures
   were each detected.

## What this proves

An **official-suite smoke** is a small but real sample of the upstream Haxe test
suite. It proves that the selected path works end to end; it does not claim that
every applicable official Haxe test already passes.

The portable compiler scorecard is the evidence record for ordinary portable
Haxe semantics. This smoke contributes only to that scorecard. It executes
three selected official Haxe tests whose behavior and expected values remain
unchanged:

- top-level `unit.TestNumericSuffixes.testFloatSuffixes`;
- generated `unitstd/IntIterator.unit.hx`;
- general regression `unit.issues.Issue6369.test`.

The checkout is pinned to Haxe commit
`e0b355c6be312c1b17382603f018cf52522ec651` (Haxe 4.3.7), and the runner is
pinned to utest commit `a94f8812e8786f2b5fec52ce9f26927591d26327`.
The machine-readable source list, hashes, and expected active tests live in
[`test/official_haxe_target_smoke/manifest.json`](../test/official_haxe_target_smoke/manifest.json).

This is not the complete applicable official `tests/unit` contract. It does
not broaden release compatibility, qualify the metal preset or explicit
Go-native APIs, or turn the current repo-local stdlib sweep into official-suite
evidence.

The small target adapter invokes the exact selected methods and uses pinned
utest in fail-fast mode. The `unitstd` test body is expanded by the pinned
official `unit.UnitBuilder.read` macro. It deliberately does not select utest's
asynchronous runner, reports, timers, or stack-trace behavior; a green smoke
therefore cannot advance claims for those independent surfaces.

The selected official methods call utest's `eq`, `t`, and `f` assertions. A
small local bridge sends those calls to pinned utest without bringing in
utest's unrelated asynchronous discovery system. For class-based tests, the
runner first verifies the upstream file hash, then creates a temporary copy
that changes only the package/class scaffolding needed to invoke it on this
target. The test bodies and expected values are unchanged, and the temporary
copies are neither committed nor published.

The generated program reports machine-readable results through a narrow typed
Go runtime observer. That observer proves the target program executed; it does
not count as evidence for the public Go-native API or the `metal` preset.

## Why active runtime records matter

Active inventory means the tests and assertions that actually execute for this
target, not simply the number of test files present in the upstream checkout.
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

The required command also exercises deliberate assertion, Go-build, runtime,
timeout, and missing-selected-source failures. The last control preserves the
valid pinned license and other selected sources, removes one actual selected
test, and requires the selected-source check itself to reject it. These controls
must each be observed as nonzero before the overall lane can pass.

## Deferred expansion

The full applicable active inventory remains intentionally deferred after this
tracer: all shared top-level classes, portable `unitstd`, general issues,
generic historical `hxcpp_issues`, and capability shards must be classified and
executed before wording can advance beyond “official-suite smoke.” That work is
tracked by Bead `haxe_go-vfp.12.10`.
