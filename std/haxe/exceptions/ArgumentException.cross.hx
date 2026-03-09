package haxe.exceptions;

/**
	What
	A staged `haxe.exceptions.ArgumentException` override for `haxe.go`.

	Why
	This direct exception path should stay library-owned and portable, but the
	upstream source shape currently pulls more staged std dependencies than this
	backend needs for the contract.

	How
	Keep the upstream constructor/message behavior while relying on the local
	`PosException` staged override and explicit string concatenation.
**/
class ArgumentException extends PosException {
	public final argument:String;

	public function new(argument:String, ?message:String, ?previous:Exception, ?pos:PosInfos):Void {
		super(message == null ? "Invalid argument \"" + argument + "\"" : message, previous, pos);
		this.argument = argument;
	}
}
