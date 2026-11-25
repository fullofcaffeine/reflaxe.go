/**
	An interface whose only `go.*` type is in the second of two anonymous
	argument types.

	In eval every anonymous structure prints as `<anonymous>`. The native
	authority gate once kept printed types as its visited set, so it skipped the
	second argument and never reported this module. The gate must report it.
**/
interface GoNativeTypeReachContractPrintFixture {
	function take(first:{count:Int}, second:{values:go.NativeSlice<Bool>}):Void;
}

/**
	An interface whose only `go.*` type is nested in its second argument type.

	Eval cuts prints short after two levels of nesting, so `Array<Array<Int>>`
	and `Array<Array<go.NativeAny>>` print the same. The gate's old printed
	visited set skipped the second argument; the gate must report it.
**/
interface GoNativeTypeReachContractNestedPrint {
	function take(plain:Array<Array<Int>>, native:Array<Array<go.NativeAny>>):Void;
}

/** A portable class next to the interfaces, with no `go.*` type at all. */
class GoNativeTypeReachContractPrintPortable {
	public var count:Int = 0;

	public function new() {}
}
