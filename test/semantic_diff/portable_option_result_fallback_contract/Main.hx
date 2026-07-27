import reflaxe.std.Option;
import reflaxe.std.Result;

class Main {
	static function renderOption(value:Option<Dynamic>):String {
		return switch value {
			case Some(item): item == null ? "some:null" : "some:" + Std.string(item);
			case None: "none";
		};
	}

	static function renderResult(value:Result<Int, Dynamic>):String {
		return switch value {
			case Ok(item): "ok:" + item;
			case Err(error): "err:" + Std.string(error);
		};
	}

	static function main() {
		final dynamicError:Dynamic = "dynamic:9";
		Sys.println("fallback.option-value=" + renderOption(Some(12)));
		Sys.println("fallback.option-null=" + renderOption(Some(null)));
		Sys.println("fallback.option-none=" + renderOption(None));
		Sys.println("fallback.result-ok=" + renderResult(Ok(8)));
		Sys.println("fallback.result-err=" + renderResult(Err(dynamicError)));
	}
}
