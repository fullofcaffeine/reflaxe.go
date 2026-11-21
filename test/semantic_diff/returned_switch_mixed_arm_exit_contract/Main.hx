enum Shape {
	Dot;
	Line(length:Int);
	Box(width:Int, height:Int);
}

class Picker {
	final base:Int;

	public function new(base:Int) {
		this.base = base;
	}

	public function pick(flag:Bool):Int {
		return switch base {
			case 0: -1;
			case _: if (flag) return base + 10 else base + 20;
		};
	}
}

/**
	Returned and stored value `switch` expressions whose arm mixes a `return`
	with a produced value (haxe_go-ss2a).

	A `return` taken inside the arm must leave the enclosing function at that
	point, skipping the rest of the arm and everything after the switch, exactly
	as `--interp` does. The other side of the same arm still produces the switch
	value. Within each function a taken return and a produced value print
	different results, so the two cannot be confused.
**/
class Main {
	static final log = new Array<String>();

	static function mayThrow(fail:Bool):Int {
		if (fail)
			throw "bad";
		return 7;
	}

	static function note(value:Int):Int {
		log.push("note" + value);
		return value;
	}

	static function mixedArm(kind:Int, flag:Bool):Int {
		final value = 1;
		return switch kind {
			case 1: if (flag) return value + 10 else value + 20;
			case _: -1;
		};
	}

	static function nestedIf(kind:Int, a:Bool, b:Bool):Int {
		return switch kind {
			case 1:
				if (a) {
					if (b)
						return 1
					else
						2;
				} else 3;
			case _: 4;
		};
	}

	static function guarded(n:Int, flag:Bool):Int {
		return switch n {
			case n if (n > 5): if (flag) return n * 100 else -n;
			case _: n;
		};
	}

	static function elseIfChain(kind:Int, a:Int):Int {
		return switch kind {
			case 1: if (a == 0) return 10 else if (a == 1) 20 else if (a == 2) return 30 else 40;
			case _: 0;
		};
	}

	static function switchInArm(kind:Int, inner:Int):Int {
		return switch kind {
			case 1:
				switch inner {
					case 0: return 100;
					case 1: if (kind == 1) return 101 else 0;
					case _: inner * 2;
				}
			case _: -kind;
		};
	}

	static function tryInArm(kind:Int, fail:Bool):Int {
		return switch kind {
			case 1: try mayThrow(fail) catch (e:Dynamic) return -7;
			case 2:
				try {
					if (fail)
						return 200;
					mayThrow(false) + 1;
				} catch (e:Dynamic)
					-2;
			case _: 0;
		};
	}

	static function stringSubject(name:String, flag:Bool):String {
		return switch name {
			case "a": if (flag) return "A!" else "a";
			case "b" | "c": if (!flag) return name + "!" else name + "?";
			case _: "other";
		};
	}

	static function enumSubject(shape:Shape, flag:Bool):Int {
		return switch shape {
			case Dot: 0;
			case Line(length): if (length > 0) return length + 10 else length + 20;
			case Box(w, h) if (flag): if (w == h) return w * h else w + h;
			case Box(w, _): w;
		};
	}

	static function sideEffects(kind:Int, flag:Bool):Int {
		log.push("start" + kind);
		final picked = switch kind {
			case 1:
				log.push("arm");
				final local = kind * 3;
				if (flag) return note(local + 1) else note(local + 2);
			case _: 0;
		};
		log.push("after" + picked);
		return picked;
	}

	static function nullableResult(kind:Int, flag:Bool):Null<Int> {
		return switch kind {
			case 1: if (flag) return null else kind;
			case _: 9;
		};
	}

	static function returnedIf(a:Bool, b:Int):Int {
		return if (a) switch b {
			case 0: return 1;
			case _: if (b > 5) return b else b + 1;
		} else 33;
	}

	static function assigned(kind:Int, flag:Bool):Int {
		var x = 0;
		x = switch kind {
			case 1: if (flag) return 99 else 1;
			case _: 2;
		};
		return x + 1000;
	}

	static function compoundAssigned(kind:Int, flag:Bool):Int {
		var x = 100;
		x += switch kind {
			case 1: if (flag) return 5 else 6;
			case _: 0;
		};
		return x;
	}

