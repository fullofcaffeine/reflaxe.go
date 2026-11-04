class Main {
	static function firstOrigin():haxe.Exception
		return new haxe.Exception("first");

	static function secondOrigin():haxe.Exception
		return new haxe.Exception("second");

	static function readLater(error:haxe.Exception):String
		return haxe.CallStack.toString(error.stack);

	static function main():Void {
		var first = firstOrigin();
		var second = secondOrigin();
		var initial = haxe.CallStack.toString(first.stack);
		var later = readLater(first);
		Sys.println("captured=" + Std.string(initial.indexOf("firstOrigin") >= 0));
		Sys.println("stable=" + Std.string(initial == later));
		Sys.println("separate=" + Std.string(readLater(second).indexOf("secondOrigin") >= 0));
		try {
			throw first;
		} catch (caught:haxe.Exception) {
			Sys.println("rethrow-same=" + Std.string(caught == first && readLater(caught) == initial));
		}
	}
}
