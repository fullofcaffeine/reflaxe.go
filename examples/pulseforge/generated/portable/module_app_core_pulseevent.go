package main

type I_app__core__PulseEvent interface {
	isAlert(minValue int) bool
}

type app__core__PulseEvent struct {
	__hx_this I_app__core__PulseEvent
	id        int
	source    *string
	region    *string
	value     int
}

func New_app__core__PulseEvent(id int, source *string, region *string, value int) *app__core__PulseEvent {
	self := &app__core__PulseEvent{}
	self.__hx_this = self
	self.id = id
	self.source = source
	self.region = region
	self.value = value
	return self
}

func (self *app__core__PulseEvent) isAlert(minValue int) bool {
	return (self.value >= minValue)
}
