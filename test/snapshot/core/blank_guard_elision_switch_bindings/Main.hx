enum Expr {
	Lit(value:Int);
	Pair(left:Int, right:Int);
}

class Main {
	static function eval(expr:Expr):Int {
		return switch (expr) {
			case Lit(value):
				value;
			case Pair(left, right):
				left + right;
		}
	}

	static function writeOnly():Void {
		var x = 0;
		x = 1;
	}

	static function main() {
		writeOnly();
		Sys.println(eval(Lit(3)));
		Sys.println(eval(Pair(2, 5)));
	}
}
