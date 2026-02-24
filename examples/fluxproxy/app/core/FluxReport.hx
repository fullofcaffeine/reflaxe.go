package app.core;

class FluxReport {
	final profile:String;
	final variant:String;
	final capability:String;
	final ingressReceived:Int;
	final ingressAccepted:Int;
	final ingressBackpressure:Int;
	final proxyResponses:Int;
	final proxyRetries:Int;
	final rateLimitedCount:Int;
	final breakerOpenCount:Int;
	final routesCount:Int;
	final routesSummary:String;
	final errorsCount:Int;
	final runtimeScore:Int;

	public function new(profile:String, variant:String, capability:String, ingressReceived:Int, ingressAccepted:Int, ingressBackpressure:Int,
			proxyResponses:Int, proxyRetries:Int, rateLimitedCount:Int, breakerOpenCount:Int, routesCount:Int, routesSummary:String, errorsCount:Int,
			runtimeScore:Int) {
		this.profile = profile;
		this.variant = variant;
		this.capability = capability;
		this.ingressReceived = ingressReceived;
		this.ingressAccepted = ingressAccepted;
		this.ingressBackpressure = ingressBackpressure;
		this.proxyResponses = proxyResponses;
		this.proxyRetries = proxyRetries;
		this.rateLimitedCount = rateLimitedCount;
		this.breakerOpenCount = breakerOpenCount;
		this.routesCount = routesCount;
		this.routesSummary = routesSummary;
		this.errorsCount = errorsCount;
		this.runtimeScore = runtimeScore;
	}

	public function lines():Array<String> {
		return [
			"fluxproxy.profile=" + profile,
			"fluxproxy.variant=" + variant,
			"runtime.capability=" + capability,
			"ingress.received=" + ingressReceived,
			"ingress.accepted=" + ingressAccepted,
			"ingress.backpressure=" + ingressBackpressure,
			"proxy.responses=" + proxyResponses,
			"proxy.retries=" + proxyRetries,
			"policy.rate_limited=" + rateLimitedCount,
			"policy.breaker_open=" + breakerOpenCount,
			"routes.count=" + routesCount,
			"routes.summary=" + routesSummary,
			"errors.count=" + errorsCount,
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

	public inline function receivedCount():Int {
		return ingressReceived;
	}

	public inline function acceptedCount():Int {
		return ingressAccepted;
	}

	public inline function backpressureCount():Int {
		return ingressBackpressure;
	}

	public inline function retriesCount():Int {
		return proxyRetries;
	}

	public inline function rateLimited():Int {
		return rateLimitedCount;
	}

	public inline function breakerOpen():Int {
		return breakerOpenCount;
	}

	public inline function errors():Int {
		return errorsCount;
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
