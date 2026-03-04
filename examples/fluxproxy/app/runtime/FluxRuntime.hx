package app.runtime;

import app.core.FluxProxyResponse;
import app.core.FluxRequest;

/**
	Execution-strategy contract for FluxProxy runtime lanes.

	`CoreRuntime` is the portable baseline implementation (loop dispatch).
	`GoNativeRuntime` is the Go-first implementation (worker/channel/select dispatch).

	Both implementations must preserve the same routing/policy result contract.
**/
interface FluxRuntime {
	public function profileId():String;
	public function variantId():String;
	public function capabilityId():String;
	public function dispatch(requests:Array<FluxRequest>, workerCount:Int):Array<FluxProxyResponse>;
	public function stageScore(responses:Array<FluxProxyResponse>, retryCount:Int, backpressureEvents:Int):Int;
}
