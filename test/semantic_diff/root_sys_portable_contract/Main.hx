class Main {
	static function main() {
		Sys.print("print=");
		Sys.println("ok");
		Sys.println("locale=" + !Sys.setTimeLocale("__haxe_go_missing_locale__"));

		var cwd = Sys.getCwd();
		Sys.setCwd(cwd);
		Sys.println("cwd=" + (Sys.getCwd() == cwd));

		var started = Sys.time();
		Sys.sleep(0.01);
		var finished = Sys.time();
		Sys.println("time=" + (started > 0 && finished >= started));

		var programPath = Sys.programPath();
		Sys.println("programPath=" + (programPath.length > 0));
		Sys.println("executableAlias=" + (Sys.executablePath() == programPath));
		Sys.println("streams=" + (Sys.stdin() != null && Sys.stdout() != null && Sys.stderr() != null));

		Sys.stdout().writeString("stdout=ok\n");
		Sys.stdout().flush();
	}
}
