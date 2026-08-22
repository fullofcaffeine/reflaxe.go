package main

import "examples_fluxproxy_metal/hxrt"

type I_app__runtime__CoreRuntime interface {
	profileId() *string
	variantId() *string
	capabilityId() *string
	dispatch(requests *hxrt.Array, workerCount int) *hxrt.Array
	stageScore(responses *hxrt.Array, retryCount int, backpressureEvents int) int
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
	return hxrt.StringFromLiteral("metal")
}

func (self *app__runtime__CoreRuntime) variantId() *string {
	return hxrt.StringFromLiteral("core")
}

func (self *app__runtime__CoreRuntime) capabilityId() *string {
	return hxrt.StringFromLiteral("loop_dispatch")
}

func (self *app__runtime__CoreRuntime) dispatch(requests *hxrt.Array, workerCount int) *hxrt.Array {
	responses := hxrt.NewArray()
	_g := 0
	for _g < requests.Len() {
		request := func(hx_value_1 any) *app__core__FluxRequest {
			if hx_value_1 == nil {
				var hx_zero_2 *app__core__FluxRequest
				return hx_zero_2
			}
			return hx_value_1.(*app__core__FluxRequest)
		}(requests.Get(_g))
		_g = int(int32((_g + 1)))
		responses.Push(app__core__FluxCodec_proxy(request, 50))
	}
	return responses
}

func (self *app__runtime__CoreRuntime) stageScore(responses *hxrt.Array, retryCount int, backpressureEvents int) int {
	successCount := 0
	errorCount := 0
	_g := 0
	for _g < responses.Len() {
		response := func(hx_value_4 any) *app__core__FluxProxyResponse {
			if hx_value_4 == nil {
				var hx_zero_5 *app__core__FluxProxyResponse
				return hx_zero_5
			}
			return hx_value_4.(*app__core__FluxProxyResponse)
		}(responses.Get(_g))
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
	score = int(int32((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(responses.Len()))))
	return score
}
