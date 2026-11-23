/**
	`Reflect` and `Dynamic` reads of Array and String members must behave as in
	`--interp` (haxe_go-rotv).

	`Reflect.field` returns the `length` of an Array or a String, and each of
	their methods as a function bound to the receiver, which `Reflect.callMethod`
	and `Dynamic` calls invoke. A field of a `Dynamic` value reads through
	`Reflect.field`, so `d.length`, `d.charAt(0)` and `d.indexOf(x)` work on an
	Array or a String held in `Dynamic`, also inside an anonymous object or a
	parsed JSON value. A missing optional argument, or a null one, reads as
	omitted. `Reflect.hasField` reports none of these members, `Reflect.fields`
	lists only `length`, `Reflect.getProperty` reads `length`, and writing
	`length` (`Reflect.setField`, `Reflect.setProperty`, `d.length = n`)
	truncates the Array or extends it with nulls.
	Left out: `iterator` and `keyValueIterator` (not read on haxe.go yet), and
	writes to other Array or String fields (`--interp` throws, haxe.go ignores
	them).
**/
class Main {
	static final arrayMembers = [
		"concat",
		"contains",
		"copy",
		"filter",
		"indexOf",
		"insert",
		"join",
		"lastIndexOf",
		"map",
		"pop",
		"push",
		"remove",
		"resize",
		"reverse",
		"shift",
		"slice",
		"sort",
		"splice",
		"toString",
		"unshift"
	];
	static final stringMembers = [
		"charAt",
		"charCodeAt",
		"indexOf",
		"lastIndexOf",
		"split",
		"substr",
		"substring",
		"toLowerCase",
		"toUpperCase",
		"toString"
	];

	static function row(label:String, value:Dynamic):Void {
		Sys.println(label + ": " + Std.string(value));
	}

	static function call(receiver:Dynamic, name:String, arguments:Array<Dynamic>):Dynamic {
		return Reflect.callMethod(receiver, Reflect.field(receiver, name), arguments);
	}

	static function main() {
		reflectReads();
		boundArrayMethods();
		boundStringMethods();
		boundIdentity();
		dynamicReads();
		lengthWrites();
	}

	static function reflectReads() {
		var numbers = [1, 2, 3];
		var text = "abcd";
		row("field Array length", Reflect.field(numbers, "length"));
		row("field String length", Reflect.field(text, "length"));
		row("field empty length", Reflect.field([], "length") + Reflect.field("", "length"));
		for (name in arrayMembers)
			row("Array " + name + " is a function", Reflect.isFunction(Reflect.field(numbers, name)));
		for (name in stringMembers)
			row("String " + name + " is a function", Reflect.isFunction(Reflect.field(text, name)));
		for (name in ["missing", "Len", "Push", "Values"])
			row("Array " + name, Reflect.field(numbers, name));
		row("String missing", Reflect.field(text, "missing"));
		row("hasField Array", [
			Reflect.hasField(numbers, "length"),
			Reflect.hasField(numbers, "push"),
			Reflect.hasField(numbers, "Len")
		]);
		row("hasField String", [Reflect.hasField(text, "length"), Reflect.hasField(text, "charAt")]);
		row("fields Array", Reflect.fields(numbers));
		row("fields String", Reflect.fields(text));
		row("getProperty", [Reflect.getProperty(numbers, "length"), Reflect.getProperty(text, "length")]);
		var length:Int = Reflect.field(text, "length");
		row("typed length", length * 2);
	}

	static function boundArrayMethods() {
		var a = [3, 1, 2, 1];
		row("indexOf", call(a, "indexOf", [1]));
		row("indexOf from", call(a, "indexOf", [1, 2]));
		row("indexOf negative from", call(a, "indexOf", [1, -1]));
		row("indexOf null from", call(a, "indexOf", [1, null]));
		row("lastIndexOf", call(a, "lastIndexOf", [1]));
		row("lastIndexOf from", call(a, "lastIndexOf", [1, 2]));
		row("contains", [call(a, "contains", [2]), call(a, "contains", [9])]);
		row("join", call(a, "join", ["|"]));
		row("slice", [call(a, "slice", [1]), call(a, "slice", [1, 2]), call(a, "slice", [1, null])]);
		row("copy", call(a, "copy", []));
		row("concat", call(a, "concat", [[9]]));
		row("toString", call(a, "toString", []));
		row("map", call(a, "map", [(x:Int) -> x * 10]));
		row("filter", call(a, "filter", [(x:Int) -> x > 1]));
		row("remove", call(a, "remove", [1]));
		row("after remove", a);
		row("shift", call(a, "shift", []));
		row("unshift", call(a, "unshift", [0]));
		row("push", call(a, "push", [7]));
		row("after push", a);
		row("sort", call(a, "sort", [(x:Int, y:Int) -> x - y]));
		row("after sort", a);
		row("insert", call(a, "insert", [1, 5]));
		row("reverse", call(a, "reverse", []));
		row("after reverse", a);
		row("splice", call(a, "splice", [1, 2]));
		row("pop", call(a, "pop", []));
		row("resize", call(a, "resize", [4]));
		row("after resize", a);
		var empty:Array<Int> = [];
		row("empty pop and shift", [call(empty, "pop", []), call(empty, "shift", [])]);
		var words = ["b", "a", "c"];
		row("String elements", [call(words, "indexOf", ["c"]), call(words, "join", ["+"])]);
	}

