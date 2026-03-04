package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.GoProfile;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.PortableNativePolicyMode;
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

		var policy = GoProfileContractAnalyzer.resolvePortableNativePolicyModeFromDefines();
		if (policy == PortableNativePolicyMode.Off) {
			return;
		}
		var allowPrefixes = GoProfileContractAnalyzer.resolvePortableNativeAllowPrefixesFromDefines();
		var projectRoot = Sys.getCwd();

		Context.onAfterTyping(types -> enforce(types, buildContext, policy, projectRoot, allowPrefixes));
	}

	static function enforce(types:Array<ModuleType>, buildContext:reflaxe.go.compiler.GoBuildContext, policy:PortableNativePolicyMode, projectRoot:String,
			allowPrefixes:Array<String>):Void {
		var diagnostics = GoProfileContractAnalyzer.analyze(types, buildContext, projectRoot, policy, allowPrefixes).diagnostics;
		for (entry in diagnostics) {
			switch (entry.severity) {
				case "error":
					Context.error(entry.message, entry.pos);
				case "warning":
					Context.warning(entry.message, entry.pos);
				case _:
			}
		}
	}

	static function isGoBuild():Bool {
		var targetName = Context.definedValue("target.name");
		return targetName == "go" || Context.defined("go_output");
	}
	#else
	public static function init():Void {}
	#end
}
