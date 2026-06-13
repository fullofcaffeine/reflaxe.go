package app.core;

import app.http.HttpRequest;
import app.http.HttpResponse;
import haxe.Json;

/**
	What: Pure Haxe request router for the incident service.
	Why: This is the code users should copy: business logic isolated from the
	target-specific socket carrier.
	How: It accepts a tiny request value and returns deterministic JSON responses.
**/
class IncidentApi {
	final config:IncidentConfig;
	final store:IncidentStore;
	var requests:Int;

	public function new(config:IncidentConfig, store:IncidentStore) {
		this.config = config;
		this.store = store;
		this.requests = 0;
	}

	public function handle(request:HttpRequest):HttpResponse {
		requests++;
		if (request.method == "GET" && request.path == "/health") {
			return HttpResponse.json(200, "{\"ok\":true,\"service\":\"" + Incident.jsonEscape(config.serviceName) + "\"}");
		}
		if (request.method == "GET" && request.path == "/incidents") {
			return HttpResponse.json(200, "{\"incidents\":" + store.listJson() + "}");
		}
		if (request.method == "POST" && request.path == "/incidents") {
			return createIncident(request.body);
		}
		if (request.method == "POST"
			&& StringTools.startsWith(request.path, "/incidents/")
			&& StringTools.endsWith(request.path, "/ack")) {
			return updateIncident(request.path, "ack");
		}
		if (request.method == "POST"
			&& StringTools.startsWith(request.path, "/incidents/")
			&& StringTools.endsWith(request.path, "/resolve")) {
			return updateIncident(request.path, "resolve");
		}
		if (request.method == "GET" && request.path == "/metrics") {
			return HttpResponse.json(200, store.metricsJson(config.serviceName, requests));
		}
		return HttpResponse.json(404, "{\"error\":\"not_found\"}");
	}

	function createIncident(body:String):HttpResponse {
		var response = HttpResponse.json(400, "{\"error\":\"invalid_json\"}");
		try {
			var raw = parseJsonBody(body);
			var title = fieldString(raw, "title", "");
			if (title == "") {
				throw new IncidentRequestException("missing_title");
			}
			var severity = fieldString(raw, "severity", "low");
			var incident = store.create(title, severity);
			response = HttpResponse.json(201, "{\"incident\":" + incident.toJson() + "}");
		} catch (error:IncidentRequestException) {
			response = HttpResponse.json(400, "{\"error\":\"" + error.code + "\"}");
		}
		return response;
	}

	static function parseJsonBody(body:String):Dynamic {
		try {
			// Json.parse returns Dynamic by design. Keep it localized at the HTTP boundary.
			return Json.parse(body == "" ? "{}" : body);
		} catch (_:haxe.Exception) {
			throw new IncidentRequestException("invalid_json");
		}
		return null;
	}

	function updateIncident(path:String, action:String):HttpResponse {
		var parts = path.split("/");
		if (parts.length < 4) {
			return HttpResponse.json(404, "{\"error\":\"not_found\"}");
		}
		var id = Std.parseInt(parts[2]);
		if (id == null) {
			return HttpResponse.json(400, "{\"error\":\"invalid_id\"}");
		}
		var incident = action == "ack" ? store.acknowledge(id) : store.resolve(id);
		if (incident == null) {
			return HttpResponse.json(404, "{\"error\":\"incident_not_found\"}");
		}
		return HttpResponse.json(200, "{\"incident\":" + incident.toJson() + "}");
	}

	static function fieldString(raw:Dynamic, name:String, fallback:String):String {
		if (raw == null || !Reflect.hasField(raw, name)) {
			return fallback;
		}
		var value = Reflect.field(raw, name);
		return value == null ? fallback : Std.string(value);
	}
}
