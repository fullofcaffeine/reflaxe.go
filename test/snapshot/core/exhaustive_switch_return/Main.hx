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
	Haxe accepts a non-Void function whose last statement returns or throws on
	every path. Go only accepts it when the last statement is "terminating" in
	Go's own sense, which an exhaustive enum switch without `default` and a
	`while (true)` loop are not. Those bodies end with an unreachable `panic`;
	bodies Go already accepts (a `default` branch, an `if`/`else` chain, a Void
	function) stay unchanged.
**/
class Main {
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
						return side;
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

	static function intDefault(value:Int):String {
		switch value {
			case 0:
				return "zero";
			default:
				return "many";
		}
	}

	static function stringWildcard(value:String):Int {
		switch value {
			case "a":
				return 1;
			case _:
				return 2;
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

	static function ifChain(value:Int):String {
		if (value < 0) {
			return "negative";
		} else if (value == 0) {
			return "zero";
		} else {
			return "positive";
		}
	}

	static function voidSwitch(shape:Shape, log:Array<String>):Void {
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
		final log = new Array<String>();
		for (shape in [Circle(11), Square(0), Square(3), Empty]) {
			Sys.println(area(shape) + " " + guarded(shape) + " " + new Box(shape).size());
			Sys.println(nested(shape, false) + " " + insideIf(shape, true) + " " + insideTry(shape) + " " + tryInsideArm(shape));
			try {
				Sys.println(throwArms(shape));
			} catch (e:String) {
				Sys.println("caught:" + e);
			}
			voidSwitch(shape, log);
		}
		Sys.println(log.join(","));
		Sys.println(colorName(Green) + " " + loopUntil(10) + " " + intDefault(0) + " " + stringWildcard("b"));
		Sys.println(boolSwitch(false) + " " + ifChain(-2) + " " + ifChain(0) + " " + ifChain(3));
		final pick = function(shape:Shape):String {
			switch shape {
				case Circle(_):
					return "lambda-circle";
				case Square(_) | Empty:
					return "lambda-other";
			}
		};
		Sys.println(pick(Circle(1)) + " " + pick(Empty));
	}
}
