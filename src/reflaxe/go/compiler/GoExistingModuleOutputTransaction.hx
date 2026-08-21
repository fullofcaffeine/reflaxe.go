package reflaxe.go.compiler;

#if (macro || eval)
import eval.luv.File.FileSync;
import eval.luv.Result;
import haxe.Json;
import haxe.io.Bytes;
import haxe.io.Path;
import reflaxe.go.compiler.GoExistingModuleOutputPlan.GoExistingModuleArtifact;
import reflaxe.go.compiler.GoExistingModuleOwnership.GoOwnedFileRecord;
import reflaxe.go.compiler.GoExistingModuleOwnership.GoExistingModuleOwnershipSnapshot;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import reflaxe.go.compiler.GoProjectMode.ExistingGoModuleProject;
import sys.FileSystem;

private typedef GoOutputTransactionJournal = {
	final projectSha256:String;
	final oldManifestSha256:String;
	final newManifestSha256:String;
}

private typedef RawGoOutputTransactionJournal = {
	final schemaVersion:Int;
	final projectSha256:String;
	final oldManifestSha256:String;
	final newManifestSha256:String;
}

/**
	Owns one digest-checked update of generated files in a caller Go module.

	Why
	Reflaxe's generic generated-file list contains paths but no content digests.
	It can safely own a standalone output tree, but it cannot decide replacement
	or cleanup inside a mixed caller/compiler module.

	What
	Preflights a complete artifact plan, stages exact bytes, preserves verified old
	bytes, installs files, and publishes the package-local ownership record last.

	How
	An immutable journal and workspace make every live step recoverable. Before a
	new commit, `recover` either rolls back a pre-commit interruption or verifies
	a manifest-last commit and removes its control files. Unknown bytes, links, or
	malformed control data stop without deleting evidence.
**/
class GoExistingModuleOutputTransaction {
	static inline final JOURNAL_SCHEMA_VERSION = 1;
	static final JOURNAL_FIELDS = ["schemaVersion", "projectSha256", "oldManifestSha256", "newManifestSha256"];

	final project:ExistingGoModuleProject;
	final boundary:GoGeneratedOutputBoundary;

	public function new(project:ExistingGoModuleProject, boundary:GoGeneratedOutputBoundary) {
		this.project = project;
		this.boundary = boundary;
	}

	public function commit(plan:GoExistingModuleOutputPlan):Void {
		recover();
		final oldOwnership = GoExistingModuleOwnership.load(project, boundary);
		verifySnapshot(oldOwnership, "a previously owned generated file changed");
		validateRuntimeDirectoryOwnership(oldOwnership);
		final artifacts = plan.artifacts();
		preflightDestinations(oldOwnership, artifacts);

		final newManifestBytes = GoExistingModuleOwnership.render(project, plan.ownershipRecords());
		prepareWorkspace(oldOwnership, artifacts, newManifestBytes);
		var journalPublished = false;
		try {
			publishJournal(oldOwnership, newManifestBytes);
			journalPublished = true;
			installArtifacts(oldOwnership, artifacts);
			installOwnership(oldOwnership);
			final installed = GoExistingModuleOwnership.load(project, boundary);
			if (installed.sourceDigest != GoExistingModuleOwnership.digest(newManifestBytes)) {
				throw transactionError("the installed ownership record differs from the staged commit");
			}
			verifySnapshot(installed, "a newly installed generated file differs from its ownership record");
			cleanupControlFiles();
		} catch (error:GoOutputPathError) {
			if (journalPublished) {
				try {
					recover();
				} catch (_:GoOutputPathError) {}
			} else {
				try {
					cleanupUnpublishedWorkspace();
				} catch (_:GoOutputPathError) {}
			}
			throw error;
		} catch (_:haxe.Exception) {
			if (journalPublished) {
				try {
					recover();
				} catch (_:GoOutputPathError) {}
			} else {
				try {
					cleanupUnpublishedWorkspace();
				} catch (_:GoOutputPathError) {}
			}
			throw new GoOutputPathError(GoOutputPathErrorKind.WriteFailed, "the existing-module output transaction failed");
		}
	}

