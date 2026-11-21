enum Kind {
	Plain;
	Count(n:Int);
}

class Main {
	static function mayThrow(fail:Bool):Int {
		if (fail)
			throw "bad";
		return 7;
	}

	static function guardedCase(kind:Kind, early:Bool):String {
		final picked = switch kind {
			case Plain if (early): return "early";
			case Plain: "plain";
			case Count(n): "count" + n;
		};
		return "picked:" + picked;
	}

	static function catchReturn(fail:Bool):Int {
		final value = try mayThrow(fail) catch (e:Dynamic) return -1;
		return value * 10;
	}

	static function tryBodyReturn(early:Bool):Int {
		final value = try {
			if (early)
				return 5;
			mayThrow(false);
		} catch (e:Dynamic)
			0;
		return value + 100;
	}

	static function ifReturn(value:Int):String {
		final half = if (value < 0) return "negative" else value * 2;
		return "half:" + half;
	}

	static function loopExits():String {
		final out = new Array<Int>();
		for (i in 0...6) {
			final doubled = if (i == 1) continue
			else
				i * 2;
			final kept = switch i {
				case 4: break;
				case _: doubled + 1;
			};
			out.push(kept);
		}
		return out.join(",");
	}

	static function nestedValueTry(fail:Bool):String {
		final outer = try {
			final inner = try mayThrow(fail) catch (e:Dynamic) return "inner-catch";
			inner * 2;
		} catch (e:Dynamic)
			-1;
		return "outer:" + outer;
	}

	static function nestedTryStatements(mode:Int):String {
		try {
			try {
				if (mode == 0)
					return "inner";
				if (mode == 1)
					throw "boom";
			} catch (e:String) {
				return "inner-catch";
			}
		} catch (e:Dynamic) {
			return "outer-catch";
		}
		return "tail";
	}

	static function main():Void {
		Sys.println(guardedCase(Plain, true));
		Sys.println(guardedCase(Plain, false));
		Sys.println(guardedCase(Count(2), true));
		Sys.println(catchReturn(false));
		Sys.println(catchReturn(true));
		Sys.println(tryBodyReturn(true));
		Sys.println(tryBodyReturn(false));
		Sys.println(ifReturn(-1));
		Sys.println(ifReturn(4));
		Sys.println(loopExits());
		Sys.println(nestedValueTry(true));
		Sys.println(nestedValueTry(false));
		Sys.println(nestedTryStatements(0));
		Sys.println(nestedTryStatements(1));
		Sys.println(nestedTryStatements(2));
	}
}
