class Main {
	static function bytesHex(value:haxe.io.Bytes):String {
		var out = new StringBuf();
		for (i in 0...value.length) {
			if (i > 0) {
				out.add(",");
			}
			out.add(value.get(i));
		}
		return out.toString();
	}

	static function main() {
		var source = haxe.io.Bytes.ofString("abc");
		var data = source.getData();
		data[1] = 'Z'.code;
		Sys.println("source1=" + source.toString());

		var alias = haxe.io.Bytes.ofData(data);
		alias.set(2, '!'.code);
		Sys.println("source2=" + source.toString());
		Sys.println("alias2=" + alias.toString());

		data[0] = 'Q'.code;
		Sys.println("source3=" + source.toString());
		Sys.println("alias3=" + alias.toString());
		Sys.println("hex=" + bytesHex(alias));
	}
}