	/** Recover a prior transaction before package ownership is inspected. */
	public function recover():Void {
		final journalBytes = boundary.readBytes(GoExistingModuleOwnership.journalPath(project));
		final journalTemp = boundary.readBytes(GoExistingModuleOwnership.journalTempPath(project));
		if (journalBytes == null) {
			if (!workspaceExists() && journalTemp == null) {
				return;
			}
			validateWorkspaceIdentity();
			if (journalTemp != null) {
				parseJournal(journalTemp);
				boundary.deleteFile(GoExistingModuleOwnership.journalTempPath(project));
			}
			removeWorkspace();
			return;
		}

		if (journalTemp != null) {
			throw transactionError("both transaction journal paths exist");
		}
		validateWorkspaceIdentity();
		final journal = parseJournal(journalBytes);
		final newManifestBytes = requiredWorkspaceFile(GoExistingModuleOwnership.NEW_MANIFEST_NAME);
		if (GoExistingModuleOwnership.digest(newManifestBytes) != journal.newManifestSha256) {
			throw transactionError("the staged ownership record differs from the transaction journal");
		}
		final newOwnership = GoExistingModuleOwnership.parse(project, boundary, newManifestBytes);
		final oldOwnership = loadJournalOldOwnership(journal);
		final liveManifest = boundary.readBytes(GoExistingModuleOwnership.ownershipPath(project));
		final liveDigest = liveManifest == null ? null : GoExistingModuleOwnership.digest(liveManifest);

		if (liveDigest == journal.newManifestSha256) {
			verifySnapshot(newOwnership, "a committed generated file differs during recovery");
			for (record in oldOwnership.files) {
				if (newOwnership.record(record.path) == null && boundary.readBytes(record.path) != null) {
					throw transactionError("a committed stale generated file still exists");
				}
			}
			cleanupControlFiles();
			return;
		}

		final oldLiveStateMatches = oldOwnership.exists() ? liveDigest == oldOwnership.sourceDigest
			|| liveManifest == null : liveManifest == null;
		if (!oldLiveStateMatches) {
			throw transactionError("the live ownership record has an unknown interrupted state");
		}
		rollbackFiles(oldOwnership, newOwnership);
		restoreOwnership(oldOwnership, newOwnership);
		verifySnapshot(oldOwnership, "a rolled-back generated file differs from prior ownership");
		cleanupControlFiles();
	}

	function preflightDestinations(oldOwnership:GoExistingModuleOwnershipSnapshot, artifacts:Array<GoExistingModuleArtifact>):Void {
		for (artifact in artifacts) {
			final previous = oldOwnership.record(artifact.path);
			final current = boundary.readBytes(artifact.path);
			if (previous == null) {
				if (oldOwnership.hasCaseAlias(artifact.path)) {
					throw conflict("a generated artifact changes the exact spelling of an owned path");
				}
				if (current != null) {
					throw conflict("a generated artifact destination is caller-owned");
				}
			} else if (current == null || GoExistingModuleOwnership.digest(current) != previous.sha256) {
				throw conflict("a previously generated artifact changed before replacement");
			}
		}
	}

	/** Reject caller-owned content inside the compiler-owned runtime package. */
	function validateRuntimeDirectoryOwnership(oldOwnership:GoExistingModuleOwnershipSnapshot):Void {
		final directory = Path.join([project.moduleRoot, project.runtimeDir.value()]);
		try {
			if (!FileSystem.exists(directory)) {
				return;
			}
			for (name in FileSystem.readDirectory(directory)) {
				final absolute = Path.join([directory, name]);
				if (FileSystem.isDirectory(absolute)) {
					throw conflict("the compiler runtime directory contains caller-owned content");
				}
				final path = project.runtimePath(name);
				final record = oldOwnership.record(path);
				final current = boundary.readBytes(path);
				if (record == null || current == null || GoExistingModuleOwnership.digest(current) != record.sha256) {
					throw conflict("the compiler runtime directory contains caller-owned content");
				}
			}
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw transactionError("the compiler runtime directory could not be inspected");
		}
	}

