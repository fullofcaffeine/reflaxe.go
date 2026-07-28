package reflaxe.go.compiler;

#if (macro || reflaxe_runtime)
import haxe.Exception;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureId;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLedgerSnapshot;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLevelId;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageTargetKind;

/**
	Stable identities for portable surfaces that may eventually admit native Go
	representations.

	Why
	A free-form module name could turn a namespace convention into semantic
	authority. Known identities must be reviewed before a contract can use them.

	What
	Names the source families owned by the Option/Result, collection, string,
	bytes, iterator, and closure admission Beads.

	How
	`validate(...)` rejects values outside this closed vocabulary. Merely being
	known does not admit a surface; production admission still requires a
	validated contract in `defaultRegistry()`.
**/
enum abstract GoSurfaceId(String) to String {
	var HaxeArray = "haxe.Array";
	var HaxeString = "haxe.String";
	var HaxeBytes = "haxe.io.Bytes";
	var HaxeStringMap = "haxe.ds.StringMap";
	var HaxeIntMap = "haxe.ds.IntMap";
	var HaxeObjectMap = "haxe.ds.ObjectMap";
	var PortableOption = "reflaxe.std.Option";
	var PortableResult = "reflaxe.std.Result";
	var HaxeIterator = "haxe.Iterator";
	var HaxeFunction = "haxe.Function";

	public function isKnown():Bool {
		return switch (this) {
			case HaxeArray | HaxeString | HaxeBytes | HaxeStringMap | HaxeIntMap | HaxeObjectMap | PortableOption | PortableResult | HaxeIterator |
				HaxeFunction:
				true;
			case _:
				false;
		};
	}
}

/**
	The source-level authority whose semantics a representation preserves.

	Why / What / How
	- Ordinary portable Haxe and a shared portable family facade are distinct
	  contracts even when both can use a native Go carrier.
	- Explicit `go.*` APIs remain native source boundaries and do not enter this
	  portable-specialization registry.
**/
enum abstract GoSourceContractKind(String) to String {
	var PortableHaxe = "portable_haxe";
	var PortableFamilyFacade = "portable_family_facade";

	public function isKnown():Bool {
		return switch (this) {
			case PortableHaxe | PortableFamilyFacade: true;
			case _: false;
		};
	}
}

/**
	Closed native Go carriers available to admitted portable contracts.

	Why / What / How
	- Reports and planners need stable values rather than target-shape strings.
	- `GoSlice` and `GoMap` name native-storage-backed semantic carriers. They
	  do not by themselves authorize naked `[]T` or `map[K]V` values; the
	  surface contract supplies the required identity, null, and iteration
	  behavior.
	- Option/Result currently select their typed carrier identities; later
	  admission Beads may enable the remaining values after their proofs land.
	- Catalog selection records representation authority, while planner
	  consumption remains a separate step.
**/
enum abstract GoNativeRepresentation(String) to String {
	var GoSlice = "go_slice";
	var GoMap = "go_map";
	var GoString = "go_string";
	var GoBytes = "go_byte_slice";
	var GoIterator = "go_iterator";
	var GoClosure = "go_closure";
	var GoOption = "go_option";
	var GoResult = "go_result";

	public function isKnown():Bool {
		return switch (this) {
			case GoSlice | GoMap | GoString | GoBytes | GoIterator | GoClosure | GoOption | GoResult: true;
			case _: false;
		};
	}
}

/**
	Closed fallback carriers that preserve source semantics when admission fails.

	Why / What / How
	- Every admitted contract must name a fallback so an optimizer cannot silently
	  invent weaker behavior.
	- The value is reported even when a particular observed shape is rejected.
**/
enum abstract GoSurfaceFallbackRepresentation(String) to String {
	var HxrtArray = "hxrt_array";
	var HxrtMap = "hxrt_map";
	var HxrtString = "hxrt_string";
	var HxrtBytes = "hxrt_bytes";
	var HxrtIterator = "hxrt_iterator";
	var HxrtClosure = "hxrt_closure";
	var PortableOption = "portable_option";
	var PortableResult = "portable_result";

	public function isKnown():Bool {
		return switch (this) {
			case HxrtArray | HxrtMap | HxrtString | HxrtBytes | HxrtIterator | HxrtClosure | PortableOption | PortableResult: true;
			case _:
				false;
		};
	}
}

/**
	What the compiler must do when the native representation is not eligible.

	Why / What / How
	- `always_available` selects the named portable fallback.
	- `reasoned_runtime_requirement` selects it and records its runtime cost.
	- `error` forbids fallback for that source contract.
**/
enum abstract GoSurfaceFallbackPolicy(String) to String {
	var AlwaysAvailable = "always_available";
	var ReasonedRuntimeRequirement = "reasoned_runtime_requirement";
	var Error = "error";

	public function isKnown():Bool {
		return switch (this) {
			case AlwaysAvailable | ReasonedRuntimeRequirement | Error: true;
			case _: false;
		};
	}
}

/**
	Whether the admitted native carrier can participate in an `hxrt`-free build.

	Why / What / How
	- The value records a reviewed contract fact, not a guess from an empty import
	  list.
	- Runtime planning remains a later consumer and may still find unrelated
	  program requirements.
**/
enum abstract GoNoHxrtStatus(String) to String {
	var Eligible = "eligible";
	var Conditional = "conditional";
	var Ineligible = "ineligible";

	public function isKnown():Bool {
		return switch (this) {
			case Eligible | Conditional | Ineligible: true;
			case _: false;
		};
	}
}

/**
	How a source-semantic contract relates to sibling Reflaxe targets.

	Why / What / How
	- `target_local` permits a Go-only source contract.
	- `shared_contract_required` requires a stable family ID and version while
	  leaving Go representation, imports, and runtime consequences target-local.
**/
enum abstract GoFamilySyncExpectation(String) to String {
	var TargetLocal = "target_local";
	var SharedContractRequired = "shared_contract_required";

	public function isKnown():Bool {
		return switch (this) {
			case TargetLocal | SharedContractRequired: true;
			case _: false;
		};
	}
}

/**
	Closed proof kinds accepted by registry governance.

	Why / What / How
	- Portable admission requires a semantic-diff fixture.
	- Generated-shape, Go-runtime, race, and performance fixtures can add
	  target-specific evidence without replacing the source-semantic proof.
**/
enum abstract GoSurfaceProofKind(String) to String {
	var SemanticDiff = "semantic_diff";
	var GeneratedShape = "generated_shape";
	var GoRuntime = "go_runtime";
	var Race = "race";
	var Performance = "performance";

	public function isKnown():Bool {
		return switch (this) {
			case SemanticDiff | GeneratedShape | GoRuntime | Race | Performance: true;
			case _: false;
		};
	}
}

/**
	Closed declaration kinds that may form a nominal admission pattern.

	Why / What / How
	- Function and anonymous shapes have dedicated typed variants.
	- Keeping this vocabulary separate prevents their ledger labels from being
	  smuggled into a nominal contract and keeps JSON Schema synchronization exact.
**/
enum abstract GoSurfaceNominalKind(String) to String {
	var Class = "class";
	var Enum = "enum";
	var Typedef = "typedef";
	var Abstract = "abstract";

	public function isKnown():Bool {
		return switch (this) {
			case Class | Enum | Typedef | Abstract: true;
			case _: false;
		};
	}
}

/**
	A recursive, macro-object-free pattern matched against `GoTypeShape`.

	Why
	Module membership alone cannot prove generic, function, or nested shape
	eligibility.

	What
	Matches exact nominal roots, function signatures, anonymous field protocols,
	and named bindings. A binding captures a complete observed child shape for
	later eligibility rules.

	How
	The registry requires an exact known root. `Bind` is permitted only below
	that root, so a contract cannot admit arbitrary typed usage.
**/
enum GoSurfaceTypePattern {
	NominalPattern(kind:GoSurfaceNominalKind, path:String, parameters:GoImmutableList<GoSurfaceTypePattern>);
	FunctionPattern(arguments:GoImmutableList<GoSurfaceFunctionArgumentPattern>, returnType:GoSurfaceTypePattern);
	AnonymousPattern(fields:GoImmutableList<GoSurfaceAnonymousFieldPattern>);
	Bind(name:String);
}

/**
	Why / What / How
	- Function optionality is observable Haxe source behavior.
	- A function pattern therefore retains the ordered optional flag beside each
	  recursive argument pattern.
**/
typedef GoSurfaceFunctionArgumentPattern = {
	final optional:Bool;
	final shape:GoSurfaceTypePattern;
}

/**
	Why / What / How
	- Haxe `Iterator<T>` is a structural protocol, so it has no nominal root.
	- An anonymous pattern retains each exact field name, optional flag, and
	  recursive shape instead of treating every anonymous object as an iterator.
**/
typedef GoSurfaceAnonymousFieldPattern = {
	final name:String;
	final optional:Bool;
	final shape:GoSurfaceTypePattern;
}

/**
	Closed post-match eligibility checks.

	Why / What / How
	- Exact root matching is necessary but not sufficient for nested `Dynamic`,
	  unresolved shapes, a bound Go-comparable map key, or a nominal map whose
	  fixed key is encoded by the typed surface.
	- Rules consume bindings created by the pattern; validation rejects references
	  to bindings the pattern did not declare.
**/
enum GoSurfaceEligibilityRule {
	NoUnknownShapes;
	ShapeContainsNoDynamic;
	BindingContainsNoDynamic(name:String);
	BindingIsGoComparable(name:String);
	BindingHasProvenCollectionCarrier(name:String);
	SurfaceHasFixedGoComparableMapKey;
}

/**
	One relative, stable proof fixture reference.

	Why / What / How
	- Proof IDs remain stable in reports while fixture paths let CI verify the
	  actual evidence.
	- Validation rejects absolute paths and traversal; repository tests own file
	  existence because packaged compiler code must not depend on checkout paths.
**/
typedef GoSurfaceProof = {
	final proofId:String;
	final kind:GoSurfaceProofKind;
	final fixturePath:String;
}

/**
	One Go import consequence of the selected representation.

	Why / What / How
	- Imports are typed contract data rather than incidental printer discovery.
	- An empty immutable list explicitly means the representation needs no Go
	  package import.
**/
typedef GoSurfaceImportRequirement = {
	final path:String;
	final reason:String;
}

