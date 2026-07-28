package reflaxe.go.compiler;

#if (macro || reflaxe_runtime)
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.compiler.TypeUsageTracker.TypeOrModuleType;
import reflaxe.compiler.TypeUsageTracker.TypeUsageLevel;
import reflaxe.compiler.TypeUsageTracker.TypeUsageMap;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureReason;
import reflaxe.helpers.TypeHelper;

using reflaxe.helpers.ModuleTypeHelper;

/**
	Stable outward names for Reflaxe's typed usage levels.

	Why
	Reflaxe represents levels as bit values, which are correct inside the tracker
	but unsuitable as an unexplained public report contract.

	What
	Names every tracker level with one deterministic report spelling.

	How
	`fromReflaxe(...)` is the only conversion point from the upstream enum
	abstract; arbitrary strings cannot enter the compiler-side ledger.
**/
enum abstract GoTypeUsageLevelId(String) to String {
	var Expression = "expression";
	var VariableType = "variable_type";
	var StaticAccess = "static_access";
	var Constructed = "constructed";
	var FunctionDeclaration = "function_declaration";
	var VariableDeclaration = "variable_declaration";
	var ExtendedFrom = "extended_from";

	public static function fromReflaxe(level:TypeUsageLevel):GoTypeUsageLevelId {
		return switch (Std.int(level)) {
			case 1: Expression;
			case 2: VariableType;
			case 4: StaticAccess;
			case 8: Constructed;
			case 16: FunctionDeclaration;
			case 32: VariableDeclaration;
			case 64: ExtendedFrom;
			case _:
				Context.fatalError("Unknown Reflaxe type-usage level: " + Std.int(level), Context.currentPos());
				Expression;
		};
	}
}

/**
	Closed target categories used by the typed usage report.

	Why
	Downstream registry and planner code must distinguish nominal types from
	function and anonymous carriers without inspecting macro objects or `Dynamic`.

	What
	Classifies the stable target facts exposed by Reflaxe's usage map.

	How
	Module types map to their Haxe declaration category; non-nominal `Type`
	entries use the final two categories.
**/
enum abstract GoTypeUsageTargetKind(String) to String {
	var Class = "class";
	var Enum = "enum";
	var Typedef = "typedef";
	var Abstract = "abstract";
	var Function = "function";
	var Anonymous = "anonymous";
}

/**
	Closed member-observation categories recorded beside type usage.

	Why
	Reflaxe's usage map proves which types participate, while registry planning
	also needs to know whether source constructed a type, selected a field, or
	called a member.

	What
	Names the three typed expression shapes this ledger records.

	How
	The collector matches `TNew`, `TField`, and `TCall` nodes and never scans
	source text.
**/
enum abstract GoMemberUsageKind(String) to String {
	var Construct = "construct";
	var FieldAccess = "field_access";
	var Call = "call";
}

/**
	A closed reason for observing a native import.

	Why
	An open string would let tracker levels, member kinds, or spelling mistakes
	enter the public report without schema validation.

	What
	Combines the seven Reflaxe usage levels with the three member observations
	that can select an import.

	How
	The two conversion functions are the only entry points from the upstream
	tracker and the member walker.
**/
enum abstract GoNativeImportUsageKind(String) to String {
	var Expression = "expression";
	var VariableType = "variable_type";
	var StaticAccess = "static_access";
	var Constructed = "constructed";
	var FunctionDeclaration = "function_declaration";
	var VariableDeclaration = "variable_declaration";
	var ExtendedFrom = "extended_from";
	var Construct = "construct";
	var FieldAccess = "field_access";
	var Call = "call";

	public static function fromTypeUsage(level:GoTypeUsageLevelId):GoNativeImportUsageKind {
		return switch (level) {
			case GoTypeUsageLevelId.Expression: Expression;
			case GoTypeUsageLevelId.VariableType: VariableType;
			case GoTypeUsageLevelId.StaticAccess: StaticAccess;
			case GoTypeUsageLevelId.Constructed: Constructed;
			case GoTypeUsageLevelId.FunctionDeclaration: FunctionDeclaration;
			case GoTypeUsageLevelId.VariableDeclaration: VariableDeclaration;
			case GoTypeUsageLevelId.ExtendedFrom: ExtendedFrom;
		};
	}

	public static function fromMemberUsage(kind:GoMemberUsageKind):GoNativeImportUsageKind {
		return switch (kind) {
			case GoMemberUsageKind.Construct: Construct;
			case GoMemberUsageKind.FieldAccess: FieldAccess;
			case GoMemberUsageKind.Call: Call;
		};
	}
}

/**
	A deeply read-only ordered collection used at the compiler authority seam.

	Why
	`haxe.ds.ReadOnlyArray` is only an aliasing view: another holder can still
	mutate its underlying `Array`. Planner evidence must not change after it is
	published on `CompilationContext`.

	What
	Exposes only length, `at(index)` reads, and iteration.

	How
	`fromArray(...)` clones its input and the stored array remains private, so no
	mutable alias crosses the boundary.
**/
class GoImmutableList<T> {
	final values:Array<T>;

	public var length(get, never):Int;

	private function new(values:Array<T>) {
		this.values = values.copy();
	}

