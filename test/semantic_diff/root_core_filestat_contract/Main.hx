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

	static function main() {
		var rootDir = "semantic_root_core_contract";
		var filePath = rootDir + "/note.txt";
		rmDirRecursive(rootDir);
		FileSystem.createDirectory(rootDir);
		File.saveContent(filePath, "hello");
		var stat:sys.FileStat = FileSystem.stat(filePath);
		Sys.println("stat.size=" + stat.size);
		Sys.println("stat.links=" + (stat.nlink >= 1));
		Sys.println("stat.mode=" + (stat.mode > 0));
		Sys.println("stat.mtime=" + (stat.mtime != null));
		rmDirRecursive(rootDir);
		Sys.println("stat.cleanup=" + !FileSystem.exists(rootDir));
	}
}
