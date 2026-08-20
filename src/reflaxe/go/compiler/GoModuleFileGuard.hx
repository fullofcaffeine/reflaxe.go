package reflaxe.go.compiler;

#if (macro || eval)
import eval.luv.File.FileSync;
import eval.luv.Result;
import haxe.crypto.Sha256;
import haxe.io.Bytes;
import haxe.io.Path;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import sys.FileSystem;
import sys.io.File;

/** Captured byte state for one caller-owned module file. */
enum GoProtectedFileState {
	Absent;
	Present(bytes:Bytes, digest:String);
}

/**
	Protects caller-owned module files with exact pre/post byte comparisons.

	The guard records absence separately from an empty file. It rejects links and
	non-files so a compiler run cannot silently exchange the protected object.
**/
class GoModuleFileGuard {
	final moduleRoot:String;
	final goMod:GoProtectedFileState;
	final goSum:GoProtectedFileState;

	public function new(moduleRoot:String) {
		this.moduleRoot = moduleRoot;
		goMod = capture("go.mod", true);
		goSum = capture("go.sum", false);
	}

	public function verify():Void {
		verifyFile("go.mod", goMod, true);
		verifyFile("go.sum", goSum, false);
	}

	function capture(name:String, required:Bool):GoProtectedFileState {
		final path = Path.join([moduleRoot, name]);
		if (isSymbolicLink(path)) {
			throw mutation("a protected module file is a symbolic link");
		}
		try {
			if (!FileSystem.exists(path)) {
				if (required) {
					throw mutation("the required module file is missing");
				}
				return Absent;
			}
			if (FileSystem.isDirectory(path)) {
				throw mutation("a protected module path is not a regular file");
			}
			final bytes = File.getBytes(path);
			return Present(bytes, Sha256.make(bytes).toHex());
		} catch (error:GoOutputPathError) {
			throw error;
		} catch (_:haxe.Exception) {
			throw mutation("a protected module file could not be read");
		}
	}

	function verifyFile(name:String, expected:GoProtectedFileState, required:Bool):Void {
		final current = capture(name, required);
		switch ([expected, current]) {
			case [Absent, Absent]:
			case [Present(expectedBytes, expectedDigest), Present(currentBytes, currentDigest)]
				if (expectedDigest == currentDigest && bytesEqual(expectedBytes, currentBytes)):
			case _:
				throw mutation("a caller-owned module file changed during generation");
		}
	}

	static function bytesEqual(left:Bytes, right:Bytes):Bool {
		if (left.length != right.length) {
			return false;
		}
		for (index in 0...left.length) {
			if (left.get(index) != right.get(index)) {
				return false;
			}
		}
		return true;
	}

	static function mutation(explanation:String):GoOutputPathError {
		return new GoOutputPathError(ProtectedCallerFile, explanation);
	}

	static function isSymbolicLink(path:String):Bool {
		return switch (FileSync.readLink(path)) {
			case Ok(_): true;
			case Error(_): false;
		};
	}
}
#else
class GoModuleFileGuard {}
#end
