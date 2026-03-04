# Lane Semantic Diff Suite

This suite verifies that `@:goMetal` lane enforcement in portable builds (`reflaxe_go_auto=auto_strict`) does not change runtime semantics for lane-clean programs.

Runner:

- `python3 test/run-semantic-diff.py --suite lanes`
- `npm run test:semantic-diff:lanes`
