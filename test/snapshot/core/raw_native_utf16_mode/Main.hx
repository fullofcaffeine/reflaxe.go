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
		var sample = "h\u00E9";
		var utf8 = haxe.io.Bytes.ofString(sample, haxe.io.Encoding.UTF8);
		var rawNative = haxe.io.Bytes.ofString(sample, haxe.io.Encoding.RawNative);
		Sys.println("utf8.len=" + utf8.length + " hex=" + bytesHex(utf8));
		Sys.println("raw.len=" + rawNative.length + " hex=" + bytesHex(rawNative));
		Sys.println("raw.get=" + rawNative.getString(0, rawNative.length, haxe.io.Encoding.RawNative));

		var output = new haxe.io.BytesOutput();
		output.writeString(sample, haxe.io.Encoding.RawNative);
		var written = output.getBytes();
		Sys.println("out.raw.hex=" + bytesHex(written));

		var input = new haxe.io.BytesInput(written);
		Sys.println("in.raw=" + input.readString(written.length, haxe.io.Encoding.RawNative));
	}
}
