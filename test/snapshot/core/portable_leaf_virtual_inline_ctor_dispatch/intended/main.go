package main

import "snapshot/hxrt"

type I_Leaf interface {
	ping() int
	pong() int
}

type Leaf struct {
	__hx_this I_Leaf
}

func New_Leaf() *Leaf {
	self := &Leaf{}
	self.__hx_this = self
	return self
}

func (self *Leaf) ping() int {
	return self.__hx_this.pong()
}

func (self *Leaf) pong() int {
	return 13
}

func main() {
	hxrt.Println(New_Leaf().__hx_this.ping())
}
