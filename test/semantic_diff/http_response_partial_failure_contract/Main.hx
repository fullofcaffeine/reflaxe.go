import haxe.io.Bytes;
import haxe.io.BytesOutput;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private class FailureOutput extends BytesOutput {
	public final events:Array<String>;

	final failure:String;

	public var closeCount(default, null):Int = 0;

	public function new(failure:String) {
		super();
		this.failure = failure;
		this.events = [];
	}

	override public function prepare(size:Int):Void {
		events.push("prepare:" + size);
		if (failure == "prepare")
			throw "prepare exploded";
		super.prepare(size);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		events.push("write:" + bytes.sub(pos, len).toString());
		if (failure == "write")
			throw "write exploded";
		return super.writeBytes(bytes, pos, len);
	}

	override public function close():Void {
		closeCount++;
		events.push("close");
		if (failure == "close")
			throw "close exploded";
		super.close();
	}
}

class Main {
	static function pythonServerScript(truncated:Bool):String {
		var declared = truncated ? 10 : 2;
		var payload = truncated ? "hello" : "ok";
		return "import http.server\n"
			+ "import socketserver\n"
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def do_GET(self):\n"
			+ "        self.send_response(200)\n"
			+ "        self.send_header('Content-Type', 'text/plain')\n"
			+ "        self.send_header('Content-Length', '"
			+ declared
			+ "')\n"
			+ "        self.end_headers()\n"
			+ "        self.wfile.write(b'"
			+ payload
			+ "')\n"
			+ "        self.wfile.flush()\n"
			+ "        self.close_connection = True\n"
			+ "    def log_message(self, fmt, *args):\n"
			+ "        return\n"
			+ "\n"
			+ "with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n"
			+ "    print(httpd.server_address[1], flush=True)\n"
			+ "    httpd.handle_request()\n";
	}

	static function startServer(truncated:Bool):ServerHandle {
		var process = new sys.io.Process("python3", ["-u", "-c", pythonServerScript(truncated)]);
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

	static function runCustomTruncated():Void {
		var server = startServer(true);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/truncated");
		var output = new FailureOutput("");
		request.onStatus = function(status) output.events.push("status:" + status);
		request.onError = function(message) output.events.push("error:" + message);
		request.customRequest(false, output);
		closeServer(server);
		Sys.println("custom.events=" + output.events.join(">"));
		Sys.println("custom.body=" + output.getBytes().toString());
		Sys.println("custom.closeCount=" + output.closeCount);
	}

	static function runRequestTruncated():Void {
		var server = startServer(true);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/truncated");
		var events = new Array<String>();
		request.onStatus = function(status) events.push("status:" + status);
		request.onError = function(message) events.push("error:" + message);
		request.request(false);
		closeServer(server);
		Sys.println("request.events=" + events.join(">"));
		Sys.println("request.bytes=" + (request.responseBytes == null ? "null" : request.responseBytes.toString()));
		Sys.println("request.data=" + request.responseData);
	}

	static function runOutputFailure(failure:String):Void {
		var server = startServer(false);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/" + failure);
		var output = new FailureOutput(failure);
		request.onStatus = function(status) output.events.push("status:" + status);
		request.onError = function(message) output.events.push("error:" + message);
		request.customRequest(false, output);
		closeServer(server);
		Sys.println(failure + ".events=" + output.events.join(">"));
		Sys.println(failure + ".closeCount=" + output.closeCount);
	}

	static function runStatusFailure():Void {
		var server = startServer(false);
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/status");
		var output = new FailureOutput("");
		request.onStatus = function(status) {
			output.events.push("status:" + status);
			throw "status exploded";
		};
		request.onError = function(message) output.events.push("error:" + message);
		request.customRequest(false, output);
		closeServer(server);
		Sys.println("status.events=" + output.events.join(">"));
		Sys.println("status.closeCount=" + output.closeCount);
	}

	static function runRequestUrlInvalid():Void {
		var result = "return";
		try {
			haxe.Http.requestUrl("not a valid URL");
		} catch (error:haxe.Exception) {
			result = "throw";
		}
		Sys.println("requestUrl.invalid=" + result);
	}

	static function runRequestUrlTruncated():Void {
		var server = startServer(true);
		var result = "returned";
		try {
			result = "return:" + haxe.Http.requestUrl("http://127.0.0.1:" + server.port + "/truncated");
		} catch (error:haxe.Exception) {
			result = "throw:" + error.message;
		}
		closeServer(server);
		Sys.println("requestUrl.truncated=" + result);
	}

	static function main() {
		runCustomTruncated();
		runRequestTruncated();
		runOutputFailure("prepare");
		runOutputFailure("write");
		runOutputFailure("close");
		runStatusFailure();
		runRequestUrlInvalid();
		runRequestUrlTruncated();
	}
}
