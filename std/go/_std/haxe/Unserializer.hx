package haxe;

import haxe.ds.List;
import hxrt.serialization.NativeSerialization;

@:noDoc
typedef TypeResolver = {
	function resolveClass(name:String):Class<Dynamic>;
	function resolveEnum(name:String):Enum<Dynamic>;
}

/**
	What:
	- Implements the complete Haxe 4.3.7 unserialization token algorithm as staged
	  source for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because decoded
	  values must initialize package-private generated Go fields and repair the
	  constructor-bypassed virtual-dispatch self pointer.
	- Cursor behavior, resolver policy, token parsing, cache order, and custom hook
	  dispatch are Haxe library semantics and must not originate in a compiler
	  declaration emitter.

	How:
	- Parse and construct all portable tokens in ordinary Haxe using the existing
	  compiler-generated `Type` metadata API.
	- Delegate only reflected field assignment and hidden-self initialization to
	  the narrow typed `NativeSerialization` boundary.
**/
class Unserializer {
	public static var DEFAULT_RESOLVER:TypeResolver = {
		resolveClass: function(name:String):Class<Dynamic> return Type.resolveClass(name),
		resolveEnum: function(name:String):Enum<Dynamic> return Type.resolveEnum(name)
	};

	static var NULL_RESOLVER:TypeResolver = {
		resolveClass: function(_name:String):Class<Dynamic> return null,
		resolveEnum: function(_name:String):Enum<Dynamic> return null
	};

	static var BASE64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:";

	var buf:String;
	var pos:Int;
	var length:Int;
	var cache:Array<Dynamic>;
	var scache:Array<String>;
	var resolver:Dynamic;

	public function new(buf:String) {
		this.buf = buf;
		length = buf.length;
		pos = 0;
		scache = [];
		cache = [];
		var current = DEFAULT_RESOLVER;
		if (current == null) {
			current = {
				resolveClass: function(name:String):Class<Dynamic> return Type.resolveClass(name),
				resolveEnum: function(name:String):Enum<Dynamic> return Type.resolveEnum(name)
			};
			DEFAULT_RESOLVER = current;
		}
		resolver = current;
	}

	public function setResolver(value:Dynamic):Void {
		resolver = value == null ? NULL_RESOLVER : value;
	}

	public function getResolver():Dynamic {
		return resolver;
	}

	inline function get(index:Int):Int {
		return StringTools.fastCodeAt(buf, index);
	}

	function readDigits():Int {
		var value = 0;
		var negative = false;
		var start = pos;
		while (true) {
			var code = get(pos);
			if (StringTools.isEof(code))
				break;
			if (code == "-".code) {
				if (pos != start)
					break;
				negative = true;
				pos++;
				continue;
			}
			if (code < "0".code || code > "9".code)
				break;
			value = value * 10 + code - "0".code;
			pos++;
		}
		return negative ? -value : value;
	}

	function readFloat():Float {
		var start = pos;
		while (true) {
			var code = get(pos);
			if (StringTools.isEof(code))
				break;
			if ((code >= 43 && code < 58) || code == "e".code || code == "E".code)
				pos++;
			else
				break;
		}
		return NativeSerialization.parseFloat(buf.substr(start, pos - start));
	}

	function unserializeObject(target:Dynamic):Void {
		while (true) {
			if (pos >= length)
				throw "Invalid object";
			if (get(pos) == "g".code)
				break;
			var key:Dynamic = unserialize();
			if (!Std.isOfType(key, String))
				throw "Invalid object key";
			NativeSerialization.setField(target, cast key, unserialize());
		}
		pos++;
	}

	function unserializeEnum<T>(declaration:Enum<T>, tag:String):T {
		if (get(pos++) != ":".code)
			throw "Invalid enum format";
		var count = readDigits();
		if (count == 0)
			return Type.createEnum(declaration, tag, []);
		var arguments = new Array<Dynamic>();
		while (count-- > 0)
			arguments.push(unserialize());
		return Type.createEnum(declaration, tag, arguments);
	}

	function decodeBytes(encoded:String):haxe.io.Bytes {
		var rest = encoded.length & 3;
		var size = (encoded.length >> 2) * 3 + (rest >= 2 ? rest - 1 : 0);
		var bytes = haxe.io.Bytes.alloc(size);
		var index = 0;
		var output = 0;
		var complete = encoded.length - rest;
		while (index < complete) {
			var first = base64Value(encoded.charCodeAt(index++));
			var second = base64Value(encoded.charCodeAt(index++));
			var third = base64Value(encoded.charCodeAt(index++));
			var fourth = base64Value(encoded.charCodeAt(index++));
			bytes.set(output++, (first << 2) | (second >> 4));
			bytes.set(output++, (second << 4) | (third >> 2));
			bytes.set(output++, (third << 6) | fourth);
		}
		if (rest >= 2) {
			var first = base64Value(encoded.charCodeAt(index++));
			var second = base64Value(encoded.charCodeAt(index++));
			bytes.set(output++, (first << 2) | (second >> 4));
			if (rest == 3) {
				var third = base64Value(encoded.charCodeAt(index++));
				bytes.set(output++, (second << 4) | (third >> 2));
			}
		}
		return bytes;
	}

