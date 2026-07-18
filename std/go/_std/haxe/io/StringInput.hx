package haxe.io;

/**
	What: Presents a String as an in-memory byte Input.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	until target BytesInput and Bytes conversion exist. Those are now staged, so a
	separate compiler-generated StringInput carrier has no responsibility.

	How: Convert with the default UTF-8 policy and delegate all reading to
	`BytesInput` through normal source inheritance.
**/
class StringInput extends BytesInput {
	public function new(value:String) {
		super(Bytes.ofString(value), null, null);
	}
}
