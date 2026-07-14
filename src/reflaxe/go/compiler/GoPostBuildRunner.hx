package reflaxe.go.compiler;

import haxe.io.Path;

/**
	What: A typed command boundary for the backend-owned post-generation Go build.
	Why: Haxe compilation must not report success after `go build` fails, and command
	launch errors must not strand the macro process in the generated output directory.
	How: Execute from the output directory, convert launch/exit/cleanup failures into
	typed results, restore the original directory, and redact known local paths before
	a diagnostic reaches the compiler surface.
**/
typedef GoPostBuildCommand = (command:String, arguments:Array<String>) -> Int;

/**
	What: The bounded failure modes of a post-generation Go build.
	Why: Launch, exit-status, and compiler-state cleanup failures require distinct,
	actionable diagnostics instead of a warning-only string.
	How: Capture only the exit code or a path-redacted exception detail at the process
	boundary; child stderr continues to flow through `Sys.command` unchanged.
**/
enum GoPostBuildFailure {
	CommandExited(code:Int);
	CommandLaunchFailed(cause:String);
	WorkingDirectoryRestoreFailed(cause:String);
}

/**
	What: Success or one structured failure from the backend-owned Go build.
	Why: The compiler caller must handle every failure before it can report success.
	How: `GoPostBuildRunner.run` returns this value only after attempting to restore
	the compiler working directory.
**/
enum GoPostBuildResult {
	BuildSucceeded;
	BuildFailed(failure:GoPostBuildFailure);
}

/**
	What: Runs and renders the post-generation `go build` phase.
	Why: Process execution, directory restoration, and path-safe diagnostics form one
	correctness boundary and should not be duplicated in the outer compiler lifecycle.
	How: Call `run`, then treat every `BuildFailed` result as a fatal compiler outcome;
	`failureMessage` produces the stable user-facing diagnostic.
**/
class GoPostBuildRunner {
	public static function run(workingDirectory:String, command:String, arguments:Array<String>, ?execute:GoPostBuildCommand):GoPostBuildResult {
		var originalCwd = Sys.getCwd();
		var commandExecutor:GoPostBuildCommand = execute == null ? (cmd, args) -> Sys.command(cmd, args) : execute;
		var result:GoPostBuildResult = BuildSucceeded;

		try {
			Sys.setCwd(workingDirectory);
			var code = commandExecutor(command, arguments);
			if (code != 0) {
				result = BuildFailed(CommandExited(code));
			}
		} catch (error:Dynamic) {
			// Haxe catch values are Dynamic by language contract. Localize that boundary
			// here and immediately convert it into a typed, path-redacted failure.
			result = BuildFailed(CommandLaunchFailed(sanitizeCause(error, workingDirectory, originalCwd, command)));
		}

		try {
			Sys.setCwd(originalCwd);
		} catch (error:Dynamic) {
			return BuildFailed(WorkingDirectoryRestoreFailed(sanitizeCause(error, workingDirectory, originalCwd, command)));
		}

		return result;
	}

	public static function failureMessage(command:String, arguments:Array<String>, result:GoPostBuildResult):String {
		var commandLabel = displayCommand(command, arguments);
		return switch (result) {
			case BuildSucceeded:
				"";
			case BuildFailed(CommandExited(code)):
				'Post-generation Go build failed: `${commandLabel}` exited with status ${code}.';
			case BuildFailed(CommandLaunchFailed(cause)):
				'Post-generation Go build failed: could not launch `${commandLabel}` (${cause}).';
			case BuildFailed(WorkingDirectoryRestoreFailed(cause)):
				'Post-generation Go build failed while restoring the compiler working directory after `${commandLabel}` (${cause}).';
		};
	}

	static function displayCommand(command:String, arguments:Array<String>):String {
		var normalizedCommand = StringTools.replace(command, "\\", "/");
		var executable = Path.withoutDirectory(normalizedCommand);
		var displayedArguments = arguments.map(argument -> Path.isAbsolute(argument) ? "<absolute-path>" : argument);
		return [executable].concat(displayedArguments).join(" ");
	}

	static function sanitizeCause(error:Dynamic, workingDirectory:String, originalCwd:String, command:String):String {
		var cause = Std.string(error);
		var redactions:Array<{value:String, replacement:String}> = [];
		addRedaction(redactions, workingDirectory, "<generated-output>");
		addRedaction(redactions, originalCwd, "<compiler-working-directory>");
		if (Path.isAbsolute(command)) {
			addRedaction(redactions, command, "<go-command>");
		}
		redactions.sort((left, right) -> right.value.length - left.value.length);
		for (redaction in redactions) {
			cause = StringTools.replace(cause, redaction.value, redaction.replacement);
		}
		return cause;
	}

	static function addRedaction(redactions:Array<{value:String, replacement:String}>, value:String, replacement:String):Void {
		if (value == null || value == "") {
			return;
		}
		redactions.push({value: value, replacement: replacement});
		var withoutTrailingSeparators = value;
		while (withoutTrailingSeparators.length > 1
			&& (StringTools.endsWith(withoutTrailingSeparators, "/") || StringTools.endsWith(withoutTrailingSeparators, "\\"))) {
			withoutTrailingSeparators = withoutTrailingSeparators.substr(0, withoutTrailingSeparators.length - 1);
		}
		if (withoutTrailingSeparators != value) {
			redactions.push({value: withoutTrailingSeparators, replacement: replacement});
		}
	}
}
