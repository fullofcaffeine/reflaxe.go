import haxe.io.Bytes;
import haxe.io.Eof;
import sys.net.Host;
import sys.net.Socket;
import sys.thread.Deque;
import sys.thread.Thread;

class Main {
	static function safeClose(socket:Socket):Void {
		if (socket == null)
			return;
		try {
			socket.close();
		} catch (_:Dynamic) {}
	}

	static function main() {
		var server = new Socket();
		server.bind(new Host("127.0.0.1"), 0);
		server.listen(1);
		var port = server.host().port;
		var serverResult = new Deque<String>();

		Thread.create(function() {
			var peer:Socket = null;
			var result = "closed";
			try {
				peer = server.accept();
				peer.output.writeString("xy");
				peer.output.flush();
			} catch (error:Dynamic) {
				result = "error:" + Std.string(error);
			}
			safeClose(peer);
			serverResult.add(result);
		});

		var client = new Socket();
		client.connect(new Host("127.0.0.1"), port);
		var bytes = Bytes.alloc(8);
		var count = client.input.readBytes(bytes, 0, bytes.length);
		var peerState = serverResult.pop(true);

		var reachedEof = false;
		var unexpected = "";
		try {
			client.input.readByte();
		} catch (_:Eof) {
			reachedEof = true;
		} catch (error:Dynamic) {
			unexpected = Std.string(error);
		}

		Sys.println("partial=" + bytes.getString(0, count) + ":" + count);
		Sys.println("server=" + peerState);
		Sys.println("eof=" + reachedEof);
		Sys.println("unexpected=" + unexpected);

		safeClose(client);
		safeClose(server);
	}
}
