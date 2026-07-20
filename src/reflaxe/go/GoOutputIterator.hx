package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import reflaxe.go.GoCompiler.GoGeneratedFile;
import reflaxe.go.compiler.GoGeneratedOutputBoundary;
import sys.FileSystem;
#end

/**
	What: Compatibility helpers for legacy iterator-driven Go output.

	Why: The registered compiler now uses `GoReflaxeCompiler`, but external macro
	callers may still invoke these public helpers and must receive the same output
	confinement guarantees.

	How: Every generated file, module file, and runtime copy is delegated to
	`GoGeneratedOutputBoundary`; this class performs no direct filesystem write.
**/
class GoOutputIterator {
	public function new() {}

	#if macro
	public static function writeFiles(outputDir:String, files:Array<GoGeneratedFile>):Void {
		var boundary = new GoGeneratedOutputBoundary(outputDir);
		for (file in files) {
			boundary.saveContent(file.relativePath, file.contents);
		}
		copyGeneratedLicenseMaterial(boundary);
	}

	public static function writeGoMod(outputDir:String, moduleName:String):Void {
		var goMod = ["module " + moduleName, "", "go 1.22", ""].join("\n");
		new GoGeneratedOutputBoundary(outputDir).saveContent("go.mod", goMod);
	}

	public static function copyRuntime(outputDir:String):Void {
		var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
		if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
			Context.fatalError("Missing packaged hxrt runtime directory", Context.currentPos());
		}

		copyDirectory(new GoGeneratedOutputBoundary(outputDir), runtimeSource, "hxrt");
	}

	static function copyDirectory(boundary:GoGeneratedOutputBoundary, source:String, targetRelative:String):Void {
		for (entry in FileSystem.readDirectory(source)) {
			var sourcePath = Path.join([source, entry]);
			var targetPath = Path.join([targetRelative, entry]);
			if (FileSystem.isDirectory(sourcePath)) {
				copyDirectory(boundary, sourcePath, targetPath);
			} else {
				boundary.copyFile(sourcePath, targetPath);
			}
		}
	}

	/**
		What: Mirrors the registered compiler's generated-project license output.

		Why: Legacy macro callers must receive the same redistribution material as the
		main compiler path instead of silently producing an under-documented tree.

		How: Copy both packaged notices through the confined direct-write boundary.
	**/
	static function copyGeneratedLicenseMaterial(boundary:GoGeneratedOutputBoundary):Void {
		var libraryRoot = findLibraryRoot();
		boundary.copyFile(Path.join([libraryRoot, "licenses", "HAXE-GO-GENERATED-MIT.txt"]), "LICENSES/HAXE-GO-GENERATED-MIT.txt");
		boundary.copyFile(Path.join([libraryRoot, "licenses", "HAXE-STDLIB-MIT.txt"]), "LICENSES/HAXE-STDLIB-MIT.txt");
	}

	static function findLibraryRoot():String {
		var thisFile = Context.resolvePath("reflaxe/go/GoOutputIterator.hx");
		var srcDir = Path.normalize(Path.directory(thisFile));
		return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
	}
	#end
}
