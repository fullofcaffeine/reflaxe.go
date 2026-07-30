import haxe.io.Error;
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

	static function isBlocked(operation:() -> Void):Bool {
		try {
			operation();
			return false;
		} catch (error:Error) {
			return switch (error) {
				case Blocked: true;
				default: throw error;
			};
		}
		return false;
	}

	static function main() {
		var host = new Host("127.0.0.1");
		var server = new Socket();
		var client = new Socket();
		var accepted:Socket = null;
		var unconnected = new Socket();
		var failed:Dynamic = null;

		try {
			server.bind(host, 0);
			server.listen(1);
			client.connect(host, server.host().port);
			accepted = server.accept();
			client.custom = "client";

			var idle = Socket.select([client], [], [unconnected], 0.0);
			Sys.println("idle=" + idle.read.length + ":" + idle.others.length);

			client.setBlocking(false);
			Sys.println("readBlocked=" + isBlocked(function() client.input.readByte()));
			client.output.writeByte("n".code);
			Sys.println("nonblockingWrite=" + accepted.input.readByte());
			client.setBlocking(true);

			server.setBlocking(false);
			Sys.println("acceptBlocked=" + isBlocked(function() server.accept()));
			server.setBlocking(true);

			accepted.output.writeString("xy");
			accepted.output.flush();
			var duplicate = Socket.select([client, client], [], [], 1.0);
			Sys.println("duplicates=" + duplicate.read.length + ":" + (duplicate.read[0] == client) + ":" + duplicate.read[1].custom);
			Sys.println("first=" + client.input.readByte());
			var buffered = Socket.select([client], [], [], 0.0);
			Sys.println("buffered=" + buffered.read.length + ":" + client.input.readByte());

			var writable = Socket.select([], [client], [], 0.0);
			Sys.println("writable=" + writable.write.length);
		} catch (error:Dynamic) {
			failed = error;
		}

		safeClose(unconnected);
		safeClose(accepted);
		safeClose(client);
		safeClose(server);
		if (failed != null)
			throw failed;
	}
}
