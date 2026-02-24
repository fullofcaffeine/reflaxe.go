package main

type I_app__core__PulseEvent interface {
	isAlert() bool
}

type app__core__PulseEvent struct {
	__hx_this I_app__core__PulseEvent
	id        int
	source    *string
	value     int
}

func New_app__core__PulseEvent(id int, source *string, value int) *app__core__PulseEvent {
	self := &app__core__PulseEvent{}
	self.__hx_this = self
	self.id = id
	self.source = source
	self.value = value
	return self
}

func (self *app__core__PulseEvent) isAlert() bool {
	return (self.value >= 8)
}
