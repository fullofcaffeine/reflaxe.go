import CgoFeature;
import CgoMode;
import PureGoFeature;
import PureGoMode;

class Main {
	static function main():Void {
		final value:Dynamic = "value";
		final target:Dynamic = String;
		if (!Std.isOfType(value, target))
			throw "dynamic type lookup failed";
		if (Std.isOfType(value, FeatureBase))
			throw "constrained subclass leaked into common type lookup";
		final feature:Dynamic = new FeatureBase();
		if (Reflect.field(feature, "label") == null || Reflect.fields(feature).indexOf("name") < 0)
			throw "common reflection metadata is missing";
		if (Type.resolveClass("Registry") == null)
			throw "unconstrained type metadata is missing";
		if (Registry.selected == "common")
			throw "build-constrained installer did not run";
	}
}
