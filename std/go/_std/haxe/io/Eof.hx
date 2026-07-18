package haxe.io;

/**
	What: Signals that an Input has no more bytes available.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because its package-private `toString` is not visible to the separate Go runtime
	formatter. EOF identity and text still need no generated compiler carrier.

	How: Use the ordinary staged class so typed catches and subclass stream loops
	flow through the compiler's normal exception lowering.
**/
class Eof {
	public function new() {}

	@:ifFeature("haxe.io.Eof.*")
	public function toString():String {
		return "Eof";
	}
}
