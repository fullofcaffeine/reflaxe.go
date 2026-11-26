/**
	A checked cast `cast(value, A)` to an abstract `A` tests the value against
	the run-time type that Haxe's typer records for the cast. An abstract has
	no run-time type of its own, so the typer follows it to its underlying
	type: `cast(n, Integer)` over Int tests for an Int, a cast to an abstract
	over a class tests for that class, and an enum abstract tests its
	underlying Int or String without checking that the value is one of its
	constants. Only an abstract marked `@:runtimeValue` is tested as itself;
	`--interp` lets no value but null through such a `@:coreType` abstract.

	The cases cover abstracts over Int, Float, String, Bool, a class, an enum
	and another abstract, enum abstracts over Int and String, `Null<T>` and
	`Dynamic` operands, null operands, casts used only as statements and
	values that must keep throwing the catchable "Class cast error". Haxe's
	own tests/unit Issue2957 is the first case: an `@:from` conversion that
	returns `cast(Std.parseInt(s), Integer)`.
**/
abstract Integer(Int) from Int to Int {
	@:from static public function fromString(s:String):Integer {
		var n:Null<Int> = Std.parseInt(s);
		return cast(n, Integer);
	}
}

abstract Real(Float) from Float to Float {}
abstract Name(String) from String to String {}
abstract Flag(Bool) from Bool to Bool {}
abstract Count(Integer) from Integer to Integer {}

class Animal {
	public function new() {}

	public function sound():String
		return "animal";
}

class Dog extends Animal {
	override public function sound():String
		return "dog";
}

class Stone {
	public function new() {}
}

abstract Pet(Animal) from Animal to Animal {
	public function sound():String
		return this.sound();
}

enum Signal {
	Go;
	Stop;
}

abstract Light(Signal) from Signal to Signal {}

enum abstract Color(Int) to Int {
	var Red = 1;
	var Green = 2;
}

enum abstract Kind(String) to String {
	var Small = "small";
	var Large = "large";
}

@:coreType @:runtimeValue abstract Opaque {}

class Main {
	static function attempt(label:String, run:() -> String) {
		try {
			Sys.println(label + ": " + run());
		} catch (e:Dynamic) {
			Sys.println(label + ": threw " + Std.string(e));
		}
	}

	static function main() {
		var fromText:Integer = "5";
		Sys.println("Issue2957 Integer from String: " + ((fromText : Int) + 1));

		var five:Null<Int> = Std.parseInt("5");
		var noInt:Null<Int> = null;
		var seven:Dynamic = 7;
		var half:Dynamic = 1.5;
		var text:Dynamic = "text";
		var yes:Dynamic = true;
		var nothing:Dynamic = null;
		var dog:Dynamic = new Dog();
		var stone:Dynamic = new Stone();

		attempt("Null<Int> 5 as Integer", () -> "" + (cast(five, Integer) : Int));
		attempt("Int 5 as Integer", () -> "" + ((cast(5, Integer) : Int) + 1));
		attempt("Dynamic 7 as Integer", () -> "" + ((cast(seven, Integer) : Int) + 1));
		attempt("null Null<Int> as Integer", () -> {
			var passed:Null<Integer> = cast(noInt, Integer);
			"passed";
		});
		attempt("null Dynamic as Integer", () -> {
			cast(nothing, Integer);
			"passed";
		});
		attempt("Dynamic text as Integer", () -> "" + (cast(text, Integer) : Int));
		attempt("Dynamic 1.5 as Integer", () -> "" + (cast(half, Integer) : Int));

		var quarter:Null<Float> = 2.25;
		attempt("Null<Float> as Real", () -> "" + ((cast(quarter, Real) : Float) * 2));
		attempt("Dynamic 7 as Real", () -> Std.string(cast(seven, Real)));
		attempt("Dynamic 1.5 as Real", () -> "" + (cast(half, Real) : Float));
		attempt("Dynamic text as Real", () -> "" + (cast(text, Real) : Float));

		var maybeName:Null<String> = "ada";
		attempt("Null<String> as Name", () -> (cast(maybeName, Name) : String).toUpperCase());
		attempt("Dynamic text as Name", () -> (cast(text, Name) : String).length + "");
		attempt("Dynamic 7 as Name", () -> (cast(seven, Name) : String));
		attempt("null Dynamic as Name", () -> "" + (cast(nothing, Name) == null));

		var maybeFlag:Null<Bool> = true;
		attempt("Null<Bool> as Flag", () -> "" + !(cast(maybeFlag, Flag) : Bool));
		attempt("Dynamic true as Flag", () -> "" + (cast(yes, Flag) : Bool));
		attempt("Dynamic 7 as Flag", () -> "" + (cast(seven, Flag) : Bool));

		attempt("Dynamic 7 as Count", () -> "" + ((cast(seven, Count) : Integer) : Int));
		attempt("Dynamic text as Count", () -> "" + ((cast(text, Count) : Integer) : Int));

		attempt("Dynamic Dog as Pet", () -> cast(dog, Pet).sound());
		attempt("Dynamic Stone as Pet", () -> cast(stone, Pet).sound());
		attempt("null Dynamic as Pet", () -> "" + (cast(nothing, Pet) == null));

		var go:Dynamic = Go;
		attempt("Dynamic Go as Light", () -> "" + ((cast(go, Light) : Signal) == Go));
		attempt("Dynamic 7 as Light", () -> "" + ((cast(seven, Light) : Signal) == Go));

		var one:Dynamic = 1;
		var many:Dynamic = 99;
		attempt("Dynamic 1 as Color", () -> "" + (cast(one, Color) == Red));
		attempt("Dynamic 99 as Color", () -> "" + (cast(many, Color) : Int));
		attempt("Dynamic text as Color", () -> "" + (cast(text, Color) : Int));
		var small:Dynamic = "small";
		attempt("Dynamic small as Kind", () -> "" + (cast(small, Kind) == Small));
		attempt("Dynamic 7 as Kind", () -> "" + (cast(seven, Kind) : String));

		attempt("Dynamic 7 as Opaque", () -> {
			cast(seven, Opaque);
			"passed";
		});
		attempt("null Dynamic as Opaque", () -> {
			cast(nothing, Opaque);
			"passed";
		});

		try {
			cast(text, Integer);
			Sys.println("statement text as Integer: not reached");
		} catch (e:haxe.Exception) {
			Sys.println("statement text as Integer: " + e.message);
		}
		cast(five, Integer);
		Sys.println("statement Null<Int> as Integer: passed");
	}
}
