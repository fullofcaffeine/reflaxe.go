package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.MetalLaneAnalyzer;
import reflaxe.go.compiler.GoAutoLoweringMode;
import reflaxe.go.compiler.GoBuildContextResolver;
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

		var enforceTypedFallbacks = buildContext.autoLoweringMode == GoAutoLoweringMode.AutoStrict;
		Context.onAfterTyping(types -> enforce(types, enforceTypedFallbacks));
	}

	static function enforce(types:Array<ModuleType>, enforceTypedFallbacks:Bool):Void {
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
					var moduleName = moduleNameForClass(classType);
					if (!laneModules.exists(moduleName)) {
						continue;
					}
					enforceClassFields(classType.fields.get(), moduleName, enforceTypedFallbacks);
					enforceClassFields(classType.statics.get(), moduleName, enforceTypedFallbacks);
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					var moduleName = moduleNameForAbstract(abstractType);
					if (!laneModules.exists(moduleName)) {
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
			scanForLaneViolations(expr, moduleName, enforceTypedFallbacks);
		}
	}

	static function scanForLaneViolations(expr:TypedExpr, moduleName:String, enforceTypedFallbacks:Bool):Void {
		if (GoProfileContractAnalyzer.isGoInjectionCall(expr)) {
			Context.error("MetalLaneEnforcer: __go__ is not allowed in @:goMetal modules when contract=portable.", expr.pos);
		}
		if (enforceTypedFallbacks) {
			var typedFallback = GoProfileContractAnalyzer.detectLaneTypedFallbackViolation(expr);
			if (typedFallback != null) {
				Context.error('MetalLaneEnforcer: typed metal fallback is not allowed in @:goMetal module "'
					+ moduleName
					+ '" when contract=portable and `-D reflaxe_go_auto=auto_strict`. '
					+ typedFallback
					+ " Use concrete generic types for go.Chan/go.Slice/go.Map/go.Result in lane modules (avoid Dynamic/Any).",
					expr.pos);
			}
		}
		TypedExprTools.iter(expr, e -> scanForLaneViolations(e, moduleName, enforceTypedFallbacks));
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
