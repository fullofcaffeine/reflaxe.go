class Main {
	static function main() {
		var envName = "HAXE_GO_ROOT_SYS_ENV";
		Sys.println("env.before=" + (Sys.getEnv(envName) == null));
		Sys.putEnv(envName, "portable");
		var env = Sys.environment();
		Sys.println("env.get=" + Sys.getEnv(envName));
		Sys.println("env.map=" + env.exists(envName));
		Sys.println("env.value=" + (env.get(envName) == "portable"));
		Sys.putEnv(envName, null);
		Sys.println("env.removed=" + (Sys.getEnv(envName) == null));
		var cwd = Sys.getCwd();
		Sys.println("cwd.present=" + (cwd != null && cwd.length > 0));
		Sys.println("system=" + Sys.systemName());
		Sys.println("shell.code=" + Sys.command("exit 6"));
	}
}
