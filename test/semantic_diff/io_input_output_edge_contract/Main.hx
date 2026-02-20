class ZeroInput extends haxe.io.Input {
	public function new() {}

	override public function readByte():Int {
		return throw new haxe.io.Eof();
	}

	override public function readBytes(buf:haxe.io.Bytes, pos:Int, len:Int):Int {
		return 0;
	}
}

class Main {
	static function main() {
		var empty = new haxe.io.BytesInput(haxe.io.Bytes.ofString(""));
		try {
			empty.readLine();
			Sys.println("lineEmpty=miss");
		} catch (_:haxe.io.Eof) {
			Sys.println("lineEmpty=eof");
		}

		var tail = new haxe.io.BytesInput(haxe.io.Bytes.ofString("tail"));
		Sys.println("lineTail=" + tail.readLine());
		try {
			tail.readLine();
			Sys.println("lineTailSecond=miss");
		} catch (_:haxe.io.Eof) {
			Sys.println("lineTailSecond=eof");
		}

		var crlf = new haxe.io.BytesInput(haxe.io.Bytes.ofString("value\r\nnext"));
		Sys.println("lineCrlf=" + crlf.readLine());
		Sys.println("lineNext=" + crlf.readLine());

		var out = new haxe.io.BytesOutput();
		try {
			out.writeInput(new ZeroInput(), 1);
			Sys.println("writeInputBlocked=miss");
		} catch (_:Dynamic) {
			Sys.println("writeInputBlocked=ok");
		}
	}
}
