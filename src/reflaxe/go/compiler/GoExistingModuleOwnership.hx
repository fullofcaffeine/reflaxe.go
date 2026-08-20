package reflaxe.go.compiler;

#if (macro || eval)
import haxe.Json;
import haxe.crypto.Sha256;
import haxe.io.Bytes;
import haxe.io.Path;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import reflaxe.go.compiler.GoProjectMode.ExistingGoModuleProject;

/** One exact compiler-owned file in a caller Go module. */
typedef GoOwnedFileRecord = {
	final path:String;
	final sha256:String;
}

private typedef RawOwnershipManifest = {
	final schemaVersion:Int;
	final manifestSchemaVersion:Int;
	final modulePath:String;
	final packageDir:String;
	final packageName:String;
	final runtimeDir:String;
	final files:Array<GoOwnedFileRecord>;
}

/**
	A validated, immutable view of one existing-module ownership record.

	Why
	Package inspection, collision checks, cleanup, and interruption recovery must
	share one meaning of ownership. A path-only list cannot distinguish generated
	bytes from a caller edit.

	What
	Retains the exact project identity, sorted module-relative paths, SHA-256
	digests, and original manifest bytes when a record exists.

	How
	`GoExistingModuleOwnership.load` validates the JSON boundary before this value
	reaches compiler logic. Callers can query records but cannot mutate the maps.
**/
class GoExistingModuleOwnershipSnapshot {
	public final files:Array<GoOwnedFileRecord>;
	public final sourceBytes:Null<Bytes>;
	public final sourceDigest:Null<String>;

	final byPath:Map<String, GoOwnedFileRecord>;

	@:allow(reflaxe.go.compiler.GoExistingModuleOwnership)
	private function new(files:Array<GoOwnedFileRecord>, sourceBytes:Null<Bytes>) {
		this.files = files;
		this.sourceBytes = sourceBytes;
		sourceDigest = sourceBytes == null ? null : GoExistingModuleOwnership.digest(sourceBytes);
		byPath = [];
		for (record in files) {
			byPath.set(GoExistingModuleOwnership.pathKey(record.path), record);
		}
	}

	public inline function exists():Bool {
		return sourceBytes != null;
	}

	public function record(path:String):Null<GoOwnedFileRecord> {
		final owned = byPath.get(GoExistingModuleOwnership.pathKey(path));
		return owned != null && owned.path == path ? owned : null;
	}

	/** Return true when ownership contains only a differently cased spelling. */
	public function hasCaseAlias(path:String):Bool {
		final owned = byPath.get(GoExistingModuleOwnership.pathKey(path));
		return owned != null && owned.path != path;
	}

	/** Return true only when the record owns the path and its current bytes. */
	public function ownsCurrent(path:String, boundary:GoGeneratedOutputBoundary):Bool {
		final owned = record(path);
		if (owned == null) {
			return false;
		}
		final current = boundary.readBytes(path);
		return current != null && GoExistingModuleOwnership.digest(current) == owned.sha256;
	}
}

/**
	Strict codec and path policy for package-local existing-module ownership.

	The host JSON value is assigned to one precise raw structure in `parse`.
	Every field is validated immediately and only concrete ownership records leave
	that boundary.
**/
class GoExistingModuleOwnership {
	public static inline final SCHEMA_VERSION = 1;
	public static inline final MANIFEST_SCHEMA_VERSION = 1;
	public static inline final OWNERSHIP_FILE_NAME = ".reflaxe-go-owned.json";
	public static inline final JOURNAL_FILE_NAME = ".reflaxe-go-transaction.json";
	public static inline final JOURNAL_TEMP_FILE_NAME = ".reflaxe-go-transaction.json.tmp";
	public static inline final WORKSPACE_NAME = ".reflaxe-go-transaction";
	public static inline final WORKSPACE_MARKER_NAME = "project.json";
	public static inline final NEW_MANIFEST_NAME = "new-ownership.json";
	public static inline final OLD_MANIFEST_NAME = "old-ownership.json";

	static final MANIFEST_FIELDS = [
		"schemaVersion",
		"manifestSchemaVersion",
		"modulePath",
		"packageDir",
		"packageName",
		"runtimeDir",
		"files"
	];
	static final FILE_FIELDS = ["path", "sha256"];

