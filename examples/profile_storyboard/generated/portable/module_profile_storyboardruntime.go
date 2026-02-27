package main

type profile__StoryboardRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(metrics *profile__StorySignalMetrics) *string
	supportsVelocityHint() bool
	velocityPerSprint() int
	riskThreshold() int
}
