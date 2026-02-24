package app.core;

class FluxProxyResponse {
	public final requestId:Int;
	public final route:String;
	public final upstream:String;
	public final status:Int;
	public final latencyMs:Int;
	public final attempts:Int;
	public final success:Bool;

	public function new(requestId:Int, route:String, upstream:String, status:Int, latencyMs:Int, attempts:Int, success:Bool) {
		this.requestId = requestId;
		this.route = route;
		this.upstream = upstream;
		this.status = status;
		this.latencyMs = latencyMs;
		this.attempts = attempts;
		this.success = success;
	}
}
