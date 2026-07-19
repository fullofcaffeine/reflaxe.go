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

import hxrt.math.NativeMathInt;
import hxrt.math.NativeRandom;
import hxrt.string.GoStringRuntime;

/**
	What:
	- Implements the complete Haxe 4.3.7 `Std` API as canonical staged source for
	  `haxe.go`.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because its `Std` declaration is an extern contract: each target
	  must supply parsing, numeric conversion, random selection, type tests, and
	  erased-value string conversion.
	- Prefix scanning, overflow policy, deprecated aliases, and downcast behavior
	  are public Haxe semantics and must remain reviewable source instead of
	  compiler call rewrites.

	How:
	- Express integer and floating-prefix parsing, aliases, downcasts, and random
	  bounds in ordinary Haxe.
	- Delegate only exact native float parsing/conversion and random generation to
	  typed bindings.
	- Keep `string` and `isOfType` as the two individually registered compiler
	  intrinsics because they consume erased target representation or a typed type
	  literal that ordinary source cannot recover.
**/
@:coreApi
class Std {
	@:deprecated('Std.is is deprecated. Use Std.isOfType instead.')
	public static inline function is(v:Dynamic, t:Dynamic):Bool {
		return isOfType(v, t);
	}

	/**
		What: Tests an erased value against one Haxe runtime type token.
		Why: Ordinary source cannot inspect the target's generated class carriers or
		core scalar representations without losing type safety.
		How: Keep this single operation as a registered typed compiler intrinsic; all
		aliases and downcast policy remain ordinary source below.
	**/
	extern public static function isOfType(v:Dynamic, t:Dynamic):Bool;

	/**
		What: Returns `value` as the requested subclass when its runtime token matches.
		Why: Haxe downcasts are nullable and must recover the canonical generated
		virtual receiver when a subclass arrived through a base-class carrier.
		How: Reuse `isOfType` for the nominal decision and perform the Haxe-proven cast
		only on the successful branch.
	**/
	public static inline function downcast<T:{}, S:T>(value:T, c:Class<S>):Null<S> {
		return isOfType(value, c) ? cast value : null;
	}

	@:deprecated('Std.instance() is deprecated. Use Std.downcast() instead.')
	public static inline function instance<T:{}, S:T>(value:T, c:Class<S>):Null<S> {
		return downcast(value, c);
	}

	/**
		What: Produces the canonical Haxe string representation of an erased value.
		Why: The representation depends on generated target carriers that staged Haxe
		source cannot inspect safely.
		How: Retain this exact operation as the second registered `Std` intrinsic.
	**/
	extern public static function string(s:Dynamic):String;

	/**
		What: Truncates a finite floating-point value toward zero as a Haxe `Int`.
		Why: Haxe source cannot spell Go's native numeric conversion directly.
		How: Delegate only the representation conversion to the typed `hxrt.math`
		binding; values outside Haxe's specified finite Int32 range remain unspecified.
	**/
	public static inline function int(x:Float):Int {
		return NativeMathInt.truncate(x);
	}

