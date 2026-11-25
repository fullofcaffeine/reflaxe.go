import Helper.native;

/**
	Why
	A caller of a `@:goNative` function whose own signature has a `go.*` type
	uses that `go.*` API, even when the callee's module is native.

	What
	A portable module that calls a module-level function returning `go.NativeAny`.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so the gate must
	fail the build for module `Main`.
**/
class Main {
	static function main():Void {
		final value = native();
		Sys.println(value == null);
	}
}
