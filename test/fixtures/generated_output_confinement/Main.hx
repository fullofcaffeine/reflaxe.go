import reflaxe.go.compiler.GoGeneratedOutputBoundary;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import sys.FileSystem;

class Main {
	static function assertThat(condition:Bool, message:String):Void {
		if (!condition) {
			throw message;
		}
	}

	static function expectRejected(boundary:GoGeneratedOutputBoundary, path:String, root:String, outside:String):Void {
		try {
			boundary.saveContent(path, "package escaped\n");
		} catch (error:GoOutputPathError) {
			assertThat(error.message.indexOf(root) == -1, "diagnostic leaked the output root");
			assertThat(error.message.indexOf(outside) == -1, "diagnostic leaked the external path");
			assertThat(error.message.indexOf(path) == -1, "diagnostic echoed the rejected path");
			return;
		}
		throw 'expected generated-output path rejection';
	}

	static function main():Void {
		var root = Sys.getEnv("GO_OUTPUT_CONFINEMENT_ROOT");
		var outside = Sys.getEnv("GO_OUTPUT_CONFINEMENT_OUTSIDE");
		assertThat(root != null && root != "", "missing generated-output root");
		assertThat(outside != null && outside != "", "missing external fixture root");
		var outputRoot:String = root;
		var outsideRoot:String = outside;
		var boundary = new GoGeneratedOutputBoundary(outputRoot);

		boundary.saveContent("nested/ok.go", "package nested\n");
		assertThat(FileSystem.exists(outputRoot + "/nested/ok.go"), "safe nested write was rejected");

		for (unsafePath in [
			"../escape.go",
			"nested/../../escape.go",
			"/private/escape.go",
			"//server/share/escape.go",
			"C:/escape.go",
			"C:\\escape.go",
			"\\\\server\\share\\escape.go",
			"nested\\escape.go",
			"nested//escape.go",
			"./escape.go",
			"nested/../escape.go",
			"NUL",
			"file.go:stream"
		]) {
			expectRejected(boundary, unsafePath, outputRoot, outsideRoot);
		}

		expectRejected(boundary, "escape-dir/escaped.go", outputRoot, outsideRoot);
		expectRejected(boundary, "escape-file.go", outputRoot, outsideRoot);
		expectRejected(boundary, "broken-file.go", outputRoot, outsideRoot);

		expectRejected(boundary, "inside-link/ok.go", outputRoot, outsideRoot);
		assertThat(!FileSystem.exists(outputRoot + "/inside/ok.go"), "contained output symlink was followed");

		Sys.println("generated-output confinement: OK");
	}
}
