#if macro
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.analyze.GoNativeTypeReach;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.GoPortableNativeImportHit;
import reflaxe.go.compiler.GoTypeMapper;
#end

/**
	Contract for the reach check in front of the native authority gate's type walk.

	Why
	The gate (`GoProfileContractAnalyzer`) asks `GoNativeTypeReach` first whether
	a type can reach a `go.*` declaration, and runs its own walk only when it can.
	The reach check remembers answers per declaration, so a remembered answer that
	is wrong for a later question would hide a `go.*` use from the gate. And the
	gate's walk can miss a `go.*` type that the reach check sees, so the gate must
	keep the walk's answer: its diagnostics must stay exactly as before.

	What
	For the field types and expression node types of real typed fixture classes,
	and for constructed lazy, monomorph, and deeply nested types, pins that:
	- the reach answer equals an independent reachability walk with no remembered
	  answers, for several question orders and for each question alone;
	- named fixture fields have the expected answers, including cycles through
	  typedefs and statics types;
	- every question the replaced walk answers true is also true here;
	- the gate answers exactly as a copy of the replaced walk, per type, per class
	  and abstract, and in the hit list, including the anonymous structure case
	  that the walk has always missed.

	How
	After the initialization macros, loads the fixture modules, collects the
	questions, and compares the answers with a copy of the replaced walk and with
	the reference walk.
**/
@:access(reflaxe.go.analyze.GoProfileContractAnalyzer)
class GoNativeTypeReachContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoNativeTypeReachContract: " + message, Context.currentPos());
	}

	static function load(path:String):ClassType {
		return switch (Context.getType(path)) {
			case TInst(ref, _): ref.get();
			case _: throw "not a class: " + path;
		};
	}

	/** The replaced gate walk (`typeHasGoNative` before the reach check), unchanged. */
	static function legacyTypeHasGoNative(t:Type):Bool {
		return legacyInner(t, []);
	}

	static function legacyInner(t:Type, seen:Array<String>):Bool {
		if (t == null) {
			return false;
		}
		var key = Std.string(t);
		if (seen.indexOf(key) != -1) {
			return false;
		}
		seen.push(key);

		return switch (t) {
			case TInst(classRef, params):
				if (isGoPath(GoTypeMapper.declarationPath(classRef))) {
					true;
				} else {
					legacyList(params, seen);
				}
			case TEnum(enumRef, params):
				if (isGoPath(GoTypeMapper.declarationPath(enumRef))) {
					true;
				} else {
					legacyList(params, seen);
				}
			case TType(typeRef, params):
				if (isGoPath(GoTypeMapper.declarationPath(typeRef))) {
					true;
				} else if (legacyList(params, seen)) {
					true;
				} else {
					legacyInner(typeRef.get().type, seen);
				}
			case TAbstract(absRef, params):
				if (isGoPath(GoTypeMapper.declarationPath(absRef))) {
					true;
				} else if (legacyList(params, seen)) {
					true;
				} else {
					legacyInner(GoTypeMapper.abstractDeclaration(absRef).type, seen);
				}
			case TFun(args, ret): var argHas = false; for (arg in args) {
					if (legacyInner(arg.t, seen)) {
						argHas = true;
						break;
					}
				} argHas || legacyInner(ret, seen);
			case TDynamic(inner): inner != null && legacyInner(inner, seen);
			case TAnonymous(anonRef):
				var anon = anonRef.get();
				var hasGo = false;
				for (field in anon.fields) {
					if (legacyInner(field.type, seen)) {
						hasGo = true;
						break;
					}
				}
				hasGo;
			case TLazy(loader):
				legacyInner(loader(), seen);
			case TMono(monoRef): var candidate = monoRef.get(); candidate != null && legacyInner(candidate, seen);
		}
	}

	static function legacyList(types:Array<Type>, seen:Array<String>):Bool {
		for (t in types) {
			if (legacyInner(t, seen)) {
				return true;
			}
		}
		return false;
	}

	static function legacyExprHasGoNative(expr:TypedExpr):Bool {
		var found = false;
		function walk(node:TypedExpr):Void {
			if (found || node == null) {
				return;
			}
			if (legacyTypeHasGoNative(node.t)) {
				found = true;
				return;
			}
			TypedExprTools.iter(node, child -> walk(child));
		}
		walk(expr);
		return found;
	}

	static function legacyFieldsUseGoNative(fields:Array<ClassField>):Bool {
		for (field in fields) {
			if (field == null) {
				continue;
			}
			if (legacyTypeHasGoNative(field.type)) {
				return true;
			}
			var expr = field.expr();
			if (expr != null && legacyExprHasGoNative(expr)) {
				return true;
			}
		}
		return false;
	}

	static function legacyClassUsesGoNative(classType:ClassType):Bool {
		if (classType.pack != null && classType.pack.length > 0 && classType.pack[0] == "go") {
			return true;
		}
		return legacyFieldsUseGoNative(classType.fields.get().concat(classType.statics.get()));
	}

	static function legacyAbstractUsesGoNative(abstractType:AbstractType):Bool {
		if (abstractType.pack != null && abstractType.pack.length > 0 && abstractType.pack[0] == "go") {
			return true;
		}
		if (legacyTypeHasGoNative(abstractType.type)) {
			return true;
		}
		if (abstractType.impl != null) {
			var impl = abstractType.impl.get();
			if (impl != null && legacyFieldsUseGoNative(impl.fields.get().concat(impl.statics.get()))) {
				return true;
			}
		}
		return false;
	}

	/** The replaced typed-hit collection, with the analyzer's unchanged helpers. */
	static function legacyCollect(types:Array<ModuleType>, projectRoot:String):Array<GoPortableNativeImportHit> {
		var hitsByModule:Map<String, GoPortableNativeImportHit> = [];
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!GoProfileContractAnalyzer.isPortableContractSource(classType.pos, projectRoot)) {
						continue;
					}
					var moduleName = GoProfileContractAnalyzer.moduleNameForClass(classType);
					if (hitsByModule.exists(moduleName)) {
						continue;
					}
					if (legacyClassUsesGoNative(classType)) {
						hitsByModule.set(moduleName, GoProfileContractAnalyzer.hitForModule(moduleName, classType.pos));
					}
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!GoProfileContractAnalyzer.isPortableContractSource(abstractType.pos, projectRoot)) {
						continue;
					}
					var moduleName = GoProfileContractAnalyzer.moduleNameForAbstract(abstractType);
					if (hitsByModule.exists(moduleName)) {
						continue;
					}
					if (legacyAbstractUsesGoNative(abstractType)) {
						hitsByModule.set(moduleName, GoProfileContractAnalyzer.hitForModule(moduleName, abstractType.pos));
					}
				case _:
			}
		}
		var out:Array<GoPortableNativeImportHit> = [for (moduleName in hitsByModule.keys()) hitsByModule.get(moduleName)];
		out.sort(GoProfileContractAnalyzer.comparePortableNativeHits);
		return out;
	}

	/** Exact reachability with no remembered answers: each declaration is expanded once per question. */
	static function reference(t:Null<Type>, visited:Map<String, Bool>):Bool {
		if (t == null) {
			return false;
		}
		function all(types:Array<Type>):Bool {
			for (entry in types)
				if (reference(entry, visited))
					return true;
			return false;
		}
		function declaration(key:String, explore:() -> Bool):Bool {
			if (visited.exists(key))
				return false;
			visited.set(key, true);
			return explore();
		}
		return switch (t) {
			case TInst(ref, params): isGoPath(ref.toString()) || all(params);
			case TEnum(ref, params): isGoPath(ref.toString()) || all(params);
			case TType(ref, params): isGoPath(ref.toString()) || all(params) || declaration("typedef " + ref.toString(),
					() -> reference(ref.get().type, visited));
			case TAbstract(ref, params): isGoPath(ref.toString()) || all(params) || declaration("abstract " + ref.toString(),
					() -> reference(ref.get().type, visited));
			case TFun(args, ret): all([for (arg in args) arg.t]) || reference(ret, visited);
			case TDynamic(inner): inner != null && reference(inner, visited);
			case TAnonymous(ref):
				final anon = ref.get();
				final fields = () -> all([for (field in anon.fields) field.type]);
				switch (anon.status) {
					case AClassStatics(owner): declaration("class statics " + owner.toString(), fields);
					case AEnumStatics(owner): declaration("enum statics " + owner.toString(), fields);
					case AAbstractStatics(owner): declaration("abstract statics " + owner.toString(), fields);
					case _: fields();
				}
			case TLazy(load): reference(load(), visited);
			case TMono(ref): reference(ref.get(), visited);
		};
	}

	static inline function isGoPath(path:String):Bool {
		return StringTools.startsWith(path, "go.");
	}

	static function referenceReaches(t:Type):Bool {
		return reference(t, new Map<String, Bool>());
	}

	/** One question: a label for messages and the type asked about. */
	static function question(label:String, type:Type):{label:String, type:Type} {
		return {label: label, type: type};
	}

	static function nested(levels:Int, leaf:Type):Type {
		var current = leaf;
		for (_ in 0...levels)
			current = TFun([], current);
		return current;
	}

	/** The statics type of a fixture abstract, with the static fields of `fieldsFrom` as its fields. */
	static function abstractStatics(name:String, fieldsFrom:ClassType):Type {
		final owner = switch (Context.getType("GoNativeTypeReachContractFixture." + name)) {
			case TAbstract(ref, _): ref;
			case _: throw "not an abstract: " + name;
		};
		final anon:AnonType = {fields: fieldsFrom.statics.get(), status: AAbstractStatics(owner)};
		return TAnonymous({get: () -> anon, toString: () -> "<anonymous>"});
	}

	static function moduleTypes(path:String):Array<ModuleType> {
		return [
			for (type in Context.getModule(path))
				switch (type) {
					case TInst(ref, _):
						TClassDecl(ref);
					case TAbstract(ref, _):
						TAbstract(ref);
					case TEnum(ref, _):
						TEnumDecl(ref);
					case TType(ref, _):
						TTypeDecl(ref);
					case _:
						throw "unexpected module type in " + path;
				}
		];
	}

	static function describe(hits:Array<GoPortableNativeImportHit>):String {
		return [for (hit in hits) hit.module + "@" + hit.location].join(",");
	}

	static function run():Void {
		final fixture = load("GoNativeTypeReachContractFixture");
		final classes = [
			fixture,
			load("GoNativeTypeReachContractFixture.ReachPortable"),
			load("GoNativeTypeReachContractFixture.ReachStatics"),
			load("GoNativeTypeReachContractFixture.ReachGoStatics"),
			load("GoNativeTypeReachContractLegacyFixture"),
			load("GoNativeTypeReachContractLegacyFixture.GoNativeTypeReachContractLegacyPortable")
		];
		final questions = new Array<{label:String, type:Type}>();
		for (source in classes) {
			for (field in source.fields.get().concat(source.statics.get())) {
				questions.push(question(source.name + "." + field.name, field.type));
				final body = field.expr();
				if (body == null)
					continue;
				var index = 0;
				function visit(node:TypedExpr):Void {
					questions.push(question(source.name + "." + field.name + "#" + index++, node.t));
					TypedExprTools.iter(node, visit);
				}
				visit(body);
			}
		}
		final goSlice = Context.getType("go.NativeSlice");
		final plainInt = Context.getType("Int");
		final constructed = [
			question("lazy go", TLazy(() -> goSlice)),
			question("lazy plain", TLazy(() -> plainInt)),
			question("bound monomorph go", TMono({get: () -> goSlice, toString: () -> "<mono>"})),
			question("bound monomorph plain", TMono({get: () -> plainInt, toString: () -> "<mono>"})),
			question("unbound monomorph", TMono({get: () -> null, toString: () -> "<mono>"})),
			question("100 levels plain", nested(100, plainInt)),
			question("100 levels go", nested(100, goSlice)),
			// Typed code reaches an abstract's statics through its implementation class, so these are built.
			question("abstract statics go", abstractStatics("ReachOverInt", load("GoNativeTypeReachContractFixture.ReachGoStatics"))),
			question("abstract statics plain", abstractStatics("ReachOverSlice", load("GoNativeTypeReachContractFixture.ReachStatics")))
		];
		for (entry in constructed)
			questions.push(entry);
		require(questions.length >= 60, "too few questions: " + questions.length);

		// Named fields have the expected answers.
		final expected = [
			"plain" => false,
			"loopB" => true,
			"loopA" => true,
			"overLoopB" => true,
			"wrapped" => true,
			"selfLoop" => false,
			"selfLoopGo" => true,
			"overSlice" => true,
			"overInt" => false,
			"overOverSlice" => true,
			"nullableOverSlice" => true,
			"arrayOfSlices" => true,
			"arrayOfInts" => false,
			"dynamicGo" => true,
			"dynamicPlain" => false,
			"anyValue" => false,
			"argumentGo" => true,
			"resultGo" => true,
			"functionPlain" => false,
			"anonymousPair" => true,
			"anonymousPlain" => false,
			"mapGo" => true,
			"mapPlain" => false,
			"kind" => false,
			"notGo" => false,
			"generic" => false
		];
		for (field in fixture.fields.get()) {
			if (!expected.exists(field.name))
				continue;
			require(new GoNativeTypeReach().reaches(field.type) == expected.get(field.name), "unexpected answer for field " + field.name);
		}
		for (name in expected.keys())
			require([for (field in fixture.fields.get()) field.name].contains(name), "fixture lost field " + name);
		final statics = [for (entry in questions) if (entry.label.indexOf("staticsCount#") != -1) entry];
		final goStatics = [
			for (entry in questions)
				if (entry.label.indexOf("goStaticsMissing#") != -1) entry
		];
		require(statics.length > 0 && !Lambda.exists(statics, entry -> new GoNativeTypeReach().reaches(entry.type)),
			"a statics type that reaches no go.* type answered true");
		require(Lambda.exists(goStatics, entry -> new GoNativeTypeReach().reaches(entry.type)), "a statics type that reaches go.* answered false");

		// Remembered answers equal the reference for several question orders and for each question alone.
		final reversed = questions.copy();
		reversed.reverse();
		final interleaved = [for (index in 0...questions.length) questions[(index * 7) % questions.length]];
		for (order in [questions, reversed, interleaved]) {
			final reach = new GoNativeTypeReach();
			for (entry in order)
				require(reach.reaches(entry.type) == referenceReaches(entry.type), "remembered answer differs from the reference for " + entry.label);
			for (entry in order)
				require(reach.reaches(entry.type) == referenceReaches(entry.type), "asked again, answer differs from the reference for " + entry.label);
		}
		final answers = [for (entry in questions) referenceReaches(entry.type)];
		for (index in 0...questions.length)
			require(new GoNativeTypeReach().reaches(questions[index].type) == answers[index],
				"answer alone differs from the reference for " + questions[index].label);
		// Every ordered pair, on a new instance and after a question that expanded a typedef: an answer
		// left behind by the first question must not change the second.
		final warmUp = Lambda.find(questions, entry -> entry.label == "GoNativeTypeReachContractFixture.plain");
		require(warmUp != null, "fixture lost field plain");
		for (warm in [false, true]) {
			for (first in 0...questions.length) {
				for (second in 0...questions.length) {
					final reach = new GoNativeTypeReach();
					if (warm)
						reach.reaches(warmUp.type);
					reach.reaches(questions[first].type);
					require(reach.reaches(questions[second].type) == answers[second],
						"after " + (warm ? "plain, " : "") + questions[first].label + ", answer differs from the reference for " + questions[second].label);
				}
			}
		}

		// Types nested beyond the depth limit answer true, so the gate's walk decides them.
		require(new GoNativeTypeReach().reaches(nested(300, goSlice)), "a deep go.* type answered false");
		require(new GoNativeTypeReach().reaches(nested(300, plainInt)), "a type beyond the depth limit answered false");
		require(!GoProfileContractAnalyzer.typeHasGoNative(nested(300, plainInt), new GoNativeTypeReach()), "the gate reported a deep type without go.*");

		// The gate answers exactly as the replaced walk; the reach check never answers false where the walk answers true.
		final shared = new GoNativeTypeReach();
		for (entry in questions) {
			final legacy = legacyTypeHasGoNative(entry.type);
			require(!legacy || shared.reaches(entry.type), "the replaced walk found go.* where the reach check did not, for " + entry.label);
			require(GoProfileContractAnalyzer.typeHasGoNative(entry.type, shared) == legacy, "the gate's answer changed for " + entry.label);
			require(GoProfileContractAnalyzer.typeHasGoNative(entry.type, new GoNativeTypeReach()) == legacy,
				"the gate's answer alone changed for " + entry.label);
		}

		// The walk misses the second anonymous structure; the reach check sees it; the gate keeps the walk's answer.
		final legacyInterface = load("GoNativeTypeReachContractLegacyFixture");
		final take = legacyInterface.fields.get()[0].type;
		require(!legacyTypeHasGoNative(take), "the replaced walk now sees the second anonymous argument");
		require(new GoNativeTypeReach().reaches(take), "the reach check missed the second anonymous argument");
		require(!GoProfileContractAnalyzer.typeHasGoNative(take, new GoNativeTypeReach()), "the gate reports the second anonymous argument");

		// Per class and abstract, and in the hit list, the gate equals the replaced collection.
		final types = moduleTypes("GoNativeTypeReachContractFixture").concat(moduleTypes("GoNativeTypeReachContractLegacyFixture"));
		final perType = new GoNativeTypeReach();
		var classCount = 0;
		var abstractCount = 0;
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(ref):
					final source = ref.get();
					classCount++;
					final legacy = legacyClassUsesGoNative(source);
					require(GoProfileContractAnalyzer.classUsesGoNative(source, perType) == legacy, "class answer changed for " + source.name);
					require(GoProfileContractAnalyzer.classUsesGoNative(source, new GoNativeTypeReach()) == legacy,
						"class answer alone changed for " + source.name);
				case TAbstract(ref):
					final source = ref.get();
					abstractCount++;
					final legacy = legacyAbstractUsesGoNative(source);
					require(GoProfileContractAnalyzer.abstractUsesGoNative(source, perType) == legacy, "abstract answer changed for " + source.name);
					require(GoProfileContractAnalyzer.abstractUsesGoNative(source, new GoNativeTypeReach()) == legacy,
						"abstract answer alone changed for " + source.name);
				case _:
			}
		}
		require(classCount >= 6 && abstractCount >= 4, "fixture modules lost classes or abstracts");
		final root = Sys.getCwd();
		final hits = describe(GoProfileContractAnalyzer.collectPortableNativeImportTypedHits(types, root, []));
		require(hits == describe(legacyCollect(types, root)), "hit list changed: " + hits);
		// The module's first type that uses go.* is the abstract ReachOverSlice; the hit points at its line.
		final overSlice = switch (Context.getType("GoNativeTypeReachContractFixture.ReachOverSlice")) {
			case TAbstract(ref, _): ref.get();
			case _: throw "not an abstract: ReachOverSlice";
		};
		final overSliceLine = PositionTools.toLocation(overSlice.pos).range.start.line;
		require(hits == "GoNativeTypeReachContractFixture@GoNativeTypeReachContractFixture:" + overSliceLine, "unexpected hit list: " + hits);
	}
	#end
}
