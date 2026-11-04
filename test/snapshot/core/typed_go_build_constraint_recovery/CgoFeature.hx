@:goBuildConstraint("cgo")
@:keep
class CgoFeature {
	static final installed = Registry.install(check());

	public var label:String = "ok";

	public function new() {}

	static function check():String {
		var erased:Dynamic = new CgoFeature();
		var restored:CgoFeature = erased;
		var source:(value:CgoFeature) -> String = value -> value.label;
		var callback:Dynamic = source;
		var recovered:(value:CgoFeature) -> String = callback;
		return recovered(restored);
	}
}
