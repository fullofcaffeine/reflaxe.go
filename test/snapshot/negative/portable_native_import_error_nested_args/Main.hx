/**
	Why
	The native authority gate once kept the printed form of each visited type as
	its visited set. Prints are cut short after two levels of nesting, so
	`Array<Array<Int>>` and `Array<Array<go.NativeAny>>` print the same, and the
	gate skipped the second argument below and never reported this module.

	What
	A portable interface whose only `go.*` type is nested in its second argument
	type.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so the gate must
	fail the build for module `Main`.
**/
interface Taker {
	function take(plain:Array<Array<Int>>, native:Array<Array<go.NativeAny>>):Void;
}

class Main {
	static function main() {}
}
