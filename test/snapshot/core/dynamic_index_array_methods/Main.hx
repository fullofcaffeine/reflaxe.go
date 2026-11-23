typedef Stack = {
	var push:Int->Int;
	var pop:Void->Null<Int>;
}

class Tally {
	public var total = 0;

	public function new() {}

	public function push(value:Int):Int {
		total += value;
		return total;
	}

	public function pop():Int {
		total--;
		return total;
	}
}

/**
	Index access and Array methods on `Dynamic` values (haxe_go-lgn0).

	haxe.go applied Go indexing to the Go `any` that holds a `Dynamic` value,
	so `d[1]` and `d[0] = x` did not build ("cannot index"). A statement call
	`d.push(x);` or `d.pop();` became a Go `append` or slice on `any`
	("invalid append"), and also on a structure with `push` and `pop` function
	fields, because the Array method lowering matched any receiver.

	Now an index of a `Dynamic` value calls `hxrt.DynamicIndexGet` and
	`hxrt.DynamicIndexSet`, which index an Array with the rules of typed
	Arrays and raise the `--interp` errors for any other value. Operators,
	conditions, `++`, `--` and compound assignment take the element as a
	`Dynamic` operand, whatever type its use gave it. The Array method
	lowering takes only Array receivers, so a `Dynamic` receiver calls the
	method through `Reflect` in statements too. Left out: a negative index
	write, `iterator` and `keyValueIterator`, a String element in an
	expression that Haxe types as Float, and fields of a String or an Array
	element.
**/
class Main {
	static var order = "";

	static function dyn(value:Dynamic):Dynamic {
		return value;
	}

	static function step(name:String, value:Dynamic):Dynamic {
		order += name;
		return value;
	}

	static function attempt(name:String, run:() -> Dynamic) {
		var text = try Std.string(run()) catch (e:String) "error " + e;
		Sys.println(name + " = " + text);
	}

	static function byValue(a:Dynamic, b:Dynamic):Int {
		return a - b > 0 ? 1 : (a == b ? 0 : -1);
	}

	static function firstInt(values:Dynamic):Int {
		return values[0];
	}

	static function firstText(values:Dynamic):String {
		return values[0];
	}

	static function main() {
		reads();
		boundTypes();
		writes();
		compound();
		nonArrays();
		evaluationOrder();
		statementMethods();
		expressionMethods();
		otherReceivers();
	}

	static function reads() {
		var a:Dynamic = dyn([10, 20, 30]);
		var i:Dynamic = dyn(1);
		var f:Dynamic = dyn(1.7);
		var n:Null<Int> = 2;
		Sys.println("read " + a[0] + " " + a[2] + " " + a[i] + " " + a[f] + " " + a[n]);
		Sys.println("outside " + a[3] + " " + a[100] + " " + a[-1]);
		var nested:Dynamic = dyn([[1, 2], [3, 4]]);
		Sys.println("nested " + nested[1][0] + " " + nested[0][5]);
		var rows:Array<Dynamic> = [dyn([5, 6]), "x"];
		Sys.println("element " + rows[0][1]);
		var json:Dynamic = haxe.Json.parse('{"items":[1,2,3],"grid":[[7,8]],"people":[{"name":"ann"}]}');
		Sys.println("json " + json.items[2] + " " + json.grid[0][1] + " " + json.people[0].name);
		var keys:Dynamic = dyn([0, 2]);
		var plain = [100, 200, 300];
		Sys.println("as index " + plain[keys[1]] + " " + a[keys[0]]);
		var maybe:Null<Dynamic> = dyn([4, 5]);
		Sys.println("nullable " + maybe[1]);
		var later = () -> a[2];
		Sys.println("closure " + later() + ' ${a[1]}');
	}

