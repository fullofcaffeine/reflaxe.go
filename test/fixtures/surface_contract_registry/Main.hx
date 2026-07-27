import reflaxe.std.Option;
import reflaxe.std.Result;

typedef HiddenDynamic = Dynamic;
abstract HiddenDynamicAbstract(Dynamic) from Dynamic to Dynamic {}

class Main {
	static function nestedLength(values:Array<Array<Int>>):Int {
		return values[1].length;
	}

	static function dynamicLabel(values:Array<Dynamic>):String {
		return Std.string(values[1]);
	}

	static function hiddenTypedefLabel(values:Array<HiddenDynamic>):String {
		return Std.string(values[0]);
	}

	static function nestedHiddenTypedefLabel(values:Array<Array<HiddenDynamic>>):String {
		return Std.string(values.length);
	}

	static function hiddenAbstractLabel(values:Array<HiddenDynamicAbstract>):String {
		return Std.string(values[0]);
	}

	static function main() {
		final values:Array<Int> = [1, 2, 3];
		final nestedValues:Array<Array<Int>> = [[1], [2, 3]];
		final dynamicValues:Array<Dynamic> = [1, "two"];
		final hiddenTypedefValues:Array<HiddenDynamic> = ["typedef"];
		final nestedHiddenTypedefValues:Array<Array<HiddenDynamic>> = [["nested-typedef"]];
		final hiddenAbstractValues:Array<HiddenDynamicAbstract> = [cast "abstract"];
		final label:String = values.join(",") + ":" + nestedLength(nestedValues) + ":" + dynamicLabel(dynamicValues) + ":"
			+ hiddenTypedefLabel(hiddenTypedefValues) + ":" + nestedHiddenTypedefLabel(nestedHiddenTypedefValues) + ":"
			+ hiddenAbstractLabel(hiddenAbstractValues);
		final bytes = haxe.io.Bytes.ofString(label);
		final option:Option<Int> = Some(values.length);
		final result:Result<Int, String> = Ok(values.length);
		final dynamicOption:Option<Dynamic> = Some(label);
		final dynamicResult:Result<Int, Dynamic> = Err(label);
		trace(nestedValues.length);
		trace(dynamicValues.length);
		trace(option);
		trace(result);
		trace(dynamicOption);
		trace(dynamicResult);
		trace(bytes.length);
	}
}
