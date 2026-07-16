package hxrt.atomic;

/**
	What
	- Typed bridge to the native integer atomic capabilities used by staged
	  `haxe.atomic.AtomicInt`.

	Why
	- Atomic memory operations require Go's `sync/atomic` implementation, while the
	  public Haxe API and its return-value contract belong in source rather than a
	  compiler-emitted shim.

	How
	- Map each operation one-for-one to an exported `runtime/hxrt/atomic_int.go`
	  function and keep the native cell behind `AtomicIntHandle`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeAtomicInt {
	@:go.name("AtomicIntNew")
	public static function create(value:Int):AtomicIntHandle;

	@:go.name("AtomicIntAdd")
	public static function add(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntSub")
	public static function sub(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntAnd")
	public static function and(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntOr")
	public static function or(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntXor")
	public static function xor(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntCompareExchange")
	public static function compareExchange(handle:AtomicIntHandle, expected:Int, replacement:Int):Int;

	@:go.name("AtomicIntExchange")
	public static function exchange(handle:AtomicIntHandle, value:Int):Int;

	@:go.name("AtomicIntLoad")
	public static function load(handle:AtomicIntHandle):Int;

	@:go.name("AtomicIntStore")
	public static function store(handle:AtomicIntHandle, value:Int):Int;
}
