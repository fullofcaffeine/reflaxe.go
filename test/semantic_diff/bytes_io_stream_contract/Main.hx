class Main {
	static function main() {
		var source = haxe.io.Bytes.ofString("abcdef");

		try {
			new haxe.io.BytesInput(source, 5, 3);
			Sys.println("ctorBounds=miss");
		} catch (_:Dynamic) {
			Sys.println("ctorBounds=ok");
		}

		var input = new haxe.io.BytesInput(source, 1, 4);
		Sys.println("len=" + input.length);
		Sys.println("pos=" + input.position);
		Sys.println("byte=" + input.readByte());

		var chunk = haxe.io.Bytes.alloc(3);
		var read = input.readBytes(chunk, 0, 3);
		Sys.println("read=" + read + ":" + chunk.toString());

		try {
			input.readByte();
			Sys.println("eof=miss");
		} catch (_:haxe.io.Eof) {
			Sys.println("eof=ok");
		}

		input.position = -2;
		Sys.println("posLow=" + input.position);
		input.position = 99;
		Sys.println("posHigh=" + input.position);

		try {
			input.readBytes(haxe.io.Bytes.alloc(2), 2, 1);
			Sys.println("readBounds=miss");
		} catch (_:Dynamic) {
			Sys.println("readBounds=ok");
		}

		var partialInput = new haxe.io.BytesInput(haxe.io.Bytes.ofString("xy"));
		var partialBuf = haxe.io.Bytes.alloc(4);
		var partialRead = partialInput.readBytes(partialBuf, 1, 3);
		Sys.println("partial=" + partialRead + ":" + partialBuf.get(1) + "," + partialBuf.get(2));

		var output = new haxe.io.BytesOutput();
		try {
			output.writeBytes(haxe.io.Bytes.ofString("x"), 2, 1);
			Sys.println("writeBounds=miss");
		} catch (_:Dynamic) {
			Sys.println("writeBounds=ok");
		}
		output.writeByte(65);
		output.writeBytes(haxe.io.Bytes.ofString("bcdef"), 1, 3);
		Sys.println("outLen=" + output.length);
		Sys.println("out=" + output.getBytes().toString());
	}
}