	function prepareWorkspace(oldOwnership:GoExistingModuleOwnershipSnapshot, artifacts:Array<GoExistingModuleArtifact>, newManifestBytes:Bytes):Void {
		if (workspaceExists()
			|| boundary.readBytes(GoExistingModuleOwnership.journalPath(project)) != null
			|| boundary.readBytes(GoExistingModuleOwnership.journalTempPath(project)) != null) {
			throw transactionError("transaction control paths already exist");
		}
		boundary.saveBytes(workspaceFile(GoExistingModuleOwnership.WORKSPACE_MARKER_NAME), GoExistingModuleOwnership.renderProjectIdentity(project));
		for (artifact in artifacts) {
			boundary.saveBytes(workspaceFile("stage/" + artifact.path), artifact.bytes);
		}
		for (record in oldOwnership.files) {
			final current = boundary.readBytes(record.path);
			if (current == null || GoExistingModuleOwnership.digest(current) != record.sha256) {
				throw conflict("a previously owned generated file changed while staging");
			}
			boundary.saveBytes(workspaceFile("backup/" + record.path), current);
		}
		if (oldOwnership.sourceBytes != null) {
			boundary.saveBytes(workspaceFile(GoExistingModuleOwnership.OLD_MANIFEST_NAME), oldOwnership.sourceBytes);
		}
		boundary.saveBytes(workspaceFile(GoExistingModuleOwnership.NEW_MANIFEST_NAME), newManifestBytes);
		boundary.saveBytes(workspaceFile("install-ownership.json"), newManifestBytes);
	}

	function publishJournal(oldOwnership:GoExistingModuleOwnershipSnapshot, newManifestBytes:Bytes):Void {
		final oldDigest = oldOwnership.sourceDigest == null ? "" : oldOwnership.sourceDigest;
		final bytes = renderJournal({
			projectSha256: GoExistingModuleOwnership.digest(GoExistingModuleOwnership.renderProjectIdentity(project)),
			oldManifestSha256: oldDigest,
			newManifestSha256: GoExistingModuleOwnership.digest(newManifestBytes)
		});
		boundary.saveBytes(GoExistingModuleOwnership.journalTempPath(project), bytes);
		boundary.moveFile(GoExistingModuleOwnership.journalTempPath(project), GoExistingModuleOwnership.journalPath(project));
	}

	function installArtifacts(oldOwnership:GoExistingModuleOwnershipSnapshot, artifacts:Array<GoExistingModuleArtifact>):Void {
		final newPaths:Map<String, Bool> = [];
		for (artifact in artifacts) {
			newPaths.set(GoExistingModuleOwnership.pathKey(artifact.path), true);
			if (oldOwnership.record(artifact.path) != null) {
				boundary.moveFile(artifact.path, workspaceFile("displaced/" + artifact.path));
			}
			boundary.moveFile(workspaceFile("stage/" + artifact.path), artifact.path);
		}
		for (record in oldOwnership.files) {
			if (!newPaths.exists(GoExistingModuleOwnership.pathKey(record.path))) {
				boundary.moveFile(record.path, workspaceFile("displaced/" + record.path));
			}
		}
	}

	function installOwnership(oldOwnership:GoExistingModuleOwnershipSnapshot):Void {
		final ownershipPath = GoExistingModuleOwnership.ownershipPath(project);
		if (oldOwnership.exists()) {
			boundary.moveFile(ownershipPath, workspaceFile("displaced-ownership.json"));
		}
		boundary.moveFile(workspaceFile("install-ownership.json"), ownershipPath);
	}

	function rollbackFiles(oldOwnership:GoExistingModuleOwnershipSnapshot, newOwnership:GoExistingModuleOwnershipSnapshot):Void {
		final paths:Map<String, String> = [];
		for (record in oldOwnership.files) {
			paths.set(GoExistingModuleOwnership.pathKey(record.path), record.path);
		}
		for (record in newOwnership.files) {
			paths.set(GoExistingModuleOwnership.pathKey(record.path), record.path);
		}
		final ordered = [for (path in paths) path];
		ordered.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		for (path in ordered) {
			final oldRecord = oldOwnership.record(path);
			final newRecord = newOwnership.record(path);
			final current = boundary.readBytes(path);
			final currentDigest = current == null ? null : GoExistingModuleOwnership.digest(current);
			if (oldRecord == null) {
				if (currentDigest == null) {
					continue;
				}
				if (newRecord == null || currentDigest != newRecord.sha256) {
					throw transactionError("an unowned file has unknown bytes during rollback");
				}
				boundary.deleteFile(path);
				continue;
			}
			if (currentDigest == oldRecord.sha256) {
				continue;
			}
			if (currentDigest != null && (newRecord == null || currentDigest != newRecord.sha256)) {
				throw transactionError("an owned file has unknown bytes during rollback");
			}
			final backup = requiredWorkspaceFile("backup/" + oldRecord.path);
			if (GoExistingModuleOwnership.digest(backup) != oldRecord.sha256) {
				throw transactionError("a transaction backup differs from prior ownership");
			}
			if (current != null) {
				boundary.deleteFile(path);
			}
			final restorePath = workspaceFile("rollback-stage/" + oldRecord.path);
			boundary.saveBytes(restorePath, backup);
			boundary.moveFile(restorePath, path);
		}
	}

