import sys.net.Address;
import sys.net.Host;
import sys.ssl.DigestAlgorithm;

class Main {
	static function main() {
		var host = new Host("127.0.0.1");
		var addr = new Address();
		addr.host = host.ip;
		addr.port = 9000;
		var clone = addr.clone();
		Sys.println("addr.host=" + addr.getHost().toString());
		Sys.println("addr.port=" + clone.port);
		Sys.println("cmp.same=" + addr.compare(clone));
		clone.port = 9001;
		Sys.println("cmp.port=" + addr.compare(clone));
		Sys.println("alg.sha256=" + DigestAlgorithm.SHA256);
		Sys.println("alg.sha512=" + DigestAlgorithm.SHA512);
	}
}
