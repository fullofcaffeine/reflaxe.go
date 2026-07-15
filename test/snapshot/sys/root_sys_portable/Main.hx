class Main {
	static function main() {
		var printFn:Dynamic->Void = Sys.print;
		var printlnFn:Dynamic->Void = Sys.println;
		var localeFn:String->Bool = Sys.setTimeLocale;
		var setCwdFn:String->Void = Sys.setCwd;
		var timeFn:Void->Float = Sys.time;
		var executablePathFn:Void->String = Sys.executablePath;
		var programPathFn:Void->String = Sys.programPath;
		var getCharFn:Bool->Int = Sys.getChar;
		var stdinFn:Void->haxe.io.Input = Sys.stdin;
		var stdoutFn:Void->haxe.io.Output = Sys.stdout;
		var stderrFn:Void->haxe.io.Output = Sys.stderr;

		printFn("print=");
		printlnFn("ok");
		printlnFn("locale=" + !localeFn("__haxe_go_missing_locale__"));

		var cwd = Sys.getCwd();
		setCwdFn(cwd);
		printlnFn("cwd=" + (Sys.getCwd() == cwd));

		var started = timeFn();
		Sys.sleep(0.01);
		var finished = timeFn();
		printlnFn("time=" + (started > 0 && finished >= started));

		var programPath = programPathFn();
		printlnFn("programPath=" + (programPath.length > 0));
		printlnFn("executableAlias=" + (executablePathFn() == programPath));

		var stdin = stdinFn();
		var stderr = stderrFn();
		printlnFn("inputFunctions=" + (stdin != null && stderr != null && getCharFn != null));
		var stdout = stdoutFn();
		stdout.writeString("stdout=ok\n");
		stdout.flush();
		stdout.close();
		stdoutFn().writeString("stdoutAfterClose=ok\n");
	}
}
