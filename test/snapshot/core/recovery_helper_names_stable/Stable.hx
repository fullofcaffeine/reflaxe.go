import Shapes;

/** Recovery helpers used here must keep their names when another module changes. */
class Stable {
	public static function run(value:Dynamic, callback:Dynamic):Int {
		var apply:Int->Int = callback;
		if (!Std.isOfType(value, Circle))
			return 0;
		var circle:Circle = value;
		return apply(circle.size);
	}
}
