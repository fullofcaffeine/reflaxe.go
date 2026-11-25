import go.NativeAny;
import go.NativeAnyMap;

/** `go.NativeAnyMap.set` lowers to a Go map assignment, not a method. */
class Main {
	static function main():Void {
		final data = new NativeAnyMap();
		final store:(String, NativeAny) -> Void = data.set;
		store("Title", "text");
	}
}
