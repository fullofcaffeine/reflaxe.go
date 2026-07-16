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

import hxrt.math.NativeMath;
import hxrt.math.NativeMathInt;
import hxrt.math.NativeRandom;

/**
	What:
	- Implements the complete Haxe 4.3.7 `Math` API as staged source for the Go
	  target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because `Math` is an extern contract whose native functions and
	  IEEE-754 constants must be bound by the target.
	- Haxe-specific rounding, finiteness, NaN propagation, and signed-zero policy
	  are source semantics rather than compiler call-matching rules.

	How:
	- Delegate one-for-one numeric operations to typed Go `math` and `math/rand`
	  externs, then express Haxe policy in ordinary source where signatures or
	  edge behavior differ.
	- Keep equal-value `min` / `max` selection in Haxe so the reference runtime's
	  operand-order signed-zero behavior is preserved exactly.
**/
@:coreApi
@:pure
class Math {
	public static var PI(default, null):Float = 3.141592653589793;
	public static var NEGATIVE_INFINITY(default, null):Float = NativeMath.inf(-1);
	public static var POSITIVE_INFINITY(default, null):Float = NativeMath.inf(1);
	public static var NaN(default, null):Float = NativeMath.nan();

	public static inline function abs(v:Float):Float {
		return NativeMath.abs(v);
	}

	/** Preserves Haxe's NaN propagation and second-operand signed-zero rule. **/
	public static function min(a:Float, b:Float):Float {
		if (isNaN(a))
			return a;
		if (isNaN(b))
			return b;
		return a < b ? a : b;
	}

	/** Preserves Haxe's NaN propagation and first-operand equal-value rule. **/
	public static function max(a:Float, b:Float):Float {
		if (isNaN(a))
			return a;
		if (isNaN(b))
			return b;
		return a < b ? b : a;
	}

	public static inline function sin(v:Float):Float {
		return NativeMath.sin(v);
	}

	public static inline function cos(v:Float):Float {
		return NativeMath.cos(v);
	}

	public static inline function tan(v:Float):Float {
		return NativeMath.tan(v);
	}

	public static inline function asin(v:Float):Float {
		return NativeMath.asin(v);
	}

	public static inline function acos(v:Float):Float {
		return NativeMath.acos(v);
	}

	public static inline function atan(v:Float):Float {
		return NativeMath.atan(v);
	}

	public static inline function atan2(y:Float, x:Float):Float {
		return NativeMath.atan2(y, x);
	}

	public static inline function exp(v:Float):Float {
		return NativeMath.exp(v);
	}

	public static inline function log(v:Float):Float {
		return NativeMath.log(v);
	}

	public static inline function pow(v:Float, exp:Float):Float {
		return NativeMath.pow(v, exp);
	}

	public static inline function sqrt(v:Float):Float {
		return NativeMath.sqrt(v);
	}

	/** Implements Haxe's ties-up rule, including `-0.5 -> 0`. **/
	public static inline function round(v:Float):Int {
		return NativeMathInt.round(v);
	}

	public static inline function floor(v:Float):Int {
		return NativeMathInt.floor(v);
	}

	public static inline function ceil(v:Float):Int {
		return NativeMathInt.ceil(v);
	}

	public static inline function random():Float {
		return NativeRandom.float64();
	}

	public static inline function ffloor(v:Float):Float {
		return NativeMath.floor(v);
	}

	public static inline function fceil(v:Float):Float {
		return NativeMath.ceil(v);
	}

	public static inline function fround(v:Float):Float {
		return NativeMath.floor(v + 0.5);
	}

	public static inline function isFinite(f:Float):Bool {
		return !NativeMath.isInf(f, 0) && !NativeMath.isNaN(f);
	}

	public static inline function isNaN(f:Float):Bool {
		return NativeMath.isNaN(f);
	}
}
