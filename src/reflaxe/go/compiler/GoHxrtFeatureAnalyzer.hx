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
	public static inline final FEATURE_FILESYSTEM = "filesystem";
	public static inline final FEATURE_PROCESS = "process";
	public static inline final FEATURE_BYTES = "bytes";
	public static inline final FEATURE_SSL = "ssl";
	public static inline final FEATURE_THREAD = "thread";
	public static inline final FEATURE_STACK = "stack";
	public static inline final FEATURE_ATOMIC_INT = "atomic_int";
	public static inline final FEATURE_ATOMIC_OBJECT = "atomic_object";

	static final FEATURE_ORDER = [
		FEATURE_CORE,
		FEATURE_STRING,
		FEATURE_PRINT,
		FEATURE_EXCEPTION,
		FEATURE_JSON,
		FEATURE_SYS,
		FEATURE_FILESYSTEM,
		FEATURE_PROCESS,
		FEATURE_BYTES,
		FEATURE_SSL,
		FEATURE_THREAD,
		FEATURE_STACK,
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
		return inferWithReasons(classPaths, enumPaths, shimGroups, requiresIoHelperSurface).features;
	}

	public static function inferWithReasons(classPaths:Array<String>, enumPaths:Array<String>, shimGroups:Array<String>,
			requiresIoHelperSurface:Bool):GoHxrtFeatureInference {
		var out = new Map<String, Bool>();
		var reasonsByFeature = new Map<String, Array<GoHxrtFeatureReason>>();

		inline function add(feature:String, sourceKind:String, source:String):Void {
			if (feature == null || !isKnownFeature(feature)) {
				return;
			}
			out.set(feature, true);
			addReason(reasonsByFeature, feature, sourceKind, source);
		}

		add(FEATURE_CORE, "baseline", "compiler_baseline");
		add(FEATURE_STRING, "baseline", "compiler_baseline");
		add(FEATURE_PRINT, "baseline", "compiler_baseline");
		add(FEATURE_EXCEPTION, "baseline", "compiler_baseline");

		for (path in classPaths) {
			if (path == "haxe.Json" || StringTools.startsWith(path, "haxe.json.")) {
				add(FEATURE_JSON, "class_usage", path);
			}

			if (path == "sys.io.Process") {
				add(FEATURE_PROCESS, "class_usage", path);
			}

			if (path == "sys.FileSystem") {
				add(FEATURE_FILESYSTEM, "class_usage", path);
			}

			if (path == "Sys" || path == "sys.io.File" || path == "sys.FileSystem" || StringTools.startsWith(path, "sys.")) {
				add(FEATURE_SYS, "class_usage", path);
			}

			if (StringTools.startsWith(path, "sys.ssl.")) {
				add(FEATURE_SSL, "class_usage", path);
			}

			if (StringTools.startsWith(path, "sys.thread.")) {
				add(FEATURE_THREAD, "class_usage", path);
			}

			if (StringTools.startsWith(path, "hxrt.stack.")) {
				add(FEATURE_STACK, "class_usage", path);
			}

			if (StringTools.startsWith(path, "haxe.io.")) {
				add(FEATURE_BYTES, "class_usage", path);
			}

			if (StringTools.startsWith(path, "haxe.atomic.")) {
				add(FEATURE_ATOMIC_INT, "class_usage", path);
				add(FEATURE_ATOMIC_OBJECT, "class_usage", path);
			}
		}

		for (path in enumPaths) {
			if (path == "haxe.io.Error") {
				add(FEATURE_EXCEPTION, "enum_usage", path);
				add(FEATURE_BYTES, "enum_usage", path);
			}
		}

		for (group in shimGroups) {
			switch (group) {
				case "atomic":
					add(FEATURE_ATOMIC_INT, "shim_group", group);
					add(FEATURE_ATOMIC_OBJECT, "shim_group", group);
				case "io":
					add(FEATURE_BYTES, "shim_group", group);
				case "sys", "http", "net_socket":
					add(FEATURE_SYS, "shim_group", group);
					add(FEATURE_PROCESS, "shim_group", group);
				case _:
			}
		}

		if (requiresIoHelperSurface) {
			add(FEATURE_BYTES, "io_helper_surface", "compiler_io_helpers");
		}

		return expandWithReasons([for (feature in out.keys()) feature], flattenReasons(reasonsByFeature));
	}

	public static function expandWithDependencies(features:Array<String>):Array<String> {
		return expandWithReasons(features, []).features;
	}

	public static function expandWithReasons(features:Array<String>, baseReasons:Array<GoHxrtFeatureReason>):GoHxrtFeatureInference {
		var selected = new Map<String, Bool>();
		for (feature in features) {
			if (feature != null && isKnownFeature(feature)) {
				selected.set(feature, true);
			}
		}
		selected.set(FEATURE_CORE, true);

		var reasonsByFeature = new Map<String, Array<GoHxrtFeatureReason>>();
		for (entry in baseReasons) {
			if (entry == null || entry.feature == null || !selected.exists(entry.feature)) {
				continue;
			}
			addReason(reasonsByFeature, entry.feature, entry.sourceKind, entry.source);
		}
		if (!reasonsByFeature.exists(FEATURE_CORE)) {
			addReason(reasonsByFeature, FEATURE_CORE, "baseline", "compiler_baseline");
		}

		var changed = true;
		while (changed) {
			changed = false;
			var keys = [for (feature in selected.keys()) feature];
			keys.sort(compareFeatureNames);
			for (feature in keys) {
				for (dependency in featureDependencies(feature)) {
					if (!selected.exists(dependency)) {
						selected.set(dependency, true);
						addReason(reasonsByFeature, dependency, "dependency_edge", feature + "->" + dependency);
						changed = true;
					}
				}
			}
		}

		var out = [for (feature in selected.keys()) feature];
		out.sort(compareFeatureNames);
		return {
			features: out,
			reasons: flattenReasons(reasonsByFeature)
		};
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
			case FEATURE_FILESYSTEM:
				[FEATURE_STRING];
			case FEATURE_PROCESS:
				[FEATURE_SYS];
			case FEATURE_BYTES:
				[FEATURE_CORE];
			case FEATURE_SSL:
				[FEATURE_STRING, FEATURE_EXCEPTION, FEATURE_BYTES];
			case FEATURE_THREAD:
				[FEATURE_CORE, FEATURE_EXCEPTION];
			case FEATURE_STACK:
				[FEATURE_STRING];
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
			case FEATURE_FILESYSTEM:
				["filesystem.go"];
			case FEATURE_PROCESS:
				["process.go"];
			case FEATURE_BYTES:
				["bytes.go"];
			case FEATURE_SSL:
				["ssl.go"];
			case FEATURE_THREAD:
				["thread.go"];
			case FEATURE_STACK:
				["stack.go"];
			case FEATURE_ATOMIC_INT:
				["atomic_int.go"];
			case FEATURE_ATOMIC_OBJECT:
				["atomic_object.go"];
			case _:
				[];
		};
	}

	static function addReason(reasonsByFeature:Map<String, Array<GoHxrtFeatureReason>>, feature:String, sourceKind:String, source:String):Void {
		var list = reasonsByFeature.get(feature);
		if (list == null) {
			list = [];
			reasonsByFeature.set(feature, list);
		}
		for (entry in list) {
			if (entry.sourceKind == sourceKind && entry.source == source) {
				return;
			}
		}
		list.push({
			feature: feature,
			sourceKind: sourceKind,
			source: source
		});
	}

	static function flattenReasons(reasonsByFeature:Map<String, Array<GoHxrtFeatureReason>>):Array<GoHxrtFeatureReason> {
		var out = new Array<GoHxrtFeatureReason>();
		var featureNames = [for (feature in reasonsByFeature.keys()) feature];
		featureNames.sort(compareFeatureNames);
		for (feature in featureNames) {
			var entries = reasonsByFeature.get(feature);
			if (entries == null) {
				continue;
			}
			entries.sort(compareReasons);
			for (entry in entries) {
				out.push({
					feature: entry.feature,
					sourceKind: entry.sourceKind,
					source: entry.source
				});
			}
		}
		return out;
	}

	static function compareFeatureNames(a:String, b:String):Int {
		var ai = FEATURE_ORDER.indexOf(a);
		var bi = FEATURE_ORDER.indexOf(b);
		if (ai == bi) {
			return Reflect.compare(a, b);
		}
		return ai - bi;
	}

	static function compareReasons(a:GoHxrtFeatureReason, b:GoHxrtFeatureReason):Int {
		var featureOrder = compareFeatureNames(a.feature, b.feature);
		if (featureOrder != 0) {
			return featureOrder;
		}
		var kindOrder = Reflect.compare(a.sourceKind, b.sourceKind);
		if (kindOrder != 0) {
			return kindOrder;
		}
		return Reflect.compare(a.source, b.source);
	}
}

typedef GoHxrtFeatureReason = {
	var feature:String;
	var sourceKind:String;
	var source:String;
}

typedef GoHxrtFeatureInference = {
	var features:Array<String>;
	var reasons:Array<GoHxrtFeatureReason>;
}
#end
