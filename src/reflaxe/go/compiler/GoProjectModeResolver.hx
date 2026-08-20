package reflaxe.go.compiler;

#if (macro || eval)
import eval.luv.File.FileSync;
import eval.luv.Result;
import haxe.Json;
import haxe.io.Path;
import haxe.macro.Context;
import reflaxe.go.ast.GoPackageName;
import reflaxe.go.compiler.GoProjectMode.ExistingGoModuleProject;
import reflaxe.go.compiler.GoProjectMode.GoBuildPolicy;
import reflaxe.go.compiler.GoProjectMode.GoEntrypointSymbol;
import reflaxe.go.compiler.GoProjectMode.GoEntrypointPolicy;
import reflaxe.go.compiler.GoProjectMode.GoProjectRelativePath;
import sys.FileSystem;
import sys.io.File;

private typedef RawEntrypoint = {
	final kind:String;
	final ?symbol:String;
}

private typedef RawBuild = {
	final kind:String;
}

private typedef RawProjectManifest = {
	final schemaVersion:Int;
	final mode:String;
	final moduleRoot:String;
	final packageDir:String;
	final packageName:String;
	final runtimeDir:String;
	final entrypoint:RawEntrypoint;
	final build:RawBuild;
}

/** Stable, path-free categories for existing-module configuration errors. */
enum abstract GoProjectModeErrorKind(String) to String {
	final InvalidManifest = "GO-EXISTING-MODULE-MANIFEST";
	final InvalidModuleFile = "GO-EXISTING-MODULE-FILE";
	final ConfigurationConflict = "GO-EXISTING-MODULE-CONFLICT";
	final UnsupportedProjectShape = "GO-EXISTING-MODULE-UNSUPPORTED";
	final InvalidPackageName = "GO-PACKAGE-NAME";
	final InvalidPackageDirectory = "GO-PACKAGE-DIR";
	final PackageMismatch = "GO-PACKAGE-MISMATCH";
	final EntrypointOwnership = "GO-ENTRYPOINT-OWNERSHIP";
}

/** A project error that can safely appear in portable compiler diagnostics. */
class GoProjectModeError extends haxe.Exception {
	public final kind:GoProjectModeErrorKind;

	public function new(kind:GoProjectModeErrorKind, explanation:String) {
		this.kind = kind;
		super('[${kind}] Existing Go module configuration is invalid: ${explanation}.');
	}
}

/**
	Parses the JSON manifest once and immediately narrows it to Haxe domain types.

	`Reflect` is confined to this parser because Haxe's JSON API returns host data.
	Every field is allowlisted and type-checked before typed fields are read.
**/
class GoProjectModeResolver {
	static final TOP_LEVEL_FIELDS = [
		"schemaVersion",
		"mode",
		"moduleRoot",
		"packageDir",
		"packageName",
		"runtimeDir",
		"entrypoint",
		"build"
	];

	public static function resolve():GoProjectMode {
		final manifestDefine = Context.definedValue(GoCompilerDefine.DefineGoProject);
		if (manifestDefine == null || StringTools.trim(manifestDefine) == "") {
			return Standalone;
		}

		try {
			return ExistingModule(parseManifest(StringTools.trim(manifestDefine)));
		} catch (error:GoProjectModeError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest could not be read or parsed");
		}
	}

