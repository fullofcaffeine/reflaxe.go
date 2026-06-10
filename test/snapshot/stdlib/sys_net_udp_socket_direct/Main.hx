import haxe.io.Bytes;
import sys.net.Address;
import sys.net.Host;
import sys.net.UdpSocket;

class Main {
	static function safeClose(socket:UdpSocket):Void {
		if (socket == null) {
			return;
		}
		try {
			socket.close();
		} catch (_:Dynamic) {}
	}

	static function main() {
		var server = new UdpSocket();
		var client = new UdpSocket();
		var failure:Dynamic = null;

		try {
			server.bind(new Host("127.0.0.1"), 0);
			var bound = server.host();
			if (bound == null || bound.port <= 0) {
				throw "missing bound udp port";
			}

			server.setBlocking(true);
			client.setBroadcast(true);

			var sent = Bytes.ofString("udp-ping");
			var target = new Address();
			target.host = new Host("127.0.0.1").ip;
			target.port = bound.port;
			var wrote = client.sendTo(sent, 0, sent.length, target);

			var recv = Bytes.alloc(32);
			var remote = new Address();
			var read = server.readFrom(recv, 0, recv.length, remote);

			Sys.println("bound.host=" + bound.host.toString());
			Sys.println("bound.port.positive=" + (bound.port > 0));
			Sys.println("wrote=" + wrote);
			Sys.println("read=" + read);
			Sys.println("payload=" + recv.sub(0, read).toString());
			Sys.println("remote.port.positive=" + (remote.port > 0));
			Sys.println("remote.host=" + remote.getHost().toString());
		} catch (error:Dynamic) {
			failure = error;
		}

		safeClose(client);
		safeClose(server);
		if (failure != null) {
			throw failure;
		}
	}
}
