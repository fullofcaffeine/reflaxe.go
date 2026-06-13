package app.core;

import haxe.Json;
import sys.FileSystem;
import sys.io.File;

/**
	What: Runtime configuration for the incident API example.
	Why: A real service example should show file-backed configuration through Haxe
	stdlib APIs, not hard-coded Go values.
	How: `load` reads a small JSON object and falls back to deterministic defaults
	when a field is absent.
**/
class IncidentConfig {
	public var serviceName:String;
	public var host:String;
	public var port:Int;
	public var statePath:String;

	public function new(serviceName:String, host:String, port:Int, statePath:String) {
		this.serviceName = serviceName;
		this.host = host;
		this.port = port;
		this.statePath = statePath;
	}

	public static function defaults():IncidentConfig {
		return new IncidentConfig("incident-api", "127.0.0.1", 0, ".incident_api_state.json");
	}

	public static function load(path:String):IncidentConfig {
		var config = defaults();
		if (!FileSystem.exists(path)) {
			return config;
		}

		// Json.parse returns Dynamic by design. Keep it localized at the config boundary.
		var raw:Dynamic = Json.parse(File.getContent(path));
		config.serviceName = stringField(raw, "serviceName", config.serviceName);
		config.host = stringField(raw, "host", config.host);
		config.port = intField(raw, "port", config.port);
		config.statePath = stringField(raw, "statePath", config.statePath);
		return config;
	}

	public static function saveExample(path:String):Void {
		File.saveContent(path, "{\"serviceName\":\"incident-api\",\"host\":\"127.0.0.1\",\"port\":0,\"statePath\":\".incident_api_state.json\"}\n");
	}

	static function stringField(raw:Dynamic, name:String, fallback:String):String {
		if (raw == null || !Reflect.hasField(raw, name)) {
			return fallback;
		}
		var value = Reflect.field(raw, name);
		if (value == null) {
			return fallback;
		}
		return Std.string(value);
	}

	static function intField(raw:Dynamic, name:String, fallback:Int):Int {
		if (raw == null || !Reflect.hasField(raw, name)) {
			return fallback;
		}
		var value = Reflect.field(raw, name);
		if (value == null) {
			return fallback;
		}
		var parsed = Std.parseInt(Std.string(value));
		return parsed == null ? fallback : parsed;
	}
}
