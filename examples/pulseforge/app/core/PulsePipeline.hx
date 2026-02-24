package app.core;

import app.runtime.BuildConfig;
import app.runtime.PulseRuntime;
import haxe.ds.StringMap;

class PulsePipeline {
	final runtime:PulseRuntime;

	public function new(runtime:PulseRuntime) {
		this.runtime = runtime;
	}

	public function run(frames:Array<PulseIngressFrame>):PulseReport {
		var ingest = ingest(frames, BuildConfig.INGEST_QUEUE_CAPACITY);
		var parsed = runtime.parse(ingest.acceptedFrames, BuildConfig.PARSE_WORKERS);
		var enriched = runtime.enrich(parsed, BuildConfig.ENRICH_WORKERS);

		var aggregates = aggregate(enriched);
		var alerts = collectAlerts(enriched, BuildConfig.ALERT_WEIGHTED_THRESHOLD);
		var runtimeScore = runtime.stageScore(parsed, enriched, alerts, ingest.backpressureEvents);
		var alertDigest = alertToken(alerts);

		return new PulseReport(runtime.profileId(), runtime.variantId(), runtime.capabilityId(), ingest.receivedCount, ingest.acceptedFrames.length,
			ingest.backpressureEvents, parsed.length, enriched.length, aggregates.sources.length, aggregates.totalValue, aggregates.totalWeighted,
			aggregates.summary, alerts.length, alertDigest, runtimeScore);
	}

	function ingest(frames:Array<PulseIngressFrame>, capacity:Int):PulseIngestResult {
		var boundedCapacity = capacity <= 0 ? 1 : capacity;
		var queue = new Array<PulseIngressFrame>();
		var queueHead = 0;
		var accepted = new Array<PulseIngressFrame>();
		var backpressureEvents = 0;
		for (frame in frames) {
			if (queue.length - queueHead >= boundedCapacity) {
				backpressureEvents++;
				accepted.push(queue[queueHead]);
				queueHead++;
			}

			queue.push(frame);
		}

		while (queueHead < queue.length) {
			accepted.push(queue[queueHead]);
			queueHead++;
		}

		return new PulseIngestResult(frames.length, accepted, backpressureEvents);
	}

	function aggregate(enriched:Array<PulseEnrichedEvent>):{
		sources:Array<PulseSourceAggregate>,
		totalValue:Int,
		totalWeighted:Int,
		summary:String
	} {
		var bySource:StringMap<PulseSourceAggregate> = new StringMap<PulseSourceAggregate>();
		var sourceKeys = new Array<String>();
		var totalValue = 0;
		var totalWeighted = 0;

		for (entry in enriched) {
			totalValue += entry.event.value;
			totalWeighted += entry.weightedValue;
			var source = entry.event.source;
			var bucket = bySource.get(source);
			if (bucket == null) {
				bucket = new PulseSourceAggregate(source);
				bySource.set(source, bucket);
				sourceKeys.push(source);
			}
			bucket.record(entry);
		}

		var sourceSummaries = new Array<PulseSourceAggregate>();
		var digest = "";
		var index = 0;
		while (index < sourceKeys.length) {
			var source = sourceKeys[index];
			var summary = bySource.get(source);
			if (summary != null) {
				sourceSummaries.push(summary);
				if (digest != "") {
					digest += ",";
				}
				digest += summary.summaryToken();
			}
			index++;
		}

		return {
			sources: sourceSummaries,
			totalValue: totalValue,
			totalWeighted: totalWeighted,
			summary: digest
		};
	}

	function collectAlerts(enriched:Array<PulseEnrichedEvent>, weightedThreshold:Int):Array<PulseAlert> {
		var alerts = new Array<PulseAlert>();
		for (entry in enriched) {
			if (entry.shouldAlert(weightedThreshold)) {
				alerts.push(PulseAlert.fromEnriched(entry));
			}
		}
		return alerts;
	}

	function alertToken(alerts:Array<PulseAlert>):String {
		if (alerts.length == 0) {
			return "none";
		}
		var digest = "";
		var index = 0;
		while (index < alerts.length) {
			if (index > 0) {
				digest += ",";
			}
			digest += Std.string(alerts[index].eventId);
			index++;
		}
		return digest;
	}
}
