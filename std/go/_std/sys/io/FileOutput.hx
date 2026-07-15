package sys.io;

import haxe.io.Bytes;
import haxe.io.Error;
import hxrt.fs.FileOutputHandle;
import hxrt.fs.NativeFile;

/**
	What
	- Implements the Haxe 4.3.7 `sys.io.FileOutput` stream contract for the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because it is an extern subclass whose writable handle and seek operations must
	  be implemented by the target. Bounds checks and public stream behavior belong
	  in staged Haxe instead of compiler-emitted method bodies.

	How
	- Store an opaque typed `hxrt` handle, validate and convert byte ranges in Haxe,
	  and delegate only typed byte-value writes, positioning, flush, and close
	  operations.
**/
class FileOutput extends haxe.io.Output {
	private var handle:FileOutputHandle;

	/**
		What: Wrap one native writable file handle.
		Why: File outputs and root standard outputs share one source-owned stream type.
		How: Retain only the opaque typed handle; `hxrt` owns the OS resource policy.
	**/
	public function new(handle:FileOutputHandle) {
		this.handle = handle;
	}

	override public function writeByte(value:Int):Void {
		NativeFile.outputWriteByte(handle, value);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;
		var values = new Array<Int>();
		for (index in 0...length)
			values.push(bytes.get(pos + index));
		return NativeFile.outputWriteValues(handle, values, 0, length);
	}

	public function seek(p:Int, pos:FileSeek):Void {
		switch (pos) {
			case SeekBegin:
				NativeFile.outputSeek(handle, p, 0);
			case SeekCur:
				NativeFile.outputSeek(handle, p, 1);
			case SeekEnd:
				NativeFile.outputSeek(handle, p, 2);
		}
	}

	public function tell():Int {
		return NativeFile.outputTell(handle);
	}

	override public function flush():Void {
		NativeFile.outputFlush(handle);
	}

	override public function close():Void {
		NativeFile.outputClose(handle);
	}
}
