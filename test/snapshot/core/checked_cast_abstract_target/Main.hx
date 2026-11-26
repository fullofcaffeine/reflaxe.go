abstract Integer(Int) from Int to Int {
	/** Haxe's tests/unit Issue2957: the `Null<Int>` holding 5 passes the cast. **/
	@:from static public function fromString(s:String):Integer {
		var n:Null<Int> = Std.parseInt(s);
		return cast(n, Integer);
	}
}

abstract Real(Float) from Float to Float {}

class Animal {
	public function new() {}

	public function sound():String
		return "animal";
}

class Dog extends Animal {
	override public function sound():String
		return "dog";
}

abstract Pet(Animal) from Animal to Animal {}

enum abstract Color(Int) to Int {
	var Red = 1;
}

/**
	A checked cast to an abstract tests the value against the abstract's
	underlying run-time type, which Haxe's typer records as the cast's module
	type: Int for `Integer` and `Color`, Float for `Real`, the class Animal for
	`Pet`. The cast still throws Haxe's catchable "Class cast error" for a value
	of another type.
**/
class Main {
	static function main():Void {
		var n:Integer = "5";
		Sys.println("Integer from String: " + ((n : Int) + 1));
		var half:Dynamic = 2.5;
		Sys.println("Dynamic as Real: " + ((cast(half, Real) : Float) * 2));
		var dog:Dynamic = new Dog();
		Sys.println("Dynamic as Pet: " + (cast(dog, Pet) : Animal).sound());
		var one:Dynamic = 1;
		Sys.println("Dynamic as Color: " + (cast(one, Color) == Red));
		var text:Dynamic = "text";
		try {
			cast(text, Integer);
			Sys.println("not reached");
		} catch (message:String) {
			Sys.println("caught: " + message);
		}
	}
}
