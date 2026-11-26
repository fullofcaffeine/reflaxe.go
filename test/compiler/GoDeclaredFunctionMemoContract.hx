#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.compiler.GoDeclaredFunctionMemo;
#end

#if macro
private typedef MemoQuery = {
	final label:String;
	final owner:ClassType;
	final fetch:() -> ClassField;
}
#end

/**
	Contract for the per-compilation memo of declared method functions.

	Why
	`GoCompiler` asks `GoDeclaredFunctionMemo` for the declared function of each
	callee instead of decoding the callee's body again for every call. A memo key
	that two methods share would hand one method's default arguments to calls of
	another, and a memo that decodes again would bring back the cost it removes.

	What
	Asks one memo about static and instance methods, a dynamic method, and a
	variable of fixture classes, in three orders, and checks that:
	- every answer prints exactly like a fresh decode of that field, including
	  its default arguments and body;
	- same-named methods of different classes, an override and the method it
	  overrides, same-named private classes of two modules, two overloads of one
	  class, and macro-generated methods that share a position (in one class,
	  in two classes, and in same-named private classes of two modules) keep
	  their own answers;
	- each field is decoded once, also when its answer is null.

	How
	After the initialization macros, loads the fixture classes. The memo and the
	fresh decodes use `declaredFunction`, a copy of
	`GoCompiler.stableDeclaredFunction`; the memo's copy counts its calls.
**/
class GoDeclaredFunctionMemoContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoDeclaredFunctionMemoContract: " + message, Context.currentPos());
	}

	static function classNamed(path:String):ClassType {
		return switch (Context.getType(path)) {
			case TInst(ref, _): ref.get();
			case _: throw "not a class: " + path;
		};
	}

	/** The class named `name` declared in `module`, also when it is private. */
	static function moduleClass(module:String, name:String):ClassType {
		for (type in Context.getModule(module))
			switch (type) {
				case TInst(ref, _) if (ref.get().name == name):
					return ref.get();
				case _:
			}
		throw "no class " + name + " in " + module;
	}

	/** Asks about field `name` of `owner`; each ask reads the field again, as lowering does. */
	static function query(owner:ClassType, name:String, isStatic:Bool):MemoQuery {
		return {
			label: owner.name + "." + name,
			owner: owner,
			fetch: () -> fieldNamed(owner, name, isStatic)
		};
	}

	static function fieldNamed(owner:ClassType, name:String, isStatic:Bool):ClassField {
		for (field in (isStatic ? owner.statics : owner.fields).get())
			if (field.name == name)
				return field;
		throw "no field " + name + " in " + owner.name;
	}

	/** `GoCompiler.stableDeclaredFunction` with its `unwrapFunction`. */
	static function declaredFunction(field:ClassField):Null<TFunc> {
		return switch (field.kind) {
			case FMethod(MethDynamic): null;
			case FMethod(_): unwrap(field.expr());
			case _: null;
		};
	}

	static function unwrap(expr:Null<TypedExpr>):Null<TFunc> {
		if (expr == null)
			return null;
		return switch (expr.expr) {
			case TFunction(func): func;
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _): unwrap(inner);
			case _: null;
		};
	}

	static function describe(func:Null<TFunc>):String {
		if (func == null)
			return "null";
		final args = [
			for (arg in func.args)
				arg.v.name + "=" + (arg.value == null ? "" : TypedExprTools.toString(arg.value))
		];
		return "(" + args.join(", ") + ") " + TypedExprTools.toString(func.expr);
	}

	static function run():Void {
		final fixture = classNamed("GoDeclaredFunctionMemoContractFixture");
		final other = classNamed("GoDeclaredFunctionMemoContractFixture.MemoOther");
		final child = classNamed("GoDeclaredFunctionMemoContractFixture.MemoChild");
		final hidden = moduleClass("GoDeclaredFunctionMemoContractFixture", "MemoHidden");
		final otherHidden = moduleClass("GoDeclaredFunctionMemoContractOtherFixture", "MemoHidden");
		final overloads = classNamed("GoDeclaredFunctionMemoContractFixture.MemoOverloads");
		final queries:Array<MemoQuery> = [
			query(fixture, "run", true),
			query(fixture, "apply", false),
			query(fixture, "hook", false),
			query(fixture, "count", false),
			query(other, "run", true),
			query(other, "apply", false),
			query(child, "apply", false),
			query(hidden, "run", true),
			query(otherHidden, "run", true),
			query(overloads, "pick", true),
			query(classNamed("GoDeclaredFunctionMemoContractFixture.MemoMadeA"), "made", true),
			query(classNamed("GoDeclaredFunctionMemoContractFixture.MemoMadeA"), "also", true),
			query(classNamed("GoDeclaredFunctionMemoContractFixture.MemoMadeB"), "made", true),
			query(hidden, "made", true),
			query(otherHidden, "made", true),
			{
				label: "MemoOverloads.pick (second overload)",
				owner: overloads,
				fetch: () -> fieldNamed(overloads, "pick", true).overloads.get()[0]
			}
		];
		final expected = [for (query in queries) describe(declaredFunction(query.fetch()))];
		final distinct = new Map<String, Bool>();
		for (answer in expected)
			distinct.set(answer, true);
		require([for (_ in distinct.keys()) true].length == queries.length - 1,
			"fixture answers must differ, apart from the two null answers: " + expected.join(" | "));

		var decodes = 0;
		final memo = new GoDeclaredFunctionMemo(field -> {
			decodes++;
			declaredFunction(field);
		});
		final forward = [for (index in 0...queries.length) index];
		final backward = forward.copy();
		backward.reverse();
		for (order in [forward, backward, forward]) {
			for (index in order) {
				final query = queries[index];
				final answer = describe(memo.get(query.owner, query.fetch()));
				require(answer == expected[index], query.label + " answers " + answer + ", fresh decode " + expected[index]);
			}
			require(decodes == queries.length, "decoded " + decodes + " times for " + queries.length + " fields");
		}
	}
	#end
}
