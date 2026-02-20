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
			+ "        length = int(self.headers.get('Content-Length', '0'))\n"
			+ "        body = b''\n"
			+ "        if length > 0:\n"
			+ "            body = self.rfile.read(length)\n"
			+ "        line = self.command + ':' + self.path\n"
			+ "        if body:\n"
			+ "            line += '|body=' + body.decode('utf-8')\n"
			+ "        with open(log_path, 'w', encoding='utf-8') as fp:\n"
			+ "            fp.write(line)\n"
			+ "        status = 404 if self.path.startswith('/missing') else 200\n"
			+ "        payload = line.encode('utf-8')\n"
			+ "        self.send_response(status)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Set-Cookie', 'a=1')\n"
			+ "        self.send_header('Set-Cookie', 'b=2')\n"
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

	static function cookiesToString(values:Null<Array<String>>):String {
		if (values == null) {
			return "null";
		}
		var out = new StringBuf();
		var first = true;
		for (value in values) {
			if (!first) {
				out.add("|");
			}
			first = false;
			out.add(value);
		}
		return out.toString();
	}

	static function main() {
		var okLog = "/tmp/reflaxe_go_http_callbacks_ok.log";
		var okServer = startServer(okLog);
		var req = new haxe.Http("http://127.0.0.1:" + okServer.port + "/ok");
		var okData = "";
		var okBytes = -1;
		var okStatus = -1;
		var okErr = "";
		req.onData = function(data) {
			okData = data;
		};
		req.onBytes = function(bytes) {
			okBytes = bytes.length;
		};
		req.onStatus = function(status) {
			okStatus = status;
		};
		req.onError = function(msg) {
			okErr = msg;
		};
		req.request(false);
		closeServer(okServer);
		Sys.println("okTrace=" + readLog(okLog));
		Sys.println("okData=" + okData);
		Sys.println("okBytes=" + okBytes);
		Sys.println("okStatus=" + okStatus);
		Sys.println("okErr=" + okErr);
		Sys.println("okResp=" + req.responseData);
		Sys.println("okCookies=" + cookiesToString(req.getResponseHeaderValues("Set-Cookie")));
		Sys.println("okContentType=" + cookiesToString(req.getResponseHeaderValues("Content-Type")));

		var postLog = "/tmp/reflaxe_go_http_callbacks_post.log";
		var postServer = startServer(postLog);
		var postReq = new haxe.Http("http://127.0.0.1:" + postServer.port + "/post");
		postReq.setPostData("payload=abc");
		var postData = "";
		var postStatus = -1;
		postReq.onData = function(data) {
			postData = data;
		};
		postReq.onStatus = function(status) {
			postStatus = status;
		};
		postReq.request(true);
		closeServer(postServer);
		Sys.println("postTrace=" + readLog(postLog));
		Sys.println("postData=" + postData);
		Sys.println("postStatus=" + postStatus);

		var missLog = "/tmp/reflaxe_go_http_callbacks_missing.log";
		var missServer = startServer(missLog);
		var missReq = new haxe.Http("http://127.0.0.1:" + missServer.port + "/missing");
		var missStatus = -1;
		var missErr = "";
		var missData = "";
		missReq.onStatus = function(status) {
			missStatus = status;
		};
		missReq.onError = function(msg) {
			missErr = msg;
		};
		missReq.onData = function(data) {
			missData = data;
		};
		missReq.request(false);
		closeServer(missServer);
		Sys.println("missTrace=" + readLog(missLog));
		Sys.println("missStatus=" + missStatus);
		Sys.println("missErr=" + missErr);
		Sys.println("missData=" + missData);
		Sys.println("missResp=" + missReq.responseData);
	}
}
