package sys.net;

import go.NativeSlice;
import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Error;
import hxrt.net.NativeSocket;
import hxrt.net.SocketHandle;
import hxrt.net.SocketIOResult;

/**
	What: Source-owned `haxe.io.Input` wrapper for one typed native socket handle.
	Why: Bounds, EOF, blocked, and Bytes mutation are Haxe stream semantics; the
	mainstream extern Socket cannot provide them and hxrt must not own generated types.
	How: Delegate byte transfer through `NativeSocket`, then translate typed status and
	copy native values into the caller's portable Bytes object.
**/
@:noCompletion
@:dox(hide)
class SocketInput extends haxe.io.Input {
	private var handle:SocketHandle;

	public function new(handle:SocketHandle) {
		this.handle = handle;
	}

	override public function readByte():Int {
		var value = NativeSocket.readByte(handle);
		if (value == NativeSocket.READ_BLOCKED)
			throw Error.Blocked;
		if (value == NativeSocket.READ_EOF)
			throw new Eof();
		return value;
	}

	override public function readBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;
		var result = NativeSocket.readValues(handle, length);
		translateReadStatus(result);
		for (index in 0...result.count)
			bytes.set(pos + index, result.values[index]);
		return result.count;
	}

	override public function close():Void {
		NativeSocket.close(handle);
	}

	/** Translate the typed native status into the public Haxe stream exceptions. **/
	private static function translateReadStatus(result:SocketIOResult):Void {
		if (result.status == NativeSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeSocket.IO_EOF)
			throw new Eof();
	}
}

/**
	What: Source-owned `haxe.io.Output` wrapper for one typed native socket handle.
	Why: Bounds, blocked/closed translation, and Haxe Output conformance are library
	semantics, while the mainstream extern Socket cannot supply a target stream.
	How: Copy Bytes into a typed native slice, delegate transfer, and translate the
	explicit native status without exposing generated objects to `hxrt`.
**/
@:noCompletion
@:dox(hide)
class SocketOutput extends haxe.io.Output {
	private var handle:SocketHandle;

	public function new(handle:SocketHandle) {
		this.handle = handle;
	}

	override public function writeByte(value:Int):Void {
		var result = NativeSocket.writeValues(handle, NativeSlice.fromArray([value]));
		translateWriteStatus(result);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;
		var values = new Array<Int>();
		for (index in 0...length)
			values.push(bytes.get(pos + index));
		var result = NativeSocket.writeValues(handle, NativeSlice.fromArray(values));
		translateWriteStatus(result);
		return result.count;
	}

	override public function flush():Void {
		NativeSocket.flush(handle);
	}

	override public function close():Void {
		NativeSocket.close(handle);
	}

	/** Translate the typed native status into the public Haxe stream exceptions. **/
	private static function translateWriteStatus(result:SocketIOResult):Void {
		if (result.status == NativeSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeSocket.IO_EOF)
			throw new Eof();
	}
}
