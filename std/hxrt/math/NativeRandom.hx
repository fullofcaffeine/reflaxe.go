package hxrt.math;

/**
	What:
	- Maps `Math.random()` to Go's process-wide pseudo-random source.

	Why:
	- Random generation belongs to the native standard library, while the Haxe
	  `[0, 1)` contract remains declared by staged `Math` source.

	How:
	- Bind directly to `math/rand.Float64`; no compiler rule, raw injection, or Go
	  runtime wrapper is required.
**/
@:go.import("math/rand")
extern class NativeRandom {
	@:go.name("Float64")
	public static function float64():Float;
}
