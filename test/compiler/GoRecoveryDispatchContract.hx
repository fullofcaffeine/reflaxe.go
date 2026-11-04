import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoASTPrinter;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoBuildConstraint;
import reflaxe.go.compiler.GoRecoveryAvailability;
import reflaxe.go.compiler.GoRecoveryDispatch;

/** Native Go checks all four tag selections against independently authored outcomes. */
class GoRecoveryDispatchContract {
	static var output:String;

	static function write(name:String, declarations:Array<GoDecl>, selection:GoBuildConstraint):Void {
		if (selection.isNever())
			return;
		var text = GoASTPrinter.printFile({packageName: "main", imports: [], decls: declarations});
		if (!selection.isAlways())
			text = "//go:build " + selection.render() + "\n\n" + text;
		sys.io.File.saveContent(output + "/" + name + ".go", text);
	}

	static function check(condition:GoExpr, message:String):GoStmt {
		return GoStmt.GoIf(condition, [
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("panic"), [GoExpr.GoStringLiteral(message)]))
		], null);
	}

	static function instance(name:String, value:Int):GoExpr {
		return GoExpr.GoUnary("&", GoExpr.GoCompositeLiteral(name, [GoCompositeElement.GoCompositeField("Value", GoExpr.GoIntLiteral(value))]));
	}

	static function main():Void {
		output = Sys.args()[0];
		var common = GoBuildConstraint.always();
		var alpha = GoBuildConstraint.parse("alpha");
		var beta = GoBuildConstraint.parse("beta");
		var availability = new GoRecoveryAvailability(_ -> null);
		availability.noteGeneratedType("Plain", "common", common);
		availability.noteGeneratedType("Alpha", "alpha", alpha);
		availability.noteGeneratedType("Beta", "beta", beta);
		availability.noteGeneratedType("Target", "target", alpha);
		var dispatcher = new GoRecoveryDispatch(availability);
		var cases = [
			for (name in ["Plain", "Alpha", "Beta"])
				{
					sourceType: GoType.pointer(GoType.named(name)),
					body: [GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("source"), "Value"))],
					dependencies: []
				}
		];
		var number = dispatcher.lower({
			symbol: "recover_number",
			params: [{name: "value", typeName: "any"}],
			result: GoType.parse("int"),
			prefix: [],
			input: GoExpr.GoIdent("value"),
			cases: cases,
			fallback: [GoStmt.GoReturn(GoExpr.GoIntLiteral(-1))]
		});
		var target = dispatcher.lower({
			symbol: "recover_target",
			params: [{name: "value", typeName: "any"}],
			result: GoType.parse("*Target"),
			prefix: [],
			input: GoExpr.GoIdent("value"),
			cases: [
				{
					sourceType: GoType.parse("*Beta"),
					dependencies: [],
					body: [
						GoStmt.GoReturn(GoExpr.GoUnary("&", GoExpr.GoCompositeLiteral("Target", [
							GoCompositeElement.GoCompositeField("Value", GoExpr.GoSelector(GoExpr.GoIdent("source"), "Value"))
						])))
					]
				}
			],
			fallback: [GoStmt.GoReturn(GoExpr.GoNil)]
		});
		var nullable = dispatcher.lower({
			symbol: "recover_nullable",
			params: [{name: "value", typeName: "any"}],
			result: GoType.parse("*Plain"),
			prefix: [],
			input: GoExpr.GoIdent("value"),
			cases: [
				{
					sourceType: GoType.parse("*Alpha"),
					dependencies: [],
					body: [
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("source")),
						GoStmt.GoReturn(GoExpr.GoNil)
					]
				}
			],
			fallback: [GoStmt.GoReturn(instance("Plain", -1))]
		});
		if (number == null || target == null || nullable == null)
			throw "known ownership was unresolved";
		var index = 0;
		for (plan in number.concat(target).concat(nullable))
			write("recovery_" + index++, [plan.declaration], plan.selection);
		write("common", [
			GoDecl.GoStructDecl("Plain", [{name: "Value", typeName: "int"}]),
			GoDecl.GoFuncDecl("main", null, [], [], [
				check(GoExpr.GoBinary("!=", GoExpr.GoCall(GoExpr.GoIdent("recover_number"), [GoExpr.GoNil]), GoExpr.GoIntLiteral(-1)), "miss lost fallback"),
				check(GoExpr.GoBinary("!=", GoExpr.GoCall(GoExpr.GoIdent("recover_number"), [instance("Plain", 7)]), GoExpr.GoIntLiteral(7)),
					"common source lost")
			])
		], common);
		for (item in [{name: "Alpha", guard: alpha, value: 11}, {name: "Beta", guard: beta, value: 22}]) {
			write("source_" + item.name.toLowerCase(), [
				GoDecl.GoStructDecl(item.name, [{name: "Value", typeName: "int"}]),
				GoDecl.GoFuncDecl("init", null, [], [], [
					check(GoExpr.GoBinary("!=", GoExpr.GoCall(GoExpr.GoIdent("recover_number"), [instance(item.name, item.value)]),
						GoExpr.GoIntLiteral(item.value)),
						"selected source lost: "
						+ item.name)
				])
			], item.guard);
		}
		write("nullable_match", [
			GoDecl.GoFuncDecl("init", null, [], [], [
				check(GoExpr.GoBinary("!=", GoExpr.GoCall(GoExpr.GoIdent("recover_nullable"), [instance("Alpha", 0)]), GoExpr.GoNil),
					"matched nil became a miss")
			])
		], alpha);
		write("nullable_miss", [
			GoDecl.GoFuncDecl("init", null, [], [], [
				GoStmt.GoVarDecl("miss", "*Plain", GoExpr.GoCall(GoExpr.GoIdent("recover_nullable"), [GoExpr.GoNil]), true),
				check(GoExpr.GoBinary("!=", GoExpr.GoSelector(GoExpr.GoIdent("miss"), "Value"), GoExpr.GoIntLiteral(-1)), "miss became matched nil")
			])
		], common);
		write("target_type", [GoDecl.GoStructDecl("Target", [{name: "Value", typeName: "int"}])], alpha);
		write("overlap", [
			GoDecl.GoFuncDecl("init", null, [], [], [
				GoStmt.GoVarDecl("result", "*Target", GoExpr.GoCall(GoExpr.GoIdent("recover_target"), [instance("Beta", 22)]), true),
				check(GoExpr.GoBinary("!=", GoExpr.GoSelector(GoExpr.GoIdent("result"), "Value"), GoExpr.GoIntLiteral(22)), "overlap lost conversion"),
				check(GoExpr.GoBinary("!=", GoExpr.GoCall(GoExpr.GoIdent("recover_target"), [GoExpr.GoNil]), GoExpr.GoNil), "conditional fallback changed null")
			])
		], alpha.and(beta));
		availability.freeze();
		Sys.println("ok");
	}
}
