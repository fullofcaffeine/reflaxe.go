package main

type app__runtime__PulseRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	parse(frames []*app__core__PulseIngressFrame, workerCount int) []*app__core__PulseEvent
	enrich(events []*app__core__PulseEvent, workerCount int) []*app__core__PulseEnrichedEvent
	stageScore(parsed []*app__core__PulseEvent, enriched []*app__core__PulseEnrichedEvent, alerts []*app__core__PulseAlert, backpressureEvents int) int
}
