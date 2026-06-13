package app.core;

import haxe.Json;
import sys.FileSystem;
import sys.io.File;

/**
	What: File-backed incident repository for the service example.
	Why: Shows real persistence through `sys.io.File` while keeping state logic
	portable Haxe code.
	How: Loads/saves one deterministic JSON document and exposes typed operations
	for the HTTP layer.
**/
class IncidentStore {
	final statePath:String;
	var incidents:Array<Incident>;
	var nextId:Int;

	public function new(statePath:String) {
		this.statePath = statePath;
		this.incidents = [];
		this.nextId = 1;
		load();
	}

	public function create(title:String, severity:String):Incident {
		var incident = new Incident(nextId, title, normalizeSeverity(severity), false, false, "2026-06-12T00:00:00Z");
		nextId++;
		incidents.push(incident);
		save();
		return incident;
	}

	public function acknowledge(id:Int):Null<Incident> {
		var incident = find(id);
		if (incident == null) {
			return null;
		}
		incident.acknowledged = true;
		save();
		return incident;
	}

	public function resolve(id:Int):Null<Incident> {
		var incident = find(id);
		if (incident == null) {
			return null;
		}
		incident.acknowledged = true;
		incident.resolved = true;
		save();
		return incident;
	}

	public function listJson():String {
		var out = new StringBuf();
		out.add("[");
		var i = 0;
		while (i < incidents.length) {
			if (i > 0) {
				out.add(",");
			}
			out.add(incidents[i].toJson());
			i++;
		}
		out.add("]");
		return out.toString();
	}

	public function metricsJson(serviceName:String, requests:Int):String {
		var open = 0;
		var acked = 0;
		var resolved = 0;
		for (incident in incidents) {
			if (incident.resolved) {
				resolved++;
			} else {
				open++;
			}
			if (incident.acknowledged) {
				acked++;
			}
		}
		return "{\"service\":\"" + Incident.jsonEscape(serviceName) + "\",\"requests\":" + requests + ",\"open\":" + open + ",\"acknowledged\":" + acked
			+ ",\"resolved\":" + resolved + "}";
	}

	function find(id:Int):Null<Incident> {
		for (incident in incidents) {
			if (incident.id == id) {
				return incident;
			}
		}
		return null;
	}

	function load():Void {
		if (!FileSystem.exists(statePath)) {
			return;
		}
		var content = StringTools.trim(File.getContent(statePath));
		if (content == "") {
			return;
		}

		// Json.parse returns Dynamic by design. Keep it localized at the file boundary.
		var raw:Dynamic = Json.parse(content);
		nextId = intField(raw, "nextId", 1);
		var loaded = new Array<Incident>();
		if (raw != null && Reflect.hasField(raw, "incidents")) {
			var values:Array<Dynamic> = cast Reflect.field(raw, "incidents");
			for (value in values) {
				loaded.push(new Incident(intField(value, "id", nextId), stringField(value, "title", "untitled"),
					normalizeSeverity(stringField(value, "severity", "low")), boolField(value, "acknowledged", false), boolField(value, "resolved", false),
					stringField(value, "createdAt", "2026-06-12T00:00:00Z")));
			}
		}
		incidents = loaded;
	}

	function save():Void {
		File.saveContent(statePath, "{\"nextId\":" + nextId + ",\"incidents\":" + listJson() + "}\n");
	}

	static function normalizeSeverity(raw:String):String {
		var value = raw.toLowerCase();
		if (value == "critical" || value == "high" || value == "medium" || value == "low") {
			return value;
		}
		return "low";
	}

	static function stringField(raw:Dynamic, name:String, fallback:String):String {
		if (raw == null || !Reflect.hasField(raw, name))
			return fallback;
		var value = Reflect.field(raw, name);
		return value == null ? fallback : Std.string(value);
	}

	static function intField(raw:Dynamic, name:String, fallback:Int):Int {
		if (raw == null || !Reflect.hasField(raw, name))
			return fallback;
		var parsed = Std.parseInt(Std.string(Reflect.field(raw, name)));
		return parsed == null ? fallback : parsed;
	}

	static function boolField(raw:Dynamic, name:String, fallback:Bool):Bool {
		if (raw == null || !Reflect.hasField(raw, name))
			return fallback;
		return Std.string(Reflect.field(raw, name)) == "true";
	}
}
