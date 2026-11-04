package main

import "snapshot/hxrt"

type I_Counter interface {
	inc(step int) int
}

type Counter struct {
	__hx_this I_Counter
	value     int
}

func (self *Counter) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Counter(start int) *Counter {
	return hx_init_Counter(hx_alloc_Counter(nil), start)
}

func hx_alloc_Counter(owner any) *Counter {
	self := &Counter{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Counter)
	}
	return self
}

func hx_init_Counter(self *Counter, start int) *Counter {
	self.value = start
	return self
}

func (self *Counter) inc(step int) int {
	self.value = int((hxrt.Int32Wrap(self.value) + hxrt.Int32Wrap(step)))
	return self.value
}

func main() {
	counter := New_Counter(5)
	var v any = any(counter.value)
	hxrt.Println(v)
	var v_1 any = any(counter.inc(2))
	hxrt.Println(v_1)
	var v_2 any = any(counter.value)
	hxrt.Println(v_2)
}
