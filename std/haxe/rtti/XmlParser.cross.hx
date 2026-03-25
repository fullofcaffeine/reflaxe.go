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

import haxe.ds.ArraySort;
import haxe.rtti.CType;

/**
	What
	A staged `haxe.rtti.XmlParser` override for `haxe.go`.

	Why
	The upstream parser is the correct semantic owner for RTTI XML processing, but
	it uses a few source shapes that `haxe.go` does not lower cleanly unchanged:
	instance `Array.sort(...)` calls and direct ordering on Go-lowered `String`
	values.

	How
	Keep the upstream parser structure and merge logic, but route sorting through
	`haxe.ds.ArraySort.sort(...)` and compare names with `Reflect.compare(...)`.
	The parser still operates on the ordinary staged `CType` / `TypeInfos`
	structures while relying on the backend-owned anonymous-record mutation fix
	below this layer.
**/
class XmlParser {
	public var root:TypeRoot;

	var curplatform:String;

	public function new() {
		root = new Array();
	}

	public function sort(?l:TypeRoot) {
		if (l == null)
			l = root;
		ArraySort.sort(l, function(e1, e2) {
			var n1 = switch e1 {
				case TPackage(p, _, _): " " + p;
				default: TypeApi.typeInfos(e1).path;
			};
			var n2 = switch e2 {
				case TPackage(p, _, _): " " + p;
				default: TypeApi.typeInfos(e2).path;
			};
			return Reflect.compare(n1, n2);
		});
		for (x in l)
			switch (x) {
				case TPackage(_, _, l):
					sort(l);
				case TClassdecl(c):
					sortFields(c.fields);
					sortFields(c.statics);
				case TEnumdecl(_):
				case TAbstractdecl(_):
				case TTypedecl(_):
			}
	}

	function sortFields(a:Array<ClassField>) {
		ArraySort.sort(a, function(f1:ClassField, f2:ClassField) {
			var v1 = TypeApi.isVar(f1.type);
			var v2 = TypeApi.isVar(f2.type);
			if (v1 && !v2)
				return -1;
			if (v2 && !v1)
				return 1;
			if (f1.name == "new")
				return -1;
			if (f2.name == "new")
				return 1;
			return Reflect.compare(f1.name, f2.name);
		});
	}

	public function process(x:Xml, platform:String) {
		curplatform = platform;
		xroot(x);
	}

	// merge inline and not inline
	function mergeRights(f1:ClassField, f2:ClassField) {
		if (f1.get == RInline && f1.set == RNo && f2.get == RNormal && f2.set == RMethod) {
			f1.get = RNormal;
			f1.set = RMethod;
			return true;
		}
		return Type.enumEq(f1.get, f2.get) && Type.enumEq(f1.set, f2.set);
	}

	function mergeDoc(f1:ClassField, f2:ClassField) {
		if (f1.doc == null)
			f1.doc = f2.doc;
		else if (f2.doc == null)
			f2.doc = f1.doc;
		return true;
	}

	function mergeFields(f:ClassField, f2:ClassField) {
		return TypeApi.fieldEq(f, f2)
			|| (f.name == f2.name && (mergeRights(f, f2) || mergeRights(f2, f)) && mergeDoc(f, f2) && TypeApi.fieldEq(f, f2));
	}

	public dynamic function newField(c:Classdef, f:ClassField) {}

	function mergeClasses(c:Classdef, c2:Classdef) {
		// todo : compare supers & interfaces
		if (c.isInterface != c2.isInterface)
			return false;
		if (curplatform != null)
			c.platforms.push(curplatform);
		if (c.isExtern != c2.isExtern)
			c.isExtern = false;

		for (f2 in c2.fields) {
			var found = null;
			for (f in c.fields)
				if (mergeFields(f, f2)) {
					found = f;
					break;
				}
			if (found == null) {
				newField(c, f2);
				c.fields.push(f2);
			} else if (curplatform != null)
				found.platforms.push(curplatform);
		}
		for (f2 in c2.statics) {
			var found = null;
			for (f in c.statics)
				if (mergeFields(f, f2)) {
					found = f;
					break;
				}
			if (found == null) {
				newField(c, f2);
				c.statics.push(f2);
			} else if (curplatform != null)
				found.platforms.push(curplatform);
		}
		return true;
	}

