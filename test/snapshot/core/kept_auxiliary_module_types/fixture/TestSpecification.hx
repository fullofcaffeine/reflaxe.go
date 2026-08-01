package fixture;

import fixture.Base;

class TestSpecification {}

@:keep class KeptAuxiliary {
	public function new() {}

	public function value():String {
		return "kept";
	}
}

class ClassWithToStringChild extends Base {}

class ClassWithToStringChild2 extends Base {
	public override function toString():String {
		return "ClassWithToStringChild2.toString()";
	}
}
