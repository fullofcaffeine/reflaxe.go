import go.NativeError;

@:go.import("errors")
extern class NativeErrors {
	@:go.name("New") public static function create(message:String):NativeError;
	@:go.name("Is") public static function matches(value:NativeError, target:NativeError):Bool;
	@:go.name("Unwrap") public static function unwrap(value:NativeError):NativeError;
	@:go.name("AsType") @:go.tupleReturn public static function asType<T>(value:NativeError):Match<T>;
}

class Match<T> {
	public final value:T;
	public final matched:Bool;

	public function new(value:T, matched:Bool) {
		this.value = value;
		this.matched = matched;
	}
}

class Holder<T> {
	public final value:T;

	public function new(value:T) {
		this.value = value;
	}
}

@:go.import("context") @:go.name("Context") extern interface NativeContext {}

@:go.import("context") extern class Contexts {
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

@:go.import("io/fs") @:go.name("FileInfo") extern interface FileInfo {}

@:go.import("io/fs") @:go.name("PathError") @:go.struct extern class PathError {
	@:go.name("Err") public var error:NativeError;
}

@:go.import("path/filepath") extern class Paths {
	@:go.name("Walk") public static function walk(path:String, visit:(String, FileInfo, NativeError) -> NativeError):NativeError;
}

@:go.import("sync") extern class NativeSync {
	@:go.name("OnceValue") public static function once<T>(factory:Void->T):Void->T;
}

@:go.import("fmt") extern class NativeFormat {
	/** Go's fmt.Errorf accepts ...any; the returned error is narrowed to NativeError. */
	@:go.name("Errorf") public static function errorf(format:String, values:haxe.Rest<Dynamic>):NativeError;
}

/** The matching API is Go 1.26; transport itself retains the Go 1.22 floor. */
@:goBuildConstraint("go1.26")
@:goNative
class ErrorProbe {
	static function check(ok:Bool, label:String):Void {
		if (!ok)
			throw label;
	}

	public static function run():Void {
		final root = NativeErrors.create("root");
		check(root.text() == "root", "explicit text");
		check(NativeErrors.matches(root, root), "direct identity");
		check(!NativeErrors.matches(root, NativeErrors.create("root")), "distinct equal text");
		check(NativeErrors.matches(null, null), "nil export");
		final create:String->NativeError = NativeErrors.create;
		final matches:(NativeError, NativeError) -> Bool = NativeErrors.matches;
		final stored = create("stored");
		check(matches(stored, stored), "stored native function");
		final pair = Contexts.withCause(Contexts.background());
		check(Contexts.cause(pair.context) == null, "nil scalar result");
		pair.cancel(root);
		check(matches(Contexts.cause(pair.context), root), "function inside native tuple");
		final once = NativeSync.once(() -> root);
		check(matches(once(), root), "returned generic callback");
		var observed:NativeError = null;
		final failure = Paths.walk("\x00", (_, _, error) -> {
			observed = error;
			return error;
		});
		check(observed != null && matches(failure, observed), "callback arguments and return");
		final typed:Match<PathError> = NativeErrors.asType(failure);
		check(typed.matched && typed.value != null, "concrete matching");
		typed.value.error = root;
		check(matches(typed.value.error, root) && matches(failure, root), "native field read/write");
		final generic:Match<NativeError> = NativeErrors.asType(root);
		check(generic.matched && matches(generic.value, root), "native generic type argument and tuple slot");
		final asNative:NativeError->Match<NativeError> = NativeErrors.asType;
		final storedMatch = asNative(root);
		check(storedMatch.matched && matches(storedMatch.value, root), "stored generic tuple function");
		final withCause:NativeContext->CancelPair = Contexts.withCause;
		final storedPair = withCause(Contexts.background());
		storedPair.cancel(root);
		check(matches(Contexts.cause(storedPair.context), root), "stored tuple function");
		final nilPointer:PathError = null;
		final wrappedNil = NativeFormat.errorf("typed nil: %w", nilPointer);
		final typedNil = NativeErrors.unwrap(wrappedNil);
		final nilMatch:Match<PathError> = NativeErrors.asType(typedNil);
		check(nilMatch.matched && nilMatch.value == null, "typed nil matches independently of value");
		check(typedNil != null
			&& new Holder(typedNil).value != null && new Holder(nilPointer).value == null, "typed nil through generic storage");
		final holder = new Holder(root);
		check(holder.value != null && matches(holder.value, root), "generic storage");
		final legacy = new go.Error("legacy");
		check(legacy.message == "legacy" && legacy.toString() == "legacy", "legacy message");
		check(go.Result.failure("old").error() == "old", "legacy result");
		Sys.println("native-error:ok");
	}
}
