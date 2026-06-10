package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;

typedef GoLambdaLoweredExpr = {
	final expr:GoExpr;
	final isStringLike:Bool;
}

typedef GoLambdaSourcePlan = {
	final domain:String;
	final elementType:String;
	final sourceExpr:GoExpr;
	final sourceType:String;
}

private typedef GoLambdaIterableLoweringConfig = {
	final lowerExpr:TypedExpr->GoLambdaLoweredExpr;
	final freshTempName:String->String;
	final isArrayType:Type->Bool;
	final arrayElementType:Type->Null<Type>;
	final arrayElementGoType:Type->String;
	final haxeDsListElementType:Type->Null<Type>;
	final scalarGoType:Type->String;
	final requireStdlibShimGroup:String->Void;
	final lowerNullableAwareTypeAssertExpr:(GoExpr, Type) -> GoExpr;
	final localVarName:TVar->String;
	final lookupLocalLambdaAlias:String->Null<String>;
}

/**
	What:
	Compiler-owned Lambda/Iterable lowering policy for Go.

	Why:
	`Lambda` itself remains source-owned stdlib code, but generic `Iterable<T>`
	call sites need backend representation glue: arrays are Go slices, `List` uses
	the current staged list carrier, and unknown iterables use the manual
	`iterator()` protocol. Keeping that policy in one helper keeps `GoCompiler`
	focused on orchestration while still making the representation-sensitive part
	explicit.

	How:
	Consumes typed callbacks from `GoCompiler` for lowering, naming, and type
	mapping, then builds the small adapter expressions used by Lambda call
	lowering. This module should not grow general Haxe stdlib behavior; it owns
	only the bridge between Haxe iterable semantics and current Go carriers.
	Go `any` stays localized here because unknown `Iterable<T>` values expose
	their elements through the manual iterator protocol instead of a statically
	known Go slice type.
**/
class GoLambdaIterableLowering {
	final lowerExpr:TypedExpr->GoLambdaLoweredExpr;
	final freshTempName:String->String;
	final isArrayType:Type->Bool;
	final arrayElementType:Type->Null<Type>;
	final arrayElementGoType:Type->String;
	final haxeDsListElementType:Type->Null<Type>;
	final scalarGoType:Type->String;
	final requireStdlibShimGroup:String->Void;
	final lowerNullableAwareTypeAssertExpr:(GoExpr, Type) -> GoExpr;
	final localVarName:TVar->String;
	final lookupLocalLambdaAlias:String->Null<String>;

	public function new(config:GoLambdaIterableLoweringConfig) {
		lowerExpr = config.lowerExpr;
		freshTempName = config.freshTempName;
		isArrayType = config.isArrayType;
		arrayElementType = config.arrayElementType;
		arrayElementGoType = config.arrayElementGoType;
		haxeDsListElementType = config.haxeDsListElementType;
		scalarGoType = config.scalarGoType;
		requireStdlibShimGroup = config.requireStdlibShimGroup;
		lowerNullableAwareTypeAssertExpr = config.lowerNullableAwareTypeAssertExpr;
		localVarName = config.localVarName;
		this.lookupLocalLambdaAlias = config.lookupLocalLambdaAlias;
	}

	public function trySourcePlan(sourceExpr:TypedExpr):Null<GoLambdaSourcePlan> {
		if (isArrayType(sourceExpr.t)) {
			var elementType = arrayElementGoType(sourceExpr.t);
			var loweredSourceExpr = lowerExpr(sourceExpr).expr;
			return {
				domain: "array",
				elementType: elementType,
				sourceExpr: loweredSourceExpr,
				sourceType: "[]" + elementType
			};
		}

		var listElement = haxeDsListElementType(sourceExpr.t);
		if (listElement != null) {
			requireStdlibShimGroup("ds");
			var loweredSourceExpr = lowerExpr(sourceExpr).expr;
			return {
				domain: "list",
				elementType: scalarGoType(listElement),
				sourceExpr: loweredSourceExpr,
				sourceType: "*haxe__ds__List"
			};
		}

		return null;
	}

