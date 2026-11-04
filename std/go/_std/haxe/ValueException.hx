package haxe;

/**
	What: Wraps arbitrary Haxe throw values while retaining their original payload.
	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because the target must coordinate its exception base with tagged Go panic transport.
	How: Extend the generated Exception class, derive the message through Std.string,
	and unwrap only the saved payload. Previous and native state remain in the base.
 */
@:keep
class ValueException extends Exception {
	// The public payload can be any Haxe value; typed consumers validate recovery.
	public var value(default, null):Any;

	public function new(value:Any, ?previous:Exception, ?native:Any) {
		super(Std.string(value), previous, native);
		this.value = value;
	}

	override function unwrap():Any
		return value;
}
