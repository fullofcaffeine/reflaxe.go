package app.core;

class PulseCodec {
	public static function parse(frame:PulseIngressFrame):PulseEvent {
		var source = normalizeToken(frame.source, "unknown");
		var region = normalizeToken(frame.region, "global");
		var value = frame.value;
		if (value < 0) {
			value = 0;
		}
		return new PulseEvent(frame.sequence, source, region, value);
	}

	public static function enrich(event:PulseEvent):PulseEnrichedEvent {
		var severity = severityFor(event.value);
		var weightedValue = event.value * severity + regionBoost(event.region);
		return new PulseEnrichedEvent(event, severity, weightedValue);
	}

	public static function severityFor(value:Int):Int {
		if (value >= 12) {
			return 3;
		}
		if (value >= 8) {
			return 2;
		}
		return 1;
	}

	public static function regionBoost(region:String):Int {
		return switch (region) {
			case "fra":
				3;
			case "iad", "gru":
				2;
			case "sfo":
				1;
			case _:
				0;
		};
	}

	static function normalizeToken(value:String, fallback:String):String {
		var trimmed = StringTools.trim(value);
		if (trimmed == "") {
			return fallback;
		}
		return trimmed;
	}
}
