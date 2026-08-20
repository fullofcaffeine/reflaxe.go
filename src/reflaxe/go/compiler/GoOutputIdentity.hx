package reflaxe.go.compiler;

import reflaxe.go.ast.GoPackageName;
import reflaxe.go.compiler.GoProjectMode.GoEntrypointSymbol;

/** Selects the stable filename family for generated Go source. */
enum GoGeneratedFileStyle {
	StandaloneFiles;
	ExistingModuleFiles;
}

typedef GoOutputIdentityData = {
	final packageName:GoPackageName;
	final entrySymbol:GoEntrypointSymbol;
	final runtimeImportPath:String;
	final fileStyle:GoGeneratedFileStyle;
}

/**
	The target-owned identity shared by naming, package printing, and lifecycle passes.

	Keeping it separate from the Haxe-selected main identity prevents source class
	names from silently deciding a caller package or exported bridge symbol.
**/
class GoOutputIdentity {
	public final packageName:GoPackageName;
	public final entrySymbol:GoEntrypointSymbol;
	public final runtimeImportPath:String;
	public final fileStyle:GoGeneratedFileStyle;

	public function new(data:GoOutputIdentityData) {
		packageName = data.packageName;
		entrySymbol = data.entrySymbol;
		runtimeImportPath = data.runtimeImportPath;
		fileStyle = data.fileStyle;
	}

	public static function standalone(runtimeImportPath:String):GoOutputIdentity {
		return new GoOutputIdentity({
			packageName: GoPackageName.named("main"),
			entrySymbol: GoEntrypointSymbol.named("main"),
			runtimeImportPath: runtimeImportPath,
			fileStyle: StandaloneFiles
		});
	}

	public inline function usesExistingModuleFiles():Bool {
		return fileStyle == ExistingModuleFiles;
	}
}
