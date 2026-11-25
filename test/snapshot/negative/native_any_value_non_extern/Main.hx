import go.NativeAny;

class Point {
	public final x:Int;

	public function new(x:Int) {
		this.x = x;
	}
}

/**
	`go.NativeAny.ofValue` stores the Go value behind a `@:go.struct` or
	`@:go.valueType` extern carrier. A Haxe class has no such Go value, so the
	compiler rejects it instead of storing the Haxe carrier.
**/
class Main {
	static function main():Void {
		final value = NativeAny.ofValue(new Point(1));
		Sys.println(value != null);
	}
}
