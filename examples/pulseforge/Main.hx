import app.core.PulsePipeline;
import app.core.PulseIngressFrame;
import app.runtime.RuntimeFactory;

class Main {
	static function workload():Array<PulseIngressFrame> {
		return [
			new PulseIngressFrame(1, "edge", 3, "iad"),
			new PulseIngressFrame(2, "api", 7, "sfo"),
			new PulseIngressFrame(3, "db", 11, "fra"),
			new PulseIngressFrame(4, "edge", 2, "iad"),
			new PulseIngressFrame(5, "auth", 13, "gru"),
			new PulseIngressFrame(6, "worker", 5, "sfo"),
			new PulseIngressFrame(7, "api", 9, "fra"),
			new PulseIngressFrame(8, "db", 4, "iad")
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
