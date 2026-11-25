class Base {
	public function new() {}
}

class Mid extends Base {}

class Other {
	public function new() {}
}

/**
	A checked cast tests its value and throws Haxe's catchable
	"Class cast error" for a value of another class, instead of a raw Go
	type-assertion panic. An upcast is proven by the operand's static type and
	keeps the plain cast lowering without a test.
**/
class Main {
	static function main():Void {
		var other:Dynamic = new Other();
		try {
			cast(other, Mid);
			Sys.println("not reached");
		} catch (message:String) {
			Sys.println("caught: " + message);
		}
		var mid = new Mid();
		var base:Base = cast(mid, Base);
		Sys.println("upcast: " + (base == mid));
	}
}
