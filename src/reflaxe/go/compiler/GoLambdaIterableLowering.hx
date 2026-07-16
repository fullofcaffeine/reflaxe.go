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

/**
	What: The target element result recovered from Haxe's anonymous `Iterator<T>` type.
	Why: The Go carrier is erased, but each generated `next` closure must retain its
	concrete function result type.
	How: `structuralIteratorShape` admits only the closed `hasNext` / `next` protocol.
**/
private typedef GoStructuralIteratorShape = {
	final nextReturnType:Type;
}

/**
	What: One generated class method viewed through both declared and applied typing.
	Why: Generic Go methods are emitted with erased declared results even when Haxe
	has substituted a concrete result at the assignment site.
	How: Keep the original field for selector metadata and both signatures for safe
	closure construction.
**/
private typedef GoConcreteIteratorMethod = {
	final field:ClassField;
	final declaredType:Type;
	final appliedType:Type;
}

private typedef GoLambdaIterableLoweringConfig = {
	final lowerExpr:TypedExpr->GoLambdaLoweredExpr;
	final freshTempName:String->String;
	final isArrayType:Type->Bool;
	final arrayElementType:Type->Null<Type>;
	final arrayElementGoType:Type->String;
	final haxeDsListElementType:Type->Null<Type>;
	final scalarGoType:Type->String;
	final functionResultGoType:Type->String;
	final lowerNullableAwareTypeAssertExpr:(GoExpr, Type) -> GoExpr;
	final interfaceFieldName:(ClassType, ClassField) -> String;
	final noteSourceOwnedStdlibUsage:ClassType->Void;
	final localVarName:TVar->String;
	final lookupLocalLambdaAlias:String->Null<String>;
}

