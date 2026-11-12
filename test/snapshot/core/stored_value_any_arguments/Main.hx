class Defaults {
	public function new() {}

	public function increment(value:Int = 5):Int {
		return value + 1;
	}

	public function enabled(value:Bool = true):Bool {
		return value;
	}
}

class Main {
	static function total(values:Array<Int>):Int {
		return values[0] + values[1];
	}

	static function main():Void {
		final defaults = new Defaults();
		final increment = defaults.increment;
		final enabled = defaults.enabled;
		Sys.println(increment());
		Sys.println(increment(8));
		Sys.println(enabled());
		Sys.println(enabled(false));
		// Any is the explicit erased-storage boundary under test. Recovering
		// its concrete array must preserve the array rather than rebox it.
		final erased:Any = [4, 7];
		Sys.println(total(erased));
	}
}
