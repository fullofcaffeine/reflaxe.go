package profile;

import domain.StoryCard;
import haxe.ds.List;

interface StoryboardRuntime {
	public function profileId():String;
	public function decorateTitle(title:String):String;
	public function highlightTag(tag:String):String;
	public function extraSignal(cards:List<StoryCard>):String;
	public function supportsVelocityHint():Bool;
	public function velocityPerSprint():Int;
	public function riskThreshold():Int;
}
