class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=" + fn());
		} catch (error:Dynamic) {
			Sys.println(label + "=!" + Std.string(error));
		}
	}

	static function main() {
		safe("fromCharCode", function() {
			return haxe.Ucs2.fromCharCode('A'.code).toNativeString();
		});
	}
}
