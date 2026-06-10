import haxe.io.Bytes;

class Main {
	static function join(values:Array<String>):String {
		var out = "";
		for (index in 0...values.length) {
			if (index > 0)
				out += "|";
			out += values[index];
		}
		return out;
	}

	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=" + fn());
		} catch (error:Dynamic) {
			Sys.println(label + "=!" + Std.string(error));
		}
	}

	static function main() {
		var utf8 = new haxe.Utf8();
		utf8.addChar('a'.code);
		utf8.addChar(0x1F600);
		utf8.addChar('é'.code);
		var value = utf8.toString();
		var sized = new haxe.Utf8(4);
		sized.addChar('z'.code);
		var iterated = [];
		haxe.Utf8.iter(value, function(code) iterated.push(Std.string(code)));

		Sys.println("string=" + value);
		Sys.println("sized=" + sized.toString());
		Sys.println("iter=" + join(iterated));
		Sys.println("charCodeAt.0=" + haxe.Utf8.charCodeAt(value, 0));
		Sys.println("charCodeAt.1=" + haxe.Utf8.charCodeAt(value, 1));
		Sys.println("validate=" + haxe.Utf8.validate(value));
		Sys.println("length=" + haxe.Utf8.length(value));
		Sys.println("compare.eq=" + haxe.Utf8.compare(value, value));
		Sys.println("compare.lt=" + haxe.Utf8.compare("abc", "abd"));
		Sys.println("sub=" + haxe.Utf8.sub(value, 1, 2));
		Sys.println("encode.ascii.hex=" + Bytes.ofString(haxe.Utf8.encode("abc")).toHex());
		Sys.println("encode.eacute.hex=" + Bytes.ofString(haxe.Utf8.encode("é")).toHex());
		Sys.println("decode.ascii.hex=" + Bytes.ofString(haxe.Utf8.decode("abc")).toHex());
		Sys.println("decode.eacute.hex=" + Bytes.ofString(haxe.Utf8.decode("é")).toHex());
		Sys.println("decode.roundtrip.hex=" + Bytes.ofString(haxe.Utf8.decode(haxe.Utf8.encode("é"))).toHex());
		Sys.println("decode.emoji.roundtrip.hex=" + Bytes.ofString(haxe.Utf8.decode(haxe.Utf8.encode("😀"))).toHex());
	}
}
