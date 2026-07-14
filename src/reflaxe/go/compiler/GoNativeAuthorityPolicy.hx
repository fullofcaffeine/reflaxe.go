package reflaxe.go.compiler;

/**
	Why
	Native API admission is a policy decision orthogonal to source semantics and
	lowering quality.

	What
	Controls whether unscoped typed `go.*` usage is diagnosed or accepted.

	How
	`Guarded` applies the configured native-import diagnostic outside explicit
	native modules; `Explicit` accepts typed native APIs without that guard.
**/
enum abstract GoNativeAuthorityPolicy(String) from String to String {
	var Guarded = "guarded";
	var Explicit = "explicit";

	public inline function label():String {
		return this;
	}

	public static inline function defaultFor(preset:GoPolicyPreset):GoNativeAuthorityPolicy {
		return preset == GoPolicyPreset.MetalCompatibility ? Explicit : Guarded;
	}
}
