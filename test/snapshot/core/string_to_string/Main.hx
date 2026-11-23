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
	`toString()` on a typed String (haxe_go-676g).

	haxe.go lowered `s.toString()` on a typed String to the Go selector
	`s.toString` on `*string`, which does not build: the String instance call
	lowering had no `toString` case.

	Now the call returns the receiver itself, which runs once, as Haxe returns
	the same String; a null String raises `Null Access`, as in `--interp`.
	Forwarding String abstracts take the same path, an abstract with its own
	`toString` keeps it, and `Std.string` agrees.
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
