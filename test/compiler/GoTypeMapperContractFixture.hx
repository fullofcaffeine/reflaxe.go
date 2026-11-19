/** A private abstract that shares its name with one in `GoTypeMapperContractOtherFixture`. */
private abstract Id(String) {}
private typedef Row = {id:Int};

enum GoTypeMapperContractKind {
	Plain;
}

/** Field types that `GoTypeMapperContract` asks the type mapper about. */
class GoTypeMapperContractFixture<T:{x:Int}, U> {
	public var id:Id;
	public var row:Row;
	public var empty:{};
	public var structural:T;
	public var nominal:U;
	public var text:String;
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
	public var nativeInt64:go.NativeInt64;
	public var map:Map<String, Int>;

	public function mapValue<S>(value:S):S {
		return value;
	}
}
