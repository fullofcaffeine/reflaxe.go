import sys.FileSystem;
import sys.io.File;

class Main {
	static function rmDirRecursive(path:String):Void {
		if (!FileSystem.exists(path))
			return;
		for (entry in FileSystem.readDirectory(path)) {
			var child = path + "/" + entry;
			if (FileSystem.isDirectory(child)) {
				rmDirRecursive(child);
			} else {
				FileSystem.deleteFile(child);
			}
		}
		FileSystem.deleteDirectory(path);
	}

	static function firstEntry(items:Array<String>):String {
		return items.length > 0 ? items[0] : "";
	}

	static function main() {
		var root = "tmp_fs_smoke";
		var fileA = root + "/a.txt";
		var fileB = root + "/b.txt";

		rmDirRecursive(root);
		Sys.println("exists0=" + FileSystem.exists(root));
		var missingThrows = false;
		try {
			FileSystem.isDirectory(root);
		} catch (_:Dynamic) {
			missingThrows = true;
		}
		Sys.println("missing.throws=" + missingThrows);
		FileSystem.createDirectory(root);
		Sys.println("dir1=" + FileSystem.isDirectory(root));
		var absolute = FileSystem.absolutePath(root);
		var canonical = FileSystem.fullPath(root);
		Sys.println("paths=" + (FileSystem.isDirectory(absolute) && FileSystem.isDirectory(canonical)));
		var missingAbsolute = FileSystem.absolutePath(root + "/missing/child.txt");
		Sys.println("absolute.missing=" + !FileSystem.exists(missingAbsolute));
		var directoryOnly = root + "/directory-only";
		FileSystem.createDirectory(directoryOnly);
		var deleteFileDirectoryThrows = false;
		try {
			FileSystem.deleteFile(directoryOnly);
		} catch (_:Dynamic) {
			deleteFileDirectoryThrows = true;
		}
		Sys.println("delete.file.directory.throws=" + deleteFileDirectoryThrows);
		FileSystem.deleteDirectory(directoryOnly);
		File.saveContent(fileA, "hello");
		FileSystem.rename(fileA, fileB);
		var deleteDirectoryFileThrows = false;
		try {
			FileSystem.deleteDirectory(fileB);
		} catch (_:Dynamic) {
			deleteDirectoryFileThrows = true;
		}
		Sys.println("delete.directory.file.throws=" + deleteDirectoryFileThrows);
		var names = FileSystem.readDirectory(root);
		Sys.println("entry=" + firstEntry(names));
		Sys.println("size=" + FileSystem.stat(fileB).size);
		Sys.println("content=" + File.getContent(fileB));
		FileSystem.deleteFile(fileB);
		FileSystem.deleteDirectory(root);
		Sys.println("exists1=" + FileSystem.exists(root));
	}
}
