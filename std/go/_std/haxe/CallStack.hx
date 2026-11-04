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

#if reflaxe_go_native_stack_trace
import haxe.NativeStackTrace;
#end

/**
	What
	A staged `haxe.CallStack` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	Direct source-owned exception and collection flows can touch `CallStack`
	through `haxe.Exception`-adjacent APIs, but the upstream implementation pulls
	in a broader transitive stdlib graph than `haxe.go` needs for the current
	portable contract. That made otherwise unrelated direct-usage fixtures fail on
	module ownership collisions instead of the actual feature under test.

	How
	Provide the same public `StackItem`/`CallStack` surface with the conservative
	behavior `haxe.go` already uses today for these paths: empty captured stacks
	and deterministic string rendering when explicitly asked. When
	`reflaxe_go_native_stack_trace` is enabled, this staged API converts
	Go-native diagnostic frames from `hxrt` into `StackItem` values. This keeps
	ownership in staged std code and avoids forcing unrelated stdlib modules into
	the build.
**/
enum StackItem {
	CFunction;
	Module(m:String);
	FilePos(s:Null<StackItem>, file:String, line:Int, ?column:Int);
	Method(classname:Null<String>, method:String);
	LocalFunction(?v:Int);
}

@:allow(haxe.Exception)
@:using(haxe.CallStack)
abstract CallStack(Array<StackItem>) from Array<StackItem> {
	public var length(get, never):Int;

	inline function get_length():Int {
		return this.length;
	}

	public static function callStack():Array<StackItem> {
		#if reflaxe_go_native_stack_trace
		return NativeStackTrace.toHaxe(NativeStackTrace.callStack(), 1);
		#else
		return [];
		#end
	}

	public static function exceptionStack(fullStack = false):Array<StackItem> {
		#if reflaxe_go_native_stack_trace
		return NativeStackTrace.toHaxe(NativeStackTrace.exceptionStack(), fullStack ? 0 : 1);
		#else
		return [];
		#end
	}

	public static function toString(stack:CallStack):String {
		var out = "";
		for (item in stack.asArray()) {
			out += "\nCalled from " + itemToString(item);
		}
		return out;
	}

	public function subtract(stack:CallStack):CallStack {
		var startIndex = -1;
		var i = -1;
		while (++i < this.length) {
			for (j in 0...stack.length) {
				if (equalItems(this[i], stack[j])) {
					if (startIndex < 0) {
						startIndex = i;
					}
					++i;
					if (i >= this.length)
						break;
				} else {
					startIndex = -1;
				}
			}
			if (startIndex >= 0)
				break;
		}
		return startIndex >= 0 ? this.slice(0, startIndex) : this;
	}

	public inline function copy():CallStack {
		return this.copy();
	}

	@:arrayAccess public inline function get(index:Int):StackItem {
		return this[index];
	}

	inline function asArray():Array<StackItem> {
		return this;
	}

	static function equalItems(item1:Null<StackItem>, item2:Null<StackItem>):Bool {
		return switch ([item1, item2]) {
			case [null, null]: true;
			case [CFunction, CFunction]: true;
			case [Module(m1), Module(m2)]:
				m1 == m2;
			case [FilePos(item1, file1, line1, col1), FilePos(item2, file2, line2, col2)]: file1 == file2 && line1 == line2 && col1 == col2 && equalItems(item1,
					item2);
			case [Method(class1, method1), Method(class2, method2)]: class1 == class2 && method1 == method2;
			case [LocalFunction(v1), LocalFunction(v2)]:
				v1 == v2;
			case _: false;
		}
	}

	static function exceptionToString(e:Exception):String {
		if (e.previous == null) {
			return 'Exception: ${e.toString()}${e.stack}';
		}
		var result = '';
		var e:Null<Exception> = e;
		var prev:Null<Exception> = null;
		while (e != null) {
			if (prev == null) {
				result = 'Exception: ${e.message}${e.stack}' + result;
			} else {
				var prevStack = @:privateAccess e.stack.subtract(prev.stack);
				result = 'Exception: ${e.message}${prevStack}\n\nNext ' + result;
			}
			prev = e;
			e = e.previous;
		}
		return result;
	}

	static function itemToString(item:StackItem):String {
		return switch (item) {
			case CFunction:
				"a C function";
			case Module(m):
				"module " + m;
			case FilePos(inner, file, line, column):
				var rendered = inner == null ? file : itemToString(inner) + " (" + file;
				rendered += " line " + line;
				if (column > 0) {
					rendered += " column " + column;
				}
				if (inner != null) {
					rendered += ")";
				}
				rendered;
			case Method(classname, method):
				(classname == null ? "<unknown>" : classname) + "." + method;
			case LocalFunction(v):
				"local function #" + Std.string(v);
		}
	}
}
