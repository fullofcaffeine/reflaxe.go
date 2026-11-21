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
	A returned `switch` whose arm is a value `if` that returns on one side and
	produces the switch value on the other (haxe_go-ss2a).
**/
class Main {
	static function mayThrow(fail:Bool):Int {
		if (fail)
			throw "bad";
		return 7;
	}

	static function mixedArm(kind:Int, flag:Bool):Int {
		final value = 1;
		return switch kind {
			case 1: if (flag) return value + 10 else value + 20;
			case _: -1;
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

	static function enumSubject(shape:Shape, flag:Bool):Int {
		return switch shape {
			case Dot: 0;
			case Line(length): if (length > 0) return length + 10 else length + 20;
			case Box(w, h) if (flag): if (w == h) return w * h else w + h;
			case Box(w, _): w;
		};
	}

	static function main():Void {
		Sys.println(mixedArm(1, true));
		Sys.println(mixedArm(1, false));
		Sys.println(mixedArm(2, true));
		for (a in 0...4)
			Sys.println(elseIfChain(1, a));
		for (inner in 0...3)
			Sys.println(switchInArm(1, inner));
		Sys.println(switchInArm(3, 0));
		Sys.println(tryInArm(1, false));
		Sys.println(tryInArm(1, true));
		Sys.println(tryInArm(2, true));
		Sys.println(tryInArm(2, false));
		for (shape in [Line(5), Line(-1), Box(3, 3), Box(3, 4)])
			Sys.println(enumSubject(shape, true));
		Sys.println(enumSubject(Box(3, 4), false));
		Sys.println(new Picker(1).pick(true));
		Sys.println(new Picker(1).pick(false));
		final offset = 5;
		final local = function(name:String, flag:Bool):String {
			return switch name {
				case "x": if (flag) return "X" + offset else "x" + offset;
				case _: name;
			};
		};
		Sys.println(local("x", true));
		Sys.println(local("x", false));
		Sys.println(local("y", true));
	}
}
