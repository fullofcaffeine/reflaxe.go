import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Input;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private class PartialInput extends Input {
	public var readCalls(default, null):Int = 0;

	final payload:Bytes;
	final maxChunk:Int;
	var position:Int = 0;

	public function new(payload:String, maxChunk:Int) {
		this.payload = Bytes.ofString(payload);
		this.maxChunk = maxChunk;
	}

	override public function readBytes(bytes:Bytes, targetPos:Int, requested:Int):Int {
		if (position >= payload.length)
			throw new Eof();
		var available = payload.length - position;
		var count = requested < available ? requested : available;
		if (count > maxChunk)
			count = maxChunk;
		bytes.blit(targetPos, payload, position, count);
		position += count;
		readCalls++;
		return count;
	}
}

private class FailingInput extends Input {
	public var readCalls(default, null):Int = 0;

	public function new() {}

	override public function readBytes(bytes:Bytes, targetPos:Int, requested:Int):Int {
		readCalls++;
		throw "upload-failed";
	}
}

class Main {
	static function pythonServerScript():String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "from email import policy\n"
			+ "from email.parser import BytesParser\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def do_POST(self):\n"
			+ "        length = int(self.headers.get('Content-Length', '0'))\n"
			+ "        body = self.rfile.read(length)\n"
			+ "        if len(body) != length:\n"
			+ "            return\n"
			+ "        prefix = ('Content-Type: ' + self.headers['Content-Type'] + '\\r\\nMIME-Version: 1.0\\r\\n\\r\\n').encode('ascii')\n"
			+ "        message = BytesParser(policy=policy.default).parsebytes(prefix + body)\n"
			+ "        values = {}\n"
			+ "        for part in message.iter_parts():\n"
			+ "            name = part.get_param('name', header='content-disposition')\n"
			+ "            values[name] = (part.get_filename(), part.get_content_type(), part.get_payload(decode=True))\n"
			+ "        note = values['note'][2].decode('utf-8')\n"
			+ "        filename, mime, payload = values['asset']\n"
			+ "        summary = 'note=' + note + ';file=' + filename + ';mime=' + mime + ';body=' + payload.decode('utf-8')\n"
			+ "        encoded = summary.encode('utf-8')\n"
			+ "        self.send_response(200)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Content-Length', str(len(encoded)))\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.write(encoded)\n"
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

	static function main() {
		var server = startServer();
		var input = new PartialInput("payload", 2);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/upload");
		request.addParameter("note", "hello");
		request.fileTransfer("asset", "demo.txt", input, 7, "text/plain");
		var data = "";
		var status = -1;
		var error = "";
		var events = new Array<String>();
		request.onData = function(value) {
			data = value;
			events.push("data");
		};
		request.onBytes = function(_) events.push("bytes");
		request.onStatus = function(value) {
			status = value;
			events.push("status");
		};
		request.onError = function(value) {
			error = value;
			events.push("error");
		};
		request.request(true);
		closeServer(server);

		Sys.println(data);
		Sys.println("readCalls=" + input.readCalls);
		Sys.println("status=" + status);
		Sys.println("error=" + error);
		Sys.println("events=" + events.join(">"));

		var failureServer = startServer();
		var failureInput = new FailingInput();
		var failureRequest = new haxe.Http("http://127.0.0.1:" + failureServer.port + "/upload");
		failureRequest.cnxTimeout = 1;
		failureRequest.fileTransfer("asset", "broken.txt", failureInput, 7, "text/plain");
		var failureStatus = -1;
		var failureError = "";
		var failureEvents = new Array<String>();
		failureRequest.onStatus = function(value) {
			failureStatus = value;
			failureEvents.push("status");
		};
		failureRequest.onError = function(value) {
			failureError = value;
			failureEvents.push("error");
		};
		failureRequest.request(true);
		closeServer(failureServer);
		Sys.println("failureError=" + failureError);
		Sys.println("failureStatus=" + failureStatus);
		Sys.println("failureEvents=" + failureEvents.join(">"));
		Sys.println("failureReadCalls=" + failureInput.readCalls);
	}
}
