import reflaxe.std.Result;

class Main {
	static function main():Void {
		final portable:Result<Int, String> = Ok(7);
		final native:go.Result<Int> = portable;
		trace(native);
	}
}
