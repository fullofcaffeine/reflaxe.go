import reflaxe.go.compiler.GoPostBuildRunner;
import reflaxe.go.compiler.GoPostBuildRunner.GoPostBuildEnvironmentEntry;
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
		Sys.putEnv("HAXE_GO_AMBIENT_SECRET", "ambient-secret-value");
		Sys.putEnv("GOFLAGS", "-ambient-flags-must-not-run");
		var originalEnvironment = environmentSnapshot();
		var path = Sys.getEnv("PATH");
		assertThat(path != null && path != "", "missing PATH test baseline");
		var governedEnvironment:Array<GoPostBuildEnvironmentEntry> = [
			{name: "PATH", value: path},
			{name: "CGO_ENABLED", value: "0"},
			{name: "GOENV", value: "off"},
			{name: "GOTOOLCHAIN", value: "local"},
			{name: "GOWORK", value: "off"},
		];
		var invokedFrom:Null<String> = null;

		var nonzero = GoPostBuildRunner.runGoverned(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], governedEnvironment, (_, _) -> {
			invokedFrom = Sys.getCwd();
			assertThat(Sys.getEnv("HAXE_GO_AMBIENT_SECRET") == null, "ambient secret reached command");
			assertThat(Sys.getEnv("GOFLAGS") == null, "ambient GOFLAGS reached command");
			assertThat(Sys.getEnv("CGO_ENABLED") == "0", "governed CGO value did not reach command");
			assertThat(Sys.getEnv("GOENV") == "off", "compiler GOENV did not reach command");
			return 7;
		});
		assertThat(invokedFrom != null && FileSystem.fullPath(invokedFrom) == FileSystem.fullPath(resolvedWorkingDirectory),
			"command did not run from generated output");
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after nonzero exit");
		assertThat(environmentSnapshot() == originalEnvironment, "environment was not restored after nonzero exit");
		switch (nonzero) {
			case BuildFailed(failures):
				assertThat(failures.length == 1, "expected one exit failure");
				switch (failures[0]) {
					case CommandExited(7):
					case _: throw "expected structured exit-7 failure";
				}
			case _:
				throw "expected structured exit-7 failure";
		}
		var nonzeroMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], nonzero);
		assertThat(nonzeroMessage == "Post-generation Go build failed: `go build .` exited with status 7.", "unexpected exit diagnostic");
		assertThat(nonzeroMessage.indexOf("/fixture/") == -1, "command path leaked into exit diagnostic");
		var absoluteOutputMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "-o", "/fixture/output/app", "."], nonzero);
		assertThat(absoluteOutputMessage.indexOf("/fixture/") == -1, "absolute output path leaked into exit diagnostic");
		assertThat(absoluteOutputMessage.indexOf("<absolute-path>") != -1, "absolute output path omitted redaction marker");

		var launchFailure = GoPostBuildRunner.runGoverned(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], governedEnvironment, (_, _) -> {
			throw "permission denied at " + resolvedWorkingDirectory;
		});
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after launch exception");
		assertThat(environmentSnapshot() == originalEnvironment, "environment was not restored after launch exception");
		switch (launchFailure) {
			case BuildFailed(failures):
				assertThat(failures.length == 1, "expected one launch failure");
				switch (failures[0]) {
					case CommandLaunchFailed(cause):
						assertThat(cause.indexOf(resolvedWorkingDirectory) == -1, "working directory leaked into structured cause");
						assertThat(cause.indexOf("<generated-output>") != -1, "structured cause omitted redaction marker");
					case _: throw "expected structured launch failure";
				}
			case _:
				throw "expected structured launch failure";
		}
		var launchMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], launchFailure);
		assertThat(launchMessage.indexOf("could not launch `go build .`") != -1, "launch diagnostic omitted phase");
		assertThat(launchMessage.indexOf(resolvedWorkingDirectory) == -1, "working directory leaked into launch diagnostic");
		assertThat(launchMessage.indexOf("/fixture/") == -1, "command path leaked into launch diagnostic");

		var failApplyOnce = true;
		var applyFailure = GoPostBuildRunner.runGoverned(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], governedEnvironment, (_, _) -> 0,
			(name, value) -> {
				Sys.putEnv(name, value);
				if (failApplyOnce && name == "GOENV" && value == "off") {
					failApplyOnce = false;
					throw new haxe.Exception("environment apply failure with ambient-secret-value");
				}
			});
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after environment apply failure");
		assertThat(environmentSnapshot() == originalEnvironment, "environment was not restored after environment apply failure");
		switch (applyFailure) {
			case BuildFailed(failures):
				assertThat(failures.length == 1, "expected one environment apply failure");
				switch (failures[0]) {
					case EnvironmentApplyFailed("GOENV"):
					case _: throw "expected structured environment apply failure";
				}
			case _:
				throw "expected structured environment apply failure";
		}
		var applyMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], applyFailure);
		assertThat(applyMessage.indexOf("GOENV") != -1, "environment apply diagnostic omitted variable");
		assertThat(applyMessage.indexOf("ambient-secret-value") == -1, "environment apply diagnostic leaked a value");

		var commandRan = false;
		var failRestoreOnce = true;
		var restoreFailure = GoPostBuildRunner.runGoverned(resolvedWorkingDirectory, "/fixture/bin/go", ["build", "."], governedEnvironment, (_, _) -> {
			commandRan = true;
			return 9;
		}, (name, value) -> {
			Sys.putEnv(name, value);
			if (failRestoreOnce && commandRan && name == "HAXE_GO_AMBIENT_SECRET" && value == "ambient-secret-value") {
				failRestoreOnce = false;
				throw new haxe.Exception("environment restore failure with ambient-secret-value");
			}
		});
		assertThat(Sys.getCwd() == originalCwd, "working directory was not restored after environment restore failure");
		assertThat(environmentSnapshot() == originalEnvironment, "environment was not restored after injected restore failure");
		switch (restoreFailure) {
			case BuildFailed(failures):
				assertThat(failures.length == 2, "expected command and restoration failures");
				switch (failures[0]) {
					case CommandExited(9):
					case _: throw "expected command exit failure before restoration failure";
				}
				switch (failures[1]) {
					case EnvironmentRestoreFailed("HAXE_GO_AMBIENT_SECRET"):
					case _: throw "expected structured environment restore failure";
				}
			case _:
				throw "expected command and environment restore failures";
		}
		var restoreMessage = GoPostBuildRunner.failureMessage("/fixture/bin/go", ["build", "."], restoreFailure);
		assertThat(restoreMessage.indexOf("status 9") != -1, "restore diagnostic lost the command failure");
		assertThat(restoreMessage.indexOf("HAXE_GO_AMBIENT_SECRET") != -1, "restore diagnostic omitted variable");
		assertThat(restoreMessage.indexOf("ambient-secret-value") == -1, "restore diagnostic leaked a value");

		Sys.println("post-generation build runner: OK");
	}

	static function environmentSnapshot():String {
		var environment = Sys.environment();
		var names = [for (name in environment.keys()) name];
		names.sort(Reflect.compare);
		return names.map(name -> name + "=" + environment.get(name)).join("\n");
	}
}
