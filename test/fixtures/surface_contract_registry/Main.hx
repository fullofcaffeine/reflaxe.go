class Main {
	static function main() {
		final values:Array<Int> = [1, 2, 3];
		final label:String = values.join(",");
		final bytes = haxe.io.Bytes.ofString(label);
		trace(bytes.length);
	}
}
