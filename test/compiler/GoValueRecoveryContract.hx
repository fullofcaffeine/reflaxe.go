#if macro
import haxe.macro.Context;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoValueRecovery;
import reflaxe.go.compiler.GoRecoveryAvailability;
import reflaxe.go.compiler.GoBuildConstraint;
#end

/** Compiler-side discovery contracts; generated runtime regressions prove behavior. */
class GoValueRecoveryContract {
	#if macro
	static function newAvailability():GoRecoveryAvailability {
		var availability = new GoRecoveryAvailability(alias -> alias.value() == "hxrt" ? "runtime" : null);
		availability.notePackage("runtime", GoBuildConstraint.always());
		return availability;
	}

	public static function check():Void {
		var owner = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var source = GoType.parse("*hxrt.Function[func(int) int]");
		var target = GoType.parse("*hxrt.Function[func(any) any]");
		owner.noteSource(source);
		owner.noteSource(GoType.parse("*hxrt.Function[func(int) int]"));
		owner.recover(GoExpr.GoIdent("input"), target);
		owner.recover(GoExpr.GoIdent("other"), target);
		if (owner.requirementCount() != 2)
			Context.fatalError("Duplicate callable requirements", Context.currentPos());
		var declarations = owner.discover();
		if (declarations.length != 1)
			Context.fatalError("Duplicate recovery helper", Context.currentPos());
		owner.freeze();
		owner.recover(GoExpr.GoIdent("again"), target);
		if (owner.discover().length != 1)
			Context.fatalError("Frozen recovery plan changed", Context.currentPos());

		// The named source is observed before its module declaration. Recovering
		// the signature from applied Haxe typing would lose this physical ABI.
		var named = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		named.noteNamedSource("optional_source");
		named.noteDeclarations([
			GoDecl.GoFuncDecl("optional_source", null, [{name: "arg", typeName: "int"}], ["int"], [])
		]);
		named.recover(GoExpr.GoIdent("stored"), target);
		if (named.requirementCount() != 2)
			Context.fatalError("Forward named source was lost", Context.currentPos());
		named.discover();
		named.freeze();

		// Recovering a callback argument discovers a second target while the first
		// helper is being built. Both demands must exist before emission freezes.
		var nested = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		nested.noteSource(GoType.parse("*hxrt.Function[func(*hxrt.Function[func(int) int]) int]"));
		nested.recover(GoExpr.GoIdent("stored"), GoType.parse("*hxrt.Function[func(*hxrt.Function[func(any) any]) any]"));
		if (nested.discover().length != 2)
			Context.fatalError("Nested recovery demand was not closed", Context.currentPos());
		nested.freeze();
		if (nested.discover().length != 2)
			Context.fatalError("Frozen nested recovery changed", Context.currentPos());

		// Helper names and declaration order depend on the demanded set, not on the
		// order in which modules first demanded each target.
		var first = GoType.parse("*hxrt.Function[func(int) int]");
		var second = GoType.parse("*hxrt.Function[func(any) any]");
		var forward = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var backward = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var forwardCalls = [
			forward.recover(GoExpr.GoIdent("a"), first),
			forward.recover(GoExpr.GoIdent("b"), second)
		];
		var backwardCalls = [
			backward.recover(GoExpr.GoIdent("b"), second),
			backward.recover(GoExpr.GoIdent("a"), first)
		];
		if (Std.string(forwardCalls[0]) != Std.string(backwardCalls[1]) || Std.string(forwardCalls[1]) != Std.string(backwardCalls[0]))
			Context.fatalError("Callable helper name depended on first use", Context.currentPos());
		if (declarationNames(forward.discover()).join(",") != declarationNames(backward.discover()).join(","))
			Context.fatalError("Callable helper placement depended on first use", Context.currentPos());

		// Alias spellings of one representation share one helper.
		var aliases = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		aliases.recover(GoExpr.GoIdent("a"), GoType.parse("*hxrt.Function[func(byte) int]"));
		aliases.recover(GoExpr.GoIdent("b"), GoType.parse("*hxrt.Function[func(uint8) int]"));
		if (aliases.requirementCount() != 1)
			Context.fatalError("Alias spellings split one callable helper", Context.currentPos());
	}

	static function declarationNames(values:Array<reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration>):Array<String> {
		return [
			for (value in values)
				switch (value.declaration) {
					case GoFuncDecl(name, _, _, _, _):
						name;
					case _:
						"";
				}
		];
	}
	#end
}
