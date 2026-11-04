typedef Narrow = {var n:Int;}
typedef Wide = {var n:Int; var m:Int;}
typedef Outer<T> = {var callbacks:Array<Void->T>;}

class C {
	public var n:Int = 1;
	public var m:Int = 2;

	public function new() {}
}

class Main {
	static function read<T>(value:Outer<T>):T
		return value.callbacks[0]();

	static function source(x:Int = 7):Int
		return x + 2;

	static function main() {
		var narrow:Narrow = new C();
		var erased:Dynamic = narrow;
		var wide:Wide = erased;
		Sys.println("wide=" + wide.m);
		Sys.println("array=" + read({callbacks: [() -> 3]}));
		var projected:(?x:Int) -> Void = source;
		projected();
		var callable:Dynamic = projected;
		Sys.println("dynamic=" + callable());
		Sys.println("reflect=" + Reflect.callMethod(null, projected, []));
		Sys.println("explicitNull=" + Reflect.callMethod(null, projected, [null]));
		var variadic:Dynamic = Reflect.makeVarArgs(args -> args.length);
		var noArgs:Void->Int = variadic;
		Sys.println("variadic=" + noArgs());
	}
}
