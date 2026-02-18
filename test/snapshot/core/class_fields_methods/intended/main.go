package main

import "snapshot/hxrt"

type Counter struct {
	value int
}

func New_Counter(start int) *Counter {
	self := &Counter{}
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
	hxrt.Println(counter.inc(2))
	hxrt.Println(counter.value)
}
