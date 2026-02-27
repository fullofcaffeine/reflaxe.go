package main

type I_profile__StorySignalMetrics interface {
}

type profile__StorySignalMetrics struct {
	__hx_this     I_profile__StorySignalMetrics
	cardCount     int
	highValue     int
	openHighValue int
}

func New_profile__StorySignalMetrics(cardCount int, highValue int, openHighValue int) *profile__StorySignalMetrics {
	self := &profile__StorySignalMetrics{}
	self.__hx_this = self
	self.cardCount = cardCount
	self.highValue = highValue
	self.openHighValue = openHighValue
	return self
}
