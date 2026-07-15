package sys;

import hxrt.fs.NativeFileSystem;

/**
	What
	- Owns the complete Haxe 4.3.7 `sys.FileSystem` API selected by the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern contract whose behavior must be supplied by
	  the target. The public API and `FileStat` construction still belong in Haxe
	  source; only native filesystem capabilities belong in `hxrt`.

	How
	- Delegate native operations through the typed `hxrt.fs.NativeFileSystem`
	  binding. Convert its typed metadata carrier into the unchanged upstream
	  `sys.FileStat` record and construct Haxe `Date` values in source.
**/
class FileSystem {
	public static function exists(path:String):Bool {
		return NativeFileSystem.exists(path);
	}

	public static function rename(path:String, newPath:String):Void {
		NativeFileSystem.rename(path, newPath);
	}

	public static function stat(path:String):FileStat {
		var value = NativeFileSystem.stat(path);
		return {
			gid: value.gid,
			uid: value.uid,
			atime: Date.fromTime(value.atimeMs),
			mtime: Date.fromTime(value.mtimeMs),
			ctime: Date.fromTime(value.ctimeMs),
			size: value.size,
			dev: value.dev,
			ino: value.ino,
			nlink: value.nlink,
			rdev: value.rdev,
			mode: value.mode
		};
	}

	public static function fullPath(relPath:String):String {
		return NativeFileSystem.fullPath(relPath);
	}

	public static function absolutePath(relPath:String):String {
		return NativeFileSystem.absolutePath(relPath);
	}

	public static function isDirectory(path:String):Bool {
		return NativeFileSystem.isDirectory(path);
	}

	public static function createDirectory(path:String):Void {
		NativeFileSystem.createDirectory(path);
	}

	public static function deleteFile(path:String):Void {
		NativeFileSystem.deleteFile(path);
	}

	public static function deleteDirectory(path:String):Void {
		NativeFileSystem.deleteDirectory(path);
	}

	public static function readDirectory(path:String):Array<String> {
		return NativeFileSystem.readDirectory(path);
	}
}
