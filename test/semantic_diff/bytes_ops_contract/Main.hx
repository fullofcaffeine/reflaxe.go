class Main {
	static function main() {
		var base = haxe.io.Bytes.ofString("abcde");
		var src = haxe.io.Bytes.ofString("XYZ");

		base.blit(1, src, 0, 2);
		Sys.println("blit=" + base.toString());

		base.fill(3, 2, 97);
		Sys.println("fill=" + base.toString());

		var mid = base.sub(1, 3);
		Sys.println("sub=" + mid.toString());

		Sys.println("cmpEq=" + base.compare(haxe.io.Bytes.ofString("aXYaa")));
		Sys.println("cmpLt=" + base.compare(haxe.io.Bytes.ofString("b0000")));
		Sys.println("cmpGt=" + base.compare(haxe.io.Bytes.ofString("aXXzz")));
	}
}
