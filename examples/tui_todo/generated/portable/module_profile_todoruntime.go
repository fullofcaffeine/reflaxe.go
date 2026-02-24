package main

type profile__TodoRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
}
