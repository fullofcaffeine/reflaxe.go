import haxe.ds.ArraySort;
import haxe.io.Path;
import reflaxe.go.naming.GoFileNameAllocator;
import sys.FileSystem;
import sys.io.File;
import sys.io.Process;

/**
	Contract for the names of generated Go source files.

	Why
	Go silently leaves a file out of the build when its name ends in `_test`
	or in a known GOOS or GOARCH value, and generated names come from Haxe
	module paths. Every generated file must build on every platform, and names
	that Go does not constrain must stay exactly as they were, because
	existing-module mode records them with their digests. Two modules can also
	want the same name: `a_b.Util` and `a.b.Util` share the base
	`module_a_b_util`, and the numbered name of the second one used to be
	`module_a_b_util_1.go`, the own name of `a.b.Util_1`, so one file silently
	replaced the other.

	What
	Pins that `GoFileNameAllocator`:
	- holds exactly the GOOS and GOARCH tables of the active Go toolchain;
	- gives names that Go's own matcher keeps in the build for every platform
	  in `go tool dist list`, for every `_test`, GOOS, GOARCH and GOOS_GOARCH
	  ending;
	- leaves other names, including `_unix`, unchanged, and keeps names unique
	  and deterministic;
	- never gives one name twice, even when a numbered name is another base's
	  own name, and gives every list the names the old counter gave whenever
	  those were already unique;
	- gives the first base of each stem its own name in any order, and makes a
	  base's name depend only on the bases that share its stem.

	How
	Runs under `--interp` with `go` on the search path. It reads `KnownOS` and
	`KnownArch` from the toolchain's `syslist.go`, writes one small Go file per
	allocated name into a scratch module, and runs `go list` once per platform
	to compare the files Go builds with all the files written. The uniqueness
	and order properties are checked on every list of up to four bases drawn
	from a small set of bases that meet, against a copy of the old counter.
**/
class GoFileNameAllocatorContract {
	static final failures = new Array<String>();

	static function check(value:Bool, message:String):Void {
		if (!value)
			failures.push(message);
	}

	static function command(args:Array<String>):String {
		final process = new Process("go", args);
		final out = process.stdout.readAll().toString();
		final err = process.stderr.readAll().toString();
		final code = process.exitCode();
		process.close();
		if (code != 0)
			throw "GoFileNameAllocatorContract: go " + args.join(" ") + " failed: " + err;
		return out;
	}

	static function sorted(values:Array<String>):Array<String> {
		final copy = values.copy();
		copy.sort((left, right) -> left < right ? -1 : left > right ? 1 : 0);
		return copy;
	}

	/** Reads the keys of one `map[string]bool` table of Go's `syslist.go`. **/
	static function goTable(source:String, name:String):Array<String> {
		final start = new EReg("var " + name + " = map\\[string\\]bool\\{", "i");
		if (!start.match(source))
			throw "GoFileNameAllocatorContract: no " + name + " table in syslist.go";
		final body = start.matchedRight();
		final table = body.substr(0, body.indexOf("}"));
		final key = ~/"([a-z0-9]+)":\s*true/g;
		final keys = new Array<String>();
		var rest = table;
		while (key.match(rest)) {
			keys.push(key.matched(1));
			rest = key.matchedRight();
		}
		return keys;
	}

	static function checkTablesMatchToolchain():Void {
		final env = command(["env", "GOROOT", "GOVERSION"]).split("\n");
		final goRoot = StringTools.trim(env[0]);
		final goVersion = StringTools.trim(env[1]);
		var syslist = Path.join([goRoot, "src", "internal", "syslist", "syslist.go"]);
		if (!FileSystem.exists(syslist))
			syslist = Path.join([goRoot, "src", "go", "build", "syslist.go"]);
		final source = File.getContent(syslist);
		final goOS = sorted(goTable(source, "KnownOS"));
		final goArch = sorted(goTable(source, "KnownArch"));
		check(goOS.length > 0 && goArch.length > 0, "empty syslist tables in " + goVersion);
		final osMessage = 'knownOS differs from $goVersion KnownOS: ${goOS.join(",")}';
		check(sorted(GoFileNameAllocator.knownOS).join(",") == goOS.join(","), osMessage);
		final archMessage = 'knownArch differs from $goVersion KnownArch: ${goArch.join(",")}';
		check(sorted(GoFileNameAllocator.knownArch).join(",") == goArch.join(","), archMessage);
	}

