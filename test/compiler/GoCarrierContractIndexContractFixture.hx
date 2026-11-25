/** A class that implements an interface with optional arguments only through a parent interface. */
class GoCarrierContractIndexContractFixture implements CarrierChildContract {
	public function new() {}

	public function get(?index:Int):Int
		return index == null ? 0 : index;

	public function plain(index:Int):Int
		return index;

	public function extra(?flag:Bool):Bool
		return flag == true;
}

/** Declares `get` with an optional argument and `plain` without one. */
interface CarrierOptionalContract {
	function get(?index:Int):Int;
	function plain(index:Int):Int;
}

/** Reaches `CarrierOptionalContract` as a parent interface and adds an optional `extra`. */
interface CarrierChildContract extends CarrierOptionalContract {
	function extra(?flag:Bool):Bool;
}

/** A second root interface with an optional `other`. */
interface CarrierOtherContract {
	function other(?text:String):String;
}

/** Has the fixture's contract only through its superclass, and adds `CarrierOtherContract`. */
class CarrierSubContract extends GoCarrierContractIndexContractFixture implements CarrierOtherContract {
	public function other(?text:String):String
		return text == null ? "" : text;
}

/** Declares an optional-argument method without any interface, so it has no contract. */
class CarrierLooseContract {
	public function new() {}

	public function get(?index:Int):Int
		return index == null ? 1 : index;
}