	function mergeEnums(e:Enumdef, e2:Enumdef) {
		if (e.isExtern != e2.isExtern)
			return false;
		if (curplatform != null)
			e.platforms.push(curplatform);
		for (c2 in e2.constructors) {
			var found = null;
			for (c in e.constructors)
				if (TypeApi.constructorEq(c, c2)) {
					found = c;
					break;
				}
			if (found == null)
				e.constructors.push(c2);
			else if (curplatform != null)
				found.platforms.push(curplatform);
		}
		return true;
	}

	function mergeTypedefs(t:Typedef, t2:Typedef) {
		if (curplatform == null)
			return false;
		t.platforms.push(curplatform);
		t.types.set(curplatform, t2.type);
		return true;
	}

	function mergeAbstracts(a:Abstractdef, a2:Abstractdef) {
		if (curplatform == null)
			return false;
		if (a.to.length != a2.to.length || a.from.length != a2.from.length)
			return false;
		for (i in 0...a.to.length)
			if (!TypeApi.typeEq(a.to[i].t, a2.to[i].t))
				return false;
		for (i in 0...a.from.length)
			if (!TypeApi.typeEq(a.from[i].t, a2.from[i].t))
				return false;
		if (a2.impl != null)
			mergeClasses(a.impl, a2.impl);
		a.platforms.push(curplatform);
		return true;
	}

	function merge(t:TypeTree) {
		var inf = TypeApi.typeInfos(t);
		var pack = splitString(inf.path, ".");
		var cur = root;
		var curpack = new Array();
		pack.pop();
		for (p in pack) {
			var found = false;
			for (pk in cur)
				switch (pk) {
					case TPackage(pname, _, subs):
						if (pname == p) {
							found = true;
							cur = subs;
							break;
						}
					default:
				}
			curpack.push(p);
			if (!found) {
				var pk = new Array();
				cur.push(TPackage(p, joinStringArray(curpack, "."), pk));
				cur = pk;
			}
		}
		for (ct in cur) {
			if (ct.match(TPackage(_)))
				continue;
			var tinf = TypeApi.typeInfos(ct);

			// compare params ?
			if (tinf.path == inf.path) {
				var sameType = true;
				if ((tinf.doc == null) != (inf.doc == null)) {
					if (inf.doc == null)
						inf.doc = tinf.doc;
					else
						tinf.doc = inf.doc;
				}
				if (tinf.path == "haxe._Int64.NativeInt64")
					continue;
				if (tinf.module == inf.module && tinf.doc == inf.doc && tinf.isPrivate == inf.isPrivate)
					switch (ct) {
						case TClassdecl(c):
							switch (t) {
								case TClassdecl(c2):
									if (mergeClasses(c, c2)) return;
								default:
									sameType = false;
							}
						case TEnumdecl(e):
							switch (t) {
								case TEnumdecl(e2):
									if (mergeEnums(e, e2)) return;
								default:
									sameType = false;
							}
						case TTypedecl(td):
							switch (t) {
								case TTypedecl(td2):
									if (mergeTypedefs(td, td2)) return;
								default:
							}
						case TAbstractdecl(a):
							switch (t) {
								case TAbstractdecl(a2):
									if (mergeAbstracts(a, a2)) return;
								default:
									sameType = false;
							}
						case TPackage(_, _, _):
							sameType = false;
					}
				// we already have a mapping, but which is incompatible
				var msg = if (tinf.module != inf.module) "module " + inf.module + " should be " + tinf.module; else if (tinf.doc != inf.doc)
					"documentation is different"; else if (tinf.isPrivate != inf.isPrivate) "private flag is different"; else if (!sameType)
					"type kind is different"; else "could not merge definition";
				throw "Incompatibilities between "
					+ tinf.path
					+ " in "
					+ joinStringArray(tinf.platforms, ",")
					+ " and "
					+ curplatform
					+ " ("
					+ msg
					+ ")";
			}
		}
		cur.push(t);
	}

	function mkPath(p:String):Path {
		return p;
	}

	function mkTypeParams(p:String):TypeParams {
		var pl = splitString(p, ":");
		if (pl[0] == "")
			return new Array();
		return pl;
	}

	function mkRights(r:String):Rights {
		if (r == "null")
			return RNo;
		if (r == "method")
			return RMethod;
		if (r == "dynamic")
			return RDynamic;
		if (r == "inline")
			return RInline;
		return RCall(r);
	}

	function xroot(x:Xml) {
		for (c in x.elements())
			merge(processElement(c));
	}

