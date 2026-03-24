@:keep
@:rtti
class Demo {
	public static var value:Int = 1;
}

class Main {
	static function main() {
		var fields = Type.getClassFields(Demo);
		Sys.println("fields.len=" + fields.length);
		Sys.println("fields.has_rtti=" + Std.string(Lambda.has(fields, "__rtti")));
		Sys.println("raw.rtti.null=" + Std.string(Reflect.field(Demo, "__rtti") == null));
		Sys.println("raw.meta.null=" + Std.string(Reflect.field(Demo, "__meta__") == null));
	}
}
