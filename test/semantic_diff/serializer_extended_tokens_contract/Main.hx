class ProbeClass {
	public var x:Int;

	public function new(?x:Int) {
		this.x = x == null ? 0 : x;
	}
}

enum ProbeEnum {
	NoArgs;
	One(v:Int);
	Two(a:Int, b:String);
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
		var list = new haxe.ds.List<Dynamic>();
		list.add(1);
		list.add("x");
		safe("list.value", function() return haxe.Serializer.run(list));
		safe("list.token", function() return haxe.Serializer.run(haxe.Unserializer.run("li1y1:xh")));

		var sm = new haxe.ds.StringMap<Dynamic>();
		sm.set("a", 1);
		safe("sm.value", function() return haxe.Serializer.run(sm));
		safe("sm.token", function() return haxe.Serializer.run(haxe.Unserializer.run("by1:ai1h")));

		var im = new haxe.ds.IntMap<Dynamic>();
		im.set(3, "v");
		safe("im.value", function() return haxe.Serializer.run(im));
		safe("im.token", function() return haxe.Serializer.run(haxe.Unserializer.run("q:3y1:vh")));

		var om = new haxe.ds.ObjectMap<Dynamic, Dynamic>();
		om.set(new ProbeClass(1), "v");
		safe("om.value", function() return haxe.Serializer.run(om));
		safe("om.token", function() return haxe.Serializer.run(haxe.Unserializer.run("Mcy10:ProbeClassy1:xi1gy1:vh")));

		safe("classref.token", function() return haxe.Serializer.run(haxe.Unserializer.run("Ay10:ProbeClass")));

		safe("enumref.token", function() return haxe.Serializer.run(haxe.Unserializer.run("By9:ProbeEnum")));

		safe("enum.j.replay.default", function() return haxe.Serializer.run(haxe.Unserializer.run("jy9:ProbeEnum:2:2i4y2:hi")));
		safe("enum.j.replay.index", function() {
			var s = new haxe.Serializer();
			s.useEnumIndex = true;
			s.serialize(haxe.Unserializer.run("jy9:ProbeEnum:2:2i4y2:hi"));
			return s.toString();
		});

		safe("unser.x", function() return haxe.Serializer.run(haxe.Unserializer.run("xi3")));

		safe("serializeException", function() {
			var s = new haxe.Serializer();
			try {
				throw "boom";
			} catch (e:Dynamic) {
				s.serializeException(e);
			}
			return s.toString();
		});
	}
}
