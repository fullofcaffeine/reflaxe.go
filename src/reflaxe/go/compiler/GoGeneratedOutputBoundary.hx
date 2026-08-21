package reflaxe.go.compiler;

#if (macro || eval)
import eval.luv.File.FileSync;
import eval.luv.Result;
import haxe.Json;
import haxe.io.Bytes;
import haxe.io.Path;
#if macro
import reflaxe.output.OutputManager;
import reflaxe.output.StringOrBytes;
#end
import sys.FileSystem;
import sys.io.File;

private typedef GoManagedFileMetadata = {
	final version:Int;
	final id:Int;
	final wasCached:Bool;
	final filesGenerated:Array<String>;
}

/**
	What: Stable categories for generated-output confinement failures.

	Why: Filesystem exceptions commonly include absolute host paths, while callers
	need actionable diagnostics that are safe to persist in snapshots and reports.

	How: The boundary converts validation, symlink, metadata, and write failures
	into one of these path-free categories before they reach compiler diagnostics.
**/
enum abstract GoOutputPathErrorKind(String) to String {
	final InvalidRelativePath = "GO-OUTPUT-PATH-001";
	final SymbolicLink = "GO-OUTPUT-PATH-002";
	final InvalidRoot = "GO-OUTPUT-PATH-003";
	final InvalidManagedMetadata = "GO-OUTPUT-PATH-004";
	final InvalidDestination = "GO-OUTPUT-PATH-005";
	final WriteFailed = "GO-OUTPUT-PATH-006";
	final ProtectedCallerFile = "GO-EXISTING-MODULE-MUTATION";
	final GeneratedFileConflict = "GO-GENERATED-FILE-CONFLICT";
	final InterruptedTransaction = "GO-OUTPUT-TRANSACTION";
}

/**
	What: A path-redacted generated-output failure.

	Why: Re-emitting the rejected value or resolved destination could expose a
	machine-local path and would make cross-host compiler evidence unstable.

	How: Retain only a stable error kind and a bounded explanation; the untrusted
	path and canonical root are intentionally never stored on the exception.
**/
class GoOutputPathError extends haxe.Exception {
	public final kind:GoOutputPathErrorKind;

	public function new(kind:GoOutputPathErrorKind, explanation:String) {
		this.kind = kind;
		super('[${kind}] Refused generated output: ${explanation}.');
	}
}

/**
	What: A compiler-relative output path that passed the cross-host lexical policy.

	Why: Plain strings make it too easy for a future writer to join an absolute,
	traversal, UNC, drive-relative, or host-dependent separator path to the output.

	How: Only `GoGeneratedOutputBoundary` can construct this value; writers receive
	the validated POSIX spelling and still perform canonical/symlink checks at use.
**/
abstract GoOutputRelativePath(String) {
	@:allow(reflaxe.go.compiler.GoGeneratedOutputBoundary)
	private inline function new(value:String) {
		this = value;
	}

	public inline function toString():String {
		return this;
	}

	public inline function segments():Array<String> {
		return this.split("/");
	}
}

/**
	What: The single typed write boundary for compiler-owned Go output.

	Why: Reflaxe's generic output manager accepts absolute and traversal paths and
	its managed-file metadata can name stale files for deletion. Go output must stay
	below the configured root even when a generated name or metadata file is wrong.

	How: Resolve the configured root once, require canonical relative path syntax,
	reject every symbolic link below that root (including broken links), prove each
	existing component remains contained, then create/write through this boundary.
	The configured root itself may have symlinked ancestors, but may not itself be a
	symlink. Concurrent hostile filesystem mutation is outside this process-local
	guard; untrusted Haxe compilation remains explicitly outside product admission.
**/
class GoGeneratedOutputBoundary {
	public static inline final MANAGED_FILE_METADATA = "_GeneratedFiles.json";

	final outputRoot:String;
	final canonicalRoot:String;
	final reservedPaths:Map<String, Bool>;
	final managedPaths:Map<String, Bool>;

