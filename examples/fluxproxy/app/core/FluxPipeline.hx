package app.core;

import app.runtime.BuildConfig;
import app.runtime.FluxRuntime;

private typedef FluxStringIntState = {
	var key:String;
	var value:Int;
}

private typedef FluxResponseById = {
	var requestId:Int;
	var response:FluxProxyResponse;
}

class FluxPipeline {
	final runtime:FluxRuntime;

	public function new(runtime:FluxRuntime) {
		this.runtime = runtime;
	}

	public function run(requests:Array<FluxRequest>):FluxReport {
		var ingest = ingest(requests, BuildConfig.INGEST_QUEUE_CAPACITY);
		var planned = applyRoutePolicies(ingest.acceptedRequests, BuildConfig.PER_ROUTE_LIMIT, BuildConfig.BREAKER_FAILURE_THRESHOLD, BuildConfig.TIMEOUT_MS);
		var dispatched = runtime.dispatch(planned.dispatchable, BuildConfig.DISPATCH_WORKERS);
		var responses = orderedResponses(planned.synthetic, dispatched, ingest.acceptedRequests);
		var aggregates = aggregate(responses);
		var retryCount = retries(responses);
		var errorCount = errors(responses);
		var score = runtime.stageScore(responses, retryCount, ingest.backpressureEvents);

		return new FluxReport(runtime.profileId(), runtime.variantId(), runtime.capabilityId(), ingest.receivedCount, ingest.acceptedRequests.length,
			ingest.backpressureEvents, responses.length, retryCount, planned.rateLimited, planned.breakerOpen, aggregates.routes.length, aggregates.summary,
			errorCount, score);
	}

	function ingest(requests:Array<FluxRequest>, capacity:Int):FluxIngestResult {
		var boundedCapacity = capacity <= 0 ? 1 : capacity;
		var queue = new Array<FluxRequest>();
		var queueHead = 0;
		var accepted = new Array<FluxRequest>();
		var backpressureEvents = 0;

		for (request in requests) {
			if (queue.length - queueHead >= boundedCapacity) {
				backpressureEvents++;
				accepted.push(queue[queueHead]);
				queueHead++;
			}
			queue.push(request);
		}

		while (queueHead < queue.length) {
			accepted.push(queue[queueHead]);
			queueHead++;
		}

		return new FluxIngestResult(requests.length, accepted, backpressureEvents);
	}

	function aggregate(responses:Array<FluxProxyResponse>):{
		routes:Array<FluxRouteAggregate>,
		summary:String
	} {
		var routes = new Array<FluxRouteAggregate>();

		for (response in responses) {
			var route = response.route;
			var bucket = findRouteAggregate(routes, route);
			if (bucket == null) {
				bucket = new FluxRouteAggregate(route);
				routes.push(bucket);
			}
			bucket.record(response);
		}

		var digest = "";
		for (item in routes) {
			if (digest != "") {
				digest += ",";
			}
			digest += item.summaryToken();
		}

		return {
			routes: routes,
			summary: digest
		};
	}

	function retries(responses:Array<FluxProxyResponse>):Int {
		var total = 0;
		for (response in responses) {
			total += response.attempts - 1;
		}
		return total;
	}

	function errors(responses:Array<FluxProxyResponse>):Int {
		var total = 0;
		for (response in responses) {
			if (!response.success) {
				total++;
			}
		}
		return total;
	}

	function applyRoutePolicies(requests:Array<FluxRequest>, perRouteLimit:Int, breakerFailureThreshold:Int, timeoutMs:Int):{
		dispatchable:Array<FluxRequest>,
		synthetic:Array<FluxProxyResponse>,
		rateLimited:Int,
		breakerOpen:Int
	} {
		var normalizedLimit = perRouteLimit <= 0 ? 1 : perRouteLimit;
		var normalizedBreaker = breakerFailureThreshold <= 0 ? 1 : breakerFailureThreshold;
		var routeCounts = new Array<FluxStringIntState>();
		var failureStreak = new Array<FluxStringIntState>();
		var dispatchable = new Array<FluxRequest>();
		var synthetic = new Array<FluxProxyResponse>();
		var rateLimited = 0;
		var breakerOpen = 0;

		for (request in requests) {
			var route = FluxCodec.normalizedRoute(request.route);
			var streak = getStringIntStateValue(failureStreak, route);
			if (streak >= normalizedBreaker) {
				synthetic.push(FluxCodec.breakerOpen(request));
				breakerOpen++;
				continue;
			}

			var routeCount = getStringIntStateValue(routeCounts, route);
			if (routeCount >= normalizedLimit) {
				synthetic.push(FluxCodec.rateLimited(request));
				rateLimited++;
				continue;
			}
			routeCounts = setStringIntStateValue(routeCounts, route, routeCount + 1);

			dispatchable.push(request);
			var predictsFailure = request.status >= 500 || request.latencyMs > timeoutMs;
			if (predictsFailure) {
				failureStreak = setStringIntStateValue(failureStreak, route, streak + 1);
			} else {
				failureStreak = setStringIntStateValue(failureStreak, route, 0);
			}
		}

		return {
			dispatchable: dispatchable,
			synthetic: synthetic,
			rateLimited: rateLimited,
			breakerOpen: breakerOpen
		};
	}

	function orderedResponses(synthetic:Array<FluxProxyResponse>, dispatched:Array<FluxProxyResponse>,
			acceptedRequests:Array<FluxRequest>):Array<FluxProxyResponse> {
		var byId = new Array<FluxResponseById>();
		for (response in synthetic) {
			byId = setResponseById(byId, response.requestId, response);
		}
		for (response in dispatched) {
			byId = setResponseById(byId, response.requestId, response);
		}

		var ordered = new Array<FluxProxyResponse>();
		for (request in acceptedRequests) {
			var response = getResponseById(byId, request.id);
			if (response != null) {
				ordered.push(response);
			}
		}
		return ordered;
	}

	function findRouteAggregate(routes:Array<FluxRouteAggregate>, route:String):Null<FluxRouteAggregate> {
		for (item in routes) {
			if (item.route == route) {
				return item;
			}
		}
		return null;
	}

	function getStringIntStateValue(states:Array<FluxStringIntState>, key:String):Int {
		for (state in states) {
			if (state.key == key) {
				return state.value;
			}
		}
		return 0;
	}

	function setStringIntStateValue(states:Array<FluxStringIntState>, key:String, value:Int):Array<FluxStringIntState> {
		for (state in states) {
			if (state.key == key) {
				state.value = value;
				return states;
			}
		}
		states.push({
			key: key,
			value: value
		});
		return states;
	}

	function getResponseById(states:Array<FluxResponseById>, requestId:Int):Null<FluxProxyResponse> {
		for (state in states) {
			if (state.requestId == requestId) {
				return state.response;
			}
		}
		return null;
	}

	function setResponseById(states:Array<FluxResponseById>, requestId:Int, response:FluxProxyResponse):Array<FluxResponseById> {
		for (state in states) {
			if (state.requestId == requestId) {
				state.response = response;
				return states;
			}
		}
		states.push({
			requestId: requestId,
			response: response
		});
		return states;
	}
}
