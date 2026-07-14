package haxe.exceptions;

/**
	What
	A staged `haxe.exceptions.NotImplementedException` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The direct exception contract here only needs the public constructor/message
	behavior. Keeping it staged and minimal avoids widening unrelated stdlib
	ownership.

	How
	Preserve the upstream default message and inheritance shape through the local
	`PosException` staged override.
**/
class NotImplementedException extends PosException {
	public function new(message:String = "Not implemented", ?previous:Exception, ?pos:PosInfos):Void {
		super(message == null ? "Not implemented" : message, previous, pos);
	}
}
