package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class InsertRuntimePreludePass implements IGoASTPass {
  public function new() {}

  public function getName():String {
    return "insert_runtime_prelude";
  }

  public function getDependencies():Array<String> {
    return ["rewrite_string_ops", "rewrite_virtual_calls"];
  }

  public function run(file:GoFile, context:CompilationContext):GoFile {
    var hasRuntime = false;
    for (path in file.imports) {
      if (path == "snapshot/hxrt") {
        hasRuntime = true;
        break;
      }
    }

    if (hasRuntime) {
      return file;
    }

    var imports = file.imports.copy();
    imports.push("snapshot/hxrt");
    return {
      packageName: file.packageName,
      imports: imports,
      decls: file.decls
    };
  }
}
