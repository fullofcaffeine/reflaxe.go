class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		safe("tok.null", function() return haxe.Serializer.run(null));
		safe("tok.true", function() return haxe.Serializer.run(true));
		safe("tok.false", function() return haxe.Serializer.run(false));
		safe("tok.zero", function() return haxe.Serializer.run(0));
		safe("tok.int", function() return haxe.Serializer.run(-7));
		safe("tok.float", function() return haxe.Serializer.run(1.5));
		safe("tok.nan", function() return haxe.Serializer.run(haxe.Unserializer.run("k")));
		safe("tok.posInf", function() return haxe.Serializer.run(haxe.Unserializer.run("p")));
		safe("tok.negInf", function() return haxe.Serializer.run(haxe.Unserializer.run("m")));
		safe("tok.negZero", function() return haxe.Serializer.run(haxe.Unserializer.run("d-0")));
		safe("tok.stringEsc", function() return haxe.Serializer.run("a b:c%\n"));
		safe("tok.array", function() return haxe.Serializer.run([null, null, 1, null]));
		safe("tok.object", function() return haxe.Serializer.run({b: 1, a: 2}));

		var shared = {x: 1};
		var withCache = new haxe.Serializer();
		withCache.useCache = true;
		withCache.serialize([shared, shared]);
		safe("tok.cacheOn", function() return withCache.toString());

		var withoutCache = new haxe.Serializer();
		withoutCache.serialize([shared, shared]);
		safe("tok.cacheOff", function() return withoutCache.toString());

		var decoder = new haxe.Unserializer("i1i2");
		safe("unseq.1", function() return Std.string(decoder.unserialize()));
		safe("unseq.2", function() return Std.string(decoder.unserialize()));
		safe("unseq.3", function() return Std.string(decoder.unserialize()));

		safe("run.single", function() return Std.string(haxe.Unserializer.run("i1i2")));
		var parsedRef:Dynamic = haxe.Unserializer.run("aoy1:xi1gr1h");
		safe("ref.roundtrip", function() return haxe.Serializer.run(parsedRef));
		var parsedRefWithCache = new haxe.Serializer();
		parsedRefWithCache.useCache = true;
		parsedRefWithCache.serialize(parsedRef);
		safe("ref.cacheOn", function() return parsedRefWithCache.toString());
	}
}
