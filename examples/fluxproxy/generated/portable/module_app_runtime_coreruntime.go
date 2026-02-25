package main

import "examples_fluxproxy_portable/hxrt"

type I_app__runtime__CoreRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	dispatch(requests []*app__core__FluxRequest, workerCount int) []*app__core__FluxProxyResponse
	stageScore(responses []*app__core__FluxProxyResponse, retryCount int, backpressureEvents int) int
}

type app__runtime__CoreRuntime struct {
	__hx_this I_app__runtime__CoreRuntime
}

func New_app__runtime__CoreRuntime() *app__runtime__CoreRuntime {
	self := &app__runtime__CoreRuntime{}
	self.__hx_this = self
	return self
}

func (self *app__runtime__CoreRuntime) profileId() *string {
	return hxrt.StringFromLiteral("portable")
}

func (self *app__runtime__CoreRuntime) variantId() *string {
	return hxrt.StringFromLiteral("core")
}

func (self *app__runtime__CoreRuntime) capabilityId() *string {
	return hxrt.StringFromLiteral("loop_dispatch")
}

func (self *app__runtime__CoreRuntime) dispatch(requests []*app__core__FluxRequest, workerCount int) []*app__core__FluxProxyResponse {
	responses := []*app__core__FluxProxyResponse{}
	_g := 0
	for _g < len(requests) {
		request := requests[_g]
		_g = int(int32((_g + 1)))
		responses = append(responses, app__core__FluxCodec_proxy(request, 50))
	}
	return responses
}

func (self *app__runtime__CoreRuntime) stageScore(responses []*app__core__FluxProxyResponse, retryCount int, backpressureEvents int) int {
	successCount := 0
	errorCount := 0
	_g := 0
	for _g < len(responses) {
		response := responses[_g]
		_g = int(int32((_g + 1)))
		if response.success {
			successCount = int(int32((successCount + 1)))
		} else {
			errorCount = int(int32((errorCount + 1)))
		}
	}
	score := 0
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(successCount) * hxrt.Int32Wrap(10))))))))
	score = int(int32((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(errorCount) * hxrt.Int32Wrap(6))))))))
	score = int(int32((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(backpressureEvents) * hxrt.Int32Wrap(2))))))))
	score = int(int32((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(retryCount) * hxrt.Int32Wrap(2))))))))
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(len(responses)))))
	return score
}
