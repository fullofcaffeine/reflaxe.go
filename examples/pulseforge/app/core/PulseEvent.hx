package app.core;

class PulseEvent {
	public final id:Int;
	public final source:String;
	public final value:Int;

	public function new(id:Int, source:String, value:Int) {
		this.id = id;
		this.source = source;
		this.value = value;
	}

	public inline function isAlert():Bool {
		return value >= 8;
	}
}
