package reflaxe.go.ast.transformers.registry.groups;

import reflaxe.go.ast.transformers.passes.CollectImportsPass;
import reflaxe.go.ast.transformers.passes.InsertRuntimePreludePass;
import reflaxe.go.ast.transformers.passes.NormalizeNamesPass;
import reflaxe.go.ast.transformers.passes.RewriteStringOpsPass;
import reflaxe.go.ast.transformers.passes.RewriteVirtualCallsPass;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class LeanBundle {
  public static function build():Array<IGoASTPass> {
    return [
      new NormalizeNamesPass(),
      new RewriteStringOpsPass(),
      new RewriteVirtualCallsPass(),
      new InsertRuntimePreludePass(),
      new CollectImportsPass()
    ];
  }
}
