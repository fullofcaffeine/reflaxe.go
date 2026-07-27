import reflaxe.std.Result;

class Main {
	static function main():Void {
		final native:go.Result<Int> = go.Go.ok(7);
		final portable:Result<Int, String> = native;
		trace(portable);
	}
}