	public static function fromArray<T>(values:Array<T>):GoImmutableList<T> {
		return new GoImmutableList<T>(values == null ? [] : values);
	}

	inline function get_length():Int {
		return values.length;
	}

	public inline function at(index:Int):T {
		return values[index];
	}

	public function iterator():Iterator<T> {
		return values.iterator();
	}
}

/**
	Closed terminal markers for a shape that cannot expand further.

	Why
	Recursive macro types and unresolved monomorphs must terminate without
	letting arbitrary diagnostic strings enter registry authority.

	What
	Names the four deterministic non-shape outcomes.

	How
	`typeShape(...)` emits these only at explicit null, recursion, or depth
	guards.
**/
enum abstract GoUnknownTypeShapeReason(String) to String {
	var Missing = "missing";
	var DepthLimit = "depth_limit";
	var Monomorph = "monomorph";
	var Recursive = "recursive";
}

/**
	A macro-object-free algebra for the typed Haxe shapes observed by the ledger.

	Why
	Flattening `UsedBox<Int>` into unrelated `UsedBox` and `Int` facts, or every
	callable into `<function>`, loses the relationships a representation registry
	must prove.

	What
	Preserves nominal parameters, type parameters, function signatures, anonymous
	fields, dynamic inner types, and deterministic recursion/unknown markers.

	How
	The collector converts `haxe.macro.Type` values immediately and stores only
	strings, booleans, enum constructors, and `GoImmutableList` children.
**/
enum GoTypeShape {
	Nominal(kind:GoTypeUsageTargetKind, path:String, parameters:GoImmutableList<GoTypeShape>);
	TypeParameter(path:String);
	Function(arguments:GoImmutableList<GoFunctionArgumentShape>, returnType:GoTypeShape);
	Anonymous(fields:GoImmutableList<GoAnonymousFieldShape>);
	DynamicShape(inner:Null<GoTypeShape>);
	UnknownShape(reason:GoUnknownTypeShapeReason);
}

/**
	Why: Optional argument semantics and ordered parameter shape affect admission.
	What: One immutable function parameter inside `GoTypeShape.Function`.
	How: Preserve the typed name/optional flag and recursively converted shape.
**/
typedef GoFunctionArgumentShape = {
	final name:String;
	final optional:Bool;
	final shape:GoTypeShape;
}

/**
	Why: Anonymous structures are admitted by field contract, not by a fake name.
	What: One sorted immutable field inside `GoTypeShape.Anonymous`.
	How: Preserve name/optionality and recursively convert the declared field type.
**/
typedef GoAnonymousFieldShape = {
	final name:String;
	final optional:Bool;
	final shape:GoTypeShape;
}

/**
	Why: A shape has different meaning when constructed, declared, or only read.
	What: One immutable shape observation at a closed Reflaxe usage level.
	How: Pair the converted algebra with `GoTypeUsageLevelId`.
**/
typedef GoTypeUsageEvidence = {
	final level:GoTypeUsageLevelId;
	final shape:GoTypeShape;
}

/**
	Why: Registry diagnostics need the operation that made a target relevant.
	What: One immutable constructor, field-access, or call observation.
	How: Store only stable target/member names and a normalized `Module:line`.
**/
typedef GoMemberUsageEvidence = {
	final kind:GoMemberUsageKind;
	final target:String;
	final member:String;
	final location:String;
}

/**
	Why: `hxrt` metadata is a token, while generated Go imports a module path.
	What: One immutable native import observation with both spellings.
	How: Resolve the token at snapshot time and retain a closed usage reason.
**/
typedef GoNativeImportUsageEvidence = {
	final target:String;
	final metadataImportPath:String;
	final resolvedImportPath:String;
	final usageKind:GoNativeImportUsageKind;
	final location:String;
}

/**
	Why: Later planners need evidence grouped by its project-owned declaration.
	What: One immutable declaration owner and all sorted observations beneath it.
	How: Deep-copy each collection into `GoImmutableList` before publication.
**/
typedef GoTypeUsageModuleEvidence = {
	final module:String;
	final kind:GoTypeUsageTargetKind;
	final location:String;
	final typeUsages:GoImmutableList<GoTypeUsageEvidence>;
	final memberUsages:GoImmutableList<GoMemberUsageEvidence>;
	final nativeImports:GoImmutableList<GoNativeImportUsageEvidence>;
}

/**
	Why: Runtime consequences must remain attributable in the optional report.
	What: One immutable selected `hxrt` feature reason.
	How: Copy the existing typed runtime inference reason into stable strings.
**/
typedef GoCapabilityUsageEvidence = {
	final feature:String;
	final sourceKind:String;
	final source:String;
}

/**
	Why: Planner authority must not change after `CompilationContext` receives it.
	What: The complete deterministic, deeply read-only usage snapshot.
	How: Publish final scalars and `GoImmutableList` trees; post-lowering capability
	reporting uses a separate snapshot rather than mutating this one.
**/
typedef GoTypeUsageLedgerSnapshot = {
	final schemaVersion:Int;
	final source:String;
	final scannerFallback:String;
	final moduleCount:Int;
	final typeUsageCount:Int;
	final memberUsageCount:Int;
	final nativeImportCount:Int;
	final capabilityCount:Int;
	final modules:GoImmutableList<GoTypeUsageModuleEvidence>;
	final capabilities:GoImmutableList<GoCapabilityUsageEvidence>;
}

