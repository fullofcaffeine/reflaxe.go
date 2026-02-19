import sys.net.Host;
import sys.net.Socket;

private typedef BoundServer = {
	var server:Socket;
	var port:Int;
}

class Main {
	static function safeClose(socket:Socket):Void {
		if (socket == null) {
			return;
		}
		try {
			socket.close();
		} catch (_:Dynamic) {}
	}

	static function bindServer():BoundServer {
		var host = new Host("127.0.0.1");
		var server = new Socket();
		server.bind(host, 0);
		server.listen(1);
		var info = server.host();
		if (info == null || info.port <= 0) {
			safeClose(server);
			throw "failed to resolve loopback port";
		}
		return {server: server, port: info.port};
	}

	static function main() {
		var bound = bindServer();
		var client = new Socket();
		var peer:Socket = null;
		var failed:Dynamic = null;

		try {
			client.connect(new Host("127.0.0.1"), bound.port);
			peer = bound.server.accept();

			client.output.writeString("ping\n");
			client.output.flush();
			var serverRead = peer.input.readLine();

			peer.output.writeString("pong:" + serverRead + "\n");
			peer.output.flush();
			var clientRead = client.input.readLine();

			Sys.println("serverRead=" + serverRead);
			Sys.println("clientRead=" + clientRead);
		} catch (error:Dynamic) {
			failed = error;
		}

		safeClose(peer);
		safeClose(client);
		safeClose(bound.server);
		if (failed != null) {
			throw failed;
		}
	}
}
