import haxe.io.Bytes;
import sys.net.Socket;

class Main {
	static function consume(socket:Socket):String {
		var line = socket.input.readLine();
		var lower = line.toLowerCase();
		var parsed = Std.parseInt("42");
		var first = socket.input.readByte();
		var buffer = Bytes.alloc(4);
		var read = socket.input.readBytes(buffer, 0, buffer.length);
		var rest = socket.input.readAll().toString();
		return lower + ":" + parsed + ":" + first + ":" + read + ":" + rest;
	}

	static function main() {
		Sys.println("socket-input-service-surface");
	}
}
