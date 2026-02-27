package profile;

class PortableRuntime implements TodoRuntime {
	public function new() {}

	public function profileId():String {
		return "portable";
	}

	public function normalizeTitle(title:String):String {
		return title;
	}

	public function normalizeTag(tag:String):String {
		return tag;
	}

	public function supportsBatchAdd():Bool {
		return false;
	}

	public function supportsDiagnostics():Bool {
		return false;
	}

	public function diagnostics(metrics:TodoRuntimeMetrics):String {
		return "off";
	}
}
