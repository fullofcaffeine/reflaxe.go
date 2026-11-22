/** The reported shape: an abstract over String with implicit casts both ways. */
abstract Name(String) from String to String {}

/** Only an implicit cast from String; reading the text needs a method. */
abstract Tag(String) from String {
	public inline function text():String {
		return this;
	}
}

/** Only an implicit cast to String; values come from the constructor. */
abstract Label(String) to String {
	public inline function new(raw:String) {
		this = raw;
	}
}

/** No implicit casts; values come from the constructor. */
abstract Code(String) {
	public inline function new(raw:String) {
		this = raw;
	}

	public inline function text():String {
		return this;
	}
}

/** Forwards selected String fields to the underlying value. */
@:forward(length, toUpperCase, charAt, indexOf)
abstract Word(String) from String to String {}

/** Forwards every String field to the underlying value. */
@:forward
abstract Text(String) from String to String {}

/** An abstract over another String-backed abstract. */
abstract Alias(Name) from Name to Name from String to String {}

/** Declared extern: Haxe generates no implementation class for it. */
extern abstract Plain(String) from String to String {
	public inline function shout():String {
		return this.toUpperCase();
	}
}

/** Applied type parameter: `Box<String>` stores a String. */
abstract Box<T>(T) from T to T {}

/** Abstract over Dynamic keeps dynamic storage. */
abstract Loose(Dynamic) from Dynamic to Dynamic {}

enum abstract Color(String) to String {
	var Red = "red";
	var Green = "green";
	var Blue = "blue";
}

enum abstract Mode(String) from String to String {
	var Fast = "fast";
	var Slow = "slow";
}

abstract Count(Int) from Int to Int {}
abstract Ratio(Float) from Float to Float {}
abstract Flag(Bool) from Bool to Bool {}

enum abstract Level(Int) to Int {
	var Low = 1;
	var High = 2;
}

enum Wrap {
	W(n:Name, c:Color);
}

typedef Entry = {
	name:Name,
	?color:Color
}

interface Namer {
	function rename(n:Name, ?suffix:Name):Name;
}

class Person implements Namer {
	public static var registry:Name = "reg";
	public static var lastColor:Null<Color> = null;

	public var name:Name;
	public var nick:Null<Name>;
	public var color:Color;

	public function new(name:Name, ?nick:Name, color:Color = Red) {
		this.name = name;
		this.nick = nick;
		this.color = color;
		lastColor = color;
	}

	public function rename(n:Name, ?suffix:Name):Name {
		if (suffix == null)
			return n;
		var head:String = n;
		var tail:String = suffix;
		return head + tail;
	}
}

/**
	Values of an abstract over String are stored as Haxe strings in every
	position (haxe_go-uvri): parameters (plain, optional, with defaults,
	`Null<T>`), returns, locals, instance and static fields, Array elements,
	Map values, enum arguments, structure fields, interface calls and function
	values. This holds with and without `from`/`to` casts, with `@:forward`,
	for enum abstracts, for abstracts over abstracts, for `extern abstract`, and
	for an applied type parameter. Abstracts over Int, Float, Bool, Dynamic and
	unapplied type parameters keep their behavior.
**/
class Main {
	static function nameToString(n:Name):String {
		return n;
	}

	static function idName(n:Name):Name {
		return n;
	}

	static function stringToName(s:String):Name {
		return s;
	}

	static function optName(?n:Name):String {
		return n == null ? "none" : n;
	}

	static function defName(n:Name = "dflt"):String {
		return n;
	}

	static function optDefName(?n:Name = "odflt"):Name {
		return n;
	}

	static function nullName(n:Null<Name>):String {
		return n == null ? "null" : n;
	}

	static function optColor(?c:Color):String {
		return c == null ? "none" : c;
	}

	static function defColor(c:Color = Green):String {
		return c;
	}

	static function defMode(m:Mode = Slow):Mode {
		return m;
	}

	static function modeOf(s:String):Mode {
		return s;
	}

