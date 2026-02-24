package app.core;

class FluxRouteAggregate {
	public final route:String;
	public var count(default, null):Int;
	public var successCount(default, null):Int;
	public var errorCount(default, null):Int;
	public var totalLatencyMs(default, null):Int;

	public function new(route:String) {
		this.route = route;
		count = 0;
		successCount = 0;
		errorCount = 0;
		totalLatencyMs = 0;
	}

	public function record(response:FluxProxyResponse):Void {
		count++;
		totalLatencyMs += response.latencyMs;
		if (response.success) {
			successCount++;
		} else {
			errorCount++;
		}
	}

	public function averageLatencyMs():Int {
		if (count == 0) {
			return 0;
		}
		var remaining = totalLatencyMs;
		var quotient = 0;
		while (remaining >= count) {
			remaining -= count;
			quotient++;
		}
		return quotient;
	}

	public function summaryToken():String {
		return route + ":" + count + "/" + successCount + "/" + errorCount + "/" + averageLatencyMs();
	}
}
