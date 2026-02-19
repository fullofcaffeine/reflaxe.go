import haxe.crypto.Base64;
import haxe.crypto.Md5;
import haxe.crypto.Sha1;
import haxe.crypto.Sha224;
import haxe.crypto.Sha256;
import haxe.io.Bytes;
import haxe.xml.Parser;
import haxe.xml.Printer;
import haxe.zip.Compress;
import haxe.zip.Uncompress;

class Main {
	static function main() {
		var payload = "ab";
		var bytes = Bytes.ofString(payload);

		Sys.println(Base64.encode(bytes));
		Sys.println(Base64.encode(bytes, false));
		Sys.println(Base64.decode("YWI=").toString());
		Sys.println(Base64.urlEncode(bytes, true));
		Sys.println(Base64.urlDecode("YWI").toString());

		Sys.println(Md5.encode(payload));
		Sys.println(Sha1.encode(payload));
		Sys.println(Sha224.encode(payload));
		Sys.println(Sha256.encode(payload));

		var doc = Parser.parse('<root><item n="1">x</item></root>');
		Sys.println(Printer.print(doc));

		var compressed = Compress.run(bytes, 9);
		var roundtrip = Uncompress.run(compressed);
		Sys.println(roundtrip.toString());
		Sys.println(compressed.length > 0);
	}
}
