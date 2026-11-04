/** Context cancellation proves that a callback in a native tuple remains callable. */
@:go.import("context")
@:go.name("Context")
extern interface NativeContext {
	@:go.name("Err") public function error():go.Error;
}

@:go.import("context")
extern class NativeContexts {
	@:go.name("Background") public static function background():NativeContext;
	@:go.name("WithCancel") @:go.tupleReturn
	public static function withCancel(parent:NativeContext):CancelPair;
}

class CancelPair {
	public final context:NativeContext;
	public final cancel:Void->Void;

	public function new(context:NativeContext, cancel:Void->Void) {
		this.context = context;
		this.cancel = cancel;
	}
}

class Main {
	static function main() {
		final pair = NativeContexts.withCancel(NativeContexts.background());
		Sys.println(pair.context.error() == null);
		final cancel = pair.cancel;
		final alias = cancel;
		cancel();
		alias();
		Sys.println(pair.context.error() != null);
		Sys.println(Reflect.isFunction(cancel));
		Sys.println(Reflect.compareMethods(cancel, alias));
	}
}
