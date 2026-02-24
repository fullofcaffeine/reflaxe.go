package app.runtime;

class RuntimeFactory {
	public static function create():FluxRuntime {
		#if fluxproxy_variant_go_native
		return new GoNativeRuntime();
		#else
		return new CoreRuntime();
		#end
	}
}
