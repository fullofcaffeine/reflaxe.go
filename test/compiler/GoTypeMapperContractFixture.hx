/** A private abstract that shares its name with one in `GoTypeMapperContractOtherFixture`. */
private abstract Id(String) {}
private typedef Row = {id:Int};

/**
	Haxe marks the implementation class of an `extern abstract` as extern, and
	dead code elimination later does the same for every empty implementation
	class. Neither changes how a value is stored, so this stays a String.
**/
private extern abstract ExternText(String) from String to String {}

enum GoTypeMapperContractKind {
	Plain;
}

/** Abstracts whose underlying types lead to Int or Float, or stop before them. */
private enum abstract Level(Int) to Int {
	var Low = 1;
}

private abstract Meters(Float) from Float from Int to Float {}
private abstract Distance(Meters) from Meters to Meters {}
private abstract Wrap<V>(V) {}
private abstract MaybeCount(Null<Int>) {}

/** Field types that `GoTypeMapperContract` asks the type mapper about. */
class GoTypeMapperContractFixture<T:{x:Int}, U> {
	public var id:Id;
	public var row:Row;
	public var empty:{};
	public var structural:T;
	public var nominal:U;
	public var text:String;
	public var externText:ExternText;
	public var kind:GoTypeMapperContractKind;
	public var nullableInt:Null<Int>;
	public var nullableBool:Null<Bool>;
	public var int32:haxe.Int32;
	public var int64:haxe.Int64;
	public var rest:haxe.Rest<Int>;
	public var vector:haxe.ds.Vector<String>;
	public var readOnly:haxe.ds.ReadOnlyArray<Int>;
	public var bytes:haxe.io.BytesData;
	public var slice:go.NativeSlice<Bool>;
	public var anyMap:go.NativeAnyMap;
	public var anyValue:go.NativeAny;
	public var anySlice:go.NativeSlice<go.NativeAny>;
	public var nativeInt64:go.NativeInt64;
	public var nativeUintptr:go.NativeUintptr;
	public var nativeUint32:go.NativeUint32;
	public var nativeUint16:go.NativeUint16;
	public var nativeUint:go.NativeUint;
	public var map:Map<String, Int>;
	public var level:Level;
	public var meters:Meters;
	public var distance:Distance;
	public var maybeMeters:Null<Meters>;
	public var wrappedInt:Wrap<Int>;
	public var wrappedText:Wrap<String>;
	public var maybeCount:MaybeCount;

	public static var created:Int = 0;

	public function mapValue<S>(value:S):S {
		return value;
	}

	/** Reads an instance and a static field so the contract can inspect field references. */
	public function describe():String {
		return text + created;
	}
}
