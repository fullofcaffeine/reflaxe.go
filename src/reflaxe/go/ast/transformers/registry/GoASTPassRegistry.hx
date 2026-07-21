package reflaxe.go.ast.transformers.registry;

#if macro
import haxe.macro.Context;
import reflaxe.go.compiler.GoCompilerDefine;
#end
import reflaxe.go.CompilationContext;
import reflaxe.go.compiler.GoAutoLoweringModeTools;
import reflaxe.go.ast.transformers.passes.CollectImportsPass;
import reflaxe.go.ast.transformers.passes.CyclicAlphaPass;
import reflaxe.go.ast.transformers.passes.CyclicBetaPass;
import reflaxe.go.ast.transformers.passes.ElideBlankIdentifierGuardsPass;
import reflaxe.go.ast.transformers.passes.InsertRuntimePreludePass;
import reflaxe.go.ast.transformers.passes.NormalizeNamesPass;
import reflaxe.go.ast.transformers.passes.RewriteStringOpsPass;
import reflaxe.go.ast.transformers.passes.RewriteVirtualCallsPass;
import reflaxe.go.ast.transformers.registry.groups.GranularBundle;
import reflaxe.go.ast.transformers.registry.groups.LeanBundle;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

typedef GoASTPassSelectionReason = {
	var pass:String;
	var reason:String;
	var source:String;
}

typedef GoASTPassSelection = {
	var source:String;
	var passes:Array<IGoASTPass>;
	var reasons:Array<GoASTPassSelectionReason>;
}

class GoASTPassRegistry {
	static inline final GRANULAR_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineGranularPassRegistry;
	static inline final LEGACY_LEAN_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineLegacyPassBundle;
	static inline final TEST_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineTestPassRegistryCase;

	public static function resolve(?context:CompilationContext):Array<IGoASTPass> {
		return select(context).passes;
	}

	public static function select(?context:CompilationContext):GoASTPassSelection {
		#if macro
		var testCase = Context.definedValue(TEST_DEFINE);
		if (testCase != null && testCase != "") {
			return selectTestCase(testCase);
		}
		#end

		#if macro
		if (Context.defined(GRANULAR_DEFINE)) {
			return selectionFromBundle(GranularBundle.build(), "legacy_granular_bundle", "Selected by compatibility define `-D " + GRANULAR_DEFINE + "`.");
		}
		if (Context.defined(LEGACY_LEAN_DEFINE)) {
			return selectionFromBundle(LeanBundle.build(), "legacy_lean_bundle", "Selected by compatibility define `-D " + LEGACY_LEAN_DEFINE + "`.");
		}
		#end

		return buildPlannerSelection(context);
	}

	static function buildPlannerSelection(context:Null<CompilationContext>):GoASTPassSelection {
		var presetLabel = "portable_default";
		var autoModeLabel = "off";
		var optimizationPreset = "portable_fast";
		if (context != null && context.buildContext != null) {
			presetLabel = context.buildContext.policyPreset.label();
			autoModeLabel = GoAutoLoweringModeTools.label(context.buildContext.autoLoweringMode);
			optimizationPreset = context.buildContext.optimizationPreset;
		}
		var plannerTag = "planner(preset=" + presetLabel + ", auto=" + autoModeLabel + ", opt=" + optimizationPreset + ")";
		var passNormalize = new NormalizeNamesPass();
		var passRewriteStrings = new RewriteStringOpsPass();
		var passRewriteVirtualCalls = new RewriteVirtualCallsPass();
		var passRuntimePrelude = new InsertRuntimePreludePass();
		var passElideBlankGuards = new ElideBlankIdentifierGuardsPass();
		var passCollectImports = new CollectImportsPass();
		var passes:Array<IGoASTPass> = [
			passNormalize,
			passRewriteStrings,
			passRewriteVirtualCalls,
			passRuntimePrelude,
			passElideBlankGuards,
			passCollectImports
		];
		var reasons:Array<GoASTPassSelectionReason> = [
			reasonForPass(passNormalize.getName(), "Canonicalize generated identifiers before rewrite passes.", plannerTag),
			reasonForPass(passRewriteStrings.getName(), "Apply planner-selected string rewrite/folding pass for deterministic code shape.", plannerTag),
			reasonForPass(passRewriteVirtualCalls.getName(), "Apply planner-selected safe virtual-call rewrite pass.", plannerTag),
			reasonForPass(passRuntimePrelude.getName(), "Inject runtime prelude declarations before cleanup/import collection.", plannerTag),
			reasonForPass(passElideBlankGuards.getName(), "Remove redundant blank-identifier consume guards after lowering.", plannerTag),
			reasonForPass(passCollectImports.getName(), "Collect final deterministic import set after all rewrites.", plannerTag)
		];
		return {
			source: "planner",
			passes: passes,
			reasons: reasons
		};
	}

	static function selectionFromBundle(passes:Array<IGoASTPass>, source:String, reasonText:String):GoASTPassSelection {
		var reasons = new Array<GoASTPassSelectionReason>();
		for (pass in passes) {
			reasons.push(reasonForPass(pass.getName(), reasonText, source));
		}
		return {
			source: source,
			passes: passes,
			reasons: reasons
		};
	}

	static function selectTestCase(testCase:String):GoASTPassSelection {
		var passes:Array<IGoASTPass> = switch (testCase) {
			case "duplicate":
				[new NormalizeNamesPass(), new NormalizeNamesPass()];
			case "missing_dep":
				[new CollectImportsPass()];
			case "cycle":
				[new CyclicAlphaPass(), new CyclicBetaPass()];
			case _:
				#if macro
				Context.fatalError('Unknown Go AST registry test case "' + testCase + '"', Context.currentPos());
				#else
				throw 'Unknown Go AST registry test case "' + testCase + '"';
				#end
				LeanBundle.build();
		};
		return selectionFromBundle(passes, "registry_test_case", 'Selected registry test case "' + testCase + '".');
	}

	static inline function reasonForPass(passName:String, reasonText:String, source:String):GoASTPassSelectionReason {
		return {
			pass: passName,
			reason: reasonText,
			source: source
		};
	}
}
