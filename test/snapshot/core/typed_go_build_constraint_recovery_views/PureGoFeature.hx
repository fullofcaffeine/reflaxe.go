@:goBuildConstraint("!cgo")
@:keep
class PureGoFeature extends FeatureBase {
	static final installed = Registry.install(check());

	public function new() {
		super();
	}

	static function check():String {
		var erased:Dynamic = new PureGoFeature();
		var restored:FeatureBase = erased;
		var view:{var label:String;} = erased;
		view.label = "changed";
		if (restored.label != "changed")
			throw "structural recovery copied source";
		restored.label = "ok";
		if (view.label != "ok")
			throw "base recovery copied source";
		var source:(value:PureGoFeature) -> String = value -> value.label;
		var callback:Dynamic = source;
		var recovered:(value:Dynamic) -> String = callback;
		if (recovered(erased) != "ok")
			throw "conditional callable projection failed";
		if (callback(erased) != "ok")
			throw "conditional erased invocation failed";
		return restored.label;
	}
}
