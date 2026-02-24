package app.core;

class PulseAlert {
	public final eventId:Int;
	public final source:String;
	public final region:String;
	public final severity:Int;
	public final weightedValue:Int;
	public final reason:String;

	public function new(eventId:Int, source:String, region:String, severity:Int, weightedValue:Int, reason:String) {
		this.eventId = eventId;
		this.source = source;
		this.region = region;
		this.severity = severity;
		this.weightedValue = weightedValue;
		this.reason = reason;
	}

	public static function fromEnriched(entry:PulseEnrichedEvent):PulseAlert {
		var label = entry.severity >= 3 ? "critical" : "warning";
		return new PulseAlert(entry.event.id, entry.event.source, entry.event.region, entry.severity, entry.weightedValue, label);
	}
}
