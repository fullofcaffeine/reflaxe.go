package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.compiler.GoBuildContextResolver;
#end

class PortableNativeImportGate {
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

		var mode = resolveMode();
		if (mode == Off) {
			return;
		}

		var projectRoot = normalizePath(Sys.getCwd());
		var allowPrefixes = resolveAllowPrefixes();
		Context.onAfterTyping(types -> enforce(types, mode, projectRoot, allowPrefixes));
	}

	static function enforce(types:Array<ModuleType>, mode:PortableNativePolicyMode, projectRoot:String, allowPrefixes:Array<String>):Void {
		var reportedModules:Map<String, Bool> = [];
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!isPortableContractSource(classType.pos, projectRoot)) {
						continue;
					}
					var moduleName = moduleNameForClass(classType);
					if (reportedModules.exists(moduleName) || isAllowedModule(moduleName, allowPrefixes)) {
						continue;
					}
					if (classUsesGoNative(classType)) {
						reportViolation(mode, moduleName, classType.pos);
						reportedModules.set(moduleName, true);
					}
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!isPortableContractSource(abstractType.pos, projectRoot)) {
						continue;
					}
					var moduleName = moduleNameForAbstract(abstractType);
					if (reportedModules.exists(moduleName) || isAllowedModule(moduleName, allowPrefixes)) {
						continue;
					}
					if (abstractUsesGoNative(abstractType)) {
						reportViolation(mode, moduleName, abstractType.pos);
						reportedModules.set(moduleName, true);
					}
				case _:
			}
		}
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

	static function reportViolation(mode:PortableNativePolicyMode, moduleName:String, pos:haxe.macro.Expr.Position):Void {
		var message = "PortableNativeImportGate: module `"
			+ moduleName
			+ "` uses target-native `go.*` surfaces while "
			+ "`reflaxe_go_profile=portable` is active. Move native usage behind adapters, "
			+ "or use `-D reflaxe_go_portable_native_policy=off|warn|error`.";
		switch (mode) {
			case Error:
				Context.error(message, pos);
			case Warn:
				Context.warning(message, pos);
			case Off:
		}
	}

	static function resolveMode():PortableNativePolicyMode {
		var value = Context.definedValue("reflaxe_go_portable_native_policy");
		if (value == null || StringTools.trim(value) == "") {
			return Warn;
		}
		switch (StringTools.trim(value).toLowerCase()) {
			case "warn":
				return Warn;
			case "error":
				return Error;
			case "off":
				return Off;
			case _:
				Context.fatalError("PortableNativeImportGate: invalid reflaxe_go_portable_native_policy `" + value + "`. Expected `warn`, `error`, or `off`.",
					Context.currentPos());
				return Warn;
		}
	}

	static function resolveAllowPrefixes():Array<String> {
		var value = Context.definedValue("reflaxe_go_portable_native_allow");
		if (value == null || StringTools.trim(value) == "") {
			return [];
		}
		var prefixes = [];
		for (raw in value.split(",")) {
			var trimmed = StringTools.trim(raw);
			if (trimmed != "") {
				prefixes.push(trimmed);
			}
		}
		return prefixes;
	}

	static function isAllowedModule(moduleName:String, allowPrefixes:Array<String>):Bool {
		for (prefix in allowPrefixes) {
			if (moduleName == prefix || StringTools.startsWith(moduleName, prefix + ".")) {
				return true;
			}
		}
		return false;
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

	static function ensureTrailingSlash(path:String):String {
		var normalized = normalizePath(path);
		return StringTools.endsWith(normalized, "/") ? normalized : normalized + "/";
	}

	static function normalizePath(path:String):String {
		return Path.normalize(path).split("\\").join("/");
	}

	static function isGoBuild():Bool {
		var targetName = Context.definedValue("target.name");
		return targetName == "go" || Context.defined("go_output");
	}
	#else
	public static function init():Void {}
	#end
}

#if macro
private enum PortableNativePolicyMode {
	Warn;
	Error;
	Off;
}
#end
