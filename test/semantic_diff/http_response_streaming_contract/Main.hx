import haxe.io.Bytes;
import haxe.io.BytesOutput;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private class StreamingOutput extends BytesOutput {
	final marker:String;
	final events:Array<String>;

	public function new(marker:String, events:Array<String>) {
		super();
		this.marker = marker;
		this.events = events;
	}

	function phase():String {
		return sys.FileSystem.exists(marker) ? "after" : "before";
	}

	override public function prepare(size:Int):Void {
		events.push("prepare:" + size + ":" + phase());
		super.prepare(size);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		events.push("write:" + bytes.sub(pos, len).toString() + ":" + phase());
		return super.writeBytes(bytes, pos, len);
	}

	override public function close():Void {
		events.push("close:" + phase());
		super.close();
	}
}

class Main {
	static function pythonServerScript():String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "import sys\n"
			+ "import time\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def do_GET(self):\n"
			+ "        self.send_response(201)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Content-Length', '11')\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.flush()\n"
			+ "        self.wfile.write(b'hello')\n"
			+ "        self.wfile.flush()\n"
			+ "        time.sleep(0.8)\n"
			+ "        open(sys.argv[1], 'w').close()\n"
			+ "        self.wfile.write(b' world')\n"
			+ "        self.wfile.flush()\n"
			+ "    def log_message(self, fmt, *args):\n"
			+ "        return\n"
			+ "\n"
			+ "with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n"
			+ "    print(httpd.server_address[1], flush=True)\n"
			+ "    httpd.handle_request()\n";
	}

	static function startServer(marker:String):ServerHandle {
		var process = new sys.io.Process("python3", ["-u", "-c", pythonServerScript(), marker]);
		var port = Std.parseInt(process.stdout.readLine());
		if (port == null) {
			process.close();
			throw "failed to read server port";
		}
		return {process: process, port: port};
	}

	static function closeServer(server:ServerHandle):Void {
		try {
			server.process.close();
		} catch (_:Dynamic) {}
	}

	static function main() {
		var marker = "/tmp/reflaxe_go_http_streaming_released";
		if (sys.FileSystem.exists(marker))
			sys.FileSystem.deleteFile(marker);
		var server = startServer(marker);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/stream");
		var events = new Array<String>();
		var output = new StreamingOutput(marker, events);
		var error = "";
		request.onStatus = function(status) events.push("status:" + status + ":" + (sys.FileSystem.exists(marker) ? "after" : "before"));
		request.onError = function(message) error = message;
		request.customRequest(false, output);
		closeServer(server);

		Sys.println("events=" + events.join(">"));
		Sys.println("body=" + output.getBytes().toString());
		Sys.println("error=" + error);
		if (sys.FileSystem.exists(marker))
			sys.FileSystem.deleteFile(marker);
	}
}
