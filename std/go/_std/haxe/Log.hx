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

package haxe;

/**
	What:
	- Implements the complete Haxe 4.3.7 `haxe.Log` formatting and mutable
	  `trace` API as canonical staged source for `haxe.go`.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because its output branch is selected through ambient target
	  conditionals; the staged Go owner must bind `Sys.println` explicitly.
	- Position formatting, custom parameters, direct function values, and dynamic
	  rebinding are public library behavior; a direct compiler rewrite to
	  `hxrt.Println` loses all four contracts.

	How:
	- Preserve the upstream formatting algorithm and mutable static dynamic method.
	- Send the final formatted string through source-owned `Sys.println`, which is
	  already the portable console boundary for this target.
**/
@:coreApi
class Log {
	/**
		What: Formats one trace value with optional source position and custom values.
		Why: This text shape is portable `haxe.Log` behavior observed by user code and
		custom trace sinks, not a target printer detail.
		How: Preserve the upstream field order and stringify every value through `Std`.
	**/
	public static function formatOutput(v:Dynamic, infos:Null<PosInfos>):String {
		var str = Std.string(v);
		if (infos == null)
			return str;
		var pstr = infos.fileName + ":" + infos.lineNumber;
		if (infos.customParams != null)
			for (v in infos.customParams)
				str += ", " + Std.string(v);
		return pstr + ": " + str;
	}

	/**
		What: Provides the mutable default trace sink used by Haxe's `trace` expression.
		Why: Callers may take, replace, restore, or null this function value; lowering it
		directly to a fixed print call would erase those public semantics.
		How: Format in source and delegate console output to the staged `Sys.println`
		boundary. Generated dynamic-function calls route null failures through Haxe
		exception semantics.
	**/
	public static dynamic function trace(v:Dynamic, ?infos:PosInfos):Void {
		var str = formatOutput(v, infos);
		Sys.println(str);
	}
}
