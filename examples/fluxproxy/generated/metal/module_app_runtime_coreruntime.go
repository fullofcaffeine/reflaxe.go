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

func (self *app__runtime__CoreRuntime) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__runtime__CoreRuntime() *app__runtime__CoreRuntime {
	return hx_init_app__runtime__CoreRuntime(hx_alloc_app__runtime__CoreRuntime(nil))
}

func hx_alloc_app__runtime__CoreRuntime(owner any) *app__runtime__CoreRuntime {
	self := &app__runtime__CoreRuntime{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__runtime__CoreRuntime)
	}
	return self
}

func hx_init_app__runtime__CoreRuntime(self *app__runtime__CoreRuntime) *app__runtime__CoreRuntime {
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
		request := hx_recover_class_0(requests.Get(_g))
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
		response := hx_recover_class_1(responses.Get(_g))
		_g = int(int32((_g + 1)))
		if response.success {
			successCount = int(int32((successCount + 1)))
		} else {
			errorCount = int(int32((errorCount + 1)))
		}
	}
	score := 0
	score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(successCount) * hxrt.Int32Wrap(10))))))
	score = int((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int((hxrt.Int32Wrap(errorCount) * hxrt.Int32Wrap(6))))))
	score = int((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int((hxrt.Int32Wrap(backpressureEvents) * hxrt.Int32Wrap(2))))))
	score = int((hxrt.Int32Wrap(score) - hxrt.Int32Wrap(int((hxrt.Int32Wrap(retryCount) * hxrt.Int32Wrap(2))))))
	score = int((hxrt.Int32Wrap(score) + hxrt.Int32Wrap(responses.Len())))
	return score
}
