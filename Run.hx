import haxe.Json;
import haxe.crypto.Sha256;
import haxe.io.Bytes;
import haxe.io.Path;
import eval.luv.File.FileSync;
import eval.luv.Result;
import sys.FileSystem;
import sys.io.File;

private typedef ReflaxeSourceMetadata = {
	final name:String;
	final abbv:String;
	final stdPaths:Array<String>;
}

/**
	What: the source-checkout Haxelib fields consumed by package staging.

	Why: Reflaxe's historical runner mutated an untyped JSON object, which made a
	misspelled or missing packaging field fail late and made metadata preservation
	implicit.

	How: JSON parsing crosses into this closed shape once, validation runs
	immediately, and the packaged manifest is reconstructed without `reflaxe`.
**/
private typedef SourceHaxelibManifest = {
	final name:String;
	final url:String;
	final license:String;
	final tags:Array<String>;
	final description:String;
	final version:String;
	final releasenote:String;
	final contributors:Array<String>;
	final classPath:String;
	final dependencies:haxe.DynamicAccess<String>;
	final reflaxe:ReflaxeSourceMetadata;
}

private typedef PackagedHaxelibManifest = {
	final name:String;
	final url:String;
	final license:String;
	final tags:Array<String>;
	final description:String;
	final version:String;
	final releasenote:String;
	final contributors:Array<String>;
	final classPath:String;
	final dependencies:haxe.DynamicAccess<String>;
}

private typedef PackageArchiveMetadata = {
	final compression:String;
	final fileMode:String;
	final ordering:String;
	final timestamp:String;
}

/**
	What: one deterministic source-authority to package-artifact mapping.

	Why: extension conversion and flattened roots must be independently auditable;
	a package file without its source and byte hashes cannot prove that mapping.

	How: entries use only repository-relative POSIX paths and SHA-256 hashes, then
	the runner sorts them by package path before writing the manifest.
**/
private typedef PackageManifestEntry = {
	final sourcePath:String;
	final packagePath:String;
	final kind:String;
	final sourceSha256:String;
	final packageSha256:String;
	final size:Int;
}

private typedef PackageManifest = {
	final schemaVersion:Int;
	final format:String;
	final archive:PackageArchiveMetadata;
	final classPath:String;
	final entries:Array<PackageManifestEntry>;
}

private typedef BuildOptions = {
	final sourceRoot:String;
	final packageRoot:String;
	final clean:Bool;
}

private enum abstract PackageEntryKind(String) to String {
	final ClassPath = "class-path";
	final Stdlib = "stdlib";
	final StdlibOverride = "stdlib-override";
	final Runtime = "runtime";
	final VendoredReflaxe = "vendored-reflaxe";
	final Metadata = "metadata";
	final PackageRunner = "package-runner";
	final Tooling = "tooling";
}

private class PackageBuildError extends haxe.Exception {
	public function new(message:String) {
		super(message);
	}
}

/**
	What: one standard-library root declared by source package metadata.

	Why: a directory merely named `_std` must not gain override semantics unless
	the package author explicitly declared that exact root.

	How: retain the normalized source root, its classification, and any nested
	declared roots that the generic Reflaxe flattening walk must skip.
**/
private class DeclaredStdPath {
	public final relativePath:String;
	public final sourceRoot:String;
	public final convertsOverrides:Bool;
	public var excludedRoots(default, null):Array<String> = [];

	public function new(relativePath:String, sourceRoot:String, convertsOverrides:Bool) {
		this.relativePath = relativePath;
		this.sourceRoot = sourceRoot;
		this.convertsOverrides = convertsOverrides;
	}
}

