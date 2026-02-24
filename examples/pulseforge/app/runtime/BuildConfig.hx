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
	public static inline final INGEST_QUEUE_CAPACITY = 3;
	public static inline final PARSE_WORKERS = 3;
	public static inline final ENRICH_WORKERS = 2;
	#else
	public static inline final VARIANT = "core";
	public static inline final INGEST_QUEUE_CAPACITY = 3;
	public static inline final PARSE_WORKERS = 1;
	public static inline final ENRICH_WORKERS = 1;
	#end

	public static inline final ALERT_WEIGHTED_THRESHOLD = 20;
}
