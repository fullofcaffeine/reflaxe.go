class Main {
	static function main() {
		var envName = "HAXE_GO_ROOT_SYS_ENV";
		Sys.println("env.before=" + (Sys.getEnv(envName) == null));
		Sys.putEnv(envName, "portable");
		var env = Sys.environment();
		Sys.println("env.get=" + Sys.getEnv(envName));
		Sys.println("env.map=" + env.exists(envName));
		var cwd = Sys.getCwd();
		Sys.println("cwd.present=" + (cwd != null && cwd.length > 0));
		Sys.println("system=" + Sys.systemName());
	}
}
