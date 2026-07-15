package hxrt.process;

/**
	What
	- Typed bridge to native child-process and pipe capabilities used by staged
	  `sys.io.Process`.

	Why
	- Process handles, OS pipes, waits, signals, and startup failures require Go
	  runtime support. Public streams, bounds/EOF translation, nullable status, and
	  lifecycle policy belong in Haxe source instead of compiler shims.

	How
	- Map one-for-one to exported `runtime/hxrt/process.go` functions. Bytes cross as
	  `Array<Int>` / `[]int`, and nonblocking status crosses through the typed
	  `ProcessExitStatus` carrier rather than `Dynamic` or `Any`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeProcess {
	@:go.name("ProcessCreate")
	public static function create(command:String, args:Array<String>):ProcessHandle;

	@:go.name("ProcessStdout")
	public static function stdout(process:ProcessHandle):ProcessOutputHandle;

	@:go.name("ProcessStderr")
	public static function stderr(process:ProcessHandle):ProcessOutputHandle;

	@:go.name("ProcessStdin")
	public static function stdin(process:ProcessHandle):ProcessInputHandle;

	@:go.name("ProcessOutputReadByteValue")
	public static function outputReadByte(handle:ProcessOutputHandle):Int;

	@:go.name("ProcessOutputReadValues")
	public static function outputReadValues(handle:ProcessOutputHandle, length:Int):Array<Int>;

	@:go.name("ProcessOutputClose")
	public static function outputClose(handle:ProcessOutputHandle):Void;

	@:go.name("ProcessInputWriteByteValue")
	public static function inputWriteByte(handle:ProcessInputHandle, value:Int):Bool;

	@:go.name("ProcessInputWriteValues")
	public static function inputWriteValues(handle:ProcessInputHandle, values:Array<Int>):Bool;

	@:go.name("ProcessInputFlush")
	public static function inputFlush(handle:ProcessInputHandle):Void;

	@:go.name("ProcessInputClose")
	public static function inputClose(handle:ProcessInputHandle):Void;

	@:go.name("ProcessPid")
	public static function pid(process:ProcessHandle):Int;

	@:go.name("ProcessExitStatusValue")
	public static function exitStatus(process:ProcessHandle, block:Bool):ProcessExitStatus;

	@:go.name("ProcessKill")
	public static function kill(process:ProcessHandle):Void;

	@:go.name("ProcessClose")
	public static function close(process:ProcessHandle):Void;
}
