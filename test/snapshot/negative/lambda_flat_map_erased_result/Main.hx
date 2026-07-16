class Main {
	static function main():Void {
		var erasedMapper:(value:Int) -> Iterable<Int> = function(value:Int):Iterable<Int> {
			return cast [value, value + 1];
		};
		Sys.println(Lambda.flatMap([1, 2], erasedMapper));
	}
}
