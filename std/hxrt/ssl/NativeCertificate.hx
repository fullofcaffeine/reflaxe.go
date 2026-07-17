package hxrt.ssl;

import go.NativeSlice;

/**
	What
	Typed bridge for native certificate values that cannot be expressed by the
	portable `sys.ssl.Certificate` source alone.

	Why
	The mainstream Haxe stdlib API returns `Array<String>`, while Go's TLS runtime
	naturally returns a native string slice. Declaring that slice as a Haxe Array
	would hide incompatible identity and growth semantics.

	How
	Expose the Go result as `NativeSlice<String>`. The staged public API performs
	the explicit copy into a portable Array at the boundary.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeCertificate {
	@:go.name("SslCertAltNames")
	public static function altNames(handle:Dynamic):NativeSlice<String>;
}
