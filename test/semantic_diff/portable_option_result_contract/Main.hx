import reflaxe.std.Option;
import reflaxe.std.Result;

class PortableFailure {
	public final code:Int;
	public final message:String;

	public function new(code:Int, message:String) {
		this.code = code;
		this.message = message;
	}
}

class Main {
	static function optionMap<T, U>(value:Option<T>, transform:T->U):Option<U> {
		return switch value {
			case Some(item): Some(transform(item));
			case None: None;
		};
	}

	static function resultMap<T, U, E>(value:Result<T, E>, transform:T->U):Result<U, E> {
		return switch value {
			case Ok(item): Ok(transform(item));
			case Err(error): Err(error);
		};
	}

	static function resultMapError<T, E, F>(value:Result<T, E>, transform:E->F):Result<T, F> {
		return switch value {
			case Ok(item): Ok(item);
			case Err(error): Err(transform(error));
		};
	}

	static function renderOption(value:Option<Null<String>>):String {
		return switch value {
			case Some(item): item == null ? "some:null" : "some:" + item;
			case None: "none";
		};
	}

	static function renderPlainStringOption(value:Option<String>):String {
		return switch value {
			case Some(item): item == null ? "some:null" : "some:" + item;
			case None: "none";
		};
	}

	static function renderResult(value:Result<Int, PortableFailure>):String {
		return switch value {
			case Ok(item): "ok:" + item;
			case Err(error): "err:" + error.code + ":" + error.message;
		};
	}

	static function renderPlainStringError(value:Result<Int, String>):String {
		return switch value {
			case Ok(item): "ok:" + item;
			case Err(error): error == null ? "err:null" : "err:" + error;
		};
	}

	static function renderNested(value:Option<Result<Int, String>>):String {
		return switch value {
			case Some(Ok(item)): "some-ok:" + item;
			case Some(Err(error)): "some-err:" + error;
			case None: "none";
		};
	}

	static function main() {
		Sys.println("option.some=" + renderOption(Some("value")));
		Sys.println("option.some-null=" + renderOption(Some(null)));
		final plainStringOption:Option<String> = Some(null);
		Sys.println("option.plain-string-null=" + renderPlainStringOption(plainStringOption));
		Sys.println("option.none=" + renderOption(None));
		Sys.println("option.generic=" + renderOption(optionMap(Some("x"), value -> value + "!")));

		Sys.println("result.ok=" + renderResult(Ok(7)));
		Sys.println("result.err=" + renderResult(Err(new PortableFailure(41, "typed"))));
		final plainStringError:Result<Int, String> = Err(null);
		Sys.println("result.plain-string-null=" + renderPlainStringError(plainStringError));
		Sys.println("result.generic-ok=" + renderResult(resultMap(Ok(5), value -> value + 2)));
		Sys.println("result.generic-err=" + renderResult(resultMapError(Err("bad"), error -> new PortableFailure(42, error))));

		Sys.println("nested.ok=" + renderNested(Some(Ok(11))));
		Sys.println("nested.err=" + renderNested(Some(Err("oops"))));
		Sys.println("nested.none=" + renderNested(None));
	}
}
