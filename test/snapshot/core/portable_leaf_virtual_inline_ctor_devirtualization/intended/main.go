package main

import "snapshot/hxrt"

type I_Leaf interface {
	ping() int
	pong() int
}

type Leaf struct {
	__hx_this I_Leaf
}

func (self *Leaf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Leaf() *Leaf {
	return hx_init_Leaf(hx_alloc_Leaf(nil))
}

func hx_alloc_Leaf(owner any) *Leaf {
	self := &Leaf{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Leaf)
	}
	return self
}

func hx_init_Leaf(self *Leaf) *Leaf {
	return self
}

func (self *Leaf) ping() int {
	return self.pong()
}

func (self *Leaf) pong() int {
	return 13
}

func main() {
	var v any = any(New_Leaf().ping())
	hxrt.Println(v)
}
