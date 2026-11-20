import haxe.io.Bytes;
import haxe.io.Path;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathErrorKind;
import reflaxe.go.compiler.GoModuleFileGuard;
import sys.FileSystem;
import sys.io.File;

/**
	Contract for the caller-owned `go.mod`/`go.sum` guard.

	Why
	Existing-module mode must never change the caller's module files. The guard
	captures them before generation and rejects the run when their bytes differ
	afterwards. A check that missed an edit with the same length, or that
	confused an empty file with a missing one, would let a caller change slip by.

	What
	Proves that `verify` accepts unchanged files and an absent optional `go.sum`,
	and rejects a same-length edit, a length change, a deletion, and a creation of
	either file, including the empty-file versus absent distinction.

	How
	The test writes the files into the directory given as the first argument,
	captures a guard, applies one change, and expects `verify` to throw the
	path-free protected-file error.
**/
class GoModuleFileGuardContract {
	static final GO_MOD = "module example.com/guard\n\ngo 1.24.0\n";
	static final GO_SUM = "example.com/dependency v1.0.0 h1:caller-owned-bytes\n";

	static var root = "";

	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function path(name:String):String {
		return Path.join([root, name]);
	}

	static function reset(goSum:Null<String>):Void {
		File.saveContent(path("go.mod"), GO_MOD);
		if (goSum == null) {
			if (FileSystem.exists(path("go.sum"))) {
				FileSystem.deleteFile(path("go.sum"));
			}
		} else {
			File.saveContent(path("go.sum"), goSum);
		}
	}

	static function expectAccepted(label:String, goSum:Null<String>, change:() -> Void):Void {
		reset(goSum);
		final guard = new GoModuleFileGuard(root);
		change();
		try {
			guard.verify();
		} catch (error:GoOutputPathError) {
			throw '$label was rejected: ${error.message}';
		}
	}

	static function expectRejected(label:String, goSum:Null<String>, change:() -> Void):Void {
		reset(goSum);
		final guard = new GoModuleFileGuard(root);
		change();
		try {
			guard.verify();
		} catch (error:GoOutputPathError) {
			require(error.kind == GoOutputPathErrorKind.ProtectedCallerFile, '$label failed with the wrong kind: ${error.message}');
			return;
		}
		throw '$label was accepted';
	}

	static function replaceByte(name:String, index:Int, code:Int):Void {
		final bytes = File.getBytes(path(name));
		require(bytes.get(index) != code, 'the $name edit would not change a byte');
		bytes.set(index, code);
		File.saveBytes(path(name), bytes);
	}

	static function main():Void {
		root = Sys.args()[0];
		require(root != null && FileSystem.isDirectory(root), "pass an existing scratch directory as the first argument");

		expectAccepted("unchanged module files", GO_SUM, () -> {});
		expectAccepted("rewritten identical module files", GO_SUM, () -> reset(GO_SUM));
		expectAccepted("an absent go.sum that stays absent", null, () -> {});
		expectAccepted("an empty go.sum that stays empty", "", () -> {});

		expectRejected("a same-length go.sum edit", GO_SUM, () -> replaceByte("go.sum", GO_SUM.length - 2, "X".code));
		expectRejected("a same-length go.mod edit", GO_SUM, () -> replaceByte("go.mod", 0, "M".code));
		expectRejected("a longer go.sum", GO_SUM, () -> File.saveContent(path("go.sum"), GO_SUM + "\n"));
		expectRejected("a shorter go.sum", GO_SUM, () -> File.saveContent(path("go.sum"), GO_SUM.substr(1)));
		expectRejected("a deleted go.sum", GO_SUM, () -> FileSystem.deleteFile(path("go.sum")));
		expectRejected("a created go.sum", null, () -> File.saveContent(path("go.sum"), GO_SUM));
		expectRejected("a deleted empty go.sum", "", () -> FileSystem.deleteFile(path("go.sum")));
		expectRejected("a created empty go.sum", null, () -> File.saveBytes(path("go.sum"), Bytes.alloc(0)));

		Sys.println("ok");
	}
}
