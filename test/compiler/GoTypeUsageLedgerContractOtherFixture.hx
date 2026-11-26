/**
	Second fixture module for `GoTypeUsageLedgerContract`.

	Why: The ledger keeps the path and native import of each enum and abstract
	by its type path. Private declarations with one name in two modules must
	keep their own module paths and import paths.

	What: A private enum `Hidden` and a private abstract `Score`, as in
	`GoTypeUsageLedgerContractFixture`, with other `@:go.import` paths, and a
	class that uses both.

	How: The contract loads this module with `Context.getModule(...)`.
**/
@:go.import("example.com/other/hidden")
private enum Hidden {
	HiddenB;
}

@:go.import("example.com/other/score")
private abstract Score(Int) from Int to Int {}

class GoTypeUsageLedgerContractOtherFixture {
	public static function total():Int {
		final hidden = HiddenB;
		final score:Score = 4;
		return (hidden == HiddenB ? 1 : 0) + score;
	}
}
