package app.runtime;

class RuntimeFactory {
	public static function create():PulseRuntime {
		#if pulseforge_variant_go_native
		return new GoNativeRuntime();
		#else
		return new CoreRuntime();
		#end
	}
}
