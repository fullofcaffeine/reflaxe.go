package app.runtime;

class BuildConfig {
	#if pulseforge_profile_portable
	public static inline final PROFILE = "portable";
	#elseif pulseforge_profile_metal
	public static inline final PROFILE = "metal";
	#else
	public static inline final PROFILE = "portable";
	#end

	#if pulseforge_variant_go_native
	public static inline final VARIANT = "go_native";
	#else
	public static inline final VARIANT = "core";
	#end
}
