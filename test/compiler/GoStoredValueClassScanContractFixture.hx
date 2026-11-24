/** A base class with an instance layout and no superclass. */
class GoStoredValueClassScanContractFixture {
	public var id:Int = 0;

	public function new() {}
}

/** A subclass: a stored-value test of it can queue and mark its superclass. */
class StoredScanChild extends GoStoredValueClassScanContractFixture {
	public var label:String = "";
}

/** A source interface, which recovery matches as the named type `N`. */
interface StoredScanShape {
	function area():Int;
}

/** A class with no superclass whose name no recovery target in the contract uses. */
class StoredScanPlain {
	public var flag:Bool = false;

	public function new() {}
}