	static function colorLetter(c:Color):String {
		return switch c {
			case Red: "R";
			case Green: "G";
			case Blue: "B";
		}
	}

	static function nameKind(n:Name):String {
		return switch n {
			case "alpha": "first";
			case "beta": "second";
			case _: "other";
		}
	}

	static function tagText(t:Tag):String {
		return t.text();
	}

	static function labelText(l:Label):String {
		return l;
	}

	static function codeText(c:Code):String {
		return c.text();
	}

	static function wordInfo(w:Word):String {
		return w.length + ":" + w.toUpperCase() + ":" + w.charAt(0) + ":" + w.indexOf("r");
	}

	static function textInfo(t:Text):String {
		return t.substr(1) + ":" + t.split("o").length;
	}

	static function aliasText(a:Alias):String {
		var n:Name = a;
		return n;
	}

	static function plainText(p:Plain):String {
		var raw:String = p;
		return p.shout() + "/" + raw;
	}

	static function boxText(b:Box<String>):String {
		return b;
	}

	static function looseText(l:Loose):String {
		return Std.string(l);
	}

	static function countOf(c:Count):Int {
		return c;
	}

	static function optCount(?c:Count):String {
		return c == null ? "none" : Std.string(c);
	}

	static function ratioOf(r:Ratio):Float {
		return r;
	}

	static function flagOf(f:Flag):Bool {
		return f;
	}

	static function levelOf(l:Level):Int {
		return l;
	}

	static function firstOf<T>(xs:Array<T>):T {
		return xs[0];
	}