	public function processElement(x:Xml) {
		var nodeName = if (x.nodeType == Xml.Document) "Document" else x.nodeName;
		if (nodeName == "class")
			return TClassdecl(xclass(x));
		if (nodeName == "enum")
			return TEnumdecl(xenum(x));
		if (nodeName == "typedef")
			return TTypedecl(xtypedef(x));
		if (nodeName == "abstract")
			return TAbstractdecl(xabstract(x));
		throw "Invalid " + nodeName;
	}

	function xmeta(x:Xml):MetaData {
		var ml = [];
		for (m in x.elementsNamed("m")) {
			var pl = [];
			for (p in m.elementsNamed("e"))
				pl.push(innerHTML(p));
			ml.push({name: requireAttr(m, "n"), params: pl});
		}
		return ml;
	}

	function xoverloads(x:Xml):Array<ClassField> {
		var l = new Array();
		for (m in x.elements()) {
			l.push(xclassfield(m));
		}
		return l;
	}

	function xpath(x:Xml):PathParams {
		var path = mkPath(requireAttr(x, "path"));
		var params = new Array();
		for (c in x.elements())
			params.push(xtype(c));
		return {
			path: path,
			params: params,
		};
	}

	function xclass(x:Xml):Classdef {
		var csuper = null;
		var doc = null;
		var tdynamic = null;
		var interfaces = new Array();
		var fields = new Array();
		var statics = new Array();
		var meta = [];
		var isInterface = x.exists("interface");
		for (c in x.elements()) {
			var nodeName = elementName(c);
			if (nodeName == "haxe_doc") {
				doc = innerData(c);
			} else if (nodeName == "extends") {
				if (isInterface) {
					interfaces.push(xpath(c));
				} else {
					csuper = xpath(c);
				}
			} else if (nodeName == "implements") {
				interfaces.push(xpath(c));
			} else if (nodeName == "haxe_dynamic") {
				tdynamic = xtype(requireFirstElement(c));
			} else if (nodeName == "meta") {
				meta = xmeta(c);
			} else if (c.exists("static")) {
				statics.push(xclassfield(c));
			} else {
				fields.push(xclassfield(c));
			}
		}
		return {
			file: x.get("file"),
			path: mkPath(requireAttr(x, "path")),
			module: if (x.exists("module")) mkPath(requireAttr(x, "module")) else null,
			doc: doc,
			isPrivate: x.exists("private"),
			isExtern: x.exists("extern"),
			isFinal: x.exists("final"),
			isInterface: isInterface,
			params: mkTypeParams(requireAttr(x, "params")),
			superClass: csuper,
			interfaces: interfaces,
			fields: fields,
			statics: statics,
			tdynamic: tdynamic,
			platforms: defplat(),
			meta: meta,
		};
	}

	function xclassfield(x:Xml, ?defPublic = false):ClassField {
		var e = x.elements();
		var t = xtype(e.next());
		var doc = null;
		var meta = [];
		var overloads = null;
		var line:Dynamic = null;
		if (x.exists("line")) {
			line = parseIntString(requireAttr(x, "line"));
		}
		for (c in e) {
			var nodeName = elementName(c);
			if (nodeName == "haxe_doc") {
				doc = innerData(c);
			} else if (nodeName == "meta") {
				meta = xmeta(c);
			} else if (nodeName == "overloads") {
				overloads = xoverloads(c);
			} else {
				throw "Invalid " + nodeName;
			}
		}
		return {
			name: elementName(x),
			type: t,
			isPublic: x.exists("public") || defPublic,
			isFinal: x.exists("final"),
			isOverride: x.exists("override"),
			line: line,
			doc: doc,
			get: if (x.exists("get")) mkRights(requireAttr(x, "get")) else RNormal,
			set: if (x.exists("set")) mkRights(requireAttr(x, "set")) else RNormal,
			params: if (x.exists("params")) mkTypeParams(requireAttr(x, "params")) else [],
			platforms: defplat(),
			meta: meta,
			overloads: overloads,
			expr: if (x.exists("expr")) requireAttr(x, "expr") else null
		};
	}

