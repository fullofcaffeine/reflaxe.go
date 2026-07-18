package haxe.io;

/**
	What: Adds a refillable byte buffer in front of another Input.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	while its base Input carrier is target-defined. This override restores the
	target-neutral buffer algorithm over the new staged hierarchy.

	How: Reuse staged `Bytes.blit` and virtual `Input.readBytes`, preserving the
	upstream available/position state machine.
**/
class BufferInput extends Input {
	public var i:Input;
	public var buf:Bytes;
	public var available:Int;
	public var pos:Int;

	public function new(input:Input, buffer:Bytes, ?pos:Int = 0, ?available:Int = 0) {
		this.i = input;
		this.buf = buffer;
		this.pos = pos;
		this.available = available;
	}

	public function refill():Void {
		if (pos > 0) {
			buf.blit(0, buf, pos, available);
			pos = 0;
		}
		available += i.readBytes(buf, available, buf.length - available);
	}

	override public function readByte():Int {
		if (available == 0)
			refill();
		var value = buf.get(pos);
		pos++;
		available--;
		return value;
	}

	override public function readBytes(bytes:Bytes, targetPos:Int, len:Int):Int {
		if (available == 0)
			refill();
		var size = len > available ? available : len;
		bytes.blit(targetPos, buf, pos, size);
		pos += size;
		available -= size;
		return size;
	}
}
