class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		safe("mix.replay.cacheOn", function() {
			var parsed = haxe.Unserializer.run("aoy1:sy1:agR0R1r0h");
			var s = new haxe.Serializer();
			s.useCache = true;
			s.serialize(parsed);
			return s.toString();
		});

		safe("mix.sequential", function() {
			var u = new haxe.Unserializer("y1:xy1:yR0R1");
			var a = u.unserialize();
			var b = u.unserialize();
			var c = u.unserialize();
			var d = u.unserialize();
			return haxe.Serializer.run([a, b, c, d]);
		});
	}
}
