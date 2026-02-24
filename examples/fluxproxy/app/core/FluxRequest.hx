package app.core;

class FluxRequest {
	public final id:Int;
	public final route:String;
	public final latencyMs:Int;
	public final status:Int;

	public function new(id:Int, route:String, latencyMs:Int, status:Int) {
		this.id = id;
		this.route = route;
		this.latencyMs = latencyMs;
		this.status = status;
	}
}
