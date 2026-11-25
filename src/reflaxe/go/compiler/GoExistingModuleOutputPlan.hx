package reflaxe.go.compiler;

#if (macro || eval)
import haxe.io.Bytes;
import reflaxe.go.compiler.GoExistingModuleOwnership.GoOwnedFileRecord;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import reflaxe.go.compiler.GoProjectMode.ExistingGoModuleProject;
import reflaxe.output.StringOrBytes;
import reflaxe.output.StringOrBytes.StringOrBytesImpl;

/** One generated source, by output path, that the plan may format ahead of `add`. */
typedef GoExistingModuleSource = {
	final path:String;
	final source:String;
}

/** A source formatted ahead of `add`, kept with the exact text that was formatted. */
private typedef GoPreformattedSource = {
	final source:String;
	final formatted:String;
}

/** Exact bytes for one validated existing-module output artifact. */
typedef GoExistingModuleArtifact = {
	final path:String;
	final bytes:Bytes;
	final sha256:String;
}

/**
	Collects a complete existing-module output before the first filesystem effect.

	Why
	Collision and stale-file checks are only atomic when every generated source,
	runtime file, report, license, and macro extra file is known first.

	What
	Stores exact bytes under validated module-relative paths and rejects duplicate,
	case-aliased, or reserved destinations.

	How
	The compiler routes its ordinary emission helpers into this plan. The output
	transaction receives one sorted immutable copy after collection completes.
	Artifact digests come from the compile's shared `GoContentDigestMemo`: an
	artifact whose bytes equal the owned file the package inspector already
	hashed reuses that digest, and the transaction later reuses the plan's.
	Go sources pass through `gofmt`; `preformat` formats the generated
	sources in one run before they are added.
**/
class GoExistingModuleOutputPlan {
	final project:ExistingGoModuleProject;
	final boundary:GoGeneratedOutputBoundary;
	final digests:GoContentDigestMemo;
	final artifactsByPath:Map<String, GoExistingModuleArtifact> = [];
	final preformatted:Map<String, GoPreformattedSource> = [];

	public function new(project:ExistingGoModuleProject, boundary:GoGeneratedOutputBoundary, digests:GoContentDigestMemo) {
		this.project = project;
		this.boundary = boundary;
		this.digests = digests;
	}

	/**
		What: Formats the Go sources among `sources` in one `gofmt` run, for the
		`add` calls that follow.

		Why: `add` formatted each Go source with its own `gofmt` process: 234
		processes and 5.4-6.2 s of the BeadsHX bootstrap compile.

		How: `GoSourceFormatter.formatAll` formats them together. `add` uses a
		result once, and only for the same path and exactly the same source text.
		When the run does not format every source, nothing is kept, and each
		`add` formats its source alone, as before: a source that `gofmt` rejects
		fails at the same `add`, with the same message. Nothing here throws or
		writes output; the run works on copies in a temporary directory.
	**/
	public function preformat(sources:Array<GoExistingModuleSource>):Void {
		final goSources = [for (entry in sources) if (isGoSource(entry.path)) entry];
		final formatted = GoSourceFormatter.formatAll([for (entry in goSources) entry.source]);
		if (formatted == null) {
			return;
		}
		for (index in 0...goSources.length) {
			preformatted.set(goSources[index].path, {source: goSources[index].source, formatted: formatted[index]});
		}
	}

	public function add(path:String, content:StringOrBytes):Void {
		final bytes = switch (content.data()) {
			case String(value):
				final rendered = isGoSource(path) ? formatGoSource(path, value) : value;
				haxe.io.Bytes.ofString(rendered);
			case Bytes(value): value;
		};
		addBytes(path, bytes);
	}

	static inline function isGoSource(path:String):Bool {
		return StringTools.endsWith(path.toLowerCase(), ".go");
	}

	function formatGoSource(path:String, source:String):String {
		final prepared = preformatted.get(path);
		if (prepared != null) {
			preformatted.remove(path);
			if (prepared.source == source) {
				return prepared.formatted;
			}
		}
		return GoSourceFormatter.format(source);
	}

	public function addBytes(path:String, bytes:Bytes):Void {
		final relative = boundary.validateDestination(path).toString();
		if (GoExistingModuleOwnership.isReservedPath(project, relative)) {
			throw new GoOutputPathError(GoOutputPathErrorKind.ProtectedCallerFile, "a generated artifact targets a reserved existing-module path");
		}
		final key = GoExistingModuleOwnership.pathKey(relative);
		if (artifactsByPath.exists(key)) {
			throw new GoOutputPathError(GoOutputPathErrorKind.GeneratedFileConflict, "generated artifacts contain a duplicate destination");
		}
		artifactsByPath.set(key, {
			path: relative,
			bytes: bytes,
			sha256: digests.digest(relative, bytes)
		});
	}

	public function artifacts():Array<GoExistingModuleArtifact> {
		final values = [for (artifact in artifactsByPath) artifact];
		values.sort((left, right) -> left.path < right.path ? -1 : (left.path > right.path ? 1 : 0));
		return values;
	}

	public function ownershipRecords():Array<GoOwnedFileRecord> {
		return [for (artifact in artifacts()) {path: artifact.path, sha256: artifact.sha256}];
	}
}
#else
class GoExistingModuleOutputPlan {}
#end
