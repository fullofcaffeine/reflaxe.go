package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.analyze.MetalLaneAnalyzer;
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
					enforceClassFields(classType.fields.get());
					enforceClassFields(classType.statics.get());
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!laneModules.exists(moduleNameForAbstract(abstractType))) {
						continue;
					}
					if (abstractType.impl != null) {
						var impl = abstractType.impl.get();
						if (impl != null) {
							enforceClassFields(impl.fields.get());
							enforceClassFields(impl.statics.get());
						}
					}
				case _:
			}
		}
	}

	static function enforceClassFields(fields:Array<ClassField>):Void {
		if (fields == null) {
			return;
		}
		for (field in fields) {
			var expr = field.expr();
			if (expr == null) {
				continue;
			}
			scanForGoInjection(expr);
		}
	}

	static function scanForGoInjection(expr:TypedExpr):Void {
		if (isGoInjectionCall(expr)) {
			Context.error("MetalLaneEnforcer: __go__ is not allowed in @:goMetal modules when contract=portable.", expr.pos);
		}
		TypedExprTools.iter(expr, e -> scanForGoInjection(e));
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
