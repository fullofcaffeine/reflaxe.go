package main

type app__runtime__FluxRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	dispatch(requests []*app__core__FluxRequest, workerCount int) []*app__core__FluxProxyResponse
	stageScore(responses []*app__core__FluxProxyResponse, retryCount int, backpressureEvents int) int
}
