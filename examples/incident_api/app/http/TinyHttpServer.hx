package app.http;

import app.core.IncidentApi;
import haxe.io.Bytes;
import haxe.io.BytesBuffer;
import sys.net.Host;
import sys.net.Socket;

/**
	What: Tiny HTTP/1.1 server built only on Haxe `sys.net.Socket`.
	Why: The example must show off the compiler and Haxe stdlib support, not Go
	`net/http`, `go.*`, externs, or raw injection.
	How: Handles one request per connection with Content-Length bodies and
	Connection: close responses.
**/
class TinyHttpServer {
	final api:IncidentApi;
	final server:Socket;

	public final host:String;
	public final port:Int;

	public function new(api:IncidentApi, host:String, port:Int) {
		this.api = api;
		this.host = host;
		this.server = new Socket();
		this.server.bind(new Host(host), port);
		this.server.listen(16);
		var bound = this.server.host();
		this.port = bound == null ? port : bound.port;
	}

	public function serveOnce():Void {
		var peer:Socket = null;
		try {
			peer = server.accept();
			var request = readRequest(peer);
			var response = api.handle(request);
			writeResponse(peer, response);
		} catch (error:haxe.Exception) {
			if (peer != null) {
				writeResponse(peer, HttpResponse.json(500, "{\"error\":\"server_error\"}"));
			}
		}
		closePeer(peer);
	}

	public function close():Void {
		try {
			server.close();
		} catch (_:haxe.Exception) {}
	}

	static function closePeer(peer:Socket):Void {
		if (peer == null) {
			return;
		}
		try {
			peer.close();
		} catch (_:haxe.Exception) {}
	}

	function readRequest(peer:Socket):HttpRequest {
		var line = peer.input.readLine();
		var first = line.split(" ");
		var method = first.length > 0 ? first[0] : "GET";
		var path = first.length > 1 ? first[1] : "/";
		var contentLength = 0;
		while (true) {
			var header = peer.input.readLine();
			if (header == "") {
				break;
			}
			var lower = header.toLowerCase();
			if (StringTools.startsWith(lower, "content-length:")) {
				var rawLength = StringTools.trim(header.substr("content-length:".length));
				var parsed = Std.parseInt(rawLength);
				contentLength = parsed == null ? 0 : parsed;
			}
		}
		return new HttpRequest(method, path, readBody(peer, contentLength));
	}

	function readBody(peer:Socket, length:Int):String {
		if (length <= 0) {
			return "";
		}
		var out = new BytesBuffer();
		var i = 0;
		while (i < length) {
			out.addByte(peer.input.readByte());
			i++;
		}
		return out.getBytes().toString();
	}

	function writeResponse(peer:Socket, response:HttpResponse):Void {
		var bodyBytes = Bytes.ofString(response.body);
		var head = "HTTP/1.1 " + response.status + " " + reason(response.status) + "\r\n" + "Content-Type: application/json\r\n" + "Content-Length: "
			+ bodyBytes.length + "\r\n" + "Connection: close\r\n" + "\r\n";
		peer.output.writeString(head + response.body);
		peer.output.flush();
	}

	static function reason(status:Int):String {
		return switch (status) {
			case 200: "OK";
			case 201: "Created";
			case 400: "Bad Request";
			case 404: "Not Found";
			case 500: "Internal Server Error";
			case _: "OK";
		}
	}
}
