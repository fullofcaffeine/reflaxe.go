package reflaxe.go.ast;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.transformers.registry.GoASTPassRegistry;
import reflaxe.go.ast.transformers.registry.RegistryCore;

class GoASTTransformer {
  public static function transform(file:GoFile, context:CompilationContext):GoFile {
    var passes = RegistryCore.validateAndOrder(GoASTPassRegistry.resolve());
    var out = file;
    for (pass in passes) {
      out = pass.run(out, context);
    }
    return out;
  }
}