/**
	What:
	Compiler-owned Lambda/Iterable lowering policy for Go.

	Why:
	`Lambda` itself remains source-owned stdlib code, but generic `Iterable<T>`
	call sites and structural `Iterator<T>` assignments need backend representation
	glue: arrays are Go slices, `List` uses the current staged list carrier,
	concrete iterator classes need typed method adapters, and unknown iterables use
	the manual `iterator()` protocol. Keeping that policy in one helper keeps
	`GoCompiler` focused on orchestration while still making the
	representation-sensitive part explicit.

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
	final functionResultGoType:Type->String;
	final lowerNullableAwareTypeAssertExpr:(GoExpr, Type) -> GoExpr;
	final interfaceFieldName:(ClassType, ClassField) -> String;
	final noteSourceOwnedStdlibUsage:ClassType->Void;
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
		functionResultGoType = config.functionResultGoType;
		lowerNullableAwareTypeAssertExpr = config.lowerNullableAwareTypeAssertExpr;
		interfaceFieldName = config.interfaceFieldName;
		noteSourceOwnedStdlibUsage = config.noteSourceOwnedStdlibUsage;
		localVarName = config.localVarName;
		this.lookupLocalLambdaAlias = config.lookupLocalLambdaAlias;
	}

	/** What: Returns the result of a zero-argument function type; rejects every other shape. **/
	function zeroArgReturnType(type:Type):Null<Type> {
		return switch (Context.follow(type)) {
			case TFun(args, returnType) if (args.length == 0): returnType;
			case _: null;
		};
	}

	/**
		What: Recognizes the canonical anonymous Haxe iterator protocol.
		Why: Other structural objects may carry additional fields that a two-closure
		adapter must not silently discard.
		How: Require exactly `hasNext():Bool` and `next():T`, preserving `T` for Go.
	**/
	function structuralIteratorShape(type:Type):Null<GoStructuralIteratorShape> {
		return switch (Context.follow(type)) {
			case TAnonymous(anonymousRef):
				var fields = anonymousRef.get().fields;
				if (fields.length != 2) {
					null;
				} else {
					var hasNextType:Null<Type> = null;
					var nextType:Null<Type> = null;
					for (field in fields) {
						switch (field.name) {
							case "hasNext": hasNextType = field.type;
							case "next": nextType = field.type;
							case _:
						}
					}
					var hasNextReturn = hasNextType == null ? null : zeroArgReturnType(hasNextType);
					var nextReturn = nextType == null ? null : zeroArgReturnType(nextType);
					if (hasNextReturn == null || scalarGoType(hasNextReturn) != "bool" || nextReturn == null) {
						null;
					} else {
						{nextReturnType: nextReturn};
					}
				}
			case _: null;
		};
	}

	/**
		What: Finds one concrete iterator method across an instantiated superclass chain.
		Why: Structural conformance may come from inherited generic methods, while the
		emitted ABI still uses the declaring class's erased signature.
		How: Apply type parameters for semantic validation but retain the declared field.
	**/
	function concreteIteratorMethod(classType:ClassType, params:Array<Type>, name:String):Null<GoConcreteIteratorMethod> {
		for (field in classType.fields.get()) {
			if (field.name != name) {
				continue;
			}
			switch (field.kind) {
				case FMethod(_):
					return {
						field: field,
						declaredType: field.type,
						appliedType: TypeTools.applyTypeParameters(field.type, classType.params, params)
					};
				case _:
			}
		}

		if (classType.superClass == null) {
			return null;
		}
		var superClass = classType.superClass.t.get();
		var superParams = [
			for (param in classType.superClass.params)
				TypeTools.applyTypeParameters(param, classType.params, params)
		];
		return concreteIteratorMethod(superClass, superParams, name);
	}

	/** What: Removes compile-time-only wrappers without changing runtime evaluation. **/
	function unwrapStructuralSourceExpr(expr:TypedExpr):TypedExpr {
		return switch (expr.expr) {
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _): unwrapStructuralSourceExpr(inner);
			case _: expr;
		};
	}

	/**
		What: Resolves the single pure alias introduced by inline `Array.iterator()`.
		Why: Accepting arbitrary block prefixes could discard observable effects.
		How: Permit zero setup expressions or exactly one matching initialized local;
		reject every larger or unrelated block for the prefix-aware follow-up bead.
	**/
	function resolveInlineBlockAlias(source:TypedExpr, setup:Array<TypedExpr>):Null<TypedExpr> {
		var current = unwrapStructuralSourceExpr(source);
		if (setup.length == 0) {
			return current;
		}
		if (setup.length != 1) {
			return null;
		}
		return switch ([current.expr, setup[0].expr]) {
			case [TLocal(variable), TVar(candidate, value)] if (candidate.id == variable.id && value != null):
				unwrapStructuralSourceExpr(value);
			case _: null;
		};
	}

	/**
		What: Recovers the typed array behind direct and safely inlined Array iterators.
		Why: The array must be captured before its generic iterator class erases `T`.
		How: Match only the standard constructor/method identities and the safe alias above.
	**/
	function nativeArrayCursorSource(expr:TypedExpr):Null<TypedExpr> {
		var source = unwrapStructuralSourceExpr(expr);
		return switch (source.expr) {
			case TBlock(exprs) if (exprs.length > 0):
				var setup = exprs.slice(0, exprs.length - 1);
				var tailSource = nativeArrayCursorSource(exprs[exprs.length - 1]);
				tailSource == null ? null : resolveInlineBlockAlias(tailSource, setup);
			case TNew(classRef, _, args): var classType = classRef.get(); classType.pack.join(".") == "haxe.iterators" && classType.name == "ArrayIterator" && args.length == 1 ? args[0] : null;
			case TCall(callee, args) if (args.length == 0):
				switch (unwrapStructuralSourceExpr(callee).expr) {
					case TField(target, FInstance(classRef, _, fieldRef)): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "Array" && fieldRef.get()
							.name == "iterator" ? target : null;
					case _: null;
				}
			case _: null;
		};
	}

	/**
		What: Builds a structural iterator map over one live typed Go slice.
		Why: Converting the slice to `[]any` would copy its slots and hide later mutation.
		How: Capture the slice header and cursor once, then emit typed closures through
		the existing structural carrier.
	**/
	function nativeArrayCursorMap(sourcePlan:GoLambdaSourcePlan, targetResultType:Type):GoExpr {
		var sourceName = freshTempName("hx_structural_array");
		var indexName = freshTempName("hx_structural_array_index");
		var mapName = freshTempName("hx_structural_iterator_map");
		var valueName = freshTempName("hx_structural_array_value");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var indexExpr = GoExpr.GoIdent(indexName);
		var targetResultGoType = functionResultGoType(targetResultType);
		var nextValue:GoExpr = GoExpr.GoIdent(valueName);
		if (targetResultGoType == "any") {
			nextValue = GoExpr.GoCall(GoExpr.GoIdent("any"), [nextValue]);
		} else if (sourcePlan.elementType != targetResultGoType) {
			nextValue = lowerNullableAwareTypeAssertExpr(GoExpr.GoCall(GoExpr.GoIdent("any"), [nextValue]), targetResultType);
		}

		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, sourcePlan.sourceType, sourcePlan.sourceExpr, true),
			GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
			GoStmt.GoVarDecl(mapName, "map[string]any", emptyDynamicFieldMapExpr(), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("hasNext")), GoExpr.GoFuncLiteral([], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoBinary("<", indexExpr, GoExpr.GoCall(GoExpr.GoIdent("len"), [sourceExpr])))
			])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("next")), GoExpr.GoFuncLiteral([], [targetResultGoType], [
				GoStmt.GoVarDecl(valueName, null, GoExpr.GoIndex(sourceExpr, indexExpr), true),
				GoStmt.GoAssign(indexExpr, GoExpr.GoBinary("+", indexExpr, GoExpr.GoIntLiteral(1))),
				GoStmt.GoReturn(nextValue)
			])),
			GoStmt.GoReturn(GoExpr.GoIdent(mapName))
		]), []);
	}

	/**
		What: Lowers a direct native-array iterator expression to the structural map.
		Why: Lowering the erased `ArrayIterator<T>` class first would require copying
		a typed Go slice to `[]any`, which would hide later indexed mutations from the
		iterator and violate Haxe's shared-array behavior.
		How: Recover the typed array before ordinary expression lowering, then capture
		the original slice and one cursor in typed `hasNext` / `next` closures.
	**/
	public function nativeArrayStructuralIteratorCoerce(sourceExpr:TypedExpr, toType:Type):Null<GoExpr> {
		var targetShape = structuralIteratorShape(toType);
		if (targetShape == null) {
			return null;
		}
		var arraySource = nativeArrayCursorSource(sourceExpr);
		if (arraySource == null) {
			return null;
		}
		var sourcePlan = trySourcePlan(arraySource);
		return sourcePlan != null && sourcePlan.domain == "array" ? nativeArrayCursorMap(sourcePlan, targetShape.nextReturnType) : null;
	}

	/**
		What:
		- Adapts one concrete generated iterator class to Haxe's structural
		  `Iterator<T>` carrier.

		Why:
		- Haxe accepts classes with typed `hasNext()` and `next()` methods as
		  iterators, but generated classes are pointers while anonymous structural
		  values are `map[string]any` on Go. Generic generated methods can also return
		  erased `any`, even when the assigned `Iterator<T>` has a concrete element
		  type.

		How:
		- Recognize only the canonical two-method structural shape, evaluate the
		  concrete source once, and expose typed closures through the existing map
		  carrier. Calls go through `__hx_this` so overrides remain virtual, and an
		  erased `next()` result is asserted back to the target element type.
	**/
	public function structuralIteratorCoerce(expr:GoExpr, fromType:Type, toType:Type):Null<GoExpr> {
		var targetShape = structuralIteratorShape(toType);
		if (targetShape == null) {
			return null;
		}

		var sourceClass:Null<ClassType> = null;
		var sourceParams:Array<Type> = [];
		switch (Context.follow(fromType)) {
			case TInst(classRef, params):
				sourceClass = classRef.get();
				sourceParams = params;
			case _:
		}
		if (sourceClass == null || sourceClass.isExtern || sourceClass.isInterface) {
			return null;
		}

		var hasNextMethod = concreteIteratorMethod(sourceClass, sourceParams, "hasNext");
		var nextMethod = concreteIteratorMethod(sourceClass, sourceParams, "next");
		if (hasNextMethod == null || nextMethod == null) {
			return null;
		}
		var appliedHasNextReturn = zeroArgReturnType(hasNextMethod.appliedType);
		var appliedNextReturn = zeroArgReturnType(nextMethod.appliedType);
		var declaredNextReturn = zeroArgReturnType(nextMethod.declaredType);
		if (appliedHasNextReturn == null
			|| scalarGoType(appliedHasNextReturn) != "bool"
			|| appliedNextReturn == null
			|| declaredNextReturn == null) {
			return null;
		}

		noteSourceOwnedStdlibUsage(sourceClass);
		var sourceName = freshTempName("hx_structural_iterator");
		var mapName = freshTempName("hx_structural_iterator_map");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var dispatchReceiver = GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "__hx_this");
		var hasNextCall = GoExpr.GoCall(GoExpr.GoSelector(dispatchReceiver, interfaceFieldName(sourceClass, hasNextMethod.field)), []);
		var nextCall = GoExpr.GoCall(GoExpr.GoSelector(dispatchReceiver, interfaceFieldName(sourceClass, nextMethod.field)), []);
		var targetResultGoType = functionResultGoType(targetShape.nextReturnType);
		var declaredResultGoType = functionResultGoType(declaredNextReturn);
		var adaptedNext:GoExpr = nextCall;
		if (targetResultGoType == "any") {
			adaptedNext = GoExpr.GoCall(GoExpr.GoIdent("any"), [nextCall]);
		} else if (declaredResultGoType != targetResultGoType) {
			adaptedNext = lowerNullableAwareTypeAssertExpr(GoExpr.GoCall(GoExpr.GoIdent("any"), [nextCall]), targetShape.nextReturnType);
		}

		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: scalarGoType(fromType)}], ["map[string]any"], [
			GoStmt.GoVarDecl(mapName, "map[string]any", emptyDynamicFieldMapExpr(), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("hasNext")),
				GoExpr.GoFuncLiteral([], ["bool"], [GoStmt.GoReturn(hasNextCall)])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("next")),
				GoExpr.GoFuncLiteral([], [targetResultGoType], [GoStmt.GoReturn(adaptedNext)])),
			GoStmt.GoReturn(GoExpr.GoIdent(mapName))
		]), [expr]);
	}

	public function trySourcePlan(sourceExpr:TypedExpr):Null<GoLambdaSourcePlan> {
		if (!isArrayType(sourceExpr.t) && haxeDsListElementType(sourceExpr.t) == null) {
			return null;
		}
		return tryValuePlan(lowerExpr(sourceExpr).expr, sourceExpr.t);
	}

	/**
		What: Builds the empty map used by erased structural object carriers.
		Why: The typed Go AST does not yet model map composite literals; duplicating
		`GoRaw` at every adapter site would grow unowned compiler debt.
		How: Keep the one syntax-only fragment here until `haxe_go-vfp.8.3` adds the
		typed composite-literal node, while all keys, closures, and calls stay typed.
	**/
	function emptyDynamicFieldMapExpr():GoExpr {
		return GoExpr.GoRaw("map[string]any{}");
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
			GoStmt.GoVarDecl(iteratorMapName, "map[string]any", emptyDynamicFieldMapExpr(), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(iteratorMapName), GoExpr.GoStringLiteral("hasNext")),
				GoExpr.GoFuncLiteral([], ["bool"], [GoStmt.GoReturn(hasNextExpr)])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(iteratorMapName), GoExpr.GoStringLiteral("next")),
				GoExpr.GoFuncLiteral([], [adaptedElementType == null ? "any" : adaptedElementType], nextBody)),
			GoStmt.GoReturn(GoExpr.GoIdent(iteratorMapName))
		]);
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, null, sourceExpr, true),
			GoStmt.GoVarDecl(wrappedName, "map[string]any", emptyDynamicFieldMapExpr(), true),
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
