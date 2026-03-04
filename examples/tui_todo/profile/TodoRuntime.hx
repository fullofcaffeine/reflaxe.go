package profile;

/**
	Runtime-contract interface for `tui_todo`.

	This example is intentionally portable-only. It focuses on stable CLI contract
	behavior rather than profile-specific codegen divergence.
**/
interface TodoRuntime {
	public function profileId():String;
	public function normalizeTitle(title:String):String;
	public function normalizeTag(tag:String):String;
	public function diagnostics(metrics:TodoRuntimeMetrics):String;
}
