package reflaxe.go.analyze;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.compiler.GoBuildContext;
import reflaxe.go.compiler.GoTypeMapper;

private typedef AnalyzerElementMethodCall = {
	final methodName:String;
	final elementType:Type;
}

private typedef AnalyzerMapMethodCall = {
	final methodName:String;
	final keyType:Type;
	final valueType:Type;
}

/**
	Central typed analyzer for profile-contract diagnostics and shared contract checks.

	Why:
	- Keep contract rules in one place so macro enforcers and report writers stay consistent.
	- Produce deterministic diagnostics payloads that can be emitted in contract reports.
**/
class GoProfileContractAnalyzer {
	public static inline final PORTABLE_NATIVE_POLICY_DEFINE = "reflaxe_go_portable_native_policy";
	public static inline final PORTABLE_NATIVE_ALLOW_DEFINE = "reflaxe_go_portable_native_allow";

	public static function analyze(types:Array<ModuleType>, buildContext:GoBuildContext, projectRoot:String, portableNativePolicy:PortableNativePolicyMode,
			portableNativeAllowPrefixes:Array<String>):GoProfileContractDiagnostics {
		var diagnostics:Array<GoContractDiagnostic> = [];
		var portableNativeImportHits:Array<String> = [];

		if (buildContext.profile == GoProfile.Portable && portableNativePolicy != PortableNativePolicyMode.Off) {
			var hits = collectPortableNativeImportHits(types, projectRoot, portableNativeAllowPrefixes);
			portableNativeImportHits = [for (hit in hits) hit.module];
			for (hit in hits) {
				diagnostics.push({
					code: "portable_native_import",
					severity: portableNativePolicy == PortableNativePolicyMode.Error ? "error" : "warning",
					module: hit.module,
					location: hit.location,
					message: portableNativeImportMessage(hit.module),
					pos: hit.pos
				});
			}
		}

		diagnostics.sort(compareDiagnostics);
		return {
			diagnostics: diagnostics,
			portableNativeImportHits: portableNativeImportHits
		};
	}

	public static function resolvePortableNativePolicyModeFromDefines():PortableNativePolicyMode {
		return resolvePortableNativePolicyMode(Context.definedValue(PORTABLE_NATIVE_POLICY_DEFINE));
	}

	public static function resolvePortableNativeAllowPrefixesFromDefines():Array<String> {
		return resolvePortableNativeAllowPrefixes(Context.definedValue(PORTABLE_NATIVE_ALLOW_DEFINE));
	}

	public static function resolvePortableNativePolicyMode(rawValue:Null<String>):PortableNativePolicyMode {
		if (rawValue == null || StringTools.trim(rawValue) == "") {
			return PortableNativePolicyMode.Warn;
		}
		var normalized = StringTools.trim(rawValue).toLowerCase();
		return switch (normalized) {
			case "warn":
				PortableNativePolicyMode.Warn;
			case "error":
				PortableNativePolicyMode.Error;
			case "off":
				PortableNativePolicyMode.Off;
			case _:
				Context.fatalError("PortableNativeImportGate: invalid " + PORTABLE_NATIVE_POLICY_DEFINE + " `" + rawValue
					+ "`. Expected `warn`, `error`, or `off`.",
					Context.currentPos());
				PortableNativePolicyMode.Warn;
		};
	}

	public static function resolvePortableNativeAllowPrefixes(rawValue:Null<String>):Array<String> {
		var prefixes:Array<String> = [];
		if (rawValue == null || StringTools.trim(rawValue) == "") {
			return prefixes;
		}
		for (raw in rawValue.split(",")) {
			var trimmed = StringTools.trim(raw);
			if (trimmed != "" && !prefixes.contains(trimmed)) {
				prefixes.push(trimmed);
			}
		}
		prefixes.sort(compareStrings);
		return prefixes;
	}

