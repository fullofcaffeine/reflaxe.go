import app.core.PulseEvent;
import app.core.PulsePipeline;
import app.runtime.RuntimeFactory;

class Main {
	static function workload():Array<PulseEvent> {
		return [
			new PulseEvent(1, "edge", 3),
			new PulseEvent(2, "api", 7),
			new PulseEvent(3, "db", 11),
			new PulseEvent(4, "edge", 2),
			new PulseEvent(5, "auth", 13),
			new PulseEvent(6, "worker", 5)
		];
	}

	static function main():Void {
		var runtime = RuntimeFactory.create();
		var pipeline = new PulsePipeline(runtime);
		var report = pipeline.run(workload());
		for (line in report.lines()) {
			Sys.println(line);
		}
	}
}
