package reflaxe.go.compiler;

/**
	Why
	Reports must explain precedence instead of exposing only an unexplained
	effective boolean.

	What
	Records which preset or compatibility/canonical define selected a policy.

	How
	The resolver assigns one source after validating explicit and legacy inputs;
	reports emit the stable label alongside the effective typed value.
**/
enum abstract GoPolicyResolutionSource(String) from String to String {
	var PolicyPreset = "policy_preset";
	var NativeAuthorityDefine = "reflaxe_go_native_authority";
	var NativeSpecializationDefine = "reflaxe_go_native_specialization";
	var NativeFallbackDefine = "reflaxe_go_native_fallback";
	var LegacyMetalFallbackDefine = "reflaxe_go_metal_allow_fallback";

	public inline function label():String {
		return this;
	}
}
