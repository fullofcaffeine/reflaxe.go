@:forward
abstract Name(String) from String to String {}

abstract Label(String) {
	public inline function new(text:String)
		this = text;

	public function toString():String
		return "label:" + this;
}

enum abstract Color(String) to String {
	var Red = "red";
}

class Holder {
	public var text:String;

	public function new(text:String)
		this.text = text;
}

/**
	`toString()` on a typed String must behave as in `--interp`
	(haxe_go-676g).

	`s.toString()` returns the same String, for a local, a call result (which
	runs once, also in a statement), a ternary, an Array element, a class field,
	a structure field and a `Dynamic` value read into a String. A forwarding
	String abstract calls it too, an abstract with its own `toString` keeps it,
	and `Std.string` agrees. On a null String it raises `Null Access`. Left
	out: a String method read as a value (`var f = s.toString`), which python
	rejects and haxe.go does not lower yet.
**/
class Main {
	static var calls = 0;

	static function next():String {
		calls++;
		return "n" + calls;
	}

	static function nothing():String {
		return null;
	}

	static function main() {
		var s = "abc";
		Sys.println(s.toString());
		Sys.println(s.toString() == s);
		Sys.println(s.toString().length + " " + s.toString().toUpperCase());
		Sys.println(next().toString() + " " + calls);
		next().toString();
		Sys.println("statement " + calls);
		Sys.println((calls > 1 ? s : "other").toString());
		Sys.println('${s.toString()}!');
		var words = ["p", "q"];
		Sys.println(words[1].toString());
		var holder = new Holder("field");
		Sys.println(holder.text.toString());
		var anon = {text: "anon"};
		Sys.println(anon.text.toString());
		var name:Name = "bob";
		Sys.println(name.toString() + " " + name.toUpperCase());
		var label = new Label("x");
		Sys.println(label.toString() + " " + Std.string(label));
		var color = Color.Red;
		Sys.println((color : String).toString());
		var dynamicText:Dynamic = "dyn";
		Sys.println(dynamicText.toString());
		var typedFromDynamic:String = dynamicText;
		Sys.println(typedFromDynamic.toString());
		Sys.println(Std.string(s) + " " + Std.string(s.toString()));
		var empty = "";
		Sys.println("[" + empty.toString() + "]");
		var unicode = "héllo";
		Sys.println(unicode.toString().length);
		var missing:String = nothing();
		Sys.println(Std.string(missing));
		try {
			Sys.println("null toString " + missing.toString());
		} catch (e:String) {
			Sys.println("null toString threw " + e);
		}
		try {
			missing.toString();
			Sys.println("null statement ran");
		} catch (e:String) {
			Sys.println("null statement threw " + e);
		}
		var optional:Null<String> = nothing();
		try {
			var value = optional.toString();
			Sys.println("null optional " + value);
		} catch (e:String) {
			Sys.println("null optional threw " + e);
		}
	}
}
