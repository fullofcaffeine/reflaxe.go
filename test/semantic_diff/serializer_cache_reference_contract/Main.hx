class Node {
	public var name:String;
	public var next:Node;

	public function new(name:String) {
		this.name = name;
	}
}

class Custom {
	public var x:Dynamic;

	public function new(?x:Int) {
		this.x = x == null ? 0 : x;
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

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		safe("cache.enum.same", function() {
			var e = ProbeEnum.One(1);
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize([e, e]);
			return s.toString();
		});

		safe("cache.enum.distinct", function() {
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize([ProbeEnum.One(1), ProbeEnum.One(1)]);
			return s.toString();
		});

		safe("cache.class.same", function() {
			var n = new Node("a");
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize([n, n]);
			return s.toString();
		});

		safe("cache.class.custom.same", function() {
			var c = new Custom(3);
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize([c, c]);
			return s.toString();
		});

		safe("cache.class.cycle", function() {
			var n = new Node("loop");
			n.next = n;
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize(n);
			return s.toString();
		});

		safe("cache.unser.enum.replay", function() {
			return haxe.Serializer.run(haxe.Unserializer.run("awy9:ProbeEnumy3:One:1i2r1h"));
		});

		safe("cache.unser.custom.replay", function() {
			return haxe.Serializer.run(haxe.Unserializer.run("acy6:Customi4gr1h"));
		});
	}
}
