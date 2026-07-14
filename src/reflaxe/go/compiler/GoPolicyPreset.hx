package reflaxe.go.compiler;

import reflaxe.go.GoProfile;

/**
	Why
	The legacy `portable|metal` selector must remain source-compatible without
	implying that haxe.go has two semantic backends.

	What
	Names the convenience preset used only to supply defaults for orthogonal
	build policies.

	How
	`GoBuildContextResolver` maps the legacy profile to one preset, then resolves
	each policy axis independently so explicit axis defines can override it.
**/
enum abstract GoPolicyPreset(String) from String to String {
	var PortableDefault = "portable_default";
	var MetalCompatibility = "metal_compatibility";

	public inline function label():String {
		return this;
	}

	public static inline function fromLegacyProfile(profile:GoProfile):GoPolicyPreset {
		return profile == GoProfile.Metal ? MetalCompatibility : PortableDefault;
	}
}
