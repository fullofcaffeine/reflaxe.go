package main

type I_profile__StorySignalMetrics interface {
}

type profile__StorySignalMetrics struct {
	__hx_this     I_profile__StorySignalMetrics
	cardCount     int
	highValue     int
	openHighValue int
}

func (self *profile__StorySignalMetrics) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_profile__StorySignalMetrics(cardCount int, highValue int, openHighValue int) *profile__StorySignalMetrics {
	return hx_init_profile__StorySignalMetrics(hx_alloc_profile__StorySignalMetrics(nil), cardCount, highValue, openHighValue)
}

func hx_alloc_profile__StorySignalMetrics(owner any) *profile__StorySignalMetrics {
	self := &profile__StorySignalMetrics{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_profile__StorySignalMetrics)
	}
	return self
}

func hx_init_profile__StorySignalMetrics(self *profile__StorySignalMetrics, cardCount int, highValue int, openHighValue int) *profile__StorySignalMetrics {
	self.cardCount = cardCount
	self.highValue = highValue
	self.openHighValue = openHighValue
	return self
}
