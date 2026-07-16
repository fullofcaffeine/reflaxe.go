class Main {
	static function main():Void {
		var erased:Iterable<Iterable<Int>> = cast [[1, 2], [3]];
		Sys.println(Lambda.flatten(erased));
	}
}
