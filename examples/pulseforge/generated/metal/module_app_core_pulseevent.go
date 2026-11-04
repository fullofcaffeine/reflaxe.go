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

func (self *app__core__PulseEvent) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__PulseEvent(id int, source *string, region *string, value int) *app__core__PulseEvent {
	return hx_init_app__core__PulseEvent(hx_alloc_app__core__PulseEvent(nil), id, source, region, value)
}

func hx_alloc_app__core__PulseEvent(owner any) *app__core__PulseEvent {
	self := &app__core__PulseEvent{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__PulseEvent)
	}
	return self
}

func hx_init_app__core__PulseEvent(self *app__core__PulseEvent, id int, source *string, region *string, value int) *app__core__PulseEvent {
	self.id = id
	self.source = source
	self.region = region
	self.value = value
	return self
}

func (self *app__core__PulseEvent) isAlert(minValue int) bool {
	return (self.value >= minValue)
}
