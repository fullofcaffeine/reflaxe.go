import haxe.io.Bytes;
import haxe.io.Encoding;

class Main {
	static function main() {
		var s:UnicodeString = "a😀bé";
		Sys.println("length=" + s.length);
		Sys.println("char1=" + s.charAt(1));
		Sys.println("code1=" + s.charCodeAt(1));
		Sys.println("substr=" + s.substr(1, 2));
		Sys.println("substring=" + s.substring(1, 3));
		Sys.println("substring.swap=" + s.substring(3, 1));
		Sys.println("substring.neg=" + s.substring(-2, 2));
		Sys.println("substring.omit=" + s.substring(2));
		Sys.println("substr.neglen=" + s.substr(1, -1));
		Sys.println("substr.negpos=" + s.substr(-2, 2));
		Sys.println("index=" + s.indexOf("bé"));
		Sys.println("index.empty=" + s.indexOf(""));
		Sys.println("index.startNeg=" + s.indexOf("bé", -2));
		Sys.println("last=" + s.lastIndexOf("a"));
		Sys.println("last.empty=" + s.lastIndexOf(""));
		Sys.println("last.start=" + s.lastIndexOf("bé", 2));
		Sys.println("valid.utf8=" + UnicodeString.validate(Bytes.ofString("ok"), Encoding.UTF8));
		Sys.println("valid.invalid=" + UnicodeString.validate(Bytes.ofHex("ff"), Encoding.UTF8));
		try {
			UnicodeString.validate(Bytes.ofString("ok"), Encoding.RawNative);
			Sys.println("valid.raw=ok");
		} catch (error:Dynamic) {
			Sys.println("valid.raw=" + Std.string(error));
		}
	}
}
