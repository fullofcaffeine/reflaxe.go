import go.NativePointer;

class Main {
	static function main():Void {
		final create:Int->NativePointer<Int> = NativePointer.of;
		Sys.println(create(7).get());
	}
}
