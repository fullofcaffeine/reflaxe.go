import haxe.io.Bytes;
import haxe.io.Encoding;

class Main {
	static function nullableCode(value:Null<Int>):String {
		return value == null ? "null" : Std.string(value);
	}

	static function validationBits(values:Array<String>):String {
		var results = [];
		for (value in values) {
			results.push(UnicodeString.validate(Bytes.ofHex(value), Encoding.UTF8) ? "1" : "0");
		}
		return results.join("");
	}

	static function rawNativeError():String {
		try {
			UnicodeString.validate(Bytes.ofString("ok"), Encoding.RawNative);
		} catch (error:Dynamic) {
			return Std.string(error);
		}
		return "missing";
	}

	static function main() {
		var value:UnicodeString = "a😀bé😀a";
		Sys.println("length=" + value.length);
		Sys.println("chars="
			+ value.charAt(-1)
			+ "|"
			+ value.charAt(0)
			+ "|"
			+ value.charAt(1)
			+ "|"
			+ value.charAt(6)
			+ "|"
			+ value.charAt(7));
		Sys.println("codes="
			+ nullableCode(value.charCodeAt(-1))
			+ "|"
			+ nullableCode(value.charCodeAt(1))
			+ "|"
			+ nullableCode(value.charCodeAt(7)));

		Sys.println("substring=" + value.substring(1, 5) + "|" + value.substring(5, 1) + "|" + value.substring(-3, 2) + "|" + value.substring(2, -1) + "|"
			+ value.substring(3, 3) + "|" + value.substring(20, 30) + "|" + value.substring(4));
		Sys.println("substr=" + value.substr(1, 3) + "|" + value.substr(-2, 2) + "|" + value.substr(-20, 2) + "|" + value.substr(20, 2) + "|"
			+ value.substr(2, 0) + "|" + value.substr(4));

		Sys.println("index=" + value.indexOf("😀") + "|" + value.indexOf("😀", 2) + "|" + value.indexOf("bé", -4) + "|" + value.indexOf("", 3) + "|"
			+ value.indexOf("", 99) + "|" + value.indexOf("missing"));
		Sys.println("overlap=" + ("aaab" : UnicodeString).indexOf("aab") + "|" + ("a😀a😀a" : UnicodeString).indexOf("a😀a", 1));
		Sys.println("last=" + value.lastIndexOf("😀") + "|" + value.lastIndexOf("😀", 1) + "|" + value.lastIndexOf("", 3) + "|" + value.lastIndexOf("", 99)
			+ "|" + value.lastIndexOf("a", -1) + "|" + value.lastIndexOf("missing"));
		Sys.println("lastOverlap=" + ("aaab" : UnicodeString).lastIndexOf("aab") + "|" + ("a😀a😀a" : UnicodeString).lastIndexOf("a😀a"));

		var iterated = [];
		for (code in value) {
			iterated.push(code);
		}
		Sys.println("iterator=" + iterated.join(","));
		var keyed = [];
		for (entry in value.keyValueIterator()) {
			keyed.push(entry.key + ":" + entry.value);
		}
		Sys.println("keyed=" + keyed.join(","));

		var left:UnicodeString = "a😀";
		var right:UnicodeString = "bé";
		Sys.println("operators=" + (left < right) + "|" + (left <= right) + "|" + (right > left) + "|" + (right >= left) + "|" + (left == left) + "|"
			+ (left != right) + "|" + (left + right));
		Sys.println("mixed=" + (left + "x") + "|" + ("x" + left));
		var assigned:UnicodeString = left;
		assigned += right;
		assigned += "x";
		Sys.println("assigned=" + assigned);

		Sys.println("valid=" + validationBits(["", "00", "7f", "c2a2", "e282ac", "f09f9880", "f48fbfbf"]));
		Sys.println("invalid=" + validationBits([
			"80",
			"c0af",
			"c2",
			"c220",
			"e08080",
			"eda080",
			"e282",
			"e28220",
			"f0808080",
			"f4908080",
			"f09f98",
			"f09f9820",
			"f5"
		]));
		Sys.println("raw=" + rawNativeError());
	}
}
