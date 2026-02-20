class Main {
	static function consumeInput(i:haxe.io.Input):String {
		i.bigEndian = true;
		var head = i.readInt16();
		var line = i.readLine();
		return head + "|" + line;
	}

	static function copyInputToOutput(i:haxe.io.Input, o:haxe.io.Output):Void {
		o.writeInput(i, 2);
	}

	static function main() {
		var out = new haxe.io.BytesOutput();
		out.bigEndian = true;
		out.writeInt16(0x0102);
		out.writeString("hello\r\n");
		out.writeInt8(33);
		var input = new haxe.io.BytesInput(out.getBytes());
		Sys.println("consume=" + consumeInput(input));
		Sys.println("tail=" + input.readInt8());

		var source = new haxe.io.BytesInput(haxe.io.Bytes.ofString("abcd"));
		var sink = new haxe.io.BytesOutput();
		copyInputToOutput(source, sink);
		Sys.println("copy=" + sink.getBytes().toString());

		var readAllInput = new haxe.io.BytesInput(haxe.io.Bytes.ofString("aa\nbb"));
		Sys.println("until=" + readAllInput.readUntil('\n'.code));
		Sys.println("all=" + readAllInput.readAll().toString());

		var readStringInput = new haxe.io.BytesInput(haxe.io.Bytes.ofString("XYZ"));
		Sys.println("str=" + readStringInput.readString(2));

		var partial = new haxe.io.BytesInput(haxe.io.Bytes.ofString("q"));
		try {
			partial.readFullBytes(haxe.io.Bytes.alloc(2), 0, 2);
			Sys.println("full=miss");
		} catch (_:Dynamic) {
			Sys.println("full=ok");
		}

		var floatOut = new haxe.io.BytesOutput();
		floatOut.bigEndian = false;
		floatOut.writeFloat(1.5);
		floatOut.writeDouble(2.25);
		var floatIn = new haxe.io.BytesInput(floatOut.getBytes());
		floatIn.bigEndian = false;
		Sys.println("float=" + floatIn.readFloat());
		Sys.println("double=" + floatIn.readDouble());

		var numOut = new haxe.io.BytesOutput();
		numOut.bigEndian = true;
		numOut.writeInt24(-2);
		numOut.writeUInt24(0x123456);
		var numIn = new haxe.io.BytesInput(numOut.getBytes());
		numIn.bigEndian = true;
		Sys.println("i24=" + numIn.readInt24());
		Sys.println("u24=" + numIn.readUInt24());

		var overflow = new haxe.io.BytesOutput();
		try {
			overflow.writeInt8(500);
			Sys.println("overflow=miss");
		} catch (_:Dynamic) {
			Sys.println("overflow=ok");
		}
	}
}
