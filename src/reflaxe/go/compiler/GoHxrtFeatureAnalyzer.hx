package reflaxe.go.compiler;

#if macro
/**
	Closed identifiers for selectively packaged `hxrt` runtime features.

	Why
	Feature names cross inference, dependency, packaging, and report boundaries.
	A typo in any one of those places can silently omit required runtime code.

	What
	Each value names one runtime slice understood by the compiler-owned registry.

	How
	The analyzer keeps compatibility constants for existing call sites, but those
	constants now reference this closed type. Strings enter only through the
	explicit manual-feature parser and must pass `isKnownFeature(...)`.
**/
enum abstract GoHxrtFeatureId(String) to String {
	var HxrtCore = "core";
	var HxrtArray = "array";
	var HxrtArraySort = "array_sort";
	var HxrtString = "string";
	var HxrtStringCompare = "string_compare";
	var HxrtEquality = "equality";
	var HxrtPrint = "print";
	var HxrtException = "exception";
	var HxrtJson = "json";
	var HxrtSys = "sys";
	var HxrtTerminal = "terminal";
	var HxrtFileIo = "file_io";
	var HxrtFilesystem = "filesystem";
	var HxrtProcess = "process";
	var HxrtSocket = "socket";
	var HxrtHttp = "http";
	var HxrtBytes = "bytes";
	var HxrtDate = "date";
	var HxrtMath = "math";
	var HxrtCrypto = "crypto";
	var HxrtZip = "zip";
	var HxrtSsl = "ssl";
	var HxrtSocketSsl = "socket_ssl";
	var HxrtThread = "thread";
	var HxrtStack = "stack";
	var HxrtTemplate = "template";
	var HxrtReflection = "reflection";
	var HxrtRegex = "regex";
	var HxrtSerialization = "serialization";
	var HxrtEnumValue = "enum_value";
	var HxrtMapInt = "map_int";
	var HxrtMapString = "map_string";
	var HxrtMapObject = "map_object";
	var HxrtAtomicInt = "atomic_int";
	var HxrtAtomicObject = "atomic_object";
}

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
	public static inline final FEATURE_CORE:GoHxrtFeatureId = GoHxrtFeatureId.HxrtCore;
	public static inline final FEATURE_ARRAY:GoHxrtFeatureId = GoHxrtFeatureId.HxrtArray;
	public static inline final FEATURE_ARRAY_SORT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtArraySort;
	public static inline final FEATURE_STRING:GoHxrtFeatureId = GoHxrtFeatureId.HxrtString;
	public static inline final FEATURE_STRING_COMPARE:GoHxrtFeatureId = GoHxrtFeatureId.HxrtStringCompare;
	public static inline final FEATURE_EQUALITY:GoHxrtFeatureId = GoHxrtFeatureId.HxrtEquality;
	public static inline final FEATURE_PRINT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtPrint;
	public static inline final FEATURE_EXCEPTION:GoHxrtFeatureId = GoHxrtFeatureId.HxrtException;
	public static inline final FEATURE_JSON:GoHxrtFeatureId = GoHxrtFeatureId.HxrtJson;
	public static inline final FEATURE_SYS:GoHxrtFeatureId = GoHxrtFeatureId.HxrtSys;
	public static inline final FEATURE_TERMINAL:GoHxrtFeatureId = GoHxrtFeatureId.HxrtTerminal;
	public static inline final FEATURE_FILE_IO:GoHxrtFeatureId = GoHxrtFeatureId.HxrtFileIo;
	public static inline final FEATURE_FILESYSTEM:GoHxrtFeatureId = GoHxrtFeatureId.HxrtFilesystem;
	public static inline final FEATURE_PROCESS:GoHxrtFeatureId = GoHxrtFeatureId.HxrtProcess;
	public static inline final FEATURE_SOCKET:GoHxrtFeatureId = GoHxrtFeatureId.HxrtSocket;
	public static inline final FEATURE_HTTP:GoHxrtFeatureId = GoHxrtFeatureId.HxrtHttp;
	public static inline final FEATURE_BYTES:GoHxrtFeatureId = GoHxrtFeatureId.HxrtBytes;
	public static inline final FEATURE_DATE:GoHxrtFeatureId = GoHxrtFeatureId.HxrtDate;
	public static inline final FEATURE_MATH:GoHxrtFeatureId = GoHxrtFeatureId.HxrtMath;
	public static inline final FEATURE_CRYPTO:GoHxrtFeatureId = GoHxrtFeatureId.HxrtCrypto;
	public static inline final FEATURE_ZIP:GoHxrtFeatureId = GoHxrtFeatureId.HxrtZip;
	public static inline final FEATURE_SSL:GoHxrtFeatureId = GoHxrtFeatureId.HxrtSsl;
	public static inline final FEATURE_SOCKET_SSL:GoHxrtFeatureId = GoHxrtFeatureId.HxrtSocketSsl;
	public static inline final FEATURE_THREAD:GoHxrtFeatureId = GoHxrtFeatureId.HxrtThread;
	public static inline final FEATURE_STACK:GoHxrtFeatureId = GoHxrtFeatureId.HxrtStack;
	public static inline final FEATURE_TEMPLATE:GoHxrtFeatureId = GoHxrtFeatureId.HxrtTemplate;
	public static inline final FEATURE_REFLECTION:GoHxrtFeatureId = GoHxrtFeatureId.HxrtReflection;
	public static inline final FEATURE_REGEX:GoHxrtFeatureId = GoHxrtFeatureId.HxrtRegex;
	public static inline final FEATURE_SERIALIZATION:GoHxrtFeatureId = GoHxrtFeatureId.HxrtSerialization;
	public static inline final FEATURE_ENUM_VALUE:GoHxrtFeatureId = GoHxrtFeatureId.HxrtEnumValue;
	public static inline final FEATURE_MAP_INT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtMapInt;
	public static inline final FEATURE_MAP_STRING:GoHxrtFeatureId = GoHxrtFeatureId.HxrtMapString;
	public static inline final FEATURE_MAP_OBJECT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtMapObject;
	public static inline final FEATURE_ATOMIC_INT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtAtomicInt;
	public static inline final FEATURE_ATOMIC_OBJECT:GoHxrtFeatureId = GoHxrtFeatureId.HxrtAtomicObject;

	static final FEATURE_ORDER:Array<String> = [
		FEATURE_CORE,
		FEATURE_ARRAY,
		FEATURE_ARRAY_SORT,
		FEATURE_STRING,
		FEATURE_STRING_COMPARE,
		FEATURE_EQUALITY,
		FEATURE_PRINT,
		FEATURE_EXCEPTION,
		FEATURE_JSON,
		FEATURE_SYS,
		FEATURE_TERMINAL,
		FEATURE_FILE_IO,
		FEATURE_FILESYSTEM,
		FEATURE_PROCESS,
		FEATURE_SOCKET,
		FEATURE_HTTP,
		FEATURE_BYTES,
		FEATURE_DATE,
		FEATURE_MATH,
		FEATURE_CRYPTO,
		FEATURE_ZIP,
		FEATURE_SSL,
		FEATURE_SOCKET_SSL,
		FEATURE_THREAD,
		FEATURE_STACK,
		FEATURE_TEMPLATE,
		FEATURE_REFLECTION,
		FEATURE_REGEX,
		FEATURE_SERIALIZATION,
		FEATURE_ENUM_VALUE,
		FEATURE_MAP_INT,
		FEATURE_MAP_STRING,
		FEATURE_MAP_OBJECT,
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
			?requiresEqualitySurface:Bool = false):Array<String> {
		return inferWithReasons(classPaths, enumPaths, shimGroups, requiresEqualitySurface).features;
	}

	public static function inferWithReasons(classPaths:Array<String>, enumPaths:Array<String>, shimGroups:Array<String>,
			?requiresEqualitySurface:Bool = false):GoHxrtFeatureInference {
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
			var isProcessSurface = path == "sys.io.Process" || StringTools.startsWith(path, "sys.io._Process.");
			var isHttpSurface = path == "sys.Http" || StringTools.startsWith(path, "hxrt.http.");
			var isSocketSurface = path == "sys.net.Host"
				|| path == "sys.net.Socket"
				|| path == "sys.net.UdpSocket"
				|| StringTools.startsWith(path, "sys.net._SocketIO.")
				|| StringTools.startsWith(path, "hxrt.net.");
			var isSocketSslSurface = path == "sys.ssl.Socket" || path == "sys.ssl._Socket.Socket_Impl_" || path == "hxrt.ssl.NativeSocket";
			if (path == "haxe.Json" || StringTools.startsWith(path, "haxe.json.")) {
				add(FEATURE_JSON, "class_usage", path);
			}

			if (isProcessSurface) {
				add(FEATURE_PROCESS, "class_usage", path);
			}
			if (isHttpSurface) {
				add(FEATURE_HTTP, "class_usage", path);
			}
			if (path == "hxrt.process.NativeProcess") {
				add(FEATURE_PROCESS, "class_usage", path);
			}
			if (isSocketSurface) {
				add(FEATURE_SOCKET, "class_usage", path);
			}
			if (isSocketSslSurface) {
				add(FEATURE_SOCKET_SSL, "class_usage", path);
			}

			if (path == "hxrt.sys.NativeSys") {
				add(FEATURE_SYS, "class_usage", path);
			}
			if (path == "hxrt.sys.NativeTerminal") {
				add(FEATURE_TERMINAL, "class_usage", path);
			}

			if (path == "hxrt.fs.NativeFile" || path == "sys.io.File" || path == "sys.io.FileInput" || path == "sys.io.FileOutput") {
				add(FEATURE_FILE_IO, "class_usage", path);
			}

			if (path == "sys.FileSystem") {
				add(FEATURE_FILESYSTEM, "class_usage", path);
			}

			var hasDedicatedRuntimeSlice = isProcessSurface || isHttpSurface || isSocketSurface || isSocketSslSurface || path == "sys.io.File"
				|| path == "sys.io.FileInput" || path == "sys.io.FileOutput";
			if (!hasDedicatedRuntimeSlice && (path == "sys.FileSystem" || StringTools.startsWith(path, "sys."))) {
				add(FEATURE_SYS, "class_usage", path);
			}

			if (StringTools.startsWith(path, "sys.ssl.")) {
				add(FEATURE_SSL, "class_usage", path);
			}
			if (StringTools.startsWith(path, "hxrt.ssl.") && path != "hxrt.ssl.NativeSocket") {
				add(FEATURE_SSL, "class_usage", path);
			}

			if (StringTools.startsWith(path, "sys.thread.")) {
				add(FEATURE_THREAD, "class_usage", path);
			}

			if (StringTools.startsWith(path, "hxrt.stack.")) {
				add(FEATURE_STACK, "class_usage", path);
			}

			if (path == "haxe.Template" || path == "hxrt.template.NativeTemplate") {
				add(FEATURE_TEMPLATE, "class_usage", path);
			}

			if (path == "Reflect" || path == "hxrt.reflect.NativeReflect") {
				add(FEATURE_REFLECTION, "class_usage", path);
			}

			if (path == "EReg" || StringTools.startsWith(path, "hxrt.regex.")) {
				add(FEATURE_REGEX, "class_usage", path);
			}

			if (path == "haxe.Serializer" || path == "haxe.Unserializer" || StringTools.startsWith(path, "hxrt.serialization.")) {
				add(FEATURE_SERIALIZATION, "class_usage", path);
			}

			if (StringTools.startsWith(path, "haxe.io.")) {
				add(FEATURE_BYTES, "class_usage", path);
			}

			if (path == "Date" || path == "hxrt.date.NativeDate" || path == "hxrt.date.DateParts") {
				add(FEATURE_DATE, "class_usage", path);
			}

			if (path == "Math" || path == "hxrt.math.NativeMathInt") {
				add(FEATURE_MATH, "class_usage", path);
			}

			if (StringTools.startsWith(path, "haxe.crypto.") || path == "hxrt.crypto.NativeCrypto") {
				add(FEATURE_CRYPTO, "class_usage", path);
			}

			if (path == "haxe.zip.Compress" || path == "haxe.zip.Uncompress" || path == "hxrt.zip.NativeZip") {
				add(FEATURE_ZIP, "class_usage", path);
			}

			var isIntMapSurface = path == "haxe.ds.IntMap"
				|| path == "hxrt.collections.IntMapHandle"
				|| path == "hxrt.collections.NativeIntMap";
			if (isIntMapSurface) {
				add(FEATURE_MAP_INT, "class_usage", path);
			}

			var isStringMapSurface = path == "haxe.ds.StringMap"
				|| path == "hxrt.collections.StringMapHandle"
				|| path == "hxrt.collections.NativeStringMap";
			if (isStringMapSurface) {
				add(FEATURE_MAP_STRING, "class_usage", path);
			}

			var isObjectMapSurface = path == "haxe.ds.ObjectMap"
				|| path == "hxrt.collections.ObjectMapHandle"
				|| path == "hxrt.collections.NativeObjectMap";
			if (isObjectMapSurface) {
				add(FEATURE_MAP_OBJECT, "class_usage", path);
			}

			if (path == "haxe.ds.EnumValueMap" || path == "hxrt.collections.NativeEnumValue") {
				add(FEATURE_ENUM_VALUE, "class_usage", path);
			}

			var isAtomicIntSurface = path == "haxe.atomic.AtomicBool"
				|| path == "haxe.atomic.AtomicInt"
				|| path == "haxe.atomic._AtomicBool.AtomicBool_Impl_"
				|| path == "haxe.atomic._AtomicInt.AtomicInt_Impl_"
				|| path == "hxrt.atomic.AtomicIntHandle"
				|| path == "hxrt.atomic.NativeAtomicInt";
			if (isAtomicIntSurface) {
				add(FEATURE_ATOMIC_INT, "class_usage", path);
			}

			var isAtomicObjectSurface = path == "haxe.atomic.AtomicObject"
				|| path == "haxe.atomic._AtomicObject.AtomicObject_Impl_"
				|| path == "hxrt.atomic.AtomicObjectHandle"
				|| path == "hxrt.atomic.NativeAtomicObject";
			if (isAtomicObjectSurface) {
				add(FEATURE_ATOMIC_OBJECT, "class_usage", path);
			}
		}

		for (path in enumPaths) {
			if (path == "haxe.io.Error") {
				add(FEATURE_EXCEPTION, "enum_usage", path);
				add(FEATURE_BYTES, "enum_usage", path);
			}
		}

		if (requiresEqualitySurface) {
			add(FEATURE_EQUALITY, "compiler_surface", "erased_haxe_equality");
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

	/**
		Returns only the files owned directly by one known capability.

		Why
		The runtime manifest lists dependencies as their own capabilities, so
		repeating dependency files under every parent would obscure ownership.

		What
		Exposes the closed feature-to-file registry without dependency expansion.

		How
		Unknown strings return an empty list; manual strings are validated before
		reaching the manifest.
	**/
	public static function directFilesForFeature(feature:String):Array<String> {
		return featureFiles(feature).copy();
	}

	static function featureDependencies(feature:String):Array<String> {
		return switch (feature) {
			case FEATURE_ARRAY:
				[FEATURE_CORE];
			case FEATURE_ARRAY_SORT:
				[FEATURE_ARRAY];
			case FEATURE_STRING:
				[FEATURE_CORE];
			case FEATURE_STRING_COMPARE:
				[FEATURE_STRING];
			case FEATURE_EQUALITY:
				[FEATURE_STRING];
			case FEATURE_PRINT:
				[FEATURE_STRING];
			case FEATURE_EXCEPTION:
				[FEATURE_STRING];
			case FEATURE_JSON:
				[FEATURE_CORE, FEATURE_ARRAY];
			case FEATURE_SYS:
				[FEATURE_STRING];
			case FEATURE_TERMINAL:
				[FEATURE_STRING];
			case FEATURE_FILE_IO:
				[FEATURE_STRING];
			case FEATURE_FILESYSTEM:
				[FEATURE_STRING];
			case FEATURE_PROCESS:
				[FEATURE_STRING];
			case FEATURE_SOCKET:
				[FEATURE_STRING, FEATURE_EXCEPTION];
			case FEATURE_HTTP:
				[FEATURE_STRING, FEATURE_BYTES, FEATURE_SOCKET];
			case FEATURE_BYTES:
				[FEATURE_CORE];
			case FEATURE_DATE:
				[FEATURE_STRING, FEATURE_EXCEPTION];
			case FEATURE_MATH:
				[];
			case FEATURE_CRYPTO:
				[FEATURE_STRING, FEATURE_EXCEPTION];
			case FEATURE_ZIP:
				[FEATURE_EXCEPTION];
			case FEATURE_SSL:
				[FEATURE_STRING, FEATURE_EXCEPTION, FEATURE_BYTES];
			case FEATURE_SOCKET_SSL:
				[FEATURE_SOCKET, FEATURE_SSL];
			case FEATURE_THREAD:
				[FEATURE_CORE, FEATURE_EXCEPTION];
			case FEATURE_STACK:
				[FEATURE_STRING];
			case FEATURE_TEMPLATE:
				[FEATURE_CORE, FEATURE_ARRAY];
			case FEATURE_REFLECTION:
				[FEATURE_STRING, FEATURE_ARRAY];
			case FEATURE_REGEX:
				[FEATURE_STRING, FEATURE_EXCEPTION];
			case FEATURE_SERIALIZATION:
				[FEATURE_STRING, FEATURE_EQUALITY];
			case FEATURE_MAP_STRING:
				[FEATURE_STRING];
			case FEATURE_MAP_OBJECT:
				[FEATURE_EXCEPTION];
			case FEATURE_ATOMIC_OBJECT:
				[FEATURE_EQUALITY];
			case _:
				[];
		};
	}

	static function featureFiles(feature:String):Array<String> {
		return switch (feature) {
			case FEATURE_CORE:
				["hxrt.go", "core.go"];
			case FEATURE_ARRAY:
				["array.go"];
			case FEATURE_ARRAY_SORT:
				["array_sort.go"];
			case FEATURE_STRING:
				["string.go"];
			case FEATURE_STRING_COMPARE:
				["string_compare.go"];
			case FEATURE_EQUALITY:
				["equality.go"];
			case FEATURE_PRINT:
				["print.go"];
			case FEATURE_EXCEPTION:
				["exception.go"];
			case FEATURE_JSON:
				["json.go"];
			case FEATURE_SYS:
				["sys.go"];
			case FEATURE_TERMINAL:
				[
					"terminal.go",
					"terminal_darwin.go",
					"terminal_linux.go",
					"terminal_posix.go",
					"terminal_unsupported.go",
					"terminal_windows.go"
				];
			case FEATURE_FILE_IO:
				["file.go"];
			case FEATURE_FILESYSTEM:
				["filesystem.go"];
			case FEATURE_PROCESS:
				["process.go"];
			case FEATURE_SOCKET:
				[
					"socket.go",
					"socket_broadcast_posix.go",
					"socket_broadcast_unsupported.go",
					"socket_broadcast_windows.go",
					"socket_listener_posix.go",
					"socket_listener_unsupported.go",
					"socket_listener_windows.go",
					"socket_readiness_darwin.go",
					"socket_readiness_linux_32.go",
					"socket_readiness_linux_64.go",
					"socket_readiness_unsupported.go"
				];
			case FEATURE_HTTP:
				["http.go"];
			case FEATURE_BYTES:
				["bytes.go"];
			case FEATURE_DATE:
				["date.go"];
			case FEATURE_MATH:
				["math.go"];
			case FEATURE_CRYPTO:
				["crypto.go"];
			case FEATURE_ZIP:
				["zip.go"];
			case FEATURE_SSL:
				["ssl.go"];
			case FEATURE_SOCKET_SSL:
				["socket_ssl.go"];
			case FEATURE_THREAD:
				["thread.go"];
			case FEATURE_STACK:
				["stack.go"];
			case FEATURE_TEMPLATE:
				["template.go"];
			case FEATURE_REFLECTION:
				["reflect.go"];
			case FEATURE_REGEX:
				["regex.go"];
			case FEATURE_SERIALIZATION:
				["serialization.go"];
			case FEATURE_ENUM_VALUE:
				["enum_value.go"];
			case FEATURE_MAP_INT:
				["map_int.go"];
			case FEATURE_MAP_STRING:
				["map_string.go"];
			case FEATURE_MAP_OBJECT:
				["map_object.go"];
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
