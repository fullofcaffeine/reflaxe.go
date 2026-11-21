enum Shape {
	Circle(r:Int);
	Square(side:Int);
	Empty;
}

enum abstract Color(Int) {
	var Red = 0;
	var Green = 1;
}

class Box {
	final shape:Shape;

	public function new(shape:Shape) {
		this.shape = shape;
	}

	public function size():Int {
		switch shape {
			case Circle(r):
				return r * 2;
			case Square(side):
				return side;
			case Empty:
				return 0;
		}
	}
}

/**
	Functions whose last statement is a statement `switch` (or a loop) that
	returns or throws on every path.

	Haxe accepts these bodies because it proves that the end of the function
	cannot be reached. Each result must match `--interp`.
**/
class Main {
	static final log = new Array<String>();

	static function area(shape:Shape):Int {
		switch shape {
			case Circle(r):
				return r * r * 3;
			case Square(side):
				return side * side;
			case Empty:
				return 0;
		}
	}

	static function guarded(shape:Shape):String {
		switch shape {
			case Circle(r) if (r > 10):
				return "big circle";
			case Circle(_):
				return "circle";
			case Square(side) if (side == 0):
				return "dot";
			case Square(_) | Empty:
				return "other";
		}
	}

	static function intDefault(value:Int):String {
		switch value {
			case 0:
				return "zero";
			case 1 | 2:
				return "small";
			default:
				return "many";
		}
	}

	static function stringWildcard(value:String):Int {
		switch value {
			case "a":
				return 1;
			case "b":
				return 2;
			case _:
				return 3;
		}
	}

	static function boolSwitch(flag:Bool):String {
		switch flag {
			case true:
				return "yes";
			case false:
				return "no";
		}
	}

	static function colorName(color:Color):String {
		switch color {
			case Red:
				return "red";
			case Green:
				return "green";
		}
	}

	static function nested(shape:Shape, flag:Bool):Int {
		switch shape {
			case Circle(r):
				switch flag {
					case true:
						return r;
					case false:
						return -r;
				}
			case Square(side):
				switch side {
					case 0:
						return 100;
					default:
						switch shape {
							case Square(_):
								return side + 1;
							case _:
								return -1;
						}
				}
			case Empty:
				return 7;
		}
	}

	static function insideIf(shape:Shape, flag:Bool):Int {
		if (flag) {
			switch shape {
				case Circle(r):
					return r;
				case Square(side):
					return side;
				case Empty:
					return 0;
			}
		} else {
			return -1;
		}
	}

	static function ifChain(value:Int):String {
		if (value < 0) {
			return "negative";
		} else if (value == 0) {
			return "zero";
		} else if (value < 10) {
			return "small";
		} else {
			return "large";
		}
	}

	static function throwArms(shape:Shape):Int {
		switch shape {
			case Circle(r):
				return r;
			case Square(_):
				throw "square";
			case Empty:
				throw "empty";
		}
	}

	static function insideTry(shape:Shape):Int {
		try {
			switch shape {
				case Circle(r):
					return r;
				case Square(side):
					return side;
				case Empty:
					throw "empty";
			}
		} catch (e:String) {
			return -1;
		}
	}

	static function tryInsideArm(shape:Shape):Int {
		switch shape {
			case Circle(r):
				try {
					if (r < 0)
						throw "negative";
					return r;
				} catch (e:String) {
					return 0;
				}
			case Square(side):
				return side;
			case Empty:
				return -5;
		}
	}

	static function loopUntil(limit:Int):Int {
		var i = 0;
		while (true) {
			if (i * i >= limit)
				return i;
			i++;
		}
	}

	static function doLoopUntil(limit:Int):Int {
		var i = 1;
		do {
			if (i >= limit)
				return i;
			i *= 2;
		} while (true);
	}

	static function voidSwitch(shape:Shape):Void {
		switch shape {
			case Circle(r):
				log.push("circle" + r);
				return;
			case Square(_):
				return;
			case Empty:
				log.push("empty");
				return;
		}
	}

	static function main():Void {
		final shapes = [Circle(2), Circle(11), Square(0), Square(4), Empty];
		for (shape in shapes) {
			Sys.println(area(shape) + " " + guarded(shape) + " " + new Box(shape).size());
			Sys.println(nested(shape, true) + " " + nested(shape, false) + " " + insideIf(shape, true) + " " + insideIf(shape, false));
			Sys.println(insideTry(shape) + " " + tryInsideArm(shape));
			try {
				Sys.println(throwArms(shape));
			} catch (e:String) {
				Sys.println("caught:" + e);
			}
			voidSwitch(shape);
		}
		Sys.println(tryInsideArm(Circle(-3)));
		Sys.println(log.join(","));
		for (value in [-4, 0, 2, 7, 30])
			Sys.println(intDefault(value) + " " + ifChain(value));
		Sys.println(stringWildcard("a") + stringWildcard("b") + stringWildcard("z"));
		Sys.println(boolSwitch(true) + " " + boolSwitch(false));
		Sys.println(colorName(Red) + " " + colorName(Green));
		Sys.println(loopUntil(10) + " " + doLoopUntil(5));
		final pick = function(shape:Shape):String {
			switch shape {
				case Circle(_):
					return "lambda-circle";
				case Square(side) if (side > 2):
					return "lambda-wide";
				case Square(_) | Empty:
					return "lambda-other";
			}
		};
		Sys.println(pick(Circle(1)) + " " + pick(Square(3)) + " " + pick(Square(1)) + " " + pick(Empty));
		final arrow = (shape:Shape) -> {
			switch shape {
				case Circle(r):
					return r + 1;
				case Square(side):
					return side - 1;
				case Empty:
					throw "arrow-empty";
			}
		};
		Sys.println(arrow(Circle(1)) + arrow(Square(5)));
	}
}
