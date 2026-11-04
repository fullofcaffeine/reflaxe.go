private class Holder {
	public var callback:Int->Void;

	public function new(callback:Int->Void) {
		this.callback = callback;
	}
}

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
		final holder = new Holder(original);
		holder.callback(7);
		holder.callback = original;
		holder.callback(8);
		final choose = true;
		final branch:Int->Void = if (choose) original else returned;
		branch(9);
		Sys.println(Reflect.compareMethods(original, holder.callback));
		Sys.println(Reflect.compareMethods(original, branch));
		final absent:Int->Int = null;
		final absentDiscarded:Int->Void = absent;
		Sys.println(absentDiscarded == null);
	}
}
