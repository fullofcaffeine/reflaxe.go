class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (_:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		var source:Array<Int> = [3, 1, 4, 1];
		var ro:haxe.ds.ReadOnlyArray<Int> = source;
		safe("ro.length", function() return Std.string(ro.length));
		safe("ro.index", function() return Std.string(ro[2]));
		safe("ro.sum", function() {
			var total = 0;
			for (i in 0...ro.length) {
				total += ro[i];
			}
			return Std.string(total);
		});
		source[0] = 9;
		safe("ro.alias", function() return Std.string(ro[0]));
	}
}
