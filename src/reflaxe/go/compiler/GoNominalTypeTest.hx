package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type.ClassType;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration;
import reflaxe.go.compiler.GoRecoveryDispatch.GoRecoveryCasePlan;

/**
	What: Tests declared Haxe class and interface membership after discovery closes.
	Why: Eager switches miss late classes and cannot name conditional classes safely.
	Go interface satisfaction also admits unrelated, same-shaped Haxe classes.
	How: Intern demanded targets, follow declared ancestry, and let the shared
	availability dispatcher place concrete cases in selected probe files. Each
	test is named from its target's generated type name (`hx_is_nominal_<Type>`)
	and declared in name order, so an edit elsewhere cannot rename or move it.
**/
class GoNominalTypeTest {
	static final symbolPrefix = "hx_is_nominal_";

	final availability:GoRecoveryAvailability;
	final dispatcher:GoRecoveryDispatch;
	final className:ClassType->String;
	final targets:Map<String, ClassType> = [];
	var targetCount = 0;
	var frozen = false;

	public function new(availability:GoRecoveryAvailability, className:ClassType->String) {
		this.availability = availability;
		this.dispatcher = new GoRecoveryDispatch(availability);
		this.className = className;
	}

	public function reset():Void {
		targets.clear();
		targetCount = 0;
		frozen = false;
	}

	public function requirementCount():Int
		return targetCount;

	public function freeze():Void
		frozen = true;

	public function matches(value:GoExpr, target:ClassType):GoExpr {
		var name = className(target);
		if (!targets.exists(name)) {
			if (frozen)
				Context.fatalError("Nominal type test discovered after freeze", Context.currentPos());
			if (!GoRecoveryDispatch.isGoIdentifier(name))
				Context.fatalError("Nominal type test requires a generated class or interface, got " + name, Context.currentPos());
			targets.set(name, target);
			targetCount++;
		}
		var symbol = GoRecoveryDispatch.entrySymbol(symbolPrefix, name);
		availability.noteHelper(symbol, GoType.functionType([GoType.parse("any")], [GoType.parse("bool")]));
		availability.requireHelper(symbol);
		return GoExpr.GoCall(GoExpr.GoIdent(symbol), [value]);
	}

	function belongsTo(source:ClassType, target:ClassType, visited:Map<String, Bool>):Bool {
		var name = className(source);
		if (name == className(target))
			return true;
		if (visited.exists(name))
			return false;
		visited.set(name, true);
		if (source.superClass != null && belongsTo(source.superClass.t.get(), target, visited))
			return true;
		for (implemented in source.interfaces)
			if (belongsTo(implemented.t.get(), target, visited))
				return true;
		return false;
	}

	/** Only emitted concrete layouts become cases; interfaces remain nominal targets. */
	public function discover(candidates:Array<ClassType>):Array<GoGuardedRecoveryDeclaration> {
		var declarations = new Array<GoGuardedRecoveryDeclaration>();
		var names = [for (name in targets.keys()) name];
		names.sort(GoRecoveryDispatch.compareSymbols);
		for (name in names) {
			var cases = new Array<GoRecoveryCasePlan>();
			for (candidate in candidates)
				if (belongsTo(candidate, targets.get(name), new Map()))
					cases.push({
						sourceType: GoType.pointer(GoType.named(className(candidate))),
						body: [GoStmt.GoReturn(GoExpr.GoBinary("!=", GoExpr.GoIdent("source"), GoExpr.GoNil))],
						dependencies: []
					});
			var plan = dispatcher.lower({
				symbol: GoRecoveryDispatch.entrySymbol(symbolPrefix, name),
				params: [{name: "value", typeName: "any"}],
				result: GoType.parse("bool"),
				prefix: [
					GoStmt.GoAssign(GoExpr.GoIdent("value"), GoExpr.GoCall(GoExpr.GoIdent("hxrt.ObjectSource"), [GoExpr.GoIdent("value")]))
				],
				input: GoExpr.GoIdent("value"),
				cases: cases,
				fallback: [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))]
			});
			if (plan != null)
				declarations = declarations.concat(plan);
		}
		return declarations;
	}
}
#end
