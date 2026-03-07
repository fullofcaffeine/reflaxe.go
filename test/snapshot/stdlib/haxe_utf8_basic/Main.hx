class Main {
	static function main() {
		var utf8 = new haxe.Utf8();
		utf8.addChar('a'.code);
		utf8.addChar(0x1F600);
		utf8.addChar('é'.code);
		var value = utf8.toString();

		Sys.println("string=" + value);
		Sys.println("length=" + haxe.Utf8.length(value));
		Sys.println("sub=" + haxe.Utf8.sub(value, 1, 2));
		Sys.println("encode=" + haxe.Utf8.encode("é"));
		Sys.println("decode=" + haxe.Utf8.decode(haxe.Utf8.encode("é")));
	}
}
