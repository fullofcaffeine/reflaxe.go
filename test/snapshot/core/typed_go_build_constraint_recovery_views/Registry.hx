class Registry {
	public static var selected:String = "pending";

	public static function install(value:String):Bool {
		selected = value;
		return true;
	}
}
