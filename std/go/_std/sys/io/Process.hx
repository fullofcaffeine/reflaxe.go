package sys.io;

import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Error;
import go.NativeSlice;
import hxrt.process.NativeProcess;
import hxrt.process.ProcessHandle;
import hxrt.process.ProcessInputHandle;
import hxrt.process.ProcessOutputHandle;

/**
	What
	- Implements the complete Haxe 4.3.7 `sys.io.Process` API for the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern contract. Process streams, nullable exit
	  status, detached rejection, and closed-state behavior must be reviewable Haxe
	  source instead of compiler-emitted structs and raw method bodies.

	How
	- Store one opaque typed process handle, construct ordinary Haxe `Input` and
	  `Output` subclasses around typed pipe handles, and delegate only native spawn,
	  byte transfer, wait, signal, and close operations to `hxrt`.
**/
class Process {
	public var stdout(default, null):haxe.io.Input;
	public var stderr(default, null):haxe.io.Input;
	public var stdin(default, null):haxe.io.Output;

	private var handle:ProcessHandle;

	/**
		What: Start the requested command immediately and expose its three pipes.
		Why: A null argument array selects shell parsing, while an explicit array must bypass the shell; detached pipes are not supported by this target.
		How: Reject detached mode in Haxe, preserve argument nullability, create one native handle, and wrap each typed pipe.
	**/
	public function new(cmd:String, ?args:Array<String>, detached:Bool = false) {
		if (detached)
			throw "Detached process is not supported on this platform";
		handle = NativeProcess.create(cmd, args == null ? null : NativeSlice.fromArray(args));
		stdout = new ProcessOutput(NativeProcess.stdout(handle));
		stderr = new ProcessOutput(NativeProcess.stderr(handle));
		stdin = new ProcessInput(NativeProcess.stdin(handle));
	}

	public function getPid():Int {
		return NativeProcess.pid(requireHandle());
	}

	/** Convert the typed native availability carrier to the public `Null<Int>` contract. **/
	public function exitCode(block:Bool = true):Null<Int> {
		var status = NativeProcess.exitStatus(requireHandle(), block);
		return status.available ? status.code : null;
	}

	/** Release pipes and wait resources without killing a still-running child. **/
	public function close():Void {
		if (handle == null)
			return;
		NativeProcess.close(handle);
		handle = null;
	}

	public function kill():Void {
		NativeProcess.kill(requireHandle());
	}

	private function requireHandle():ProcessHandle {
		if (handle == null)
			throw "Process is closed";
		return handle;
	}
}

/** Source-owned `haxe.io.Input` wrapper for a child stdout or stderr pipe. **/
private class ProcessOutput extends haxe.io.Input {
	private var handle:ProcessOutputHandle;

	public function new(handle:ProcessOutputHandle) {
		this.handle = handle;
	}

	override public function readByte():Int {
		if (handle == null)
			throw "Process output is closed";
		var value = NativeProcess.outputReadByte(handle);
		if (value < 0)
			throw new Eof();
		return value;
	}

	override public function readBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;
		if (handle == null)
			throw "Process output is closed";

		var values = NativeProcess.outputReadValues(handle, length);
		if (values.length == 0)
			throw new Eof();
		for (index in 0...values.length)
			bytes.set(pos + index, values[index]);
		return values.length;
	}

	override public function close():Void {
		if (handle == null)
			return;
		NativeProcess.outputClose(handle);
		handle = null;
	}
}

/** Source-owned `haxe.io.Output` wrapper for a child stdin pipe. **/
private class ProcessInput extends haxe.io.Output {
	private var handle:ProcessInputHandle;

	public function new(handle:ProcessInputHandle) {
		this.handle = handle;
	}

	override public function writeByte(value:Int):Void {
		if (handle == null || !NativeProcess.inputWriteByte(handle, value))
			throw new Eof();
	}

	override public function writeBytes(bytes:Bytes, pos:Int, length:Int):Int {
		if (pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;

		var values = new Array<Int>();
		for (index in 0...length)
			values.push(bytes.get(pos + index));
		if (handle == null || !NativeProcess.inputWriteValues(handle, NativeSlice.fromArray(values)))
			throw new Eof();
		return length;
	}

	override public function flush():Void {
		if (handle != null)
			NativeProcess.inputFlush(handle);
	}

	override public function close():Void {
		if (handle == null)
			return;
		NativeProcess.inputClose(handle);
		handle = null;
	}
}
