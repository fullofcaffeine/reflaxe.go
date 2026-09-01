final class Registry {
	public static var selected:String = "common";

	public static function install(value:String):Bool {
		selected = value;
		return true;
	}
}
