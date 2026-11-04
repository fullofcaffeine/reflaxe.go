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
	availability dispatcher place concrete cases in selected probe files.
**/
class GoNominalTypeTest {
	final availability:GoRecoveryAvailability;
	final dispatcher:GoRecoveryDispatch;
	final className:ClassType->String;
	final targets:Array<ClassType> = [];
	var frozen = false;

	public function new(availability:GoRecoveryAvailability, className:ClassType->String) {
		this.availability = availability;
		this.dispatcher = new GoRecoveryDispatch(availability);
		this.className = className;
	}

	public function reset():Void {
		targets.resize(0);
		frozen = false;
	}

	public function requirementCount():Int
		return targets.length;

	public function freeze():Void
		frozen = true;

	public function matches(value:GoExpr, target:ClassType):GoExpr {
		var index = -1;
		for (i in 0...targets.length)
			if (className(targets[i]) == className(target)) {
				index = i;
				break;
			}
		if (index < 0) {
			if (frozen)
				Context.fatalError("Nominal type test discovered after freeze", Context.currentPos());
			index = targets.length;
			targets.push(target);
		}
		var symbol = "hx_is_nominal_" + index;
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
		for (index in 0...targets.length) {
			var cases = new Array<GoRecoveryCasePlan>();
			for (candidate in candidates)
				if (belongsTo(candidate, targets[index], new Map()))
					cases.push({
						sourceType: GoType.pointer(GoType.named(className(candidate))),
						body: [GoStmt.GoReturn(GoExpr.GoBinary("!=", GoExpr.GoIdent("source"), GoExpr.GoNil))],
						dependencies: []
					});
			var plan = dispatcher.lower({
				symbol: "hx_is_nominal_" + index,
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
