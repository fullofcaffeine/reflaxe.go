import haxe.io.BytesOutput;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private class TrackingOutput extends BytesOutput {
	public var wasClosed(default, null):Bool = false;

	public function new() {
		super();
	}

	override public function close():Void {
		wasClosed = true;
		super.close();
	}
}

class Main {
	static function pythonServerScript():String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def do_GET(self):\n"
			+ "        status = 101 if self.path == '/switching' else (404 if self.path == '/missing' else 200)\n"
			+ "        payload = b'' if status == 101 else ('body:' + self.path).encode('utf-8')\n"
			+ "        self.send_response(status)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Content-Length', str(len(payload)))\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.write(payload)\n"
			+ "    def log_message(self, fmt, *args):\n"
			+ "        return\n"
			+ "\n"
			+ "with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n"
			+ "    print(httpd.server_address[1], flush=True)\n"
			+ "    httpd.handle_request()\n";
	}

	static function startServer():ServerHandle {
		var process = new sys.io.Process("python3", ["-u", "-c", pythonServerScript()]);
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

	static function run(path:String):Void {
		var server = startServer();
		var request = new haxe.Http("http://127.0.0.1:" + server.port + path);
		var output = new TrackingOutput();
		var status = -1;
		var error = "";
		var events = new Array<String>();
		request.onStatus = function(value) {
			status = value;
			events.push("status");
		};
		request.onError = function(value) {
			error = value;
			events.push("error");
		};
		request.onData = function(_) events.push("data");
		request.onBytes = function(_) events.push("bytes");
		request.customRequest(false, output);
		closeServer(server);

		Sys.println(path + ".body=" + output.getBytes().toString());
		Sys.println(path + ".closed=" + output.wasClosed);
		Sys.println(path + ".status=" + status);
		Sys.println(path + ".error=" + error);
		Sys.println(path + ".events=" + events.join(">"));
		Sys.println(path + ".response=" + request.responseData);
	}

	static function main() {
		run("/ok");
		run("/missing");
		run("/switching");
	}
}
