import haxe.io.Bytes;
import sys.ssl.Digest;
import sys.ssl.DigestAlgorithm;

class Main {
	static function main() {
		var out = Digest.make(Bytes.ofString("ssl"), DigestAlgorithm.SHA256);
		Sys.println(out.toHex());
	}
}
