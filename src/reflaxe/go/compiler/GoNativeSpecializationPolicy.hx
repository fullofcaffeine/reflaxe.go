package reflaxe.go.compiler;

/**
	Why
	Attempting a typed Go representation is an optimization policy; it must not
	redefine the source semantics selected by typed APIs and module boundaries.

	What
	Controls when typed native specialization is attempted.

	How
	`Proven` uses capability planners and enabled portable fast paths; `Eager`
	attempts supported typed `go.*` specializations at every eligible call site.
**/
enum abstract GoNativeSpecializationPolicy(String) from String to String {
	var Proven = "proven";
	var Eager = "eager";

	public inline function label():String {
		return this;
	}

	public static inline function defaultFor(preset:GoPolicyPreset):GoNativeSpecializationPolicy {
		return preset == GoPolicyPreset.MetalCompatibility ? Eager : Proven;
	}
}
