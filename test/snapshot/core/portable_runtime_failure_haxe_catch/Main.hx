class Main {
	static function main():Void {
		var value:Dynamic = null;
		try {
			var typed:Int = value;
			Sys.println(typed + 1);
		} catch (_:Dynamic) {
			Sys.println("caught-portable-runtime-failure");
		}
	}
}
