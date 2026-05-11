class Main {
	static function main() {
		var code = Sys.command("sh", ["-c", "printf child-out; exit 7"]);
		Sys.println("");
		Sys.println("code=" + code);
	}
}
