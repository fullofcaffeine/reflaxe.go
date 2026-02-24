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

func New_app__core__PulseIngressFrame(sequence int, source *string, value int, region *string) *app__core__PulseIngressFrame {
	self := &app__core__PulseIngressFrame{}
	self.__hx_this = self
	self.sequence = sequence
	self.source = source
	self.value = value
	self.region = region
	return self
}
