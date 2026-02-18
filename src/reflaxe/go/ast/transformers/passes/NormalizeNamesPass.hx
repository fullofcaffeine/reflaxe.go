package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class NormalizeNamesPass implements IGoASTPass {
  public function new() {}

  public function getName():String {
    return "normalize_names";
  }

  public function getDependencies():Array<String> {
    return [];
  }

  public function run(file:GoFile, context:CompilationContext):GoFile {
    return file;
  }
}