	static function insideLoop(kind:Int, flag:Bool):Int {
		var total = 0;
		for (i in 0...3) {
			total += i;
			final step = switch kind {
				case 1: if (flag && i == 1) return total * 10 else i;
				case _: 0;
			};
			total += step;
		}
		return total;
	}

	static function generic<T>(value:T, alt:T, kind:Int, flag:Bool):T {
		return switch kind {
			case 1: if (flag) return alt else value;
			case _: value;
		};
	}

	static function voidArm(kind:Int, flag:Bool):Void {
		final picked = switch kind {
			case 1: if (flag) return
				else
					"one";
			case _: "other";
		};
		log.push("void:" + picked);
	}

	static function main():Void {
		Sys.println(mixedArm(1, true));
		Sys.println(mixedArm(1, false));
		Sys.println(mixedArm(2, true));
		for (a in [true, false])
			for (b in [true, false])
				Sys.println(nestedIf(1, a, b));
		Sys.println(nestedIf(2, true, true));
		Sys.println(guarded(7, true));
		Sys.println(guarded(7, false));
		Sys.println(guarded(3, true));
		for (a in 0...4)
			Sys.println(elseIfChain(1, a));
		Sys.println(elseIfChain(2, 0));
		for (inner in 0...3)
			Sys.println(switchInArm(1, inner));
		Sys.println(switchInArm(3, 0));
		Sys.println(tryInArm(1, false));
		Sys.println(tryInArm(1, true));
		Sys.println(tryInArm(2, true));
		Sys.println(tryInArm(2, false));
		Sys.println(tryInArm(3, false));
		for (name in ["a", "b", "d"]) {
			Sys.println(stringSubject(name, true));
			Sys.println(stringSubject(name, false));
		}
		for (shape in [Dot, Line(5), Line(-1), Box(3, 3), Box(3, 4)])
			Sys.println(enumSubject(shape, true));
		Sys.println(enumSubject(Box(3, 4), false));
		Sys.println(sideEffects(1, true));
		Sys.println(sideEffects(1, false));
		Sys.println(sideEffects(2, true));
		Sys.println(log.join(","));
		log.resize(0);
		Sys.println(nullableResult(1, true));
		Sys.println(nullableResult(1, false));
		Sys.println(nullableResult(2, false));
		Sys.println(returnedIf(true, 0));
		Sys.println(returnedIf(true, 7));
		Sys.println(returnedIf(true, 2));
		Sys.println(returnedIf(false, 0));
		Sys.println(assigned(1, true));
		Sys.println(assigned(1, false));
		Sys.println(assigned(2, false));
		Sys.println(compoundAssigned(1, true));
		Sys.println(compoundAssigned(1, false));
		Sys.println(insideLoop(1, true));
		Sys.println(insideLoop(1, false));
		Sys.println(insideLoop(2, true));
		Sys.println(generic("value", "alt", 1, true));
		Sys.println(generic("value", "alt", 1, false));
		Sys.println(generic(1, 2, 2, true));
		voidArm(1, true);
		voidArm(1, false);
		voidArm(2, true);
		Sys.println(log.join(","));
		Sys.println(new Picker(1).pick(true));
		Sys.println(new Picker(1).pick(false));
		Sys.println(new Picker(0).pick(true));
		final offset = 5;
		final local = function(kind:Int, flag:Bool):Int {
			return switch kind {
				case 1: if (flag) return offset + 10 else offset + 20;
				case _: -offset;
			};
		};
		Sys.println(local(1, true));
		Sys.println(local(1, false));
		Sys.println(local(2, true));
		final arrow = (name:String, flag:Bool) -> switch name {
			case "x": if (flag) return "X" + offset else "x" + offset;
			case _: name;
		};
		Sys.println(arrow("x", true));
		Sys.println(arrow("x", false));
		Sys.println(arrow("y", false));
		function named(shape:Shape):String {
			return switch shape {
				case Line(n): if (n == offset) return "match" else "line" + n;
				case _: "shape";
			};
		}
		Sys.println(named(Line(5)));
		Sys.println(named(Line(2)));
		Sys.println(named(Dot));
	}
}
