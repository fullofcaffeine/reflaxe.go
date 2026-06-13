import app.core.IncidentApi;
import app.core.IncidentConfig;
import app.core.IncidentStore;
import app.http.TinyHttpServer;

class Main {
	static function argValue(name:String, fallback:String):String {
		var args = Sys.args();
		var i = 0;
		while (i < args.length - 1) {
			if (args[i] == name) {
				return args[i + 1];
			}
			i++;
		}
		return fallback;
	}

	static function hasArg(name:String):Bool {
		for (arg in Sys.args()) {
			if (arg == name)
				return true;
		}
		return false;
	}

	static function printHelp():Void {
		Sys.println("incident_api commands:");
		Sys.println("  --scripted");
		Sys.println("  init-config --config <path>");
		Sys.println("  serve --config <path>");
		Sys.println("curl examples:");
		Sys.println("  curl http://127.0.0.1:8080/health");
		Sys.println("  curl -X POST -d '{\"title\":\"Database lag\",\"severity\":\"high\"}' http://127.0.0.1:8080/incidents");
	}

	static function serve(configPath:String):Void {
		var config = IncidentConfig.load(configPath);
		var api = new IncidentApi(config, new IncidentStore(config.statePath));
		var server = new TinyHttpServer(api, config.host, config.port);
		Sys.println("incident_api listening on http://" + server.host + ":" + server.port);
		while (true) {
			server.serveOnce();
		}
	}

	static function main() {
		#if example_ci
		Sys.println(Harness.run());
		#else
		if (hasArg("--scripted")) {
			Sys.println(Harness.run());
			return;
		}
		if (hasArg("init-config")) {
			var configPath = argValue("--config", "config.json");
			IncidentConfig.saveExample(configPath);
			Sys.println("wrote " + configPath);
			return;
		}
		if (hasArg("serve")) {
			serve(argValue("--config", "config.json"));
			return;
		}
		printHelp();
		#end
	}
}
