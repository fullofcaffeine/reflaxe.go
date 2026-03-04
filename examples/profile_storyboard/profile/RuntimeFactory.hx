package profile;

class RuntimeFactory {
	public static function create():StoryboardRuntime {
		return new PortableRuntime();
	}
}
