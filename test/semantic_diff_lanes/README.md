# Lane Semantic Diff Suite

This suite verifies that `@:goMetal` lane enforcement in portable builds (`reflaxe_go_auto=auto_strict`) does not change runtime semantics for lane-clean programs.

Lane coverage includes:

- typed `go.*` lane-clean behavior (`collections_lane_clean_contract`, `result_lane_clean_contract`)
- portable iterator/list/map/string behavior inside `@:goMetal` modules (`portable_surfaces_lane_invariance_contract`)

Runner:

- `python3 test/run-semantic-diff.py --suite lanes`
- `npm run test:semantic-diff:lanes`
