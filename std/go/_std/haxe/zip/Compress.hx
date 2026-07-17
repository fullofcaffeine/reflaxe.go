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
	- Implements `haxe.zip.Compress` as staged Haxe source for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because its portable compressor intentionally throws
	  `NotImplementedException`; zlib execution requires a target runtime
	  capability.
	- Compression levels, Haxe `Bytes` ownership, and the public one-shot API are
	  library behavior and must not be emitted as compiler-owned Go functions.

	How:
	- Retain the whole-buffer `execute` contract used by established Haxe targets:
	  consume the remaining source, write one complete zlib stream, and keep
	  `setFlushMode` / `close` as no-ops because no streaming handle is retained.
	- Convert `Bytes` to integer values in Haxe and delegate only zlib execution to
	  the typed `NativeZip` capability.
**/
@:coreApi
class Compress {
	var level:Int;

	public function new(level:Int):Void {
		validateLevel(level);
		this.level = level;
	}

	public function execute(src:Bytes, srcPos:Int, dst:Bytes, dstPos:Int):{done:Bool, read:Int, write:Int} {
		var input = src.sub(srcPos, src.length - srcPos);
		var data = fromValues(NativeZip.compress(toValues(input), level));
		dst.blit(dstPos, data, 0, data.length);
		return {
			done: true,
			read: input.length,
			write: data.length
		};
	}

	public function setFlushMode(f:FlushMode):Void {}

	public function close():Void {}

	public static function run(s:Bytes, level:Int):Bytes {
		validateLevel(level);
		return fromValues(NativeZip.compress(toValues(s), level));
	}

	static function validateLevel(level:Int):Void {
		if (level < -1 || level > 9)
			throw 'Invalid zlib compression level: $level';
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
