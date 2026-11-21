/**
	Paired control for `sys_environment_method_selection` (haxe_go-ymtk): the
	same program with a reachable `Sys.getChar` call must still emit that method
	and copy the terminal runtime. The call runs only with a `read` argument, so
	the runtime smoke stays non-interactive.
**/
class Main {
	static function main() {
		Sys.putEnv("HAXE_GO_SYS_METHOD_SELECTION", "environment");
		var environment = Sys.environment();
		Sys.println(environment.get("HAXE_GO_SYS_METHOD_SELECTION"));
		var arguments = Sys.args();
		if (arguments.length > 0 && arguments[0] == "read")
			Sys.println(Sys.getChar(false));
	}
}
