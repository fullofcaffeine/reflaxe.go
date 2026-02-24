class Main {
	static function main() {
		var cwd = Sys.getCwd();
		var process = new sys.io.Process("echo", ["ok"]);
		process.close();
		Sys.println(cwd);
	}
}