	static function checkNames(label:String, bases:Array<String>, expected:Array<String>):Void {
		final actual = GoFileNameAllocator.allocate(bases);
		check(actual.join(" ") == expected.join(" "), label + ": expected " + expected.join(" ") + ", got " + actual.join(" "));
	}

	static function checkStableNames():Void {
		checkNames("unconstrained names", [
			"main",
			"support",
			"recovery_0",
			"recovery_1",
			"module_haxe_int64",
			"module_x_unix",
			"module_linux_tools",
			"module_x_testing",
			"module_x_amd64x",
			"module_x_linux_tools",
			"linux",
			"test",
			"haxego_generated_main"
		], [
			"main.go",
			"support.go",
			"recovery_0.go",
			"recovery_1.go",
			"module_haxe_int64.go",
			"module_x_unix.go",
			"module_linux_tools.go",
			"module_x_testing.go",
			"module_x_amd64x.go",
			"module_x_linux_tools.go",
			"linux.go",
			"test.go",
			"haxego_generated_main.go"
		]);
		checkNames("constrained names", [
			"module_unit_test",
			"module_tools_linux",
			"module_tools_windows_amd64",
			"module_x_arm64",
			"module_x_js",
			"module_x_linux_test",
			"recovery_386",
			"haxego_generated_module_tools_linux"
		], [
			"module_unit_test_hx.go",
			"module_tools_linux_hx.go",
			"module_tools_windows_amd64_hx.go",
			"module_x_arm64_hx.go",
			"module_x_js_hx.go",
			"module_x_linux_test_hx.go",
			"recovery_386_hx.go",
			"haxego_generated_module_tools_linux_hx.go"
		]);
		checkNames("collisions", [
			"module_a_b_util",
			"module_a_b_util",
			"Module_A_B_Util",
			"module_tools_linux",
			"module_tools_linux_hx",
			"module_tools_linux"
		], [
			"module_a_b_util.go",
			"module_a_b_util_1.go",
			"Module_A_B_Util_2.go",
			"module_tools_linux_hx.go",
			"module_tools_linux_hx_1.go",
			"module_tools_linux_hx_2.go"
		]);

		// The use counter must step over 386, a GOARCH value.
		final names = GoFileNameAllocator.allocate([for (_ in 0...388) "module_y"]);
		check(names[385] == "module_y_385.go" && names[386] == "module_y_387.go" && names[387] == "module_y_388.go",
			"the use counter must skip 386, got " + names.slice(385).join(" "));
		final unique = new Map<String, Bool>();
		for (name in names)
			unique.set(name, true);
		check([for (_ in unique.keys()) 1].length == names.length, "repeated bases must give distinct names");
	}

	/**
		Bases whose numbered names are other bases' own names, as the modules of
		`test/snapshot/core/go_file_name_collisions` give them in sorted module
		order.
	**/
	static function checkCollidingNames():Void {
		checkNames("a numbered name skips another base's own name", ["module_a_b_util", "module_a_b_util_1", "module_a_b_util"],
			["module_a_b_util.go", "module_a_b_util_1.go", "module_a_b_util_2.go"]);
		checkNames("a base keeps its own name after the bases it collides with", ["module_m_n_util", "module_m_n_util", "module_m_n_util_1"],
			["module_m_n_util.go", "module_m_n_util_2.go", "module_m_n_util_1.go"]);
		checkNames("a numbered escaped name skips another base's own name", ["module_tools_linux", "module_tools_linux_hx", "module_tools_linux_hx_1"], [
			"module_tools_linux_hx.go",
			"module_tools_linux_hx_2.go",
			"module_tools_linux_hx_1.go"
		]);
		checkNames("numbered names skip every taken number", ["x", "x", "x", "x_1", "x_2", "x_1"],
			["x.go", "x_3.go", "x_4.go", "x_1.go", "x_2.go", "x_1_1.go"]);
		checkNames("taken names are compared without letter case", ["module_a", "Module_A_1", "module_a"], ["module_a.go", "Module_A_1.go", "module_a_2.go"]);
	}

	/** Every order of `values`, repeats included. **/
	static function permutations(values:Array<String>):Array<Array<String>> {
		if (values.length <= 1)
			return [values.copy()];
		final result = new Array<Array<String>>();
		for (index in 0...values.length) {
			final rest = values.slice(0, index).concat(values.slice(index + 1));
			for (tail in permutations(rest))
				result.push([values[index]].concat(tail));
		}
		return result;
	}

