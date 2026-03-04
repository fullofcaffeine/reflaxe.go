package profile;

class RuntimeFactory {
	public static function create():TodoRuntime {
		return new PortableRuntime();
	}
}
