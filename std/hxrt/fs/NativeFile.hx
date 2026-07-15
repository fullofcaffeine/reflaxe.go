package hxrt.fs;

/**
	What
	- Typed bridge to native file and file-stream capabilities used by staged `sys.io`.

	Why
	- OS handles and byte transfer require Go runtime support, but the public Haxe API,
	  bounds checks, EOF construction, and seek-origin selection do not belong in the
	  compiler or in raw injection.

	How
	- Map each operation one-for-one to `runtime/hxrt/file.go`. Arbitrary bytes cross
	  as `Array<Int>` / `[]int`, keeping generated `haxe.io.Bytes` internals out of hxrt.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeFile {
	@:go.name("SysStdin")
	public static function stdin():FileInputHandle;

	@:go.name("SysStdout")
	public static function stdout():FileOutputHandle;

	@:go.name("SysStderr")
	public static function stderr():FileOutputHandle;

	@:go.name("FileReadContent")
	public static function readContent(path:String):String;

	@:go.name("FileWriteContent")
	public static function writeContent(path:String, content:String):Void;

	@:go.name("FileReadByteValues")
	public static function readByteValues(path:String):Array<Int>;

	@:go.name("FileWriteByteValues")
	public static function writeByteValues(path:String, values:Array<Int>):Void;

	@:go.name("FileCopyContents")
	public static function copyContents(srcPath:String, dstPath:String):Void;

	@:go.name("FileOpenInput")
	public static function openInput(path:String):FileInputHandle;

	@:go.name("FileOpenWrite")
	public static function openWrite(path:String):FileOutputHandle;

	@:go.name("FileOpenAppend")
	public static function openAppend(path:String):FileOutputHandle;

	@:go.name("FileOpenUpdate")
	public static function openUpdate(path:String):FileOutputHandle;

	@:go.name("FileInputReadByteValue")
	public static function inputReadByte(handle:FileInputHandle):Int;

	@:go.name("FileInputReadValues")
	public static function inputReadValues(handle:FileInputHandle, length:Int):Array<Int>;

	@:go.name("FileInputTell")
	public static function inputTell(handle:FileInputHandle):Int;

	@:go.name("FileInputSeek")
	public static function inputSeek(handle:FileInputHandle, offset:Int, whence:Int):Void;

	@:go.name("FileInputEof")
	public static function inputEof(handle:FileInputHandle):Bool;

	@:go.name("FileInputClose")
	public static function inputClose(handle:FileInputHandle):Void;

	@:go.name("FileOutputWriteByteValue")
	public static function outputWriteByte(handle:FileOutputHandle, value:Int):Void;

	@:go.name("FileOutputWriteValues")
	public static function outputWriteValues(handle:FileOutputHandle, values:Array<Int>, pos:Int, len:Int):Int;

	@:go.name("FileOutputTell")
	public static function outputTell(handle:FileOutputHandle):Int;

	@:go.name("FileOutputSeek")
	public static function outputSeek(handle:FileOutputHandle, offset:Int, whence:Int):Void;

	@:go.name("FileOutputFlush")
	public static function outputFlush(handle:FileOutputHandle):Void;

	@:go.name("FileOutputClose")
	public static function outputClose(handle:FileOutputHandle):Void;
}
