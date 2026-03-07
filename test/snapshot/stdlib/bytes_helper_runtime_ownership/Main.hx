class Main {
	static function main() {
		var parsed = haxe.io.Bytes.ofHex("0fDA");
		Sys.println(parsed.toHex());

		var buffer = new haxe.io.BytesBuffer();
		buffer.addByte(260);
		buffer.addBytes(parsed, 1, 1);
		buffer.addString("Z");
		var out = buffer.getBytes();
		Sys.println(out.toHex());
	}
}