/**
	What: the fully validated, typed configuration for one package build.

	Why: package layout is release evidence. Inferring roots while copying could
	convert an undeclared directory or delete an unsafe destination.

	How: load the source manifest once, validate every repository-relative root,
	classify only declared paths whose final segment is `_std`, and reject outputs
	that overlap an input tree before any file is written.
**/
private class PackageBuildConfig {
	public final sourceRoot:String;
	public final packageRoot:String;
	public final classPath:String;
	public final classPathRoot:String;
	public final runtimeRoot:String;
	public final vendoredReflaxeRoot:String;
	public final stdPaths:Array<DeclaredStdPath>;
	public final sourceManifest:SourceHaxelibManifest;
	public final clean:Bool;

	public static function load(options:BuildOptions):PackageBuildConfig {
		final sourceRoot = PackagePathTools.absolute(options.sourceRoot);
		final packageRoot = PackagePathTools.absolute(options.packageRoot);
		PackagePathTools.requireDirectory(sourceRoot, "source root");

		final manifestPath = Path.join([sourceRoot, "haxelib.json"]);
		PackagePathTools.requireFile(manifestPath, "source haxelib.json");
		final sourceManifest = readSourceManifest(manifestPath);
		validateSourceManifest(sourceManifest);

		final classPath = PackagePathTools.safeRelative(sourceManifest.classPath, "classPath");
		if (classPath != "src") {
			throw new PackageBuildError('haxelib classPath must be "src" for canonical package staging; found "${classPath}"');
		}
		final classPathRoot = Path.join([sourceRoot, classPath]);
		final runtimeRoot = Path.join([sourceRoot, "runtime"]);
		final vendoredReflaxeRoot = Path.join([sourceRoot, "vendor", "reflaxe"]);
		PackagePathTools.requireDirectory(classPathRoot, "classPath root");
		PackagePathTools.requireDirectory(runtimeRoot, "runtime root");
		PackagePathTools.requireDirectory(vendoredReflaxeRoot, "vendored Reflaxe root");

		final stdPaths:Array<DeclaredStdPath> = [];
		final seenStdPaths:Map<String, Bool> = [];
		for (declared in sourceManifest.reflaxe.stdPaths) {
			final relative = PackagePathTools.safeRelative(declared, "reflaxe.stdPaths entry");
			if (seenStdPaths.exists(relative)) {
				throw new PackageBuildError('duplicate reflaxe.stdPaths entry: ${relative}');
			}
			seenStdPaths.set(relative, true);
			final source = Path.join([sourceRoot, relative]);
			PackagePathTools.requireDirectory(source, 'declared std path "${relative}"');
			final segments = relative.split("/");
			final convertsOverrides = segments[segments.length - 1] == "_std";
			stdPaths.push(new DeclaredStdPath(relative, source, convertsOverrides));
		}

		for (current in stdPaths) {
			for (candidate in stdPaths) {
				if (candidate != current && PackagePathTools.isWithin(candidate.sourceRoot, current.sourceRoot)) {
					current.excludedRoots.push(candidate.sourceRoot);
				}
			}
			current.excludedRoots.sort(PackagePathTools.compareUtf8);
		}

		final inputRoots = [classPathRoot, runtimeRoot, vendoredReflaxeRoot];
		for (stdPath in stdPaths) {
			inputRoots.push(stdPath.sourceRoot);
		}
		for (inputRoot in inputRoots) {
			if (PackagePathTools.isWithin(packageRoot, inputRoot)) {
				throw new PackageBuildError("package output must not overlap an input tree");
			}
		}
		if (PackagePathTools.isWithin(sourceRoot, packageRoot)) {
			throw new PackageBuildError("package output must not contain the source root");
		}

		for (required in [
			"LICENSE",
			"LICENSING.md",
			"README.md",
			"extraParams.hxml",
			"license-policy.json",
			"licenses/HAXE-GO-GENERATED-MIT.txt",
			"licenses/HAXE-STDLIB-MIT.txt",
			"Run.hx",
			"scripts/dev/go-hx.sh",
			"scripts/dev/haxe_go_watch.py"
		]) {
			PackagePathTools.requireFile(Path.join([sourceRoot, required]), 'required package file "${required}"');
		}
		for (required in ["haxelib.json", "LICENSE", "PATCHES.md", "FUTURE_MODIFICATIONS.md"]) {
			PackagePathTools.requireFile(Path.join([vendoredReflaxeRoot, required]), 'required vendored Reflaxe file "${required}"');
		}

		return new PackageBuildConfig(sourceRoot, packageRoot, classPath, classPathRoot, runtimeRoot, vendoredReflaxeRoot, stdPaths, sourceManifest,
			options.clean);
	}

