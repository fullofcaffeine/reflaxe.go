import haxe.io.Path;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoSourceFormatter;
import sys.FileSystem;
import sys.io.File;

/**
	Contract for formatting many generated Go sources in one `gofmt` run.

	Why
	`GoSourceFormatter.formatAll` formats the generated sources of an
	existing-module compile in one `gofmt -w` run over files in a temporary
	directory, instead of one `gofmt` process per source. Its results replace
	those of `format` (standard input) in the output, so they must be identical,
	and any case it cannot format exactly must answer null, so that `format`
	reports the failure exactly as before.

	What
	Pins that `formatAll` returns, in order, exactly what `format` returns for
	each source, including sources `gofmt` changes and long-line layout; that it
	returns null for a source that is only a fragment (which standard input
	accepts), for a rejected source, when it cannot make its directory, and when
	`gofmt` is missing; and that it leaves no files behind.

	How
	Runs under `--interp` with `TMPDIR` pointed at a fresh directory of its own,
	so it can check that the directory is empty after each call.
**/
class GoSourceFormatterContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw "GoSourceFormatterContract: " + message;
	}

	static function same(left:Null<Array<String>>, right:Array<String>):Bool {
		if (left == null || left.length != right.length)
			return false;
		for (index in 0...right.length)
			if (left[index] != right[index])
				return false;
		return true;
	}

	static function formatEach(sources:Array<String>):Array<String> {
		return [for (source in sources) GoSourceFormatter.format(source)];
	}

	static function main():Void {
		final scratch = Path.join([Sys.getCwd(), ".cache", "go-source-formatter-contract-" + Std.random(0x3fffffff)]);
		FileSystem.createDirectory(scratch);
		final originalTemporary = Sys.getEnv("TMPDIR");
		Sys.putEnv("TMPDIR", scratch);
		try {
			run(scratch);
		} catch (error:haxe.Exception) {
			cleanUp(scratch, originalTemporary);
			throw error;
		}
		cleanUp(scratch, originalTemporary);
		Sys.println("GoSourceFormatterContract: ok");
	}

	static function cleanUp(scratch:String, originalTemporary:Null<String>):Void {
		Sys.putEnv("TMPDIR", originalTemporary == null ? "" : originalTemporary);
		final notADirectory = Path.join([scratch, "file"]);
		if (FileSystem.exists(notADirectory))
			FileSystem.deleteFile(notADirectory);
		FileSystem.deleteDirectory(scratch);
	}

	static function requireEmpty(scratch:String, after:String):Void {
		require(FileSystem.readDirectory(scratch).length == 0, "files left behind after " + after);
	}

	static function run(scratch:String):Void {
		final empty = GoSourceFormatter.formatAll([]);
		require(empty != null && empty.length == 0, "no sources must give no results");

		final sources = [
			"package main\n\nfunc main() {}\n",
			"package main\nimport \"fmt\"\nfunc  main( ) {\n\tif true{fmt.Println( \"x\" )}\n}\n",
			"package demo\n\nconst (\n\tA = 1 // one\n\tLongerName = 22 // twenty-two\n)\n\ntype T struct {\n\tA int // a\n\tVeryLongFieldName string // b\n}\n",
			"package main\nimport \"fmt\"\nfunc  main( ) {\n\tif true{fmt.Println( \"x\" )}\n}\n",
			"package text\n\nvar Greeting = \"héllo 世界\"\nvar   Other=1\n"
		];
		require(same(GoSourceFormatter.formatAll(sources), formatEach(sources)), "results differ from formatting each source alone");
		requireEmpty(scratch, "formatting");
		final changed = GoSourceFormatter.format(sources[1]);
		require(changed != sources[1], "the fixture must contain a source that gofmt changes");

		final fragment = "x := 1\nprintln(x)\n";
		require(GoSourceFormatter.format(fragment).length > 0, "standard input must accept a statement list");
		require(GoSourceFormatter.formatAll([sources[0], fragment]) == null, "a fragment must leave the decision to format");
		requireEmpty(scratch, "a fragment");

		final rejected = "package main\n\nfunc {\n";
		require(GoSourceFormatter.formatAll([sources[0], rejected, sources[1]]) == null, "a rejected source must give null");
		requireEmpty(scratch, "a rejected source");
		var message = "";
		try {
			GoSourceFormatter.format(rejected);
		} catch (error:GoOutputPathError) {
			message = error.message;
		}
		require(StringTools.startsWith(message, "[GO-OUTPUT-PATH-006] Refused generated output: gofmt failed: <standard input>:3:"),
			"format must still report the standard input diagnostic, got: " + message);

		// A temporary directory that cannot hold a new directory.
		final notADirectory = Path.join([scratch, "file"]);
		File.saveContent(notADirectory, "");
		Sys.putEnv("TMPDIR", notADirectory);
		require(GoSourceFormatter.formatAll(sources) == null, "an unusable temporary directory must give null");
		Sys.putEnv("TMPDIR", scratch);
		FileSystem.deleteFile(notADirectory);

		// No gofmt on the search path.
		final searchPath = Sys.getEnv("PATH");
		Sys.putEnv("PATH", scratch);
		final missing = GoSourceFormatter.formatAll(sources);
		Sys.putEnv("PATH", searchPath == null ? "" : searchPath);
		require(missing == null, "a missing gofmt must give null");
		requireEmpty(scratch, "a missing gofmt");
	}
}
