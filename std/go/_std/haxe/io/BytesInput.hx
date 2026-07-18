package haxe.io;

/**
	What: Reads a bounded window from one Bytes value.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	without a target Bytes carrier. Position clamping, EOF, partial reads, and
	aliasing nevertheless remain portable stream semantics.

	How: Retain the aliased `BytesData`, track the current window, and override only
	the primitive read methods inherited from staged `Input`.
**/
class BytesInput extends Input {
	var b:BytesData;
	var pos:Int;
	var len:Int;
	var totlen:Int;

	public var position(get, set):Int;
	public var length(get, never):Int;

	public function new(bytes:Bytes, ?pos:Int, ?len:Int) {
		if (pos == null)
			pos = 0;
		if (len == null)
			len = bytes.length - pos;
		if (pos < 0 || len < 0 || pos + len > bytes.length)
			throw Error.OutsideBounds;
		this.b = bytes.getData();
		this.pos = pos;
		this.len = len;
		this.totlen = len;
	}

	inline function get_position():Int {
		return pos;
	}

	inline function get_length():Int {
		return totlen;
	}

	function set_position(value:Int):Int {
		if (value < 0)
			value = 0;
		else if (value > length)
			value = length;
		len = totlen - value;
		return pos = value;
	}

	override public function readByte():Int {
		if (len == 0)
			throw new Eof();
		len--;
		return b[pos++];
	}

	override public function readBytes(bytes:Bytes, targetPos:Int, requested:Int):Int {
		if (targetPos < 0 || requested < 0 || targetPos + requested > bytes.length)
			throw Error.OutsideBounds;
		if (len == 0 && requested > 0)
			throw new Eof();
		if (requested > len)
			requested = len;
		for (index in 0...requested)
			bytes.set(targetPos + index, b[pos + index]);
		pos += requested;
		len -= requested;
		return requested;
	}
}
