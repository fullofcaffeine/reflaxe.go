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

package haxe.rtti;

import haxe.rtti.CType;

/**
	What
	A staged `haxe.rtti.Rtti` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The public RTTI helper API is portable, but its direct `__rtti` access needs
	to flow through the backend-owned class-token metadata bridge on Go.

	How
	Keep the upstream control flow and error behavior, but read `__rtti` through
	`Reflect.field` and normalize the payload with `Std.string(...)` before
	parsing. That preserves the public Haxe API while using the backend contract
	established in `haxe.go-14as.59`.
**/
class Rtti {
	/**
		Returns the `haxe.rtti.CType.Classdef` corresponding to class `c`.

		If `c` has no runtime type information, e.g. because no `@:rtti` was
		added, an exception of type `String` is thrown.

		If `c` is `null`, the result is unspecified.
	**/
	static public function getRtti<T>(c:Class<T>):Classdef {
		var rtti = Reflect.field(c, "__rtti");
		if (rtti == null) {
			throw 'Class ${Type.getClassName(c)} has no RTTI information, consider adding @:rtti';
		}
		var x = Xml.parse(Std.string(rtti)).firstElement();
		var infos = new haxe.rtti.XmlParser().processElement(x);
		switch (infos) {
			case TClassdecl(c):
				return c;
			case var t:
				throw 'Enum mismatch: expected TClassDecl but found $t';
		}
	}

	/**
		Tells if `c` has runtime type information.

		If `c` is `null`, the result is unspecified.
	**/
	static public function hasRtti<T>(c:Class<T>):Bool {
		for (fieldName in Type.getClassFields(c)) {
			if (fieldName == "__rtti") {
				return true;
			}
		}
		return false;
	}
}
