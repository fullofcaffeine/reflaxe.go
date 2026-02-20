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
		var root = "semantic_fs_contract";
		var fileA = root + "/a.txt";
		var fileB = root + "/b.txt";

		rmDirRecursive(root);
		Sys.println("exists0=" + FileSystem.exists(root));
		FileSystem.createDirectory(root);
		Sys.println("dir1=" + FileSystem.isDirectory(root));

		File.saveContent(fileA, "alpha");
		FileSystem.rename(fileA, fileB);
		var names = FileSystem.readDirectory(root);
		Sys.println("entry=" + firstEntry(names));
		Sys.println("content=" + File.getContent(fileB));

		FileSystem.deleteFile(fileB);
		FileSystem.deleteDirectory(root);
		Sys.println("exists1=" + FileSystem.exists(root));
	}
}
