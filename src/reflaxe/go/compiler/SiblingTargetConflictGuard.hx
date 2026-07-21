package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
#end

#if macro
/**
	Why: source and packaged sibling targets expose different activation signals.

	What: names one sibling compiler, its supported defines, and its canonical
	source override directory when it has one.

	How: the ordered descriptor table doubles as stable diagnostic ordering, so
	users never see filesystem-dependent conflict lists.
**/
private typedef SiblingTargetDescriptor = {
	/** The stable name shown to users; descriptor order is diagnostic order. */
	final displayName:String;

	/** Defines that mean the sibling compiler was activated as a library. */
	final defineNames:Array<String>;

	/** Canonical source `_std` directory, or null when that sibling has none. */
	final sourceStdDirectory:Null<String>;
}
#end

/**
	Why: two Reflaxe target libraries can provide different implementations for
	the same Haxe stdlib module. Letting classpath order choose one makes builds
	nondeterministic and can pair one compiler with another target's stdlib.

	What: rejects a Go compilation when a known sibling compiler define or
	canonical sibling `std/<target>/_std` root is active. Ordinary helper
	classpaths and non-Go use of the bootstrap remain valid.

	How: validate once when Go compiler initialization starts, then validate
	again after all initialization macros have run. The second pass catches a
	sibling library whose HXML was expanded later in the same command while
	still running before application modules are typed.
**/
class SiblingTargetConflictGuard {
	#if macro
	static var initialized:Bool = false;

	static final SIBLING_TARGETS:Array<SiblingTargetDescriptor> = [
		{
			displayName: "genes",
			defineNames: ["genes", "genes.ts", "genes-ts"],
			sourceStdDirectory: null
		},
		{
			displayName: "reflaxe.c",
			defineNames: ["reflaxe.c"],
			sourceStdDirectory: "c"
		},
		{
			displayName: "reflaxe.elixir",
			defineNames: ["reflaxe.elixir"],
			sourceStdDirectory: "elixir"
		},
		{
			displayName: "reflaxe.ocaml",
			defineNames: ["reflaxe.ocaml"],
			sourceStdDirectory: "ocaml"
		},
		{
			displayName: "reflaxe.ruby",
			defineNames: ["reflaxe.ruby"],
			sourceStdDirectory: "ruby"
		},
		{
			displayName: "reflaxe.rust",
			defineNames: ["reflaxe.rust"],
			sourceStdDirectory: "rust"
		}
	];

	/**
		Registers the initial-configuration check for a Go build.

		The immediate pass catches paths and defines already present. The
		`onAfterInitMacros` pass closes the ordering gap between multiple `-lib`
		expansions without delaying the diagnostic until ordinary module typing.
	**/
	public static function init():Void {
		if (initialized) {
			return;
		}
		initialized = true;

		validate();
		Context.onAfterInitMacros(validate);
	}

	static function validate():Void {
		var conflicts = activeSiblingTargets();
		if (conflicts.length == 0) {
			return;
		}

		Context.fatalError('Reflaxe.Go cannot compile with competing sibling targets: ${conflicts.join(", ")}. '
			+ 'Use exactly one Reflaxe target compiler per Haxe invocation; remove the sibling library or its canonical _std classpath.',
			Context.currentPos());
	}

	static function activeSiblingTargets():Array<String> {
		var active = new Map<String, Bool>();
		var classPaths = Context.getClassPath();

		for (target in SIBLING_TARGETS) {
			for (defineName in target.defineNames) {
				if (Context.defined(defineName)) {
					active.set(target.displayName, true);
					break;
				}
			}

			if (target.sourceStdDirectory == null) {
				continue;
			}
			for (classPath in classPaths) {
				if (isCanonicalStdRoot(classPath, target.sourceStdDirectory)) {
					active.set(target.displayName, true);
					break;
				}
			}
		}

		return [
			for (target in SIBLING_TARGETS) if (active.exists(target.displayName)) target.displayName
		];
	}

	static function isCanonicalStdRoot(classPath:String, targetDirectory:String):Bool {
		var normalized = StringTools.replace(classPath, "\\", "/");
		while (StringTools.endsWith(normalized, "/")) {
			normalized = normalized.substr(0, normalized.length - 1);
		}
		normalized = normalized.toLowerCase();
		var relativeRoot = 'std/${targetDirectory.toLowerCase()}/_std';
		return normalized == relativeRoot || StringTools.endsWith(normalized, "/" + relativeRoot);
	}
	#else
	public static function init():Void {}
	#end
}
