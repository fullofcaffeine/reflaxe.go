import haxe.crypto.Base64;
import haxe.crypto.Md5;
import haxe.crypto.Sha1;
import haxe.crypto.Sha224;
import haxe.crypto.Sha256;
import haxe.io.Bytes;

class Main {
	static function invalidBase64Throws(value:String):Bool {
		var threw = false;
		try {
			Base64.decode(value);
		} catch (_:Dynamic) {
			threw = true;
		}
		return threw;
	}

	static function main() {
		var bytes = Bytes.ofString("ab");
		var binary = Bytes.alloc(4);
		binary.set(0, 0);
		binary.set(1, 127);
		binary.set(2, 128);
		binary.set(3, 255);

		Sys.println(Base64.encode(bytes));
		Sys.println(Base64.encode(bytes, false));
		Sys.println(Base64.decode("YWI=").toString());
		Sys.println(Base64.urlEncode(bytes, true));
		Sys.println(Base64.urlDecode("YWI").toString());
		Sys.println(Base64.encode(binary));
		Sys.println(Base64.decode("").length);
		Sys.println(invalidBase64Throws("%%%"));

		Sys.println(Md5.encode("ab"));
		Sys.println(Sha1.encode("ab"));
		Sys.println(Sha224.encode("ab"));
		Sys.println(Sha256.encode("ab"));
		Sys.println(Md5.make(bytes).toHex());
		Sys.println(Sha1.make(bytes).toHex());
		Sys.println(Sha224.make(bytes).toHex());
		Sys.println(Sha256.make(bytes).toHex());
	}
}
