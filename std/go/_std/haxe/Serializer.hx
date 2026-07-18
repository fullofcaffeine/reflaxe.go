package haxe;

import haxe.ds.List;
import hxrt.serialization.NativeSerialization;

/**
	What:
	- Implements the complete Haxe 4.3.7 serialization token algorithm as staged
	  source for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because generated Go
	  instance fields are package-private and require one typed runtime snapshot
	  capability. Token choice, caches, recursive traversal, custom hooks, and wire
	  compatibility are ordinary Haxe library behavior, not compiler intrinsics.

	How:
	- Follow the upstream wire-format algorithm over `Type`, collections, Date,
	  Bytes, and enum APIs.
	- Delegate only erased generated-field access to `NativeSerialization`; its
	  unavoidable `Dynamic` values are immediately restored to the public
	  serialization traversal.
**/
class Serializer {
	public static var USE_CACHE = false;
	public static var USE_ENUM_INDEX = false;

	static var BASE64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:";

	var buf:StringBuf;
	var cache:Array<Dynamic>;
	var shash:haxe.ds.StringMap<Int>;
	var scount:Int;

	public var useCache:Bool;
	public var useEnumIndex:Bool;

	public function new() {
		buf = new StringBuf();
		cache = [];
		useCache = USE_CACHE;
		useEnumIndex = USE_ENUM_INDEX;
		shash = new haxe.ds.StringMap();
		scount = 0;
	}

	public function toString():String {
		return buf.toString();
	}

	function serializeString(value:String):Void {
		var known = shash.get(value);
		if (known != null) {
			buf.add("R");
			buf.add(known);
			return;
		}
		shash.set(value, scount++);
		buf.add("y");
		var encoded = StringTools.urlEncode(value);
		buf.add(encoded.length);
		buf.add(":");
		buf.add(encoded);
	}

	/** Uses portable Haxe reference equality; erased comparison is isolated in hxrt. **/
	function serializeRef(value:Dynamic):Bool {
		for (index in 0...cache.length) {
			if (cache[index] == value) {
				buf.add("r");
				buf.add(index);
				return true;
			}
		}
		cache.push(value);
		return false;
	}

	function serializeFields(value:Dynamic):Void {
		var fields = NativeSerialization.fields(value);
		for (index in 0...fields.length) {
			var field = fields[index];
			serializeString(field.name);
			serialize(field.value);
		}
		buf.add("g");
	}

	function serializeArray(value:Array<Dynamic>):Void {
		var nullCount = 0;
		buf.add("a");
		for (item in value) {
			if (item == null) {
				nullCount++;
				continue;
			}
			flushNulls(nullCount);
			nullCount = 0;
			serialize(item);
		}
		flushNulls(nullCount);
		buf.add("h");
	}

	inline function flushNulls(count:Int):Void {
		if (count == 1) {
			buf.add("n");
		} else if (count > 1) {
			buf.add("u");
			buf.add(count);
		}
	}

	function serializeBytes(value:haxe.io.Bytes):Void {
		buf.add("s");
		buf.add(Math.ceil((value.length * 8) / 6));
		buf.add(":");
		var index = 0;
		var max = value.length - 2;
		while (index < max) {
			var first = value.get(index++);
			var second = value.get(index++);
			var third = value.get(index++);
			buf.addChar(BASE64.charCodeAt(first >> 2));
			buf.addChar(BASE64.charCodeAt(((first << 4) | (second >> 4)) & 63));
			buf.addChar(BASE64.charCodeAt(((second << 2) | (third >> 6)) & 63));
			buf.addChar(BASE64.charCodeAt(third & 63));
		}
		if (index == max) {
			var first = value.get(index++);
			var second = value.get(index++);
			buf.addChar(BASE64.charCodeAt(first >> 2));
			buf.addChar(BASE64.charCodeAt(((first << 4) | (second >> 4)) & 63));
			buf.addChar(BASE64.charCodeAt((second << 2) & 63));
		} else if (index == max + 1) {
			var first = value.get(index++);
			buf.addChar(BASE64.charCodeAt(first >> 2));
			buf.addChar(BASE64.charCodeAt((first << 4) & 63));
		}
	}

