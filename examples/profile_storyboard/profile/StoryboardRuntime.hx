package profile;

/**
	Runtime-contract interface for `profile_storyboard`.

	This example is intentionally portable-only. A previous metal runtime was
	removed because it introduced synthetic differences that did not provide real
	profile value for users.
**/
interface StoryboardRuntime {
	public function profileId():String;
	public function decorateTitle(title:String):String;
	public function highlightTag(tag:String):String;
	public function extraSignal(metrics:StorySignalMetrics):String;
	public function supportsVelocityHint():Bool;
	public function velocityPerSprint():Int;
	public function riskThreshold():Int;
}
