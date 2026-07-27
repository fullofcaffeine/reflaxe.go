package main

type I_UsedBox interface {
}

type UsedBox struct {
	__hx_this I_UsedBox
	value     any
}

func New_UsedBox(value any) *UsedBox {
	self := &UsedBox{}
	self.__hx_this = self
	self.value = value
	return self
}

type I_UsedSibling interface {
}

type UsedSibling struct {
	__hx_this I_UsedSibling
	delta     int
}

func New_UsedSibling(delta int) *UsedSibling {
	self := &UsedSibling{}
	self.__hx_this = self
	self.delta = delta
	return self
}
