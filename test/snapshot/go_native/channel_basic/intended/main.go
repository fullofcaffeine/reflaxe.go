package main

import "snapshot/hxrt"

func main() {
	ch := New_go___Chan()
	_ = ch
	ch.__hx_this.send(10)
	ch.__hx_this.send(20)
	hxrt.Println(ch.__hx_this.recv())
	hxrt.Println(ch.__hx_this.recv())
	hxrt.Println(ch.__hx_this.recv())
}

type I_go___Chan interface {
	send(value any)
	recv() any
}

type go___Chan struct {
	__hx_this I_go___Chan
	queue     []any
	readIndex int
}

func New_go___Chan() *go___Chan {
	self := &go___Chan{}
	self.__hx_this = self
	self.queue = []any{}
	self.readIndex = 0
	return self
}

func (self *go___Chan) send(value any) {
	self.queue = append(self.queue, value)
}

func (self *go___Chan) recv() any {
	if self.readIndex >= len(self.queue) {
		return nil
	}
	value := self.queue[self.readIndex]
	_ = value
	self.readIndex = (self.readIndex + 1)
	return value
}
