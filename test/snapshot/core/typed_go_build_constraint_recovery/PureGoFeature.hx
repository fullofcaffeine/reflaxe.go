@:goBuildConstraint("!cgo")
@:keep
class PureGoFeature {
	static final installed = Registry.install(check());

	public var label:String = "ok";

	public function new() {}

	static function check():String {
		var erased:Dynamic = new PureGoFeature();
		var restored:PureGoFeature = erased;
		var source:(value:PureGoFeature) -> String = value -> value.label;
		var callback:Dynamic = source;
		var recovered:(value:PureGoFeature) -> String = callback;
		return recovered(restored);
	}
}