private typedef MutableModuleEvidence = {
	var module:String;
	var kind:GoTypeUsageTargetKind;
	var location:String;
	var typeUsages:Map<String, GoTypeUsageEvidence>;
	var memberUsages:Map<String, GoMemberUsageEvidence>;
	var nativeImports:Map<String, MutableNativeImportEvidence>;
}

private typedef MutableNativeImportEvidence = {
	final target:String;
	final metadataImportPath:String;
	final usageKind:GoNativeImportUsageKind;
	final location:String;
}

private typedef ResolvedMember = {
	final target:String;
	final member:String;
	final moduleType:Null<ModuleType>;
}

/**
	Collects Reflaxe type usage into a deterministic, path-safe compiler ledger.

	Why
	Runtime slicing and portable specialization cannot safely use source-text
	import scans as semantic authority. Reflaxe already observes the post-typing,
	post-DCE program, but haxe.go previously discarded that evidence.

	What
	Stores project-owner modules, their typed target usage, typed member/call
	observations, native import metadata, and later runtime-capability reasons.

	How
	Compiler callbacks supply `getTypeUsage()` while Reflaxe's current module is
	active. This collector removes macro objects at the boundary, normalizes
	locations to `Module:line`, deduplicates by stable keys, and sorts every
	outward array. The immutable snapshot is placed in `CompilationContext` for
	later registry/planner consumption.
**/
class GoTypeUsageLedger {
	public static inline final SCHEMA_VERSION = 1;
	public static inline final SOURCE = "reflaxe_type_usage_tracker";
	public static inline final SCANNER_FALLBACK = "transitional_contract_diagnostics_only";

	final modules:Map<String, MutableModuleEvidence> = [];

	public function new() {}

	public function collect(moduleType:ModuleType, usage:Null<TypeUsageMap>):Void {
		if (!isProjectOwner(moduleType)) {
			return;
		}

		var module = moduleName(moduleType);
		var entry = modules.get(module);
		if (entry == null) {
			entry = {
				module: module,
				kind: moduleKind(moduleType),
				location: locationLabel(module, modulePosition(moduleType)),
				typeUsages: [],
				memberUsages: [],
				nativeImports: []
			};
			modules.set(module, entry);
		}

		collectTypeUsage(entry, usage);
		collectMemberUsage(entry, moduleType);
	}

	public function snapshot(capabilityReasons:Array<GoHxrtFeatureReason>, runtimeImportPath:String):GoTypeUsageLedgerSnapshot {
		var moduleEntries = new Array<GoTypeUsageModuleEvidence>();
		var typeUsageCount = 0;
		var memberUsageCount = 0;
		var nativeImportCount = 0;
		var moduleNames = [for (name in modules.keys()) name];
		moduleNames.sort(compareStrings);

		for (name in moduleNames) {
			var mutable = modules.get(name);
			if (mutable == null) {
				continue;
			}
			var typeUsages = [for (value in mutable.typeUsages) value];
			typeUsages.sort(compareTypeUsage);
			var memberUsages = [for (value in mutable.memberUsages) value];
			memberUsages.sort(compareMemberUsage);
			var mutableNativeImports = [for (value in mutable.nativeImports) value];
			mutableNativeImports.sort(compareMutableNativeImport);
			var nativeImports = [
				for (value in mutableNativeImports)
					{
						target: value.target,
						metadataImportPath: value.metadataImportPath,
						resolvedImportPath: resolveImportPath(value.metadataImportPath, runtimeImportPath),
						usageKind: value.usageKind,
						location: value.location
					}
			];
			if (typeUsages.length == 0 && memberUsages.length == 0 && nativeImports.length == 0) {
				continue;
			}
			typeUsageCount += typeUsages.length;
			memberUsageCount += memberUsages.length;
			nativeImportCount += nativeImports.length;
			moduleEntries.push({
				module: mutable.module,
				kind: mutable.kind,
				location: mutable.location,
				typeUsages: GoImmutableList.fromArray(typeUsages),
				memberUsages: GoImmutableList.fromArray(memberUsages),
				nativeImports: GoImmutableList.fromArray(nativeImports)
			});
		}

		var capabilitiesByKey = new Map<String, GoCapabilityUsageEvidence>();
		for (reason in capabilityReasons) {
			if (reason == null) {
				continue;
			}
			var evidence:GoCapabilityUsageEvidence = {
				feature: reason.feature,
				sourceKind: reason.sourceKind,
				source: reason.source
			};
			capabilitiesByKey.set(capabilityKey(evidence), evidence);
		}
		var capabilities = [for (value in capabilitiesByKey) value];
		capabilities.sort(compareCapabilities);

		return {
			schemaVersion: SCHEMA_VERSION,
			source: SOURCE,
			scannerFallback: SCANNER_FALLBACK,
			moduleCount: moduleEntries.length,
			typeUsageCount: typeUsageCount,
			memberUsageCount: memberUsageCount,
			nativeImportCount: nativeImportCount,
			capabilityCount: capabilities.length,
			modules: GoImmutableList.fromArray(moduleEntries),
			capabilities: GoImmutableList.fromArray(capabilities)
		};
	}

