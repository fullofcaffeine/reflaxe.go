class InlineThrowBox {
	public final label:String;

	public function new(label:String) {
		this.label = label;
	}
}

class InlineThrowAccessors<T> {
	final valid:Bool;
	final textValue:String;
	final intValue:Int;
	final boolValue:Bool;
	final nullableValue:Null<InlineThrowBox>;
	final genericValue:T;

	public var text(get, never):String;
	public var number(get, never):Int;
	public var flag(get, never):Bool;
	public var maybe(get, never):Null<InlineThrowBox>;
	public var item(get, never):T;

	public function new(valid:Bool, textValue:String, intValue:Int, boolValue:Bool, nullableValue:Null<InlineThrowBox>, genericValue:T) {
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

	inline function get_maybe():Null<InlineThrowBox> {
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
	What: Exercises inline throwing accessors through different enclosing expressions.
	Why: A throw fallback must use the accessor's immediate result type, not the
	comparison, interpolation, or return expression that receives the accessor.
	How: Cover String, Int, Bool, nullable-reference, and generic results on both
	valid and caught-throw paths.
**/
class Main {
	static function stringContext(value:InlineThrowAccessors<String>):String {
		return "string=" + (value.text == "text");
	}

	static function intContext(value:InlineThrowAccessors<String>):String {
		return "int=" + (value.number == 7);
	}

	static function boolContext(value:InlineThrowAccessors<String>):String {
		return 'bool=${value.flag}';
	}

	static function nullableContext(value:InlineThrowAccessors<String>):String {
		return "nullable=" + (value.maybe != null);
	}

	static function genericValue<T>(value:InlineThrowAccessors<T>):T {
		return value.item;
	}

	static function genericContext(value:InlineThrowAccessors<String>):String {
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
		var valid = new InlineThrowAccessors(true, "text", 7, true, new InlineThrowBox("box"), "generic");
		var invalid = new InlineThrowAccessors(false, "text", 7, true, new InlineThrowBox("box"), "generic");

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