	function new(sourceRoot:String, packageRoot:String, classPath:String, classPathRoot:String, runtimeRoot:String, vendoredReflaxeRoot:String,
			stdPaths:Array<DeclaredStdPath>, sourceManifest:SourceHaxelibManifest, clean:Bool) {
		this.sourceRoot = sourceRoot;
		this.packageRoot = packageRoot;
		this.classPath = classPath;
		this.classPathRoot = classPathRoot;
		this.runtimeRoot = runtimeRoot;
		this.vendoredReflaxeRoot = vendoredReflaxeRoot;
		this.stdPaths = stdPaths;
		this.sourceManifest = sourceManifest;
		this.clean = clean;
	}

	static function readSourceManifest(path:String):SourceHaxelibManifest {
		try {
			final parsed:SourceHaxelibManifest = Json.parse(File.getContent(path));
			return parsed;
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("could not read or parse haxelib.json");
		}
	}

	static function validateSourceManifest(manifest:SourceHaxelibManifest):Void {
		if (manifest == null) {
			throw new PackageBuildError("haxelib.json must contain an object");
		}
		for (field in [
			manifest.name,
			manifest.url,
			manifest.license,
			manifest.description,
			manifest.version,
			manifest.releasenote,
			manifest.classPath
		]) {
			if (field == null || StringTools.trim(field) == "") {
				throw new PackageBuildError("haxelib.json contains a missing or empty required string field");
			}
		}
		if (manifest.tags == null || manifest.contributors == null || manifest.dependencies == null) {
			throw new PackageBuildError("haxelib.json tags, contributors, and dependencies are required");
		}
		if (manifest.reflaxe == null || manifest.reflaxe.stdPaths == null || manifest.reflaxe.stdPaths.length == 0) {
			throw new PackageBuildError("haxelib.json must declare at least one reflaxe.stdPaths entry");
		}
		if (manifest.reflaxe.name == null || manifest.reflaxe.abbv == null) {
			throw new PackageBuildError("haxelib.json reflaxe name and abbreviation are required");
		}
	}
}

/**
	What: canonical path, traversal, and recursive-clean operations for staging.

	Why: platform separators and ambiguous relative paths must never alter package
	member names or widen the directory a clean build may remove.

	How: normalize filesystem paths at the boundary, expose only validated POSIX
	relative paths to the manifest, and compare traversal names by UTF-8 bytes.
**/
private class PackagePathTools {
	public static function absolute(path:String):String {
		return Path.normalize(FileSystem.absolutePath(path));
	}

	public static function safeRelative(value:String, label:String):String {
		if (value == null || value == "") {
			throw new PackageBuildError('${label} must be a safe repository-relative path');
		}
		final portable = StringTools.replace(value, "\\", "/");
		final normalized = StringTools.replace(Path.normalize(portable), "\\", "/");
		final driveAbsolute = ~/^[A-Za-z]:/.match(portable);
		if (Path.isAbsolute(portable)
			|| driveAbsolute
			|| portable.indexOf(":") != -1
			|| normalized == "."
			|| normalized == ".."
			|| StringTools.startsWith(normalized, "../")
			|| normalized != portable
			|| portable.split("/").contains("")) {
			throw new PackageBuildError('${label} must be a safe repository-relative path');
		}
		for (segment in normalized.split("/")) {
			if (segment == "." || segment == ".." || StringTools.endsWith(segment, ".") || StringTools.endsWith(segment, " ")
				|| isWindowsDeviceName(segment) || containsControlCharacter(segment)) {
				throw new PackageBuildError('${label} must be a safe repository-relative path');
			}
		}
		return normalized;
	}

