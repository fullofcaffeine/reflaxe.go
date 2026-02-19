import sys.net.Host;

class Main {
	static function renderReverse(label:String, host:Host):Void {
		var ok = true;
		var value = "";
		try {
			value = host.reverse();
		} catch (_:Dynamic) {
			ok = false;
		}
		Sys.println(label + ".reverse_ok=" + ok);
		if (ok) {
			Sys.println(label + ".reverse=" + value);
		}
	}

	static function main() {
		var loop = new Host("127.0.0.1");
		Sys.println("loop.to=" + loop.toString());
		renderReverse("loop", loop);

		var named = new Host("localhost");
		Sys.println("named.to=" + named.toString());

		var zero = new Host("0.0.0.0");
		renderReverse("zero", zero);

		var invalidConstructOk = true;
		try {
			new Host("256.256.256.256");
		} catch (_:Dynamic) {
			invalidConstructOk = false;
		}
		Sys.println("invalid.construct_ok=" + invalidConstructOk);

		Sys.println("localhost_nonempty=" + (Host.localhost() != ""));
	}
}
