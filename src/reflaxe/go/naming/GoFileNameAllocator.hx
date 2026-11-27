package reflaxe.go.naming;

/**
	Gives each generated Go source file a unique name that every platform builds.

	Why
	The Go tool reads some file name suffixes as build constraints and silently
	leaves the file out of a build that does not match them. A name that ends in
	`_test` makes a test file, which `go build` and `go run` never compile. A
	name that ends in a known GOOS or GOARCH value, such as `_linux`, `_js` or
	`_arm64` (which also covers a `_windows_amd64` pair), builds only on that
	platform. Generated file names come from Haxe module paths, so a module
	`unit.Test` or `tools.Linux` used to produce such a name, and the program
	then failed to build with "undefined" errors.

	What
	`next` returns `<base>.go` for the first use of a base and `<base>_<n>.go`
	for later uses, as the compiler always did. When the base would end in a
	constrained suffix, `next` first appends `_hx` to it, so
	`module_tools_linux` becomes `module_tools_linux_hx.go`. A counter value
	that is itself constrained (`386`) is skipped. Other names do not change.

	How
	`hasBuildConstraintSuffix` mirrors Go's own test-file rule and
	`goodOSArchFile` from `go/build`: Go ignores everything before the first
	underscore, then looks only at the last underscore-separated word. The
	escaped base is also the key of the use counter, so a module whose own name
	already ends in `_hx` still gets a distinct file. Callers pass bases made only
	of letters, digits and underscores (module paths are sanitized first), so
	Go's cut of the name at its first dot never applies.
**/
class GoFileNameAllocator {
	/**
		Every GOOS value Go recognizes in file names.

		Verbatim copy of the keys of `KnownOS` in Go 1.26.11
		`src/internal/syslist/syslist.go`, the table `go/build` and `cmd/go` use
		for file name matching. Go 1.27.4 and Go 1.22 through 1.25 have the
		same keys. Go only adds entries, so update this copy when a new Go
		release adds one; `GoFileNameAllocatorContract` compares it with the
		active toolchain.

		`unix` is not here on purpose: Go's `UnixOS` table feeds only the
		`unix` build tag and "is not used for filename matching", so a file
		named `x_unix.go` builds on every platform.
	**/
	public static final knownOS:Array<String> = [
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"hurd",
		"illumos",
		"ios",
		"js",
		"linux",
		"nacl",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"wasip1",
		"windows",
		"zos"
	];

	/**
		Every GOARCH value Go recognizes in file names.

		Verbatim copy of the keys of `KnownArch` in Go 1.26.11
		`src/internal/syslist/syslist.go`; see `knownOS` for the version notes.
	**/
	public static final knownArch:Array<String> = [
		"386",
		"amd64",
		"amd64p32",
		"arm",
		"armbe",
		"arm64",
		"arm64be",
		"loong64",
		"mips",
		"mipsle",
		"mips64",
		"mips64le",
		"mips64p32",
		"mips64p32le",
		"ppc",
		"ppc64",
		"ppc64le",
		"riscv",
		"riscv64",
		"s390",
		"s390x",
		"sparc",
		"sparc64",
		"wasm"
	];

	final uses = new Map<String, Int>();

	public function new() {}

	/**
		Returns the next unique `.go` file name for `base`.

		Bases are compared without letter case, so `Module_A` and `module_a`
		share one counter.
	**/
	public function next(base:String):String {
		final stem = hasBuildConstraintSuffix(base) ? base + "_hx" : base;
		final key = stem.toLowerCase();
		var count = uses.exists(key) ? uses.get(key) : 0;
		while (count > 0 && hasBuildConstraintSuffix(stem + "_" + count)) {
			count++;
		}
		uses.set(key, count + 1);
		if (count == 0) {
			return stem + ".go";
		}
		return stem + "_" + count + ".go";
	}

	/**
		Whether Go would build a file named `<stem>.go` only for some targets,
		or only as a test.

		The stem must have an underscore, because Go ignores everything before
		the first one (`linux.go` builds everywhere). Then its last
		underscore-separated word decides: `test` makes a test file, and a known
		GOOS or GOARCH makes a platform file. When the last word is not one of
		these, Go applies no constraint, even to an earlier GOOS word.
	**/
	public static function hasBuildConstraintSuffix(stem:String):Bool {
		if (stem.indexOf("_") < 0) {
			return false;
		}
		final last = stem.substr(stem.lastIndexOf("_") + 1);
		return last == "test" || knownOS.contains(last) || knownArch.contains(last);
	}
}
