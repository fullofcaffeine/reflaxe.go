@:goBuildConstraint("!cgo")
@:keep
final class PureGoFeature {
	public static function selected():String {
		return "pure-go";
	}
}