	public function manualIteratorProtocolSource(sourceExpr:GoExpr, ?sourcePlan:Null<GoLambdaSourcePlan>):GoExpr {
		var sourceName = freshTempName("hx_lambda_source");
		var wrappedName = freshTempName("hx_lambda_wrapped");
		var iteratorFactoryBody:Array<GoStmt> = switch (sourcePlan == null ? "generic" : sourcePlan.domain) {
			case "array":
				var indexName = freshTempName("hx_lambda_index");
				var valueName = freshTempName("hx_lambda_value");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + indexName + " < len(" + sourceName
					+ ") }, \"next\": func() any { " + valueName + " := " + sourceName + "[" + indexName + "]; " + indexName + "++; return " + valueName +
					" }}";
				[
					GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
			case "list":
				var indexName = freshTempName("hx_lambda_index");
				var valueName = freshTempName("hx_lambda_value");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + indexName + " < len(" + sourceName
					+ ".items) }, \"next\": func() any { " + valueName + " := " + sourceName + ".items[" + indexName + "]; " + indexName + "++; return "
					+ valueName + " }}";
				[
					GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
			case _:
				var iteratorName = freshTempName("hx_lambda_iterator");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + iteratorName + ".hasNext() }, \"next\": func() any { return "
					+ iteratorName + ".next() }}";
				[
					GoStmt.GoVarDecl(iteratorName, null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "iterator"), []), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
		};
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, null, sourceExpr, true),
			GoStmt.GoVarDecl(wrappedName, "map[string]any", GoExpr.GoRaw("map[string]any{}"), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(wrappedName), GoExpr.GoStringLiteral("iterator")),
				GoExpr.GoFuncLiteral([], ["map[string]any"], iteratorFactoryBody)),
			GoStmt.GoReturn(GoExpr.GoIdent(wrappedName))
		]), []);
	}

	public function dynamicIterableSource(sourceExpr:TypedExpr):GoExpr {
		return manualIteratorProtocolSource(lowerExpr(sourceExpr).expr);
	}

	public function firstFunctionArgType(type:Type):Null<Type> {
		return switch (Context.follow(type)) {
			case TFun(args, _):
				args.length > 0 ? args[0].t : null;
			case _:
				null;
		};
	}

	public function functionAliasName(expr:TypedExpr):Null<String> {
		return switch (expr.expr) {
			case TField(_, FStatic(classRef, fieldRef)):
				var classType = classRef.get();
				var field = fieldRef.get();
				if (classType.pack.length == 0 && classType.name == "Lambda" && (field.name == "map" || field.name == "fold")) {
					field.name;
				} else {
					null;
				}
			case TLocal(variable):
				lookupLocalLambdaAlias(localVarName(variable));
			case TMeta(_, inner):
				functionAliasName(inner);
			case TParenthesis(inner):
				functionAliasName(inner);
			case TCast(inner, _):
				functionAliasName(inner);
			case _:
				null;
		};
	}

	public function predicateAnyAdapter(predicateExpr:GoExpr, predicateType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(predicateType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], ["bool"], [GoStmt.GoReturn(GoExpr.GoCall(predicateExpr, [adaptedArgExpr]))]);
	}

	public function mapperAnyAdapter(mapperExpr:GoExpr, mapperType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(mapperType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], ["any"], [GoStmt.GoReturn(GoExpr.GoCall(mapperExpr, [adaptedArgExpr]))]);
	}

	public function consumerAnyAdapter(consumerExpr:GoExpr, consumerType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(consumerType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], [], [GoStmt.GoExprStmt(GoExpr.GoCall(consumerExpr, [adaptedArgExpr]))]);
	}

	public function folderAnyAdapter(folderExpr:GoExpr, folderType:Type):GoExpr {
		var rawValueName = freshTempName("hx_lambda_value");
		var rawAccName = freshTempName("hx_lambda_acc");
		var adaptedValueExpr:GoExpr = GoExpr.GoIdent(rawValueName);
		var adaptedAccExpr:GoExpr = GoExpr.GoIdent(rawAccName);
		switch (Context.follow(folderType)) {
			case TFun(args, _):
				if (args.length > 0) {
					adaptedValueExpr = lowerNullableAwareTypeAssertExpr(adaptedValueExpr, args[0].t);
				}
				if (args.length > 1) {
					adaptedAccExpr = lowerNullableAwareTypeAssertExpr(adaptedAccExpr, args[1].t);
				}
			case _:
		}
		return GoExpr.GoFuncLiteral([{name: rawValueName, typeName: "any"}, {name: rawAccName, typeName: "any"}], ["any"],
			[GoStmt.GoReturn(GoExpr.GoCall(folderExpr, [adaptedValueExpr, adaptedAccExpr]))]);
	}

	public function anyArrayCoerce(anySliceExpr:GoExpr, targetArrayType:Type):GoExpr {
		if (!isArrayType(targetArrayType)) {
			return anySliceExpr;
		}
		var targetElementType = arrayElementType(targetArrayType);
		if (targetElementType == null) {
			return anySliceExpr;
		}
		var targetElementGoType = arrayElementGoType(targetArrayType);
		if (targetElementGoType == "any") {
			return anySliceExpr;
		}
		var rawName = freshTempName("hx_lambda_raw");
		var outName = freshTempName("hx_lambda_out");
		var itemName = freshTempName("hx_lambda_item");
		var convertedItemExpr = lowerNullableAwareTypeAssertExpr(GoExpr.GoIdent(itemName), targetElementType);
		var outType = "[]" + targetElementGoType;
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: rawName, typeName: "[]any"}], [outType], [
			GoStmt.GoVarDecl(outName, outType, GoExpr.GoRaw("make(" + outType + ", 0, len(" + rawName + "))"), true),
			GoStmt.GoRaw("for _, " + itemName + " := range " + rawName + " {"),
			GoStmt.GoAssign(GoExpr.GoIdent(outName), GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(outName), convertedItemExpr])),
			GoStmt.GoRaw("}"),
			GoStmt.GoReturn(GoExpr.GoIdent(outName))
		]), [anySliceExpr]);
	}

	public function isGeneratedCall(callee:TypedExpr, methodName:String):Bool {
		return switch (callee.expr) {
			case TIdent(name):
				name == ("Lambda_" + methodName);
			case TMeta(_, inner):
				isGeneratedCall(inner, methodName);
			case TParenthesis(inner):
				isGeneratedCall(inner, methodName);
			case TCast(inner, _):
				isGeneratedCall(inner, methodName);
			case _:
				false;
		};
	}
}
#end
