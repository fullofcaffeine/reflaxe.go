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
import go.NativeSlice;
import hxrt.zip.NativeZip;

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
	- Preserve the 64 KiB default and pass positive buffer sizes explicitly to a
	  typed `NativeZip` capability. A negative constructor `windowBits` selects
	  raw DEFLATE; other values select a zlib stream.
	- Retain the established whole-buffer target contract for `execute`, with
	  no-op flush and close methods because no runtime handle is retained.
**/
@:coreApi
class Uncompress {
	var raw:Bool;

	public function new(?windowBits:Int):Void {
		raw = windowBits != null && windowBits < 0;
	}

	public function execute(src:Bytes, srcPos:Int, dst:Bytes, dstPos:Int):{done:Bool, read:Int, write:Int} {
		var input = src.sub(srcPos, src.length - srcPos);
		var bufferSize = dst.length - dstPos;
		if (bufferSize <= 0)
			return {done: false, read: 0, write: 0};
		var data = fromValues(NativeZip.uncompress(toValues(input), raw, bufferSize));
		dst.blit(dstPos, data, 0, data.length);
		return {
			done: true,
			read: input.length,
			write: data.length
		};
	}

	public function setFlushMode(f:FlushMode):Void {}

	public function close():Void {}

	public static function run(src:Bytes, ?bufsize:Int):Bytes {
		var resolvedBufferSize = bufsize == null ? 65536 : bufsize;
		if (resolvedBufferSize <= 0)
			throw 'Invalid zlib buffer size: $resolvedBufferSize';
		return fromValues(NativeZip.uncompress(toValues(src), false, resolvedBufferSize));
	}

	static function toValues(bytes:Bytes):NativeSlice<Int> {
		var values = new Array<Int>();
		for (index in 0...bytes.length)
			values.push(bytes.get(index));
		return NativeSlice.fromArray(values);
	}

	static function fromValues(values:NativeSlice<Int>):Bytes {
		var bytes = Bytes.alloc(values.length);
		for (index in 0...values.length)
			bytes.set(index, values[index]);
		return bytes;
	}
}
