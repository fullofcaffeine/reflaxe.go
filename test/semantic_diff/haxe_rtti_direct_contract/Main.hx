import haxe.rtti.CType.CTypeTools;
import haxe.rtti.CType.TypeTree;

@:rtti
class Demo {
	public static var field:String = "value";
}

class Main {
	static function main() {
		var typeMeta = haxe.rtti.Meta.getType(Demo);
		var fieldMeta = haxe.rtti.Meta.getFields(Demo);
		var info = haxe.rtti.Rtti.getRtti(Demo);
		var rawRtti:Dynamic = Reflect.field(Demo, "__rtti");
		var parsed = new haxe.rtti.XmlParser().processElement(Xml.parse(Std.string(rawRtti)).firstElement());

		Sys.println("typeMetaKeys=" + Lambda.count(Reflect.fields(typeMeta)));
		Sys.println("fieldMetaKeys=" + Lambda.count(Reflect.fields(fieldMeta)));
		Sys.println("hasRtti=" + Std.string(haxe.rtti.Rtti.hasRtti(Demo)));
		Sys.println("rttiPath=" + info.path);
		Sys.println("rttiStaticCount=" + info.statics.length);
		Sys.println("staticType=" + CTypeTools.toString(info.statics[0].type));
		switch (parsed) {
			case TClassdecl(c):
				Sys.println("parsedPath=" + c.path);
				Sys.println("parsedStaticCount=" + c.statics.length);
			case _:
				Sys.println("parsed=unexpected");
		}
	}
}
