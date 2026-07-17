class SnapshotInlineThrowBox {
	public final label:String;

	public function new(label:String) {
		this.label = label;
	}
}

class SnapshotInlineThrowAccessors<T> {
	final valid:Bool;
	final textValue:String;
	final intValue:Int;
	final boolValue:Bool;
	final nullableValue:Null<SnapshotInlineThrowBox>;
	final genericValue:T;

	public var text(get, never):String;
	public var number(get, never):Int;
	public var flag(get, never):Bool;
	public var maybe(get, never):Null<SnapshotInlineThrowBox>;
	public var item(get, never):T;

	public function new(valid:Bool, textValue:String, intValue:Int, boolValue:Bool, nullableValue:Null<SnapshotInlineThrowBox>, genericValue:T) {
		this.valid = valid;
		this.textValue = textValue;
		this.intValue = intValue;
		this.boolValue = boolValue;
		this.nullableValue = nullableValue;
		this.genericValue = genericValue;
	}

	inline function get_text():String {
		if (!valid) {
			throw "text";
		}
		return textValue;
	}

	inline function get_number():Int {
		if (!valid) {
			throw "number";
		}
		return intValue;
	}

	inline function get_flag():Bool {
		return valid ? boolValue : throw "flag";
	}

	inline function get_maybe():Null<SnapshotInlineThrowBox> {
		if (!valid) {
			throw "maybe";
		}
		return nullableValue;
	}

	inline function get_item():T {
		return valid ? genericValue : throw "item";
	}
}

/**
	What: Pins fallback types for inline throwing accessors in generated Go.
	Why: A surrounding comparison or interpolation can have a different result type
	from the accessor whose invalid branch throws.
	How: Cover String, Int, Bool, nullable-reference, and generic values on valid
	and caught-throw paths.
**/
class Main {
	static function stringContext(value:SnapshotInlineThrowAccessors<String>):String {
		return "string=" + (value.text == "text");
	}

	static function intContext(value:SnapshotInlineThrowAccessors<String>):String {
		return "int=" + (value.number == 7);
	}

	static function boolContext(value:SnapshotInlineThrowAccessors<String>):String {
		return 'bool=${value.flag}';
	}

	static function nullableContext(value:SnapshotInlineThrowAccessors<String>):String {
		return "nullable=" + (value.maybe != null);
	}

	static function genericValue<T>(value:SnapshotInlineThrowAccessors<T>):T {
		return value.item;
	}

	static function genericContext(value:SnapshotInlineThrowAccessors<String>):String {
		return "generic=" + Std.string(genericValue(value));
	}

	static function nestedFunctionContext():String {
		var action = function():String {
			throw "nested";
		};
		return capture("nested", action);
	}

	static function capture(label:String, action:Void->String):String {
		try {
			return label + ":miss:" + action();
		} catch (_:Dynamic) {
			return label + ":throw";
		}
		return label + ":unreachable";
	}

	static function main() {
		var valid = new SnapshotInlineThrowAccessors(true, "text", 7, true, new SnapshotInlineThrowBox("box"), "generic");
		var invalid = new SnapshotInlineThrowAccessors(false, "text", 7, true, new SnapshotInlineThrowBox("box"), "generic");

		Sys.println(stringContext(valid));
		Sys.println(intContext(valid));
		Sys.println(boolContext(valid));
		Sys.println(nullableContext(valid));
		Sys.println(genericContext(valid));
		Sys.println(nestedFunctionContext());

		Sys.println(capture("string", () -> stringContext(invalid)));
		Sys.println(capture("int", () -> intContext(invalid)));
		Sys.println(capture("bool", () -> boolContext(invalid)));
		Sys.println(capture("nullable", () -> nullableContext(invalid)));
		Sys.println(capture("generic", () -> genericContext(invalid)));
	}
}
