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
	}
	#end
}
