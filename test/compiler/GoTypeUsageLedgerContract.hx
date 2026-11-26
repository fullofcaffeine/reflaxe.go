#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.compiler.GoTypeUsageLedger;
import reflaxe.go.compiler.GoTypeUsageTracker;
import reflaxe.helpers.TypeHelper;

using reflaxe.helpers.ModuleTypeHelper;
#end

/**
	Contract for how the type-usage ledger names member targets and native
	imports.

	Why
	The ledger reads a class once for both its path and its `@:go.import`
	metadata, keeps the path and import of each enum, typedef, and abstract by
	type path, and names a referenced field by printing its reference instead of
	decoding it. These shortcuts must name exactly the targets, members, and
	imports of the plain reads. Otherwise the member and native-import evidence
	of the type-usage report, and the runtime and registry decisions built on
	it, would change.

	What
	Proves, with one shared ledger and in both orders:
	- for every module type of the two fixture modules and some standard
	  library modules, `declarationTarget(...)` gives `ModuleTypeHelper.getPath(...)`
	  and `plainImportPath(...)`;
	- for every field access, call, and construction in the fixture bodies,
	  `resolveMember(...)` and `resolveCallee(...)` give the target, member, and
	  declaration of `plainResolveMember(...)`; the fixtures reach every kind of
	  field access;
	- the ledger kept targets only for enums, typedefs, and abstracts, and kept
	  the same-named private declarations of the two modules apart;
	- a report collected from the fixtures records the native imports of a class,
	  an enum, an abstract, and a typedef, with the trimmed import path and the
	  first valid `@:go.import` entry.

	How
	After the initialization macros, loads the fixture modules and compares each
	answer with the plain functions below, which are the ledger's reads as they
	were before it kept declaration targets.
**/
@:access(reflaxe.go.compiler.GoTypeUsageLedger)
class GoTypeUsageLedgerContract {
	#if macro
	static final FIXTURE_MODULES = ["GoTypeUsageLedgerContractFixture", "GoTypeUsageLedgerContractOtherFixture"];
	static final STD_MODULES = ["haxe.ds.Option", "haxe.ds.Vector", "haxe.Int64", "haxe.io.Bytes"];

	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoTypeUsageLedgerContract: " + message, Context.currentPos());
	}

	static function moduleTypes(module:String):Array<ModuleType> {
		final result:Array<ModuleType> = [];
		for (type in Context.getModule(module)) {
			final moduleType = TypeHelper.toModuleType(type);
			require(moduleType != null, "module type of " + Std.string(type));
			if (moduleType != null)
				result.push(moduleType);
		}
		return result;
	}

	static function bodies(moduleType:ModuleType):Array<TypedExpr> {
		final result:Array<TypedExpr> = [];
		switch (moduleType) {
			case TClassDecl(classRef):
				final classType = classRef.get();
				final fields = classType.fields.get().concat(classType.statics.get());
				if (classType.constructor != null)
					fields.push(classType.constructor.get());
				for (field in fields) {
					final body = field.expr();
					if (body != null)
						result.push(body);
				}
			case _:
		}
		return result;
	}

	static function label(target:String, member:String, path:Null<String>, importPath:Null<String>):String {
		return target + " " + member + " " + path + " " + importPath;
	}

	static function run():Void {
		final fixtureTypes:Array<ModuleType> = [];
		for (module in FIXTURE_MODULES)
			for (moduleType in moduleTypes(module))
				fixtureTypes.push(moduleType);
		final all = fixtureTypes.copy();
		for (module in STD_MODULES)
			for (moduleType in moduleTypes(module))
				all.push(moduleType);

		// Type every fixture body before the comparison, as Haxe has done before the compiler runs.
		for (moduleType in fixtureTypes)
			bodies(moduleType);

		final ledger = new GoTypeUsageLedger();
		final accessKinds = new Map<String, Bool>();
		var castCalls = 0;
		for (round in 0...2) {
			final ordered = round == 0 ? all : [for (index in 0...all.length) all[all.length - 1 - index]];
			for (moduleType in ordered) {
				final target = ledger.declarationTarget(moduleType);
				require(target.path == moduleType.getPath(), "path of " + moduleType.getPath() + ": " + target.path);
				require(target.importPath == plainImportPath(moduleType), "import path of " + moduleType.getPath() + ": " + target.importPath);
			}
			for (moduleType in fixtureTypes)
				for (body in bodies(moduleType)) {
					function visit(expr:TypedExpr):Void {
						switch (expr.expr) {
							case TField(targetExpr, access):
								accessKinds.set(accessKind(access), true);
								final actual = ledger.resolveMember(access, targetExpr.t);
								final expected = plainResolveMember(access, targetExpr.t);
								final actualDeclaration = actual == null ? null : actual.declaration;
								final actualLabel = actual == null ? "null" : label(actual.target, actual.member,
									actualDeclaration == null ? null : actualDeclaration.path, actualDeclaration == null ? null : actualDeclaration.importPath);
								require(actualLabel == expected, "member of " + TypedExprTools.toString(expr) + ": " + actualLabel + " instead of " + expected);
							case TCall(callee, _):
								switch (callee.expr) {
									case TParenthesis(_) | TCast(_, _) | TMeta(_, _):
										castCalls++;
									case _:
								}
								final actual = ledger.resolveCallee(callee);
								final expected = plainResolveCallee(callee);
								final actualDeclaration = actual == null ? null : actual.declaration;
								final actualLabel = actual == null ? "null" : label(actual.target, actual.member,
									actualDeclaration == null ? null : actualDeclaration.path, actualDeclaration == null ? null : actualDeclaration.importPath);
								require(actualLabel == expected, "callee of " + TypedExprTools.toString(expr) + ": " + actualLabel + " instead of " + expected);
							case TNew(classRef, _, _):
								final moduleType:ModuleType = TClassDecl(classRef);
								final actual = ledger.declarationTarget(moduleType);
								require(actual.path == moduleType.getPath() && actual.importPath == plainImportPath(moduleType),
									"constructed " + moduleType.getPath());
							case _:
						}
						TypedExprTools.iter(expr, visit);
					}
					visit(body);
				}
		}
		for (kind in ["instance", "static", "enum", "closure", "anonymous", "dynamic"])
			require(accessKinds.exists(kind), "fixture reaches a " + kind + " field access");
		require(castCalls > 0, "fixture calls through a cast");

		// Targets are kept only for enums, typedefs, and abstracts, one per declaration.
		final keys = [for (key in ledger.declarationTargetsByPath.keys()) key];
		require(keys.length > 5, "the ledger kept declaration targets (" + keys.length + ")");
		for (key in keys)
			require(["E", "T", "A"].indexOf(key.charAt(0)) != -1, "kept the target of " + key);
		for (name in ["EHidden", "AScore"]) {
			final kind = name.charAt(0);
			final suffix = "." + name.substr(1);
			final imports = [
				for (key in keys)
					if (key.charAt(0) == kind && StringTools.endsWith(key, suffix)) ledger.declarationTargetsByPath.get(key).importPath
			];
			imports.sort((a, b) -> a < b ? -1 : (a > b ? 1 : 0));
			final base = name == "EHidden" ? "hidden" : "score";
			require(imports.join(",") == "example.com/first/" + base + ",example.com/other/" + base,
				"private "
				+ name.substr(1)
				+ " of both modules: "
				+ imports);
		}

		// The report records native imports of every declaration kind.
		final tracker = new GoTypeUsageTracker();
		final collector = new GoTypeUsageLedger();
		for (moduleType in fixtureTypes)
			collector.collect(moduleType, tracker.track(moduleType), tracker.classFields);
		final report = GoTypeUsageLedger.renderJson(collector.snapshot([], "hxrt"));
		for (expected in [
			["GoTypeUsageLedgerContractNativeBox", "example.com/native/box", "construct"],
			["GoTypeUsageLedgerContractNativeBox", "example.com/native/box", "call"],
			[
				"GoTypeUsageLedgerContractNativeChoice",
				"example.com/native/choice",
				"field_access"
			],
			[
				"GoTypeUsageLedgerContractNativeCount",
				"example.com/native/count",
				"variable_declaration"
			],
			[
				"GoTypeUsageLedgerContractNativeShape",
				"example.com/native/shape",
				"variable_declaration"
			],
			["GoTypeUsageLedgerContractLateImport", "example.com/native/late", "call"]
		]) {
			final entry = '"target": "GoTypeUsageLedgerContractFixture.' + expected[0] + '",\n\t\t\t\t\t"metadataImportPath": "' + expected[1]
				+ '",\n\t\t\t\t\t"resolvedImportPath": "' + expected[1] + '",\n\t\t\t\t\t"usageKind": "' + expected[2] + '"';
			require(report.indexOf(entry) != -1, "native import " + expected.join(" ") + " in the report");
		}

		Sys.println("type usage ledger contract passed");
	}

	static function accessKind(access:FieldAccess):String {
		return switch (access) {
			case FInstance(_, _, _): "instance";
			case FStatic(_, _): "static";
			case FEnum(_, _): "enum";
			case FClosure(_, _): "closure";
			case FAnon(_): "anonymous";
			case FDynamic(_): "dynamic";
		};
	}

	/**
		The ledger's native import read before it kept declaration targets, kept
		as the oracle for this contract. Do not optimize it.
	**/
	static function plainImportPath(moduleType:ModuleType):Null<String> {
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

	/** The ledger's callee resolution before it kept declaration targets. Do not optimize it. */
	static function plainResolveCallee(expr:TypedExpr):String {
		return switch (expr.expr) {
			case TField(target, access):
				plainResolveMember(access, target.t);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				plainResolveCallee(inner);
			case _:
				"null";
		};
	}

	/** The ledger's member resolution before it kept declaration targets. Do not optimize it. */
	static function plainResolveMember(access:FieldAccess, targetType:Type):String {
		final resolved:{target:String, member:String, moduleType:Null<ModuleType>} = switch (access) {
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
		final moduleType = resolved.moduleType;
		return label(resolved.target, resolved.member, moduleType == null ? null : moduleType.getPath(),
			moduleType == null ? null : plainImportPath(moduleType));
	}
	#end
}
