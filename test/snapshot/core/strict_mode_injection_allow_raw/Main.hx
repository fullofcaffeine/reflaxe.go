@:goAllowRaw
class Main {
	static function main() {
		var base = 1;
		var direct:Int = cast untyped __go__("2");
		var interpolated:Int = cast reflaxe.go.macros.GoInjection.__go__("{0} + 2", base);
		Sys.println(direct);
		Sys.println(interpolated);
	}
}