	/** The `base=name` pairs of one allocation, sorted, so equal bases may swap names. **/
	static function pairs(bases:Array<String>):String {
		final names = GoFileNameAllocator.allocate(bases);
		return sorted([for (index in 0...bases.length) bases[index] + "=" + names[index]]).join(" ");
	}

	/**
		Checks that every order of `bases` gives each base the same name. Bases
		that share a stem must be equal strings here: of two different bases with
		one stem, the first in the list gets the stem's own name.
	**/
	static function checkSameNamesInEveryOrder(label:String, bases:Array<String>):Void {
		final expected = pairs(bases);
		for (order in permutations(bases)) {
			final actual = pairs(order);
			if (actual != expected) {
				check(false, label + ": order " + order.join(" ") + " gives " + actual + ", expected " + expected);
				return;
			}
		}
	}

	static function checkOrderIndependence():Void {
		checkSameNamesInEveryOrder("module_a_b_util_1 keeps its own name in every order", ["module_a_b_util", "module_a_b_util", "module_a_b_util_1"]);
		checkSameNamesInEveryOrder("escaped names in every order", ["module_tools_linux", "module_tools_linux", "module_tools_linux_hx_1"]);
		checkSameNamesInEveryOrder("chained numbers in every order", ["x", "x", "x", "x_1", "x_2", "x_1"]);
	}

	/**
		The allocator before haxe_go-eh1u: one use counter per stem, and no
		check against the names of other stems. Kept only to prove that every
		list it already named without a repeat keeps exactly those names.
	**/
	static function legacyNames(bases:Array<String>):Array<String> {
		final uses = new Map<String, Int>();
		return [
			for (base in bases) {
				final stem = stemOf(base);
				final key = stem.toLowerCase();
				var count = uses.exists(key) ? uses.get(key) : 0;
				while (count > 0 && GoFileNameAllocator.hasBuildConstraintSuffix(stem + "_" + count))
					count++;
				uses.set(key, count + 1);
				count == 0 ? stem + ".go" : stem + "_" + count + ".go";
			}
		];
	}

	static function stemOf(base:String):String {
		return GoFileNameAllocator.hasBuildConstraintSuffix(base) ? base + "_hx" : base;
	}

	static function distinctIgnoringCase(names:Array<String>):Bool {
		final seen = new Map<String, Bool>();
		for (name in names) {
			final key = name.toLowerCase();
			if (seen.exists(key))
				return false;
			seen.set(key, true);
		}
		return true;
	}

	/**
		Checks every list of one to four bases drawn from bases that meet:
		shared stems, a stem in other letter case, numbered names that are other
		stems, and the `_hx` escape. For each list:
		- no two names are equal without letter case;
		- the names equal the old counter's names whenever those were unique;
		- the first base of each stem gets `<stem>.go`;
		- moving the bases of other stems does not change any name: the list
		  gives the same name to each base as its arrangement grouped by stem,
		  which keeps the order within each stem.
	**/
	static function checkEveryShortList():Void {
		final alphabet = ["x", "X", "x_1", "x_2", "x_1_1", "x_linux", "x_linux_hx", "x_linux_hx_1"];
		final firstFailure = new Map<String, String>();
		function fail(property:String, bases:Array<String>, detail:String):Void {
			if (!firstFailure.exists(property))
				firstFailure.set(property, property + ": [" + bases.join(" ") + "] " + detail);
		}
		var lists = [new Array<String>()];
		for (_ in 0...4) {
			final longer = new Array<Array<String>>();
			for (list in lists)
				for (base in alphabet)
					longer.push(list.concat([base]));
			lists = longer;
			for (bases in lists) {
				final names = GoFileNameAllocator.allocate(bases);
				if (!distinctIgnoringCase(names))
					fail("unique names", bases, "gives " + names.join(" "));
				final legacy = legacyNames(bases);
				if (distinctIgnoringCase(legacy) && legacy.join(" ") != names.join(" "))
					fail("unchanged names", bases, "gives " + names.join(" ") + ", the old counter gave " + legacy.join(" "));
				final seenStems = new Map<String, Bool>();
				for (index in 0...bases.length) {
					final stem = stemOf(bases[index]);
					if (!seenStems.exists(stem.toLowerCase()) && names[index] != stem + ".go")
						fail("own name of the first base", bases, "gives " + names.join(" "));
					seenStems.set(stem.toLowerCase(), true);
				}
				final grouped = [for (index in 0...bases.length) index];
				ArraySort.sort(grouped, (left, right) -> {
					final leftKey = stemOf(bases[left]).toLowerCase();
					final rightKey = stemOf(bases[right]).toLowerCase();
					leftKey < rightKey ? -1 : leftKey > rightKey ? 1 : 0;
				});
				final groupedNames = GoFileNameAllocator.allocate([for (index in grouped) bases[index]]);
				for (position in 0...grouped.length)
					if (groupedNames[position] != names[grouped[position]]) {
						fail("order independence", bases, "gives " + names.join(" ") + ", grouped by stem it gives " + groupedNames.join(" "));
						break;
					}
			}
		}
		for (message in firstFailure)
			check(false, message);
	}

