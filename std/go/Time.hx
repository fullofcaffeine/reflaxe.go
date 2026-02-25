package go;

/**
	Typed wrapper for the Go `time` package `Time` API used by baseline interop examples.
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
