/** Pins first-class closure transport independently of method dispatch. */
class Main {
	static function counter():Int->Int {
		var total = 0;
		return delta -> total += delta;
	}

	static function reuse<T>(value:T):T {
		return value;
	}

	static function main() {
		final first = counter();
		final alias = first;
		final other = counter();
		Sys.println("identity=" + Reflect.compareMethods(first, alias) + ":" + Reflect.compareMethods(first, other));
		final stored = [first, alias, other];
		Sys.println("transport=" + Reflect.compareMethods(stored[0], reuse(stored[1])) + ":" + Reflect.compareMethods(stored[0], stored[2]));
		Sys.println("capture=" + first(2) + ":" + alias(3) + ":" + other(1));
		final record = {callback: first};
		final extracted = record.callback;
		final boxed:Dynamic = extracted;
		final restored:Int->Int = cast boxed;
		Sys.println("record=" + record.callback(0) + ":" + Reflect.compareMethods(first, restored));
		final missing:Null<Int->Int> = null;
		var evaluated = 0;
		var caught = false;
		try {
			missing(++evaluated);
		} catch (_:Dynamic) {
			caught = true;
		}
		Sys.println("null=" + caught + ":" + evaluated);
		// A callable has no source-level String member, even if its runtime carrier does.
		Sys.println("opaque=" + (Reflect.field(first, "String") == null) + ":" + Reflect.hasField(first, "String"));
	}
}
