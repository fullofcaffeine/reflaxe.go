package profile;

class RuntimeFactory {
	public static function create():StoryboardRuntime {
		#if example_profile_metal
		return new MetalRuntime();
		#else
		return new PortableRuntime();
		#end
	}
}
