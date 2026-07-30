import haxe.io.Bytes;
import haxe.io.BytesOutput;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private class TrackingOutput extends BytesOutput {
	public final events:Array<String> = [];
	public var closeCount(default, null):Int = 0;

	public function new() {
		super();
	}

	override public function prepare(size:Int):Void {
		events.push("prepare:" + size);
		super.prepare(size);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		events.push("write:" + bytes.sub(pos, len).toString());
		return super.writeBytes(bytes, pos, len);
	}

	override public function close():Void {
		closeCount++;
		events.push("close");
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
			+ "        payload = b'hello world'\n"
			+ "        self.send_response(200)\n"
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
		} catch (_:haxe.Exception) {}
	}

	static function run(url:String):TrackingOutput {
		var request = new haxe.Http(url);
		var output = new TrackingOutput();
		request.onStatus = function(status) output.events.push("status:" + status);
		request.onError = function(message) output.events.push("error:" + message);
		request.customRequest(false, output);
		return output;
	}

	static function runNetwork():TrackingOutput {
		var server = startServer();
		var output = run("http://127.0.0.1:" + server.port + "/body");
		closeServer(server);
		return output;
	}

	static function main() {
		#if reflaxe_go_test_target
		var first = run("data:text/plain,hello%20world");
		#else
		var first = runNetwork();
		#end
		var second = runNetwork();

		var firstEvents = first.events.join(">");
		var secondEvents = second.events.join(">");
		Sys.println("first.events=" + firstEvents);
		Sys.println("second.events=" + secondEvents);
		Sys.println("same.events=" + (firstEvents == secondEvents));
		Sys.println("first.body=" + first.getBytes().toString());
		Sys.println("second.body=" + second.getBytes().toString());
		Sys.println("same.close=" + (first.closeCount == second.closeCount));
	}
}
