package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.ProfileResolver;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import sys.FileSystem;
import sys.io.File;
#end

class BoundaryEnforcer {
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
		if (!Context.defined("reflaxe_go_strict_examples")) {
			return;
		}

		var allowFrameworkTypedInjections = ProfileResolver.resolve() == GoProfile.Metal;
		var preflightFindings = preflightScanForGoInjections();
		if (preflightFindings.length > 0) {
			Context.fatalError("BoundaryEnforcer: __go__ is not allowed in strict examples (" + preflightFindings[0] + ")", Context.currentPos());
		}
		Context.onAfterTyping(types -> enforceExampleBoundaries(types, allowFrameworkTypedInjections));
	}

	static function enforceExampleBoundaries(types:Array<ModuleType>, allowFrameworkTypedInjections:Bool):Void {
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!isExampleSource(classType.pos)) {
						continue;
					}
					enforceNoGoInjectionInClass(classType, allowFrameworkTypedInjections);
				case _:
			}
		}
	}

	static function enforceNoGoInjectionInClass(classType:ClassType, allowFrameworkTypedInjections:Bool):Void {
		var allFields = classType.fields.get().concat(classType.statics.get());
		for (field in allFields) {
			var expr = field.expr();
			if (expr == null) {
				continue;
			}
			scanForGoInjection(expr, allowFrameworkTypedInjections);
		}
	}

	static function scanForGoInjection(expr:TypedExpr, allowFrameworkTypedInjections:Bool):Void {
		if (GoProfileContractAnalyzer.isGoInjectionCall(expr)) {
			if (allowFrameworkTypedInjections && isFrameworkTypedInjectionExpr(expr.pos)) {
				TypedExprTools.iter(expr, e -> scanForGoInjection(e, allowFrameworkTypedInjections));
				return;
			}
			Context.error("BoundaryEnforcer: __go__ is not allowed in strict examples. "
				+ "Implement the feature in Haxe or add a reusable framework wrapper in `std/`.",
				expr.pos);
		}

		TypedExprTools.iter(expr, e -> scanForGoInjection(e, allowFrameworkTypedInjections));
	}

	static function isExampleSource(pos:haxe.macro.Expr.Position):Bool {
		var file = Context.getPosInfos(pos).file;
		if (file == null || file == "") {
			return false;
		}

		var cwd = normalizePath(Sys.getCwd());
		var normalized = normalizePath(file);
		if (!Path.isAbsolute(normalized)) {
			normalized = normalizePath(Path.join([cwd, normalized]));
		}

		return normalized.indexOf("/examples/") != -1 || normalized.indexOf("/test/snapshot/") != -1;
	}

	static function isFrameworkTypedInjectionExpr(pos:haxe.macro.Expr.Position):Bool {
		var file = Context.getPosInfos(pos).file;
		if (file == null || file == "") {
			return false;
		}

		var cwd = normalizePath(Sys.getCwd());
		var normalized = normalizePath(file);
		if (!Path.isAbsolute(normalized)) {
			normalized = normalizePath(Path.join([cwd, normalized]));
		}

		return normalized.indexOf("/src/reflaxe/go/macros/") != -1 || normalized.indexOf("/std/go/metal/") != -1;
	}

	static function preflightScanForGoInjections():Array<String> {
		var files = new Array<String>();
		var cwd = normalizePath(Sys.getCwd());

		for (classPath in Context.getClassPath()) {
			var full = absolutePath(classPath);
			if (!StringTools.startsWith(full, cwd)) {
				continue;
			}
			if (full.indexOf("/examples/") == -1 && full.indexOf("/test/") == -1) {
				continue;
			}
			collectHxFiles(full, files);
		}

		var findings = new Array<String>();
		for (path in files) {
			var content = File.getContent(path);
			if (StringTools.contains(content, "__go__(")) {
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
