package go;

/**
	What
	- Typed facade for the Go `time.Time` API used by baseline interop examples.

	Why
	- `time.Time` is a real Go-native type, not an upstream Haxe stdlib override.
	  The facade therefore stays publicly importable from ordinary `std/go`.

	How
	- Type, member, and receiver metadata map construction and Unix timestamp
	  access directly to the native Go API.
**/
@:go.import("time")
@:go.name("Time")
extern class Time {
	@:go.name("Now")
	public static function now():Time;

	@:go.name("Unix")
	public function unix():Int;

	@:go.receiver
	@:go.name("Unix")
	public static function unixViaReceiver(value:Time):Int;
}