	/**
		What: Parses Haxe's signed decimal or hexadecimal integer prefix contract.
		Why: Whitespace, accepted prefixes, stop position, Int32 overflow, and null
		behavior are public library semantics and do not belong in a Go runtime parser.
		How: Scan and bounds-check in staged Haxe, accumulating decimal values
		negatively so `-2147483648` remains representable.
	**/
	public static function parseInt(x:String):Null<Int> {
		if (x == null)
			return null;

		var index = 0;
		var length = x.length;
		while (index < length && isSpaceCode(x.charCodeAt(index)))
			index++;

		var negative = false;
		if (index < length) {
			var sign = x.charCodeAt(index);
			if (sign == '-'.code || sign == '+'.code) {
				negative = sign == '-'.code;
				index++;
			}
		}

		var hexadecimal = index + 1 < length
			&& x.charCodeAt(index) == '0'.code
			&& (x.charCodeAt(index + 1) == 'x'.code || x.charCodeAt(index + 1) == 'X'.code);
		if (hexadecimal)
			index += 2;

		var digitStart = index;
		if (hexadecimal) {
			var value = 0;
			var significantDigits = 0;
			var sawNonZero = false;
			while (index < length) {
				var digit = digitValue(x.charCodeAt(index), true);
				if (digit < 0)
					break;
				if (digit != 0 || sawNonZero) {
					sawNonZero = true;
					significantDigits++;
					if (significantDigits > 8)
						return null;
				}
				value = (value << 4) | digit;
				index++;
			}
			if (index == digitStart)
				return null;
			return negative ? -value : value;
		}

		// Accumulate negatively so -2147483648 remains representable while the
		// positive limit can still reject 2147483648 before arithmetic wraps.
		var result = 0;
		while (index < length) {
			var digit = digitValue(x.charCodeAt(index), false);
			if (digit < 0)
				break;
			var lastAllowedDigit = negative ? 8 : 7;
			if (result < -214748364 || (result == -214748364 && digit > lastAllowedDigit))
				return null;
			result = result * 10 - digit;
			index++;
		}
		if (index == digitStart)
			return null;
		return negative ? result : -result;
	}

	/**
		What: Parses the longest valid signed decimal/exponent prefix as a Haxe Float.
		Why: Prefix and malformed-exponent policy are portable library behavior, while
		exact IEEE-754 conversion is native representation work.
		How: Validate the token in source, then pass only that exact substring through
		the narrow typed string-runtime conversion.
	**/
	public static function parseFloat(x:String):Float {
		if (x == null)
			return invalidFloat();

		var index = 0;
		var length = x.length;
		while (index < length && isSpaceCode(x.charCodeAt(index)))
			index++;
		var start = index;

		if (index < length) {
			var sign = x.charCodeAt(index);
			if (sign == '-'.code || sign == '+'.code)
				index++;
		}

		var digits = 0;
		while (index < length && isDecimalDigit(x.charCodeAt(index))) {
			digits++;
			index++;
		}
		if (index < length && x.charCodeAt(index) == '.'.code) {
			index++;
			while (index < length && isDecimalDigit(x.charCodeAt(index))) {
				digits++;
				index++;
			}
		}
		if (digits == 0)
			return invalidFloat();

		if (index < length && (x.charCodeAt(index) == 'e'.code || x.charCodeAt(index) == 'E'.code)) {
			index++;
			if (index < length && (x.charCodeAt(index) == '-'.code || x.charCodeAt(index) == '+'.code))
				index++;
			var exponentStart = index;
			while (index < length && isDecimalDigit(x.charCodeAt(index)))
				index++;
			if (index == exponentStart)
				return invalidFloat();
		}

		return GoStringRuntime.parseFloatExact(x.substring(start, index));
	}

	/**
		What: Returns a pseudo-random integer in `[0, x)`, or zero when `x <= 1`.
		Why: Bound policy is Haxe API behavior; random-number generation is a native
		standard-library capability.
		How: Guard the bound in source and call the typed Go random binding only when
		`math/rand.Intn` can accept it.
	**/
	public static inline function random(x:Int):Int {
		return x <= 1 ? 0 : NativeRandom.intn(x);
	}

	static inline function invalidFloat():Float {
		return GoStringRuntime.parseFloatExact("");
	}

	static inline function isSpaceCode(code:Int):Bool {
		return (code >= 9 && code <= 13) || code == 32;
	}

	static inline function isDecimalDigit(code:Int):Bool {
		return code >= '0'.code && code <= '9'.code;
	}

	static inline function digitValue(code:Int, hexadecimal:Bool):Int {
		if (isDecimalDigit(code))
			return code - '0'.code;
		if (hexadecimal && code >= 'a'.code && code <= 'f'.code)
			return code - 'a'.code + 10;
		if (hexadecimal && code >= 'A'.code && code <= 'F'.code)
			return code - 'A'.code + 10;
		return -1;
	}
}