	public static function requireDirectory(path:String, label:String):Void {
		try {
			if (FileSystem.exists(path) && FileSystem.isDirectory(path) && !isSymbolicLink(path)) {
				return;
			}
		} catch (_:haxe.Exception) {}
		throw new PackageBuildError('${label} is not a directory');
	}

	public static function requireFile(path:String, label:String):Void {
		try {
			if (FileSystem.exists(path) && !FileSystem.isDirectory(path) && !isSymbolicLink(path)) {
				return;
			}
		} catch (_:haxe.Exception) {}
		throw new PackageBuildError('${label} is not a file');
	}

	public static function isWithin(candidate:String, parent:String):Bool {
		var normalizedCandidate = StringTools.replace(Path.normalize(candidate), "\\", "/");
		var normalizedParent = Path.removeTrailingSlashes(StringTools.replace(Path.normalize(parent), "\\", "/"));
		if (Sys.systemName() == "Windows") {
			normalizedCandidate = normalizedCandidate.toLowerCase();
			normalizedParent = normalizedParent.toLowerCase();
		}
		return normalizedCandidate == normalizedParent || StringTools.startsWith(normalizedCandidate, normalizedParent + "/");
	}

	public static function relativeTo(path:String, root:String):String {
		final normalizedPath = StringTools.replace(Path.normalize(path), "\\", "/");
		final normalizedRoot = Path.removeTrailingSlashes(StringTools.replace(Path.normalize(root), "\\", "/"));
		if (!StringTools.startsWith(normalizedPath, normalizedRoot + "/")) {
			throw new PackageBuildError("package source path is outside its declared root");
		}
		return safeRelative(normalizedPath.substr(normalizedRoot.length + 1), "generated package path");
	}

	public static function compareUtf8(left:String, right:String):Int {
		final leftBytes = Bytes.ofString(left);
		final rightBytes = Bytes.ofString(right);
		final commonLength = leftBytes.length < rightBytes.length ? leftBytes.length : rightBytes.length;
		for (index in 0...commonLength) {
			final delta = leftBytes.get(index) - rightBytes.get(index);
			if (delta != 0) {
				return delta;
			}
		}
		return leftBytes.length - rightBytes.length;
	}

	public static function ensureParent(path:String):Void {
		final parent = Path.directory(path);
		if (parent == null || parent == "" || FileSystem.exists(parent)) {
			return;
		}
		ensureParent(parent);
		FileSystem.createDirectory(parent);
	}

	/**
		What: Resolves one package member below a freshly prepared staging root.

		Why: A validated archive name is insufficient if an existing destination
		component is a symlink or canonically resolves outside the staging tree.

		How: Revalidate the relative spelling, reject all destination symlinks,
		prove existing components against the canonical root, prepare parents, and
		repeat the proof immediately before the caller writes.
	**/
	public static function confinedDestination(root:String, relative:String):String {
		final safe = safeRelative(StringTools.replace(relative, "\\", "/"), "package path");
		if (!FileSystem.exists(root) || !FileSystem.isDirectory(root) || isSymbolicLink(root)) {
			throw new PackageBuildError("package output root is not a plain directory");
		}
		final canonicalRoot = canonicalExisting(root);
		validateDestinationComponents(root, safe, canonicalRoot);
		final destination = Path.join([root, safe]);
		try {
			ensureParent(destination);
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package output parent could not be prepared");
		}
		validateDestinationComponents(root, safe, canonicalRoot);
		return destination;
	}

