package reflaxe.go.compiler;

#if macro
/**
	GoHxrtFeatureAnalyzer

	Why
	- Runtime slicing rules should live in one deterministic planner instead of being spread
	  across emitter/compiler plumbing.
	- Selective runtime copy needs stable feature ordering and dependency expansion.

	What
	- Infers `hxrt` feature requirements from used class/enum paths and shim groups.
	- Expands feature dependencies and maps features to concrete runtime file names.

	How
	- `inferFromUsage(...)` applies conservative rules so selective mode stays safe.
	- `expandWithDependencies(...)` normalizes + sorts feature sets.
	- `filesForFeatures(...)` returns deterministic runtime file lists for emitter copy.
**/
class GoHxrtFeatureAnalyzer {
	public static inline final FEATURE_CORE = "core";
	public static inline final FEATURE_STRING = "string";
	public static inline final FEATURE_PRINT = "print";
	public static inline final FEATURE_EXCEPTION = "exception";
	public static inline final FEATURE_JSON = "json";
	public static inline final FEATURE_SYS = "sys";
	public static inline final FEATURE_PROCESS = "process";
	public static inline final FEATURE_BYTES = "bytes";
	public static inline final FEATURE_ATOMIC_INT = "atomic_int";
	public static inline final FEATURE_ATOMIC_OBJECT = "atomic_object";

	static final FEATURE_ORDER = [
		FEATURE_CORE,
		FEATURE_STRING,
		FEATURE_PRINT,
		FEATURE_EXCEPTION,
		FEATURE_JSON,
		FEATURE_SYS,
		FEATURE_PROCESS,
		FEATURE_BYTES,
		FEATURE_ATOMIC_INT,
		FEATURE_ATOMIC_OBJECT
	];

	public static function knownFeatures():Array<String> {
		return FEATURE_ORDER.copy();
	}

	public static function isKnownFeature(feature:String):Bool {
		return FEATURE_ORDER.indexOf(feature) >= 0;
	}

	public static function inferFromUsage(classPaths:Array<String>, enumPaths:Array<String>, shimGroups:Array<String>,
			requiresIoHelperSurface:Bool):Array<String> {
		var out = new Map<String, Bool>();
		add(out, FEATURE_CORE);
		add(out, FEATURE_STRING);
		add(out, FEATURE_PRINT);
		add(out, FEATURE_EXCEPTION);

		for (path in classPaths) {
			if (path == "haxe.Json" || StringTools.startsWith(path, "haxe.json.")) {
				add(out, FEATURE_JSON);
			}

			if (path == "sys.io.Process") {
				add(out, FEATURE_PROCESS);
			}

			if (path == "Sys" || path == "sys.io.File" || path == "sys.FileSystem" || StringTools.startsWith(path, "sys.")) {
				add(out, FEATURE_SYS);
			}

			if (StringTools.startsWith(path, "haxe.io.")) {
				add(out, FEATURE_BYTES);
			}

			if (StringTools.startsWith(path, "haxe.atomic.")) {
				add(out, FEATURE_ATOMIC_INT);
				add(out, FEATURE_ATOMIC_OBJECT);
			}
		}

		for (path in enumPaths) {
			if (path == "haxe.io.Error") {
				add(out, FEATURE_EXCEPTION);
				add(out, FEATURE_BYTES);
			}
		}

		for (group in shimGroups) {
			switch (group) {
				case "atomic":
					add(out, FEATURE_ATOMIC_INT);
					add(out, FEATURE_ATOMIC_OBJECT);
				case "io":
					add(out, FEATURE_BYTES);
				case "sys", "filesystem", "http", "net_socket":
					add(out, FEATURE_SYS);
					add(out, FEATURE_PROCESS);
				case _:
			}
		}

		if (requiresIoHelperSurface) {
			add(out, FEATURE_BYTES);
		}

		return expandWithDependencies([for (feature in out.keys()) feature]);
	}

	public static function expandWithDependencies(features:Array<String>):Array<String> {
		var selected = new Map<String, Bool>();
		for (feature in features) {
			if (feature != null && isKnownFeature(feature)) {
				selected.set(feature, true);
			}
		}
		selected.set(FEATURE_CORE, true);

		var changed = true;
		while (changed) {
			changed = false;
			var keys = [for (feature in selected.keys()) feature];
			for (feature in keys) {
				for (dependency in featureDependencies(feature)) {
					if (!selected.exists(dependency)) {
						selected.set(dependency, true);
						changed = true;
					}
				}
			}
		}

		var out = [for (feature in selected.keys()) feature];
		out.sort(function(a, b) {
			var ai = FEATURE_ORDER.indexOf(a);
			var bi = FEATURE_ORDER.indexOf(b);
			if (ai == bi) {
				return Reflect.compare(a, b);
			}
			return ai - bi;
		});
		return out;
	}

	public static function filesForFeatures(features:Array<String>):Array<String> {
		var selected = expandWithDependencies(features);
		var files = new Map<String, Bool>();
		for (feature in selected) {
			for (fileName in featureFiles(feature)) {
				files.set(fileName, true);
			}
		}
		var out = [for (fileName in files.keys()) fileName];
		out.sort(Reflect.compare);
		return out;
	}

	static function add(map:Map<String, Bool>, feature:String):Void {
		map.set(feature, true);
	}

	static function featureDependencies(feature:String):Array<String> {
		return switch (feature) {
			case FEATURE_STRING:
				[FEATURE_CORE];
			case FEATURE_PRINT:
				[FEATURE_STRING];
			case FEATURE_EXCEPTION:
				[FEATURE_STRING];
			case FEATURE_JSON:
				[FEATURE_CORE];
			case FEATURE_SYS:
				[FEATURE_STRING];
			case FEATURE_PROCESS:
				[FEATURE_SYS];
			case FEATURE_BYTES:
				[FEATURE_CORE];
			case _:
				[];
		};
	}

	static function featureFiles(feature:String):Array<String> {
		return switch (feature) {
			case FEATURE_CORE:
				["hxrt.go", "core.go"];
			case FEATURE_STRING:
				["string.go"];
			case FEATURE_PRINT:
				["print.go"];
			case FEATURE_EXCEPTION:
				["exception.go"];
			case FEATURE_JSON:
				["json.go"];
			case FEATURE_SYS:
				["sys.go"];
			case FEATURE_PROCESS:
				["process.go"];
			case FEATURE_BYTES:
				["bytes.go"];
			case FEATURE_ATOMIC_INT:
				["atomic_int.go"];
			case FEATURE_ATOMIC_OBJECT:
				["atomic_object.go"];
			case _:
				[];
		};
	}
}
#end