	public static function collectPortableNativeImportHits(types:Array<ModuleType>, projectRoot:String,
			allowPrefixes:Array<String>):Array<GoPortableNativeImportHit> {
		var hitsByModule:Map<String, GoPortableNativeImportHit> = [];
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!isPortableContractSource(classType.pos, projectRoot)) {
						continue;
					}
					var moduleName = moduleNameForClass(classType);
					if (hitsByModule.exists(moduleName) || isAllowedModule(moduleName, allowPrefixes)) {
						continue;
					}
					if (classUsesGoNative(classType)) {
						hitsByModule.set(moduleName, hitForModule(moduleName, classType.pos));
					}
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!isPortableContractSource(abstractType.pos, projectRoot)) {
						continue;
					}
					var moduleName = moduleNameForAbstract(abstractType);
					if (hitsByModule.exists(moduleName) || isAllowedModule(moduleName, allowPrefixes)) {
						continue;
					}
					if (abstractUsesGoNative(abstractType)) {
						hitsByModule.set(moduleName, hitForModule(moduleName, abstractType.pos));
					}
				case _:
			}
		}

		var out:Array<GoPortableNativeImportHit> = [for (moduleName in hitsByModule.keys()) hitsByModule.get(moduleName)];
		out.sort(comparePortableNativeHits);
		return out;
	}

	public static function isGoInjectionCall(expr:TypedExpr):Bool {
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
		};
	}

	public static function detectLaneTypedFallbackViolation(expr:TypedExpr):Null<String> {
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
				detectLaneTypedFallbackViolationFromCall(callee, expr.t);
			case _:
				null;
		};
	}

	static function detectLaneTypedFallbackViolationFromCall(callee:TypedExpr, returnType:Type):Null<String> {
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

	static function classUsesGoNative(classType:ClassType):Bool {
		if (classType.pack != null && classType.pack.length > 0 && classType.pack[0] == "go") {
			return true;
		}
		var fields = classType.fields.get().concat(classType.statics.get());
		for (field in fields) {
			if (field == null) {
				continue;
			}
			if (typeHasGoNative(field.type)) {
				return true;
			}
			var expr = field.expr();
			if (expr != null && exprHasGoNative(expr)) {
				return true;
			}
		}
		return false;
	}

	static function abstractUsesGoNative(abstractType:AbstractType):Bool {
		if (abstractType.pack != null && abstractType.pack.length > 0 && abstractType.pack[0] == "go") {
			return true;
		}

		if (typeHasGoNative(abstractType.type)) {
			return true;
		}

		if (abstractType.impl != null) {
			var impl = abstractType.impl.get();
			if (impl != null) {
				var fields = impl.fields.get().concat(impl.statics.get());
				for (field in fields) {
					if (field == null) {
						continue;
					}
					if (typeHasGoNative(field.type)) {
						return true;
					}
					var expr = field.expr();
					if (expr != null && exprHasGoNative(expr)) {
						return true;
					}
				}
			}
		}
		return false;
	}

	static function exprHasGoNative(expr:TypedExpr):Bool {
		var found = false;
		function walk(node:TypedExpr):Void {
			if (found || node == null) {
				return;
			}
			if (typeHasGoNative(node.t)) {
				found = true;
				return;
			}
			TypedExprTools.iter(node, child -> walk(child));
		}
		walk(expr);
		return found;
	}

	static function typeHasGoNative(t:Type):Bool {
		return typeHasGoNativeInner(t, []);
	}

	static function typeHasGoNativeInner(t:Type, seen:Array<String>):Bool {
		if (t == null) {
			return false;
		}
		var key = Std.string(t);
		if (seen.indexOf(key) != -1) {
			return false;
		}
		seen.push(key);

		return switch (t) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoPath(classType.pack)) {
					true;
				} else {
					typeListHasGoNative(params, seen);
				}
			case TEnum(enumRef, params):
				var enumType = enumRef.get();
				if (isGoPath(enumType.pack)) {
					true;
				} else {
					typeListHasGoNative(params, seen);
				}
			case TType(typeRef, params):
				var typeDecl = typeRef.get();
				if (isGoPath(typeDecl.pack)) {
					true;
				} else if (typeListHasGoNative(params, seen)) {
					true;
				} else {
					typeHasGoNativeInner(typeDecl.type, seen);
				}
			case TAbstract(absRef, params):
				var absType = absRef.get();
				if (isGoPath(absType.pack)) {
					true;
				} else if (typeListHasGoNative(params, seen)) {
					true;
				} else {
					typeHasGoNativeInner(absType.type, seen);
				}
			case TFun(args, ret): var argHas = false; for (arg in args) {
					if (typeHasGoNativeInner(arg.t, seen)) {
						argHas = true;
						break;
					}
				} argHas || typeHasGoNativeInner(ret, seen);
			case TDynamic(inner): inner != null && typeHasGoNativeInner(inner, seen);
			case TAnonymous(anonRef):
				var anon = anonRef.get();
				var hasGo = false;
				for (field in anon.fields) {
					if (typeHasGoNativeInner(field.type, seen)) {
						hasGo = true;
						break;
					}
				}
				hasGo;
			case TLazy(loader):
				typeHasGoNativeInner(loader(), seen);
			case TMono(monoRef): var candidate = monoRef.get(); candidate != null && typeHasGoNativeInner(candidate, seen);
		}
	}

	static function typeListHasGoNative(types:Array<Type>, seen:Array<String>):Bool {
		for (t in types) {
			if (typeHasGoNativeInner(t, seen)) {
				return true;
			}
		}
		return false;
	}

	static inline function isGoPath(pack:Array<String>):Bool {
		return pack != null && pack.length > 0 && pack[0] == "go";
	}

	static function portableNativeImportMessage(moduleName:String):String {
		return "PortableNativeImportGate: module `"
			+ moduleName
			+ "` uses target-native `go.*` surfaces while "
			+ "`reflaxe_go_profile=portable` is active. Move native usage behind adapters, "
			+ "or use `-D reflaxe_go_portable_native_policy=off|warn|error`.";
	}

	static function isPortableContractSource(pos:haxe.macro.Expr.Position, projectRoot:String):Bool {
		var root = ensureTrailingSlash(projectRoot);
		var file = normalizePath(Context.getPosInfos(pos).file);
		if (file == null || file == "") {
			return false;
		}

		if (!Path.isAbsolute(file)) {
			file = normalizePath(Path.join([root, file]));
		}

		if (!StringTools.startsWith(file, root)) {
			return false;
		}

		if (file.indexOf("/src/reflaxe/") != -1 || file.indexOf("/std/") != -1 || file.indexOf("/src/go/") != -1) {
			return false;
		}

		return true;
	}

	static function isAllowedModule(moduleName:String, allowPrefixes:Array<String>):Bool {
		for (prefix in allowPrefixes) {
			if (moduleName == prefix || StringTools.startsWith(moduleName, prefix + ".")) {
				return true;
			}
		}
		return false;
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

	static function hitForModule(moduleName:String, pos:haxe.macro.Expr.Position):GoPortableNativeImportHit {
		return {
			module: moduleName,
			pos: pos,
			location: locationLabel(moduleName, pos)
		};
	}

	static function locationLabel(moduleName:String, pos:haxe.macro.Expr.Position):String {
		var line = 1;
		var location = PositionTools.toLocation(pos);
		if (location != null && location.range != null && location.range.start != null && location.range.start.line > 0) {
			line = location.range.start.line;
		}
		return moduleName + ":" + line;
	}

	static function comparePortableNativeHits(a:GoPortableNativeImportHit, b:GoPortableNativeImportHit):Int {
		var moduleOrder = compareStrings(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		return compareStrings(a.location, b.location);
	}

	static function compareDiagnostics(a:GoContractDiagnostic, b:GoContractDiagnostic):Int {
		var moduleOrder = compareStrings(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var severityOrder = compareStrings(a.severity, b.severity);
		if (severityOrder != 0) {
			return severityOrder;
		}
		var codeOrder = compareStrings(a.code, b.code);
		if (codeOrder != 0) {
			return codeOrder;
		}
		var locationOrder = compareStrings(a.location, b.location);
		if (locationOrder != 0) {
			return locationOrder;
		}
		return compareStrings(a.message, b.message);
	}

	static inline function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function ensureTrailingSlash(path:String):String {
		var normalized = normalizePath(path);
		return StringTools.endsWith(normalized, "/") ? normalized : normalized + "/";
	}

	static function normalizePath(path:String):String {
		return Path.normalize(path).split("\\").join("/");
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

	static function asGoChanMethodCall(callee:TypedExpr):Null<AnalyzerElementMethodCall> {
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

	static function asGoSliceMethodCall(callee:TypedExpr):Null<AnalyzerElementMethodCall> {
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

	static function asGoMapMethodCall(callee:TypedExpr):Null<AnalyzerMapMethodCall> {
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

	static function asGoResultMethodCall(callee:TypedExpr):Null<AnalyzerElementMethodCall> {
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
}

enum PortableNativePolicyMode {
	Off;
	Warn;
	Error;
}

typedef GoPortableNativeImportHit = {
	var module:String;
	var location:String;
	var pos:haxe.macro.Expr.Position;
}

typedef GoContractDiagnostic = {
	var code:String;
	var severity:String;
	var module:String;
	var location:String;
	var message:String;
	var pos:haxe.macro.Expr.Position;
}

typedef GoProfileContractDiagnostics = {
	var diagnostics:Array<GoContractDiagnostic>;
	var portableNativeImportHits:Array<String>;
}
#end
