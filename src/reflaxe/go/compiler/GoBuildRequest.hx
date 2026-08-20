package reflaxe.go.compiler;

/** Inputs for one validated `go build` request in a caller-owned module. */
private typedef GoBuildRequestData = {
	final packageTarget:String;
	final output:String;
	final tags:Array<String>;
	final ldflags:Array<String>;
	final trimpath:Bool;
	final race:Bool;
	final arguments:Array<String>;
}

/** The exact process boundary recorded and executed for a typed Go build. */
class GoBuildInvocation {
	public final command:String;
	public final arguments:Array<String>;

	public function new(command:String, arguments:Array<String>) {
		this.command = command;
		this.arguments = arguments.copy();
	}

	/** Render a path-independent record of the effective process invocation. */
	public function renderJson():String {
		final lines = [
			"{",
			'\t"schemaVersion": 1,',
			'\t"contract": "reflaxe.go/build-invocation",',
			'\t"workingDirectory": ".",',
			'\t"command": "${jsonEscape(command)}",',
			'\t"arguments": ['
		];
		for (index in 0...arguments.length) {
			final suffix = index == arguments.length - 1 ? "" : ",";
			lines.push('\t\t"${jsonEscape(arguments[index])}"${suffix}');
		}
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	static function jsonEscape(value:String):String {
		var escaped = StringTools.replace(value, "\\", "\\\\");
		escaped = StringTools.replace(escaped, '"', '\\"');
		escaped = StringTools.replace(escaped, "\n", "\\n");
		escaped = StringTools.replace(escaped, "\r", "\\r");
		return StringTools.replace(escaped, "\t", "\\t");
	}
}

/**
	A build request narrowed from the project manifest before generation starts.

	The request owns canonical argument ordering. Linker arguments use the same
	quoting contract as Go's `cmd/internal/quoted.Join`, so they remain individual
	linker tokens without passing through a shell.
**/
class GoBuildRequest {
	public static inline final REPORT_PATH = "reflaxe_go_build.json";

	final packageTarget:String;
	final output:String;
	final tags:Array<String>;
	final ldflags:Array<String>;
	final trimpath:Bool;
	final race:Bool;
	final additionalArguments:Array<String>;

	@:allow(reflaxe.go.compiler.GoProjectModeResolver)
	private function new(data:GoBuildRequestData) {
		packageTarget = data.packageTarget;
		output = data.output;
		tags = data.tags.copy();
		ldflags = data.ldflags.copy();
		trimpath = data.trimpath;
		race = data.race;
		additionalArguments = data.arguments.copy();
	}

	/** Build the exact, deterministic process invocation. */
	public function invocation():GoBuildInvocation {
		final arguments = ["build"];
		if (trimpath) {
			arguments.push("-trimpath");
		}
		if (race) {
			arguments.push("-race");
		}
		if (tags.length > 0) {
			arguments.push("-tags=" + tags.join(","));
		}
		if (ldflags.length > 0) {
			arguments.push("-ldflags=" + joinQuoted(ldflags));
		}
		for (argument in additionalArguments) {
			arguments.push(argument);
		}
		arguments.push("-o");
		arguments.push(output);
		arguments.push(packageTarget);
		return new GoBuildInvocation("go", arguments);
	}

	static function joinQuoted(values:Array<String>):String {
		return values.map(quoteIfNeeded).join(" ");
	}

	static function quoteIfNeeded(value:String):String {
		var hasSpace = false;
		var hasSingleQuote = false;
		var hasDoubleQuote = false;
		for (index in 0...value.length) {
			switch (value.charAt(index)) {
				case " ", "\t", "\n", "\r":
					hasSpace = true;
				case "'":
					hasSingleQuote = true;
				case '"':
					hasDoubleQuote = true;
				case _:
			}
		}
		if (!hasSpace && !hasSingleQuote && !hasDoubleQuote) {
			return value;
		}
		return hasSingleQuote ? '"${value}"' : "'" + value + "'";
	}
}
