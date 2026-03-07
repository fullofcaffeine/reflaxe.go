package reflaxe.go.macros;

#if macro
import haxe.macro.Expr;
#end

/**
	GoInjection: a typed macro shim around `untyped __go__(...)`.

	Why:
	- Framework-owned std/runtime layers sometimes need a small raw Go escape hatch for
	  low-level abstractions that would be worse as compiler-owned `GoRaw` blobs.
	- Calling `untyped __go__("...")` directly works, but it forces every callsite to use
	  `untyped` and makes the intended ownership boundary less obvious.
	- `__go__` does not infer package imports; package-level interop should still flow through
	  typed extern metadata such as `@:go.import`, `@:go.name`, and `@:go.receiver`.

	What:
	- `GoInjection.__go__(code, args)` expands to `untyped __go__(code, ...args)`.
	- The `haxe.go` backend recognizes that call and emits the raw Go snippet directly.

	How:
	- Use `{0}`, `{1}`, ... placeholders inside `code`; the backend lowers the referenced
	  Haxe expressions to Go expressions and prints them into the snippet.
	- If there are no placeholders, `code` is emitted verbatim as raw Go.
	- Keep this in framework-owned layers (`std/`, runtime bindings, narrow abstraction
	  helpers), not in application business logic.
**/
class GoInjection {
	public static macro function __go__(code:String, args:Array<Expr>):Expr {
		var callArgs = [macro $v{code}].concat(args);
		return macro untyped __go__($a{callArgs});
	}
}
