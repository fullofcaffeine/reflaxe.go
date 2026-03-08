class Main {
	static function main() {
		var sample = "h\u00E9";
		var raw = haxe.io.Bytes.ofString(sample, haxe.io.Encoding.RawNative);
		Sys.println("raw.len=" + raw.length);
		Sys.println("raw.hex=" + raw.toHex());
		Sys.println("raw.get=" + raw.getString(0, raw.length, haxe.io.Encoding.RawNative));
		Sys.println("raw.base64.before=" + haxe.crypto.Base64.encode(raw));
		raw.set(0, "A".code);
		Sys.println("raw.hex.after=" + raw.toHex());
		Sys.println("raw.get.after=" + raw.getString(0, raw.length, haxe.io.Encoding.RawNative));
		Sys.println("raw.base64.after=" + haxe.crypto.Base64.encode(raw));

		var output = new haxe.io.BytesOutput();
		output.writeString(sample, haxe.io.Encoding.RawNative);
		var written = output.getBytes();
		Sys.println("out.hex=" + written.toHex());
		var input = new haxe.io.BytesInput(written);
		Sys.println("in.raw=" + input.readString(written.length, haxe.io.Encoding.RawNative));
	}
}
