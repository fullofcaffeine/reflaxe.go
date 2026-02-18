package main

import "snapshot/hxrt"

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

type I_Child interface {
}

type Child struct {
	*Base
	__hx_this I_Child
}

func New_Child() *Child {
	self := &Child{}
	self.Base = New_Base()
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func main() {
	child := New_Child().Base
	_ = child
	base := New_Base()
	_ = base
	hxrt.Println(func(hx_value *Base) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*Child)
		return ok
	}(child))
	hxrt.Println((child != nil))
	hxrt.Println(func(hx_value *Base) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*Child)
		return ok
	}(base))
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case *Child:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(nil)))
	hxrt.Println(true)
	hxrt.Println(true)
	hxrt.Println(false)
	hxrt.Println((hxrt.StringFromLiteral("x") != nil))
	hxrt.Println(true)
	hxrt.Println(false)
}
