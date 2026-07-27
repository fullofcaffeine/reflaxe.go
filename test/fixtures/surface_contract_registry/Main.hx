import reflaxe.std.Option;
import reflaxe.std.Result;

class Main {
	static function main() {
		final values:Array<Int> = [1, 2, 3];
		final label:String = values.join(",");
		final bytes = haxe.io.Bytes.ofString(label);
		final option:Option<Int> = Some(values.length);
		final result:Result<Int, String> = Ok(values.length);
		final dynamicOption:Option<Dynamic> = Some(label);
		final dynamicResult:Result<Int, Dynamic> = Err(label);
		trace(option);
		trace(result);
		trace(dynamicOption);
		trace(dynamicResult);
		trace(bytes.length);
	}
}
