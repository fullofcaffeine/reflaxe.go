package main

type I_profile__TodoRuntimeMetrics interface {
}

type profile__TodoRuntimeMetrics struct {
	__hx_this I_profile__TodoRuntimeMetrics
	total     int
	open      int
	done      int
	p1        int
}

func New_profile__TodoRuntimeMetrics(total int, open int, done int, p1 int) *profile__TodoRuntimeMetrics {
	self := &profile__TodoRuntimeMetrics{}
	self.__hx_this = self
	self.total = total
	self.open = open
	self.done = done
	self.p1 = p1
	return self
}
