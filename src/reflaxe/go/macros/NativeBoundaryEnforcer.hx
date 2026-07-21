package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.analyze.GoNativeBoundaryAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.compiler.GoAutoLoweringMode;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoCompilerDefine;
#end

/**
	Why
	A native module boundary is a source-level contract and must behave the same
	under every build policy preset.

	What
	Enforces controlled interop and strict typed fallback rules in modules marked
	with `@:goNative` or its `@:goMetal` compatibility alias.

	How
	After typing, scans only declared native-boundary modules. Raw injection is
	always rejected; `auto_strict` additionally rejects unresolved typed native
	representations before lowering.
**/
class NativeBoundaryEnforcer {
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
		var enforceTypedFallbacks = buildContext.autoLoweringMode == GoAutoLoweringMode.AutoStrict;
		Context.onAfterTyping(types -> enforce(types, enforceTypedFallbacks));
	}

	static function enforce(types:Array<ModuleType>, enforceTypedFallbacks:Bool):Void {
		var boundarySnapshot = GoNativeBoundaryAnalyzer.collect(types);
		if (boundarySnapshot.modules.length == 0) {
			return;
		}

		var boundaryModules:Map<String, Bool> = [];
		for (moduleName in boundarySnapshot.modules) {
			boundaryModules.set(moduleName, true);
		}

		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					var moduleName = moduleNameForClass(classType);
					if (!boundaryModules.exists(moduleName)) {
						continue;
					}
					enforceClassFields(classType.fields.get(), moduleName, enforceTypedFallbacks);
					enforceClassFields(classType.statics.get(), moduleName, enforceTypedFallbacks);
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					var moduleName = moduleNameForAbstract(abstractType);
					if (!boundaryModules.exists(moduleName)) {
						continue;
					}
					if (abstractType.impl != null) {
						var impl = abstractType.impl.get();
						if (impl != null) {
							enforceClassFields(impl.fields.get(), moduleName, enforceTypedFallbacks);
							enforceClassFields(impl.statics.get(), moduleName, enforceTypedFallbacks);
						}
					}
				case _:
			}
		}
	}

	static function enforceClassFields(fields:Array<ClassField>, moduleName:String, enforceTypedFallbacks:Bool):Void {
		if (fields == null) {
			return;
		}
		for (field in fields) {
			var expr = field.expr();
			if (expr == null) {
				continue;
			}
			scanForBoundaryViolations(expr, moduleName, enforceTypedFallbacks);
		}
	}

	static function scanForBoundaryViolations(expr:TypedExpr, moduleName:String, enforceTypedFallbacks:Bool):Void {
		if (GoProfileContractAnalyzer.isGoInjectionCall(expr)) {
			Context.error("NativeBoundaryEnforcer: __go__ is not allowed in @:goNative modules (@:goMetal is a compatibility alias).", expr.pos);
		}
		if (enforceTypedFallbacks) {
			var typedFallback = GoProfileContractAnalyzer.detectNativeBoundaryTypedFallbackViolation(expr);
			if (typedFallback != null) {
				Context.error('NativeBoundaryEnforcer: typed native fallback is not allowed in native-boundary module "'
					+ moduleName
					+ '" with `-D reflaxe_go_auto=auto_strict`. '
					+ typedFallback
					+ " Use concrete generic types for go.Chan/go.Slice/go.Map/go.Result in native-boundary modules (avoid Dynamic/Any).",
					expr.pos);
			}
		}
		TypedExprTools.iter(expr, e -> scanForBoundaryViolations(e, moduleName, enforceTypedFallbacks));
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
		var targetName = Context.definedValue(GoCompilerDefine.DefineTargetName);
		return targetName == GoCompilerDefine.DefineGoTarget || Context.defined(GoCompilerDefine.DefineGoOutput);
	}
	#else
	public static function init():Void {}
	#end
}