	static function boundTypes() {
		var a:Dynamic = dyn([10, 20, 30]);
		var words:Dynamic = dyn(["ab", "cd"]);
		var typedInt:Int = a[0];
		var typedFloat:Float = a[1];
		var typedString:String = words[1];
		var missing:Null<Int> = a[5];
		Sys.println("typed " + (typedInt + 1) + " " + (typedFloat / 8) + " " + typedString.toUpperCase() + " " + (missing == null));
		Sys.println("returns " + (firstInt(a) + 1) + " " + firstText(words));
		Sys.println("operators "
			+ (a[0] + a[1])
			+ " "
			+ (a[2] - 5)
			+ " "
			+ (a[0] * 2)
			+ " "
			+ (a[1] > a[0])
			+ " "
			+ (a[0] == 10)
			+ " "
			+ (-a[0]));
		var flags:Dynamic = dyn([true, false]);
		Sys.println("conditions "
			+ (flags[0] ? "yes" : "no")
			+ " "
			+ (flags[1] ? "yes" : "no")
			+ " "
			+ (flags[0] && !flags[1])
			+ " "
			+ (a[5] == null));
		var big:Dynamic = dyn([2147483647]);
		Sys.println("int " + (big[0] + 1) + " " + Std.int(a[0] / 3) + " " + Math.abs(a[0]));
		switch (a[1]) {
			case 20:
				Sys.println("switch twenty");
			case _:
				Sys.println("switch other");
		}
		var mixed:Dynamic = dyn([1, "x", 2.5, null]);
		Sys.println("mixed " + (mixed[0] + mixed[2]) + " " + (mixed[2] * 2) + " " + (mixed[0] < mixed[2]));
		attempt("null operand", () -> mixed[3] * 2);
		var total:Dynamic = 0;
		for (k in 0...a.length)
			total += a[k];
		Sys.println("loop " + total);
	}

	static function writes() {
		var a:Dynamic = dyn([10, 20, 30]);
		a[1] = 5;
		Sys.println("write " + a + " " + a.length);
		a[5] = 7;
		Sys.println("grow " + a + " " + a.length);
		Sys.println("value " + (a[0] = 1) + " " + a);
		var i:Dynamic = dyn(2);
		a[i] = "s";
		var f:Dynamic = dyn(3.9);
		a[f] = 3;
		a[4] = 2.0;
		Sys.println("dynamic index " + a + " " + Type.typeof(a[4]));
		a[0] = a[3];
		Sys.println("copy " + a);
		var typed = [1, 2];
		var alias:Dynamic = typed;
		alias[0] = 100;
		alias[3] = 300;
		Sys.println("alias " + typed + " " + typed.length);
		var nested:Dynamic = dyn([[1, 2], [3, 4]]);
		nested[1][0] = 9;
		Sys.println("nested " + nested);
		var json:Dynamic = haxe.Json.parse('{"items":[1,2,3]}');
		json.items[0] = 7;
		Sys.println("json " + haxe.Json.stringify(json));
		var rows:Array<Dynamic> = [dyn([5, 6])];
		rows[0][1] = 60;
		Sys.println("element " + rows);
		var target = [0, 0];
		target[1] = a[3];
		Sys.println("typed target " + target);
	}

	static function compound() {
		var a:Dynamic = dyn([10, 2.5, "s", null]);
		a[0] += 1;
		a[1] *= 2;
		a[2] += "t";
		Sys.println("compound " + a);
		Sys.println("value " + (a[0] -= 4) + " " + a[0]);
		Sys.println("postfix " + (a[0]++) + " " + a[0]);
		Sys.println("prefix " + (++a[0]) + " " + a[0]);
		a[0]--;
		--a[0];
		a[0]++;
		Sys.println("steps " + a[0]);
		a[0] += a[0];
		Sys.println("self " + a[0]);
		attempt("null element", () -> a[3] += 1);
		Sys.println("after " + a);
		var holder:Dynamic = dyn({list: [1, 2]});
		holder.list[0] += 5;
		holder.list[1]++;
		Sys.println("field element " + holder.list);
	}

