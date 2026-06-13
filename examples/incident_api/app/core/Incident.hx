package app.core;

/**
	What: Incident domain record used by the service example.
	Why: Keeps the service business object ordinary Haxe code instead of hiding it
	behind Go-native request handlers.
	How: Instances are stored in memory, persisted as JSON text, and rendered into
	deterministic JSON response snippets.
**/
class Incident {
	public var id:Int;
	public var title:String;
	public var severity:String;
	public var acknowledged:Bool;
	public var resolved:Bool;
	public var createdAt:String;

	public function new(id:Int, title:String, severity:String, acknowledged:Bool, resolved:Bool, createdAt:String) {
		this.id = id;
		this.title = title;
		this.severity = severity;
		this.acknowledged = acknowledged;
		this.resolved = resolved;
		this.createdAt = createdAt;
	}

	public function toJson():String {
		return "{\"id\":" + id + ",\"title\":\"" + jsonEscape(title) + "\"" + ",\"severity\":\"" + jsonEscape(severity) + "\"" + ",\"acknowledged\":"
			+ boolJson(acknowledged) + ",\"resolved\":" + boolJson(resolved) + ",\"createdAt\":\"" + jsonEscape(createdAt) + "\"}";
	}

	static function boolJson(value:Bool):String {
		return value ? "true" : "false";
	}

	public static function jsonEscape(value:String):String {
		var out = new StringBuf();
		var i = 0;
		while (i < value.length) {
			var code = value.charCodeAt(i);
			if (code == 34) {
				out.add('\\\"');
			} else if (code == 92) {
				out.add('\\\\');
			} else if (code == 10) {
				out.add('\\n');
			} else if (code == 13) {
				out.add('\\r');
			} else if (code == 9) {
				out.add('\\t');
			} else {
				out.addChar(code);
			}
			i++;
		}
		return out.toString();
	}
}
