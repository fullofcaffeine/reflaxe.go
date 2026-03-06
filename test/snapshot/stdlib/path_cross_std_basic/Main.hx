import haxe.io.Path;

class Main {
	static function main() {
		var unix = new Path("/tmp/demo.txt");
		Sys.println(unix.dir);
		Sys.println(unix.file);
		Sys.println(unix.ext);
		Sys.println(unix.toString());

		var dot = new Path(".");
		Sys.println(dot.dir);
		Sys.println(dot.file);
		Sys.println(dot.ext);
		Sys.println(dot.toString());

		Sys.println(Path.withoutExtension("/tmp/demo.txt"));
		Sys.println(Path.withoutDirectory("/tmp/demo.txt"));
		Sys.println(Path.directory("demo.txt"));
		Sys.println(Path.extension("/tmp/demo.txt"));
		Sys.println(Path.withExtension("/tmp/demo.txt", "log"));
		Sys.println(Path.join(["/tmp", "demo", "..", "out", "file.txt"]));
		Sys.println(Path.normalize("/usr/local/../lib//./a\\b"));
		Sys.println(Path.addTrailingSlash("a\\b"));
		Sys.println(Path.addTrailingSlash("a/b"));
		Sys.println(Path.removeTrailingSlashes("a///"));
		Sys.println(Path.isAbsolute("/tmp/demo.txt"));
		Sys.println(Path.isAbsolute("C:\\tmp\\demo.txt"));
		Sys.println(Path.isAbsolute("\\\\server\\share"));
		Sys.println(Path.isAbsolute("relative/path"));
	}
}
