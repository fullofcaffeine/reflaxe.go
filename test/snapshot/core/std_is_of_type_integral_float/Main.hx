/**
	`Std.isOfType(value, Int)` accepts a Float with an Int value (haxe_go-jxxd).

	`--interp` answers true for a Float that is integral and within the 32-bit
	range (expected.stdout is the `--interp` output). haxe.go answered false for
	every Float: a `Dynamic` value used a Go type switch over the integer types,
	and a value typed Float gave the constant false. The generated Go now calls
	`hxrt.IsInt` for a `Dynamic`, `Any` or `Null<Float>` value and
	`hxrt.FloatIsInt` for a value typed Float. The runtime `Int` token test
	calls `hxrt.IsInt` too.
**/
class Main {
	static function main() {
		var values:Array<Dynamic> = [4.0, 4.5, 2147483648.0, 7];
		for (value in values)
			Sys.println(value + " isInt=" + Std.isOfType(value, Int) + " typeof=" + Type.typeof(value));
		var typed:Float = 4.0;
		var optional:Null<Float> = 5.0;
		var any:Any = 6.0;
		var token:Dynamic = Int;
		Sys.println("Float=" + Std.isOfType(typed, Int) + " Null<Float>=" + Std.isOfType(optional, Int) + " Any=" + Std.isOfType(any, Int) + " token="
			+ Std.isOfType(7.0, token));
		var whole:Dynamic = 9.0;
		var castWhole:Int = cast(whole, Int);
		Sys.println("cast=" + castWhole);
	}
}
