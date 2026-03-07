import haxe.ds.Either;

class Main {
	static function main() {
		var left:Either<String, Int> = Left("go");
		var right:Either<String, Int> = Right(7);
		Sys.println(switch left {
			case Left(value): "left=" + value;
			case Right(value): "right=" + Std.string(value);
		});
		Sys.println(switch right {
			case Left(value): "left=" + value;
			case Right(value): "right=" + Std.string(value);
		});
	}
}
