package reflaxe.go.compiler;

import reflaxe.go.ast.GoPackageName;

/** Selects whether haxe.go owns a standalone module or joins a caller module. */
enum GoProjectMode {
	Standalone;
	ExistingModule(project:ExistingGoModuleProject);
}

/** The entry point requested by a typed existing-module manifest. */
enum GoEntrypointPolicy {
	CompilerMain;
	CallerBridge(symbol:GoEntrypointSymbol);
}

/** A validated Go function identifier used for the generated process entry. */
abstract GoEntrypointSymbol(String) {
	private inline function new(value:String) {
		this = value;
	}

	public static function named(value:String):GoEntrypointSymbol {
		if (!GoPackageName.isIdentifier(value)) {
			throw "invalid Go entrypoint symbol";
		}
		return new GoEntrypointSymbol(value);
	}

	public inline function value():String {
		return this;
	}
}

/** A module-root-relative path validated by the project manifest boundary. */
abstract GoProjectRelativePath(String) {
	private inline function new(value:String) {
		this = value;
	}

	@:allow(reflaxe.go.compiler.GoProjectModeResolver)
	private static function validated(value:String):GoProjectRelativePath {
		return new GoProjectRelativePath(value);
	}

	public inline function value():String {
		return this;
	}

	public inline function resolve(child:String):String {
		return this == "." ? child : this + "/" + child;
	}
}

/** Build behavior requested by an existing-module manifest. */
enum GoBuildPolicy {
	NoBuild;
	GoBuild(request:GoBuildRequest);
}

typedef ExistingGoModuleProjectData = {
	final manifestPath:String;
	final moduleRoot:String;
	final modulePath:String;
	final packageDir:GoProjectRelativePath;
	final packageName:GoPackageName;
	final runtimeDir:GoProjectRelativePath;
	final entrypoint:GoEntrypointPolicy;
	final build:GoBuildPolicy;
}

/**
	Validated inputs for one package inside a caller-owned Go module.

	The constructor accepts a named record so path and policy values cannot be
	accidentally exchanged at a long positional call site.
**/
class ExistingGoModuleProject {
	public final manifestPath:String;
	public final moduleRoot:String;
	public final modulePath:String;
	public final packageDir:GoProjectRelativePath;
	public final packageName:GoPackageName;
	public final runtimeDir:GoProjectRelativePath;
	public final entrypoint:GoEntrypointPolicy;
	public final build:GoBuildPolicy;

	public function new(data:ExistingGoModuleProjectData) {
		manifestPath = data.manifestPath;
		moduleRoot = data.moduleRoot;
		modulePath = data.modulePath;
		packageDir = data.packageDir;
		packageName = data.packageName;
		runtimeDir = data.runtimeDir;
		entrypoint = data.entrypoint;
		build = data.build;
	}

	/** Return one generated source path relative to the caller module root. */
	public inline function generatedSourcePath(fileName:String):String {
		return packageDir.resolve(fileName);
	}

	/** Return one runtime support path relative to the caller module root. */
	public inline function runtimePath(fileName:String):String {
		return runtimeDir.resolve(fileName);
	}

	/** Return the import path for the compiler-owned runtime package. */
	public function runtimeImportPath():String {
		return runtimeDir.value() == "." ? modulePath : modulePath + "/" + runtimeDir.value();
	}
}
