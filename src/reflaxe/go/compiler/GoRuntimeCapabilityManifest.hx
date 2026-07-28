package reflaxe.go.compiler;

#if (macro || reflaxe_runtime)
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureId;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureReason;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;

/**
	One selected runtime capability and the evidence that selected it.

	Why
	A flat feature list cannot prove which files belong to a capability or why
	that capability entered the generated project.

	What
	Groups one closed `hxrt` feature ID with its deterministic file list and all
	typed-usage, surface-contract, dependency, define, or compatibility reasons.

	How
	`GoRuntimeCapabilityManifest.build(...)` is the only production constructor.
	It expands dependencies first, then publishes deeply read-only lists.
**/
typedef GoRuntimeCapabilitySelection = {
	final id:GoHxrtFeatureId;
	final files:GoImmutableList<String>;
	final reasons:GoImmutableList<GoHxrtFeatureReason>;
}

/**
	The immutable authority consumed by runtime copying and reporting.

	Why
	When the copier and report independently expand strings, the report can claim
	a file set that differs from the generated project.

	What
	Contains the resolved copy mode, typed feature sets, exact runtime files, and
	per-capability evidence for one build.

	How
	The compiler builds it once after typed lowering has supplied final runtime
	requirements. Both output and `hxrt_plan.*` consume the same snapshot.
**/
typedef GoRuntimeCapabilityManifestSnapshot = {
	final schemaVersion:Int;
	final authority:String;
	final fullCopy:Bool;
	final selectiveEnabled:Bool;
	final inferenceDisabled:Bool;
	final manualFeatures:GoImmutableList<GoHxrtFeatureId>;
	final inferredFeatures:GoImmutableList<GoHxrtFeatureId>;
	final selectedFeatures:GoImmutableList<GoHxrtFeatureId>;
	final files:GoImmutableList<String>;
	final reasons:GoImmutableList<GoHxrtFeatureReason>;
	final capabilities:GoImmutableList<GoRuntimeCapabilitySelection>;
}

/**
	Builds the single typed runtime-capability manifest.

	Why
	Runtime dependencies, footprint-explicit files, full-copy compatibility, and
	report provenance previously lived in separate branches.

	What
	Normalizes manual and inferred feature evidence, expands dependencies, assigns
	every selected capability at least one reason, and resolves the exact files.

	How
	Selective mode includes only evidenced capabilities. Full-copy compatibility
	adds the historically broad capabilities with an explicit compatibility
	contract reason, while footprint-explicit capabilities still require typed,
	define, manual, or inference-disabled evidence.
**/
class GoRuntimeCapabilityManifest {
	public static inline final SCHEMA_VERSION = 1;
	public static inline final AUTHORITY = "typed_usage_plus_surface_plan_runtime_manifest";

	static final FULL_COPY_BASE_FEATURES:Array<String> = [
		GoHxrtFeatureAnalyzer.FEATURE_CORE,
		GoHxrtFeatureAnalyzer.FEATURE_STRING,
		GoHxrtFeatureAnalyzer.FEATURE_EQUALITY,
		GoHxrtFeatureAnalyzer.FEATURE_PRINT,
		GoHxrtFeatureAnalyzer.FEATURE_EXCEPTION,
		GoHxrtFeatureAnalyzer.FEATURE_JSON,
		GoHxrtFeatureAnalyzer.FEATURE_SYS,
		GoHxrtFeatureAnalyzer.FEATURE_FILE_IO,
		GoHxrtFeatureAnalyzer.FEATURE_FILESYSTEM,
		GoHxrtFeatureAnalyzer.FEATURE_PROCESS,
		GoHxrtFeatureAnalyzer.FEATURE_BYTES,
		GoHxrtFeatureAnalyzer.FEATURE_SSL,
		GoHxrtFeatureAnalyzer.FEATURE_THREAD,
		GoHxrtFeatureAnalyzer.FEATURE_ENUM_VALUE,
		GoHxrtFeatureAnalyzer.FEATURE_MAP_INT,
		GoHxrtFeatureAnalyzer.FEATURE_MAP_STRING,
		GoHxrtFeatureAnalyzer.FEATURE_MAP_OBJECT,
		GoHxrtFeatureAnalyzer.FEATURE_ATOMIC_INT,
		GoHxrtFeatureAnalyzer.FEATURE_ATOMIC_OBJECT
	];

