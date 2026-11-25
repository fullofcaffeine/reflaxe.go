typedef ReachSlice = go.NativeSlice<Bool>;
typedef ReachWrapped = {value:ReachSlice};
typedef ReachPlain = {value:Int, name:String};

/** Reaches `go.*` only after its first field leads into `ReachLoopB`, which leads back here. */
typedef ReachLoopA = {a_next:Null<ReachLoopB>, z_slice:go.NativeSlice<Bool>};

typedef ReachLoopB = {back:Null<ReachLoopA>, count:Int};
typedef ReachSelfLoop = {next:Null<ReachSelfLoop>, count:Int};
typedef ReachSelfLoopGo = {next:Null<ReachSelfLoopGo>, value:go.NativeAny};
abstract ReachOverSlice(go.NativeSlice<Bool>) {}
abstract ReachOverInt(Int) {}
abstract ReachOverOverSlice(ReachOverSlice) {}
abstract ReachOverLoopB(ReachLoopB) {}

enum ReachKind {
	Plain;
}

/** An enum whose constructor takes a `go.*` value, so its statics type reaches `go.*`. */
enum ReachGoKind {
	Holding(values:go.NativeSlice<Bool>);
}

/** A class whose statics type refers to itself and reaches no `go.*` type. */
class ReachStatics {
	public static var count:Int = 0;
	public static var self = ReachStatics;
}

/**
	A class with a `go.*` static field whose statics type refers to itself.

	Its statics type does not reach `go.*`: only the statics of a `go.*` owner
	count, and an access to `values` has the field's own `go.*` type.
**/
class ReachGoStatics {
	public static var values:Null<go.NativeSlice<Bool>> = null;
	public static var self = ReachGoStatics;
}

/** A class with a go-free static next to one whose result type is `go.*`. */
class ReachMixedStatics {
	public static function plain():Int {
		return 1;
	}

	public static function native():Null<go.NativeSlice<Bool>> {
		return null;
	}
}

/** Calls only go-free statics and module-level functions whose siblings have `go.*` types. */
class ReachSiblingCaller {
	public static function run():Int {
		return ReachMixedStatics.plain() + plainModuleFunction();
	}
}

/** Calls a static whose own result type is `go.*`. */
class ReachNativeCaller {
	public static function run():Bool {
		return ReachMixedStatics.native() == null;
	}
}

/** Calls a module-level function whose own result type is `go.*`. */
class ReachNativeModuleCaller {
	public static function run():Bool {
		return nativeModuleFunction() == null;
	}
}

/** Names a `go.*` class: a call to its go-free static counts while another of its statics has a `go.*` type. */
class ReachGoOwnerCaller {
	public static function run():Void {
		go.Go.spawn(() -> {});
	}
}

/** A go-free module-level function next to one whose result type is `go.*`. */
function plainModuleFunction():Int {
	return 2;
}

function nativeModuleFunction():Null<go.NativeSlice<Bool>> {
	return null;
}

/** A class with no `go.*` type in any field type or expression. */
class ReachPortable {
	public var plain:ReachPlain;
	public var loop:ReachSelfLoop;

	public function new() {}

	public function total():Int {
		return plain.value + ReachStatics.count;
	}
}

/** Field and expression types that `GoNativeTypeReachContract` asks about. */
class GoNativeTypeReachContractFixture<T> {
	public var plain:ReachPlain;
	public var loopB:ReachLoopB;
	public var loopA:ReachLoopA;
	public var overLoopB:ReachOverLoopB;
	public var wrapped:ReachWrapped;
	public var selfLoop:ReachSelfLoop;
	public var selfLoopGo:ReachSelfLoopGo;
	public var overSlice:ReachOverSlice;
	public var overInt:ReachOverInt;
	public var overOverSlice:ReachOverOverSlice;
	public var nullableOverSlice:Null<ReachOverSlice>;
	public var arrayOfSlices:Array<go.NativeSlice<Bool>>;
	public var arrayOfInts:Array<Int>;
	public var dynamicGo:Dynamic<go.NativeAny>;
	public var dynamicPlain:Dynamic<Int>;
	public var anyValue:Dynamic;
	public var argumentGo:(go.NativeAny) -> Int;
	public var resultGo:(Int) -> ReachWrapped;
	public var functionPlain:(Int, ReachPlain) -> String;
	public var anonymousPair:{first:{count:Int}, second:{values:go.NativeSlice<Bool>}};
	public var anonymousPlain:{first:{count:Int}, second:{name:String}};
	public var mapGo:Map<String, go.NativeAny>;
	public var mapPlain:Map<String, ReachPlain>;
	public var kind:ReachKind;
	public var notGo:gonot.GoNativeTypeReachNotGo;
	public var generic:T;

	public function new() {}

	public function staticsCount():Int {
		return ReachStatics.count;
	}

	public function goStaticsMissing():Bool {
		return ReachGoStatics.values == null;
	}

	public function goKind():ReachGoKind {
		return ReachGoKind.Holding(null);
	}
}
