package reflaxe.go.compiler;

/**
	Why
	Effective policy without provenance would make compatibility precedence
	impossible to audit in reports and tests.

	What
	Typed result of resolving the compatibility preset and native policy axes.

	How
	The resolver constructs it once at compile start so macros, lowering, runtime
	planning, and reports consume identical values and sources.
**/
typedef GoPolicyResolution = {
	final preset:GoPolicyPreset;
	final semanticBoundarySource:GoSemanticBoundarySource;
	final nativeAuthority:GoNativeAuthorityPolicy;
	final nativeAuthoritySource:GoPolicyResolutionSource;
	final nativeSpecialization:GoNativeSpecializationPolicy;
	final nativeSpecializationSource:GoPolicyResolutionSource;
	final nativeFallback:GoNativeFallbackPolicy;
	final nativeFallbackSource:GoPolicyResolutionSource;
}
