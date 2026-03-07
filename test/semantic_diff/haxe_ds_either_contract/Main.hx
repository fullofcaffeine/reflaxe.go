import haxe.ds.Either;

class Main {
	static function render(value:Either<String, Int>):String {
		return switch value {
			case Left(text): "left=" + text;
			case Right(number): "right=" + Std.string(number);
		};
	}

	static function main() {
		Sys.println(render(Left("go")));
		Sys.println(render(Right(7)));
	}
}
