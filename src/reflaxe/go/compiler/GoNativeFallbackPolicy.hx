package reflaxe.go.compiler;

/**
	Why
	Failure to prove a native representation needs an explicit response that is
	independent from both source semantics and specialization eagerness.

	What
	Selects whether a failed typed native specialization may fall back or must
	fail compilation.

	How
	`Allow` preserves the semantics-safe representation; `Error` emits a
	compile-time diagnostic for user-owned fallback sites.
**/
enum abstract GoNativeFallbackPolicy(String) from String to String {
	var Allow = "allow";
	var Error = "error";

	public inline function label():String {
		return this;
	}

	public static inline function defaultFor(preset:GoPolicyPreset):GoNativeFallbackPolicy {
		return preset == GoPolicyPreset.MetalCompatibility ? Error : Allow;
	}
}
