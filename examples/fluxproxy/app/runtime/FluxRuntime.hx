package app.runtime;

import app.core.FluxProxyResponse;
import app.core.FluxRequest;

interface FluxRuntime {
	public function profileId():String;
	public function variantId():String;
	public function capabilityId():String;
	public function dispatch(requests:Array<FluxRequest>, workerCount:Int):Array<FluxProxyResponse>;
	public function stageScore(responses:Array<FluxProxyResponse>, retryCount:Int, backpressureEvents:Int):Int;
}
