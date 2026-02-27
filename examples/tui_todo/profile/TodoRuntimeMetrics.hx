package profile;

class TodoRuntimeMetrics {
	public final total:Int;
	public final open:Int;
	public final done:Int;
	public final p1:Int;

	public function new(total:Int, open:Int, done:Int, p1:Int) {
		this.total = total;
		this.open = open;
		this.done = done;
		this.p1 = p1;
	}
}
