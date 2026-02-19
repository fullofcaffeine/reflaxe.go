import sys.net.Host;

class Main {
	static function main() {
		var loop = new Host("127.0.0.1");
		Sys.println("loop.to_nonempty=" + (loop.toString() != ""));

		var named = new Host("localhost");
		Sys.println("named.to_nonempty=" + (named.toString() != ""));

		var invalidThrows = false;
		try {
			new Host("256.256.256.256");
		} catch (_:Dynamic) {
			invalidThrows = true;
		}
		Sys.println("invalid_throws=" + invalidThrows);

		Sys.println("localhost_nonempty=" + (Host.localhost() != ""));

		try {
			loop.reverse();
		} catch (_:Dynamic) {}
		Sys.println("reverse_called=true");
	}
}
