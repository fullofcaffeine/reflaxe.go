enum Color {
	Red;
	Rgb(r:Int, g:Int, b:Int);
}

class Box {
	public static var createdWith:Int = -1;

	public var x:Int;

	public function new(x:Int) {
		createdWith = x;
		this.x = x;
	}
}

class Main {
	static function main() {
		var boxClass = Type.getClass(new Box(7));
		Sys.println("class.name=" + Type.getClassName(boxClass));

		var resolvedClass = Type.resolveClass("Box");
		Sys.println("class.resolve=" + Type.getClassName(resolvedClass));

		var classForCreate:Class<Box> = cast(resolvedClass == null ? boxClass : resolvedClass);
		var createdBox = Type.createInstance(classForCreate, [3]);
		Sys.println("class.createdWith=" + Box.createdWith);
		Sys.println("class.instance.class=" + Type.getClassName(Type.getClass(createdBox)));

		var enumType = Type.getEnum(Color.Red);
		Sys.println("enum.name=" + Type.getEnumName(enumType));

		var resolvedEnum = Type.resolveEnum("Color");
		Sys.println("enum.resolve=" + Type.getEnumName(resolvedEnum));

		var enumForCreate:Enum<Color> = cast(resolvedEnum == null ? enumType : resolvedEnum);
		var createdEnum = Type.createEnum(enumForCreate, "Rgb", [1, 2, 3]);
		Sys.println("enum.ctor=" + Type.enumConstructor(createdEnum));
		Sys.println("enum.index=" + Type.enumIndex(createdEnum));

		var params = Type.enumParameters(createdEnum);
		Sys.println("enum.params=" + params.length + ":" + params[0] + ":" + params[1] + ":" + params[2]);

		Sys.println("enum.eq.true=" + Std.string(Type.enumEq(createdEnum, Color.Rgb(1, 2, 3))));
		Sys.println("enum.eq.false=" + Std.string(Type.enumEq(createdEnum, Color.Rgb(1, 2, 4))));
	}
}
