package main

type I_app__core__PulseIngressFrame interface {
}

type app__core__PulseIngressFrame struct {
	__hx_this I_app__core__PulseIngressFrame
	sequence  int
	source    *string
	value     int
	region    *string
}

func (self *app__core__PulseIngressFrame) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__PulseIngressFrame(sequence int, source *string, value int, region *string) *app__core__PulseIngressFrame {
	return hx_init_app__core__PulseIngressFrame(hx_alloc_app__core__PulseIngressFrame(nil), sequence, source, value, region)
}

func hx_alloc_app__core__PulseIngressFrame(owner any) *app__core__PulseIngressFrame {
	self := &app__core__PulseIngressFrame{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__PulseIngressFrame)
	}
	return self
}

func hx_init_app__core__PulseIngressFrame(self *app__core__PulseIngressFrame, sequence int, source *string, value int, region *string) *app__core__PulseIngressFrame {
	self.sequence = sequence
	self.source = source
	self.value = value
	self.region = region
	return self
}
