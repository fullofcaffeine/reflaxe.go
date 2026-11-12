import go.NativePointer;

class Main {
	static function usePointer(value:NativePointer<Float>):Void {
		Sys.println(value == null);
	}

	static function main():Void {
		usePointer(null);
	}
}
