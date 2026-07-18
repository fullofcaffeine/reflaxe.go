package haxe.io;

/**
	What: Collects one Output stream into an in-memory Bytes buffer.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	until Output and BytesBuffer have target definitions. This override composes
	those staged APIs without a compiler-emitted duplicate.

	How: Override primitive writes to append to `BytesBuffer`, then expose the
	completed bytes through `getBytes`.
**/
class BytesOutput extends Output {
	var b:BytesBuffer;

	public var length(get, never):Int;

	public function new() {
		b = new BytesBuffer();
	}

	inline function get_length():Int {
		return b.length;
	}

	override public function writeByte(value:Int):Void {
		b.addByte(value);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		b.addBytes(bytes, pos, len);
		return len;
	}

	public function getBytes():Bytes {
		return b.getBytes();
	}
}
