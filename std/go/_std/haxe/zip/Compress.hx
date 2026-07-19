/*
 * Copyright (C)2005-2019 Haxe Foundation
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */

package haxe.zip;

import haxe.io.Bytes;
import haxe.io.Error;
import go.NativeSlice;
import hxrt.zip.NativeZip;
import hxrt.zip.ZipDeflateHandle;

/**
	What:
	- Implements `haxe.zip.Compress` as staged Haxe source for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because its portable compressor intentionally throws
	  `NotImplementedException`; zlib execution requires a target runtime
	  capability.
	- Compression levels, Haxe `Bytes` ownership, and the public one-shot API are
	  library behavior and must not be emitted as compiler-owned Go functions.

	How:
	- Retain an opaque typed deflate handle across calls, convert only the
	  remaining source bytes, and ask `NativeZip` for no more output than fits in
	  the destination. Haxe performs the final blit and public result assembly.
	- Support NO, SYNC, and FINISH exactly. Go's standard compressor exposes no
	  honest equivalent for FULL's dictionary reset or BLOCK's boundary stop, so
	  selecting either fails explicitly instead of silently changing semantics.
	- Make close idempotent and reject later mutation or execution in staged
	  source, with the runtime independently defending the opaque handle.
**/
@:coreApi
class Compress {
	var handle:ZipDeflateHandle;
	var flushMode:Int;
	var closed:Bool;

	public function new(level:Int):Void {
		validateLevel(level);
		handle = NativeZip.createDeflate(level);
		flushMode = NativeZip.FLUSH_NO;
		closed = false;
	}

	public function execute(src:Bytes, srcPos:Int, dst:Bytes, dstPos:Int):{done:Bool, read:Int, write:Int} {
		ensureOpen();
		validatePosition(srcPos, src.length);
		validatePosition(dstPos, dst.length);
		var outputLimit = dst.length - dstPos;
		if (outputLimit == 0)
			return {done: false, read: 0, write: 0};
		var step = NativeZip.executeDeflate(handle, toValuesFrom(src, srcPos), outputLimit, flushMode);
		var write = writeValues(dst, dstPos, step.values);
		return {
			done: step.done,
			read: step.read,
			write: write
		};
	}

	public function setFlushMode(f:FlushMode):Void {
		ensureOpen();
		flushMode = flushModeCode(f);
	}

	public function close():Void {
		if (closed)
			return;
		closed = true;
		NativeZip.closeDeflate(handle);
	}

	public static function run(s:Bytes, level:Int):Bytes {
		validateLevel(level);
		return fromValues(NativeZip.compress(toValues(s), level));
	}

	static function validateLevel(level:Int):Void {
		if (level < -1 || level > 9)
			throw 'Invalid zlib compression level: $level';
	}

	/**
		What: Validate one public source or destination cursor.
		Why: Native slices must never receive a position outside its owning Haxe
		`Bytes`; otherwise target-specific slice failures would leak through.
		How: Accept the end cursor for empty/drain calls and use Haxe's canonical
		`OutsideBounds` error for every other invalid position.
	**/
	static function validatePosition(position:Int, length:Int):Void {
		if (position < 0 || position > length)
			throw Error.OutsideBounds;
	}

	/**
		What: Translate portable flush constructors to the typed runtime protocol.
		Why: Go exposes exact NO, SYNC, and FINISH operations but no FULL dictionary
		reset or BLOCK boundary operation.
		How: Use stable private integer codes for supported modes and fail at the
		public setter for unsupported semantics.
	**/
	static function flushModeCode(mode:FlushMode):Int {
		return switch (mode) {
			case NO: NativeZip.FLUSH_NO;
			case SYNC: NativeZip.FLUSH_SYNC;
			case FINISH: NativeZip.FLUSH_FINISH;
			case FULL: throw "haxe.zip.FlushMode.FULL is not supported by Go's standard compressor";
			case BLOCK: throw "haxe.zip.FlushMode.BLOCK is not supported by Go's standard compressor";
		};
	}

	/**
		What: Guard methods that require a live native compressor.
		Why: A closed opaque handle cannot safely accept more input or policy changes.
		How: Keep the public lifecycle error deterministic before crossing into Go.
	**/
	function ensureOpen():Void {
		if (closed)
			throw "haxe.zip.Compress is closed";
	}

	static function toValues(bytes:Bytes):NativeSlice<Int> {
		return toValuesFrom(bytes, 0);
	}

	/**
		What: Copy the remaining Haxe bytes into the typed native slice boundary.
		Why: Constructing an intermediate `Bytes.sub` would copy the same input twice
		for every progressive call.
		How: Traverse directly from the validated source cursor to the end.
	**/
	static function toValuesFrom(bytes:Bytes, position:Int):NativeSlice<Int> {
		var values = new Array<Int>();
		for (index in position...bytes.length)
			values.push(bytes.get(index));
		return NativeSlice.fromArray(values);
	}

	/**
		What: Copy one bounded native output step into its Haxe destination.
		Why: The staged layer owns `Bytes`, but allocating a temporary `Bytes` value
		before every partial write would add a redundant copy.
		How: Store each typed integer value at the already-validated destination
		cursor and return the public write count.
	**/
	static function writeValues(destination:Bytes, position:Int, values:NativeSlice<Int>):Int {
		for (index in 0...values.length)
			destination.set(position + index, values[index]);
		return values.length;
	}

	static function fromValues(values:NativeSlice<Int>):Bytes {
		var bytes = Bytes.alloc(values.length);
		for (index in 0...values.length)
			bytes.set(index, values[index]);
		return bytes;
	}
}