	static function validateDestinationComponents(root:String, relative:String, canonicalRoot:String):Void {
		try {
			var current = root;
			final segments = relative.split("/");
			for (index in 0...segments.length) {
				current = Path.join([current, segments[index]]);
				if (isSymbolicLink(current)) {
					throw new PackageBuildError("package output contains a path outside its canonical root");
				}
				if (!FileSystem.exists(current)) {
					continue;
				}
				if (!isWithin(canonicalExisting(current), canonicalRoot)) {
					throw new PackageBuildError("package output contains a path outside its canonical root");
				}
				if (index < segments.length - 1 && !FileSystem.isDirectory(current)) {
					throw new PackageBuildError("package output parent is not a directory");
				}
				if (index == segments.length - 1 && FileSystem.isDirectory(current)) {
					throw new PackageBuildError("package file destination is already a directory");
				}
			}
		} catch (error:PackageBuildError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package output destination could not be validated");
		}
	}

	/**
		What: Deletes a package staging tree only after a complete confinement pass.

		Why: Recursive deletion through a directory symlink can mutate files outside
		the user-selected package root before the builder has written anything.

		How: Resolve the canonical root, reject every symlink or escaped descendant
		in a read-only preflight, then perform the already-proven recursive deletion.
	**/
	public static function deleteTree(path:String):Void {
		if (isSymbolicLink(path)) {
			throw new PackageBuildError("package output contains a path outside its canonical root");
		}
		if (!FileSystem.exists(path)) {
			return;
		}
		try {
			final canonicalRoot = canonicalExisting(path);
			validateDeleteTree(path, canonicalRoot);
			deleteValidatedTree(path);
		} catch (error:PackageBuildError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package output tree could not be cleaned");
		}
	}

	static function validateDeleteTree(path:String, canonicalRoot:String):Void {
		if (isSymbolicLink(path) || !isWithin(canonicalExisting(path), canonicalRoot)) {
			throw new PackageBuildError("package output contains a path outside its canonical root");
		}
		if (!FileSystem.isDirectory(path)) {
			return;
		}
		final entries = FileSystem.readDirectory(path);
		entries.sort(compareUtf8);
		for (entry in entries) {
			validateDeleteTree(Path.join([path, entry]), canonicalRoot);
		}
	}

	static function deleteValidatedTree(path:String):Void {
		if (!FileSystem.isDirectory(path)) {
			FileSystem.deleteFile(path);
			return;
		}
		final entries = FileSystem.readDirectory(path);
		entries.sort(compareUtf8);
		for (entry in entries) {
			deleteValidatedTree(Path.join([path, entry]));
		}
		FileSystem.deleteDirectory(path);
	}

	public static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}

	static function canonicalExisting(path:String):String {
		try {
			return StringTools.replace(Path.normalize(FileSystem.fullPath(path)), "\\", "/");
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package path could not be canonically resolved");
		}
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
}

/**
	What: the deterministic implementation of the established Reflaxe build copy.

	Why: release packages must flatten declared std roots without also capturing
	untracked caches, editor files, or undeclared target assets.

	How: copy compiler `.hx`, runtime `.go`, vendored `.hx`, and explicit metadata
	allowlists; reserve every output path once; then hash and sort the mappings.
**/
private class PackageBuilder {
	static inline final PACKAGE_MANIFEST = "reflaxe-package-manifest.json";

	final config:PackageBuildConfig;
	final entries:Array<PackageManifestEntry> = [];
	final pathOwners:Map<String, String> = [];

	public function new(config:PackageBuildConfig) {
		this.config = config;
	}

	public function build():Void {
		prepareOutput();
		copyTree(config.classPathRoot, config.classPath, PackageEntryKind.ClassPath, [], ".hx");
		for (stdPath in config.stdPaths) {
			copyStdPath(stdPath);
		}
		copyTree(config.runtimeRoot, "runtime", PackageEntryKind.Runtime, [], ".go");
		copyTree(Path.join([config.vendoredReflaxeRoot, "src"]), "vendor/reflaxe/src", PackageEntryKind.VendoredReflaxe, [], ".hx");
		copyVendoredReflaxeFile("haxelib.json");
		copyVendoredReflaxeFile("LICENSE");
		copyVendoredReflaxeFile("PATCHES.md");
		copyVendoredReflaxeFile("FUTURE_MODIFICATIONS.md");
		copyRequiredFile("LICENSE", PackageEntryKind.Metadata);
		copyRequiredFile("LICENSING.md", PackageEntryKind.Metadata);
		copyRequiredFile("README.md", PackageEntryKind.Metadata);
		copyRequiredFile("extraParams.hxml", PackageEntryKind.Metadata);
		copyRequiredFile("license-policy.json", PackageEntryKind.Metadata);
		copyRequiredFile("licenses/HAXE-GO-GENERATED-MIT.txt", PackageEntryKind.Metadata);
		copyRequiredFile("licenses/HAXE-STDLIB-MIT.txt", PackageEntryKind.Metadata);
		copyRequiredFile("Run.hx", PackageEntryKind.PackageRunner);
		copyDevelopmentTooling();
		writePackagedHaxelib();
		writePackageManifest();
	}

