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
import hxrt.zip.ZipInflateHandle;

/**
	What:
	- Implements `haxe.zip.Uncompress` as staged Haxe source for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go`: its generic static path is a Haxe inflater, while the target API
	  also needs native zlib and negative-window raw-DEFLATE execution used by
	  `haxe.zip.Tools`.
	- Optional buffer defaults, Haxe `Bytes` conversion, and the public one-shot
	  instance contract are source-level policy rather than compiler behavior.

	How:
	- A negative constructor `windowBits` selects raw DEFLATE; other values select
	  zlib. Retain the resulting opaque typed inflater across partial calls.
	- Pass only remaining integer byte values plus available destination capacity
	  to `NativeZip`; Haxe owns offsets, destination writes, and the public result.
	- Support NO, SYNC, and FINISH policy. FULL and BLOCK are rejected explicitly
	  because Go's standard inflater cannot promise their zlib boundary behavior.
	- Preserve the 64 KiB static-run default while making instance close
	  idempotent and use-after-close deterministic.
**/
@:coreApi
class Uncompress {
	var raw:Bool;
	var handle:ZipInflateHandle;
	var flushMode:Int;
	var closed:Bool;

	public function new(?windowBits:Int):Void {
		raw = windowBits != null && windowBits < 0;
		handle = NativeZip.createInflate(raw);
		flushMode = NativeZip.FLUSH_NO;
		closed = false;
	}

	public function execute(src:Bytes, srcPos:Int, dst:Bytes, dstPos:Int):{done:Bool, read:Int, write:Int} {
		ensureOpen();
		validatePosition(srcPos, src.length);
		validatePosition(dstPos, dst.length);
		var bufferSize = dst.length - dstPos;
		if (bufferSize == 0)
			return {done: false, read: 0, write: 0};
		var step = NativeZip.executeInflate(handle, toValuesFrom(src, srcPos), bufferSize, flushMode);
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
		NativeZip.closeInflate(handle);
	}

	public static function run(src:Bytes, ?bufsize:Int):Bytes {
		var resolvedBufferSize = bufsize == null ? 65536 : bufsize;
		if (resolvedBufferSize <= 0)
			throw 'Invalid zlib buffer size: $resolvedBufferSize';
		return fromValues(NativeZip.uncompress(toValues(src), false, resolvedBufferSize));
	}

	/**
		What: Validate one public source or destination cursor.
		Why: Partial native execution must not receive a position outside its owning
		Haxe `Bytes` value.
		How: Permit the end cursor for drain calls and otherwise throw the canonical
		Haxe `OutsideBounds` error.
	**/
	static function validatePosition(position:Int, length:Int):Void {
		if (position < 0 || position > length)
			throw Error.OutsideBounds;
	}

	/**
		What: Translate portable flush constructors to the native integer protocol.
		Why: Go's inflater cannot expose exact FULL or BLOCK boundary behavior.
		How: Map supported modes and reject unsupported semantics at policy selection
		rather than silently downgrading them.
	**/
	static function flushModeCode(mode:FlushMode):Int {
		return switch (mode) {
			case NO: NativeZip.FLUSH_NO;
			case SYNC: NativeZip.FLUSH_SYNC;
			case FINISH: NativeZip.FLUSH_FINISH;
			case FULL: throw "haxe.zip.FlushMode.FULL is not supported by Go's standard inflater";
			case BLOCK: throw "haxe.zip.FlushMode.BLOCK is not supported by Go's standard inflater";
		};
	}

	/**
		What: Guard methods that require a live native inflater.
		Why: Retained native state is released by `close` and cannot be reused safely.
		How: Throw one deterministic public lifecycle error before crossing into Go.
	**/
	function ensureOpen():Void {
		if (closed)
			throw "haxe.zip.Uncompress is closed";
	}

	static function toValues(bytes:Bytes):NativeSlice<Int> {
		return toValuesFrom(bytes, 0);
	}

	/**
		What: Copy source bytes from one validated cursor into the native boundary.
		Why: Progressive execution should not allocate and copy an intermediate
		`Bytes.sub` before converting the same remaining input.
		How: Read directly from the cursor through the source end.
	**/
	static function toValuesFrom(bytes:Bytes, position:Int):NativeSlice<Int> {
		var values = new Array<Int>();
		for (index in position...bytes.length)
			values.push(bytes.get(index));
		return NativeSlice.fromArray(values);
	}

	/**
		What: Write one bounded native decode step into its Haxe destination.
		Why: Haxe owns `Bytes`, but a temporary output `Bytes` plus `blit` would copy
		every progressive result twice.
		How: Store the typed native values at the validated destination cursor and
		return their count.
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
