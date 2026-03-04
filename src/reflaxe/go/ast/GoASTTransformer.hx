package reflaxe.go.ast;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.transformers.registry.GoASTPassRegistry;
import reflaxe.go.ast.transformers.registry.RegistryCore;

class GoASTTransformer {
	public static function transform(file:GoFile, context:CompilationContext):GoFile {
		var selection = GoASTPassRegistry.select(context);
		var passes = RegistryCore.validateAndOrder(selection.passes);
		context.selectedGoAstPassSource = selection.source;
		var reasonByPass = new Map<String, {reason:String, source:String}>();
		for (reason in selection.reasons) {
			reasonByPass.set(reason.pass, {
				reason: reason.reason,
				source: reason.source
			});
		}
		context.selectedGoAstPassReasons = [];
		for (pass in passes) {
			var name = pass.getName();
			if (context.appliedGoAstPassNames.indexOf(name) == -1) {
				context.appliedGoAstPassNames.push(name);
			}
			var reason = reasonByPass.get(name);
			context.selectedGoAstPassReasons.push({
				pass: name,
				reason: reason == null ? "No planner reason recorded; pass selected by registry ordering." : reason.reason,
				source: reason == null ? selection.source : reason.source
			});
		}
		var out = file;
		for (pass in passes) {
			out = pass.run(out, context);
		}
		return out;
	}
}
