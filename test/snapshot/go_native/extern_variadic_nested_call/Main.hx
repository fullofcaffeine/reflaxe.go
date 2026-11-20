/**
	Regression for haxe_go-jvzb: values packed into an imported Go `...T`
	parameter must reach Go as the declared native element type.

	The typer packs simple values directly, but hoists the pack into a temporary
	when a value is a call, and passes an existing Array for a spread. Both forms
	used to keep the Haxe `[]*string` storage, which Go rejects for `...string`.
**/
@:go.import("strings") @:go.name("Replacer")
extern class NativeReplacer {
	@:go.name("Replace") public function replace(value:String):String;
}

@:go.import("strings")
extern class NativeStrings {
	@:go.name("NewReplacer") public static function newReplacer(oldNew:haxe.Rest<String>):NativeReplacer;
	@:go.name("ToUpper") public static function toUpper(value:String):String;
}

@:go.import("path/filepath")
extern class NativeFilepath {
	@:go.name("Join") public static function join(elements:haxe.Rest<String>):String;
	@:go.name("Dir") public static function dir(path:String):String;
	@:go.name("Base") public static function base(path:String):String;
}

@:go.import("os/exec") @:go.name("Cmd")
extern class NativeCmd {
	@:go.name("String") public function describe():String;
}

/** `exec.Command(name, arg...)`: a fixed parameter before the variadic one. */
@:go.import("os/exec")
extern class NativeExec {
	@:go.name("Command") public static function command(name:String, arguments:haxe.Rest<String>):NativeCmd;
}

/** An erased `...any` parameter still needs Haxe strings as Go string values. */
@:go.import("fmt")
extern class NativeFormat {
	@:go.name("Sprint") public static function sprint(values:haxe.Rest<Dynamic>):String;
}

/** A narrower type than the declared element: it must convert to a Go string. */
abstract FileName(String) to String {
	public inline function new(value:String)
		this = value;
}

class Main {
	static var calls = new Array<String>();

	static function name():String {
		calls.push("name");
		return "config.yaml";
	}

	static function fileName():FileName {
		calls.push("fileName");
		return new FileName("local.yaml");
	}

	static function arguments():Array<String> {
		return ["-n", name()];
	}

	/** Forwards a Haxe rest parameter, whose `[]*string` storage must convert too. */
	static function joinAll(...parts:String):String {
		return NativeFilepath.join(...parts);
	}

	static function main():Void {
		Sys.println(NativeFilepath.join(NativeFilepath.dir("/a/b/c.txt"), "d.txt"));
		Sys.println(NativeFilepath.join("/root", name()));
		Sys.println(NativeFilepath.join("/root", fileName()));
		Sys.println(NativeFilepath.join(name(), "x", NativeFilepath.base("/q/y")));
		Sys.println(NativeFilepath.join(...[NativeFilepath.dir("/a/b/c.txt"), name()]));
		final rest = ["e", "f"];
		Sys.println(NativeExec.command(NativeFilepath.join("/bin", "echo"), ...rest).describe());
		Sys.println(NativeExec.command("/bin/echo", ...arguments()).describe());
		Sys.println(NativeExec.command("/bin/echo", NativeFilepath.base("/q/y"), name()).describe());
		final replacer = NativeStrings.newReplacer(NativeStrings.toUpper("a"), "1", fileName(), "2");
		Sys.println(replacer.replace("A local.yaml"));
		Sys.println(joinAll("p", name(), fileName()));
		Sys.println(NativeFormat.sprint("<", name(), "|", fileName(), ">"));
		calls = [];
		NativeFilepath.join(fileName(), name(), fileName());
		Sys.println("order=" + calls.join(","));
	}
}
