import sys.net.Address;
import sys.net.Host;
import sys.ssl.DigestAlgorithm;

class Main {
	static function main() {
		var host = new Host("127.0.0.1");
		var addr = new Address();
		addr.host = host.ip;
		addr.port = 3210;
		Sys.println("snapshot.host=" + addr.getHost().toString());
		Sys.println("snapshot.compare=" + addr.compare(addr.clone()));
		Sys.println("snapshot.alg=" + DigestAlgorithm.SHA224 + "," + DigestAlgorithm.SHA384 + "," + DigestAlgorithm.RIPEMD160);
	}
}
