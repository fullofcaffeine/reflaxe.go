import go.Chan;

class Main {
	static function main() {
		var s:String = "héllo";
		var out = s.charAt(1) + Std.string(s.charCodeAt(1)) + s.substring(0, 3) + s.substr(-2);
		Sys.println(s.length + out.length);

		var ch = new Chan<Int>();
		ch.send(7);
		Sys.println(ch.recv());
	}
}
