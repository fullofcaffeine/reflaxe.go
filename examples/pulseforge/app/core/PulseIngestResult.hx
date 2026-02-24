package app.core;

class PulseIngestResult {
	public final receivedCount:Int;
	public final acceptedFrames:Array<PulseIngressFrame>;
	public final backpressureEvents:Int;

	public function new(receivedCount:Int, acceptedFrames:Array<PulseIngressFrame>, backpressureEvents:Int) {
		this.receivedCount = receivedCount;
		this.acceptedFrames = acceptedFrames;
		this.backpressureEvents = backpressureEvents;
	}
}