	function xenum(x:Xml):Enumdef {
		var cl = new Array();
		var doc = null;
		var meta = [];
		for (c in x.elements())
			if (elementName(c) == "haxe_doc")
				doc = innerData(c);
			else if (elementName(c) == "meta")
				meta = xmeta(c);
			else
				cl.push(xenumfield(c));
		return {
			file: x.get("file"),
			path: mkPath(requireAttr(x, "path")),
			module: if (x.exists("module")) mkPath(requireAttr(x, "module")) else null,
			doc: doc,
			isPrivate: x.exists("private"),
			isExtern: x.exists("extern"),
			params: mkTypeParams(requireAttr(x, "params")),
			constructors: cl,
			platforms: defplat(),
			meta: meta,
		};
	}

	function xenumfield(x:Xml):EnumField {
		var args = null;
		var docElements = x.elementsNamed("haxe_doc");
		var xdoc = if (docElements.hasNext()) docElements.next() else null;
		var meta = if (hasNamedElement(x, "meta")) xmeta(requireNamedElement(x, "meta")) else [];
		if (x.exists("a")) {
			var names = splitString(requireAttr(x, "a"), ":");
			var elts = x.elements();
			args = new Array();
			for (c in names) {
				var opt = false;
				if (c.charAt(0) == "?") {
					opt = true;
					c = c.substr(1);
				}
				args.push({
					name: c,
					opt: opt,
					t: xtype(elts.next()),
				});
			}
		}
		return {
			name: elementName(x),
			args: args,
			doc: if (xdoc == null) null else innerData(xdoc),
			meta: meta,
			platforms: defplat(),
		};
	}

	function xabstract(x:Xml):Abstractdef {
		var doc = null, impl = null, athis = null;
		var meta = [], to = [], from = [];
		for (c in x.elements()) {
			var nodeName = elementName(c);
			if (nodeName == "haxe_doc") {
				doc = innerData(c);
			} else if (nodeName == "meta") {
				meta = xmeta(c);
			} else if (nodeName == "to") {
				for (t in c.elements())
					to.push({t: xtype(requireFirstElement(t)), field: t.get("field")});
			} else if (nodeName == "from") {
				for (t in c.elements())
					from.push({t: xtype(requireFirstElement(t)), field: t.get("field")});
			} else if (nodeName == "impl") {
				impl = xclass(requireNamedElement(c, "class"));
			} else if (nodeName == "this") {
				athis = xtype(requireFirstElement(c));
			} else {
				throw "Invalid " + nodeName;
			}
		}
		return {
			file: x.get("file"),
			path: mkPath(requireAttr(x, "path")),
			module: if (x.exists("module")) mkPath(requireAttr(x, "module")) else null,
			doc: doc,
			isPrivate: x.exists("private"),
			params: mkTypeParams(requireAttr(x, "params")),
			platforms: defplat(),
			meta: meta,
			athis: athis,
			to: to,
			from: from,
			impl: impl
		};
	}

	function xtypedef(x:Xml):Typedef {
		var doc = null;
		var t = null;
		var meta = [];
		for (c in x.elements())
			if (elementName(c) == "haxe_doc")
				doc = innerData(c);
			else if (elementName(c) == "meta")
				meta = xmeta(c);
			else
				t = xtype(c);
		var types = new haxe.ds.StringMap();
		if (curplatform != null)
			types.set(curplatform, t);
		return {
			file: x.get("file"),
			path: mkPath(requireAttr(x, "path")),
			module: if (x.exists("module")) mkPath(requireAttr(x, "module")) else null,
			doc: doc,
			isPrivate: x.exists("private"),
			params: mkTypeParams(requireAttr(x, "params")),
			type: t,
			types: types,
			platforms: defplat(),
			meta: meta,
		};
	}