/**
	The complete compiler-owned admission contract for one portable surface.

	Why
	Native representation changes can preserve output shape while breaking Haxe
	identity, mutation, nil, iteration, Unicode, callback, or error behavior.

	What
	Records versioned source semantics, an exact typed shape pattern, additional
	eligibility rules, native and fallback carriers, import/runtime consequences,
	no-`hxrt` status, proof fixtures, and family synchronization expectations.

	How
	Only `GoSurfaceContractRegistry.create(...)` can turn these records into
	authority. It validates and deep-copies them before any decision is exposed.
**/
typedef GoSurfaceContract = {
	final surfaceId:GoSurfaceId;
	final contractVersion:Int;
	final sourceContract:GoSourceContractKind;
	final sourceSemanticsId:String;
	final sourceSemanticsVersion:Int;
	final sourceSemantics:String;
	final eligibleShape:GoSurfaceTypePattern;
	final eligibilityRules:GoImmutableList<GoSurfaceEligibilityRule>;
	final nativeRepresentation:GoNativeRepresentation;
	final nativeImports:GoImmutableList<GoSurfaceImportRequirement>;
	final nativeRuntimeRequirements:GoImmutableList<GoHxrtFeatureId>;
	final fallbackRepresentation:GoSurfaceFallbackRepresentation;
	final fallbackPolicy:GoSurfaceFallbackPolicy;
	final fallbackImports:GoImmutableList<GoSurfaceImportRequirement>;
	final fallbackRuntimeRequirements:GoImmutableList<GoHxrtFeatureId>;
	final noHxrtStatus:GoNoHxrtStatus;
	final proofs:GoImmutableList<GoSurfaceProof>;
	final familySyncExpectation:GoFamilySyncExpectation;
	final familyContractId:String;
	final familyContractVersion:Int;
}

/**
	Stable validation failures for malformed registry entries.

	Why / What / How
	- Catalog mistakes fail closed before planning.
	- Typed codes make negative tests and CI governance independent from prose.
**/
enum abstract GoSurfaceValidationCode(String) to String {
	var DuplicateSurface = "duplicate_surface";
	var UnknownSurface = "unknown_surface";
	var InvalidVersion = "invalid_version";
	var UnknownSourceContract = "unknown_source_contract";
	var InvalidSourceSemantics = "invalid_source_semantics";
	var InvalidShape = "invalid_shape";
	var UnknownEligibilityRule = "unknown_eligibility_rule";
	var UnboundEligibilityRule = "unbound_eligibility_rule";
	var UnknownNativeRepresentation = "unknown_native_representation";
	var UnknownFallbackRepresentation = "unknown_fallback_representation";
	var UnknownFallbackPolicy = "unknown_fallback_policy";
	var UnknownNoHxrtStatus = "unknown_no_hxrt_status";
	var MissingProof = "missing_proof";
	var MissingSemanticProof = "missing_semantic_proof";
	var DuplicateProof = "duplicate_proof";
	var UnknownProofKind = "unknown_proof_kind";
	var UnsafeProofPath = "unsafe_proof_path";
	var UnknownRuntimeRequirement = "unknown_runtime_requirement";
	var InvalidImportRequirement = "invalid_import_requirement";
	var InvalidNoHxrtContract = "invalid_no_hxrt_contract";
	var InvalidFamilySync = "invalid_family_sync";
	var MalformedContract = "malformed_contract";
}

/** One deterministic registry validation failure. */
typedef GoSurfaceValidationIssue = {
	final code:GoSurfaceValidationCode;
	final surfaceId:String;
	final detail:String;
}

/** Stable result of evaluating one compiler-observed known surface shape. */
enum abstract GoSurfaceDecisionOutcome(String) to String {
	var Admitted = "admitted";
	var Rejected = "rejected";
}

/** Stable reason for an admission or rejection. */
enum abstract GoSurfaceDecisionReason(String) to String {
	var ContractAdmitted = "contract_admitted";
	var ContractMissing = "contract_missing";
	var ShapeMismatch = "shape_mismatch";
	var EligibilityRejected = "eligibility_rejected";
}

/**
	One immutable and explainable registry decision.

	Why / What / How
	- Later planners can consume this record without rerunning admission.
	- A rejected entry retains the known fallback contract when one exists.
**/
typedef GoSurfaceDecision = {
	final module:String;
	final location:String;
	final usageLevel:GoTypeUsageLevelId;
	final surfaceId:GoSurfaceId;
	final shape:GoTypeShape;
	final outcome:GoSurfaceDecisionOutcome;
	final reason:GoSurfaceDecisionReason;
	final detail:String;
	final contractVersion:Int;
	final selectedRepresentation:Null<GoNativeRepresentation>;
	final nativeImports:GoImmutableList<GoSurfaceImportRequirement>;
	final runtimeRequirements:GoImmutableList<GoHxrtFeatureId>;
	final fallbackRepresentation:Null<GoSurfaceFallbackRepresentation>;
	final fallbackPolicy:Null<GoSurfaceFallbackPolicy>;
	final fallbackImports:GoImmutableList<GoSurfaceImportRequirement>;
	final fallbackRuntimeRequirements:GoImmutableList<GoHxrtFeatureId>;
	final noHxrtStatus:Null<GoNoHxrtStatus>;
	final proofIds:GoImmutableList<String>;
	final familySyncExpectation:Null<GoFamilySyncExpectation>;
	final familyContractId:String;
	final familyContractVersion:Int;
}

/**
	The deeply read-only registry authority published on `CompilationContext`.

	Why / What / How
	- The source ledger and validated catalog are evaluated once before lowering.
	- Reports and future planners consume the same sorted snapshot.
**/
typedef GoSurfaceContractRegistrySnapshot = {
	final schemaVersion:Int;
	final registryVersion:Int;
	final authority:String;
	final profileAdmission:String;
	final catalogCount:Int;
	final decisionCount:Int;
	final admittedCount:Int;
	final rejectedCount:Int;
	final contracts:GoImmutableList<GoSurfaceContract>;
	final decisions:GoImmutableList<GoSurfaceDecision>;
}

/**
	A typed exception raised only when invalid catalog data is promoted to
	compiler authority.

	Why / What / How
	- `validate(...)` supports negative governance tests without throwing.
	- `create(...)` refuses to construct a usable registry when any issue exists.
**/
class GoSurfaceContractRegistryException extends Exception {
	public final issues:GoImmutableList<GoSurfaceValidationIssue>;

	public function new(issues:GoImmutableList<GoSurfaceValidationIssue>) {
		this.issues = issues;
		var first = issues.length == 0 ? "unknown validation failure" : issues.at(0).detail;
		super("Invalid Go surface contract registry: " + first);
	}
}

private typedef PatternMatch = {
	final matched:Bool;
	final bindings:Map<String, GoTypeShape>;
}

/**
	The single fail-closed authority for portable native-representation admission.

	Why
	Profile selection, module prefixes, and source-text scans are not semantic
	proof. A representation must combine compiler-observed typed usage with a
	versioned validated contract.

	What
	Validates a catalog, matches exact typed shapes, evaluates closed eligibility
	rules, and emits deterministic machine/human reports for every known observed
	surface.

	How
	`defaultRegistry()` contains only surface-family entries whose dependent
	admission Beads have landed. Unknown observed types are ignored; known
	observed surfaces without a contract are explicitly rejected as
	`contract_missing`.
**/
class GoSurfaceContractRegistry {
	public static inline final SCHEMA_VERSION = 2;
	public static inline final REGISTRY_VERSION = 1;
	public static inline final AUTHORITY = "typed_usage_plus_versioned_surface_contract";
	public static inline final PROFILE_ADMISSION = "forbidden";

	final contracts:GoImmutableList<GoSurfaceContract>;
	final contractsById:Map<String, GoSurfaceContract>;

	private function new(contracts:Array<GoSurfaceContract>) {
		var copied = [for (contract in contracts) copyContract(contract)];
		copied.sort(compareContracts);
		this.contracts = GoImmutableList.fromArray(copied);
		this.contractsById = [];
		for (contract in copied) {
			contractsById.set(contract.surfaceId, contract);
		}
	}

	/**
		Create validated registry authority.

		Why / What / How
		- Validation is mandatory, not a report-only warning.
		- The input array and every nested collection are copied before publication.
	**/
	public static function create(contracts:Array<GoSurfaceContract>):GoSurfaceContractRegistry {
		var source = contracts == null ? [] : contracts;
		var issues = validate(contracts);
		if (issues.length > 0) {
			throw new GoSurfaceContractRegistryException(issues);
		}
		try {
			return new GoSurfaceContractRegistry(source);
		} catch (_:Exception) {
			throw new GoSurfaceContractRegistryException(GoImmutableList.fromArray([
				{
					code: GoSurfaceValidationCode.MalformedContract,
					surfaceId: "",
					detail: "Catalog data changed or remained malformed after validation."
				}
			]));
		}
	}

	/**
		Production catalog.

		Why / What / How
		- Array, String, Bytes, StringMap, IntMap, Iterator, Option, and Result
		  have exact portable contracts and semantic evidence.
		- ObjectMap and functions remain absent until their identity proofs land.
		- Catalog membership is profile-independent and does not itself change
		  lowering; `.7.6` owns planner consumption.
	**/
	public static function defaultRegistry():GoSurfaceContractRegistry {
		return create([
			haxeArrayContract(),
			haxeStringContract(),
			haxeBytesContract(),
			haxeStringMapContract(),
			haxeIntMapContract(),
			haxeIteratorContract(),
			portableOptionContract(),
			portableResultContract()
		]);
	}

