import go.NativePointer;

class Main {
	static function main():Void {
		final pointer = NativePointer.of(7);
		final read = pointer.get;
		Sys.println(read());
	}
}
