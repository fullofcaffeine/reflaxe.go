package reflaxe.go.analyze;

#if (macro || reflaxe_runtime)
import haxe.macro.Type;
import reflaxe.go.compiler.GoTypeMapper;

/**
	What: Answers whether a macro type can reach a `go.*` declaration, and
	remembers the answer for each typedef, abstract, and `go.*` statics type it
	expands.

	Why: The native authority gate (`GoProfileContractAnalyzer`) asks this about
	the type of every field and expression node in every portable class: about
	96,000 questions on the BeadsHX bootstrap compile. Its earlier walk started
	again for each question and kept the `Std.string` print of every visited type
	as its visited set. That cost about 3.9 s of the compile, and it missed `go.*`
	types: in eval every anonymous structure prints as `<anonymous>` and prints
	are cut short after two levels of nesting, so the walk skipped any type that
	printed like one it had met before.

	How: The walk follows the path of each class, enum, typedef, and abstract,
	its type parameters, the declared type of each typedef and abstract, function
	arguments and results, `Dynamic<T>`, anonymous fields, lazy types, and bound
	monomorphs. A typedef, an abstract, and the statics type of a class, enum, or
	abstract are declarations: an expansion is the same wherever the declaration
	appears, so its answer is kept by path. A path names one declaration: Haxe
	rejects two types with one path, private types get module-qualified paths,
	and the typedef that wraps a statics type (such as `Class<Main>`) names its
	owner.

	A statics type is the type of an expression that names a class, enum, or
	abstract, such as `Helper` in `Helper.plain()`; its fields are all the
	statics of its owner. Only the statics type of a `go.*` owner is expanded.
	For any other owner it answers false: a member used through it is a separate
	expression with its own type, so a caller of a go-free `Helper.plain()` is
	not counted as using `go.*` because another static of `Helper` has a `go.*`
	type. Module-level functions are statics of their module's `_Fields_` class,
	so the same holds for them. The statics of a `go.*` owner are walked like
	any other expansion, so a go-free static of a `go.*` type counts while
	another of its statics has a `go.*` type (haxe_go-jo2j tracks the `go.*`
	types whose statics are all go-free).

	An expansion that meets a declaration it is still expanding, or one it has
	already expanded and found false, answers false for it. That is safe for the
	outermost expansion, whose walk covers everything the inner ones can reach,
	but an inner false answer can depend on a declaration further out that is
	still open. So a true answer is kept at once (it names a real `go.*`
	declaration), and false answers are kept only when the outermost expansion
	ends false, for every declaration it met. Other inner answers are found again
	when asked later.

	The answer is therefore exact reachability, with no depth limit of its own,
	and it is the gate's final answer. Only declarations are remembered: other
	types are trees (Haxe refuses to bind a monomorph to a type that contains it),
	so the walk ends without a visited set for them. Declarations do not change
	after typing; make one instance per analysis.
**/
class GoNativeTypeReach {
	final typedefs:Map<String, GoNativeReachAnswer> = [];
	final abstracts:Map<String, GoNativeReachAnswer> = [];
	final classStatics:Map<String, GoNativeReachAnswer> = [];
	final enumStatics:Map<String, GoNativeReachAnswer> = [];
	final abstractStatics:Map<String, GoNativeReachAnswer> = [];

	/** Declarations met by the open outermost expansion, as (answer map, path) pairs. */
	final openMemos:Array<Map<String, GoNativeReachAnswer>> = [];

	final openPaths:Array<String> = [];

	public function new() {}

	/** True when `type` can reach a declaration whose dotted path starts with `go.`. */
	public function reaches(type:Null<Type>):Bool {
		return walk(type);
	}

	function walk(type:Null<Type>):Bool {
		if (type == null) {
			return false;
		}
		return switch (type) {
			case TInst(ref, params): isGoPath(GoTypeMapper.declarationPath(ref)) || walkAll(params);
			case TEnum(ref, params): isGoPath(GoTypeMapper.declarationPath(ref)) || walkAll(params);
			case TType(ref, params): final path = GoTypeMapper.declarationPath(ref); isGoPath(path) || walkAll(params) || expand(typedefs, path,
					() -> walk(ref.get().type));
			case TAbstract(ref, params): final path = GoTypeMapper.declarationPath(ref); isGoPath(path) || walkAll(params) || expand(abstracts, path,
					() -> walk(GoTypeMapper.abstractDeclaration(ref).type));
			case TFun(args, ret): var found = false; for (arg in args) {
					if (walk(arg.t)) {
						found = true;
						break;
					}
				} found || walk(ret);
			case TDynamic(inner): inner != null && walk(inner);
			case TAnonymous(ref):
				final anon = ref.get();
				switch (anon.status) {
					case AClassStatics(owner):
						walkStatics(classStatics, GoTypeMapper.declarationPath(owner), anon.fields);
					case AEnumStatics(owner):
						walkStatics(enumStatics, GoTypeMapper.declarationPath(owner), anon.fields);
					case AAbstractStatics(owner):
						walkStatics(abstractStatics, GoTypeMapper.declarationPath(owner), anon.fields);
					case _:
						walkFields(anon.fields);
				}
			case TLazy(load):
				walk(load());
			case TMono(ref):
				walk(ref.get());
		};
	}

	function walkAll(types:Array<Type>):Bool {
		for (type in types) {
			if (walk(type)) {
				return true;
			}
		}
		return false;
	}

	function walkFields(fields:Array<ClassField>):Bool {
		for (field in fields) {
			if (walk(field.type)) {
				return true;
			}
		}
		return false;
	}

	/** Answers for the statics type of the declaration at `ownerPath`: false unless the owner is `go.*`. */
	function walkStatics(memo:Map<String, GoNativeReachAnswer>, ownerPath:String, fields:Array<ClassField>):Bool {
		return isGoPath(ownerPath) && expand(memo, ownerPath, () -> walkFields(fields));
	}

	/** Answers for one declaration, expanding it with `explore` the first time it is met. */
	function expand(memo:Map<String, GoNativeReachAnswer>, path:String, explore:() -> Bool):Bool {
		switch (memo.get(path)) {
			case Reaches:
				return true;
			case Unreached | Open:
				return false;
			case null:
		}
		final outermost = openPaths.length == 0;
		memo.set(path, Open);
		openMemos.push(memo);
		openPaths.push(path);
		final found = explore();
		if (found) {
			memo.set(path, Reaches);
		}
		if (outermost) {
			for (index in 0...openPaths.length) {
				final openMemo = openMemos[index];
				final openPath = openPaths[index];
				if (openMemo.get(openPath) == Open) {
					// The outermost walk covered all it can reach: false is final, but with a
					// true answer an open declaration may still reach `go.*` another way.
					if (found) {
						openMemo.remove(openPath);
					} else {
						openMemo.set(openPath, Unreached);
					}
				}
			}
			openMemos.resize(0);
			openPaths.resize(0);
		}
		return found;
	}

	/** True when a dotted declaration path (`pack.Name`) lies in the `go` package tree. */
	static inline function isGoPath(path:String):Bool {
		return StringTools.startsWith(path, "go.");
	}
}

/**
	The state of one declaration in `GoNativeTypeReach`.

	`Open` means the open outermost expansion has met the declaration and found
	no `go.*` declaration through it so far.
**/
private enum abstract GoNativeReachAnswer(Int) {
	var Reaches = 1;
	var Unreached = 0;
	var Open = 2;
}
#end
