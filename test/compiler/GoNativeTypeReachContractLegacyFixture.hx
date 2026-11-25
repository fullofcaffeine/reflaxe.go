/**
	An interface whose only `go.*` type is in the second of two anonymous
	argument types.

	The native authority gate's walk keeps printed types as its visited set, and
	every anonymous structure prints the same, so the walk skips the second
	argument and has never reported this module. `GoNativeTypeReach` sees the
	`go.*` type; the gate must still not report it.
**/
interface GoNativeTypeReachContractLegacyFixture {
	function take(first:{count:Int}, second:{values:go.NativeSlice<Bool>}):Void;
}

/** A portable class next to the interface, with no `go.*` type at all. */
class GoNativeTypeReachContractLegacyPortable {
	public var count:Int = 0;

	public function new() {}
}
