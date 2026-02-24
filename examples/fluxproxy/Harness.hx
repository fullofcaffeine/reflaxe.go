import app.core.FluxPipeline;
import app.core.FluxReport;
import app.core.FluxRequest;
import app.runtime.FluxRuntime;

class Harness {
	public static function baselineRequests():Array<FluxRequest> {
		return [
			new FluxRequest(1, "/v1/items", 30, 200),
			new FluxRequest(2, "/v1/items", 70, 503),
			new FluxRequest(3, "/assets/logo.png", 12, 200),
			new FluxRequest(4, "/health", 4, 200),
			new FluxRequest(5, "/v1/auth", 40, 502),
			new FluxRequest(6, "/v1/items", 18, 200),
			new FluxRequest(7, "/assets/main.css", 9, 200),
			new FluxRequest(8, "/v1/auth", 28, 200)
		];
	}

	static function cloneRequests(requests:Array<FluxRequest>):Array<FluxRequest> {
		var out = new Array<FluxRequest>();
		for (request in requests) {
			out.push(new FluxRequest(request.id, request.route, request.latencyMs, request.status));
		}
		return out;
	}

	static function runReport(runtime:FluxRuntime, requests:Array<FluxRequest>):FluxReport {
		var pipeline = new FluxPipeline(runtime);
		return pipeline.run(cloneRequests(requests));
	}

	public static function run(runtime:FluxRuntime):String {
		return runReport(runtime, baselineRequests()).render();
	}

	public static function runWithRequests(runtime:FluxRuntime, requests:Array<FluxRequest>):String {
		return runReport(runtime, requests).render();
	}

	static function assertBreakerScenario(runtime:FluxRuntime):Void {
		var breakerRequests = [
			new FluxRequest(1, "/breaker/api", 60, 503),
			new FluxRequest(2, "/breaker/api", 65, 502),
			new FluxRequest(3, "/breaker/api", 20, 200),
			new FluxRequest(4, "/breaker/api", 18, 200)
		];
		var report = runReport(runtime, breakerRequests);
		if (report.breakerOpen() != 2) {
			throw "policy.breaker_open scenario drift";
		}
		if (report.rateLimited() != 0) {
			throw "policy.rate_limited scenario drift";
		}
		if (report.retriesCount() != 2) {
			throw "proxy.retries scenario drift";
		}
	}

	public static function assertContract(runtime:FluxRuntime):String {
		var report = runReport(runtime, baselineRequests());
		if (report.profileId() != runtime.profileId()) {
			throw "profile drift";
		}
		if (report.variantId() != runtime.variantId()) {
			throw "variant drift";
		}
		if (report.capabilityId() != runtime.capabilityId()) {
			throw "capability drift";
		}
		if (report.receivedCount() != 8) {
			throw "ingress.received drift";
		}
		if (report.acceptedCount() != 8) {
			throw "ingress.accepted drift";
		}
		if (report.backpressureCount() != 5) {
			throw "ingress.backpressure drift";
		}
		if (report.retriesCount() != 2) {
			throw "proxy.retries drift";
		}
		if (report.rateLimited() != 1) {
			throw "policy.rate_limited drift";
		}
		if (report.breakerOpen() != 0) {
			throw "policy.breaker_open drift";
		}
		if (report.errors() != 3) {
			throw "errors.count drift";
		}

		var expectedScore = runtime.variantId() == "go_native" ? 35 : 26;
		if (report.score() != expectedScore) {
			throw "runtime.score drift";
		}
		assertBreakerScenario(runtime);
		return report.render();
	}
}
