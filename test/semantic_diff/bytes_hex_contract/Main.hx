class Main {
	static function bytesHex(b:haxe.io.Bytes):String {
		var out = new StringBuf();
		for (i in 0...b.length) {
			if (i > 0) {
				out.add(",");
			}
			out.add(b.get(i));
		}
		return out.toString();
	}

	static function main() {
		var bytes = haxe.io.Bytes.ofString("Az");
		Sys.println("toHex=" + bytes.toHex());
		Sys.println("ofHex.upper=" + haxe.io.Bytes.ofHex("0FDA").toHex());
		Sys.println("ofHex.invalid=" + bytesHex(haxe.io.Bytes.ofHex("gg")));
		try {
			haxe.io.Bytes.ofHex("abc");
			Sys.println("odd=miss");
		} catch (e:Dynamic) {
			Sys.println("odd=" + Std.string(e));
		}
	}
}
