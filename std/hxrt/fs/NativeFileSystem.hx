package hxrt.fs;

/**
	What
	- Typed bridge to the narrow `hxrt` filesystem capability surface used by
	  staged `sys.FileSystem` on `haxe.go`.

	Why
	- Filesystem operations need Go's native `os` and `path/filepath` packages,
	  but the Haxe stdlib API and record construction should remain ordinary
	  source. This binding avoids compiler shims and raw injection.

	How
	- Each static extern maps one-for-one to an exported `hxrt` function. Native
	  failures cross the existing Haxe exception boundary inside the runtime.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeFileSystem {
	@:go.name("FileSystemExists")
	public static function exists(path:String):Bool;

	@:go.name("FileSystemRename")
	public static function rename(path:String, newPath:String):Void;

	@:go.name("FileSystemStatPath")
	public static function stat(path:String):FileSystemStat;

	@:go.name("FileSystemFullPath")
	public static function fullPath(path:String):String;

	@:go.name("FileSystemAbsolutePath")
	public static function absolutePath(path:String):String;

	@:go.name("FileSystemIsDirectory")
	public static function isDirectory(path:String):Bool;

	@:go.name("FileSystemCreateDirectory")
	public static function createDirectory(path:String):Void;

	@:go.name("FileSystemDeleteFile")
	public static function deleteFile(path:String):Void;

	@:go.name("FileSystemDeleteDirectory")
	public static function deleteDirectory(path:String):Void;

	@:go.name("FileSystemReadDirectory")
	public static function readDirectory(path:String):Array<String>;
}