	/**
		Every ending Go reads from a file name: each GOOS, GOARCH and `test`
		word alone, each with `_test`, and each GOOS_GOARCH pair, in module,
		recovery and existing-module shapes.
	**/
	static function constrainedBases():Array<String> {
		final words = GoFileNameAllocator.knownOS.concat(GoFileNameAllocator.knownArch).concat(["test"]);
		final bases = [
			"main",
			"support",
			"recovery_0",
			"recovery_386",
			"haxego_generated_module_x_linux",
			"module_x_unix"
		];
		for (word in words) {
			bases.push("module_x_" + word);
			bases.push("module_x_" + word + "_test");
		}
		for (os in GoFileNameAllocator.knownOS)
			for (arch in GoFileNameAllocator.knownArch)
				bases.push("module_x_" + os + "_" + arch);
		return bases;
	}

	static function checkGoBuildsEveryFile(scratch:String):Void {
		final names = GoFileNameAllocator.allocate(constrainedBases());
		File.saveContent(Path.join([scratch, "go.mod"]), "module example.com/filenames\n\ngo 1.22\n");
		for (index in 0...names.length)
			File.saveContent(Path.join([scratch, names[index]]), "package filenames\n\nconst c" + index + " = " + index + "\n");
		final expected = sorted(names).join(" ");

		final saved = [
			for (name in ["GOOS", "GOARCH", "CGO_ENABLED", "GOWORK", "GOFLAGS"])
				name => Sys.getEnv(name)
		];
		Sys.putEnv("CGO_ENABLED", "0");
		Sys.putEnv("GOWORK", "off");
		Sys.putEnv("GOFLAGS", "");
		final platforms = [
			for (line in command(["tool", "dist", "list"]).split("\n"))
				if (StringTools.trim(line) != "") StringTools.trim(line)
		];
		check(platforms.length > 0, "go tool dist list returned no platforms");
		try {
			for (platform in platforms) {
				final parts = platform.split("/");
				Sys.putEnv("GOOS", parts[0]);
				Sys.putEnv("GOARCH", parts[1]);
				final built = command(["-C", scratch, "list", "-f", "{{join .GoFiles \" \"}}", "."]);
				final actual = sorted(StringTools.trim(built).split(" ")).join(" ");
				if (actual != expected) {
					final kept = new Map<String, Bool>();
					for (name in actual.split(" "))
						kept.set(name, true);
					final dropped = [for (name in names) if (!kept.exists(name)) name];
					check(false, platform + " leaves out " + dropped.length + " file(s), e.g. " + dropped.slice(0, 4).join(" "));
				}
			}
		} catch (error:haxe.Exception) {
			restore(saved);
			throw error;
		}
		restore(saved);
	}

	static function restore(saved:Map<String, Null<String>>):Void {
		for (name => value in saved)
			Sys.putEnv(name, value == null ? "" : value);
	}

	static function deleteTree(path:String):Void {
		for (entry in FileSystem.readDirectory(path)) {
			final child = Path.join([path, entry]);
			if (FileSystem.isDirectory(child))
				deleteTree(child);
			else
				FileSystem.deleteFile(child);
		}
		FileSystem.deleteDirectory(path);
	}

	static function main():Void {
		checkTablesMatchToolchain();
		checkStableNames();
		checkCollidingNames();
		checkOrderIndependence();
		checkEveryShortList();
		final scratch = Path.join([
			Sys.getCwd(),
			".cache",
			"go-file-name-allocator-contract-" + Std.random(0x3fffffff)
		]);
		FileSystem.createDirectory(scratch);
		try {
			checkGoBuildsEveryFile(scratch);
		} catch (error:haxe.Exception) {
			deleteTree(scratch);
			throw error;
		}
		deleteTree(scratch);
		if (failures.length > 0)
			throw "GoFileNameAllocatorContract:\n" + failures.join("\n");
		Sys.println("GoFileNameAllocatorContract: ok");
	}
}
