/**
	Regression for haxe_go-ymtk: one non-inline `Sys` call must emit only that
	method. `Sys.getChar` stays out, so the selective runtime must not copy the
	terminal support files. `sys_get_char_method_selection` is the paired control.
**/
class Main {
	static function main() {
		Sys.putEnv("HAXE_GO_SYS_METHOD_SELECTION", "environment");
		var environment = Sys.environment();
		Sys.println(environment.get("HAXE_GO_SYS_METHOD_SELECTION"));
	}
}
