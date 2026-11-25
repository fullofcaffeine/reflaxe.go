#if macro
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.analyze.GoNativeTypeReach;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.GoPortableNativeImportHit;
#end

/**
	Contract for the native authority gate's answer to "can this type reach a
	`go.*` declaration?".

	Why
	The gate (`GoProfileContractAnalyzer`) answers with `GoNativeTypeReach`,
	which remembers answers per declaration. A remembered answer that is wrong
	for a later question would hide a `go.*` use from the gate, or report one
	that is not there. The walk it replaced kept printed types as its visited
	set, so it missed a `go.*` type that prints like a type met before: every
	anonymous structure prints the same, and prints are cut short after two
	levels of nesting.

	What
	For the field types and expression node types of real typed fixture classes,
	and for constructed lazy, monomorph, and deeply nested types, pins that:
	- the answer equals an independent reachability walk with no remembered
	  answers, for several question orders, for each question alone, and for
	  every ordered pair of questions;
	- named fixture fields have the expected answers, including cycles through
	  typedefs and statics types;
	- naming a type counts only the statics of a `go.*` type, so a caller of a
	  go-free static or module-level function is not counted because a sibling
	  has a `go.*` type;
	- the gate's answers per class and abstract, and its hit list, equal the
	  same collection over the independent walk, and the gate reports the
	  `go.*` types behind types that print alike.

	How
	After the initialization macros, loads the fixture modules, collects the
	questions, and compares the answers with the independent walk.
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
				// Only the statics of a go.* owner are walked.
				switch (anon.status) {
					case AClassStatics(owner): isGoPath(owner.toString()) && declaration("class statics " + owner.toString(), fields);
					case AEnumStatics(owner): isGoPath(owner.toString()) && declaration("enum statics " + owner.toString(), fields);
					case AAbstractStatics(owner): isGoPath(owner.toString()) && declaration("abstract statics " + owner.toString(), fields);
					case _: fields();
				}
			case TLazy(load): reference(load(), visited);
			case TMono(ref): reference(ref.get(), visited);
		};
	}

	static inline function isGoPath(path:String):Bool {
		return StringTools.startsWith(path, "go.");
	}

	static function isGoPack(pack:Array<String>):Bool {
		return pack != null && pack.length > 0 && pack[0] == "go";
	}

	static function referenceReaches(t:Type):Bool {
		return reference(t, new Map<String, Bool>());
	}

	static function referenceExprReaches(expr:TypedExpr):Bool {
		var found = false;
		function visit(node:TypedExpr):Void {
			if (found || node == null) {
				return;
			}
			if (referenceReaches(node.t)) {
				found = true;
				return;
			}
			TypedExprTools.iter(node, visit);
		}
		visit(expr);
		return found;
	}

	static function referenceFieldsUseGoNative(fields:Array<ClassField>):Bool {
		for (field in fields) {
			if (field == null) {
				continue;
			}
			if (referenceReaches(field.type)) {
				return true;
			}
			final body = field.expr();
			if (body != null && referenceExprReaches(body)) {
				return true;
			}
		}
		return false;
	}

	static function referenceClassUsesGoNative(classType:ClassType):Bool {
		return isGoPack(classType.pack) || referenceFieldsUseGoNative(classType.fields.get().concat(classType.statics.get()));
	}

	static function referenceAbstractUsesGoNative(abstractType:AbstractType):Bool {
		if (isGoPack(abstractType.pack) || referenceReaches(abstractType.type)) {
			return true;
		}
		final impl = abstractType.impl == null ? null : abstractType.impl.get();
		return impl != null && referenceFieldsUseGoNative(impl.fields.get().concat(impl.statics.get()));
	}

	/** The gate's typed-hit collection over the reference walk, with the analyzer's helpers for the rest. */
	static function referenceCollect(types:Array<ModuleType>, projectRoot:String):Array<GoPortableNativeImportHit> {
		var hitsByModule:Map<String, GoPortableNativeImportHit> = [];
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					if (!GoProfileContractAnalyzer.isPortableContractSource(classType.pos, projectRoot)) {
						continue;
					}
					var moduleName = GoProfileContractAnalyzer.moduleNameForClass(classType);
					if (!hitsByModule.exists(moduleName) && referenceClassUsesGoNative(classType)) {
						hitsByModule.set(moduleName, GoProfileContractAnalyzer.hitForModule(moduleName, classType.pos));
					}
				case TAbstract(abstractRef):
					var abstractType = abstractRef.get();
					if (!GoProfileContractAnalyzer.isPortableContractSource(abstractType.pos, projectRoot)) {
						continue;
					}
					var moduleName = GoProfileContractAnalyzer.moduleNameForAbstract(abstractType);
					if (!hitsByModule.exists(moduleName) && referenceAbstractUsesGoNative(abstractType)) {
						hitsByModule.set(moduleName, GoProfileContractAnalyzer.hitForModule(moduleName, abstractType.pos));
					}
				case _:
			}
		}
		var out:Array<GoPortableNativeImportHit> = [for (moduleName in hitsByModule.keys()) hitsByModule.get(moduleName)];
		out.sort(GoProfileContractAnalyzer.comparePortableNativeHits);
		return out;
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

	/** The statics type of the abstract at `path`, with the static fields of `fieldsFrom` as its fields. */
	static function abstractStatics(path:String, fieldsFrom:ClassType):Type {
		final owner = switch (Context.getType(path)) {
			case TAbstract(ref, _): ref;
			case _: throw "not an abstract: " + path;
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

	static function declarationLine(pos:haxe.macro.Expr.Position):Int {
		return PositionTools.toLocation(pos).range.start.line;
	}

	static function run():Void {
		final fixture = load("GoNativeTypeReachContractFixture");
		final classes = [
			fixture,
			load("GoNativeTypeReachContractFixture.ReachPortable"),
			load("GoNativeTypeReachContractFixture.ReachStatics"),
			load("GoNativeTypeReachContractFixture.ReachGoStatics"),
			load("GoNativeTypeReachContractFixture.ReachMixedStatics"),
			load("GoNativeTypeReachContractFixture.ReachSiblingCaller"),
			load("GoNativeTypeReachContractFixture.ReachNativeCaller"),
			load("GoNativeTypeReachContractFixture.ReachNativeModuleCaller"),
			load("GoNativeTypeReachContractFixture.ReachGoOwnerCaller"),
			load("GoNativeTypeReachContractPrintFixture"),
			load("GoNativeTypeReachContractPrintFixture.GoNativeTypeReachContractNestedPrint"),
			load("GoNativeTypeReachContractPrintFixture.GoNativeTypeReachContractPrintPortable")
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
			question("300 levels plain", nested(300, plainInt)),
			question("300 levels go", nested(300, goSlice)),
			// Typed code reaches an abstract's statics through its implementation class, so these are built.
			question("go abstract statics with go fields", abstractStatics("go.NativeAny", load("GoNativeTypeReachContractFixture.ReachGoStatics"))),
			// Each owner keeps one set of fields, as in typed code: a path names one statics type.
			question("go abstract statics with plain fields", abstractStatics("go.Go.__ChanHandle", load("GoNativeTypeReachContractFixture.ReachStatics"))),
			question("abstract statics with go fields",
				abstractStatics("GoNativeTypeReachContractFixture.ReachOverInt", load("GoNativeTypeReachContractFixture.ReachGoStatics"))),
			question("abstract statics with plain fields",
				abstractStatics("GoNativeTypeReachContractFixture.ReachOverSlice", load("GoNativeTypeReachContractFixture.ReachStatics")))
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
		final constructedAnswers = [
			"go abstract statics with go fields" => true,
			"go abstract statics with plain fields" => false,
			"abstract statics with go fields" => false,
			"abstract statics with plain fields" => false
		];
		for (label => want in constructedAnswers) {
			final entry = Lambda.find(questions, entry -> entry.label == label);
			require(entry != null && new GoNativeTypeReach().reaches(entry.type) == want, "unexpected answer for " + label);
		}

		// Naming a type counts only the statics of a go.* type; each member used through it has its own type.
		final callers = [
			"ReachSiblingCaller" => false,
			"ReachNativeCaller" => true,
			"ReachNativeModuleCaller" => true,
			"ReachGoOwnerCaller" => true
		];
		for (name => want in callers)
			require(GoProfileContractAnalyzer.classUsesGoNative(load("GoNativeTypeReachContractFixture." + name), new GoNativeTypeReach()) == want,
				"unexpected class answer for " + name);

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

		// Deep types answer exactly: no depth limit answers for them.
		require(new GoNativeTypeReach().reaches(nested(300, goSlice)), "a deep go.* type answered false");
		require(!new GoNativeTypeReach().reaches(nested(300, plainInt)), "a deep type without go.* answered true");

		// A go.* type behind a type that prints the same as an earlier one is found.
		for (path in [
			"GoNativeTypeReachContractPrintFixture",
			"GoNativeTypeReachContractPrintFixture.GoNativeTypeReachContractNestedPrint"
		]) {
			final take = load(path).fields.get()[0];
			require(take.name == "take", "print fixture lost method take in " + path);
			require(new GoNativeTypeReach().reaches(take.type), "the gate missed the go.* argument behind a like-printed one in " + path);
		}

		// Per class and abstract, and in the hit list, the gate equals the collection over the reference.
		final types = moduleTypes("GoNativeTypeReachContractFixture").concat(moduleTypes("GoNativeTypeReachContractPrintFixture"));
		final perType = new GoNativeTypeReach();
		var classCount = 0;
		var abstractCount = 0;
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(ref):
					final source = ref.get();
					classCount++;
					final want = referenceClassUsesGoNative(source);
					require(GoProfileContractAnalyzer.classUsesGoNative(source, perType) == want, "class answer differs for " + source.name);
					require(GoProfileContractAnalyzer.classUsesGoNative(source, new GoNativeTypeReach()) == want,
						"class answer alone differs for " + source.name);
				case TAbstract(ref):
					final source = ref.get();
					abstractCount++;
					final want = referenceAbstractUsesGoNative(source);
					require(GoProfileContractAnalyzer.abstractUsesGoNative(source, perType) == want, "abstract answer differs for " + source.name);
					require(GoProfileContractAnalyzer.abstractUsesGoNative(source, new GoNativeTypeReach()) == want,
						"abstract answer alone differs for " + source.name);
				case _:
			}
		}
		require(classCount >= 13 && abstractCount >= 4, "fixture modules lost classes or abstracts");
		final root = Sys.getCwd();
		final hits = describe(GoProfileContractAnalyzer.collectPortableNativeImportTypedHits(types, root, []));
		require(hits == describe(referenceCollect(types, root)), "hit list differs from the reference: " + hits);
		// Each module's hit points at its first type that uses go.*: the abstract ReachOverSlice, and the
		// interface whose second anonymous argument is its only go.* type.
		final overSlice = switch (Context.getType("GoNativeTypeReachContractFixture.ReachOverSlice")) {
			case TAbstract(ref, _): ref.get();
			case _: throw "not an abstract: ReachOverSlice";
		};
		final printLine = declarationLine(load("GoNativeTypeReachContractPrintFixture").pos);
		require(hits == "GoNativeTypeReachContractFixture@GoNativeTypeReachContractFixture:"
			+ declarationLine(overSlice.pos)
			+ ",GoNativeTypeReachContractPrintFixture@GoNativeTypeReachContractPrintFixture:"
			+ printLine,
			"unexpected hit list: "
			+ hits);
	}
	#end
}
