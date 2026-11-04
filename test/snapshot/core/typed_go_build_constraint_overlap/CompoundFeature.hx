@:go.import("image")
@:go.name("Point")
@:go.struct
extern class CompoundFeaturePoint {
	public function new();
	@:go.name("X") public var x:Int;
}

@:goBuildConstraint("(alpha || beta) && !gamma")
@:keep
class CompoundFeature {
	static final installed = check();

	public function new() {}

	static function check():Bool {
		var point = new CompoundFeaturePoint();
		point.x = 42;
		var native:(value:CompoundFeaturePoint) -> Int = value -> value.x;
		var erased:Dynamic = native;
		var recovered:(value:Dynamic) -> Int = erased;
		if (recovered(point) != 42)
			throw "native recovery failed";
		if (erased(point) != 42)
			throw "native erased invocation failed";
		var source:(callback:(value:CompoundFeature) -> String) -> String = callback -> callback(new CompoundFeature());
		var stored:Dynamic = source;
		var nested:(callback:(value:Dynamic) -> String) -> String = stored;
		var calls = 0;
		var result = nested(value -> {
			calls++;
			return "ok";
		});
		if (result != "ok" || calls != 1)
			throw "nested callback recovery changed evaluation";
		return Registry.install();
	}
}
