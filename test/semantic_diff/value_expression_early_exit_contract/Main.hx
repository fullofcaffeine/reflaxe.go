enum Kind {
	Plain;
	Count(n:Int);
}

/**
	Early exits from value-producing `switch`, `if` and `try` expressions.

	Each exit must leave the enclosing function or loop at the point where it
	runs, exactly as `--interp` does, instead of producing a value.
**/
class Main {
	static final log = new Array<String>();

	static function mayThrow(fail:Bool):Int {
		if (fail)
			throw "bad";
		return 7;
	}

	static function guardedCase(kind:Kind, early:Bool):String {
		final picked = switch kind {
			case Plain if (early): return "early-plain";
			case Count(n) if (n < 0): return "early-negative";
			case Plain: "plain";
			case Count(n): "count" + n;
		};
		return "picked:" + picked;
	}

	static function guardedInt(value:Int, early:Bool):Int {
		final doubled = switch value {
			case 1 if (early): return -100;
			case _: value * 2;
		};
		return doubled + 1;
	}

	static function unguardedCase(value:Int):String {
		final name = switch value {
			case 0: return "zero";
			case 1: "one";
			case _: "many";
		};
		return "name:" + name;
	}

	static function catchReturn(fail:Bool):Int {
		final value = try mayThrow(fail) catch (e:Dynamic) return -1;
		return value * 10;
	}

	static function tryBodyReturn(early:Bool, fail:Bool):Int {
		final value = try {
			if (early)
				return 5;
			mayThrow(fail);
		} catch (e:Dynamic)
			0;
		return value + 100;
	}

	static function ifReturn(value:Int):String {
		final half = if (value < 0) return "negative" else value * 2;
		return "half:" + half;
	}

	static function voidExits(value:Int, early:Bool):Void {
		final a = switch value {
			case 0 if (early): return;
			case _: value + 1;
		};
		final b = if (a > 3) return
		else
			a;
		final c = try {
			if (b == 2)
				throw "two";
			b;
		} catch (e:Dynamic)
			return;
		log.push("void:" + a + ":" + b + ":" + c);
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
		var j = 0;
		while (j < 6) {
			j++;
			final picked = switch j {
				case 2 if (out.length > 0): continue;
				case 5: break;
				case _: j;
			};
			out.push(picked * 100);
		}
		return out.join(",");
	}

	static function throwExits(value:Int):String {
		final a = if (value < 0) throw "negative" else value;
		final b = switch a {
			case 0: throw "zero";
			case 1 if (value > 0): throw "one";
			case _: a * 3;
		};
		final c = try {
			if (b > 10)
				throw "big";
			b;
		} catch (e:String)
			throw "rethrown:" + e;
		return "ok:" + c;
	}

	static function nestedValues(mode:Int, flag:Bool):String {
		final value = switch mode {
			case 0: if (flag) return "if-return" else 1;
			case 1: try mayThrow(flag) catch (e:Dynamic) return "catch-return";
			case 2:
				final inner = try {
					switch flag {
						case true: return "switch-in-try";
						case false: 2;
					}
				} catch (e:Dynamic)
					0;
				inner + 1;
			case 3:
				final outer = try {
					final innerValue = try mayThrow(flag) catch (e:Dynamic) return "try-in-try";
					innerValue * 2;
				} catch (e:Dynamic)
					-1;
				outer;
			case _: mode;
		};
		return "value:" + value;
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
			if (mode == 2)
				return "outer";
		} catch (e:Dynamic) {
			return "outer-catch";
		}
		return "tail";
	}

	static function main():Void {
		Sys.println(guardedCase(Plain, true));
		Sys.println(guardedCase(Plain, false));
		Sys.println(guardedCase(Count(-3), true));
		Sys.println(guardedCase(Count(4), true));
		Sys.println(guardedInt(1, true));
		Sys.println(guardedInt(1, false));
		Sys.println(guardedInt(5, true));
		Sys.println(unguardedCase(0));
		Sys.println(unguardedCase(1));
		Sys.println(unguardedCase(9));
		Sys.println(catchReturn(false));
		Sys.println(catchReturn(true));
		Sys.println(tryBodyReturn(true, false));
		Sys.println(tryBodyReturn(false, false));
		Sys.println(tryBodyReturn(false, true));
		Sys.println(ifReturn(-1));
		Sys.println(ifReturn(4));
		voidExits(0, true);
		voidExits(0, false);
		voidExits(5, false);
		voidExits(1, false);
		voidExits(2, false);
		Sys.println(log.join(","));
		Sys.println(loopExits());
		for (value in [-1, 0, 1, 2, 5]) {
			try {
				Sys.println(throwExits(value));
			} catch (e:String) {
				Sys.println("caught:" + e);
			}
		}
		for (mode in 0...5) {
			Sys.println(nestedValues(mode, true));
			Sys.println(nestedValues(mode, false));
		}
		for (mode in 0...4)
			Sys.println(nestedTryStatements(mode));
		final lambda = function(value:Int):String {
			final picked = switch value {
				case 0 if (value == 0): return "lambda-early";
				case _: value;
			};
			return "lambda:" + picked;
		};
		Sys.println(lambda(0));
		Sys.println(lambda(2));
		final arrow = (value:Int) -> {
			final checked = try {
				if (value > 1)
					throw "big";
				value;
			} catch (e:Dynamic)
				return -1;
			checked * 2;
		};
		Sys.println(arrow(1));
		Sys.println(arrow(5));
	}
}
