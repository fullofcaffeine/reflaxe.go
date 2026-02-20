class Main {
	static function readVersion(args:Array<String>):String {
		var process = new sys.io.Process("haxe", args);
		var line = StringTools.trim(process.stdout.readLine());
		process.close();
		return line;
	}

	static function main() {
		var versionA = readVersion(["--version"]);
		var versionB = readVersion(["-D", "reflaxe_go_process_contract_probe=1", "--version"]);

		Sys.println("versionA.non_empty=" + (versionA.length > 0));
		Sys.println("versionB.non_empty=" + (versionB.length > 0));
		Sys.println("version.eq=" + (versionA == versionB));
	}
}
