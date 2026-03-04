package main

type profile__TodoRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	diagnostics(metrics *profile__TodoRuntimeMetrics) *string
}
