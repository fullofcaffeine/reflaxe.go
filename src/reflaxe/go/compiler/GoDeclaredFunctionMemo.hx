package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type.ClassField;
import haxe.macro.Type.ClassType;
import haxe.macro.Type.TFunc;

/**
	What: Remembers, for one compilation, the declared function of each class
	method: the `TFunc` that the `decode` function given to the constructor
	returns for it, or null.

	Why: `GoCompiler.resolveFunctionInfo` reads the declared function of the
	callee of every call to apply source default arguments, and
	`buildStaticFunctionInfoTable` reads every static method of the project. Each
	read decoded the whole method body with `ClassField.expr()`, also when the
	static table already held the answer. On the BeadsHX bootstrap compile that
	was 9,090 calls and 8,413 decoded bodies, 2.1-2.5 s of the compile. Reflaxe
	compiles after typing, filters, and DCE are finished, so the typed AST of a
	method does not change during the compilation, and its first decode gives the
	answer of every later one.

	How: A method is identified by its declaring class (module and name), its
	name, and the offsets of its source position. A static and an instance field
	of one class cannot share a name, and overloads (`overload extern inline`)
	share their class and name but not their position in the class's file, so two
	methods never share a key, also when a macro generates them at one position. The first `get` for a key
	decodes the field and keeps the result, also null; later calls return the
	kept result. The kept `TFunc` objects are shared, so callers must only read
	them, as `GoCompiler` does.
**/
class GoDeclaredFunctionMemo {
	final decode:ClassField->Null<TFunc>;
	final functions:Map<String, Null<TFunc>> = [];

	public function new(decode:ClassField->Null<TFunc>) {
		this.decode = decode;
	}

	/** The declared function of `field`, a static or instance field declared by `owner`. */
	public function get(owner:ClassType, field:ClassField):Null<TFunc> {
		final key = keyOf(owner, field);
		if (functions.exists(key)) {
			return functions.get(key);
		}
		final func = decode(field);
		functions.set(key, func);
		return func;
	}

	static function keyOf(owner:ClassType, field:ClassField):String {
		final position = Context.getPosInfos(field.pos);
		return owner.module + "|" + owner.name + "|" + field.name + "|" + position.min + "|" + position.max;
	}
}
#else
class GoDeclaredFunctionMemo {}
#end
