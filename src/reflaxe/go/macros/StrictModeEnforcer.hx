package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoRawInjectionAuthorityAnalyzer;
import reflaxe.go.compiler.GoBuildContextResolver;
import sys.FileSystem;
import sys.io.File;
#end

class StrictModeEnforcer {
	#if macro
	static var initialized = false;
	static final FRAMEWORK_TYPED_INJECTION_PATHS = ["/src/reflaxe/go/macros/", "/src/go/"];

	public static function init():Void {
		if (initialized) {
			return;
		}
		initialized = true;

		if (!isGoBuild()) {
			return;
		}

		var projectRoot = normalizePath(Sys.getCwd());
		var allowFrameworkTypedInjections = GoBuildContextResolver.resolve().hasExplicitNativeAuthority();
		var preflightFindings = preflightScanForGoInjections(projectRoot);
		if (preflightFindings.length > 0) {
			Context.fatalError("StrictModeEnforcer: __go__ is not allowed in strict mode (" + preflightFindings[0] + ")", Context.currentPos());
		}
		Context.onAfterTyping(types -> enforce(types, projectRoot, allowFrameworkTypedInjections, allowedRawInjectionModules(types)));
	}

	static function enforce(types:Array<ModuleType>, projectRoot:String, allowFrameworkTypedInjections:Bool, allowedRawModules:Map<String, Bool>):Void {
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!isStrictProjectSource(classType.pos, projectRoot)) {
						continue;
					}
					enforceNoGoInjectionInClass(classType, projectRoot, allowFrameworkTypedInjections, allowedRawModules.exists(moduleNameForClass(classType)));
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!isStrictProjectSource(abstractType.pos, projectRoot) || abstractType.impl == null) {
						continue;
					}
					var impl = abstractType.impl.get();
					if (impl != null) {
						enforceNoGoInjectionInFields(impl.fields.get().concat(impl.statics.get()), projectRoot, allowFrameworkTypedInjections,
							allowedRawModules.exists(moduleNameForAbstract(abstractType)));
					}
				case _:
			}
		}
	}

	static function enforceNoGoInjectionInClass(classType:ClassType, projectRoot:String, allowFrameworkTypedInjections:Bool,
			allowScopedRawAuthority:Bool):Void {
		enforceNoGoInjectionInFields(classType.fields.get().concat(classType.statics.get()), projectRoot, allowFrameworkTypedInjections,
			allowScopedRawAuthority);
	}

	static function enforceNoGoInjectionInFields(fields:Array<ClassField>, projectRoot:String, allowFrameworkTypedInjections:Bool,
			allowScopedRawAuthority:Bool):Void {
		for (field in fields) {
			var expr = field.expr();
			if (expr == null) {
				continue;
			}
			scanForGoInjection(expr, projectRoot, allowFrameworkTypedInjections, allowScopedRawAuthority);
		}
	}

	static function scanForGoInjection(expr:TypedExpr, projectRoot:String, allowFrameworkTypedInjections:Bool, allowScopedRawAuthority:Bool):Void {
		if (GoProfileContractAnalyzer.isGoInjectionCall(expr)) {
			if (allowScopedRawAuthority) {
				TypedExprTools.iter(expr, e -> scanForGoInjection(e, projectRoot, allowFrameworkTypedInjections, allowScopedRawAuthority));
				return;
			}
			if (allowFrameworkTypedInjections && isFrameworkTypedInjectionExpr(expr.pos, projectRoot)) {
				TypedExprTools.iter(expr, e -> scanForGoInjection(e, projectRoot, allowFrameworkTypedInjections, allowScopedRawAuthority));
				return;
			}
			Context.error("StrictModeEnforcer: __go__ is not allowed in strict mode. "
				+ "Prefer a typed wrapper or move target-specific interop into `std/`.", expr.pos);
		}

		TypedExprTools.iter(expr, e -> scanForGoInjection(e, projectRoot, allowFrameworkTypedInjections, allowScopedRawAuthority));
	}

	static function isStrictProjectSource(pos:haxe.macro.Expr.Position, projectRoot:String):Bool {
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

		if (file.indexOf("/src/reflaxe/") != -1 || file.indexOf("/std/") != -1) {
			return false;
		}

		return true;
	}

	static function isFrameworkTypedInjectionExpr(pos:haxe.macro.Expr.Position, projectRoot:String):Bool {
		var root = ensureTrailingSlash(projectRoot);
		var file = normalizePath(Context.getPosInfos(pos).file);
		if (file == null || file == "") {
			return false;
		}

		if (!Path.isAbsolute(file)) {
			file = normalizePath(Path.join([root, file]));
		}

		if (!StringTools.startsWith(file, root)) {
			return true;
		}

		for (authorityPath in FRAMEWORK_TYPED_INJECTION_PATHS) {
			if (file.indexOf(authorityPath) != -1) {
				return true;
			}
		}
		return false;
	}

	static function preflightScanForGoInjections(projectRoot:String):Array<String> {
		var root = ensureTrailingSlash(projectRoot);
		var files = new Array<String>();

		for (classPath in Context.getClassPath()) {
			var full = absolutePath(classPath);
			var fullWithSlash = ensureTrailingSlash(full);
			if (!StringTools.startsWith(full, root) && !StringTools.startsWith(fullWithSlash, root)) {
				continue;
			}
			if (full.indexOf("/src/reflaxe/") != -1 || full.indexOf("/std/") != -1) {
				continue;
			}
			collectHxFiles(full, files);
		}

		var findings = new Array<String>();
		for (path in files) {
			var content = File.getContent(path);
			if (StringTools.contains(content, "__go__(") && !GoRawInjectionAuthorityAnalyzer.sourceTextHasRawAuthorityMarker(content)) {
				findings.push(path);
			}
		}

		return findings;
	}

	static function absolutePath(path:String):String {
		if (Path.isAbsolute(path)) {
			return normalizePath(path);
		}
		return normalizePath(Path.join([Sys.getCwd(), path]));
	}

	static function collectHxFiles(path:String, out:Array<String>):Void {
		if (!FileSystem.exists(path)) {
			return;
		}
		if (FileSystem.isDirectory(path)) {
			for (entry in FileSystem.readDirectory(path)) {
				collectHxFiles(Path.join([path, entry]), out);
			}
			return;
		}
		if (StringTools.endsWith(path, ".hx")) {
			out.push(path);
		}
	}

	static function ensureTrailingSlash(path:String):String {
		var normalized = normalizePath(path);
		return StringTools.endsWith(normalized, "/") ? normalized : normalized + "/";
	}

	static function allowedRawInjectionModules(types:Array<ModuleType>):Map<String, Bool> {
		var out:Map<String, Bool> = [];
		var snapshot = GoRawInjectionAuthorityAnalyzer.collect(types);
		for (module in snapshot.modules) {
			out.set(module, true);
		}
		return out;
	}

	static function moduleNameForClass(classType:ClassType):String {
		if (classType.module != null && classType.module.length > 0) {
			return classType.module;
		}
		return pathFromPack(classType.pack, classType.name);
	}

	static function moduleNameForAbstract(abstractType:AbstractType):String {
		if (abstractType.module != null && abstractType.module.length > 0) {
			return abstractType.module;
		}
		return pathFromPack(abstractType.pack, abstractType.name);
	}

	static function pathFromPack(pack:Array<String>, name:String):String {
		return pack == null || pack.length == 0 ? name : pack.join(".") + "." + name;
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
