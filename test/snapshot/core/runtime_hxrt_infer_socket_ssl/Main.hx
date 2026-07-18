import sys.ssl.Socket;

class Main {
	static function main() {
		var socket = new Socket();
		socket.verifyCert = false;
		socket.close();
		Sys.println("tls-socket-ready");
	}
}
