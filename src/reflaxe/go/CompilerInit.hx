package reflaxe.go;

#if macro
import haxe.macro.Compiler as MacroCompiler;
import haxe.macro.Context;
import reflaxe.BaseCompiler.BaseCompilerFileOutputType;
import reflaxe.ReflectCompiler;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoCompilerDefine;
import reflaxe.go.compiler.SiblingTargetConflictGuard;
import reflaxe.go.macros.AutoEmptyConstructor;
import reflaxe.go.macros.BoundaryEnforcer;
import reflaxe.go.macros.NativeBoundaryEnforcer;
import reflaxe.go.macros.NativeAuthorityGate;
import reflaxe.go.macros.SourceOwnedStdlibRetention;
import reflaxe.go.macros.StrictModeEnforcer;
#end

class CompilerInit {
	#if macro
	static var initialized = false;

	public static function Start():Void {
		if (!BuildDetection.isGoBuild()) {
			return;
		}
		SiblingTargetConflictGuard.init();

		if (initialized) {
			return;
		}
		initialized = true;

		var buildContext = GoBuildContextResolver.resolve();
		SourceOwnedStdlibRetention.init();
		AutoEmptyConstructor.init();
		if (buildContext.strictExamples) {
			BoundaryEnforcer.init();
		}
		if (buildContext.strictUserBoundaries) {
			StrictModeEnforcer.init();
		}
		NativeBoundaryEnforcer.init();
		NativeAuthorityGate.init();

		// What: Declare the target capabilities implemented by staged std/hxrt.
		// Why: Packaged `.cross.hx` dependencies can inspect these defines before
		// their types are built; a source-checkout-only `_std` path can hide a
		// missing capability declaration.
		// How: Publish the capability before Reflaxe starts typing user and library
		// modules. Focused target-capability snapshots keep the declarations honest.
		MacroCompiler.define(GoCompilerDefine.DefineTargetAtomics);
		MacroCompiler.define(GoCompilerDefine.DefineTargetThreaded);

		ReflectCompiler.Start();
		ReflectCompiler.AddCompiler(new GoReflaxeCompiler(), {
			outputDirDefineName: GoCompilerDefine.DefineGoOutput,
			fileOutputType: Manual,
			fileOutputExtension: ".go",
			targetCodeInjectionName: "__go__",
			expressionPreprocessors: [],
			ignoreBodilessFunctions: false,
			ignoreExterns: true,
			trackUsedTypes: true
		});
	}
	#else
	public static function Start():Void {}
	#end
}
