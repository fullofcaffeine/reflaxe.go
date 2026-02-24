package app.core;

class PulseSourceAggregate {
	public final source:String;
	public var count(default, null):Int;
	public var totalValue(default, null):Int;
	public var totalWeighted(default, null):Int;
	public var maxValue(default, null):Int;
	public var maxSeverity(default, null):Int;

	public function new(source:String) {
		this.source = source;
		count = 0;
		totalValue = 0;
		totalWeighted = 0;
		maxValue = 0;
		maxSeverity = 0;
	}

	public function record(entry:PulseEnrichedEvent):Void {
		count++;
		totalValue += entry.event.value;
		totalWeighted += entry.weightedValue;
		if (entry.event.value > maxValue) {
			maxValue = entry.event.value;
		}
		if (entry.severity > maxSeverity) {
			maxSeverity = entry.severity;
		}
	}

	public function summaryToken():String {
		return source + ":" + count + "/" + totalValue + "/" + totalWeighted + "/sev" + maxSeverity;
	}
}