	public static inline function ownershipPath(project:ExistingGoModuleProject):String {
		return project.packageDir.resolve(OWNERSHIP_FILE_NAME);
	}

	public static inline function journalPath(project:ExistingGoModuleProject):String {
		return project.packageDir.resolve(JOURNAL_FILE_NAME);
	}

	public static inline function journalTempPath(project:ExistingGoModuleProject):String {
		return project.packageDir.resolve(JOURNAL_TEMP_FILE_NAME);
	}

	public static inline function workspacePath(project:ExistingGoModuleProject):String {
		return project.packageDir.resolve(WORKSPACE_NAME);
	}

	public static inline function workspaceChild(project:ExistingGoModuleProject, child:String):String {
		return workspacePath(project) + "/" + child;
	}

	public static function empty():GoExistingModuleOwnershipSnapshot {
		return new GoExistingModuleOwnershipSnapshot([], null);
	}

	public static function load(project:ExistingGoModuleProject, boundary:GoGeneratedOutputBoundary):GoExistingModuleOwnershipSnapshot {
		final bytes = boundary.readBytes(ownershipPath(project));
		return bytes == null ? empty() : parse(project, boundary, bytes);
	}

	public static function parse(project:ExistingGoModuleProject, boundary:GoGeneratedOutputBoundary, bytes:Bytes):GoExistingModuleOwnershipSnapshot {
		final raw:RawOwnershipManifest = try {
			Json.parse(bytes.toString());
		} catch (_:haxe.Exception) {
			throw invalid("the ownership record is not valid JSON");
		}
		validateObject(raw, MANIFEST_FIELDS, "the ownership record has an invalid shape");
		if (!isIntField(raw, "schemaVersion")
			|| raw.schemaVersion != SCHEMA_VERSION
			|| !isIntField(raw, "manifestSchemaVersion")
			|| raw.manifestSchemaVersion != MANIFEST_SCHEMA_VERSION) {
			throw invalid("the ownership record schema is not supported");
		}
		for (field in ["modulePath", "packageDir", "packageName", "runtimeDir"]) {
			if (!isStringField(raw, field)) {
				throw invalid("the ownership record has an invalid identity");
			}
		}
		if (raw.modulePath != project.modulePath
			|| raw.packageDir != project.packageDir.value()
			|| raw.packageName != project.packageName.value()
			|| raw.runtimeDir != project.runtimeDir.value()) {
			throw invalid("the ownership record belongs to another project shape");
		}

		final rawFiles:Array<{}> = Reflect.field(raw, "files");
		if (!Std.isOfType(rawFiles, Array)) {
			throw invalid("the ownership record files field is not an array");
		}
		final files:Array<GoOwnedFileRecord> = [];
		final seen:Map<String, Bool> = [];
		var priorPath:Null<String> = null;
		for (rawFile in rawFiles) {
			validateObject(rawFile, FILE_FIELDS, "an ownership file record has an invalid shape");
			if (!isStringField(rawFile, "path") || !isStringField(rawFile, "sha256")) {
				throw invalid("an ownership file record has an invalid value");
			}
			final path:String = Reflect.field(rawFile, "path");
			final sha256:String = Reflect.field(rawFile, "sha256");
			boundary.validateDestination(path);
			if (isReservedPath(project, path)) {
				throw invalid("the ownership record names a reserved path");
			}
			if (!~/^[0-9a-f]{64}$/.match(sha256)) {
				throw invalid("an ownership file record has an invalid digest");
			}
			final key = pathKey(path);
			if (seen.exists(key)) {
				throw invalid("the ownership record contains a duplicate path");
			}
			if (priorPath != null && priorPath > path) {
				throw invalid("the ownership record paths are not sorted");
			}
			seen.set(key, true);
			files.push({path: path, sha256: sha256});
			priorPath = path;
		}
		return new GoExistingModuleOwnershipSnapshot(files, bytes);
	}

