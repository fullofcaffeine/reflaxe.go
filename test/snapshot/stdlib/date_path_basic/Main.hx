import haxe.io.Path;

class Main {
	static function main() {
		var d = Date.fromString("2024-02-03 04:05:06");
		Sys.println(d.getFullYear());
		Sys.println(d.getMonth());
		Sys.println(d.getDate());
		Sys.println(d.getHours());

		Sys.println(Path.join(["a", "b", "c.txt"]));

		var p = new Path("/tmp/demo.txt");
		Sys.println(p.dir);
		Sys.println(p.file);
		Sys.println(p.ext);
	}
}
