class Main {
	static function main() {
		var lines = new haxe.io.BytesInput(haxe.io.Bytes.ofString("first\r\nsecond"));
		Sys.println(lines.readLine());
		Sys.println(lines.readLine());

		var replay = new haxe.io.BytesInput(haxe.io.Bytes.ofString("012345"));
		var buf = haxe.io.Bytes.alloc(4);
		replay.readFullBytes(buf, 0, 4);
		Sys.println(buf.toString());

		var copyOut = new haxe.io.BytesOutput();
		copyOut.writeInput(new haxe.io.BytesInput(haxe.io.Bytes.ofString("xy")));
		Sys.println(copyOut.getBytes().toString());

		var all = new haxe.io.BytesInput(haxe.io.Bytes.ofString("zz")).readAll();
		Sys.println(all.toString());
	}
}
