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

	static function findLibraryRoot():String {
		var thisFile = Context.resolvePath("reflaxe/go/GoOutputIterator.hx");
		var srcDir = Path.normalize(Path.directory(thisFile));
		return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
	}
	#end
}
