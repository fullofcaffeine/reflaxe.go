package reflaxe.go.compiler;

/** Stable, path-free categories for existing-module configuration errors. */
enum abstract GoProjectModeErrorKind(String) to String {
	final InvalidManifest = "GO-EXISTING-MODULE-MANIFEST";
	final InvalidModuleFile = "GO-EXISTING-MODULE-FILE";
	final ConfigurationConflict = "GO-EXISTING-MODULE-CONFLICT";
	final UnsupportedProjectShape = "GO-EXISTING-MODULE-UNSUPPORTED";
	final InvalidPackageName = "GO-PACKAGE-NAME";
	final InvalidPackageDirectory = "GO-PACKAGE-DIR";
	final PackageMismatch = "GO-PACKAGE-MISMATCH";
	final EntrypointOwnership = "GO-ENTRYPOINT-OWNERSHIP";
	final InvalidBuildTarget = "GO-BUILD-TARGET";
	final InvalidBuildOutput = "GO-BUILD-OUTPUT";
	final InvalidBuildTag = "GO-BUILD-TAG";
	final InvalidLinkerArgument = "GO-BUILD-LDFLAG";
	final InvalidBuildArgument = "GO-BUILD-ARGUMENT";
	final InvalidBuildEnvironment = "GO-BUILD-ENVIRONMENT";
}

/** A project error that can safely appear in portable compiler diagnostics. */
class GoProjectModeError extends haxe.Exception {
	public final kind:GoProjectModeErrorKind;

	public function new(kind:GoProjectModeErrorKind, explanation:String) {
		this.kind = kind;
		super('[${kind}] Existing Go module configuration is invalid: ${explanation}.');
	}
}