	public static function render(project:ExistingGoModuleProject, files:Array<GoOwnedFileRecord>):Bytes {
		final sorted = files.copy();
		sorted.sort((left, right) -> left.path < right.path ? -1 : (left.path > right.path ? 1 : 0));
		final lines = [
			"{",
			'  "schemaVersion": ${SCHEMA_VERSION},',
			'  "manifestSchemaVersion": ${MANIFEST_SCHEMA_VERSION},',
			'  "modulePath": ${jsonString(project.modulePath)},',
			'  "packageDir": ${jsonString(project.packageDir.value())},',
			'  "packageName": ${jsonString(project.packageName.value())},',
			'  "runtimeDir": ${jsonString(project.runtimeDir.value())},',
			'  "files": ['
		];
		for (index in 0...sorted.length) {
			final record = sorted[index];
			final suffix = index == sorted.length - 1 ? "" : ",";
			lines.push('    {"path": ${jsonString(record.path)}, "sha256": ${jsonString(record.sha256)}}${suffix}');
		}
		lines.push("  ]");
		lines.push("}");
		lines.push("");
		return Bytes.ofString(lines.join("\n"));
	}

	/** Exact deterministic project identity used by the transaction workspace. */
	public static function renderProjectIdentity(project:ExistingGoModuleProject):Bytes {
		return Bytes.ofString([
			"{",
			'  "schemaVersion": ${SCHEMA_VERSION},',
			'  "modulePath": ${jsonString(project.modulePath)},',
			'  "packageDir": ${jsonString(project.packageDir.value())},',
			'  "packageName": ${jsonString(project.packageName.value())},',
			'  "runtimeDir": ${jsonString(project.runtimeDir.value())}',
			"}",
			""
		].join("\n"));
	}

	public static function isReservedPath(project:ExistingGoModuleProject, path:String):Bool {
		final key = pathKey(path);
		final workspace = pathKey(workspacePath(project));
		if (key == "go.mod"
			|| key == "go.sum"
			|| key == pathKey(GoGeneratedOutputBoundary.MANAGED_FILE_METADATA)
			|| key == pathKey(ownershipPath(project))
			|| key == pathKey(journalPath(project))
			|| key == pathKey(journalTempPath(project))
			|| key == workspace
			|| StringTools.startsWith(key, workspace + "/")) {
			return true;
		}
		final manifestRelative = moduleRelativePath(project.moduleRoot, project.manifestPath);
		return manifestRelative != null && key == pathKey(manifestRelative);
	}

	public static inline function digest(bytes:Bytes):String {
		return Sha256.make(bytes).toHex();
	}

	public static inline function pathKey(path:String):String {
		return path.toLowerCase();
	}

	static function moduleRelativePath(moduleRoot:String, absolutePath:String):Null<String> {
		final root = normalizePath(moduleRoot);
		final candidate = normalizePath(absolutePath);
		if (candidate == root) {
			return null;
		}
		final prefix = root + "/";
		return StringTools.startsWith(candidate, prefix) ? candidate.substr(prefix.length) : null;
	}

	static function normalizePath(path:String):String {
		var value = StringTools.replace(Path.normalize(path), "\\", "/");
		while (value.length > 1 && StringTools.endsWith(value, "/")) {
			value = value.substr(0, value.length - 1);
		}
		return value;
	}

	static function validateObject(value:{}, allowed:Array<String>, explanation:String):Void {
		if (value == null || Std.isOfType(value, Array) || !Reflect.isObject(value)) {
			throw invalid(explanation);
		}
		final fields = Reflect.fields(value);
		fields.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		final expected = allowed.copy();
		expected.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		if (fields.length != expected.length) {
			throw invalid(explanation);
		}
		for (index in 0...fields.length) {
			if (fields[index] != expected[index]) {
				throw invalid(explanation);
			}
		}
	}

	static inline function isStringField(value:{}, field:String):Bool {
		return Std.isOfType(Reflect.field(value, field), String);
	}

	static inline function isIntField(value:{}, field:String):Bool {
		return Std.isOfType(Reflect.field(value, field), Int);
	}

	static inline function jsonString(value:String):String {
		return Json.stringify(value);
	}

	static function invalid(explanation:String):GoOutputPathError {
		return new GoOutputPathError(GoOutputPathErrorKind.InvalidManagedMetadata, explanation);
	}
}
#else
class GoExistingModuleOwnership {}
class GoExistingModuleOwnershipSnapshot {}
#end