	function serializeClass(value:Dynamic, declaration:Class<Dynamic>):Void {
		var className = Type.getClassName(declaration);
		if (className == "String") {
			serializeString(value);
			return;
		}
		if (useCache && serializeRef(value))
			return;

		switch (className) {
			case "Array":
				serializeArray(cast value);
			case "haxe.ds.List":
				buf.add("l");
				var list:List<Dynamic> = cast value;
				for (item in list)
					serialize(item);
				buf.add("h");
			case "Date":
				var date:Date = cast value;
				buf.add("v");
				buf.add(date.getTime());
			case "haxe.ds.StringMap":
				buf.add("b");
				var map:haxe.ds.StringMap<Dynamic> = cast value;
				for (key in map.keys()) {
					serializeString(key);
					serialize(map.get(key));
				}
				buf.add("h");
			case "haxe.ds.IntMap":
				buf.add("q");
				var map:haxe.ds.IntMap<Dynamic> = cast value;
				for (key in map.keys()) {
					buf.add(":");
					buf.add(key);
					serialize(map.get(key));
				}
				buf.add("h");
			case "haxe.ds.ObjectMap":
				buf.add("M");
				var map:haxe.ds.ObjectMap<Dynamic, Dynamic> = cast value;
				for (key in map.keys()) {
					serialize(key);
					serialize(map.get(key));
				}
				buf.add("h");
			case "haxe.io.Bytes":
				serializeBytes(cast value);
			default:
				if (useCache)
					cache.pop();
				if (GoSerializationBridge.hasSerializeHook(value)) {
					buf.add("C");
					serializeString(className);
					if (useCache)
						cache.push(value);
					GoSerializationBridge.callSerializeHook(value, this);
					buf.add("g");
				} else {
					buf.add("c");
					serializeString(className);
					if (useCache)
						cache.push(value);
					serializeFields(value);
				}
		}
	}

	function serializeEnum(value:Dynamic, declaration:Enum<Dynamic>):Void {
		if (useCache) {
			if (serializeRef(value))
				return;
			cache.pop();
		}
		buf.add(useEnumIndex ? "j" : "w");
		serializeString(Type.getEnumName(declaration));
		if (useEnumIndex) {
			buf.add(":");
			buf.add(Type.enumIndex(value));
		} else {
			serializeString(Type.enumConstructor(value));
		}
		buf.add(":");
		var parameters = Type.enumParameters(value);
		buf.add(parameters.length);
		for (parameter in parameters)
			serialize(parameter);
		if (useCache)
			cache.push(value);
	}

	public function serialize(value:Dynamic):Void {
		if (Std.isOfType(value, haxe.io.Bytes)) {
			if (useCache && serializeRef(value))
				return;
			serializeBytes(cast value);
			return;
		}
		switch (Type.typeof(value)) {
			case TNull:
				buf.add("n");
			case TInt:
				var integer:Int = value;
				if (integer == 0)
					buf.add("z");
				else {
					buf.add("i");
					buf.add(integer);
				}
			case TFloat:
				var number:Float = value;
				if (Math.isNaN(number))
					buf.add("k");
				else if (!Math.isFinite(number))
					buf.add(number < 0 ? "m" : "p");
				else {
					buf.add("d");
					buf.add(number);
				}
			case TBool:
				var boolean:Bool = value;
				buf.add(boolean ? "t" : "f");
			case TClass(declaration):
				serializeClass(value, declaration);
			case TObject:
				if (Std.isOfType(value, Class)) {
					buf.add("A");
					serializeString(Type.getClassName(value));
				} else if (Std.isOfType(value, Enum)) {
					buf.add("B");
					serializeString(Type.getEnumName(value));
				} else {
					if (useCache && serializeRef(value))
						return;
					buf.add("o");
					serializeFields(value);
				}
			case TEnum(declaration):
				serializeEnum(value, declaration);
			case TFunction:
				throw "Cannot serialize function";
			case TUnknown:
				throw "Cannot serialize " + Std.string(value);
		}
	}

	public function serializeException(value:Dynamic):Void {
		buf.add("x");
		serialize(value);
	}

	public static function run(value:Dynamic):String {
		var serializer = new Serializer();
		serializer.serialize(value);
		return serializer.toString();
	}
}