	public static function emptySnapshot():GoTypeUsageLedgerSnapshot {
		return {
			schemaVersion: SCHEMA_VERSION,
			source: SOURCE,
			scannerFallback: SCANNER_FALLBACK,
			moduleCount: 0,
			typeUsageCount: 0,
			memberUsageCount: 0,
			nativeImportCount: 0,
			capabilityCount: 0,
			modules: GoImmutableList.fromArray([]),
			capabilities: GoImmutableList.fromArray([])
		};
	}

	public static function renderJson(snapshot:GoTypeUsageLedgerSnapshot):String {
		var lines = [
			"{",
			'\t"schemaVersion": ' + snapshot.schemaVersion + ",",
			'\t"source": "' + jsonEscape(snapshot.source) + '",',
			'\t"scannerFallback": "' + jsonEscape(snapshot.scannerFallback) + '",',
			'\t"moduleCount": ' + snapshot.moduleCount + ",",
			'\t"typeUsageCount": ' + snapshot.typeUsageCount + ",",
			'\t"memberUsageCount": ' + snapshot.memberUsageCount + ",",
			'\t"nativeImportCount": ' + snapshot.nativeImportCount + ",",
			'\t"capabilityCount": ' + snapshot.capabilityCount + ",",
			'\t"modules": ['
		];

		for (moduleIndex in 0...snapshot.modules.length) {
			var module = snapshot.modules.at(moduleIndex);
			lines.push("\t\t{");
			lines.push('\t\t\t"module": "' + jsonEscape(module.module) + '",');
			lines.push('\t\t\t"kind": "' + jsonEscape(module.kind) + '",');
			lines.push('\t\t\t"location": "' + jsonEscape(module.location) + '",');
			lines.push('\t\t\t"typeUsages": [');
			for (index in 0...module.typeUsages.length) {
				var usage = module.typeUsages.at(index);
				lines.push("\t\t\t\t{");
				lines.push('\t\t\t\t\t"level": "' + jsonEscape(usage.level) + '",');
				lines.push('\t\t\t\t\t"shape": ' + renderShapeJson(usage.shape));
				lines.push("\t\t\t\t}" + (index + 1 < module.typeUsages.length ? "," : ""));
			}
			lines.push("\t\t\t],");
			lines.push('\t\t\t"memberUsages": [');
			for (index in 0...module.memberUsages.length) {
				var usage = module.memberUsages.at(index);
				lines.push("\t\t\t\t{");
				lines.push('\t\t\t\t\t"kind": "' + jsonEscape(usage.kind) + '",');
				lines.push('\t\t\t\t\t"target": "' + jsonEscape(usage.target) + '",');
				lines.push('\t\t\t\t\t"member": "' + jsonEscape(usage.member) + '",');
				lines.push('\t\t\t\t\t"location": "' + jsonEscape(usage.location) + '"');
				lines.push("\t\t\t\t}" + (index + 1 < module.memberUsages.length ? "," : ""));
			}
			lines.push("\t\t\t],");
			lines.push('\t\t\t"nativeImports": [');
			for (index in 0...module.nativeImports.length) {
				var usage = module.nativeImports.at(index);
				lines.push("\t\t\t\t{");
				lines.push('\t\t\t\t\t"target": "' + jsonEscape(usage.target) + '",');
				lines.push('\t\t\t\t\t"metadataImportPath": "' + jsonEscape(usage.metadataImportPath) + '",');
				lines.push('\t\t\t\t\t"resolvedImportPath": "' + jsonEscape(usage.resolvedImportPath) + '",');
				lines.push('\t\t\t\t\t"usageKind": "' + jsonEscape(usage.usageKind) + '",');
				lines.push('\t\t\t\t\t"location": "' + jsonEscape(usage.location) + '"');
				lines.push("\t\t\t\t}" + (index + 1 < module.nativeImports.length ? "," : ""));
			}
			lines.push("\t\t\t]");
			lines.push("\t\t}" + (moduleIndex + 1 < snapshot.modules.length ? "," : ""));
		}
		lines.push("\t],");
		lines.push('\t"capabilities": [');
		for (index in 0...snapshot.capabilities.length) {
			var capability = snapshot.capabilities.at(index);
			lines.push("\t\t{");
			lines.push('\t\t\t"feature": "' + jsonEscape(capability.feature) + '",');
			lines.push('\t\t\t"sourceKind": "' + jsonEscape(capability.sourceKind) + '",');
			lines.push('\t\t\t"source": "' + jsonEscape(capability.source) + '"');
			lines.push("\t\t}" + (index + 1 < snapshot.capabilities.length ? "," : ""));
		}
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	/**
		Canonical JSON identity for one typed shape.

		Why / What / How
		- Registry and planner reports must identify the same used type without
		  maintaining independent encoders.
		- The result contains only the stable macro-object-free shape algebra.
	**/
	public static function renderShapeJson(shape:GoTypeShape):String {
		return switch (shape) {
			case Nominal(kind, path, parameters):
				shapeObjectJson(kind, path, shapeListJson(parameters), "[]", "null", "[]");
			case TypeParameter(path):
				shapeObjectJson("type_parameter", path, "[]", "[]", "null", "[]");
			case Function(arguments, returnType):
				var argumentJson = new Array<String>();
				for (argument in arguments) {
					argumentJson.push('{"name":"' + jsonEscape(argument.name) + '","optional":' + argument.optional + ',"shape":'
						+ renderShapeJson(argument.shape) + "}");
				}
				shapeObjectJson("function", "", "[]", "[" + argumentJson.join(",") + "]", renderShapeJson(returnType), "[]");
			case Anonymous(fields):
				var fieldJson = new Array<String>();
				for (field in fields) {
					fieldJson.push('{"name":"'
						+ jsonEscape(field.name)
						+ '","optional":'
						+ field.optional
						+ ',"shape":'
						+ renderShapeJson(field.shape)
						+ "}");
				}
				shapeObjectJson("anonymous", "", "[]", "[]", "null", "[" + fieldJson.join(",") + "]");
			case DynamicShape(inner):
				shapeObjectJson("dynamic", "", "[]", "[]", inner == null ? "null" : renderShapeJson(inner), "[]");
			case UnknownShape(reason):
				shapeObjectJson("unknown", reason, "[]", "[]", "null", "[]");
		};
	}

	static function shapeObjectJson(kind:String, path:String, parameters:String, arguments:String, returnType:String, fields:String):String {
		return '{"kind":"' + jsonEscape(kind) + '","path":"' + jsonEscape(path) + '","parameters":' + parameters + ',"arguments":' + arguments
			+ ',"returnType":' + returnType + ',"fields":' + fields + "}";
	}

	static function shapeListJson(shapes:GoImmutableList<GoTypeShape>):String {
		var values = new Array<String>();
		for (shape in shapes) {
			values.push(renderShapeJson(shape));
		}
		return "[" + values.join(",") + "]";
	}

	public static function renderMarkdown(snapshot:GoTypeUsageLedgerSnapshot):String {
		var lines = [
			"# Typed Usage Ledger",
			"",
			"- Schema: `" + snapshot.schemaVersion + "`",
			"- Source: `" + snapshot.source + "`",
			"- Scanner fallback: `" + snapshot.scannerFallback + "`",
			"- Modules: `" + snapshot.moduleCount + "`",
			"- Type usages: `" + snapshot.typeUsageCount + "`",
			"- Member usages: `" + snapshot.memberUsageCount + "`",
			"- Native imports: `" + snapshot.nativeImportCount + "`",
			"- Runtime capability reasons: `" + snapshot.capabilityCount + "`",
			"",
			"## Modules",
			""
		];
		for (module in snapshot.modules) {
			lines.push("### `" + module.module + "`");
			lines.push("");
			lines.push("- Kind: `" + module.kind + "`");
			lines.push("- Location: `" + module.location + "`");
			lines.push("- Type usages: `" + module.typeUsages.length + "`");
			lines.push("- Member usages: `" + module.memberUsages.length + "`");
			lines.push("- Native imports: `" + module.nativeImports.length + "`");
			lines.push("");
		}
		lines.push("## Runtime capabilities");
		lines.push("");
		if (snapshot.capabilities.length == 0) {
			lines.push("- None.");
		} else {
			for (capability in snapshot.capabilities) {
				lines.push("- `" + capability.feature + "` from `" + capability.sourceKind + ":" + capability.source + "`");
			}
		}
		return lines.join("\n") + "\n";
	}

	function collectTypeUsage(module:MutableModuleEvidence, usage:Null<TypeUsageMap>):Void {
		if (usage == null) {
			return;
		}
		for (level in usage.keys()) {
			var values = usage.get(level);
			if (values == null) {
				continue;
			}
			var outwardLevel = GoTypeUsageLevelId.fromReflaxe(level);
			for (value in values) {
				var evidence = typeUsageEvidence(value, outwardLevel);
				if (evidence == null) {
					continue;
				}
				module.typeUsages.set(typeUsageKey(evidence), evidence);
				var targetModule = moduleTypeForUsage(value);
				if (targetModule != null) {
					addNativeImport(module, targetModule, GoNativeImportUsageKind.fromTypeUsage(outwardLevel), module.location);
				}
			}
		}
	}

	function collectMemberUsage(module:MutableModuleEvidence, owner:ModuleType):Void {
		switch (owner) {
			case TClassDecl(classRef):
				var classType = classRef.get();
				if (classType.superClass != null) {
					addTypeShape(module, GoTypeUsageLevelId.ExtendedFrom, TInst(classType.superClass.t, classType.superClass.params));
				}
				for (implemented in classType.interfaces) {
					addTypeShape(module, GoTypeUsageLevelId.ExtendedFrom, TInst(implemented.t, implemented.params));
				}
				if (classType.constructor != null) {
					collectClassFieldExpr(module, classType.constructor.get());
				}
				for (field in classType.fields.get()) {
					collectClassFieldExpr(module, field);
				}
				for (field in classType.statics.get()) {
					collectClassFieldExpr(module, field);
				}
			case TEnumDecl(_) | TTypeDecl(_) | TAbstract(_):
		}
	}

	function collectClassFieldExpr(module:MutableModuleEvidence, field:ClassField):Void {
		switch (field.kind) {
			case FVar(_, _):
				addTypeShape(module, GoTypeUsageLevelId.VariableDeclaration, field.type);
			case FMethod(_):
				addTypeShape(module, GoTypeUsageLevelId.FunctionDeclaration, field.type);
		}
		var expression = field.expr();
		if (expression == null) {
			return;
		}
		function visit(expr:TypedExpr):Void {
			switch (expr.expr) {
				case TNew(classRef, parameters, _):
					addTypeShape(module, GoTypeUsageLevelId.Constructed, TInst(classRef, parameters));
					var target:ModuleType = TClassDecl(classRef);
					addMemberUsage(module, GoMemberUsageKind.Construct, target.getPath(), "new", target, expr.pos);
				case TVar(variable, _):
					addTypeShape(module, GoTypeUsageLevelId.VariableType, variable.t);
				case TField(targetExpr, access):
					addTypeShape(module, GoTypeUsageLevelId.Expression, expr.t);
					var resolved = resolveMember(access, targetExpr.t);
					if (resolved != null) {
						addMemberUsage(module, GoMemberUsageKind.FieldAccess, resolved.target, resolved.member, resolved.moduleType, expr.pos);
					}
				case TCall(callee, _):
					addTypeShape(module, GoTypeUsageLevelId.Expression, expr.t);
					var resolved = resolveCallee(callee);
					if (resolved != null) {
						addMemberUsage(module, GoMemberUsageKind.Call, resolved.target, resolved.member, resolved.moduleType, expr.pos);
					}
				case _:
					addTypeShape(module, GoTypeUsageLevelId.Expression, expr.t);
			}
			TypedExprTools.iter(expr, visit);
		}
		visit(expression);
	}

	function addTypeShape(module:MutableModuleEvidence, level:GoTypeUsageLevelId, type:Type):Void {
		if (type == null) {
			return;
		}
		var evidence:GoTypeUsageEvidence = {
			level: level,
			shape: typeShape(type)
		};
		module.typeUsages.set(typeUsageKey(evidence), evidence);
	}

	function addMemberUsage(module:MutableModuleEvidence, kind:GoMemberUsageKind, target:String, member:String, targetModule:Null<ModuleType>,
			pos:haxe.macro.Expr.Position):Void {
		var location = locationLabel(module.module, pos);
		var evidence:GoMemberUsageEvidence = {
			kind: kind,
			target: target,
			member: member,
			location: location
		};
		module.memberUsages.set(memberUsageKey(evidence), evidence);
		if (targetModule != null) {
			addNativeImport(module, targetModule, GoNativeImportUsageKind.fromMemberUsage(kind), location);
		}
	}

	function addNativeImport(module:MutableModuleEvidence, target:ModuleType, usageKind:GoNativeImportUsageKind, location:String):Void {
		var importPath = goImportPath(target);
		if (importPath == null) {
			return;
		}
		var evidence:MutableNativeImportEvidence = {
			target: target.getPath(),
			metadataImportPath: importPath,
			usageKind: usageKind,
			location: location
		};
		module.nativeImports.set(nativeImportKey(evidence), evidence);
	}

	static function typeUsageEvidence(value:TypeOrModuleType, level:GoTypeUsageLevelId):Null<GoTypeUsageEvidence> {
		return switch (value) {
			case EModuleType(moduleType):
				{
					level: level,
					shape: moduleTypeShape(moduleType)
				};
			case EType(type):
				{
					level: level,
					shape: typeShape(type)
				};
		};
	}

	static function typeShape(type:Type, ?trail:Array<String>, depth:Int = 0):GoTypeShape {
		if (type == null) {
			return UnknownShape(GoUnknownTypeShapeReason.Missing);
		}
		if (depth >= 64) {
			return UnknownShape(GoUnknownTypeShapeReason.DepthLimit);
		}
		var activeTrail = trail == null ? [] : trail;
		var id = TypeHelper.getUniqueId(type);
		if (activeTrail.indexOf(id) != -1) {
			return UnknownShape(GoUnknownTypeShapeReason.Recursive);
		}
		var nextTrail = activeTrail.copy();
		nextTrail.push(id);

		return switch (type) {
			case TMono(reference):
				var resolved = reference.get();
				resolved == null ? UnknownShape(GoUnknownTypeShapeReason.Monomorph) : typeShape(resolved, nextTrail, depth + 1);
			case TEnum(enumRef, parameters):
				nominalShape(GoTypeUsageTargetKind.Enum, TEnumDecl(enumRef).getPath(), parameters, activeTrail, depth);
			case TInst(classRef, parameters):
				var classType = classRef.get();
				switch (classType.kind) {
					case KTypeParameter(_):
						TypeParameter(TClassDecl(classRef).getPath());
					case _:
						nominalShape(GoTypeUsageTargetKind.Class, TClassDecl(classRef).getPath(), parameters, activeTrail, depth);
				}
			case TType(typeRef, parameters):
				final path = TTypeDecl(typeRef).getPath();
				if (path == "StdTypes.Iterator" && parameters.length == 1) {
					/*
						Why: Haxe Iterator<T> is the canonical structural
						hasNext/next protocol, but retaining only its typedef name
						hides those fields from exact registry matching.
						What: Follow this one compiler-known standard typedef into
						its anonymous shape while preserving the applied T.
						How: All user typedefs remain nominal and opaque, so this
						does not let aliases hide Dynamic or other storage.
					 */
					typeShape(Context.follow(type), nextTrail, depth + 1);
				} else {
					nominalShape(GoTypeUsageTargetKind.Typedef, path, parameters, activeTrail, depth);
				}
			case TFun(arguments, returnType):
				var outwardArguments = new Array<GoFunctionArgumentShape>();
				for (argument in arguments) {
					outwardArguments.push({
						name: argument.name,
						optional: argument.opt,
						shape: typeShape(argument.t, nextTrail, depth + 1)
					});
				}
				Function(GoImmutableList.fromArray(outwardArguments), typeShape(returnType, nextTrail, depth + 1));
			case TAnonymous(anonymousRef):
				var sourceFields = anonymousRef.get().fields.copy();
				sourceFields.sort((a, b) -> compareStrings(a.name, b.name));
				var outwardFields = new Array<GoAnonymousFieldShape>();
				for (field in sourceFields) {
					outwardFields.push({
						name: field.name,
						optional: field.meta != null && field.meta.has(":optional"),
						shape: typeShape(field.type, nextTrail, depth + 1)
					});
				}
				Anonymous(GoImmutableList.fromArray(outwardFields));
			case TDynamic(inner):
				DynamicShape(inner == null ? null : typeShape(inner, nextTrail, depth + 1));
			case TLazy(resolve):
				typeShape(resolve(), nextTrail, depth + 1);
			case TAbstract(abstractRef, parameters):
				var moduleType:ModuleType = TAbstract(abstractRef);
				nominalShape(GoTypeUsageTargetKind.Abstract, moduleType.getPath(), parameters, activeTrail, depth);
		};
	}

	static function nominalShape(kind:GoTypeUsageTargetKind, path:String, parameters:Array<Type>, trail:Array<String>, depth:Int):GoTypeShape {
		var outwardParameters = new Array<GoTypeShape>();
		for (parameter in parameters) {
			outwardParameters.push(typeShape(parameter, trail, depth + 1));
		}
		return Nominal(kind, path, GoImmutableList.fromArray(outwardParameters));
	}

	static function moduleTypeShape(moduleType:ModuleType):GoTypeShape {
		return Nominal(moduleKind(moduleType), moduleType.getPath(), GoImmutableList.fromArray([]));
	}

	static function moduleTypeForUsage(value:TypeOrModuleType):Null<ModuleType> {
		return switch (value) {
			case EModuleType(moduleType):
				moduleType;
			case EType(type):
				TypeHelper.toModuleType(type);
		};
	}

	static function resolveCallee(expr:TypedExpr):Null<ResolvedMember> {
		return switch (expr.expr) {
			case TField(target, access):
				resolveMember(access, target.t);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				resolveCallee(inner);
			case _:
				null;
		};
	}

	static function resolveMember(access:FieldAccess, targetType:Type):Null<ResolvedMember> {
		return switch (access) {
			case FInstance(classRef, _, fieldRef):
				var moduleType:ModuleType = TClassDecl(classRef);
				{target: moduleType.getPath(), member: fieldRef.get().name, moduleType: moduleType};
			case FStatic(classRef, fieldRef):
				var moduleType:ModuleType = TClassDecl(classRef);
				{target: moduleType.getPath(), member: fieldRef.get().name, moduleType: moduleType};
			case FEnum(enumRef, enumField):
				var moduleType:ModuleType = TEnumDecl(enumRef);
				{target: moduleType.getPath(), member: enumField.name, moduleType: moduleType};
			case FClosure(classRef, fieldRef):
				var moduleType:Null<ModuleType> = classRef == null ? TypeHelper.toModuleType(targetType) : TClassDecl(classRef.c);
				{
					target: moduleType == null ? "<function>" : moduleType.getPath(),
					member: fieldRef.get().name,
					moduleType: moduleType
				};
			case FAnon(fieldRef):
				{target: "<anonymous>", member: fieldRef.get().name, moduleType: null};
			case FDynamic(name):
				var moduleType = TypeHelper.toModuleType(targetType);
				{target: moduleType == null ? "<dynamic>" : moduleType.getPath(), member: name, moduleType: moduleType};
		};
	}

	static function goImportPath(moduleType:ModuleType):Null<String> {
		var metadata = switch (moduleType) {
			case TClassDecl(classRef):
				classRef.get().meta;
			case TEnumDecl(enumRef):
				enumRef.get().meta;
			case TTypeDecl(typeRef):
				typeRef.get().meta;
			case TAbstract(abstractRef):
				abstractRef.get().meta;
		};
		if (metadata == null) {
			return null;
		}
		for (entry in metadata.extract(":go.import")) {
			if (entry.params == null || entry.params.length != 1) {
				continue;
			}
			switch (entry.params[0].expr) {
				case EConst(CString(value, _)) if (value != null && StringTools.trim(value) != ""):
					return StringTools.trim(value);
				case _:
			}
		}
		return null;
	}

	static function isProjectOwner(moduleType:ModuleType):Bool {
		var module = moduleName(moduleType);
		if (module == ""
			|| module == "StdTypes"
			|| StringTools.startsWith(module, "haxe.")
			|| StringTools.startsWith(module, "sys.")
			|| StringTools.startsWith(module, "go.")
			|| StringTools.startsWith(module, "hxrt.")
			|| StringTools.startsWith(module, "reflaxe.")) {
			return false;
		}
		var file = normalizePath(Context.getPosInfos(modulePosition(moduleType)).file);
		return file.indexOf("/std/") == -1 && file.indexOf("/vendor/") == -1 && file.indexOf("/src/reflaxe/") == -1;
	}

	static function moduleName(moduleType:ModuleType):String {
		return moduleType.getPath();
	}

	static function modulePosition(moduleType:ModuleType):haxe.macro.Expr.Position {
		return switch (moduleType) {
			case TClassDecl(classRef):
				classRef.get().pos;
			case TEnumDecl(enumRef):
				enumRef.get().pos;
			case TTypeDecl(typeRef):
				typeRef.get().pos;
			case TAbstract(abstractRef):
				abstractRef.get().pos;
		};
	}

	static function moduleKind(moduleType:ModuleType):GoTypeUsageTargetKind {
		return switch (moduleType) {
			case TClassDecl(_): GoTypeUsageTargetKind.Class;
			case TEnumDecl(_): GoTypeUsageTargetKind.Enum;
			case TTypeDecl(_): GoTypeUsageTargetKind.Typedef;
			case TAbstract(_): GoTypeUsageTargetKind.Abstract;
		};
	}

	static function locationLabel(module:String, pos:haxe.macro.Expr.Position):String {
		var line = 1;
		var location = PositionTools.toLocation(pos);
		if (location != null && location.range != null && location.range.start != null && location.range.start.line > 0) {
			line = location.range.start.line;
		}
		return module + ":" + line;
	}

	static function typeUsageKey(entry:GoTypeUsageEvidence):String {
		return entry.level + "\u0000" + shapeKey(entry.shape);
	}

	static function memberUsageKey(entry:GoMemberUsageEvidence):String {
		return entry.kind + "\u0000" + entry.target + "\u0000" + entry.member + "\u0000" + entry.location;
	}

	static function nativeImportKey(entry:MutableNativeImportEvidence):String {
		return entry.target + "\u0000" + entry.metadataImportPath + "\u0000" + entry.usageKind + "\u0000" + entry.location;
	}

	static function capabilityKey(entry:GoCapabilityUsageEvidence):String {
		return entry.feature + "\u0000" + entry.sourceKind + "\u0000" + entry.source;
	}

	static function compareTypeUsage(a:GoTypeUsageEvidence, b:GoTypeUsageEvidence):Int {
		var levelOrder = compareStrings(a.level, b.level);
		return levelOrder != 0 ? levelOrder : compareStrings(shapeKey(a.shape), shapeKey(b.shape));
	}

	static function compareMemberUsage(a:GoMemberUsageEvidence, b:GoMemberUsageEvidence):Int {
		var locationOrder = compareStrings(a.location, b.location);
		if (locationOrder != 0) {
			return locationOrder;
		}
		var kindOrder = compareStrings(a.kind, b.kind);
		if (kindOrder != 0) {
			return kindOrder;
		}
		var targetOrder = compareStrings(a.target, b.target);
		return targetOrder != 0 ? targetOrder : compareStrings(a.member, b.member);
	}

	static function compareMutableNativeImport(a:MutableNativeImportEvidence, b:MutableNativeImportEvidence):Int {
		var targetOrder = compareStrings(a.target, b.target);
		if (targetOrder != 0) {
			return targetOrder;
		}
		var importOrder = compareStrings(a.metadataImportPath, b.metadataImportPath);
		if (importOrder != 0) {
			return importOrder;
		}
		var usageOrder = compareStrings(a.usageKind, b.usageKind);
		return usageOrder != 0 ? usageOrder : compareStrings(a.location, b.location);
	}

	static function shapeKey(shape:GoTypeShape):String {
		return renderShapeJson(shape);
	}

	static function resolveImportPath(metadataImportPath:String, runtimeImportPath:String):String {
		if (metadataImportPath == "hxrt") {
			return runtimeImportPath == null || StringTools.trim(runtimeImportPath) == "" ? "hxrt" : StringTools.trim(runtimeImportPath);
		}
		return metadataImportPath;
	}

	static function compareCapabilities(a:GoCapabilityUsageEvidence, b:GoCapabilityUsageEvidence):Int {
		var featureOrder = compareStrings(a.feature, b.feature);
		if (featureOrder != 0) {
			return featureOrder;
		}
		var kindOrder = compareStrings(a.sourceKind, b.sourceKind);
		return kindOrder != 0 ? kindOrder : compareStrings(a.source, b.source);
	}

	static inline function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function normalizePath(path:String):String {
		return path == null ? "" : path.split("\\").join("/");
	}

	static function jsonEscape(value:String):String {
		if (value == null) {
			return "";
		}
		return value.split("\\")
			.join("\\\\")
			.split('"')
			.join('\\"')
			.split("\n")
			.join("\\n")
			.split("\r")
			.join("\\r")
			.split("\t")
			.join("\\t");
	}
}
#end
