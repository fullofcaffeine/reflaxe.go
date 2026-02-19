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

class ResolverLog {
	public static var entries:Array<String> = [];

	public static function reset():Void {
		entries = [];
	}

	public static function push(entry:String):Void {
		entries.push(entry);
	}

	public static function dump():String {
		return haxe.Serializer.run(entries);
	}

	public static function record(prefix:String, name:Dynamic):Void {
		entries.push(prefix + ":" + haxe.Serializer.run(name));
	}
}

class DynamicArgResolver {
	public function new() {}

	public function resolveClass(name:Dynamic):Class<Dynamic> {
		ResolverLog.record("dc", name);
		return null;
	}

	public function resolveEnum(name:Dynamic):Enum<Dynamic> {
		ResolverLog.record("de", name);
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

	static function runObjectResolver(raw:String):String {
		ResolverLog.reset();
		var u = new haxe.Unserializer(raw);
		u.setResolver({
			resolveClass: function(name:String):Class<Dynamic> {
				ResolverLog.record("oc", name);
				return null;
			},
			resolveEnum: function(name:String):Enum<Dynamic> {
				ResolverLog.record("oe", name);
				return null;
			}
		});
		try {
			u.unserialize();
		} catch (_:Dynamic) {}
		return ResolverLog.dump();
	}

	static function runDynamicResolver(raw:String):String {
		ResolverLog.reset();
		var u = new haxe.Unserializer(raw);
		u.setResolver(new DynamicArgResolver());
		try {
			u.unserialize();
		} catch (_:Dynamic) {}
		return ResolverLog.dump();
	}

	static function main() {
		safe("poly.obj.class", function() return runObjectResolver("cy5:Plainy1:xi4g"));
		safe("poly.obj.enum", function() return runObjectResolver("wy9:ProbeEnumy3:One:1i7"));
		safe("poly.obj.classRef", function() return runObjectResolver("Ay5:Plain"));
		safe("poly.obj.enumRef", function() return runObjectResolver("By9:ProbeEnum"));
		safe("poly.dyn.class", function() return runDynamicResolver("cy5:Plainy1:xi6g"));
		safe("poly.dyn.enum", function() return runDynamicResolver("wy9:ProbeEnumy3:One:1i8"));
		safe("poly.dyn.classRef", function() return runDynamicResolver("Ay5:Plain"));
		safe("poly.dyn.enumRef", function() return runDynamicResolver("By9:ProbeEnum"));
	}
}