	static function parseManifest(configuredPath:String):ExistingGoModuleProject {
		final manifestPath = canonicalRegularFile(configuredPath, InvalidManifest, "the project manifest is not a regular file");
		final raw:RawProjectManifest = try {
			Json.parse(File.getContent(manifestPath));
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest is not valid JSON");
		}

		validateObject(raw, TOP_LEVEL_FIELDS, TOP_LEVEL_FIELDS, "the project manifest has an invalid shape");
		validateInt(raw, "schemaVersion");
		for (field in ["mode", "moduleRoot", "packageDir", "packageName", "runtimeDir"]) {
			validateString(raw, field);
		}
		validateObject(raw.entrypoint, ["kind", "symbol"], ["kind"], "entrypoint has an invalid shape");
		validateString(raw.entrypoint, "kind");
		if (Reflect.hasField(raw.entrypoint, "symbol")) {
			validateString(raw.entrypoint, "symbol");
		}
		validateObject(raw.build, ["kind"], ["kind"], "build has an invalid shape");
		validateString(raw.build, "kind");

		if (raw.schemaVersion != 1 || raw.mode != "existing-module") {
			throw new GoProjectModeError(InvalidManifest, "the manifest schema or mode is not supported");
		}

		final moduleRoot = resolveModuleRoot(manifestPath, raw.moduleRoot);
		final packageDir = resolveProjectDirectory(moduleRoot, raw.packageDir, "packageDir");
		final runtimeDir = resolveProjectDirectory(moduleRoot, raw.runtimeDir, "runtimeDir");
		if (packageDir.value().toLowerCase() == runtimeDir.value().toLowerCase()) {
			throw new GoProjectModeError(InvalidPackageDirectory, "packageDir and runtimeDir must identify different Go packages");
		}
		if (!GoPackageName.isIdentifier(raw.packageName)) {
			throw new GoProjectModeError(InvalidPackageName, "packageName is not a valid Go package identifier");
		}
		final packageName = GoPackageName.named(raw.packageName);

		final entrypoint = switch (raw.entrypoint.kind) {
			case "compiler-main" if (!Reflect.hasField(raw.entrypoint, "symbol")):
				if (packageName.value() != "main") {
					throw new GoProjectModeError(InvalidPackageName, "compiler-main requires packageName main");
				}
				GoEntrypointPolicy.CompilerMain;
			case "caller-bridge" if (raw.entrypoint.symbol != null):
				if (!GoPackageName.isIdentifier(raw.entrypoint.symbol)) {
					throw new GoProjectModeError(InvalidManifest, "caller-bridge symbol is not a valid Go identifier");
				}
				final symbol = GoEntrypointSymbol.named(raw.entrypoint.symbol);
				GoEntrypointPolicy.CallerBridge(symbol);
			case _:
				throw new GoProjectModeError(InvalidManifest, "entrypoint has an unsupported kind or fields");
		};
		if (raw.build.kind != "none") {
			throw new GoProjectModeError(UnsupportedProjectShape, "typed Go build requests are implemented by a later compatibility slice");
		}

		GoPackageDirectoryInspector.validate(moduleRoot, packageDir, packageName, entrypoint);
		validateLegacyDefines(moduleRoot, packageDir);
		final modulePath = readModulePath(moduleRoot);
		final assertedModule = Context.definedValue(GoCompilerDefine.DefineGoModule);
		if (assertedModule != null && StringTools.trim(assertedModule) != "" && StringTools.trim(assertedModule) != modulePath) {
			throw new GoProjectModeError(ConfigurationConflict, "go_module does not match the caller module path");
		}

		return new ExistingGoModuleProject({
			manifestPath: manifestPath,
			moduleRoot: moduleRoot,
			modulePath: modulePath,
			packageDir: packageDir,
			packageName: packageName,
			runtimeDir: runtimeDir,
			entrypoint: entrypoint,
			build: GoBuildPolicy.NoBuild
		});
	}

	static function resolveProjectDirectory(moduleRoot:String, configuredPath:String, field:String):GoProjectRelativePath {
		validateRelativePath(configuredPath, field);
		final relative = GoProjectRelativePath.validated(configuredPath);
		var current = moduleRoot;
		final segments:Array<String> = configuredPath == "." ? [] : configuredPath.split("/");
		for (segment in segments) {
			current = Path.join([current, segment]);
			if (isSymbolicLink(current)) {
				throw new GoProjectModeError(InvalidPackageDirectory, field + " contains a symbolic link");
			}
			try {
				if (FileSystem.exists(current) && !FileSystem.isDirectory(current)) {
					throw new GoProjectModeError(InvalidPackageDirectory, field + " is not a directory");
				}
			} catch (error:GoProjectModeError) {
				throw error;
			} catch (_:haxe.Exception) {
				throw new GoProjectModeError(InvalidPackageDirectory, field + " could not be inspected");
			}
		}
		return relative;
	}

	static function resolveModuleRoot(manifestPath:String, configuredRoot:String):String {
		validateRelativePath(configuredRoot, "moduleRoot");
		final candidate = Path.join([Path.directory(manifestPath), configuredRoot]);
		if (isSymbolicLink(candidate)) {
			throw new GoProjectModeError(InvalidModuleFile, "moduleRoot is a symbolic link");
		}
		try {
			if (!FileSystem.exists(candidate) || !FileSystem.isDirectory(candidate)) {
				throw new GoProjectModeError(InvalidModuleFile, "moduleRoot is not a directory");
			}
			return normalizePath(FileSystem.fullPath(candidate));
		} catch (error:GoProjectModeError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(InvalidModuleFile, "moduleRoot could not be resolved");
		}
	}

	static function validateLegacyDefines(moduleRoot:String, packageDir:GoProjectRelativePath):Void {
		final configuredOutput = Context.definedValue(GoCompilerDefine.DefineGoOutput);
		if (configuredOutput == null || StringTools.trim(configuredOutput) == "") {
			throw new GoProjectModeError(ConfigurationConflict, "go_output is required");
		}
		final canonicalOutput = try {
			canonicalPathAllowMissing(StringTools.trim(configuredOutput));
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(ConfigurationConflict, "go_output could not be resolved");
		}
		final packageRoot = canonicalPathAllowMissing(Path.join([moduleRoot, packageDir.value()]));
		if (!samePath(canonicalOutput, packageRoot)) {
			throw new GoProjectModeError(ConfigurationConflict, "go_output does not match moduleRoot/packageDir");
		}
		if (Context.defined(GoCompilerDefine.DefineGoCommand) || Context.defined(GoCompilerDefine.DefineGoBuildOutput)) {
			throw new GoProjectModeError(ConfigurationConflict, "legacy Go build defines conflict with build.kind none");
		}
	}

