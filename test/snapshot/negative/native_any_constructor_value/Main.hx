import go.NativeAny;

/** `go.NativeAny` constructors are compiler-lowered conversions, not functions. */
class Main {
	static function main():Void {
		final convert:String->NativeAny = NativeAny.ofString;
		convert("text");
	}
}
