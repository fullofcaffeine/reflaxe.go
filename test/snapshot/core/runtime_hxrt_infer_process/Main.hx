class Main {
	static function main() {
		var process = new sys.io.Process("echo", ["ok"]);
		process.close();
	}
}
