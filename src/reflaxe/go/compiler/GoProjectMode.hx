package reflaxe.go.compiler;

/** Selects whether haxe.go owns a standalone module or joins a caller module. */
enum GoProjectMode {
	Standalone;
	ExistingModule(project:ExistingGoModuleProject);
}

/** The entry point requested by a typed existing-module manifest. */
enum GoEntrypointPolicy {
	CompilerMain;
	CallerBridge(symbol:String);
}

/** Build behavior supported by the first existing-module safety slice. */
enum GoBuildPolicy {
	NoBuild;
}

typedef ExistingGoModuleProjectData = {
	final manifestPath:String;
	final moduleRoot:String;
	final modulePath:String;
	final packageDir:String;
	final packageName:String;
	final runtimeDir:String;
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
	public final packageDir:String;
	public final packageName:String;
	public final runtimeDir:String;
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
}
