package app.runtime;

class BuildConfig {
	#if fluxproxy_profile_portable
	public static inline final PROFILE = "portable";
	#elseif fluxproxy_profile_metal
	public static inline final PROFILE = "metal";
	#else
	public static inline final PROFILE = "portable";
	#end

	#if fluxproxy_variant_go_native
	public static inline final VARIANT = "go_native";
	public static inline final DISPATCH_WORKERS = 3;
	#else
	public static inline final VARIANT = "core";
	public static inline final DISPATCH_WORKERS = 1;
	#end

	public static inline final INGEST_QUEUE_CAPACITY = 3;
	public static inline final PER_ROUTE_LIMIT = 2;
	public static inline final BREAKER_FAILURE_THRESHOLD = 2;
	public static inline final TIMEOUT_MS = 50;
}
