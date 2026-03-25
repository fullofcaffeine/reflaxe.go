import haxe.rtti.CType.CTypeTools;
import haxe.rtti.CType.TypeTree;

@:rtti
class Demo {
	public static var field:String = "value";
}

class Main {
	static function main() {
		var info = haxe.rtti.Rtti.getRtti(Demo);
		Sys.println("path=" + info.path);
		Sys.println("staticType=" + CTypeTools.toString(info.statics[0].type));

		var rawRtti:Dynamic = Reflect.field(Demo, "__rtti");
		var parsed = new haxe.rtti.XmlParser().processElement(Xml.parse(Std.string(rawRtti)).firstElement());
		switch (parsed) {
			case TClassdecl(c):
				Sys.println("parsedPath=" + c.path);
			case _:
				Sys.println("parsed=unexpected");
		}
	}
}
