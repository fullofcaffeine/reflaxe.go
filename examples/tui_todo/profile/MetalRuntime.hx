package profile;

class MetalRuntime implements TodoRuntime {
	public function new() {}

	public function profileId():String {
		return "metal";
	}

	public function normalizeTitle(title:String):String {
		return title;
	}

	public function normalizeTag(tag:String):String {
		return "metal-" + tag;
	}

	public function supportsBatchAdd():Bool {
		return true;
	}

	public function supportsDiagnostics():Bool {
		return true;
	}

	public function diagnostics(metrics:TodoRuntimeMetrics):String {
		return "p1=" + metrics.p1 + ",completed=" + metrics.done;
	}
}