	function restoreOwnership(oldOwnership:GoExistingModuleOwnershipSnapshot, newOwnership:GoExistingModuleOwnershipSnapshot):Void {
		final path = GoExistingModuleOwnership.ownershipPath(project);
		final current = boundary.readBytes(path);
		final currentDigest = current == null ? null : GoExistingModuleOwnership.digest(current);
		if (!oldOwnership.exists()) {
			if (currentDigest == null) {
				return;
			}
			if (currentDigest != newOwnership.sourceDigest) {
				throw transactionError("the ownership record has unknown bytes during rollback");
			}
			boundary.deleteFile(path);
			return;
		}
		if (currentDigest == oldOwnership.sourceDigest) {
			return;
		}
		if (currentDigest != null && currentDigest != newOwnership.sourceDigest) {
			throw transactionError("the ownership record has unknown bytes during rollback");
		}
		final oldBytes = requiredWorkspaceFile(GoExistingModuleOwnership.OLD_MANIFEST_NAME);
		if (GoExistingModuleOwnership.digest(oldBytes) != oldOwnership.sourceDigest) {
			throw transactionError("the ownership backup differs from the journal");
		}
		if (current != null) {
			boundary.deleteFile(path);
		}
		final restorePath = workspaceFile("rollback-ownership.json");
		boundary.saveBytes(restorePath, oldBytes);
		boundary.moveFile(restorePath, path);
	}

	function loadJournalOldOwnership(journal:GoOutputTransactionJournal):GoExistingModuleOwnershipSnapshot {
		if (journal.oldManifestSha256 == "") {
			return GoExistingModuleOwnership.empty();
		}
		final bytes = requiredWorkspaceFile(GoExistingModuleOwnership.OLD_MANIFEST_NAME);
		if (GoExistingModuleOwnership.digest(bytes) != journal.oldManifestSha256) {
			throw transactionError("the prior ownership backup differs from the transaction journal");
		}
		return GoExistingModuleOwnership.parse(project, boundary, bytes);
	}

	function verifySnapshot(snapshot:GoExistingModuleOwnershipSnapshot, explanation:String):Void {
		for (record in snapshot.files) {
			final bytes = boundary.readBytes(record.path);
			if (bytes == null || GoExistingModuleOwnership.digest(bytes) != record.sha256) {
				throw conflict(explanation);
			}
		}
	}

	function cleanupUnpublishedWorkspace():Void {
		if (boundary.readBytes(GoExistingModuleOwnership.journalTempPath(project)) != null) {
			boundary.deleteFile(GoExistingModuleOwnership.journalTempPath(project));
		}
		if (workspaceExists()) {
			validateWorkspaceIdentity();
			removeWorkspace();
		}
	}

	function cleanupControlFiles():Void {
		boundary.deleteFile(GoExistingModuleOwnership.journalPath(project));
		if (workspaceExists()) {
			validateWorkspaceIdentity();
			removeWorkspace();
		}
	}

	function validateWorkspaceIdentity():Void {
		if (!workspaceExists()) {
			throw transactionError("the transaction workspace is missing");
		}
		final marker = boundary.readBytes(workspaceFile(GoExistingModuleOwnership.WORKSPACE_MARKER_NAME));
		final expected = GoExistingModuleOwnership.renderProjectIdentity(project);
		if (marker == null || GoExistingModuleOwnership.digest(marker) != GoExistingModuleOwnership.digest(expected)) {
			throw transactionError("the transaction workspace identity is invalid");
		}
	}

	function requiredWorkspaceFile(path:String):Bytes {
		final bytes = boundary.readBytes(workspaceFile(path));
		if (bytes == null) {
			throw transactionError("a required transaction file is missing");
		}
		return bytes;
	}

	inline function workspaceFile(path:String):String {
		return GoExistingModuleOwnership.workspaceChild(project, path);
	}

