@:goBuildConstraint("cgo")
@:keep
final class CgoFeature extends FeatureBase {
	static final installed:Bool = Registry.install("cgo");

	public var mode:String = "cgo";

	public function new() {
		super();
	}

	public static function selected():String {
		return "cgo";
	}

	public override function label():String {
		return mode;
	}
}
