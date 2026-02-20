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

	static function encodingName(value:haxe.io.Encoding):String {
		return switch (value) {
			case UTF8:
				"UTF8";
			case RawNative:
				"RawNative";
		};
	}

	static function main() {
		var sample = "h\u00E9";
		var utf8 = haxe.io.Bytes.ofString(sample, haxe.io.Encoding.UTF8);
		var raw = haxe.io.Bytes.ofString(sample, haxe.io.Encoding.RawNative);

		Sys.println("enum.utf8=" + encodingName(haxe.io.Encoding.UTF8));
		Sys.println("enum.raw=" + encodingName(haxe.io.Encoding.RawNative));

		Sys.println("utf8.len=" + utf8.length + " hex=" + bytesHex(utf8));
		Sys.println("raw.len=" + raw.length + " hex=" + bytesHex(raw));
		Sys.println("utf8.get=" + utf8.getString(0, utf8.length, haxe.io.Encoding.UTF8));
		Sys.println("raw.get=" + raw.getString(0, raw.length, haxe.io.Encoding.RawNative));

		var output = new haxe.io.BytesOutput();
		output.writeString(sample, haxe.io.Encoding.RawNative);
		var written = output.getBytes();
		var input = new haxe.io.BytesInput(written);
		Sys.println("read.raw=" + input.readString(written.length, haxe.io.Encoding.RawNative));

		try {
			utf8.getString(-1, 1, haxe.io.Encoding.UTF8);
			Sys.println("bounds=miss");
		} catch (e:haxe.io.Error) {
			switch (e) {
				case OutsideBounds:
					Sys.println("bounds=ok");
				default:
					Sys.println("bounds=bad");
			}
		}
	}
}