	public function new(configuredRoot:String, ?callerOwnedPaths:Array<String>) {
		reservedPaths = [];
		managedPaths = [];
		if (callerOwnedPaths != null) {
			for (path in callerOwnedPaths) {
				reservedPaths.set(reservedPathKey(path), true);
			}
		}
		if (configuredRoot == null || StringTools.trim(configuredRoot) == "") {
			throw new GoOutputPathError(InvalidRoot, "the configured root is empty");
		}

		outputRoot = normalizeFilesystemPath(FileSystem.absolutePath(configuredRoot));
		if (isSymbolicLink(outputRoot)) {
			throw new GoOutputPathError(SymbolicLink, "the configured root is a symbolic link");
		}
		try {
			if (FileSystem.exists(outputRoot)) {
				if (!FileSystem.isDirectory(outputRoot)) {
					throw new GoOutputPathError(InvalidRoot, "the configured root is not a directory");
				}
			} else {
				FileSystem.createDirectory(outputRoot);
			}
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidRoot, "the configured root could not be prepared");
		}

		if (isSymbolicLink(outputRoot)) {
			throw new GoOutputPathError(SymbolicLink, "the configured root became a symbolic link");
		}
		try {
			canonicalRoot = normalizeFilesystemPath(FileSystem.fullPath(outputRoot));
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidRoot, "the configured root could not be resolved");
		}
	}

	/**
		What: Validates one destination without changing the filesystem.

		Why: Reflaxe extra-file keys and old managed-file entries must be rejected
		before any safe output is written or any stale file is deleted.

		How: Apply lexical validation, then inspect every existing or broken-link
		component below the canonical root.
	**/
	public function validateDestination(path:String):GoOutputRelativePath {
		final relative = validateRelativePath(path);
		if (reservedPaths.exists(reservedPathKey(relative.toString()))) {
			throw new GoOutputPathError(ProtectedCallerFile, "the destination is caller-owned in existing-module mode");
		}
		assertDestinationSafe(relative);
		return relative;
	}

	/**
		What: Preflights Reflaxe's old generated-file inventory.

		Why: The generic manager later deletes every remaining inventory entry; a
		poisoned traversal or symlink entry would otherwise widen that deletion.

		How: Validate the metadata destination itself, parse its closed shape, and
		prove every listed stale-file destination before generation begins.
	**/
	public function validateManagedFileMetadata():Void {
		final metadataRelative = validateDestination(MANAGED_FILE_METADATA);
		final metadataPath = absolutePath(metadataRelative);

		try {
			if (!FileSystem.exists(metadataPath)) {
				return;
			}
			final metadata:GoManagedFileMetadata = Json.parse(File.getContent(metadataPath));
			if (metadata == null || metadata.filesGenerated == null) {
				throw new GoOutputPathError(InvalidManagedMetadata, "managed-file metadata has an invalid shape");
			}
			for (managedPath in metadata.filesGenerated) {
				try {
					final relative = validateDestination(managedPath);
					managedPaths.set(reservedPathKey(relative.toString()), true);
				} catch (error:GoOutputPathError) {
					if (error.kind == ProtectedCallerFile) {
						throw error;
					}
					throw new GoOutputPathError(InvalidManagedMetadata, "managed-file metadata contains an unsafe destination");
				}
			}
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidManagedMetadata, "managed-file metadata could not be validated");
		}
	}

	/**
		Rejects replacement of a file that the current Reflaxe inventory does not own.

		This is the narrow first-generation collision rule for existing-module source
		files. The digest-backed transaction and cleanup authority remain owned by the
		later mixed-owner output contract.
	**/
	public function validateManagedReplacement(path:String):Void {
		final relative = validateDestination(path);
		final destination = absolutePath(relative);
		try {
			if (FileSystem.exists(destination) && !managedPaths.exists(reservedPathKey(relative.toString()))) {
				throw new GoOutputPathError(GeneratedFileConflict, "a generated source destination is caller-owned");
			}
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidDestination, "a generated source destination could not be inspected");
		}
	}

	/**
		What: Saves generated text through a direct filesystem writer.

		Why: The legacy iterator flow predates Reflaxe's output manager but must obey
		the same confinement contract.

		How: Prepare and re-check the destination, then convert any host write error
		to a path-redacted boundary error.
	**/
	public function saveContent(path:String, content:String):Void {
		final destination = prepareDestination(path);
		try {
			File.saveContent(destination.absolutePath, content);
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "the destination could not be written");
		}
	}

	/**
		Reads one confined regular file without treating absence as an error.

		Why
		Existing-module ownership must compare exact bytes before replacement or
		cleanup. A plain `File.getBytes` call would separate that check from the
		symlink and containment rules used by every generated-output effect.

		What
		Returns `null` when the confined path is absent and exact bytes when it is
		a regular file. Links and directories remain hard errors.

		How
		Validate the destination immediately before inspecting and reading it.
	**/
	public function readBytes(path:String):Null<Bytes> {
		final relative = validateDestination(path);
		final destination = absolutePath(relative);
		try {
			if (!FileSystem.exists(destination)) {
				return null;
			}
			if (FileSystem.isDirectory(destination)) {
				throw new GoOutputPathError(InvalidDestination, "a file destination is already a directory");
			}
			return File.getBytes(destination);
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidDestination, "a generated file could not be read");
		}
	}

	/** Save exact generated bytes through the confined direct writer. **/
	public function saveBytes(path:String, bytes:Bytes):Void {
		final destination = prepareDestination(path);
		try {
			File.saveBytes(destination.absolutePath, bytes);
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "the destination could not be written");
		}
	}

	/**
		Moves one compiler-owned file between two confined, already validated paths.

		The destination must be absent. This keeps replacement portable on hosts
		whose rename primitive does not replace an existing file.
	**/
	public function moveFile(sourcePath:String, targetPath:String):Void {
		final source = validateDestination(sourcePath);
		final target = prepareDestination(targetPath);
		final sourceAbsolute = absolutePath(source);
		try {
			if (!FileSystem.exists(sourceAbsolute) || FileSystem.isDirectory(sourceAbsolute)) {
				throw new GoOutputPathError(InvalidDestination, "the transaction source is not a regular file");
			}
			if (FileSystem.exists(target.absolutePath)) {
				throw new GoOutputPathError(InvalidDestination, "the transaction destination already exists");
			}
			FileSystem.rename(sourceAbsolute, target.absolutePath);
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "a generated file could not be moved");
		}
	}

	/** Delete one exact confined file; absence is idempotent. **/
	public function deleteFile(path:String):Void {
		final relative = validateDestination(path);
		final destination = absolutePath(relative);
		try {
			if (!FileSystem.exists(destination)) {
				return;
			}
			if (FileSystem.isDirectory(destination)) {
				throw new GoOutputPathError(InvalidDestination, "a file destination is already a directory");
			}
			FileSystem.deleteFile(destination);
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "a generated file could not be removed");
		}
	}

	/**
		What: Copies one compiler support file through the direct writer boundary.

		Why: Runtime copy helpers must not reconstruct unchecked target paths while
		walking a support directory.

		How: Prepare the target first, then copy with path-redacted failure handling.
	**/
	public function copyFile(sourcePath:String, targetPath:String):Void {
		final destination = prepareDestination(targetPath);
		try {
			File.copy(sourcePath, destination.absolutePath);
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "a compiler support file could not be copied");
		}
	}

	/**
		What: Saves generated content through Reflaxe after confinement succeeds.

		Why: Reflaxe owns change detection and managed-file recording, so the Go
		target must retain that behavior without exposing its permissive path parser.

		How: Prepare the destination using this root, then pass only the validated
		relative spelling to `OutputManager.saveFile`.
	**/
	#if macro
	public function saveFile(output:OutputManager, path:String, content:StringOrBytes):Void {
		final destination = prepareDestination(path);
		try {
			output.saveFile(destination.relativePath.toString(), content);
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "the destination could not be written");
		}
	}

	/**
		What: Copies a support file through Reflaxe's managed writer.

		Why: Reading and saving runtime files in the outer compiler would split path
		validation from the actual managed write.

		How: Prepare once, read the trusted compiler support file, and save only its
		validated relative destination through the output manager.
	**/
	public function copyManagedFile(output:OutputManager, sourcePath:String, targetPath:String):Void {
		final destination = prepareDestination(targetPath);
		try {
			output.saveFile(destination.relativePath.toString(), File.getContent(sourcePath));
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(WriteFailed, "a compiler support file could not be copied");
		}
	}
	#end

	function prepareDestination(path:String):{relativePath:GoOutputRelativePath, absolutePath:String} {
		final relative = validateDestination(path);
		final destination = absolutePath(relative);
		final parent = Path.directory(destination);
		try {
			if (!FileSystem.exists(parent)) {
				FileSystem.createDirectory(parent);
			}
		} catch (_:haxe.Exception) {
			throw new GoOutputPathError(InvalidDestination, "the destination parent could not be prepared");
		}
		assertDestinationSafe(relative);
		return {relativePath: relative, absolutePath: destination};
	}

	function assertDestinationSafe(relative:GoOutputRelativePath):Void {
		var current = outputRoot;
		final segments = relative.segments();
		for (index in 0...segments.length) {
			current = Path.join([current, segments[index]]);
			try {
				if (isSymbolicLink(current)) {
					throw new GoOutputPathError(SymbolicLink, "a destination component is a symbolic link");
				}
				if (!FileSystem.exists(current)) {
					continue;
				}
				if (index < segments.length - 1 && !FileSystem.isDirectory(current)) {
					throw new GoOutputPathError(InvalidDestination, "a destination parent is not a directory");
				}
				if (index == segments.length - 1 && FileSystem.isDirectory(current)) {
					throw new GoOutputPathError(InvalidDestination, "a file destination is already a directory");
				}
				final canonical = normalizeFilesystemPath(FileSystem.fullPath(current));
				if (!isWithin(canonical, canonicalRoot)) {
					throw new GoOutputPathError(InvalidDestination, "a destination resolves outside the configured root");
				}
			} catch (error:GoOutputPathError) {
				throw error;
			} catch (_:haxe.Exception) {
				throw new GoOutputPathError(InvalidDestination, "a destination component could not be resolved");
			}
		}
	}

	function absolutePath(relative:GoOutputRelativePath):String {
		return Path.join([outputRoot, relative.toString()]);
	}

	static function validateRelativePath(path:String):GoOutputRelativePath {
		if (path == null || path == "" || Path.isAbsolute(path) || StringTools.startsWith(path, "/") || StringTools.startsWith(path, "\\")
			|| ~/^[A-Za-z]:/.match(path) || path.indexOf("\\") != -1 || path.indexOf(":") != -1) {
			throw new GoOutputPathError(InvalidRelativePath, "the destination is not a canonical compiler-relative path");
		}

		final segments = path.split("/");
		for (segment in segments) {
			if (segment == "" || segment == "." || segment == ".." || StringTools.endsWith(segment, ".") || StringTools.endsWith(segment, " ")
				|| isWindowsDeviceName(segment) || containsControlCharacter(segment)) {
				throw new GoOutputPathError(InvalidRelativePath, "the destination is not a canonical compiler-relative path");
			}
		}
		if (StringTools.replace(Path.normalize(path), "\\", "/") != path) {
			throw new GoOutputPathError(InvalidRelativePath, "the destination is not a canonical compiler-relative path");
		}
		return new GoOutputRelativePath(path);
	}

	static function containsControlCharacter(value:String):Bool {
		for (index in 0...value.length) {
			final code = value.charCodeAt(index);
			if (code != null && (code < 32 || code == 127)) {
				return true;
			}
		}
		return false;
	}

	static function isWindowsDeviceName(segment:String):Bool {
		return ~/^(con|prn|aux|nul|com[1-9]|lpt[1-9])(\..*)?$/i.match(segment);
	}

	static function reservedPathKey(path:String):String {
		return path.toLowerCase();
	}

	static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}

	static function normalizeFilesystemPath(path:String):String {
		var normalized = StringTools.replace(Path.normalize(path), "\\", "/");
		while (normalized.length > 1 && StringTools.endsWith(normalized, "/") && !~/^[A-Za-z]:\/$/.match(normalized)) {
			normalized = normalized.substr(0, normalized.length - 1);
		}
		return normalized;
	}

	static function isWithin(candidate:String, root:String):Bool {
		var comparedCandidate = candidate;
		var comparedRoot = root;
		if (Sys.systemName() == "Windows") {
			comparedCandidate = comparedCandidate.toLowerCase();
			comparedRoot = comparedRoot.toLowerCase();
		}
		return comparedCandidate == comparedRoot || StringTools.startsWith(comparedCandidate, comparedRoot + "/");
	}
}
#else
class GoGeneratedOutputBoundary {}
#end
