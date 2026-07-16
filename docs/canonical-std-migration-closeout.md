# Canonical `_std` Migration Closeout

This document records the final source, package, and consumer contract for the
canonical Reflaxe standard-library migration. It is derived evidence, not a
second source of ownership truth: per-file ownership remains authoritative in
`docs/stdlib-provenance-ledger.json`.

## What changed

The source checkout and an installed Haxelib package intentionally use two
different file shapes for the same target overrides:

| Context | Override shape | Why |
| --- | --- | --- |
| Source checkout | ordinary `.hx` under `std/go/_std` | Review and edit normal Haxe modules in the canonical target root. |
| Installed package | generated `.cross.hx` under `src` | Let Haxe select target-specific replacements from the flattened package classpath. |
| Both contexts | ordinary support under `std/haxe`, `std/sys`, `std/hxrt`, and `std/go` | Support, typed runtime bindings, and public Go facades are not upstream replacements. |

The package runner is the only owner of the `.hx` to `.cross.hx` name change.
File bytes do not change, and the package manifest records the source path,
package path, ownership kind, size, and both SHA-256 hashes.

## Why the migration is behavior-neutral

This migration changes selection and package paths, not Haxe or generated-Go
semantics. The intended snapshot result is therefore zero generated-Go or
stdout snapshot changes. Any later generated output change must be explained
as a separate semantic or compiler change; it cannot be accepted as incidental
layout churn.

Both source and packaged selection are behavior-tested. Source compilation
uses the initial `std/go/_std` classpath. Package compilation uses only the
flattened `src` classpath and its generated `.cross.hx` files. Both paths run
`go test ./...`, execute the generated program, and reject machine-local path
leaks.

## Closed inventory

`test/canonical_std_layout_status.json` records the fail-closed inventory:

- 0 tracked `.cross.hx` files;
- 69 canonical override sources under `std/go/_std`;
- 104 ledger entries: 69 upstream overrides, 5 staged support modules, 25 typed
  `hxrt` bindings, and 5 public `go.*` facades;
- 279 source-to-package manifest entries in the Haxelib ZIP, including the
  native-policy types, generated-output boundary, and contract documentation
  added after closeout;
- 280 ZIP members after adding the embedded package manifest itself.

The closeout contract compares every canonical override path with the
`upstream_std_override` ledger set. A new, removed, or reclassified file must
update the ledger and the declared inventory in the same reviewed change.

## Isolated installed-package proof

`scripts/ci/run-isolated-haxelib-smoke.py` exercises the consumer boundary:

1. build the real deterministic package ZIP;
2. create a temporary local Haxelib repository and temporary Go caches while
   reusing the already-installed Haxe toolchain;
3. install only that ZIP, without repository development links;
4. compile `stdlib/stringtools_cross_std_basic` with `-lib reflaxe.go` and no
   checkout classpath;
5. run `go test ./...`, run the generated app, compare exact stdout, and scan
   generated files for checkout or temporary paths.

The command emits only relative names, counts, and pass/fail state as JSON.
Failures are classified by package, install, compile, Go test, execution,
stdout, or path-scan stage so a package-layout regression is not mistaken for a
compiler failure.

Run the focused proof with:

```bash
npm run test:canonical-std-closeout
```

The later release-candidate matrix expands this smoke to both product profiles
and more interop/runtime surfaces. This closeout test is deliberately the
smallest real installed-package proof needed to show that the `_std` migration
does not depend on the source checkout.

## Full closeout gates

Run all migration evidence with:

```bash
npm run test:stdlib:governance
npm run test:canonical-std-layout
npm run test:canonical-std-closeout
npm test
npm run test:semantic-diff
npm run test:stdlib-sweep:go-test
npm run test:examples
python3 test/test_raw_injection_hygiene_contract.py
```

The full snapshot, semantic-differential, upstream-stdlib, and example gates
remain authoritative for behavior. The focused closeout contract adds the
missing installed-package and inventory evidence; it does not replace those
broader suites.
