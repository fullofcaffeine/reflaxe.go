package main

import "examples_pulseforge_metal/hxrt"

type app__runtime__PulseRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	parse(frames *hxrt.Array, workerCount int) *hxrt.Array
	enrich(events *hxrt.Array, workerCount int) *hxrt.Array
	stageScore(parsed *hxrt.Array, enriched *hxrt.Array, alerts *hxrt.Array, backpressureEvents int) int
}
