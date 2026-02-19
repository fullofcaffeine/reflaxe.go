class Main {
	static function main() {
		var obj:Dynamic = {
			name: "Ada",
			count: 1,
			nilField: null
		};

		Sys.println(Reflect.hasField(obj, "name"));
		Sys.println(Reflect.field(obj, "name"));
		Sys.println(Reflect.hasField(obj, "missing"));
		Sys.println(Std.string(Reflect.field(obj, "missing")));

		Reflect.setField(obj, "name", "Bea");
		Reflect.setField(obj, "extra", 42);
		Sys.println(Reflect.field(obj, "name"));
		Sys.println(Reflect.field(obj, "extra"));
		Sys.println(Reflect.hasField(obj, "extra"));

		Sys.println(Reflect.hasField(obj, "nilField"));
		Sys.println(Std.string(Reflect.field(obj, "nilField")));
		Reflect.setField(obj, "nilField", "x");
		Sys.println(Reflect.field(obj, "nilField"));
		Reflect.setField(obj, "nilField", null);
		Sys.println(Std.string(Reflect.field(obj, "nilField")));

		Sys.println(Reflect.hasField(null, "x"));
		Sys.println(Std.string(Reflect.field(null, "x")));
		try {
			Reflect.setField(null, "x", 1);
			Sys.println("set-null-ok");
		} catch (e:Dynamic) {
			Sys.println("set-null-err:" + Std.string(e));
		}
		Sys.println("done");
	}
}
