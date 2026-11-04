package reflaxe.go.compiler;

import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryAvailability.GoRecoveryDependency;

typedef GoRecoveryCasePlan = {
	final sourceType:GoType;
	final body:Array<GoStmt>;
	final dependencies:Array<GoRecoveryDependency>;
}

typedef GoGuardedRecoveryDeclaration = {
	final selection:GoBuildConstraint;
	final declaration:GoDecl;
}

typedef GoRecoveryEntry = {
	final symbol:String;
	final params:Array<GoParam>;
	final result:GoType;
	final prefix:Array<GoStmt>;
	final input:GoExpr;
	final cases:Array<GoRecoveryCasePlan>;
	final fallback:Array<GoStmt>;
}

private typedef RecoveryGroup = {
	final selection:GoBuildConstraint;
	final cases:Array<GoTypeSwitchCase>;
}

/**
	Keeps a recovery entry available wherever its signature exists. Conditional
	source cases live in complementary probe files; their explicit match bit
	distinguishes a recovered nil from a miss. No runtime registration is needed.
**/
class GoRecoveryDispatch {
	final availability:GoRecoveryAvailability;

	public function new(availability:GoRecoveryAvailability) {
		this.availability = availability;
	}

	/** Returns null until all source and signature owners have been discovered. */
	public function lower(entry:GoRecoveryEntry):Null<Array<GoGuardedRecoveryDeclaration>> {
		if (entry.params.length == 0 || !entry.params[0].typeName.sameRepresentation(GoType.parse("any")))
			throw new haxe.Exception("Recovery entry requires an erased first parameter");
		var signature = GoType.functionType([for (param in entry.params) param.typeName], [entry.result]);
		availability.noteHelper(entry.symbol, signature);
		var presence = availability.conditionForType(signature);
		if (presence == null)
			return null;
		if (presence.isNever())
			return [];
		var direct = new Array<GoTypeSwitchCase>();
		var groups = new Map<String, RecoveryGroup>();
		var complete = true;
		for (candidate in entry.cases) {
			var dependencies = candidate.dependencies.copy();
			dependencies.push(TypeReference(candidate.sourceType));
			var selection = availability.conditionFor(dependencies);
			if (selection == null) {
				complete = false;
				continue;
			}
			if (presence.and(selection).isNever())
				continue;
			var arm:GoTypeSwitchCase = {typeName: candidate.sourceType, body: candidate.body};
			if (selection.isAlways() || presence.and(selection).key() == presence.key()) {
				direct.push(arm);
			} else {
				var key = selection.key();
				var group = groups.get(key);
				if (group == null) {
					group = {selection: selection, cases: []};
					groups.set(key, group);
				}
				group.cases.push(arm);
			}
		}
		if (!complete)
			return null;
		var output = new Array<GoGuardedRecoveryDeclaration>();
		var body = entry.prefix.copy();
		var keys = [for (key in groups.keys()) key];
		keys.sort((a, b) -> a < b ? -1 : a > b ? 1 : 0);
		if (keys.length == 0) {
			if (direct.length == 0)
				body = body.concat(entry.fallback);
			else
				body.push(GoTypeSwitch(entry.input, "source", direct, entry.fallback));
		} else {
			if (direct.length > 0)
				body.push(GoTypeSwitch(entry.input, "source", direct, null));
			for (index in 0...keys.length) {
				var group = groups.get(keys[index]);
				var symbol = entry.symbol + "__probe_" + index;
				var results = [entry.result, GoType.parse("bool")];
				availability.noteHelper(symbol, GoType.functionType([for (param in entry.params) param.typeName], results));
				var probeCases = [for (arm in group.cases) {typeName: arm.typeName, body: matchedBody(arm.body)}];
				output.push({
					selection: presence.and(group.selection),
					declaration: GoFuncDecl(symbol, null, entry.params, results, [
						GoTypeSwitch(GoIdent(entry.params[0].name), "source", probeCases, missedBody(entry.result))
					])
				});
				var stub = presence.and(group.selection.negate());
				if (!stub.isNever())
					output.push({selection: stub, declaration: GoFuncDecl(symbol, null, entry.params, results, missedBody(entry.result))});
				var arguments = [entry.input];
				for (index in 1...entry.params.length)
					arguments.push(GoIdent(entry.params[index].name));
				body.push(GoMultiAssign(["hx_recovered", "hx_matched"], GoCall(GoIdent(symbol), arguments), index == 0));
				body.push(GoIf(GoIdent("hx_matched"), [GoReturn(GoIdent("hx_recovered"))], null));
			}
			body = body.concat(entry.fallback);
		}
		output.unshift({selection: presence, declaration: GoFuncDecl(entry.symbol, null, entry.params, [entry.result], body)});
		return output;
	}

	/** Builders end each case with one outer return; nested closures retain their ABI. */
	static function matchedBody(body:Array<GoStmt>):Array<GoStmt> {
		var result = body.copy();
		if (result.length == 0)
			throw new haxe.Exception("Recovery case has no result");
		result[result.length - 1] = switch (result[result.length - 1]) {
			case GoReturn(value) if (value != null): GoReturnValues([value, GoIdent("true")]);
			case _: throw new haxe.Exception("Recovery case must end with a value return");
		};
		return result;
	}

	static function missedBody(result:GoType):Array<GoStmt> {
		return [
			GoVarDecl("zero", result, null, false),
			GoReturnValues([GoIdent("zero"), GoIdent("false")])
		];
	}
}
