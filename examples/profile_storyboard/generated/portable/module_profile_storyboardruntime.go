package main

type profile__StoryboardRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
	supportsVelocityHint() bool
	velocityPerSprint() int
	riskThreshold() int
}
