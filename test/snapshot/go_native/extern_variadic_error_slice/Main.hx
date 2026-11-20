/**
	Regression for haxe_go-ww4u: error values packed into an imported Go
	`...error` parameter must reach Go as `[]error`.

	Haxe stores `go.Error` as a pointer to its own class and `go.NativeError` as
	a pointer to a runtime carrier, so both packs used to build a slice of those
	pointers, which Go rejects for `...error`. A null value must also arrive as
	a nil interface: a typed nil pointer would make `errors.Join` keep it.
**/

import go.NativeError;

/** `go.Error` is a message snapshot: each crossing copies only the text. */
@:go.import("errors")
extern class LegacyErrors {
	@:go.name("Join") public static function join(errors:haxe.Rest<go.Error>):go.Error;
	@:go.name("New") public static function create(message:String):go.Error;
	@:go.name("Unwrap") public static function unwrap(error:go.Error):go.Error;
}

/** `go.NativeError` keeps the original Go value, so `errors.Is` still matches. */
@:go.import("errors")
extern class NativeErrors {
	@:go.name("Join") public static function join(errors:haxe.Rest<NativeError>):NativeError;
	@:go.name("New") public static function create(message:String):NativeError;
	@:go.name("Is") public static function matches(error:NativeError, target:NativeError):Bool;
}

@:goNative
class Main {
	static var calls = new Array<String>();

	static function legacy(message:String):go.Error {
		calls.push(message);
		return new go.Error(message);
	}

	static function describe(error:Null<go.Error>):String {
		return error == null ? "null" : StringTools.replace(error.message, "\n", "|");
	}

	/** Forwards a Haxe rest parameter, whose Haxe slice storage must convert too. */
	static function joinAll(...errors:go.Error):go.Error {
		return LegacyErrors.join(...errors);
	}

	/** A local named `error` hides Go's `error` type, so the slice must not name it. */
	static function joinWithNull(error:go.Error):go.Error {
		return LegacyErrors.join(error, null);
	}

	static function main():Void {
		final first = new go.Error("first");
		final second = LegacyErrors.create("second");
		final none:go.Error = null;
		Sys.println("literal=" + describe(LegacyErrors.join(first, none, second)));
		Sys.println("all-null=" + describe(LegacyErrors.join(none, null)));
		final list = [first, none, second];
		Sys.println("spread=" + describe(LegacyErrors.join(...list)));
		final nulls:Array<go.Error> = [null, none];
		Sys.println("spread-null=" + describe(LegacyErrors.join(...nulls)));
		Sys.println("nested=" + describe(LegacyErrors.join(LegacyErrors.join(first, none), second)));
		Sys.println("rest=" + describe(joinAll(first, none, second)));
		Sys.println("shadow=" + describe(joinWithNull(second)));
		Sys.println("unwrap=" + describe(LegacyErrors.unwrap(first)));
		Sys.println("order=" + describe(LegacyErrors.join(legacy("a"), none, legacy("b"))) + " " + calls.join(","));

		final x = NativeErrors.create("x");
		final y = NativeErrors.create("y");
		final missing:NativeError = null;
		final joined = NativeErrors.join(x, missing, y);
		Sys.println("native=" + StringTools.replace(joined.text(), "\n", "|"));
		Sys.println("native-is=" + NativeErrors.matches(joined, x) + "," + NativeErrors.matches(joined, y));
		Sys.println("native-is-copy=" + NativeErrors.matches(joined, NativeErrors.create("x")));
		final natives = [x, missing, y];
		Sys.println("native-spread-is=" + NativeErrors.matches(NativeErrors.join(...natives), y));
		Sys.println("native-nested-is=" + NativeErrors.matches(NativeErrors.join(NativeErrors.join(x), y), x));
		Sys.println("native-null=" + (NativeErrors.join(missing, null) == null));
	}
}