	static function nonArrays() {
		var text:Dynamic = dyn("hello");
		attempt("string read", () -> text[0]);
		attempt("string write", () -> text[0] = "j");
		var object:Dynamic = dyn({x: 1});
		attempt("object read", () -> object[0]);
		attempt("object write", () -> object[0] = 3);
		attempt("object compound", () -> object[0] += 1);
		attempt("object step", () -> object[0]++);
		var nothing:Dynamic = dyn(null);
		attempt("null read", () -> nothing[0]);
		attempt("null write", () -> nothing[0] = 1);
		var number:Dynamic = dyn(5);
		attempt("int read", () -> number[0]);
		var flag:Dynamic = dyn(true);
		attempt("bool read", () -> flag[0]);
		var map:Dynamic = dyn([1 => "one"]);
		var mapText = try Std.string(map[1]) catch (e:String) "error";
		Sys.println("map read = " + mapText);
		Sys.println("unchanged " + text + " " + object.x);
	}

	static function evaluationOrder() {
		var a:Dynamic = dyn([1, 2]);
		step("t", a)[step("i", 0)] = step("v", 10);
		step("t", a)[step("i", 1)] += step("v", 5);
		var old = step("t", a)[step("i", 0)]++;
		Sys.println("order " + order + " " + a + " " + old);
		order = "";
		var read = step("t", a)[step("i", 1)];
		Sys.println("read order " + order + " " + read);
	}

	static function statementMethods() {
		var a:Dynamic = dyn([3, 1, 2]);
		a.push(4);
		a.push(5);
		a.pop();
		a.unshift(0);
		a.shift();
		a.insert(1, 9);
		a.remove(9);
		a.reverse();
		Sys.println("statements " + a);
		a.sort(byValue);
		a.resize(5);
		a.splice(4, 1);
		a.concat([5]);
		a.copy();
		a.filter(x -> x != 1);
		a.map(x -> x * 2);
		a.indexOf(2);
		a.lastIndexOf(2);
		a.contains(2);
		a.join(",");
		a.slice(1);
		a.toString();
		Sys.println("statements " + a);
	}

	static function expressionMethods() {
		var a:Dynamic = dyn([3, 1, 2]);
		Sys.println("push " + a.push(4) + " " + a);
		Sys.println("pop " + a.pop() + " " + a);
		Sys.println("unshift " + a.unshift(0) + " " + a);
		Sys.println("shift " + a.shift() + " " + a);
		Sys.println("insert " + a.insert(1, 9) + " " + a);
		Sys.println("remove " + a.remove(9) + " " + a.remove(9) + " " + a);
		Sys.println("reverse " + a.reverse() + " " + a);
		Sys.println("sort " + a.sort(byValue) + " " + a);
		Sys.println("resize " + a.resize(4) + " " + a);
		Sys.println("splice " + a.splice(3, 1) + " " + a);
		Sys.println("concat " + a.concat([5]) + " copy " + a.copy());
		Sys.println("filter " + a.filter(x -> x != 1) + " map " + a.map(x -> x * 2));
		Sys.println("search " + a.indexOf(2) + " " + a.lastIndexOf(2) + " " + a.contains(2) + " " + a.contains(7));
		Sys.println("join " + a.join("-") + " slice " + a.slice(1) + " text " + a.toString());
		var length:Int = a.push(8);
		var popped:Int = a.pop();
		Sys.println("typed " + (length + popped));
	}

	static function otherReceivers() {
		var items = [];
		useStack({push: value -> items.push(value), pop: () -> items.pop()});
		Sys.println("structure " + items);
		var tally:Dynamic = dyn(new Tally());
		tally.push(5);
		tally.push(2);
		tally.pop();
		Sys.println("class " + tally.total);
	}

	static function useStack(stack:Stack) {
		stack.push(1);
		stack.push(2);
		stack.pop();
		Sys.println("structure push " + stack.push(3));
	}
}
