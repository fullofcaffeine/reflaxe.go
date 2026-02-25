package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.analyze.MetalLaneAnalyzer;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoTypeMapper;

private typedef LaneElementMethodCall = {
	final methodName:String;
	final elementType:Type;
}

private typedef LaneMapMethodCall = {
	final methodName:String;
	final keyType:Type;
	final valueType:Type;
}
#end

class MetalLaneEnforcer {
	#if macro
	static var initialized = false;

	public static function init():Void {
		if (initialized) {
			return;
		}
		initialized = true;

		if (!isGoBuild()) {
			return;
		}

		var buildContext = GoBuildContextResolver.resolve();
		if (buildContext.profile != GoProfile.Portable) {
			return;
		}

		Context.onAfterTyping(types -> enforce(types));
	}

	static function enforce(types:Array<ModuleType>):Void {
		var laneSnapshot = MetalLaneAnalyzer.collect(types);
		if (laneSnapshot.modules.length == 0) {
			return;
		}

		var laneModules:Map<String, Bool> = [];
		for (moduleName in laneSnapshot.modules) {
			laneModules.set(moduleName, true);
		}

		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!laneModules.exists(moduleNameForClass(classType))) {
						continue;
					}
					enforceClassFields(classType.fields.get(), moduleNameForClass(classType));
					enforceClassFields(classType.statics.get(), moduleNameForClass(classType));
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!laneModules.exists(moduleNameForAbstract(abstractType))) {
						continue;
					}
					if (abstractType.impl != null) {
						var impl = abstractType.impl.get();
						if (impl != null) {
							enforceClassFields(impl.fields.get(), moduleNameForAbstract(abstractType));
							enforceClassFields(impl.statics.get(), moduleNameForAbstract(abstractType));
						}
					}
				case _:
			}
		}
	}

	static function enforceClassFields(fields:Array<ClassField>, moduleName:String):Void {
		if (fields == null) {
			return;
		}
		for (field in fields) {
			var expr = field.expr();
			if (expr == null) {
				continue;
			}
			scanForLaneViolations(expr, moduleName);
		}
	}

	static function scanForLaneViolations(expr:TypedExpr, moduleName:String):Void {
		if (isGoInjectionCall(expr)) {
			Context.error("MetalLaneEnforcer: __go__ is not allowed in @:goMetal modules when contract=portable.", expr.pos);
		}
		var typedFallback = detectTypedFallbackViolation(expr);
		if (typedFallback != null) {
			Context.error('MetalLaneEnforcer: typed metal fallback is not allowed in @:goMetal module "'
				+ moduleName
				+ '" when contract=portable. '
				+ typedFallback
				+ " Use concrete generic types for go.Chan/go.Slice/go.Map/go.Result in lane modules (avoid Dynamic/Any).",
				expr.pos);
		}
		TypedExprTools.iter(expr, e -> scanForLaneViolations(e, moduleName));
	}

	static function isGoInjectionCall(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TCall(callTarget, _):
				switch (callTarget.expr) {
					case TIdent(name):
						name == "__go__";
					case TLocal(variable):
						variable.name == "__go__";
					case TField(_, fieldAccess):
						switch (fieldAccess) {
							case FInstance(_, _, classField) | FStatic(_, classField) | FAnon(classField) | FClosure(_, classField):
								classField.get().name == "__go__";
							case FEnum(_, enumField):
								enumField.name == "__go__";
							case FDynamic(name):
								name == "__go__";
						}
					case _:
						false;
				}
			case _:
				false;
		}
	}

	static function detectTypedFallbackViolation(expr:TypedExpr):Null<String> {
		return switch (expr.expr) {
			case TNew(classRef, _, _):
				var classType = classRef.get();
				if (isGoClass(classType, "Chan")) {
					var elementType = goChanElementType(expr.t);
					if (elementType == null || !isMonomorphizableMetalType(elementType)) {
						"Could not monomorphize go.Chan element type for constructor specialization.";
					} else {
						null;
					}
				} else {
					null;
				}
			case TCall(callee, _):
				detectTypedFallbackViolationFromCall(callee, expr.t);
			case _:
				null;
		};
	}

	static function detectTypedFallbackViolationFromCall(callee:TypedExpr, returnType:Type):Null<String> {
		if (isGoStaticCall(callee, "Go", "newChan")) {
			var elementType = goChanElementType(returnType);
			if (elementType == null || !isMonomorphizableMetalType(elementType)) {
				return "Could not monomorphize go.Go.newChan return type for metal specialization.";
			}
		}

		if (isGoStaticCall(callee, "Result", "ok") || isGoStaticCall(callee, "Go", "ok")) {
			var elementType = goResultElementType(returnType);
			if (elementType == null || !isMonomorphizableMetalType(elementType)) {
				return "Could not monomorphize go.Result<T>.ok return type for metal specialization.";
			}
		}

		if (isGoStaticCall(callee, "Result", "failure") || isGoStaticCall(callee, "Go", "fail")) {
			var elementType = goResultElementType(returnType);
			if (elementType == null || !isMonomorphizableMetalType(elementType)) {
				return "Could not monomorphize go.Result<T>.failure return type for metal specialization.";
			}
		}

		var chanMethod = asGoChanMethodCall(callee);
		if (chanMethod != null && !isMonomorphizableMetalType(chanMethod.elementType)) {
			return "Could not monomorphize go.Chan method call (element type resolves to any).";
		}

		var sliceMethod = asGoSliceMethodCall(callee);
		if (sliceMethod != null && !isMonomorphizableMetalType(sliceMethod.elementType)) {
			return "Could not monomorphize go.Slice element type for metal specialization.";
		}

		var mapMethod = asGoMapMethodCall(callee);
		if (mapMethod != null && (!isMonomorphizableMetalType(mapMethod.keyType) || !isMonomorphizableMetalType(mapMethod.valueType))) {
			return "Could not monomorphize go.Map key/value types for metal specialization.";
		}

		var resultMethod = asGoResultMethodCall(callee);
		if (resultMethod != null && !isMonomorphizableMetalType(resultMethod.elementType)) {
			return "Could not monomorphize go.Result<T> method receiver for metal specialization.";
		}

		return null;
	}

	static function isGoStaticCall(callee:TypedExpr, className:String, fieldName:String):Bool {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, fieldRef)): var classType = classRef.get(); classType.pack.join(".") == "go" && classType.name == className && fieldRef.get()
					.name == fieldName;
			case TMeta(_, inner):
				isGoStaticCall(inner, className, fieldName);
			case TParenthesis(inner):
				isGoStaticCall(inner, className, fieldName);
			case TCast(inner, _):
				isGoStaticCall(inner, className, fieldName);
			case _:
				false;
		};
	}

	static function asGoChanMethodCall(callee:TypedExpr):Null<LaneElementMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, fieldRef)):
				var classType = classRef.get();
				var elementType = goChanElementType(target.t);
				if (isGoClass(classType, "Chan") && elementType != null) {
					{
						methodName: fieldRef.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoChanMethodCall(inner);
			case TParenthesis(inner):
				asGoChanMethodCall(inner);
			case TCast(inner, _):
				asGoChanMethodCall(inner);
			case _:
				null;
		};
	}

	static function asGoSliceMethodCall(callee:TypedExpr):Null<LaneElementMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, fieldRef)):
				var classType = classRef.get();
				var elementType = goSliceElementType(target.t);
				if (isGoClass(classType, "Slice") && elementType != null) {
					{
						methodName: fieldRef.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoSliceMethodCall(inner);
			case TParenthesis(inner):
				asGoSliceMethodCall(inner);
			case TCast(inner, _):
				asGoSliceMethodCall(inner);
			case _:
				null;
		};
	}

	static function asGoMapMethodCall(callee:TypedExpr):Null<LaneMapMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, fieldRef)):
				var classType = classRef.get();
				var pair = goMapTypePair(target.t);
				if (isGoClass(classType, "Map") && pair != null) {
					{
						methodName: fieldRef.get().name,
						keyType: pair.keyType,
						valueType: pair.valueType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoMapMethodCall(inner);
			case TParenthesis(inner):
				asGoMapMethodCall(inner);
			case TCast(inner, _):
				asGoMapMethodCall(inner);
			case _:
				null;
		};
	}

	static function asGoResultMethodCall(callee:TypedExpr):Null<LaneElementMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, fieldRef)):
				var classType = classRef.get();
				var elementType = goResultElementType(target.t);
				if (isGoClass(classType, "Result") && elementType != null) {
					{
						methodName: fieldRef.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoResultMethodCall(inner);
			case TParenthesis(inner):
				asGoResultMethodCall(inner);
			case TCast(inner, _):
				asGoResultMethodCall(inner);
			case _:
				null;
		};
	}

	static function goChanElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoClass(classType, "Chan") && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goChanElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goChanElementType(resolved);
			case _:
				null;
		};
	}

	static function goSliceElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoClass(classType, "Slice") && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goSliceElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goSliceElementType(resolved);
			case _:
				null;
		};
	}

	static function goMapTypePair(type:Type):Null<{keyType:Type, valueType:Type}> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoClass(classType, "Map") && params.length == 2) {
					{
						keyType: params[0],
						valueType: params[1]
					};
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goMapTypePair(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goMapTypePair(resolved);
			case _:
				null;
		};
	}

	static function goResultElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoClass(classType, "Result") && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goResultElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goResultElementType(resolved);
			case _:
				null;
		};
	}

	static function isMonomorphizableMetalType(type:Type):Bool {
		var goType = GoTypeMapper.scalarGoType(type, _ -> "_", _ -> "_");
		return goType != "any";
	}

	static inline function isGoClass(classType:ClassType, className:String):Bool {
		return classType.pack.join(".") == "go" && classType.name == className;
	}

	static inline function moduleNameForClass(classType:ClassType):String {
		if (classType.module != null && classType.module.length > 0) {
			return classType.module;
		}
		return classType.pack == null || classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
	}

	static inline function moduleNameForAbstract(abstractType:AbstractType):String {
		if (abstractType.module != null && abstractType.module.length > 0) {
			return abstractType.module;
		}
		return abstractType.pack == null
			|| abstractType.pack.length == 0 ? abstractType.name : abstractType.pack.join(".") + "." + abstractType.name;
	}

	static function isGoBuild():Bool {
		var targetName = Context.definedValue("target.name");
		return targetName == "go" || Context.defined("go_output");
	}
	#else
	public static function init():Void {}
	#end
}
