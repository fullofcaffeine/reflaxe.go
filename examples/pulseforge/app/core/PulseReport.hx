package app.core;

class PulseReport {
	final profile:String;
	final variant:String;
	final capability:String;
	final ingestCount:Int;
	final totalValue:Int;
	final alertCount:Int;
	final runtimeScore:Int;

	public function new(profile:String, variant:String, capability:String, ingestCount:Int, totalValue:Int, alertCount:Int, runtimeScore:Int) {
		this.profile = profile;
		this.variant = variant;
		this.capability = capability;
		this.ingestCount = ingestCount;
		this.totalValue = totalValue;
		this.alertCount = alertCount;
		this.runtimeScore = runtimeScore;
	}

	public function lines():Array<String> {
		return [
			"pulseforge.profile=" + profile,
			"pulseforge.variant=" + variant,
			"runtime.capability=" + capability,
			"ingest.events=" + ingestCount,
			"pipeline.total=" + totalValue,
			"pipeline.alerts=" + alertCount,
			"runtime.score=" + runtimeScore
		];
	}
}
