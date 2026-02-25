package reflaxe.go;

#if macro
import haxe.macro.Compiler as MacroCompiler;
import haxe.macro.Context;
import reflaxe.BaseCompiler.BaseCompilerFileOutputType;
import reflaxe.ReflectCompiler;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.macros.AutoEmptyConstructor;
import reflaxe.go.macros.BoundaryEnforcer;
import reflaxe.go.macros.MetalLaneEnforcer;
import reflaxe.go.macros.StrictModeEnforcer;
#end

class CompilerInit {
	#if macro
	static var initialized = false;

	public static function Start():Void {
		if (!BuildDetection.isGoBuild()) {
			return;
		}

		if (initialized) {
			return;
		}
		initialized = true;

		var buildContext = GoBuildContextResolver.resolve();
		AutoEmptyConstructor.init();
		if (buildContext.strictExamples) {
			BoundaryEnforcer.init();
		}
		if (buildContext.strictUserBoundaries) {
			StrictModeEnforcer.init();
		}
		MetalLaneEnforcer.init();

		// Enable stdlib atomic surfaces guarded behind target.atomics.
		MacroCompiler.define("target.atomics");

		ReflectCompiler.Start();
		ReflectCompiler.AddCompiler(new GoReflaxeCompiler(), {
			outputDirDefineName: "go_output",
			fileOutputType: Manual,
			fileOutputExtension: ".go",
			targetCodeInjectionName: "__go__",
			expressionPreprocessors: [],
			ignoreBodilessFunctions: false,
			ignoreExterns: true
		});
	}
	#else
	public static function Start():Void {}
	#end
}
