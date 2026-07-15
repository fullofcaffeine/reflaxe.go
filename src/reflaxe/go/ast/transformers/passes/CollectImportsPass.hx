package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoImportPath;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class CollectImportsPass implements IGoASTPass {
	public function new() {}

	public function getName():String {
		return "collect_imports";
	}

	public function getDependencies():Array<String> {
		return ["insert_runtime_prelude"];
	}

	public function run(file:GoFile, context:CompilationContext):GoFile {
		var seen = new Map<String, Bool>();
		var imports = new Array<GoImportPath>();

		for (path in file.imports) {
			if (path == null) {
				continue;
			}
			var value = path.value();
			if (!seen.exists(value)) {
				seen.set(value, true);
				imports.push(path);
			}
		}

		imports.sort(function(a, b) return Reflect.compare(a.value(), b.value()));

		return {
			packageName: file.packageName,
			imports: imports,
			decls: file.decls
		};
	}
}
