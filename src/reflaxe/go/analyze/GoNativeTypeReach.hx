package reflaxe.go.analyze;

#if (macro || reflaxe_runtime)
import haxe.macro.Type;
import reflaxe.go.compiler.GoTypeMapper;

/**
	What: Answers whether a macro type can reach a `go.*` declaration, and
	remembers the answer for each typedef, abstract, and statics type it expands.

	Why: The native authority gate asks this about the type of every expression
	node in every portable class: about 96,000 questions on the BeadsHX bootstrap
	compile. The gate's own walk (`GoProfileContractAnalyzer.typeHasGoNativeInner`)
	starts again for each question, expands the same typedefs and abstracts each
	time, and prints every visited type with `Std.string` to detect cycles. That
	cost about 3.9 s of the compile, and most of these types reach no `go.*`
	declaration at all.

	How: The walk follows exactly the edges of the gate's walk: the path of each
	class, enum, typedef, and abstract, its type parameters, the declared type of
	each typedef and abstract, function arguments and results, `Dynamic<T>`,
	anonymous fields, lazy types, and bound monomorphs. A typedef, an abstract,
	and the statics type of a class, enum, or abstract are declarations: an
	expansion is the same wherever the declaration appears, so its answer is kept
	by path. A path names one declaration: Haxe rejects two types with one path,
	private types get module-qualified paths, and the typedef that wraps a statics
	type (such as `Class<Main>`) names its owner.

	An expansion that meets a declaration it is still expanding, or one it has
	already expanded and found false, answers false for it. That is safe for the
	outermost expansion, whose walk covers everything the inner ones can reach,
	but an inner false answer can depend on a declaration further out that is
	still open. So a true answer is kept at once (it names a real `go.*`
	declaration), and false answers are kept only when the outermost expansion
	ends false, for every declaration it met. Other inner answers are found again
	when asked later.

	The answer is therefore exact reachability. Every question answered false
	here is also false in the gate's walk, because that walk follows the same
	edges and can only skip some of them. Types nested deeper than `MAX_DEPTH`
	answer true, which can only send a caller to its slower walk. Declarations do
	not change after typing; make one instance per analysis.
**/
class GoNativeTypeReach {
	/** Nesting depth at which the walk stops and answers true, so a caller checks the type itself. */
	static inline final MAX_DEPTH = 256;

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
		return walk(type, 0);
	}

	function walk(type:Null<Type>, depth:Int):Bool {
		if (type == null) {
			return false;
		}
		if (depth > MAX_DEPTH) {
			return true;
		}
		final next = depth + 1;
		return switch (type) {
			case TInst(ref, params): isGoPath(GoTypeMapper.declarationPath(ref)) || walkAll(params, next);
			case TEnum(ref, params): isGoPath(GoTypeMapper.declarationPath(ref)) || walkAll(params, next);
			case TType(ref, params): final path = GoTypeMapper.declarationPath(ref); isGoPath(path) || walkAll(params,
					next) || expand(typedefs, path, inner -> walk(ref.get().type, inner), next);
			case TAbstract(ref, params): final path = GoTypeMapper.declarationPath(ref); isGoPath(path) || walkAll(params,
					next) || expand(abstracts, path, inner -> walk(GoTypeMapper.abstractDeclaration(ref).type, inner), next);
			case TFun(args, ret): var found = false; for (arg in args) {
					if (walk(arg.t, next)) {
						found = true;
						break;
					}
				} found || walk(ret, next);
			case TDynamic(inner): inner != null && walk(inner, next);
			case TAnonymous(ref):
				final anon = ref.get();
				switch (anon.status) {
					case AClassStatics(owner):
						expand(classStatics, GoTypeMapper.declarationPath(owner), inner -> walkFields(anon.fields, inner), next);
					case AEnumStatics(owner):
						expand(enumStatics, GoTypeMapper.declarationPath(owner), inner -> walkFields(anon.fields, inner), next);
					case AAbstractStatics(owner):
						expand(abstractStatics, GoTypeMapper.declarationPath(owner), inner -> walkFields(anon.fields, inner), next);
					case _:
						walkFields(anon.fields, next);
				}
			case TLazy(load):
				walk(load(), next);
			case TMono(ref):
				walk(ref.get(), next);
		};
	}

	function walkAll(types:Array<Type>, depth:Int):Bool {
		for (type in types) {
			if (walk(type, depth)) {
				return true;
			}
		}
		return false;
	}

	function walkFields(fields:Array<ClassField>, depth:Int):Bool {
		for (field in fields) {
			if (walk(field.type, depth)) {
				return true;
			}
		}
		return false;
	}

	/** Answers for one declaration, expanding it with `explore` the first time it is met. */
	function expand(memo:Map<String, GoNativeReachAnswer>, path:String, explore:Int->Bool, depth:Int):Bool {
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
		final found = explore(depth);
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