	static final FULL_COPY_INFERENCE_DISABLED_FEATURES:Array<String> = [
		GoHxrtFeatureAnalyzer.FEATURE_ARRAY,
		GoHxrtFeatureAnalyzer.FEATURE_TERMINAL,
		GoHxrtFeatureAnalyzer.FEATURE_HTTP,
		GoHxrtFeatureAnalyzer.FEATURE_DATE,
		GoHxrtFeatureAnalyzer.FEATURE_MATH,
		GoHxrtFeatureAnalyzer.FEATURE_CRYPTO,
		GoHxrtFeatureAnalyzer.FEATURE_ZIP,
		GoHxrtFeatureAnalyzer.FEATURE_SOCKET,
		GoHxrtFeatureAnalyzer.FEATURE_SOCKET_SSL,
		GoHxrtFeatureAnalyzer.FEATURE_TEMPLATE,
		GoHxrtFeatureAnalyzer.FEATURE_REFLECTION,
		GoHxrtFeatureAnalyzer.FEATURE_REGEX,
		GoHxrtFeatureAnalyzer.FEATURE_SERIALIZATION
	];

	public static function build(buildContext:GoBuildContext, inferredFeatures:Array<String>,
			inferredReasons:Array<GoHxrtFeatureReason>):GoRuntimeCapabilityManifestSnapshot {
		var fullCopy = buildContext.hxrtForceFullCopy || !buildContext.isHxrtSelectiveEnabled();
		var manual = normalizeKnown(buildContext.hxrtManualFeatures);
		var inferred = buildContext.hxrtNoFeatureInfer ? [] : normalizeKnown(inferredFeatures);
		var selected = inferred.copy();
		var reasons = copyReasonsForSelected(inferredReasons, inferred);

		for (feature in manual) {
			addUnique(selected, feature);
			reasons.push({
				feature: feature,
				sourceKind: "manual_define",
				source: GoBuildContextResolver.HXRT_FEATURES_DEFINE
			});
		}

		if (buildContext.nativeStackTraceEnabled) {
			addUnique(selected, GoHxrtFeatureAnalyzer.FEATURE_STACK);
			reasons.push({
				feature: GoHxrtFeatureAnalyzer.FEATURE_STACK,
				sourceKind: "define",
				source: GoBuildContextResolver.NATIVE_STACK_TRACE_DEFINE
			});
		}

		var expanded = GoHxrtFeatureAnalyzer.expandWithReasons(selected, reasons);
		var finalFeatures = expanded.features.copy();
		var finalReasons = expanded.reasons.copy();

		if (fullCopy) {
			for (feature in FULL_COPY_BASE_FEATURES) {
				addUnique(finalFeatures, feature);
				finalReasons.push({
					feature: feature,
					sourceKind: "compatibility_contract",
					source: "default_full_copy"
				});
			}
			if (buildContext.hxrtNoFeatureInfer) {
				for (feature in FULL_COPY_INFERENCE_DISABLED_FEATURES) {
					addUnique(finalFeatures, feature);
					finalReasons.push({
						feature: feature,
						sourceKind: "compatibility_contract",
						source: "inference_disabled_full_copy"
					});
				}
			}
		}
		finalFeatures.sort(compareFeatureIds);
		finalReasons = sortedUniqueReasons(finalReasons);
		var capabilities = new Array<GoRuntimeCapabilitySelection>();
		var filesByName = new Map<String, Bool>();
		for (feature in finalFeatures) {
			var typedId = knownId(feature);
			if (typedId == null) {
				continue;
			}
			var featureFiles = GoHxrtFeatureAnalyzer.directFilesForFeature(feature);
			featureFiles.sort(compareStrings);
			for (fileName in featureFiles) {
				filesByName.set(fileName, true);
			}
			var featureReasons = [
				for (reason in finalReasons)
					if (reason.feature == feature) {
						feature: reason.feature,
						sourceKind: reason.sourceKind,
						source: reason.source
					}
			];
			capabilities.push({
				id: typedId,
				files: GoImmutableList.fromArray(featureFiles),
				reasons: GoImmutableList.fromArray(featureReasons)
			});
		}

		var files = [for (fileName in filesByName.keys()) fileName];
		files.sort(compareStrings);
		return {
			schemaVersion: SCHEMA_VERSION,
			authority: AUTHORITY,
			fullCopy: fullCopy,
			selectiveEnabled: buildContext.isHxrtSelectiveEnabled(),
			inferenceDisabled: buildContext.hxrtNoFeatureInfer,
			manualFeatures: typedList(manual),
			inferredFeatures: typedList(inferred),
			selectedFeatures: typedList(finalFeatures),
			files: GoImmutableList.fromArray(files),
			reasons: GoImmutableList.fromArray(finalReasons),
			capabilities: GoImmutableList.fromArray(capabilities)
		};
	}

