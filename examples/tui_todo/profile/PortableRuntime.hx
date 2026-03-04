package profile;

/**
	Portable runtime implementation for `tui_todo`.

	This runtime owns title/tag normalization and diagnostics string formatting for
	the portable scripted contract. No separate metal runtime is kept because it
	did not provide meaningful user-facing value in this app.
**/
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

	public function diagnostics(metrics:TodoRuntimeMetrics):String {
		return "p1=" + metrics.p1 + ",completed=" + metrics.done;
	}
}
