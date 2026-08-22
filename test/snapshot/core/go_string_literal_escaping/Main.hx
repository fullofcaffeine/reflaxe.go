import go.Fmt;

class Main {
	static function main() {
		final value = "\x00A\x01\x07\x08\x0c\n\r\t\x0b\x1f\x7f\"\\é🙂";
		final bytes = haxe.io.Bytes.ofString(value);
		Fmt.println(bytes.length);
		for (index in 0...bytes.length) {
			Fmt.println(bytes.get(index));
		}
	}
}
