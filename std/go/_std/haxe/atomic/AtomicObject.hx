package haxe.atomic;

import hxrt.atomic.AtomicObjectHandle;
import hxrt.atomic.NativeAtomicObject;

/**
	What
	- Implements the upstream `haxe.atomic.AtomicObject<T>` API as an ordinary staged
	  Haxe abstract over a typed opaque runtime handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is a `@:coreType` declaration with no operation bodies.
	  The target must supply those bodies, but they do not require compiler context.

	How
	- Keep every public operation inline, delegate atomic storage to
	  `hxrt.atomic.NativeAtomicObject`, and immediately cast the bridge's erased value
	  back to `T`. Reference comparison and synchronization remain runtime-owned.
**/
abstract AtomicObject<T:{}>(AtomicObjectHandle) {
	public inline function new(value:T):Void {
		this = NativeAtomicObject.create(value);
	}

	public inline function compareExchange(expected:T, replacement:T):T {
		return cast NativeAtomicObject.compareExchange(this, expected, replacement);
	}

	public inline function exchange(value:T):T {
		return cast NativeAtomicObject.exchange(this, value);
	}

	public inline function load():T {
		return cast NativeAtomicObject.load(this);
	}

	public inline function store(value:T):T {
		return cast NativeAtomicObject.store(this, value);
	}
}
