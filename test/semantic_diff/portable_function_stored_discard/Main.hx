class Main {
	static function discard(callback:Int->Int):Int->Void {
		return callback;
	}

	static function main() {
		final original = function(value:Int):Int {
			Sys.println(value);
			return value;
		};
		final returned = discard(original);
		final stored:Int->Void = original;
		returned(3);
		stored(5);
		Sys.println(Reflect.compareMethods(original, returned));
		Sys.println(Reflect.compareMethods(original, stored));
	}
}
