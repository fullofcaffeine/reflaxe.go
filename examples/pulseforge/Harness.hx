import app.core.PulseIngressFrame;
import app.core.PulsePipeline;
import app.core.PulseReport;
import app.runtime.PulseRuntime;

class Harness {
	public static function baselineFrames():Array<PulseIngressFrame> {
		return [
			new PulseIngressFrame(1, "edge", 3, "iad"),
			new PulseIngressFrame(2, "api", 7, "sfo"),
			new PulseIngressFrame(3, "db", 11, "fra"),
			new PulseIngressFrame(4, "edge", 2, "iad"),
			new PulseIngressFrame(5, "auth", 13, "gru"),
			new PulseIngressFrame(6, "worker", 5, "sfo"),
			new PulseIngressFrame(7, "api", 9, "fra"),
			new PulseIngressFrame(8, "db", 4, "iad")
		];
	}

	static function cloneFrames(frames:Array<PulseIngressFrame>):Array<PulseIngressFrame> {
		var out = new Array<PulseIngressFrame>();
		for (frame in frames) {
			out.push(new PulseIngressFrame(frame.sequence, frame.source, frame.value, frame.region));
		}
		return out;
	}

	static function runReport(runtime:PulseRuntime, frames:Array<PulseIngressFrame>):PulseReport {
		var pipeline = new PulsePipeline(runtime);
		return pipeline.run(cloneFrames(frames));
	}

	public static function run(runtime:PulseRuntime):String {
		return runReport(runtime, baselineFrames()).render();
	}

	public static function runWithFrames(runtime:PulseRuntime, frames:Array<PulseIngressFrame>):String {
		return runReport(runtime, frames).render();
	}

	public static function assertContract(runtime:PulseRuntime):String {
		var report = runReport(runtime, baselineFrames());

		if (report.profileId() != runtime.profileId()) {
			throw "profile drift";
		}
		if (report.variantId() != runtime.variantId()) {
			throw "variant drift";
		}
		if (report.capabilityId() != runtime.capabilityId()) {
			throw "capability drift";
		}
		if (report.ingestReceivedCount() != 8) {
			throw "ingest.received drift";
		}
		if (report.ingestAcceptedCount() != 8) {
			throw "ingest.accepted drift";
		}
		if (report.ingestBackpressureCount() != 5) {
			throw "ingest.backpressure drift";
		}
		if (report.alertEventCount() != 2) {
			throw "alert.count drift";
		}
		if (report.alertEventDigest() != "3,5") {
			throw "alert.events drift";
		}

		var expectedScore = runtime.variantId() == "go_native" ? 123 : 108;
		if (report.score() != expectedScore) {
			throw "runtime.score drift";
		}

		return report.render();
	}
}
