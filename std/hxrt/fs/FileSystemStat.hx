package hxrt.fs;

/**
	What
	- Typed runtime carrier for filesystem metadata used by staged
	  `sys.FileSystem` on `haxe.go`.

	Why
	- Go's `os.FileInfo` and native time values cannot implement the upstream
	  Haxe `sys.FileStat` anonymous record directly. This target support belongs
	  under `std/hxrt`, not in the compiler or the override-only `_std` tree.

	How
	- Metadata maps fields directly to `hxrt.FileSystemStat`; staged Haxe converts
	  the millisecond timestamps to `Date` and constructs the public record.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("FileSystemStat")
extern class FileSystemStat {
	@:go.name("Gid")
	public var gid:Int;

	@:go.name("Uid")
	public var uid:Int;

	@:go.name("AtimeMs")
	public var atimeMs:Float;

	@:go.name("MtimeMs")
	public var mtimeMs:Float;

	@:go.name("CtimeMs")
	public var ctimeMs:Float;

	@:go.name("Size")
	public var size:Int;

	@:go.name("Dev")
	public var dev:Int;

	@:go.name("Ino")
	public var ino:Int;

	@:go.name("Nlink")
	public var nlink:Int;

	@:go.name("Rdev")
	public var rdev:Int;

	@:go.name("Mode")
	public var mode:Int;
}
