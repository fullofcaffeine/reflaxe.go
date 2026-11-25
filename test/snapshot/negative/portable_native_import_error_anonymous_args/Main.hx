/**
	Why
	The native authority gate once kept the printed form of each visited type as
	its visited set. Every anonymous structure prints the same, so the gate
	skipped the second anonymous argument below and never reported this module.

	What
	A portable interface whose only `go.*` type is in its second anonymous
	argument type.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so the gate must
	fail the build for module `Main`.
**/
interface Taker {
	function take(first:{count:Int}, second:{values:go.NativeSlice<Bool>}):Void;
}

class Main {
	static function main() {}
}
