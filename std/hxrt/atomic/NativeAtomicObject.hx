package hxrt.atomic;

/**
	What
	- Typed-handle bridge to the native object atomic capabilities used by staged
	  `haxe.atomic.AtomicObject`.

	Why
	- Atomic reference storage and comparison require Go runtime support. The stored
	  value is deliberately `Dynamic` only at this narrow bridge because one runtime
	  cell must carry every `AtomicObject<T>` instantiation after Haxe generic erasure.

	How
	- Keep the cell itself typed as `AtomicObjectHandle`, map operations one-for-one
	  to `runtime/hxrt/atomic_object.go`, and let the staged generic wrapper cast each
	  returned value immediately back to its declared `T`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeAtomicObject {
	@:go.name("AtomicObjectNew")
	public static function create(value:Dynamic):AtomicObjectHandle;

	@:go.name("AtomicObjectCompareExchange")
	public static function compareExchange(handle:AtomicObjectHandle, expected:Dynamic, replacement:Dynamic):Dynamic;

	@:go.name("AtomicObjectExchange")
	public static function exchange(handle:AtomicObjectHandle, value:Dynamic):Dynamic;

	@:go.name("AtomicObjectLoad")
	public static function load(handle:AtomicObjectHandle):Dynamic;

	@:go.name("AtomicObjectStore")
	public static function store(handle:AtomicObjectHandle, value:Dynamic):Dynamic;
}
