package hxrt.math;

/**
	What:
	- Maps `Math.random()` and the guarded `Std.random(max)` call to Go's
	  process-wide pseudo-random source.

	Why:
	- Random generation belongs to the native standard library, while the Haxe
	  `[0, 1)` and guarded integer-bound contracts remain declared by staged
	  `Math` and `Std` source.

	How:
	- Bind directly to `math/rand.Float64` and `math/rand.Intn`; no compiler rule,
	  raw injection, or Go runtime wrapper is required.
**/
@:go.import("math/rand")
extern class NativeRandom {
	@:go.name("Float64")
	public static function float64():Float;

	@:go.name("Intn")
	public static function intn(max:Int):Int;
}
