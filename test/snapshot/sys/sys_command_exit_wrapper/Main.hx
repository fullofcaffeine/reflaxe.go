class Main {
	static function main() {
		var code = Sys.command("sh", ["-c", "printf 'wrapper-out\\n'; exit 7"]);
		Sys.exit(code);
	}
}
