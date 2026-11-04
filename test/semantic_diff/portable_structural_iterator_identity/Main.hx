class Cursor {
	public function new() {}

	public function hasNext():Bool
		return false;

	public function next():Int
		return 7;
}

class Main {
	// Erasure deliberately tests object identity, not compile-time view equality.
	static function same(left:Dynamic, right:Dynamic):Bool
		return left == right;

	static function missing():Cursor
		return null;

	static function main() {
		final source = new Cursor();
		final left:Iterator<Int> = source;
		final right:Iterator<Int> = source;
		Sys.println(same(left, right));
		final a = left.next;
		final b = left.next;
		Sys.println(a == b);
		Sys.println(Reflect.compareMethods(a, b));
		final absent:Iterator<Int> = missing();
		Sys.println(absent == null);
	}
}
