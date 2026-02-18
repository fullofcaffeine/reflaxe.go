package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class CyclicAlphaPass implements IGoASTPass {
  public function new() {}

  public function getName():String {
    return "cycle_alpha";
  }

  public function getDependencies():Array<String> {
    return ["cycle_beta"];
  }

  public function run(file:GoFile, context:CompilationContext):GoFile {
    return file;
  }
}
