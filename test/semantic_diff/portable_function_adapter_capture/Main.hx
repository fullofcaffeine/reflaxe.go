/**
	Checks that generic callback adapters capture their source exactly once.
	Factory effects occur at argument evaluation, and later reassignment must
	not replace a callback that the callee already received.
**/
class Main {
	static var created = 0;

	static function mapper():Int->Int {
		created++;
		return value -> value + 1;
	}

	static function invokeTwice(callback:Int->Void):Void {
		callback(1);
		callback(2);
	}

	static function main() {
		final first = Lambda.map([1, 2], mapper());
		Sys.println("factory=" + created + ":" + first.join(","));
		var callback:Int->Int = null;
		callback = value -> {
			if (value == 1)
				callback = _ -> 99;
			return value + 1;
		};
		final original = callback;
		final second = Lambda.map([1, 2], callback);
		Sys.println("inline=" + second.join(",") + ":" + callback(3) + ":" + original(3));
		final calls:Array<Int> = [];
		callback = value -> {
			calls.push(value);
			callback = value -> {
				calls.push(99);
				return 99;
			};
			return value;
		};
		invokeTwice(callback);
		Sys.println("retained=" + calls.join(","));
	}
}