	static function compareNames(x:Name, y:Name):Int {
		var a:String = x;
		var b:String = y;
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function main() {
		Sys.println("param.plain=" + nameToString("alpha") + "," + nameKind("alpha") + "," + nameKind("beta") + "," + nameKind("zeta"));
		Sys.println("param.id=" + idName("beta") + "," + stringToName("gamma"));
		Sys.println("param.optional=" + optName() + "," + optName(null) + "," + optName("set"));
		Sys.println("param.default=" + defName() + "," + defName(null) + "," + defName("given") + "," + defName(""));
		Sys.println("param.optionalDefault=" + optDefName() + "," + optDefName(null) + "," + optDefName("given"));
		var noName:Null<Name> = null;
		Sys.println("param.nullable=" + nullName(noName) + "," + nullName("here") + "," + (noName == null));

		var local:Name = "local";
		var copy:String = local;
		var back:Name = copy;
		var reassigned:Name = "first";
		reassigned = "second";
		Sys.println("local="
			+ local
			+ ","
			+ copy
			+ ","
			+ back
			+ ","
			+ reassigned
			+ ","
			+ '$local!'
			+ ","
			+ StringTools.startsWith(local, "lo"));

		var joined:Name = "ab" + Std.string(1);
		var literal:Name = "ab1";
		Sys.println("equality=" + (joined == literal) + "," + (joined != literal) + "," + (local == "local") + "," + ("local" == local) + "," +
			(local == copy));

		var p = new Person("ann");
		var q = new Person("bob", "bobby", Blue);
		p.name = "anna";
		Person.registry = "people";
		var fieldText:String = p.name;
		Sys.println("field=" + fieldText + "," + p.nick + "," + p.color + "," + q.name + "," + q.nick + "," + q.color + "," + Person.registry + ","
			+ Person.lastColor);

		var namer:Namer = q;
		Sys.println("interface=" + namer.rename("x") + "," + namer.rename("x", "-y") + "," + namer.rename("x", null));

		var names:Array<Name> = ["c", "a", "b"];
		names.push("d");
		names.sort(compareNames);
		var first:Name = names[0];
		var firstText:String = names[1];
		var generic:Name = firstOf(names);
		var lengths = names.map(n -> (n : String).length);
		var kept = names.filter(n -> n != "a");
		Sys.println("array=" + names.join("|") + "," + names.length + "," + first + "," + firstText + "," + generic + "," + names.indexOf("b") + ","
			+ names.contains("zz") + "," + lengths.join("|") + "," + kept.join("|"));
		names[2] = "bb";
		var loopOut = new Array<String>();
		for (n in names) {
			var s:String = n;
			loopOut.push(s.toUpperCase());
		}
		Sys.println("array.loop=" + loopOut.join("|"));

		var byKey = new Map<String, Name>();
		byKey.set("k1", "v1");
		byKey["k2"] = "v2";
		var got:Name = byKey.get("k1");
		var gotText:String = byKey["k2"];
		var missing:Null<Name> = byKey.get("none");
		var keys = [for (k in byKey.keys()) k];
		keys.sort(compareNames);
		var values = [for (v in byKey) v];
		values.sort(compareNames);
		Sys.println("map=" + got + "," + gotText + "," + missing + "," + byKey.exists("k1") + "," + keys.join("|") + "," + values.join("|"));
		var byName = new Map<Name, Int>();
		byName.set("n1", 1);
		var nameKey:Name = "n1";
		Sys.println("map.key=" + byName.get(nameKey) + "," + byName.exists("n1") + "," + byName.exists("n2"));

		var c:Color = Green;
		var colorText:String = c;
		var colors:Array<Color> = [Red, Blue];
		colors.push(Green);
		var colorMap = new Map<String, Color>();
		colorMap.set("sky", Blue);
		Sys.println("enumAbstract=" + colorLetter(c) + "," + colorText + "," + colors.join("|") + "," + (c == Green) + "," + (colors[1] == Blue) + ","
			+ optColor() + "," + optColor(Blue) + "," + defColor() + "," + defColor(null) + "," + defColor(Red) + "," + colorMap.get("sky"));
		var m:Mode = modeOf("fast");
		var modeLetter = switch m {
			case Fast: "F";
			case Slow: "S";
		}
		Sys.println("enumAbstract.from=" + m + "," + (m == Fast) + "," + defMode() + "," + defMode(null) + "," + modeLetter);

		var fn:Name->String = nameToString;
		var fromText:String->Name = stringToName;
		var closure = function(n:Name):Name {
			var s:String = n;
			return s + "!";
		};
		function pick(n:Name, ?extra:Name):Name {
			return extra == null ? n : extra;
		}
		Sys.println("function=" + fn("f1") + "," + fromText("f2") + "," + closure("f3") + "," + pick("f4") + "," + pick("f4", "f5"));

		var entry:Entry = {name: "anon", color: Red};
		var entry2:Entry = {name: "anon2"};
		Sys.println("anon=" + entry.name + "," + entry.color + "," + entry2.name + "," + entry2.color);

		var wrapped = W("wrapped", Blue);
		switch wrapped {
			case W(n, wc):
				Sys.println("enum=" + nameToString(n) + "," + wc);
		}

		var d:Dynamic = local;
		var fromDyn:Name = d;
		var fromDynText:String = fromDyn;
		Sys.println("dynamic=" + fromDynText + "," + Std.string(d) + "," + Std.isOfType(d, String));

		Sys.println("shapes=" + tagText("tag") + "," + labelText(new Label("label")) + "," + codeText(new Code("code")) + "," + wordInfo("word") + ","
			+ textInfo("foo boo") + "," + aliasText("alias") + "," + plainText("plain") + "," + boxText("box"));

		var counts:Array<Count> = [1, 2, 3];
		var total = 0;
		for (k in counts) {
			var kv:Int = k;
			total += kv;
		}
		var ratios:Map<String, Ratio> = ["r" => 0.5];
		var ratio:Float = ratios.get("r");
		Sys.println("int=" + countOf(5) + "," + total + "," + optCount() + "," + optCount(4) + "," + levelOf(High));
		Sys.println("float.bool=" + ratioOf(2.5) + "," + ratio + "," + flagOf(true) + "," + flagOf(false));
		var intBox:Box<Int> = 9;
		var intBoxValue:Int = intBox;
		Sys.println("dynamicAbstract=" + looseText("text") + "," + looseText(42) + "," + intBoxValue);
	}
}
