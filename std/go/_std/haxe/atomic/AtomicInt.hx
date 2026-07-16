package haxe.atomic;

import hxrt.atomic.AtomicIntHandle;
import hxrt.atomic.NativeAtomicInt;

/**
	What
	- Implements the upstream `haxe.atomic.AtomicInt` API as an ordinary staged Haxe
	  abstract over a typed opaque runtime handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is a `@:coreType` declaration with no operation bodies.
	  The target must supply those bodies, but they do not require compiler context.

	How
	- Keep every public operation inline and delegate the native memory operation to
	  `hxrt.atomic.NativeAtomicInt`. The runtime preserves the upstream contract that
	  mutating operations return the value from before the mutation, except `store`,
	  which returns the value stored.
**/
abstract AtomicInt(AtomicIntHandle) {
	public inline function new(value:Int):Void {
		this = NativeAtomicInt.create(value);
	}

	public inline function add(value:Int):Int {
		return NativeAtomicInt.add(this, value);
	}

	public inline function sub(value:Int):Int {
		return NativeAtomicInt.sub(this, value);
	}

	public inline function and(value:Int):Int {
		return NativeAtomicInt.and(this, value);
	}

	public inline function or(value:Int):Int {
		return NativeAtomicInt.or(this, value);
	}

	public inline function xor(value:Int):Int {
		return NativeAtomicInt.xor(this, value);
	}

	public inline function compareExchange(expected:Int, replacement:Int):Int {
		return NativeAtomicInt.compareExchange(this, expected, replacement);
	}

	public inline function exchange(value:Int):Int {
		return NativeAtomicInt.exchange(this, value);
	}

	public inline function load():Int {
		return NativeAtomicInt.load(this);
	}

	public inline function store(value:Int):Int {
		return NativeAtomicInt.store(this, value);
	}
}