	/**
		What: ships the canonical one-shot and managed-watch commands under the
		installed package's `tools/` directory.

		Why: scaffolded projects resolve their compiler through Haxelib and may not
		have this source checkout beside them. Omitting the commands would make the
		documented `npm run dev` work only for repository contributors.

		How: copy the two explicit source authorities and record each mapping as
		tooling provenance in the deterministic package manifest.
	**/
	function copyDevelopmentTooling():Void {
		copyMappedFile(Path.join([config.sourceRoot, "scripts/dev/go-hx.sh"]), "tools/go-hx.sh", PackageEntryKind.Tooling);
		copyMappedFile(Path.join([config.sourceRoot, "scripts/dev/haxe_go_watch.py"]), "tools/haxe_go_watch.py", PackageEntryKind.Tooling);
	}

	function prepareOutput():Void {
		try {
			if (PackagePathTools.isSymbolicLink(config.packageRoot)) {
				throw new PackageBuildError("package output contains a path outside its canonical root");
			}
			if (FileSystem.exists(config.packageRoot)) {
				if (!FileSystem.isDirectory(config.packageRoot)) {
					throw new PackageBuildError("package output already exists as a file");
				}
				if (!config.clean) {
					throw new PackageBuildError("package output already exists; pass --clean to replace it");
				}
				PackagePathTools.deleteTree(config.packageRoot);
			}
			PackagePathTools.ensureParent(config.packageRoot);
			FileSystem.createDirectory(config.packageRoot);
		} catch (error:PackageBuildError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package output root could not be prepared");
		}
	}

	function copyStdPath(stdPath:DeclaredStdPath):Void {
		final files = collectFiles(stdPath.sourceRoot, stdPath.excludedRoots);
		for (source in files) {
			final sourceRelative = PackagePathTools.relativeTo(source, stdPath.sourceRoot);
			if (!StringTools.endsWith(sourceRelative, ".hx")) {
				continue;
			}
			if (StringTools.endsWith(sourceRelative, ".cross.hx")) {
				throw new PackageBuildError('source std paths must contain ordinary .hx authority, not .cross.hx: ${sourceRelative}');
			}
			final packageRelative = if (stdPath.convertsOverrides) {
				Path.withoutExtension(sourceRelative) + ".cross.hx";
			} else {
				sourceRelative;
			}
			final kind = stdPath.convertsOverrides ? PackageEntryKind.StdlibOverride : PackageEntryKind.Stdlib;
			copyMappedFile(source, Path.join([config.classPath, packageRelative]), kind);
		}
	}

	function copyTree(sourceRoot:String, packagePrefix:String, kind:PackageEntryKind, excludedRoots:Array<String>, requiredSuffix:Null<String>):Void {
		for (source in collectFiles(sourceRoot, excludedRoots)) {
			final sourceRelative = PackagePathTools.relativeTo(source, sourceRoot);
			if (requiredSuffix != null && !StringTools.endsWith(sourceRelative, requiredSuffix)) {
				continue;
			}
			if (isRepositoryOnlyRuntimeTest(sourceRelative, kind)) {
				continue;
			}
			if ((kind == PackageEntryKind.ClassPath || kind == PackageEntryKind.Stdlib)
				&& StringTools.endsWith(sourceRelative, ".cross.hx")) {
				throw new PackageBuildError('source class paths must not contain generated .cross.hx files: ${sourceRelative}');
			}
			copyMappedFile(source, Path.join([packagePrefix, sourceRelative]), kind);
		}
	}

