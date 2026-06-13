import app.core.IncidentApi;
import app.core.IncidentConfig;
import app.core.IncidentStore;
import app.http.TinyHttpServer;
import sys.FileSystem;
import sys.io.File;
import sys.net.Host;
import sys.net.Socket;

class Harness {
	static inline final CONFIG_FILE = ".incident_api_scripted_config.json";
	static inline final STATE_FILE = ".incident_api_scripted_state.json";

	static function cleanup():Void {
		for (path in [CONFIG_FILE, STATE_FILE]) {
			try {
				if (FileSystem.exists(path)) {
					FileSystem.deleteFile(path);
				}
			} catch (_:haxe.Exception) {}
		}
	}

	static function request(server:TinyHttpServer, method:String, path:String, body:String):String {
		var client = new Socket();
		var result = "";
		try {
			client.connect(new Host(server.host), server.port);
			var contentLength = haxe.io.Bytes.ofString(body).length;
			client.output.writeString(method + " " + path + " HTTP/1.1\r\n");
			client.output.writeString("Host: " + server.host + ":" + server.port + "\r\n");
			client.output.writeString("Content-Type: application/json\r\n");
			client.output.writeString("Content-Length: " + contentLength + "\r\n");
			client.output.writeString("Connection: close\r\n\r\n");
			client.output.writeString(body);
			client.output.flush();
			server.serveOnce();
			result = summarize(client.input.readAll().toString());
		} catch (error:haxe.Exception) {
			result = "HTTP/1.1 000 Client Error body={\"error\":\"client_error\"}";
		}
		try {
			client.close();
		} catch (_:haxe.Exception) {}
		return result;
	}

	static function summarize(raw:String):String {
		var normalized = StringTools.replace(raw, "\r\n", "\n");
		var sections = normalized.split("\n\n");
		var headerLines = sections.length > 0 ? sections[0].split("\n") : [];
		var status = headerLines.length > 0 ? headerLines[0] : "HTTP/1.1 000 Missing";
		var body = sections.length > 1 ? sections[1] : "";
		return status + " body=" + body;
	}

	public static function run():String {
		cleanup();
		IncidentConfig.saveExample(CONFIG_FILE);
		var rawConfig = File.getContent(CONFIG_FILE);
		var config = IncidentConfig.load(CONFIG_FILE);
		config.statePath = STATE_FILE;
		var store = new IncidentStore(config.statePath);
		var api = new IncidentApi(config, store);
		var server = new TinyHttpServer(api, config.host, config.port);
		var out = new Array<String>();
		try {
			out.push("config=" + StringTools.trim(rawConfig));
			out.push("listen=" + server.host + ":<ephemeral>");
			out.push("health=" + request(server, "GET", "/health", ""));
			out.push("create=" + request(server, "POST", "/incidents", "{\"title\":\"Database lag\",\"severity\":\"high\"}"));
			out.push("list=" + request(server, "GET", "/incidents", ""));
			out.push("ack=" + request(server, "POST", "/incidents/1/ack", ""));
			out.push("resolve=" + request(server, "POST", "/incidents/1/resolve", ""));
			out.push("metrics=" + request(server, "GET", "/metrics", ""));
			out.push("state=" + StringTools.trim(File.getContent(STATE_FILE)));
		} catch (error:haxe.Exception) {
			out.push("error=" + Std.string(error));
		}
		server.close();
		cleanup();
		return out.join("\n");
	}
}
