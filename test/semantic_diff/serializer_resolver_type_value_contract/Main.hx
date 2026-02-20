class Plain {
	public var x:Int;

	public function new(?x:Int) {
		this.x = x == null ? 0 : x;
	}
}

enum ProbeEnum {
	NoArgs;
	One(v:Int);
}

class TypeValueResolver {
	public function new() {}

	public function resolveClass(name:String):Class<Dynamic> {
		return switch (name) {
			case "Plain": cast Plain;
			case _: null;
		}
	}

	public function resolveEnum(name:String):Enum<Dynamic> {
		return switch (name) {
			case "ProbeEnum": cast ProbeEnum;
			case _: null;
		}
	}
}

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (_:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function runWith(raw:String):Dynamic {
		var u = new haxe.Unserializer(raw);
		u.setResolver(new TypeValueResolver());
		return u.unserialize();
	}

	static function main() {
		safe("type.class", function() return haxe.Serializer.run(runWith("cy5:Plainy1:xi3g")));
		safe("type.classRef", function() return haxe.Serializer.run(runWith("Ay5:Plain")));
		safe("type.enum", function() return haxe.Serializer.run(runWith("wy9:ProbeEnumy3:One:1i4")));
		safe("type.enumIndex", function() return haxe.Serializer.run(runWith("jy9:ProbeEnum:1:1i7")));
		safe("type.enumRef", function() return haxe.Serializer.run(runWith("By9:ProbeEnum")));
	}
}
