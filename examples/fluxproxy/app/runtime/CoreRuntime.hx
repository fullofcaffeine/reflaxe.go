package app.runtime;

import app.core.FluxCodec;
import app.core.FluxProxyResponse;
import app.core.FluxRequest;

/**
	Portable-baseline runtime adapter for FluxProxy.

	This runtime keeps dispatch logic simple and deterministic (loop-based). It is
	the default reference behavior for portability and contract verification.

	See `GoNativeRuntime` for the Go-first runtime lane that changes dispatch
	strategy while preserving the same policy results.
**/
class CoreRuntime implements FluxRuntime {
	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "loop_dispatch";
	}

	public function dispatch(requests:Array<FluxRequest>, workerCount:Int):Array<FluxProxyResponse> {
		var responses = new Array<FluxProxyResponse>();
		for (request in requests) {
			responses.push(FluxCodec.proxy(request, BuildConfig.TIMEOUT_MS));
		}
		return responses;
	}

	public function stageScore(responses:Array<FluxProxyResponse>, retryCount:Int, backpressureEvents:Int):Int {
		var successCount = 0;
		var errorCount = 0;
		for (response in responses) {
			if (response.success) {
				successCount++;
			} else {
				errorCount++;
			}
		}

		var score = 0;
		score += successCount * 10;
		score -= errorCount * 6;
		score -= backpressureEvents * 2;
		score -= retryCount * 2;
		score += responses.length;
		return score;
	}
}
