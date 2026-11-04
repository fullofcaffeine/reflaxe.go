/** Named native error type used only as a generic result argument. */
@:go.import("io/fs")
@:go.name("PathError")
@:go.struct
extern class PathError {}

@:go.import("errors")
extern class NativeErrors {
	@:go.name("AsType")
	@:go.tupleReturn
	public static function asType<T>(error:go.Error):ErrorMatch<T>;
}

/** Haxe owns the tuple carrier; the native call owns its concrete type argument. */
class ErrorMatch<T> {
	public final value:T;
	public final matched:Bool;

	public function new(value:T, matched:Bool) {
		this.value = value;
		this.matched = matched;
	}
}

/** Phantom T records TypeFor's requested type without changing reflect.Type's ABI. */
@:go.import("reflect")
@:go.name("Type")
extern interface NativeType<T> {
	@:go.name("String")
	public function name():String;
}

@:go.import("reflect")
extern class NativeReflect {
	@:go.name("TypeFor")
	public static function typeFor<T>():NativeType<T>;
}

@:goNative
class Main {
	static function main() {
		NativeErrorProbe.check();
		final scalar:NativeType<Int> = NativeReflect.typeFor();
		Sys.println(scalar.name());
		final named:NativeType<PathError> = NativeReflect.typeFor();
		Sys.println(named.name());
		final factory:Void->NativeType<String> = NativeReflect.typeFor;
		Sys.println(factory().name());
	}
}
