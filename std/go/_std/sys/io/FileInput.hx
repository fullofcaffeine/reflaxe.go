package sys.io;

import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Error;
import hxrt.fs.FileInputHandle;
import hxrt.fs.NativeFile;

/**
	What
	- Implements the Haxe 4.3.7 `sys.io.FileInput` stream contract for the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because it is an extern subclass whose file handle, EOF behavior, and seek
	  operations must be supplied by the target. Those public semantics should be
	  reviewable Haxe source instead of File-specific compiler branches.

	How
	- Store an opaque typed `hxrt` handle, perform bounds and EOF translation in
	  Haxe, and delegate only typed byte-value reads, position changes, and close
	  operations.
**/
class FileInput extends haxe.io.Input {
	private var handle:FileInputHandle;

	/**
		What: Wrap one native readable file handle.
		Why: `File.read` and root `Sys.stdin` need the same source-owned stream type.
		How: Retain only the opaque typed handle; `hxrt` owns the OS resource.
	**/
	public function new(handle:FileInputHandle) {
		this.handle = handle;
	}

	override public function readByte():Int {
		var value = NativeFile.inputReadByte(handle);
		if (value < 0)
			throw new Eof();
		return value;
	}

	override public function readBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;

		var values = NativeFile.inputReadValues(handle, length);
		if (values.length == 0)
			throw new Eof();
		for (index in 0...values.length)
			bytes.set(pos + index, values[index]);
		return values.length;
	}

	public function seek(p:Int, pos:FileSeek):Void {
		switch (pos) {
			case SeekBegin:
				NativeFile.inputSeek(handle, p, 0);
			case SeekCur:
				NativeFile.inputSeek(handle, p, 1);
			case SeekEnd:
				NativeFile.inputSeek(handle, p, 2);
		}
	}

	public function tell():Int {
		return NativeFile.inputTell(handle);
	}

	public function eof():Bool {
		return NativeFile.inputEof(handle);
	}

	override public function close():Void {
		NativeFile.inputClose(handle);
	}
}
