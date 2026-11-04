/** fmt accepts arbitrary native values; keep the erased variadic API at this extern boundary. */
@:go.import("fmt")
@:go.name("fmt")
extern class NativeFormat {
	@:go.name("Sprintf")
	static function format(pattern:String, values:haxe.Rest<Dynamic>):String;
}

class Main {
	static var evaluations = 0;

	static function next():Int {
		evaluations++;
		return evaluations;
	}

	static function main():Void {
		final record = {number: 7};
		Sys.println(NativeFormat.format("%d", 3));
		Sys.println(NativeFormat.format("%d", record.number));
		final local = 11;
		Sys.println(NativeFormat.format("%d", local));
		Sys.println(NativeFormat.format("%d:%d", next(), next()));
		Sys.println("evaluations=" + evaluations);
		Sys.println(NativeFormat.format("empty"));
	}
}
