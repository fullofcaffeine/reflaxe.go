package profile;

interface StoryboardRuntime {
	public function profileId():String;
	public function decorateTitle(title:String):String;
	public function highlightTag(tag:String):String;
	public function extraSignal(metrics:StorySignalMetrics):String;
	public function supportsVelocityHint():Bool;
	public function velocityPerSprint():Int;
	public function riskThreshold():Int;
}
