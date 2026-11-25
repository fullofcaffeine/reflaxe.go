package reflaxe.go.compiler;

#if (macro || eval)
import haxe.io.Bytes;
import haxe.io.Path;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import sys.FileSystem;
import sys.io.File;
import sys.io.Process;

/**
	Canonicalizes compiler-owned Go source before exact output bytes are recorded.

	Why
	`gofmt` has context-sensitive layout rules, including line-length-sensitive
	spacing. Reimplementing those rules in the AST printer would create a second,
	incomplete Go formatter and make ownership hashes unstable under normal tooling.

	What
	Runs the canonical Go formatter over one generated source string and returns its
	exact stdout. Missing tools and rejected source fail compilation without exposing
	the generated source or machine-local paths. `formatAll` formats many sources in
	one formatter run and leaves every failure to `format`.

	How
	The source crosses one process boundary through standard input. The result is read
	back as a concrete string before it enters the existing-module output plan.
**/
class GoSourceFormatter {
	public static function format(source:String):String {
		var process:Null<Process> = null;
		try {
			process = new Process("gofmt", []);
			process.stdin.writeString(source);
			process.stdin.close();
			final formatted = process.stdout.readAll().toString();
			final diagnostics = StringTools.trim(process.stderr.readAll().toString());
			final exitCode = process.exitCode();
			process.close();
			process = null;
			if (exitCode != 0) {
				final detail = diagnostics == "" ? "the generated source was rejected" : diagnostics;
				throw formatError('gofmt failed: ${detail}');
			}
			return formatted;
		} catch (error:GoOutputPathError) {
			closeQuietly(process);
			throw error;
		} catch (_:Dynamic) {
			// Haxe process APIs expose launch and pipe failures as Dynamic. Contain that
			// host boundary here and return only a typed, path-free compiler failure.
			closeQuietly(process);
			throw formatError("gofmt could not format the generated source");
		}
	}

	/**
		What: Formats many sources with one `gofmt` run and returns them in the same
		order, or returns null when that run does not format every source.

		Why: Starting one `gofmt` process per generated file cost 5.4-6.2 s of the
		BeadsHX bootstrap compile (234 files). One run starts one process and
		formats the files on all cores.

		How: Writes each source to its own file in a new temporary directory, runs
		`gofmt -w` on the directory, and reads the files back. For a source that
		parses as a whole Go file, the file holds exactly what `format` returns:
		`gofmt` parses a file and standard input the same way and prints both with
		the same settings. Only standard input also accepts declaration or statement
		lists without a package clause. So any failure (such a fragment, a rejected
		source, a missing tool, or a file system error) returns null, and the caller
		formats each source with `format`, which reports the first failure exactly
		as it always has. The directory is removed afterwards. The files are written
		and read as UTF-8 bytes, as `format` writes and reads its pipes.
	**/
	public static function formatAll(sources:Array<String>):Null<Array<String>> {
		if (sources.length == 0) {
			return [];
		}
		var directory:Null<String> = null;
		var process:Null<Process> = null;
		try {
			directory = createScratchDirectory();
			if (directory == null) {
				return null;
			}
			final paths = [
				for (index in 0...sources.length) Path.join([directory, StringTools.lpad(Std.string(index), "0", 6) + ".go"])
			];
			for (index in 0...sources.length) {
				File.saveBytes(paths[index], Bytes.ofString(sources[index]));
			}
			process = new Process("gofmt", ["-w", directory]);
			process.stdin.close();
			// `gofmt -w` writes nothing to standard output; read the diagnostics first so
			// that many rejected files cannot fill that pipe while this waits.
			final diagnostics = process.stderr.readAll();
			process.stdout.readAll();
			final exitCode = process.exitCode();
			process.close();
			process = null;
			final formatted = exitCode == 0 && diagnostics.length == 0 ? [for (path in paths) File.getBytes(path).toString()] : null;
			removeScratchDirectory(directory);
			return formatted;
		} catch (_:haxe.Exception) {
			closeQuietly(process);
			if (directory != null) {
				removeScratchDirectory(directory);
			}
			return null;
		}
	}

	/** A new, empty directory under the host's temporary directory, or null when none can be made. */
	static function createScratchDirectory():Null<String> {
		final root = temporaryRoot();
		if (root == null) {
			return null;
		}
		for (_ in 0...8) {
			final candidate = Path.join([
				root,
				"haxe-go-gofmt-" + Std.string(Std.random(0x3fffffff)) + "-" + Std.string(Std.random(0x3fffffff))
			]);
			if (!FileSystem.exists(candidate)) {
				FileSystem.createDirectory(candidate);
				return candidate;
			}
		}
		return null;
	}

	static function temporaryRoot():Null<String> {
		for (name in ["TMPDIR", "TEMP", "TMP"]) {
			final value = Sys.getEnv(name);
			if (value != null && value != "") {
				return value;
			}
		}
		return Sys.systemName() == "Windows" ? null : "/tmp";
	}

	/** Removes the scratch directory and every file in it, ignoring failures: it holds only copies. */
	static function removeScratchDirectory(directory:String):Void {
		try {
			for (name in FileSystem.readDirectory(directory)) {
				FileSystem.deleteFile(Path.join([directory, name]));
			}
			FileSystem.deleteDirectory(directory);
		} catch (_:haxe.Exception) {}
	}

	static function closeQuietly(process:Null<Process>):Void {
		if (process == null) {
			return;
		}
		try {
			process.close();
		} catch (_:Dynamic) {
			// Cleanup cannot replace the original formatter failure.
		}
	}

	static function formatError(message:String):GoOutputPathError {
		return new GoOutputPathError(GoOutputPathErrorKind.WriteFailed, message);
	}
}
#else
class GoSourceFormatter {}
#end
