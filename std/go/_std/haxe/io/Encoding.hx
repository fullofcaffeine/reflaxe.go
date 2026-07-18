package haxe.io;

/**
	What: The portable string encodings accepted by Haxe IO APIs.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	while RawNative is selected by a target define. The public constructors remain
	ordinary Haxe policy rather than compiler-emitted tags.

	How: Lower the upstream constructors from staged source. `Bytes` interprets
	`RawNative` using the documented `reflaxe_go_raw_native_mode` policy.
**/
enum Encoding {
	UTF8;
	RawNative;
}
