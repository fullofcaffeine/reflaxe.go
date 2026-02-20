enum FlagEnum {
	NoArgs;
	WithArgs(a:Int, b:String);
}

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (_:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function withFlags(useCache:Bool, useEnumIndex:Bool, fn:Void->String):String {
		var prevCache = haxe.Serializer.USE_CACHE;
		var prevEnumIndex = haxe.Serializer.USE_ENUM_INDEX;
		haxe.Serializer.USE_CACHE = useCache;
		haxe.Serializer.USE_ENUM_INDEX = useEnumIndex;
		var out = fn();
		haxe.Serializer.USE_CACHE = prevCache;
		haxe.Serializer.USE_ENUM_INDEX = prevEnumIndex;
		return out;
	}

	static function cachePayload():Array<Dynamic> {
		var shared = {v: 7};
		return [shared, shared];
	}

	static function main() {
		safe("global.defaults", function() {
			return Std.string(haxe.Serializer.USE_CACHE) + "," + Std.string(haxe.Serializer.USE_ENUM_INDEX);
		});

		safe("global.cache.off.run", function() {
			return withFlags(false, false, function() return haxe.Serializer.run(cachePayload()));
		});

		safe("global.cache.on.run", function() {
			return withFlags(true, false, function() return haxe.Serializer.run(cachePayload()));
		});

		safe("global.cache.instance.new", function() {
			return withFlags(true, false, function() {
				var s = new haxe.Serializer();
				s.serialize(cachePayload());
				return s.toString();
			});
		});

		safe("global.cache.instance.old", function() {
			var s = new haxe.Serializer();
			return withFlags(true, false, function() {
				s.serialize(cachePayload());
				return s.toString();
			});
		});

		safe("global.enumIndex.off.run", function() {
			return withFlags(false, false, function() return haxe.Serializer.run(FlagEnum.WithArgs(2, "x")));
		});

		safe("global.enumIndex.on.run", function() {
			return withFlags(false, true, function() return haxe.Serializer.run(FlagEnum.WithArgs(2, "x")));
		});

		safe("global.exception.flags", function() {
			return withFlags(true, true, function() {
				var shared = {msg: "boom"};
				var s = new haxe.Serializer();
				s.serializeException(shared);
				s.serializeException(shared);
				return s.toString();
			});
		});

		safe("global.restored", function() {
			return Std.string(haxe.Serializer.USE_CACHE) + "," + Std.string(haxe.Serializer.USE_ENUM_INDEX);
		});
	}
}
