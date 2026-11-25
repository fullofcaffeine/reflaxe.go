import go.NativeAny;
import go.NativeAnyMap;
import go.NativeSlice;
import go.NativeStringSlice;

/**
	Contract for haxe_go-ya0z: `go.NativeAnyMap` builds the native Go
	`map[string]any` that template- and reflection-heavy Go APIs expect, and
	`go.NativeAny` puts native Go values in it.

	The data root is upstream bd's `list --format` root
	(cmd/bd/list_output.go:134-140). The expected output is what upstream Go
	code prints for the same `map[string]interface{}` of native values: printf
	padding, `eq` against a string constant, nested field access, `range`,
	`index`, `len`, `%T`, missing keys and encoding/json all agree.
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

class Main {
	static final trail = new Array<String>();

	static function note(step:String):String {
		trail.push(step);
		return step;
	}

	static function target(map:NativeAnyMap):NativeAnyMap {
		trail.push("map");
		return map;
	}

	static function issue(id:String, title:String, priority:Int, labels:Array<String>):NativeAnyMap {
		final data = new NativeAnyMap();
		data.set("ID", id);
		data.set("Title", title);
		data.set("Priority", priority);
		data.set("Labels", NativeStringSlice.fromArray(labels));
		return data;
	}

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
		final parent = issue("bd-1", "Fix the parser", 1, ["parser", "p1"]);
		final dependency = new NativeAnyMap();
		dependency.set("IssueID", "bd-1");
		dependency.set("DependsOnID", "bd-22");
		dependency.set("Type", "blocks");
		final data = new NativeAnyMap();
		data.set("IssueID", "bd-1");
		data.set("DependsOnID", "bd-22");
		data.set("Type", "blocks");
		data.set("Issue", parent);
		data.set("Dependency", dependency);
		// An integral Float stays a Go float64.
		data.set("Score", NativeAny.ofFloat(2));
		data.set("Ready", true);
		data.set("Missing", NativeAny.nil());
		final nested:Array<NativeAny> = ["nested"];
		final items:Array<NativeAny> = ["one", 2, 3.5, false, null, NativeSlice.fromArray(nested)];
		data.set("Items", NativeSlice.fromArray(items));

		render("{{.IssueID}} {{.DependsOnID}}", data);
		render('{{printf "%-14s" .IssueID}}|{{if eq .Type "blocks"}}blocks{{else}}other{{end}}|{{.Issue.Title}}|{{.Dependency.DependsOnID}}', data);
		render('{{range $$i, $$label := .Issue.Labels}}{{if $$i}},{{end}}{{$$label}}{{end}} {{len .Issue.Labels}} {{index .Issue.Labels 1}}', data);
		render('{{printf "%T %T %T %T %T %T %T %T %T" . .IssueID .Issue.Priority .Score .Ready .Missing .Issue.Labels .Items .Issue}}', data);
		render('{{range .Items}}[{{printf "%T" .}} {{.}}]{{end}}', data);
		render('{{.Missing}} {{.Score}} {{printf "%.1f" .Score}} {{.Issue.Priority | printf "P%d"}} {{if .Ready}}ready{{end}} {{.Nope}}', data);
		render("{{.Issue.Title.Missing}}", data);
		final encoder = NativeJson.newEncoder(NativeOs.stdout);
		final error = encoder.encode(data);
		if (error != null)
			Sys.println("json: " + error.text());

		// `set` runs the map, the key and the value in source order, also when
		// the value needs statements of its own.
		final ordered = new NativeAnyMap();
		target(ordered).set(note("key"), if (trail.length > 0) note("value") else "other");
		Sys.println(trail.join(","));
		encoder.encode(ordered);
	}
}
