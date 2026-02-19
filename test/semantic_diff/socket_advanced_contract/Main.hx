import sys.net.Host;
import sys.net.Socket;

private typedef BoundServer = {
	var server:Socket;
	var port:Int;
}

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

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
		var client:Socket = null;
		var peer:Socket = null;
		var failure:Dynamic = null;

		try {
			client = new Socket();
			client.connect(new Host("127.0.0.1"), bound.port);
			peer = bound.server.accept();
			client.custom = "client";

			safe("setTimeout.noThrow", function() {
				client.setTimeout(0.05);
				return "ok";
			});
			safe("setBlocking.noThrow", function() {
				client.setBlocking(false);
				client.setBlocking(true);
				return "ok";
			});
			safe("setFastSend.noThrow", function() {
				client.setFastSend(true);
				client.setFastSend(false);
				return "ok";
			});

			safe("select.write.ready", function() {
				var ready = Socket.select([], [client], [], 0.0);
				return ready.read.length + ":" + ready.write.length + ":" + ready.others.length;
			});
			safe("select.read.empty", function() {
				var ready = Socket.select([client], [], [], 0.0);
				return "read=" + ready.read.length;
			});

			peer.output.writeString("alpha\n");
			peer.output.flush();
			safe("waitForRead", function() {
				client.waitForRead();
				return client.input.readLine();
			});

			peer.output.writeString("beta\n");
			peer.output.flush();
			safe("select.read.ready.custom", function() {
				var ready = Socket.select([client], [], [], 1.0);
				var first = ready.read.length > 0 ? ready.read[0] : null;
				var custom = first == null ? "none" : Std.string(first.custom);
				var line = first == null ? "" : first.input.readLine();
				return ready.read.length + ":" + custom + ":" + line;
			});

			safe("shutdown.noThrow", function() {
				client.shutdown(false, true);
				return "ok";
			});
		} catch (error:Dynamic) {
			failure = error;
		}

		safeClose(peer);
		safeClose(client);
		safeClose(bound.server);
		if (failure != null) {
			throw failure;
		}
	}
}
