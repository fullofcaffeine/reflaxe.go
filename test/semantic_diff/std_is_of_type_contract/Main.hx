enum Token {
	Alpha;
	Beta(value:Int);
}

class Base {
	public function new() {}
}

class Child extends Base {
	public function new() {
		super();
	}
}

class Main {
	static function emit(label:String, value:Bool):Void {
		Sys.println(label + "=" + (value ? "true" : "false"));
	}

	static function main() {
		var child:Base = new Child();
		var base:Base = new Base();

		emit("typed.child_is_child", Std.isOfType(child, Child));
		emit("typed.child_is_base", Std.isOfType(child, Base));
		emit("typed.base_is_child", Std.isOfType(base, Child));
		emit("typed.null_is_child", Std.isOfType(null, Child));

		emit("typed.int_is_int", Std.isOfType(1, Int));
		emit("typed.int_is_float", Std.isOfType(1, Float));
		emit("typed.float_is_int", Std.isOfType(1.5, Int));
		emit("typed.string_is_string", Std.isOfType("x", String));
		emit("typed.bool_is_bool", Std.isOfType(true, Bool));
		emit("typed.null_is_dynamic", Std.isOfType(null, Dynamic));

		var dynamicValue:Dynamic = new Child();
		emit("dynamic.child_is_base", Std.isOfType(dynamicValue, Base));
		emit("dynamic.child_is_child", Std.isOfType(dynamicValue, Child));

		dynamicValue = new Base();
		emit("dynamic.base_is_child", Std.isOfType(dynamicValue, Child));

		dynamicValue = [1, 2];
		emit("dynamic.array_is_array", Std.isOfType(dynamicValue, Array));

		dynamicValue = 1;
		emit("dynamic.int_is_array", Std.isOfType(dynamicValue, Array));

		dynamicValue = Token.Beta(3);
		emit("dynamic.enum_is_token", Std.isOfType(dynamicValue, Token));

		var token:Token = Token.Alpha;
		emit("typed.enum_is_token", Std.isOfType(token, Token));
	}
}
