class Local {
	public var value:Int;

	public function new()
		value = 7;

	public static function read():Int {
		final bytes = new Local();
		final strings = new Local();
		return bytes.value + strings.value - 7;
	}

	public static function argument(bytes:Local):Int
		return bytes.value;

	public static function loops():Int {
		var sum = 0;
		for (bytes in [1, 2, 3])
			sum += bytes;
		return sum;
	}

	public static function catchValue():String {
		try {
			throw "caught";
		} catch (bytes:haxe.Exception) {
			return bytes.message;
		}
	}

	public static function closure():Int {
		final strings = 7;
		final fn = function() {
			final strings_local = 3;
			return strings * 10 + strings_local;
		};
		return fn();
	}

	public static function runtimeAlias():Int {
		final hxrt = new Local();
		return hxrt.value + "abc".length;
	}

	public static function suffix():Int {
		var total = 0;
		{
			final bytes = new Local();
			total += bytes.value;
		}
		{
			final bytes_local = new Local();
			total += bytes_local.value;
		}
		{
			final bytes_local = 1;
			total += bytes_local;
		}
		return total;
	}
}
