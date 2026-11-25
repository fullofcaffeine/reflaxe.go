import Helper.plain;

/**
	Why
	The native authority gate once counted every static of a type that an
	expression names, so calling a go-free function of a `@:goNative` module
	reported the caller when a sibling function had a `go.*` signature.

	What
	A portable module that calls only go-free functions of two `@:goNative`
	modules: a module-level function and a class static, each with a sibling
	that takes a `go.*` value.

	How
	Compiles under `reflaxe_go_portable_native_policy=error`, so any gate report
	fails the build; the contract report pins that the gate has no hits.
**/
class Main {
	static function main():Void {
		Sys.println(plain());
		Sys.println(Tools.plain());
	}
}
