package main

type I_UsedBox interface {
}

type UsedBox struct {
	__hx_this I_UsedBox
	value     any
}

func (self *UsedBox) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_UsedBox(value any) *UsedBox {
	return hx_init_UsedBox(hx_alloc_UsedBox(nil), value)
}

func hx_alloc_UsedBox(owner any) *UsedBox {
	self := &UsedBox{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_UsedBox)
	}
	return self
}

func hx_init_UsedBox(self *UsedBox, value any) *UsedBox {
	self.value = value
	return self
}

type I_UsedSibling interface {
}

type UsedSibling struct {
	__hx_this I_UsedSibling
	delta     int
}

func (self *UsedSibling) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_UsedSibling(delta int) *UsedSibling {
	return hx_init_UsedSibling(hx_alloc_UsedSibling(nil), delta)
}

func hx_alloc_UsedSibling(owner any) *UsedSibling {
	self := &UsedSibling{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_UsedSibling)
	}
	return self
}

func hx_init_UsedSibling(self *UsedSibling, delta int) *UsedSibling {
	self.delta = delta
	return self
}
