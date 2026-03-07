package haxe;

/**
	Elements returned by `CallStack` methods.

	Why
	- `haxe.Exception` and user debug helpers need the public `StackItem` enum shape to exist.
	- The mainstream stdlib version depends on `NativeStackTrace` integration that `haxe.go`
	  does not provide yet.

	What
	- Declares the public `StackItem` enum and a deterministic `CallStack` abstract.

	How
	- `callStack()` and `exceptionStack()` currently return empty arrays on `haxe.go`.
	- `toString()` is intentionally empty until native stack capture exists.
	- This keeps the API available without pretending Go stack integration is portable today.
**/
enum StackItem {
	CFunction;
	Module(m:String);
	FilePos(s:Null<StackItem>, file:String, line:Int, ?column:Int);
	Method(classname:Null<String>, method:String);
	LocalFunction(?v:Int);
}

/**
	Get information about the call stack.

	Why
	- Upstream `haxe.CallStack` cannot be reused unchanged because it relies on target-owned
	  `NativeStackTrace` data that does not exist on `haxe.go` yet.

	What
	- Exposes the stdlib `CallStack` API with deterministic Go-specific fallback behavior.

	How
	- Stack queries return empty arrays.
	- Formatting returns `""`.
	- Copying and array access still behave like the upstream abstract surface.
**/
@:allow(haxe.Exception)
@:using(haxe.CallStack)
abstract CallStack(Array<StackItem>) from Array<StackItem> {
	public var length(get, never):Int;

	inline function get_length():Int
		return this.length;

	public static function callStack():Array<StackItem> {
		return [];
	}

	public static function exceptionStack(_fullStack:Bool = false):Array<StackItem> {
		return [];
	}

	static public function toString(_stack:CallStack):String {
		return "";
	}

	static function exceptionToString(e:Exception):String {
		return "Exception: " + Std.string(e);
	}

	public function subtract(_stack:CallStack):CallStack {
		return this;
	}

	public inline function copy():CallStack {
		return this.copy();
	}

	@:arrayAccess public inline function get(index:Int):StackItem {
		return this[index];
	}

	inline function asArray():Array<StackItem> {
		return this;
	}
}
