package haxe.exceptions;

/**
	What
	A staged `haxe.exceptions.PosException` override for `haxe.go`.

	Why
	The upstream implementation is semantically correct, but its source shape
	pulls in extra staged std dependencies that are not necessary for the direct
	exception contract this backend needs to preserve here.

	How
	Keep the public API and message/position behavior, but express construction and
	`toString` with explicit concatenation so the class lowers through the existing
	`haxe.Exception` runtime carrier without widening the stdlib ownership graph.
**/
class PosException extends Exception {
	public final posInfos:PosInfos;

	public function new(message:String, ?previous:Exception, ?pos:PosInfos):Void {
		super(message, previous);
		if (pos == null) {
			posInfos = {
				fileName: "(unknown)",
				lineNumber: 0,
				className: "(unknown)",
				methodName: "(unknown)"
			};
		} else {
			posInfos = pos;
		}
	}

	override function toString():String {
		return message
			+ " in "
			+ posInfos.className
			+ "."
			+ posInfos.methodName
			+ " at "
			+ posInfos.fileName
			+ ":"
			+ posInfos.lineNumber;
	}
}
