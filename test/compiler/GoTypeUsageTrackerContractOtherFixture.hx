/**
	Second fixture module for `GoTypeUsageTrackerContract`.

	Why: Declarations of two modules can print alike. Private enums, typedefs,
	and abstracts of the same name get module-specific paths, and the local
	type parameter `pick.V` prints the same in both modules although Reflaxe
	gives each module's `V` its own id.

	What: Private `Hidden`, `Score`, and `Choice` that reuse the first fixture's
	names, a local generic function `pick<V>`, and two inline methods. Haxe
	inlines them into the first fixture's static `local()`, which brings this
	module's `pick<V>` and private declarations into the first fixture's map
	next to its own.

	How: The contract tracks both modules with one tracker. A declaration id
	that one module stored must not answer for the other module's declaration,
	and the first fixture's map must keep both modules' declarations apart.
**/
private typedef Hidden = {
	final label:String;
}

private abstract Score(String) from String {}

private enum Choice {
	Only;
}

class GoTypeUsageTrackerContractOtherFixture {
	public var hidden:Hidden;
	public var score:Score;
	public var choice:Choice;

	public function new() {
		hidden = {label: "label"};
		score = "score";
		choice = Only;
	}

	public static inline function picked():Int {
		function pick<V>(value:V):V
			return value;
		return pick(1);
	}

	public static inline function hiddenLength():Int {
		final hidden:Hidden = {label: "inline"};
		final score:Score = "inline";
		final choice:Choice = Only;
		return hidden.label.length + (choice == Only && score != null ? 1 : 0);
	}

	public static function local():String {
		function pick<V>(value:V):V
			return value;
		return pick("text");
	}
}
