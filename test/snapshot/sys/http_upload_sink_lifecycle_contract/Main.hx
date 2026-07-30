import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Input;
import sys.thread.Tls;

private typedef ServerHandle = {
	var process:sys.io.Process;
	var port:Int;
}

private typedef ScenarioResult = {
	var status:Int;
	var error:String;
	var data:String;
	var reads:Int;
	var wrongThread:Bool;
	var afterReturnReads:Int;
}

private class TrackingInput extends Input {
	public var reads(default, null):Int = 0;
	public var wrongThread(default, null):Bool = false;
	public var afterReturnReads(default, null):Int = 0;
	public var requestReturned:Bool = false;

	final mode:String;
	final callerMarker:Tls<String>;
	final payload:Bytes;
	var position:Int = 0;

	public function new(mode:String, payload = "") {
		this.mode = mode;
		this.payload = Bytes.ofString(payload);
		callerMarker = new Tls<String>();
		callerMarker.value = "source-caller";
	}

	override public function readBytes(bytes:Bytes, targetPos:Int, requested:Int):Int {
		reads++;
		if (callerMarker.value != "source-caller")
			wrongThread = true;
		if (requestReturned)
			afterReturnReads++;

		switch (mode) {
			case "source-error":
				if (position > 0)
					throw "source-exploded";
				bytes.set(targetPos, 120);
				position++;
				return 1;
			case "early-eof":
				if (position > 0)
					throw new Eof();
				bytes.set(targetPos, 120);
				position++;
				return 1;
			case "zero":
				return 0;
			case "stream":
				bytes.fill(targetPos, requested, 97);
				position += requested;
				return requested;
			default:
				if (position >= payload.length)
					throw new Eof();
				var count = payload.length - position;
				if (count > 2)
					count = 2;
				if (count > requested)
					count = requested;
				bytes.blit(targetPos, payload, position, count);
				position += count;
				return count;
		}
	}
}

class Main {
	static function pythonServerScript(mode:String, count:Int):String {
		return "import http.server\n"
			+ "import socketserver\n"
			+ "import time\n"
			+ "MODE = "
			+ haxe.Json.stringify(mode)
			+ "\n"
			+ "COUNT = "
			+ count
			+ "\n"
			+ "class Handler(http.server.BaseHTTPRequestHandler):\n"
			+ "    def do_POST(self):\n"
			+ "        if MODE == 'early':\n"
			+ "            body = b''\n"
			+ "            self.send_response(413)\n"
			+ "            self.send_header('Content-Length', str(len(body)))\n"
			+ "            self.end_headers()\n"
			+ "            self.wfile.write(body)\n"
			+ "            self.wfile.flush()\n"
			+ "            self.close_connection = True\n"
			+ "            self.connection.settimeout(0.5)\n"
			+ "            remaining = int(self.headers.get('Content-Length', '0'))\n"
			+ "            try:\n"
			+ "                while remaining > 0:\n"
			+ "                    chunk = self.rfile.read(min(65536, remaining))\n"
			+ "                    if not chunk:\n"
			+ "                        break\n"
			+ "                    remaining -= len(chunk)\n"
			+ "            except (TimeoutError, ConnectionResetError):\n"
			+ "                pass\n"
			+ "            return\n"
			+ "        if MODE == 'timeout':\n"
			+ "            time.sleep(0.5)\n"
			+ "            return\n"
			+ "        if MODE == 'close':\n"
			+ "            self.connection.shutdown(2)\n"
			+ "            self.connection.close()\n"
			+ "            return\n"
			+ "        length = int(self.headers.get('Content-Length', '0'))\n"
			+ "        self.rfile.read(length)\n"
			+ "        body = b'ok'\n"
			+ "        try:\n"
			+ "            self.send_response(200)\n"
			+ "            self.send_header('Content-Length', str(len(body)))\n"
			+ "            self.end_headers()\n"
			+ "            self.wfile.write(body)\n"
			+ "        except (BrokenPipeError, ConnectionResetError):\n"
			+ "            pass\n"
			+ "    def log_message(self, fmt, *args):\n"
			+ "        return\n"
			+ "with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n"
			+ "    print(httpd.server_address[1], flush=True)\n"
			+ "    for _ in range(COUNT):\n"
			+ "        httpd.handle_request()\n";
	}

