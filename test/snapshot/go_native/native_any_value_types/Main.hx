import go.NativeAny;
import go.NativeAnyMap;
import go.NativeSlice;

/**
	Contract for haxe_go-37b9 option (a): `go.NativeAny.ofValue` stores the Go
	value of a precise extern value type, so Go reflection APIs see the real Go
	type and its methods.

	`time.Time` is a `@:go.struct` extern, and `time.Month`, `time.Duration`
	and the named string type `regexp/syntax.ErrorCode` are `@:go.valueType`
	externs. Haxe keeps these externs pointer-backed (`*time.Time`), and a
	`Dynamic` value keeps that carrier, so `%T` printed `*time.Time`. The
	expected output is what upstream Go code prints for the same
	`map[string]interface{}` of Go values: method calls such as
	`.CreatedAt.Year` and `.Format`, `%T`, Stringer output, `eq` on a named
	string, a struct zero value and encoding/json all agree.
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
	@:go.name("Execute") public function execute(writer:NativeWriter, data:NativeAny):go.NativeError;
}

@:go.import("text/template")
extern class NativeTemplates {
	@:go.name("New") public static function create(name:String):NativeTemplate;
}

@:go.import("encoding/json") @:go.name("Encoder")
extern class NativeEncoder {
	@:go.name("Encode") public function encode(value:NativeAny):go.NativeError;
}

@:go.import("encoding/json")
extern class NativeJson {
	@:go.name("NewEncoder") public static function newEncoder(writer:NativeWriter):NativeEncoder;
}

@:go.import("time") @:go.name("Month") @:go.valueType
extern class GoMonth {}

@:go.import("time") @:go.name("Duration") @:go.valueType
extern class GoDuration {}

@:go.import("time") @:go.name("Time") @:go.struct
extern class GoTime {
	public function new();

	@:go.valueReturn @:go.name("UTC") public function utc():GoTime;
	@:go.valueReturn @:go.name("Month") public function month():GoMonth;
}

final class ParseDurationResult {
	public var value(default, null):GoDuration;
	public var error(default, null):Null<go.Error>;

	public function new(value:GoDuration, error:Null<go.Error>) {
		this.value = value;
		this.error = error;
	}
}

@:go.import("time")
extern class GoTimePkg {
	@:go.valueReturn @:go.name("Unix") public static function unix(seconds:Int, nanoseconds:Int):GoTime;

	@:go.tupleReturn @:go.tupleValueResults("0") @:go.name("ParseDuration")
	public static function parseDuration(value:String):ParseDurationResult;
}

/** A named Go string type: `type ErrorCode string`. */
@:go.import("regexp/syntax") @:go.name("ErrorCode") @:go.valueType
extern class SyntaxErrorCode {}

@:go.import("regexp/syntax")
extern class SyntaxPkg {
	/** A typed Go constant of the named string type. */
	@:go.name("ErrMissingBracket") @:go.valueField public static var errMissingBracket:SyntaxErrorCode;
}

class Main {
	static function render(source:String, data:NativeAnyMap):Void {
		final parsed = NativeTemplates.create("format").parse(source);
		if (!parsed.isOk()) {
			Sys.println("invalid format template: " + parsed.error());
			return;
		}
		final error = parsed.unwrap().execute(NativeOs.stdout, data);
		if (error != null) {
			Sys.println("template execution error: " + error.text());
			return;
		}
		Sys.println("");
	}

	static function main():Void {
		final created = GoTimePkg.unix(1795608000, 0).utc();
		final timeout = GoTimePkg.parseDuration("1m30s").value;
		final issue = new NativeAnyMap();
		issue.set("Title", "Fix the parser");
		issue.set("CreatedAt", NativeAny.ofValue(created));
		issue.set("Month", NativeAny.ofValue(created.month()));
		final data = new NativeAnyMap();
		data.set("Issue", issue);
		data.set("Timeout", NativeAny.ofValue(timeout));
		data.set("Code", NativeAny.ofValue(SyntaxPkg.errMissingBracket));
		// The Go zero value of a struct extern, stored by value.
		data.set("Zero", NativeAny.ofValue(new GoTime()));
		final values:Array<NativeAny> = [
			NativeAny.ofValue(created.month()),
			NativeAny.ofValue(SyntaxPkg.errMissingBracket)
		];
		data.set("Values", NativeSlice.fromArray(values));

		render('{{.Issue.CreatedAt.Year}} {{.Issue.CreatedAt.Format "2006-01-02"}} {{.Issue.CreatedAt}}', data);
		render('{{printf "%T %T %T %T %T" .Issue.CreatedAt .Issue.Month .Timeout .Code .Values}}', data);
		render('{{.Issue.Month}} {{.Timeout}} {{.Code}} {{if eq .Code "missing closing ]"}}eq{{end}} {{range .Values}}[{{printf "%T" .}}]{{end}} {{.Zero.IsZero}}',
			data);
		final error = NativeJson.newEncoder(NativeOs.stdout).encode(data);
		if (error != null)
			Sys.println("json: " + error.text());
	}
}
