private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

class Main {
	static function pythonServerScript():String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def _reply(self):\n"
			+ "        length = int(self.headers.get('Content-Length', '0'))\n"
			+ "        body = self.rfile.read(length).decode('utf-8') if length else ''\n"
			+ "        repeated = ','.join(self.headers.get_all('X-Repeat') or [])\n"
			+ "        summary = self.command + ' ' + self.path + '|repeat=' + repeated + '|body=' + body\n"
			+ "        payload = summary.encode('utf-8')\n"
			+ "        self.send_response(200)\n"
			+ "        self.send_header('Content-Length', str(len(payload)))\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.write(payload)\n"
			+ "    def do_GET(self):\n"
			+ "        self._reply()\n"
			+ "    def do_POST(self):\n"
			+ "        self._reply()\n"
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

	static function runGet():Void {
		var server = startServer();
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/get?base=one%20two&base=two%2Bthree");
		request.addParameter("field", "first");
		request.addParameter("field", "second");
		request.setParameter("field", "replaced");
		request.setParameter("other", "initial");
		request.addParameter("other", "x+y");
		request.setHeader("X-Repeat", "first");
		request.addHeader("X-Repeat", "second");
		request.setHeader("X-Repeat", "replaced");
		var result = "";
		request.onData = function(value) result = value;
		request.request(false);
		closeServer(server);
		Sys.println(result);
	}

	static function runPost():Void {
		var server = startServer();
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/post?base=from%20url&base=second");
		request.addParameter("field", "first value");
		request.addParameter("field", "second");
		request.setParameter("field", "replaced value");
		request.setParameter("other", "initial");
		request.addParameter("other", "x+y");
		request.setHeader("X-Repeat", "one");
		request.addHeader("X-Repeat", "two");
		var result = "";
		request.onData = function(value) result = value;
		request.request(true);
		closeServer(server);
		Sys.println(result);
	}

	static function main() {
		runGet();
		runPost();
	}
}
