package app.core;

class PulseReport {
	final profile:String;
	final variant:String;
	final capability:String;
	final ingestReceived:Int;
	final ingestAccepted:Int;
	final backpressureEvents:Int;
	final parseCount:Int;
	final enrichCount:Int;
	final aggregateSourceCount:Int;
	final aggregateTotalValue:Int;
	final aggregateWeightedTotal:Int;
	final aggregateDigest:String;
	final alertCount:Int;
	final alertDigest:String;
	final runtimeScore:Int;

	public function new(profile:String, variant:String, capability:String, ingestReceived:Int, ingestAccepted:Int, backpressureEvents:Int, parseCount:Int,
			enrichCount:Int, aggregateSourceCount:Int, aggregateTotalValue:Int, aggregateWeightedTotal:Int, aggregateDigest:String, alertCount:Int,
			alertDigest:String, runtimeScore:Int) {
		this.profile = profile;
		this.variant = variant;
		this.capability = capability;
		this.ingestReceived = ingestReceived;
		this.ingestAccepted = ingestAccepted;
		this.backpressureEvents = backpressureEvents;
		this.parseCount = parseCount;
		this.enrichCount = enrichCount;
		this.aggregateSourceCount = aggregateSourceCount;
		this.aggregateTotalValue = aggregateTotalValue;
		this.aggregateWeightedTotal = aggregateWeightedTotal;
		this.aggregateDigest = aggregateDigest;
		this.alertCount = alertCount;
		this.alertDigest = alertDigest;
		this.runtimeScore = runtimeScore;
	}

	public function lines():Array<String> {
		return [
			"pulseforge.profile=" + profile,
			"pulseforge.variant=" + variant,
			"runtime.capability=" + capability,
			"ingest.received=" + ingestReceived,
			"ingest.accepted=" + ingestAccepted,
			"ingest.backpressure=" + backpressureEvents,
			"parse.events=" + parseCount,
			"enrich.events=" + enrichCount,
			"aggregate.sources=" + aggregateSourceCount,
			"aggregate.total=" + aggregateTotalValue,
			"aggregate.weighted_total=" + aggregateWeightedTotal,
			"aggregate.summary=" + aggregateDigest,
			"alert.count=" + alertCount,
			"alert.events=" + alertDigest,
			"runtime.score=" + runtimeScore
		];
	}

	public inline function profileId():String {
		return profile;
	}

	public inline function variantId():String {
		return variant;
	}

	public inline function capabilityId():String {
		return capability;
	}

	public inline function ingestReceivedCount():Int {
		return ingestReceived;
	}

	public inline function ingestAcceptedCount():Int {
		return ingestAccepted;
	}

	public inline function ingestBackpressureCount():Int {
		return backpressureEvents;
	}

	public inline function alertEventCount():Int {
		return alertCount;
	}

	public inline function alertEventDigest():String {
		return alertDigest;
	}

	public inline function score():Int {
		return runtimeScore;
	}

	public function render():String {
		var out = "";
		var values = lines();
		var i = 0;
		while (i < values.length) {
			if (i > 0) {
				out += "\n";
			}
			out += values[i];
			i++;
		}
		return out;
	}
}
