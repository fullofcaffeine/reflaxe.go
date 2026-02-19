class Main {
	static function risky(v:Int):Int {
		if (v == 0) {
			throw "bad";
		}
		return v + 1;
	}

	static function main() {
		var a = try risky(0) catch (e:Dynamic) 11;
		var b = try risky(4) catch (e:Dynamic) 11;
		Sys.println(Std.string(a));
		Sys.println(Std.string(b));
	}
}