	function workspaceExists():Bool {
		final path = Path.join([project.moduleRoot, GoExistingModuleOwnership.workspacePath(project)]);
		if (isSymbolicLink(path)) {
			throw transactionError("the transaction workspace is a symbolic link");
		}
		try {
			return FileSystem.exists(path);
		} catch (_:haxe.Exception) {
			throw transactionError("the transaction workspace could not be inspected");
		}
	}

	function removeWorkspace():Void {
		final path = Path.join([project.moduleRoot, GoExistingModuleOwnership.workspacePath(project)]);
		try {
			removeTree(path);
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw transactionError("the transaction workspace could not be removed");
		}
	}

	static function removeTree(path:String):Void {
		if (isSymbolicLink(path)) {
			throw transactionError("a transaction workspace entry is a symbolic link");
		}
		if (!FileSystem.exists(path)) {
			return;
		}
		if (FileSystem.isDirectory(path)) {
			for (name in FileSystem.readDirectory(path)) {
				removeTree(Path.join([path, name]));
			}
			FileSystem.deleteDirectory(path);
		} else {
			FileSystem.deleteFile(path);
		}
	}

	function renderJournal(journal:GoOutputTransactionJournal):Bytes {
		return Bytes.ofString([
			"{",
			'  "schemaVersion": ${JOURNAL_SCHEMA_VERSION},',
			'  "projectSha256": ${Json.stringify(journal.projectSha256)},',
			'  "oldManifestSha256": ${Json.stringify(journal.oldManifestSha256)},',
			'  "newManifestSha256": ${Json.stringify(journal.newManifestSha256)}',
			"}",
			""
		].join("\n"));
	}

	function parseJournal(bytes:Bytes):GoOutputTransactionJournal {
		final raw:RawGoOutputTransactionJournal = try {
			Json.parse(bytes.toString());
		} catch (_:haxe.Exception) {
			throw transactionError("the transaction journal is not valid JSON");
		}
		if (raw == null || Std.isOfType(raw, Array) || !Reflect.isObject(raw)) {
			throw transactionError("the transaction journal has an invalid shape");
		}
		final fields = Reflect.fields(raw);
		fields.sort(compareStrings);
		final expected = JOURNAL_FIELDS.copy();
		expected.sort(compareStrings);
		if (fields.length != expected.length) {
			throw transactionError("the transaction journal has an invalid shape");
		}
		for (index in 0...fields.length) {
			if (fields[index] != expected[index]) {
				throw transactionError("the transaction journal has an invalid shape");
			}
		}
		if (!Std.isOfType(Reflect.field(raw, "schemaVersion"), Int) || raw.schemaVersion != JOURNAL_SCHEMA_VERSION) {
			throw transactionError("the transaction journal schema is not supported");
		}
		for (field in ["projectSha256", "oldManifestSha256", "newManifestSha256"]) {
			if (!Std.isOfType(Reflect.field(raw, field), String)) {
				throw transactionError("the transaction journal has an invalid digest");
			}
		}
		final journal:GoOutputTransactionJournal = {
			projectSha256: raw.projectSha256,
			oldManifestSha256: raw.oldManifestSha256,
			newManifestSha256: raw.newManifestSha256
		};
		if (!isDigest(journal.projectSha256)
			|| (journal.oldManifestSha256 != "" && !isDigest(journal.oldManifestSha256))
			|| !isDigest(journal.newManifestSha256)) {
			throw transactionError("the transaction journal has an invalid digest");
		}
		final expectedProject = GoExistingModuleOwnership.digest(GoExistingModuleOwnership.renderProjectIdentity(project));
		if (journal.projectSha256 != expectedProject) {
			throw transactionError("the transaction journal belongs to another project shape");
		}
		return journal;
	}

	static inline function isDigest(value:String):Bool {
		return ~/^[0-9a-f]{64}$/.match(value);
	}

	static function compareStrings(left:String, right:String):Int {
		return left < right ? -1 : (left > right ? 1 : 0);
	}

	static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}

	static function conflict(explanation:String):GoOutputPathError {
		return new GoOutputPathError(GoOutputPathErrorKind.GeneratedFileConflict, explanation);
	}

	static function transactionError(explanation:String):GoOutputPathError {
		return new GoOutputPathError(GoOutputPathErrorKind.InterruptedTransaction, explanation);
	}
}
#else
class GoExistingModuleOutputTransaction {}
#end
