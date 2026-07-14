package app.runtime;

import app.core.FluxCodec;
import app.core.FluxProxyResponse;
import app.core.FluxRequest;
import go.Chan;
import go.Go;
import go.Select;

/**
	Go-first runtime adapter for FluxProxy `go_native` variant.

	Why this differs from `CoreRuntime`:
	- uses worker/channel/select dispatch paths,
	- stresses typed specialization in generated Go for hot request flows,
	- keeps the same routing/policy output contract as `CoreRuntime`.

	Use this lane when benchmarking or tuning Go-native execution behavior.
**/
@:goNative
class GoNativeRuntime implements FluxRuntime {
	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "worker_chan_fanout";
	}

	public function dispatch(requests:Array<FluxRequest>, workerCount:Int):Array<FluxProxyResponse> {
		if (requests.length == 0) {
			return [];
		}

		var workers = normalizedWorkers(workerCount);
		var inbox:Chan<FluxRequest> = Go.newChan(requests.length);
		var out:Chan<FluxProxyResponse> = Go.newChan(requests.length);
		var done:Chan<Int> = Go.newChan(workers);

		for (request in requests) {
			inbox.send(request);
		}
		inbox.close();

		var index = 0;
		while (index < workers) {
			Go.spawn(function() {
				var processed = 0;
				while (true) {
					var request = inbox.recv();
					if (request == null) {
						break;
					}
					out.send(FluxCodec.proxy(cast request, BuildConfig.TIMEOUT_MS));
					processed++;
				}
				done.send(processed);
			});
			index++;
		}

		waitForWorkers(done, workers);
		return orderResponses(drain(out, requests.length), requests);
	}

	public function stageScore(responses:Array<FluxProxyResponse>, retryCount:Int, backpressureEvents:Int):Int {
		var latency:Chan<Int> = Go.newChan(responses.length);
		var score = 0;

		for (response in responses) {
			var step = response.success ? 5 : -8;
			score += switch (Select.send(latency, response.latencyMs)) {
				case Sent: step;
				case Defaulted: -20;
			};
		}

		var remaining = responses.length;
		while (remaining > 0) {
			score += switch (Select.recv(latency)) {
				case Received(value): divFloorPositive(value, 20);
				case Defaulted:
					0;
			};
			remaining--;
		}

		latency.close();
		score += responses.length * 4;
		score -= backpressureEvents * 3;
		score += retryCount * 5;
		return score;
	}

	function divFloorPositive(numerator:Int, denominator:Int):Int {
		if (denominator <= 0) {
			return 0;
		}
		var remaining = numerator;
		var quotient = 0;
		while (remaining >= denominator) {
			remaining -= denominator;
			quotient++;
		}
		return quotient;
	}

	function normalizedWorkers(workerCount:Int):Int {
		if (workerCount <= 0) {
			return 1;
		}
		return workerCount;
	}

	function waitForWorkers(done:Chan<Int>, workers:Int):Void {
		var completed = 0;
		while (completed < workers) {
			done.recv();
			completed++;
		}
		done.close();
	}

	function drain(out:Chan<FluxProxyResponse>, expected:Int):Array<FluxProxyResponse> {
		var responses = new Array<FluxProxyResponse>();
		var remaining = expected;
		while (remaining > 0) {
			switch (Select.recv(out)) {
				case Received(response):
					responses.push(response);
					remaining--;
				case Defaulted:
					var fallback = out.recv();
					if (fallback != null) {
						responses.push(cast fallback);
						remaining--;
					}
			}
		}
		out.close();
		return responses;
	}

	function orderResponses(items:Array<FluxProxyResponse>, requests:Array<FluxRequest>):Array<FluxProxyResponse> {
		var maxId = 0;
		for (request in requests) {
			if (request.id > maxId) {
				maxId = request.id;
			}
		}

		var byId = new Array<Null<FluxProxyResponse>>();
		var size = maxId + 1;
		var i = 0;
		while (i < size) {
			byId.push(null);
			i++;
		}

		for (item in items) {
			if (item.requestId >= 0 && item.requestId <= maxId) {
				byId[item.requestId] = item;
			}
		}

		var ordered = new Array<FluxProxyResponse>();
		for (request in requests) {
			var response:Null<FluxProxyResponse> = null;
			if (request.id >= 0 && request.id <= maxId) {
				response = byId[request.id];
			}
			if (response != null) {
				ordered.push(cast response);
			}
		}
		return ordered;
	}
}