	static function canonicalPathAllowMissing(path:String):String {
		var current = normalizePath(FileSystem.absolutePath(path));
		final missing:Array<String> = [];
		while (!FileSystem.exists(current)) {
			final parent = normalizePath(Path.directory(current));
			if (parent == current || parent == "") {
				throw new haxe.Exception("path has no existing ancestor");
			}
			missing.unshift(Path.withoutDirectory(current));
			current = parent;
		}
		var canonical = normalizePath(FileSystem.fullPath(current));
		for (segment in missing) {
			canonical = normalizePath(Path.join([canonical, segment]));
		}
		return canonical;
	}

	static function readModulePath(moduleRoot:String):String {
		final goModPath = Path.join([moduleRoot, "go.mod"]);
		canonicalRegularFile(goModPath, InvalidModuleFile, "go.mod is missing or is not a regular file");
		final contents = try {
			File.getContent(goModPath);
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(InvalidModuleFile, "go.mod could not be read");
		}
		var modulePath:Null<String> = null;
		final moduleLine = ~/^module[ \t]+([^ \t]+)(?:[ \t]+\/\/.*)?$/;
		for (line in contents.split("\n")) {
			final trimmed = StringTools.trim(line);
			if (moduleLine.match(trimmed)) {
				if (modulePath != null) {
					throw new GoProjectModeError(InvalidModuleFile, "go.mod contains more than one module directive");
				}
				modulePath = moduleLine.matched(1);
			}
		}
		if (modulePath == null || modulePath == "") {
			throw new GoProjectModeError(InvalidModuleFile, "go.mod has no valid module directive");
		}
		return modulePath;
	}

	static function canonicalRegularFile(path:String, kind:GoProjectModeErrorKind, explanation:String):String {
		if (isSymbolicLink(path)) {
			throw new GoProjectModeError(kind, explanation);
		}
		try {
			if (!FileSystem.exists(path) || FileSystem.isDirectory(path)) {
				throw new GoProjectModeError(kind, explanation);
			}
			return normalizePath(FileSystem.fullPath(path));
		} catch (error:GoProjectModeError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(kind, explanation);
		}
	}

	static function validateObject(value:{}, allowed:Array<String>, required:Array<String>, explanation:String):Void {
		if (value == null || !Reflect.isObject(value) || Std.isOfType(value, Array)) {
			throw new GoProjectModeError(InvalidManifest, explanation);
		}
		for (field in Reflect.fields(value)) {
			if (!allowed.contains(field)) {
				throw new GoProjectModeError(InvalidManifest, "the project manifest contains an unknown field");
			}
		}
		for (field in required) {
			if (!Reflect.hasField(value, field)) {
				throw new GoProjectModeError(InvalidManifest, explanation);
			}
		}
	}

	static function validateString(value:{}, field:String):Void {
		if (!Reflect.hasField(value, field) || !Std.isOfType(Reflect.field(value, field), String)) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest contains a non-string field");
		}
	}

	static function validateInt(value:{}, field:String):Void {
		if (!Reflect.hasField(value, field) || !Std.isOfType(Reflect.field(value, field), Int)) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest contains a non-integer field");
		}
	}

	static function validateRelativePath(value:String, field:String):Void {
		if (value == ".") {
			return;
		}
		if (value == null || value == "" || Path.isAbsolute(value) || value.indexOf("\\") != -1 || value.indexOf(":") != -1) {
			throw new GoProjectModeError(InvalidManifest, field + " is not a safe relative path");
		}
		for (segment in value.split("/")) {
			if (segment == "" || segment == "." || segment == "..") {
				throw new GoProjectModeError(InvalidManifest, field + " is not a safe relative path");
			}
		}
	}

	static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}

	static function normalizePath(path:String):String {
		var normalized = StringTools.replace(Path.normalize(path), "\\", "/");
		while (normalized.length > 1 && StringTools.endsWith(normalized, "/")) {
			normalized = normalized.substr(0, normalized.length - 1);
		}
		return normalized;
	}

	static function samePath(left:String, right:String):Bool {
		return Sys.systemName() == "Windows" ? left.toLowerCase() == right.toLowerCase() : left == right;
	}
}
#else
class GoProjectModeResolver {}
#end
