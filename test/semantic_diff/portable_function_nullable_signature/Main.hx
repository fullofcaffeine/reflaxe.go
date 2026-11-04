class Main {
	static function inspect(value:Null<Int>->Bool):Bool {
		return value(null);
	}

	static function main() {
		final integer:Void->Null<Int> = function():Null<Int> return null;
		final decimal:Void->Null<Float> = function():Null<Float> return null;
		final flag:Void->Null<Bool> = function():Null<Bool> return null;
		Sys.println(integer() == null);
		Sys.println(decimal() == null);
		Sys.println(flag() == null);
		Sys.println(inspect(function(value:Null<Int>):Bool return value == null));
	}
}
