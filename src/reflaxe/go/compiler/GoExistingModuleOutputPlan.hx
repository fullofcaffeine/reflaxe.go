package reflaxe.go.compiler;

#if (macro || eval)
import haxe.io.Bytes;
import reflaxe.go.compiler.GoExistingModuleOwnership.GoOwnedFileRecord;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import reflaxe.go.compiler.GoProjectMode.ExistingGoModuleProject;
import reflaxe.output.StringOrBytes;
import reflaxe.output.StringOrBytes.StringOrBytesImpl;

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
**/
class GoExistingModuleOutputPlan {
	final project:ExistingGoModuleProject;
	final boundary:GoGeneratedOutputBoundary;
	final artifactsByPath:Map<String, GoExistingModuleArtifact> = [];

	public function new(project:ExistingGoModuleProject, boundary:GoGeneratedOutputBoundary) {
		this.project = project;
		this.boundary = boundary;
	}

	public function add(path:String, content:StringOrBytes):Void {
		final bytes = switch (content.data()) {
			case String(value):
				final rendered = StringTools.endsWith(path.toLowerCase(), ".go") ? GoSourceFormatter.format(value) : value;
				haxe.io.Bytes.ofString(rendered);
			case Bytes(value): value;
		};
		addBytes(path, bytes);
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
			sha256: GoExistingModuleOwnership.digest(bytes)
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
