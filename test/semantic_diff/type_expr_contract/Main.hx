enum DemoEnum {
	One;
}

class Main {
	static function main() {
		var classRef:Dynamic = Main;
		var enumRef:Dynamic = DemoEnum;

		Sys.println("class.is_class=" + Std.isOfType(classRef, Class));
		Sys.println("class.is_enum=" + Std.isOfType(classRef, Enum));

		Sys.println("enum.is_class=" + Std.isOfType(enumRef, Class));
		Sys.println("enum.is_enum=" + Std.isOfType(enumRef, Enum));

		var ctorRef:Dynamic = DemoEnum.One;
		Sys.println("ctor.is_enum=" + Std.isOfType(ctorRef, DemoEnum));
	}
}
