package haxe.io;

/**
	What
	- Source-owned helper surface for the inherited `haxe.io.Input` and
	  `haxe.io.Output` loop helpers on `haxe.go`.
	- Owns the generic stream-control helpers such as `readAll`, `readLine`,
	  `readUntil`, `readFullBytes`, `write`, `writeFullBytes`, `writeInput`,
	  and `writeString`.

	Why
	- The mainstream Haxe stdlib helper logic could not be reused unchanged here
	  because `haxe.go` still owns the generated Go type shapes for the base IO
	  surface in `GoCompiler`, and those helper bodies were historically emitted
	  there as raw compiler declarations instead of normal source-owned code.
	- Keeping those generic loops in the compiler is bloat: they do not depend on
	  profile resolution or backend-only representation policy in the same way the
	  RawNative/string-cache bytes paths still do.

	How
	- Keep the generated `haxe__io__input_*` and `haxe__io__output_*` wrappers as
	  thin compiler forwarders.
	- Implement the loop bodies here in ordinary Haxe so they lower in the same
	  generated package as the rest of the IO surface.
	- Use narrow framework-owned raw `__go__` writes for `Bytes` / `Input` /
	  `Output` operations that the current staged Haxe surface does not expose
	  cleanly, while still keeping the generic loop ownership out of `GoCompiler`.
	  The generated Go `haxe__io__Bytes` backing storage is real, but this backend
	  does not currently expose a reusable Haxe-level mutation surface for the
	  staged helper layer.
	- Pull this helper class in only when the IO helper surface is actually
	  required.
**/
@:goAllowRaw
@:keep
class GoIoHelpers {
	/**
		What
		- Returns the staged `BytesOutput` contents as `Bytes`.

		Why
		- The mainstream `BytesOutput.getBytes()` API exists, but the current
		  `haxe.go` compile path does not always expose that method at the staged
		  helper typing surface even though the generated Go method exists.

		How
		- Bridge to the generated Go method through the framework-owned raw helper
		  path instead of keeping the whole IO loop body in `GoCompiler`.
	**/
	static inline function bytesOutputGetBytes(out:BytesOutput):Bytes {
		return untyped __go__("{0}.getBytes()", out);
	}

	/**
		What
		- Calls the generated `Input.readBytes` method from the staged helper layer.

		Why
		- The generated Go IO contract exposes `readBytes`, but the current staged
		  Haxe typing surface does not consistently expose that method on
		  `haxe.io.Input` across all compile contexts.

		How
		- Bridge to the generated method through the framework-owned raw helper
		  path so `GoIoHelpers` can preserve the original `Blocked` semantics
		  without moving the loop back into `GoCompiler`.
	**/
	static inline function inputReadBytes(self:Input, buf:Bytes, pos:Int, len:Int):Int {
		return untyped __go__("{0}.readBytes({1}, {2}, {3})", self, buf, pos, len);
	}

	/**
		What
		- Calls the generated `Output.writeBytes` method from the staged helper layer.

		Why
		- The generated Go IO contract exposes `writeBytes`, but the staged Haxe
		  typing surface does not consistently expose that method on
		  `haxe.io.Output` in every compile context.

		How
		- Bridge to the generated method through the framework-owned raw helper
		  path so `GoIoHelpers` can preserve `Blocked` semantics without restoring
		  compiler-owned raw loop bodies.
	**/
	static inline function outputWriteBytes(self:Output, buf:Bytes, pos:Int, len:Int):Int {
		return untyped __go__("{0}.writeBytes({1}, {2}, {3})", self, buf, pos, len);
	}

	public static function inputReadAll(self:Input, bufsize:Int):Bytes {
		if (self == null) {
			return Bytes.alloc(0);
		}
		var buf = Bytes.alloc(bufsize);
		var total = new BytesOutput();
		while (true) {
			var done = false;
			try {
				var chunk = inputReadBytes(self, buf, 0, bufsize);
				if (chunk == 0) {
					throw Error.Blocked;
				}
				outputWriteFullBytes(total, buf, 0, chunk);
			} catch (_:Eof) {
				done = true;
			}
			if (done) {
				break;
			}
		}
		return bytesOutputGetBytes(total);
	}

	public static function inputReadFullBytes(self:Input, s:Bytes, pos:Int, len:Int):Void {
		if (self == null) {
			throw Error.Blocked;
		}
		while (len > 0) {
			var read = inputReadBytes(self, s, pos, len);
			if (read == 0) {
				throw Error.Blocked;
			}
			pos += read;
			len -= read;
		}
	}

	public static function inputRead(self:Input, nbytes:Int):Bytes {
		var out = Bytes.alloc(nbytes);
		inputReadFullBytes(self, out, 0, nbytes);
		return out;
	}

	public static function inputReadUntil(self:Input, end:Int):String {
		var buf = new BytesOutput();
		while (true) {
			var last = self.readByte();
			if (last == end) {
				break;
			}
			untyped __go__("func() int { {0}.writeByte({1}); return 0 }()", buf, last);
		}
		return bytesOutputGetBytes(buf).toString();
	}

	public static function inputReadLine(self:Input):String {
		var buf = new BytesOutput();
		while (true) {
			var last = 0;
			var ended = false;
			try {
				last = self.readByte();
				if (last == 10) {
					ended = true;
				} else {
					untyped __go__("func() int { {0}.writeByte({1}); return 0 }()", buf, last);
				}
			} catch (e:Eof) {
				var partial = bytesOutputGetBytes(buf).toString();
				if (partial.length == 0) {
					throw e;
				}
				return partial;
			}
			if (ended) {
				break;
			}
		}
		var out = bytesOutputGetBytes(buf).toString();
		if (out.length > 0 && out.charCodeAt(out.length - 1) == 13) {
			out = out.substr(0, out.length - 1);
		}
		return out;
	}

	public static function outputWrite(self:Output, s:Bytes):Void {
		if (self == null || s == null) {
			return;
		}
		var remaining = s.length;
		var pos = 0;
		while (remaining > 0) {
			var wrote = outputWriteBytes(self, s, pos, remaining);
			if (wrote == 0) {
				throw Error.Blocked;
			}
			pos += wrote;
			remaining -= wrote;
		}
	}

	public static function outputWriteFullBytes(self:Output, s:Bytes, pos:Int, len:Int):Void {
		while (len > 0) {
			var wrote = outputWriteBytes(self, s, pos, len);
			if (wrote == 0) {
				throw Error.Blocked;
			}
			pos += wrote;
			len -= wrote;
		}
	}

	public static function outputWriteInput(self:Output, i:Input, bufsize:Int):Void {
		if (self == null || i == null) {
			return;
		}
		var buf = Bytes.alloc(bufsize);
		while (true) {
			var done = false;
			try {
				var lenRead = inputReadBytes(i, buf, 0, bufsize);
				if (lenRead == 0) {
					throw Error.Blocked;
				}
				outputWriteFullBytes(self, buf, 0, lenRead);
			} catch (_:Eof) {
				done = true;
			}
			if (done) {
				break;
			}
		}
	}

	public static function outputWriteString(self:Output, s:String, encoding:Null<Encoding>):Void {
		if (s == null) {
			s = "";
		}
		var bytes = encoding == null ? Bytes.ofString(s) : Bytes.ofString(s, encoding);
		outputWriteFullBytes(self, bytes, 0, bytes.length);
	}
}
