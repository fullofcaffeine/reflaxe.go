package reflaxe.go.compiler;

#if (macro || eval)
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
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
	the generated source or machine-local paths.

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
