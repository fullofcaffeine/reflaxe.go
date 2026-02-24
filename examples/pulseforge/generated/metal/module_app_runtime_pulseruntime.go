package main

type app__runtime__PulseRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	processScore(events []*app__core__PulseEvent) int
}