	static function sortedUniqueReasons(reasons:Array<GoHxrtFeatureReason>):Array<GoHxrtFeatureReason> {
		var byKey = new Map<String, GoHxrtFeatureReason>();
		for (reason in reasons) {
			if (reason != null && GoHxrtFeatureAnalyzer.isKnownFeature(reason.feature)) {
				byKey.set(reason.feature + "\n" + reason.sourceKind + "\n" + reason.source, {
					feature: reason.feature,
					sourceKind: reason.sourceKind,
					source: reason.source
				});
			}
		}
		var out = [for (reason in byKey) reason];
		out.sort((a, b) -> {
			var featureOrder = compareFeatureIds(a.feature, b.feature);
			if (featureOrder != 0) {
				return featureOrder;
			}
			var kindOrder = compareStrings(a.sourceKind, b.sourceKind);
			return kindOrder != 0 ? kindOrder : compareStrings(a.source, b.source);
		});
		return out;
	}

	static function copyReasonsForSelected(reasons:Array<GoHxrtFeatureReason>, selected:Array<String>):Array<GoHxrtFeatureReason> {
		var out = new Array<GoHxrtFeatureReason>();
		if (reasons == null) {
			return out;
		}
		for (reason in reasons) {
			if (reason != null && selected.indexOf(reason.feature) >= 0) {
				out.push({
					feature: reason.feature,
					sourceKind: reason.sourceKind,
					source: reason.source
				});
			}
		}
		return out;
	}

	static function normalizeKnown(features:Array<String>):Array<String> {
		var out = new Array<String>();
		if (features != null) {
			for (feature in features) {
				if (feature != null && GoHxrtFeatureAnalyzer.isKnownFeature(feature)) {
					addUnique(out, feature);
				}
			}
		}
		out.sort(compareStrings);
		return out;
	}

	static function typedList(features:Array<String>):GoImmutableList<GoHxrtFeatureId> {
		var out = new Array<GoHxrtFeatureId>();
		for (feature in features) {
			var typedId = knownId(feature);
			if (typedId != null) {
				out.push(typedId);
			}
		}
		return GoImmutableList.fromArray(out);
	}

