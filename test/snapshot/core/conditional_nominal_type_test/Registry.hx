class Registry {
	public static var count:Int = 0;

	public static function install():Bool {
		count++;
		return true;
	}
}
