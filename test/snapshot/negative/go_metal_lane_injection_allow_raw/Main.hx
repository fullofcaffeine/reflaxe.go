@:goAllowRaw
@:goMetal
class Main {
	static function main() {
		var value:Int = cast untyped __go__("1");
		Sys.println(value);
	}
}
