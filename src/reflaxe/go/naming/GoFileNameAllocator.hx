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

	Two modules can also want one name. `a_b.Util` and `a.b.Util` both have the
	base `module_a_b_util`, and the old counter gave the second one
	`module_a_b_util_1.go` without checking that this is the own name of a
	module `a.b.Util_1`. One file then silently replaced the other.

	What
	`allocate` names all the files of one output together. When the base would
	end in a constrained suffix, it first appends `_hx` to it, so
	`module_tools_linux` becomes `module_tools_linux_hx.go`; the result is the
	stem. The first base of each stem gets `<stem>.go`. Later bases of that stem
	get `<stem>_<n>.go` with the smallest free `n` from 1 up: `n` skips a
	constrained counter value (`386`) and every name that is the `<stem>.go` of
	another base in the list. So the module `a.b.Util_1` keeps
	`module_a_b_util_1.go` and the second `module_a_b_util` base gets
	`module_a_b_util_2.go`. Every list that the old counter named without a
	repeat keeps exactly the same names.

	How
	`hasBuildConstraintSuffix` mirrors Go's own test-file rule and
	`goodOSArchFile` from `go/build`: Go ignores everything before the first
	underscore, then looks only at the last underscore-separated word. The
	escaped base is also the key of the counter, so a module whose own name
	already ends in `_hx` still gets a distinct file. Callers pass bases made only
	of letters, digits and underscores (module paths are sanitized first), so
	Go's cut of the name at its first dot never applies.

	All names are compared without letter case, because macOS and Windows file
	systems do. `allocate` first marks every stem's own name as taken, then
	hands out names in list order. A numbered name `<stem>_<n>` has the digits
	`n` after its last underscore, so it can only equal a numbered name of the
	same stem or another stem's own name, and both are skipped. As a result the
	names cannot repeat, the first base of each stem always gets its own name,
	and a base's name depends only on the order of the bases that share its
	stem, never on where the other bases are in the list.
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

	/**
		Returns one distinct `.go` file name for each of `bases`, in the same
		order.

		Bases are compared without letter case, so `Module_A` and `module_a`
		share one stem. The caller passes every file of the output in one call;
		see the class documentation for the rule.
	**/
	public static function allocate(bases:Array<String>):Array<String> {
		final stems = [for (base in bases) hasBuildConstraintSuffix(base) ? base + "_hx" : base];
		final taken = new Map<String, Bool>();
		for (stem in stems) {
			taken.set(stem.toLowerCase(), true);
		}
		final nextCount = new Map<String, Int>();
		final names = new Array<String>();
		for (stem in stems) {
			final key = stem.toLowerCase();
			final number = nextCount.get(key);
			if (number == null) {
				nextCount.set(key, 1);
				names.push(stem + ".go");
				continue;
			}
			var count:Int = number;
			while (hasBuildConstraintSuffix(stem + "_" + count) || taken.exists(key + "_" + count)) {
				count++;
			}
			nextCount.set(key, count + 1);
			taken.set(key + "_" + count, true);
			names.push(stem + "_" + count + ".go");
		}
		return names;
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
