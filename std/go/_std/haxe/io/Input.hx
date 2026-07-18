package haxe.io;

/**
	What: Defines the portable base contract and algorithms for byte inputs.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	without a target base carrier and byte representation. The algorithms are
	ordinary Haxe behavior, not compiler interface wrappers or raw helpers.

	How: Lower the stream loops and numeric decoding as a normal staged class. The
	compiler's standard `__hx_this` dispatch makes calls from these base methods
	reach `BytesInput`, file, process, socket, and user overrides.
**/
class Input {
	public var bigEndian(default, set):Bool;

	public function readByte():Int {
		return throw new haxe.exceptions.NotImplementedException();
	}

	public function readBytes(bytes:Bytes, pos:Int, len:Int):Int {
		if (pos < 0 || len < 0 || pos + len > bytes.length)
			throw Error.OutsideBounds;
		var remaining = len;
		try {
			while (remaining > 0) {
				bytes.set(pos, readByte());
				pos++;
				remaining--;
			}
		} catch (_:Eof) {}
		return len - remaining;
	}

	public function close():Void {}

	function set_bigEndian(value:Bool):Bool {
		bigEndian = value;
		return value;
	}

	public function readAll(?bufsize:Int):Bytes {
		if (bufsize == null)
			bufsize = 1 << 14;
		var buffer = Bytes.alloc(bufsize);
		var total = new BytesBuffer();
		try {
			while (true) {
				var count = readBytes(buffer, 0, bufsize);
				if (count == 0)
					throw Error.Blocked;
				total.addBytes(buffer, 0, count);
			}
		} catch (_:Eof) {}
		return total.getBytes();
	}

	public function readFullBytes(bytes:Bytes, pos:Int, len:Int):Void {
		while (len > 0) {
			var count = readBytes(bytes, pos, len);
			if (count == 0)
				throw Error.Blocked;
			pos += count;
			len -= count;
		}
	}

	public function read(nbytes:Int):Bytes {
		var bytes = Bytes.alloc(nbytes);
		readFullBytes(bytes, 0, nbytes);
		return bytes;
	}

	public function readUntil(end:Int):String {
		var buffer = new BytesBuffer();
		var value:Int;
		while ((value = readByte()) != end)
			buffer.addByte(value);
		return buffer.getBytes().toString();
	}

	public function readLine():String {
		var buffer = new BytesBuffer();
		var value:Int;
		var result:String;
		try {
			while ((value = readByte()) != 10)
				buffer.addByte(value);
			result = buffer.getBytes().toString();
			if (result.length > 0 && result.charCodeAt(result.length - 1) == 13)
				result = result.substr(0, -1);
		} catch (error:Eof) {
			result = buffer.getBytes().toString();
			if (result.length == 0)
				throw error;
		}
		return result;
	}

	public function readFloat():Float {
		return FPHelper.i32ToFloat(readInt32());
	}

	public function readDouble():Float {
		var first = readInt32();
		var second = readInt32();
		return bigEndian ? FPHelper.i64ToDouble(second, first) : FPHelper.i64ToDouble(first, second);
	}

	public function readInt8():Int {
		var value = readByte();
		return value >= 128 ? value - 256 : value;
	}

	public function readInt16():Int {
		var first = readByte();
		var second = readByte();
		var value = bigEndian ? second | (first << 8) : first | (second << 8);
		return (value & 0x8000) != 0 ? value - 0x10000 : value;
	}

	public function readUInt16():Int {
		var first = readByte();
		var second = readByte();
		return bigEndian ? second | (first << 8) : first | (second << 8);
	}

	public function readInt24():Int {
		var first = readByte();
		var second = readByte();
		var third = readByte();
		var value = bigEndian ? third | (second << 8) | (first << 16) : first | (second << 8) | (third << 16);
		return (value & 0x800000) != 0 ? value - 0x1000000 : value;
	}

	public function readUInt24():Int {
		var first = readByte();
		var second = readByte();
		var third = readByte();
		return bigEndian ? third | (second << 8) | (first << 16) : first | (second << 8) | (third << 16);
	}

	public function readInt32():Int {
		var first = readByte();
		var second = readByte();
		var third = readByte();
		var fourth = readByte();
		return bigEndian ? fourth | (third << 8) | (second << 16) | (first << 24) : first | (second << 8) | (third << 16) | (fourth << 24);
	}

	public function readString(len:Int, ?encoding:Encoding):String {
		var bytes = Bytes.alloc(len);
		readFullBytes(bytes, 0, len);
		return bytes.getString(0, len, encoding);
	}
}