	/**
		The portable shared Haxe `Array<T>` contract.

		Why
		A naked Go `[]T` copies its slice header when assigned, so a later append
		can change one alias's length without changing the other. Primitive slices
		also cannot represent sparse Haxe holes without extra presence state.

		What
		Admits recursively typed element shapes to `GoSlice`, meaning a shared,
		slice-backed semantic carrier. It does not authorize replacing Haxe
		`Array<T>` with a naked Go slice.

		How
		The carrier must preserve object identity, alias-visible mutation,
		sparse/null slots, distinct null-versus-empty values, nested collection
		identity, iteration, and callback-visible mutation. Dynamic or unresolved
		nested shapes, named type parameters, and opaque typedef/abstract storage
		retain the existing `hxrt.Array` fallback.
	**/
	static function haxeArrayContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeArray,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.array.shared",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Ordered mutable Array object with shared identity, alias-visible length/content mutation, sparse null slots, distinct null and empty values, nested identity, iteration, and callback-visible mutation.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "Array",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("element")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("element"),
				GoSurfaceEligibilityRule.BindingHasProvenCollectionCarrier("element")
			]),
			nativeRepresentation: GoNativeRepresentation.GoSlice,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtArray,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtArray]),
			noHxrtStatus: GoNoHxrtStatus.Conditional,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-collections-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_collections_contract"
				},
				{
					proofId: "array-shared-identity-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/array_identity_contract"
				},
				{
					proofId: "portable-collections-generated-fallback-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/fixtures/surface_contract_registry"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable Haxe `String` contract.

		Why
		A naked Go `string` cannot represent Haxe null, and byte indexing would
		change non-ASCII length, character, and slicing behavior.

		What
		`GoString` means the compiler's nullable pointer-backed Go string carrier,
		not a non-null raw Go value. The carrier preserves value equality,
		null-aware concatenation/stringification, and haxe.go's Unicode-scalar
		indexing contract.

		How
		String literals and operations continue to use the reviewed `hxrt` string
		boundary. `RawNative` UTF-16LE is a Bytes encoding policy and never grants
		byte offsets authority over String operations.
	**/
	static function haxeStringContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeString,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.string.nullable-unicode-scalar",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Nullable immutable string value with value equality, null-aware concatenation and stringification, and Unicode-scalar length, character, code-point, split, and slice operations.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "String", GoImmutableList.fromArray([])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.ShapeContainsNoDynamic
			]),
			nativeRepresentation: GoNativeRepresentation.GoString,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtString]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtString,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtString]),
			noHxrtStatus: GoNoHxrtStatus.Ineligible,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-string-bytes-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_string_bytes_contract"
				},
				{
					proofId: "unicode-string-source-owned-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/unicode_string_source_owned"
				},
				{
					proofId: "null-string-concat-generated-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/snapshot/core/string_concat_null_semantics"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable `haxe.io.Bytes` contract.

		Why
		A naked `[]byte` cannot preserve the public integer-valued `BytesData`
		alias, object identity, cached native views, and mutation visibility in
		both directions.

		What
		`GoBytes` means the shared staged Bytes object: authoritative aliased
		`[]int` data plus an opaque cached native `ByteView`. It does not authorize
		replacing the source value with a naked Go byte slice.

		How
		Mutations invalidate or validate the view as needed; `ofData`/`getData`
		remain bidirectional aliases, `sub` copies, overlapping `blit` is safe,
		values stay in 0...255, and UTF-8/RawNative policies remain explicit.
	**/
	static function haxeBytesContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeBytes,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.io.bytes.shared-data-view",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Shared mutable Bytes object with aliased integer BytesData, cache-coherent native byte views, 0...255 values, overlap-safe blit, copying subranges, distinct null and empty values, and explicit UTF-8/RawNative encoding.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "haxe.io.Bytes", GoImmutableList.fromArray([])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.ShapeContainsNoDynamic
			]),
			nativeRepresentation: GoNativeRepresentation.GoBytes,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtBytes]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtBytes,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtBytes]),
			noHxrtStatus: GoNoHxrtStatus.Ineligible,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-string-bytes-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_string_bytes_contract"
				},
				{
					proofId: "bytes-normalization-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/bytes_normalization_contract"
				},
				{
					proofId: "bytes-staged-carrier-generated-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/snapshot/stdlib/bytes_basic"
				},
				{
					proofId: "raw-native-utf16-generated-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/snapshot/core/raw_native_utf16_mode"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The structural Haxe `Iterator<T>` contract.

		Why
		An iterator is shared mutable progress, not a repeatable collection. A Go
		`range` rewrite can evaluate the source again, snapshot changing data, or
		give aliases independent cursors.

		What
		`GoIterator` is the exact non-optional zero-argument `hasNext():Bool` and
		`next():T` protocol whose two closures share one state owner. It is not a
		nominal type, a Go `range`, or a promise about `next()` after exhaustion.

		How
		Only recursively proven `T` shapes admit. Dynamic, unresolved, named
		generic, and opaque typedef/abstract element shapes retain the structural
		`hxrt` fallback, preserving order, live mutation, and alias-shared
		exhaustion.
	**/
	static function haxeIteratorContract():GoSurfaceContract {
		final noArguments = GoImmutableList.fromArray([]);
		return {
			surfaceId: GoSurfaceId.HaxeIterator,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.iterator.shared-cursor",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Exact hasNext/next structural iterator with single source evaluation, ordered delivery, repeated hasNext stability, live source visibility, and alias-shared cursor/exhaustion state.",
			eligibleShape: GoSurfaceTypePattern.AnonymousPattern(GoImmutableList.fromArray([
				{
					name: "hasNext",
					optional: false,
					shape: GoSurfaceTypePattern.FunctionPattern(noArguments,
						GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Abstract, "StdTypes.Bool", GoImmutableList.fromArray([])))
				},
				{
					name: "next",
					optional: false,
					shape: GoSurfaceTypePattern.FunctionPattern(noArguments, GoSurfaceTypePattern.Bind("element"))
				}
			])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("element"),
				GoSurfaceEligibilityRule.BindingHasProvenCollectionCarrier("element")
			]),
			nativeRepresentation: GoNativeRepresentation.GoIterator,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtIterator,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtCore]),
			noHxrtStatus: GoNoHxrtStatus.Eligible,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-iterator-closure-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_iterator_closure_contract"
				},
				{
					proofId: "structural-iterator-assignment-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/structural_iterator_assignment_contract"
				},
				{
					proofId: "structural-iterator-generated-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/snapshot/core/structural_iterator_assignment"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable fixed-string-key `haxe.ds.StringMap<V>` contract.

		Why / What / How
		- The nominal surface fixes the key to Haxe `String`; a typed rule proves
		  that its equality can use a Go-comparable string key without coercion.
		- A semantic carrier around `map[string]V` must retain shared mutation,
		  present-null versus missing, copy, iteration, nested values, and
		  callback-visible updates.
		- Dynamic, unresolved, or opaque alias value shapes retain the staged
		  StringMap plus `hxrt` fallback. Explicit `go.Map<K,V>` is not this
		  contract.
	**/
	static function haxeStringMapContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeStringMap,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.ds.string-map",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Shared mutable String-key map with exact string equality, present-null distinct from missing through exists(), shallow copy, complete iteration, nested identity, and callback-visible mutation.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.StringMap",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("value")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.SurfaceHasFixedGoComparableMapKey,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("value"),
				GoSurfaceEligibilityRule.BindingHasProvenCollectionCarrier("value")
			]),
			nativeRepresentation: GoNativeRepresentation.GoMap,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtMap,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtMapString]),
			noHxrtStatus: GoNoHxrtStatus.Conditional,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-collections-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_collections_contract"
				},
				{
					proofId: "portable-map-family-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/ds_maps_list_contract"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable fixed-integer-key `haxe.ds.IntMap<V>` contract.

		Why / What / How
		- The nominal surface fixes the key to Haxe `Int`; its equality is the
		  compiler's typed, Go-comparable integer equality rather than a string
		  conversion.
		- A semantic carrier around `map[int]V` retains shared mutation,
		  present-null versus missing, copy, iteration, nested values, and
		  callback-visible updates.
		- Dynamic, unresolved, or opaque alias value shapes retain the staged
		  IntMap plus its exact `hxrt` fallback.
	**/
	static function haxeIntMapContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeIntMap,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.ds.int-map",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Shared mutable Int-key map with exact integer equality, present-null distinct from missing through exists(), shallow copy, complete iteration, nested identity, and callback-visible mutation.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.IntMap",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("value")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.SurfaceHasFixedGoComparableMapKey,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("value"),
				GoSurfaceEligibilityRule.BindingHasProvenCollectionCarrier("value")
			]),
			nativeRepresentation: GoNativeRepresentation.GoMap,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtMap,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtMapInt]),
			noHxrtStatus: GoNoHxrtStatus.Conditional,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-collections-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_collections_contract"
				},
				{
					proofId: "portable-map-family-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/ds_maps_list_contract"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable `Some(value) | None` source contract.

		Why / What / How
		- Explicit presence means `Some(null)` is observably different from `None`.
		- Fully typed generic parameters can use a typed tagged Go carrier.
		- `Dynamic` or unresolved nested shapes retain the ordinary portable enum
		  fallback. This is not a conversion to nullable Go data.
	**/
	static function portableOptionContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.PortableOption,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableFamilyFacade,
			sourceSemanticsId: "reflaxe.std.option",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Tagged Some(value) or None; explicit presence preserves Some(null) as distinct from None.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Enum, "reflaxe.std.Option",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("value")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("value")
			]),
			nativeRepresentation: GoNativeRepresentation.GoOption,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.PortableOption,
			fallbackPolicy: GoSurfaceFallbackPolicy.AlwaysAvailable,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([]),
			noHxrtStatus: GoNoHxrtStatus.Eligible,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-option-result-typed-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_option_result_contract"
				},
				{
					proofId: "portable-option-result-fallback-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_option_result_fallback_contract"
				},
				{
					proofId: "portable-option-result-generated-fallback-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/fixtures/surface_contract_registry"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/**
		The portable `Ok(value) | Err(error)` source contract.

		Why / What / How
		- `T` and `E` are independent and must survive representation selection.
		- `GoResult` means a typed two-parameter tagged carrier; it never means
		  native `go.Result<T>` or Go `(T, error)`.
		- `Dynamic` or unresolved nested shapes retain the ordinary portable enum
		  fallback instead of erasing the error parameter.
	**/
	static function portableResultContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.PortableResult,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableFamilyFacade,
			sourceSemanticsId: "reflaxe.std.result",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Tagged Ok(value) or Err(error) with independent T and E parameters and no implicit Go-error conversion.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Enum, "reflaxe.std.Result",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("value"), GoSurfaceTypePattern.Bind("error")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("value"),
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("error")
			]),
			nativeRepresentation: GoNativeRepresentation.GoResult,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.PortableResult,
			fallbackPolicy: GoSurfaceFallbackPolicy.AlwaysAvailable,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([]),
			noHxrtStatus: GoNoHxrtStatus.Eligible,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "portable-option-result-typed-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_option_result_contract"
				},
				{
					proofId: "portable-option-result-fallback-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/portable_option_result_fallback_contract"
				},
				{
					proofId: "portable-option-result-generated-fallback-shape",
					kind: GoSurfaceProofKind.GeneratedShape,
					fixturePath: "test/fixtures/surface_contract_registry"
				}
			]),
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	/** Validate catalog data without promoting it to compiler authority. */
	public static function validate(contracts:Array<GoSurfaceContract>):GoImmutableList<GoSurfaceValidationIssue> {
		var issues = new Array<GoSurfaceValidationIssue>();
		var seenSurfaces = new Map<String, Bool>();
		if (contracts == null) {
			addIssue(issues, GoSurfaceValidationCode.MalformedContract, "", "The catalog must be a non-null array.");
		} else {
			for (contract in contracts) {
				if (contract == null) {
					addIssue(issues, GoSurfaceValidationCode.UnknownSurface, "", "Null surface contract entries are not allowed.");
					continue;
				}
				try {
					validateContract(contract, seenSurfaces, issues);
				} catch (_:Exception) {
					addIssue(issues, GoSurfaceValidationCode.MalformedContract, contract.surfaceId, "Malformed nested catalog data could not be validated.");
				}
			}
		}
		issues.sort(compareValidationIssues);
		return GoImmutableList.fromArray(issues);
	}

	static function validateContract(contract:GoSurfaceContract, seenSurfaces:Map<String, Bool>, issues:Array<GoSurfaceValidationIssue>):Void {
		var surfaceId:String = contract.surfaceId;
		if (seenSurfaces.exists(surfaceId)) {
			addIssue(issues, GoSurfaceValidationCode.DuplicateSurface, surfaceId, 'Surface "$surfaceId" appears more than once.');
		} else {
			seenSurfaces.set(surfaceId, true);
		}
		if (!contract.surfaceId.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownSurface, surfaceId, 'Surface "$surfaceId" is outside the closed registry vocabulary.');
		}
		if (contract.contractVersion < 1 || contract.sourceSemanticsVersion < 1) {
			addIssue(issues, GoSurfaceValidationCode.InvalidVersion, surfaceId, "Contract and source-semantics versions must both be positive.");
		}
		if (!contract.sourceContract.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownSourceContract, surfaceId, "Source contract kind is not recognized.");
		}
		if (isBlank(contract.sourceSemanticsId) || isBlank(contract.sourceSemantics)) {
			addIssue(issues, GoSurfaceValidationCode.InvalidSourceSemantics, surfaceId, "Source semantics require a stable ID and non-empty description.");
		}
		var patternIsValid = isValidPattern(contract.eligibleShape);
		if (!patternIsValid || !patternRootMatchesSurface(contract.surfaceId, contract.eligibleShape)) {
			addIssue(issues, GoSurfaceValidationCode.InvalidShape, surfaceId, "Eligible shape must use a valid recursive pattern with the exact known root.");
		}
		if (patternIsValid) {
			validateEligibilityRules(contract, issues);
		}
		if (!contract.nativeRepresentation.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownNativeRepresentation, surfaceId, "Native representation is not recognized.");
		}
		if (!contract.fallbackRepresentation.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownFallbackRepresentation, surfaceId, "Fallback representation is not recognized.");
		}
		if (!contract.fallbackPolicy.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownFallbackPolicy, surfaceId, "Fallback policy is not recognized.");
		}
		if (!contract.noHxrtStatus.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.UnknownNoHxrtStatus, surfaceId, "No-hxrt status is not recognized.");
		}
		validateImports(surfaceId, contract.nativeImports, issues);
		validateImports(surfaceId, contract.fallbackImports, issues);
		validateRuntimeRequirements(surfaceId, contract.nativeRuntimeRequirements, issues);
		validateRuntimeRequirements(surfaceId, contract.fallbackRuntimeRequirements, issues);
		if (contract.noHxrtStatus == GoNoHxrtStatus.Eligible
			&& contract.nativeRuntimeRequirements != null
			&& contract.nativeRuntimeRequirements.length > 0) {
			addIssue(issues, GoSurfaceValidationCode.InvalidNoHxrtContract, surfaceId,
				"A no-hxrt eligible native representation cannot require hxrt features.");
		}
		validateProofs(contract, issues);
		validateFamilySync(contract, issues);
	}

	/**
		Evaluate compiler-observed typed usage once.

		Why / What / How
		- Candidate classification uses only exact `GoTypeShape` roots.
		- Profile/preset state is not an input.
		- Known surfaces without a contract reject explicitly; unrelated shapes do
		  not become registry noise.
	**/
	public function snapshot(ledger:GoTypeUsageLedgerSnapshot):GoSurfaceContractRegistrySnapshot {
		var decisionsByKey = new Map<String, GoSurfaceDecision>();
		if (ledger != null) {
			for (module in ledger.modules) {
				for (usage in module.typeUsages) {
					var surfaceId = surfaceForShape(usage.shape);
					if (surfaceId == null) {
						continue;
					}
					var decision = evaluate(module.module, module.location, usage.level, surfaceId, usage.shape);
					decisionsByKey.set(decisionKey(decision), decision);
				}
			}
		}
		var decisions = [for (decision in decisionsByKey) decision];
		decisions.sort(compareDecisions);
		var admittedCount = 0;
		for (decision in decisions) {
			if (decision.outcome == GoSurfaceDecisionOutcome.Admitted) {
				admittedCount++;
			}
		}
		return {
			schemaVersion: SCHEMA_VERSION,
			registryVersion: REGISTRY_VERSION,
			authority: AUTHORITY,
			profileAdmission: PROFILE_ADMISSION,
			catalogCount: contracts.length,
			decisionCount: decisions.length,
			admittedCount: admittedCount,
			rejectedCount: decisions.length - admittedCount,
			contracts: GoImmutableList.fromArray([for (contract in contracts) copyContract(contract)]),
			decisions: GoImmutableList.fromArray(decisions)
		};
	}

	/** Empty immutable snapshot used before typed usage is available. */
	public static function emptySnapshot():GoSurfaceContractRegistrySnapshot {
		return {
			schemaVersion: SCHEMA_VERSION,
			registryVersion: REGISTRY_VERSION,
			authority: AUTHORITY,
			profileAdmission: PROFILE_ADMISSION,
			catalogCount: 0,
			decisionCount: 0,
			admittedCount: 0,
			rejectedCount: 0,
			contracts: GoImmutableList.fromArray([]),
			decisions: GoImmutableList.fromArray([])
		};
	}

	function evaluate(module:String, location:String, usageLevel:GoTypeUsageLevelId, surfaceId:GoSurfaceId, shape:GoTypeShape):GoSurfaceDecision {
		var contract = contractsById.get(surfaceId);
		if (contract == null) {
			return missingContractDecision(module, location, usageLevel, surfaceId, shape);
		}
		var patternMatch = matchPattern(contract.eligibleShape, shape, []);
		if (!patternMatch.matched) {
			return contractDecision(module, location, usageLevel, surfaceId, shape, contract, GoSurfaceDecisionOutcome.Rejected,
				GoSurfaceDecisionReason.ShapeMismatch, "Observed type shape does not match the contract pattern.", false);
		}
		var failedRule = firstFailedEligibilityRule(contract.surfaceId, contract.eligibilityRules, shape, patternMatch.bindings);
		if (failedRule != null) {
			return contractDecision(module, location, usageLevel, surfaceId, shape, contract, GoSurfaceDecisionOutcome.Rejected,
				GoSurfaceDecisionReason.EligibilityRejected, "Eligibility rule rejected the shape: " + failedRule, false);
		}
		return contractDecision(module, location, usageLevel, surfaceId, shape, contract, GoSurfaceDecisionOutcome.Admitted,
			GoSurfaceDecisionReason.ContractAdmitted, "Contract admitted this exact typed shape.", true);
	}

	static function missingContractDecision(module:String, location:String, usageLevel:GoTypeUsageLevelId, surfaceId:GoSurfaceId,
			shape:GoTypeShape):GoSurfaceDecision {
		return {
			module: module,
			location: location,
			usageLevel: usageLevel,
			surfaceId: surfaceId,
			shape: shape,
			outcome: GoSurfaceDecisionOutcome.Rejected,
			reason: GoSurfaceDecisionReason.ContractMissing,
			detail: "Known portable surface has no admitted contract.",
			contractVersion: 0,
			selectedRepresentation: null,
			nativeImports: GoImmutableList.fromArray([]),
			runtimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: null,
			fallbackPolicy: null,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([]),
			noHxrtStatus: null,
			proofIds: GoImmutableList.fromArray([]),
			familySyncExpectation: null,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	static function contractDecision(module:String, location:String, usageLevel:GoTypeUsageLevelId, surfaceId:GoSurfaceId, shape:GoTypeShape,
			contract:GoSurfaceContract, outcome:GoSurfaceDecisionOutcome, reason:GoSurfaceDecisionReason, detail:String, admitted:Bool):GoSurfaceDecision {
		var proofIds = [for (proof in contract.proofs) proof.proofId];
		return {
			module: module,
			location: location,
			usageLevel: usageLevel,
			surfaceId: surfaceId,
			shape: shape,
			outcome: outcome,
			reason: reason,
			detail: detail,
			contractVersion: contract.contractVersion,
			selectedRepresentation: admitted ? contract.nativeRepresentation : null,
			nativeImports: admitted ? copyImports(contract.nativeImports) : GoImmutableList.fromArray([]),
			runtimeRequirements: admitted ? copyRuntimeRequirements(contract.nativeRuntimeRequirements) : GoImmutableList.fromArray([]),
			fallbackRepresentation: contract.fallbackRepresentation,
			fallbackPolicy: contract.fallbackPolicy,
			fallbackImports: copyImports(contract.fallbackImports),
			fallbackRuntimeRequirements: copyRuntimeRequirements(contract.fallbackRuntimeRequirements),
			noHxrtStatus: contract.noHxrtStatus,
			proofIds: GoImmutableList.fromArray(proofIds),
			familySyncExpectation: contract.familySyncExpectation,
			familyContractId: contract.familyContractId,
			familyContractVersion: contract.familyContractVersion
		};
	}

	static function surfaceForShape(shape:GoTypeShape):Null<GoSurfaceId> {
		return switch (shape) {
			case Nominal(_, path, parameters):
				switch (path) {
					case "Array": GoSurfaceId.HaxeArray;
					case "String": GoSurfaceId.HaxeString;
					case "haxe.io.Bytes": GoSurfaceId.HaxeBytes;
					case "haxe.ds.StringMap": GoSurfaceId.HaxeStringMap;
					case "haxe.ds.IntMap": GoSurfaceId.HaxeIntMap;
					case "haxe.ds.ObjectMap": GoSurfaceId.HaxeObjectMap;
					// Reflaxe also reports bare enum declaration/constructor identities
					// with no applied parameters. They do not choose a representation;
					// the corresponding applied typed usage carries that authority.
					case "reflaxe.std.Option" if (parameters.length > 0): GoSurfaceId.PortableOption;
					case "reflaxe.std.Result" if (parameters.length > 0): GoSurfaceId.PortableResult;
					case _: null;
				}
			case Function(_, _):
				GoSurfaceId.HaxeFunction;
			case Anonymous(_) if (isIteratorProtocolShape(shape)):
				GoSurfaceId.HaxeIterator;
			case TypeParameter(_) | Anonymous(_) | DynamicShape(_) | UnknownShape(_):
				null;
		};
	}

	/**
		Recognize only Haxe's exact structural `Iterator<T>` protocol.

		Why / What / How
		- Anonymous objects have no nominal identity, so candidate classification
		  must prove both required fields before consulting the catalog.
		- The ledger sorts fields. Exactly two non-optional zero-argument methods
		  are required: `hasNext():Bool` and `next():T`.
		- Extra fields, optional methods, parameters, or a non-Bool `hasNext`
		  remain ordinary anonymous structures and produce no registry decision.
	**/
	static function isIteratorProtocolShape(shape:GoTypeShape):Bool {
		return switch (shape) {
			case Anonymous(fields) if (fields.length == 2):
				final hasNext = fields.at(0);
				final next = fields.at(1);
				if (hasNext == null || next == null || hasNext.name != "hasNext" || hasNext.optional || next.name != "next" || next.optional) {
					false;
				} else {
					final validHasNext = switch (hasNext.shape) {
						case Function(arguments, Nominal(GoTypeUsageTargetKind.Abstract, "StdTypes.Bool", parameters)): arguments.length == 0 && parameters.length == 0;
						case _:
							false;
					};
					final validNext = switch (next.shape) {
						case Function(arguments, _): arguments.length == 0;
						case _: false;
					};
					validHasNext && validNext
					;
				}
			case _:
				false;
		};
	}

	static function patternRootMatchesSurface(surfaceId:GoSurfaceId, pattern:GoSurfaceTypePattern):Bool {
		if (pattern == null) {
			return false;
		}
		try {
			return switch [surfaceId, pattern] {
				case [GoSurfaceId.HaxeArray, NominalPattern(GoSurfaceNominalKind.Class, "Array", _)]:
					true;
				case [GoSurfaceId.HaxeString, NominalPattern(GoSurfaceNominalKind.Class, "String", _)]:
					true;
				case [
					GoSurfaceId.HaxeBytes,
					NominalPattern(GoSurfaceNominalKind.Class, "haxe.io.Bytes", _)
				]:
					true;
				case [
					GoSurfaceId.HaxeStringMap,
					NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.StringMap", _)
				]:
					true;
				case [
					GoSurfaceId.HaxeIntMap,
					NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.IntMap", _)
				]:
					true;
				case [
					GoSurfaceId.HaxeObjectMap,
					NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.ObjectMap", _)
				]:
					true;
				case [
					GoSurfaceId.PortableOption,
					NominalPattern(GoSurfaceNominalKind.Enum, "reflaxe.std.Option", _)
				]:
					true;
				case [
					GoSurfaceId.PortableResult,
					NominalPattern(GoSurfaceNominalKind.Enum, "reflaxe.std.Result", _)
				]:
					true;
				case [GoSurfaceId.HaxeIterator, AnonymousPattern(_)]:
					isIteratorProtocolPattern(pattern);
				case [GoSurfaceId.HaxeFunction, FunctionPattern(_, _)]:
					true;
				case _:
					false;
			};
		} catch (_:Exception) {
			return false;
		}
	}

	/**
		Validate the contract-side form of Haxe's exact Iterator protocol.

		Why / What / How
		- Candidate recognition alone is insufficient: a malformed catalog entry
		  must not become authority merely because it happens never to match.
		- Require the same two ordered non-optional zero-argument methods as the
		  typed shape, with Bool for `hasNext` and one bound result for `next`.
	**/
	static function isIteratorProtocolPattern(pattern:GoSurfaceTypePattern):Bool {
		return switch (pattern) {
			case AnonymousPattern(fields) if (fields != null && fields.length == 2):
				final hasNext = fields.at(0);
				final next = fields.at(1);
				if (hasNext == null || next == null || hasNext.name != "hasNext" || hasNext.optional || next.name != "next" || next.optional) {
					false;
				} else {
					final validHasNext = switch (hasNext.shape) {
						case FunctionPattern(arguments, NominalPattern(GoSurfaceNominalKind.Abstract, "StdTypes.Bool", parameters)): arguments.length == 0 && parameters.length == 0;
						case _:
							false;
					};
					final validNext = switch (next.shape) {
						case FunctionPattern(arguments, Bind(_)): arguments.length == 0;
						case _: false;
					};
					validHasNext && validNext
					;
				}
			case _:
				false;
		};
	}

	static function isValidPattern(pattern:GoSurfaceTypePattern):Bool {
		if (pattern == null) {
			return false;
		}
		try {
			return switch (pattern) {
				case Bind(name):
					isValidBindingName(name);
				case NominalPattern(kind, path, parameters):
					if (!kind.isKnown() || isBlank(path) || parameters == null) {
						false;
					} else {
						var valid = true;
						for (parameter in parameters) {
							if (!isValidPattern(parameter)) {
								valid = false;
								break;
							}
						}
						valid;
					}
				case FunctionPattern(arguments, returnType):
					if (arguments == null || returnType == null) {
						false;
					} else {
						var valid = isValidPattern(returnType);
						if (valid) {
							for (argument in arguments) {
								if (argument == null || !isValidPattern(argument.shape)) {
									valid = false;
									break;
								}
							}
						}
						valid;
					}
				case AnonymousPattern(fields):
					if (fields == null) {
						false;
					} else {
						var valid = true;
						var seen = new Map<String, Bool>();
						for (field in fields) {
							if (field == null || isBlank(field.name) || field.shape == null || seen.exists(field.name) || !isValidPattern(field.shape)) {
								valid = false;
								break;
							}
							seen.set(field.name, true);
						}
						valid;
					}
			};
		} catch (_:Exception) {
			return false;
		}
	}

	static function isValidBindingName(name:String):Bool {
		if (isBlank(name)) {
			return false;
		}
		for (index in 0...name.length) {
			var code = name.charCodeAt(index);
			var valid = code != null
				&& (code >= 65 && code <= 90 || code >= 97 && code <= 122 || code == 95 || index > 0 && code >= 48 && code <= 57);
			if (!valid) {
				return false;
			}
		}
		return true;
	}

	static function matchPattern(pattern:GoSurfaceTypePattern, shape:GoTypeShape, bindings:Map<String, GoTypeShape>):PatternMatch {
		return switch (pattern) {
			case Bind(name):
				var existing = bindings.get(name);
				if (existing == null) {
					bindings.set(name, shape);
					{matched: true, bindings: bindings};
				} else {
					{matched: shapeKey(existing) == shapeKey(shape), bindings: bindings};
				}
			case NominalPattern(expectedKind, expectedPath, expectedParameters):
				switch (shape) {
					case Nominal(actualKind, actualPath, actualParameters):
						if ((expectedKind : String) != (actualKind : String)
							|| expectedPath != actualPath
							|| expectedParameters.length != actualParameters.length) {
							{matched: false, bindings: bindings};
						} else {
							var matched = true;
							for (index in 0...expectedParameters.length) {
								var child = matchPattern(expectedParameters.at(index), actualParameters.at(index), bindings);
								if (!child.matched) {
									matched = false;
									break;
								}
							}
							{matched: matched, bindings: bindings};
						}
					case _:
						{matched: false, bindings: bindings};
				}
			case FunctionPattern(expectedArguments, expectedReturn):
				switch (shape) {
					case Function(actualArguments, actualReturn):
						if (expectedArguments.length != actualArguments.length) {
							{matched: false, bindings: bindings};
						} else {
							var matched = true;
							for (index in 0...expectedArguments.length) {
								var expected = expectedArguments.at(index);
								var actual = actualArguments.at(index);
								if (expected.optional != actual.optional
									|| !matchPattern(expected.shape, actual.shape, bindings).matched) {
									matched = false;
									break;
								}
							}
							if (matched && !matchPattern(expectedReturn, actualReturn, bindings).matched) {
								matched = false;
							}
							{matched: matched, bindings: bindings};
						}
					case _:
						{matched: false, bindings: bindings};
				}
			case AnonymousPattern(expectedFields):
				switch (shape) {
					case Anonymous(actualFields):
						if (expectedFields.length != actualFields.length) {
							{matched: false, bindings: bindings};
						} else {
							var matched = true;
							for (index in 0...expectedFields.length) {
								var expected = expectedFields.at(index);
								var actual = actualFields.at(index);
								if (expected.name != actual.name
									|| expected.optional != actual.optional
									|| !matchPattern(expected.shape, actual.shape, bindings).matched) {
									matched = false;
									break;
								}
							}
							{matched: matched, bindings: bindings};
						}
					case _:
						{matched: false, bindings: bindings};
				}
		};
	}

	static function firstFailedEligibilityRule(surfaceId:GoSurfaceId, rules:GoImmutableList<GoSurfaceEligibilityRule>, shape:GoTypeShape,
			bindings:Map<String, GoTypeShape>):Null<String> {
		for (rule in rules) {
			var failure = switch (rule) {
				case NoUnknownShapes:
					containsUnknown(shape) ? "no_unknown_shapes" : null;
				case ShapeContainsNoDynamic:
					containsDynamic(shape) ? "shape_contains_no_dynamic" : null;
				case BindingContainsNoDynamic(name): var bound = bindings.get(name); bound == null || containsDynamic(bound) ? "binding_contains_no_dynamic:" + name : null;
				case BindingIsGoComparable(name): var bound = bindings.get(name); bound == null || !isConservativelyGoComparable(bound) ? "binding_is_go_comparable:" + name : null;
				case BindingHasProvenCollectionCarrier(name): var bound = bindings.get(name); bound == null || !hasProvenCollectionCarrier(bound) ? "binding_has_proven_collection_carrier:" + name : null;
				case SurfaceHasFixedGoComparableMapKey:
					hasFixedGoComparableMapKey(surfaceId) ? null : "surface_has_fixed_go_comparable_map_key";
			};
			if (failure != null) {
				return failure;
			}
		}
		return null;
	}

	/**
		Whether a portable map surface fixes a key whose Haxe equality is exactly
		representable by Go `==`.

		Why / What / How
		- `StringMap` and `IntMap` encode their key types in the typed nominal
		  surface even though the ledger shape carries only the value parameter.
		- `ObjectMap` deliberately fails: its contract is object identity, not
		  structural Go comparability.
		- Explicit `go.Map<K,V>` never reaches this portable registry, so its
		  compatibility string coercion cannot satisfy this proof.
	**/
	static function hasFixedGoComparableMapKey(surfaceId:GoSurfaceId):Bool {
		return surfaceId == GoSurfaceId.HaxeStringMap || surfaceId == GoSurfaceId.HaxeIntMap;
	}

	/**
		Whether the ledger contains enough recursive facts to choose collection
		storage without guessing through an erased alias.

		Why
		`GoTypeShape` retains a typedef or abstract's nominal identity and applied
		parameters, but it does not currently retain the followed underlying type.
		A typedef or user abstract over `Dynamic` could therefore look free of a
		`DynamicShape` node.

		What
		Accepts classes, enums, anonymous records, and functions only when their
		recorded child shapes are also proven. Core primitive and `Null<T>`
		abstracts are closed known cases. Opaque typedefs, user abstracts, named
		type parameters, Dynamic, and unresolved shapes fail closed.

		How
		A later ledger schema may carry followed underlying-type evidence and
		broaden this rule. Until then, conservative fallback preserves behavior
		and the report names this exact rejection.
	**/
	static function hasProvenCollectionCarrier(shape:GoTypeShape):Bool {
		return switch (shape) {
			case Nominal(kind, path, parameters):
				switch (kind) {
					case GoTypeUsageTargetKind.Class | GoTypeUsageTargetKind.Enum:
						allShapes(parameters, hasProvenCollectionCarrier);
					case GoTypeUsageTargetKind.Abstract:
						switch (path) {
							case "StdTypes.Int" | "StdTypes.Float" | "StdTypes.Bool" | "StdTypes.UInt" | "StdTypes.Void":
								parameters.length == 0;
							case "StdTypes.Null": parameters.length == 1 && hasProvenCollectionCarrier(parameters.at(0));
							case _:
								false;
						}
					case GoTypeUsageTargetKind.Typedef:
						false;
					case _:
						false;
				}
			case Function(arguments, returnType):
				var proven = hasProvenCollectionCarrier(returnType);
				if (proven) {
					for (argument in arguments) {
						if (!hasProvenCollectionCarrier(argument.shape)) {
							proven = false;
							break;
						}
					}
				}
				proven;
			case Anonymous(fields):
				var proven = true;
				for (field in fields) {
					if (!hasProvenCollectionCarrier(field.shape)) {
						proven = false;
						break;
					}
				}
				proven;
			case TypeParameter(_) | DynamicShape(_) | UnknownShape(_):
				false;
		};
	}

	static function containsUnknown(shape:GoTypeShape):Bool {
		return switch (shape) {
			case UnknownShape(_):
				true;
			case Nominal(_, _, parameters):
				anyShape(parameters, containsUnknown);
			case Function(arguments, returnType):
				var found = containsUnknown(returnType);
				if (!found) {
					for (argument in arguments) {
						if (containsUnknown(argument.shape)) {
							found = true;
							break;
						}
					}
				}
				found;
			case Anonymous(fields):
				var found = false;
				for (field in fields) {
					if (containsUnknown(field.shape)) {
						found = true;
						break;
					}
				}
				found;
			case DynamicShape(inner): inner != null && containsUnknown(inner);
			case TypeParameter(_):
				false;
		};
	}

	static function containsDynamic(shape:GoTypeShape):Bool {
		return switch (shape) {
			case DynamicShape(_):
				true;
			case Nominal(_, _, parameters):
				anyShape(parameters, containsDynamic);
			case Function(arguments, returnType):
				var found = containsDynamic(returnType);
				if (!found) {
					for (argument in arguments) {
						if (containsDynamic(argument.shape)) {
							found = true;
							break;
						}
					}
				}
				found;
			case Anonymous(fields):
				var found = false;
				for (field in fields) {
					if (containsDynamic(field.shape)) {
						found = true;
						break;
					}
				}
				found;
			case TypeParameter(_) | UnknownShape(_):
				false;
		};
	}

	static function anyShape(shapes:GoImmutableList<GoTypeShape>, predicate:GoTypeShape->Bool):Bool {
		for (shape in shapes) {
			if (predicate(shape)) {
				return true;
			}
		}
		return false;
	}

	static function allShapes(shapes:GoImmutableList<GoTypeShape>, predicate:GoTypeShape->Bool):Bool {
		for (shape in shapes) {
			if (!predicate(shape)) {
				return false;
			}
		}
		return true;
	}

	static function isConservativelyGoComparable(shape:GoTypeShape):Bool {
		return switch (shape) {
			case Nominal(GoTypeUsageTargetKind.Abstract, path, parameters) if (parameters.length == 0): path == "StdTypes.Int" || path == "StdTypes.Float" || path == "StdTypes.Bool" || path == "StdTypes.UInt";
			case Nominal(GoTypeUsageTargetKind.Class, "String", parameters) if (parameters.length == 0):
				true;
			case TypeParameter(_) | Nominal(_, _, _) | Function(_, _) | Anonymous(_) | DynamicShape(_) | UnknownShape(_):
				false;
		};
	}

	static function validateEligibilityRules(contract:GoSurfaceContract, issues:Array<GoSurfaceValidationIssue>):Void {
		if (contract.eligibilityRules == null) {
			addIssue(issues, GoSurfaceValidationCode.UnknownEligibilityRule, contract.surfaceId, "Eligibility rules must be an immutable list.");
			return;
		}
		var bindings = new Map<String, Bool>();
		collectPatternBindings(contract.eligibleShape, bindings);
		for (rule in contract.eligibilityRules) {
			if (rule == null) {
				addIssue(issues, GoSurfaceValidationCode.UnknownEligibilityRule, contract.surfaceId, "Null eligibility rules are not allowed.");
				continue;
			}
			switch (rule) {
				case NoUnknownShapes | ShapeContainsNoDynamic | SurfaceHasFixedGoComparableMapKey:
				case BindingContainsNoDynamic(name) | BindingIsGoComparable(name) | BindingHasProvenCollectionCarrier(name):
					if (isBlank(name) || !bindings.exists(name)) {
						addIssue(issues, GoSurfaceValidationCode.UnboundEligibilityRule, contract.surfaceId,
							'Eligibility rule references unknown binding "$name".');
					}
				case _:
					addIssue(issues, GoSurfaceValidationCode.UnknownEligibilityRule, contract.surfaceId, "Eligibility rule is not recognized.");
			}
		}
	}

	static function collectPatternBindings(pattern:GoSurfaceTypePattern, bindings:Map<String, Bool>):Void {
		switch (pattern) {
			case Bind(name):
				if (!isBlank(name)) {
					bindings.set(name, true);
				}
			case NominalPattern(_, _, parameters):
				for (parameter in parameters) {
					collectPatternBindings(parameter, bindings);
				}
			case FunctionPattern(arguments, returnType):
				for (argument in arguments) {
					collectPatternBindings(argument.shape, bindings);
				}
				collectPatternBindings(returnType, bindings);
			case AnonymousPattern(fields):
				for (field in fields) {
					collectPatternBindings(field.shape, bindings);
				}
		}
	}

	static function validateProofs(contract:GoSurfaceContract, issues:Array<GoSurfaceValidationIssue>):Void {
		if (contract.proofs == null) {
			addIssue(issues, GoSurfaceValidationCode.MissingProof, contract.surfaceId, "Proofs must be an immutable list.");
			addIssue(issues, GoSurfaceValidationCode.MissingSemanticProof, contract.surfaceId, "Portable admission requires a semantic_diff proof.");
			return;
		}
		if (contract.proofs.length == 0) {
			addIssue(issues, GoSurfaceValidationCode.MissingProof, contract.surfaceId, "Admitted contracts require at least one proof fixture.");
		}
		var hasSemanticDiff = false;
		var seenProofs = new Map<String, Bool>();
		for (proof in contract.proofs) {
			if (proof == null) {
				addIssue(issues, GoSurfaceValidationCode.MissingProof, contract.surfaceId, "Null proof entries are not allowed.");
				continue;
			}
			if (proof.kind == GoSurfaceProofKind.SemanticDiff) {
				hasSemanticDiff = true;
			}
			if (!proof.kind.isKnown()) {
				addIssue(issues, GoSurfaceValidationCode.UnknownProofKind, contract.surfaceId, "Proof kind is not recognized.");
			}
			if (isBlank(proof.proofId)) {
				addIssue(issues, GoSurfaceValidationCode.MissingProof, contract.surfaceId, "Proof IDs must be non-empty.");
			} else if (seenProofs.exists(proof.proofId)) {
				addIssue(issues, GoSurfaceValidationCode.DuplicateProof, contract.surfaceId, 'Proof ID "${proof.proofId}" appears more than once.');
			} else {
				seenProofs.set(proof.proofId, true);
			}
			if (!isSafeRelativePath(proof.fixturePath)) {
				addIssue(issues, GoSurfaceValidationCode.UnsafeProofPath, contract.surfaceId,
					'Proof path "${proof.fixturePath}" must be repository-relative without traversal.');
			}
		}
		if (!hasSemanticDiff) {
			addIssue(issues, GoSurfaceValidationCode.MissingSemanticProof, contract.surfaceId, "Portable admission requires at least one semantic_diff proof.");
		}
	}

	static function validateRuntimeRequirements(surfaceId:String, requirements:GoImmutableList<GoHxrtFeatureId>, issues:Array<GoSurfaceValidationIssue>):Void {
		if (requirements == null) {
			addIssue(issues, GoSurfaceValidationCode.UnknownRuntimeRequirement, surfaceId, "Runtime requirements must be an immutable list.");
			return;
		}
		for (feature in requirements) {
			if (feature == null || !GoHxrtFeatureAnalyzer.isKnownFeature(feature)) {
				addIssue(issues, GoSurfaceValidationCode.UnknownRuntimeRequirement, surfaceId, 'Runtime requirement "$feature" is not a known hxrt feature.');
			}
		}
	}

	static function validateImports(surfaceId:String, imports:GoImmutableList<GoSurfaceImportRequirement>, issues:Array<GoSurfaceValidationIssue>):Void {
		if (imports == null) {
			addIssue(issues, GoSurfaceValidationCode.InvalidImportRequirement, surfaceId, "Import requirements must be an immutable list.");
			return;
		}
		for (requirement in imports) {
			if (requirement == null) {
				addIssue(issues, GoSurfaceValidationCode.InvalidImportRequirement, surfaceId, "Null import requirements are not allowed.");
				continue;
			}
			var path = requirement.path == null ? "" : StringTools.trim(requirement.path);
			if (path == "" || isBlank(requirement.reason) || StringTools.startsWith(path, "/") || path.indexOf("\\") >= 0 || path.indexOf("..") >= 0
				|| path.indexOf(" ") >= 0) {
				addIssue(issues, GoSurfaceValidationCode.InvalidImportRequirement, surfaceId,
					"Import requirements need a safe Go path and a non-empty reason.");
			}
		}
	}

	static function validateFamilySync(contract:GoSurfaceContract, issues:Array<GoSurfaceValidationIssue>):Void {
		if (!contract.familySyncExpectation.isKnown()) {
			addIssue(issues, GoSurfaceValidationCode.InvalidFamilySync, contract.surfaceId, "Family synchronization expectation is not recognized.");
			return;
		}
		switch (contract.familySyncExpectation) {
			case GoFamilySyncExpectation.TargetLocal:
				if (!isBlank(contract.familyContractId) || contract.familyContractVersion != 0) {
					addIssue(issues, GoSurfaceValidationCode.InvalidFamilySync, contract.surfaceId,
						"Target-local contracts must not claim a family contract ID/version.");
				}
			case GoFamilySyncExpectation.SharedContractRequired:
				if (isBlank(contract.familyContractId) || contract.familyContractVersion < 1) {
					addIssue(issues, GoSurfaceValidationCode.InvalidFamilySync, contract.surfaceId,
						"Shared contracts require a stable family contract ID and positive version.");
				}
			case _:
		}
	}

	static function copyContract(contract:GoSurfaceContract):GoSurfaceContract {
		var rules = [for (rule in contract.eligibilityRules) rule];
		rules.sort((a, b) -> compareStrings(eligibilityRuleKey(a), eligibilityRuleKey(b)));
		var proofs:Array<GoSurfaceProof> = [
			for (proof in contract.proofs)
				{
					proofId: proof.proofId,
					kind: proof.kind,
					fixturePath: proof.fixturePath
				}
		];
		proofs.sort((a, b) -> compareStrings(proofKey(a), proofKey(b)));
		return {
			surfaceId: contract.surfaceId,
			contractVersion: contract.contractVersion,
			sourceContract: contract.sourceContract,
			sourceSemanticsId: contract.sourceSemanticsId,
			sourceSemanticsVersion: contract.sourceSemanticsVersion,
			sourceSemantics: contract.sourceSemantics,
			eligibleShape: copyPattern(contract.eligibleShape),
			eligibilityRules: GoImmutableList.fromArray(rules),
			nativeRepresentation: contract.nativeRepresentation,
			nativeImports: copyImports(contract.nativeImports),
			nativeRuntimeRequirements: copyRuntimeRequirements(contract.nativeRuntimeRequirements),
			fallbackRepresentation: contract.fallbackRepresentation,
			fallbackPolicy: contract.fallbackPolicy,
			fallbackImports: copyImports(contract.fallbackImports),
			fallbackRuntimeRequirements: copyRuntimeRequirements(contract.fallbackRuntimeRequirements),
			noHxrtStatus: contract.noHxrtStatus,
			proofs: GoImmutableList.fromArray(proofs),
			familySyncExpectation: contract.familySyncExpectation,
			familyContractId: contract.familyContractId,
			familyContractVersion: contract.familyContractVersion
		};
	}

	static function copyPattern(pattern:GoSurfaceTypePattern):GoSurfaceTypePattern {
		return switch (pattern) {
			case Bind(name):
				Bind(name);
			case NominalPattern(kind, path, parameters):
				NominalPattern(kind, path, GoImmutableList.fromArray([for (parameter in parameters) copyPattern(parameter)]));
			case FunctionPattern(arguments, returnType):
				var copiedArguments = [
					for (argument in arguments)
						{
							optional: argument.optional,
							shape: copyPattern(argument.shape)
						}
				];
				FunctionPattern(GoImmutableList.fromArray(copiedArguments), copyPattern(returnType));
			case AnonymousPattern(fields):
				var copiedFields = [
					for (field in fields)
						{
							name: field.name,
							optional: field.optional,
							shape: copyPattern(field.shape)
						}
				];
				AnonymousPattern(GoImmutableList.fromArray(copiedFields));
		};
	}

	static function copyImports(imports:GoImmutableList<GoSurfaceImportRequirement>):GoImmutableList<GoSurfaceImportRequirement> {
		var copied:Array<GoSurfaceImportRequirement> = [
			for (requirement in imports)
				{
					path: requirement.path,
					reason: requirement.reason
				}
		];
		copied.sort((a, b) -> compareStrings(importKey(a), importKey(b)));
		return GoImmutableList.fromArray(copied);
	}

	static function copyRuntimeRequirements(requirements:GoImmutableList<GoHxrtFeatureId>):GoImmutableList<GoHxrtFeatureId> {
		var copied = [for (feature in requirements) feature];
		copied.sort((a, b) -> compareStrings(a, b));
		var unique = new Array<GoHxrtFeatureId>();
		for (feature in copied) {
			if (unique.length == 0 || unique[unique.length - 1] != feature) {
				unique.push(feature);
			}
		}
		return GoImmutableList.fromArray(unique);
	}

	/** Render the registry authority as deterministic JSON. */
	public static function renderJson(snapshot:GoSurfaceContractRegistrySnapshot):String {
		var lines = [
			"{",
			'\t"schemaVersion": ' + snapshot.schemaVersion + ",",
			'\t"registryVersion": ' + snapshot.registryVersion + ",",
			'\t"authority": "' + jsonEscape(snapshot.authority) + '",',
			'\t"profileAdmission": "' + jsonEscape(snapshot.profileAdmission) + '",',
			'\t"catalogCount": ' + snapshot.catalogCount + ",",
			'\t"decisionCount": ' + snapshot.decisionCount + ",",
			'\t"admittedCount": ' + snapshot.admittedCount + ",",
			'\t"rejectedCount": ' + snapshot.rejectedCount + ",",
			'\t"contracts": ['
		];
		for (index in 0...snapshot.contracts.length) {
			appendContractJson(lines, snapshot.contracts.at(index), "\t\t");
			if (index + 1 < snapshot.contracts.length) {
				lines[lines.length - 1] += ",";
			}
		}
		lines.push("\t],");
		lines.push('\t"decisions": [');
		for (index in 0...snapshot.decisions.length) {
			appendDecisionJson(lines, snapshot.decisions.at(index), "\t\t");
			if (index + 1 < snapshot.decisions.length) {
				lines[lines.length - 1] += ",";
			}
		}
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	/** Render the same authority as a concise human-readable report. */
	public static function renderMarkdown(snapshot:GoSurfaceContractRegistrySnapshot):String {
		var lines = [
			"# Go Surface Contract Registry",
			"",
			"- Schema version: `" + snapshot.schemaVersion + "`",
			"- Registry version: `" + snapshot.registryVersion + "`",
			"- Authority: `" + snapshot.authority + "`",
			"- Profile admission: `" + snapshot.profileAdmission + "`",
			"- Catalog entries: `" + snapshot.catalogCount + "`",
			"- Decisions: `" + snapshot.decisionCount + "`",
			"- Admitted: `" + snapshot.admittedCount + "`",
			"- Rejected: `" + snapshot.rejectedCount + "`",
			"",
			"## Contracts",
			""
		];
		if (snapshot.contracts.length == 0) {
			lines.push("- None admitted.");
		} else {
			for (contract in snapshot.contracts) {
				lines.push("- `" + contract.surfaceId + "` v" + contract.contractVersion + ": `" + contract.nativeRepresentation + "`; fallback `"
					+ contract.fallbackRepresentation + "`; source semantics `" + contract.sourceSemanticsId + "` v" + contract.sourceSemanticsVersion + ".");
			}
		}
		lines.push("");
		lines.push("## Decisions");
		lines.push("");
		if (snapshot.decisions.length == 0) {
			lines.push("- No known portable surface usage observed.");
		} else {
			for (decision in snapshot.decisions) {
				var label = switch (decision.reason) {
					case GoSurfaceDecisionReason.ContractAdmitted: "Contract admitted";
					case GoSurfaceDecisionReason.ContractMissing: "Contract missing";
					case GoSurfaceDecisionReason.ShapeMismatch: "Shape mismatch";
					case GoSurfaceDecisionReason.EligibilityRejected: "Eligibility rejected";
					case _: "Registry decision";
				};
				lines.push("- " + label + " `" + decision.surfaceId + "` in `" + decision.module + "` at `" + decision.location + "` (`"
					+ decision.usageLevel + "`, `" + shapeKey(decision.shape) + "`): " + decision.detail);
			}
		}
		return lines.join("\n") + "\n";
	}

	static function appendContractJson(lines:Array<String>, contract:GoSurfaceContract, indent:String):Void {
		lines.push(indent + "{");
		lines.push(indent + '\t"surfaceId": "' + jsonEscape(contract.surfaceId) + '",');
		lines.push(indent + '\t"contractVersion": ' + contract.contractVersion + ",");
		lines.push(indent + '\t"sourceContract": "' + jsonEscape(contract.sourceContract) + '",');
		lines.push(indent + '\t"sourceSemanticsId": "' + jsonEscape(contract.sourceSemanticsId) + '",');
		lines.push(indent + '\t"sourceSemanticsVersion": ' + contract.sourceSemanticsVersion + ",");
		lines.push(indent + '\t"sourceSemantics": "' + jsonEscape(contract.sourceSemantics) + '",');
		lines.push(indent + '\t"eligibleShape": ' + patternJson(contract.eligibleShape) + ",");
		lines.push(indent + '\t"eligibilityRules": ' + eligibilityRulesJson(contract.eligibilityRules) + ",");
		lines.push(indent + '\t"nativeRepresentation": "' + jsonEscape(contract.nativeRepresentation) + '",');
		lines.push(indent + '\t"nativeImports": ' + importsJson(contract.nativeImports) + ",");
		lines.push(indent + '\t"nativeRuntimeRequirements": ' + runtimeRequirementsJson(contract.nativeRuntimeRequirements) + ",");
		lines.push(indent + '\t"fallbackRepresentation": "' + jsonEscape(contract.fallbackRepresentation) + '",');
		lines.push(indent + '\t"fallbackPolicy": "' + jsonEscape(contract.fallbackPolicy) + '",');
		lines.push(indent + '\t"fallbackImports": ' + importsJson(contract.fallbackImports) + ",");
		lines.push(indent + '\t"fallbackRuntimeRequirements": ' + runtimeRequirementsJson(contract.fallbackRuntimeRequirements) + ",");
		lines.push(indent + '\t"noHxrtStatus": "' + jsonEscape(contract.noHxrtStatus) + '",');
		lines.push(indent + '\t"proofs": ' + proofsJson(contract.proofs) + ",");
		lines.push(indent + '\t"familySyncExpectation": "' + jsonEscape(contract.familySyncExpectation) + '",');
		lines.push(indent + '\t"familyContractId": "' + jsonEscape(contract.familyContractId) + '",');
		lines.push(indent + '\t"familyContractVersion": ' + contract.familyContractVersion);
		lines.push(indent + "}");
	}

	static function appendDecisionJson(lines:Array<String>, decision:GoSurfaceDecision, indent:String):Void {
		lines.push(indent + "{");
		lines.push(indent + '\t"module": "' + jsonEscape(decision.module) + '",');
		lines.push(indent + '\t"location": "' + jsonEscape(decision.location) + '",');
		lines.push(indent + '\t"usageLevel": "' + jsonEscape(decision.usageLevel) + '",');
		lines.push(indent + '\t"surfaceId": "' + jsonEscape(decision.surfaceId) + '",');
		lines.push(indent + '\t"shape": ' + shapeJson(decision.shape) + ",");
		lines.push(indent + '\t"outcome": "' + jsonEscape(decision.outcome) + '",');
		lines.push(indent + '\t"reason": "' + jsonEscape(decision.reason) + '",');
		lines.push(indent + '\t"detail": "' + jsonEscape(decision.detail) + '",');
		lines.push(indent + '\t"contractVersion": ' + decision.contractVersion + ",");
		lines.push(indent + '\t"selectedRepresentation": ' + nullableStringJson(decision.selectedRepresentation) + ",");
		lines.push(indent + '\t"nativeImports": ' + importsJson(decision.nativeImports) + ",");
		lines.push(indent + '\t"runtimeRequirements": ' + runtimeRequirementsJson(decision.runtimeRequirements) + ",");
		lines.push(indent + '\t"fallbackRepresentation": ' + nullableStringJson(decision.fallbackRepresentation) + ",");
		lines.push(indent + '\t"fallbackPolicy": ' + nullableStringJson(decision.fallbackPolicy) + ",");
		lines.push(indent + '\t"fallbackImports": ' + importsJson(decision.fallbackImports) + ",");
		lines.push(indent + '\t"fallbackRuntimeRequirements": ' + runtimeRequirementsJson(decision.fallbackRuntimeRequirements) + ",");
		lines.push(indent + '\t"noHxrtStatus": ' + nullableStringJson(decision.noHxrtStatus) + ",");
		lines.push(indent + '\t"proofIds": ' + stringListJson(decision.proofIds) + ",");
		lines.push(indent + '\t"familySyncExpectation": ' + nullableStringJson(decision.familySyncExpectation) + ",");
		lines.push(indent + '\t"familyContractId": "' + jsonEscape(decision.familyContractId) + '",');
		lines.push(indent + '\t"familyContractVersion": ' + decision.familyContractVersion);
		lines.push(indent + "}");
	}

	static function patternJson(pattern:GoSurfaceTypePattern):String {
		return switch (pattern) {
			case Bind(name):
				'{"kind":"bind","name":"' + jsonEscape(name) + '"}';
			case NominalPattern(kind, path, parameters):
				'{"kind":"nominal","targetKind":"'
				+ jsonEscape(kind)
				+ '","path":"'
				+ jsonEscape(path)
				+ '","parameters":'
				+ patternListJson(parameters)
				+ "}";
			case FunctionPattern(arguments, returnType):
				var values = new Array<String>();
				for (argument in arguments) {
					values.push('{"optional":' + argument.optional + ',"shape":' + patternJson(argument.shape) + "}");
				}
				'{"kind":"function","arguments":['
				+ values.join(",")
				+ '],"returnType":'
				+ patternJson(returnType)
				+ "}";
			case AnonymousPattern(fields):
				var values = new Array<String>();
				for (field in fields) {
					values.push('{"name":"' + jsonEscape(field.name) + '","optional":' + field.optional + ',"shape":' + patternJson(field.shape) + "}");
				}
				'{"kind":"anonymous","fields":[' + values.join(",") + "]}";
		};
	}

	static function shapeJson(shape:GoTypeShape):String {
		return switch (shape) {
			case Nominal(kind, path, parameters):
				shapeObjectJson(kind, path, shapeListJson(parameters), "[]", "null", "[]");
			case TypeParameter(path):
				shapeObjectJson("type_parameter", path, "[]", "[]", "null", "[]");
			case Function(arguments, returnType):
				var values = new Array<String>();
				for (argument in arguments) {
					values.push('{"name":"'
						+ jsonEscape(argument.name)
						+ '","optional":'
						+ argument.optional
						+ ',"shape":'
						+ shapeJson(argument.shape)
						+ "}");
				}
				shapeObjectJson("function", "", "[]", "[" + values.join(",") + "]", shapeJson(returnType), "[]");
			case Anonymous(fields):
				var values = new Array<String>();
				for (field in fields) {
					values.push('{"name":"' + jsonEscape(field.name) + '","optional":' + field.optional + ',"shape":' + shapeJson(field.shape) + "}");
				}
				shapeObjectJson("anonymous", "", "[]", "[]", "null", "[" + values.join(",") + "]");
			case DynamicShape(inner):
				shapeObjectJson("dynamic", "", "[]", "[]", inner == null ? "null" : shapeJson(inner), "[]");
			case UnknownShape(reason):
				shapeObjectJson("unknown", reason, "[]", "[]", "null", "[]");
		};
	}

	static function shapeObjectJson(kind:String, path:String, parameters:String, arguments:String, returnType:String, fields:String):String {
		return '{"kind":"' + jsonEscape(kind) + '","path":"' + jsonEscape(path) + '","parameters":' + parameters + ',"arguments":' + arguments
			+ ',"returnType":' + returnType + ',"fields":' + fields + "}";
	}

	static function patternListJson(patterns:GoImmutableList<GoSurfaceTypePattern>):String {
		var values = [for (pattern in patterns) patternJson(pattern)];
		return "[" + values.join(",") + "]";
	}

	static function shapeListJson(shapes:GoImmutableList<GoTypeShape>):String {
		var values = [for (shape in shapes) shapeJson(shape)];
		return "[" + values.join(",") + "]";
	}

	static function eligibilityRulesJson(rules:GoImmutableList<GoSurfaceEligibilityRule>):String {
		var values = [for (rule in rules) '"' + jsonEscape(eligibilityRuleKey(rule)) + '"'];
		return "[" + values.join(",") + "]";
	}

	static function importsJson(imports:GoImmutableList<GoSurfaceImportRequirement>):String {
		var values = [
			for (requirement in imports)
				'{"path":"'
				+ jsonEscape(requirement.path)
				+ '","reason":"'
				+ jsonEscape(requirement.reason)
				+ '"}'];
		return "[" + values.join(",") + "]";
	}

	static function runtimeRequirementsJson(requirements:GoImmutableList<GoHxrtFeatureId>):String {
		var values = [for (feature in requirements) '"' + jsonEscape(feature) + '"'];
		return "[" + values.join(",") + "]";
	}

	static function proofsJson(proofs:GoImmutableList<GoSurfaceProof>):String {
		var values = [
			for (proof in proofs)
				'{"proofId":"'
				+ jsonEscape(proof.proofId)
				+ '","kind":"'
				+ jsonEscape(proof.kind)
				+ '","fixturePath":"'
				+ jsonEscape(proof.fixturePath)
				+ '"}'];
		return "[" + values.join(",") + "]";
	}

	static function stringListJson(values:GoImmutableList<String>):String {
		var encoded = [for (value in values) '"' + jsonEscape(value) + '"'];
		return "[" + encoded.join(",") + "]";
	}

	static function nullableStringJson(value:Null<String>):String {
		return value == null ? "null" : '"' + jsonEscape(value) + '"';
	}

	static function shapeKey(shape:GoTypeShape):String {
		return shapeJson(shape);
	}

	static function decisionKey(decision:GoSurfaceDecision):String {
		return decision.module
			+ "\n"
			+ decision.location
			+ "\n"
			+ decision.usageLevel
			+ "\n"
			+ decision.surfaceId
			+ "\n"
			+ shapeKey(decision.shape);
	}

	static function eligibilityRuleKey(rule:GoSurfaceEligibilityRule):String {
		return switch (rule) {
			case NoUnknownShapes: "no_unknown_shapes";
			case ShapeContainsNoDynamic: "shape_contains_no_dynamic";
			case BindingContainsNoDynamic(name): "binding_contains_no_dynamic:" + name;
			case BindingIsGoComparable(name): "binding_is_go_comparable:" + name;
			case BindingHasProvenCollectionCarrier(name): "binding_has_proven_collection_carrier:" + name;
			case SurfaceHasFixedGoComparableMapKey: "surface_has_fixed_go_comparable_map_key";
		};
	}

	static function proofKey(proof:GoSurfaceProof):String {
		return proof.proofId + "\n" + proof.kind + "\n" + proof.fixturePath;
	}

	static function importKey(requirement:GoSurfaceImportRequirement):String {
		return requirement.path + "\n" + requirement.reason;
	}

	static function compareContracts(a:GoSurfaceContract, b:GoSurfaceContract):Int {
		return compareStrings(a.surfaceId, b.surfaceId);
	}

	static function compareDecisions(a:GoSurfaceDecision, b:GoSurfaceDecision):Int {
		return compareStrings(decisionKey(a), decisionKey(b));
	}

	static function compareValidationIssues(a:GoSurfaceValidationIssue, b:GoSurfaceValidationIssue):Int {
		return compareStrings(a.surfaceId + "\n" + a.code + "\n" + a.detail, b.surfaceId + "\n" + b.code + "\n" + b.detail);
	}

	static function addIssue(issues:Array<GoSurfaceValidationIssue>, code:GoSurfaceValidationCode, surfaceId:String, detail:String):Void {
		issues.push({
			code: code,
			surfaceId: surfaceId == null ? "" : surfaceId,
			detail: detail
		});
	}

	static function isSafeRelativePath(path:String):Bool {
		if (isBlank(path)
			|| StringTools.startsWith(path, "/")
			|| path.indexOf("\\") >= 0
			|| (path.length >= 2 && isAsciiLetter(path.charCodeAt(0)) && path.charAt(1) == ":")) {
			return false;
		}
		for (part in path.split("/")) {
			if (part == ".." || part == "") {
				return false;
			}
		}
		return true;
	}

	static function isAsciiLetter(code:Null<Int>):Bool {
		return code != null && (code >= 65 && code <= 90 || code >= 97 && code <= 122);
	}

	static function isBlank(value:String):Bool {
		return value == null || StringTools.trim(value) == "";
	}

	static inline function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function jsonEscape(value:String):String {
		if (value == null) {
			return "";
		}
		var escaped = new StringBuf();
		for (index in 0...value.length) {
			var code = value.charCodeAt(index);
			switch (code) {
				case 8:
					escaped.add("\\b");
				case 9:
					escaped.add("\\t");
				case 10:
					escaped.add("\\n");
				case 12:
					escaped.add("\\f");
				case 13:
					escaped.add("\\r");
				case 34:
					escaped.add('\\"');
				case 92:
					escaped.add("\\\\");
				case control if (control != null && control < 32):
					escaped.add("\\u" + StringTools.hex(control, 4).toLowerCase());
				case _:
					escaped.addChar(code);
			}
		}
		return escaped.toString();
	}
}
#end