	static function knownId(feature:String):Null<GoHxrtFeatureId> {
		return switch (feature) {
			case GoHxrtFeatureAnalyzer.FEATURE_CORE: GoHxrtFeatureId.HxrtCore;
			case GoHxrtFeatureAnalyzer.FEATURE_ARRAY: GoHxrtFeatureId.HxrtArray;
			case GoHxrtFeatureAnalyzer.FEATURE_STRING: GoHxrtFeatureId.HxrtString;
			case GoHxrtFeatureAnalyzer.FEATURE_EQUALITY: GoHxrtFeatureId.HxrtEquality;
			case GoHxrtFeatureAnalyzer.FEATURE_PRINT: GoHxrtFeatureId.HxrtPrint;
			case GoHxrtFeatureAnalyzer.FEATURE_EXCEPTION: GoHxrtFeatureId.HxrtException;
			case GoHxrtFeatureAnalyzer.FEATURE_JSON: GoHxrtFeatureId.HxrtJson;
			case GoHxrtFeatureAnalyzer.FEATURE_SYS: GoHxrtFeatureId.HxrtSys;
			case GoHxrtFeatureAnalyzer.FEATURE_TERMINAL: GoHxrtFeatureId.HxrtTerminal;
			case GoHxrtFeatureAnalyzer.FEATURE_FILE_IO: GoHxrtFeatureId.HxrtFileIo;
			case GoHxrtFeatureAnalyzer.FEATURE_FILESYSTEM: GoHxrtFeatureId.HxrtFilesystem;
			case GoHxrtFeatureAnalyzer.FEATURE_PROCESS: GoHxrtFeatureId.HxrtProcess;
			case GoHxrtFeatureAnalyzer.FEATURE_SOCKET: GoHxrtFeatureId.HxrtSocket;
			case GoHxrtFeatureAnalyzer.FEATURE_HTTP: GoHxrtFeatureId.HxrtHttp;
			case GoHxrtFeatureAnalyzer.FEATURE_BYTES: GoHxrtFeatureId.HxrtBytes;
			case GoHxrtFeatureAnalyzer.FEATURE_DATE: GoHxrtFeatureId.HxrtDate;
			case GoHxrtFeatureAnalyzer.FEATURE_MATH: GoHxrtFeatureId.HxrtMath;
			case GoHxrtFeatureAnalyzer.FEATURE_CRYPTO: GoHxrtFeatureId.HxrtCrypto;
			case GoHxrtFeatureAnalyzer.FEATURE_ZIP: GoHxrtFeatureId.HxrtZip;
			case GoHxrtFeatureAnalyzer.FEATURE_SSL: GoHxrtFeatureId.HxrtSsl;
			case GoHxrtFeatureAnalyzer.FEATURE_SOCKET_SSL: GoHxrtFeatureId.HxrtSocketSsl;
			case GoHxrtFeatureAnalyzer.FEATURE_THREAD: GoHxrtFeatureId.HxrtThread;
			case GoHxrtFeatureAnalyzer.FEATURE_STACK: GoHxrtFeatureId.HxrtStack;
			case GoHxrtFeatureAnalyzer.FEATURE_TEMPLATE: GoHxrtFeatureId.HxrtTemplate;
			case GoHxrtFeatureAnalyzer.FEATURE_REFLECTION: GoHxrtFeatureId.HxrtReflection;
			case GoHxrtFeatureAnalyzer.FEATURE_REGEX: GoHxrtFeatureId.HxrtRegex;
			case GoHxrtFeatureAnalyzer.FEATURE_SERIALIZATION: GoHxrtFeatureId.HxrtSerialization;
			case GoHxrtFeatureAnalyzer.FEATURE_ENUM_VALUE: GoHxrtFeatureId.HxrtEnumValue;
			case GoHxrtFeatureAnalyzer.FEATURE_MAP_INT: GoHxrtFeatureId.HxrtMapInt;
			case GoHxrtFeatureAnalyzer.FEATURE_MAP_STRING: GoHxrtFeatureId.HxrtMapString;
			case GoHxrtFeatureAnalyzer.FEATURE_MAP_OBJECT: GoHxrtFeatureId.HxrtMapObject;
			case GoHxrtFeatureAnalyzer.FEATURE_ATOMIC_INT: GoHxrtFeatureId.HxrtAtomicInt;
			case GoHxrtFeatureAnalyzer.FEATURE_ATOMIC_OBJECT: GoHxrtFeatureId.HxrtAtomicObject;
			case _: null;
		};
	}

	static function addUnique(values:Array<String>, value:String):Void {
		if (values.indexOf(value) < 0) {
			values.push(value);
		}
	}

	static function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function compareFeatureIds(a:String, b:String):Int {
		var order = GoHxrtFeatureAnalyzer.knownFeatures();
		var ai = order.indexOf(a);
		var bi = order.indexOf(b);
		return ai == bi ? compareStrings(a, b) : ai - bi;
	}
}
#end
