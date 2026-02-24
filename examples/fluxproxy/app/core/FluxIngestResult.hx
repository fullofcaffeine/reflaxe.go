package app.core;

class FluxIngestResult {
	public final receivedCount:Int;
	public final acceptedRequests:Array<FluxRequest>;
	public final backpressureEvents:Int;

	public function new(receivedCount:Int, acceptedRequests:Array<FluxRequest>, backpressureEvents:Int) {
		this.receivedCount = receivedCount;
		this.acceptedRequests = acceptedRequests;
		this.backpressureEvents = backpressureEvents;
	}
}