	/**
		What: identifies Go tests that belong only to this compiler repository.

		Why: Haxelib consumers need the runtime implementation, but shipping hxrt's
		unit tests would add development-only processes and fixtures to every install.

		How: exclude only conventional `_test.go` files while walking the explicitly
		classified runtime tree; ordinary runtime `.go` files remain packaged.
	**/
	function isRepositoryOnlyRuntimeTest(sourceRelative:String, kind:PackageEntryKind):Bool {
		return kind == PackageEntryKind.Runtime && StringTools.endsWith(sourceRelative, "_test.go");
	}

	function collectFiles(root:String, excludedRoots:Array<String>):Array<String> {
		final files:Array<String> = [];
		visit(root, excludedRoots, files);
		files.sort(PackagePathTools.compareUtf8);
		return files;
	}

	function visit(directory:String, excludedRoots:Array<String>, files:Array<String>):Void {
		final children = FileSystem.readDirectory(directory);
		children.sort(PackagePathTools.compareUtf8);
		for (child in children) {
			final path = Path.normalize(Path.join([directory, child]));
			if (PackagePathTools.isSymbolicLink(path)) {
				throw new PackageBuildError("package source trees must not contain symbolic links");
			}
			if (excludedRoots.contains(path)) {
				continue;
			}
			if (FileSystem.isDirectory(path)) {
				visit(path, excludedRoots, files);
			} else {
				files.push(path);
			}
		}
	}

	function copyRequiredFile(relative:String, kind:PackageEntryKind):Void {
		copyMappedFile(Path.join([config.sourceRoot, relative]), relative, kind);
	}

	function copyVendoredReflaxeFile(relative:String):Void {
		copyMappedFile(Path.join([config.vendoredReflaxeRoot, relative]), Path.join(["vendor/reflaxe", relative]), PackageEntryKind.VendoredReflaxe);
	}

	function copyMappedFile(source:String, packageRelative:String, kind:PackageEntryKind):Void {
		final safePackagePath = PackagePathTools.safeRelative(StringTools.replace(packageRelative, "\\", "/"), "package path");
		reservePath(safePackagePath, PackagePathTools.relativeTo(source, config.sourceRoot));
		final destination = PackagePathTools.confinedDestination(config.packageRoot, safePackagePath);
		try {
			File.copy(source, destination);
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package file could not be copied");
		}
		recordEntry(source, destination, safePackagePath, kind);
	}

	function reservePath(packagePath:String, sourcePath:String):Void {
		final previous = pathOwners.get(packagePath);
		if (previous != null) {
			throw new PackageBuildError('package path collision at ${packagePath}: ${previous} and ${sourcePath}');
		}
		pathOwners.set(packagePath, sourcePath);
	}

	function recordEntry(source:String, destination:String, packagePath:String, kind:PackageEntryKind):Void {
		final packageBytes = File.getBytes(destination);
		entries.push({
			sourcePath: PackagePathTools.relativeTo(source, config.sourceRoot),
			packagePath: packagePath,
			kind: kind,
			sourceSha256: digest(source),
			packageSha256: Sha256.make(packageBytes).toHex(),
			size: packageBytes.length
		});
	}

