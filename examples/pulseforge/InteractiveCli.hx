import app.core.PulseIngressFrame;
import app.core.PulsePipeline;
import app.core.PulseReport;
import app.runtime.PulseRuntime;
import haxe.io.Bytes;

class InteractiveCli {
	static function parsePositiveInt(raw:String):Int {
		if (raw == "") {
			return -1;
		}
		var bytes = Bytes.ofString(raw);
		var value = 0;
		var i = 0;
		while (i < bytes.length) {
			var code = bytes.get(i);
			if (code < 48 || code > 57) {
				return -1;
			}
			value = (value * 10) + (code - 48);
			i++;
		}
		return value;
	}

	static function decodeToken(raw:String):String {
		return StringTools.replace(raw, "_", " ");
	}

	static function nextSequence(frames:Array<PulseIngressFrame>):Int {
		var next = 1;
		for (frame in frames) {
			if (frame.sequence >= next) {
				next = frame.sequence + 1;
			}
		}
		return next;
	}

	static function runReport(runtime:PulseRuntime, frames:Array<PulseIngressFrame>):PulseReport {
		var pipeline = new PulsePipeline(runtime);
		return pipeline.run(frames);
	}

	static function liveLine(report:PulseReport):String {
		return "live ingest.received=" + report.ingestReceivedCount() + ",ingest.backpressure=" + report.ingestBackpressureCount() + ",alert.count="
			+ report.alertEventCount() + ",runtime.score=" + report.score();
	}

	static function printHelp(runtime:PulseRuntime):Void {
		Sys.println("commands:");
		Sys.println("  help");
		Sys.println("  profile");
		Sys.println("  reset");
		Sys.println("  status");
		Sys.println("  scripted");
		Sys.println("  ingest <source_token> <value> <region_token>");
		Sys.println("token note: use '_' for spaces");
		Sys.println("runtime=" + runtime.profileId() + "/" + runtime.variantId() + "/" + runtime.capabilityId());
	}

	static function printUsage(runtime:PulseRuntime):Void {
		Sys.println("pulseforge interactive command session (" + runtime.profileId() + ")");
		Sys.println("run scripted contract mode with: --scripted");
		Sys.println("commands:");
		Sys.println("  pulseforge help");
		Sys.println("  pulseforge profile");
		Sys.println("  pulseforge status");
		Sys.println("  pulseforge ingest edge 8 iad status");
		Sys.println("generated-source invocation:");
		Sys.println("  go run . --scripted");
		Sys.println("  go run . status");
	}

	static function failUsage(message:String):Void {
		Sys.println("error: " + message);
		Sys.println("run `help` for command syntax");
	}

	public static function run(runtime:PulseRuntime):Void {
		var frames = Harness.baselineFrames();
		var args = Sys.args();
		if (args.length == 0) {
			printUsage(runtime);
			return;
		}

		var i = 0;
		while (i < args.length) {
			var cmd = args[i];
			if (cmd == "help") {
				printHelp(runtime);
				i++;
				continue;
			}
			if (cmd == "profile") {
				Sys.println("profile=" + runtime.profileId() + ",variant=" + runtime.variantId() + ",capability=" + runtime.capabilityId());
				i++;
				continue;
			}
			if (cmd == "reset") {
				frames = Harness.baselineFrames();
				var resetReport = runReport(runtime, frames);
				Sys.println("ok reset");
				Sys.println(liveLine(resetReport));
				i++;
				continue;
			}
			if (cmd == "status") {
				var statusReport = runReport(runtime, frames);
				Sys.println(statusReport.render());
				i++;
				continue;
			}
			if (cmd == "scripted") {
				Sys.println(Harness.runWithFrames(runtime, frames));
				i++;
				continue;
			}
			if (cmd == "ingest") {
				if (i + 3 >= args.length) {
					failUsage("ingest requires <source_token> <value> <region_token>");
					return;
				}
				var source = decodeToken(args[i + 1]);
				var value = parsePositiveInt(args[i + 2]);
				if (value < 0) {
					failUsage("invalid value: " + args[i + 2]);
					return;
				}
				var region = decodeToken(args[i + 3]);
				var sequence = nextSequence(frames);
				frames.push(new PulseIngressFrame(sequence, source, value, region));
				var ingestReport = runReport(runtime, frames);
				Sys.println("ok ingest seq=" + sequence);
				Sys.println(liveLine(ingestReport));
				i += 4;
				continue;
			}

			failUsage("unknown command: " + cmd);
			return;
		}
	}
}
