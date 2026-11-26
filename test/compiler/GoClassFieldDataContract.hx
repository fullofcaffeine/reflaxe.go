#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.GoReflaxeCompiler;
#end

/**
	Contract for the per-field data that the Go compiler asks Reflaxe to build.

	Why
	Before it calls `compileClassImpl`, Reflaxe builds a `ClassFuncData` for every
	method, decoding the method's body again, and a `ClassVarData` for every
	variable. `GoReflaxeCompiler.compileClassImpl` ignores both, because the Go
	lowering reads fields from the `ClassType`. Answering false from
	`shouldGenerateClassField` turns that work off. Reflaxe's one visible effect
	of building the data must stay: the warning "Function information not
	found." for a method whose type is not a function type.

	What
	- Every field of `GoTypeUsageTrackerContractFixture` (constructor, instance
	  and static methods, instance and static variables) gets false, and no
	  warning.
	- A method whose type is not a function type gets false and exactly one
	  such warning at its position, and no warning when
	  `ignoreBodilessFunctions` is on.

	How
	After the initialization macros, types every fixture field, asks new
	`GoReflaxeCompiler` instances, and counts the warnings that
	`Context.getMessages()` reports at each position.
**/
class GoClassFieldDataContract {
	#if macro
	static inline final WARNING = "Function information not found.";

	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoClassFieldDataContract: " + message, Context.currentPos());
	}

	static function warningsAt(pos:haxe.macro.Expr.Position):Int {
		var count = 0;
		final wanted = Context.getPosInfos(pos);
		for (message in Context.getMessages())
			switch (message) {
				case Warning(text, at) if (text == WARNING):
					final infos = Context.getPosInfos(at);
					if (infos.file == wanted.file && infos.min == wanted.min && infos.max == wanted.max)
						count++;
				case _:
			}
		return count;
	}

	/** A copy of `field` whose type is `type`. */
	static function withType(field:ClassField, type:Type):ClassField {
		return {
			name: field.name,
			type: type,
			isPublic: field.isPublic,
			isExtern: field.isExtern,
			isFinal: field.isFinal,
			isAbstract: field.isAbstract,
			params: field.params,
			meta: field.meta,
			kind: field.kind,
			expr: field.expr,
			pos: field.pos,
			doc: field.doc,
			overloads: field.overloads
		};
	}

	static function fixtureFields():Array<ClassField> {
		final owner = switch (Context.getType("GoTypeUsageTrackerContractFixture")) {
			case TInst(ref, _): ref.get();
			case _: Context.fatalError("GoClassFieldDataContract: fixture is not a class", Context.currentPos());
		};
		final fields = owner.fields.get().concat(owner.statics.get());
		if (owner.constructor != null)
			fields.push(owner.constructor.get());
		return fields;
	}

	static function run():Void {
		// Reflaxe asks about fields after Haxe has typed them. Here typing is still
		// lazy: `field.expr()` types a field, and a later read of the class sees
		// its function type. Read every body once, then read the class again.
		for (field in fixtureFields())
			field.expr();
		final fields = fixtureFields();

		final compiler = new GoReflaxeCompiler();
		var methods = 0;
		var variables = 0;
		var method:Null<ClassField> = null;
		for (field in fields) {
			switch (field.kind) {
				case FMethod(_):
					methods++;
					method = field;
				case FVar(_, _):
					variables++;
			}
			require(!compiler.shouldGenerateClassField(field), "no Reflaxe data for " + field.name);
			require(warningsAt(field.pos) == 0, "no warning for " + field.name);
		}
		require(methods >= 5 && variables >= 5, "fixture fields were asked about (" + methods + " methods, " + variables + " variables)");
		if (method == null)
			return;

		// Reflaxe warns once for a method it cannot describe; so does the compiler.
		final untypedMethod = withType(method, TDynamic(null));
		require(!compiler.shouldGenerateClassField(untypedMethod), "no Reflaxe data for a method without a function type");
		require(warningsAt(untypedMethod.pos) == 1, "one warning for a method without a function type (" + warningsAt(untypedMethod.pos) + ")");

		final quiet = new GoReflaxeCompiler();
		quiet.setOptions({ignoreBodilessFunctions: true});
		require(!quiet.shouldGenerateClassField(untypedMethod), "no Reflaxe data with ignoreBodilessFunctions");
		require(warningsAt(untypedMethod.pos) == 1, "no further warning with ignoreBodilessFunctions");

		// The one expected warning is this contract's own; keep it out of the output.
		final expected = Context.getPosInfos(untypedMethod.pos);
		Context.filterMessages(message -> switch (message) {
			case Warning(text, at) if (text == WARNING):
				final infos = Context.getPosInfos(at);
				!(infos.file == expected.file && infos.min == expected.min && infos.max == expected.max);
			case _: true;
		});

		Sys.println("class field data contract passed");
	}
	#end
}
