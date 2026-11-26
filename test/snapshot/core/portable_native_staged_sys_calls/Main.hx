/**
	Why
	Staged `Sys.args` and `Sys.command` convert between a Haxe `Array` and
	`go.NativeSlice` for the runtime. While they were inline, Haxe copied that
	conversion into the caller, so the native authority gate reported a portable
	module that used only the portable `Sys` API.

	What
	A portable module that calls `Sys.args()`, `Sys.command` with an argument
	array, and the shell form of `Sys.command`.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so any gate report
	fails the build; the contract report pins that the gate has no hits.
**/
class Main {
	static function main():Void {
		Sys.println("args=" + (Sys.args() != null));
		Sys.println("direct=" + Sys.command("sh", ["-c", "exit 3"]));
		Sys.println("shell=" + Sys.command("exit 4"));
	}
}
