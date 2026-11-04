package haxe;

/**
	What: Owns exception state, conversion policy and saved stack lifetime in Haxe.
	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because its extern contract requires a target implementation; a native carrier
	cannot preserve ordinary generated inheritance, previous and native identity.
	How: Keep typed state in generated fields, capture the selected diagnostic stack
	once, and validate arbitrary payloads at caught/thrown boundaries. hxrt only
	transports tagged Go panics; native attachments remain separate from payloads.
 */
@:keep
class Exception {
	public var message(get, never):String;
	public var stack(get, never):CallStack;
	public var previous(get, never):Null<Exception>;
	public var native(get, never):Any;

	var savedMessage:String;
	var savedStack:CallStack;
	var savedPrevious:Null<Exception>;
	// Any is required by the public native-attachment API, not used for typed state.
	var savedNative:Any;

	public function new(message:String, ?previous:Exception, ?native:Any) {
		savedMessage = message;
		savedPrevious = previous;
		savedNative = native == null ? this : native;
		savedStack = CallStack.callStack();
	}

	function get_message():String
		return savedMessage;

	function get_stack():CallStack
		return savedStack;

	function get_previous():Null<Exception>
		return savedPrevious;

	final function get_native():Any
		return savedNative;

	function unwrap():Any
		return savedNative;

	public function toString():String
		return message;

	public function details():String
		return CallStack.exceptionToString(this);

	/** Validate arbitrary caught values before restoring the nominal source type. */
	static function caught(value:Any):Exception {
		if (Std.isOfType(value, Exception)) {
			return cast value;
		}
		return new ValueException(value, null, value);
	}

	/** Go transports Haxe objects themselves; native attachments are not Go panics. */
	static function thrown(value:Any):Any {
		if (Std.isOfType(value, Exception))
			return value;
		return new ValueException(value);
	}

	/** Non-exception catches observe a ValueException payload, not its wrapper. */
	static function catchValue(value:Any):Any {
		if (Std.isOfType(value, ValueException)) {
			var wrapped:ValueException = cast value;
			return wrapped.value;
		}
		return value;
	}
}
