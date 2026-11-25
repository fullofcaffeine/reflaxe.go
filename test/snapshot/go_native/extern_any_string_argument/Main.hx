/**
	Regression for haxe_go-ya0z: a Haxe string passed to a non-variadic Go
	`any` parameter must reach Go as a Go `string`, as it already did through
	`...any`. Before the fix, Go received the Haxe `*string` carrier, so
	text/template printed `*string` for `%T`, a pointer for `%v`, and failed on
	`eq . "42"` with "incompatible types for comparison".

	The expected output is what upstream Go code prints for the same calls with
	Go string values.
**/
@:go.import("io") @:go.name("Writer")
extern interface NativeWriter {}

@:go.import("os")
extern class NativeOs {
	@:go.name("Stdout") public static var stdout:NativeWriter;
}

@:go.import("text/template") @:go.name("Template")
extern class NativeTemplate {
	@:go.name("Parse") @:go.valueError public function parse(text:String):go.Result<NativeTemplate>;
	@:go.name("Execute") public function execute(writer:NativeWriter, data:Dynamic):go.NativeError;
	@:go.name("Execute") public function executeAny(writer:NativeWriter, data:Any):go.NativeError;
}

@:go.import("text/template")
extern class NativeTemplates {
	@:go.name("New") public static function create(name:String):NativeTemplate;
}

@:go.import("sync/atomic") @:go.name("Value") @:go.struct
extern class NativeAtomicValue {
	public function new();
	@:go.name("Store") public function store(value:Dynamic):Void;
	@:go.name("Load") public function load():Dynamic;
}

/**
	Go names on an extern abstract are not a Go type: its values are Haxe
	strings (see GoTypeMapper.hasOrdinaryStringBacking), so it crosses as a Go
	`string`. A port that needs `json.Number` itself declares a
	`@:go.valueType extern class`.
**/
@:go.import("encoding/json") @:go.name("Number")
extern abstract JsonNumber(String) {}

abstract Label(String) from String {}

class Main {
	static function compile(source:String):NativeTemplate {
		return NativeTemplates.create("t").parse(source).unwrap();
	}

	static function report(error:Null<go.NativeError>):Void {
		if (error != null)
			Sys.println("error: " + error.text());
	}

	static function text():String {
		return "call";
	}

	static function main():Void {
		final kind = compile('{{printf "%T|%v|%-6s|" . . .}}{{if eq . "42"}}eq{{else}}ne{{end}}\n');
		final plain = "plain";
		final number:JsonNumber = cast "42";
		final label:Label = "label";
		final missing:Null<String> = null;
		report(kind.execute(NativeOs.stdout, plain));
		report(kind.execute(NativeOs.stdout, number));
		report(kind.execute(NativeOs.stdout, label));
		report(kind.execute(NativeOs.stdout, "literal"));
		report(kind.execute(NativeOs.stdout, text()));
		// A null Haxe String crosses as the text "null", as it does for a Go
		// string parameter and for `...any`.
		report(kind.execute(NativeOs.stdout, missing));
		report(kind.executeAny(NativeOs.stdout, number));

		// Non-string scalars already were native Go values; keep them so.
		final scalar = compile('{{printf "%T|%v" . .}}\n');
		report(scalar.execute(NativeOs.stdout, 7));
		report(scalar.execute(NativeOs.stdout, 2.5));
		report(scalar.execute(NativeOs.stdout, true));

		// A Go string that comes back through `any` is still a Haxe String.
		final box = new NativeAtomicValue();
		box.store(plain);
		final back:String = box.load();
		Sys.println(back + "/" + back.length + "/" + (back == "plain"));
	}
}
