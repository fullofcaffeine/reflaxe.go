package reflaxe.go.ast.transformers.registry;

#if macro
import haxe.macro.Context;
#end
import reflaxe.go.CompilationContext;
import reflaxe.go.compiler.GoAutoLoweringModeTools;
import reflaxe.go.GoProfile;
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
	static inline final GRANULAR_DEFINE = "go_granular_pass_registry";
	static inline final LEGACY_LEAN_DEFINE = "reflaxe_go_legacy_pass_bundle";
	static inline final TEST_DEFINE = "reflaxe_go_test_registry_case";

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
		var contractLabel = "portable";
		var autoModeLabel = "off";
		var optimizationPreset = "portable_fast";
		if (context != null && context.buildContext != null) {
			contractLabel = context.profile == GoProfile.Metal ? "metal" : "portable";
			autoModeLabel = GoAutoLoweringModeTools.label(context.buildContext.autoLoweringMode);
			optimizationPreset = context.buildContext.optimizationPreset;
		}
		var plannerTag = "planner(contract=" + contractLabel + ", auto=" + autoModeLabel + ", opt=" + optimizationPreset + ")";
		var passes:Array<IGoASTPass> = [
			new NormalizeNamesPass(),
			new RewriteStringOpsPass(),
			new RewriteVirtualCallsPass(),
			new InsertRuntimePreludePass(),
			new ElideBlankIdentifierGuardsPass(),
			new CollectImportsPass()
		];
		var reasons:Array<GoASTPassSelectionReason> = [
			reasonForPass("NormalizeNamesPass", "Canonicalize generated identifiers before rewrite passes.", plannerTag),
			reasonForPass("RewriteStringOpsPass", "Apply planner-selected string rewrite/folding pass for deterministic code shape.", plannerTag),
			reasonForPass("RewriteVirtualCallsPass", "Apply planner-selected safe virtual-call rewrite pass.", plannerTag),
			reasonForPass("InsertRuntimePreludePass", "Inject runtime prelude declarations before cleanup/import collection.", plannerTag),
			reasonForPass("ElideBlankIdentifierGuardsPass", "Remove redundant blank-identifier consume guards after lowering.", plannerTag),
			reasonForPass("CollectImportsPass", "Collect final deterministic import set after all rewrites.", plannerTag)
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
