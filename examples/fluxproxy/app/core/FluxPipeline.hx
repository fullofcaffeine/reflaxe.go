package app.core;

import app.runtime.BuildConfig;
import app.runtime.FluxRuntime;
import haxe.ds.IntMap;
import haxe.ds.StringMap;

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
		var byRoute:StringMap<FluxRouteAggregate> = new StringMap<FluxRouteAggregate>();
		var routeKeys = new Array<String>();

		for (response in responses) {
			var route = response.route;
			var bucket = byRoute.get(route);
			if (bucket == null) {
				bucket = new FluxRouteAggregate(route);
				byRoute.set(route, bucket);
				routeKeys.push(route);
			}
			bucket.record(response);
		}

		var routes = new Array<FluxRouteAggregate>();
		var digest = "";
		var i = 0;
		while (i < routeKeys.length) {
			var route = routeKeys[i];
			var item = byRoute.get(route);
			if (item != null) {
				routes.push(item);
				if (digest != "") {
					digest += ",";
				}
				digest += item.summaryToken();
			}
			i++;
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
		var routeCounts:StringMap<Int> = new StringMap<Int>();
		var failureStreak:StringMap<Int> = new StringMap<Int>();
		var dispatchable = new Array<FluxRequest>();
		var synthetic = new Array<FluxProxyResponse>();
		var rateLimited = 0;
		var breakerOpen = 0;

		for (request in requests) {
			var route = FluxCodec.normalizedRoute(request.route);
			var streak = 0;
			if (failureStreak.exists(route)) {
				streak += failureStreak.get(route);
			}
			if (streak >= normalizedBreaker) {
				synthetic.push(FluxCodec.breakerOpen(request));
				breakerOpen++;
				continue;
			}

			var routeCount = 0;
			if (routeCounts.exists(route)) {
				routeCount += routeCounts.get(route);
			}
			if (routeCount >= normalizedLimit) {
				synthetic.push(FluxCodec.rateLimited(request));
				rateLimited++;
				continue;
			}
			routeCounts.set(route, routeCount + 1);

			dispatchable.push(request);
			var predictsFailure = request.status >= 500 || request.latencyMs > timeoutMs;
			if (predictsFailure) {
				failureStreak.set(route, streak + 1);
			} else {
				failureStreak.set(route, 0);
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
		var byId:IntMap<FluxProxyResponse> = new IntMap<FluxProxyResponse>();
		for (response in synthetic) {
			byId.set(response.requestId, response);
		}
		for (response in dispatched) {
			byId.set(response.requestId, response);
		}

		var ordered = new Array<FluxProxyResponse>();
		for (request in acceptedRequests) {
			var response = byId.get(request.id);
			if (response != null) {
				ordered.push(response);
			}
		}
		return ordered;
	}
}
