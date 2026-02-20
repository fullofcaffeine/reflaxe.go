private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Dynamic;
};

class Main {
	static function pythonServerScript():String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "import sys\n"
			+ "\n"
			+ "log_path = sys.argv[1]\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def _reply(self):\n"
			+ "        line = self.command + ':' + self.path\n"
			+ "        with open(log_path, 'w', encoding='utf-8') as fp:\n"
			+ "            fp.write(line)\n"
			+ "        body = line.encode('utf-8')\n"
			+ "        self.send_response(200)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Content-Length', str(len(body)))\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.write(body)\n"
			+ "    def do_GET(self):\n"
			+ "        self._reply()\n"
			+ "    def do_PATCH(self):\n"
			+ "        self._reply()\n"
			+ "    def log_message(self, fmt, *args):\n"
			+ "        return\n"
			+ "\n"
			+ "with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n"
			+ "    print(httpd.server_address[1], flush=True)\n"
			+ "    httpd.handle_request()\n";
	}

	static function startServer(logPath:String):ServerHandle {
		var script = pythonServerScript();
		var proc = new sys.io.Process("python3", ["-u", "-c", script, logPath]);
		var line = proc.stdout.readLine();
		var port:Dynamic = haxe.Json.parse(line);
		if (port == null) {
			proc.close();
			throw "failed to read server port";
		}
		return {process: proc, port: port};
	}

	static function closeServer(server:ServerHandle):Void {
		try {
			server.process.close();
		} catch (_:Dynamic) {}
	}

	static function readLog(path:String):String {
		return StringTools.trim(sys.io.File.getContent(path));
	}

	static function main() {
		var directLog = "/tmp/reflaxe_go_http_direct.log";
		var directServer = startServer(directLog);
		var direct = haxe.Http.requestUrl("http://127.0.0.1:" + directServer.port + "/direct");
		closeServer(directServer);
		Sys.println("direct=" + direct);
		Sys.println("directTrace=" + readLog(directLog));

		var customLog = "/tmp/reflaxe_go_http_custom.log";
		var customServer = startServer(customLog);
		var req = new haxe.Http("http://127.0.0.1:" + customServer.port + "/method");
		var sink = new haxe.io.BytesBuffer();
		req.customRequest(false, cast sink, null, "PATCH");
		closeServer(customServer);
		Sys.println("customTrace=" + readLog(customLog));

		var socketLog = "/tmp/reflaxe_go_http_socket.log";
		var socketServer = startServer(socketLog);
		var socketReq = new haxe.Http("http://127.0.0.1:" + socketServer.port + "/socket");
		var socketSink = new haxe.io.BytesBuffer();
		socketReq.customRequest(false, cast socketSink, new sys.net.Socket(), "PATCH");
		closeServer(socketServer);
		Sys.println("socketTrace=" + readLog(socketLog));

		var proxyLog = "/tmp/reflaxe_go_http_proxy.log";
		var proxyServer = startServer(proxyLog);
		haxe.Http.PROXY = cast {
			host: "127.0.0.1",
			port: proxyServer.port,
			auth: null
		};
		var proxied = haxe.Http.requestUrl("http://example.invalid/proxy-check");
		haxe.Http.PROXY = null;
		closeServer(proxyServer);
		Sys.println("proxy=" + proxied);
		Sys.println("proxyTrace=" + readLog(proxyLog));
	}
}
