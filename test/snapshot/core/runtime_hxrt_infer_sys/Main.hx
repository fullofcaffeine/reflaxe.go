class Main {
	static function main() {
		var cwd = Sys.getCwd();
		if (cwd.length == 0)
			throw "expected a current working directory";
	}
}
