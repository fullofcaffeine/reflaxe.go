class LocalException extends haxe.Exception {
	public var code:Int;

	public function new(message:String, code:Int, ?previous:haxe.Exception) {
		super(message, previous);
		this.code = code;
	}

	override public function toString():String {
		return message + ":" + code;
	}
}

class Main {
	static function main():Void {
		var cause = new haxe.Exception("root cause");
		var source = new LocalException("boom", 7, cause);
		// Deliberate erased boundary: structural recovery must retain the source.
		var erased:Dynamic = source;
		var view:{
			var message(get, never):String;
			var previous(get, never):Null<haxe.Exception>;
			final code:Int;
			function toString():String;
			function details():String;
		} = erased;
		Sys.println("message=" + view.message);
		Sys.println("code=" + view.code);
		Sys.println("string=" + view.toString());
		source.code = 8;
		Sys.println("live=" + view.code);
		var returned:Dynamic = view;
		var restored:LocalException = returned;
		Sys.println("same=" + Std.string(restored == source));
		Sys.println("direct-previous=" + Std.string(source.previous == cause));
		Sys.println("previous=" + Std.string(view.previous == cause));
		Sys.println("native=" + Std.string(source.native == source));
		var directDetails = source.details();
		var details = view.details();
		Sys.println("details-same=" + Std.string(details == directDetails));
		Sys.println("details-message=" + Std.string(details.indexOf("boom") >= 0));
		Sys.println("details-cause=" + Std.string(details.indexOf("root cause") >= 0));
	}
}
