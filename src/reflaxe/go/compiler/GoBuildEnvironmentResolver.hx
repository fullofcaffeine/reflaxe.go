package reflaxe.go.compiler;

#if (macro || eval)
import haxe.io.Path;
import reflaxe.go.compiler.GoBuildEnvironment.GoBuildEnvironmentEntry;
import reflaxe.go.compiler.GoBuildEnvironment.GoBuildEnvironmentReportPolicy;
import reflaxe.go.compiler.GoBuildEnvironment.GoBuildEnvironmentSource;
import reflaxe.go.compiler.GoProjectModeError;
import reflaxe.go.compiler.GoProjectModeError.GoProjectModeErrorKind;
#end

/** The narrow manifest-boundary input for one requested environment value. */
typedef GoBuildEnvironmentInput = {
	final name:String;
	final source:String;
	final ?value:String;
}

#if (macro || eval)
private enum abstract GoBuildEnvironmentVariable(String) to String {
	final CgoEnabled = "CGO_ENABLED";
	final GoOs = "GOOS";
	final GoArch = "GOARCH";
	final GoAmd64 = "GOAMD64";
	final GoArm = "GOARM";
	final GoArm64 = "GOARM64";
	final Go386 = "GO386";
	final GoMips = "GOMIPS";
	final GoMips64 = "GOMIPS64";
	final GoPpc64 = "GOPPC64";
	final GoRiscv64 = "GORISCV64";
	final GoWasm = "GOWASM";
	final Cc = "CC";
	final Cxx = "CXX";
	final PkgConfig = "PKG_CONFIG";
	final GoCache = "GOCACHE";
	final GoModCache = "GOMODCACHE";
	final GoPath = "GOPATH";
	final TmpDir = "TMPDIR";
	final Temp = "TEMP";
	final Tmp = "TMP";
	final GoProxy = "GOPROXY";
	final GoPrivate = "GOPRIVATE";
	final GoNoProxy = "GONOPROXY";
	final GoNoSumDb = "GONOSUMDB";
	final GoSumDb = "GOSUMDB";
	final HttpProxy = "HTTP_PROXY";
	final HttpsProxy = "HTTPS_PROXY";
	final NoProxy = "NO_PROXY";
	final SslCertFile = "SSL_CERT_FILE";
	final SslCertDir = "SSL_CERT_DIR";
}

/**
	What: Resolves manifest environment entries into one minimal child environment.
	Why: Go build behavior must be explicit without copying ambient flags or secrets.
	How: Validate a closed variable set, resolve literal or same-name inherited values,
	add fixed Go controls and host launch requirements, then return an immutable value.
**/
class GoBuildEnvironmentResolver {
	static final GO_OPERATING_SYSTEMS = [
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"illumos",
		"ios",
		"js",
		"linux",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"wasip1",
		"windows"
	];
	static final GO_ARCHITECTURES = [
		"386",
		"amd64",
		"arm",
		"arm64",
		"loong64",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
		"ppc64",
		"ppc64le",
		"riscv64",
		"s390x",
		"wasm"
	];
	static final FIXED_NAMES = ["GOENV", "GOTOOLCHAIN", "GOWORK"];
	static final WINDOWS_BASELINE_NAMES = ["SystemRoot", "WINDIR", "ComSpec", "PATHEXT"];

	public static function resolve(inputs:Array<GoBuildEnvironmentInput>, race:Bool):GoBuildEnvironment {
		final entries:Array<GoBuildEnvironmentEntry> = [];
		final seen:Map<String, Bool> = [];
		for (input in inputs) {
			final key = input.name.toUpperCase();
			if (seen.exists(key)) {
				fail(input.name, "is repeated or aliases another entry");
			}
			seen.set(key, true);
			if (FIXED_NAMES.contains(key)) {
				fail(input.name, "is compiler-owned and cannot be configured");
			}
			final variable = parseVariable(input.name);
			if (variable == null) {
				fail(input.name, "is unknown or forbidden");
			}
			final resolvedVariable:GoBuildEnvironmentVariable = variable;
			final source = parseSource(input, resolvedVariable);
			final value = switch (source) {
				case Literal: input.value;
				case Inherit: inheritedValue(resolvedVariable);
				case Compiler, HostBaseline: null;
			};
			if (value == null || value == "" || hasNul(value)) {
				fail(resolvedVariable, "has an empty or invalid value");
			}
			validateValue(resolvedVariable, value);
			entries.push({
				name: resolvedVariable,
				value: value,
				source: source,
				reportPolicy: reportPolicy(resolvedVariable)
			});
		}

		validateCombinations(entries, race);
		if (!hasEntry(entries, GoCache)) {
			fail(GoCache, "must be configured explicitly for a governed build");
		}
		addHostBaseline(entries);
		entries.push(fixedEntry("GOENV", "off"));
		entries.push(fixedEntry("GOTOOLCHAIN", "local"));
		entries.push(fixedEntry("GOWORK", "off"));
		entries.sort((left, right) -> Reflect.compare(left.name.toUpperCase(), right.name.toUpperCase()));
		return new GoBuildEnvironment(entries);
	}

