import go.NativePointer;

@:go.import("fmt")
extern class NativeFmt {
	// fmt.Sscan declares ...any; only the native pointer carriers are concrete.
	@:go.name("Sscan") static function scan(input:String, destinations:haxe.Rest<Dynamic>):Void;
}

@:goNative
class Main {
	static function readNumber(pointer:NativePointer<Int>):Int
		return pointer.get();

	static function main():Void {
		var original = 3;
		final number = NativePointer.of(original);
		final independent = NativePointer.of(number.get());
		final word = NativePointer.of("before");
		final flag = NativePointer.of(false);

		NativeFmt.scan("7 after true", number, word, flag);
		Sys.println("mutation=" + (number.get() == 7 && word.get() == "after" && flag.get()));
		Sys.println("independent=" + (original == 3 && independent.get() == 3));

		final alias = number;
		NativeFmt.scan("9", alias);
		Sys.println("identity=" + (number.get() == 9));
		Sys.println("typed-parameter=" + (readNumber(number) == 9));

		final captured = word.get();
		NativeFmt.scan("latest", word);
		Sys.println("string-snapshot=" + (captured == "after" && word.get() == "latest"));

		var inputCalls = 0;
		final evaluated = NativePointer.of(++inputCalls);
		var receiverCalls = 0;
		final read = NativePointer.of(++receiverCalls).get();
		Sys.println("evaluation=" + (inputCalls == 1 && evaluated.get() == 1 && receiverCalls == 1 && read == 1));
	}
}
