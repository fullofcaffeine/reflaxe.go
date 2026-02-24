package app.core;

class PulseEnrichedEvent {
	public final event:PulseEvent;
	public final severity:Int;
	public final weightedValue:Int;

	public function new(event:PulseEvent, severity:Int, weightedValue:Int) {
		this.event = event;
		this.severity = severity;
		this.weightedValue = weightedValue;
	}

	public inline function shouldAlert(weightedThreshold:Int):Bool {
		return weightedValue >= weightedThreshold;
	}
}
