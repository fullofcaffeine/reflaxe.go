import go.NativeSlice;

/**
	Why
	Staged `Sys.command` keeps its `go.NativeSlice` conversion inside the
	staged `Sys` module, so a portable caller of `Sys.command` is go-free. A
	portable module that also builds a `go.NativeSlice` itself still uses a
	`go.*` API, and the gate must keep reporting it.

	What
	A portable module that calls `Sys.command` and converts an array to
	`go.NativeSlice` directly.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so the gate must
	fail the build for module `Main`.
**/
class Main {
	static function main():Void {
		final args = ["-c", "exit 0"];
		final slice = NativeSlice.fromArray(args);
		Sys.println(Sys.command("sh", args));
		Sys.println(slice == null);
	}
}
