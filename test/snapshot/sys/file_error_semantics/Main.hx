import sys.FileSystem;
import sys.io.File;

class Main {
	static function throws(action:() -> Void):Bool {
		try {
			action();
			return false;
		} catch (_:Dynamic) {
			return true;
		}
		return false;
	}

	static function main() {
		var root = "tmp_file_error_semantics";
		if (FileSystem.exists(root)) {
			FileSystem.deleteDirectory(root);
		}
		FileSystem.createDirectory(root);

		Sys.println("missing.read.throws=" + throws(() -> {
			File.getContent(root + "/missing.txt");
		}));
		Sys.println("directory.read.throws=" + throws(() -> {
			File.getContent(root);
		}));
		Sys.println("directory.write.throws=" + throws(() -> {
			File.saveContent(root, "not-a-file");
		}));
		Sys.println("environment.invalid.throws=" + throws(() -> {
			Sys.putEnv("HAXE_GO=INVALID", "value");
		}));

		var locked = root + "/locked.txt";
		File.saveContent(locked, "secret");
		Sys.command("chmod", ["000", locked]);
		Sys.println("permission.read.throws=" + throws(() -> {
			File.getContent(locked);
		}));
		Sys.println("permission.write.throws=" + throws(() -> {
			File.saveContent(locked, "replacement");
		}));
		Sys.command("chmod", ["600", locked]);
		FileSystem.deleteFile(locked);

		FileSystem.deleteDirectory(root);
	}
}
