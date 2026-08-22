/** Exercises cross-module Array.slice lowering under explicit Go-native class authority. */
@:goNative
final class NativeSlicer {
	public static function middle(values:Array<Int>):Array<Int> {
		return values.slice(1, -1);
	}
}
