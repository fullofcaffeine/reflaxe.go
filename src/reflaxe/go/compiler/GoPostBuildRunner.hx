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

/** One exact name/value pair installed for a governed Go child process. */
typedef GoPostBuildEnvironmentEntry = {
	final name:String;
	final value:String;
}

/** Testable boundary for setting or removing one compiler-process variable. */
typedef GoPostBuildEnvironmentMutation = (name:String, value:Null<String>) -> Void;

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
	EnvironmentApplyFailed(variable:String);
	EnvironmentRestoreFailed(variable:String);
}

/**
	What: Success or ordered structured failures from the backend-owned Go build.
	Why: The compiler caller must handle command and cleanup failures before success.
	How: The runner returns this value after it attempts environment and directory
	restoration, so a cleanup error does not replace the primary command error.
**/
enum GoPostBuildResult {
	BuildSucceeded;
	BuildFailed(failures:Array<GoPostBuildFailure>);
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
		return runInternal(workingDirectory, command, arguments, null, execute, null);
	}

	/**
		Run with only the supplied environment, then restore the parent exactly.

		The optional mutation function exists for focused failure injection. Normal
		callers use `Sys.putEnv` through the default boundary.
	**/
	public static function runGoverned(workingDirectory:String, command:String, arguments:Array<String>, environment:Array<GoPostBuildEnvironmentEntry>,
			?execute:GoPostBuildCommand, ?mutateEnvironment:GoPostBuildEnvironmentMutation):GoPostBuildResult {
		return runInternal(workingDirectory, command, arguments, environment, execute, mutateEnvironment);
	}

	static function runInternal(workingDirectory:String, command:String, arguments:Array<String>, environment:Null<Array<GoPostBuildEnvironmentEntry>>,
			?execute:GoPostBuildCommand, ?mutateEnvironment:GoPostBuildEnvironmentMutation):GoPostBuildResult {
		var originalCwd = Sys.getCwd();
		var commandExecutor:GoPostBuildCommand = execute == null ? (cmd, args) -> Sys.command(cmd, args) : execute;
		var environmentMutator:GoPostBuildEnvironmentMutation = mutateEnvironment == null ? (name, value) -> Sys.putEnv(name, value) : mutateEnvironment;
		final failures:Array<GoPostBuildFailure> = [];
		var originalEnvironment:Null<Array<GoPostBuildEnvironmentEntry>> = null;
		if (environment != null) {
			originalEnvironment = environmentEntries();
		}

		try {
			Sys.setCwd(workingDirectory);
			if (environment != null) {
				final failedVariable = replaceEnvironment(environment, environmentMutator);
				if (failedVariable != null)
					failures.push(EnvironmentApplyFailed(failedVariable));
			}
			if (failures.length == 0) {
				var code = commandExecutor(command, arguments);
				if (code != 0)
					failures.push(CommandExited(code));
			}
		} catch (error:Dynamic) {
			// Haxe catch values are Dynamic by language contract. Localize that boundary
			// here and immediately convert it into a typed, path-redacted failure.
			failures.push(CommandLaunchFailed(sanitizeCause(error, workingDirectory, originalCwd, command)));
		}

		if (originalEnvironment != null) {
			final failedVariable = replaceEnvironment(originalEnvironment, environmentMutator);
			if (failedVariable != null)
				failures.push(EnvironmentRestoreFailed(failedVariable));
		}

		try {
			Sys.setCwd(originalCwd);
		} catch (error:Dynamic) {
			failures.push(WorkingDirectoryRestoreFailed(sanitizeCause(error, workingDirectory, originalCwd, command)));
		}

		return failures.length == 0 ? BuildSucceeded : BuildFailed(failures);
	}

	public static function failureMessage(command:String, arguments:Array<String>, result:GoPostBuildResult):String {
		var commandLabel = displayCommand(command, arguments);
		return switch (result) {
			case BuildSucceeded:
				"";
			case BuildFailed(failures):
				failures.map(failure -> failureDetail(commandLabel, failure)).join(" ");
		};
	}

	static function failureDetail(commandLabel:String, failure:GoPostBuildFailure):String {
		return switch (failure) {
			case CommandExited(code):
				'Post-generation Go build failed: `${commandLabel}` exited with status ${code}.';
			case CommandLaunchFailed(cause):
				'Post-generation Go build failed: could not launch `${commandLabel}` (${cause}).';
			case WorkingDirectoryRestoreFailed(cause):
				'Post-generation Go build failed while restoring the compiler working directory after `${commandLabel}` (${cause}).';
			case EnvironmentApplyFailed(variable):
				'Post-generation Go build failed: could not apply governed environment variable ${variable} before `${commandLabel}`.';
			case EnvironmentRestoreFailed(variable):
				'Post-generation Go build failed while restoring compiler environment variable ${variable} after `${commandLabel}`.';
		};
	}

	static function environmentEntries():Array<GoPostBuildEnvironmentEntry> {
		final environment = Sys.environment();
		final entries = [for (name in environment.keys()) {name: name, value: environment.get(name)}];
		entries.sort((left, right) -> Reflect.compare(normalizedEnvironmentName(left.name), normalizedEnvironmentName(right.name)));
		return entries;
	}

	static function replaceEnvironment(target:Array<GoPostBuildEnvironmentEntry>, mutate:GoPostBuildEnvironmentMutation):Null<String> {
		final current = environmentEntries();
		var firstFailure:Null<String> = null;
		for (entry in current) {
			if (!containsEnvironmentName(target, entry.name)) {
				try {
					mutate(entry.name, null);
				} catch (_:haxe.Exception) {
					if (firstFailure == null)
						firstFailure = entry.name;
				}
			}
		}
		final orderedTarget = target.copy();
		orderedTarget.sort((left, right) -> Reflect.compare(normalizedEnvironmentName(left.name), normalizedEnvironmentName(right.name)));
		for (entry in orderedTarget) {
			try {
				mutate(entry.name, entry.value);
			} catch (_:haxe.Exception) {
				if (firstFailure == null)
					firstFailure = entry.name;
			}
		}
		return firstFailure;
	}

	static function containsEnvironmentName(entries:Array<GoPostBuildEnvironmentEntry>, name:String):Bool {
		final normalized = normalizedEnvironmentName(name);
		for (entry in entries)
			if (normalizedEnvironmentName(entry.name) == normalized)
				return true;
		return false;
	}

	static function normalizedEnvironmentName(name:String):String {
		return Sys.systemName() == "Windows" ? name.toUpperCase() : name;
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
