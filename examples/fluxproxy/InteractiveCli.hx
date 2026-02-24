import app.core.FluxPipeline;
import app.core.FluxReport;
import app.core.FluxRequest;
import app.runtime.FluxRuntime;
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

	static function nextId(requests:Array<FluxRequest>):Int {
		var next = 1;
		for (request in requests) {
			if (request.id >= next) {
				next = request.id + 1;
			}
		}
		return next;
	}

	static function runReport(runtime:FluxRuntime, requests:Array<FluxRequest>):FluxReport {
		var pipeline = new FluxPipeline(runtime);
		return pipeline.run(requests);
	}

	static function liveLine(report:FluxReport):String {
		return "live ingress.received=" + report.receivedCount() + ",ingress.backpressure=" + report.backpressureCount() + ",proxy.retries="
			+ report.retriesCount() + ",errors.count=" + report.errors() + ",runtime.score=" + report.score();
	}

	static function printHelp(runtime:FluxRuntime):Void {
		Sys.println("commands:");
		Sys.println("  help");
		Sys.println("  profile");
		Sys.println("  reset");
		Sys.println("  status");
		Sys.println("  scripted");
		Sys.println("  ingest <route_token> <latency_ms> <status_code>");
		Sys.println("token note: use '_' for spaces");
		Sys.println("runtime=" + runtime.profileId() + "/" + runtime.variantId() + "/" + runtime.capabilityId());
	}

	static function printUsage(runtime:FluxRuntime):Void {
		Sys.println("fluxproxy interactive command session (" + runtime.profileId() + ")");
		Sys.println("run scripted contract mode with: --scripted");
		Sys.println("commands:");
		Sys.println("  fluxproxy help");
		Sys.println("  fluxproxy profile");
		Sys.println("  fluxproxy status");
		Sys.println("  fluxproxy ingest /v1/items 45 200 status");
		Sys.println("generated-source invocation:");
		Sys.println("  go run . --scripted");
		Sys.println("  go run . status");
	}

	static function failUsage(message:String):Void {
		Sys.println("error: " + message);
		Sys.println("run `help` for command syntax");
	}

	public static function run(runtime:FluxRuntime):Void {
		var requests = Harness.baselineRequests();
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
				requests = Harness.baselineRequests();
				var resetReport = runReport(runtime, requests);
				Sys.println("ok reset");
				Sys.println(liveLine(resetReport));
				i++;
				continue;
			}
			if (cmd == "status") {
				var statusReport = runReport(runtime, requests);
				Sys.println(statusReport.render());
				i++;
				continue;
			}
			if (cmd == "scripted") {
				Sys.println(Harness.runWithRequests(runtime, requests));
				i++;
				continue;
			}
			if (cmd == "ingest") {
				if (i + 3 >= args.length) {
					failUsage("ingest requires <route_token> <latency_ms> <status_code>");
					return;
				}

				var route = decodeToken(args[i + 1]);
				var latency = parsePositiveInt(args[i + 2]);
				if (latency < 0) {
					failUsage("invalid latency_ms: " + args[i + 2]);
					return;
				}
				var status = parsePositiveInt(args[i + 3]);
				if (status < 100 || status > 599) {
					failUsage("invalid status_code: " + args[i + 3]);
					return;
				}

				var requestId = nextId(requests);
				requests.push(new FluxRequest(requestId, route, latency, status));
				var ingestReport = runReport(runtime, requests);
				Sys.println("ok ingest id=" + requestId);
				Sys.println(liveLine(ingestReport));
				i += 4;
				continue;
			}

			failUsage("unknown command: " + cmd);
			return;
		}
	}
}