	static function startServer(mode:String, count = 1):ServerHandle {
		var process = new sys.io.Process("python3", ["-u", "-c", pythonServerScript(mode, count)]);
		var port = Std.parseInt(process.stdout.readLine());
		if (port == null) {
			process.close();
			throw "failed to read server port";
		}
		return {process: process, port: port};
	}

	static function run(server:ServerHandle, mode:String, declaredSize:Int, timeout = 2.0):ScenarioResult {
		var input = new TrackingInput(mode, mode == "payload" ? "payload" : "");
		var request = new haxe.Http("http://127.0.0.1:" + server.port + "/upload");
		request.cnxTimeout = timeout;
		request.fileTransfer("asset", "demo.bin", input, declaredSize, "application/octet-stream");
		var status = -1;
		var error = "";
		var data = "";
		request.onStatus = function(value) status = value;
		request.onError = function(value) error = value;
		request.onData = function(value) data = value;
		request.request(true);
		input.requestReturned = true;
		Sys.sleep(0.01);
		return {
			status: status,
			error: error,
			data: data,
			reads: input.reads,
			wrongThread: input.wrongThread,
			afterReturnReads: input.afterReturnReads
		};
	}

	static function finishServer(server:ServerHandle):Void {
		try {
			server.process.close();
		} catch (_:Dynamic) {}
	}

	static function main() {
		var successServer = startServer("consume");
		var success = run(successServer, "payload", 7);
		finishServer(successServer);
		Sys.println("success="
			+ (success.status == 200 && success.error == "" && success.data == "ok" && success.reads == 4 && !success.wrongThread
				&& success.afterReturnReads == 0));

		var sourceServer = startServer("consume");
		var source = run(sourceServer, "source-error", 7);
		finishServer(sourceServer);
		Sys.println("source="
			+ (source.status == -1 && source.error == "source-exploded" && source.reads == 2 && !source.wrongThread && source.afterReturnReads == 0));

		var eofServer = startServer("consume");
		var eof = run(eofServer, "early-eof", 7);
		finishServer(eofServer);
		Sys.println("earlyEof=" + (eof.status == -1 && eof.error == "Transfer aborted" && eof.reads == 2 && !eof.wrongThread && eof.afterReturnReads == 0));

		var zeroServer = startServer("consume");
		var zero = run(zeroServer, "zero", 7);
		finishServer(zeroServer);
		Sys.println("zero="
			+ (zero.status == -1
				&& zero.error == "multipart upload made no progress"
				&& zero.reads == 1
				&& !zero.wrongThread
				&& zero.afterReturnReads == 0));

		var closeServer = startServer("close");
		var closed = run(closeServer, "stream", 8 * 1024 * 1024);
		finishServer(closeServer);
		Sys.println("serverClose=" + (closed.error != "" && !closed.wrongThread && closed.afterReturnReads == 0));

		var timeoutServer = startServer("timeout");
		var timed = run(timeoutServer, "stream", 16 * 1024 * 1024, 0.1);
		finishServer(timeoutServer);
		Sys.println("timeout=" + (timed.error != "" && !timed.wrongThread && timed.afterReturnReads == 0));

		var repeatedServer = startServer("early", 12);
		var repeatedStatus = 0;
		var repeatedError = 0;
		var repeatedThread = 0;
		var repeatedAfterReturn = 0;
		for (_ in 0...12) {
			var early = run(repeatedServer, "stream", 8 * 1024 * 1024);
			if (early.status == 413)
				repeatedStatus++;
			if (early.error == "Http Error #413")
				repeatedError++;
			if (!early.wrongThread)
				repeatedThread++;
			if (early.afterReturnReads == 0)
				repeatedAfterReturn++;
		}
		finishServer(repeatedServer);
		Sys.println("earlyRepeated=status:"
			+ repeatedStatus
			+ ";error:"
			+ repeatedError
			+ ";thread:"
			+ repeatedThread
			+ ";after:"
			+ repeatedAfterReturn);
	}
}
