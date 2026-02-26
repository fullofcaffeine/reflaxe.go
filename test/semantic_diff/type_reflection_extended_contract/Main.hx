enum Flag {
	A;
	B(v:Int);
	C;
}

class Base {
	public static var baseStatic:Int = 1;

	public var x:Int;

	public function new(x:Int) {
		this.x = x;
	}

	public function ping():String {
		return "base";
	}
}

class Child extends Base {
	public static var childStatic:Int = 2;

	public var y:Int;

	public function new(x:Int, y:Int) {
		super(x);
		this.y = y;
	}

	public function pong():String {
		return "child";
	}
}

class Main {
	static function has(values:Array<String>, needle:String):Bool {
		for (value in values) {
			if (value == needle) {
				return true;
			}
		}
		return false;
	}

	static function join(values:Array<String>):String {
		var out = "";
		for (index in 0...values.length) {
			if (index > 0) {
				out += "|";
			}
			out += values[index];
		}
		return out;
	}

	static function main() {
		var cls = Type.getClass(new Child(3, 4));
		var superCls = Type.getSuperClass(cls);
		Sys.println("super=" + Type.getClassName(superCls));

		var classFields = Type.getClassFields(cls);
		Sys.println("classFields.has.childStatic=" + Std.string(has(classFields, "childStatic")));
		Sys.println("classFields.has.baseStatic=" + Std.string(has(classFields, "baseStatic")));

		var instanceFields = Type.getInstanceFields(cls);
		Sys.println("instanceFields.has.x=" + Std.string(has(instanceFields, "x")));
		Sys.println("instanceFields.has.y=" + Std.string(has(instanceFields, "y")));
		Sys.println("instanceFields.has.ping=" + Std.string(has(instanceFields, "ping")));
		Sys.println("instanceFields.has.pong=" + Std.string(has(instanceFields, "pong")));

		var enumType = Type.getEnum(Flag.A);
		var constructs = Type.getEnumConstructs(enumType);
		Sys.println("enumConstructs=" + join(constructs));

		var empty = Type.createEmptyInstance(cls);
		Sys.println("empty.class=" + Type.getClassName(Type.getClass(empty)));

		var all = Type.allEnums(enumType);
		Sys.println("allEnums.length=" + all.length);
		for (value in all) {
			Sys.println("allEnums.item=" + Type.enumConstructor(value));
		}

		var typeNull = Type.typeof(null);
		Sys.println("typeof.null=" + Type.enumConstructor(typeNull));
		Sys.println("typeof.int=" + Type.enumConstructor(Type.typeof(1)));
		Sys.println("typeof.float=" + Type.enumConstructor(Type.typeof(1.5)));
		Sys.println("typeof.bool=" + Type.enumConstructor(Type.typeof(true)));
		Sys.println("typeof.string=" + Type.enumConstructor(Type.typeof("x")));
		Sys.println("typeof.array=" + Type.enumConstructor(Type.typeof([1, 2, 3])));
		Sys.println("typeof.object=" + Type.enumConstructor(Type.typeof({foo: 1})));
		Sys.println("typeof.function=" + Type.enumConstructor(Type.typeof(function() return 1)));

		var classType = Type.typeof(new Child(1, 2));
		Sys.println("typeof.class=" + Type.enumConstructor(classType));
		Sys.println("typeof.class.param=" + Type.getClassName(cast Type.enumParameters(classType)[0]));

		var enumValueType = Type.typeof(Flag.B(10));
		Sys.println("typeof.enum=" + Type.enumConstructor(enumValueType));
		Sys.println("typeof.enum.param=" + Type.getEnumName(cast Type.enumParameters(enumValueType)[0]));
	}
}
