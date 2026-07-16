package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypeTools;
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
		lowerNullableAwareTypeAssertExpr = config.lowerNullableAwareTypeAssertExpr;
		localVarName = config.localVarName;
		this.lookupLocalLambdaAlias = config.lookupLocalLambdaAlias;
	}

	public function trySourcePlan(sourceExpr:TypedExpr):Null<GoLambdaSourcePlan> {
		if (!isArrayType(sourceExpr.t) && haxeDsListElementType(sourceExpr.t) == null) {
			return null;
		}
		return tryValuePlan(lowerExpr(sourceExpr).expr, sourceExpr.t);
	}

	function tryValuePlan(sourceExpr:GoExpr, sourceType:Type):Null<GoLambdaSourcePlan> {
		if (isArrayType(sourceType)) {
			var elementType = arrayElementGoType(sourceType);
			return {
				domain: "array",
				elementType: elementType,
				sourceExpr: sourceExpr,
				sourceType: "[]" + elementType
			};
		}

		var listElement = haxeDsListElementType(sourceType);
		if (listElement != null) {
			return {
				domain: "list",
				elementType: scalarGoType(listElement),
				sourceExpr: sourceExpr,
				sourceType: "*haxe__ds__List"
			};
		}

		return null;
	}

	public function manualIteratorProtocolSource(sourceExpr:GoExpr, ?sourcePlan:Null<GoLambdaSourcePlan>, ?adaptElement:GoExpr->GoExpr,
			?adaptedElementType:String):GoExpr {
		var sourceName = freshTempName("hx_lambda_source");
		var wrappedName = freshTempName("hx_lambda_wrapped");
		var iteratorMapName = freshTempName("hx_lambda_iterator_map");
		var valueName = freshTempName("hx_lambda_value");
		var iteratorFactoryBody = new Array<GoStmt>();
		var hasNextExpr:GoExpr;
		var nextValueExpr:GoExpr;
		var nextPrefix = new Array<GoStmt>();

		if (sourcePlan != null && sourcePlan.domain == "array") {
			var indexName = freshTempName("hx_lambda_index");
			var indexExpr = GoExpr.GoIdent(indexName);
			iteratorFactoryBody.push(GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true));
			hasNextExpr = GoExpr.GoBinary("<", indexExpr, GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent(sourceName)]));
			nextValueExpr = GoExpr.GoIndex(GoExpr.GoIdent(sourceName), indexExpr);
			nextPrefix.push(GoStmt.GoAssign(indexExpr, GoExpr.GoBinary("+", indexExpr, GoExpr.GoIntLiteral(1))));
		} else {
			var iteratorName = freshTempName("hx_lambda_iterator");
			var iteratorExpr = GoExpr.GoIdent(iteratorName);
			iteratorFactoryBody.push(GoStmt.GoVarDecl(iteratorName, null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "iterator"), []), true));
			hasNextExpr = GoExpr.GoCall(GoExpr.GoSelector(iteratorExpr, "hasNext"), []);
			nextValueExpr = GoExpr.GoCall(GoExpr.GoSelector(iteratorExpr, "next"), []);
		}

		var returnedValue:GoExpr = GoExpr.GoIdent(valueName);
		if (adaptElement != null) {
			returnedValue = adaptElement(returnedValue);
		}
		var nextBody = [GoStmt.GoVarDecl(valueName, null, nextValueExpr, true)].concat(nextPrefix).concat([GoStmt.GoReturn(returnedValue)]);
		iteratorFactoryBody = iteratorFactoryBody.concat([
			GoStmt.GoVarDecl(iteratorMapName, "map[string]any", GoExpr.GoRaw("map[string]any{}"), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(iteratorMapName), GoExpr.GoStringLiteral("hasNext")),
				GoExpr.GoFuncLiteral([], ["bool"], [GoStmt.GoReturn(hasNextExpr)])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(iteratorMapName), GoExpr.GoStringLiteral("next")),
				GoExpr.GoFuncLiteral([], [adaptedElementType == null ? "any" : adaptedElementType], nextBody)),
			GoStmt.GoReturn(GoExpr.GoIdent(iteratorMapName))
		]);
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, null, sourceExpr, true),
			GoStmt.GoVarDecl(wrappedName, "map[string]any", GoExpr.GoRaw("map[string]any{}"), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(wrappedName), GoExpr.GoStringLiteral("iterator")),
				GoExpr.GoFuncLiteral([], ["map[string]any"], iteratorFactoryBody)),
			GoStmt.GoReturn(GoExpr.GoIdent(wrappedName))
		]), []);
	}

	public function dynamicIterableSource(sourceExpr:TypedExpr):GoExpr {
		var sourcePlan = trySourcePlan(sourceExpr);
		return manualIteratorProtocolSource(sourcePlan == null ? lowerExpr(sourceExpr).expr : sourcePlan.sourceExpr, sourcePlan);
	}

	function dynamicIterableValue(sourceExpr:GoExpr, sourceType:Type):GoExpr {
		var sourcePlan = tryValuePlan(sourceExpr, sourceType);
		return manualIteratorProtocolSource(sourcePlan == null ? sourceExpr : sourcePlan.sourceExpr, sourcePlan);
	}

	function classInstanceFieldType(classType:ClassType, params:Array<Type>, fieldName:String):Null<Type> {
		for (field in classType.fields.get()) {
			if (field.name == fieldName) {
				return TypeTools.applyTypeParameters(field.type, classType.params, params);
			}
		}
		if (classType.superClass != null) {
			var superParams = [
				for (param in classType.superClass.params)
					TypeTools.applyTypeParameters(param, classType.params, params)
			];
			return classInstanceFieldType(classType.superClass.t.get(), superParams, fieldName);
		}
		return null;
	}

	function instanceFieldType(type:Type, fieldName:String):Null<Type> {
		return switch (Context.follow(type)) {
			case TInst(classRef, params):
				classInstanceFieldType(classRef.get(), params, fieldName);
			case TAnonymous(anonymousRef):
				var found:Null<Type> = null;
				for (field in anonymousRef.get().fields) {
					if (field.name == fieldName) {
						found = field.type;
						break;
					}
				}
				found;
			case _:
				null;
		};
	}

	function functionReturnType(type:Type):Null<Type> {
		return switch (Context.follow(type)) {
			case TFun(_, returnType): returnType;
			case _: null;
		};
	}

	function iterableElementType(type:Type):Null<Type> {
		var arrayElement = arrayElementType(type);
		if (arrayElement != null) {
			return arrayElement;
		}
		var listElement = haxeDsListElementType(type);
		if (listElement != null) {
			return listElement;
		}
		var iteratorField = instanceFieldType(type, "iterator");
		if (iteratorField == null) {
			return null;
		}
		var iteratorType = functionReturnType(iteratorField);
		if (iteratorType == null) {
			return null;
		}
		var nextField = instanceFieldType(iteratorType, "next");
		return nextField == null ? null : functionReturnType(nextField);
	}

	public function dynamicNestedIterableSource(sourceExpr:TypedExpr):GoExpr {
		var nestedType = iterableElementType(sourceExpr.t);
		if (nestedType == null || scalarGoType(nestedType) == "any" || scalarGoType(nestedType) == "map[string]any") {
			Context.fatalError("Lambda.flatten requires a statically concrete nested iterable carrier on Go", sourceExpr.pos);
		}
		var sourcePlan = trySourcePlan(sourceExpr);
		var loweredSource = sourcePlan == null ? lowerExpr(sourceExpr).expr : sourcePlan.sourceExpr;
		return manualIteratorProtocolSource(loweredSource, sourcePlan, function(rawElement:GoExpr):GoExpr {
			var boxedElement = GoExpr.GoCall(GoExpr.GoIdent("any"), [rawElement]);
			var nestedElement = lowerNullableAwareTypeAssertExpr(boxedElement, nestedType);
			var dynamicNestedElement = dynamicIterableValue(nestedElement, nestedType);
			return GoExpr.GoCall(GoExpr.GoIdent("Lambda_goIterableCarrierAdapter"), [dynamicNestedElement]);
		}, "interface{iterator() map[string]any}");
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

	public function indexedMapperAnyAdapter(mapperExpr:GoExpr, mapperType:Type):GoExpr {
		var rawIndexName = freshTempName("hx_lambda_index");
		var rawValueName = freshTempName("hx_lambda_value");
		var adaptedValueExpr:GoExpr = GoExpr.GoIdent(rawValueName);
		switch (Context.follow(mapperType)) {
			case TFun(args, _):
				if (args.length > 1) {
					adaptedValueExpr = lowerNullableAwareTypeAssertExpr(adaptedValueExpr, args[1].t);
				}
			case _:
		}
		return GoExpr.GoFuncLiteral([{name: rawIndexName, typeName: "int"}, {name: rawValueName, typeName: "any"}], ["any"], [
			GoStmt.GoReturn(GoExpr.GoCall(mapperExpr, [GoExpr.GoIdent(rawIndexName), adaptedValueExpr]))
		]);
	}

	public function iterableMapperAnyAdapter(mapperExpr:GoExpr, mapperType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var returnType:Null<Type> = null;
		switch (Context.follow(mapperType)) {
			case TFun(args, mappedType):
				if (args.length > 0) {
					adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, args[0].t);
				}
				returnType = mappedType;
			case _:
		}
		if (returnType == null || scalarGoType(returnType) == "any" || scalarGoType(returnType) == "map[string]any") {
			Context.fatalError("Lambda.flatMap requires a callback with a statically concrete iterable result on Go", Context.currentPos());
		}
		var mappedExpr = GoExpr.GoCall(mapperExpr, [adaptedArgExpr]);
		var dynamicMappedExpr = dynamicIterableValue(mappedExpr, returnType);
		var carrierExpr = GoExpr.GoCall(GoExpr.GoIdent("Lambda_goIterableCarrierAdapter"), [dynamicMappedExpr]);
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], ["interface{iterator() map[string]any}"], [GoStmt.GoReturn(carrierExpr)]);
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

	public function indexedFolderAnyAdapter(folderExpr:GoExpr, folderType:Type):GoExpr {
		var rawValueName = freshTempName("hx_lambda_value");
		var rawAccName = freshTempName("hx_lambda_acc");
		var rawIndexName = freshTempName("hx_lambda_index");
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
		return GoExpr.GoFuncLiteral([
			{name: rawValueName, typeName: "any"},
			{name: rawAccName, typeName: "any"},
			{name: rawIndexName, typeName: "int"}
		], ["any"], [
			GoStmt.GoReturn(GoExpr.GoCall(folderExpr, [adaptedValueExpr, adaptedAccExpr, GoExpr.GoIdent(rawIndexName)]))
		]);
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
