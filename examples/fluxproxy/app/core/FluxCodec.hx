package app.core;

class FluxCodec {
	public static function proxy(request:FluxRequest, timeoutMs:Int):FluxProxyResponse {
		var route = normalizedRoute(request.route);
		var latency = request.latencyMs;
		if (latency < 0) {
			latency = 0;
		}
		var status = request.status;
		var attempts = status >= 500 ? 2 : 1;
		if (latency > timeoutMs) {
			status = 504;
			attempts = 2;
		}
		var success = status < 500;
		return new FluxProxyResponse(request.id, route, upstreamForRoute(route), status, latency, attempts, success);
	}

	public static function normalizedRoute(route:String):String {
		var trimmed = StringTools.trim(route);
		if (trimmed == "") {
			return "/unknown";
		}
		return trimmed;
	}

	public static function rateLimited(request:FluxRequest):FluxProxyResponse {
		var route = normalizedRoute(request.route);
		var latency = request.latencyMs;
		if (latency < 0) {
			latency = 0;
		}
		return new FluxProxyResponse(request.id, route, "rate-limit", 429, latency, 1, false);
	}

	public static function breakerOpen(request:FluxRequest):FluxProxyResponse {
		var route = normalizedRoute(request.route);
		return new FluxProxyResponse(request.id, route, "breaker-open", 503, 0, 1, false);
	}

	static function upstreamForRoute(route:String):String {
		if (StringTools.startsWith(route, "/assets")) {
			return "cdn";
		}
		if (route == "/health") {
			return "healthz";
		}
		return "core-api";
	}
}
