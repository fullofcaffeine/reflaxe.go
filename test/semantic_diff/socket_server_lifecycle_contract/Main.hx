import sys.net.Host;
import sys.net.Socket;

class Main {
	static function safeClose(socket:Socket):Void {
		if (socket == null)
			return;
		try {
			socket.close();
		} catch (_:Dynamic) {}
	}

	static function main() {
		var host = new Host("127.0.0.1");
		var server = new Socket();
		server.bind(host, 0);
		var port = server.host().port;

		var beforeListen = new Socket();
		beforeListen.setTimeout(0.02);
		var connectedBeforeListen = true;
		try {
			beforeListen.connect(host, port);
		} catch (_:Dynamic) {
			connectedBeforeListen = false;
		}
		Sys.println("connectedBeforeListen=" + connectedBeforeListen);
		safeClose(beforeListen);
		if (connectedBeforeListen) {
			safeClose(server);
			return;
		}

		server.listen(1);
		server.listen(2);
		var client = new Socket();
		var accepted:Socket = null;
		var failed:Dynamic = null;
		try {
			client.connect(host, port);
			accepted = server.accept();
			client.output.writeString("ping\n");
			client.output.flush();
			var request = accepted.input.readLine();
			accepted.output.writeString("pong:" + request + "\n");
			accepted.output.flush();
			Sys.println("roundTrip=" + client.input.readLine());
		} catch (error:Dynamic) {
			failed = error;
		}
		safeClose(accepted);
		safeClose(client);
		safeClose(server);
		if (failed != null)
			throw failed;
	}
}
