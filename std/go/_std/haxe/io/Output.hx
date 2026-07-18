package haxe.io;

/**
	What: Defines the portable base contract and algorithms for byte outputs.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	without target byte and base carriers. Numeric encoding, blocked-write loops,
	and string conversion are Haxe stdlib semantics rather than compiler wrappers.

	How: Express the upstream behavior as staged Haxe and rely on normal virtual
	dispatch for target-specific `writeByte` and `writeBytes` overrides.
**/
class Output {
	public var bigEndian(default, set):Bool;

	public function writeByte(value:Int):Void {
		throw new haxe.exceptions.NotImplementedException();
	}

	public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		if (pos < 0 || len < 0 || pos + len > bytes.length)
			throw Error.OutsideBounds;
		var total = len;
		while (len > 0) {
			writeByte(bytes.get(pos));
			pos++;
			len--;
		}
		return total;
	}

	public function flush():Void {}

	public function close():Void {}

	function set_bigEndian(value:Bool):Bool {
		bigEndian = value;
		return value;
	}

	public function write(bytes:Bytes):Void {
		writeFullBytes(bytes, 0, bytes.length);
	}

	public function writeFullBytes(bytes:Bytes, pos:Int, len:Int):Void {
		while (len > 0) {
			var count = writeBytes(bytes, pos, len);
			if (count == 0)
				throw Error.Blocked;
			pos += count;
			len -= count;
		}
	}

	public function writeFloat(value:Float):Void {
		writeInt32(FPHelper.floatToI32(value));
	}

	public function writeDouble(value:Float):Void {
		var bits = FPHelper.doubleToI64(value);
		if (bigEndian) {
			writeInt32(bits.high);
			writeInt32(bits.low);
		} else {
			writeInt32(bits.low);
			writeInt32(bits.high);
		}
	}

	public function writeInt8(value:Int):Void {
		if (value < -0x80 || value >= 0x80)
			throw Error.Overflow;
		writeByte(value & 0xFF);
	}

	public function writeInt16(value:Int):Void {
		if (value < -0x8000 || value >= 0x8000)
			throw Error.Overflow;
		writeUInt16(value & 0xFFFF);
	}

	public function writeUInt16(value:Int):Void {
		if (value < 0 || value >= 0x10000)
			throw Error.Overflow;
		if (bigEndian) {
			writeByte(value >> 8);
			writeByte(value & 0xFF);
		} else {
			writeByte(value & 0xFF);
			writeByte(value >> 8);
		}
	}

	public function writeInt24(value:Int):Void {
		if (value < -0x800000 || value >= 0x800000)
			throw Error.Overflow;
		writeUInt24(value & 0xFFFFFF);
	}

	public function writeUInt24(value:Int):Void {
		if (value < 0 || value >= 0x1000000)
			throw Error.Overflow;
		if (bigEndian) {
			writeByte(value >> 16);
			writeByte((value >> 8) & 0xFF);
			writeByte(value & 0xFF);
		} else {
			writeByte(value & 0xFF);
			writeByte((value >> 8) & 0xFF);
			writeByte(value >> 16);
		}
	}

	public function writeInt32(value:Int):Void {
		if (bigEndian) {
			writeByte(value >>> 24);
			writeByte((value >> 16) & 0xFF);
			writeByte((value >> 8) & 0xFF);
			writeByte(value & 0xFF);
		} else {
			writeByte(value & 0xFF);
			writeByte((value >> 8) & 0xFF);
			writeByte((value >> 16) & 0xFF);
			writeByte(value >>> 24);
		}
	}

	public function prepare(nbytes:Int):Void {}

	public function writeInput(input:Input, ?bufsize:Int):Void {
		if (bufsize == null)
			bufsize = 4096;
		var buffer = Bytes.alloc(bufsize);
		try {
			while (true) {
				var count = input.readBytes(buffer, 0, bufsize);
				if (count == 0)
					throw Error.Blocked;
				writeFullBytes(buffer, 0, count);
			}
		} catch (_:Eof) {}
	}

	public function writeString(value:String, ?encoding:Encoding):Void {
		var bytes = Bytes.ofString(value, encoding);
		writeFullBytes(bytes, 0, bytes.length);
	}
}
