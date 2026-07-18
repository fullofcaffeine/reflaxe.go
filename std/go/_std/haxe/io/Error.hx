package haxe.io;

/**
	What: Enumerates the portable failures raised by Haxe IO operations.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	without a target IO hierarchy that transports these values. Bounds, blocking,
	overflow, and custom-error policy still belong to the Haxe stdlib.

	How: Retain the upstream enum constructors so ordinary pattern matching and
	exception transport use the compiler's shared enum machinery.
**/
enum Error {
	Blocked;
	Overflow;
	OutsideBounds;
	Custom(e:Dynamic);
}
