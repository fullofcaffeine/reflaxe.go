package hxrt.zip;

import go.NativeSlice;

/**
	What
	- Typed carrier for one bounded deflate or inflate step.

	Why
	- Public `execute` results and destination writes belong to staged Haxe, so
	  the runtime must not construct anonymous Haxe objects or generated
	  `haxe.io.Bytes` instances.

	How
	- Carry only native output values, consumed-source count, and completion;
	  staged source performs the `Bytes.blit` and creates the public result.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ZipCodecStep")
extern class ZipCodecStep {
	@:go.name("Values")
	public var values:NativeSlice<Int>;

	@:go.name("Read")
	public var read:Int;

	@:go.name("Done")
	public var done:Bool;
}
