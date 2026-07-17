# Lane Semantic Diff Suite

This suite verifies that native-boundary metadata in portable builds does not
change runtime semantics for source whose executable behavior remains portable.
The interpreter is the reference, so target-native `go.*` APIs do not belong in
this suite.

Lane coverage includes:

- portable iterator/list/map/string behavior inside `@:goMetal` modules (`portable_surfaces_lane_invariance_contract`)

Concrete `go.*` behavior under `reflaxe_go_auto=auto_strict` is covered by
`test/snapshot/go_native/native_boundary_collections_strict`, where generated Go
is compiled and run against `expected.stdout` without inventing an interpreter
implementation.

Runner:

- `python3 test/run-semantic-diff.py --suite lanes`
- `npm run test:semantic-diff:lanes`
