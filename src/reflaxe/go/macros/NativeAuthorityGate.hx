package reflaxe.go.macros;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.PortableNativePolicyMode;
import reflaxe.go.analyze.GoProfileContractAnalyzer.PortableNativeScanMode;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoNativeAuthorityPolicy;
#end

/**
	Why
	Admission diagnostics for typed Go APIs are an explicit policy axis, not a
	consequence of a semantic profile.

	What
	Applies the configured warning/error policy to unscoped `go.*` usage when
	native authority is `guarded`.

	How
	Runs the shared typed analyzer after typing. `@:goNative` modules and the
	`@:goMetal` compatibility alias are explicit boundaries and are exempt.
**/
class NativeAuthorityGate {
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
		if (buildContext.nativeAuthorityPolicy != GoNativeAuthorityPolicy.Guarded) {
			return;
		}

		var policy = GoProfileContractAnalyzer.resolvePortableNativePolicyModeFromDefines();
		if (policy == PortableNativePolicyMode.Off) {
			return;
		}
		var scanMode = GoProfileContractAnalyzer.resolvePortableNativeScanModeFromDefines();
		var allowPrefixes = GoProfileContractAnalyzer.resolvePortableNativeAllowPrefixesFromDefines();
		var projectRoot = Sys.getCwd();

		Context.onAfterTyping(types -> enforce(types, buildContext, policy, scanMode, projectRoot, allowPrefixes));
	}

	static function enforce(types:Array<ModuleType>, buildContext:reflaxe.go.compiler.GoBuildContext, policy:PortableNativePolicyMode,
			scanMode:PortableNativeScanMode, projectRoot:String, allowPrefixes:Array<String>):Void {
		var diagnostics = GoProfileContractAnalyzer.analyze(types, buildContext, projectRoot, policy, scanMode, allowPrefixes).diagnostics;
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
