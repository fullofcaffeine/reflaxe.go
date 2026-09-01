@:goBuildConstraint("!cgo")
@:keep
final class PureGoFeature extends FeatureBase {
	static final installed:Bool = Registry.install("pure-go");

	public var mode:String = "pure-go";

	public function new() {
		super();
	}

	public static function selected():String {
		return "pure-go";
	}

	public override function label():String {
		return mode;
	}
}
