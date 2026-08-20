package reflaxe.go.compiler;

import reflaxe.go.compiler.GoPostBuildRunner.GoPostBuildEnvironmentEntry;

/** Where one governed build-environment value came from. */
enum abstract GoBuildEnvironmentSource(String) to String {
	final Literal = "literal";
	final Inherit = "inherit";
	final Compiler = "compiler";
	final HostBaseline = "host-baseline";
}

/** How one effective value may appear in the generated build report. */
enum abstract GoBuildEnvironmentReportPolicy(String) {
	final Plain;
	final PathRedacted;
	final SecretRedacted;
	final Hidden;
}

/** One validated effective variable at the Go process boundary. */
typedef GoBuildEnvironmentEntry = {
	final name:String;
	final value:String;
	final source:GoBuildEnvironmentSource;
	final reportPolicy:GoBuildEnvironmentReportPolicy;
}

/** One value that is safe to serialize in the generated build report. */
typedef GoBuildEnvironmentReportEntry = {
	final name:String;
	final source:String;
	final value:String;
}

/**
	What: The complete, minimal environment for one existing-module Go build.
	Why: Passing the compiler process environment to Go would let unrelated flags,
	credentials, and workspace settings silently change a recorded build.
	How: The resolver creates this value from a closed manifest allowlist, compiler
	fixed controls, and the small host launch baseline. Consumers receive copies.
**/
class GoBuildEnvironment {
	final entries:Array<GoBuildEnvironmentEntry>;

	@:allow(reflaxe.go.compiler.GoBuildEnvironmentResolver)
	private function new(entries:Array<GoBuildEnvironmentEntry>) {
		this.entries = entries.copy();
	}

	/** Return the exact name/value pairs to install for the Go child process. */
	public function processEntries():Array<GoPostBuildEnvironmentEntry> {
		return entries.map(entry -> {name: entry.name, value: entry.value});
	}

	/** Return deterministic report values after applying each variable's policy. */
	public function reportEntries():Array<GoBuildEnvironmentReportEntry> {
		final report:Array<GoBuildEnvironmentReportEntry> = [];
		for (entry in entries) {
			final value = switch (entry.reportPolicy) {
				case Plain: entry.value;
				case PathRedacted: "<path>";
				case SecretRedacted: "<redacted>";
				case Hidden: continue;
			};
			report.push({name: entry.name, source: entry.source, value: value});
		}
		report.sort((left, right) -> Reflect.compare(left.name, right.name));
		return report;
	}
}
