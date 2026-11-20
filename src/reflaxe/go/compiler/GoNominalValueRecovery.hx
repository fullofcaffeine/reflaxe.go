package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type.ClassType;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryDispatch.GoRecoveryCasePlan;
import reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration;
import reflaxe.go.compiler.GoRecoveryAvailability.GoRecoveryDependency;

/**
	What: Recovers generated class storage after a value crosses an erased boundary.
	Why: A boxed subclass is not a Go pointer to its embedded base. Structural and
	base views must also resolve their canonical object before recovering a class.
	How: Intern demanded targets, then emit typed superclass selector paths after
	the compiler has discovered every reachable generated layout. Each helper is
	named from its target's generated type name (`hx_recover_class_<Type>`) and
	declared in name order, so an edit elsewhere cannot rename or move it.
**/
class GoNominalValueRecovery {
	static final symbolPrefix = "hx_recover_class_";

	final availability:GoRecoveryAvailability;
	final dispatcher:GoRecoveryDispatch;
	final className:ClassType->String;
	final inheritancePath:(ClassType, ClassType) -> Null<Array<ClassType>>;
	final targets:Map<String, ClassType> = [];
	var targetCount = 0;
	var frozen = false;

	public function new(availability:GoRecoveryAvailability, className:ClassType->String, inheritancePath:(ClassType, ClassType) -> Null<Array<ClassType>>) {
		this.availability = availability;
		this.dispatcher = new GoRecoveryDispatch(availability);
		this.className = className;
		this.inheritancePath = inheritancePath;
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

	public function recover(value:GoExpr, target:ClassType):GoExpr {
		var name = className(target);
		if (!targets.exists(name)) {
			if (frozen)
				Context.fatalError("Nominal recovery discovered after freeze", Context.currentPos());
			if (!GoRecoveryDispatch.isGoIdentifier(name))
				Context.fatalError("Nominal recovery requires a generated class, got " + name, Context.currentPos());
			targets.set(name, target);
			targetCount++;
		}
		var symbol = GoRecoveryDispatch.entrySymbol(symbolPrefix, name);
		availability.noteHelper(symbol, GoType.functionType([GoType.parse("any")], [GoType.pointer(GoType.named(name))]));
		availability.requireHelper(symbol);
		return GoExpr.GoCall(GoExpr.GoIdent(symbol), [value]);
	}

	/** Candidates come from the compiler's emitted, same-package class layouts. */
	public function discover(candidates:Array<ClassType>):Array<GoGuardedRecoveryDeclaration> {
		var declarations = new Array<GoGuardedRecoveryDeclaration>();
		var names = [for (name in targets.keys()) name];
		names.sort(GoRecoveryDispatch.compareSymbols);
		for (name in names) {
			var target = targets.get(name);
			var targetType = GoType.pointer(GoType.named(name));
			var cases = new Array<GoRecoveryCasePlan>();
			for (candidate in candidates) {
				var path = inheritancePath(candidate, target);
				if (path == null)
					continue;
				var result = GoExpr.GoIdent("source");
				for (parent in path)
					result = GoExpr.GoSelector(result, className(parent));
				cases.push({
					sourceType: GoType.pointer(GoType.named(className(candidate))),
					body: [GoStmt.GoReturn(result)],
					dependencies: [
						for (parent in path)
							TypeReference(GoType.pointer(GoType.named(className(parent))))
					]
				});
			}
			var plan = dispatcher.lower({
				symbol: GoRecoveryDispatch.entrySymbol(symbolPrefix, name),
				params: [{name: "value", typeName: "any"}],
				result: targetType,
				prefix: [
					GoStmt.GoAssign(GoExpr.GoIdent("value"), GoExpr.GoCall(GoExpr.GoIdent("hxrt.ObjectSource"), [GoExpr.GoIdent("value")])),
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null)
				],
				input: GoExpr.GoIdent("value"),
				cases: cases,
				fallback: [GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent("value"), targetType))]
			});
			if (plan != null)
				declarations = declarations.concat(plan);
		}
		return declarations;
	}
}
#end
