package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Compiler;
import haxe.macro.Context;
import sys.FileSystem;
#end

/**
	What: validates or supplies the vendored Reflaxe framework before compiler initialization.

	Why: source builds declare std override precedence in `haxe_libraries/reflaxe.go.hxml`,
	so changing compiler configuration after macro startup would be too late and unsupported.
	Direct `extraParams.hxml` consumers still need a typed fallback for the vendored framework.

	How: source-checkout library HXML places every precedence-sensitive path on the initial
	classpath. This bootstrap only appends the non-conflicting vendored Reflaxe path when the
	local `reflaxe` library was not loaded, accepts a release package where Reflaxe is already
	flattened into `src`, and otherwise emits a typed fatal diagnostic.
**/
class CompilerBootstrap {
	#if macro
	static var bootstrapped:Bool = false;

	public static function Start():Void {
		if (bootstrapped) {
			return;
		}
		bootstrapped = true;

		var root = findLibraryRoot();
		var vendoredReflaxe = Path.normalize(Path.join([root, "vendor", "reflaxe", "src"]));
		if (FileSystem.exists(vendoredReflaxe) && FileSystem.isDirectory(vendoredReflaxe)) {
			if (!Context.defined("reflaxe")) {
				Compiler.addClassPath(vendoredReflaxe);
			}
			return;
		}

		var packagedReflaxe = Path.normalize(Path.join([root, "src", "reflaxe", "ReflectCompiler.hx"]));
		if (FileSystem.exists(packagedReflaxe) && !FileSystem.isDirectory(packagedReflaxe)) {
			return;
		}

		Context.fatalError("Reflaxe.Go could not resolve its vendored Reflaxe framework. Use the checked-in library HXML or a package staged with vendored Reflaxe sources.",
			Context.currentPos());
	}

	static function findLibraryRoot():String {
		var thisFile = Context.resolvePath("reflaxe/go/CompilerBootstrap.hx");
		var srcDir = Path.normalize(Path.directory(thisFile));
		return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
	}
	#else
	public static function Start():Void {}
	#end
}
