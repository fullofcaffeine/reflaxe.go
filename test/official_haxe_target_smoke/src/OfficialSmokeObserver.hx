/** Target-only output observer; it does not contribute portable test evidence. */
@:goNative
class OfficialSmokeObserver {
	public static function line(value:String):Void {
		OfficialNativeFmt.println(value);
	}
}

/** Typed model of the framework display boundary consumed by the observer. */
@:go.import("hxrt")
@:go.package("hxrt")
private extern class OfficialNativeFmt {
	@:go.name("Println")
	public static function println(value:Dynamic):Void;
}
