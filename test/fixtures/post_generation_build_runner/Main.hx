import reflaxe.go.compiler.GoPostBuildRunner;
import reflaxe.go.compiler.GoPostBuildRunner.GoPostBuildFailure;
import reflaxe.go.compiler.GoPostBuildRunner.GoPostBuildResult;
import sys.FileSystem;

class Main {
	static function assertThat(condition:Bool, message:String):Void {
		if (!condition) {
			throw message;
		}
	}

	static function main():Void {
		var workingDirectory = Sys.getEnv("GO_POST_BUILD_TEST_WORKDIR");
		assertThat(workingDirectory != null && workingDirectory != "", "missing test working directory");
		var resolvedWorkingDirectory:String = workingDirectory;
		var originalCwd = Sys.getCwd();
		var invokedFrom:Null<String> = null;

		var nonzero = GoPostBuildRunner.run(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], (_, _) -> {
			invokedFrom = Sys.getCwd();
			return 7;
		});
		assertThat(invokedFrom != null && FileSystem.fullPath(invokedFrom) == FileSystem.fullPath(resolvedWorkingDirectory),
			"command did not run from generated output");
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after nonzero exit");
		switch (nonzero) {
			case BuildFailed(CommandExited(7)):
			case _:
				throw "expected structured exit-7 failure";
		}
		var nonzeroMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], nonzero);
		assertThat(nonzeroMessage == "Post-generation Go build failed: `go build .` exited with status 7.", "unexpected exit diagnostic");
		assertThat(nonzeroMessage.indexOf("/fixture/") == -1, "command path leaked into exit diagnostic");
		var absoluteOutputMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "-o", "/fixture/output/app", "."], nonzero);
		assertThat(absoluteOutputMessage.indexOf("/fixture/") == -1, "absolute output path leaked into exit diagnostic");
		assertThat(absoluteOutputMessage.indexOf("<absolute-path>") != -1, "absolute output path omitted redaction marker");

		var launchFailure = GoPostBuildRunner.run(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], (_, _) -> {
			throw "permission denied at " + resolvedWorkingDirectory;
		});
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after launch exception");
		switch (launchFailure) {
			case BuildFailed(CommandLaunchFailed(cause)):
				assertThat(cause.indexOf(resolvedWorkingDirectory) == -1, "working directory leaked into structured cause");
				assertThat(cause.indexOf("<generated-output>") != -1, "structured cause omitted redaction marker");
			case _:
				throw "expected structured launch failure";
		}
		var launchMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], launchFailure);
		assertThat(launchMessage.indexOf("could not launch `go build .`") != -1, "launch diagnostic omitted phase");
		assertThat(launchMessage.indexOf(resolvedWorkingDirectory) == -1, "working directory leaked into launch diagnostic");
		assertThat(launchMessage.indexOf("/fixture/") == -1, "command path leaked into launch diagnostic");

		Sys.println("post-generation build runner: OK");
	}
}
