package sys.io;

import haxe.io.Bytes;
import hxrt.fs.NativeFile;

/**
	What
	- Owns the complete Haxe 4.3.7 `sys.io.File` static API selected by the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because it is an extern contract whose implementation must be supplied by the
	  target. File algorithms and stream construction belong in typed Haxe source,
	  not in compiler-emitted Go declarations.

	How
	- Delegate OS operations to `hxrt.fs.NativeFile`, convert typed integer byte
	  arrays through the public `Bytes` API, and wrap opaque native handles in the
	  staged `FileInput` and `FileOutput` classes.
**/
class File {
	public static function getContent(path:String):String {
		return NativeFile.readContent(path);
	}

	public static function saveContent(path:String, content:String):Void {
		NativeFile.writeContent(path, content);
	}

	public static function getBytes(path:String):Bytes {
		var values = NativeFile.readByteValues(path);
		var bytes = Bytes.alloc(values.length);
		for (index in 0...values.length)
			bytes.set(index, values[index]);
		return bytes;
	}

	public static function saveBytes(path:String, bytes:Bytes):Void {
		var values = new Array<Int>();
		for (index in 0...bytes.length)
			values.push(bytes.get(index));
		NativeFile.writeByteValues(path, values);
	}

	public static function read(path:String, binary:Bool = true):FileInput {
		return new FileInput(NativeFile.openInput(path));
	}

	public static function write(path:String, binary:Bool = true):FileOutput {
		return new FileOutput(NativeFile.openWrite(path));
	}

	public static function append(path:String, binary:Bool = true):FileOutput {
		return new FileOutput(NativeFile.openAppend(path));
	}

	public static function update(path:String, binary:Bool = true):FileOutput {
		return new FileOutput(NativeFile.openUpdate(path));
	}

	public static function copy(srcPath:String, dstPath:String):Void {
		NativeFile.copyContents(srcPath, dstPath);
	}
}
