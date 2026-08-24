@:go.import("image")
@:go.name("Point")
@:go.struct
extern class GoPoint {
	public function new();

	@:go.name("X")
	public var x:Int;

	@:go.name("Y")
	public var y:Int;
}

@:go.import("fmt")
extern class GoFmt {
	@:go.name("Println")
	public static function println(value:Int):Void;
}

@:go.import("net/url")
@:go.name("URL")
@:go.struct
extern class GoURL {
	public function new();

	@:go.name("Scheme")
	public var scheme:String;

	@:go.name("Path")
	public var path:String;
}

class Main {
	static function main() {
		var point = new GoPoint();
		point.x = 20;
		point.y = 22;
		var url = new GoURL();
		var assignedScheme = url.scheme = "http";
		url.path = "";
		url.scheme += "s";
		var appendedPath = url.path += "/beads";
		if (assignedScheme == "http" && url.scheme == "https" && appendedPath == "/beads" && url.path == "/beads") {
			GoFmt.println(point.x + point.y);
		} else {
			GoFmt.println(-1);
		}
	}
}
