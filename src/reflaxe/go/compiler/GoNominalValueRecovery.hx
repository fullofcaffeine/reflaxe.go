package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type.ClassType;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoType;

/**
	What: Recovers generated class storage after a value crosses an erased boundary.
	Why: A boxed subclass is not a Go pointer to its embedded base. Structural and
	base views must also resolve their canonical object before recovering a class.
	How: Intern demanded targets, then emit typed superclass selector paths after
	the compiler has discovered every reachable generated layout.
**/
class GoNominalValueRecovery {
	final className:ClassType->String;
	final inheritancePath:(ClassType, ClassType) -> Null<Array<ClassType>>;
	final targets:Array<ClassType> = [];
	var frozen = false;

	public function new(className:ClassType->String, inheritancePath:(ClassType, ClassType) -> Null<Array<ClassType>>) {
		this.className = className;
		this.inheritancePath = inheritancePath;
	}

	public function reset():Void {
		targets.resize(0);
		frozen = false;
	}

	public function requirementCount():Int
		return targets.length;

	public function freeze():Void
		frozen = true;

	public function recover(value:GoExpr, target:ClassType):GoExpr {
		var index = -1;
		for (i in 0...targets.length)
			if (className(targets[i]) == className(target)) {
				index = i;
				break;
			}
		if (index < 0) {
			if (frozen)
				Context.fatalError("Nominal recovery discovered after freeze", Context.currentPos());
			index = targets.length;
			targets.push(target);
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hx_recover_class_" + index), [value]);
	}

	/** Candidates come from the compiler's emitted, same-package class layouts. */
	public function discover(candidates:Array<ClassType>):Array<GoDecl> {
		var declarations = new Array<GoDecl>();
		for (index in 0...targets.length) {
			var target = targets[index];
			var targetType = GoType.pointer(GoType.named(className(target)));
			var cases = new Array<GoTypeSwitchCase>();
			for (candidate in candidates) {
				var path = inheritancePath(candidate, target);
				if (path == null)
					continue;
				var result = GoExpr.GoIdent("source");
				for (parent in path)
					result = GoExpr.GoSelector(result, className(parent));
				cases.push({typeName: GoType.pointer(GoType.named(className(candidate))), body: [GoStmt.GoReturn(result)]});
			}
			declarations.push(GoDecl.GoFuncDecl("hx_recover_class_" + index, null, [{name: "value", typeName: "any"}], [targetType], [
				GoStmt.GoAssign(GoExpr.GoIdent("value"), GoExpr.GoCall(GoExpr.GoIdent("hxrt.ObjectSource"), [GoExpr.GoIdent("value")])),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoTypeSwitch(GoExpr.GoIdent("value"), "source", cases, [GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent("value"), targetType))])
			]));
		}
		return declarations;
	}
}
#end
