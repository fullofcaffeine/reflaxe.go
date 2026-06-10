package haxe;

#if reflaxe_go_native_stack_trace
import haxe.NativeStackTrace;
#end

/**
	What
	A staged `haxe.CallStack` override for `haxe.go`.

	Why
	Direct source-owned exception and collection flows can touch `CallStack`
	through `haxe.Exception`-adjacent APIs, but the upstream implementation pulls
	in a broader transitive stdlib graph than `haxe.go` needs for the current
	portable contract. That made otherwise unrelated direct-usage fixtures fail on
	module ownership collisions instead of the actual feature under test.

	How
	Provide the same public `StackItem`/`CallStack` surface with the conservative
	behavior `haxe.go` already uses today for these paths: empty captured stacks
	and deterministic string rendering when explicitly asked. When
	`reflaxe_go_native_stack_trace` is enabled, this staged API converts
	Go-native diagnostic frames from `hxrt` into `StackItem` values. This keeps
	ownership in staged std code and avoids forcing unrelated stdlib modules into
	the build.
**/
enum StackItem {
	CFunction;
	Module(m:String);
	FilePos(s:Null<StackItem>, file:String, line:Int, ?column:Int);
	Method(classname:Null<String>, method:String);
	LocalFunction(?v:Int);
}

@:allow(haxe.Exception)
@:using(haxe.CallStack)
abstract CallStack(Array<StackItem>) from Array<StackItem> {
	public var length(get, never):Int;

	inline function get_length():Int {
		return this.length;
	}

	public static function callStack():Array<StackItem> {
		#if reflaxe_go_native_stack_trace
		return NativeStackTrace.toHaxe(NativeStackTrace.callStack(), 1);
		#else
		return [];
		#end
	}

	public static function exceptionStack(fullStack = false):Array<StackItem> {
		#if reflaxe_go_native_stack_trace
		return NativeStackTrace.toHaxe(NativeStackTrace.exceptionStack(), fullStack ? 0 : 1);
		#else
		return [];
		#end
	}

	public static function toString(stack:CallStack):String {
		var out = "";
		for (item in stack.asArray()) {
			out += "\nCalled from " + itemToString(item);
		}
		return out;
	}

	public function subtract(stack:CallStack):CallStack {
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

	static function exceptionToString(e:Exception):String {
		return "Exception: " + e.toString();
	}

	static function itemToString(item:StackItem):String {
		return switch (item) {
			case CFunction:
				"a C function";
			case Module(m):
				"module " + m;
			case FilePos(inner, file, line, column):
				var rendered = inner == null ? file : itemToString(inner) + " (" + file;
				rendered += " line " + line;
				if (column > 0) {
					rendered += " column " + column;
				}
				if (inner != null) {
					rendered += ")";
				}
				rendered;
			case Method(classname, method):
				(classname == null ? "<unknown>" : classname) + "." + method;
			case LocalFunction(v):
				"local function #" + Std.string(v);
		}
	}
}
