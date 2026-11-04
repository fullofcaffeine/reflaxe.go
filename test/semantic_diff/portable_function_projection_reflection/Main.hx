class Main {
	static var calls:Int = 0;

	static function source(x:Int = 7):Int {
		calls++;
		return x + 2;
	}

	static function higher(callback:Int->Int):Int
		return callback(3);

	static function main() {
		var projected:(?x:Int) -> Void = source;
		projected();
		Sys.println(Reflect.callMethod(null, projected, [null]));
		Sys.println(Reflect.callMethod(null, projected, [3]));
		Sys.println(calls);
		var literal:(?x:Int) -> Void = function(x:Int = 7):Int return x + 2;
		Sys.println(Reflect.callMethod(null, literal, []));
		var higherView:(Dynamic->Dynamic) -> Void = higher;
		Sys.println(Reflect.callMethod(null, higherView, [(x:Int) -> x + 2]));
		var dynamicSource:Dynamic = source;
		Sys.println(dynamicSource());
	}
}
