package main

import "examples_fluxproxy_portable/hxrt"

type app__runtime__FluxRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	dispatch(requests *hxrt.Array, workerCount int) *hxrt.Array
	stageScore(responses *hxrt.Array, retryCount int, backpressureEvents int) int
}
