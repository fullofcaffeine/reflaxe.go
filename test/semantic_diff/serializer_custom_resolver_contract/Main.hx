class Plain {
	public var x:Int;

	public function new(?x:Int) {
		this.x = x == null ? 0 : x;
	}
}

class Custom {
	public var x:Dynamic;

	public function new(?x:Int) {
		x = x == null ? 0 : x;
		this.x = x;
	}

	public function hxSerialize(s:haxe.Serializer):Void {
		s.serialize(x);
	}

	public function hxUnserialize(u:haxe.Unserializer):Void {
		x = u.unserialize();
	}
}

enum ProbeEnum {
	NoArgs;
	One(v:Int);
}

class NullResolver {
	public function new() {}

	public function resolveClass(name:String):Class<Dynamic> {
		return null;
	}

	public function resolveEnum(name:String):Enum<Dynamic> {
		return null;
	}
}

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		safe("custom.ser", function() return haxe.Serializer.run(new Custom(2)));
		safe("custom.replay", function() {
			var value = haxe.Unserializer.run("Cy6:Customi3g");
			return haxe.Serializer.run(value);
		});

		safe("plain.c.replay", function() {
			var value = haxe.Unserializer.run("cy5:Plainy1:xi7g");
			return haxe.Serializer.run(value);
		});

		safe("enum.w.replay", function() {
			var value:ProbeEnum = cast haxe.Unserializer.run("wy9:ProbeEnumy3:One:1i9");
			return haxe.Serializer.run(value);
		});

		safe("enum.j.replay", function() {
			var value:ProbeEnum = cast haxe.Unserializer.run("jy9:ProbeEnum:1:1i9");
			return haxe.Serializer.run(value);
		});

		safe("classref.default", function() return haxe.Serializer.run(haxe.Unserializer.run("Ay5:Plain")));
		safe("enumref.default", function() return haxe.Serializer.run(haxe.Unserializer.run("By9:ProbeEnum")));

		safe("resolver.null.class", function() {
			var u = new haxe.Unserializer("cy5:Plainy1:xi1g");
			u.setResolver(null);
			u.unserialize();
			return "unreachable";
		});

		safe("resolver.null.enum", function() {
			var u = new haxe.Unserializer("wy9:ProbeEnumy6:NoArgs:0");
			u.setResolver(new NullResolver());
			u.unserialize();
			return "unreachable";
		});
	}
}
