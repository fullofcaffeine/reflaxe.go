import go.Result;

@:go.import("strconv")
extern class StrconvPkg {
	@:go.name("Atoi")
	@:go.valueError
	static function atoi(value:String):Result<Int>;
}

@:go.import("net/url")
extern class URLPkg {
	@:go.name("QueryUnescape")
	@:go.valueError
	static function queryUnescape(value:String):Result<String>;
}

class Main {
	static function main() {
		var ok = StrconvPkg.atoi("42");
		Sys.println("ok.isOk=" + ok.isOk());
		Sys.println("ok.isErr=" + ok.isErr());
		Sys.println("ok.unwrap=" + ok.unwrap());

		var err = StrconvPkg.atoi("x");
		Sys.println("err.isOk=" + err.isOk());
		Sys.println("err.isErr=" + err.isErr());
		Sys.println("err.hasError=" + (err.error() != null));

		var text = URLPkg.queryUnescape("hello+world");
		Sys.println("text.unwrap=" + text.unwrap());
	}
}