	static function parseVariable(name:String):Null<GoBuildEnvironmentVariable> {
		return switch (name) {
			case "CGO_ENABLED": CgoEnabled;
			case "GOOS": GoOs;
			case "GOARCH": GoArch;
			case "GOAMD64": GoAmd64;
			case "GOARM": GoArm;
			case "GOARM64": GoArm64;
			case "GO386": Go386;
			case "GOMIPS": GoMips;
			case "GOMIPS64": GoMips64;
			case "GOPPC64": GoPpc64;
			case "GORISCV64": GoRiscv64;
			case "GOWASM": GoWasm;
			case "CC": Cc;
			case "CXX": Cxx;
			case "PKG_CONFIG": PkgConfig;
			case "GOCACHE": GoCache;
			case "GOMODCACHE": GoModCache;
			case "GOPATH": GoPath;
			case "TMPDIR": TmpDir;
			case "TEMP": Temp;
			case "TMP": Tmp;
			case "GOPROXY": GoProxy;
			case "GOPRIVATE": GoPrivate;
			case "GONOPROXY": GoNoProxy;
			case "GONOSUMDB": GoNoSumDb;
			case "GOSUMDB": GoSumDb;
			case "HTTP_PROXY": HttpProxy;
			case "HTTPS_PROXY": HttpsProxy;
			case "NO_PROXY": NoProxy;
			case "SSL_CERT_FILE": SslCertFile;
			case "SSL_CERT_DIR": SslCertDir;
			case _: null;
		};
	}

	static function parseSource(input:GoBuildEnvironmentInput, variable:GoBuildEnvironmentVariable):GoBuildEnvironmentSource {
		return switch (input.source) {
			case "literal" if (input.value != null): Literal;
			case "inherit" if (input.value == null): Inherit;
			case "literal":
				fail(variable, "uses literal source without a value");
			case "inherit":
				fail(variable, "uses inherit source with a value");
			case _:
				fail(variable, "has an unsupported source");
		};
	}

	static function inheritedValue(variable:GoBuildEnvironmentVariable):String {
		final value = Sys.getEnv(variable);
		if (value == null || value == "") {
			fail(variable, "requests inheritance but is absent from the compiler environment");
		}
		return value;
	}

	static function validateValue(variable:GoBuildEnvironmentVariable, value:String):Void {
		switch (variable) {
			case CgoEnabled:
				if (value != "0" && value != "1")
					fail(variable, "must be 0 or 1");
			case GoOs:
				if (!GO_OPERATING_SYSTEMS.contains(value))
					fail(variable, "is not a supported Go operating system");
			case GoArch:
				if (!GO_ARCHITECTURES.contains(value))
					fail(variable, "is not a supported Go architecture");
			case GoAmd64:
				if (!~/^v[1-4]$/.match(value))
					fail(variable, "is not a supported amd64 tuning value");
			case GoArm:
				if (!~/^[567](,(softfloat|hardfloat))?$/.match(value))
					fail(variable, "is not a supported arm tuning value");
			case GoArm64:
				if (!~/^v(8\.[0-9]|9\.[0-5])(,(lse|crypto)){0,2}$/.match(value))
					fail(variable, "is not a supported arm64 tuning value");
			case Go386:
				if (value != "sse2" && value != "softfloat")
					fail(variable, "is not a supported 386 tuning value");
			case GoMips, GoMips64:
				if (value != "hardfloat" && value != "softfloat")
					fail(variable, "is not a supported MIPS tuning value");
			case GoPpc64:
				if (!["power8", "power9", "power10"].contains(value))
					fail(variable, "is not a supported ppc64 tuning value");
			case GoRiscv64:
				if (!["rva20u64", "rva22u64", "rva23u64"].contains(value))
					fail(variable, "is not a supported riscv64 tuning value");
			case GoWasm:
				validateWasmFeatures(variable, value);
			case Cc, Cxx, PkgConfig:
				validateExecutable(variable, value);
			case GoCache, GoModCache, GoPath, TmpDir, Temp, Tmp, SslCertFile, SslCertDir:
				if (!Path.isAbsolute(value))
					fail(variable, "must be one absolute path");
			case GoProxy, GoPrivate, GoNoProxy, GoNoSumDb, GoSumDb, HttpProxy, HttpsProxy, NoProxy:
		}
	}

	static function validateWasmFeatures(variable:GoBuildEnvironmentVariable, value:String):Void {
		final seen:Map<String, Bool> = [];
		for (feature in value.split(",")) {
			if ((feature != "satconv" && feature != "signext") || seen.exists(feature)) {
				fail(variable, "contains an unsupported or repeated WebAssembly feature");
			}
			seen.set(feature, true);
		}
	}

