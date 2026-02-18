package main

import "snapshot/hxrt"

type I_Counter interface {
	inc(step int) int
}

type Counter struct {
	__hx_this I_Counter
	value     int
}

func New_Counter(start int) *Counter {
	self := &Counter{}
	self.__hx_this = self
	self.value = start
	return self
}

func (self *Counter) inc(step int) int {
	self.value = (self.value + step)
	return self.value
}

func main() {
	counter := New_Counter(5)
	hxrt.Println(counter.value)
	hxrt.Println(counter.__hx_this.inc(2))
	hxrt.Println(counter.value)
}
