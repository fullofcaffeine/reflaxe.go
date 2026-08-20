package reflaxe.go.compiler;

#if (macro || eval)
import eval.luv.File.FileSync;
import eval.luv.Result;
import haxe.Json;
import haxe.io.Path;
import reflaxe.go.ast.GoPackageName;
import reflaxe.go.compiler.GoProjectMode.GoEntrypointPolicy;
import reflaxe.go.compiler.GoProjectMode.GoProjectRelativePath;
import reflaxe.go.compiler.GoProjectModeError;
import reflaxe.go.compiler.GoProjectModeError.GoProjectModeErrorKind;
import sys.FileSystem;
import sys.io.File;

private typedef LegacyManagedMetadata = {
	final filesGenerated:Array<String>;
}

/**
	Validates caller Go package clauses before existing-module output begins.

	The scanner reads only the required leading `package` clause. It skips Go
	comments without trying to model declarations or infer symbol ownership.
	Compiler-owned `main` is admitted only for an empty directory or for a rerun
	whose Go sources all appear in the current generated-file inventory.
**/
class GoPackageDirectoryInspector {
	public static function validate(moduleRoot:String, packageDir:GoProjectRelativePath, packageName:GoPackageName, entrypoint:GoEntrypointPolicy):Void {
		final directory = Path.join([moduleRoot, packageDir.value()]);
		if (!FileSystem.exists(directory)) {
			return;
		}

		final goFiles = try {
			[
				for (name in FileSystem.readDirectory(directory))
					if (isGoSourceName(name) && !FileSystem.isDirectory(Path.join([directory, name]))) name
			];
		} catch (_:haxe.Exception) {
			throw new GoProjectModeError(GoProjectModeErrorKind.InvalidPackageDirectory, "packageDir could not be inspected");
		}
		goFiles.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		final managedPaths = provisionalManagedPaths(moduleRoot);
		final callerGoFiles = [
			for (name in goFiles)
				if (!isManagedGeneratedSource(packageDir, name, managedPaths)) name
		];

		for (name in goFiles) {
			final path = Path.join([directory, name]);
			if (isSymbolicLink(path)) {
				throw new GoProjectModeError(GoProjectModeErrorKind.InvalidPackageDirectory, "packageDir contains a symbolic-link Go source");
			}
			final declared = try {
				readPackageName(File.getContent(path));
			} catch (_:haxe.Exception) {
				throw new GoProjectModeError(GoProjectModeErrorKind.PackageMismatch, "an existing Go source has no readable package clause");
			}
			final externalTestPackage = StringTools.endsWith(name, "_test.go") && declared == packageName.value() + "_test";
			if (declared != packageName.value() && !externalTestPackage) {
				throw new GoProjectModeError(GoProjectModeErrorKind.PackageMismatch, "an existing Go source declares another package");
			}
		}

		switch (entrypoint) {
			case CompilerMain if (callerGoFiles.length > 0):
				throw new GoProjectModeError(GoProjectModeErrorKind.EntrypointOwnership,
					"compiler-main requires an empty package directory until compiler ownership is established");
			case _:
		}
	}

	/**
		Reads only a provisional ownership hint for the preflight decision.

		The generated-output boundary validates the full metadata before any write or
		deletion. An absent or malformed hint grants no ownership here.
	**/
	static function provisionalManagedPaths(moduleRoot:String):Map<String, Bool> {
		final paths:Map<String, Bool> = [];
		final metadataPath = Path.join([moduleRoot, GoGeneratedOutputBoundary.MANAGED_FILE_METADATA]);
		try {
			if (!FileSystem.exists(metadataPath) || FileSystem.isDirectory(metadataPath) || isSymbolicLink(metadataPath)) {
				return paths;
			}
			final metadata:LegacyManagedMetadata = Json.parse(File.getContent(metadataPath));
			if (metadata == null || metadata.filesGenerated == null) {
				return paths;
			}
			for (path in metadata.filesGenerated) {
				if (path != null) {
					paths.set(path.toLowerCase(), true);
				}
			}
		} catch (_:haxe.Exception) {}
		return paths;
	}

	static function isManagedGeneratedSource(packageDir:GoProjectRelativePath, name:String, managedPaths:Map<String, Bool>):Bool {
		return StringTools.startsWith(name.toLowerCase(), "haxego_generated_")
			&& managedPaths.exists(packageDir.resolve(name).toLowerCase());
	}

	static function isGoSourceName(name:String):Bool {
		return name != null && name != "" && !StringTools.startsWith(name, ".") && !StringTools.startsWith(name, "_") && StringTools.endsWith(name, ".go");
	}

	static function readPackageName(contents:String):String {
		var index = 0;
		if (contents.length > 0 && contents.charCodeAt(0) == 0xFEFF) {
			index++;
		}
		index = skipTrivia(contents, index);
		if (contents.substr(index, 7) != "package" || isIdentifierCode(charCodeAt(contents, index + 7))) {
			throw new haxe.Exception("missing package clause");
		}
		index += 7;
		if (index >= contents.length || !isWhitespace(contents.charCodeAt(index))) {
			throw new haxe.Exception("invalid package clause");
		}
		while (index < contents.length && isWhitespace(contents.charCodeAt(index))) {
			index++;
		}
		final start = index;
		while (index < contents.length && isIdentifierCode(contents.charCodeAt(index))) {
			index++;
		}
		final name = contents.substring(start, index);
		if (!GoPackageName.isIdentifier(name)) {
			throw new haxe.Exception("invalid package name");
		}
		return name;
	}

	static function skipTrivia(contents:String, start:Int):Int {
		var index = start;
		while (index < contents.length) {
			if (isWhitespace(contents.charCodeAt(index))) {
				index++;
				continue;
			}
			if (contents.substr(index, 2) == "//") {
				final newline = contents.indexOf("\n", index + 2);
				index = newline < 0 ? contents.length : newline + 1;
				continue;
			}
			if (contents.substr(index, 2) == "/*") {
				final close = contents.indexOf("*/", index + 2);
				if (close < 0) {
					throw new haxe.Exception("unterminated comment");
				}
				index = close + 2;
				continue;
			}
			break;
		}
		return index;
	}

	static inline function charCodeAt(value:String, index:Int):Int {
		return index >= value.length ? -1 : value.charCodeAt(index);
	}

	static inline function isWhitespace(code:Int):Bool {
		return code == 32 || code == 9 || code == 10 || code == 13;
	}

	static inline function isIdentifierCode(code:Int):Bool {
		return code == 95 || (code >= 48 && code <= 57) || (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
	}

	static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}
}
#else
class GoPackageDirectoryInspector {}
#end
