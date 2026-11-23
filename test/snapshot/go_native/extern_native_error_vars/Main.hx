import go.NativeError;

/**
	Go package-level variables bound as extern static vars.
	A read crosses the same import boundary as a native call result, so `io.EOF`
	keeps its native identity for `errors.Is`. A write exports the carrier again.
**/
@:go.import("io")
@:go.package("io")
extern class IoVars {
	@:go.name("EOF") public static var eof(default, never):NativeError;
}

@:go.import("context") @:go.name("Context") extern interface NativeContext {}

@:go.import("context")
extern class ContextVars {
	@:go.name("Canceled") public static var canceled(default, never):NativeError;
	@:go.name("Background") public static function background():NativeContext;
	@:go.name("WithCancelCause") @:go.tupleReturn public static function withCause(context:NativeContext):CancelPair;
	@:go.name("Cause") public static function cause(context:NativeContext):NativeError;
}

class CancelPair {
	public final context:NativeContext;
	public final cancel:NativeError->Void;

	public function new(context:NativeContext, cancel:NativeError->Void) {
		this.context = context;
		this.cancel = cancel;
	}
}

@:go.import("errors")
extern class ErrorsPkg {
	/** A writable package variable; this program replaces it and restores it. */
	@:go.name("ErrUnsupported") public static var unsupported:NativeError;

	@:go.name("New") public static function create(message:String):NativeError;
	@:go.name("Is") public static function is(error:NativeError, target:NativeError):Bool;
	@:go.name("Join") public static function join(errors:haxe.Rest<NativeError>):NativeError;
}

/** Go's image.ZP is a Point value, not a pointer, so it is marked as a value. */
@:go.import("image")
@:go.name("Point")
@:go.struct
extern class GoPoint {
	public function new();

	@:go.name("X") public var x:Int;
	@:go.name("Y") public var y:Int;
}

@:go.import("image")
extern class ImageVars {
	@:go.name("ZP") @:go.valueField public static var origin:GoPoint;
}

@:go.import("time")
extern class TimeVars {
	@:go.name("RFC3339") public static var rfc3339(default, never):String;
}

class Holder {
	public var error:NativeError;

	public function new(error:NativeError) {
		this.error = error;
	}
}

@:goNative
class Main {
	static var last:NativeError = null;

	static function check(ok:Bool, label:String):Void {
		if (!ok)
			throw label;
	}

	static function describe(error:NativeError):String {
		return error == null ? "nil" : error.text();
	}

	static function measure(text:String):Int {
		return text.length;
	}

	static function sum(point:GoPoint):Int {
		return point.x + point.y;
	}

	static function main():Void {
		final eof = IoVars.eof;
		Sys.println(eof.text());
		Sys.println(ErrorsPkg.is(eof, IoVars.eof));

		Sys.println(IoVars.eof.text());
		check(ErrorsPkg.is(IoVars.eof, IoVars.eof), "direct identity");
		check(!ErrorsPkg.is(IoVars.eof, ContextVars.canceled), "distinct variables");
		check(IoVars.eof != null && !(IoVars.eof == null), "non-null read");
		Sys.println(describe(ContextVars.canceled));

		final holder = new Holder(IoVars.eof);
		check(ErrorsPkg.is(holder.error, IoVars.eof), "field storage");
		holder.error = ContextVars.canceled;
		check(ErrorsPkg.is(holder.error, ContextVars.canceled), "field write");
		last = IoVars.eof;
		check(ErrorsPkg.is(last, eof), "static storage");
		final stored:Array<NativeError> = [IoVars.eof, ContextVars.canceled];
		stored.push(IoVars.eof);
		check(stored.length == 3
			&& ErrorsPkg.is(stored[0], IoVars.eof)
			&& ErrorsPkg.is(stored[1], ContextVars.canceled)
			&& ErrorsPkg.is(stored[2], eof),
			"array storage");

		final joined = ErrorsPkg.join(IoVars.eof, ContextVars.canceled);
		check(ErrorsPkg.is(joined, IoVars.eof) && ErrorsPkg.is(joined, ContextVars.canceled), "variadic native argument");
		final pair = ContextVars.withCause(ContextVars.background());
		pair.cancel(null);
		check(ErrorsPkg.is(ContextVars.cause(pair.context), ContextVars.canceled), "native result matches variable");

		final original = ErrorsPkg.unsupported;
		final replacement = ErrorsPkg.create("replacement");
		ErrorsPkg.unsupported = replacement;
		check(ErrorsPkg.is(ErrorsPkg.unsupported, replacement), "variable write");
		final restored = (ErrorsPkg.unsupported = original);
		check(ErrorsPkg.is(restored, original) && ErrorsPkg.is(ErrorsPkg.unsupported, original), "variable write result");
		ErrorsPkg.unsupported = null;
		check(ErrorsPkg.unsupported == null, "nil variable write");
		ErrorsPkg.unsupported = original;

		final origin = ImageVars.origin;
		origin.x = 7;
		check(origin.x == 7 && ImageVars.origin.x == 0 && sum(ImageVars.origin) == 0, "struct value variable copy");
		final moved = new GoPoint();
		moved.y = 5;
		ImageVars.origin = moved;
		moved.y = 9;
		check(sum(ImageVars.origin) == 5, "struct value variable write");
		ImageVars.origin = new GoPoint();

		Sys.println(TimeVars.rfc3339);
		check(measure(TimeVars.rfc3339) == 25, "string variable");
		Sys.println("native-error-vars:ok");
	}
}
