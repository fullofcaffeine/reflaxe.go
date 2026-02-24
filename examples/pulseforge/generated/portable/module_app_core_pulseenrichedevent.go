package main

type I_app__core__PulseEnrichedEvent interface {
	shouldAlert(weightedThreshold int) bool
}

type app__core__PulseEnrichedEvent struct {
	__hx_this     I_app__core__PulseEnrichedEvent
	event         *app__core__PulseEvent
	severity      int
	weightedValue int
}

func New_app__core__PulseEnrichedEvent(event *app__core__PulseEvent, severity int, weightedValue int) *app__core__PulseEnrichedEvent {
	self := &app__core__PulseEnrichedEvent{}
	self.__hx_this = self
	self.event = event
	self.severity = severity
	self.weightedValue = weightedValue
	return self
}

func (self *app__core__PulseEnrichedEvent) shouldAlert(weightedThreshold int) bool {
	return (self.weightedValue >= weightedThreshold)
}
