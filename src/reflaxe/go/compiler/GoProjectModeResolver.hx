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
	final ?packageTarget:String;
	final ?output:String;
	final ?tags:Array<String>;
	final ?ldflags:Array<String>;
	final ?trimpath:Bool;
	final ?race:Bool;
	final ?arguments:Array<String>;
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
	final InvalidBuildTarget = "GO-BUILD-TARGET";
	final InvalidBuildOutput = "GO-BUILD-OUTPUT";
	final InvalidBuildTag = "GO-BUILD-TAG";
	final InvalidLinkerArgument = "GO-BUILD-LDFLAG";
	final InvalidBuildArgument = "GO-BUILD-ARGUMENT";
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
		validateObject(raw.build, [
			"kind",
			"packageTarget",
			"output",
			"tags",
			"ldflags",
			"trimpath",
			"race",
			"arguments"
		], ["kind"], "build has an invalid shape");
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
		final build = resolveBuildPolicy(moduleRoot, raw.build);

		GoPackageDirectoryInspector.validate(moduleRoot, packageDir, packageName, entrypoint);
		validateLegacyDefines(moduleRoot, packageDir, build);
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
			build: build
		});
	}

	static function resolveBuildPolicy(moduleRoot:String, raw:RawBuild):GoBuildPolicy {
		if (raw.kind == "none") {
			validateObject(raw, ["kind"], ["kind"], "build kind none has an invalid shape");
			return GoBuildPolicy.NoBuild;
		}
		if (raw.kind != "go-build") {
			throw new GoProjectModeError(UnsupportedProjectShape, "build has an unsupported kind");
		}

		final fields = [
			"kind",
			"packageTarget",
			"output",
			"tags",
			"ldflags",
			"trimpath",
			"race",
			"arguments"
		];
		validateObject(raw, fields, fields, "go-build has an invalid shape");
		for (field in ["packageTarget", "output"]) {
			validateString(raw, field);
		}
		for (field in ["tags", "ldflags", "arguments"]) {
			validateStringArray(raw, field);
		}
		for (field in ["trimpath", "race"]) {
			validateBool(raw, field);
		}

		final packageTarget = validateBuildTarget(moduleRoot, raw.packageTarget);
		final output = validateBuildOutput(moduleRoot, raw.output);
		final tags = canonicalBuildTags(raw.tags);
		validateLinkerArguments(raw.ldflags);
		final arguments = canonicalBuildArguments(raw.arguments);
		return GoBuildPolicy.GoBuild(new GoBuildRequest({
			packageTarget: packageTarget,
			output: output,
			tags: tags,
			ldflags: raw.ldflags,
			trimpath: raw.trimpath,
			race: raw.race,
			arguments: arguments
		}));
	}

	static function validateBuildTarget(moduleRoot:String, value:String):String {
		if (value == ".") {
			return value;
		}
		if (!StringTools.startsWith(value, "./")) {
			throw new GoProjectModeError(InvalidBuildTarget, "the package target must be dot or a module-relative package");
		}
		final relative = value.substr(2);
		try {
			resolveProjectDirectory(moduleRoot, relative, "build.packageTarget");
		} catch (_:GoProjectModeError) {
			throw new GoProjectModeError(InvalidBuildTarget, "the package target is not a safe module-relative package");
		}
		return value;
	}

	static function validateBuildOutput(moduleRoot:String, value:String):String {
		try {
			validateRelativePath(value, "build.output");
			if (value == ".") {
				throw new GoProjectModeError(InvalidBuildOutput, "the output must name a module-relative file");
			}
			var current = moduleRoot;
			final segments = value.split("/");
			for (index in 0...segments.length) {
				current = Path.join([current, segments[index]]);
				if (isSymbolicLink(current)) {
					throw new GoProjectModeError(InvalidBuildOutput, "the output path contains a symbolic link");
				}
				if (index < segments.length - 1 && FileSystem.exists(current) && !FileSystem.isDirectory(current)) {
					throw new GoProjectModeError(InvalidBuildOutput, "the output parent is not a directory");
				}
				if (index == segments.length - 1 && FileSystem.exists(current) && FileSystem.isDirectory(current)) {
					throw new GoProjectModeError(InvalidBuildOutput, "the output must name a file");
				}
			}
		} catch (error:GoProjectModeError) {
			if (error.kind == InvalidBuildOutput) {
				throw error;
			}
			throw new GoProjectModeError(InvalidBuildOutput, "the output is not a safe module-relative file");
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(InvalidBuildOutput, "the output could not be inspected");
		}
		return value;
	}

	static function canonicalBuildTags(values:Array<String>):Array<String> {
		final unique:Map<String, Bool> = [];
		for (value in values) {
			if (value == null || value == "" || !~/^[A-Za-z0-9_.]+$/.match(value)) {
				throw new GoProjectModeError(InvalidBuildTag, "a build tag is not a valid standalone tag");
			}
			unique.set(value, true);
		}
		final result = [for (value in unique.keys()) value];
		result.sort(Reflect.compare);
		return result;
	}

	static function validateLinkerArguments(values:Array<String>):Void {
		for (value in values) {
			var hasControl = false;
			if (value != null) {
				for (index in 0...value.length) {
					final code = value.charCodeAt(index);
					if (code != null && (code < 32 || code == 127)) {
						hasControl = true;
					}
				}
			}
			if (value == null || value == "" || hasControl || (value.indexOf("'") != -1 && value.indexOf('"') != -1)) {
				throw new GoProjectModeError(InvalidLinkerArgument, "a linker argument cannot be represented by the Go linker argument grammar");
			}
		}
	}

	static function canonicalBuildArguments(values:Array<String>):Array<String> {
		final byFamily:Map<String, String> = [];
		for (value in values) {
			final family = approvedBuildArgumentFamily(value);
			if (family == null || byFamily.exists(family)) {
				throw new GoProjectModeError(InvalidBuildArgument, "an additional Go build argument is unapproved or repeated");
			}
			byFamily.set(family, value);
		}
		if (!byFamily.exists("mod")) {
			byFamily.set("mod", "-mod=readonly");
		}
		final result = [for (value in byFamily) value];
		result.sort(Reflect.compare);
		return result;
	}

	static function approvedBuildArgumentFamily(value:String):Null<String> {
		return switch (value) {
			case "-a": "a";
			case "-v": "v";
			case "-x": "x";
			case _ if (~/^-buildvcs=(auto|false|true)$/.match(value)): "buildvcs";
			case _ if (~/^-buildmode=(archive|c-archive|c-shared|default|exe|pie|plugin|shared)$/.match(value)): "buildmode";
			case _ if (~/^-mod=(readonly|vendor)$/.match(value)): "mod";
			case _ if (~/^-p=[1-9][0-9]*$/.match(value)): "p";
			case _: null;
		};
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

	static function validateLegacyDefines(moduleRoot:String, packageDir:GoProjectRelativePath, build:GoBuildPolicy):Void {
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
			throw new GoProjectModeError(ConfigurationConflict, "legacy Go build defines conflict with the typed build policy");
		}
		switch (build) {
			case NoBuild:
			case GoBuild(_) if (Context.defined(GoCompilerDefine.DefineGoNoBuild)
				|| Context.defined(GoCompilerDefine.DefineGoCodegenOnly)):
				throw new GoProjectModeError(ConfigurationConflict, "legacy no-build defines conflict with build.kind go-build");
			case GoBuild(_):
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

	static function validateBool(value:{}, field:String):Void {
		if (!Reflect.hasField(value, field) || !Std.isOfType(Reflect.field(value, field), Bool)) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest contains a non-boolean field");
		}
	}

	static function validateStringArray(value:{}, field:String):Void {
		if (!Reflect.hasField(value, field) || !Std.isOfType(Reflect.field(value, field), Array)) {
			throw new GoProjectModeError(InvalidManifest, "the project manifest contains a non-array field");
		}
		final values:Array<{}> = Reflect.field(value, field);
		for (item in values) {
			if (!Std.isOfType(item, String)) {
				throw new GoProjectModeError(InvalidManifest, "the project manifest contains a non-string array item");
			}
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
