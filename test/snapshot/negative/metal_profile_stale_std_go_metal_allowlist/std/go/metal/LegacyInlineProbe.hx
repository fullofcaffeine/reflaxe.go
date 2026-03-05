package go.metal;

class LegacyInlineProbe {
	public static inline function one():Int {
		return untyped __go__ /*probe*/ ("1");
	}
}