	static function boundStringMethods() {
		var s = "x,y,z";
		row("charAt", [call(s, "charAt", [2]), call(s, "charAt", [10])]);
		row("charCodeAt", [call(s, "charCodeAt", [0]), call(s, "charCodeAt", [10])]);
		row("indexOf", [
			call(s, "indexOf", [","]),
			call(s, "indexOf", [",", 2]),
			call(s, "indexOf", [",", null])
		]);
		row("lastIndexOf", [call(s, "lastIndexOf", [","]), call(s, "lastIndexOf", [",", 2])]);
		row("split", call(s, "split", [","]));
		row("substr", [call(s, "substr", [2]), call(s, "substr", [2, 1]), call(s, "substr", [2, null])]);
		row("substring", [
			call(s, "substring", [2]),
			call(s, "substring", [2, 3]),
			call(s, "substring", [2, null])
		]);
		row("case", [call(s, "toUpperCase", []), call("ABC", "toLowerCase", [])]);
		row("toString", call(s, "toString", []));
	}

	static function boundIdentity() {
		var a = [1];
		var alias = a;
		var push:Dynamic = Reflect.field(a, "push");
		a.push(2);
		push(3);
		Reflect.callMethod(null, push, [4]);
		row("bound push", alias);
		var length:Dynamic = Reflect.field(a, "length");
		a.push(5);
		row("length read is a value", length);
		var again = Reflect.field(a, "push");
		row("compareMethods same", Reflect.compareMethods(push, again));
		row("compareMethods other Array", Reflect.compareMethods(push, Reflect.field([1], "push")));
		row("compareMethods other method", Reflect.compareMethods(push, Reflect.field(a, "pop")));
		row("fresh value", push == again);
		var s = "ab";
		row("String compareMethods", Reflect.compareMethods(Reflect.field(s, "charAt"), Reflect.field(s, "charAt")));
	}

	static function dynamicReads() {
		var d:Dynamic = [1, 2, 3];
		row("d.length", d.length);
		var pushed:Int = d.push(4);
		row("d.push", pushed);
		row("d.indexOf", d.indexOf(2));
		row("d.join", d.join(","));
		row("d.pop", d.pop());
		row("d.length + 1", d.length + 1);
		var n:Int = d.length;
		row("typed Int", n * 2);
		row("interpolation", 'size ${d.length}');
		row("condition", d.length > 2 ? "big" : "small");
		row("comparison", d.length == 3);
		var s:Dynamic = "x,y,z";
		row("s.length", s.length);
		row("s.charAt", s.charAt(0));
		row("s.split", s.split(","));
		row("s.split.length", s.split(",").length);
		row("s.indexOf", s.indexOf("y"));
		row("s.toUpperCase", s.toUpperCase());
		row("s.substr", s.substr(2));
		row("s.charCodeAt", s.charCodeAt(0));
		var c:String = s.charAt(4);
		row("typed String", c + "!");
		var o:Dynamic = {items: [1, 2], name: "ab"};
		row("nested", o.items.length + o.name.length);
		row("nested call", o.name.charAt(1) + o.items.indexOf(2));
		var parsed:Dynamic = haxe.Json.parse('{"tags":["a","b","c"],"title":"Hello"}');
		row("JSON", parsed.tags.length + " " + parsed.tags.join("/") + " " + parsed.title.toLowerCase());
	}

	static function lengthWrites() {
		var w = [5, 6, 7];
		Reflect.setField(w, "length", 1);
		row("setField length 1", w);
		// A fresh Array: `--interp` shows the old elements again when a
		// truncated Array grows, and the Haxe API leaves them unspecified.
		var grown = [8];
		Reflect.setField(grown, "length", 3);
		row("setField length 3", grown);
		Reflect.setProperty(grown, "length", 2);
		row("setProperty length 2", grown);
		var d:Dynamic = [5, 6, 7];
		var alias:Array<Int> = d;
		row("d.length = 0", d.length = 0);
		row("after d.length = 0", [d, alias]);
	}
}