	function base64Value(code:Int):Int {
		for (index in 0...BASE64.length)
			if (BASE64.charCodeAt(index) == code)
				return index;
		return -1;
	}

	public function unserialize():Dynamic {
		switch (get(pos++)) {
			case "n".code:
				return null;
			case "t".code:
				return true;
			case "f".code:
				return false;
			case "z".code:
				return 0;
			case "i".code:
				return readDigits();
			case "d".code:
				return readFloat();
			case "y".code:
				var stringLength = readDigits();
				if (get(pos++) != ":".code || length - pos < stringLength)
					throw "Invalid string length";
				var value = StringTools.urlDecode(buf.substr(pos, stringLength));
				pos += stringLength;
				scache.push(value);
				return value;
			case "k".code:
				return Math.NaN;
			case "m".code:
				return Math.NEGATIVE_INFINITY;
			case "p".code:
				return Math.POSITIVE_INFINITY;
			case "a".code:
				var array = new Array<Dynamic>();
				cache.push(array);
				while (true) {
					var code = get(pos);
					if (code == "h".code) {
						pos++;
						break;
					}
					if (code == "u".code) {
						pos++;
						var count = readDigits();
						array[array.length + count - 1] = null;
					} else {
						array.push(unserialize());
					}
				}
				return array;
			case "o".code:
				var object:Dynamic = {};
				cache.push(object);
				unserializeObject(object);
				return object;
			case "r".code:
				var index = readDigits();
				if (index < 0 || index >= cache.length)
					throw "Invalid reference";
				return cache[index];
			case "R".code:
				var index = readDigits();
				if (index < 0 || index >= scache.length)
					throw "Invalid string reference";
				return scache[index];
			case "x".code:
				throw unserialize();
			case "c".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveClass(resolver, name);
				if (declaration == null)
					throw "Class not found " + name;
				var object = Type.createEmptyInstance(declaration);
				NativeSerialization.bindSelf(object);
				cache.push(object);
				unserializeObject(object);
				return object;
			case "w".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveEnum(resolver, name);
				if (declaration == null)
					throw "Enum not found " + name;
				var value = unserializeEnum(declaration, unserialize());
				cache.push(value);
				return value;
			case "j".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveEnum(resolver, name);
				if (declaration == null)
					throw "Enum not found " + name;
				pos++;
				var index = readDigits();
				var tag = Type.getEnumConstructs(declaration)[index];
				if (tag == null)
					throw "Unknown enum index " + name + "@" + index;
				var value = unserializeEnum(declaration, tag);
				cache.push(value);
				return value;
			case "l".code:
				var list = new List<Dynamic>();
				cache.push(list);
				while (get(pos) != "h".code)
					list.add(unserialize());
				pos++;
				return list;
			case "b".code:
				var map = new haxe.ds.StringMap<Dynamic>();
				cache.push(map);
				while (get(pos) != "h".code) {
					var key:String = unserialize();
					map.set(key, unserialize());
				}
				pos++;
				return map;
			case "q".code:
				var map = new haxe.ds.IntMap<Dynamic>();
				cache.push(map);
				var code = get(pos++);
				while (code == ":".code) {
					map.set(readDigits(), unserialize());
					code = get(pos++);
				}
				if (code != "h".code)
					throw "Invalid IntMap format";
				return map;
			case "M".code:
				var map = new haxe.ds.ObjectMap<Dynamic, Dynamic>();
				cache.push(map);
				while (get(pos) != "h".code)
					map.set(unserialize(), unserialize());
				pos++;
				return map;
			case "v".code:
				var date:Date;
				if (isLegacyDate()) {
					date = Date.fromString(buf.substr(pos, 19));
					pos += 19;
				} else {
					date = Date.fromTime(readFloat());
				}
				cache.push(date);
				return date;
			case "s".code:
				var bytesLength = readDigits();
				if (get(pos++) != ":".code || length - pos < bytesLength)
					throw "Invalid bytes length";
				var bytes = decodeBytes(buf.substr(pos, bytesLength));
				pos += bytesLength;
				cache.push(bytes);
				return bytes;
			case "C".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveClass(resolver, name);
				if (declaration == null)
					throw "Class not found " + name;
				var object:Dynamic = Type.createEmptyInstance(declaration);
				NativeSerialization.bindSelf(object);
				cache.push(object);
				if (!GoSerializationBridge.callUnserializeHook(object, this))
					throw "Class " + name + " has no hxUnserialize hook";
				if (get(pos++) != "g".code)
					throw "Invalid custom data";
				return object;
			case "A".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveClass(resolver, name);
				if (declaration == null)
					throw "Class not found " + name;
				return declaration;
			case "B".code:
				var name:String = unserialize();
				var declaration = GoSerializationBridge.resolveEnum(resolver, name);
				if (declaration == null)
					throw "Enum not found " + name;
				return declaration;
			default:
		}
		pos--;
		throw "Invalid char " + buf.charAt(pos) + " at position " + pos;
	}

	function isLegacyDate():Bool {
		for (offset in 0...4) {
			var code = get(pos + offset);
			if (code < "0".code || code > "9".code)
				return false;
		}
		return get(pos + 4) == "-".code;
	}

	public static function run(value:String):Dynamic {
		return new Unserializer(value).unserialize();
	}
}
