package app.core;

class PulseIngressFrame {
	public final sequence:Int;
	public final source:String;
	public final value:Int;
	public final region:String;

	public function new(sequence:Int, source:String, value:Int, region:String) {
		this.sequence = sequence;
		this.source = source;
		this.value = value;
		this.region = region;
	}
}