	function xtype(x:Xml):CType {
		var nodeName = elementName(x);
		if (nodeName == "unknown")
			return CUnknown;
		if (nodeName == "e")
			return CEnum(mkPath(requireAttr(x, "path")), xtypeparams(x));
		if (nodeName == "c")
			return CClass(mkPath(requireAttr(x, "path")), xtypeparams(x));
		if (nodeName == "t")
			return CTypedef(mkPath(requireAttr(x, "path")), xtypeparams(x));
		if (nodeName == "x")
			return CAbstract(mkPath(requireAttr(x, "path")), xtypeparams(x));
		if (nodeName == "f") {
			var args = new Array();
			var aname = splitString(requireAttr(x, "a"), ":");
			var argIndex = 0;
			var evalues = x.exists("v") ? splitString(requireAttr(x, "v"), ":") : null;
			var valueIndex = 0;
			for (e in x.elements()) {
				var opt = false;
				var a = argIndex < aname.length ? aname[argIndex] : null;
				argIndex++;
				if (a == null)
					a = "";
				if (a.charAt(0) == "?") {
					opt = true;
					a = a.substr(1);
				}
				var v = evalues == null || valueIndex >= evalues.length ? null : evalues[valueIndex++];
				args.push({
					name: a,
					opt: opt,
					t: xtype(e),
					value: v == "" ? null : v
				});
			}
			var ret = args[args.length - 1];
			var callArgs = new Array();
			for (i in 0...(args.length - 1)) {
				callArgs.push(args[i]);
			}
			return CFunction(callArgs, ret.t);
		}
		if (nodeName == "a") {
			var fields = new Array();
			for (f in x.elements()) {
				var f = xclassfield(f, true);
				f.platforms = new Array(); // platforms selection are on the type itself, not on fields
				fields.push(f);
			}
			return CAnonymous(fields);
		}
		if (nodeName == "d") {
			var t = null;
			var tx = x.firstElement();
			if (tx != null)
				t = xtype(tx);
			return CDynamic(t);
		}
		throw "Invalid " + nodeName;
	}

	function xtypeparams(x:Xml):Array<CType> {
		var p = new Array();
		for (c in x.elements())
			p.push(xtype(c));
		return p;
	}

	function defplat() {
		var l = new Array();
		if (curplatform != null)
			l.push(curplatform);
		return l;
	}

	function joinStringArray(values:Array<String>, separator:String):String {
		var buf = new StringBuf();
		for (i in 0...values.length) {
			if (i > 0) {
				buf.add(separator);
			}
			buf.add(values[i]);
		}
		return buf.toString();
	}

	function splitString(value:String, separator:String):Array<String> {
		if (separator == "") {
			return [value];
		}
		var parts = new Array<String>();
		var start = 0;
		while (true) {
			var index = findSeparator(value, separator, start);
			if (index == -1) {
				parts.push(value.substr(start));
				break;
			}
			parts.push(value.substr(start, index - start));
			start = index + separator.length;
		}
		return parts;
	}

	function findSeparator(value:String, separator:String, start:Int):Int {
		var limit = value.length - separator.length;
		var index = start;
		while (index <= limit) {
			if (value.substr(index, separator.length) == separator) {
				return index;
			}
			index++;
		}
		return -1;
	}

	function requireAttr(x:Xml, name:String):String {
		var value = x.get(name);
		return value == null ? "" : value;
	}

	function hasNamedElement(x:Xml, name:String):Bool {
		return x.elementsNamed(name).hasNext();
	}

	function requireNamedElement(x:Xml, name:String):Xml {
		var elements = x.elementsNamed(name);
		if (!elements.hasNext()) {
			throw nodeDisplayName(x) + " is missing element " + name;
		}
		return elements.next();
	}

	function requireFirstElement(x:Xml):Xml {
		var first = x.firstElement();
		if (first == null) {
			throw nodeDisplayName(x) + " is missing first element";
		}
		return first;
	}

	function nodeDisplayName(x:Xml):String {
		return if (x.nodeType == Xml.Document) "Document" else elementName(x);
	}

	function elementName(x:Xml):String {
		var name = x.nodeName;
		return name == null ? "" : name;
	}

	function innerData(x:Xml):String {
		var it = x.iterator();
		if (!it.hasNext()) {
			throw nodeDisplayName(x) + " does not have data";
		}
		var value = it.next();
		if (it.hasNext()) {
			throw nodeDisplayName(x) + " does not only have data";
		}
		if (value.nodeType != Xml.PCData && value.nodeType != Xml.CData) {
			throw nodeDisplayName(x) + " does not have data";
		}
		return value.nodeValue;
	}

	function innerHTML(x:Xml):String {
		var buf = new StringBuf();
		for (child in x) {
			buf.add(child.toString());
		}
		return buf.toString();
	}

	function parseIntString(value:String):Int {
		if (value == null || value == "") {
			return 0;
		}
		var zeroCode = 48;
		var nineCode = 57;
		var negative = false;
		var index = 0;
		if (value.charAt(0) == "-") {
			negative = true;
			index = 1;
		}
		var result = 0;
		while (index < value.length) {
			var code:Int = value.charCodeAt(index);
			if (code < zeroCode || code > nineCode) {
				return 0;
			}
			result = result * 10 + (code - zeroCode);
			index++;
		}
		return negative ? -result : result;
	}
}
