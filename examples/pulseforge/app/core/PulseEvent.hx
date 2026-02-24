package app.core;

class PulseEvent {
	public final id:Int;
	public final source:String;
	public final region:String;
	public final value:Int;

	public function new(id:Int, source:String, region:String, value:Int) {
		this.id = id;
		this.source = source;
		this.region = region;
		this.value = value;
	}

	public inline function isAlert(minValue:Int):Bool {
		return value >= minValue;
	}
}