	static function validateExecutable(variable:GoBuildEnvironmentVariable, value:String):Void {
		if (~/[\x00-\x20\x7f'"=`$;&|<>]/.match(value)) {
			fail(variable, "must be one executable path or basename without arguments");
		}
		if (!Path.isAbsolute(value) && (value.indexOf("/") != -1 || value.indexOf("\\") != -1 || !~/^[A-Za-z0-9._+-]+$/.match(value))) {
			fail(variable, "must be one executable path or basename without arguments");
		}
	}

	static function validateCombinations(entries:Array<GoBuildEnvironmentEntry>, race:Bool):Void {
		final cgo = valueOf(entries, CgoEnabled);
		if (cgo == "0") {
			for (compiler in [Cc, Cxx, PkgConfig]) {
				if (hasEntry(entries, compiler)) {
					throw new GoProjectModeError(GoProjectModeErrorKind.InvalidBuildEnvironment,
						'variables CGO_ENABLED and ${compiler} cannot be combined when CGO is disabled');
				}
			}
			if (race) {
				fail(CgoEnabled, "cannot disable CGO while race mode is enabled");
			}
		}

		final selectedArch = valueOf(entries, GoArch);
		for (entry in entries) {
			final requiredArch = tuningArchitecture(entry.name);
			if (requiredArch != null
				&& selectedArch != requiredArch
				&& !(requiredArch == "mips" && selectedArch == "mipsle")
				&& !(requiredArch == "mips64" && selectedArch == "mips64le")
				&& !(requiredArch == "ppc64" && selectedArch == "ppc64le")) {
				fail(entry.name, selectedArch == null ? "requires an explicit matching GOARCH" : "does not match GOARCH");
			}
		}
	}

	static function tuningArchitecture(name:String):Null<String> {
		return switch (name) {
			case "GOAMD64": "amd64";
			case "GOARM": "arm";
			case "GOARM64": "arm64";
			case "GO386": "386";
			case "GOMIPS": "mips";
			case "GOMIPS64": "mips64";
			case "GOPPC64": "ppc64";
			case "GORISCV64": "riscv64";
			case "GOWASM": "wasm";
			case _: null;
		};
	}

	static function reportPolicy(variable:GoBuildEnvironmentVariable):GoBuildEnvironmentReportPolicy {
		return switch (variable) {
			case CgoEnabled, GoOs, GoArch, GoAmd64, GoArm, GoArm64, Go386, GoMips, GoMips64, GoPpc64, GoRiscv64, GoWasm:
				Plain;
			case Cc, Cxx, PkgConfig, GoCache, GoModCache, GoPath, TmpDir, Temp, Tmp:
				PathRedacted;
			case GoProxy, GoPrivate, GoNoProxy, GoNoSumDb, GoSumDb, HttpProxy, HttpsProxy, NoProxy, SslCertFile, SslCertDir:
				SecretRedacted;
		};
	}

	static function addHostBaseline(entries:Array<GoBuildEnvironmentEntry>):Void {
		final path = inheritedHostEntry("PATH");
		if (path == null || path.value == "") {
			fail("PATH", "is required to launch the governed Go command");
		}
		entries.push(hostEntry(path.name, path.value));
		if (Sys.systemName() == "Windows") {
			for (name in WINDOWS_BASELINE_NAMES) {
				final entry = inheritedHostEntry(name);
				if (entry != null && entry.value != "")
					entries.push(hostEntry(entry.name, entry.value));
			}
		}
	}

	static function inheritedHostEntry(name:String):Null<{name:String, value:String}> {
		final environment = Sys.environment();
		for (candidate in environment.keys()) {
			if (candidate == name || (Sys.systemName() == "Windows" && candidate.toUpperCase() == name.toUpperCase())) {
				return {name: candidate, value: environment.get(candidate)};
			}
		}
		return null;
	}

	static function fixedEntry(name:String, value:String):GoBuildEnvironmentEntry {
		return {
			name: name,
			value: value,
			source: Compiler,
			reportPolicy: Plain
		};
	}

	static function hostEntry(name:String, value:String):GoBuildEnvironmentEntry {
		return {
			name: name,
			value: value,
			source: HostBaseline,
			reportPolicy: Hidden
		};
	}

	static function hasEntry(entries:Array<GoBuildEnvironmentEntry>, variable:GoBuildEnvironmentVariable):Bool {
		return valueOf(entries, variable) != null;
	}

	static function valueOf(entries:Array<GoBuildEnvironmentEntry>, variable:GoBuildEnvironmentVariable):Null<String> {
		for (entry in entries)
			if (entry.name == variable)
				return entry.value;
		return null;
	}

	static function hasNul(value:String):Bool {
		return value.indexOf(String.fromCharCode(0)) != -1;
	}

	static function fail<T>(name:String, explanation:String):T {
		throw new GoProjectModeError(GoProjectModeErrorKind.InvalidBuildEnvironment, 'variable ${name} ${explanation}');
	}
}
#else
class GoBuildEnvironmentResolver {}
#end
