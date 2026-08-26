package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypeTools;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoType;

typedef GoLambdaLoweredExpr = {
	final expr:GoExpr;
	final isStringLike:Bool;
}

/**
	What: One lowered expression plus statements that must execute before it.
	Why: Inline iterator methods can perform observable work before producing the
	concrete iterator value that needs structural adaptation.
	How: The iterable owner returns the ordered prefix to `GoCompiler`, which emits
	it directly or materializes it inside an expression context.
**/
typedef GoLambdaLoweredExprWithPrefix = {
	final prefix:Array<GoStmt>;
	final expr:GoExpr;
	final isStringLike:Bool;
}

typedef GoLambdaSourcePlan = {
	final domain:String;
	final elementType:String;
	final sourceExpr:GoExpr;
	final sourceType:String;
	final sharedArray:Bool;
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

/**
	What: The validated class and method authority for one structural iterator adapter.
	Why: Inline-tail discovery must prove an adapter exists before lowering retained
	setup statements, while ordinary coercion needs the same proof to build closures.
	How: Resolve the exact anonymous target shape and concrete source methods once.
**/
private typedef GoConcreteStructuralIteratorPlan = {
	final targetShape:GoStructuralIteratorShape;
	final sourceClass:ClassType;
	final hasNextMethod:GoConcreteIteratorMethod;
	final nextMethod:GoConcreteIteratorMethod;
	final declaredNextReturn:Type;
}

/**
	What: One concrete terminal iterator plus expressions that precede it in an inline block.
	Why: The enclosing block is typed as anonymous `Iterator<T>`, but its terminal
	expression retains the concrete class needed by the Go adapter.
	How: Keep the terminal typed expression intact and lower the ordered setup later.
**/
private typedef GoInlineConcreteIteratorTailPlan = {
	final setup:Array<TypedExpr>;
	final sourceExpr:TypedExpr;
}

/**
	What: The typed array source and ordered setup expressions recovered from an
	inlined Array iterator.
	Why: Replacing the final erased iterator constructor must not discard effects
	that appeared earlier in the inline block.
	How: `nativeArrayCursorPlan` separates the terminal array from its prefix while
	retaining every non-alias setup expression.
**/
private typedef GoNativeArrayCursorPlan = {
	final setup:Array<TypedExpr>;
	final sourceExpr:TypedExpr;
}

private typedef GoLambdaIterableLoweringConfig = {
	final lowerExpr:TypedExpr->GoLambdaLoweredExpr;
	final lowerToStatements:TypedExpr->Array<GoStmt>;
	final freshTempName:String->String;
	final isArrayType:Type->Bool;
	final isHaxeArrayType:Type->Bool;
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
	glue: portable Arrays use the shared `hxrt.Array` identity, explicit/fixed
	array-like views may remain native Go slices, `List` uses the current staged
	list carrier, concrete iterator classes need typed method adapters, and unknown
	iterables use the manual `iterator()` protocol. Keeping that policy in one helper keeps
	`GoCompiler` focused on orchestration while still making the
	representation-sensitive part explicit.

	How:
	Consumes typed callbacks from `GoCompiler` for lowering, naming, and type
	mapping, then builds the small adapter expressions and ordered setup statements
	used by iterable lowering. This module should not grow general Haxe stdlib
	behavior; it owns only the bridge between Haxe iterable semantics and current
	Go carriers.
	Go `any` stays localized here because unknown `Iterable<T>` values expose
	their elements through the manual iterator protocol instead of statically known
	shared-Array or native-slice storage.
**/
class GoLambdaIterableLowering {
	final lowerExpr:TypedExpr->GoLambdaLoweredExpr;
	final lowerToStatements:TypedExpr->Array<GoStmt>;
	final freshTempName:String->String;
	final isArrayType:Type->Bool;
	final isHaxeArrayType:Type->Bool;
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
		lowerToStatements = config.lowerToStatements;
		freshTempName = config.freshTempName;
		isArrayType = config.isArrayType;
		isHaxeArrayType = config.isHaxeArrayType;
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

	/**
		What: Validates a concrete generated class against the closed iterator target.
		Why: Both ordinary values and recovered inline tails must make the same typed
		decision before emitting a structural carrier.
		How: Require a non-extern class plus zero-argument `hasNext():Bool` and `next()`
		methods, retaining both applied validation and the declared emitted result ABI.
	**/
	function concreteStructuralIteratorPlan(fromType:Type, toType:Type):Null<GoConcreteStructuralIteratorPlan> {
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

		return {
			targetShape: targetShape,
			sourceClass: sourceClass,
			hasNextMethod: hasNextMethod,
			nextMethod: nextMethod,
			declaredNextReturn: declaredNextReturn
		};
	}

	/** What: Removes compile-time-only wrappers without changing runtime evaluation. **/
	function unwrapStructuralSourceExpr(expr:TypedExpr):TypedExpr {
		return switch (expr.expr) {
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _): unwrapStructuralSourceExpr(inner);
			case _: expr;
		};
	}

	/**
		What: Folds trailing local aliases into an array cursor plan.
		Why: Haxe's inline `Array.iterator()` introduces compile-only aliases that do
		not need generated locals, while earlier effects must remain in the prefix.
		How: Remove only a final `var alias = value` referenced by the terminal source;
		recurse across consecutive trailing aliases and preserve every other statement.
	**/
	function foldTrailingArrayAliases(sourceExpr:TypedExpr, setup:Array<TypedExpr>):GoNativeArrayCursorPlan {
		var current = unwrapStructuralSourceExpr(sourceExpr);
		if (setup.length == 0) {
			return {setup: setup, sourceExpr: current};
		}
		var last = setup[setup.length - 1];
		return switch ([current.expr, last.expr]) {
			case [TLocal(variable), TVar(candidate, value)] if (candidate.id == variable.id && value != null):
				foldTrailingArrayAliases(value, setup.slice(0, setup.length - 1));
			case _:
				{setup: setup, sourceExpr: current};
		};
	}

	/**
		What: Recovers the typed array and ordered prefix behind an Array iterator.
		Why: The array must be captured before its generic iterator class erases `T`,
		without dropping effects introduced by an inline caller.
		How: Match only the standard constructor/method identities, recursively collect
		block setup expressions, and fold only safe trailing aliases.
	**/
	function nativeArrayCursorPlan(expr:TypedExpr):Null<GoNativeArrayCursorPlan> {
		var source = unwrapStructuralSourceExpr(expr);
		return switch (source.expr) {
			case TBlock(exprs) if (exprs.length > 0):
				var setup = exprs.slice(0, exprs.length - 1);
				var tailPlan = nativeArrayCursorPlan(exprs[exprs.length - 1]);
				tailPlan == null ? null : foldTrailingArrayAliases(tailPlan.sourceExpr, setup.concat(tailPlan.setup));
			case TNew(classRef, _, args): var classType = classRef.get(); classType.pack.join(".") == "haxe.iterators" && classType.name == "ArrayIterator" && args.length == 1 ? {
					setup: [],
					sourceExpr: args[0]
				} : null;
			case TCall(callee, args) if (args.length == 0):
				switch (unwrapStructuralSourceExpr(callee).expr) {
					case TField(target, FInstance(classRef, _, fieldRef)): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "Array" && fieldRef.get()
							.name == "iterator" ? {
								setup: [],
								sourceExpr: target
							} : null;
					case _: null;
				}
			case _: null;
		};
	}

	/**
		What: Builds a structural iterator map over one live array storage value.
		Why: Portable Arrays must retain their shared carrier, while copying a native
		slice to `[]any` would hide later indexed mutation.
		How: Capture the shared carrier or slice header and cursor once, then emit typed
		closures through the existing structural iterator carrier.
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
		} else if (sourcePlan.sharedArray || sourcePlan.elementType != targetResultGoType) {
			nextValue = lowerNullableAwareTypeAssertExpr(GoExpr.GoCall(GoExpr.GoIdent("any"), [nextValue]), targetResultType);
		}
		var lengthExpr = sourcePlan.sharedArray ? GoExpr.GoCall(GoExpr.GoSelector(sourceExpr, "Len"), []) : GoExpr.GoCall(GoExpr.GoIdent("len"), [sourceExpr]);
		var indexedValue = sourcePlan.sharedArray ? GoExpr.GoCall(GoExpr.GoSelector(sourceExpr, "Get"), [indexExpr]) : GoExpr.GoIndex(sourceExpr, indexExpr);

		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, sourcePlan.sourceType, sourcePlan.sourceExpr, true),
			GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
			GoStmt.GoVarDecl(mapName, "map[string]any", emptyDynamicFieldMapExpr(), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("hasNext")),
				GoExpr.GoFuncLiteral([], ["bool"], [GoStmt.GoReturn(GoExpr.GoBinary("<", indexExpr, lengthExpr))])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("next")), GoExpr.GoFuncLiteral([], [targetResultGoType], [
				GoStmt.GoVarDecl(valueName, null, indexedValue, true),
				GoStmt.GoAssign(indexExpr, GoExpr.GoBinary("+", indexExpr, GoExpr.GoIntLiteral(1))),
				GoStmt.GoReturn(nextValue)
			])),
			GoStmt.GoReturn(GoExpr.GoIdent(mapName))
		]), []);
	}

	/**
		What: Lowers an Array iterator expression and its ordered setup prefix
		to the structural map.
		Why: Lowering the erased `ArrayIterator<T>` class first would either discard the
		portable shared carrier or copy native slice storage to `[]any`, hiding later
		mutations from the iterator; replacing an inline block's final constructor alone
		would also discard earlier effects.
		How: Recover the typed array before ordinary expression lowering, lower every
		retained setup expression to statements, and return that prefix with one live
		carrier/slice cursor for the caller to emit or materialize.
	**/
	public function nativeArrayStructuralIteratorCoerce(sourceExpr:TypedExpr, toType:Type):Null<GoLambdaLoweredExprWithPrefix> {
		var targetShape = structuralIteratorShape(toType);
		if (targetShape == null) {
			return null;
		}
		var cursorPlan = nativeArrayCursorPlan(sourceExpr);
		if (cursorPlan == null) {
			return null;
		}
		var sourcePlan = trySourcePlan(cursorPlan.sourceExpr);
		if (sourcePlan == null || sourcePlan.domain != "array") {
			return null;
		}
		var prefix = new Array<GoStmt>();
		for (setupExpr in cursorPlan.setup) {
			prefix = prefix.concat(lowerToStatements(setupExpr));
		}
		return {
			prefix: prefix,
			expr: nativeArrayCursorMap(sourcePlan, targetShape.nextReturnType),
			isStringLike: false
		};
	}

	/**
		What: Recovers the concrete terminal iterator from a nested inline block.
		Why: Haxe assigns the block its declared anonymous `Iterator<T>` type, hiding
		the terminal class from ordinary post-lowering coercion.
		How: Retain every preceding expression in order and admit the terminal value
		only when the shared concrete structural plan validates it.
	**/
	function inlineConcreteIteratorTailPlan(expr:TypedExpr, toType:Type):Null<GoInlineConcreteIteratorTailPlan> {
		var source = unwrapStructuralSourceExpr(expr);
		return switch (source.expr) {
			case TBlock(exprs) if (exprs.length > 0):
				var tailPlan = inlineConcreteIteratorTailPlan(exprs[exprs.length - 1], toType);
				tailPlan == null ? null : {
					setup: exprs.slice(0, exprs.length - 1).concat(tailPlan.setup),
					sourceExpr: tailPlan.sourceExpr
				};
			case _:
				concreteStructuralIteratorPlan(source.t, toType) == null ? null : {
					setup: [],
					sourceExpr: source
				};
		};
	}

	/**
		What: Adapts an effectful inline block ending in a concrete generated iterator.
		Why: Lowering the enclosing anonymous block first loses the class authority
		required by `structuralIteratorCoerce` and leaves an invalid Go pointer-to-map
		assignment.
		How: Keep Array on its dedicated shared-carrier/native-slice path, lower the
		retained setup once, then pass the typed terminal value to the existing
		class-agnostic adapter.
	**/
	public function inlineConcreteStructuralIteratorCoerce(sourceExpr:TypedExpr, toType:Type):Null<GoLambdaLoweredExprWithPrefix> {
		var source = unwrapStructuralSourceExpr(sourceExpr);
		switch (source.expr) {
			case TBlock(_):
			case _:
				return null;
		}
		var tailPlan = inlineConcreteIteratorTailPlan(source, toType);
		if (tailPlan == null) {
			return null;
		}
		var loweredTail = lowerExpr(tailPlan.sourceExpr);
		var adaptedTail = structuralIteratorCoerce(loweredTail.expr, tailPlan.sourceExpr.t, toType);
		if (adaptedTail == null) {
			return null;
		}
		var prefix = new Array<GoStmt>();
		for (setupExpr in tailPlan.setup) {
			prefix = prefix.concat(lowerToStatements(setupExpr));
		}
		return {
			prefix: prefix,
			expr: adaptedTail,
			isStringLike: false
		};
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
		var plan = concreteStructuralIteratorPlan(fromType, toType);
		if (plan == null) {
			return null;
		}
		var targetShape = plan.targetShape;
		var sourceClass = plan.sourceClass;
		var hasNextMethod = plan.hasNextMethod;
		var nextMethod = plan.nextMethod;
		noteSourceOwnedStdlibUsage(sourceClass);
		var sourceName = freshTempName("hx_structural_iterator");
		var mapName = freshTempName("hx_structural_iterator_map");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var dispatchReceiver = GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "__hx_this");
		var hasNextCall = GoExpr.GoCall(GoExpr.GoSelector(dispatchReceiver, interfaceFieldName(sourceClass, hasNextMethod.field)), []);
		var nextCall = GoExpr.GoCall(GoExpr.GoSelector(dispatchReceiver, interfaceFieldName(sourceClass, nextMethod.field)), []);
		var targetResultGoType = functionResultGoType(targetShape.nextReturnType);
		var declaredResultGoType = functionResultGoType(plan.declaredNextReturn);
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

	/**
		What: Adapts an erased generic method result to its applied structural
		`Iterator<T>` closure ABI.

		Why: A generated generic method such as `StringMap<T>.iterator()` is emitted
		once, so its `next` field stores `func() any`. The applied Haxe call can expose
		`next():Concrete`; asserting the stored function itself to `func() Concrete`
		would panic because Go function types are invariant.

		How: Admit only the canonical two-field iterator shape and only when the
		declared result is erased while the applied result is concrete. Capture the
		raw iterator once, retain its `hasNext` closure, and wrap `next` with the
		existing typed result assertion.
	**/
	public function adaptErasedGenericStructuralIteratorResult(expr:GoExpr, declaredFunctionType:Type, appliedFunctionType:Type):Null<GoExpr> {
		var declaredReturn = functionReturnType(declaredFunctionType);
		var appliedReturn = functionReturnType(appliedFunctionType);
		if (declaredReturn == null || appliedReturn == null) {
			return null;
		}
		var declaredShape = structuralIteratorShape(declaredReturn);
		var appliedShape = structuralIteratorShape(appliedReturn);
		if (declaredShape == null || appliedShape == null) {
			return null;
		}

		var declaredResultType = functionResultGoType(declaredShape.nextReturnType);
		var appliedResultType = functionResultGoType(appliedShape.nextReturnType);
		if (declaredResultType != "any" || appliedResultType == "any") {
			return null;
		}

		var rawName = freshTempName("hx_erased_iterator");
		var hasNextName = freshTempName("hx_erased_iterator_has_next");
		var nextName = freshTempName("hx_erased_iterator_next");
		var mapName = freshTempName("hx_applied_iterator");
		var rawExpr = GoExpr.GoIdent(rawName);
		var nextCall = GoExpr.GoCall(GoExpr.GoIdent(nextName), []);
		var adaptedNext = lowerNullableAwareTypeAssertExpr(nextCall, appliedShape.nextReturnType);

		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: rawName, typeName: "map[string]any"}], ["map[string]any"], [
			GoStmt.GoVarDecl(hasNextName, "func() bool", GoExpr.GoTypeAssert(GoExpr.GoIndex(rawExpr, GoExpr.GoStringLiteral("hasNext")), "func() bool"), true),
			GoStmt.GoVarDecl(nextName, "func() any", GoExpr.GoTypeAssert(GoExpr.GoIndex(rawExpr, GoExpr.GoStringLiteral("next")), "func() any"), true),
			GoStmt.GoVarDecl(mapName, "map[string]any", emptyDynamicFieldMapExpr(), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("hasNext")),
				GoExpr.GoFuncLiteral([], ["bool"], [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(hasNextName), []))])),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(mapName), GoExpr.GoStringLiteral("next")),
				GoExpr.GoFuncLiteral([], [appliedResultType], [GoStmt.GoReturn(adaptedNext)])),
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
		Why: Iterator callbacks use an erased field map, but its initialization should
		remain visible to import analysis and transforms.
		How: Construct the empty value through the same typed composite-literal node
		used by ordinary anonymous-object lowering.
	**/
	function emptyDynamicFieldMapExpr():GoExpr {
		var mapType = GoType.map(GoType.builtin(GoBuiltinType.StringType), GoType.builtin(GoBuiltinType.AnyType));
		return GoExpr.GoCompositeLiteral(mapType, []);
	}

	function tryValuePlan(sourceExpr:GoExpr, sourceType:Type):Null<GoLambdaSourcePlan> {
		if (isArrayType(sourceType)) {
			var elementType = arrayElementGoType(sourceType);
			return {
				domain: "array",
				elementType: elementType,
				sourceExpr: sourceExpr,
				sourceType: isHaxeArrayType(sourceType) ? "*hxrt.Array" : "[]" + elementType,
				sharedArray: isHaxeArrayType(sourceType)
			};
		}

		var listElement = haxeDsListElementType(sourceType);
		if (listElement != null) {
			return {
				domain: "list",
				elementType: scalarGoType(listElement),
				sourceExpr: sourceExpr,
				sourceType: "*haxe__ds__List",
				sharedArray: false
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
			hasNextExpr = GoExpr.GoBinary("<", indexExpr,
				sourcePlan.sharedArray ? GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "Len"),
					[]) : GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent(sourceName)]));
			nextValueExpr = sourcePlan.sharedArray ? GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "Get"),
				[indexExpr]) : GoExpr.GoIndex(GoExpr.GoIdent(sourceName), indexExpr);
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
		if (isHaxeArrayType(targetArrayType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.ArrayFromValues"), [anySliceExpr]);
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
			GoStmt.GoVarDecl(outName, outType,
				GoExpr.GoMakeSlice(targetElementGoType, GoExpr.GoIntLiteral(0), GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent(rawName)])), true),
			GoStmt.GoRangeStmt(null, itemName, GoExpr.GoIdent(rawName), true, [
				GoStmt.GoAssign(GoExpr.GoIdent(outName), GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(outName), convertedItemExpr]))
			]),
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
