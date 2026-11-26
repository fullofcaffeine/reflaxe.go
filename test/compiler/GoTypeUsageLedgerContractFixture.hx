/**
	First fixture module for `GoTypeUsageLedgerContract`.

	Why: The contract compares how the type-usage ledger names member targets,
	member names, and native imports with the plain reads it replaced, so the
	fixture must reach every kind of field access and every kind of declaration
	that can carry `@:go.import`.

	What: A native extern class, enum, abstract, and typedef with `@:go.import`
	(the enum's value has spaces the ledger trims), a class whose first two
	`@:go.import` entries are ignored (empty, not a string), and private
	declarations whose names and kinds `GoTypeUsageLedgerContractOtherFixture`
	reuses with other import paths. `exercise(...)` constructs, reads instance
	and static fields, takes closures of a class method and of a structure
	method, selects enum constructors, reads a dynamic field, calls through
	parentheses and a cast, and calls an abstract method.

	How: The contract loads this module and the other fixture module with
	`Context.getModule(...)`; nothing here is compiled to Go.
**/
@:go.import("example.com/native/box")
extern class GoTypeUsageLedgerContractNativeBox {
	function new();
	static function make():GoTypeUsageLedgerContractNativeBox;
	function value():Int;
}

@:go.import(" example.com/native/choice ")
enum GoTypeUsageLedgerContractNativeChoice {
	Plain;
	Tagged(value:Int);
}

@:go.import("example.com/native/count")
abstract GoTypeUsageLedgerContractNativeCount(Int) from Int to Int {
	public function doubled():Int
		return this * 2;
}

@:go.import("example.com/native/shape")
typedef GoTypeUsageLedgerContractNativeShape = {
	final x:Int;
	function describe():String;
}

@:go.import("")
@:go.import(1)
@:go.import("example.com/native/late")
class GoTypeUsageLedgerContractLateImport {
	public static function ping():Int
		return 1;
}

@:go.import("example.com/first/hidden")
private enum Hidden {
	HiddenA;
}

@:go.import("example.com/first/score")
private abstract Score(Int) from Int to Int {}

class GoTypeUsageLedgerContractFixture {
	public var box:GoTypeUsageLedgerContractNativeBox;
	public var count:GoTypeUsageLedgerContractNativeCount;
	public var shape:GoTypeUsageLedgerContractNativeShape;
	public var hidden:Hidden;
	public var score:Score;

	public function new() {
		box = new GoTypeUsageLedgerContractNativeBox();
		count = 3;
		shape = {x: 1, describe: () -> "shape"};
		hidden = HiddenA;
		score = 2;
	}

	public function exercise(dynamicValue:Dynamic, items:Array<Int>):Int {
		final made = GoTypeUsageLedgerContractNativeBox.make();
		final read = made.value();
		final bound = made.value;
		final describe = shape.describe;
		final choice = GoTypeUsageLedgerContractNativeChoice.Tagged(read);
		final plain = GoTypeUsageLedgerContractNativeChoice.Plain;
		final dynamicField = dynamicValue.anything;
		final pushed = items.push(read);
		final viaCast = (cast made : GoTypeUsageLedgerContractNativeBox).value() + (made.value)() + (cast made.value : () -> Int)();
		final late = GoTypeUsageLedgerContractLateImport.ping();
		final doubled = count.doubled();
		final scored:Score = score;
		return read + bound() + shape.x + describe().length + pushed + viaCast + late + doubled + scored + (dynamicField == null ? 0 : 1)
			+ (choice == plain ? 1 : 0) + (hidden == HiddenA ? 1 : 0);
	}
}