	function writePackagedHaxelib():Void {
		final dependencies:haxe.DynamicAccess<String> = {};
		final dependencyNames = [for (name in config.sourceManifest.dependencies.keys()) name];
		dependencyNames.sort(PackagePathTools.compareUtf8);
		for (name in dependencyNames) {
			final requirement = config.sourceManifest.dependencies.get(name);
			if (requirement == null) {
				throw new PackageBuildError('haxelib dependency "${name}" must have a string requirement');
			}
			dependencies.set(name, requirement);
		}
		final packaged:PackagedHaxelibManifest = {
			name: config.sourceManifest.name,
			url: config.sourceManifest.url,
			license: config.sourceManifest.license,
			tags: config.sourceManifest.tags.copy(),
			description: config.sourceManifest.description,
			version: config.sourceManifest.version,
			releasenote: config.sourceManifest.releasenote,
			contributors: config.sourceManifest.contributors.copy(),
			classPath: config.classPath,
			dependencies: dependencies
		};
		final packagePath = "haxelib.json";
		final source = Path.join([config.sourceRoot, packagePath]);
		reservePath(packagePath, packagePath);
		final destination = PackagePathTools.confinedDestination(config.packageRoot, packagePath);
		try {
			File.saveContent(destination, Json.stringify(packaged, null, "  ") + "\n");
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package metadata could not be written");
		}
		recordEntry(source, destination, packagePath, PackageEntryKind.Metadata);
	}

	function writePackageManifest():Void {
		entries.sort((left, right) -> {
			final packageOrder = PackagePathTools.compareUtf8(left.packagePath, right.packagePath);
			return packageOrder == 0 ? PackagePathTools.compareUtf8(left.sourcePath, right.sourcePath) : packageOrder;
		});
		final manifest:PackageManifest = {
			schemaVersion: 1,
			format: "reflaxe.go-haxelib-package",
			archive: {
				compression: "stored",
				fileMode: "0644",
				ordering: "utf8-bytewise",
				timestamp: "2000-01-01T00:00:00Z"
			},
			classPath: config.classPath,
			entries: entries
		};
		final destination = PackagePathTools.confinedDestination(config.packageRoot, PACKAGE_MANIFEST);
		try {
			File.saveContent(destination, Json.stringify(manifest, null, "  ") + "\n");
		} catch (_:haxe.Exception) {
			throw new PackageBuildError("package manifest could not be written");
		}
	}

	static function digest(path:String):String {
		return Sha256.make(File.getBytes(path)).toHex();
	}
}

/**
	What: stages the canonical Reflaxe.Go Haxelib package.

	Why: source overrides are ordinary `.hx` files, while Haxelib selection needs
	only the explicitly declared `_std` root flattened as `.cross.hx`.

	How: retain Reflaxe `Run.buildProject` semantics for the `build` command,
	declared `stdPaths`, `_std` conversion, metadata sanitation, and the historical
	`--deleteOldFolder` alias. Local adaptations add typed configuration, source
	allowlists, UTF-8 ordering, collision rejection, a hash manifest, and a
	deterministic ZIP companion. Invoke it with
	`haxe --run Run build <output> --source-root <checkout> --clean`.
**/
class Run {
	static function main():Void {
		try {
			final options = parseOptions(Sys.args());
			final config = PackageBuildConfig.load(options);
			new PackageBuilder(config).build();
			Sys.println('[package-runner] staged ${config.packageRoot}');
		} catch (error:haxe.Exception) {
			Sys.stderr().writeString('[package-runner] ERROR: ${error.message}\n');
			Sys.exit(2);
		}
	}

	static function parseOptions(args:Array<String>):BuildOptions {
		if (args.length < 2 || args.shift() != "build") {
			throw new PackageBuildError("usage: haxe --run Run build <output> [--source-root <checkout>] [--clean]");
		}
		final packageRoot = args.shift();
		var sourceRoot = Sys.getCwd();
		var clean = false;
		while (args.length > 0) {
			final option = args.shift();
			switch (option) {
				case "--source-root":
					if (args.length == 0) {
						throw new PackageBuildError("--source-root requires a path");
					}
					sourceRoot = args.shift();
				case "--clean" | "--deleteOldFolder":
					clean = true;
				case _:
					throw new PackageBuildError("unknown package runner option");
			}
		}
		return {
			sourceRoot: sourceRoot,
			packageRoot: packageRoot,
			clean: clean
		};
	}
}
